package usecases

import (
	"context"
	"fmt"
	"time"

	"github.com/xamust/ekom-test/internal/entity"
	"github.com/xamust/ekom-test/internal/interfaces"
	"github.com/xamust/ekom-test/internal/types"
)

var _ interfaces.BannersUsecases = (*BannersUsecases)(nil)

type BannersUsecases struct {
	log    interfaces.Logger
	mapper interfaces.Mappers
	repo   interfaces.Repository
}

func newBannersUsecases(deps *Dependencies) interfaces.BannersUsecases {
	return &BannersUsecases{
		log:    deps.Log,
		mapper: deps.Mapper,
		repo:   deps.Repo,
	}
}

func (b *BannersUsecases) GetByFilter(
	ctx context.Context,
	bannerID uint64,
	filter *types.Filter,
) (*types.Stats, error) {
	byFilter, err := b.repo.BannersRepository().ListByFilter(ctx, &entity.Filter{
		BannerID: bannerID,
		TSFrom:   filter.TSFrom,
		TSTo:     filter.TSTo,
	})
	if err != nil {
		return nil, err
	}
	return b.mapper.BannerMappers().ToStats(byFilter), nil
}

func (b *BannersUsecases) IncreaseBannerCount(
	ctx context.Context,
	bannerID uint64,
) error {
	return b.repo.
		BannersRepository().
		IncrementCount(
			ctx,
			&entity.Banner{
				ID:        bannerID,
				Name:      fmt.Sprintf("Banner-%d", bannerID),
				Timestamp: time.Now().Truncate(time.Minute),
			},
		)
}
