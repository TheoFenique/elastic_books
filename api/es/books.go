package es

import (
	"context"
	"elastic_books/api/config"
	"elastic_books/api/models"
	"encoding/json"
	"errors"
	"github.com/olivere/elastic/v7"
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

func SearchBooks(ctx context.Context, key, value string) ([]models.Book, error) {
	searchSource := elastic.NewSearchSource()
	searchSource.Query(elastic.NewMatchQuery(key, value))

	searchRes, err := config.EsClient.Search().Index("books").SearchSource(searchSource).Do(ctx)
	if err != nil {
		log.Println("Error searching book")
		return nil, err
	}

	var books []models.Book
	for _, res := range searchRes.Hits.Hits {
		var book models.Book
		err := json.Unmarshal(res.Source, &book)
		if err != nil {
			return nil, errors.New("cannot decode book")
		}

		books = append(books, book)
	}

	return nil, nil
}
