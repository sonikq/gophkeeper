package grpc

import (
	"context"
	"fmt"
	"github.com/sonikq/gophkeeper/internal/app/models"
	pb "github.com/sonikq/gophkeeper/internal/delivery/grpc/v1"
)

func (s *GophKeeperServer) SaveCredentials(ctx context.Context, in *pb.SaveCredentialsDataRequest) (*pb.SaveCredentialsDataResponse, error) {
	var response pb.SaveCredentialsDataResponse

	validationError := s.ValidateRequest(ctx, in.Token)
	response.Error = validationError
	if response.Error != "" {
		return &response, models.ErrInvalidToken
	}

	newCredentials := models.CredentialsData{
		UUID:     in.Data.Uuid,
		Login:    in.Data.Login,
		Password: in.Data.Password,
	}

	if in.Data.Meta != nil {
		newCredentials.Meta = in.Data.Meta.Content
	}

	err := s.Usecase.SaveCredentials(ctx, newCredentials)
	if err != nil {
		response.Error = fmt.Sprintf("internal server error for data %s", in.Data.Uuid)
		return &response, models.ErrDatabaseError
	}
	return &response, nil
}

func (s *GophKeeperServer) LoadCredentials(ctx context.Context, in *pb.LoadCredentialsDataRequest) (*pb.LoadCredentialsDataResponse, error) {
	var response pb.LoadCredentialsDataResponse

	validationError := s.ValidateRequest(ctx, in.Token)
	response.Error = validationError
	if response.Error != "" {
		return &response, models.ErrInvalidToken
	}

	credentials, err := s.Usecase.LoadCredentials(ctx, in.Uuid)
	if err != nil {
		response.Error = fmt.Sprintf("internal server error for data %s", in.Uuid)
		return &response, models.ErrDatabaseError
	}
	response.Data = &pb.CredentialsData{
		Uuid:     credentials.UUID,
		Login:    credentials.Login,
		Password: credentials.Password,
		Meta: &pb.Meta{
			Content: credentials.Meta,
		},
	}
	return &response, nil
}
