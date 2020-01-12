package api

import (
	"github.com/Zygimantass/beer-backend/app"
	"github.com/go-chi/chi"
)

type API struct {
	App    *app.App
	Config *Config
	TRM    *TripRouteManager
}

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

func (api *API) Init() chi.Router {
	r := chi.NewRouter()

	r.Mount("/trip/", api.TRM.TripRoutes())

	return r
}
