package controllers

import (
	"encoding/json"
	"net/http"

	"palwithpen.github.com/airbnb/src/ro"
	"palwithpen.github.com/airbnb/src/utils"
)

func SaveUserHandler(w http.ResponseWriter, r *http.Request) {
	var user ro.User
	decode := json.NewDecoder(r.Body)
	if err := decode.Decode(&user); err != nil {
		utils.JSONResponse(w, "error", http.StatusBadRequest, nil, map[string]string{"error": err.Error()})
		return
	}
	defer r.Body.Close()

	if err := utils.ValidateStruct(user); err != nil {
		validationErrors := utils.ValidationErrorToMap(err)
		utils.JSONResponse(w, "error", http.StatusBadRequest, nil, validationErrors)
		return
	}
	utils.JSONResponse(w, "success", http.StatusOK, user, nil)

}

func GetUserDetailsHandler(w http.ResponseWriter, r *http.Request) {

}

func DeleteUserDetailsHandler(w http.ResponseWriter, r *http.Request) {

}
