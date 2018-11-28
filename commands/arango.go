package commands

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	cli "gopkg.in/urfave/cli.v1"
)

// StoreCSVinDB reads CSV file, converts to User struct, saves in ArangoDB
func StoreCSVinDB(c *cli.Context) error {
	// initialize connection to arangodb
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{"http://" + c.String("host") + c.String("port")},
	})
	if err != nil {
		log.Fatal(err)
	}
	// initialize client
	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(c.String("user"), c.String("pw")),
	})
	if err != nil {
		log.Fatal(err)
	}

	ctx := driver.WithQueryCount(context.Background())
	db, err := client.Database(ctx, c.String("db"))
	if err != nil {
		log.Fatal(err)
	}

	col, err := db.Collection(ctx, c.String("collection"))
	if err != nil {
		log.Fatal(err)
	}

	// open input file
	file, err := os.Open(c.String("file"))
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
	docs, errs, err := col.CreateDocuments(nil, users)
	if err != nil {
		log.Fatalf("Failed to create documents: %v", err)
	} else if err := errs.FirstNonNil(); err != nil {
		log.Fatalf("Failed to create documents: first error: %v", err)
	}

	fmt.Printf("Created %d documents in collection '%s' in database '%s'\n", len(docs), col.Name(), db.Name())

	return nil
}
