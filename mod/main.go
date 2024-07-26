package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("mod details");
	r:=mux.NewRouter();
	r.HandleFunc("/",serveHome).Methods("GET")
	log.Fatal(http.ListenAndServe(":4006",r))
}

func serveHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Header 1</h1><h1>Header 2</h1><h1>Header 3</h1>"))
}