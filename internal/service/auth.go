package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/martynove/gophermart/internal/models"
	"github.com/martynove/gophermart/internal/repository"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GetUserByLogin(user models.User) (bool, error) {
	return s.repo.GetUserByLogin(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(models.PasswordSalt)))
}
