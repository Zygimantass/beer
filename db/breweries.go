package db

import (
	"github.com/Zygimantass/beer-backend/models"
	"log"
	"strconv"
)

func (db *Database) GetBreweries() ([]models.Brewery, error) {
	rows, err := db.Connection.Query("SELECT breweries.id, coordinates.latitude, " +
		"coordinates.longitude, breweries.name, breweries.address1," +
		"COUNT(breweries.id) FROM beer.coordinates\n" +
		"INNER JOIN beer.breweries ON breweries.id = coordinates.brewery_id\n" +
		"INNER JOIN beer.beers ON breweries.id = beers.brewery_id " +
		"GROUP BY breweries.id;")

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
			&brewery.Address1, &brewery.BeerTypeCount)
		if err != nil {
			log.Fatal(err)
			continue
		}

		breweries = append(breweries, brewery)
	}

	return breweries, nil
}

func (db *Database) GetBeerCount(breweries []models.Brewery) (int, error) {
	query := "SELECT COUNT(DISTINCT beers.id) FROM breweries\n" +
		"INNER JOIN beers ON breweries.id = beers.brewery_id " +
		"WHERE"

	for i, brewery := range breweries {
		whereStmt := " breweries.id = " + strconv.Itoa(brewery.Id)
		if i != len(breweries)-1 { // add OR statement if it isn't the last brewery
			whereStmt += " OR"
		}
		query += whereStmt
	}

	query += ";"

	println(query)

	rows, err := db.Connection.Query(query)
	if err != nil {
		return 0, err
	}

	defer rows.Close()

	var beerCount int
	for rows.Next() {
		err := rows.Scan(&beerCount)
		if err != nil {
			log.Fatal(err)
			return 0, err
		}
	}

	return beerCount, nil
}

func (db *Database) GetBeerTypes(breweries []models.Brewery) ([]string, error) {
	query := "SELECT beers.name FROM breweries\n" +
		"INNER JOIN beers ON breweries.id = beers.brewery_id " +
		"WHERE"

	for i, brewery := range breweries {
		whereStmt := " breweries.id = " + strconv.Itoa(brewery.Id)
		if i != len(breweries)-1 { // add OR statement if it isn't the last brewery
			whereStmt += " OR"
		}
		query += whereStmt
	}

	query += ";"

	rows, err := db.Connection.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var beerTypes []string

	for rows.Next() {
		var beerType string

		err := rows.Scan(&beerType)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		beerTypes = append(beerTypes, beerType)
	}

	return beerTypes, nil
}
