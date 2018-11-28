package commands

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

// User data structure
type User struct {
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Email         string `json:"email"`
	Organization  string `json:"organization"`
	GroupName     string `json:"group_name"`
	FirstAddress  string `json:"first_address"`
	SecondAddress string `json:"second_address"`
	City          string `json:"city"`
	State         string `json:"state"`
	Zipcode       string `json:"zipcode"`
	Country       string `json:"country"`
	Phone         string `json:"phone"`
	IsActive      bool   `json:"is_active"`
	// created_at
	// updated_at
}

// CSVtoJSON reads csv file, converts it to JSON and writes it to new file
func CSVtoJSON() {
	filePtr := flag.String("file", "users.csv", "input file")
	outputPtr := flag.String("output", "users.json", "output file")

	flag.Parse()

	// open input file
	file, err := os.Open(*filePtr)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var users []User

	// read file into variable
	reader := csv.NewReader(file)

	for {
		line, err := reader.Read()
		if err == io.EOF {
			// break the loop at end of file
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		users = append(users, User{
			FirstName:     line[1],
			LastName:      line[2],
			Email:         line[0],
			Organization:  line[6],
			GroupName:     line[4], // is this right?
			FirstAddress:  line[7],
			SecondAddress: line[8],
			City:          line[9],
			State:         line[10],
			Zipcode:       line[13],
			Country:       line[12],
			Phone:         line[15],
			IsActive:      true,
		})
	}
	// convert to JSON
	usersJSON, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("total number of users =", len(users))

	jsonFile, err := os.Create(*outputPtr)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(usersJSON)

	fmt.Println("JSON successfully written to", *outputPtr)
}
