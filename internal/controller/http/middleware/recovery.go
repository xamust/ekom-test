package middleware

import (
	"fmt"
	"runtime/debug"
	"strings"

	"github.com/gofiber/fiber/v2"
	fiberRecover "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/xamust/ekom-test/internal/interfaces"
)

func Recovery(log interfaces.Logger) func(ctx *fiber.Ctx) error {
	return fiberRecover.New(
		fiberRecover.Config{
			EnableStackTrace:  true,
			StackTraceHandler: logPanic(log),
		})
}

func logPanic(log interfaces.Logger) func(c *fiber.Ctx, e interface{}) {
	return func(c *fiber.Ctx, e interface{}) {
		log.Error(buildPanicMessage(c, e))
	}
}

func buildPanicMessage(ctx *fiber.Ctx, err interface{}) string {
	var result strings.Builder

	result.WriteString(ctx.IP())
	result.WriteString(" - ")
	result.WriteString(ctx.Method())
	result.WriteString(" ")
	result.WriteString(ctx.OriginalURL())
	result.WriteString(" PANIC DETECTED: ")
	result.WriteString(fmt.Sprintf("%v\n%s\n", err, debug.Stack()))

	return result.String()
}
