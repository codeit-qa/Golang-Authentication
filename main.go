package main

import (
	"net/http"

	routes "GO/routes"
)

func main() {

	http.HandleFunc("/login", routes.HandleSignin)

	http.ListenAndServe(":8080", nil)

}
