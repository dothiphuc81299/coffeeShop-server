package model

import (
	"encoding/json"
	"time"
)

const dateISOFormat = "2006-01-02T15:04:05.000Z"

// TimeResponse ...
type TimeResponse struct {
	Time time.Time
}

// UnmarshalJSON ...
func (t *TimeResponse) UnmarshalJSON(b []byte) error {
	if string(b) == "" || string(b) == "\"\"" {
		return nil
	}
	return json.Unmarshal(b, &t.Time)
}

// MarshalJSON ...
func (t TimeResponse) MarshalJSON() ([]byte, error) {
	if t.Time.IsZero() {
		return json.Marshal("")
	}
	return json.Marshal(t.Time.Format(dateISOFormat))
}

// TimeResponseInit ...
func TimeResponseInit(t time.Time) TimeResponse {
	return TimeResponse{Time: t}
}

// TimeResponsePointInit ...
func TimeResponsePointInit(t time.Time) *TimeResponse {
	return &TimeResponse{Time: t}
}

// IsAfter ...
func (t *TimeResponse) IsAfter(c TimeResponse) bool {
	return t.Time.After(c.Time)
}
