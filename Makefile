clear:
	ls | grep -v runtime | grep -v builtin | grep -v go.mod | grep -v Makefile | xargs rm -rf
