package persistence

import (
	"context"
	"github.com/NetfluxESIR/backend/internal/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Engine struct {
	dsn    string
	logger *log.Entry
	db     *gorm.DB
}

func New(ctx context.Context, dsn string, logger *log.Entry) (*Engine, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	engine := &Engine{
		dsn:    dsn,
		logger: logger,
		db:     db,
	}
	if err := engine.Migrate(ctx); err != nil {
		return nil, err
	}
	return engine, nil
}

func (p *Engine) Close(ctx context.Context) error {
	db, err := p.db.DB()
	if err != nil {
		return err
	}
	return db.Close()
}

func (p *Engine) Migrate(ctx context.Context) error {
	return p.db.AutoMigrate(models.Models())
}
