package app

import "github.com/Zygimantass/beer/models"

// Path represents the path found by our algorithm
type Path struct {
	FuelUsed  float64          `json:"fuelUsed"` // how much fuel was used along the way
	Points    []models.Brewery `json:"points"` // breweries in the path
	BeerCount int              `json:"beerCount"` // how many beers can be tasted
	BeerTypes []string		   `json:"beerTypes"` // what were the beers tasted
}

// FindPath returns the optimal route given the edges and fuel constraints
func FindPath(origin models.Brewery, edges []models.Brewery, fuel float64) Path {
	currentPoint := origin

	var points []models.Brewery
	points = append(points, currentPoint)

	fuelUsed := 0.0

	visited := make(map[int]bool)
	visited[currentPoint.ID] = true

	for {
		if fuel < 0 {
			break // check if more fuel is left
		}

		minWeight := 99999.0
		minPoint := models.Brewery{}

		for _, point := range edges {
			id := point.ID

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

		if minWeight == 99999.0 || minPoint.ID == 0 {
			break
		}

		distance := currentPoint.Location.Distance(minPoint.Location)
		currentPoint = minPoint
		visited[minPoint.ID] = true

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
