package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend/types"
)

var ZendExecuteEx func(executeData *ZendExecuteData)
var ZendExecuteInternal func(executeData *ZendExecuteData, return_value *types.Zval)

/* export zend_pass_function to allow comparisons against it */

/* dedicated Zend executor functions - do not use! */

var ZEND_VM_STACK_HEADER_SLOTS *types.Zval = (ZEND_MM_ALIGNED_SIZE(b.SizeOf("struct _zend_vm_stack")) +
	ZEND_MM_ALIGNED_SIZE(b.SizeOf("zval")) - 1) / ZEND_MM_ALIGNED_SIZE(b.SizeOf("zval"))

/* services */

const ZEND_USER_OPCODE_CONTINUE = 0
const ZEND_USER_OPCODE_RETURN = 1
const ZEND_USER_OPCODE_DISPATCH = 2
const ZEND_USER_OPCODE_ENTER = 3
const ZEND_USER_OPCODE_LEAVE = 4
const ZEND_USER_OPCODE_DISPATCH_TO = 0x100

/* former zend_execute_locks.h */

type ZendFreeOp *types.Zval

const CACHE_SPECIAL = 1 << 0

/* Virtual current working directory support */
const _CONST_CODE = 0
const _TMP_CODE = 1
const _VAR_CODE = 2
const _UNUSED_CODE = 3
const _CV_CODE = 4

type IncdecT func(*types.Zval) int

var ZendPassFunction = MakeInternalFunctionSimplify(ZifPass)

const ZEND_VM_STACK_PAGE_SLOTS = 16 * 1024

var ZEND_VM_STACK_PAGE_SIZE = ZEND_VM_STACK_PAGE_SLOTS * b.SizeOf("zval")

var ZEND_FAKE_OP_ARRAY *ZendOpArray = (*ZendOpArray)(zend_intptr_t - 1)
var _zendVmStackPushCallFrameEx = ZendVmStackPushCallFrameEx
var _zendVmStackPushCallFrame = ZendVmStackPushCallFrame

const VM_SMART_OPCODES = 1
