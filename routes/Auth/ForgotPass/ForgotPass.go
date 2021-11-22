package routes

import (
	database "GO/database"
	helpers "GO/helpers"
	models "GO/models"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
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

		generatedCode := helpers.HandleCodeGenerator(6)
		code, _ := strconv.Atoi(generatedCode)

		helpers.HandleEmailService(forgot.Email, code)

		token, _, _ := helpers.JWTTokenGenerator(forgot.Email, "", "", "")

		database.HandleInsertToken("GO", "tokens", token, code, time.Now())

	}
}
