// <<generate>>

package core

const SAPI_OPTION_NO_CHDIR = 1
const SAPI_POST_BLOCK_SIZE = 0x4000

type sapi_module_struct = _sapiModule

var sapi_module sapi_module_struct

type _sapiGlobals = sapi_globals_struct

var sapi_globals sapi_globals_struct

type SapiHeaderOpEnum = int

const (
	SAPI_HEADER_REPLACE = iota
	SAPI_HEADER_ADD
	SAPI_HEADER_DELETE
	SAPI_HEADER_DELETE_ALL
	SAPI_HEADER_SET_STATUS
)
const SAPI_HEADER_ADD SapiHeaderOpEnum = 1 << 0
const SAPI_HEADER_SENT_SUCCESSFULLY = 1
const SAPI_HEADER_DO_SEND = 2
const SAPI_HEADER_SEND_FAILED = 3
const SAPI_DEFAULT_MIMETYPE = "text/html"
const SAPI_DEFAULT_CHARSET *byte = PHP_DEFAULT_CHARSET
const SAPI_PHP_VERSION_HEADER *byte = "X-Powered-By: PHP/" + PHP_VERSION
