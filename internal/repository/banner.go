package repository

import (
	"context"

	"github.com/xamust/ekom-test/internal/entity"
	"github.com/xamust/ekom-test/internal/interfaces"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ interfaces.BannersRepository = (*BannersRepository)(nil)

type BannersRepository struct {
	dbClient *mongo.Collection
	log      interfaces.Logger
}

func newBannersRepository(
	mongoDB *mongo.Database,
	log interfaces.Logger,
	collectionName string,
) interfaces.BannersRepository {
	return &BannersRepository{
		dbClient: mongoDB.Collection(collectionName),
		log:      log,
	}

}

func (b *BannersRepository) ListByFilter(
	ctx context.Context,
	filter *entity.Filter,
) ([]entity.Banner, error) {
	cursor, err := b.dbClient.Find(
		ctx,
		bson.M{
			"banner_id": filter.BannerID,
			"timestamp": bson.M{
				"$gte": filter.TSFrom,
				"$lte": filter.TSTo,
			},
		},
	)
	if err != nil {
		return []entity.Banner{}, err
	}
	defer cursor.Close(ctx)
	var result []entity.Banner
	if err := cursor.All(context.Background(), &result); err != nil {
		return []entity.Banner{}, err
	}
	return result, nil
}

func (b *BannersRepository) IncrementCount(
	ctx context.Context,
	in *entity.Banner,
) error {
	res, err := b.dbClient.UpdateOne(
		ctx,
		bson.M{
			"banner_id":   in.ID,
			"banner_name": in.Name,
			"timestamp":   in.Timestamp,
		},
		bson.M{
			"$inc": bson.M{"count": 1},
		},
		options.Update().SetUpsert(true),
	)
	if err != nil {
		return err
	}
	_ = res
	return nil
}
