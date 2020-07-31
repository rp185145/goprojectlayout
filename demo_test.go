package goprojectlayout

import "testing"

func ExampleDemo() {
	Demo()
}

func TestDemo(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test Demo()",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Demo()
		})
	}
}
