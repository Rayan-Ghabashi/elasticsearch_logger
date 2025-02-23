package models

import (
	"github.com/elastic/go-elasticsearch/v8"
)

type LogElastic struct {
	Level     string `json:"level"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
	Service   string `json:"service"`
}

type JsonFilter struct {
	Level     *string `json:"level,omitempty"`
	Message   *string `json:"message,omitempty"`
	Timestamp *string `json:"timestamp,omitempty"`
	Service   *string `json:"service,omitempty"`
}

type elasticInterface interface {
	logToElasticSearch(level string, message string, service string) error
}
type ElasticService struct {
	Client *elasticsearch.Client
}
