package handlers

import (
	"apartment_search_service/internal/handlers/flat"
	"apartment_search_service/internal/handlers/house"
	"apartment_search_service/internal/handlers/user"
	"github.com/gorilla/mux"
	"net/http"
)

func SetupRoutes(user *user.Handler, house *house.Handler, flat *flat.Handler) *mux.Router {
	r := mux.NewRouter()

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})

	r.HandleFunc("/dummyLogin", user.DummyLogin).Methods("GET", "OPTIONS")
	r.HandleFunc("/login", user.Login).Methods("POST", "OPTIONS")
	r.HandleFunc("/register", user.Register).Methods("POST", "OPTIONS")

	houser := r.PathPrefix("/house").Subrouter()
	houser.Use(user.Authorize)
	houser.HandleFunc("/{id}", house.GetFlatsInDeHouse).Methods("GET", "OPTIONS")
	//houser.HandleFunc("/{id}/subscribe", user.Register).Methods("POST", "OPTIONS")
	houser.Handle("/create",
		user.ModeratorMiddleware(http.HandlerFunc(house.CreateHouse))).Methods("POST", "OPTIONS") // admin

	flatr := r.PathPrefix("/flat").Subrouter()
	flatr.Use(user.Authorize)
	flatr.HandleFunc("/create", flat.CreateFlat).Methods("POST", "OPTIONS")
	flatr.Handle("/update",
		user.ModeratorMiddleware(http.HandlerFunc(flat.UpdateFlat))).Methods("POST", "OPTIONS") // admin

	return r
}
