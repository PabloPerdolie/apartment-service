package repositories

import "apartment_search_service/internal/models"

type UserRepository interface {
	Insert(user *models.User) error
	GetById(id string) (*models.User, error)
	Get(email, password string) (*models.User, error)
}

type HouseRepository interface {
	Insert(house *models.House) error
	GetById(id int32) (*models.House, error)
}

type FlatRepository interface {
	Insert(flat *models.Flat) error
	UpdateStatus(flatId int32, moderId, status string) (*models.Flat, error)
	GetFlatsByHouseId(houseId int32, isModer bool) ([]*models.Flat, error)
}

type SubscriptionRepository interface {
	AddSubscriber(houseId int32, email string) error
	GetSubscribers(houseId int32) ([]string, error)
}
