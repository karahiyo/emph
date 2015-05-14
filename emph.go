package main

import (
	"github.com/codegangsta/cli"
	"github.com/mgutz/ansi"
	"encoding/csv"
	"regexp"
	"bufio"
	"fmt"
	"os"
	"io"
)

const (
	REPLACE_ALL = -1
)

type Item struct {
	RegexpObj	*regexp.Regexp
	ColorCode	string
}

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
		fp, err := os.Open(c.String("conf"))
		if err != nil {
			panic(err)
		}
		defer fp.Close()

		var conf []Item
		reader := csv.NewReader(fp)
		reader.Comma = '\t'
		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				panic(err)
			}
			re, err := regexp.Compile(record[0])
			if err != nil {
				continue
			}
			cc := ansi.ColorCode(record[1])
			if cc == "" {
				continue
			}
			i := Item { RegexpObj: re, ColorCode: cc }
			conf = append(conf, i)
		}

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			for _, i := range conf {
				r := i.RegexpObj
				cc := i.ColorCode
				line = r.ReplaceAllStringFunc(line,
				func(m string) string {
					return cc + m + ansi.Reset
				})
			}
			fmt.Println(line)
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
	}

	app.Run(os.Args)
}


