package response

import (
	"time"
	"tirenn/test-efishery/auth-app/models"
)

type Auth struct {
	Phone     string    `json:"phone"`
	Name      string    `json:"name"`
	Role      string    `json:"role"`
	Password  string    `json:"password"`
	Timestamp time.Time `json:"timestamp"`
}

func ToAuth(a models.Auth) Auth {
	return Auth{
		Phone:     a.Phone,
		Name:      a.Name,
		Role:      a.Role,
		Password:  a.Password,
		Timestamp: a.CreatedAt,
	}
}

type Login struct {
	Token string `json:"token"`
}

type GetToken struct {
	Phone string `json:"phone"`
	Name  string `json:"name"`
	Role  string `json:"role"`
}
