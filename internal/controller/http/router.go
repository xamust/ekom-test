package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/xamust/ekom-test/config"
	_ "github.com/xamust/ekom-test/docs"
	"github.com/xamust/ekom-test/internal/controller/http/middleware"
	v1 "github.com/xamust/ekom-test/internal/controller/http/v1"
	"github.com/xamust/ekom-test/internal/interfaces"
	"github.com/xamust/ekom-test/internal/mappers"
)

// NewRouter new router with fiber
// Swagger spec:
// @title ekom-test app
// @description Сервис, который будет считать клики и собирать их в поминутную статистику
// @version 1.0
// @host localhost:8080
// @BasePath /api/v1
func NewRouter(
	app *fiber.App,
	cfg *config.Config,
	usecases interfaces.Usecases,
	log interfaces.Logger) {
	app.Use(
		cors.New(),
		middleware.Logger(log),
		middleware.Recovery(log),
	)

	// Swagger
	if cfg.Swagger.Enabled {
		app.Get("/swagger/*", swagger.HandlerDefault)
	}

	// Routers v1
	apiV1Group := app.Group("/api/v1")

	v1.NewHandlersV1(&v1.Dependencies{
		APIGroup: apiV1Group,
		Usecases: usecases,
		Logger:   log,
		Mappers:  mappers.NewMappers(),
	})
}
