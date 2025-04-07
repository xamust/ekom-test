package mappers

import "github.com/xamust/ekom-test/internal/interfaces"

var _ interfaces.Mappers = (*Mappers)(nil)

type Mappers struct {
	id     interfaces.IDMappers
	banner interfaces.BannerMappers
}

func NewMappers() interfaces.Mappers {
	return &Mappers{
		id:     newIDMappers(),
		banner: newBannerMappers(),
	}
}

func (m *Mappers) IDMappers() interfaces.IDMappers {
	return m.id
}

func (m *Mappers) BannerMappers() interfaces.BannerMappers {
	return m.banner
}
