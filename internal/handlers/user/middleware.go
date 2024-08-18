package user

import (
	"apartment_search_service/internal/errors"
	openapi "apartment_search_service/internal/openapi/gen"
	"apartment_search_service/internal/utils"
	cont "github.com/gorilla/context"
	"net/http"
	"strings"
)

func (h *Handler) Authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			utils.RespondWithError400(w, r, h.logger, "Empty authorization header", errors.CodeEmptyAuthorizationHeader)
			return
		}
		headerPart := strings.Split(header, " ")
		if len(headerPart) != 2 {
			utils.RespondWithError400(w, r, h.logger, "Invalid authorization header", errors.CodeInvalidAuthorizationHeader)
			return
		}

		userId, err := h.service.ParseToken(headerPart[1])
		if err != nil {
			utils.RespondWithError401(w, r, h.logger, err.Error(), errors.CodeUnauthorized)
			return
		}

		user, err := h.service.GetUser(userId)
		if err != nil {
			utils.RespondWithError401(w, r, h.logger, err.Error(), errors.CodeUnauthorized)
			return
		}

		cont.Set(r, "role", user.UserType)
		cont.Set(r, "userId", user.UserId)
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) ModeratorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role := cont.Get(r, "role").(string)

		if role != string(openapi.MODERATOR) {
			utils.RespondWithError401(w, r, h.logger, "Not enough rights", errors.CodeForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
