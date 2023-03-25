package zif

import (
	"io/ioutil"
	"strings"
)

func lcName(name string) string {
	var buf strings.Builder
	for i, c := range name {
		if c >= 'A' && c <= 'Z' {
			if i > 0 {
				buf.WriteByte('_')
			}
			buf.WriteRune(c - 'A' + 'a')
		} else {
			buf.WriteRune(c)
		}
	}
	return buf.String()
}

func writeFileIfChanged(fileName string, content string) (changed bool, err error) {
	existContent, err := ioutil.ReadFile(fileName)
	if err == nil && string(existContent) == content {
		return false, nil
	}

	err = ioutil.WriteFile(fileName, []byte(content), 0644)
	return true, err
}
