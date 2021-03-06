package util

import (
	"crypto/rand"
	"crypto/rsa"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"

	"github.com/scSZn/blog/internal/model"
)

const TokenValidTime = 24 * time.Hour
const Issuer = "scSZn"

type Claims struct {
	jwt.RegisteredClaims
	Uid      string `json:"uid"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Role     string `json:"role"`
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
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenValidTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    Issuer,
		},
		Uid:      user.Uid,
		Username: user.Username,
		Nickname: user.Nickname,
		Role:     user.Role,
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
