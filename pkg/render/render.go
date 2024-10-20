package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func Template(w http.ResponseWriter, tmplName string) {
	files, err := template.ParseFiles("./templates/" + tmplName)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	if err = files.Execute(w, nil); err != nil {
		fmt.Println("error executing template:", err)
	}
}
