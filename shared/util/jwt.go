package util

import (
	"douyin/shared/constant"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWT struct {
	SigningKey []byte
}

func NewJWT(secretKey string) *JWT {
	return &JWT{
		SigningKey: []byte(secretKey),
	}
}

type Claims struct {
	UserID int64
	jwt.StandardClaims
}

func (s *JWT) GenerateToken(userID int64) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(constant.TokenExpireTime)
	claims := Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(s.SigningKey)
	return token, err
}

func (s *JWT) ParseToken(token string) (*Claims, error) {
	claims := new(Claims)
	tokenClaims, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return s.SigningKey, nil
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

	return nil, errors.New("token time out")
}
