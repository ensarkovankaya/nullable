package nullable

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
)

func TestInt_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name      string
		buf       *bytes.Buffer
		expect    Int
		expectErr error
	}{
		{
			name: "null value",
			buf:  bytes.NewBufferString(`{"value":null}`),
			expect: Int{
				Present: true,
			},
			expectErr: nil,
		},
		{
			name: "valid value",
			buf:  bytes.NewBufferString(`{"value":1}`),
			expect: Int{
				Present: true,
				Valid:   true,
				Value:   1,
			},
			expectErr: nil,
		},
		{
			name:      "empty",
			buf:       bytes.NewBufferString(`{}`),
			expect:    Int{},
			expectErr: nil,
		},
		{
			name: "unmarshallable",
			buf:  bytes.NewBufferString(`{"value":"wat"}`),
			expect: Int{
				Present: true,
			},
			expectErr: &json.UnmarshalTypeError{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str := struct {
				Value Int `json:"value"`
			}{}

			if err := json.Unmarshal(tt.buf.Bytes(), &str); !typeMatch(tt.expectErr, err) {
				t.Fatalf("unexpected unmarshaling error: %s", err)
			}

			got := str.Value
			if got.Present != tt.expect.Present || got.Valid != tt.expect.Valid || got.Value != tt.expect.Value {
				t.Errorf("expected value to be %#v got %#v", tt.expect, got)
			}
		})
	}
}

func TestIntSlice_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name      string
		buf       *bytes.Buffer
		expect    IntSlice
		expectErr error
	}{
		{
			name: "null value",
			buf:  bytes.NewBufferString(`{"value":null}`),
			expect: IntSlice{
				Present: true,
			},
			expectErr: nil,
		},
		{
			name: "valid value",
			buf:  bytes.NewBufferString(`{"value":[1, 2]}`),
			expect: IntSlice{
				Present: true,
				Valid:   true,
				Value:   []int64{1, 2},
			},
			expectErr: nil,
		},
		{
			name:      "empty",
			buf:       bytes.NewBufferString(`{}`),
			expect:    IntSlice{},
			expectErr: nil,
		},
		{
			name: "unmarshallable",
			buf:  bytes.NewBufferString(`{"value":"wat"}`),
			expect: IntSlice{
				Present: true,
			},
			expectErr: &json.UnmarshalTypeError{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str := struct {
				Value IntSlice `json:"value"`
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
