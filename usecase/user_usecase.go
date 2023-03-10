package usecase

import (
	"context"

	"github.com/alramdein/user-service/model"
	"github.com/alramdein/user-service/utils"
	log "github.com/sirupsen/logrus"
)

type userUsecase struct {
	userRepo model.UserRepository
}

func NewUserUsecase(userRepo model.UserRepository) model.UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (u *userUsecase) FindByUsernameAndPassword(ctx context.Context, username string, password string) (*model.User, error) {
	user, err := u.userRepo.FindByUsernameAndPassword(ctx, username, password)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if user == nil {
		log.Error(ErrNotFound)
		return nil, ErrNotFound
	}

	return user, nil
}

func (u *userUsecase) FindByID(ctx context.Context, userID int64) (*model.User, error) {
	user, err := u.userRepo.FindByID(ctx, userID)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if user == nil {
		log.Error(ErrNotFound)
		return nil, ErrNotFound
	}

	return user, nil
}

func (u *userUsecase) Create(ctx context.Context, input model.CreateUserInput) error {
	err := u.userRepo.Create(ctx, model.User{
		ID:       utils.GenerateUID(),
		Username: input.Username,
		Password: input.Password,
		Email:    input.Email,
		RoleID:   input.RoleID,
	})
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (u *userUsecase) Update(ctx context.Context, input model.User) error {
	user, err := u.userRepo.FindByID(ctx, input.ID)
	if err != nil {
		log.Error(err)
		return err
	}
	if user == nil {
		return ErrUserNotFound
	}

	err = u.userRepo.Update(ctx, input)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (u *userUsecase) Delete(ctx context.Context, userID int64) error {
	user, err := u.userRepo.FindByID(ctx, userID)
	if err != nil {
		log.Error(err)
		return err
	}
	if user == nil {
		return ErrUserNotFound
	}

	err = u.userRepo.Delete(ctx, userID)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
