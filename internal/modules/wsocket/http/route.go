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
			http.MethodGet, "/conn",
			endpoint.Handler(connectWs),
			endpoint.Tags("socket"),
			endpoint.Summary("测试socket连接"),
			endpoint.Description("测试socket连接"),
			endpoint.Response(http.StatusOK, "success", endpoint.Schema(map[string]interface{}{})),
		),
	}
}
