package server

import (
	"context"
	"fmt"
	"time"
)

func (s *Server) ValidateToken(ctx context.Context, token string) (string, error) {
	// remove Bearer
	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}
	getToken, err := s.db.GetToken(ctx, token)
	if err != nil {
		return "", err
	}
	if getToken.Token == "" {
		return "", fmt.Errorf("token not found")
	}
	if getToken.Expire && getToken.ExpirationDate.Before(time.Now()) {
		return "", fmt.Errorf("token expired")
	}
	return getToken.AccountId, nil
}
