package modules

import (
	"go-base/internal/middleware"
	"go-base/internal/modules/example/api"
	example "go-base/internal/modules/example/http"
	wsocket "go-base/internal/modules/wsocket/http"

	"github.com/gin-gonic/gin"
	"github.com/zc2638/swag"
)

var consoleIgnoreAuthPaths = []string{
	"console/example/test-ignore-auth",
	"console/example/test-get-token",
}

var Rg = []RouterGroup{
	{Group: "console", Endpoints: SwagEndpoints(), Mw: []gin.HandlerFunc{middleware.CommonTokenMw(consoleIgnoreAuthPaths...)}},
	{Group: "service", Endpoints: ServiceEndpoints(), Mw: nil},
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

func ServiceEndpoints() []*swag.Endpoint {
	var endpoints []*swag.Endpoint
	endpoints = append(endpoints, api.SwagEndpoints()...)
	return endpoints
}
