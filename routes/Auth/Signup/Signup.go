package routes

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	database "GO/database"
	helper "GO/helpers"
	model "GO/models"

	"github.com/avct/uasurfer"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func HandleSignup(response http.ResponseWriter, request *http.Request) {

	if request.Method != "POST" {
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte("{\"message\": \"Method not allowed\"}"))
		return
	}

	var user model.UserModel
	var result model.ResponseModel

	err := json.NewDecoder(request.Body).Decode(&user)

	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	userAgent := request.Header.Get("User-Agent")
	DeviceInfo := uasurfer.Parse(userAgent)

	var agent model.UserAgent

	agent.UserAgent = userAgent
	agent.OS.Name = DeviceInfo.OS.Name.String()
	agent.OS.Platform = DeviceInfo.OS.Platform.String()
	agent.OS.Version = string(rune(DeviceInfo.OS.Version.Major)) + "." + string(rune(DeviceInfo.OS.Version.Minor)) + "." + string(rune(DeviceInfo.OS.Version.Patch))
	agent.Browser.Name = DeviceInfo.Browser.Name.String()
	agent.Browser.Version = string(rune(DeviceInfo.Browser.Version.Major)) + "." + string(rune(DeviceInfo.Browser.Version.Minor)) + "." + string(rune(DeviceInfo.Browser.Version.Patch))
	agent.DeviceType = DeviceInfo.DeviceType.String()

	user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.ID = primitive.NewObjectID()
	user.User_id = user.ID.Hex()
	token, refreshToken, _ := helper.JWTTokenGenerator(user.Email, user.First_name, user.Last_name, user.User_id)
	user.Token = token
	user.Refresh_token = refreshToken
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(encryptedPassword)
	result.Token = token
	result.Expires_in = time.Now().Local().Add(time.Hour * time.Duration(24)).Unix()
	generatedCode := helper.HandleCodeGenerator(6)
	code, _ := strconv.Atoi(generatedCode)
	insertErr := database.HandleDatabaseInsert("GO", "users", user.Email, user.Phone, user.Password, user.First_name, user.Last_name, user.User_id, user.Created_at, user.Updated_at, user.Token, code, agent)

	if insertErr {
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(&result)
		helper.HandleEmailService(user.Email, code)

	} else {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("{\"message\": \"Duplicate Data\"}"))
		return
	}

}
