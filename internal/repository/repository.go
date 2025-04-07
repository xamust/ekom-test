package repository

import (
	"github.com/xamust/ekom-test/config"
	"github.com/xamust/ekom-test/internal/interfaces"
	"go.mongodb.org/mongo-driver/mongo"
)

var _ interfaces.Repository = (*Repository)(nil)

type Repository struct {
	banner interfaces.BannersRepository
}

func NewRepository(
	log interfaces.Logger,
	mongoDB *mongo.Database,
	cfg *config.Config,
) *Repository {
	return &Repository{
		// TODO : think about it (collections)...
		banner: newBannersRepository(
			mongoDB,
			log,
			cfg.DB.Mongo.CollectionName,
		),
	}
}

func (r *Repository) BannersRepository() interfaces.BannersRepository {
	return r.banner
}
