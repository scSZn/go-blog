package dao

import (
	"context"
	"github.com/pkg/errors"
	"github.com/scSZn/blog/consts"
	"github.com/scSZn/blog/internal/model"
	"gorm.io/gorm"
)

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{
		db: db,
	}
}

func (d *UserDAO) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	var user *model.User

	db := d.db.Table(model.UserTableName).Where("username = ? AND is_del = ?", username, consts.NoDelStatus)
	if err := db.First(&user).Error; err != nil {
		return nil, errors.Wrap(err, "UserDAO.GetUserByUsername: get user fail, username is "+username)
	}

	return user, nil
}
