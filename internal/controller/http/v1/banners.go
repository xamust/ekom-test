package v1

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/xamust/ekom-test/internal/interfaces"
	"github.com/xamust/ekom-test/internal/types"
)

const (
	bannerID = "bannerID"
)

var _ interfaces.BannersHandlers = (*BannersHandlers)(nil)

type BannersHandlers struct {
	log      interfaces.Logger
	mapper   interfaces.Mappers
	usecases interfaces.Usecases
}

func newBannersHandlers(deps *Dependencies) interfaces.BannersHandlers {
	b := &BannersHandlers{
		log:      deps.Logger,
		mapper:   deps.Mappers,
		usecases: deps.Usecases,
	}
	bannersGroup := deps.APIGroup.Group("/banners")
	bannersGroup.Get("/counter/:bannerID", b.Counter)
	bannersGroup.Post("/stats/:bannerID", b.Stats)
	return b
}

// Stats godoc
// @Summary     Статистика показов по bannerID за указанный промежуток времени
// @Description Banners Stat
// @ID          Banners_stat
// @Tags  	    banners
// @Accept      json
// @Produce     json
// @Param bannerID path int true "banner id"
// @Param request body types.Filter true "stat"
// @Success     200 {object} types.Stats
// @Failure     400 {object} ErrorResponse
// @Failure     500 {object} ErrorResponse
// @Router      /banners/stats/{bannerID} [post]
func (b *BannersHandlers) Stats(ctx *fiber.Ctx) error {
	id, err := b.mapper.
		IDMappers().
		FromString(ctx.Params(bannerID))
	if err != nil {
		b.log.Error(err, "op", "http - v1 - Stats - Banners")
		return errorResponse(ctx, fiber.StatusBadRequest, "incorrect banner ID")
	}
	var request types.Filter
	if err := ctx.BodyParser(&request); err != nil {
		b.log.Error(err, "op", "http - v1 - Stats - Banners")
		return errorResponse(ctx, http.StatusBadRequest, "invalid request body")
	}
	stats, err := b.usecases.
		BannersUsecases().
		GetByFilter(
			ctx.Context(),
			id,
			&request)
	if err != nil {
		b.log.Error(err, "op", "http - v1 - Stats - Banners")
		return errorResponse(ctx, http.StatusInternalServerError, "internal error")
	}
	return ctx.Status(http.StatusOK).JSON(stats)
}

// Counter godoc
// @Summary     Счетчик показов баннера по bannerID
// @Description Banner counter inc
// @ID          Banners_counter
// @Tags  	    banners
// @Accept      json
// @Produce     json
// @Param bannerID path string true "banner id"
// @Success     204
// @Failure     500 {object} ErrorResponse
// @Router      /banners/counter/{bannerID} [get]
func (b *BannersHandlers) Counter(ctx *fiber.Ctx) error {
	id, err := b.mapper.
		IDMappers().
		FromString(ctx.Params(bannerID))
	if err != nil {
		b.log.Error(err, "op", "http - v1 - Counter - Banners")
		return errorResponse(ctx, fiber.StatusBadRequest, "incorrect banner ID")
	}
	if err := b.usecases.
		BannersUsecases().
		IncreaseBannerCount(ctx.Context(), id); err != nil {
		b.log.Error(err, "op", "http - v1 - Counter - Banners")
		return errorResponse(ctx, fiber.StatusInternalServerError, "internal error")
	}
	return ctx.Status(http.StatusNoContent).JSON(fiber.Map{})
}
