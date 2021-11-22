package EmailVerification

import (
	"fmt"
	"net/http"
)

func HandleEmailVerification(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	if request.Method != "POST" {
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte("{\"message\": \"Method not allowed\"}"))
		return
	}

	fmt.Fprintf(response, "Email verification")
}
