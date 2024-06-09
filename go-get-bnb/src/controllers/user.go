package controllers

import (
	"encoding/json"
	"net/http"

	"palwithpen.github.com/airbnb/src/ro"
	"palwithpen.github.com/airbnb/src/services"
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

	if _, err := services.SaveUser(user); err != nil {
		utils.JSONResponse(w, "error", 1003, nil, err)
	}

	utils.JSONResponse(w, "success", http.StatusOK, user, nil)

}

func GetUserDetailsHandler(w http.ResponseWriter, r *http.Request) {

}

func DeleteUserDetailsHandler(w http.ResponseWriter, r *http.Request) {

}

/*

package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "yourproject/db"
    "github.com/go-chi/chi/v5"
)

// CreateUserHandler handles creating a new user
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
    var user db.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := db.CreateUser(&user); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

// GetUserHandler handles fetching a user by ID
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
    idStr := chi.URLParam(r, "id")
    id, err := strconv.ParseInt(idStr, 10, 64)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    user, err := db.GetUserByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(user)
}

// UpdateUserHandler handles updating an existing user
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
    var user db.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := db.UpdateUser(&user); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(user)
}

// DeleteUserHandler handles deleting a user by ID
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
    idStr := chi.URLParam(r, "id")
    id, err := strconv.ParseInt(idStr, 10, 64)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := db.DeleteUser(id); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}


*/
