package service

import (
	"github.com/jonatasmateus/e-commerce-fc-17/go-api/internal/database"
	"github.com/jonatasmateus/e-commerce-fc-17/go-api/internal/entity"
)

type ProductService struct {
	ProductDB database.ProductDB
}

func NewProductService(ProductDB database.ProductDB) *ProductService {
	return &ProductService{ProductDB: ProductDB}
}

func (cs *ProductService) CreateProduct(name, description string, price float64, category_id, image_url string) (*entity.Product, error) {
	product := entity.NewProduct(name, description, price, category_id, image_url)
	_, err := cs.ProductDB.CreateProduct(product)
	if err != nil {
		return nil, err
	}
	return product, err
}

func (cs *ProductService) GetProduct(id string) (*entity.Product, error) {
	product, err := cs.ProductDB.GetProduct(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (cs *ProductService) GetProducts() ([]*entity.Product, error) {
	products, err := cs.ProductDB.GetProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (cs *ProductService) GetProductsByCategoryID(categoryID string) ([]*entity.Product, error) {
	products, err := cs.ProductDB.GetProductsByCategoryID(categoryID)
	if err != nil {
		return nil, err
	}
	return products, nil
}
