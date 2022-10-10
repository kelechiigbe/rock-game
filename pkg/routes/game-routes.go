package routes

import (
	"github.com/gorilla/mux"
	"github.com/kelechiigbe/rock-game/pkg/controllers"
)

var RegisterGameRoutes = func(router *mux.Router){
	router.HandleFunc("/choices", controllers.GetAllChoices).Methods("GET")
	router.HandleFunc("/choice", controllers.GetRandomChoice).Methods("GET")
	router.HandleFunc("/choices/{choice_id}", controllers.GetChoiceById).Methods("GET")
	router.HandleFunc("/play", controllers.PlayGame).Methods("POST","OPTIONS")
}