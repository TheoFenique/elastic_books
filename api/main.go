package main

import (
	"api/elastic_books/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main () {
	es := config.InitDB()
	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	defer res.Body.Close()
	if res.IsError() {
		log.Fatalf("Error: %s", res.String())
	}

	log.Println("IT IS WORKING SOMEHOW")

	r := mux.NewRouter()
	r.HandleFunc("/", simpleRoot)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func simpleRoot(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Working")
}


