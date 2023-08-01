package logger

import (
	"bytes"
	"encoding/json"
	"log"
	"mailerservice/internal/config"
	"mailerservice/internal/utils"
	"net/http"
)

// SendElasticLog sends log to elastic
func SendElasticLog(data interface{}) {

	url := config.EnvConfigs.LoggerElasticUrl

	dataReq := map[string]interface{}{}

	dataReq["index"] = "mailerservice"
	dataReq["data"] = data

	reqJson, err := json.Marshal(dataReq)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqJson))
	utils.LogErr("error creating request", err)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	_ = resp.Body.Close() // close body when send request we don't need to read the response
	utils.LogErr("error sending request", err)

	log.Default().Println("Log sent successfully!")

}
