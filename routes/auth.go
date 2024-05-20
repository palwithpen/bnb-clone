package routes

import (
	"net/http"

	"palwithpen.github.com/airbnb/controllers"
)

func SetupAuthRoutes(mux *http.ServeMux, prefix string) {
	mux.HandleFunc(prefix+"/login", controllers.LoginHandler)
	mux.HandleFunc(prefix+"/reset", controllers.ResetHandler)
}