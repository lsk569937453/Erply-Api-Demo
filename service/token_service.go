package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetToken(c *gin.Context) (*string, error) {
	userName := c.Query("userName")
	if userName == "" {
		return nil, fmt.Errorf("userName can not be empty!")
	}
	// 生成Token
	tokenString, err := GenToken(userName)
	return &tokenString, err
}
