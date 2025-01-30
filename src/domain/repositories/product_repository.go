package repositories

import "API_REST/src/domain/entities"

type ProductRepository interface {
	Save(product *entities.Product) error
	FindByID(id int32) (*entities.Product, error)
	Update(product *entities.Product) error
	Delete(id int32) error
}
