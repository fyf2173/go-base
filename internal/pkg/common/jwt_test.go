package common

import (
	"testing"

	"github.com/dgrijalva/jwt-go"
)

func TestNewToken(t *testing.T) {
	t.Log(jwt.NewWithClaims(jwt.SigningMethodHS256, Claim{UserClaim: UserClaim{UserId: 10000, Openid: "abcd"}}).SignedString([]byte(SignedString)))
	// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEwMDAwLCJPcGVuaWQiOiJhYmNkIn0.SEDgidMh6l6G7F_4hBL2d_8zWcuGoBaG2PV2xBEgjKs
}

func TestParseToken(t *testing.T) {
	var (
		tokenStr = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEwMDAwLCJPcGVuaWQiOiJhYmNkIn0.SEDgidMh6l6G7F_4hBL2d_8zWcuGoBaG2PV2xBEgjKs"
		claim    Claim
	)
	tokenInfo, err := jwt.ParseWithClaims(tokenStr, &claim, func(t *jwt.Token) (interface{}, error) {
		return []byte(SignedString), nil
	})
	t.Log(err)
	t.Logf("%#v", tokenInfo)
	t.Logf("%+v", claim.Valid())
	t.Logf("valide:%v %#v", tokenInfo.Claims.Valid(), tokenInfo.Claims.(*Claim))
	t.Log(claim)
}
