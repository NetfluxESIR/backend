package server

import (
	"context"
	"time"
)

func (s *Server) ValidateToken(ctx context.Context, token string) (bool, error) {
	getToken, err := s.db.GetToken(ctx, token)
	if err != nil {
		return false, err
	}
	if getToken.Token == "" {
		return false, nil
	}
	if getToken.Expire && getToken.ExpirationDate.Before(time.Now()) {
		return false, nil
	}
	return true, nil
}
