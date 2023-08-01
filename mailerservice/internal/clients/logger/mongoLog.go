package logger

import (
	"bytes"
	"encoding/json"
	"log"
	"mailerservice/internal/config"
	"mailerservice/internal/utils"
	"net/http"
)

type Log struct {
	Collection     string      `json:"collection"`
	Source         string      `json:"source"`
	Method         string      `json:"method"`
	Request        interface{} `json:"request"`
	RequestHeader  interface{} `json:"request_header"`
	Response       interface{} `json:"response"`
	ResponseHeader interface{} `json:"response_header"`
	Duration       string      `json:"duration"`
	Status         int         `json:"status"`
}

// SendMongoLog sends log to mongo
func SendMongoLog(successLog Log) {

	url := config.EnvConfigs.LoggerMongoUrl
	successLog.Collection = "mailerservice"

	data, err := json.Marshal(successLog)
	utils.LogErr("error marshalling log", err)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	utils.LogErr("error creating request", err)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	_ = resp.Body.Close() // close body when send request we don't need to read the response
	utils.LogErr("error sending request", err)

	log.Default().Println("Log sent successfully!")

}
