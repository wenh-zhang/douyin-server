package util

import (
	"douyin/pkg/constant"
	"douyin/pkg/errno"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	UserID int64
	jwt.StandardClaims
}

func GenerateToken(userID int64) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(constant.TokenExpireTime)
	claims := Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(constant.TokenSignedKey)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	claims := new(Claims)
	tokenClaims, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return constant.TokenSignedKey, nil
	})
	if err != nil {
		return nil, err
	}

	if tokenClaims != nil {
		if tokenClaims.Valid {
			return claims, nil
		}

		//todo 超时处理逻辑
	}

	return nil, errno.TokenTimeOutErr
}
