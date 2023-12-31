package scripts

import (
	"cronitor/internal/data"
	"cronitor/internal/messages/producer"
	"log"
	"time"
)

// DeleteExpiredVerifications is a function that deletes expired verifications
func DeleteExpiredVerifications() {

	esLog := make(map[string]interface{})
	esLog["status"] = "success"
	esLog["tag"] = "[ExpiredVerificationsDeleted]"
	esLog["message"] = "Expired verifications deleted!"
	esLog["timestamp"] = time.Now()

	db := data.InitPostgresDB()
	defer data.ClosePostgresDB(db)

	// Before two hours
	expirationTime := time.Now().Add(-2 * time.Hour)

	result := db.Exec("DELETE FROM verifications WHERE created_at < ?", expirationTime)
	if result.Error != nil {
		log.Default().Println("Error deleting expired verifications:", result.Error)
		esLog["error"] = result.Error.Error()
		esLog["status"] = "failed"
		return
	}

	producer.PublishElasticLogMessage(producer.ElasticLogMessage{
		Index: "cronitor",
		Data:  esLog,
	})
}
