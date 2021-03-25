package service

import (
	"context"
	"erply-api/client"
	"erply-api/constants"
	"erply-api/dao"
	"erply-api/vojo"
	"fmt"
	"time"

	"github.com/fatih/structs"

	"github.com/gin-gonic/gin"

	"github.com/erply/api-go-wrapper/pkg/api/customers"
)

func GetCustomersBulk(c *gin.Context) (*customers.GetCustomersResponseBulk, error, *int) {
	customerCli := client.ApiClient.CustomerManager
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
	cli := client.ApiClient.CustomerManager

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
	cli := client.ApiClient.CustomerManager
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

func GetCustomerByCustomerId(c *gin.Context) (*dao.CustomerDao, error, *int) {
	customerId := c.Query("customerId")
	if customerId == "" {
		return nil, fmt.Errorf("userName can not be empty!"), nil
	}
	var res dao.CustomerDao
	err := dao.CronDb.Where("customer_id=?", customerId).First(&res).Error
	if err != nil {
		errCode := constants.DATA_BASE_CALL_ERROR
		return nil, err, &errCode
	}
	return &res, nil, nil
}
