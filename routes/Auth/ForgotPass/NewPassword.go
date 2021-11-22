package routes

import (
	models "GO/models"
	"encoding/json"
	"net/http"
)

func HandleNewPassword(response http.ResponseWriter, request *http.Request) {

	if request.Method == "POST" {
		if request.Method != "POST" {
			response.WriteHeader(http.StatusMethodNotAllowed)
			response.Write([]byte("{\"message\": \"Method not allowed\"}"))
			return
		}
	}

	var auth models.NewPassword

	err := json.NewDecoder(request.Body).Decode(&auth)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	if auth.NewPassword != auth.ConfirmPassword {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("{\"message\": \"Password and Confirm Password do not match\"}"))
		return
	}

}
