package utils

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
)

//jwt service
type JWTService interface {
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtServices struct {
	secretKey string
}

//auth-jwt
func JWTAuthService() JWTService {
	return &jwtServices{
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	secret := os.Getenv("JWT_SECRET")
	return secret
}

func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])

		}
		return []byte(service.secretKey), nil
	})

}
