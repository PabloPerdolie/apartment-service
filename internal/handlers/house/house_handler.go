package house

import (
	"apartment_search_service/internal/errors"
	"apartment_search_service/internal/models"
	openapi "apartment_search_service/internal/openapi/gen"
	"apartment_search_service/internal/services"
	"apartment_search_service/internal/utils"
	"database/sql"
	"encoding/json"
	errors1 "errors"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

type Handler struct {
	service     services.HouseService
	logger      *logrus.Logger
	subscribers services.SubscriptionService
}

func NewHandler(service services.HouseService, logger *logrus.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}

func (h *Handler) CreateHouse(w http.ResponseWriter, r *http.Request) {
	var req openapi.HouseCreatePostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError400(w, r, h.logger, "Invalid request body", errors.CodeInvalidInput)
		return

	}

	validate := validator.New()

	if err := validate.Struct(req); err != nil {
		utils.RespondWithError400(w, r, h.logger, "Invalid input data", errors.CodeInvalidInput)
		return
	}

	house := models.House{
		Address:   req.GetAddress(),
		Year:      req.GetYear(),
		Developer: req.GetDeveloper(),
		CreatedAt: time.Now(),
	}

	if err := h.service.CreateHouse(&house); err != nil {
		utils.RespondWithError500(w, r, h.logger, http.StatusInternalServerError, "House already exists", errors.CodeHouseAlreadyExists)
		return
	}
	h.logger.WithFields(logrus.Fields{
		"house_id": house.Id,
	}).Info("House created successfully")

	resp := openapi.House{
		Id:        house.Id,
		Address:   house.Address,
		Year:      house.Year,
		Developer: *openapi.NewNullableString(&house.Developer),
		CreatedAt: &house.CreatedAt,
		UpdateAt:  &house.UpdatedAt.Time,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) GetFlatsInDeHouse(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || id < 1 {
		utils.RespondWithError500(w, r, h.logger, http.StatusNotImplemented, "Invalid id", errors.CodeBadRequest)
		return
	}

	// role := context.Get(r, "role").(string)
	role := r.Context().Value("role").(string)
	var isModer bool
	if role != string(openapi.CLIENT) {
		isModer = true
	}

	flats, err := h.service.GetFlatsByHouseId(int32(id), isModer)
	if err != nil {
		if errors1.Is(err, sql.ErrNoRows) {
			utils.RespondWithError500(w, r, h.logger, http.StatusInternalServerError, "Flats in de house not found", errors.CodeHouseNotFound)
		} else {
			utils.RespondWithError500(w, r, h.logger, http.StatusInternalServerError, err.Error(), errors.CodeServiceError)
		}
		return
	}

	h.logger.WithFields(logrus.Fields{
		"house_id": id,
	}).Info("Flats received successfully")

	resp := openapi.HouseIdGet200Response{
		Flats: make([]openapi.Flat, len(flats)),
	}

	for i, flat := range flats {
		resp.Flats[i] = openapi.Flat{
			Id:      flat.Id,
			HouseId: flat.HouseId,
			Price:   flat.Price,
			Rooms:   flat.Rooms,
			Status:  openapi.Status(flat.Status),
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) SubscribeToNewFlatsInDeHouse(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || id < 1 {
		utils.RespondWithError500(w, r, h.logger, http.StatusNotImplemented, "Invalid id", errors.CodeBadRequest)
		return
	}

	var req openapi.HouseIdSubscribePostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError400(w, r, h.logger, "Invalid request body", errors.CodeInvalidInput)
		return
	}

	validate := validator.New()

	if err := validate.Struct(req); err != nil {
		utils.RespondWithError400(w, r, h.logger, "Invalid input data", errors.CodeInvalidCredentials)
		return
	}

	err = h.subscribers.AddSubscriber(int32(id), req.GetEmail())
	if err != nil {
		if errors1.Is(err, sql.ErrNoRows) {
			utils.RespondWithError500(w, r, h.logger, http.StatusInternalServerError, "House not found", errors.CodeHouseNotFound)
		} else {
			utils.RespondWithError500(w, r, h.logger, http.StatusInternalServerError, err.Error(), errors.CodeServiceError)
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Successfully subscribed")
}
