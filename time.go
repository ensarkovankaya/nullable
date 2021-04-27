package nullable

import (
	"bytes"
	"encoding/json"
	"time"

	"github.com/go-openapi/strfmt"
)

// Time represents a time that may be null or not
// present in json at all.
type Time struct {
	Present bool // Present is true if key is present in json
	Valid   bool // Valid is true if value is not null and valid time
	Value   time.Time
}

// Returns nil if not present or valid. Otherwise it will
// return a pointer to the value.
func (t *Time) Ptr() *time.Time {
	if t.Present && t.Valid {
		return &t.Value
	}

	return nil
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (t *Time) UnmarshalJSON(data []byte) error {
	t.Present = true

	if bytes.Equal(data, null) {
		return nil
	}

	if err := json.Unmarshal(data, &t.Value); err != nil {
		return err
	}

	t.Valid = true
	return nil
}

// Validate implements runtime.Validateable interface for go-swagger generation.
func (t *Time) Validate(formats strfmt.Registry) error {
	return nil
}
