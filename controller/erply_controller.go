package controller

import (
	"erply-api/constants"
	"erply-api/log"
	"erply-api/service"
	"erply-api/vojo"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @param default  Authorization header string true "Authorization" default(eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3RBY2NvdW50IiwiZXhwIjoxNjE5MjE5Mzk2LCJpc3MiOiJlcnBseS1hcGkifQ._D6JcnV-5FeOhv6be1M5-6uaDLMqCOmiaoqyAnStjcE)
// @Summary AddCustomerRewardPoints
// @Description add customer reward points
// @Tags AddCustomerRewardPoints
// @Produce  json
// @Accept json
// @Param - body vojo.AddCustomerRewardPointsReq false " "
// @Router /api/AddCustomerRewardPoints [post]
// @Success 200 {string} string "{"resCode": 0,"message": {"transactionID": 24,"customerID": 17,"points": 90,"createdUnixTime": 1616634980,"expiryUnixTime": 0}}"
func AddCustomerRewardPoints(c *gin.Context) {
	var res vojo.BaseRes
	res.Rescode = constants.NORMAL_RESPONSE_STATUS

	tt, err, errCode := service.AddCustomerRewardPoints(c)
	if err != nil {
		if errCode != nil {
			res.Rescode = *errCode
		} else {
			res.Rescode = constants.ERROR_RESPONSE_STATUS
		}
		res.Message = err.Error()
		log.Error("AdminLogin error", err.Error())
	} else {
		res.Message = tt
	}
	c.JSON(http.StatusOK, res)
}

// @Security ApiKeyAuth
// @param default  Authorization header string true "Authorization" default(eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3RBY2NvdW50IiwiZXhwIjoxNjE5MjE5Mzk2LCJpc3MiOiJlcnBseS1hcGkifQ._D6JcnV-5FeOhv6be1M5-6uaDLMqCOmiaoqyAnStjcE)
// @Summary AddCustomerRewardPointsBulk
// @Description add customer reward points bulk
// @Tags AddCustomerRewardPointsBulk
// @Produce  json
// @Accept json
// @Param - body vojo.AddCustomerRewardPointsBulkReq false " "
// @Router /api/AddCustomerRewardPointsBulk [post]
// @Success 200 {string} string "{"resCode": 0,"message": {"transactionID": 24,"customerID": 17,"points": 90,"createdUnixTime": 1616634980,"expiryUnixTime": 0}}"
func AddCustomerRewardPointsBulk(c *gin.Context) {
	var res vojo.BaseRes
	res.Rescode = constants.NORMAL_RESPONSE_STATUS

	tt, err, errCode := service.AddCustomerRewardPointsBulk(c)
	if err != nil {
		if errCode != nil {
			res.Rescode = *errCode
		} else {
			res.Rescode = constants.ERROR_RESPONSE_STATUS
		}
		res.Message = err.Error()
		log.Error("AdminLogin error", err.Error())
	} else {
		res.Message = tt
	}
	c.JSON(http.StatusOK, res)
}

// @Security ApiKeyAuth
// @param default  Authorization header string true "Authorization" default(eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3RBY2NvdW50IiwiZXhwIjoxNjE5MjE5Mzk2LCJpc3MiOiJlcnBseS1hcGkifQ._D6JcnV-5FeOhv6be1M5-6uaDLMqCOmiaoqyAnStjcE)
// @Summary GetCustomersBulk
// @Description get customers bulk
// @Tags GetCustomersBulk
// @Produce  json
// @Accept json
// @Param - body vojo.GetCustomersBulkReq true " "
// @Router /api/GetCustomersBulk [post]
// @Success 200 {string} string "{"resCode": 0,"message": {"transactionID": 24,"customerID": 17,"points": 90,"createdUnixTime": 1616634980,"expiryUnixTime": 0}}"
func GetCustomersBulk(c *gin.Context) {
	var res vojo.BaseRes
	res.Rescode = constants.NORMAL_RESPONSE_STATUS

	tt, err, errCode := service.GetCustomersBulk(c)
	if err != nil {
		if errCode != nil {
			res.Rescode = *errCode
		} else {
			res.Rescode = constants.ERROR_RESPONSE_STATUS
		}
		res.Message = err.Error()
		log.Error("AdminLogin error", err.Error())
	} else {
		res.Message = tt
	}
	c.JSON(http.StatusOK, res)
}

// @Security ApiKeyAuth
// @param default  Authorization header string true "Authorization" default(eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3RBY2NvdW50IiwiZXhwIjoxNjE5MjE5Mzk2LCJpc3MiOiJlcnBseS1hcGkifQ._D6JcnV-5FeOhv6be1M5-6uaDLMqCOmiaoqyAnStjcE)
// @Summary GetCustomerByCustomerId
// @Description get customers by customerId
// @Tags GetCustomerByCustomerId
// @Produce  json
// @Accept json
// @Param default customerId query string true "customerId" default(13)
// @Router /api/GetCustomerByCustomerId [get]
// @Success 200 {string} string "{"resCode": 0,"message": {"transactionID": 24,"customerID": 17,"points": 90,"createdUnixTime": 1616634980,"expiryUnixTime": 0}}"
func GetCustomerByCustomerId(c *gin.Context) {
	var res vojo.BaseRes
	res.Rescode = constants.NORMAL_RESPONSE_STATUS

	tt, err, errCode := service.GetCustomerByCustomerId(c)
	if err != nil {
		if errCode != nil {
			res.Rescode = *errCode
		} else {
			res.Rescode = constants.ERROR_RESPONSE_STATUS
		}
		res.Message = err.Error()
		log.Error("AdminLogin error", err.Error())
	} else {
		res.Message = tt
	}
	c.JSON(http.StatusOK, res)
}
