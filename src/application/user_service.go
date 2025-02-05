package application

import (
	"API_REST/src/domain/entities"
	"API_REST/src/domain/repositories"
)

type UserService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(username, email string) error {
	user := entities.NewUser(username, email)
	return s.repo.Create(user)
}

func (s *UserService) GetAllUsers() ([]entities.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) UpdateUser(id int32, username, email string) error {
	user := &entities.User{ID: id, Username: username, Email: email}
	return s.repo.Update(user)
}

func (s *UserService) DeleteUser(id int32) error {
	return s.repo.Delete(id)
}
