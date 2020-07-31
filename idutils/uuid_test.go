package idutils

import (
	"fmt"
	"testing"
)

func ExampleGetUUID() {
	fmt.Println("UUID:", GetUUID())
}

func TestGetUUID(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Test GetUUID()",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if len(GetUUID()) == 0 {
				t.Errorf("GetUUID() should not return empty UUID")
			}
		})
	}
}
