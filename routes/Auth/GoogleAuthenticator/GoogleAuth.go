package routes

import (
	"encoding/base32"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	models "GO/models"

	dgoogauth "github.com/dgryski/dgoogauth"
	qr "rsc.io/qr"
)

const (
	qrFilename = "/tmp/qr.png"
)

var secret = []byte{'H', 'e', 'l', 'l', 'o', '!', 0xDE, 0xAD, 0xBE, 0xEF}

func HandleGoogleAuthenticate(response http.ResponseWriter, request *http.Request) {

	var google models.GoogleAuthenticate

	errBody := json.NewDecoder(request.Body).Decode(&google)

	if errBody != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("{\"message\": \"Invalid request body\"}"))
		return
	}

	secretBase32 := base32.StdEncoding.EncodeToString(secret)

	account := google.Email
	issuer := "Go-Authentication"

	URL, err := url.Parse("otpauth://totp")
	if err != nil {
		panic(err)
	}

	URL.Path += "/" + url.PathEscape(issuer) + ":" + url.PathEscape(account)

	params := url.Values{}
	params.Add("secret", secretBase32)
	params.Add("issuer", issuer)

	URL.RawQuery = params.Encode()
	fmt.Printf("URL is %s\n", URL.String())

	code, err := qr.Encode(URL.String(), qr.Q)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
	b := code.PNG()
	err = ioutil.WriteFile(qrFilename, b, 0600)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}

	fmt.Fprintf(response, "QR code is in %s. Please scan it into Google Authenticator app.\n", qrFilename)
	response.WriteHeader(http.StatusOK)

}

func HandleCodeAuth(response http.ResponseWriter, request *http.Request) {

	var google models.GoogleAuthenticate

	errBody := json.NewDecoder(request.Body).Decode(&google)

	if errBody != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("{\"message\": \"Invalid request body\"}"))
		return
	} else if request.Method != "POST" {
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte("{\"message\": \"Method not allowed\"}"))
		return
	} else if google.Code == "" {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("{\"message\": \"Unathorized\"}"))
		return
	}

	secretBase32 := base32.StdEncoding.EncodeToString(secret)

	otpc := &dgoogauth.OTPConfig{
		Secret:      secretBase32,
		WindowSize:  3,
		HotpCounter: 0,
	}

	val, err := otpc.Authenticate(google.Code)
	if err != nil {
		fmt.Println(err)
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	if !val {
		response.WriteHeader(http.StatusUnauthorized)
		response.Write([]byte("{\"message\": \"Invalid Code, make sure your system date and time is valid also\"}"))
		return
	}
	response.WriteHeader(http.StatusOK)
	response.Write([]byte("{\"message\": \"Authenticated Sucessfully\"}"))

}
