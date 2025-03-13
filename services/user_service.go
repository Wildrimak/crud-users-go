package services

import (
	"cadastro-usuarios-go/models"
	"cadastro-usuarios-go/repositories"
)

type UserService struct {
	Repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) CreateUser(user models.User) (int, error) {
	return s.Repo.Create(user)
}

func (s *UserService) GetUsers() ([]models.User, error) {
	return s.Repo.GetAll()
}

func (s *UserService) UpdateUser(user models.User) error {
	return s.Repo.Update(user)
}

func (s *UserService) DeleteUser(id int) error {
	return s.Repo.Delete(id)
}
