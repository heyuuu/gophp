package tests

import (
	"errors"
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
)

func parseTestCase(file string, srcDir string) (*TestCase, error) {
	sections, err := parseTestCaseSections(file)
	if err != nil {
		return nil, err
	}

	shortFileName := file
	if strings.HasPrefix(file, srcDir+"/") {
		shortFileName = file[len(srcDir)+1:]
	}

	tc := &TestCase{
		File:          file,
		ShortFileName: shortFileName,
		sections:      sections,
	}

	// parse fields
	tc.TestName = strings.TrimSpace(sections["TEST"])

	if capture, ok := sections["CAPTURE_STDIO"]; ok {
		lcCapture := strings.ToLower(capture)
		tc.CaptureStdin = strings.Contains(lcCapture, "stdin")
		tc.CaptureStdout = strings.Contains(lcCapture, "stdout")
		tc.CaptureStderr = strings.Contains(lcCapture, "stderr")
	}

	return tc, nil
}

var allowSections = map[string]bool{
	"EXPECT":               true,
	"EXPECTF":              true,
	"EXPECTREGEX":          true,
	"EXPECTREGEX_EXTERNAL": true,
	"EXPECT_EXTERNAL":      true,
	"EXPECTF_EXTERNAL":     true,
	"EXPECTHEADERS":        true,
	"POST":                 true,
	"POST_RAW":             true,
	"GZIP_POST":            true,
	"DEFLATE_POST":         true,
	"PUT":                  true,
	"GET":                  true,
	"COOKIE":               true,
	"ARGS":                 true,
	"FILE":                 true,
	"FILEEOF":              true,
	"FILE_EXTERNAL":        true,
	"REDIRECTTEST":         true,
	"CAPTURE_STDIO":        true,
	"STDIN":                true,
	"CGI":                  true,
	"PHPDBG":               true,
	"INI":                  true,
	"ENV":                  true,
	"EXTENSIONS":           true,
	"SKIPIF":               true,
	"XFAIL":                true,
	"XLEAK":                true,
	"CLEAN":                true,
	"CREDITS":              true,
	"DESCRIPTION":          true,
	"CONFLICTS":            true,
	"WHITESPACE_SENSITIVE": true,
}

var regSectionTitle = regexp.MustCompile(`^--([_A-Z]+)--`)
var regSectionDone = regexp.MustCompile(`^===DONE===\s*$`)

func parseTestCaseSections(file string) (map[string]string, error) {
	lines, err := readLines(file)
	if err != nil {
		return nil, fmt.Errorf("cannot open test file: %s", file)
	} else if len(lines) == 0 {
		return nil, fmt.Errorf("empty test [%s]", file)
	} else if !strings.HasPrefix(lines[0], "--TEST--") {
		return nil, fmt.Errorf("tests must start with --TEST-- [%s]", file)
	}

	// build section map
	sections := map[string]string{"TEST": ""}
	section := "TEST"
	sectionFile := false
	sectionDone := false
	for _, line := range lines[1:] {
		// Match the beginning of a section.
		if match := regSectionTitle.FindStringSubmatch(line); len(match) > 0 {
			section = match[1]

			if sections[section] != "" {
				return nil, fmt.Errorf("duplicated %s section", section)
			}

			// check for unknown sections
			if !allowSections[section] {
				return nil, fmt.Errorf(`unknown section "%s"`, section)
			}

			sections[section] = ""
			sectionFile = section == "FILE" || section == "FILEEOF" || section == "FILE_EXTERNAL"
			sectionDone = false
			continue
		}

		// Add to the section text.
		if !sectionDone {
			sections[section] += line
		}

		// End of actual test?
		if sectionFile && regSectionDone.MatchString(line) {
			sectionDone = true
		}
	}

	// check sections
	if existKeys(sections, "PHPDBG") == 0 && existKeys(sections, "FILE", "FILEEOF", "FILE_EXTERNAL") != 1 {
		return nil, errors.New("missing section --FILE--")
	}
	if sections["FILEEOF"] != "" {
		sections["FILE"] = strings.TrimRight(sections["FILEEOF"], "\r\n")
		delete(sections, "FILEEOF")
	}
	for _, prefix := range []string{"FILE", "EXPECT", "EXPECTF", "EXPECTREGEX"} {
		key := prefix + "_EXTERNAL"
		if sections[key] != "" {
			sections[key] = filepath.Join(filepath.Dir(file), strings.TrimSpace(strings.ReplaceAll(sections[key], "..", "")))
			sections[prefix], err = fileGetContents(sections[key])
			if err == nil {
				delete(sections, key)
			} else {
				return nil, fmt.Errorf("could not load --%s-- %s", section, sections[key])
			}
		}
	}

	if existKeys(sections, "EXPECT", "EXPECTF", "EXPECTREGEX") != 1 {
		return nil, errors.New("missing section --EXPECT--, --EXPECTF-- or --EXPECTREGEX--")
	}

	return sections, nil
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
