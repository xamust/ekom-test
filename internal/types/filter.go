package types

import "time"

type Filter struct {
	TSFrom time.Time `json:"ts_from" example:"2025-04-07T00:00:00Z"`
	TSTo   time.Time `json:"ts_to" example:"2025-04-07T23:59:00Z"`
}
