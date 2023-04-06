package standard

import (
	"fmt"
	"testing"
)

func TestPhpCharmaskEx(t *testing.T) {
	tests := []struct {
		input string
		want  string
		want1 bool
	}{
		{"abc", "abc", true},
		{"a..c", "abc", true},
		{"abc..cd", "abcd", true},
		{"abc..ad", "", false},
	}
	for i, tt := range tests {
		ttName := fmt.Sprintf("case-%d", i)
		t.Run(ttName, func(t *testing.T) {
			got, got1 := PhpCharmaskEx(tt.input)
			if got != tt.want {
				t.Errorf("PhpCharmaskEx() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("PhpCharmaskEx() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
