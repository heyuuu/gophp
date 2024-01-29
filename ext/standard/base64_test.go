package standard

import (
	"encoding/hex"
	"testing"
)

func TestPhpBase64Decode(t *testing.T) {
	tests := []struct {
		name   string
		str    string
		strict bool
		want   string
		want1  bool
	}{
		// base
		{"base-0", "", false, "", true},
		{"base-1", "a", false, "", true},
		{"base-2", "aa", false, "69", true},
		{"base-3", "aaa", false, "69a6", true},
		{"base-4", "aaaa", false, "69a69a", true},
		{"base-5", "aaaaa", false, "69a69a", true},
		{"base-6", "aaaaaa", false, "69a69a69", true},
		{"base-7", "aaaaaaa", false, "69a69a69a6", true},
		{"base-8", "aaaaaaaa", false, "69a69a69a69a", true},
		{"base-9", "aaaaaaaaa", false, "69a69a69a69a", true},
		// padding
		{"padding-0", "=", false, "", true},
		{"padding-1", "a=", false, "", true},
		{"padding-2", "aa=", false, "69", true},
		{"padding-3", "aaa=", false, "69a6", true},
		{"padding-4", "aaaa=", false, "69a69a", true},
		{"padding-5", "aaaaa=", false, "69a69a", true},
		{"padding-6", "aaaaaa=", false, "69a69a69", true},
		{"padding-7", "aaaaaaa=", false, "69a69a69a6", true},
		{"padding-8", "aaaaaaaa=", false, "69a69a69a69a", true},
		{"padding-9", "aaaaaaaaa=", false, "69a69a69a69a", true},
		// space
		{"space-0", " a\r \ta=", false, "69", true},
		{"space-1", "a a\r \ta=", false, "69a6", true},
		{"space-2", "aa a\r \ta=", false, "69a69a", true},
		{"space-3", "aaa a\r \ta=", false, "69a69a", true},
		{"space-4", "aaaa a\r \ta=", false, "69a69a69", true},
		{"space-5", "aaaaa a\r \ta=", false, "69a69a69a6", true},
		{"space-6", "aaaaaa a\r \ta=", false, "69a69a69a69a", true},
		{"space-7", "aaaaaaa a\r \ta=", false, "69a69a69a69a", true},
		{"space-8", "aaaaaaaa a\r \ta=", false, "69a69a69a69a69", true},
		{"space-9", "aaaaaaaaa a\r \ta=", false, "69a69a69a69a69a6", true},
		// space-in-padding
		{"space-in-padding-0", "= =", false, "", true},
		{"space-in-padding-1", "a= =", false, "", true},
		{"space-in-padding-2", "aa= =", false, "69", true},
		{"space-in-padding-3", "aaa= =", false, "69a6", true},
		{"space-in-padding-4", "aaaa= =", false, "69a69a", true},
		{"space-in-padding-5", "aaaaa= =", false, "69a69a", true},
		{"space-in-padding-6", "aaaaaa= =", false, "69a69a69", true},
		{"space-in-padding-7", "aaaaaaa= =", false, "69a69a69a6", true},
		{"space-in-padding-8", "aaaaaaaa= =", false, "69a69a69a69a", true},
		{"space-in-padding-9", "aaaaaaaaa= =", false, "69a69a69a69a", true},
		// padding-in-word
		{"padding-in-word-0", "=aaaaa", false, "69a69a", true},
		{"padding-in-word-1", "a=aaaaa", false, "69a69a69", true},
		{"padding-in-word-2", "aa=aaaaa", false, "69a69a69a6", true},
		{"padding-in-word-3", "aaa=aaaaa", false, "69a69a69a69a", true},
		{"padding-in-word-4", "aaaa=aaaaa", false, "69a69a69a69a", true},
		{"padding-in-word-5", "aaaaa=aaaaa", false, "69a69a69a69a69", true},
		{"padding-in-word-6", "aaaaaa=aaaaa", false, "69a69a69a69a69a6", true},
		{"padding-in-word-7", "aaaaaaa=aaaaa", false, "69a69a69a69a69a69a", true},
		{"padding-in-word-8", "aaaaaaaa=aaaaa", false, "69a69a69a69a69a69a", true},
		{"padding-in-word-9", "aaaaaaaaa=aaaaa", false, "69a69a69a69a69a69a69", true},
		// strict-base
		{"strict-base-0", "", true, "", true},
		{"strict-base-1", "a", true, "", false},
		{"strict-base-2", "aa", true, "69", true},
		{"strict-base-3", "aaa", true, "69a6", true},
		{"strict-base-4", "aaaa", true, "69a69a", true},
		{"strict-base-5", "aaaaa", true, "", false},
		{"strict-base-6", "aaaaaa", true, "69a69a69", true},
		{"strict-base-7", "aaaaaaa", true, "69a69a69a6", true},
		{"strict-base-8", "aaaaaaaa", true, "69a69a69a69a", true},
		{"strict-base-9", "aaaaaaaaa", true, "", false},
		// strict-padding
		{"strict-padding-0", "=", true, "", false},
		{"strict-padding-1", "a=", true, "", false},
		{"strict-padding-2", "aa=", true, "", false},
		{"strict-padding-3", "aaa=", true, "69a6", true},
		{"strict-padding-4", "aaaa=", true, "", false},
		{"strict-padding-5", "aaaaa=", true, "", false},
		{"strict-padding-6", "aaaaaa=", true, "", false},
		{"strict-padding-7", "aaaaaaa=", true, "69a69a69a6", true},
		{"strict-padding-8", "aaaaaaaa=", true, "", false},
		{"strict-padding-9", "aaaaaaaaa=", true, "", false},
		// strict-space
		{"strict-space-0", " a\r \ta=", true, "", false},
		{"strict-space-1", "a a\r \ta=", true, "69a6", true},
		{"strict-space-2", "aa a\r \ta=", true, "", false},
		{"strict-space-3", "aaa a\r \ta=", true, "", false},
		{"strict-space-4", "aaaa a\r \ta=", true, "", false},
		{"strict-space-5", "aaaaa a\r \ta=", true, "69a69a69a6", true},
		{"strict-space-6", "aaaaaa a\r \ta=", true, "", false},
		{"strict-space-7", "aaaaaaa a\r \ta=", true, "", false},
		{"strict-space-8", "aaaaaaaa a\r \ta=", true, "", false},
		{"strict-space-9", "aaaaaaaaa a\r \ta=", true, "69a69a69a69a69a6", true},
		// strict-space-in-padding
		{"strict-space-in-padding-0", "= =", true, "", false},
		{"strict-space-in-padding-1", "a= =", true, "", false},
		{"strict-space-in-padding-2", "aa= =", true, "69", true},
		{"strict-space-in-padding-3", "aaa= =", true, "", false},
		{"strict-space-in-padding-4", "aaaa= =", true, "", false},
		{"strict-space-in-padding-5", "aaaaa= =", true, "", false},
		{"strict-space-in-padding-6", "aaaaaa= =", true, "69a69a69", true},
		{"strict-space-in-padding-7", "aaaaaaa= =", true, "", false},
		{"strict-space-in-padding-8", "aaaaaaaa= =", true, "", false},
		{"strict-space-in-padding-9", "aaaaaaaaa= =", true, "", false},
		// strict-padding-in-word
		{"strict-padding-in-word-0", "=aaaaa", true, "", false},
		{"strict-padding-in-word-1", "a=aaaaa", true, "", false},
		{"strict-padding-in-word-2", "aa=aaaaa", true, "", false},
		{"strict-padding-in-word-3", "aaa=aaaaa", true, "", false},
		{"strict-padding-in-word-4", "aaaa=aaaaa", true, "", false},
		{"strict-padding-in-word-5", "aaaaa=aaaaa", true, "", false},
		{"strict-padding-in-word-6", "aaaaaa=aaaaa", true, "", false},
		{"strict-padding-in-word-7", "aaaaaaa=aaaaa", true, "", false},
		{"strict-padding-in-word-8", "aaaaaaaa=aaaaa", true, "", false},
		{"strict-padding-in-word-9", "aaaaaaaaa=aaaaa", true, "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := PhpBase64Decode(tt.str, tt.strict)
			got = hex.EncodeToString([]byte(got))
			if got != tt.want {
				t.Errorf("PhpBase64Decode() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("PhpBase64Decode() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
