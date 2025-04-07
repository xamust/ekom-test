package middleware

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/xamust/ekom-test/internal/interfaces"
)

func Logger(log interfaces.Logger) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		err := ctx.Next()
		log.Info(buildRequestMessage(ctx))
		return err
	}
}

func buildRequestMessage(ctx *fiber.Ctx) string {
	var result strings.Builder

	result.WriteString(ctx.IP())
	result.WriteString(" - ")
	result.WriteString(strings.ToUpper(ctx.Method()))
	result.WriteString(" ")
	result.WriteString(ctx.OriginalURL())
	result.WriteString(" - ")
	result.WriteString(strconv.Itoa(ctx.Response().StatusCode()))
	result.WriteString(" ")
	result.WriteString(strconv.Itoa(len(ctx.Response().Body())))

	return result.String()
}
