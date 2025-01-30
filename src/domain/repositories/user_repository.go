package repositories

import "API_REST/domain/entities"

type UserRepository interface {
	Save(user *entities.User) error
	FindByID(id int32) (*entities.User, error)
	Update(user *entities.User) error
	Delete(id int32) error
}
