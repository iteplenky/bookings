package main

import (
	"github.com/iteplenky/bookings/pkg/config"
	"github.com/iteplenky/bookings/pkg/handlers"
	"github.com/iteplenky/bookings/pkg/render"
	"log"
	"net/http"
)

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache", err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	handlers.NewHandlers(handlers.NewRepository(&app))
	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	_ = http.ListenAndServe(":8080", nil)
}
