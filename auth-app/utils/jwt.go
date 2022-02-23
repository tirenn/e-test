package utils

import (
	"fmt"
	"os"
	"time"
	"tirenn/test-efishery/auth-app/models"

	"github.com/dgrijalva/jwt-go"
)

//jwt service
type JWTService interface {
	GenerateToken(auth models.Auth) string
	ValidateToken(token string) (*jwt.Token, error)
}

type authCustomClaims struct {
	Phone string `json:"phone"`
	Name  string `json:"name"`
	Role  string `json:"role"`
	jwt.StandardClaims
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

func (service *jwtServices) GenerateToken(auth models.Auth) string {
	claims := &authCustomClaims{
		auth.Phone,
		auth.Name,
		auth.Role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])

		}
		return []byte(service.secretKey), nil
	})

}
