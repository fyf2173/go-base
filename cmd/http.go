/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"go-base/internal/middleware"
	"go-base/internal/modules"
	"go-base/internal/orm"

	"github.com/fyf2173/ysdk-go/cmder"
	"github.com/fyf2173/ysdk-go/xctx"
	"github.com/fyf2173/ysdk-go/xlog"

	"net/http"
	_ "net/http/pprof"

	"github.com/fyf2173/ysdk-go/apisdk/ginplus"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		option.Host("http://0.0.0.0:2222"),
		option.BasePath("/godemo/v1"),
	)
	return func(r *gin.Engine) {
		router := r.Group(doc.BasePath)
		router.Use(middleware.Access(), gin.Recovery())
		for _, rg := range rgs {
			g := router.Group(rg.Group, rg.Mw...)
			for _, ep := range rg.Endpoints {
				g.Handle(ep.Method, ep.Path, ep.Handler.(func(*gin.Context)))
			}

			doc.WithGroup("/" + rg.Group).AddEndpoint(rg.Endpoints...)
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
			xlog.Init(viper.GetString("logger.level"), viper.GetBool("logger.add_source"))
			go func() { http.ListenAndServe(":6060", nil) }()
			orm.InitConn()
			srv := ginplus.NewGinServer()
			srv.RegisterHandler(
				swaggerDocHandler(modules.Rg...),
			)
			cc := xctx.New()
			xlog.Error(cc, srv.Start(":2222", func() {
				xlog.Info(cc, "do nothing before exit")
			}))
		},
	})
	return cc
}
