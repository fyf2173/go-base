package main

import (
	"context"
	"go-base/cmd"
	"go-base/internal/config"
	"go-base/internal/pkg/job"

	"github.com/fyf2173/ysdk-go/util"
	"github.com/fyf2173/ysdk-go/xdb"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var subCmdFlags = "subcmd"

// jobCmd represents the http command
var jobCmd = &cobra.Command{
	Use:   "run",
	Short: "script脚本命令",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.SetFormatter(&logrus.JSONFormatter{})

		config.InitConfigData()
		util.Assert(xdb.InitGorm(config.CfgData.Mode, config.CfgData.Mysql))

		runner := job.NewRunner()
		runner.Add("test", func(ctx context.Context) error {
			return nil
		})

		cmdHandler, _ := cmd.Flags().GetString(subCmdFlags)
		if err := runner.Exec(cmdHandler); err != nil {
			logrus.Printf("执行子命令[%s]失败，err=%s", cmdHandler, err)
			return
		}
		logrus.Printf("---------------------- 任务[%s]执行完成 ----------------------", cmdHandler)
	},
}

func init() {
	jobCmd.Flags().String(subCmdFlags, "", "需要执行的子命令")

	cmd.RootCmd.AddCommand(jobCmd)

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
