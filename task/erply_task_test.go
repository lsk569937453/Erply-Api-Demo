package task

import (
	"erply-api/dao"
	"testing"

	"github.com/erply/api-go-wrapper/pkg/api/customers"
	"github.com/stretchr/testify/assert"
)

func TestDoReqSuccess(t *testing.T) {
	assert := assert.New(t)

	err := doReq()
	assert.Equal(err, nil)
}

func TestInsertDataSuccess(t *testing.T) {
	assert := assert.New(t)
	customArray := make([]customers.Customer, 0)
	customer1 := customers.Customer{
		FullName:   "ssss",
		CustomerID: 1487,
	}
	customArray = append(customArray, customer1)
	err := UpdateData(customArray)

	var res dao.CustomerDao
	err = dao.CronDb.Where("customer_id=?", 1487).First(&res).Error
	assert.Equal(err, nil)
	assert.Equal(res.CustomerID, 1487)
	assert.Equal(res.FullName, "ssss")
}

func TestUpdateDataSuccess(t *testing.T) {
	assert := assert.New(t)
	customArray := make([]customers.Customer, 0)
	customer1 := customers.Customer{
		FullName:   "ssss2",
		CustomerID: 1487,
	}
	customArray = append(customArray, customer1)
	err := UpdateData(customArray)
	assert.Equal(err, nil)

	var res dao.CustomerDao
	err = dao.CronDb.Where("customer_id=?", 1487).First(&res).Error
	assert.Equal(err, nil)
	assert.Equal(res.CustomerID, 1487)
	assert.Equal(res.FullName, "ssss2")
}
