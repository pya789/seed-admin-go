package utils

import (
	"seed-admin/common"

	"github.com/golang-jwt/jwt"
)

type Jwt struct {
	SigningKey []byte
}
type CustomerClaims struct {
	*jwt.StandardClaims     //标准字段
	UserId              int `json:"user_id"`
}

func NewJwt(key ...[]byte) *Jwt {
	if len(key) != 0 {
		return &Jwt{key[0]}
	}
	return &Jwt{[]byte(common.CONFIG.String("jwt.signingKey"))}
}

// 创建Token
func (j *Jwt) CreateToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// 解析Token
func (j *Jwt) ParseToken(tokenString string) (*CustomerClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomerClaims{}, func(t *jwt.Token) (any, error) { return j.SigningKey, nil })
	if err != nil {
		return nil, err
	}
	return token.Claims.(*CustomerClaims), nil
}
