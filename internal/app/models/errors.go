package models

import "errors"

var (
	ErrDatabaseError       = errors.New("database internal error")
	ErrInvalidToken        = errors.New("invalid token")
	ErrUserAlreadyExists   = errors.New("user already exists")
	ErrInMemoryDB          = errors.New("InMemoryDB internal error")
	ErrUserNotFound        = errors.New("user not found")
	ErrContextTimeout      = errors.New("context timeout called")
	ErrDatabaseUnreachable = errors.New("database unreachable")
)
