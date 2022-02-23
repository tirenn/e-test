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

func Aggregate(k []Komoditas) (res []Komoditas) {
	id := make(map[string]bool)
	areaProvinsi := make(map[string]string)

	if len(k) > 0 {
		for _, v := range k {
			id[v.UUID] = true
			areaProvinsi[v.UUID] = v.AreaProvinsi
		}

		for i, v := range k {
			if id[v.UUID] {
				k[i].AreaProvinsi = areaProvinsi[v.UUID]
			}
		}
	}

	return k
}

func FindMax(k []Komoditas) (max Komoditas) {
	max = k[0]
	for _, d := range k {
		if d.Price > max.Price {
			max = d
		}
	}
	return max
}

func FindMin(k []Komoditas) (min Komoditas) {
	min = k[0]
	for _, d := range k {
		if d.Price < min.Price {
			min = d
		}
	}
	return min
}
