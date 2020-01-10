package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {

	csvFile, _ := os.Open("geocodes.csv")
	csvReader := csv.NewReader(bufio.NewReader(csvFile))

	var coordinates []Coordinate

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

		id, err := strconv.Atoi(line[1])

		if err != nil {
			log.Println(err.Error())
			continue
		}

		lat, err := strconv.ParseFloat(line[2], 64)

		if err != nil {
			log.Println(err.Error())
			continue
		}

		lon, err := strconv.ParseFloat(line[3], 64)

		if err != nil {
			log.Println(err.Error())
			continue
		}

		coordinates = append(coordinates, Coordinate{
			id,
			lat,
			lon,
		})
	}

	firstPoint := Coordinate {
		id: -1,
		lat: 51.355468,
		lon: 11.100790,
	}

	//firstPoint := Coordinate {
	//	id: -1,
	//	lat: 54.895589,
	//	lon: 23.886463,
	//}

	path := findPath(firstPoint, coordinates, 2000)

	//breweries := getBreweriesFromCsv("breweries.csv")

	for _, brewery := range path.points {
		if brewery.id == -1 {
			//log.Println("HOME")
			//continue
		}
		fmt.Printf("/+%f,+%f", brewery.lat, brewery.lon)
		//log.Printf("%v, %s, %f, %f", brewery.id, breweries[brewery.id].name, brewery.lat, brewery.lon)
	}
}
