package services

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/direwen/flashpaper/internal/config"
	"github.com/direwen/flashpaper/internal/models"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) RegisterUser(email, password string) error {
	db := config.GetDB()

	//Checking if there's already a user with these credentials
	var existingUser models.User
	if err := db.Where("email = ?", email).First(&existingUser).Error; err == nil {
		return errors.New("user with this email already exists")
	}

	//Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	//Create the user
	user := models.User{
		Email:    email,
		Password: string(hashedPassword),
	}

	if err := db.Create(&user).Error; err != nil {
		return err
	}

	return nil

}
