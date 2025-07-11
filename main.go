package main

import (
	"context"
	"log"
	"os"

	"github.com/liaozzzzzz/code-push-server/cmd"
	"github.com/urfave/cli/v3"
)

func main() {
	app := cli.Command{
		Name:  "code-push",
		Usage: "code push",
		Commands: []*cli.Command{
			cmd.Start(),
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}

}
