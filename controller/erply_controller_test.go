package controller

import (
	"encoding/json"
	"erply-api/constants"
	"erply-api/util"
	"erply-api/vojo"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

const TOKEN = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3RBY2NvdW50IiwiZXhwIjoxNjE5MjE5Mzk2LCJpc3MiOiJlcnBseS1hcGkifQ._D6JcnV-5FeOhv6be1M5-6uaDLMqCOmiaoqyAnStjcE"

func TestAddCustomerRewardPointsSuccess(t *testing.T) {
	var w *httptest.ResponseRecorder
	assert := assert.New(t)
	// 1.测试 index 请求
	urlIndex := "/api/AddCustomerRewardPoints"
	customId := "13"
	points := "50"
	body := vojo.AddCustomerRewardPointsReq{
		CustomerID: &customId,
		Points:     &points,
	}
	headers := map[string]string{
		"Authorization": TOKEN,
	}
	w = util.PostJson(urlIndex, body, router, headers)
	assert.Equal(200, w.Code)
	resBody := w.Body.String()
	var baseRes vojo.BaseRes
	err := json.Unmarshal([]byte(resBody), &baseRes)
	assert.Equal(err, nil)
	assert.Equal(constants.NORMAL_RESPONSE_STATUS, baseRes.Rescode)
	fmt.Println("req result is:" + resBody)
}
func TestAddCustomerRewardPointsError(t *testing.T) {

	var w *httptest.ResponseRecorder
	assert := assert.New(t)
	// 1.测试 index 请求
	urlIndex := "/api/AddCustomerRewardPoints"
	customId := "1111111111111111"
	points := "50"
	body := vojo.AddCustomerRewardPointsReq{
		CustomerID: &customId,
		Points:     &points,
	}
	headers := map[string]string{
		"Authorization": TOKEN,
	}
	w = util.PostJson(urlIndex, body, router, headers)
	assert.Equal(200, w.Code)
	resBody := w.Body.String()
	var baseRes vojo.BaseRes
	err := json.Unmarshal([]byte(resBody), &baseRes)
	assert.Equal(err, nil)
	assert.Equal(constants.CALL_ERPLY_API_ERROR, baseRes.Rescode)
}
func TestAddCustomerRewardPointsBulkSuccess(t *testing.T) {
	var w *httptest.ResponseRecorder
	assert := assert.New(t)
	// 1.测试 index 请求
	urlIndex := "/api/AddCustomerRewardPointsBulk"
	customId1 := "13"
	points1 := "50"
	custom1 := vojo.AddCustomerRewardPointsBulkSingle{
		CustomerID: &customId1,
		Points:     &points1,
	}
	customId2 := "17"
	points2 := "300"
	custom2 := vojo.AddCustomerRewardPointsBulkSingle{
		CustomerID: &customId2,
		Points:     &points2,
	}
	customList := make([]vojo.AddCustomerRewardPointsBulkSingle, 0)
	customList = append(customList, custom1)
	customList = append(customList, custom2)

	body := vojo.AddCustomerRewardPointsBulkReq{
		CustomerList: customList,
	}

	headers := map[string]string{
		"Authorization": TOKEN,
	}
	w = util.PostJson(urlIndex, body, router, headers)
	assert.Equal(200, w.Code)
	resBody := w.Body.String()
	var baseRes vojo.BaseRes
	err := json.Unmarshal([]byte(resBody), &baseRes)
	assert.Equal(err, nil)
	assert.Equal(constants.NORMAL_RESPONSE_STATUS, baseRes.Rescode)
	fmt.Println("req result is:" + resBody)
}
func TestAddCustomerRewardPointsBulkError(t *testing.T) {
	var w *httptest.ResponseRecorder
	assert := assert.New(t)
	// 1.测试 index 请求
	urlIndex := "/api/AddCustomerRewardPointsBulk"
	customId1 := "111111111111"
	points1 := "50"
	custom1 := vojo.AddCustomerRewardPointsBulkSingle{
		CustomerID: &customId1,
		Points:     &points1,
	}
	customId2 := "17"
	points2 := "2222222222222"
	custom2 := vojo.AddCustomerRewardPointsBulkSingle{
		CustomerID: &customId2,
		Points:     &points2,
	}
	customList := make([]vojo.AddCustomerRewardPointsBulkSingle, 0)
	customList = append(customList, custom1)
	customList = append(customList, custom2)

	body := vojo.AddCustomerRewardPointsBulkReq{
		CustomerList: customList,
	}

	headers := map[string]string{
		"Authorization": TOKEN,
	}
	w = util.PostJson(urlIndex, body, router, headers)
	assert.Equal(200, w.Code)
	resBody := w.Body.String()
	var baseRes vojo.BaseRes
	err := json.Unmarshal([]byte(resBody), &baseRes)
	assert.Equal(err, nil)
	assert.Equal(constants.CALL_ERPLY_API_ERROR, baseRes.Rescode)
	fmt.Println("req result is:" + resBody)
}
func TestGetCustomersBulkSuccess(t *testing.T) {
	var w *httptest.ResponseRecorder
	assert := assert.New(t)
	// 1.测试 index 请求
	urlIndex := "/api/GetCustomersBulk"
	recordsOnPage1 := "100"
	pageNo1 := "0"
	page1 := vojo.GetCustomersSingle{
		RecordsOnPage: &recordsOnPage1,
		PageNo:        &pageNo1,
	}
	recordsOnPage2 := "100"
	pageNo2 := "1"
	pge2 := vojo.GetCustomersSingle{
		RecordsOnPage: &recordsOnPage2,
		PageNo:        &pageNo2,
	}
	pageList := make([]vojo.GetCustomersSingle, 0)
	pageList = append(pageList, page1)
	pageList = append(pageList, pge2)

	body := vojo.GetCustomersBulkReq{
		PageList: pageList,
	}
	headers := map[string]string{
		"Authorization": TOKEN,
	}
	w = util.PostJson(urlIndex, body, router, headers)
	assert.Equal(200, w.Code)
	resBody := w.Body.String()
	var baseRes vojo.BaseRes
	err := json.Unmarshal([]byte(resBody), &baseRes)
	assert.Equal(err, nil)
	assert.Equal(constants.NORMAL_RESPONSE_STATUS, baseRes.Rescode)
}
func TestGetCustomersBulkError(t *testing.T) {
	var w *httptest.ResponseRecorder
	assert := assert.New(t)
	// 1.测试 index 请求
	urlIndex := "/api/GetCustomersBulk"
	recordsOnPage1 := "100"
	pageNo1 := "0"
	page1 := vojo.GetCustomersSingle{
		RecordsOnPage: &recordsOnPage1,
		PageNo:        &pageNo1,
	}
	recordsOnPage2 := "100"
	pageNo2 := "-1"
	pge2 := vojo.GetCustomersSingle{
		RecordsOnPage: &recordsOnPage2,
		PageNo:        &pageNo2,
	}
	pageList := make([]vojo.GetCustomersSingle, 0)
	pageList = append(pageList, page1)
	pageList = append(pageList, pge2)

	body := vojo.GetCustomersBulkReq{
		PageList: pageList,
	}
	headers := map[string]string{
		"Authorization": TOKEN,
	}
	w = util.PostJson(urlIndex, body, router, headers)
	assert.Equal(200, w.Code)
	resBody := w.Body.String()
	var baseRes vojo.BaseRes
	err := json.Unmarshal([]byte(resBody), &baseRes)
	assert.Equal(err, nil)
	assert.Equal(constants.CALL_ERPLY_API_ERROR, baseRes.Rescode)
}
