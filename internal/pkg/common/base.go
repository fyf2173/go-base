package common

import "github.com/gin-gonic/gin"

const (
	CtxUserInfo  = "userinfo"
	CtxAdminInfo = "admin_info"
)

var DefaultAdminClaim = AdminClaim{Name: "系统", Id: 0}

func GetAdminInfo(ctx *gin.Context) *AdminClaim {
	if v, ok := ctx.Get(CtxAdminInfo); ok {
		return &v.(*Claim).AdminClaim
	}
	return nil
}

func GetUserInfo(ctx *gin.Context) *UserClaim {
	if userinfo, ok := ctx.Get(CtxUserInfo); ok {
		return &userinfo.(*Claim).UserClaim
	}
	return nil
}
