package app

import (
	"github.com/MortalSC/fastgo/cmd/fg-apiserver/app/options"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// 配置文件路径
var configFile string

func NewFastGOCommand() *cobra.Command {
	// 创建默认的命令行选项
	opts := options.NewServerOptions()

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
			return run(opts)

		},
		// 设置命令运行时的参数检查，不需要指定命令行参数。eg：./fg-apiserver param1 param2
		Args: cobra.NoArgs,
	}

	// 初始化配置函数，在每个命令运行时调用
	cobra.OnInitialize(onInitialize)

	cmd.PersistentFlags().StringVarP(&configFile, "config", "c", filePath(), "Path to the fg-apiserver configuration file.")

	return cmd
}

func run(opts *options.ServerOptions) error {
	// 使用 viper 将配置解析到 opts 中
	if err := viper.Unmarshal(opts); err != nil {
		return err
	}

	// 对命令行选项值进行校验
	if err := opts.Validate(); err != nil {
		return err
	}

	// 获取应用配置
	cfg, err := opts.Config()
	if err != nil {
		return err
	}

	// 创建服务器实例
	server, err := cfg.NewServer()
	if err != nil {
		return err
	}

	return server.Run()
}
