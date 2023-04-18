package standard

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
)

const PHP_STREAM_CLIENT_PERSISTENT = 1
const PHP_STREAM_CLIENT_ASYNC_CONNECT = 2
const PHP_STREAM_CLIENT_CONNECT = 4

// todo
var ZifStreamWrapperRegister func(executeData *zend.ZendExecuteData, return_value *types.Zval)
var ZifStreamWrapperUnregister func(executeData *zend.ZendExecuteData, return_value *types.Zval)
var ZifStreamWrapperRestore func(executeData *zend.ZendExecuteData, return_value *types.Zval)

type PhpTimeoutUll = unsigned__long__long
