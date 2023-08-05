package common

import (
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

func TestMiniappToken(t *testing.T) {
	miniappClaim := Claim{UserClaim: UserClaim{UserId: 10000, Openid: "abcd"}, AuthClient: "miniapp"}
	jwtToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, miniappClaim).SignedString([]byte(SignedString))
	assert.Nil(t, err, "NewWithClaims is not nil")

	var newClaim Claim
	tokenInfo, err := jwt.ParseWithClaims(jwtToken, &newClaim, func(t *jwt.Token) (interface{}, error) { return []byte(SignedString), nil })
	assert.Nil(t, err, "ParseWithClaims return err")

	assert.True(t, tokenInfo.Valid, "tokenInfo is not valid")

	assert.Equal(t, int64(10000), newClaim.UserClaim.UserId)
	assert.Equal(t, "abcd", newClaim.UserClaim.Openid)
}

func TestConsoleToken(t *testing.T) {
	miniappClaim := Claim{AdminClaim: AdminClaim{Id: 10000, Name: "abcd"}, AuthClient: "miniapp"}
	jwtToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, miniappClaim).SignedString([]byte(SignedString))
	assert.Nil(t, err, "NewWithClaims is not nil")

	var newClaim Claim
	tokenInfo, err := jwt.ParseWithClaims(jwtToken, &newClaim, func(t *jwt.Token) (interface{}, error) { return []byte(SignedString), nil })
	assert.Nil(t, err, "ParseWithClaims return err")

	assert.True(t, tokenInfo.Valid, "tokenInfo is not valid")

	assert.Equal(t, int64(10000), newClaim.AdminClaim.Id)
	assert.Equal(t, "abcd", newClaim.AdminClaim.Name)
}
