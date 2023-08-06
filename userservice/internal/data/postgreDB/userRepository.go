package postgreDB

import (
	"log"
)

// InsertUser is a function that inserts a user into the database
func InsertUser(user User) (uint64, error) {

	db := InitDB()
	defer Close(db)

	result := db.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}

	log.Default().Println("User created successfully:", user.Id)

	return user.Id, nil
}

// GetUserWithEmail is a function that returns a user with the given username
func GetUserWithEmail(email string) (User, error) {

	db := InitDB()
	defer Close(db)

	var user User

	result := db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

// GetUserWithId is a function that returns a user with the given id
func GetUserWithId(id uint64) (User, error) {

	db := InitDB()
	defer Close(db)

	var user User

	result := db.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

// IncrementLoginAttemptCount is a function that increments the login attempt count of a user
func IncrementLoginAttemptCount(user *User) error {

	db := InitDB()
	defer Close(db)

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
func ResetLoginAttemptCount(user *User) error {

	db := InitDB()
	defer Close(db)

	user.LoginAttemptCount = 0

	result := db.Save(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// UpdateUser is a function that updates a user
func UpdateUser(user *User) error {

	db := InitDB()
	defer Close(db)

	result := db.Save(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
