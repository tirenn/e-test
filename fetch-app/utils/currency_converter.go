package utils

import (
	"encoding/json"
	"errors"
	"os"
	"time"

	"github.com/ReneKroon/ttlcache/v2"
)

type CurrencyConverterResponse struct {
	USDIDR float32 `json:"USD_IDR"`
}

var cache ttlcache.SimpleCache = ttlcache.NewCache()
var id string = "usdidr"

func USDToIDR() (result float32, err error) {
	url := os.Getenv("BASE_URL_CURRENCY_CONVERTER")
	if url == "" {
		url = "https://free.currconv.com/api/v7/convert"
	}

	apiKey := os.Getenv("API_KEY_CURRENCY_CONVERTER")

	queryParams := map[string]string{
		"q":       "USD_IDR",
		"compact": "ultra",
		"apiKey":  apiKey,
	}

	res, err := Get(url, queryParams)

	if res.StatusCode != 200 && res.StatusCode != 201 {
		err = errors.New("Get Data Error")
		return
	}

	currencyConverterResponse := CurrencyConverterResponse{}
	json.Unmarshal([]byte(res.Body), &currencyConverterResponse)

	cache.SetTTL(time.Duration(24 * time.Hour))
	cache.Set(id, currencyConverterResponse.USDIDR)

	result = currencyConverterResponse.USDIDR
	return
}

func GetUSDtoIDR() (result float32) {
	res, err := cache.Get(id)
	if err != nil {
		result, _ = USDToIDR()
		return
	}

	result = res.(float32)

	if result == 0 {
		result, _ = USDToIDR()
		return
	}

	return
}

func IDRToUSDConverter(idr float32) (usd float32) {
	conv := GetUSDtoIDR()
	if conv == 0 {
		usd = idr / 14000
		return
	}

	usd = idr / conv
	return
}
