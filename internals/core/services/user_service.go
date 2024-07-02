package services

import (
	"golangApiBoilerplate/internals/core/domain"
	"golangApiBoilerplate/internals/repositories"
)

type IUserService interface {
	Get(email string) (domain.User, error)
	GetList() ([]domain.User, error)
}
type UserService struct {
	userRepository repositories.IUserRepository
}

// This line is for get feedback in case we are not implementing the interface correctly
var _ IUserService = (*UserService)(nil)

func NewUserService(repository repositories.IUserRepository) *UserService {
	return &UserService{
		userRepository: repository,
	}
}

func (s *UserService) Get(email string) (domain.User, error) {
	data, err := s.userRepository.Get(email)
	if err != nil {
		return domain.User{}, err
	}
	return data, nil
}

func (s *UserService) GetList() ([]domain.User, error) {
	data, err := s.userRepository.GetList()
	if err != nil {
		return nil, err
	}
	return data, nil
}
