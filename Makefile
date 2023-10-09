GO=go
GOFMT=gofmt
BULIDPATH=./temp/

gophp-ui:
	${GO} build -o ${BULIDPATH} ./internal/cmd/gp-ui
	${BULIDPATH}gp-ui

gophp-dev:
	${GO} build -o ${BULIDPATH} ./cmd/gophp
	${BULIDPATH}gophp -r "echo 123, 456, 'abc';"
