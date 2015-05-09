package main

import (
	"github.com/codegangsta/cli"
	"github.com/mgutz/ansi"
	"bufio"
	"fmt"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "emph"
	app.Version = "0.1"
	app.Usage = "The kind and sincere usage"
	app.Action = func(c *cli.Context) {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			colored_line := ansi.Color(line, "cyan+u")
			fmt.Println(colored_line)
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
	}

	app.Run(os.Args)
}


