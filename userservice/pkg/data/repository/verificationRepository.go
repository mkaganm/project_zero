package repository

import (
	"log"
	"userservice/pkg/data"
	"userservice/pkg/data/entity"
)

// InsertVerificationCode is a function that inserts a verification code into the database
func InsertVerificationCode(verification entity.Verification) (uint64, error) {

	db := data.Init()
	defer data.Close(db)

	result := db.Create(&verification)
	if result.Error != nil {
		return 0, result.Error
	}

	log.Default().Println("Verification code created successfully:", verification.Id)

	return verification.Id, nil
}

// GetVerificationCodeWithUserId is a function that gets a verification code with a user id
func GetVerificationCodeWithUserId(userId uint64) (entity.Verification, error) {

	db := data.Init()
	defer data.Close(db)

	var verification entity.Verification

	result := db.Where("user_id = ?", userId).First(&verification)
	if result.Error != nil {
		return verification, result.Error
	}

	return verification, nil
}

// DeleteVerificationWithId is a function that deletes a verification code with an id
func DeleteVerificationWithId(id uint64) error {

	db := data.Init()
	defer data.Close(db)

	result := db.Where("id = ?", id).Delete(&entity.Verification{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// DeleteVerificationWithUserId is a function that deletes a verification code with a user id
func DeleteVerificationWithUserId(userId uint64) error {

	db := data.Init()
	defer data.Close(db)

	result := db.Where("user_id = ?", userId).Delete(&entity.Verification{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
