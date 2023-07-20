package data

import "time"

type User struct {
	Id                uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	Username          string    `json:"username" gorm:"not null;unique"`
	Password          string    `json:"password" gorm:"not null"`
	FirstName         string    `json:"first_name" gorm:"not null"`
	LastName          string    `json:"last_name" gorm:"not null"`
	Email             string    `json:"email" gorm:"not null;unique"`
	PhoneNumber       string    `json:"phone_number" gorm:"not null"`
	IsBlocked         bool      `json:"is_blocked" gorm:"default:false"`
	LoginAttemptCount int       `json:"login_attempt_count" gorm:"default:0"`
	CreatedAt         time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt         time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type Verification struct {
	Id                   uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID               uint64    `json:"user_id" gorm:"references:users(id);not null"`
	VerificationCodeHash string    `json:"verification_code_hash" gorm:"not null"`
	CreatedAt            time.Time `json:"created_at" gorm:"autoCreateTime"`
}
