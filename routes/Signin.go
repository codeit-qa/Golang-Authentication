package routes

import (
	"fmt"
	"net/http"
)

func HandleSignin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hellow World")
}
