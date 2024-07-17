package services

import (
	"go_project_structure/internal/models"
	"go_project_structure/internal/repository"
)

type UserService interface {
	GetAllUsers() ([]models.User, error)
	FindOneByKey(id uint) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	DeleteUser(user models.User) error
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		userRepository: repo,
	}
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.userRepository.FindAll()
}

func (s *userService) FindOneByKey(id uint) (models.User, error) {
	return s.userRepository.FindOneByKey(id)
}

func (s *userService) CreateUser(user models.User) (models.User, error) {
	return s.userRepository.CreateUser(user)
}

func (s *userService) DeleteUser(user models.User) error {
	return s.userRepository.DeleteUser(user)
}
