package request

import "tirenn/test-efishery/auth-app/models"

type Auth struct {
	Phone string `json:"phone" validate:"numeric,min=5,max=13"`
	Name  string `json:"name" validate:"required"`
	Role  string `json:"role" validate:"required,roles"`
}

func ToAuth(r Auth) models.Auth {
	return models.Auth{
		Phone: r.Phone,
		Name:  r.Name,
		Role:  r.Role,
	}
}

type Login struct {
	Phone    string `json:"phone" validate:"numeric,min=5,max=13"`
	Password string `json:"password" validate:"required"`
}

type JWT struct {
	JWT string `json:"jwt" validate:"required"`
}
