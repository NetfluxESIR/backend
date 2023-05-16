package server

import (
	"context"
	"github.com/NetfluxESIR/backend/internal/persistence"
	"github.com/NetfluxESIR/backend/pkg/api"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	cfg    *Config
	logger *log.Entry
	db     *persistence.Engine
	api    *api.API
}

func New(ctx context.Context, cfg *Config) (*Server, error) {
	db, err := persistence.New(ctx, cfg.DSN, cfg.Logger.WithField("component", "persistence"))
	if err != nil {
		return nil, err
	}
	return &Server{
		cfg:    cfg,
		logger: cfg.Logger,
		db:     db,
		api:    api.New(ctx, cfg.Host, cfg.Port, cfg.Logger.WithField("component", "api")),
	}, nil
}

func (s *Server) Run(ctx context.Context) error {
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	return nil
}
