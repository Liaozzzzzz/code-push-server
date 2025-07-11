package bootstrap

import (
	"context"
	"fmt"
	"strings"

	"github.com/liaozzzzzz/code-push-server/internal/config"
)

type RunConfig struct {
	ConfigDir string
	Env       string
}

func Run(ctx context.Context, runCfg RunConfig) error {

	config.MustLoad(runCfg.ConfigDir, strings.Split(runCfg.Env, ",")...)

	fmt.Println(config.C)
	return nil
}
