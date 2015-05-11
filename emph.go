package main

import (
	"github.com/codegangsta/cli"
	"github.com/mgutz/ansi"
	"encoding/csv"
	"strings"
	"bufio"
	"fmt"
	"os"
	"io"
)

const (
	REPLACE_ALL = -1
)

func main() {

	app := cli.NewApp()
	app.Name = "emph"
	app.Version = "0.1"
	app.Usage = "The kind and sincere usage"
	app.Flags = []cli.Flag{
		cli.StringFlag {
			Name:"conf,c",
			Usage: "supecify hilight setting",
		},
	}
	app.Action = func(c *cli.Context) {
		var fp *os.File
		if len(os.Args) < 2 {
			fp = os.Stdin
		} else {
			var err error
			fp, err = os.Open(c.String("conf"))
			if err != nil {
				panic(err)
			}
			defer fp.Close()
		}

		conf := make(map[string]string)
		reader := csv.NewReader(fp)
		reader.Comma = '\t'
		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				panic(err)
			}
			conf[record[0]] = record[1]
		}
		keys := make([]string, len(conf))
		for k, _ := range conf {
			keys = append(keys, k)
		}

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			for _, k := range keys {
				if strings.Contains(line, k) {
					line = strings.Replace(line, k, ansi.Color(k, conf[k]), REPLACE_ALL)
				}
			}
			fmt.Println(line)
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
	}

	app.Run(os.Args)
}


