package server

import (
	"context"
	"github.com/sonikq/gophkeeper/internal/app/models"
	pb "github.com/sonikq/gophkeeper/internal/delivery/grpc/v1"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"testing"
)

func TestStorageServer(t *testing.T) {
	go Run()
	conn, err := grpc.Dial(":3200", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewGophKeeperHandlerClient(conn)

	ctx := context.Background()
	// register new user
	response, errResp := client.RegisterUser(ctx, &pb.RegisterUserRequest{
		User: &pb.User{
			Login:    "user",
			Password: "pwd",
		},
	})
	require.NoError(t, errResp)
	require.Equal(t, "", response.Error)

	// now try login with user which not exists
	resp, err := client.LoginUser(ctx, &pb.LoginUserRequest{
		User: &pb.User{
			Login:    "unknown_user",
			Password: "pwd",
		},
	})
	require.Error(t, models.ErrUserNotFound, err)

	// then we login with proper user
	resp, err = client.LoginUser(ctx, &pb.LoginUserRequest{
		User: &pb.User{
			Login:    "user",
			Password: "pwd",
		},
	})
	require.NoError(t, err)
	require.Equal(t, "", resp.Error)
	require.Equal(t, "user/pwd/salt", resp.Token)
	token := resp.Token

	// let's save some data
	r, errResp := client.SaveCredentials(ctx, &pb.SaveCredentialsDataRequest{
		Token: token,
		Data: &pb.CredentialsData{
			Uuid:     "new uuid",
			Login:    "login",
			Password: "password",
			Meta: &pb.Meta{
				Content: "metadata",
			},
		},
	})
	require.NoError(t, errResp)
	require.Equal(t, "", r.Error)

	// then let's load this data
	loadResp, err := client.LoadCredentials(ctx, &pb.LoadCredentialsDataRequest{
		Token: token,
		Uuid:  "new uuid",
	})
	require.NoError(t, err)
	require.EqualValues(t, pb.CredentialsData{
		Uuid:     "new uuid",
		Login:    "login",
		Password: "password",
		Meta: &pb.Meta{
			Content: "metadata",
		},
	}, loadResp.Data)
}
