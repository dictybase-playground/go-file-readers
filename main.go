package main

import (
	"os"

	"github.com/dictybase-playground/go-file-readers/commands"
	cli "gopkg.in/urfave/cli.v1"
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
					Value: "pseudogene",
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
					Value: "users.csv",
				},
				cli.StringFlag{
					Name:  "output",
					Usage: "output file",
					Value: "users.json",
				},
			},
		},
		{
			Name:   "txt",
			Usage:  "reads txt file, converts it to JSON and writes it to new file",
			Action: commands.TxtToJSON,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "file",
					Usage: "txt file to convert",
				},
				cli.StringFlag{
					Name:  "chr",
					Usage: "the chromosome to extract",
				},
				cli.StringFlag{
					Name:  "output",
					Usage: "output file",
					Value: "output.json",
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
					Value: "users.csv",
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
		{
			Name:   "ddanat",
			Usage:  "converts data given DDANAT IDs into JSON file with circos coordinates and other data",
			Action: commands.ConvertXlsxToJSON,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "xlsx",
					Usage: "xlsx file to convert",
				},
				cli.StringFlag{
					Name:  "output",
					Usage: "output json file",
					Value: "output.json",
				},
				cli.StringFlag{
					Name:  "json",
					Usage: "genes json file to parse",
					Value: "data/genes.json",
				},
				cli.StringFlag{
					Name:  "txt",
					Usage: "txt file containing DDANAT IDs",
				},
			},
		},
		{
			Name:   "plasmid",
			Usage:  "reads csv and tsv files, converts plasmid name to ID and writes it to new json file",
			Action: commands.OrderCSVtoJSON,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "order-csv",
					Usage: "csv file containing orders data",
					Value: "data/stock_orders.csv",
				},
				cli.StringFlag{
					Name:  "plasmid-tsv",
					Usage: "tsv file containing plasmid data",
					Value: "data/plasmid_plasmid.tsv",
				},
				cli.StringFlag{
					Name:  "output",
					Usage: "output file",
					Value: "orders.json",
				},
			},
		},
	}
	app.Run(os.Args)
}
