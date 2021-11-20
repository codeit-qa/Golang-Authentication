package routes

import (
	"encoding/json"
	"net/http"
	"time"

	database "GO/database"
	helper "GO/helpers"
	model "GO/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func HandleSignup(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	var user model.UserModel

	err := json.NewDecoder(request.Body).Decode(&user)

	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.ID = primitive.NewObjectID()
	user.User_id = user.ID.Hex()
	token, refreshToken, _ := helper.JWTTokenGenerator(user.Email, user.First_name, user.Last_name, user.User_id)
	user.Token = token
	user.Refresh_token = refreshToken

	database.HandleDatabase("GO", "users", user)

	json.NewEncoder(response).Encode(user)

}