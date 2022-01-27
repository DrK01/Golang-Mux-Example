package main

import (
	"app/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.IndexHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
