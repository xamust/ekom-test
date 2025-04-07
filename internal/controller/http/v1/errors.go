package v1

import "github.com/gofiber/fiber/v2"

type ErrorResponse struct {
	Error string `json:"error" example:"internal error"`
}

func errorResponse(ctx *fiber.Ctx, code int, msg string) error {
	return ctx.Status(code).JSON(ErrorResponse{msg})
}
