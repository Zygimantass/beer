package api

import (
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

type TripRouteManager struct {
	api *API
}

func (trm *TripRouteManager) TripRoutes () *chi.Mux {
	r := chi.NewRouter()

	return r
}

func (trm *TripRouteManager) GetTrip(w http.ResponseWriter, r *http.Request) {
	latitude, ok := r.URL.Query()["lat"]

	if !ok || len(latitude) < 1 {
		http.Error(w, "You have to include the latitude of your coordinates", http.StatusBadRequest)
		return
	}

	latitudeFloat, err := strconv.ParseFloat(latitude[0], 64)

	if (err != nil) {
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

	_ := trm.api.App.GetTrip(latitudeFloat, longitudeFloat)
}


