package main

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func CreateTemplate(name string) *template.Template {
	templates := template.New("_layout.html")
	return template.Must(templates.ParseFiles(
		"templates/_layout.html",
		filepath.Join("templates", name+".html"),
	))

}

type Controller struct {
}

func (ctrl *Controller) HomeHandler(writer http.ResponseWriter, request *http.Request) {
	CreateTemplate("index").Execute(writer, map[string]interface{}{
		"name": "home",
	})
}

func (ctrl *Controller) FormHandler(writer http.ResponseWriter, request *http.Request) {
	http.Redirect(writer, request, "/", http.StatusFound)
}
