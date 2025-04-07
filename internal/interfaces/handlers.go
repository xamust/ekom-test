package interfaces

import "github.com/gofiber/fiber/v2"

type Handlers interface {
	BannersHandlers() BannersHandlers
}

type BannersHandlers interface {
	Stats(ctx *fiber.Ctx) error
	Counter(ctx *fiber.Ctx) error
}
