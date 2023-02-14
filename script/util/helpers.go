package util

import (
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func MustReadFileString(filename string) string {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}

func MustWriteFileString(filename string, text string) {
	err := ioutil.WriteFile(filename, []byte(text), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func RegexReplaceAll(rule *regexp.Regexp, src string, repl func(matches []string) string) string {
	allIndexes := rule.FindAllStringSubmatchIndex(src, -1)

	var buf strings.Builder
	var pos = 0
	for _, indexes := range allIndexes {
		start := indexes[0]
		end := indexes[1]

		// 计算匹配组
		matches := make([]string, len(indexes)/2)
		for i := range matches {
			subStart := indexes[2*i]
			subEnd := indexes[2*i+1]
			if subStart < 0 {
				matches[i] = ""
			} else {
				matches[i] = src[subStart:subEnd]
			}
		}

		// 未匹配部分
		if start > pos {
			buf.WriteString(src[pos:start])
		}

		// 匹配部分
		buf.WriteString(repl(matches))

		// 偏移
		pos = end
	}
	buf.WriteString(src[pos:])

	return buf.String()
}

func EqualsAny[T comparable](value T, expected ...T) bool {
	for _, v := range expected {
		if value == v {
			return true
		}
	}
	return false
}
