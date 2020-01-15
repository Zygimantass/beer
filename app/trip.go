package app

import "github.com/Zygimantass/beer/models"

// MaxFuel is the fuel constraint
const MaxFuel = 2000

// GetTrip returns the best path given the latitude and longitude
func (a *App) GetTrip(lat float64, lon float64) (*Path, error) {
	breweries, err := a.Database.GetBreweries()
	if err != nil {
		return nil, err
	}

	path := FindPath(models.Brewery{
		ID:   -1,
		Name: "Home",
		Location: models.Coordinate{
			Lat: lat,
			Lon: lon,
		},
	}, breweries, MaxFuel)

	return &path, nil
}

// GetBeerCount returns the count of beers tasted in the breweries given
func (a *App) GetBeerCount(breweries []models.Brewery) (int, error) {
	return a.Database.GetBeerCount(breweries)
}

// GetBeerTypes returns the beers tasted in the given breweries
func (a *App) GetBeerTypes(breweries []models.Brewery) ([]string, error) {
	return a.Database.GetBeerTypes(breweries)
}
