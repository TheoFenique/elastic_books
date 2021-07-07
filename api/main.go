package main

import (
	"api/elastic_books/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main () {
	r := mux.NewRouter()
	r.HandleFunc("/", simpleRoot)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func simpleRoot(w http.ResponseWriter, r *http.Request) {
	es := config.InitDB()
	log.Println(es.Info())
	json.NewEncoder(w).Encode("Working")
}


