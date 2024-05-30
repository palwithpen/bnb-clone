package controllers

import (
	"encoding/json"
	"net/http"

	"palwithpen.github.com/airbnb/src/ro"
	"palwithpen.github.com/airbnb/src/utils"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var login ro.Login
	decode := json.NewDecoder(r.Body)

	if err := decode.Decode(&login); err != nil {
		utils.JSONResponse(w, "error", http.StatusBadRequest, nil, map[string]string{"error": err.Error()})
	}

	defer r.Body.Close()

	if err := utils.ValidateStruct(login); err != nil {
		validationErrors := utils.ValidationErrorToMap(err)
		utils.JSONResponse(w, "error", http.StatusBadRequest, nil, validationErrors)
		return
	}
	utils.JSONResponse(w, "success", http.StatusOK, login, nil)
}

func ResetHandler(w http.ResponseWriter, r *http.Request) {
	var reset ro.Reset

	decode := json.NewDecoder(r.Body)

	if err := decode.Decode(&reset); err != nil {
		utils.JSONResponse(w, "error", http.StatusBadRequest, nil, map[string]string{"error": err.Error()})
	}
	defer r.Body.Close()

	if err := utils.ValidateStruct(reset); err != nil {
		validationErrors := utils.ValidationErrorToMap(err)
		utils.JSONResponse(w, "error", http.StatusBadRequest, nil, validationErrors)
		return
	}

	utils.JSONResponse(w, "success", http.StatusOK, reset, nil)
}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {

}
