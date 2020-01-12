package models

type Brewery struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Address1      string `json:"address"`
	BeerTypeCount int
	Location      Coordinate
}
