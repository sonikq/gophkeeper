package grpc

import (
	"context"
	"fmt"
	"github.com/sonikq/gophkeeper/internal/app/models"
	pb "github.com/sonikq/gophkeeper/internal/delivery/grpc/v1"
)

func (s *GophKeeperServer) SaveText(ctx context.Context, in *pb.SaveTextDataRequest) (*pb.SaveTextDataResponse, error) {
	var response pb.SaveTextDataResponse

	validationError := s.ValidateRequest(ctx, in.Token)
	response.Error = validationError
	if response.Error != "" {
		return &response, models.ErrInvalidToken
	}

	newText := models.TextData{
		UUID: in.Data.Uuid,
		Data: in.Data.Data,
		Meta: in.Data.Meta.Content,
	}
	err := s.Usecase.SaveText(ctx, newText)
	if err != nil {
		response.Error = fmt.Sprintf("internal server error for data %s", in.Data.Uuid)
		return &response, models.ErrDatabaseError
	}
	return &response, nil
}

func (s *GophKeeperServer) LoadText(ctx context.Context, in *pb.LoadTextDataRequest) (*pb.LoadTextDataResponse, error) {
	var response pb.LoadTextDataResponse

	validationError := s.ValidateRequest(ctx, in.Token)
	response.Error = validationError
	if response.Error != "" {
		return &response, models.ErrInvalidToken
	}

	text, err := s.Usecase.LoadText(ctx, in.Uuid)
	if err != nil {
		response.Error = fmt.Sprintf("internal server error for data %s", in.Uuid)
		return &response, models.ErrDatabaseError
	}
	response.Data = &pb.TextData{
		Uuid: text.UUID,
		Data: text.Data,
		Meta: &pb.Meta{
			Content: text.Meta,
		},
	}
	return &response, nil
}
