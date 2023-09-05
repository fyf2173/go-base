package api

import (
	"github.com/fyf2173/ysdk-go/apisdk/ginplus"
	"github.com/gin-gonic/gin"
)

func test(ctx *gin.Context) {
	ginplus.ExitSuccess(ctx, "Gin hello world from service api [linux]")
	return
}
