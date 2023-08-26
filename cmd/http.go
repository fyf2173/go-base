/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/fyf2173/ysdk-go/cmder"
	"go-base/internal/modules"
	"go-base/internal/orm"
	"log"
	"log/slog"
	"os"

	"github.com/fyf2173/ysdk-go/apisdk/ginplus"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/zc2638/swag"
	"github.com/zc2638/swag/option"
)

type httpcmd struct {
	*cmder.BaseCmd
}

func swaggerDocHandler(rgs ...modules.RouterGroup) func(r *gin.Engine) {
	doc := swag.New(
		option.Title("demo接口文档"),
		option.Version(Version),
		option.Host("http://localhost:2222"),
		option.BasePath("/godemo/v1"),
	)
	return func(r *gin.Engine) {
		router := r.Group(doc.BasePath)
		for _, rg := range rgs {
			g := router.Group(rg.Group, rg.Mw...)
			for _, ep := range rg.Endpoints {
				g.Handle(ep.Method, ep.Path, ep.Handler.(func(*gin.Context)))
			}

			doc.WithGroup(rg.Group).AddEndpoint(rg.Endpoints...)
		}
		r.GET("/swagger/json", gin.WrapH(doc.Handler()))
		r.GET("/swagger/ui/*any", gin.WrapH(swag.UIHandler("/swagger/ui", "/swagger/json", true)))
	}
}

func newhttpcmd() *httpcmd {
	cc := &httpcmd{}
	cc.BaseCmd = cmder.NewBaseCmd(&cobra.Command{
		Use:   "srv",
		Short: "HTTP对外接口服务",
		Run: func(cmd *cobra.Command, args []string) {
			slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr, nil)))
			orm.InitConn()
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
