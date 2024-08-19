package tests

import (
	"apartment_search_service/internal/repositories/postgres"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetFlatsByHouseIdClient(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	flatRepo := postgres.NewFlatRepository(db)
	houseId := int32(1)

	rows := sqlmock.NewRows([]string{"id", "house_id", "price", "rooms", "status"}).
		AddRow(1, houseId, 100000, 3, "approved").
		AddRow(2, houseId, 150000, 4, "approved")

	mock.ExpectQuery(`SELECT id, house_id, price, rooms, status FROM flats WHERE house_id=\$1? AND status = 'approved'`).
		WithArgs(houseId).
		WillReturnRows(rows)

	flats, err := flatRepo.GetFlatsByHouseId(houseId, false)
	assert.NoError(t, err)
	assert.Len(t, flats, 2)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestGetFlatsByHouseIdModer(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	flatRepo := postgres.NewFlatRepository(db)
	houseId := int32(1)

	rows := sqlmock.NewRows([]string{"id", "house_id", "price", "rooms", "status"}).
		AddRow(1, houseId, 100000, 3, "approved").
		AddRow(2, houseId, 150000, 4, "approved")

	mock.ExpectQuery(`SELECT id, house_id, price, rooms, status FROM flats WHERE house_id=\$1?`).
		WithArgs(houseId).
		WillReturnRows(rows)

	flats, err := flatRepo.GetFlatsByHouseId(houseId, true)
	assert.NoError(t, err)
	assert.Len(t, flats, 2)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestGetFlatsByHouseIdQueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	flatRepo := postgres.NewFlatRepository(db)
	houseId := int32(1)

	mock.ExpectQuery(`SELECT id, house_id, price, rooms, status FROM flats WHERE house_id=\$1`).
		WithArgs(houseId).
		WillReturnError(fmt.Errorf("column does not exist"))

	flats, err := flatRepo.GetFlatsByHouseId(houseId, true)
	assert.Error(t, err)
	assert.Nil(t, flats)
	assert.NoError(t, mock.ExpectationsWereMet())
}
