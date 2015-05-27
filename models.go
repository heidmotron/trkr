package trkr

import (
	"encoding/json"
	"time"
)

type Activity struct {
	Kind       string `json:"kind"`
	Changes    json.RawMessage
	OccurredAt time.Time `json:"occurred_at"`
}
