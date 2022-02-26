package util

import (
	"crypto/rand"
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/scSZn/blog/internal/model"
	"time"
)

const TokenValidTime = 24 * time.Hour
const Issuer = "scSZn"

type Claims struct {
	jwt.StandardClaims
	Uid      string `json:"uid"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
}

var privateKey, publicKey = genRsaKey()

//RSA公钥私钥产生
func genRsaKey() (*rsa.PrivateKey, *rsa.PublicKey) {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}

	return privateKey, &privateKey.PublicKey
}

func GenerateToken(user *model.User) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodRS256, &Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenValidTime).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    Issuer,
		},
		Uid:      user.Uid,
		Username: user.Username,
		Nickname: user.Nickname,
	})

	signedString, err := claims.SignedString(privateKey)
	if err != nil {
		return "", errors.Wrapf(err, "util.GenerateToken: generate jwt token fail, user: %+v", user)
	}

	return signedString, nil
}

func ValidToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*Claims)
		if ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
