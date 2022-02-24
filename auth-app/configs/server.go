package configs

import (
	"fmt"
	"os"
)

// GetAddress will return formatted address
func GetAddress() string {
	host := os.Getenv("APP_HOST")
	app_port := os.Getenv("APP_PORT")

	return fmt.Sprintf("%s:%s", host, app_port)
}
