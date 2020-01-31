package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func redirect(w http.ResponseWriter, r *http.Request) {
	//take short code as GET parameter
	param := mux.Vars(r)
	shortCode := param["shortCode"]

	//identify URL associated with it and redirect user to appropriate url
	//redirect to 404 if no url found

	if val, ok := urls[shortCode]; ok {

		http.Redirect(w, r, val, 301)
		return

	} else {

		//we can redirect user to page designed for frontend
		w.Write([]byte("<h1>Invalid URL</h1>"))
		w.WriteHeader(404)
		return
	}
}
