package argparse

import "log"

/* Parameter parsing API -- andrei */

const ZEND_PARSE_PARAMS_QUIET = 1 << 1
const ZEND_PARSE_PARAMS_THROW = 1 << 2

/* Fast parameter parsing API */

type ZendExpectedType int

func (t ZendExpectedType) String() string {
	switch t {
	case Z_EXPECTED_LONG:
		return "int"
	case Z_EXPECTED_BOOL:
		return "bool"
	case Z_EXPECTED_STRING:
		return "string"
	case Z_EXPECTED_ARRAY:
		return "array"
	case Z_EXPECTED_FUNC:
		return "array"
	case Z_EXPECTED_RESOURCE:
		return "array"
	case Z_EXPECTED_PATH:
		return "array"
	case Z_EXPECTED_OBJECT:
		return "array"
	case Z_EXPECTED_DOUBLE:
		return "array"
	default:
		log.Printf("unexpected ZendExpectedType: %d\n", t)
		return ""
	}
}

const (
	Z_EXPECTED_LONG ZendExpectedType = iota
	Z_EXPECTED_BOOL
	Z_EXPECTED_STRING
	Z_EXPECTED_ARRAY
	Z_EXPECTED_FUNC
	Z_EXPECTED_RESOURCE
	Z_EXPECTED_PATH
	Z_EXPECTED_OBJECT
	Z_EXPECTED_DOUBLE
)

const ZPP_ERROR_OK = 0
const ZPP_ERROR_FAILURE = 1
const ZPP_ERROR_WRONG_CALLBACK = 2
const ZPP_ERROR_WRONG_CLASS = 3
const ZPP_ERROR_WRONG_ARG = 4
const ZPP_ERROR_WRONG_COUNT = 5
