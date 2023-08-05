package modules

import (
	example "go-base/internal/modules/example/http"
	wsocket "go-base/internal/modules/wsocket/http"

	"github.com/gin-gonic/gin"
	"github.com/zc2638/swag"
)

var Rg = []RouterGroup{
	{Group: "console", Endpoints: SwagEndpoints(), Mw: nil},
	{Group: "app", Endpoints: SwagEndpoints(), Mw: nil},
}

type RouterGroup struct {
	Group     string
	Endpoints []*swag.Endpoint
	Mw        []gin.HandlerFunc
}

func SwagEndpoints() []*swag.Endpoint {
	var endpoints []*swag.Endpoint
	endpoints = append(endpoints, example.SwagEndpoints()...)
	endpoints = append(endpoints, wsocket.SwagEndpoints()...)
	return endpoints
}
