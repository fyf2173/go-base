package api

import (
	"net/http"

	"github.com/zc2638/swag"
	"github.com/zc2638/swag/endpoint"
)

func SwagEndpoints() []*swag.Endpoint {
	return []*swag.Endpoint{
		endpoint.New(
			http.MethodGet, "/example/test",
			endpoint.Handler(test),
			endpoint.Tags("service api"),
			endpoint.Summary("内部服务无鉴权接口"),
			endpoint.Description("内部服务无鉴权接口"),
			// endpoint.Body(nil, "Pet object that needs to be added to the store", true),
			endpoint.Response(http.StatusOK, "success", endpoint.Schema(map[string]interface{}{})),
		),
	}
}
