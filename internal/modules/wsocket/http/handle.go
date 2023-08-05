package http

import (
	"go-base/internal/modules/wsocket"

	"github.com/gin-gonic/gin"
)

func connectWs(ctx *gin.Context) {
	wsocket.ServeWs(wsocket.HubInstance, ctx.Writer, ctx.Request)
	return
}
