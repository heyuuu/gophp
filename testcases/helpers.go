package main

import (
	"io"
	"os"
	"path/filepath"
)

func eachFile(dir string, deep bool, fileHandler func(string) error) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {

		path := filepath.Join(dir, file.Name())
		if file.IsDir() {
			if deep {
				err = eachFile(path, deep, fileHandler)
			}
		} else {
			err = fileHandler(path)
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func findFiles(dir string, deep bool) []string {
	var files []string
	_ = eachFile(dir, deep, func(s string) error {
		files = append(files, s)
		return nil
	})
	return files
}

func copyFile(to string, from string) (err error) {
	// check dir
	toDir := filepath.Dir(to)
	if _, err = os.Stat(toDir); os.IsNotExist(err) {
		err = os.MkdirAll(toDir, 0755)
	}
	if err != nil {
		return err
	}

	//
	in, err := os.Open(from)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(to)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	return nil
}
