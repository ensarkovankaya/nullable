package nullable

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"
)

var referenceDate = time.Date(1991, 5, 23, 1, 2, 3, 4, time.UTC)

func TestTime_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name      string
		buf       *bytes.Buffer
		expect    Time
		expectErr error
	}{
		{
			name: "null value",
			buf:  bytes.NewBufferString(`{"value":null}`),
			expect: Time{
				Present: true,
			},
			expectErr: nil,
		},
		{
			name: "valid value",
			buf:  bytes.NewBufferString(`{"value":"` + referenceDate.Format(time.RFC3339Nano) + `"}`),
			expect: Time{
				Present: true,
				Valid:   true,
				Value:   referenceDate,
			},
			expectErr: nil,
		},
		{
			name:      "empty",
			buf:       bytes.NewBufferString(`{}`),
			expect:    Time{},
			expectErr: nil,
		},
		{
			name: "unmarshallable",
			buf:  bytes.NewBufferString(`{"value":42}`),
			expect: Time{
				Present: true,
			},
			expectErr: &time.ParseError{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str := struct {
				Value Time `json:"value"`
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
