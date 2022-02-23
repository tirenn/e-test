package utils

import (
	"fmt"
	"strconv"
)

func StringToFloat64(buf string, dval float64) float64 {
	res, err := strconv.ParseFloat(buf, 32)

	if err == nil {
		return float64(res)
	}

	return dval
}

func Float64ToString(val float64) string {
	return strconv.FormatFloat(
		float64(val),
		'f',
		-1,
		64,
	)
}

func InterfaceToString(i interface{}) string {
	res := fmt.Sprintf("%v", i)
	return res
}
