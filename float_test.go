package nullable

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
)

func TestFloat_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name           string
		buf            *bytes.Buffer
		expect         Float64
		expectErr      error
		expectedPtrNil bool
	}{
		{
			name: "null value",
			buf:  bytes.NewBufferString(`{"value":null}`),
			expect: Float64{
				Present: true,
			},
			expectErr:      nil,
			expectedPtrNil: true,
		},
		{
			name: "valid value",
			buf:  bytes.NewBufferString(`{"value":1.1}`),
			expect: Float64{
				Present: true,
				Valid:   true,
				Value:   1.1,
			},
			expectErr:      nil,
			expectedPtrNil: false,
		},
		{
			name:           "empty",
			buf:            bytes.NewBufferString(`{}`),
			expect:         Float64{},
			expectErr:      nil,
			expectedPtrNil: true,
		},
		{
			name: "unmarshallable",
			buf:  bytes.NewBufferString(`{"value":"wat"}`),
			expect: Float64{
				Present: true,
			},
			expectErr:      &json.UnmarshalTypeError{},
			expectedPtrNil: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str := struct {
				Value Float64 `json:"value"`
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

func TestFloatSlice_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name      string
		buf       *bytes.Buffer
		expect    Float64Slice
		expectErr error
	}{
		{
			name: "null value",
			buf:  bytes.NewBufferString(`{"value":null}`),
			expect: Float64Slice{
				Present: true,
			},
			expectErr: nil,
		},
		{
			name: "valid value",
			buf:  bytes.NewBufferString(`{"value":[1.1, 1.2]}`),
			expect: Float64Slice{
				Present: true,
				Valid:   true,
				Value:   []float64{1.1, 1.2},
			},
			expectErr: nil,
		},
		{
			name:      "empty",
			buf:       bytes.NewBufferString(`{}`),
			expect:    Float64Slice{},
			expectErr: nil,
		},
		{
			name: "unmarshallable",
			buf:  bytes.NewBufferString(`{"value":"wat"}`),
			expect: Float64Slice{
				Present: true,
			},
			expectErr: &json.UnmarshalTypeError{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str := struct {
				Value Float64Slice `json:"value"`
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
