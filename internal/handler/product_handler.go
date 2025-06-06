package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/miloalej-dev/supermarket/internal/model"
	"github.com/miloalej-dev/supermarket/internal/repository"
	"net/http"
	"strconv"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	data, err := json.MarshalIndent(repository.Products, "", "\t")
	_, err = w.Write(data)
	if err != nil {
		return
	}
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		http.Error(w, "Product ID is required", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	product, err := repository.FindProductByID(id)

	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	data, err := json.MarshalIndent(product, "", "\t")

	if err != nil {
		http.Error(w, "Error encoding product", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(data)

	if err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
		return
	}
}

func GetProductsByPriceGreaterThan(w http.ResponseWriter, r *http.Request) {
	priceGreater, err := strconv.ParseFloat(r.URL.Query().Get("priceGt"), 64)

	if err != nil {
		http.Error(w, "Invalid price greater than value", http.StatusBadRequest)
		return
	}

	items := repository.FindProductsByPriceGreaterThan(priceGreater)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	data, err := json.MarshalIndent(items, "", "\t")
	if err != nil {
		http.Error(w, "Error encoding products", http.StatusInternalServerError)
		return
	}
	_, err = w.Write(data)

	if err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
		return
	}
}

func PostProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var product model.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid requests payload", http.StatusBadRequest)
		return
	}

	repository.AddProduct(product)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	data, err := json.MarshalIndent(product, "", "\t")
	if err != nil {
		http.Error(w, "Error encoding product", http.StatusInternalServerError)
		return
	}
	_, err = w.Write(data)

	if err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
		return
	}
}
