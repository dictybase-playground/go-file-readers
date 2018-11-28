package main

import (
	"os"

	"github.com/dictybase-playground/go-file-reader/commands"
	"gopkg.in/codegangsta/cli.v2"
)

func main() {
	app := &cli.App{
		Name:    "go-file-readers",
		Usage:   "reads and converts data files",
		Version: "1.0.0",
		Commands: []cli.Command{
			{
				Name:   "gff3",
				Usage:  "reads GFF3 file and extracts data by given type",
				Action: commands.GFF3ReadAndWrite,
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "file",
						Usage: "gff3 file to convert",
						Value: "sample.gff3",
					},
					cli.StringFlag{
						Name:  "type",
						Usage: "type of data to extract (i.e. pseudogenes)",
						Value: "pseudogenes",
					},
					cli.StringFlag{
						Name:  "output",
						Usage: "output file",
						Value: "pseudogenes.gff3",
					},
				},
			},
			{
				Name:   "csv",
				Usage:  "reads csv file, converts it to JSON and writes it to new file",
				Action: commands.CSVtoJSON,
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "file",
						Usage: "csv file to convert",
						Value: "users.gff3",
					},
					cli.StringFlag{
						Name:  "output",
						Usage: "output file",
						Value: "users.json",
					},
				},
			},
		},
	}
	app.Run(os.Args)
}
