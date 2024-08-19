package postgres

import (
	"apartment_search_service/internal/models"
	"apartment_search_service/internal/repositories"
	"database/sql"
)

type houseRepository struct {
	db *sql.DB
}

func NewHouseRepository(db *sql.DB) repositories.HouseRepository {
	return &houseRepository{db: db}
}

func (hr *houseRepository) Insert(house *models.House) error {
	query := `INSERT INTO houses (address, year, developer, created_at) VALUES ($1, $2, $3, $4) RETURNING id`
	err := hr.db.QueryRow(query, house.Address, house.Year, house.Developer, house.CreatedAt).
		Scan(&house.Id)
	return err

}

func (hr *houseRepository) GetById(id int32) (*models.House, error) {
	house := &models.House{}
	query := `SELECT id, address, year, developer, created_at, updated_at FROM houses WHERE id = $1`
	if err := hr.db.QueryRow(query, id).
		Scan(&house.Id, &house.Address, &house.Year, &house.Developer, &house.CreatedAt, &house.UpdatedAt); err != nil {
		return nil, err
	}
	return house, nil
}
