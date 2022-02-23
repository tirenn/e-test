// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package domain

import (
	"gorm.io/gorm"
	"tirenn/test-efishery/auth-app/domain/auth"
)

// Injectors from wire.go:

func InitAuthAPI(db *gorm.DB) auth.ControllerContract {
	repositoryContract := auth.NewRepository(db)
	serviceContract := auth.NewService(repositoryContract)
	controllerContract := auth.NewController(serviceContract)
	return controllerContract
}
