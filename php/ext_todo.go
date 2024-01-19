package php

import (
	"math"
	"path/filepath"
	"strings"
)

const LongMax = math.MaxInt

func ForbidDynamicCall(ctx *Context, s string) bool {
	// todo
	return true
}

func ZendDirname(path string) string {
	if path == "" {
		return ""
	}

	/* Strip trailing slashes */
	path = strings.TrimRight(path, "/")
	if path == "" {
		/* The path only contained slashes */
		return string(filepath.Separator)
	}

	/* Strip filename */
	if pos := strings.LastIndexByte(path, '/'); pos >= 0 {
		path = path[:pos]
	} else {
		/* No slash found, therefore return '.' */
		return "."
	}

	/* Strip slashes which came before the file name */
	path = strings.TrimRight(path, "/")
	if path == "" {
		return string(filepath.Separator)
	}
	return path
}
