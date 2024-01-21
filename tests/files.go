package tests

import (
	"github.com/heyuuu/gophp/shim/cmp"
	"github.com/heyuuu/gophp/shim/slices"
	"os"
	"path/filepath"
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
