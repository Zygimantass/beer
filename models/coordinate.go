package models

import "math"

type Coordinate struct {
	Lat float64 `json:"latitude"`
	Lon float64 `json:"longitude"`
}

func deg2rad(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func (c1 *Coordinate) Distance(c2 Coordinate) float64 {
	const earthRadius = 6371

	distLat := deg2rad(c2.Lat - c1.Lat) // distance between latitude points in radians
	distLon := deg2rad(c2.Lon - c1.Lon) // distance between longitude points in radians
	lat1 := deg2rad(c1.Lat)
	lat2 := deg2rad(c2.Lat)

	a := math.Pow(math.Sin(distLat/2), 2) + math.Pow(math.Sin(distLon/2), 2)*math.Cos(lat1)*math.Cos(lat2) // haversine formula of two points

	distance := 2 * earthRadius * math.Asin(math.Sqrt(a))

	return distance
}
