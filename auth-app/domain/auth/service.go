package auth

import (
	"tirenn/test-efishery/auth-app/dto/http/request"
	"tirenn/test-efishery/auth-app/dto/http/response"
	"tirenn/test-efishery/auth-app/utils"

	"github.com/dgrijalva/jwt-go"
)

type ServiceContract interface {
	Create(req request.Auth) (res response.Auth, err error)
	Login(req request.Login) (response.Login, error)
	GetJWT(reqToken string) (res response.GetToken, err error)
}

type Service struct {
	repository RepositoryContract
}

func NewService(repo RepositoryContract) ServiceContract {
	return &Service{
		repository: repo,
	}
}

func (s *Service) Create(req request.Auth) (res response.Auth, err error) {
	auth := request.ToAuth(req)
	auth.Password = utils.RandomString(4)

	err = s.repository.Create(&auth)
	if err != nil {
		return
	}

	res = response.ToAuth(auth)
	return
}

func (s *Service) Login(req request.Login) (res response.Login, err error) {
	auth, err := s.repository.Login(req.Phone, req.Password)
	if err != nil {
		return
	}

	token := utils.JWTAuthService().GenerateToken(auth)
	res.Token = token
	return
}

func (s *Service) GetJWT(reqToken string) (res response.GetToken, err error) {
	token, err := utils.JWTAuthService().ValidateToken(reqToken)
	if err != nil {
		return
	}

	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		res.Name = utils.InterfaceToString(claims["name"])
		res.Phone = utils.InterfaceToString(claims["phone"])
		res.Role = utils.InterfaceToString(claims["role"])
	}

	return
}
