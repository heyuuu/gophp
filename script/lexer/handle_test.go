package lexer

import (
	"io/ioutil"
	"log"
	"testing"
)

func TestHandle(t *testing.T) {
	var srcFile = "/Users/heyu/Code/sik/sik-go-gen-2-parser/parser/zend_language_scanner.l"
	var distFile = "/Users/heyu/Code/sik/sik-go-gen-2-parser/parser/zend_language_scanner_gen.l"
	var funcFile = "/Users/heyu/Code/sik/sik-go-gen-2-parser/parser/zend_language_scanner_func.h"
	var text = MustReadFileString(srcFile)
	var result, functions = handle(text)
	MustWriteFileString(distFile, result)
	MustWriteFileString(funcFile, functions)
}

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
