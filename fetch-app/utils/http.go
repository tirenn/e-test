package utils

import (
	"os"

	"github.com/go-resty/resty/v2"
)

type ResponseAPI struct {
	StatusCode int
	Body       string
}

func Get(url string, queryParams map[string]string) (ResponseAPI, error) {
	client := resty.New()

	token := os.Getenv("token")

	response, err := client.R().
		SetQueryParams(queryParams).
		SetHeader("Accept", "application/json").
		SetAuthToken(token).
		Get(url)
	if err != nil {
		return ResponseAPI{}, err
	}

	res := ResponseAPI{
		StatusCode: response.StatusCode(),
		Body:       string(response.Body()),
	}

	return res, nil
}
