package http

import (
	"net/http"

	"github.com/zc2638/swag"
	"github.com/zc2638/swag/endpoint"
)

func SwagEndpoints() []*swag.Endpoint {
	return []*swag.Endpoint{
		endpoint.New(
			http.MethodGet, "/example/test",
			endpoint.Handler(Test),
			endpoint.Summary("Add a new test api"),
			endpoint.Tags("demo"),
			endpoint.Description("Additional information on adding a pet to the store"),
			//endpoint.Body(nil, "Pet object that needs to be added to the store", true),
			endpoint.Response(http.StatusOK, "Successfully added pet", endpoint.Schema(map[string]interface{}{})),
		),
	}
}
