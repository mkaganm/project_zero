package elastic

import (
	"bytes"
	"cronitor/internal/config"
	"encoding/json"
	"log"
	"net/http"
)

// SendLog is a function that sends log to elastic
func SendLog(data interface{}) {

	url := config.EnvConfigs.LoggerElasticUrl

	dataReq := make(map[string]interface{})
	dataReq["index"] = "cronitor"
	dataReq["data"] = data

	reqJson, _ := json.Marshal(dataReq)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(reqJson))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, _ := client.Do(req)
	_ = resp.Body.Close() // close body when send request we don't need to read the response

	log.Default().Println("Log sent successfully!")

}
