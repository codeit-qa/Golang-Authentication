package routes

import (
	"fmt"
	"net/http"
)

func HandleSignin(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	if request.Method != "POST" {
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte("{\"message\": \"Method not allowed\"}"))
		return
	}

	fmt.Fprintf(response, "HandleSignin")
}
