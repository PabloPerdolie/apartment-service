package repositories

import "apartment_search_service/internal/models"

type UserRepository interface {
	Insert(user *models.User) error
	GetById(id string) (*models.User, error)
	Get(email, password string) (*models.User, error)
}
