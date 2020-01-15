package app_test

import (
	"github.com/Zygimantass/beer/app"
	"github.com/Zygimantass/beer/models"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetTrip(t *testing.T) {
	assert := assert.New(t)

	homePoint := models.Brewery{
		ID:   -1,
		Name: "Home",
		Location: models.Coordinate{
			Lat: 0,
			Lon: 0,
		},
	}

	edges := []models.Brewery{
		models.Brewery{
			ID: 1,
            Name: "Point1",
            Location: models.Coordinate{
				Lat: 1,
				Lon: 1,
			},
			BeerTypeCount: 20,
		}, // in range of homePoint
		models.Brewery{
			ID: 2,
			Name: "Point2",
			Location: models.Coordinate{
				Lat: 2,
				Lon: 2,
			},
			BeerTypeCount: 1,
		},
		models.Brewery{
			ID: 3,
			Name: "Point3",
			Location: models.Coordinate{
				Lat: 50,
				Lon: 50,
			},
		},
	}

	path := app.FindPath(homePoint, edges, 2000.0)

	assert.Equal(4, len(path.Points), "the path length should be four")
	assert.Equal(1, path.Points[1].ID, "the first point should be the one with the higher beer count")
	assert.Equal(2, path.Points[2].ID, "the second point should be the one with the lower beer count")

	invalidPath := app.FindPath(homePoint, edges[2:], 2000.0)

	assert.Equal(2, len(invalidPath.Points), "the invalid's path length should be two")
}
