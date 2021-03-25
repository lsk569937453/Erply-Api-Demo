package service

import (
	"erply-api/vojo"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const TokenExpireDuration = time.Hour * 24 * 30

var MySecret = []byte("so9Yb_*(62~/jghi")

// GenToken 生成JWT
func GenToken(userLoginName string) (string, error) {

	// 创建一个我们自己的声明
	c := vojo.UserClaims{
		userLoginName, // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "erply-api",                                // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*vojo.UserClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &vojo.UserClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*vojo.UserClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
