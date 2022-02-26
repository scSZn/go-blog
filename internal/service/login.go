package service

import (
	"context"
	"database/sql"
	"encoding/hex"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/scSZn/blog/global"
	"github.com/scSZn/blog/internal/dao"
	"github.com/scSZn/blog/pkg/errcode"
	"github.com/scSZn/blog/pkg/util"
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

func (ls *LoginService) Login(request *LoginRequest) (string, error) {
	userDao := dao.NewUserDAO(ls.db)

	user, err := userDao.GetUserByUsername(ls.ctx, request.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			global.Logger.Errorf(ls.ctx, "LoginService.Login: login fail, no such user,request: %+v， err: %+v", request, err)
			return "", errcode.LoginFail
		} else {
			global.Logger.Errorf(ls.ctx, "LoginService.Login: login fail, request: %+v, err: %+v", request, err)
			return "", errcode.ServerError
		}
	}

	realPassport, err := hex.DecodeString(user.Passport)
	if err != nil {
		global.Logger.Errorf(ls.ctx, "LoginService.Login: decode user real passport fail, request: %+v, err: %+v", request, err)
		return "", errcode.ServerError
	}

	if err = bcrypt.CompareHashAndPassword(realPassport, []byte(request.Password)); err != nil {
		global.Logger.Errorf(ls.ctx, "LoginService.Login: login fail, passport incorrect, request: %+v, err: %+v", request, err)
		return "", errcode.ServerError
	}

	token, err := util.GenerateToken(user)
	if err != nil {
		global.Logger.Errorf(ls.ctx, "LoginService.Login: generate token fail, user: %+v, err: %+v", user, err)
		return "", errcode.ServerError
	}

	return token, nil
}
