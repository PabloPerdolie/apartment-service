package handlers

import (
	"apartment_search_service/internal/handlers/user"
	"github.com/gorilla/mux"
)

func SetupRoutes(user *user.Handler) *mux.Router {
	r := mux.NewRouter()

	//r.Use(middleware.Cors)

	r.HandleFunc("/dummyLogin", user.DummyLogin).Methods("GET", "OPTIONS")
	r.HandleFunc("/login", user.Login).Methods("POST", "OPTIONS")
	r.HandleFunc("/register", user.Register).Methods("POST", "OPTIONS")

	//house := r.PathPrefix("/house").Subrouter()
	//
	//flat := r.PathPrefix("/flat").Subrouter()

	return r
}
