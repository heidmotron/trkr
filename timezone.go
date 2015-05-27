package trkr

// TimeZone represents a Tracker timezone
type TimeZone struct {
	OlsonName *string `json:"olson_name,omitempty"`
	Offset    *string `json:"offset,omitempty"`
}
