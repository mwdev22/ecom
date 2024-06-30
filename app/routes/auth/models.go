package auth

import (
	"fmt"

	"github.com/mwdev22/ecom/app/types"
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

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) Store {
	return Store{db: db}
}

func (s *Store) CreateUser(user *types.RegisterUserPayload) error {
	var existingUser User
	result := s.db.Where("email = ?", user.Email).First(&existingUser)
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

	if err := s.db.Create(user).Error; err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (s *Store) GetUserByEmail(email string) (*User, error) {
	var user User
	if err := s.db.First(&user, email).Error; err != nil {
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
