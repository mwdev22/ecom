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

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) UserStore {
	return UserStore{db: db}
}

func (s *UserStore) CreateUser(payload *types.RegisterUserPayload) error {
	result, err := s.GetUserByEmail(payload.Email)
	if err == nil {
		return fmt.Errorf("user with email %s already exists", result.Email)
	} else if err != gorm.ErrRecordNotFound {
		return fmt.Errorf("error checking existing user: %w", err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}
	user := User{
		Email:     payload.Email,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Password:  string(hashedPassword),
	}

	if err := s.db.Create(&user).Error; err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (s *UserStore) GetUserByEmail(email string) (*User, error) {
	var user User
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
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
