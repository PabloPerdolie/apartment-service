package services

import "apartment_search_service/internal/models"

type AuthService interface {
	CreateUser(user *models.User) (id string, err error)
	GetUser(id string) (*models.User, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (id string, err error)
}
