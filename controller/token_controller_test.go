package controller

import (
	"encoding/json"
	"erply-api/constants"
	"erply-api/midware"
	"erply-api/util"
	"erply-api/vojo"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func TestMain(m *testing.M) {

	fmt.Println("test begin")
	InitRouter()
	m.Run()
	fmt.Println("test end")
}
func InitRouter() {
	r := gin.New()
	r.GET("/getToken", GetToken)

	v1 := r.Group("/api")
	{
		v1.Use(midware.JWTAuthMiddleware())

		v1.POST("/AddCustomerRewardPoints", AddCustomerRewardPoints)
		v1.POST("/AddCustomerRewardPointsBulk", AddCustomerRewardPointsBulk)
		v1.POST("/GetCustomersBulk", GetCustomersBulk)

	}
	router = r
}
func TestGetTokenSuccess(t *testing.T) {

	var w *httptest.ResponseRecorder
	assert := assert.New(t)

	// 1.测试 index 请求
	urlIndex := "/getToken?userName=testUser"
	w = util.Get(urlIndex, router)
	assert.Equal(200, w.Code)
	body := w.Body.String()
	var baseRes vojo.BaseRes
	err := json.Unmarshal([]byte(body), &baseRes)
	assert.Equal(err, nil)
	assert.Equal(constants.NORMAL_RESPONSE_STATUS, baseRes.Rescode)
	fmt.Println("user Token is:" + baseRes.Message.(string))
}
func TestGetTokenError(t *testing.T) {

	var w *httptest.ResponseRecorder
	assert := assert.New(t)

	// 1.测试 index 请求
	urlIndex := "/getToken"
	w = util.Get(urlIndex, router)
	assert.Equal(200, w.Code)
	body := w.Body.String()
	var baseRes vojo.BaseRes
	err := json.Unmarshal([]byte(body), &baseRes)
	assert.Equal(err, nil)
	assert.Equal(constants.ERROR_RESPONSE_STATUS, baseRes.Rescode)
}
