clear:
	ls | grep -v builtin | grep -v go.mod | grep -v Makefile | xargs rm -rf
lexer:
	go generate ./parser/zend_language_scanner.go
	go fmt ./zend/zend_language_scanner_gen.go