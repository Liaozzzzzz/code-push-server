package cmd

import (
	"context"

	"github.com/liaozzzzzz/code-push-server/internal/bootstrap"
	"github.com/urfave/cli/v3"
)

func Start() *cli.Command {
	return &cli.Command{
		Name:  "start",
		Usage: "Start server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				Usage:       "Config directory",
				DefaultText: "config",
				Value:       "config",
			},
			&cli.StringFlag{
				Name:        "env",
				Aliases:     []string{"e"},
				Usage:       "Runtime environment",
				DefaultText: "dev",
				Value:       "dev",
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			configDir := c.String("config")
			env := c.String("env")

			err := bootstrap.Run(ctx, bootstrap.RunConfig{
				ConfigDir: configDir,
				Env:       env,
			})
			if err != nil {
				panic(err)
			}
			return nil
		},
	}
}
