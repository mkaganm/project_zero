package mailer

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"userservice/internal/config"
	"userservice/internal/utils"
)

// SendMailRequest is a struct that represents a request to send an email
type SendMailRequest struct {
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Body    string   `json:"body"`
}

// SendMail is a function that sends an email
func SendMail(request SendMailRequest) {

	// Marshal request
	url := config.EnvConfigs.MailerUrl
	data, err := json.Marshal(request)
	utils.LogErr("Error while marshalling request: ", err)

	// Create request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	utils.LogErr("Error while creating request: ", err)

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	_ = resp.Body.Close() // close body when send request we don't need to read the response
	utils.LogErr("Error while sending request: ", err)

	log.Default().Println("Email sent successfully!")

}
