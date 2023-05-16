package persistence

import (
	"context"
	"github.com/NetfluxESIR/backend/internal/models"
)

func (p *Engine) GetToken(ctx context.Context, token string) (models.Token, error) {
	var t models.Token
	tx := p.db.WithContext(ctx).Where("token = ?", token).First(&t)
	return t, tx.Error
}
