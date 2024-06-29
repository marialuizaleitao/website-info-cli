package main

import (
	"fmt"
	"github.com/marialuizaleitao/website-info-cli/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Starting point")

	application := app.Generate()
	if err := application.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
