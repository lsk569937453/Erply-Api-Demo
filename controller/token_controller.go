package controller

import (
	"erply-api/constants"
	"erply-api/log"
	"erply-api/service"
	"erply-api/vojo"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Name will print hello name
// @Summary GetToken
// @Description Get the token
// @Tags GetToken
// @Security Bearer
// @Produce  json
// @Param userName query string true "userName"
// @Resource GetToken
// @Router /getToken [get]
// @Success 200 {string} string "{"resCode": 0,"message": "xxxxxx"}"
func GetToken(c *gin.Context) {

	var res vojo.BaseRes
	res.Rescode = constants.NORMAL_RESPONSE_STATUS

	tt, err := service.GetToken(c)
	if err != nil {
		res.Rescode = constants.ERROR_RESPONSE_STATUS
		res.Message = err.Error()
		log.Error("GetToken error", err.Error())
	} else {
		res.Message = tt
	}
	c.JSON(http.StatusOK, res)
}
