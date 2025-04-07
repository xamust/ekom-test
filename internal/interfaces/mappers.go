package interfaces

import (
	"github.com/xamust/ekom-test/internal/entity"
	"github.com/xamust/ekom-test/internal/types"
)

type Mappers interface {
	IDMappers() IDMappers
	BannerMappers() BannerMappers
}

type IDMappers interface {
	FromString(in string) (uint64, error)
}

type BannerMappers interface {
	ToStats(in []entity.Banner) *types.Stats
}
