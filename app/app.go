package app

import "github.com/Zygimantass/beer/db"

// App provides business functionality and connection to the database
type App struct {
	Database *db.Database
}

// New returns an App object
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
