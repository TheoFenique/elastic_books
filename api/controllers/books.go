package controllers

import (
	"context"
	"elastic_books/api/es"
	"elastic_books/api/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func PostBook(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		log.Printf("Err decode body: %s", err.Error())
		json.NewEncoder(w).Encode("Error decoding body")
	}

	err = es.CreateBook(ctx, &book)
	if err != nil {
		log.Printf("Error create book: %s", err.Error())
	}

	json.NewEncoder(w).Encode(fmt.Sprintf("Sucessful insertion: %v", book))
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("GetBooks")
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("GetBooks")
}
