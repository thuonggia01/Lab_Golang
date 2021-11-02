package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	f, err := os.Open("users.csv")

	if err != nil {

		log.Fatal(err)
	}

	r := csv.NewReader(f)

	for {

		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		for value := range record {
			fmt.Printf("%s\n", record[value])
		}
	}
}
