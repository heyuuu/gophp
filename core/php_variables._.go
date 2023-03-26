package core

import (
	"sik/zend/types"
)

const PARSE_POST = 0
const PARSE_GET = 1
const PARSE_COOKIE = 2
const PARSE_STRING = 3
const PARSE_ENV = 4
const PARSE_SERVER = 5
const PARSE_SESSION = 6

/* binary-safe version */

const NUM_TRACK_VARS = 6

var PhpImportEnvironmentVariables func(array_ptr *types.Zval) = _phpImportEnvironmentVariables

/* binary-safe version */

type PostVarData = PostVarDataT
