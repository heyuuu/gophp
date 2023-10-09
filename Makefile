GO=go
GOFMT=gofmt
BULIDPATH=./temp/

gophp-ui:
	${GO} build -o ${BULIDPATH} ./internal/cmd/gp-ui
	${BULIDPATH}gp-ui

gophp-dev:
	${GO} build -o ${BULIDPATH} ./cmd/gophp
	${BULIDPATH}gophp -r "echo 123, 456, 'abc';"

lexer:
	${GO} generate ./parser/zend_language_scanner.go
	${GOFMT} -w ./zend/zend_language_scanner_gen.go

sikgen:
	${GO} build -o ${BULIDPATH} ./internal/cmd/sikgen
	${BULIDPATH}sikgen -cmd gen-func
	${GOFMT} -w .

sikgen-clear:
	${GO} build -o ${BULIDPATH} ./internal/cmd/sikgen
	${BULIDPATH}sikgen -cmd clear-func

gophp:
	${GO} build -o ${BULIDPATH} ./cmd/gophp/
	${BULIDPATH}gophp
