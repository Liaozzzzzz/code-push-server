package bootstrap

import (
	"context"
	"fmt"

	"github.com/liaozzzzzz/code-push-server/internal/config"
	"github.com/liaozzzzzz/code-push-server/internal/server"
)

type RunConfig struct {
	ConfigDir string
	Env       string
}

func Run(ctx context.Context, runCfg RunConfig) error {

	// 加载配置
	config.MustLoad(runCfg.ConfigDir, runCfg.Env)

	// 这里可以添加更多的启动逻辑
	if config.C.General.Debug {
		fmt.Println("🐛 调试模式已启用")
	}

	// 创建并启动HTTP服务器
	srv := server.NewServer()
	return srv.Start()
}
