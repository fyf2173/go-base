/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"go-base/internal/middleware"
	"go-base/internal/modules"
	"log"

	"github.com/fyf2173/ysdk-go/apisdk/ginplus"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

type httpcmd struct {
	*baseCmd
}

func consoleHandler(r *gin.Engine) {
	for _, v := range modules.ConsoleRoutes() {
		var mw []gin.HandlerFunc
		if middleware.CheckAuthIgnoreRegPath(middleware.AuthIgnorePath, v.Path) == false {
			mw = append(mw, middleware.ParseTokenMw())
		}
		r.Group("console").Handle(v.Method, v.Path, v.Handler.(gin.HandlerFunc)).Use(mw...)
	}
	return
}

func appHandler(r *gin.Engine) {
	for _, v := range modules.AppRoutes() {
		r.Group("app").Handle(v.Method, v.Path, v.Handler.(gin.HandlerFunc))
	}
	return
}

func newhttpcmd() *httpcmd {
	cc := &httpcmd{}
	cc.baseCmd = newBaseCmd(&cobra.Command{
		Use:   "srv",
		Short: "HTTP对外接口服务",
		Run: func(cmd *cobra.Command, args []string) {
			// config.InitConfigData()
			// if err := xdb.InitGorm(config.CfgData.Mode, config.CfgData.Mysql); err != nil {
			// 	panic(err)
			// }

			srv := ginplus.NewGinServer()
			srv.RegisterHandler(consoleHandler, appHandler, func(r *gin.Engine) {
				r.GET("/test", func(ctx *gin.Context) {
					ginplus.ExitSuccess(ctx, "ok")
				})
			})
			log.Fatalln(srv.Start(":2222", func() {
				log.Println("do nothing before exit")
			}))
		},
	})
	return cc
}
