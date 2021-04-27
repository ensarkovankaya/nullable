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
	_ runtime.Validatable        = &Int64{}
	_ runtime.ContextValidatable = &Int64{}
	_ runtime.Validatable        = &Int64Slice{}
	_ runtime.ContextValidatable = &Int64Slice{}
)

// Int64 represents a int64 that may be null or not
// present in json at all.
type Int64 struct {
	Present bool // Present is true if key is present in json
	Valid   bool // Valid is true if value is not null and valid int64
	Value   int64
}

// Returns nil if not present or valid. Otherwise it will
// return a pointer to the value.
func (i *Int64) Ptr() *int64 {
	if i.Present && i.Valid {
		return &i.Value
	}

	return nil
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (i *Int64) UnmarshalJSON(data []byte) error {
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
func (i *Int64) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate implements runtime.ContextValidatable from github.com/go-openapi/runtime
func (i *Int64) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// Int64Slice represents a []int64 that may be null or not
// present in json at all.
type Int64Slice struct {
	Present bool // Present is true if key is present in json
	Valid   bool // Valid is true if value is not null
	Value   []int64
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (i *Int64Slice) UnmarshalJSON(data []byte) error {
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
func (i *Int64Slice) Ptr() *[]int64 {
	if i.Present && i.Valid {
		return &i.Value
	}

	return nil
}

// Validate implements runtime.Validateable interface from github.com/go-openapi/runtime
func (i *Int64Slice) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate implements runtime.ContextValidatable from github.com/go-openapi/runtime
func (i *Int64Slice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
