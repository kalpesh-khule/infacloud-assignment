package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//initializing the mux router
	r := mux.NewRouter()

	//to allow access for API which does not require authentication
	r.HandleFunc("/fetchPublicInfo", fetchPublicInfo).Methods("GET")

	//to allow access to only logged in users, role does not matter.
	//Any authenticated user should be allowed to access this API
	r.HandleFunc("/anyLoggedInUser", authenticate(anyLoggedInUser, "any")).Methods("GET")

	//Access restricted API. Only to be allowed to access by certain user group.
	//in this case only users with admin authentication will be allowed access
	r.HandleFunc("/onlyAdminCanAccess", authenticate(onlyAdminCanAccess, "admin")).Methods("POST")

	//Access restricted API, where multiple user groups are allowed access
	//in this example, admin and employee are allowed access to this API
	r.HandleFunc("/editSensitiveInfo", authenticate(editSensitiveInfo, "admin,employee")).Methods("POST")

	http.ListenAndServe(":8080", r)

}
