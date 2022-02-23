package fetch

import (
	"tirenn/test-efishery/fetch-app/dto/http/response"
	"tirenn/test-efishery/fetch-app/utils"

	"github.com/dgrijalva/jwt-go"

	"github.com/montanaflynn/stats"
)

type ServiceContract interface {
	Get() (res []response.Komoditas, err error)
	GetAggregate() (res []response.KomoditasAggregate, err error)
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
		res[i].PriceUSD = utils.Float64ToString(utils.IDRToUSDConverter(utils.StringToFloat64(res[i].Price, 0)))
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

func (s *Service) GetAggregate() (res []response.KomoditasAggregate, err error) {
	kom, err := s.repository.Get()

	data := response.ToKomoditasArray(kom)

	for i := range data {
		data[i].PriceUSD = utils.Float64ToString(utils.IDRToUSDConverter(utils.StringToFloat64(data[i].Price, 0)))
	}

	res = s.getAggregate(data)

	return
}

func (s *Service) getAggregate(data []response.Komoditas) []response.KomoditasAggregate {
	res := []response.KomoditasAggregate{}

	if len(data) > 0 {
		res = s.getProvinsi(data)

		for _, v := range data {
			for i, a := range res {
				if a.AreaProvinsi == v.AreaProvinsi {
					res[i].Data = append(res[i].Data, v)
					break
				}
			}
		}

		s.setAggregate(res)
	}

	return res
}

func (s *Service) getProvinsi(data []response.Komoditas) []response.KomoditasAggregate {
	res := []response.KomoditasAggregate{}
	for _, v := range data {
		if !s.containsProvinsi(res, v.AreaProvinsi) {
			r := response.KomoditasAggregate{}
			r.AreaProvinsi = v.AreaProvinsi
			res = append(res, r)
		}
	}

	return res
}

func (s *Service) containsProvinsi(data []response.KomoditasAggregate, provinsi string) bool {
	for _, a := range data {
		if a.AreaProvinsi == provinsi {
			return true
		}
	}
	return false
}

func (s *Service) setAggregate(k []response.KomoditasAggregate) {
	for i := range k {
		data := []float64{}
		for _, d := range k[i].Data {
			data = append(data, utils.StringToFloat64(d.Price, 0))
		}

		k[i].MaxPrice = s.getMaxPrice(data)
		k[i].MinPrice = s.getMinPrice(data)
		k[i].AveragePrice = s.getAvgPrice(data)
		k[i].MedianPrice = s.getMedianPrice(data)
	}
}

func (s *Service) getMaxPrice(data []float64) string {
	max, _ := stats.Max(data)
	return utils.Float64ToString(max)
}

func (s *Service) getMinPrice(data []float64) string {
	min, _ := stats.Min(data)
	return utils.Float64ToString(min)
}

func (s *Service) getAvgPrice(data []float64) string {
	total := float64(0)
	for _, v := range data {
		total += v
	}

	avg := total / float64(len(data))
	return utils.Float64ToString(avg)
}

func (s *Service) getMedianPrice(data []float64) string {
	median, _ := stats.Median(data)
	return utils.Float64ToString(median)
}
