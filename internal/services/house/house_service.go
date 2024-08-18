package house

import (
	"apartment_search_service/internal/models"
	"apartment_search_service/internal/repositories"
	"apartment_search_service/internal/services"
)

type houseService struct {
	flatRepo  repositories.FlatRepository
	houseRepo repositories.HouseRepository
}

func NewHouseService(h repositories.HouseRepository, f repositories.FlatRepository) services.HouseService {
	return &houseService{
		flatRepo:  f,
		houseRepo: h,
	}
}

func (h *houseService) CreateHouse(house *models.House) error {
	return h.houseRepo.Insert(house)
}

func (h *houseService) GetFlatsByHouseId(houseId int32, isModer bool) ([]*models.Flat, error) {
	_, err := h.houseRepo.GetById(houseId)
	if err != nil {
		return nil, err
	}
	return h.flatRepo.GetFlatsByHouseId(houseId, isModer)
}
