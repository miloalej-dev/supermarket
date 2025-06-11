package handlers

import (
	"encoding/json"
	"github.com/miloalej-dev/supermarket/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockProductRepository struct {
	mock.Mock
}

// --- Mock Implementation ---

func (m *MockProductRepository) Save(product domain.Product) (domain.Product, error) {
	args := m.Called(product)
	return args.Get(0).(domain.Product), args.Error(1)
}

func (m *MockProductRepository) FindById(id int) *domain.Product {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*domain.Product)
}

func (m *MockProductRepository) FindAll() []domain.Product {
	args := m.Called()
	return args.Get(0).([]domain.Product)
}

func (m *MockProductRepository) Update(id int, product domain.Product) *domain.Product {
	args := m.Called(id, product)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*domain.Product)
}

func (m *MockProductRepository) UpdatePartial(id int, updates map[string]interface{}) *domain.Product {
	args := m.Called(id, updates)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*domain.Product)
}

func (m *MockProductRepository) Delete(product domain.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductRepository) FindByPriceGreaterThan(price float64) []domain.Product {
	args := m.Called(price)
	return args.Get(0).([]domain.Product)
}

// --- Handler Tests ---

func TestGetProducts(t *testing.T) {
	mockRepo := new(MockProductRepository)
	handler := NewProductHandler(mockRepo)
	expected := []domain.Product{{Id: 1, Name: "Apple"}}
	mockRepo.On("FindAll").Return(expected)

	req := httptest.NewRequest("GET", "/products", nil)
	w := httptest.NewRecorder()

	handler.GetProducts(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var got []domain.Product
	_ = json.Unmarshal(w.Body.Bytes(), &got)
	assert.Equal(t, expected, got)
	mockRepo.AssertExpectations(t)
}

/*
func TestGetProductById_Found(t *testing.T) {
	mockRepo := new(MockProductRepository)
	handler := NewProductHandler(mockRepo)
	expected := &domain.Product{Id: 2, Name: "Banana"}
	mockRepo.On("FindById", 2).Return(expected)

	req := httptest.NewRequest("GET", "/products/2", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "2")
	req = req.WithContext(chi.NewRouteContextContext(req.Context(), rctx))
	w := httptest.NewRecorder()

	handler.GetProductById(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var got domain.Product
	_ = json.Unmarshal(w.Body.Bytes(), &got)
	assert.Equal(t, *expected, got)
	mockRepo.AssertExpectations(t)
}


func TestGetProductById_NotFound(t *testing.T) {
	mockRepo := new(MockProductRepository)
	handler := NewProductHandler(mockRepo)
	mockRepo.On("FindById", 99).Return(nil)

	req := httptest.NewRequest("GET", "/products/99", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "99")
	req = req.WithContext(chi.NewRouteContextContext(req.Context(), rctx))
	w := httptest.NewRecorder()

	handler.GetProductById(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	mockRepo.AssertExpectations(t)
}


func TestCreateProduct(t *testing.T) {
	mockRepo := new(MockProductRepository)
	handler := NewProductHandler(mockRepo)
	input := domain.Product{Name: "Orange"}
	saved := domain.Product{Id: 3, Name: "Orange"}
	mockRepo.On("Save", input).Return(saved, nil)

	body, _ := json.Marshal(input)
	req := httptest.NewRequest("POST", "/products", httptest.NewBody(body))
	w := httptest.NewRecorder()

	handler.CreateProduct(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var got domain.Product
	_ = json.Unmarshal(w.Body.Bytes(), &got)
	assert.Equal(t, saved, got)
	mockRepo.AssertExpectations(t)
}
*/

/*
func TestDeleteProduct_NotFound(t *testing.T) {
	mockRepo := new(MockProductRepository)
	handler := NewProductHandler(mockRepo)
	mockRepo.On("FindById", 42).Return(nil)

	req := httptest.NewRequest("DELETE", "/products/42", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "42")
	req = req.WithContext(chi.NewRouteContextContext(req.Context(), rctx))
	w := httptest.NewRecorder()

	handler.DeleteProduct(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestDeleteProduct_Success(t *testing.T) {
	mockRepo := new(MockProductRepository)
	handler := NewProductHandler(mockRepo)
	product := &domain.Product{Id: 5, Name: "Grape"}
	mockRepo.On("FindById", 5).Return(product)
	mockRepo.On("Delete", *product).Return(nil)

	req := httptest.NewRequest("DELETE", "/products/5", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "5")
	req = req.WithContext(chi.NewRouteContextContext(req.Context(), rctx))
	w := httptest.NewRecorder()

	handler.DeleteProduct(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestDeleteProduct_Fail(t *testing.T) {
	mockRepo := new(MockProductRepository)
	handler := NewProductHandler(mockRepo)
	product := &domain.Product{Id: 6, Name: "Pear"}
	mockRepo.On("FindById", 6).Return(product)
	mockRepo.On("Delete", *product).Return(assert.AnError)

	req := httptest.NewRequest("DELETE", "/products/6", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "6")
	req = req.WithContext(chi.NewRouteContext(req.Context(), rctx))
	w := httptest.NewRecorder()

	handler.DeleteProduct(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockRepo.AssertExpectations(t)
}
*/
