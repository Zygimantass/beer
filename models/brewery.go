package models

// Brewery represents a brewery from the data
type Brewery struct {
	Id            int        `json:"id"`
	Name          string     `json:"name"`
	Address1      string     `json:"address"`
	BeerTypeCount int        `json:"beerTypeCount"`
	Location      Coordinate `json:"location"`
}
