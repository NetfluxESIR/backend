package server

import (
	"context"
	"github.com/NetfluxESIR/backend/internal/models"
	"github.com/NetfluxESIR/backend/internal/persistence"
	"github.com/NetfluxESIR/backend/pkg/api"
	"github.com/NetfluxESIR/backend/pkg/api/gen"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	cfg    *Config
	logger *log.Entry
	db     *persistence.Engine
	api    *api.API
	s3     *session.Session
}

func New(ctx context.Context, cfg *Config) (*Server, error) {
	db, err := persistence.New(ctx, cfg.DSN, cfg.Logger.WithField("component", "persistence"))
	if err != nil {
		return nil, err
	}
	s := &Server{
		cfg:    cfg,
		logger: cfg.Logger,
		db:     db,
		api:    nil,
	}
	password, err := hashAndSalt([]byte(cfg.AdminPassword))
	if err != nil {
		return nil, err
	}
	account := &models.Account{
		Email:          cfg.AdminAccount,
		HashedPassword: password,
		Role:           string(gen.ADMIN),
	}
	s.db.RegisterUser(ctx, account)
	s.s3, err = session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(cfg.S3AccessKey, cfg.S3SecretKey, ""),
		Region:      aws.String(cfg.S3Region),
	})
	if err != nil {
		return nil, err
	}
	s.api = api.New(ctx, cfg.Host, cfg.Port, s, s.logger.WithField("component", "api"))
	return s, nil
}

func (s *Server) Run(ctx context.Context) error {
	return s.api.Run(ctx)
}

func (s *Server) Stop(ctx context.Context) error {
	err := s.db.Close(ctx)
	if err != nil {
		return err
	}
	return s.api.Stop(ctx)
}
