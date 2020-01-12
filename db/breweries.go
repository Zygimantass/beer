package db

import (
	"github.com/Zygimantass/beer-backend/models"
	"log"
)

func (db *Database) GetBreweries() ([]models.Brewery, error) {
	rows, err := db.Connection.Query("SELECT breweries.id, coordinates.latitude, " +
                          			 "coordinates.longitude, breweries.name, breweries.address1 " +
									 "FROM beer.coordinates\n" +
                           			 "INNER JOIN beer.breweries ON breweries.id = coordinates.brewery_id;")

	if err != nil {
		return nil, err
	}

	var breweries []models.Brewery

	defer rows.Close()
	for rows.Next() {
		brewery := models.Brewery{
			Location: models.Coordinate{},
		}

		err := rows.Scan(&brewery.Id, &brewery.Location.Lat,
						 &brewery.Location.Lon, &brewery.Name,
						 &brewery.Address1)
		if err != nil {
			log.Fatal(err)
			continue
		}

		breweries = append(breweries, brewery)
	}

	return breweries, nil
}
