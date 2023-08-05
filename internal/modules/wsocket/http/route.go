package http

import (
	"go-base/internal/modules/wsocket"
	"net/http"

	"github.com/zc2638/swag"
	"github.com/zc2638/swag/endpoint"
)

func SwagEndpoints() []*swag.Endpoint {
	wsocket.InitHub()
	return []*swag.Endpoint{
		endpoint.New(
			http.MethodGet, "/ws",
			endpoint.Handler(connectWs),
			endpoint.Summary("测试socket连接"),
			endpoint.Tags("demo"),
			endpoint.Description("Additional information on adding a pet to the store"),
			endpoint.Response(http.StatusOK, "Successfully added pet", endpoint.Schema(map[string]interface{}{})),
		),
	}
}
