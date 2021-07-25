package config

import (
	"context"
	"elastic_books/api/models"
	"github.com/olivere/elastic/v7"
	"log"
	"time"
)

var EsClient *elastic.Client

func InitDB() {
	ctx := context.Background()
	getESClient()

	exists, err := EsClient.IndexExists("books").Do(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if !exists {
		_, err := EsClient.CreateIndex("books").BodyString(models.BookMapping).Do(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func getESClient() {
	for i := 0; i < 3; i++ {
		log.Println("Trying to connect to ES...")
		client, err := elastic.NewClient(elastic.SetURL("http://es01:9200"),
			elastic.SetSniff(false),
			elastic.SetHealthcheck(false))
		if err != nil {
			log.Println("Failed to connect, retrying in 5 seconds...")
			time.Sleep(5 * time.Second)
			continue
		}
		EsClient = client
		log.Println("Connected to the DB!")
		return
	}
	log.Fatal("Cannot connect to DB")
}
