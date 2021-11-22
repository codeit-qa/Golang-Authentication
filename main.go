package main

import (
	"net/http"

	EmailVerification "GO/routes/Auth/EmailVerification"
	Signin "GO/routes/Auth/Signin"
	Signup "GO/routes/Auth/Signup"
)

func main() {

	http.HandleFunc("/signup", Signup.HandleSignup)
	http.HandleFunc("/signin", Signin.HandleSignin)
	http.HandleFunc("/emailVerify", EmailVerification.HandleEmailVerification)

	http.ListenAndServe(":8080", nil)

}
