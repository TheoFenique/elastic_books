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

	res, err := es.CreateBook(ctx, &book)
	if err != nil {
		log.Printf("Error create book: %s", err.Error())
	}

	json.NewEncoder(w).Encode(fmt.Sprintf("Sucessful insertion: %v", res))
}

func SearchBooks(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	key := "title"
	value := "title"

	books, err := es.SearchBooks(ctx, key, value)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}

	json.NewEncoder(w).Encode(books)
}
