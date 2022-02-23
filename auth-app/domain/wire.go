//+build wireinject

package domain

import (
	"tirenn/test-efishery/auth-app/domain/auth"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitAuthAPI(db *gorm.DB) auth.ControllerContract {
	wire.Build(auth.NewRepository, auth.NewService, auth.NewController)
	return &auth.Controller{}
}
