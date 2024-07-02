package services

import (
	"golangApiBoilerplate/internals/core/domain"
	"golangApiBoilerplate/internals/core/ports"
)

type UserService struct {
	userRepository ports.IUserRepository
}

// This line is for get feedback in case we are not implementing the interface correctly
var _ ports.IUserService = (*UserService)(nil)

func NewUserService(repository ports.IUserRepository) *UserService {
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
