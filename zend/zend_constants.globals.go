// <<generate>>

package zend

const CONST_CS = 1 << 0
const CONST_PERSISTENT = 1 << 1
const CONST_CT_SUBST = 1 << 2
const CONST_NO_FILE_CACHE = 1 << 3
const PHP_USER_CONSTANT = 0x7fffff
const ZEND_GET_CONSTANT_NO_DEPRECATION_CHECK = 0x1000

const ZEND_CONSTANT_DTOR DtorFuncT = FreeZendConstant
const IS_CONSTANT_VISITED_MARK = 0x80
