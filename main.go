package main

import (
	"net/http"

	EmailVerification "GO/routes/Auth/EmailVerification"
	ForgotPass "GO/routes/Auth/ForgotPass"
	Signin "GO/routes/Auth/Signin"
	Signup "GO/routes/Auth/Signup"
)

func main() {

	http.HandleFunc("/signup", Signup.HandleSignup)
	http.HandleFunc("/signin", Signin.HandleSignin)
	http.HandleFunc("/emailVerify", EmailVerification.HandleEmailVerification)
	http.HandleFunc("/forgotPass", ForgotPass.HandleForgotPass)
	http.HandleFunc("/codeAuth", ForgotPass.HandleCodeAuth)
	http.HandleFunc("/newPass", ForgotPass.HandleNewPassword)

	http.ListenAndServe(":8080", nil)

}
