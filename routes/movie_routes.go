package router

import (
	"log"

	"github.com/gorilla/mux"
	controller "github.com/maazmaahi/mongoapi/controllers"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	log.Println("Router invoked ...")
	router.HandleFunc("/api/movies", controller.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.DeleteAMovie).Methods("DELETE")
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("DELETE")

	return router
}
