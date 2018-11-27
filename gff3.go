package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

// GFF3ReadAndWrite reads GFF3 file and writes new one based on type
func GFF3ReadAndWrite() {
	filePtr := flag.String("file", "sample.gff3", "input file")
	dataPtr := flag.String("type", "pseudogenes", "type of data to extract")
	outputPtr := flag.String("output", "pseudogenes.gff3", "output file")

	flag.Parse()

	// open input file
	file, err := os.Open(*filePtr)
	if err != nil {
		log.Fatal(err) // equivalent to Print() followed by os.Exit(1)
	}
	// close file when done
	defer file.Close()

	// create output file
	output, err := os.Create(*outputPtr)
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
		if parts[2] == *dataPtr {
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
