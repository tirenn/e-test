//+build wireinject

package domain

import (
	"tirenn/test-efishery/fetch-app/domain/fetch"

	"github.com/google/wire"
)

func InitFetchAPI(baseUrl string) fetch.ControllerContract {
	wire.Build(fetch.NewRepository, fetch.NewService, fetch.NewController)
	return &fetch.Controller{}
}
