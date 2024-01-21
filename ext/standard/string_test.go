package standard

import (
	"github.com/heyuuu/gophp/php"
	"testing"
)

func TestPhpCharmaskEx(t *testing.T) {
	tests := []struct {
		name  string
		want  string
		want1 bool
	}{
		{"abc", "abc", true},
		{"a..c", "abc", true},
		{"abc..cd", "abcd", true},
		{"abc..ad", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := php.MockContext()
			got, got1 := PhpCharmaskEx(ctx, tt.name)
			if got != tt.want {
				t.Errorf("PhpCharmaskEx() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("PhpCharmaskEx() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestZifBin2hex(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{string([]byte{128})}, "80"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ZifBin2hex(tt.args.data); got != tt.want {
				t.Errorf("ZifBin2hex() = %v, want %v", got, tt.want)
			}
		})
	}
}
