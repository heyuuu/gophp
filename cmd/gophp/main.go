package main

import (
	"flag"
	"github.com/heyuuu/gophp/compile"
	"log"
)

func main() {
	var code string
	flag.StringVar(&code, "r", "", "code")
	flag.Parse()

	source := compile.NewSourcesByCode(code)
	compiler := compile.Compiler{}
	proj, err := compiler.Compile(source)
	if err != nil {
		log.Fatalln(err)
	}



	println(code)
}
