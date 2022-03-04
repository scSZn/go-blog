package dao

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/scSZn/blog/consts"
	"github.com/scSZn/blog/internal/model"
)

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(ctx context.Context, db *gorm.DB) *UserDAO {
	return &UserDAO{
		db: db.WithContext(ctx),
	}
}

func (d *UserDAO) GetUserByUsername(username string) (*model.User, error) {
	var user *model.User

	db := d.db.Table(model.UserTableName).Where("username = ? AND is_del = ?", username, consts.NoDelStatus)
	if err := db.First(&user).Error; err != nil {
		return nil, errors.Wrap(err, "UserDAO.GetUserByUsername: get user fail, username is "+username)
	}

	return user, nil
}
