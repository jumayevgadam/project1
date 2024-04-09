package service

import (
	"Project1/internal/users/model"
	"Project1/internal/users/repository"
	"crypto/sha1"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

var (
	salt       = "sdfsbsbh34hj34h"
	signingKey = []byte("###%5645646566")
)

type tokenClaims struct {
	jwt.StandardClaims

	UserId int `json:"user_id"`
}

func (s *UserService) CreateUser(user model.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *UserService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})

	return token.SignedString(signingKey)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *UserService) DeleteUser(userId int) error {
	return s.repo.DeleteUser(userId)
}
