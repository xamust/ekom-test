package interfaces

import (
	"context"

	"github.com/xamust/ekom-test/internal/entity"
)

type Repository interface {
	BannersRepository() BannersRepository
}

type BannersRepository interface {
	// ListByFilter list entity.Banner from repo by filter
	ListByFilter(ctx context.Context, filter *entity.Filter) ([]entity.Banner, error)
	// IncrementCount $inc с upsert атомарная операция
	IncrementCount(ctx context.Context, in *entity.Banner) error
}
