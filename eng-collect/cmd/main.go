package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags) // code position + date-time
}

func main() {
	parseTextCommand := &cli.Command{
		Name:    "parse",
		Aliases: []string{"load", "analysis"},
		Usage:   "parse text to retrieve and insert new words",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "url",
				Aliases:  []string{"link", "s", "youtube"},
				Usage:    "youtube video link to parse",
				Required: false,
			},
			&cli.StringFlag{
				Name:     "text",
				Aliases:  []string{"t"},
				Usage:    "text(direct) to parse",
				Required: false,
			},
			&cli.StringFlag{
				Name:     "file",
				Aliases:  []string{"f"},
				Usage:    "path of local text file to parse",
				Required: false,
			},
		},
		Action: func(cliCtx *cli.Context) error {
			url := cliCtx.String("url")
			text := cliCtx.String("text")
			file := cliCtx.String("file")

			if url == "" && text == "" && file == "" {
				return ErrNothingToParse
			}

			// parse directly from text
			if text != "" {
				fmt.Println(text)
				return nil
			}

			// parse from local file, input is path to that file
			if file != "" {
				fmt.Println(file)
				return nil
			}

			// last option: parse from youtube URL
			return nil
		},
	}
	app := &cli.App{
		Name:  "engcollect",
		Usage: "collect and store english words from text or URL",
		Commands: []*cli.Command{
			parseTextCommand,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

var ErrNothingToParse = errors.New("nothing to parse, use help to see avaliable choices for parsing")
