package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xamust/ekom-test/internal/interfaces"
)

var _ interfaces.Handlers = (*apiHandlersV1)(nil)

type Dependencies struct {
	APIGroup fiber.Router
	Usecases interfaces.Usecases
	Logger   interfaces.Logger
	Mappers  interfaces.Mappers
}
type apiHandlersV1 struct {
	banners interfaces.BannersHandlers
}

func NewHandlersV1(deps *Dependencies) interfaces.Handlers {
	return &apiHandlersV1{
		banners: newBannersHandlers(deps),
	}
}

func (a *apiHandlersV1) BannersHandlers() interfaces.BannersHandlers {
	return a.banners
}
