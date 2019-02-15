package commands

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/tealeg/xlsx"
	"github.com/thedevsaddam/gojsonq"
	cli "gopkg.in/urfave/cli.v1"
)

// This command needs the following:
// 1. Text file with DDANAT IDs (one per line)
// 2. XLSX file that shows DDANAT IDs and Gene ID (i.e. anatomy_spatial_expression.xlsx)
// 3. JSON file containing the Circos gene coordinates
// 4. Output JSON file (default is output.json)

// The goal here is to take a DDANAT ID, match it to its gene ID, then create a JSON file
// containing its Circos coordinates as well as other relevant data.

// Spatial is the main data struct for json
type Spatial struct {
	Type       string       `json:"type"`
	ID         string       `json:"id"`
	Attributes *SpatialAttr `json:"attributes"`
}

// SpatialAttr is for the attributes field in the json
type SpatialAttr struct {
	BlockID     string `json:"block_id"`
	Start       int    `json:"start"`
	End         int    `json:"end"`
	Strand      string `json:"strand"`
	GeneName    string `json:"gene_name"`
	AnatomyTerm string `json:"anatomy_term"`
	Protein     bool   `json:"protein"`
	RNA         bool   `json:"rna"`
}

// Gene is the main data struct for genes json
type Gene struct {
	Type       string        `json:"type"`
	ID         string        `json:"id"`
	Attributes *GeneDataAttr `json:"attributes"`
}

// GeneDataAttr is for the attributes field in the genes json
type GeneDataAttr struct {
	SeqID   string `json:"seqid"`
	BlockID string `json:"block_id"`
	Source  string `json:"source"`
	Start   int    `json:"start"`
	End     int    `json:"end"`
	Strand  string `json:"strand"`
}

// ConvertXlsxToJSON reads an XLSX file and TXT file with DDANAT IDs,
// then writes the customized output into a JSON file.
func ConvertXlsxToJSON(c *cli.Context) error {
	// 1. Open the specified XLSX file
	x, err := xlsx.OpenFile(c.String("xlsx"))
	if err != nil {
		return fmt.Errorf("error reading xlsx file %s", err)
	}
	var genes []Spatial
	// 2. Open the specified TXT file and get slice of DDANAT IDs.
	m, err := readTxtAndGetIDs(c.String("txt"))
	if err != nil {
		return fmt.Errorf("error reading txt file %s", err)
	}
	// 3. Loop through the Excel sheets.
	for _, s := range x.Sheets {
		// 4. Loop through the rows of the Excel sheet.
		for _, r := range s.Rows {
			// 5. Loop through the slice of DDANAT IDs.
			for _, v := range m {
				// 6. If the DDANAT IDs from the Excel sheet and the slice match up, continue.
				if r.Cells[2].String() == v {
					// 7. Read the JSON content from specified file.
					jq := gojsonq.New().File(c.String("json"))
					// 8. Run query to find the matching gene ID for that DDANAT ID.
					q := jq.From("data").WhereEqual("id", r.Cells[0].String()).First()
					// 9. Get the JSON encoding from query response.
					b, err := json.Marshal(q)
					if err != nil {
						return fmt.Errorf("error converting to json %s", err)
					}
					g := Gene{}
					// 10. Parse the JSON data and store it in a pointer for our Gene struct.
					json.Unmarshal(b, &g)
					// 11. Append the data into the customized Spatial struct, getting the
					// desired data from both the JSON and the XLSX.
					genes = append(genes, Spatial{
						Type: "genes",
						ID:   g.ID,
						Attributes: &SpatialAttr{
							BlockID:     g.Attributes.BlockID,
							Start:       g.Attributes.Start,
							End:         g.Attributes.End,
							Strand:      g.Attributes.Strand,
							GeneName:    r.Cells[1].String(),
							AnatomyTerm: r.Cells[3].String(),
							Protein:     convertToBool(r.Cells[5].String()),
							RNA:         convertToBool(r.Cells[6].String()),
						},
					})
				}
			}
		}
	}
	// 12. Convert the genes slice to JSON.
	ga, err := json.Marshal(genes)
	if err != nil {
		return fmt.Errorf("unable to convert genes slice to json %s", err)
	}
	fmt.Println("total number of genes =", len(genes))
	// 13. Create JSON file to store our new data.
	jf, err := os.Create(c.String("output"))
	if err != nil {
		return fmt.Errorf("unable to create output json file %s", err)
	}
	defer jf.Close()
	// 14. Write the data to our JSON file.
	jf.WriteString("{\"data\": ")
	jf.Write(ga)
	jf.WriteString("}")
	return nil
}

func convertToBool(s string) bool {
	if s == "Yes" {
		return true
	}
	return false
}

// readTxtAndGetIDs reads a whole file into memory
// and returns a slice of its lines.
func readTxtAndGetIDs(p string) ([]string, error) {
	file, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "\t")
		lines = append(lines, parts[0])
	}
	return lines, scanner.Err()
}
