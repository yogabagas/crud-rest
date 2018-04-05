package routers

import (
	"day14gaji/controllers"

	"github.com/gorilla/mux"
)

func setGajiRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/gaji", controllers.GetGaji).Methods("GET")
	return router
}
