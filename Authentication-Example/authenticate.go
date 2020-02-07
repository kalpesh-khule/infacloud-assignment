package main

import (
	"net/http"
	"strings"
)

func authenticate(function http.HandlerFunc, allowedRole string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		//fetch authentication information from API request headers

		//Database connecton
		db, err := connectDB()
		if err != nil {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.WriteHeader(500)
			return
		}
		defer db.Close()

		//code to fetch validation information related to authentication token
		result, err := db.Query("")
		if err != nil {
			w.WriteHeader(500)
			return
		}
		defer result.Close()

		//code to scan authentication information and user role from result
		//status is validity of authentication token
		//role is user role, whether user is end-user, employee or admin etc
		var status bool
		var userRole string
		result.Next()
		err = result.Scan(&status, &userRole)

		//if there is error in scanning, it means no record is found for auth information
		if err != nil {
			w.WriteHeader(401)
			return
		}

		//code to check if authentication information is valid (true) or not
		// AND if user's role is allowed to access certain functionality or not
		if status == true && (strings.Contains(allowedRole, userRole) || allowedRole == "any") {
			function(w, r)
		} else {
			w.WriteHeader(401)
			return
		}

	}

}
