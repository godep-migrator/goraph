package main

import (
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
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

	weight, err := strconv.ParseFloat(request.FormValue("weight"), 32)
	if err != nil {
		panic(err)
	}

	model := Connect("mongodb://heroku_app18793481:ljai9380ctc05odrm34rnr58gg@ds051658.mongolab.com:51658/heroku_app18793481", "heroku_app18793481")
	model.InsertWeight(weight)

	http.Redirect(writer, request, "/", http.StatusFound)
}
