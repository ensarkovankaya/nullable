package nullable

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"
)

func TestBool_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name      string
		buf       *bytes.Buffer
		expect    Bool
		expectErr error
	}{
		{
			name: "null value",
			buf:  bytes.NewBufferString(`{"value":null}`),
			expect: Bool{
				Present: true,
			},
			expectErr: nil,
		},
		{
			name: "valid value",
			buf:  bytes.NewBufferString(`{"value":true}`),
			expect: Bool{
				Present: true,
				Valid:   true,
				Value:   true,
			},
			expectErr: nil,
		},
		{
			name:      "empty",
			buf:       bytes.NewBufferString(`{}`),
			expect:    Bool{},
			expectErr: nil,
		},
		{
			name: "unmarshallable",
			buf:  bytes.NewBufferString(`{"value":42}`),
			expect: Bool{
				Present: true,
			},
			expectErr: &json.UnmarshalTypeError{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str := struct {
				Value Bool `json:"value"`
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

func typeMatch(expected, actual interface{}) bool {
	if expected != nil && actual != nil {
		expectedType := reflect.TypeOf(expected)
		actualType := reflect.TypeOf(actual)

		return actualType.AssignableTo(expectedType)
	} else if expected == nil && actual == nil {
		return true
	} else {
		// One of them was nil, while the other was not
		// -> never matches
		return false
	}
}
