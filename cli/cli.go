package cli

import (
	"test/cli/tool/controllers"

	"github.com/urfave/cli/v2"
)

type CLIApp struct {
	App *cli.App
}

func Init() *CLIApp {
	return &CLIApp{
		App: &cli.App{
			Description: "Zoom CLI Toolkit",
			Flags: []cli.Flag{
				&cli.BoolFlag{Name: "single", Aliases: []string{"s"}},
				&cli.BoolFlag{Name: "batch", Aliases: []string{"b"}},
			},
			Commands: []*cli.Command{
				{
					Name:  "did",
					Usage: "commands to manage byoc DIDs",
					Subcommands: []*cli.Command{
						{
							Name:   "add",
							Usage:  "Add byoc DID to account",
							Action: controllers.AddDID,
						},
					},
				},
			},
		},
	}
}
