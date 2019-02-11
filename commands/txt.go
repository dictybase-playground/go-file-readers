package commands

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	cli "gopkg.in/urfave/cli.v1"
)

// Chromosome is the main data struct for json
type Chromosome struct {
	Type       string          `json:"type"`
	Attributes *ChromosomeAttr `json:"attributes"`
}

// ChromosomeAttr is for the attributes field in the json
type ChromosomeAttr struct {
	Chromosome string `json:"chromosome"`
	Start      string `json:"start"`
	End        string `json:"end"`
	Value      string `json:"value"`
}

// TxtToJSON reads txt file, converts it to JSON and writes it to new file
func TxtToJSON(c *cli.Context) error {
	file, err := os.Open(c.String("file"))
	if err != nil {
		return fmt.Errorf("error opening file %s", err)
	}
	defer file.Close()
	var data []Chromosome
	chrMap := map[string]string{
		"DD1": "DDB0232428",
		"DD2": "DDB0232429",
		"DD3": "DDB0232430",
		"DD4": "DDB0232431",
		"DD5": "DDB0232432",
		"DD6": "DDB0232433",
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if err != nil {
			return fmt.Errorf("error reading line %s", err)
		}
		// split line by tab
		parts := strings.Split(line, "\t")
		data = append(data, Chromosome{
			Type: "chromosomes",
			Attributes: &ChromosomeAttr{
				Chromosome: chrMap[parts[0]],
				Start:      parts[1],
				End:        parts[2],
				Value:      parts[3],
			},
		})
	}
	j, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error converting to json %s", err)
	}
	fmt.Printf("total number of lines converted = %d", len(data))
	f, err := os.Create(c.String("output"))
	if err != nil {
		return fmt.Errorf("error creating output file %s", err)
	}
	defer f.Close()
	f.Write(j)
	fmt.Println("JSON successfully written to", c.String("output"))
	return nil
}
