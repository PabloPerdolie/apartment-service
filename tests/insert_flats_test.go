package tests

import (
	"apartment_search_service/internal/models"
	"apartment_search_service/internal/repositories/postgres"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertFlat(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	flatRepo := postgres.NewFlatRepository(db)

	flat := &models.Flat{
		HouseId: 1,
		Price:   100000,
		Rooms:   3,
		Status:  "created",
	}

	mock.ExpectQuery(`INSERT INTO flats \(house_id, price, rooms, status\) VALUES \(\$1, \$2, \$3, \$4\) RETURNING id`).
		WithArgs(flat.HouseId, flat.Price, flat.Rooms, flat.Status).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	err = flatRepo.Insert(flat)

	assert.NoError(t, err)
	assert.Equal(t, int32(1), flat.Id)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestInsertFlat_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	flatRepo := postgres.NewFlatRepository(db)

	flat := &models.Flat{
		HouseId: 1,
		Price:   100000,
		Rooms:   3,
		Status:  "created",
	}

	mock.ExpectQuery(`INSERT INTO flats \(house_id, price, rooms, status\) VALUES \(\$1, \$2, \$3, \$4\) RETURNING id`).
		WithArgs(flat.HouseId, flat.Price, flat.Rooms, flat.Status).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	err = flatRepo.Insert(flat)

	assert.NoError(t, err)
	assert.Equal(t, int32(1), flat.Id)
	assert.NoError(t, mock.ExpectationsWereMet())
}
