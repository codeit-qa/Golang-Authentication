package EmailVerification

import (
	database "GO/database"
	model "GO/models"
	"encoding/json"
	"net/http"
)

func HandleEmailVerification(response http.ResponseWriter, request *http.Request) {

	AuthToken := request.Header.Get("Authenticate")

	var code model.Code
	err := json.NewDecoder(request.Body).Decode(&code)

	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	if request.Method != "POST" {
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte("{\"message\": \"Method not allowed\"}"))
		return
	} else if AuthToken == "" {
		response.WriteHeader(http.StatusUnauthorized)
		response.Write([]byte("{\"message\": \"Authorization Token is not provided\"}"))
		return
	}

	status := database.HandleTokenAuthentication("GO", "tokens", AuthToken, code.Code)

	if !status {
		response.WriteHeader(http.StatusUnauthorized)
		response.Write([]byte("{\"message\": \"Unathorized\"}"))
		return
	} else if status {
		response.WriteHeader(http.StatusOK)
		response.Write([]byte("{\"message\": \"Sucess\"}"))

	}

}
