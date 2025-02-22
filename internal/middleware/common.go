package middleware

import (
	"go-base/internal/pkg/common"
	"net/http"
	"regexp"
	"time"

	"github.com/fyf2173/ysdk-go/xctx"
	"github.com/fyf2173/ysdk-go/xhttp"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

// Access 生成唯一的traceId.
func Access() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		traceID := ctx.GetHeader(xhttp.HeaderTraceId)
		if traceID == "" {
			traceID = uuid.New().String()
		}
		ctx.Set(xctx.TraceId, traceID)
		ctx.Next()
	}
}

func CheckAuthIgnoreRegPath(regSets []string, path string) bool {
	for _, rstr := range regSets {
		if ok, _ := regexp.MatchString(rstr, path); ok {
			return true
		}
	}
	return false
}

func CommonTokenMw(ignoreAuthPaths ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			tokenStr = ctx.GetHeader("token")
			claim    common.Claim
		)
		if CheckAuthIgnoreRegPath(ignoreAuthPaths, ctx.Request.URL.Path) {
			ctx.Next()
			return
		}
		if tokenStr == "" {
			tokenStr = ctx.Query("token")
		}
		token, err := jwt.ParseWithClaims(tokenStr, &claim, func(*jwt.Token) (interface{}, error) {
			return []byte(common.SignedString), nil
		})
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, "Token not valid")
			ctx.Abort()
			return
		}
		if token.Claims.Valid() != nil {
			ctx.JSON(http.StatusUnauthorized, "Token not valid")
			ctx.Abort()
			return
		}
		if claim.AuthClient == common.AuthClientApp && (claim.UserClaim.UserId == 0 || claim.UserClaim.Openid == "") {
			ctx.JSON(http.StatusUnauthorized, claim.AuthClient+" "+http.StatusText(http.StatusUnauthorized))
			ctx.Abort()
			return
		}
		if claim.AuthClient == common.AuthClientConsole && claim.AdminClaim.Id == 0 {
			ctx.JSON(http.StatusUnauthorized, claim.AuthClient+" "+http.StatusText(http.StatusUnauthorized))
			ctx.Abort()
			return
		}
		if !claim.VerifyExpiresAt(time.Now().Unix(), false) {
			ctx.JSON(http.StatusUnauthorized, "Token expired")
			ctx.Abort()
			return
		}
		ctx.Set(common.CtxUserInfo, &claim)
		ctx.Next()
		return
	}
}
