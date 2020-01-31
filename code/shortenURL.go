package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"net/url"
)

func shortenURL(w http.ResponseWriter, r *http.Request) {

	//to take URL as request body
	decoder := json.NewDecoder(r.Body)
	var body requestURL
	err := decoder.Decode(&body)
	if err != nil {
		panic(err)
	}

	//Check if the URL passed is a valid URL or not
	longURL, err := url.ParseRequestURI(body.URL)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	//shorten URL (generate random string)

	var resp ShortCode
	resp.ShortCode = generateNewShortCode()
	resp.URL = longURL.String()

	//save it into the map
	urls[resp.ShortCode] = resp.URL

	//send shortcode in response

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(resp)
	return

}

func generateNewShortCode() string {

	for true {
		shortCode := randStringBytes(8)
		if _, ok := urls[shortCode]; !ok {
			return shortCode
		}
	}
	return ""
}

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
