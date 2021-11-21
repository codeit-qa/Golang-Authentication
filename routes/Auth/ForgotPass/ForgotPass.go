package routes

import (
	database "GO/database"
	models "GO/models"
	"encoding/json"
	"net/http"
)

func HandleForgotPass(response http.ResponseWriter, request *http.Request) {

	if request.Method != "POST" {
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte("{\"message\": \"Method not allowed\"}"))
		return
	}

	var forgot models.ForgotPass

	err := json.NewDecoder(request.Body).Decode(&forgot)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	status := database.HandleForgotPass(forgot.Email, "GO", "users")

	if !status {
		response.WriteHeader(http.StatusUnauthorized)
		response.Write([]byte("{\"message\": \"Email not exists\"}"))
		return
	} else if status {
		response.WriteHeader(http.StatusOK)

	}
}
