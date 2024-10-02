package grpc

import (
	"context"
	"github.com/sonikq/gophkeeper/internal/app/models"
)

func ValidateToken(incoming, stored string) bool {
	return incoming == stored
}

func (s *GophKeeperServer) ValidateRequest(ctx context.Context, token string) string {
	decodedToken := DecodeToken(token)
	user, err := s.Usecase.FindUser(ctx, decodedToken.Login, decodedToken.Password)
	switch err {
	case models.ErrUserNotFound:
		return "authorization error: wrong username"
	case models.ErrContextTimeout:
		return "context timeout. internal server error"
	case models.ErrInMemoryDB:
		return "internal server error"
	}
	for _, session := range user.Sessions {
		if ValidateToken(token, session.Token) {
			// In case of successfull validation value for response.Error is empty
			return ""
		}
	}

	// No tokens registered for the user, therefore token is invalid
	return "authorization error: invalid token"
}
