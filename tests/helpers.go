package tests

import (
	"compress/gzip"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"slices"
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

func isDir(dir string) bool {
	s, err := os.Stat(dir)
	return err == nil && s.IsDir()
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
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}

	return os.WriteFile(file, []byte(text), 0644)
}

func unlink(filename string) error {
	return os.Remove(filename)
}

func basename(path string, suffix string) string {
	base := filepath.Base(path)
	if suffix != "" && strings.HasSuffix(base, suffix) {
		base = base[:len(base)-len(suffix)]
	}
	return base
}

func existKey[K comparable, V any](m map[K]V, key K) bool {
	_, exists := m[key]
	return exists
}

func existAnyKey[K comparable, V any](m map[K]V, keys ...K) bool {
	for _, key := range keys {
		if _, exists := m[key]; exists {
			return true
		}
	}
	return false
}

func existKeys[K comparable, V any](m map[K]V, keys ...K) int {
	count := 0
	for _, key := range keys {
		if _, exists := m[key]; exists {
			count++
		}
	}
	return count
}

func strpos(s string, substr string, offset int) int {
	if offset < 0 {
		offset += len(s)
	}
	if offset < 0 || offset >= len(s) {
		return -1
	}
	if idx := strings.Index(s, substr[offset:]); idx >= 0 {
		return idx + offset
	} else {
		return -1
	}
}

func pregQuote(s string) string {
	return regexp.QuoteMeta(s)
}

func addslashes(s string) string {
	if s == "" {
		return ""
	}
	replacer := strings.NewReplacer(
		"\000", "\\0",
		`'`, `\'`,
		`"`, `\"`,
		`\`, `\\`,
	)
	return replacer.Replace(s)
}

func gzencode(str string) (string, error) {
	var buf strings.Builder

	gz := gzip.NewWriter(&buf)
	_, err := gz.Write([]byte(str))
	if err != nil {
		return "", err
	}

	err = gz.Close()
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func gzcompress(s string) (string, error) {
	// todo
	return s, nil
}

func pregMatch(rule string, s string) bool {
	reg, err := regexp.Compile(rule)
	if err != nil {
		return false
	}
	return reg.MatchString(s)
}
