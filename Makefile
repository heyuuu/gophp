GO=go
GOFMT=gofmt
BULIDPATH=./temp/

version:
	${GO} version

gp-gen:
	gp-gen -cmd gen-func

gp-lombok:
	gp-lombok -cmd generate

gophp-dev:
	${GO} build -o ${BULIDPATH} ./cmd/gophp
	${BULIDPATH}gophp -r "echo 123, 456, 'abc';"

dump-all: version
	rm -rf ./log/dump
	${GO} run ./cmd/gophp-test > ./log/dump-all.log