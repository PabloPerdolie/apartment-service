package flat

import (
	"apartment_search_service/internal/models"
	openapi "apartment_search_service/internal/openapi/gen"
	"apartment_search_service/internal/services"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/context"
	"log"
	"net/http"
)

type Handler struct {
	service services.FlatService
}

func NewHandler(service services.FlatService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateFlat(w http.ResponseWriter, r *http.Request) {
	var req openapi.FlatCreatePostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
		// todo error
		return
	}

	validate := validator.New()

	if err := validate.Struct(req); err != nil {
		log.Println(err)
		// todo error
		return
	}

	flat := models.Flat{
		HouseId: req.GetHouseId(),
		Price:   req.GetPrice(),
		Rooms:   req.GetRooms(),
		Status:  string(openapi.CREATED),
	}

	if err := h.service.CreateFlat(&flat); err != nil {
		log.Println(err)
		// todo err500
		return
	}

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
		log.Println(err)
		//todo error400
		return
	}

	validate := validator.New()

	if err := validate.Struct(req); err != nil {
		log.Println(err)
		// todo error400
		return
	}

	id := context.Get(r, "userId").(string)

	flat, err := h.service.UpdateStatus(req.GetId(), id, string(req.GetStatus()))
	if err != nil {
		log.Println(err)
		// todo error500
		return
	}

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
