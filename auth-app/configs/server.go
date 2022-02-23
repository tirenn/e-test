package configs

import (
	"fmt"
	"os"
)

// GetAddress will return formatted address
func GetAddress() string {
	host := os.Getenv("APP_HOST")
	if host == "" {
		host = ""
	}

	app_port := os.Getenv("APP_PORT")
	if app_port == "" {
		app_port = "8080"
	}

	return fmt.Sprintf("%s:%s", host, app_port)
}
