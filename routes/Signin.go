package routes

import (
	"encoding/json"
	"net/http"
	"time"

	database "GO/database"
	helper "GO/helpers"
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
	var result model.ResponseModel

	err := json.NewDecoder(request.Body).Decode(&user)

	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	auth, email, fname, lname, userid := database.HandleAuthentication(user.Email, user.Password, "GO", "users")
	token, _, _ := helper.JWTTokenGenerator(email, fname, lname, userid)

	result.Token = token
	result.Expires_in = time.Now().Local().Add(time.Hour * time.Duration(24)).Unix()

	if !auth {
		response.WriteHeader(http.StatusUnauthorized)
		response.Write([]byte("{\"message\": \"Invalid credentials\"}"))
		return
	}
	if auth {
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(&result)
	}
}
