package tests

import (
	"apartment_search_service/internal/handlers/flat"
	"apartment_search_service/internal/models"
	"context"
	"errors"
	"strings"

	openapi "apartment_search_service/internal/openapi/gen"
	"apartment_search_service/internal/services/mocks"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateFlat_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockFlatService(ctrl)
	mockNotifier := mocks.NewMockSubscriptionService(ctrl)
	logger := logrus.New()

	handler := flat.NewHandler(mockService, mockNotifier, logger)

	reqBody := `{
		"house_id": 1,
		"price": 100000,
		"rooms": 3
	}`

	req, err := http.NewRequest("POST", "/flat/create", strings.NewReader(reqBody))
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	flat := &models.Flat{
		HouseId: 1,
		Price:   100000,
		Rooms:   3,
		Status:  string(openapi.CREATED),
	}

	mockService.EXPECT().CreateFlat(flat).DoAndReturn(func(f *models.Flat) error {
		f.Id = 0
		return nil
	})
	mockNotifier.EXPECT().NotifySubscribers(flat.HouseId, flat.Id).Return()

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/flat/create", handler.CreateFlat)

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var resp openapi.Flat
	err = json.NewDecoder(rr.Body).Decode(&resp)
	if err != nil {
		t.Fatalf("could not decode response: %v", err)
	}

	expectedResponse := openapi.Flat{
		Id:      0,
		HouseId: 1,
		Price:   100000,
		Rooms:   3,
		Status:  openapi.Status(flat.Status),
	}

	assert.Equal(t, expectedResponse, resp)
}

func TestCreateFlat_DecodeError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logrus.New()

	handler := flat.NewHandler(nil, nil, logger)

	reqBody := `invalid json`

	req, err := http.NewRequest("POST", "/flat/create", strings.NewReader(reqBody))
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	requestId := "12bh3b4h3b4"
	ctx := context.WithValue(req.Context(), "requestId", requestId)
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/flat/create", handler.CreateFlat)

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestCreateFlat_ValidationError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logrus.New()

	handler := flat.NewHandler(nil, nil, logger)

	reqBody := `{
		"house_id": 0,
		"price": 100000,
		"rooms": 3
	}`

	req, err := http.NewRequest("POST", "/flat/create", strings.NewReader(reqBody))
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	requestId := "12bh3b4h3b4"
	ctx := context.WithValue(req.Context(), "requestId", requestId)
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/flat/create", handler.CreateFlat)

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestCreateFlat_ServiceError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockFlatService(ctrl)
	logger := logrus.New()

	handler := flat.NewHandler(mockService, nil, logger)

	reqBody := `{
		"house_id": 1,
		"price": 100000,
		"rooms": 3
	}`

	req, err := http.NewRequest("POST", "/flat/create", strings.NewReader(reqBody))
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	mockService.EXPECT().CreateFlat(gomock.Any()).Return(errors.New("service error"))

	requestId := "12bh3b4h3b4"
	ctx := context.WithValue(req.Context(), "requestId", requestId)
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/flat/create", handler.CreateFlat)

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}
