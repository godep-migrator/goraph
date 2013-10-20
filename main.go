package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func HomeHandler(writer http.ResponseWriter, request *http.Request) {

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

	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler).Name("home")
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	http.Handle("/", router)
	if err := http.ListenAndServe(":8000", router); err != nil {
		panic(err)
	}
}
