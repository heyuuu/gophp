package main

import (
	"flag"
	"github.com/heyuuu/gophp/php"
	"log"
)

func main() {
	var code string
	flag.StringVar(&code, "r", "", "code")
	flag.Parse()

	r := php.Default()
	err := r.RunCode("<?php " + code)
	if err != nil {
		log.Fatalln(err)
	}
}
