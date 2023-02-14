package finder

import (
	"io/fs"
	"io/ioutil"
	"log"
	"strings"
)

type FileInfo struct {
	fs.FileInfo
	Path         string
	RelativePath string
}

func (f FileInfo) RelativeFirstDir() string {
	return strings.SplitN(strings.Trim(f.RelativePath, "/"), "/", 2)[0]
}

type Filter func(file FileInfo) bool

var (
	hiddenFilter Filter = func(file FileInfo) bool {
		return !strings.HasPrefix(file.Name(), ".")
	}
	goFileFilter Filter = func(file FileInfo) bool {
		if file.IsDir() {
			return true
		} else {
			return strings.HasSuffix(file.Name(), ".go")
		}
	}
)

type Finder struct {
	root    string
	filters []Filter
}

func NewFinder(root string) *Finder {
	return &Finder{
		root: root,
		filters: []Filter{
			hiddenFilter,
			goFileFilter,
		},
	}
}

func (f *Finder) AddFilter(filter Filter) {
	f.filters = append(f.filters, filter)
}

func (f *Finder) Walk(handler func(file FileInfo)) {
	f.walkDir(f.root, handler)
}

func (f *Finder) Find() []FileInfo {
	var result []FileInfo
	f.Walk(func(file FileInfo) {
		result = append(result, file)
	})
	return result
}

func (f *Finder) walkDir(dir string, handler func(file FileInfo)) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		var fileInfo = f.newFileInfo(dir, file)
		if !f.filter(fileInfo) {
			continue
		}

		if fileInfo.IsDir() {
			f.walkDir(fileInfo.Path, handler)
		} else {
			handler(fileInfo)
		}
	}
}

func (f *Finder) filter(file FileInfo) bool {
	for _, filter := range f.filters {
		if !filter(file) {
			return false
		}
	}
	return true
}

func (f *Finder) newFileInfo(dir string, file fs.FileInfo) FileInfo {
	var path = dir + "/" + file.Name()
	var relativePath = path[len(f.root):]
	return FileInfo{FileInfo: file, Path: path, RelativePath: relativePath}
}
