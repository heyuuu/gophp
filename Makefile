GO=/Users/heyu/go/go1.18.2/bin/go
GOFMT=/Users/heyu/go/go1.18.2/bin/gofmt
BULIDPATH=./debug/

lexer:
	go generate ./parser/zend_language_scanner.go
	go fmt ./zend/zend_language_scanner_gen.go

sikgen:
	${GO} build -o ${BULIDPATH} ./internal/cmd/sikgen
	${BULIDPATH}sikgen -cmd gen-func
	${GOFMT} -w .

sikgen-clear:
	${GO} build -o ${BULIDPATH} ./internal/cmd/sikgen
	${BULIDPATH}sikgen -cmd clear-func

build-php:
	${GO} build -o ${BULIDPATH} ./cmd/gophp/
	${BULIDPATH}gophp