package internal

import (
	"bytes"
	"github.com/rayan-ghabashi/elasticsearch_logger/config"
	"github.com/rayan-ghabashi/elasticsearch_logger/models"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
	"time"
)

func LogToElasticSearch(es *models.ElasticService, level string, message string, service string) (*esapi.Response, error) {
	if es.Client == nil {
		log.Fatalf("Elasticsearch client not initialized")
	}

	logMessage := models.LogElastic{
		Level:     level,
		Message:   message,
		Timestamp: time.Now().Format(time.RFC3339),
		Service:   service,
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
