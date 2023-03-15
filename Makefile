GOBIN=/Users/heyu/go/go1.18.2/bin/go

clear:
	ls | grep -v builtin | grep -v go.mod | grep -v Makefile | xargs rm -rf
lexer:
	go generate ./parser/zend_language_scanner.go
	go fmt ./zend/zend_language_scanner_gen.go

sikgen-func:
	${GOBIN} build ./internal/cmd/sikgen
	./sikgen gen-func -d zend

sikgen-func-clear:
	${GOBIN} build ./internal/cmd/sikgen
	./sikgen clear-func