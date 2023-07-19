package repository

import (
	"log"
	"userservice/pkg/data"
	"userservice/pkg/data/entity"
)

// InsertUser is a function that inserts a user into the database
func InsertUser(user entity.User) (uint64, error) {

	db := data.Init()
	defer data.Close(db)

	result := db.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}

	log.Default().Println("User created successfully:", user.Id)

	return user.Id, nil
}

// GetUserWithEmail is a function that returns a user with the given username
func GetUserWithEmail(email string) (entity.User, error) {

	db := data.Init()
	defer data.Close(db)

	var user entity.User

	result := db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

// GetUserWithId is a function that returns a user with the given id
func GetUserWithId(id uint64) (entity.User, error) {

	db := data.Init()
	defer data.Close(db)

	var user entity.User

	result := db.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

// IncrementLoginAttemptCount is a function that increments the login attempt count of a user
func IncrementLoginAttemptCount(user *entity.User) error {

	db := data.Init()
	defer data.Close(db)

	user.LoginAttemptCount++

	if user.LoginAttemptCount >= 3 {
		user.IsBlocked = true
	}

	result := db.Save(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// ResetLoginAttemptCount is a function that resets the login attempt count of a user
func ResetLoginAttemptCount(user *entity.User) error {

	db := data.Init()
	defer data.Close(db)

	user.LoginAttemptCount = 0

	result := db.Save(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// UpdateUser is a function that updates a user
func UpdateUser(user *entity.User) error {

	db := data.Init()
	defer data.Close(db)

	result := db.Save(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
