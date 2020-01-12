package app

import "github.com/Zygimantass/beer-backend/models"

const MAX_FUEL = 2000

func (a *App) GetTrip(lat float64, lon float64) (*Path, error) {
	breweries, err := a.Database.GetBreweries()
	if err != nil {
		return nil, err
	}

	path := findPath(models.Brewery{
		Id:   -1,
		Name: "Home",
		Location: models.Coordinate{
			Lat: lat,
			Lon: lon,
		},
	}, breweries, MAX_FUEL)

	return &path, nil
}
