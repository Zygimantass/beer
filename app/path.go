package app

import "github.com/Zygimantass/beer-backend/models"

type Path struct {
	FuelUsed  float64          `json:"fuelUsed"`
	Points    []models.Brewery `json:"points"`
	BeerCount int              `json:"beerCount"`
	BeerTypes []string		   `json:"beerTypes"`
}

func findPath(origin models.Brewery, edges []models.Brewery, fuel float64) Path {
	currentPoint := origin

	var points []models.Brewery
	points = append(points, currentPoint)

	fuelUsed := 0.0

	visited := make(map[int]bool)
	visited[currentPoint.Id] = true

	for {
		if fuel < 0 {
			break // check if more fuel is left
		}

		minWeight := 99999.0
		minPoint := models.Brewery{}

		for _, point := range edges {
			id := point.Id

			if point == currentPoint {
				continue
			}

			dist := currentPoint.Location.Distance(point.Location)
			distHome := origin.Location.Distance(point.Location)

			if dist+distHome > fuel { // check if we can make it back home from the next point
				continue
			}

			weight := (dist + distHome) / float64(point.BeerTypeCount) // weight is based on distance versus count of beer types

			if dist != 0 && weight < minWeight && !visited[id] { // check if we have been in the point and if it is the nearest point so far
				minPoint = point
				minWeight = weight
			}
		}

		if minWeight == 99999.0 || minPoint.Id == 0 {
			break
		}

		distance := currentPoint.Location.Distance(minPoint.Location)
		currentPoint = minPoint
		visited[minPoint.Id] = true

		points = append(points, currentPoint)
		fuel -= distance
		fuelUsed += distance
	}

	distHome := currentPoint.Location.Distance(origin.Location)

	fuel -= distHome
	fuelUsed += distHome

	points = append(points, origin)

	return Path{
		FuelUsed: fuelUsed,
		Points:   points,
	}
}
