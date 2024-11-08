package grpc

import (
	"context"
	"fmt"
	"github.com/sonikq/gophkeeper/internal/app/models"
	pb "github.com/sonikq/gophkeeper/internal/delivery/grpc/v1"
)

func (s *GophKeeperServer) SaveBinary(ctx context.Context, in *pb.SaveBinaryDataRequest) (*pb.SaveBinaryDataResponse, error) {
	var response pb.SaveBinaryDataResponse

	validationError := s.ValidateRequest(ctx, in.Token)
	response.Error = validationError
	if response.Error != "" {
		return &response, models.ErrInvalidToken
	}

	newBytes := models.BinaryData{
		UUID: in.Data.Uuid,
		Data: in.Data.Data,
		Meta: in.Data.Meta.Content,
	}
	err := s.Usecase.SaveBinary(ctx, newBytes)
	if err != nil {
		response.Error = fmt.Sprintf("internal server error for data %s", in.Data.Uuid)
		return &response, models.ErrDatabaseError
	}
	return &response, nil
}

func (s *GophKeeperServer) LoadBinary(ctx context.Context, in *pb.LoadBinaryDataRequest) (*pb.LoadBinaryDataResponse, error) {
	var response pb.LoadBinaryDataResponse

	validationError := s.ValidateRequest(ctx, in.Token)
	response.Error = validationError
	if response.Error != "" {
		return &response, models.ErrInvalidToken
	}

	bin, err := s.Usecase.LoadBinary(ctx, in.Uuid)
	if err != nil {
		response.Error = fmt.Sprintf("internal server error for data %s", in.Uuid)
		return &response, models.ErrDatabaseError
	}
	response.Data = &pb.BinaryData{
		Uuid: bin.UUID,
		Data: bin.Data,
		Meta: &pb.Meta{
			Content: bin.Meta,
		},
	}
	return &response, nil
}
