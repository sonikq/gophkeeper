package grpc

import (
	"context"
	"github.com/sonikq/gophkeeper/internal/app/models"
	pb "github.com/sonikq/gophkeeper/internal/delivery/grpc/v1"
	"github.com/sonikq/gophkeeper/internal/repository"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"sync"
	"testing"
)

type BinaryTestSuite struct {
	suite.Suite

	Server GophKeeperServer
	ctx    context.Context
}

func (suite *BinaryTestSuite) SetupTest() {
	suite.Server = GophKeeperServer{
		Usecase: &repository.InMemoryRepo{
			Users:       sync.Map{},
			Credentials: sync.Map{},
		},
	}
	suite.ctx = context.Background()
	sessions := []models.Session{
		{
			UUID:  "test session 1",
			Token: "initial login/initial password/salt",
		},
		{
			UUID:  "test session 2",
			Token: "initial login/initial password/test salt 2",
		},
	}
	err := suite.Server.Usecase.SaveUser(suite.ctx, models.User{
		Login:    "initial login",
		Password: "initial password",
		Sessions: sessions,
	})
	if err != nil {
		suite.T().Errorf("Error setup - saving user: %e", err)
	}
	err = suite.Server.Usecase.SaveBinary(suite.ctx, models.BinaryData{
		UUID: "initial UUID",
		Data: []byte{255, 255, 1},
		Meta: "initial Meta",
	})
	if err != nil {
		suite.T().Errorf("Error setup - saving text: %e", err)
	}
}

func (suite *BinaryTestSuite) TestSaveBinaryInvalidToken() {
	req := &pb.SaveBinaryDataRequest{
		Token: "initial login/initial password/invalid salt",
		Data: &pb.BinaryData{
			Uuid: "test uuid",
			Data: []byte{255, 255, 1},
			Meta: &pb.Meta{},
		},
	}
	resp, err := suite.Server.SaveBinary(suite.ctx, req)
	require.Error(suite.T(), models.ErrInvalidToken, err)
	require.Equal(suite.T(), "authorization error: invalid token", resp.Error)
}

func (suite *BinaryTestSuite) TestSaveBinarySuccess() {
	req := &pb.SaveBinaryDataRequest{
		Token: "initial login/initial password/salt",
		Data: &pb.BinaryData{
			Uuid: "test uuid",
			Data: []byte{255, 255, 2},
			Meta: &pb.Meta{},
		},
	}
	resp, err := suite.Server.SaveBinary(suite.ctx, req)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), "", resp.Error)
}

func (suite *BinaryTestSuite) TestLoadBinarySuccess() {
	req := &pb.LoadBinaryDataRequest{
		Token: "initial login/initial password/salt",
		Uuid:  "initial UUID",
	}
	resp, err := suite.Server.LoadBinary(suite.ctx, req)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), "", resp.Error)
	require.Equal(suite.T(), "initial UUID", resp.Data.Uuid)
	require.Equal(suite.T(), []byte{255, 255, 1}, resp.Data.Data)
	require.Equal(suite.T(), "initial Meta", resp.Data.Meta.Content)
}

func (suite *BinaryTestSuite) TestLoadBinaryAuthError() {
	req := &pb.LoadBinaryDataRequest{
		Token: "initial login/initial password/wrong salt",
		Uuid:  "initial UUID",
	}
	resp, err := suite.Server.LoadBinary(suite.ctx, req)
	require.Error(suite.T(), models.ErrInvalidToken, err)
	require.Equal(suite.T(), "authorization error: invalid token", resp.Error)
}

func (suite *BinaryTestSuite) TestLoadBinaryNoSuchID() {
	req := &pb.LoadBinaryDataRequest{
		Token: "initial login/initial password/salt",
		Uuid:  "wrong UUID",
	}
	resp, err := suite.Server.LoadBinary(suite.ctx, req)
	require.Error(suite.T(), models.ErrDatabaseError, err)
	require.Equal(suite.T(), "internal server error for data wrong UUID", resp.Error)
}

func TestBinaryTestSuite(t *testing.T) {
	suite.Run(t, new(BinaryTestSuite))
}
