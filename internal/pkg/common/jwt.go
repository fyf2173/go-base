package common

import "github.com/dgrijalva/jwt-go"

const SignedString = "microkepler"

const (
	AuthClientApp     = "miniapp"
	AuthClientConsole = "console"
)

type Claim struct {
	AuthClient string
	UserClaim
	AdminClaim
	jwt.StandardClaims
}

type UserClaim struct {
	UserId int64
	Openid string
}

type AdminClaim struct {
	Id   int64
	Name string
}
