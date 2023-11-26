package php

import (
	"fmt"
	"strings"
	"testing"
)

type testParseCallback struct {
	buf strings.Builder
}

func (cb *testParseCallback) Reset()         { cb.buf.Reset() }
func (cb *testParseCallback) String() string { return cb.buf.String() }
func (cb *testParseCallback) Comment(comment string) {
	cb.buf.WriteString(fmt.Sprintf("Com: %s\n", comment))
}
func (cb *testParseCallback) SectionStart(section string) {
	cb.buf.WriteString(fmt.Sprintf("Sec: %s\n", section))
}
func (cb *testParseCallback) Pair(section string, key string, value string) {
	cb.buf.WriteString(fmt.Sprintf("K-V: Section=%s, key=%s, value=%s\n", section, key, value))
}

func TestIniParseStr(t *testing.T) {
	tests := []struct {
		name   string
		ini    string
		output string
		err    error
	}{
		{
			"1",
			"a=b", `
K-V: Section=, key=a, value=b
`,
			nil},
		{
			"section",
			`
a=b
[ part1 ]
prefix.name = value
`,
			`
K-V: Section=, key=a, value=b
Sec: part1
K-V: Section=part1, key=prefix.name, value=value
`,
			nil},
	}
	cb := &testParseCallback{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.output = strings.TrimSpace(tt.output)
			cb.Reset()
			err := IniScan(tt.ini, cb)
			result := strings.TrimSpace(cb.String())
			if result != tt.output {
				t.Errorf("IniScan() ouput = ```\n%v\n```, \nwantOutput ```\n%v\n```", result, tt.output)
			}
			if tt.err != nil && err == nil || tt.err == nil && err != nil || tt.err != nil && err != nil && tt.err.Error() != err.Error() {
				t.Errorf("IniScan() error = %v, wantErr %v", err, tt.err)
			}
		})
	}
}

func Test_iniCutPair(t *testing.T) {
	tests := []struct {
		name  string
		line  string
		key   string
		value string
		ok    bool
	}{
		{"basic-1", "a=b", "a", "b", true},
		{"basic-2", "a.b.c   =     this-is-a-value", "a.b.c", "this-is-a-value", true},
		{"quote-value-1", `a="a;b\";c";comment`, "a", `"a;b\";c"`, true},
		{"quote-value-2", `a='a;b\';c';comment`, "a", `'a;b\';c'`, true},
		{"no-value-1", `a=;comment`, "a", ``, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := iniCutPair(tt.line)
			if got != tt.key {
				t.Errorf("iniCutPair() key = %v, want %v", got, tt.key)
			}
			if got1 != tt.value {
				t.Errorf("iniCutPair() value = %v, want %v", got1, tt.value)
			}
			if got2 != tt.ok {
				t.Errorf("iniCutPair() ok = %v, want %v", got2, tt.ok)
			}
		})
	}
}
