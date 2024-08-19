package tests

import (
	errors1 "apartment_search_service/internal/errors"
	"apartment_search_service/internal/handlers/house"
	"apartment_search_service/internal/models"
	openapi "apartment_search_service/internal/openapi/gen"
	"apartment_search_service/internal/services/mocks"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetFlatsInDeHouse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockHouseService := mocks.NewMockHouseService(ctrl)

	logger := logrus.New()

	handler := house.NewHandler(mockHouseService, nil, logger)

	expectedFlats := []*models.Flat{
		{
			Id:      1,
			HouseId: 1,
			Price:   1000,
			Rooms:   2,
			Status:  string(openapi.APPROVED),
		},
		{
			Id:      2,
			HouseId: 1,
			Price:   2000,
			Rooms:   3,
			Status:  string(openapi.APPROVED),
		},
	}

	mockHouseService.EXPECT().GetFlatsByHouseId(int32(1), true).Return(expectedFlats, nil)

	req, err := http.NewRequest("GET", "/house/1", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rr := httptest.NewRecorder()

	ctx := context.WithValue(req.Context(), "role", string(openapi.MODERATOR))
	req = req.WithContext(ctx)

	router := mux.NewRouter()
	router.HandleFunc("/house/{id}", handler.GetFlatsInDeHouse)

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var resp openapi.HouseIdGet200Response
	err = json.NewDecoder(rr.Body).Decode(&resp)
	if err != nil {
		t.Fatalf("could not decode response: %v", err)
	}

	expectedResponse := openapi.HouseIdGet200Response{
		Flats: make([]openapi.Flat, len(expectedFlats)),
	}
	for i, flat := range expectedFlats {
		expectedResponse.Flats[i] = openapi.Flat{
			Id:      flat.Id,
			HouseId: flat.HouseId,
			Price:   flat.Price,
			Rooms:   flat.Rooms,
			Status:  openapi.Status(flat.Status),
		}
	}

	assert.Equal(t, expectedResponse, resp)
}

func TestGetFlatsInDeHouse_Error500SqlNoRows(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logrus.New()

	mockHouseService := mocks.NewMockHouseService(ctrl)
	handler := house.NewHandler(mockHouseService, nil, logger)

	mockHouseService.EXPECT().GetFlatsByHouseId(int32(1), false).Return(nil, sql.ErrNoRows)

	req, err := http.NewRequest("GET", "/house/1", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rr := httptest.NewRecorder()

	requestId := "12bh3b4h3b4"
	ctx := context.WithValue(req.Context(), "requestId", requestId)
	ctx = context.WithValue(ctx, "role", string(openapi.CLIENT))
	req = req.WithContext(ctx)

	router := mux.NewRouter()
	router.HandleFunc("/house/{id}", handler.GetFlatsInDeHouse)

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)

	var resp openapi.DummyLoginGet500Response
	err = json.NewDecoder(rr.Body).Decode(&resp)
	if err != nil {
		t.Fatalf("could not decode response: %v", err)
	}

	code := int32(errors1.CodeHouseNotFound)
	expectedResponse := openapi.DummyLoginGet500Response{
		Message:   "Flats in de house not found",
		RequestId: &requestId,
		Code:      &code,
	}

	assert.Equal(t, expectedResponse, resp)
}

func TestGetFlatsInDeHouse_Error500DbError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logrus.New()

	mockHouseService := mocks.NewMockHouseService(ctrl)
	handler := house.NewHandler(mockHouseService, nil, logger)

	mockHouseService.EXPECT().GetFlatsByHouseId(int32(1), false).Return(nil, errors.New("database error"))

	req, err := http.NewRequest("GET", "/house/1", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rr := httptest.NewRecorder()

	requestId := "12bh3b4h3b4"
	ctx := context.WithValue(req.Context(), "requestId", requestId)
	ctx = context.WithValue(ctx, "role", string(openapi.CLIENT))
	req = req.WithContext(ctx)

	router := mux.NewRouter()
	router.HandleFunc("/house/{id}", handler.GetFlatsInDeHouse)

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)

	var resp openapi.DummyLoginGet500Response
	err = json.NewDecoder(rr.Body).Decode(&resp)
	if err != nil {
		t.Fatalf("could not decode response: %v", err)
	}

	code := int32(errors1.CodeServiceError)
	expectedResponse := openapi.DummyLoginGet500Response{
		Message:   "database error",
		RequestId: &requestId,
		Code:      &code,
	}

	assert.Equal(t, expectedResponse, resp)
}

func TestGetFlatsInDeHouse_Error500InvalidId(t *testing.T) {
	logger := logrus.New()

	handler := house.NewHandler(nil, nil, logger)

	req, err := http.NewRequest("GET", "/house/afdgd", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rr := httptest.NewRecorder()

	requestId := "12bh3b4h3b4"
	ctx := context.WithValue(req.Context(), "requestId", requestId)
	req = req.WithContext(ctx)

	router := mux.NewRouter()
	router.HandleFunc("/house/{id}", handler.GetFlatsInDeHouse)

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotImplemented, rr.Code)

	var resp openapi.DummyLoginGet500Response
	err = json.NewDecoder(rr.Body).Decode(&resp)
	if err != nil {
		t.Fatalf("could not decode response: %v", err)
	}

	code := int32(errors1.CodeBadRequest)
	expectedResponse := openapi.DummyLoginGet500Response{
		Message:   "Invalid id",
		RequestId: &requestId,
		Code:      &code,
	}

	assert.Equal(t, expectedResponse, resp)
}
