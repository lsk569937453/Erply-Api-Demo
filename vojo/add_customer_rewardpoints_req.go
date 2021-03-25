package vojo

type AddCustomerRewardPointsReq struct {
	CustomerID *string `form:"customerID" binding:"required" json:"customerID" example:"17"`
	Points     *string `form:"points" binding:"required" json:"points"  example:"90"`
}
