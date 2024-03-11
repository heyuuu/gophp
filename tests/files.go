package tests

import (
	"cmp"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
)

func EachTestFile(dir string, handler func(file string) error) error {
	return EachTestFileEx(dir, false, handler)
}

func EachTestFileEx(dir string, cleanTmp bool, handler func(file string) error) error {
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
			err = EachTestFileEx(path, cleanTmp, handler)
		} else if strings.HasSuffix(path, ".phpt") {
			err = handler(path)
		} else if cleanTmp && strings.HasSuffix(path, ".tmp") {
			// 清理上次执行未清理的 .tmp 文件，忽略清理异常
			_ = os.Remove(path)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func EachTestCase(srcDir string, dir string, handler func(tc *TestCase) error) error {
	return EachTestCaseEx(srcDir, dir, false, handler)
}

func EachTestCaseEx(srcDir string, dir string, cleanTmp bool, handler func(tc *TestCase) error) error {
	return EachTestFileEx(dir, cleanTmp, func(file string) error {
		name, _ := filepath.Rel(srcDir, file)
		tc := NewTestCase(name, file)
		return handler(tc)
	})
}

func FindTestCases(srcDir string, dir string) ([]*TestCase, error) {
	var cases []*TestCase
	err := EachTestCase(srcDir, dir, func(tc *TestCase) error {
		cases = append(cases, tc)
		return nil
	})
	if err != nil {
		return nil, err
	}

	sortTestCases(cases)
	return cases, nil
}

func FindTestCasesInSrcDir(srcDir string, cleanTmp bool) ([]*TestCase, error) {
	var cases []*TestCase
	var subDirs = []string{"Zend", "tests", "sapi"}
	for _, subDir := range subDirs {
		dir := filepath.Join(srcDir, subDir)
		err := EachTestCaseEx(srcDir, dir, cleanTmp, func(tc *TestCase) error {
			cases = append(cases, tc)
			return nil
		})
		if err != nil {
			return nil, err
		}
	}

	sortTestCases(cases)

	return cases, nil
}

func eachTestPath(dir string, handler func(dir string)) bool {
	files, err := os.ReadDir(dir)
	if err != nil {
		return false
	}

	var isTestPath bool
	for _, file := range files {
		if file.Name() == "" || file.Name()[0] == '.' {
			continue
		}

		if file.IsDir() {
			if eachTestPath(filepath.Join(dir, file.Name()), handler) {
				isTestPath = true
			}
		} else if !isTestPath && strings.HasSuffix(file.Name(), ".phpt") {
			isTestPath = true
		}
	}
	if isTestPath {
		handler(dir)
	}
	return isTestPath
}

func FindTestPathsInSrcDir(srcDir string, includeExt bool) []string {
	var subDirs = []string{"Zend", "tests", "sapi"}
	if includeExt {
		subDirs = append(subDirs, "ext")
	}

	var paths []string
	for _, subDir := range subDirs {
		dir := filepath.Join(srcDir, subDir)
		eachTestPath(dir, func(dir string) {
			path, _ := filepath.Rel(srcDir, dir)
			paths = append(paths, path)
		})
	}

	slices.SortFunc(paths, compareFileName)

	return paths
}

func sortTestCases(cases []*TestCase) {
	slices.SortStableFunc(cases, func(c1, c2 *TestCase) int {
		return compareFileName(c1.FileName(), c2.FileName())
	})
}

func scoreFileName(fileName string) int {
	if strings.HasPrefix(fileName, "tests/run-test") {
		return 2
	} else if strings.HasPrefix(fileName, "tests") {
		return 1
	}
	return 0
}

func compareFileName(file1, file2 string) int {
	score1, score2 := scoreFileName(file1), scoreFileName(file2)
	if score1 == score2 {
		return cmp.Compare(file1, file2)
	} else {
		return -cmp.Compare(score1, score2)
	}
}

var allowSections = map[string]bool{
	"TEST":                 true,
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
	"CAPTURE_STDIO":        true,
	"STDIN":                true,
	"CGI":                  true,
	"INI":                  true,
	"ENV":                  true,
	"EXTENSIONS":           true,
	"SKIPIF":               true,
	"XFAIL":                true,
	//"XLEAK":true,
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

	err = checkFileSections(file, sections)
	if err != nil {
		return nil, err
	}

	return sections, nil
}

func checkFileSections(file string, sections map[string]string) error {
	for section, _ := range sections {
		if !allowSections[section] {
			return fmt.Errorf(`unknown section "%s"`, section)
		}
	}

	// check sections
	if existKeys(sections, "FILE", "FILEEOF", "FILE_EXTERNAL") != 1 {
		return errors.New("missing section --FILE--")
	}
	if existKey(sections, "FILEEOF") {
		sections["FILE"] = strings.TrimRight(sections["FILEEOF"], "\r\n")
		delete(sections, "FILEEOF")
	}
	for _, prefix := range []string{"FILE", "EXPECT", "EXPECTF", "EXPECTREGEX"} {
		key := prefix + "_EXTERNAL"
		if existKey(sections, key) {
			// don't allow tests to retrieve files from anywhere but this subdirectory
			path := filepath.Join(filepath.Dir(file), strings.TrimSpace(strings.ReplaceAll(sections[key], "..", "")))
			content, err := fileGetContents(path)
			if err != nil {
				return fmt.Errorf("could not load --%s-- %s", key, path)
			}

			sections[prefix] = content
			delete(sections, key)
		}
	}

	if existKeys(sections, "EXPECT", "EXPECTF", "EXPECTREGEX") != 1 {
		return errors.New("missing section --EXPECT--, --EXPECTF-- or --EXPECTREGEX--")
	}

	return nil
}
