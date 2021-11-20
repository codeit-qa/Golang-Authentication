package routes

import (
	"encoding/json"
	"net/http"

	database "GO/database"
	model "GO/models"
)

func HandleSignin(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	if request.Method != "POST" {
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte("{\"message\": \"Method not allowed\"}"))
		return
	}

	var user model.AuthenticationModel

	err := json.NewDecoder(request.Body).Decode(&user)

	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	auth := database.HandleAuthentication(user.Email, user.Password, "GO", "users", &user)

	if !auth {
		response.WriteHeader(http.StatusUnauthorized)
		response.Write([]byte("{\"message\": \"Invalid credentials\"}"))
		return
	}
	if auth {
		response.WriteHeader(http.StatusOK)
		response.Write([]byte("{\"message\": \"Successfully signed in\"}"))

	}
}
