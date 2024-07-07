package main

import (
	"cli/app"
	"log"
	"os"
)

func main() {
	cli := app.Generate()

	if error := cli.Run(os.Args); error != nil {
		log.Fatal(error)
	}
}
