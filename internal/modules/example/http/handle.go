package http

import (
	"github.com/fyf2173/ysdk-go/apisdk/ginplus"
	"github.com/gin-gonic/gin"
)

func Test(ctx *gin.Context) {
	ginplus.ExitSuccess(ctx, "Gin hello world")
	return
}
