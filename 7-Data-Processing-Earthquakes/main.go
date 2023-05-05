package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type record struct {
	datetime     string
	text         string
	coordinates  string
	retweetCount string
	likeCount    string
	rescue       string
}

func main() {
	// Open the CSV file
	file, err := os.Open("izmir_earthquake_labeled.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Set the delimiter
	reader.Comma = ','

	// Create an empty slice of records
	var records []record

	// Loop through the records
	for {
		// Read each record from CSV
		row, err := reader.Read()

		// Exit if we reach end of file
		if err == io.EOF {
			break
		}

		// Handle other errors
		if err != nil {
			log.Fatal(err)
		}

		// Create a new record object
		rec := record{}

		// Assign values to the record object
		rec.datetime = row[0]
		rec.text = row[1]
		rec.coordinates = row[2]
		rec.retweetCount = row[3]
		rec.likeCount = row[4]
		rec.rescue = row[5]

		// Append the record to the records slice
		records = append(records, rec)
	}

	// Process the records
	for _, rec := range records {
		// Print the record to console
		fmt.Printf("Datetime: %s\nText: %s\nCoordinates: %s\nRetweet Count: %s\nLike Count: %s\nRescue: %s\n\n",
			rec.datetime, rec.text, rec.coordinates, rec.retweetCount, rec.likeCount, rec.rescue)
	}
}
