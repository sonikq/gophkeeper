package server

import (
	"fmt"
	"github.com/sonikq/gophkeeper/internal/app/server/config"
	gophkeeper "github.com/sonikq/gophkeeper/internal/delivery/grpc/v1"
	pb "github.com/sonikq/gophkeeper/internal/delivery/grpc/v1"
	"github.com/sonikq/gophkeeper/internal/repository"
	"github.com/sonikq/gophkeeper/internal/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
)

type GophKeeperServer struct {
	pb.UnimplementedGophKeeperHandlerServer
	Usecase usecase.GophKeeperUseCase
}

func newGophKeeperServerGrpc(gserver *grpc.Server, usecase usecase.GophKeeperUseCase) {

	gophKeeperServer := &GophKeeperServer{
		Usecase: usecase,
	}

	gophkeeper.RegisterGophKeeperHandlerServer(gserver, gophKeeperServer)
	reflection.Register(gserver)
}

func Run() {
	idleConnsClosed := make(chan struct{})
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to initialize config: %s", err.Error())
	}

	//ctx, cancel := context.WithTimeout(context.Background(), cfg.CtxTimeout)
	//defer cancel()

	listen, err := net.Listen("tcp", cfg.RunAddress)
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	go func() {
		<-sigint
		fmt.Println("received shutdown signal")
		fmt.Println("running gracefull shutdown")
		server.GracefulStop()
		fmt.Println("server shutted down")
		close(idleConnsClosed)
	}()

	repo := repository.NewGophKeeperRepository()

	usecaseManager := usecase.NewGophKeeperUseCase(repo)

	newGophKeeperServerGrpc(server, usecaseManager)

	fmt.Println("gRPC server starts working")
	if err = server.Serve(listen); err != nil {
		log.Fatal(err)
	}
	<-idleConnsClosed
}
