package models

import "errors"

var (
	ErrDatabaseError     = errors.New("database internal error")
	ErrInvalidToken      = errors.New("invalid token")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrTimeout           = errors.New("timeout error")
	ErrUserNotFound      = errors.New("user not found")
)
