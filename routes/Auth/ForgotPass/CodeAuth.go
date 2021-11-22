package routes

import (
	database "GO/database"
	models "GO/models"
	"encoding/json"
	"net/http"
	"time"
)

func HandleCodeAuth(response http.ResponseWriter, request *http.Request) {

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

	var code models.Code

	err := json.NewDecoder(request.Body).Decode(&code)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("{\"message\": \"Invalid request body\"}"))
		return
	}

	auth := database.HandleTokenAuthentication("GO", "tokens", AuthToken, code.Code)

	if !auth {
		response.WriteHeader(http.StatusUnauthorized)
		response.Write([]byte("{\"message\": \"Invalid Code\"}"))
		return
	} else if auth {

		//first we will remove the token and code from the database if its true because we dont want user to use the same code again
		var result models.ResponseModel
		result.Token = AuthToken
		result.Expires_in = time.Now().Add(time.Hour * 24).Unix()

		errDel := database.HandleRemoveCode("GO", "tokens", code.Code, AuthToken)
		if !errDel {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte("{\"message\": \"Internal Server Error\"}"))
			return
		}

		//now i will add the token again to database in order to extract email from it later

		err := database.HandleInsertToken("GO", "tokens", AuthToken, 0, time.Now())
		if !err {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte("{\"message\": \"Internal Server Error\"}"))
			return
		}

		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(&result)
		return
	}

}
