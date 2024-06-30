package auth

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string `gorm:"unique"`
	Password  string
}

func CreateUser(db *gorm.DB, user *User) error {
	var existingUser User
	result := db.Where("email = ?", user.Email).First(&existingUser)
	if result.Error == nil {
		return fmt.Errorf("user with email %s already exists", user.Email)
	} else if result.Error != gorm.ErrRecordNotFound {
		return fmt.Errorf("error checking existing user: %w", result.Error)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}
	user.Password = string(hashedPassword)

	if err := db.Create(user).Error; err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func GetUser(db *gorm.DB, id uint) (*User, error) {
	var user User
	if err := db.First(&user, id).Error; err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return &user, nil
}

func UpdateUser(db *gorm.DB, user *User) error {
	if err := db.Save(user).Error; err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

func DeleteUser(db *gorm.DB, id uint) error {
	if err := db.Delete(&User{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}
