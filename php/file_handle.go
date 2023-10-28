package php

import (
	"io"
	"io/ioutil"
	"strings"
)

type FileHandle struct {
	reader     io.Reader
	openedPath string
}

func (f *FileHandle) OpenedPath() string {
	return f.openedPath
}

func (f *FileHandle) ReadAll() (string, error) {
	data, err := ioutil.ReadAll(f.reader)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func NewFileHandleByString(code string) *FileHandle {
	return &FileHandle{
		reader: strings.NewReader(code),
	}
}
