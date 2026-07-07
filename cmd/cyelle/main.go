package main

import (
	"os"

	"cyelle/internal/cli"
)

func main() {
	app := cli.New()

	os.Exit(app.Run(os.Args))
}