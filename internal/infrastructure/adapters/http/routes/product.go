package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/miloalej-dev/supermarket/internal/infrastructure/adapters/http/handlers"
	"github.com/miloalej-dev/supermarket/internal/infrastructure/adapters/http/middlewares"
)

// ProductRoutes sets up the routes for product-related operations.
func ProductRoutes(router chi.Router, handler *handlers.ProductHandler) {
	router.Route("/products", func(r chi.Router) {
		r.Use(middlewares.AuthMiddleware)

		r.Get("/", handler.GetProducts)
		r.Get("/{id}", handler.GetProductById)
		r.Get("/search", handler.GetProductsBySearch)
		r.Post("/", handler.CreateProduct)
		r.Put("/{id}", handler.UpdateProduct)
		r.Patch("/{id}", handler.PatchProduct)
		r.Delete("/{id}", handler.DeleteProduct)
	})
}
