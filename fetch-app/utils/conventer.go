package utils

import (
	"fmt"
	"strconv"
)

func StringToFloat32(buf string, dval float32) float32 {
	res, err := strconv.ParseFloat(buf, 32)

	if err == nil {
		return float32(res)
	}

	return dval
}

func Float32ToString(val float32) string {
	return strconv.FormatFloat(
		float64(val),
		'f',
		-1,
		32,
	)
}

func InterfaceToString(i interface{}) string {
	res := fmt.Sprintf("%v", i)
	return res
}
