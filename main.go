package main

import (
	"bytes"
	"elastic-logging/config"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type LogElastic struct {
	Level     string `json:"level"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
	Service   string `json:"service"`
}

type elasticInterface interface {
	logToElasticSearch(level string, message string, service string) error
}
type ElasticService struct {
	Client *elasticsearch.Client
}

func GetEsClient() *ElasticService {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating elasticsearch Client: %+v\n", err)
	}
	_, err = es.Info()
	if err != nil {
		log.Fatalf("Error connecting to elasticsearch Client %+v\n", err)
	}
	return &ElasticService{Client: es}
}

func (es *ElasticService) LogToElasticSearch(level string, message string, service string) (*esapi.Response, error) {
	if es.Client == nil {
		log.Fatalf("Elasticsearch client not initialized")
	}

	logMessage := LogElastic{
		level,
		message,
		time.Now().Format(time.RFC3339),
		service,
	}
	jsonLog, err := json.Marshal(logMessage)
	if err != nil {
		log.Printf("Error Marshalling json: %+v\n", err)
		return nil, err
	}
	res, err := es.Client.Index(config.IndexName, bytes.NewReader(jsonLog))
	if err != nil {
		log.Printf("Error indexing log message: %+v\n", err)
		return nil, err
	}
	return res, err
}
func main() {
	fmt.Printf("Starting main program\n")
	es := GetEsClient()
	es.LogToElasticSearch("info", "started main program", "elastic_logger")
	for i := range []int{1, 2, 3, 4, 5, 6} {
		es.LogToElasticSearch("critical", fmt.Sprintf("Printing to main number %d", i), "elastic_logger")
	}
	es.LogToElasticSearch("weird", "ending program", "elastic_logger")
	return
}
