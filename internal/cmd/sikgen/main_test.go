package main

import (
	"strings"
	"testing"
)

func TestGenFunc(t *testing.T) {
	args := strings.Split("./sikgen -cmd gen-func -d ../../../zend", " ")
	run(args)
}
