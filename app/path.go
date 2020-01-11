package app

type Path struct {
	fuelUsed float64
	points   []Coordinate
}

func findPath(origin Coordinate, edges []Coordinate, fuel float64) Path {
	currentPoint := origin

	var points []Coordinate
	points = append(points, currentPoint)

	fuelUsed := 0.0

	visited := make(map[int]bool)
	visited[currentPoint.id] = true

	for {
		if fuel < 0 {
			break // check if more fuel is left
		}

		minDistance := 99999.0
		minPoint := Coordinate{}

		for _, point := range edges {
			id := point.id

			if point == currentPoint {
				continue
			}

			dist := currentPoint.distance(point)
			distHome := origin.distance(point)

			if dist+distHome > fuel { // check if we can make it back home from the next point
				continue
			}

			if dist != 0 && dist < minDistance && !visited[id] { // check if we have been in the point and if it is the nearest point so far
				minPoint = point
				minDistance = dist + distHome
			}
		}

		if minDistance == 99999.0 || minPoint.id == 0 {
			break
		}

		distance := currentPoint.distance(minPoint)
		currentPoint = minPoint
		visited[minPoint.id] = true

		points = append(points, currentPoint)
		fuel -= distance
		fuelUsed += distance
	}

	distHome := currentPoint.distance(origin)

	fuel -= distHome
	fuelUsed += distHome

	points = append(points, origin)

	return Path{
		fuelUsed,
		points,
	}
}
