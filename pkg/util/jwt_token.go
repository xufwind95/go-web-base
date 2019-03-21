package util

import (
	"errors"
	"time"

	log "github.com/cihub/seelog"
	"github.com/dgrijalva/jwt-go"

	"github.com/xufwind95/go-web-base/config"
)

var (
	jwtSecret     string
	jwtExpiration int64
)

func InitJwtConfig(conf *config.Config) {
	jwtSecret = conf.JWT.Privatekey
	jwtExpiration = conf.JWT.Expiration
}

// 使用指定的密码生成token
func SignToken(id uint) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiration*10)).Unix(),
		"nbf": time.Now().Unix(),
		"iat": time.Now().Unix(),
	})

	tokenString, err = token.SignedString([]byte(jwtSecret))

	if err != nil {
		log.Info("Failed to sign a new token, due to " + err.Error())
	}

	return
}

// 验证秘钥格式是否正确
func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secret), nil
	}
}

var TokenError = errors.New("parse token failed")

// 验证token是否合法，合法的返回用户名和密码
func ParseToken(tokenString string) (uint64, error) {
	token, err := jwt.Parse(tokenString, secretFunc(jwtSecret))

	if err != nil {
		return 0, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id := uint64(claims["id"].(float64))
		return id, nil
	} else {
		return 0, TokenError
	}
}
