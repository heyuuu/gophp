GO=go
GOFMT=gofmt
BULIDPATH=./temp/

gp-gen:
	gp-gen -cmd gen-func

gp-lombok:
	gp-lombok -cmd generate

gp-ui: gp-gen
	${GO} build -o ${BULIDPATH} ./internal/cmd/gp-ui
	${BULIDPATH}gp-ui

gophp-dev:
	${GO} build -o ${BULIDPATH} ./cmd/gophp
	${BULIDPATH}gophp -r "echo 123, 456, 'abc';"
