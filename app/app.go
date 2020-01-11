package app

type App struct {

}

func New() (app *App, err error) {
	app = &App{}

	return app, nil
}
