package user

import (
	"apartment_search_service/internal/errors"
	"apartment_search_service/internal/models"
	openapi "apartment_search_service/internal/openapi/gen"
	"apartment_search_service/internal/services"
	"apartment_search_service/internal/utils"
	"database/sql"
	"encoding/json"
	errors1 "errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Handler struct {
	service services.AuthService
	logger  *logrus.Logger
}

func NewHandler(service services.AuthService, logger *logrus.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}

func (h *Handler) DummyLogin(w http.ResponseWriter, r *http.Request) {
	userType, err := openapi.NewUserTypeFromValue(r.URL.Query().Get("user_type"))
	if err != nil {
		utils.RespondWithError500(w, r, h.logger, http.StatusNotImplemented, "Invalid user type", errors.CodeInvalidUserType)
		return
	}

	id := uuid.New().String()
	email := fmt.Sprintf("dummy_email_%s@gmail.com", id)
	password := "12345678"
	user := models.User{
		UserId:   id,
		Email:    email,
		Password: password,
		UserType: string(*userType),
	}

	_, err = h.service.CreateUser(&user)
	if err != nil {
		utils.RespondWithError500(w, r, h.logger, http.StatusInternalServerError, "Failed to create user", errors.CodeServiceError)
		return
	}

	token, err := h.service.GenerateToken(id, password)
	if err != nil {
		utils.RespondWithError500(w, r, h.logger, http.StatusInternalServerError, "Failed to generate token", errors.CodeServiceError)
		return
	}

	h.logger.WithFields(logrus.Fields{
		"user_id": id,
	}).Info("User successfully logged in")

	resp := openapi.DummyLoginGet200Response{Token: &token}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req openapi.LoginPostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError400(w, r, h.logger, "Invalid request body", errors.CodeInvalidInput)
		return
	}

	validate := validator.New()

	if err := validate.Struct(req); err != nil {
		utils.RespondWithError400(w, r, h.logger, "Invalid input data", errors.CodeInvalidCredentials)
		return
	}

	token, err := h.service.GenerateToken(req.GetId(), req.GetPassword())
	if err != nil {
		if errors1.Is(err, sql.ErrNoRows) {
			utils.RespondWithError404(w, r, h.logger, "User not found", errors.CodeUserNotFound)
		} else {
			utils.RespondWithError500(w, r, h.logger, http.StatusInternalServerError, err.Error(), errors.CodeServiceError)
		}
		return
	}

	h.logger.WithFields(logrus.Fields{
		"user_id": req.GetId(),
	}).Info("User successfully sign in")

	resp := openapi.DummyLoginGet200Response{Token: &token}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req openapi.RegisterPostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError400(w, r, h.logger, "Invalid request body", errors.CodeInvalidInput)
		return
	}

	validate := validator.New()

	if err := validate.Struct(req); err != nil {
		utils.RespondWithError400(w, r, h.logger, "Invalid input data", errors.CodeInvalidInput)
		return
	}

	user := models.User{
		UserId:   uuid.New().String(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
		UserType: string(req.GetUserType()),
	}

	id, err := h.service.CreateUser(&user)
	if err != nil {
		utils.RespondWithError500(w, r, h.logger, http.StatusInternalServerError, "User already exists", errors.CodeUserAlreadyExists)
		return
	}

	h.logger.WithFields(logrus.Fields{
		"user_id": id,
	}).Info("User successfully sign up")

	resp := openapi.RegisterPost200Response{UserId: &id}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
