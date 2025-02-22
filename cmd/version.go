/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/fyf2173/ysdk-go/cmder"

	"github.com/spf13/cobra"
)

var (
	AppName string
	Branch  string
	Commit  string
	Author  string
	Date    string
	Version string
)

func String() string {
	return fmt.Sprintf("AppName: %s\nVersion: %s\nBranch: %s\nCommit: %s\nAuthor: %s\nDate: %s", AppName, Version, Branch, Commit, Author, Date)
}

type versioncmd struct {
	*cmder.BaseCmd
}

func newVersioncmd() *versioncmd {
	cc := &versioncmd{}
	cc.BaseCmd = cmder.NewBaseCmd(&cobra.Command{
		Use:   "version",
		Short: "查看当前编译版本号",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(String())
		},
	})
	return cc
}
