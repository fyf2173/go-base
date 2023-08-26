/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"github.com/fyf2173/ysdk-go/cmder"
	"go-base/internal/pkg/job"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type jobcmd struct {
	*cmder.BaseCmd
	subcmd  string
	subspec string
}

func newJobcmd() *jobcmd {
	cc := &jobcmd{}
	cc.BaseCmd = cmder.NewBaseCmd(&cobra.Command{
		Use:   "job",
		Short: "script脚本命令",
		Run: func(cmd *cobra.Command, args []string) {
			logrus.SetFormatter(&logrus.JSONFormatter{})
			runner := job.NewRunner()
			runner.Add("test", func(ctx context.Context) error {
				fmt.Println("------------------", cc.subspec)
				return nil
			})

			if err := runner.Exec(cc.subcmd); err != nil {
				logrus.Printf("执行子命令[%s]失败，err=%s", cc.subcmd, err)
				return
			}
			logrus.Printf("---------------------- 任务[%s]执行完成 ----------------------", cc.subcmd)
		},
	})
	cc.BaseCmd.Cmd.Flags().StringVar(&cc.subcmd, "subcmd", "", "需要执行的子命令")
	cc.BaseCmd.Cmd.Flags().StringVar(&cc.subspec, "subspec", "", "需要执行的子命令的参数")
	return cc
}
