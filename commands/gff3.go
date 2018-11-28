package commands

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"gopkg.in/urfave/cli.v1"
)

// GFF3ReadAndWrite reads GFF3 file and writes new one based on type
func GFF3ReadAndWrite(c *cli.Context) {
	// open input file
	file, err := os.Open(c.String("file"))
	if err != nil {
		log.Fatal(err) // equivalent to Print() followed by os.Exit(1)
	}
	// close file when done
	defer file.Close()

	// create output file
	output, err := os.Create(c.String("output"))
	if err != nil {
		log.Fatal(err)
	}

	var dataArr []string
	isFasta := false

	// create new Scanner to read from file
	scanner := bufio.NewScanner(file)
	// use scan method to read file
	for scanner.Scan() {
		// if FASTA line, continue
		if isFasta {
			continue
		}

		// store line in variable
		line := scanner.Text()

		// ignore if line is comment or FASTA
		if strings.HasPrefix(line, "##") {
			if strings.HasPrefix(line, "###") {
				isFasta = true
			}
			continue
		}

		// split line by tab
		parts := strings.Split(line, "\t")

		// if pseudogene, add it to slice
		if parts[2] == c.String("type") {
			dataArr = append(dataArr, line)
			// print line to output file
			fmt.Fprintln(output, line)
		}

	}
	// print length of array (# of pseudogenes)
	fmt.Println(len(dataArr))

	// if some type of error, log it
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
