package controllers

import (
	"context"
	"elastic_books/api/es"
	"elastic_books/api/helpers"
	"elastic_books/api/models"
	"encoding/json"
	"errors"
	"net/http"
)

func PostBook(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		helpers.EncodeError(w, err)
		return
	}

	_, err = es.CreateBook(ctx, &book)
	if err != nil {
		helpers.EncodeError(w, err)
		return
	}

	helpers.EncodeData(w, book)
}

func SearchBooks(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	key := r.URL.Query().Get("key")
	value := r.URL.Query().Get("key")

	if key != "" && value != "" {
		books, err := es.SearchBooks(ctx, key, value)
		if err != nil {
			helpers.EncodeError(w, err)
			return
		}

		helpers.EncodeData(w, books)
	} else {
		helpers.EncodeError(w, errors.New("need 'key' and 'value' GET parameters"))
	}

}
