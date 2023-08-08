package mail

import (
	"log"
	"mailerservice/internal/clients/logger"
	"mailerservice/internal/config"
	"net/smtp"
	"strconv"
	"time"
)

type Mail struct {
	From    string   `json:"from"`
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Body    string   `json:"body"`
}

// SendMail is a function to send email
func SendMail(mail Mail) error {

	esLog := make(map[string]interface{})
	esLog["status"] = "success"
	esLog["tag"] = "[MailSent]"
	esLog["message"] = "Mail sent!"
	esLog["timestamp"] = time.Now()

	smtpHost := config.EnvConfigs.SmtpHost
	smtpPort := config.EnvConfigs.SmtpPort
	username := config.EnvConfigs.MailerSenderAddress
	password := config.EnvConfigs.MailerSenderPassword
	auth := smtp.PlainAuth("", username, password, smtpHost)

	message := []byte("Subject: " + mail.Subject + "\r\n" + "\r\n" + mail.Body)

	err := smtp.SendMail(smtpHost+":"+strconv.Itoa(smtpPort), auth, mail.From, mail.To, message)
	if err != nil {
		log.Default().Println("Error while sending email: " + err.Error())
		esLog["error"] = err.Error()
		esLog["status"] = "failed"
		return err
	}

	log.Default().Println("Email sent successfully!")
	logger.SendElasticLog(esLog)
	return nil

}
