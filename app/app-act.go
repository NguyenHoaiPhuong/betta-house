package app

import (
	"log"
	"net/http"
	"time"
)

// PrintConfig prints all configurations
func (app *App) PrintConfig() {
	app.cfg.Print()
}

// Run App
func (app *App) Run() {
	listenPort := ":" + *app.cfg.APIPort
	srv := &http.Server{
		Handler:      app.api.Router,
		Addr:         listenPort,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
