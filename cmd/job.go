/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"go-base/internal/pkg/job"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type jobcmd struct {
	*baseCmd
	subcmd string
}

func newJobcmd() *jobcmd {
	cc := &jobcmd{}
	cc.baseCmd = newBaseCmd(&cobra.Command{
		Use:   "job",
		Short: "script脚本命令",
		Run: func(cmd *cobra.Command, args []string) {
			logrus.SetFormatter(&logrus.JSONFormatter{})
			runner := job.NewRunner()
			runner.Add("test", func(ctx context.Context) error {
				return nil
			})

			if err := runner.Exec(cc.subcmd); err != nil {
				logrus.Printf("执行子命令[%s]失败，err=%s", cc.subcmd, err)
				return
			}
			logrus.Printf("---------------------- 任务[%s]执行完成 ----------------------", cc.subcmd)
		},
	})
	cc.cmd.Flags().StringVar(&cc.subcmd, "subcmd", "", "需要执行的子命令")
	return cc
}
