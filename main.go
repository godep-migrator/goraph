package main

import (
	"html/template"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Controller struct {
}

func (ctrl *Controller) HomeHandler(writer http.ResponseWriter, request *http.Request) {

	templates := template.New("_layout.html")
	templates = template.Must(templates.ParseFiles(
		"templates/_layout.html",
		"templates/index.html",
	))

	templates.Execute(writer, map[string]interface{}{
		"name": "home",
	})
}

func main() {

	ctrl := new(Controller)

	router := mux.NewRouter()
	router.HandleFunc("/", ctrl.HomeHandler).Name("home")
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	http.Handle("/", router)
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), router); err != nil {
		panic(err)
	}
}
