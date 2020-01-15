package api

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"strconv"
)

// TripRouteManager manages routes related to trip finding
type TripRouteManager struct {
	api *API
}

// TripRoutes mounts trip endpoints to the router
func (trm *TripRouteManager) TripRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/find", trm.GetTrip)
	return r
}

// GetTrip parses request's arguments and returns the nearest path given by App.GetTrip
func (trm *TripRouteManager) GetTrip(w http.ResponseWriter, r *http.Request) {
	latitude, ok := r.URL.Query()["lat"]

	if !ok || len(latitude) < 1 {
		http.Error(w, "You have to include the latitude of your coordinates", http.StatusBadRequest)
		return
	}

	latitudeFloat, err := strconv.ParseFloat(latitude[0], 64)

	if err != nil {
		http.Error(w, "Your latitude argument is invalid", http.StatusBadRequest)
		return
	}

	longitude, ok := r.URL.Query()["lon"]
	if !ok || len(longitude) < 1 {
		http.Error(w, "You have to include the longitude of your coordinates", http.StatusBadRequest)
	}

	longitudeFloat, err := strconv.ParseFloat(longitude[0], 64)
	if err != nil {
		http.Error(w, "Your longitude argument is invalid", http.StatusBadRequest)
		return
	}

	path, err := trm.api.App.GetTrip(latitudeFloat, longitudeFloat)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		log.Fatal(err.Error())
	}

	beerCount, err := trm.api.App.GetBeerCount(path.Points)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		log.Fatal(err.Error())
	}

	path.BeerCount = beerCount

	beerTypes, err := trm.api.App.GetBeerTypes(path.Points)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		log.Fatal(err.Error())
	}

	path.BeerTypes = beerTypes

	pathJSON, err := json.Marshal(path)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		log.Fatal(err.Error())
	}

	w.Write(pathJSON)
}
