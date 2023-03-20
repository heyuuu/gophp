package zpp

/* Parameter parsing API -- andrei */

const ZEND_PARSE_PARAMS_QUIET = 1 << 1
const ZEND_PARSE_PARAMS_THROW = 1 << 2

/* Fast parameter parsing API */
const (
	Z_EXPECTED_LONG     = "int"
	Z_EXPECTED_BOOL     = "bool"
	Z_EXPECTED_STRING   = "string"
	Z_EXPECTED_ARRAY    = "array"
	Z_EXPECTED_FUNC     = "valid callback"
	Z_EXPECTED_RESOURCE = "resource"
	Z_EXPECTED_PATH     = "a valid path"
	Z_EXPECTED_OBJECT   = "object"
	Z_EXPECTED_DOUBLE   = "float"
)

const ZPP_ERROR_OK = 0
const ZPP_ERROR_FAILURE = 1
const ZPP_ERROR_WRONG_CALLBACK = 2
const ZPP_ERROR_WRONG_CLASS = 3
const ZPP_ERROR_WRONG_ARG = 4
const ZPP_ERROR_WRONG_COUNT = 5
