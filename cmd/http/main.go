package main

import (
	"go-base/cmd"
	"go-base/cmd/http/api"
	"go-base/internal/config"
	"log"

	"github.com/fyf2173/ysdk-go/apisdk/ginplus"
	"github.com/fyf2173/ysdk-go/xdb"

	"github.com/spf13/cobra"
)

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "srv",
	Short: "HTTP对外接口服务",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		config.InitConfigData()
		if err := xdb.InitGorm(config.CfgData.Mode, config.CfgData.Mysql); err != nil {
			panic(err)
		}

		srv := ginplus.NewGinServer()
		srv.RegisterHandler(api.ConsoleHandler, api.AppHandler)
		log.Fatalln(srv.Start(":2222", func() {
			log.Println("do nothing before exit")
		}))
	},
}

func init() {
	cmd.RootCmd.AddCommand(httpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// httpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// httpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func main() {
	cmd.Execute()
}
