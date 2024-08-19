package flat

import (
	"apartment_search_service/internal/errors"
	"apartment_search_service/internal/models"
	openapi "apartment_search_service/internal/openapi/gen"
	"apartment_search_service/internal/services"
	"apartment_search_service/internal/utils"
	"database/sql"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"net/http"

	errors1 "errors"
)

type Handler struct {
	service  services.FlatService
	logger   *logrus.Logger
	notifier services.SubscriptionService
}

func NewHandler(service services.FlatService, notifier services.SubscriptionService, logger *logrus.Logger) *Handler {
	return &Handler{
		service:  service,
		logger:   logger,
		notifier: notifier,
	}
}

func (h *Handler) CreateFlat(w http.ResponseWriter, r *http.Request) {
	var req openapi.FlatCreatePostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError400(w, r, h.logger, "Invalid request body", errors.CodeInvalidInput)
		return
	}

	validate := validator.New()

	if err := validate.Struct(req); err != nil {
		utils.RespondWithError400(w, r, h.logger, "Invalid input data", errors.CodeInvalidInput)
		return
	}

	flat := models.Flat{
		HouseId: req.GetHouseId(),
		Price:   req.GetPrice(),
		Rooms:   req.GetRooms(),
		Status:  string(openapi.CREATED),
	}

	if err := h.service.CreateFlat(&flat); err != nil {
		utils.RespondWithError500(w, r, h.logger, http.StatusInternalServerError, err.Error(), errors.CodeServiceError)
		return
	}

	h.notifier.NotifySubscribers(flat.HouseId, flat.Id)

	h.logger.WithFields(logrus.Fields{
		"flat_id": flat.Id,
	}).Info("Flat created successfully")

	resp := openapi.Flat{
		Id:      flat.Id,
		HouseId: flat.HouseId,
		Price:   flat.Price,
		Rooms:   flat.Rooms,
		Status:  openapi.Status(flat.Status),
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) UpdateFlat(w http.ResponseWriter, r *http.Request) {
	var req openapi.FlatUpdatePostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError400(w, r, h.logger, "Invalid request body", errors.CodeInvalidInput)
		return
	}

	validate := validator.New()

	if err := validate.Struct(req); err != nil {
		utils.RespondWithError400(w, r, h.logger, "Invalid input data", errors.CodeInvalidInput)
		return
	}

	id := r.Context().Value("userId").(string)

	flat, err := h.service.UpdateStatus(req.GetId(), id, string(req.GetStatus()))
	if err != nil {
		if errors1.Is(err, sql.ErrNoRows) {
			utils.RespondWithError500(w, r, h.logger, http.StatusInternalServerError, "Flat not found", errors.CodeFlatNotFound)
		} else {
			utils.RespondWithError500(w, r, h.logger, http.StatusInternalServerError, err.Error(), errors.CodeFlatStatusError)
		}
		return
	}

	h.logger.WithFields(logrus.Fields{
		"flat_id": flat.Id,
		"status":  flat.Status,
	}).Info("Flat updated successfully")

	resp := openapi.Flat{
		Id:      flat.Id,
		HouseId: flat.HouseId,
		Price:   flat.Price,
		Rooms:   flat.Rooms,
		Status:  openapi.Status(flat.Status),
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
