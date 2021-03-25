package vojo

type GetCustomersSingle struct {
	RecordsOnPage *string `form:"recordsOnPage" binding:"required" structs:"recordsOnPage"`
	PageNo        *string `form:"pageNo" binding:"required" structs:"pageNo"`
}
type GetCustomersBulkReq struct {
	PageList []GetCustomersSingle `form:"pageList" binding:"required,dive"`
}
