package routes

import (
	"net/http"

	"palwithpen.github.com/airbnb/src/controllers"
	"palwithpen.github.com/airbnb/src/utils"
)

func SetupUserDetailsRoutes(mux *http.ServeMux, prefix string) {
	mux.HandleFunc(prefix+"/save", utils.MethodHandler("POST", controllers.SaveUserHandler))
	mux.HandleFunc(prefix+"/get/", utils.MethodHandler("GET", controllers.GetUserDetailsHandler))
	mux.HandleFunc(prefix+"/delete/", utils.MethodHandler("DELETE", controllers.DeleteUserDetailsHandler))
}
