package services

import (
	"apartment_search_service/internal/models"
	"apartment_search_service/internal/repositories"
)

type flatService struct {
	flatRepo  repositories.FlatRepository
	houseRepo repositories.HouseRepository
}

func NewFlatService(f repositories.FlatRepository, h repositories.HouseRepository) FlatService {
	return &flatService{
		flatRepo:  f,
		houseRepo: h,
	}
}

func (f *flatService) CreateFlat(flat *models.Flat) error {
	_, err := f.houseRepo.GetById(flat.HouseId)
	if err != nil {
		return err
	}
	return f.flatRepo.Insert(flat)
}

func (f *flatService) UpdateStatus(flatId int32, moderId, status string) (*models.Flat, error) {
	return f.flatRepo.UpdateStatus(flatId, moderId, status)
}
