package main

import "math"

type Coordinate struct {
	id int
	lat, lon float64
}

func deg2rad(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func (c1 *Coordinate) distance(c2 Coordinate) float64 {
	const earthRadius = 6371

	distLat := deg2rad(c2.lat - c1.lat) // distance between latitude points in radians
	distLon := deg2rad(c2.lon - c1.lon) // distance between longitude points in radians
	lat1 := deg2rad(c1.lat)
	lat2 := deg2rad(c2.lat)

	a := math.Pow(math.Sin(distLat / 2), 2) + math.Pow(math.Sin(distLon / 2), 2) * math.Cos(lat1) * math.Cos(lat2) // haversine formula of two points

	distance := 2 * earthRadius * math.Asin(math.Sqrt(a))

	return distance
}
