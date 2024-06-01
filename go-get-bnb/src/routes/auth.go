package routes

import (
	"net/http"

	"palwithpen.github.com/airbnb/src/controllers"
	"palwithpen.github.com/airbnb/src/utils"
)

func SetupAuthRoutes(mux *http.ServeMux, prefix string) {
	mux.HandleFunc(prefix+"/login", utils.MethodHandler("POST", controllers.LoginHandler))
	mux.HandleFunc(prefix+"/reset", utils.MethodHandler("POST", controllers.ResetHandler))
	mux.HandleFunc(prefix+"/signup", utils.MethodHandler("POST", controllers.SignUpHandler))
}
