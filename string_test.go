package nullable

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
)

func TestString_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name           string
		buf            *bytes.Buffer
		expect         String
		expectErr      error
		expectedPtrNil bool
	}{
		{
			name: "null value",
			buf:  bytes.NewBufferString(`{"value":null}`),
			expect: String{
				Present: true,
			},
			expectErr:      nil,
			expectedPtrNil: true,
		},
		{
			name: "valid value",
			buf:  bytes.NewBufferString(`{"value":"string"}`),
			expect: String{
				Present: true,
				Valid:   true,
				Value:   "string",
			},
			expectErr:      nil,
			expectedPtrNil: false,
		},
		{
			name:           "empty",
			buf:            bytes.NewBufferString(`{}`),
			expect:         String{},
			expectErr:      nil,
			expectedPtrNil: true,
		},
		{
			name: "unmarshallable",
			buf:  bytes.NewBufferString(`{"value":42}`),
			expect: String{
				Present: true,
			},
			expectErr:      &json.UnmarshalTypeError{},
			expectedPtrNil: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str := struct {
				Value String `json:"value"`
			}{}

			if err := json.Unmarshal(tt.buf.Bytes(), &str); !typeMatch(tt.expectErr, err) {
				t.Fatalf("unexpected unmarshaling error: %s", err)
			}

			got := str.Value
			if got.Present != tt.expect.Present || got.Valid != tt.expect.Valid || got.Value != tt.expect.Value || got.Ptr() == nil != tt.expectedPtrNil {
				t.Errorf("expected value to be %#v got %#v", tt.expect, got)
			}
		})
	}
}

func TestStringSlice_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name      string
		buf       *bytes.Buffer
		expect    StringSlice
		expectErr error
	}{
		{
			name: "null value",
			buf:  bytes.NewBufferString(`{"value":null}`),
			expect: StringSlice{
				Present: true,
			},
			expectErr: nil,
		},
		{
			name: "valid value",
			buf:  bytes.NewBufferString(`{"value":["string", "string2"]}`),
			expect: StringSlice{
				Present: true,
				Valid:   true,
				Value:   []string{"string", "string2"},
			},
			expectErr: nil,
		},
		{
			name:      "empty",
			buf:       bytes.NewBufferString(`{}`),
			expect:    StringSlice{},
			expectErr: nil,
		},
		{
			name: "unmarshallable",
			buf:  bytes.NewBufferString(`{"value":42}`),
			expect: StringSlice{
				Present: true,
			},
			expectErr: &json.UnmarshalTypeError{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str := struct {
				Value StringSlice `json:"value"`
			}{}

			if err := json.Unmarshal(tt.buf.Bytes(), &str); !typeMatch(tt.expectErr, err) {
				t.Fatalf("unexpected unmarshaling error: %s", err)
			}

			got := str.Value
			if got.Present != tt.expect.Present || got.Valid != tt.expect.Valid || fmt.Sprintf("%v", got.Value) != fmt.Sprintf("%v", tt.expect.Value) {
				t.Errorf("expected value to be %#v got %#v", tt.expect, got)
			}
		})
	}
}
