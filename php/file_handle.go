package php

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
)

const CommandLineFileName = "Command line code"

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

func (f *FileHandle) Close() error {
	if c, ok := f.reader.(io.Closer); ok {
		return c.Close()
	}
	return nil
}

func NewFileHandleByFilename(file string) (*FileHandle, error) {
	fp, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	return &FileHandle{
		reader:     fp,
		openedPath: file,
	}, nil
}

func NewFileHandleByString(code string) *FileHandle {
	return &FileHandle{
		reader: strings.NewReader(code),
	}
}
