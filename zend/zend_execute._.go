// <<generate>>

package zend

import (
	b "sik/builtin"
)

// Source: <Zend/zend_execute.h>

/*
   +----------------------------------------------------------------------+
   | Zend Engine                                                          |
   +----------------------------------------------------------------------+
   | Copyright (c) Zend Technologies Ltd. (http://www.zend.com)           |
   +----------------------------------------------------------------------+
   | This source file is subject to version 2.00 of the Zend license,     |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.zend.com/license/2_00.txt.                                |
   | If you did not receive a copy of the Zend license and are unable to  |
   | obtain it through the world-wide-web, please send a note to          |
   | license@zend.com so we can mail you a copy immediately.              |
   +----------------------------------------------------------------------+
   | Authors: Andi Gutmans <andi@php.net>                                 |
   |          Zeev Suraski <zeev@php.net>                                 |
   |          Dmitry Stogov <dmitry@php.net>                              |
   +----------------------------------------------------------------------+
*/

var ZendExecuteEx func(execute_data *ZendExecuteData)
var ZendExecuteInternal func(execute_data *ZendExecuteData, return_value *Zval)

/* export zend_pass_function to allow comparisons against it */

/* dedicated Zend executor functions - do not use! */

const ZEND_VM_STACK_HEADER_SLOTS *Zval = (ZEND_MM_ALIGNED_SIZE(b.SizeOf("struct _zend_vm_stack")) + ZEND_MM_ALIGNED_SIZE(b.SizeOf("zval")) - 1) / ZEND_MM_ALIGNED_SIZE(b.SizeOf("zval"))

/*
 * In general in RELEASE build ZEND_ASSERT() must be zero-cost, but for some
 * reason, GCC generated worse code, performing CSE on assertion code and the
 * following "slow path" and moving memory read operatins from slow path into
 * common header. This made a degradation for the fast path.
 * The following "#if ZEND_DEBUG" eliminates it.
 */

/* services */

const ZEND_USER_OPCODE_CONTINUE = 0
const ZEND_USER_OPCODE_RETURN = 1
const ZEND_USER_OPCODE_DISPATCH = 2
const ZEND_USER_OPCODE_ENTER = 3
const ZEND_USER_OPCODE_LEAVE = 4
const ZEND_USER_OPCODE_DISPATCH_TO = 0x100

/* former zend_execute_locks.h */

type ZendFreeOp *Zval

const CACHE_SPECIAL = 1 << 0

// Source: <Zend/zend_execute.c>

/*
   +----------------------------------------------------------------------+
   | Zend Engine                                                          |
   +----------------------------------------------------------------------+
   | Copyright (c) Zend Technologies Ltd. (http://www.zend.com)           |
   +----------------------------------------------------------------------+
   | This source file is subject to version 2.00 of the Zend license,     |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.zend.com/license/2_00.txt.                                |
   | If you did not receive a copy of the Zend license and are unable to  |
   | obtain it through the world-wide-web, please send a note to          |
   | license@zend.com so we can mail you a copy immediately.              |
   +----------------------------------------------------------------------+
   | Authors: Andi Gutmans <andi@php.net>                                 |
   |          Zeev Suraski <zeev@php.net>                                 |
   |          Dmitry Stogov <dmitry@php.net>                              |
   +----------------------------------------------------------------------+
*/

const ZEND_INTENSIVE_DEBUGGING = 0

/* Virtual current working directory support */

const EXECUTE_DATA_D = zend_execute_data * execute_data
const EXECUTE_DATA_C EXECUTE_DATA_D = execute_data
const OPLINE_C *ZendOp = opline
const _CONST_CODE = 0
const _TMP_CODE = 1
const _VAR_CODE = 2
const _UNUSED_CODE = 3
const _CV_CODE = 4

type IncdecT func(*Zval) int

var ZendPassFunction = MakeInternalFunctionSimplify(ZifPass)

const ZEND_VM_STACK_PAGE_SLOTS = 16 * 1024
const ZEND_VM_STACK_PAGE_SIZE = ZEND_VM_STACK_PAGE_SLOTS * b.SizeOf("zval")

/* this should modify object only if it's empty */

/* Utility Functions for Extensions */

/* Checks whether an array can be assigned to the reference. Returns conflicting property if
 * assignment is not possible, NULL otherwise. */

/* Checks whether an stdClass can be assigned to the reference. Returns conflicting property if
 * assignment is not possible, NULL otherwise. */

/* 1: valid, 0: invalid, -1: may be valid after type coercion */

/*
 * Stack Frame Layout (the whole stack frame is allocated at once)
 * ==================
 *
 *                             +========================================+
 * EG(current_execute_data) -> | zend_execute_data                      |
 *                             +----------------------------------------+
 *     EX_VAR_NUM(0) --------> | VAR[0] = ARG[1]                        |
 *                             | ...                                    |
 *                             | VAR[op_array->num_args-1] = ARG[N]     |
 *                             | ...                                    |
 *                             | VAR[op_array->last_var-1]              |
 *                             | VAR[op_array->last_var] = TMP[0]       |
 *                             | ...                                    |
 *                             | VAR[op_array->last_var+op_array->T-1]  |
 *                             | ARG[N+1] (extra_args)                  |
 *                             | ...                                    |
 *                             +----------------------------------------+
 */

const ZEND_FAKE_OP_ARRAY *ZendOpArray = (*ZendOpArray)(zend_intptr_t - 1)
const _zendVmStackPushCallFrameEx = ZendVmStackPushCallFrameEx
const _zendVmStackPushCallFrame = ZendVmStackPushCallFrame
const VM_SMART_OPCODES = 1
