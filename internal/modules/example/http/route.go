package http

import (
	"net/http"

	"github.com/zc2638/swag"
	"github.com/zc2638/swag/endpoint"
)

func SwagEndpoints() []*swag.Endpoint {
	return []*swag.Endpoint{
		endpoint.New(
			http.MethodGet, "/example/test-ignore-auth",
			endpoint.Handler(TestIgnoreAuth),
			endpoint.Tags("console api"),
			endpoint.Summary("test-ignore-auth-api"),
			endpoint.Description("后台服务需鉴权接口"),
			// endpoint.Body(nil, "Pet object that needs to be added to the store", true),
			endpoint.Response(http.StatusOK, "success", endpoint.Schema(map[string]interface{}{})),
		),
		endpoint.New(
			http.MethodGet, "/example/test-auth",
			endpoint.Handler(TestWithAuth),
			endpoint.Tags("console api"),
			endpoint.Summary("test-auth-api"),
			endpoint.Description("后台服务需鉴权接口"),
			// endpoint.Body(nil, "Pet object that needs to be added to the store", true),
			endpoint.Response(http.StatusOK, "success", endpoint.Schema(map[string]interface{}{})),
		),
		endpoint.New(
			http.MethodGet, "/example/test-get-token",
			endpoint.Handler(TestGetToken),
			endpoint.Tags("console api"),
			endpoint.Summary("test-get-token"),
			endpoint.Description("后台服务需鉴权接口"),
			// endpoint.Body(nil, "Pet object that needs to be added to the store", true),
			endpoint.Response(http.StatusOK, "success", endpoint.Schema(map[string]interface{}{})),
		),
	}
}
