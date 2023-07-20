package mailer

import (
	"bytes"
	"encoding/json"
	"io"
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
func SendMail(request SendMailRequest) error {

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
	utils.LogErr("Error while sending request: ", err)

	// Close request body
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			utils.LogErr("Error while closing request body: ", err)
			return
		}
	}(resp.Body)

	log.Default().Println("Email sent successfully!")

	return nil

}
