package services

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/direwen/flashpaper/internal/models"
	"github.com/direwen/flashpaper/pkg/utils"
)

type AuthService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{
		db: db,
	}
}

func (s *AuthService) RegisterUser(ctx context.Context, email, password string) error {

	//Checking if there's already a user with these credentials
	var existingUser models.User
	if err := s.db.WithContext(ctx).Where("email = ?", email).First(&existingUser).Error; err == nil {
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

	if err := s.db.WithContext(ctx).Create(&user).Error; err != nil {
		return err
	}

	return nil

}

func (s *AuthService) LoginUser(ctx context.Context, email, password string) (string, error) {
	var user models.User

	// Find the user record and populate "user" variable with all data included in that found record
	if err := s.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return "", errors.New("invalid credentials")
	}

	// Compared hashed user password and the provided password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Generate jwt token
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil

}
