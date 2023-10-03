package php

type Sources interface {
	LoadSource(relativePath string) (string, bool)
}

const DefaultSourcePath = ""

func NewSourcesByCode(code string) Sources {
	return &baseSources{
		loadSource: func(relativePath string) (string, bool) {
			if relativePath == DefaultSourcePath {
				return code, true
			}
			return "", false
		},
	}
}

// baseSources
type baseSources struct {
	loadSource func(relativePath string) (string, bool)
}

var _ Sources = (*baseSources)(nil)

func (s baseSources) LoadSource(relativePath string) (string, bool) {
	return s.loadSource(relativePath)
}
