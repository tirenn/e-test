package auth

import (
	"errors"
	"testing"
	"tirenn/test-efishery/auth-app/dto/http/request"
	"tirenn/test-efishery/auth-app/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Init() (*RepositoryMock, Service) {
	var repo = &RepositoryMock{Mock: mock.Mock{}}
	return repo, Service{repository: repo}
}

func TestCreateSuccess(t *testing.T) {
	repo, service := Init()

	req := request.Auth{
		Phone: "0812345",
		Name:  "name",
		Role:  "staff",
	}

	auth := models.Auth{
		BaseModel: models.BaseModel{
			ID: "auth-1",
		},
		Phone:    "0812345",
		Name:     "name",
		Role:     "staff",
		Password: "asdf",
	}

	repo.Mock.On("Create", mock.AnythingOfType("*models.Auth")).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*models.Auth)
		*arg = auth
	})

	res, err := service.Create(req)
	assert.Equal(t, nil, err)
	assert.Equal(t, auth.Name, res.Name)
}

func TestCreateFail(t *testing.T) {
	repo, service := Init()

	req := request.Auth{
		Phone: "0812345",
		Name:  "name",
		Role:  "staff",
	}

	err := errors.New("Cannot Create")
	repo.Mock.On("Create", mock.AnythingOfType("*models.Auth")).Return(err)

	res, errRes := service.Create(req)
	assert.Equal(t, errRes, err)
	assert.Equal(t, "", res.Name)
}

func TestLoginSuccess(t *testing.T) {
	repo, service := Init()

	req := request.Login{
		Phone:    "08123456",
		Password: "asdf",
	}

	auth := models.Auth{
		BaseModel: models.BaseModel{
			ID: "auth-1",
		},
		Phone:    "0812345",
		Name:     "name",
		Role:     "staff",
		Password: "asdf",
	}

	repo.Mock.On("Login", req.Phone, req.Password).Return(auth, nil)

	res, err := service.Login(req)
	assert.Equal(t, nil, err)
	assert.NotEmpty(t, res.Token)
}

func TestLoginFail(t *testing.T) {
	repo, service := Init()

	req := request.Login{
		Phone:    "08123456",
		Password: "asdf",
	}

	auth := models.Auth{}

	err := errors.New("Login Error")
	repo.Mock.On("Login", req.Phone, req.Password).Return(auth, err)

	res, errRes := service.Login(req)
	assert.Equal(t, errRes, err)
	assert.Empty(t, res.Token)
}
