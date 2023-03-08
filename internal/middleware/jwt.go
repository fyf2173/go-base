package middleware

import (
	"go-base/internal/pkg/common"
	"net/http"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var AuthIgnorePath = map[string]bool{
	"/v1/user/gettoken": true,
	"/v1/example/test":  false,
	"/v1/sku/.*":        true,
	"/v1/goods/.*":      true,
}

func CheckAuthIgnoreRegPath(regSets map[string]bool, path string) bool {
	for rstr := range regSets {
		if ok, _ := regexp.MatchString(rstr, path); ok {
			return true
		}
	}
	return false
}

func ParseTokenMw() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			tokenStr = ctx.GetHeader("token")
			claim    common.Claim
		)
		token, err := jwt.ParseWithClaims(tokenStr, &claim, func(t *jwt.Token) (interface{}, error) {
			return []byte(common.SignedString), nil
		})
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, "Token not valid")
			return
		}
		if token.Claims.Valid() != nil {
			ctx.JSON(http.StatusUnauthorized, "Token not valid")
			return
		}
		if claim.AuthClient == common.AuthClientApp && (claim.UserClaim.UserId == 0 || claim.UserClaim.Openid == "") {
			ctx.JSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
			return
		}
		if claim.AuthClient == common.AuthClientConsole && claim.AdminClaim.Id == 0 {
			ctx.JSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
			return
		}
		if !claim.VerifyExpiresAt(time.Now().Unix(), false) {
			ctx.JSON(http.StatusUnauthorized, "Token expired")
			return
		}
		ctx.Set(common.CtxUserInfo, &claim)
		ctx.Next()
		return
	}
}
