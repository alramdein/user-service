package repository

import (
	"context"

	"github.com/alramdein/user-service/model"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) model.UserRepository {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) FindByUsernameAndPassword(ctx context.Context, username string, password string) (*model.User, error) {
	var user *model.User
	err := u.db.WithContext(ctx).Where("username = ? AND password = ?", username, password).Take(&user).Error
	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return nil, nil
	default:
		return nil, err
	}

	return user, nil
}

func (u *userRepo) FindByID(ctx context.Context, id int64) (*model.User, error) {
	var user *model.User
	err := u.db.WithContext(ctx).Where("id = ?", id).Take(&user).Error
	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return nil, nil
	default:
		return nil, err
	}

	return user, nil
}

func (u *userRepo) Create(ctx context.Context, user model.User) error {
	err := u.db.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (u *userRepo) Update(ctx context.Context, user model.User) error {
	err := u.db.Where("id = ?", user.ID).Updates(&model.User{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		RoleID:   user.RoleID,
	}).Error
	if err != nil {
		return err
	}

	return nil
}

func (u *userRepo) Delete(ctx context.Context, userID int64) error {
	err := u.db.Delete(&model.User{}, userID).Error
	if err != nil {
		return err
	}

	return nil
}
