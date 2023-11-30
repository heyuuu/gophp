package php

import "strings"

const NsSeparator = '\\'

func CleanNsName(name string) string {
	if name != "" && name[0] == NsSeparator {
		name = name[1:]
	}
	return name
}

func SplitNsName(name string) (ns string, baseName string) {
	name = CleanNsName(name)
	if idx := strings.LastIndexByte(name, NsSeparator); idx >= 0 {
		return name[:idx], name[idx+1:]
	} else {
		return "", name
	}
}
