// <<generate>>

package standard

import (
	"sik/zend"
)

const PHP_STREAM_CLIENT_PERSISTENT = 1
const PHP_STREAM_CLIENT_ASYNC_CONNECT = 2
const PHP_STREAM_CLIENT_CONNECT = 4

var ZifStreamWrapperRegister func(execute_data *zend.ZendExecuteData, return_value *zend.Zval)
var ZifStreamWrapperUnregister func(execute_data *zend.ZendExecuteData, return_value *zend.Zval)
var ZifStreamWrapperRestore func(execute_data *zend.ZendExecuteData, return_value *zend.Zval)

type PhpTimeoutUll = unsigned__long__long
