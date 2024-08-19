package utils

import (
	openapi "apartment_search_service/internal/openapi/gen"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

func RespondWithError500(w http.ResponseWriter, r *http.Request, logger *logrus.Logger, statusCode int, message string, code int32) {
	//requestId := context.Get(r, "requestId").(string)
	requestId := r.Context().Value("requestId").(string)

	response := openapi.DummyLoginGet500Response{
		Message:   message,
		RequestId: &requestId,
		Code:      &code,
	}

	logger.Errorf("RequestId: %s, Error: %s, Code: %d", requestId, message, code)

	w.Header().Set("Retry-After", "5")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

func RespondWithError400(w http.ResponseWriter, r *http.Request, logger *logrus.Logger, message string, code int32) {
	requestId := r.Context().Value("requestId").(string)
	logger.Errorf("RequestId: %s, Error: %s, Code: %d", requestId, message, code)

	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode("Invalid input data")
}

func RespondWithError401(w http.ResponseWriter, r *http.Request, logger *logrus.Logger, message string, code int32) {
	requestId := r.Context().Value("requestId").(string)
	logger.Errorf("RequestId: %s, Error: %s, Code: %d", requestId, message, code)

	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode("Unauthorized access")
}

func RespondWithError404(w http.ResponseWriter, r *http.Request, logger *logrus.Logger, message string, code int32) {
	requestId := r.Context().Value("requestId").(string)
	logger.Errorf("RequestId: %s, Error: %s, Code: %d", requestId, message, code)

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("User not found")
}
