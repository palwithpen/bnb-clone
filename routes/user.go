package routes

import (
	"net/http"

	"palwithpen.github.com/airbnb/controllers"
)

func SetupUserDetailsRoutes(mux *http.ServeMux, prefix string) {
	mux.HandleFunc(prefix+"/save", controllers.SaveUserHandler)
}
