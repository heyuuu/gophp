// <<generate>>

package zend

import (
	b "sik/builtin"
)

var ZendExecuteEx func(execute_data *ZendExecuteData)
var ZendExecuteInternal func(execute_data *ZendExecuteData, return_value *Zval)

const ZEND_VM_STACK_HEADER_SLOTS *Zval = (ZEND_MM_ALIGNED_SIZE(b.SizeOf("struct _zend_vm_stack")) + ZEND_MM_ALIGNED_SIZE(b.SizeOf("zval")) - 1) / ZEND_MM_ALIGNED_SIZE(b.SizeOf("zval"))
const ZEND_USER_OPCODE_CONTINUE = 0
const ZEND_USER_OPCODE_RETURN = 1
const ZEND_USER_OPCODE_DISPATCH = 2
const ZEND_USER_OPCODE_ENTER = 3
const ZEND_USER_OPCODE_LEAVE = 4
const ZEND_USER_OPCODE_DISPATCH_TO = 0x100

type ZendFreeOp *Zval

const CACHE_SPECIAL = 1 << 0
const ZEND_INTENSIVE_DEBUGGING = 0
const EXECUTE_DATA_D = zend_execute_data * execute_data
const EXECUTE_DATA_C EXECUTE_DATA_D = execute_data
const OPLINE_C *ZendOp = opline
const _CONST_CODE = 0
const _TMP_CODE = 1
const _VAR_CODE = 2
const _UNUSED_CODE = 3
const _CV_CODE = 4

type IncdecT func(*Zval) int

var ZendPassFunction ZendInternalFunction = ZendInternalFunction{ZEND_INTERNAL_FUNCTION, {0, 0, 0}, 0, nil, nil, nil, 0, 0, nil, ZifPass, nil, {nil, nil, nil, nil}}

const ZEND_VM_STACK_PAGE_SLOTS = 16 * 1024
const ZEND_VM_STACK_PAGE_SIZE = ZEND_VM_STACK_PAGE_SLOTS * b.SizeOf("zval")
const ZEND_FAKE_OP_ARRAY *ZendOpArray = (*ZendOpArray)(zend_intptr_t - 1)
const _zendVmStackPushCallFrameEx = ZendVmStackPushCallFrameEx
const _zendVmStackPushCallFrame = ZendVmStackPushCallFrame
const VM_SMART_OPCODES = 1
