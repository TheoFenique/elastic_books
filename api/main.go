package main

import (
	"elastic_books/api/config"
	"elastic_books/api/controllers"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	config.InitDB()
	r.HandleFunc("/book", controllers.PostBook).Methods("POST")
	r.HandleFunc("/book", controllers.SearchBooks).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func simpleRoot(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Working")
}
