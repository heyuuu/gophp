package tests

import (
	"fmt"
	"github.com/heyuuu/gophp/shim/slices"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func timeFormat(t time.Time, format string) string {
	var buf strings.Builder
	for _, r := range format {
		switch r {
		case 'Y':
			_, _ = fmt.Fprintf(&buf, "%04d", t.Year())
		case 'm':
			_, _ = fmt.Fprintf(&buf, "%02d", t.Month())
		case 'd':
			_, _ = fmt.Fprintf(&buf, "%02d", t.Day())
		case 'H':
			_, _ = fmt.Fprintf(&buf, "%02d", t.Hour())
		case 'i':
			_, _ = fmt.Fprintf(&buf, "%02d", t.Minute())
		case 's':
			_, _ = fmt.Fprintf(&buf, "%02d", t.Second())
		default:
			buf.WriteRune(r)
		}
	}
	return buf.String()
}

func sliceUnique[T comparable](slice []T) []T {
	existItems := map[T]struct{}{}
	return slices.DeleteFunc(slice, func(e T) bool {
		if _, exists := existItems[e]; exists {
			return true
		}

		existItems[e] = struct{}{}
		return false
	})
}

func readLines(file string) ([]string, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	content := string(data)
	lineCount := strings.Count(content, "\n") + 1
	lines := make([]string, 0, lineCount)
	for len(content) > 0 {
		if idx := strings.IndexByte(content, '\n'); idx >= 0 {
			lines = append(lines, content[:idx+1])
			content = content[idx+1:]
		} else {
			lines = append(lines, content)
			break
		}
	}

	return lines, nil
}

func fileExists(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}

func fileGetContents(file string) (string, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func filePutContents(file string, text string) error {
	dir := filepath.Dir(file)
	if !fileExists(dir) {
		err := os.MkdirAll(dir, 0644)
		if err != nil {
			return err
		}
	}

	return os.WriteFile(file, []byte(text), 0644)
}

func basename(path string, suffix string) string {
	base := filepath.Base(path)
	if suffix != "" && strings.HasSuffix(base, suffix) {
		base = base[:len(base)-len(suffix)]
	}
	return base
}