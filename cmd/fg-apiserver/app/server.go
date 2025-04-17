package app

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewFastGOCommand() *cobra.Command {
	cmd := &cobra.Command{
		// 指定命令名称 -> 将出现在帮助信息中
		Use: "fg-apiserver",
		// 命令描述
		Short: "A very lightweight full go project",
		Long:  "A very lightweight full go project, designed to help beginners quickly learn Go project development.",
		// 设置不打印帮助信息，以快速确认错误信息
		SilenceUsage: true,
		// 指定调用 cmd.Execute() 时，执行的 Run 函数
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Hello FastGo")
			return nil
		},
		// 设置命令运行时的参数检查，不需要指定命令行参数。eg：./fg-apiserver param1 param2
		Args: cobra.NoArgs,
	}
	return cmd
}
