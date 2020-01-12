package app

import "github.com/Zygimantass/beer-backend/db"

type App struct {
	Database *db.Database
}

func New() (app *App, err error) {
	app = &App{}

	dbConfig, err := db.InitDatabaseConfig()
	if err != nil {
		return nil, err
	}

	database, err := db.New(dbConfig)
	if err != nil {
		return nil, err
	}

	app.Database = database
	return app, nil
}
