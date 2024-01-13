package tests

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestConfig_runTestReal(t *testing.T) {
	srcDir := "/Users/heyu/Code/src/php-7.4.33"
	tests := []struct {
		name   string
		output string
	}{
		{"tests/run-test/test001.phpt", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// init Config
			var buf strings.Builder
			events := NewDefaultEventHandler(func(verbose int, message string) {
				buf.WriteString(message)
			})
			c := Config{SrcDir: srcDir, Events: events}

			// run function
			testFile := filepath.Join(srcDir, tt.name)
			_, err := c.runTest(0, testFile)
			if err != nil {
				t.Errorf("runTest()  error = %v", err)
				return
			}

			// check output
			output := buf.String()
			if output != tt.output {
				t.Errorf("runTest() output = %s, expect = %s", output, tt.output)
				return
			}
		})
	}
}
