package nullable

import (
	"bytes"
	"encoding/json"
	"time"
)

// Time represents a time that may be null or not
// present in json at all.
type Time struct {
	Present bool // Present is true if key is present in json
	Valid   bool // Valid is true if value is not null and valid time
	Value   time.Time
}

// UnmarshalJSON implements json.Marshaler interface.
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
