package service

import (
	"context"
	"database/sql"
	"encoding/hex"
	"github.com/scSZn/blog/pkg/logger"

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
	userDao := dao.NewUserDAO(ls.ctx, ls.db)

	user, err := userDao.GetUserByUsername(request.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			global.Logger.Errorf(ls.ctx, logger.Fields{"params": request}, err, "login fail, no such user")
			return "", errcode.LoginFail
		} else {
			global.Logger.Errorf(ls.ctx, logger.Fields{"params": request}, err, "login fail, mysql error")
			return "", errcode.ServerError
		}
	}

	realPassport, err := hex.DecodeString(user.Passport)
	if err != nil {
		global.Logger.Errorf(ls.ctx, logger.Fields{"params": request}, err, "decode user real passport fail")
		return "", errcode.ServerError
	}

	if err = bcrypt.CompareHashAndPassword(realPassport, []byte(request.Password)); err != nil {
		global.Logger.Errorf(ls.ctx, logger.Fields{"params": request}, err, "login fail, passport incorrect")
		return "", errcode.ServerError
	}

	token, err := util.GenerateToken(user)
	if err != nil {
		global.Logger.Errorf(ls.ctx, logger.Fields{"params": request}, err, "generate token fail")
		return "", errcode.ServerError
	}

	return token, nil
}
