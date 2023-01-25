clear:
	ls | grep -v builtin | grep -v go.mod | grep -v Makefile | xargs rm -rf
