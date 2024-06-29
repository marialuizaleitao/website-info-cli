package main

import (
	"github.com/marialuizaleitao/website-info-cli/app"
	"log"
	"os"
)

func main() {
	application := app.Generate()
	if err := application.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
