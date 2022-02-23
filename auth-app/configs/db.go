package configs

import (
	"log"
	"os"

	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	// "kata.ai/chatcommerce/inventory-service/models"
	"tirenn/test-efishery/auth-app/models"
	"tirenn/test-efishery/auth-app/utils"
)

// InitDB will initialize DB connection
func InitDB() *gorm.DB {
	is_debug := logger.Error
	if utils.StringToBool(os.Getenv("DB_DEBUG")) {
		is_debug = logger.Info
	}

	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Millisecond, // Slow SQL threshold
			LogLevel:      is_debug,         // Log level
			Colorful:      true,             // Disable color
		},
	)

	db, err := gorm.Open(sqlite.Open("./auth.db"), &gorm.Config{
		Logger: dbLogger,
	})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&models.Auth{},
	)

	return db
}
