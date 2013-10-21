package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var port string = "8000"

func main() {

	ctrl := new(Controller)

	router := mux.NewRouter()
	router.HandleFunc("/", ctrl.HomeHandler).Name("home")
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	http.Handle("/", router)

	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	if err := http.ListenAndServe(":"+port, router); err != nil {
		panic(err)
	}
}
