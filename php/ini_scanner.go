package php

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"
)

// IniScanCallback
type IniScanCallback interface {
	Comment(comment string)
	SectionStart(section string)
	Pair(section string, key string, value string)
}

// IniScanCallbackFunc
type IniScanCallbackFunc func(section string, key string, value string)

func (cb IniScanCallbackFunc) Comment(comment string)      {}
func (cb IniScanCallbackFunc) SectionStart(section string) {}
func (cb IniScanCallbackFunc) Pair(section string, key string, value string) {
	cb(section, key, value)
}

func IniScan(str string, cb IniScanCallback) error {
	buf := bufio.NewReader(strings.NewReader(str))

	var line string
	var err error
	var currSection string
	var lineno int
	for err == nil {
		lineno++
		line, err = buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		} else if line[0] == ';' || line[0] == '#' {
			cb.Comment(line)
		} else if section, ok := iniCutSection(line); ok {
			currSection = section
			cb.SectionStart(section)
		} else if key, value, ok := iniCutPair(line); ok {
			cb.Pair(currSection, key, value)
		} else {
			return fmt.Errorf("syntax err line: [%d] %s", lineno, line)
		}
	}
	if err == io.EOF {
		err = nil
	}
	return err
}

var iniSectionReg = regexp.MustCompile(`^\[[ \t]*([\w\.\-]+)[ \t]*\](;.*)?$`)

func iniCutSection(line string) (string, bool) {
	if matches := iniSectionReg.FindStringSubmatch(line); len(matches) > 0 {
		return matches[1], true
	}
	return "", false
}

var iniPairReg = regexp.MustCompile(`^([\w\.\-]+)\s*=\s*('(\\'|[^'])*'|"(\\"|[^"])*"|[^'";]*)(;.*)?$`)

func iniCutPair(line string) (string, string, bool) {
	if matches := iniPairReg.FindStringSubmatch(line); len(matches) > 0 {
		key := matches[1]
		val := matches[2]
		return key, val, true
	}
	return "", "", false
}
