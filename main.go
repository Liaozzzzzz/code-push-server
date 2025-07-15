package main

import (
	"context"
	"log"
	"os"

	"github.com/liaozzzzzz/code-push-server/cmd"
	_ "github.com/liaozzzzzz/code-push-server/docs"
	"github.com/urfave/cli/v3"
)

// @title           code-push-server
// @version         0.0.1
// @description     code-push-server
// @host      		localhost:8042
// @BasePath  		/api/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	app := cli.Command{
		Name:  "code-push-server",
		Usage: "code push server",
		Commands: []*cli.Command{
			cmd.Start(),
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}

}
