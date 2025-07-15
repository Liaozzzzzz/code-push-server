package bootstrap

import (
	"context"
	"fmt"
	"log"

	"github.com/liaozzzzzz/code-push-server/internal/config"
	"github.com/liaozzzzzz/code-push-server/internal/utils/crypto"
)

type RunConfig struct {
	ConfigDir string
	Env       string
}

func Run(ctx context.Context, runCfg RunConfig) error {

	// 加载配置
	config.MustLoad(runCfg.ConfigDir, runCfg.Env)

	// 初始化加密工具
	if err := crypto.InitCrypto(); err != nil {
		log.Fatal("Failed to initialize crypto:", err)
	}

	// 这里可以添加更多的启动逻辑
	if config.C.General.Debug {
		fmt.Println("🐛 调试模式已启用")
		fmt.Println("🔐 加密工具已初始化")
	}

	// 创建并启动HTTP服务器
	srv := NewServer()
	return srv.StartServer()
}
