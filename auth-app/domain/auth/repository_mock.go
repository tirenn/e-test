package auth

import (
	"tirenn/test-efishery/auth-app/models"

	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	Mock mock.Mock
}

func (r *RepositoryMock) Create(auth *models.Auth) error {
	args := r.Mock.Called(auth)
	return args.Error(0)
}

func (r *RepositoryMock) Get(id string) (models.Auth, error) {
	args := r.Mock.Called(id)
	result := args.Get(0)
	return result.(models.Auth), args.Error(1)
}
