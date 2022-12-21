package model

import "context"

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleID   int64  `json:"role_id"`
}

type CreateUserInput struct {
	Username string
	Email    string
	Password string
	RoleID   int64
}

type UserRepository interface {
	FindByUsernameAndPassword(ctx context.Context, username string, password string) (*User, error)
	FindByID(ctx context.Context, userID int64) (*User, error)
	Create(ctx context.Context, user User) error
}

type UserUsecase interface {
	FindByUsernameAndPassword(ctx context.Context, username string, password string) (*User, error)
	FindByID(ctx context.Context, userID int64) (*User, error)
	Create(ctx context.Context, input CreateUserInput) error
}
