package house

import (
	"apartment_search_service/internal/models"
	openapi "apartment_search_service/internal/openapi/gen"
	"apartment_search_service/internal/services"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Handler struct {
	service services.HouseService
}

func NewHandler(service services.HouseService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateHouse(w http.ResponseWriter, r *http.Request) {
	var req openapi.HouseCreatePostRequest
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

	house := models.House{
		Address:   req.GetAddress(),
		Year:      req.GetYear(),
		Developer: req.GetDeveloper(),
		CreatedAt: time.Now(),
	}

	if err := h.service.CreateHouse(&house); err != nil {
		log.Println(err)
		// todo err500
		return
	}

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
		log.Println(err)
		// todo error400
		return
	}

	role := context.Get(r, "role").(string)
	var isModer bool
	if role != string(openapi.CLIENT) {
		isModer = true
	}

	flats, err := h.service.GetFlatsByHouseId(int32(id), isModer)
	if err != nil {
		// todo error500
		return
	}

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
