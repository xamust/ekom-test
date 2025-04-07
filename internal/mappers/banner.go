package mappers

import (
	"github.com/xamust/ekom-test/internal/entity"
	"github.com/xamust/ekom-test/internal/interfaces"
	"github.com/xamust/ekom-test/internal/types"
)

var _ interfaces.BannerMappers = (*BannerMapper)(nil)

type BannerMapper struct{}

func newBannerMappers() interfaces.BannerMappers {
	return &BannerMapper{}
}
func (b *BannerMapper) ToStats(in []entity.Banner) *types.Stats {
	result := make([]types.Stat, 0, len(in))
	for _, banner := range in {
		result = append(result, b.toStat(banner))
	}
	return &types.Stats{
		Stats: result,
	}
}

func (b *BannerMapper) toStat(in entity.Banner) types.Stat {
	return types.Stat{
		TS: in.Timestamp,
		V:  in.Count,
	}
}
