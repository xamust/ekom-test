package usecases

import (
	"github.com/xamust/ekom-test/internal/interfaces"
)

var _ interfaces.Usecases = (*Usecases)(nil)

type Dependencies struct {
	Log    interfaces.Logger
	Mapper interfaces.Mappers
	Repo   interfaces.Repository
}

type Usecases struct {
	banner interfaces.BannersUsecases
}

func NewUsecases(deps *Dependencies) *Usecases {
	return &Usecases{
		banner: newBannersUsecases(deps),
	}
}

func (u *Usecases) BannersUsecases() interfaces.BannersUsecases {
	return u.banner
}
