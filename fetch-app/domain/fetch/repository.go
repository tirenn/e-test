package fetch

import (
	"encoding/json"
	"errors"
	"tirenn/test-efishery/fetch-app/models"
	"tirenn/test-efishery/fetch-app/utils"
)

type RepositoryContract interface {
	Get() (komoditas []models.Komoditas, err error)
}

type Repository struct {
	baseUrl string
}

func NewRepository(baseUrl string) RepositoryContract {
	return &Repository{
		baseUrl: baseUrl,
	}
}

func (r *Repository) Get() (komoditas []models.Komoditas, err error) {
	url := r.baseUrl + "/storages/5e1edf521073e315924ceab4/list"

	res, err := utils.Get(url, nil)
	if err != nil {
		return
	}

	if res.StatusCode != 200 && res.StatusCode != 201 {
		err = errors.New("Error Get Data")
		return
	}

	json.Unmarshal([]byte(res.Body), &komoditas)
	return komoditas, err
}
