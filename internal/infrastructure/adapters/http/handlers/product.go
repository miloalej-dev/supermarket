package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/miloalej-dev/supermarket/internal/application/ports/outbound"
	"github.com/miloalej-dev/supermarket/internal/domain"
	"github.com/miloalej-dev/supermarket/internal/infrastructure/adapters/http/responses"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	productRepository outbound.ProductRepository
}

func NewProductHandler(productRepository outbound.ProductRepository) *ProductHandler {
	return &ProductHandler{
		productRepository: productRepository,
	}
}

func (p *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, _ := json.MarshalIndent(p.productRepository.FindAll(), "", "\t")

	w.WriteHeader(http.StatusOK)

	_, _ = w.Write(data)
}

func (p *ProductHandler) GetProductById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		badRequestError, _ := json.MarshalIndent(responses.BadRequestError("Product ID is required or invalid"), "", "\t")
		_, _ = w.Write(badRequestError)
		return
	}

	product := p.productRepository.FindById(id)

	if product == nil {
		w.WriteHeader(http.StatusNotFound)
		notFoundError, _ := json.MarshalIndent(responses.NotFoundError("Product does not exist"), "", "\t")
		_, _ = w.Write(notFoundError)
		return
	}

	w.WriteHeader(http.StatusOK)
	data, _ := json.MarshalIndent(product, "", "\t")

	_, _ = w.Write(data)
}

func (p *ProductHandler) GetProductsBySearch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	priceStr := r.URL.Query().Get("priceGt")

	if priceStr == "" {
		badRequestError, _ := json.MarshalIndent(responses.BadRequestError("Query param 'PriceGt' is required"), "", "\t")
		_, _ = w.Write(badRequestError)
		return
	}

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		badRequestError, _ := json.MarshalIndent(responses.BadRequestError("Query param 'PriceGt' must be a number"), "", "\t")
		_, _ = w.Write(badRequestError)
		return
	}

	products := p.productRepository.FindByPriceGreaterThan(price)

	w.WriteHeader(http.StatusOK)
	data, _ := json.MarshalIndent(products, "", "\t")

	_, _ = w.Write(data)
}

func (p *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var product domain.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		badRequestError, _ := json.MarshalIndent(responses.BadRequestError("Invalid request body"), "", "\t")
		_, _ = w.Write(badRequestError)
		return
	}

	savedProduct, _ := p.productRepository.Save(product)

	w.WriteHeader(http.StatusCreated)
	data, _ := json.MarshalIndent(savedProduct, "", "\t")

	_, _ = w.Write(data)
}

func (p *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		badRequestError, _ := json.MarshalIndent(responses.BadRequestError("Product ID is required or invalid"), "", "\t")
		_, _ = w.Write(badRequestError)
		return
	}

	var product domain.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		badRequestError, _ := json.MarshalIndent(responses.BadRequestError("Invalid request body"), "", "\t")
		_, _ = w.Write(badRequestError)
		return
	}

	update := p.productRepository.Update(id, product)

	if update == nil {
		w.WriteHeader(http.StatusNotFound)
		notFoundError, _ := json.MarshalIndent(responses.NotFoundError("Product does not exist"), "", "\t")
		_, _ = w.Write(notFoundError)
		return
	}

	w.WriteHeader(http.StatusOK)
	data, _ := json.MarshalIndent(update, "", "\t")

	_, _ = w.Write(data)
}

func (p *ProductHandler) PatchProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		badRequestError, _ := json.MarshalIndent(responses.BadRequestError("Product ID is required or invalid"), "", "\t")
		_, _ = w.Write(badRequestError)
		return
	}

	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		badRequestError, _ := json.MarshalIndent(responses.BadRequestError("Invalid request body"), "", "\t")
		_, _ = w.Write(badRequestError)
		return
	}

	product := p.productRepository.UpdatePartial(id, updates)

	if product == nil {
		w.WriteHeader(http.StatusNotFound)
		notFoundError, _ := json.MarshalIndent(responses.NotFoundError("Product does not exist"), "", "\t")
		_, _ = w.Write(notFoundError)
	}

	w.WriteHeader(http.StatusOK)
	data, _ := json.MarshalIndent(product, "", "\t")

	_, _ = w.Write(data)
}

func (p *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		badRequestError, _ := json.MarshalIndent(responses.BadRequestError("Product ID is required or invalid"), "", "\t")
		_, _ = w.Write(badRequestError)
		return
	}

	product := p.productRepository.FindById(id)
	if product == nil {
		w.WriteHeader(http.StatusNotFound)
		notFoundError, _ := json.MarshalIndent(responses.NotFoundError("Product does not exist"), "", "\t")
		_, _ = w.Write(notFoundError)
		return
	}

	err = p.productRepository.Delete(*product)
	if err != nil {
		http.Error(w, "Failed to delete product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
