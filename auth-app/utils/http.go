package utils

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
)

type ResponseAPI struct {
	StatusCode int
	Body       string
}

func Post(url string, queryParams map[string]string, data interface{}) (ResponseAPI, error) {
	client := resty.New()

	reqBody, err := json.Marshal(data)
	if err != nil {
		return ResponseAPI{}, err
	}

	response, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetQueryParams(queryParams).
		SetBody(string(reqBody)).
		Post(url)
	if err != nil {
		return ResponseAPI{}, err
	}

	res := ResponseAPI{
		StatusCode: response.StatusCode(),
		Body:       string(response.Body()),
	}

	return res, nil
}
