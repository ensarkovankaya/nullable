package nullable

import (
	"bytes"
	"encoding/json"
	"github.com/go-openapi/strfmt"
)

var null = []byte("null")

// String represents a string that may be null or not
// present in json at all.
type String struct {
	Present bool // Present is true if key is present in json
	Valid   bool // Valid is true if value is not null and valid string
	Value   string
}

// Returns nil if not present or valid. Otherwise it will
// return a pointer to the value.
func (s *String) Ptr() *string {
	if s.Present && s.Valid {
		return &s.Value
	}

	return nil
}

// UnmarshalJSON implements json.Marshaler interface.
func (s *String) UnmarshalJSON(data []byte) error {
	s.Present = true

	if bytes.Equal(data, null) {
		return nil
	}

	if err := json.Unmarshal(data, &s.Value); err != nil {
		return err
	}

	s.Valid = true
	return nil
}

// Validate implements runtime.Validateable interface for go-swagger generation.
func (s *String) Validate(formats strfmt.Registry) error {
	return nil
}

// StringSlice represents a []string that may be null or not
// present in json at all.
type StringSlice struct {
	Present bool // Present is true if key is present in json
	Valid   bool // Valid is true if value is not null
	Value   []string
}

// UnmarshalJSON implements json.Marshaler interface.
func (s *StringSlice) UnmarshalJSON(data []byte) error {
	s.Present = true

	if bytes.Equal(data, null) {
		return nil
	}

	if err := json.Unmarshal(data, &s.Value); err != nil {
		return err
	}

	s.Valid = true
	return nil
}
