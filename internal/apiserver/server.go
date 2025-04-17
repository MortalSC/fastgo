package apiserver

import (
	"fmt"

	genericoptions "github.com/MortalSC/fastgo/pkg/options"
)

// Config 配置结构体，用于存储应用相关配置
type Config struct {
	MySQLOptions *genericoptions.MySQLOptions
}

// Server 定义一个服务器结构体类型
type Server struct {
	cfg *Config
}

// NewServer 根据配置创建一个服务器实例
func (cfg *Config) NewServer() (*Server, error) {
	return &Server{
		cfg: cfg,
	}, nil
}

func (s *Server) Run() error {
	fmt.Printf("Read Mysql host from config: %s\n", s.cfg.MySQLOptions.Addr)

	// 使用 select 防止进程退出
	select {}
}
