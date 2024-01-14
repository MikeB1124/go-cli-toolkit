package controllers

import (
	"fmt"
	"test/cli/tool/csvUtil"

	"github.com/urfave/cli/v2"
)

func AddDID(cCtx *cli.Context) error {
	single := cCtx.Bool("s")
	batch := cCtx.Bool("b")

	if single {
		did := cCtx.Args().First()
		fmt.Println("added single DID:", did)
	} else if batch {
		filePath := cCtx.Args().First()
		list, err := csvUtil.ReadCSVFile(filePath)
		if err != nil {
			fmt.Println("Error:", err)
		}

		for _, recordDict := range list {
			fmt.Println(recordDict["DID"])
		}
	} else {
		fmt.Println("Please provide either --single or --batch flag")
	}

	return nil
}
