package app

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	// defaultHomeDir 定义放置 fastgo 服务配置的默认目录
	defaultHomeDir = ".fastgo"

	// defaultConfigName 指定 fastgo 服务的默认配置文件名
	defaultConfigName = "fg-apiserver.yaml"
)

// onInitialze 设置需要读取的配置文件名、环境变量，并将其内容读取到 viper 中
func onInitialize() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		// 使用默认配置文件路径和名称
		for _, dir := range searchDirs() {
			// dir 目录加入到配置文件的搜索路径
			viper.AddConfigPath(dir)
		}

		// 设置读取配置文件的格式为 YAML
		viper.SetConfigType("yaml")

		// 配置文件名称（不需要文件拓展名）
		viper.SetConfigName(defaultConfigName)
	}

	// 读取环境变量并设置前缀
	setupEnviromentVariables()

	// 读取配置文件，如果指定了配置文件名，则使用指定的，否则就在注册的路径中搜索
	_ = viper.ReadInConfig()
}

// setupEnviromentVariables 配置环境变量规则
func setupEnviromentVariables() {
	// 允许 viper 自定匹配环境变量
	viper.AutomaticEnv()
	// 设置环境变量前缀
	viper.SetEnvPrefix("FASTGO")
	// 替换环境变量 key 中的分隔符 '.' 和 '-' 为 '_'
	replacer := strings.NewReplacer(".", "_", "-", "_")
	viper.SetEnvKeyReplacer(replacer)
}

// searchDirs 返回默认配置文件搜索目录
func searchDirs() []string {
	// 获取用户主目录
	homeDir, err := os.UserHomeDir()
	// 如果获取用户主目失败，则打印错误信息并退出程序
	cobra.CheckErr(err)
	return []string{filepath.Join(homeDir, defaultHomeDir), "."}
}

// filePath 获取默认配置文件的完整路径
func filePath() string {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)
	return filepath.Join(home, defaultHomeDir, defaultConfigName)
}
