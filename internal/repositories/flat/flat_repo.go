package flat

import (
	"apartment_search_service/internal/models"
	openapi "apartment_search_service/internal/openapi/gen"
	"apartment_search_service/internal/repositories"
	"database/sql"
	"fmt"
)

type flatRepository struct {
	db *sql.DB
}

func NewFlatRepository(db *sql.DB) repositories.FlatRepository {
	return &flatRepository{db: db}
}

func (f *flatRepository) Insert(flat *models.Flat) error {
	query := `INSERT INTO flats (house_id, price, rooms, status) VALUES ($1, $2, $3, $4) RETURNING id`
	err := f.db.QueryRow(query, flat.HouseId, flat.Price, flat.Rooms, flat.Status).
		Scan(&flat.Id)
	return err
}

func (f *flatRepository) UpdateStatus(flatId int32, moderId, status string) (*models.Flat, error) {
	flat := &models.Flat{}
	query := `
		UPDATE flats SET status = $1, moderator_id = $2
		WHERE id = $3 
		AND ((status = $4 AND moderator_id = $2) -- end moderation for current moderator
		OR (status = $5 OR status != $4))        -- start new moderation
		RETURNING id, house_id, price, rooms, status, moderator_id`
	err := f.db.QueryRow(query, status, moderId, flatId, openapi.ON_MODERATION, openapi.CREATED).
		Scan(&flat.Id, &flat.HouseId, &flat.Price, &flat.Rooms, &flat.Status, &flat.ModeratorId)
	if err != nil {
		return nil, err
	}
	return flat, nil
}

func (f *flatRepository) GetFlatsByHouseId(houseId int32, isModer bool) ([]*models.Flat, error) {
	query := `SELECT id, house_id, price, rooms, status FROM flats WHERE house_id=$1`
	if !isModer {
		query = fmt.Sprintf(`%s AND status = '%s'`, query, openapi.APPROVED)
	}

	rows, err := f.db.Query(query, houseId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	flats := make([]*models.Flat, 0, 32)
	for rows.Next() {
		flat := &models.Flat{}
		err := rows.Scan(&flat.Id, &flat.HouseId, &flat.Price, &flat.Rooms, &flat.Status)
		if err != nil {
			return nil, err
		}
		flats = append(flats, flat)
	}

	return flats, nil
}
