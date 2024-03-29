package tests

import (
	"errors"
	"fmt"
	"github.com/heyuuu/gophp/shim/cmp"
	"github.com/heyuuu/gophp/shim/slices"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func EachTestFile(dir string, deep bool, handle func(file string) error) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.Name() == "" || file.Name()[0] == '.' {
			continue
		}

		path := filepath.Join(dir, file.Name())
		if file.IsDir() {
			if deep {
				err = EachTestFile(path, deep, handle)
			}
		} else {
			if strings.HasSuffix(path, ".phpt") {
				err = handle(path)
			}
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func FindTestFiles(dir string, deep bool) ([]string, error) {
	var files []string
	err := EachTestFile(dir, deep, func(file string) error {
		files = append(files, file)
		return nil
	})
	if err != nil {
		return nil, err
	}

	sortTestFiles(files, dir)
	return files, nil
}

func sortTestFiles(files []string, srcDir string) {
	runTestDir := filepath.Join(srcDir, "tests/run-test")
	testDir := filepath.Join(srcDir, "tests")
	scorer := func(file string) int {
		if strings.HasPrefix(file, runTestDir) {
			return 2
		} else if strings.HasPrefix(file, testDir) {
			return 1
		}
		return 0
	}
	slices.SortStableFunc(files, func(file1, file2 string) int {
		score1, score2 := scorer(file1), scorer(file2)
		if score1 == score2 {
			return cmp.Compare(file1, file2)
		} else {
			return -cmp.Compare(score1, score2)
		}
	})
}

func ParseTestFile(name string, file string) (*TestCase, error) {
	sections, err := parseTestFileSections(file)
	if err != nil {
		return nil, err
	}
	return NewTestCase(name, file, sections), nil
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

func parseTestFileSections(file string) (map[string]string, error) {
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
