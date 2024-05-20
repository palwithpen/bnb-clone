package controllers

import (
	"fmt"
	"net/http"
)

func SaveUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Save user endpoint")
}
