package api

import (
	example "go-base/internal/modules/example/http"

	"github.com/fyf2173/ysdk-go/apisdk"
)

func ConsoleRoutes() []*apisdk.Route {
	var routes []*apisdk.Route
	routes = append(routes, example.Entries()...)
	return routes
}

func AppRoutes() []*apisdk.Route {
	var routes []*apisdk.Route
	return routes
}
