package application

import (
	"API_REST/src/domain/entities"
	"API_REST/src/domain/repositories"
)

type ProductService struct {
	repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(name string, price float32) error {
	product := entities.NewProduct(name, price)
	return s.repo.Create(product)
}

func (s *ProductService) GetAllProducts() ([]entities.Product, error) {
	return s.repo.GetAll()
}

func (s *ProductService) UpdateProduct(id int32, name string, price float32) error {
	product := &entities.Product{ID: id, Name: name, Price: price}
	return s.repo.Update(product)
}

func (s *ProductService) DeleteProduct(id int32) error {
	return s.repo.Delete(id)
}
