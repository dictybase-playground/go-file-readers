package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// set variables for opening file based on CLI arg
	file, err := os.Open(os.Args[1])

	// if error isn't empty, do this:
	if err != nil {
		// equivalent to Print() followed by os.Exit(1)
		log.Fatal(err)
	}
	// close file when done
	defer file.Close()

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
		if parts[2] == "pseudogene" {
			dataArr = append(dataArr, line)
		}

	}
	// print length of array (# of pseudogenes)
	fmt.Println(len(dataArr))

	// if some type of error, log it
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
