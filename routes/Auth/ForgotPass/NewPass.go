package routes

import (
	database "GO/database"
	helpers "GO/helpers"
	models "GO/models"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func HandleNewPassword(response http.ResponseWriter, request *http.Request) {

	AuthToken := request.Header.Get("Authenticate")

	var forgot models.NewPassword

	if request.Method == "POST" {
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte("{\"message\": \"Method not allowed\"}"))
		return
	} else if AuthToken == "" {
		response.WriteHeader(http.StatusUnauthorized)
		response.Write([]byte("{\"message\": \"Authorization Token is not provided\"}"))
		return
	}

	err := json.NewDecoder(request.Body).Decode(&forgot)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("{\"message\": \"Invalid request body\"}"))
		return
	}

	if forgot.NewPassword == "" || forgot.ConfirmPassword == "" {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("{\"message\": \"Both password fields are required\"}"))
		return
	} else if forgot.NewPassword != forgot.ConfirmPassword {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("{\"message\": \"Both password should match\"}"))
		return
	}

	auth := database.HandleTokenAuthentication("GO", "tokens", AuthToken, 0)
	if !auth {
		response.WriteHeader(http.StatusUnauthorized)
		response.Write([]byte("{\"message\": \"Unathorized\"}"))
		return
	} else if auth {
		claim, errToken := helpers.ValidateToken(AuthToken)
		if !errToken {
			response.WriteHeader(http.StatusUnauthorized)
			response.Write([]byte("{\"message\": \"Unathorized\"}"))
			return
		}
		encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(forgot.NewPassword), 10)
		updateStatus := database.HandleUpdatePassword("GO", "users", claim.Email, string(encryptedPassword))

		if updateStatus {
			response.WriteHeader(http.StatusOK)
			response.Write([]byte("{\"message\": \"Sucess\"}"))
		} else if !updateStatus {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte("{\"message\": \"Internal Server Error\"}"))
		}

	}

}
