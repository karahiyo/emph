package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "emph"
	app.Usage = "The kind and sincere usage"
	app.Action = func(c *cli.Context) {
		println("he, he, hello?...")
	}

	app.Run(os.Args)
}


