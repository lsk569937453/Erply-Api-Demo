package vojo

type GetTokenReq struct {
	UserAccount *string `form:"userAccount" binding:"required"`
}
