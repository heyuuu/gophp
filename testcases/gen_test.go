package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"testing"
)

var phpCaseRoot = "/Users/heyu/Code/sik/gophp/testcases"

type dirInfo struct {
	dir   string
	files []string
}

func TestGen(t *testing.T) {
	w, err := os.Create("php_test.go")
	if err != nil {
		log.Panicln(err)
	}

	fmt.Fprint(w, `
package main

import "testing"
`)
	handleDir(w, filepath.Join(phpCaseRoot, "php/tests"))
	handleDir(w, filepath.Join(phpCaseRoot, "php/Zend"))

	w.Close()
}

func handleDir(w io.Writer, root string) {
	var dirs []dirInfo
	eachDirFiles(root, func(dir string, filePaths []string) {
		dirs = append(dirs, dirInfo{dir: dir, files: filePaths})
	})
	sort.Slice(dirs, func(i, j int) bool {
		return dirs[i].dir < dirs[j].dir
	})

	for _, info := range dirs {
		genCode(w, info.dir, info.files)
	}
}

func makeName(dir string) string {
	var buf strings.Builder
	words := regexp.MustCompile(`[^\w]+`).Split(dir[len(phpCaseRoot):], -1)
	for _, word := range words {
		if word == "" {
			continue
		}
		buf.WriteString(strings.ToUpper(word[:1]))
		buf.WriteString(word[1:])
	}
	return buf.String()
}

func genCode(w io.Writer, dir string, files []string) {
	name := makeName(dir)

	fmt.Fprintf(w, `
func Test%s(t *testing.T) {
	tests := []struct {
		name string
	}{`, name)
	for _, file := range files {
		fmt.Fprintf(w, `
		{"%s"},`, file[len(phpCaseRoot)+1:])

	}

	fmt.Fprint(w, `
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt.name)
		})
	}
}
`)
}

func eachDirFiles(dir string, handler func(dir string, filePaths []string)) {
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Panicf("ReadDir fail")
	}

	var filePaths []string
	for _, file := range files {
		path := filepath.Join(dir, file.Name())
		if file.IsDir() {
			eachDirFiles(path, handler)
		} else {
			filePaths = append(filePaths, path)
		}
	}
	if len(filePaths) > 0 {
		handler(dir, filePaths)
	}
}
