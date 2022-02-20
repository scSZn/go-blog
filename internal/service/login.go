package service

import (
	"context"
	"github.com/scSZn/blog/global"
	"gorm.io/gorm"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginService struct {
	ctx context.Context
	db  *gorm.DB
}

func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{
		ctx: ctx,
		db:  global.DB,
	}
}

func (ls *LoginService) Login(request *LoginRequest) error {
	return nil
}
