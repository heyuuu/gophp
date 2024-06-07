.PHONY: install

PHP=/opt/homebrew/bin/php

install:
	${PHP} install.php

gen-ast-node:
	${PHP} parser.php generate --mode=ast-node

gen-ast-decode:
	${PHP} parser.php generate --mode=ast-decode

gen-token-type:
	${PHP} parser.php generate --mode=token-type


