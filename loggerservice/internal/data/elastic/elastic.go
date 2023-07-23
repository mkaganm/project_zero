package elastic

import (
	"bytes"
	"encoding/json"
	"loggerservice/internal/config"
	"loggerservice/internal/utils"
	"net/http"
)

// SendLog2Elastic sends log to Elasticsearch
func SendLog2Elastic(data interface{}) {

	// Convert data to JSON
	requestJSON, err := json.Marshal(data)
	utils.LogErr("Error marshaling request to JSON:", err)

	// Send log to Elasticsearch
	_, err = http.Post(config.EnvConfigs.ElasticUrl, "application/json", bytes.NewBuffer(requestJSON))
	utils.LogErr("Error sending log to Elasticsearch:", err)

}
