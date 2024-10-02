package grpc

import (
	"fmt"
	"strings"
)

type Token struct {
	Login    string
	Password string
	salt     string
}

func EncodeToken(token Token) string {
	return fmt.Sprintf("%s/%s/%s", token.Login, token.Password, token.salt)
}

func DecodeToken(input string) Token {
	result := strings.Split(input, "/")
	login, password, salt := result[0], result[1], result[2]
	return Token{
		Login:    login,
		Password: password,
		salt:     salt,
	}
}
