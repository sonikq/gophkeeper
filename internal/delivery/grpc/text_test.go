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

type TextTestSuite struct {
	suite.Suite

	Server GophKeeperServer
	ctx    context.Context
}

func (suite *TextTestSuite) SetupTest() {
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
	err = suite.Server.Usecase.SaveText(suite.ctx, models.TextData{
		UUID: "initial UUID",
		Data: "initial text",
		Meta: "initial Meta",
	})
	if err != nil {
		suite.T().Errorf("Error setup - saving text: %e", err)
	}
}

func (suite *TextTestSuite) TestSaveTextInvalidToken() {
	req := &pb.SaveTextDataRequest{
		Token: "initial login/initial password/invalid salt",
		Data: &pb.TextData{
			Uuid: "test uuid",
			Data: "new text",
			Meta: &pb.Meta{},
		},
	}
	resp, err := suite.Server.SaveText(suite.ctx, req)
	require.Error(suite.T(), models.ErrInvalidToken, err)
	require.Equal(suite.T(), "authorization error: invalid token", resp.Error)
}

func (suite *TextTestSuite) TestSaveTextSuccess() {
	req := &pb.SaveTextDataRequest{
		Token: "initial login/initial password/salt",
		Data: &pb.TextData{
			Uuid: "test uuid",
			Data: "new text",
			Meta: &pb.Meta{},
		},
	}
	resp, err := suite.Server.SaveText(suite.ctx, req)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), "", resp.Error)
}

func (suite *TextTestSuite) TestLoadTextSuccess() {
	req := &pb.LoadTextDataRequest{
		Token: "initial login/initial password/salt",
		Uuid:  "initial UUID",
	}
	resp, err := suite.Server.LoadText(suite.ctx, req)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), "", resp.Error)
	require.Equal(suite.T(), "initial UUID", resp.Data.Uuid)
	require.Equal(suite.T(), "initial text", resp.Data.Data)
	require.Equal(suite.T(), "initial Meta", resp.Data.Meta.Content)
}

func (suite *TextTestSuite) TestLoadTextAuthError() {
	req := &pb.LoadTextDataRequest{
		Token: "initial login/initial password/wrong salt",
		Uuid:  "initial UUID",
	}
	resp, err := suite.Server.LoadText(suite.ctx, req)
	require.Error(suite.T(), models.ErrInvalidToken, err)
	require.Equal(suite.T(), "authorization error: invalid token", resp.Error)
}

func (suite *TextTestSuite) TestLoadTextNoSuchID() {
	req := &pb.LoadTextDataRequest{
		Token: "initial login/initial password/salt",
		Uuid:  "wrong UUID",
	}
	resp, err := suite.Server.LoadText(suite.ctx, req)
	require.Error(suite.T(), models.ErrDatabaseError, err)
	require.Equal(suite.T(), "internal server error for data wrong UUID", resp.Error)
}

func TestTextTestSuite(t *testing.T) {
	suite.Run(t, new(TextTestSuite))
}
