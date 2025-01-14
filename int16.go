// Code generated by go generate; DO NOT EDIT.
// This file was generated by scripts/models/gen_nullable_types.go
package nullable

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// Compile time validation that our types implement the expected interfaces
var (
	_ runtime.Validatable        = &Int16{}
	_ runtime.ContextValidatable = &Int16{}
	_ runtime.Validatable        = &Int16Slice{}
	_ runtime.ContextValidatable = &Int16Slice{}
)

// Int16 represents a int16 that may be null or not
// present in json at all.
type Int16 struct {
	Present bool // Present is true if key is present in json
	Valid   bool // Valid is true if value is not null and valid int16
	Value   int16
}

// Returns nil if not present or valid. Otherwise it will
// return a pointer to the value.
func (i *Int16) Ptr() *int16 {
	if i.Present && i.Valid {
		return &i.Value
	}

	return nil
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (i *Int16) UnmarshalJSON(data []byte) error {
	i.Present = true

	if bytes.Equal(data, null) {
		return nil
	}

	if err := json.Unmarshal(data, &i.Value); err != nil {
		return err
	}

	i.Valid = true
	return nil
}

// Validate implements runtime.Validateable interface from github.com/go-openapi/runtime
func (i *Int16) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate implements runtime.ContextValidatable from github.com/go-openapi/runtime
func (i *Int16) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// Int16Slice represents a []int16 that may be null or not
// present in json at all.
type Int16Slice struct {
	Present bool // Present is true if key is present in json
	Valid   bool // Valid is true if value is not null
	Value   []int16
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (i *Int16Slice) UnmarshalJSON(data []byte) error {
	i.Present = true

	if bytes.Equal(data, null) {
		return nil
	}

	if err := json.Unmarshal(data, &i.Value); err != nil {
		return err
	}

	i.Valid = true
	return nil
}

// Returns nil if not present or valid. Otherwise it will
// return a pointer to the value.
func (i *Int16Slice) Ptr() *[]int16 {
	if i.Present && i.Valid {
		return &i.Value
	}

	return nil
}

// Validate implements runtime.Validateable interface from github.com/go-openapi/runtime
func (i *Int16Slice) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate implements runtime.ContextValidatable from github.com/go-openapi/runtime
func (i *Int16Slice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
