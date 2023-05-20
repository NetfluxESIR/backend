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

func (p *Engine) CreateToken(ctx context.Context, token models.Token) error {
	tx := p.db.WithContext(ctx).Create(&token)
	return tx.Error
}

func (p *Engine) UpdateToken(ctx context.Context, token models.Token) error {
	tx := p.db.WithContext(ctx).Save(&token)
	return tx.Error
}

func (p *Engine) DeleteToken(ctx context.Context, token string) error {
	tx := p.db.WithContext(ctx).Where("token = ?", token).Delete(&models.Token{})
	return tx.Error
}

func (p *Engine) GetVideo(ctx context.Context, userId, videoId string) (models.Video, error) {
	var v models.Video
	tx := p.db.WithContext(ctx).
		Preload("Processing").
		Where("video_id = ?", videoId).
		Where("user_id = ?", userId).
		First(&v)
	return v, tx.Error
}

func (p *Engine) CreateVideo(ctx context.Context, video *models.Video) error {
	tx := p.db.WithContext(ctx).Create(video)
	return tx.Error
}

func (p *Engine) UpdateVideo(ctx context.Context, video *models.Video) error {
	tx := p.db.WithContext(ctx).Save(video)
	return tx.Error
}

func (p *Engine) DeleteVideo(ctx context.Context, userId, videoId string) error {
	tx := p.db.WithContext(ctx).Where("video_id = ?", videoId).Where("user_id = ?", userId).Delete(&models.Video{})
	return tx.Error
}

func (p *Engine) GetVideoList(ctx context.Context, userId string, limit int, offset int) ([]models.Video, error) {
	var v []models.Video
	tx := p.db.WithContext(ctx).Preload("Processing").Where("user_id = ?", userId).Find(&v).Limit(limit).Offset(offset)
	return v, tx.Error
}

func (p *Engine) RegisterUser(ctx context.Context, user *models.Account) error {
	tx := p.db.WithContext(ctx).Create(user)
	return tx.Error
}

func (p *Engine) GetUser(ctx context.Context, email string) (models.Account, error) {
	var u models.Account
	tx := p.db.WithContext(ctx).Where("email = ?", email).First(&u)
	return u, tx.Error
}

func (p *Engine) GetProcessing(ctx context.Context, videoId string) (models.Processing, error) {
	var processing models.Processing
	tx := p.db.WithContext(ctx).
		Preload("Steps").
		Where("video_id = ?", videoId).
		First(&processing)
	return processing, tx.Error
}

func (p *Engine) UpdateProcessing(ctx context.Context, processing models.Processing) error {
	tx := p.db.WithContext(ctx).Save(&processing)
	return tx.Error
}
