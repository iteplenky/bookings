package handlers

import (
	"github.com/iteplenky/bookings/pkg/render"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "main.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "about.page.tmpl")
}
