package finder

import (
	"sik/script/util"
)

func DefaultProjectFinder() *Finder {
	projectRoot := "/Users/heyu/Code/sik/sik-go-gen-2/"
	finder := NewFinder(projectRoot)
	finder.AddFilter(func(file FileInfo) bool {
		if file.IsDir() {
			firstDir := file.RelativeFirstDir()
			return util.EqualsAny(firstDir, "core", "ext", "sapi", "zend")
		}
		return true
	})
	return finder
}
