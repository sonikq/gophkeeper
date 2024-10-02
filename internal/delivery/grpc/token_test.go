package grpc

import (
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TokenTestSuite struct {
	suite.Suite
}

func (suite *TokenTestSuite) TestEncodeToken() {
	sampleToken := Token{
		Login:    "test login",
		Password: "test password",
		salt:     "test salt",
	}
	result := EncodeToken(sampleToken)
	require.Equal(suite.T(), "test login/test password/test salt", result)
}

func (suite *TokenTestSuite) TestDecodeToken() {
	testToken := "test login/test password/test salt"
	result := DecodeToken(testToken)
	require.Equal(suite.T(), Token{
		Login:    "test login",
		Password: "test password",
		salt:     "test salt",
	}, result)
}

func TestTokenTestSuite(t *testing.T) {
	suite.Run(t, new(TokenTestSuite))
}
