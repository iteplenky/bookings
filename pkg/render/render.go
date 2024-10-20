package render

import (
	"bytes"
	"github.com/iteplenky/bookings/pkg/config"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func Template(w http.ResponseWriter, tmplName string) {
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmplName]
	if !ok {
		http.Error(w, "Template not found", http.StatusNotFound)
		log.Println("error getting template:", tmplName)
		return
	}

	buf := new(bytes.Buffer)
	if err := t.Execute(buf, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("error executing template:", tmplName, err)
		return
	}

	_, err := buf.WriteTo(w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("error writing template:", tmplName, err)
		return
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := make(map[string]*template.Template)

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return cache, err
	}

	for _, page := range pages {

		name := filepath.Base(page)

		t, err := template.New(name).ParseFiles(page)
		if err != nil {
			return cache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return cache, err
		}

		if len(matches) > 0 {
			t, err = t.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return cache, err
			}
		}
		cache[name] = t
	}
	return cache, nil
}
