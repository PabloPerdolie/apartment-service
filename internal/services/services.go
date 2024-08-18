package services

import "apartment_search_service/internal/models"

type AuthService interface {
	CreateUser(user *models.User) (id string, err error)
	GetUser(id string) (*models.User, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (id string, err error)
}

type HouseService interface {
	CreateHouse(house *models.House) error
	GetFlatsByHouseId(houseId int32, isModer bool) ([]*models.Flat, error)
	// SubscribeHouse(id string)
}

type FlatService interface {
	CreateFlat(flat *models.Flat) error
	UpdateStatus(flatId int32, moderId, status string) (*models.Flat, error)
}
