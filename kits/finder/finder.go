package finder

import (
	"io/fs"
	"path"
	"path/filepath"
)

// flags
type Mode int

const (
	ModeDefault Mode = iota
	ModeOnlyFiles
	ModeOnlyDirs
)

type Ignore int

const (
	IgnoreVcsFiles Ignore = 1 << iota
	IgnoreDotFiles
	IgnoreVcsIgnoredFiles
)

// Option
type Option func(f *Finder)

func OptOnlyFiles() Option            { return func(f *Finder) { f.mode = ModeOnlyFiles } }
func OptOnlyDirs() Option             { return func(f *Finder) { f.mode = ModeOnlyDirs } }
func OptMaxDepth(maxDepth int) Option { return func(f *Finder) { f.maxDepth = maxDepth } }

type File struct {
	Path         string
	RelativePath string
	Info         fs.FileInfo
}

// Finder
type Finder struct {
	mode     Mode
	ignore   Ignore
	maxDepth int
	path     string
}

func NewFinder(pathStr string, options ...Option) *Finder {
	pathStr = path.Clean(pathStr)
	f := &Finder{
		path:     pathStr,
		mode:     ModeDefault,
		ignore:   IgnoreVcsFiles | IgnoreDotFiles,
		maxDepth: -1,
	}
	for _, option := range options {
		option(f)
	}
	return f
}

func (f *Finder) Files() *Finder {
	f.mode = ModeOnlyFiles
	return f
}
func (f *Finder) Dirs() *Finder {
	f.mode = ModeOnlyDirs
	return f
}

func (f *Finder) Walk(fn func(File) error) error {
	return filepath.Walk(f.path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if f.mode == ModeOnlyFiles && info.IsDir() {
			return nil
		}
		if f.mode == ModeOnlyDirs && !info.IsDir() {
			return nil
		}

		base := filepath.Base(path)
		if (f.ignore&IgnoreDotFiles != 0) && base[0] == '.' {
			return fs.SkipDir
		}

		relativePath := path[len(f.path):]
		file := File{Path: path, RelativePath: relativePath, Info: info}
		return fn(file)
	})
}
