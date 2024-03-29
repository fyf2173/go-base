/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"go-base/internal/pkg/job"
	"log/slog"

	"github.com/fyf2173/ysdk-go/cmder"
	"github.com/fyf2173/ysdk-go/xctx"
	"github.com/fyf2173/ysdk-go/xlog"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type jobcmd struct {
	*cmder.BaseCmd
	job.Job
}

func newJobcmd() *jobcmd {
	cc := &jobcmd{}
	cc.BaseCmd = cmder.NewBaseCmd(&cobra.Command{
		Use:   "job",
		Short: "script脚本命令",
		Run: func(cmd *cobra.Command, args []string) {
			xlog.Init(viper.GetString("logger.level"), viper.GetBool("logger.add_source"))
			ctx := xctx.New()
			runner := job.NewRunner()
			runner.Add("test", func(ctx context.Context, args []string) error {
				fmt.Println("------------------", args, xctx.CtxId(ctx))
				return nil
			})

			if err := runner.ExecJob(ctx, cc.Job); err != nil {
				xlog.Error(ctx, fmt.Errorf("执行子命令[%s]失败，err=%s", cc.Job, err), slog.Any("subspec", cc.Args))
				return
			}
			xlog.Info(ctx, "---------------------- 任务[%s]执行完成 ----------------------", slog.Any("job", cc.Job))
		},
	})
	cc.BaseCmd.Cmd.Flags().StringVar(&cc.Job.Cmd, "subcmd", "", "需要执行的子命令")
	cc.BaseCmd.Cmd.Flags().StringArrayVar(&cc.Job.Args, "subspec", nil, "需要执行的子命令的参数")
	return cc
}
