package main

import (
	"elastic_books/api/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	es := config.InitDB()
	log.Println(es.Info())
	r.HandleFunc("/", simpleRoot)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func simpleRoot(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Working")
}
