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
	// Open the specified XLSX file
	x, err := xlsx.OpenFile(c.String("xlsx"))
	if err != nil {
		return fmt.Errorf("error reading xlsx file %s", err)
	}
	var genes []Spatial
	// Open the specified TXT file and get slice of DDANAT IDs.
	m, err := readTxtAndGetIDs(c.String("txt"))
	if err != nil {
		return fmt.Errorf("error reading txt file %s", err)
	}
	// Loop through the Excel sheets.
	for _, s := range x.Sheets {
		// Loop through the rows of the Excel sheet.
		for _, r := range s.Rows {
			// Loop through the slice of DDANAT IDs.
			for _, v := range m {
				// If the DDANAT IDs from the Excel sheet and the slice match up, continue.
				if r.Cells[2].String() == v {
					g, err := readAndParseJSON(c.String("json"), r.Cells[0].String())
					if err != nil {
						return fmt.Errorf("error reading and parsing json %s", err)
					}
					// Append the data into the customized Spatial struct, getting the
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
	// Convert the genes slice to JSON.
	ga, err := json.Marshal(genes)
	if err != nil {
		return fmt.Errorf("unable to convert genes slice to json %s", err)
	}
	fmt.Println("total number of genes =", len(genes))
	// Create JSON file to store our new data.
	jf, err := os.Create(c.String("output"))
	if err != nil {
		return fmt.Errorf("unable to create output json file %s", err)
	}
	defer jf.Close()
	// Write the data to our JSON file.
	jf.WriteString("{\"data\": ")
	jf.Write(ga)
	jf.WriteString("}")
	return nil
}

// readAndParseJSON reads from a specified JSON file,
// finds the matching gene ID for a DDANAT ID,
// and parses this data into our Gene struct.
func readAndParseJSON(j, id string) (Gene, error) {
	// Read the JSON content from specified file.
	jq := gojsonq.New().File(j)
	// Run query to find the matching gene ID for that DDANAT ID.
	q := jq.From("data").WhereEqual("id", id).First()
	// Get the JSON encoding from query response.
	b, err := json.Marshal(q)
	if err != nil {
		return Gene{}, fmt.Errorf("error converting to json %s", err)
	}
	g := Gene{}
	// Parse the JSON data and store it in a pointer for our Gene struct.
	json.Unmarshal(b, &g)
	return g, nil
}

// convertToBool simply converts the Yes/No cells to a boolean
func convertToBool(s string) bool {
	if s == "Yes" {
		return true
	}
	return false
}

// readTxtAndGetIDs reads a whole file into memory
// and returns a slice of its lines (the DDANAT IDs).
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
