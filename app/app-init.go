package app

import (
	"github.com/NguyenHoaiPhuong/betta-house/api"
	"github.com/NguyenHoaiPhuong/betta-house/config"
)

// Init initializes internal members inside App struct
func (app *App) Init() {
	app.initConfig()
	app.initAPI()
}

func (app *App) initConfig() {
	app.cfg = config.NewConfig("resources/conf.json")
}

func (app *App) initAPI() {
	app.api = api.NewAPI()
	app.api.RegisterHandleFunction("GET", "/", home)
}
