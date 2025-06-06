package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/miloalej-dev/supermarket/internal/handler"
	"net/http"
)

func main() {

	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Route("/products", func(r chi.Router) {
		r.Get("/", handler.GetProducts)
		r.Get("/{id}", handler.GetProductById)
		r.Get("/search", handler.GetProductsByPriceGreaterThan)
		r.Post("/", handler.PostProduct)
	})

	defer func() {
		if r := recover(); r != nil {
			http.Error(http.ResponseWriter(nil), "Internal Server Error", http.StatusInternalServerError)
		}

		fmt.Println("Server stopped gracefully")
	}()

	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}

}
