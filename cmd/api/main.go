package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/miloalej-dev/supermarket/internal/infrastructure/adapters/database/repositories"
	"github.com/miloalej-dev/supermarket/internal/infrastructure/adapters/http/handlers"
	"github.com/miloalej-dev/supermarket/internal/infrastructure/adapters/http/middlewares"
	"github.com/miloalej-dev/supermarket/internal/infrastructure/adapters/http/routes"
	"net/http"
)

func main() {
	// Initialize the in-memory product repository
	productRepository := repositories.NewFileRepository()

	productHandler := handlers.NewProductHandler(productRepository)

	router := chi.NewRouter()

	router.Use(middlewares.RequestLogger)

	routes.DefaultRoutes(router)
	routes.ProductRoutes(router, productHandler)

	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
		return
	}

}
