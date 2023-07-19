package scripts

import (
	"cronitor/pkg/data"
	"fmt"
	"log"
	"time"
)

// ResetLoginAttemptCounts is a function that resets login attempt counts
func ResetLoginAttemptCounts() {
	db := data.Init()
	defer data.Close(db)

	// Before two hours
	twoHoursAgo := time.Now().Add(-2 * time.Hour)

	query := fmt.Sprintf(
		"UPDATE users SET login_attempt_count = 0 WHERE login_attempt_count > 0 AND updated_at < '%s'",
		twoHoursAgo.Format("2006-01-02 15:04:05"))

	result := db.Exec(query)
	if result.Error != nil {
		log.Default().Println("Error while resetting login attempt counts:", result.Error)
		return
	}

	log.Default().Println("Login attempt counts reset.")
}
