package models_test

import (
	"github.com/Zygimantass/beer/models"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestDistance(t *testing.T) {
	assert := assert.New(t)

	c1 := models.Coordinate{
		Lat: 30.2234001,
		Lon: -97.769699,
	}

	c2 := models.Coordinate{
		Lat: 37.7825012,
		Lon: -122.3929977,
	}

	dist1 := c1.Distance(c2)
	dist2 := c2.Distance(c1)

	assert.Equal(2411.587845213825, dist1,"The distance should be equal to the constant")
	assert.Equal(dist1, dist2, "Distances between flipped order coordinates should be equal")
}
