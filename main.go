package main

import (
	"ipf/app"
	"log"
	"os"
)

func main() {
	app := app.Gen()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
