package models

type Book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Abstract string `json:"abstract"`
}

const BookMapping = `{
	"settings": {
		"number_of_shards": 1,
		"number_of_replicas": 1
	},
	"mappings": {
		"properties": {
			"name": {
				"type": "text"
			},
			"author": {
				"type": "text"
			},
			"Abstract": {
				"type": "text"
			}
		}
	}
}`
