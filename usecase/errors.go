package usecase

import "errors"

var (
	ErrNotFound       = errors.New("not found")
	ErrInvalidBalance = errors.New("invalid balance")
	ErrUserNotFound   = errors.New("user not found")
)
