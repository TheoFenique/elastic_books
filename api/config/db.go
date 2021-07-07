package config

import (
	"github.com/elastic/go-elasticsearch/v7"
	"log"
)

func InitDB() elasticsearch.Client{
	client, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Cannot connect to ES: %s", err)
	}
	return *client
}
