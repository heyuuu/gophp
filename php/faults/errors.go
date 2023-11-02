package faults

import "github.com/heyuuu/gophp/php/types"

/**
 * constants and global variables
 */
const E_ERROR = 1 << 0
const E_WARNING = 1 << 1
const E_PARSE = 1 << 2
const E_NOTICE = 1 << 3
const E_CORE_ERROR = 1 << 4
const E_CORE_WARNING = 1 << 5
const E_COMPILE_ERROR = 1 << 6
const E_COMPILE_WARNING = 1 << 7
const E_USER_ERROR = 1 << 8
const E_USER_WARNING = 1 << 9
const E_USER_NOTICE = 1 << 10
const E_STRICT = 1 << 11
const E_RECOVERABLE_ERROR = 1 << 12
const E_DEPRECATED = 1 << 13
const E_USER_DEPRECATED = 1 << 14
const E_ALL = E_ERROR | E_WARNING | E_PARSE | E_NOTICE | E_CORE_ERROR | E_CORE_WARNING | E_COMPILE_ERROR | E_COMPILE_WARNING | E_USER_ERROR | E_USER_WARNING | E_USER_NOTICE | E_RECOVERABLE_ERROR | E_DEPRECATED | E_USER_DEPRECATED | E_STRICT
const E_CORE = E_CORE_ERROR | E_CORE_WARNING

func Error(typ int, message string) {
	// todo
}

func ThrowError(exceptionCe *types.Class, message string) {
	// todo
}
