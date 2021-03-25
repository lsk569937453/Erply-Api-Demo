package vojo

type AddCustomerRewardPointsBulkSingle struct {
	CustomerID *string `form:"customerID" binding:"required" structs:"customerID" json:"customerID"`
	Points     *string `form:"points" binding:"required" structs:"points" json:"points"`
}
type AddCustomerRewardPointsBulkReq struct {
	CustomerList []AddCustomerRewardPointsBulkSingle `form:"customerList" binding:"required,dive"`
}
