package main

import (
	"log"
	"os"
	"test/cli/tool/cli"
)

func main() {
	cliApp := cli.Init()
	if err := cliApp.App.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
