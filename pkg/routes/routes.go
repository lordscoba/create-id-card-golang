package routes

import (
	"github.com/gorilla/mux"
	"github.com/lordscoba/create-id-card-golang/pkg/controllers"
)

var RegisterCardRoutes = func(router *mux.Router){
	router.HandleFunc("/create", controllers.CreateCard).Methods("POST")
	router.HandleFunc("/", controllers.GetCard).Methods("GET")
	router.HandleFunc("/show/{cardId}", controllers.GetCardById).Methods("GET")
	router.HandleFunc("/update/{cardId}", controllers.UpdateCard).Methods("PUT")
	router.HandleFunc("/delete/{cardId}", controllers.DeleteCard).Methods("DELETE")
}
