package utils

import (
	"fmt"
	"strconv"
)

func InterfaceToString(i interface{}) string {
	res := fmt.Sprintf("%v", i)
	return res
}

func StringToBool(val string) bool {
	out, _ := strconv.ParseBool(val)
	return out
}
