package main

import (
	"github.com/heyuuu/gophp/php"
)

func doCli() {
	engine := php.NewEngine()
	engine.Start()
}
