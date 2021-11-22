package main

import (
	"net/http"

	Signin "GO/routes/Auth/Signin"
	Signup "GO/routes/Auth/Signup"
	EmailVerification "GO/routes/EmailVerification"
)

func main() {

	http.HandleFunc("/signup", Signup.HandleSignup)
	http.HandleFunc("/signin", Signin.HandleSignin)
	http.HandleFunc("/emailVerify", EmailVerification.HandleEmailVerification)

	http.ListenAndServe(":8080", nil)

}
