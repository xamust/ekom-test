package entity

import "time"

type Filter struct {
	BannerID uint64    `json:"banner_id"`
	TSFrom   time.Time `json:"ts_from"`
	TSTo     time.Time `json:"ts_to"`
}
