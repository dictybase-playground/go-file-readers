package main

import (
	"os"

	"github.com/dictybase-playground/go-file-reader/commands"
	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "go-file-readers"
	app.Usage = "reads and converts data files"
	app.Version = "1.0.0"
	app.Commands = []cli.Command{
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
		{
			Name:   "arango",
			Usage:  "reads csv file, converts to user data structure and stores in arangodb",
			Action: commands.StoreCSVinDB,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "file",
					Usage: "csv file to convert",
					Value: "users.gff3",
				},
				cli.StringFlag{
					Name:  "user",
					Usage: "arangodb username",
				},
				cli.StringFlag{
					Name:  "pw",
					Usage: "arangodb password",
				},
				cli.StringFlag{
					Name:  "host",
					Usage: "arangodb host",
					Value: "localhost",
				},
				cli.StringFlag{
					Name:  "port",
					Usage: "arangodb port",
					Value: "8529",
				},
				cli.StringFlag{
					Name:  "db",
					Usage: "arangodb database",
					Value: "colleagues",
				},
				cli.StringFlag{
					Name:  "collection",
					Usage: "arangodb collection",
					Value: "users",
				},
			},
		},
	}
	app.Run(os.Args)
}
