package routers

import (
	"encoding/json"
	"net/http"

	"github.com/AVVKavvk/mongoWithGO/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter();
	router.HandleFunc("/",serveHome).Methods("GET");
	router.HandleFunc("/api/movies",controllers.GetAllMovies).Methods("GET");
	router.HandleFunc("/api/movie",controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}",controllers.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}",controllers.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/api/movie",controllers.DeleteAllMovie).Methods("DELETE")

	return router;
}

func serveHome(w http.ResponseWriter,r *http.Request){
	
	json.NewEncoder(w).Encode("welcome to home page of netflix")
}