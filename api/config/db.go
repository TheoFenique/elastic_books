package config

import (
	"github.com/elastic/go-elasticsearch/v7"
	"log"
	"time"
)

func InitDB() *elasticsearch.Client{
	for i := 0; i < 3; i++ {
		log.Println("Trying to connect to ES...")
		cfg := elasticsearch.Config{
			Addresses: []string{
				"http://elasticsearch:9200",
			},
		}
		client, err := elasticsearch.NewClient(cfg)
		if err != nil {
			log.Println("Failed to connect, retrying in 15 seconds...")
			time.Sleep(15 * time.Second)
			continue
		}
		log.Println("Connected to the DB!")
		return client
	}
	return nil
}
