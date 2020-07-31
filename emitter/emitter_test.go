package emitter

import "testing"

func ExampleEmitter() {
	Emitter("Hello Emitter")
}

func TestEmitter(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test Emitter()",
			args: args{
				s: "test args",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Emitter(tt.args.s)
		})
	}
}
