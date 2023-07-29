package elastic

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"loggerservice/internal/config"
	"loggerservice/internal/utils"
	"net/http"
	"strconv"
	"time"
)

type ElasticLog struct {
	Timestamp time.Time   `json:"timestamp"`
	Data      interface{} `json:"data"`
}

// SendLog2Elastic sends log to Elasticsearch
func SendLog2Elastic(data interface{}, index string) error {
	// Elasticsearch URL
	elasticURL := config.EnvConfigs.ElasticUrl + "/" + index + "/_doc"

	elasticLog := ElasticLog{
		Timestamp: time.Now(),
		Data:      data,
	}

	// Convert data to JSON
	requestJSON, err := json.Marshal(elasticLog)
	if err != nil {
		utils.LogErr("Error marshaling data to JSON:", err)
		return err
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", elasticURL, bytes.NewBuffer(requestJSON))
	if err != nil {
		utils.LogErr("Error creating HTTP request:", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	// Create HTTP client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		utils.LogErr("Error sending HTTP request:", err)
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			utils.LogErr("Error closing response body:", err)
		}
	}(resp.Body)

	// Check response status code
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		utils.LogInfo("Log sent to Elasticsearch successfully.")
	} else {
		utils.LogInfo("Error sending log to Elasticsearch. Status Code:")
		return errors.New("Error sending log to Elasticsearch. Status Code:" + strconv.Itoa(resp.StatusCode))
	}

	return nil
}
