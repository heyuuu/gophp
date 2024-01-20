package standard

import (
	"testing"
)

func TestZifBin2hex(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{string([]byte{128})}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ZifBin2hex(tt.args.data); got != tt.want {
				t.Errorf("ZifBin2hex() = %v, want %v", got, tt.want)
			}
		})
	}
}
