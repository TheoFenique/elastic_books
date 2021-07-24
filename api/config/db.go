package config

import (
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esutil"
	"log"
	"time"
)

func InitDB() *elasticsearch.Client {
	for i := 0; i < 3; i++ {
		log.Println("Trying to connect to ES...")
		cfg := elasticsearch.Config{
			Addresses: []string{
				"http://elasticsearch:9200/",
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

func CreateBulkIndexer(client *elasticsearch.Client) *esutil.BulkIndexer {
	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:         "books",          // The default index name
		Client:        client,           // The Elasticsearch client
		NumWorkers:    numWorkers,       // The number of worker goroutines
		FlushBytes:    int(flushBytes),  // The flush threshold in bytes
		FlushInterval: 30 * time.Second, // The periodic flush interval
	})
	if err != nil {
		log.Fatalf("Error creating the indexer: %s", err)
	}
	return bi
}
