package compile

import "io/ioutil"

// sources
type Source struct {
	RelativePath string
	Content      string
}

type Sources interface {
	EachSource(func(Source) bool) (bool, error)
	LoadSource(relativePath string) *Source
	LoadSourceByClass(className string) []Source
}

func NewSourcesByCode(code string) Sources {
	return &baseSources{
		eachSource: func(f func(Source) bool) (bool, error) {
			return f(Source{"__main__", code}), nil
		},
	}
}

func NewSourcesByFile(file string) (Sources, error) {
	code, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return &baseSources{
		eachSource: func(f func(Source) bool) (bool, error) {
			return f(Source{"__main__", string(code)}), nil
		},
	}, nil
}

type baseSources struct {
	eachSource        func(func(Source) bool) (bool, error)
	loadSource        func(relativePath string) *Source
	loadSourceByClass func(className string) []Source
}

func (s *baseSources) EachSource(f func(Source) bool) (bool, error) {
	if s.eachSource == nil {
		return true, nil
	}
	return s.eachSource(f)
}

func (s *baseSources) LoadSource(relativePath string) *Source {
	if s.loadSource == nil {
		return nil
	}
	return s.loadSource(relativePath)
}

func (s *baseSources) LoadSourceByClass(className string) []Source {
	if s.loadSourceByClass == nil {
		return nil
	}
	return s.loadSourceByClass(className)
}
