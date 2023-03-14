/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"go-base/internal/modules"
	"log"
	"strings"

	"github.com/fyf2173/ysdk-go/apisdk/ginplus"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/zc2638/swag"
	"github.com/zc2638/swag/option"
)

type httpcmd struct {
	*baseCmd
}

func swaggerDocHandler(rgs ...modules.RouterGroup) func(r *gin.Engine) {
	api := swag.New(
		option.Title("demo接口文档"),
		option.Version(Version),
		option.Host("http://localhost:2222"),
		option.BasePath("/godemo/v1"),
	)

	var tmpRoutes []*swag.Endpoint
	for _, rg := range rgs {
		for _, ep := range rg.Endpoints {
			ep.Path = fmt.Sprintf("/%s/%s", rg.Group, ep.Path)
		}
		tmpRoutes = append(tmpRoutes, rg.Endpoints...)
	}
	api.AddEndpoint(tmpRoutes...)

	return func(r *gin.Engine) {
		api.Walk(func(path string, e *swag.Endpoint) {
			h := e.Handler.(func(ctx *gin.Context))
			path = strings.TrimPrefix(swag.ColonPath(path), api.BasePath)
			r.Group(api.BasePath).Handle(e.Method, path, h)
		})

		r.GET("/swagger/json", gin.WrapH(api.Handler()))
		r.GET("/swagger/ui/*any", gin.WrapH(swag.UIHandler("/swagger/ui", "/swagger/json", true)))
	}
}

func newhttpcmd() *httpcmd {
	cc := &httpcmd{}
	cc.baseCmd = newBaseCmd(&cobra.Command{
		Use:   "srv",
		Short: "HTTP对外接口服务",
		Run: func(cmd *cobra.Command, args []string) {
			srv := ginplus.NewGinServer()
			srv.RegisterHandler(
				swaggerDocHandler(modules.Rg...),
			)
			log.Fatalln(srv.Start(":2222", func() {
				log.Println("do nothing before exit")
			}))
		},
	})
	return cc
}