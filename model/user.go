package model

import (
	"context"

	"github.com/alramdein/user-service/pb"
)

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
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, userID int64) error
}

type UserUsecase interface {
	FindByUsernameAndPassword(ctx context.Context, username string, password string) (*User, error)
	FindByID(ctx context.Context, userID int64) (*User, error)
	Create(ctx context.Context, input CreateUserInput) error
	Update(ctx context.Context, input User) error
	Delete(ctx context.Context, userID int64) error
}

func (u *User) ToProto() *pb.User {
	return &pb.User{
		Id:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		RoleId:   u.RoleID,
	}
}
