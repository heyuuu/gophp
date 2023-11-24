package php

import "testing"

func TestPathAbsJoin(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{"1", []string{"a/aa", "../b", "./c"}, "a/b/c"},
		{"1", []string{"/a/aa", "../b", "./c"}, "/a/b/c"},
		{"1", []string{"a/aa", "../b", "/c"}, "/c"},
		{"1", []string{"a/aa", "/b", "/c", "d"}, "/c/d"},
		{"1", []string{"a/aa", ""}, "a/aa"},
		{"1", []string{"a", "b/"}, "a/b/"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PathAbsJoin(tt.args...); got != tt.want {
				t.Errorf("PathAbsJoin() = %v, want %v", got, tt.want)
			}
		})
	}
}
