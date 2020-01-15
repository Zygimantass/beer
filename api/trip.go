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
		http.Error(w, "The latitude of the coordinates has to be included in the request", http.StatusBadRequest)
		return
	}

	latitudeFloat, err := strconv.ParseFloat(latitude[0], 64)

	if err != nil {
		http.Error(w, "The latitude argument is invalid", http.StatusBadRequest)
		return
	}

	if latitude <= -90 || latitude > 90 {
		http.Error(w, "The latitude has to be between -90 and 90 degrees", http.StatusBadRequest)
	}

	longitude, ok := r.URL.Query()["lon"]
	if !ok || len(longitude) < 1 {
		http.Error(w, "The longitude of the coordinates has to be included in the request", http.StatusBadRequest)
	}

	longitudeFloat, err := strconv.ParseFloat(longitude[0], 64)
	if err != nil {
		http.Error(w, "The longitude argument is invalid", http.StatusBadRequest)
		return
	}

	if longitude <= -180 || longitude >= 180 {
		http.Error(w, "The latitude has to be between -180 and 180 degrees", http.StatusBadRequest)
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
