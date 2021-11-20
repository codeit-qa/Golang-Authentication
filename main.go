package main

import (
	"net/http"

	routes "GO/routes"
)

func main() {

	http.HandleFunc("/signup", routes.HandleSignup)
	http.HandleFunc("/signin", routes.HandleSignin)

	http.ListenAndServe(":8080", nil)

}
