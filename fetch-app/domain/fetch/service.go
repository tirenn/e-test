package fetch

import (
	"tirenn/test-efishery/fetch-app/dto/http/response"
	"tirenn/test-efishery/fetch-app/utils"

	"github.com/dgrijalva/jwt-go"
)

type ServiceContract interface {
	Get() (res []response.Komoditas, err error)
	GetAggregate() (res []response.Komoditas, err error)
	GetJWT(reqToken string) (res response.Token, err error)
}

type Service struct {
	repository RepositoryContract
}

func NewService(repo RepositoryContract) ServiceContract {
	return &Service{
		repository: repo,
	}
}

func (s *Service) Get() (res []response.Komoditas, err error) {
	kom, err := s.repository.Get()

	res = response.ToKomoditasArray(kom)

	for i := range res {
		res[i].PriceUSD = utils.Float32ToString(utils.IDRToUSDConverter(utils.StringToFloat32(res[i].Price, 0)))
	}

	return
}

func (s *Service) GetJWT(reqToken string) (res response.Token, err error) {
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

func (s *Service) GetAggregate() (res []response.Komoditas, err error) {
	kom, err := s.repository.Get()

	res = response.ToKomoditasArray(kom)

	for i := range res {
		res[i].PriceUSD = utils.Float32ToString(utils.IDRToUSDConverter(utils.StringToFloat32(res[i].Price, 0)))
	}

	return
}
