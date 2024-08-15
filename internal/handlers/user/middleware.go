package user

import (
	openapi "apartment_search_service/internal/openapi/gen"
	cont "github.com/gorilla/context"
	"net/http"
	"strings"
)

func (h *Handler) Authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			// todo error
			http.Error(w, "Empty authorization header", http.StatusUnauthorized)
			return
		}
		headerPart := strings.Split(header, " ")
		if len(headerPart) != 2 {
			// todo error
			http.Error(w, "Invalid authorization header", http.StatusUnauthorized)
			return
		}

		userId, err := h.service.ParseToken(headerPart[1])
		if err != nil {
			// todo error
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		user, err := h.service.GetUser(userId)
		if err != nil {
			// todo error
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		cont.Set(r, "role", user.UserType)
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) ModeratorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role, ok := cont.Get(r, "role").(string)
		if !ok {
			// todo error
			http.Error(w, "Role not found", http.StatusForbidden)
			return
		}

		if role != string(openapi.MODERATOR) {
			// todo error
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
