package repositories

import (
	"golangApiBoilerplate/internals/core/domain"
	"golangApiBoilerplate/internals/core/ports"
)

type UserRepository struct{}

var _ ports.IUserRepository = (*UserRepository)(nil)

func NewUserRepository() (*UserRepository, error) {
	return &UserRepository{}, nil
}

func (r *UserRepository) Get(email string) (domain.User, error) {
	return domain.User{ID: 1, Email: "test1@test.com", Password: "111"}, nil
}

func (r *UserRepository) GetList() ([]domain.User, error) {
	return []domain.User{
		{ID: 1, Email: "test1@test.com", Password: "111"},
		{ID: 2, Email: "test2@test.com", Password: "222"},
		{ID: 3, Email: "test3@test.com", Password: "333"},
		{ID: 4, Email: "test4@test.com", Password: "444"},
	}, nil
}
