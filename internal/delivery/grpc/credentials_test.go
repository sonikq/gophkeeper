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

type CredentialsTestSuite struct {
	suite.Suite

	Server GophKeeperServer
	ctx    context.Context
}

func (suite *CredentialsTestSuite) SetupTest() {
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
	err = suite.Server.Usecase.SaveCredentials(suite.ctx, models.CredentialsData{
		UUID:     "initial UUID",
		Login:    "initial login",
		Password: "initial password",
		Meta:     "initial Meta",
	})
	if err != nil {
		suite.T().Errorf("Error setup - saving credentials: %e", err)
	}
}

func (suite *CredentialsTestSuite) TestSaveCredentialsInvalidToken() {
	req := &pb.SaveCredentialsDataRequest{
		Token: "invalid user/invalid password/invalid salt",
		Data: &pb.CredentialsData{
			Uuid:     "test uuid",
			Login:    "invalid login",
			Password: "new password",
			Meta:     nil,
		},
	}
	resp, err := suite.Server.SaveCredentials(suite.ctx, req)
	require.Error(suite.T(), models.ErrInvalidToken, err)
	require.Equal(suite.T(), "authorization error: wrong username", resp.Error)
}

func (suite *CredentialsTestSuite) TestSaveCredentialsSuccess() {
	req := &pb.SaveCredentialsDataRequest{
		Token: "initial login/initial password/salt",
		Data: &pb.CredentialsData{
			Uuid:     "test uuid",
			Login:    "new login",
			Password: "new password",
			Meta:     nil,
		},
	}
	resp, err := suite.Server.SaveCredentials(suite.ctx, req)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), "", resp.Error)
}

func (suite *CredentialsTestSuite) TestLoadCredentialsSuccess() {
	req := &pb.LoadCredentialsDataRequest{
		Token: "initial login/initial password/salt",
		Uuid:  "initial UUID",
	}
	resp, err := suite.Server.LoadCredentials(suite.ctx, req)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), "", resp.Error)
	require.Equal(suite.T(), "initial UUID", resp.Data.Uuid)
	require.Equal(suite.T(), "initial login", resp.Data.Login)
	require.Equal(suite.T(), "initial password", resp.Data.Password)
	require.Equal(suite.T(), "initial Meta", resp.Data.Meta.Content)
}

func (suite *CredentialsTestSuite) TestLoadCredentialsAuthError() {
	req := &pb.LoadCredentialsDataRequest{
		Token: "w/w/w",
		Uuid:  "initial UUID",
	}
	resp, err := suite.Server.LoadCredentials(suite.ctx, req)
	require.Error(suite.T(), models.ErrInvalidToken, err)
	require.Equal(suite.T(), "authorization error: wrong username", resp.Error)
}

func (suite *CredentialsTestSuite) TestLoadCredentialsNoSuchID() {
	req := &pb.LoadCredentialsDataRequest{
		Token: "initial login/initial password/salt",
		Uuid:  "wrong UUID",
	}
	resp, err := suite.Server.LoadCredentials(suite.ctx, req)
	require.Error(suite.T(), models.ErrDatabaseError, err)
	require.Equal(suite.T(), "internal server error for data wrong UUID", resp.Error)
}

func TestCredentialsTestSuite(t *testing.T) {
	suite.Run(t, new(CredentialsTestSuite))
}
