package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {

	//data, err := ioutil.ReadFile("geocodes.csv")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Print(string(data))

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
	currPoint := firstPoint
	fuel := 2000.0
	visited := make(map[int]bool)
	visited[firstPoint.id] = true

	for {
		if fuel < 0 {
			break
		}

		minDistance := 99999.0
		minPoint := Coordinate {

		}

		for _, point := range coordinates {
			id := point.id

			if point != currPoint {
				dist := currPoint.distance(point)
				distHome := firstPoint.distance(point)

				if dist + distHome > fuel {
					continue
				}

				if dist != 0 && dist < minDistance && !visited[id] {
					minPoint = point
					minDistance = dist + distHome
				}
			}
		}


		if minDistance == 99999.0 {
			break
		}

		distance := currPoint.distance(minPoint)
		currPoint = minPoint
		visited[minPoint.id] = true

		fuel -= distance
	}
}