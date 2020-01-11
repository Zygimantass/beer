package app

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

type Brewery struct {
	id       int
	name     string
	address1 string
	address2 string
	city     string
	state    string
	code     string
	country  string
	phone    string
	website  string
	filepath string
	descript string
}

func getBreweriesFromCsv(filename string) map[int]Brewery {
	breweries := make(map[int]Brewery)

	csvFile, _ := os.Open(filename)
	csvReader := csv.NewReader(bufio.NewReader(csvFile))

	for {
		line, err := csvReader.Read()

		if len(line) > 0 {
			if line[0] == "id" {
				continue // first line skipped
			}
		}

		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		id, err := strconv.Atoi(line[0])

		if err != nil {
			log.Println(err.Error())
			continue
		}

		breweries[id] = Brewery{
			id:       id,
			name:     line[1],
			address1: line[2],
		}
	}

	return breweries
}
