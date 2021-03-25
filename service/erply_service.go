package service

import (
	"context"
	"crypto/tls"
	"erply-api/constants"
	"erply-api/internal/common"
	"erply-api/vojo"
	"net/http"

	"github.com/fatih/structs"

	"github.com/gin-gonic/gin"

	"github.com/erply/api-go-wrapper/pkg/api/auth"

	"github.com/erply/api-go-wrapper/pkg/api/customers"

	"github.com/erply/api-go-wrapper/pkg/api"

	"time"
)

var apiClient *api.Client

func init() {
	const (
		username   = "lsk569937453@gmail.com"
		password   = "demo1234"
		clientCode = "113073"
		partnerKey = "113073"
	)

	apiClientInit, err := buildClient()
	common.Die(err)
	apiClient = apiClientInit

}
func buildClient() (*api.Client, error) {

	username := "lsk569937453@gmail.com"
	password := "demo1234"
	clientCode := "113073"

	connectionTimeout := 60 * time.Second
	transport := &http.Transport{
		DisableKeepAlives:     true,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
		ResponseHeaderTimeout: connectionTimeout,
	}
	httpCl := &http.Client{Transport: transport}

	sessionKey, err := auth.VerifyUser(username, password, clientCode, http.DefaultClient)
	if err != nil {
		panic(err)
	}

	apiClient, err := api.NewClient(sessionKey, clientCode, httpCl)
	if err != nil {
		panic(err)
	}

	return apiClient, nil
}

func GetCustomersBulk(c *gin.Context) (*customers.GetCustomersResponseBulk, error, *int) {
	customerCli := apiClient.CustomerManager
	var reqx vojo.GetCustomersBulkReq
	err := c.Bind(&reqx)
	if err != nil {
		return nil, err, nil
	}
	reqArray := make([]map[string]interface{}, 0)
	for _, value := range reqx.PageList {
		mapResult := structs.Map(value)
		reqArray = append(reqArray, mapResult)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	bulkResp, err := customerCli.GetCustomersBulk(ctx, reqArray, map[string]string{})
	if err != nil {
		errCode := constants.CALL_ERPLY_API_ERROR
		return &bulkResp, err, &errCode
	}
	return &bulkResp, err, nil

}

func AddCustomerRewardPoints(c *gin.Context) (*customers.AddCustomerRewardPointsResult, error, *int) {

	var reqx vojo.AddCustomerRewardPointsReq
	err := c.Bind(&reqx)
	if err != nil {
		return nil, err, nil
	}
	cli := apiClient.CustomerManager

	req := map[string]string{
		"customerID": *reqx.CustomerID,
		"points":     *reqx.Points,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := cli.AddCustomerRewardPoints(ctx, req)
	if err != nil {
		errCode := constants.CALL_ERPLY_API_ERROR
		return &resp, err, &errCode
	}
	return &resp, err, nil
}

func AddCustomerRewardPointsBulk(c *gin.Context) (*customers.AddCustomerRewardPointsResponseBulk, error, *int) {
	cli := apiClient.CustomerManager
	var reqx vojo.AddCustomerRewardPointsBulkReq
	err := c.Bind(&reqx)
	if err != nil {
		return nil, err, nil
	}
	reqArray := make([]map[string]interface{}, 0)
	for _, value := range reqx.CustomerList {
		mapResult := structs.Map(value)
		reqArray = append(reqArray, mapResult)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := cli.AddCustomerRewardPointsBulk(ctx, reqArray, map[string]string{})

	if err != nil {
		errCode := constants.CALL_ERPLY_API_ERROR
		return &resp, err, &errCode
	}
	return &resp, err, nil

}
