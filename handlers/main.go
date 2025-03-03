package handlers

import (
	"github.com/rayan-ghabashi/elasticsearch_logger/models"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/elastic/go-elasticsearch/v8"
)

func getLogsHandler(es *elasticsearch.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jsonBody, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading json body: %+v\n", err)
			return
		}
		defer r.Body.Close()

		var filter models.JsonFilter
		json.Unmarshal(jsonBody, &filter)

	}
}
