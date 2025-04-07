package types

import "time"

type Stats struct {
	Stats []Stat `json:"stats"`
}
type Stat struct {
	TS time.Time `json:"ts"`
	V  uint64    `json:"v"`
}
