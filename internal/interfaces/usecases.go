package interfaces

import (
	"context"

	"github.com/xamust/ekom-test/internal/types"
)

type Usecases interface {
	BannersUsecases() BannersUsecases
}

type BannersUsecases interface {
	// GetByFilter get stat
	GetByFilter(
		ctx context.Context,
		bannerID uint64,
		filter *types.Filter,
	) (*types.Stats, error)
	// IncreaseBannerCount _
	IncreaseBannerCount(ctx context.Context, bannerID uint64) error
}
