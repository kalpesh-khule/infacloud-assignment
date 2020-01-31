package main

type requestURL struct {
	URL string `json:"url"`
}

//ShortCode response
type ShortCode struct {
	URL       string `json:"url"`
	ShortCode string `json:"shortCode"`
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

var urls = make(map[string]string)
