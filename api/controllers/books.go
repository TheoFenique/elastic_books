package controllers

import (
	"encoding/json"
	"net/http"
)

func PostBook(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("GetBooks")
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("GetBooks")
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("GetBooks")
}


