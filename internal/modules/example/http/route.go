package http

import (
	"github.com/fyf2173/ysdk-go/apisdk"
	"github.com/fyf2173/ysdk-go/apisdk/ginplus"
)

func Entries() []*apisdk.Route {
	return []*apisdk.Route{
		ginplus.Get("/v1/example/test", Test, ""),
	}
}
