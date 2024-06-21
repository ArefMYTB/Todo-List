package routes

import (
	"Todo-List/controllers"
	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", controllers.Show)
	router.HandleFunc("/add", controllers.Add).Methods("POST")
	//router.HandleFunc("/update/{id}", controllers.Update).Methods("PUT")
	router.HandleFunc("/delete/{id}", controllers.Delete)
	router.HandleFunc("/complete/{id}", controllers.Complete)

	return router
}
