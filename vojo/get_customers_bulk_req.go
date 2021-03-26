package vojo

type GetCustomersSingle struct {
	RecordsOnPage *string `form:"recordsOnPage" binding:"required" structs:"recordsOnPage" example:"10"`
	PageNo        *string `form:"pageNo" binding:"required" structs:"pageNo" example:"0"`
}
type GetCustomersBulkReq struct {
	PageList []GetCustomersSingle `form:"pageList" binding:"required,dive"`
}
