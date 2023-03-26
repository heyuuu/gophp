package core

import (
	"sik/zend/types"
)

const FCGI_VERSION_1 = 1
const FCGI_MAX_LENGTH = 0xffff
const FCGI_KEEP_CONN = 1

type FcgiRole = int

const (
	FCGI_RESPONDER  = 1
	FCGI_AUTHORIZER = 2
	FCGI_FILTER     = 3
)
const (
	FCGI_DEBUG   = 1
	FCGI_NOTICE  = 2
	FCGI_WARNING = 3
	FCGI_ERROR   = 4
	FCGI_ALERT   = 5
)

type FcgiRequestType = int

const (
	FCGI_BEGIN_REQUEST     = 1
	FCGI_ABORT_REQUEST     = 2
	FCGI_END_REQUEST       = 3
	FCGI_PARAMS            = 4
	FCGI_STDIN             = 5
	FCGI_STDOUT            = 6
	FCGI_STDERR            = 7
	FCGI_DATA              = 8
	FCGI_GET_VALUES        = 9
	FCGI_GET_VALUES_RESULT = 10
)

type DcgiProtocolStatus = int
type _fcgiProtocolStatus = DcgiProtocolStatus

const (
	FCGI_REQUEST_COMPLETE = 0
	FCGI_CANT_MPX_CONN    = 1
	FCGI_OVERLOADED       = 2
	FCGI_UNKNOWN_ROLE     = 3
)

/* FastCGI client API */

type FcgiApplyFunc func(var_ *byte, var_len uint, val *byte, val_len uint, arg any)

const FCGI_HASH_TABLE_SIZE = 128
const FCGI_HASH_TABLE_MASK = FCGI_HASH_TABLE_SIZE - 1
const FCGI_HASH_SEG_SIZE = 4096

type FcgiLogger func(type_ int, fmt *byte, _ ...any)

const INADDR_NONE = uint64(-1)

var FcgiLog FcgiLogger

var FcgiMgmtVars types.Array
var IsInitialized int = 0
var IsFastcgi int = 0
var InShutdown int = 0
var AllowedClients *SaT = nil
var ClientSa SaT
