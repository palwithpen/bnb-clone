package controllers

import (
	"fmt"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Your login logic here
	fmt.Fprintf(w, "Login endpoint")
}

func ResetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Reset endpoint")
}
