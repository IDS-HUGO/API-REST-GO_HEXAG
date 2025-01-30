package application

import (
	oyecto/domain/entities"
	oyecto/domain/repositories"
)

type UserService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(name, email string) error {
	user := entities.NewUser(name, email)
	return s.repo.Save(user)
}

func (s *UserService) GetUser(id int32) (*entities.User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) UpdateUser(id int32, name, email string) error {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	user.SetName(name)
	user.SetEmail(email)
	return s.repo.Update(user)
}

func (s *UserService) DeleteUser(id int32) error {
	return s.repo.Delete(id)
}
