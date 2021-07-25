package es

import (
	"context"
	"elastic_books/api/config"
	"elastic_books/api/models"
	"encoding/json"
	"errors"
	"log"
)

func prepareForES(value interface{}) (string, error) {
	b, err := json.Marshal(value)
	if err != nil {
		return "", errors.New("error marshal book")
	}

	return string(b), nil
}

func CreateBook(ctx context.Context, book *models.Book) error {
	s, err := prepareForES(book)
	if err != nil {
		log.Println("Cannot marshal book into bytes")
		return errors.New("cannot create book")
	}

	_, err = config.EsClient.Index().Index("books").BodyJson(s).Do(ctx)
	if err != nil {
		return errors.New("error insert book")
	}
	return nil
}
