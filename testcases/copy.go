package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var srcDir, distDir string

	sc := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter php src dir:")
	if sc.Scan() {
		srcDir = sc.Text()
	}
	fmt.Println("Enter copy dist dir:")
	if sc.Scan() {
		distDir = sc.Text()
	}

	fmt.Printf("Start: srcDir=%s, distDir=%s\n", srcDir, distDir)
	doCopy(srcDir, distDir)
	fmt.Println("Finish.")
}

func doCopy(srcDir string, distDir string) {
	err := eachFile(srcDir, func(file string) error {
		if !strings.HasSuffix(file, ".phpt") {
			return nil
		}
		if !strings.HasPrefix(file, srcDir) {
			return nil
		}

		distFile := filepath.Join(distDir, file[len(srcDir):])
		return copyFile(distFile, file)
	})
	if err != nil {
		log.Panicln(err)
	}
}

func eachFile(dir string, fileHandler func(string) error) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {

		path := filepath.Join(dir, file.Name())
		if file.IsDir() {
			err = eachFile(path, fileHandler)
		} else {
			err = fileHandler(path)
		}
		if err != nil {
			return err
		}
	}

	return nil
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
