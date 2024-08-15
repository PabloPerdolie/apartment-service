package handlers

import (
	"apartment_search_service/internal/handlers/user"
	"github.com/gorilla/mux"
	"net/http"
)

func SetupRoutes(user *user.Handler) *mux.Router {
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

	//house := r.PathPrefix("/house").Subrouter()
	//house.HandleFunc("/{id}", user.Register).Methods("POST", "OPTIONS")
	//house.HandleFunc("/{id}/subscribe", user.Register).Methods("POST", "OPTIONS")
	//house.Handle("/create", user.ModeratorMiddleware(http.HandlerFunc(user.Register))).Methods("POST", "OPTIONS") // admin
	//
	//flat := r.PathPrefix("/flat").Subrouter()
	//flat.HandleFunc("/create", user.Register).Methods("POST", "OPTIONS")
	//flat.Handle("/update", user.ModeratorMiddleware(http.HandlerFunc(user.Register))).Methods("POST", "OPTIONS") // admin

	return r
}
