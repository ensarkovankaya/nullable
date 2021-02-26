package nullable

import (
	"fmt"
	"testing"

	"github.com/go-openapi/runtime"
)

func TestInt16ImplementsRuntimeValidatable(t *testing.T) {

	tests := []interface{}{
		&Bool{},
		&BoolSlice{},
		&String{},
		&StringSlice{},
		&Int{},
		&IntSlice{},
		&Int16{},
		&Int16Slice{},
		&Int32{},
		&Int32Slice{},
		&Int64{},
		&Int64Slice{},
		&Float32{},
		&Float32Slice{},
		&Float64{},
		&Float64Slice{},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%T", tt), func(t *testing.T) {

			_, ok := interface{}(tt).(runtime.Validatable)

			if !ok {
				t.Errorf("expected %T to implement github.com/go-openapi/runtime runtime.Validatable", tt)
			}

			_, ok = interface{}(tt).(runtime.ContextValidatable)

			if !ok {
				t.Errorf("expected %T to implement github.com/go-openapi/runtime runtime.ContextValidatable", tt)
			}
		})
	}
}
