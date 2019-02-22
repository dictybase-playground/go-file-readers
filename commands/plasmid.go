package commands

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	cli "gopkg.in/urfave/cli.v1"
)

// Order data structure
type Order struct {
	Date      string   `json:"date"`
	Purchaser string   `json:"purchaser"`
	Items     []string `json:"items"`
}

type Plasmid struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// OrderCSVtoJSON reads csv and tsv files, converts plasmid name to ID and writes output to json file
func OrderCSVtoJSON(c *cli.Context) error {
	o, err := os.Open(c.String("order-csv"))
	if err != nil {
		log.Fatal(err)
	}
	defer o.Close()
	var orders []Order
	var items []string
	var date, purchaser string
	reader := csv.NewReader(o)
	for {
		line, err := reader.Read()
		if err == io.EOF {
			// break the loop at end of file
			break
		}
		if err != nil {
			fmt.Errorf("error reading csv %s", err)
			// cannot return here due to csv reader throwing error
			// when csv lines have different number of fields
		}
		// if err, ok := err.(*csv.ParseError); ok && err.Err == csv.ErrFieldCount {
		// 	return nil
		// }
		for _, n := range line {
			if n[:2] == "20" {
				items = nil
				date = n
			} else if strings.Contains(n, "@") {
				purchaser = n
			} else if n[:1] == "p" {
				a, err := convertPlasmidIDToName(c.String("plasmid-tsv"), n)
				if err != nil {
					return fmt.Errorf("error converting plasmid id %s", err)
				}
				items = append(items, a)
			} else {
				items = append(items, n)
			}
		}
		orders = append(orders, Order{
			Date:      date,
			Purchaser: purchaser,
			Items:     items,
		})
	}
	j, err := json.Marshal(orders)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("total number of orders =", len(orders))
	jf, err := os.Create(c.String("output"))
	if err != nil {
		fmt.Println(err)
	}
	defer jf.Close()
	// Write the data to our JSON file.
	jf.WriteString("{\"data\": ")
	jf.Write(j)
	jf.WriteString("}")
	fmt.Println("JSON successfully written to", c.String("output"))
	return nil
}

func convertPlasmidIDToName(file string, name string) (string, error) {
	t, err := os.Open(file)
	if err != nil {
		return name, fmt.Errorf("error opening file %s", err)
	}
	defer t.Close()
	var n string
	scanner := bufio.NewScanner(t)
	for scanner.Scan() {
		line := scanner.Text()
		if err != nil {
			return name, fmt.Errorf("error reading line %s", err)
		}
		parts := strings.Split(line, "\t")
		if parts[1] == name {
			n = parts[0]
		}
	}
	return n, nil
}
