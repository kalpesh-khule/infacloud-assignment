package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	//initializing the mux router
	r := mux.NewRouter()

	r.HandleFunc("/shorten", shortenURL).Methods("POST")
	r.HandleFunc("/{shortCode}", redirect).Methods("GET")

	handler := cors.AllowAll().Handler(r)
	http.ListenAndServe(":8080", handler)

}
