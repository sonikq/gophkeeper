package grpc

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/sonikq/gophkeeper/internal/app/models"
	pb "github.com/sonikq/gophkeeper/internal/delivery/grpc/v1"
)

func (s *GophKeeperServer) LoginUser(ctx context.Context, in *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	var response pb.LoginUserResponse

	token := Token{
		Login:    in.User.Login,
		Password: in.User.Password,
		salt:     "salt",
	}
	encodedToken := EncodeToken(token)
	_, err := s.Usecase.FindUser(ctx, in.User.Login, in.User.Password)
	switch err {
	case nil:
		response.Token = encodedToken
		return &response, nil
	case models.ErrUserNotFound:
		response.Error = fmt.Sprintf("user %s not found", in.User.Login)
		return &response, models.ErrUserNotFound
	default:
		response.Error = fmt.Sprintf("error loading user %s: %e", in.User.Login, err)
		return &response, models.ErrDatabaseError
	}
}

func (s *GophKeeperServer) RegisterUser(ctx context.Context, in *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	var response pb.RegisterUserResponse

	token := Token{
		Login:    in.User.Login,
		Password: in.User.Password,
		salt:     "salt",
	}
	encodedToken := EncodeToken(token)
	user, err := s.Usecase.FindUser(ctx, in.User.Login, in.User.Password)
	switch err {
	case nil:
		response.Error = fmt.Sprintf("user already exists %s", in.User.Login)
		return &response, models.ErrUserAlreadyExists
	case models.ErrContextTimeout:
		response.Error = fmt.Sprintf("registration. find user timeout %s: %e", in.User.Login, err)
		return &response, models.ErrContextTimeout
	case models.ErrInMemoryDB:
		response.Error = fmt.Sprintf("internal server error %s: %e", in.User.Login, err)
		return &response, models.ErrDatabaseError
	}

	sessions := user.Sessions
	sessions = append(sessions, models.Session{
		UUID:  uuid.NewString(),
		Token: encodedToken,
	})

	newUser := models.User{
		UUID:     uuid.NewString(),
		Login:    in.User.Login,
		Password: in.User.Password,
		Sessions: sessions,
	}
	err = s.Usecase.SaveUser(ctx, newUser)
	if err != nil {
		response.Error = fmt.Sprintf("registration error with %s: %e", in.User.Login, err)
		return &response, models.ErrDatabaseError
	}
	return &response, nil
}
