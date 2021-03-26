package vojo

type AddCustomerRewardPointsBulkSingle struct {
	CustomerID *string `form:"customerID" binding:"required" structs:"customerID" json:"customerID" example:"13"`
	Points     *string `form:"points" binding:"required" structs:"points" json:"points" example:"90"`
}
type AddCustomerRewardPointsBulkReq struct {
	CustomerList []AddCustomerRewardPointsBulkSingle `form:"customerList" binding:"required,dive"`
}
