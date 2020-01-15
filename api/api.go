package api

import (
	"github.com/Zygimantass/beer/app"
	"github.com/go-chi/chi"
)

// API implements route handling and configuration of the REST API
type API struct {
	App    *app.App
	Config *Config
	TRM    *TripRouteManager
}

// New creates an API object
func New(a *app.App) (api *API, err error) {
	api = &API{
		App: a,
	}
	api.Config, err = InitConfig()

	if err != nil {
		return nil, err
	}

	api.TRM = &TripRouteManager{
		api: api,
	}

	return api, nil
}

// Init implements mounting the routes to the router
func (api *API) Init() chi.Router {
	r := chi.NewRouter()

	r.Mount("/trip/", api.TRM.TripRoutes())

	return r
}
