package response

import (
	"tirenn/test-efishery/fetch-app/models"
)

type Komoditas struct {
	UUID         string `json:"uuid"`
	Komoditas    string `json:"komoditas"`
	AreaProvinsi string `json:"area_provinsi"`
	AreaKota     string `json:"area_kota"`
	Size         string `json:"size"`
	Price        string `json:"price"`
	PriceUSD     string `json:"price_usd"`
	TglParsed    string `json:"tgl_parsed"`
	Timestamp    string `json:"timestamp"`
}

func ToKomoditas(k models.Komoditas) Komoditas {
	return Komoditas{
		UUID:         k.UUID,
		Komoditas:    k.Komoditas,
		AreaProvinsi: k.AreaProvinsi,
		AreaKota:     k.AreaKota,
		Size:         k.Size,
		Price:        k.Price,
		TglParsed:    k.TglParsed,
		Timestamp:    k.Timestamp,
	}
}

func ToKomoditasArray(k []models.Komoditas) []Komoditas {
	komoditasArray := []Komoditas{}

	for _, v := range k {
		komoditasArray = append(komoditasArray, ToKomoditas(v))
	}

	return komoditasArray
}

type KomoditasAggregate struct {
	AreaProvinsi string      `json:"area_provinsi"`
	MaxPrice     string      `json:"max_price"`
	MinPrice     string      `json:"min_price"`
	AveragePrice string      `json:"average_price"`
	MedianPrice  string      `json:"median_price"`
	Data         []Komoditas `json:"data"`
}
