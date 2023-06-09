package api

import (
	"context"
	"fmt"
	"github.com/NetfluxESIR/backend/pkg/api/gen"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

//go:generate oapi-codegen -generate types -o gen/types.gen.go -package gen specs/pkg/video-api/openapi.yaml
//go:generate oapi-codegen -generate gin -o gen/server.gen.go -package gen specs/pkg/video-api/openapi.yaml
//go:generate oapi-codegen -generate client -o gen/client.gen.go -package gen specs/pkg/video-api/openapi.yaml

type API struct {
	host   string
	port   int
	logger *log.Entry
	server *gen.ServerInterfaceWrapper
}

type Config struct {
	Host    string
	Port    int
	Handler ServerInterface
}

type ServerInterface interface {
	Auth
	gen.ServerInterface
}
type Auth interface {
	ValidateToken(ctx context.Context, token string) (string, error)
}

func New(ctx context.Context, host string, port int, handler ServerInterface, logger *log.Entry) *API {
	return &API{
		host:   host,
		port:   port,
		logger: logger,
		server: &gen.ServerInterfaceWrapper{
			Handler: handler,
			HandlerMiddlewares: []gen.MiddlewareFunc{
				func(c *gin.Context) {
					if c.Request.URL.Path == "/api/v1/users/login" {
						c.Next()
						return
					}
					token := c.Request.Header.Get("Authorization")
					if token == "" {
						c.AbortWithStatusJSON(401, gin.H{"error": "missing token"})
						return
					}
					uuid, err := handler.ValidateToken(c, token)
					if err != nil {
						c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
						return
					}
					c.Set("userId", uuid)
					c.Next()
				},
			},
			ErrorHandler: func(c *gin.Context, err error, code int) {
				c.AbortWithStatusJSON(code, gin.H{"error": err.Error()})
			},
		},
	}
}

func (a *API) Stop(ctx context.Context) error {
	return nil
}

func (a *API) Run(ctx context.Context) error {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, "Authorization", "authorization", "Content-Type", "content-type", "*")
	router.Use(cors.New(config))
	gen.RegisterHandlersWithOptions(router, a.server.Handler, gen.GinServerOptions{
		Middlewares: a.server.HandlerMiddlewares,
		ErrorHandler: func(c *gin.Context, err error, code int) {
			a.server.ErrorHandler(c, err, code)
		},
	})
	a.logger.Infof("API listening on %s:%d", a.host, a.port)
	if err := router.Run(fmt.Sprintf("%s:%d", a.host, a.port)); err != nil {
		return err
	}
	return nil
}
