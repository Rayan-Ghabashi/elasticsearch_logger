package main

import (
	"elastic-logging/models"
	"fmt"
	"log"

	"elastic-logging/internal"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type LogElastic struct {
	Level     string `json:"level"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
	Service   string `json:"service"`
}

func GetEsClient() *models.ElasticService {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating elasticsearch Client: %+v\n", err)
	}
	_, err = es.Info()
	if err != nil {
		log.Fatalf("Error connecting to elasticsearch Client %+v\n", err)
	}
	return &models.ElasticService{Client: es}
}

func LogToElasticSearch(es *models.ElasticService, level string, message string, service string) (*esapi.Response, error) {
	return internal.LogToElasticSearch(es, level, message, service)
}

func main() {
	fmt.Printf("Starting main program\n")
	es := GetEsClient()
	internal.LogToElasticSearch(es, "info", "started main program", "elastic_logger")
	for i := range []int{1, 2, 3, 4, 5, 6} {
		internal.LogToElasticSearch(es, "critical", fmt.Sprintf("Printing to main number %d", i), "elastic_logger")
	}
	internal.LogToElasticSearch(es, "weird", "ending program", "elastic_logger")
	return
}
