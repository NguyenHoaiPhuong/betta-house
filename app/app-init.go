package app

import (
	"net/http"

	"github.com/NguyenHoaiPhuong/betta-house/api"
	"github.com/NguyenHoaiPhuong/betta-house/config"
	"github.com/NguyenHoaiPhuong/betta-house/repo"
)

// Init initializes internal members inside App struct
func (app *App) Init() {
	app.initConfig()
	app.initDatabase()
	app.initAPI()
	app.registerHandleFunctions()
}

func (app *App) initConfig() {
	app.cfg = config.NewConfig("resources/conf.json")
}

func (app *App) initDatabase() {
	app.db = repo.NewMongoDB(*app.cfg.MongodbServerHost, *app.cfg.MongodbPort)
}

func (app *App) initAPI() {
	app.api = api.NewAPI()
}

func (app *App) registerHandleFunctions() {
	app.api.Router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	app.api.RegisterHandleFunction("GET", "/", home)
}
