package EmailVerification

import (
	database "GO/database"
	"fmt"
	"net/http"
)

func HandleEmailVerification(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("Content-type", "application/json")
	AuthToken := request.Header.Get("Authenticate")

	if request.Method != "POST" {
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte("{\"message\": \"Method not allowed\"}"))
		return
	} else if AuthToken == "" {
		response.WriteHeader(http.StatusUnauthorized)
		response.Write([]byte("{\"message\": \"Authorization Token is not provided\"}"))
		return
	}

	status, dbToken := database.HandleTokenAuthentication("GO", "tokens", AuthToken)

	if !status {
		response.WriteHeader(http.StatusUnauthorized)
		response.Write([]byte("{\"message\": \"Token expired\"}"))
		return
	} else if status {
		fmt.Fprintf(response, dbToken)
	}

}
