package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

// User data structure
type User struct {
	firstName     string `json:"first_name"`
	lastName      string `json:"last_name"`
	email         string `json:"email"`
	organization  string `json:"organization"`
	groupName     string `json:"group_name"`
	firstAddress  string `json:"first_address"`
	secondAddress string `json:"second_address"`
	city          string `json:"city"`
	state         string `json:"state"`
	zipcode       string `json:"zipcode"`
	country       string `json:"country"`
	phone         string `json:"phone"`
	isActive      bool   `json:"is_active"`
	// created_at
	// updated_at
}

func main() {
	filePtr := flag.String("file", "users.csv", "input file")

	flag.Parse()

	// open input file
	file, err := os.Open(*filePtr)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var users []User

	// read file into variable
	r := csv.NewReader(file)
	_, err = r.Read()
	if err != nil {
		log.Fatal(err)
	}

	for {
		r, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		data := User{
			firstName:     r[1],
			lastName:      r[2],
			email:         r[0],
			organization:  r[6],
			groupName:     r[4], // is this right?
			firstAddress:  r[7],
			secondAddress: r[8],
			city:          r[9],
			state:         r[10],
			zipcode:       r[13],
			country:       r[12],
			phone:         r[15],
			isActive:      true,
		}

		users = append(users, data)

		fmt.Println(data.firstName + " " + data.lastName)
	}
	// Convert to JSON
	// jsonData, err := json.Marshal(users)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// jsonFile, err := os.Create("./data.json")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer jsonFile.Close()

	// jsonFile.Write(jsonData)
	// jsonFile.Close()
}
