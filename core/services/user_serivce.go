package services

import (
	"errors"

	"github.com/horThana/Backend/core/domain"
	"github.com/horThana/Backend/core/ports"
)


type UserService interface {
	CreateUser(user domain.User) error
	FindUseById(id string) (domain.User, error)
	FindAllUsers() ([]domain.User, error)
	DeleteUser(id string) error
}

type UserServiceImpl struct {
	repo ports.UsersRepository
}

func NewUserService(repo ports.UsersRepository) UserService {
	return &UserServiceImpl{repo: repo}
}

//CreateUser is a method to create user
func (s *UserServiceImpl) CreateUser(user domain.User) error {
	if user.Username == "" {
		return errors.New("name is required")
	}
	if err := s.repo.SaveUser(user); err != nil {
		return err
	}
	return nil
}

//FindUseById is a method to get user by id
func (s *UserServiceImpl) FindUseById(id string) (domain.User, error) {
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

//FindAllUser is a method to get all user
func (s *UserServiceImpl) FindAllUsers() ([]domain.User, error) {
	user, err := s.repo.GetAllUser()
	if err != nil {
		return []domain.User{}, err
	}
	return user, nil
}

//DeleteUser is a method to delete user
func (s *UserServiceImpl) DeleteUser(id string) error {
	err := s.repo.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}

