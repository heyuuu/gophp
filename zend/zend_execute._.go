package zend

import (
	"github.com/heyuuu/gophp/zend/types"
)

var ZendExecuteEx func(executeData *ZendExecuteData)
var ZendExecuteInternal func(executeData *ZendExecuteData, return_value *types.Zval)

const ZEND_USER_OPCODE_CONTINUE = 0
const ZEND_USER_OPCODE_RETURN = 1
const ZEND_USER_OPCODE_DISPATCH = 2
const ZEND_USER_OPCODE_ENTER = 3
const ZEND_USER_OPCODE_LEAVE = 4
const ZEND_USER_OPCODE_DISPATCH_TO = 0x100

/* former zend_execute_locks.h */

type ZendFreeOp = *types.Zval

const CACHE_SPECIAL = 1 << 0

/* Virtual current working directory support */
const _CONST_CODE = 0
const _TMP_CODE = 1
const _VAR_CODE = 2
const _UNUSED_CODE = 3
const _CV_CODE = 4

type IncdecT func(*types.Zval) int

var ZendPassFunction = types.MakeInternalFunctionSimplify(ZifPass)

var ZEND_FAKE_OP_ARRAY *types.ZendOpArray = (*types.ZendOpArray)(zend_intptr_t - 1)
