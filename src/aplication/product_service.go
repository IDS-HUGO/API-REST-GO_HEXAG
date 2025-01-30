package application

import (
	oyecto/domain/entities"
	oyecto/domain/repositories"
)

type ProductService struct {
	repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(name string, price float32) error {
	product := entities.NewProduct(name, price)
	return s.repo.Save(product)
}

func (s *ProductService) GetProduct(id int32) (*entities.Product, error) {
	return s.repo.FindByID(id)
}

func (s *ProductService) UpdateProduct(id int32, name string, price float32) error {
	product, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	product.SetName(name)
	product.SetPrice(price)
	return s.repo.Update(product)
}

func (s *ProductService) DeleteProduct(id int32) error {
	return s.repo.Delete(id)
}
