package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jonatasmateus/e-commerce-fc-17/go-api/internal/entity"
	"github.com/jonatasmateus/e-commerce-fc-17/go-api/internal/service"
)

type WebProductHandler struct {
	ProductService *service.ProductService
}

func NewWebProductHandler(ProductService service.ProductService) *WebProductHandler {
	return &WebProductHandler{ProductService: &ProductService}
}

func (wch *WebProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := wch.ProductService.CreateProduct(product.Name, product.Description, product.Price, product.CategoryID, product.ImageURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func (wch *WebProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}
	product, err := wch.ProductService.GetProduct(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(product)
}

func (wch *WebProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := wch.ProductService.GetProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}

func (wch *WebProductHandler) GetProductsByCategoryID(w http.ResponseWriter, r *http.Request) {
	categoryID := chi.URLParam(r, "id")
	if categoryID == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}
	products, err := wch.ProductService.GetProductsByCategoryID(categoryID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}
