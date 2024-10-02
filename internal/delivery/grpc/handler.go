package grpc

import (
	gophkeeper "github.com/sonikq/gophkeeper/internal/delivery/grpc/v1"
	pb "github.com/sonikq/gophkeeper/internal/delivery/grpc/v1"
	"github.com/sonikq/gophkeeper/internal/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GophKeeperServer struct {
	pb.UnimplementedGophKeeperHandlerServer
	Usecase usecase.GophKeeperUseCase
}

func NewGophKeeperServer(gserver *grpc.Server, usecase usecase.GophKeeperUseCase) {

	gophKeeperServer := &GophKeeperServer{
		Usecase: usecase,
	}

	gophkeeper.RegisterGophKeeperHandlerServer(gserver, gophKeeperServer)
	reflection.Register(gserver)
}
