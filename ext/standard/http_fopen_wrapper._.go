package standard

import (
	"sik/core"
)

const HTTP_HEADER_BLOCK_SIZE = 1024
const PHP_URL_REDIRECT_MAX = 20
const HTTP_HEADER_USER_AGENT = 1
const HTTP_HEADER_HOST = 2
const HTTP_HEADER_AUTH = 4
const HTTP_HEADER_FROM = 8
const HTTP_HEADER_CONTENT_LENGTH = 16
const HTTP_HEADER_TYPE = 32
const HTTP_HEADER_CONNECTION = 64
const HTTP_WRAPPER_HEADER_INIT = 1
const HTTP_WRAPPER_REDIRECTED = 2

var HttpStreamWops core.PhpStreamWrapperOps = core.MakePhpStreamWrapperOps(PhpStreamUrlWrapHttp, nil, PhpStreamHttpStreamStat, nil, nil, "http", nil, nil, nil, nil, nil)
var PhpStreamHttpWrapper core.PhpStreamWrapper = core.MakePhpStreamWrapper(&HttpStreamWops, nil, 1)
