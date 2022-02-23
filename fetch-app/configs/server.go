package configs

import (
	"fmt"
	"os"
)

// GetAddress will return formatted address
func GetAddress() string {
	host := ""
	if os.Getenv("APP_HOST") != "" {
		host = os.Getenv("APP_HOST")
	}

	app_port := "8081"
	if os.Getenv("APP_PORT") != "" {
		app_port = os.Getenv("APP_PORT")
	}

	return fmt.Sprintf("%s:%s", host, app_port)
}
