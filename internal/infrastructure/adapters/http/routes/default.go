package routes

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/miloalej-dev/supermarket/internal/infrastructure/adapters/http/responses"
	"net/http"
)

// DefaultRoutes sets up the default routes for the API, such as the root endpoint and error handling for not found
// and method not allowed requests.
func DefaultRoutes(router chi.Router) {
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		_, _ = w.Write([]byte(`{"message": "Welcome to the Supermarket API"}`))
	})

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")

		notFoundError, _ := json.MarshalIndent(responses.ErrorResponse{
			Status:  http.StatusNotFound,
			Message: "Resource not found, please check the URL",
		}, "", "\t")
		_, _ = w.Write(notFoundError)
	})

	router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Header().Set("Content-Type", "application/json")

		methodNotAllowedError, _ := json.MarshalIndent(responses.ErrorResponse{
			Status:  http.StatusMethodNotAllowed,
			Message: "Method not allowed, please check the request method",
		}, "", "\t")
		_, _ = w.Write(methodNotAllowedError)
	})
}
