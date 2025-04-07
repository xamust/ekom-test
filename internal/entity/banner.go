package entity

import "time"

type Banner struct {
	ID        uint64    `bson:"banner_id"`
	Name      string    `bson:"banner_name"`
	Timestamp time.Time `bson:"timestamp"`
	Count     uint64    `bson:"count"`
}
