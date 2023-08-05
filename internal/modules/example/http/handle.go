package http

import (
	"go-base/internal/pkg/common"

	"github.com/dgrijalva/jwt-go"
	"github.com/fyf2173/ysdk-go/apisdk/ginplus"
	"github.com/gin-gonic/gin"
)

func TestIgnoreAuth(ctx *gin.Context) {
	ginplus.ExitSuccess(ctx, "Gin hello world from console api ignore auth")
	return
}

func TestWithAuth(ctx *gin.Context) {
	ginplus.ExitSuccess(ctx, "Gin hello world from console api with auth")
	return
}

func TestGetToken(ctx *gin.Context) {
	miniappClaim := common.Claim{AdminClaim: common.AdminClaim{Id: 10000, Name: "abcd"}, AuthClient: "console"}
	jwtToken, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, miniappClaim).SignedString([]byte(common.SignedString))
	ginplus.ExitSuccess(ctx, jwtToken)
	return
}
