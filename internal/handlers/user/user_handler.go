package user

import (
	"apartment_search_service/internal/models"
	openapi "apartment_search_service/internal/openapi/gen"
	"apartment_search_service/internal/services"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"log"
	"net/http"
)

type Handler struct {
	service services.AuthService
}

func NewHandler(service services.AuthService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) DummyLogin(w http.ResponseWriter, r *http.Request) {
	userType, err := openapi.NewUserTypeFromValue(r.URL.Query().Get("user_type"))
	if err != nil {
		log.Println(err)
		//todo error
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
		log.Println(err)
		// todo error
		//http.Error(w, "Unable to create user", http.StatusInternalServerError)
		return
	}

	token, err := h.service.GenerateToken(id, password)
	if err != nil {
		log.Println(err)
		//todo error
		return
	}

	resp := openapi.DummyLoginGet200Response{Token: &token}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req openapi.LoginPostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		// todo error
		return
	}

	validate := validator.New()

	if err := validate.Struct(req); err != nil {
		log.Println(err)
		// todo error
		//errs := err.(validator.ValidationErrors)
		//newErrorResponse(w, http.StatusBadRequest, errs.Error())
		return
	}

	token, err := h.service.GenerateToken(req.GetId(), req.GetPassword())
	if err != nil {
		log.Println(err)
		//todo error
		return
	}

	resp := openapi.DummyLoginGet200Response{Token: &token}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req openapi.RegisterPostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		// todo error
		return
	}

	validate := validator.New()

	if err := validate.Struct(req); err != nil {
		log.Println(err)
		// todo error
		//errs := err.(validator.ValidationErrors)
		//newErrorResponse(w, http.StatusBadRequest, errs.Error())
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
		log.Println(err)
		// todo error
		return
	}
	resp := openapi.RegisterPost200Response{UserId: &id}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
