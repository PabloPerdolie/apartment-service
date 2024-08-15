package services

import (
	"apartment_search_service/internal/models"
	"apartment_search_service/internal/repositories"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	salt       = "jdfgn45n64b3h45b3hb4"
	signingKey = "5h6b76h7534i5f3ga7i6"
	tokenTTL   = 1 * time.Hour
)

type authService struct {
	repo repositories.UserRepository
}

func NewAuthService(repo repositories.UserRepository) AuthService {
	return &authService{repo: repo}
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId string `json:"user_id"`
}

func (as *authService) CreateUser(user *models.User) (id string, err error) {
	user.Password = HashPassword(user.Password)
	_, err = as.repo.Get(user.Email, user.Password)
	if err == nil {
		return id, errors.New("email is not free")
	}
	err = as.repo.Insert(user)
	if err != nil {
		return id, err
	}
	return user.UserId, err
}

func (as *authService) GetUser(id string) (*models.User, error) {
	return as.repo.GetById(id)
}

func (as *authService) GenerateToken(id, password string) (string, error) {
	_, err := as.repo.Get(id, HashPassword(password))
	if err != nil {
		return "", err
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
		},
		id,
	}).SignedString([]byte(signingKey))
	return token, err
}

func (as *authService) ParseToken(accessToken string) (id string, err error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return id, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return id, errors.New("token claims not the *tokenClaims")
	}
	return claims.UserId, err
}

func HashPassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
