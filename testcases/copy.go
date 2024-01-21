package main

import (
	"bufio"
	"fmt"
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
	err := eachFile(srcDir, true, func(file string) error {
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
