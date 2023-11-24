package php

import (
	"os"
	"path/filepath"
)

func PathAbsJoin(paths ...string) string {
	for i := len(paths) - 1; i > 0; i-- {
		if filepath.IsAbs(paths[i]) {
			return filepath.Join(paths[i:]...)
		}
	}
	return filepath.Join(paths...)
}

const LINK_MAX = 10

func Readlink(p string) (string, bool) {
	p = filepath.Clean(p)
	for i := 0; i < LINK_MAX; i++ {
		if link, err := os.Readlink(p); err == nil {
			p = PathAbsJoin(filepath.Dir(p), link)
		} else {
			return p, true
		}
	}
	return "", false // link times
}
