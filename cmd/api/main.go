package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/miloalej-dev/supermarket/internal/infrastructure/adapters/database/repositories"
	"github.com/miloalej-dev/supermarket/internal/infrastructure/adapters/http/handlers"
	"github.com/miloalej-dev/supermarket/internal/infrastructure/adapters/http/middlewares"
	"github.com/miloalej-dev/supermarket/internal/infrastructure/adapters/http/responses"
	"net/http"
)

func main() {
	// Initialize the in-memory product repository
	productRepository := repositories.NewFileRepository()

	productHandler := handlers.NewProductHandler(productRepository)

	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Route("/products", func(r chi.Router) {
		r.Use(middlewares.AuthMiddleware)
		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProductById)
		r.Get("/search", productHandler.GetProductsBySearch)
		r.Post("/", productHandler.CreateProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Patch("/{id}", productHandler.PatchProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)

		r.NotFound(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Header().Set("Content-Type", "application/json")

			notFoundError, _ := json.MarshalIndent(responses.ErrorResponse{
				Status:  http.StatusNotFound,
				Message: "Resource not found, please check the URL",
			}, "", "\t")
			_, _ = w.Write(notFoundError)
		})

		r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Header().Set("Content-Type", "application/json")

			methodNotAllowedError, _ := json.MarshalIndent(responses.ErrorResponse{
				Status:  http.StatusMethodNotAllowed,
				Message: "Method not allowed, please check the request method",
			}, "", "\t")
			_, _ = w.Write(methodNotAllowedError)
		})
	})

	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
		return
	}

}
