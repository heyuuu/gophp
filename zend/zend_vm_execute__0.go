// <<generate>>

package zend

import g "sik/runtime/grammar"

// Source: <Zend/zend_vm_execute.h>

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

var ZendUserOpcodeHandlers []UserOpcodeHandlerT = []UserOpcodeHandlerT{UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil), UserOpcodeHandlerT(nil)}
var ZendUserOpcodes []ZendUchar = []ZendUchar{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, 125, 126, 127, 128, 129, 130, 131, 132, 133, 134, 135, 136, 137, 138, 139, 140, 141, 142, 143, 144, 145, 146, 147, 148, 149, 150, 151, 152, 153, 154, 155, 156, 157, 158, 159, 160, 161, 162, 163, 164, 165, 166, 167, 168, 169, 170, 171, 172, 173, 174, 175, 176, 177, 178, 179, 180, 181, 182, 183, 184, 185, 186, 187, 188, 189, 190, 191, 192, 193, 194, 195, 196, 197, 198, 199, 200, 201, 202, 203, 204, 205, 206, 207, 208, 209, 210, 211, 212, 213, 214, 215, 216, 217, 218, 219, 220, 221, 222, 223, 224, 225, 226, 227, 228, 229, 230, 231, 232, 233, 234, 235, 236, 237, 238, 239, 240, 241, 242, 243, 244, 245, 246, 247, 248, 249, 250, 251, 252, 253, 254, 255}

// #define SPEC_START_MASK       0x0000ffff

// #define SPEC_EXTRA_MASK       0xfffc0000

// #define SPEC_RULE_OP1       0x00010000

// #define SPEC_RULE_OP2       0x00020000

// #define SPEC_RULE_OP_DATA       0x00040000

// #define SPEC_RULE_RETVAL       0x00080000

// #define SPEC_RULE_QUICK_ARG       0x00100000

// #define SPEC_RULE_SMART_BRANCH       0x00200000

// #define SPEC_RULE_COMMUTATIVE       0x00800000

// #define SPEC_RULE_ISSET       0x01000000

var ZendSpecHandlers *uint32
var ZendOpcodeHandlers *any
var ZendHandlersCount int

// #define zend_vm_get_opcode_handler_func       zend_vm_get_opcode_handler

// #define VM_TRACE(op)

// #define VM_TRACE_START()

// #define VM_TRACE_END()

// #define ZEND_OPCODE_HANDLER_ARGS       zend_execute_data * execute_data

// #define ZEND_OPCODE_HANDLER_ARGS_PASSTHRU       execute_data

// #define ZEND_OPCODE_HANDLER_ARGS_DC       , ZEND_OPCODE_HANDLER_ARGS

// #define ZEND_OPCODE_HANDLER_ARGS_PASSTHRU_CC       , ZEND_OPCODE_HANDLER_ARGS_PASSTHRU

// #define ZEND_OPCODE_HANDLER_RET       int

// #define ZEND_VM_TAIL_CALL(call) return call

// #define ZEND_VM_CONTINUE() return 0

// #define ZEND_VM_RETURN() return - 1

// #define ZEND_VM_HOT

// #define ZEND_VM_COLD       ZEND_COLD ZEND_OPT_SIZE

type OpcodeHandlerT func(execute_data *ZendExecuteData) int

// #define DCL_OPLINE

// #define OPLINE       EX ( opline )

// #define USE_OPLINE       const zend_op * opline = EX ( opline ) ;

// #define LOAD_OPLINE()

// #define LOAD_OPLINE_EX()

// #define LOAD_NEXT_OPLINE() ZEND_VM_INC_OPCODE ( )

// #define SAVE_OPLINE()

// #define SAVE_OPLINE_EX()

// #define HANDLE_EXCEPTION() LOAD_OPLINE ( ) ; ZEND_VM_CONTINUE ( )

// #define HANDLE_EXCEPTION_LEAVE() LOAD_OPLINE ( ) ; ZEND_VM_LEAVE ( )

// #define ZEND_VM_ENTER_EX() return 1

// #define ZEND_VM_ENTER() return 1

// #define ZEND_VM_LEAVE() return 2

// #define ZEND_VM_INTERRUPT() ZEND_VM_TAIL_CALL ( zend_interrupt_helper_SPEC ( ZEND_OPCODE_HANDLER_ARGS_PASSTHRU ) ) ;

// #define ZEND_VM_LOOP_INTERRUPT() zend_interrupt_helper_SPEC ( ZEND_OPCODE_HANDLER_ARGS_PASSTHRU ) ;

// #define ZEND_VM_DISPATCH(opcode,opline) ZEND_VM_TAIL_CALL ( ( ( opcode_handler_t ) zend_vm_get_opcode_handler_func ( opcode , opline ) ) ( ZEND_OPCODE_HANDLER_ARGS_PASSTHRU ) ) ;

func zend_add_helper_SPEC(op_1 *Zval, op_2 *Zval, execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	if op_1.GetTypeInfo() == 0 {
		op_1 = _zvalUndefinedOp1(execute_data)
	}
	if op_2.GetTypeInfo() == 0 {
		op_2 = _zvalUndefinedOp2(execute_data)
	}
	AddFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op_1, op_2)
	if (opline.GetOp1Type() & (1<<1 | 1<<2)) != 0 {
		ZvalPtrDtorNogc(op_1)
	}
	if (opline.GetOp2Type() & (1<<1 | 1<<2)) != 0 {
		ZvalPtrDtorNogc(op_2)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func zend_sub_helper_SPEC(op_1 *Zval, op_2 *Zval, execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	if op_1.GetTypeInfo() == 0 {
		op_1 = _zvalUndefinedOp1(execute_data)
	}
	if op_2.GetTypeInfo() == 0 {
		op_2 = _zvalUndefinedOp2(execute_data)
	}
	SubFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op_1, op_2)
	if (opline.GetOp1Type() & (1<<1 | 1<<2)) != 0 {
		ZvalPtrDtorNogc(op_1)
	}
	if (opline.GetOp2Type() & (1<<1 | 1<<2)) != 0 {
		ZvalPtrDtorNogc(op_2)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func zend_mul_helper_SPEC(op_1 *Zval, op_2 *Zval, execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	if op_1.GetTypeInfo() == 0 {
		op_1 = _zvalUndefinedOp1(execute_data)
	}
	if op_2.GetTypeInfo() == 0 {
		op_2 = _zvalUndefinedOp2(execute_data)
	}
	MulFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op_1, op_2)
	if (opline.GetOp1Type() & (1<<1 | 1<<2)) != 0 {
		ZvalPtrDtorNogc(op_1)
	}
	if (opline.GetOp2Type() & (1<<1 | 1<<2)) != 0 {
		ZvalPtrDtorNogc(op_2)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func zend_mod_by_zero_helper_SPEC(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	ZendThrowExceptionEx(ZendCeDivisionByZeroError, 0, "Modulo by zero")
	(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
	return 0
}
func zend_mod_helper_SPEC(op_1 *Zval, op_2 *Zval, execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	if op_1.GetTypeInfo() == 0 {
		op_1 = _zvalUndefinedOp1(execute_data)
	}
	if op_2.GetTypeInfo() == 0 {
		op_2 = _zvalUndefinedOp2(execute_data)
	}
	ModFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op_1, op_2)
	if (opline.GetOp1Type() & (1<<1 | 1<<2)) != 0 {
		ZvalPtrDtorNogc(op_1)
	}
	if (opline.GetOp2Type() & (1<<1 | 1<<2)) != 0 {
		ZvalPtrDtorNogc(op_2)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func zend_shift_left_helper_SPEC(op_1 *Zval, op_2 *Zval, execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	if op_1.GetTypeInfo() == 0 {
		op_1 = _zvalUndefinedOp1(execute_data)
	}
	if op_2.GetTypeInfo() == 0 {
		op_2 = _zvalUndefinedOp2(execute_data)
	}
	ShiftLeftFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op_1, op_2)
	if (opline.GetOp1Type() & (1<<1 | 1<<2)) != 0 {
		ZvalPtrDtorNogc(op_1)
	}
	if (opline.GetOp2Type() & (1<<1 | 1<<2)) != 0 {
		ZvalPtrDtorNogc(op_2)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func zend_shift_right_helper_SPEC(op_1 *Zval, op_2 *Zval, execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	if op_1.GetTypeInfo() == 0 {
		op_1 = _zvalUndefinedOp1(execute_data)
	}
	if op_2.GetTypeInfo() == 0 {
		op_2 = _zvalUndefinedOp2(execute_data)
	}
	ShiftRightFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op_1, op_2)
	if (opline.GetOp1Type() & (1<<1 | 1<<2)) != 0 {
		ZvalPtrDtorNogc(op_1)
	}
	if (opline.GetOp2Type() & (1<<1 | 1<<2)) != 0 {
		ZvalPtrDtorNogc(op_2)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func zend_is_equal_helper_SPEC(op_1 *Zval, op_2 *Zval, execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	if op_1.GetTypeInfo() == 0 {
		op_1 = _zvalUndefinedOp1(execute_data)
	}
	if op_2.GetTypeInfo() == 0 {
		op_2 = _zvalUndefinedOp2(execute_data)
	}
	CompareFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op_1, op_2)
	if (opline.GetOp1Type() & (1<<1 | 1<<2)) != 0 {
		ZvalPtrDtorNogc(op_1)
	}
	if (opline.GetOp2Type() & (1<<1 | 1<<2)) != 0 {
		ZvalPtrDtorNogc(op_2)
	}
	if EG.GetException() != nil {
		return 0
	}
	if (*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())).GetValue().GetLval() == 0 {
		if (opline + 1).GetOpcode() == 44 {
			execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
			if EG.GetVmInterrupt() != 0 {
				return zend_interrupt_helper_SPEC(execute_data)
			}
			return 0
		} else if (opline + 1).GetOpcode() == 43 {
			execute_data.SetOpline(opline + 2)
			return 0
		}
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	} else {
		if (opline + 1).GetOpcode() == 44 {
			execute_data.SetOpline(opline + 2)
			return 0
		} else if (opline + 1).GetOpcode() == 43 {
			execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
			if EG.GetVmInterrupt() != 0 {
				return zend_interrupt_helper_SPEC(execute_data)
			}
			return 0
		}
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
}
func zend_is_not_equal_helper_SPEC(op_1 *Zval, op_2 *Zval, execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	if op_1.GetTypeInfo() == 0 {
		op_1 = _zvalUndefinedOp1(execute_data)
	}
	if op_2.GetTypeInfo() == 0 {
		op_2 = _zvalUndefinedOp2(execute_data)
	}
	CompareFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op_1, op_2)
	if (opline.GetOp1Type() & (1<<1 | 1<<2)) != 0 {
		ZvalPtrDtorNogc(op_1)
	}
	if (opline.GetOp2Type() & (1<<1 | 1<<2)) != 0 {
		ZvalPtrDtorNogc(op_2)
	}
	if EG.GetException() != nil {
		return 0
	}
	if (*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())).GetValue().GetLval() != 0 {
		if (opline + 1).GetOpcode() == 44 {
			execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
			if EG.GetVmInterrupt() != 0 {
				return zend_interrupt_helper_SPEC(execute_data)
			}
			return 0
		} else if (opline + 1).GetOpcode() == 43 {
			execute_data.SetOpline(opline + 2)
			return 0
		}
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	} else {
		if (opline + 1).GetOpcode() == 44 {
			execute_data.SetOpline(opline + 2)
			return 0
		} else if (opline + 1).GetOpcode() == 43 {
			execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
			if EG.GetVmInterrupt() != 0 {
				return zend_interrupt_helper_SPEC(execute_data)
			}
			return 0
		}
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
}
func zend_is_smaller_helper_SPEC(op_1 *Zval, op_2 *Zval, execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	if op_1.GetTypeInfo() == 0 {
		op_1 = _zvalUndefinedOp1(execute_data)
	}
	if op_2.GetTypeInfo() == 0 {
		op_2 = _zvalUndefinedOp2(execute_data)
	}
	CompareFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op_1, op_2)
	if (opline.GetOp1Type() & (1<<1 | 1<<2)) != 0 {
		ZvalPtrDtorNogc(op_1)
	}
	if (opline.GetOp2Type() & (1<<1 | 1<<2)) != 0 {
		ZvalPtrDtorNogc(op_2)
	}
	if EG.GetException() != nil {
		return 0
	}
	if (*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())).GetValue().GetLval() < 0 {
		if (opline + 1).GetOpcode() == 44 {
			execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
			if EG.GetVmInterrupt() != 0 {
				return zend_interrupt_helper_SPEC(execute_data)
			}
			return 0
		} else if (opline + 1).GetOpcode() == 43 {
			execute_data.SetOpline(opline + 2)
			return 0
		}
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	} else {
		if (opline + 1).GetOpcode() == 44 {
			execute_data.SetOpline(opline + 2)
			return 0
		} else if (opline + 1).GetOpcode() == 43 {
			execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
			if EG.GetVmInterrupt() != 0 {
				return zend_interrupt_helper_SPEC(execute_data)
			}
			return 0
		}
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
}
func zend_is_smaller_or_equal_helper_SPEC(op_1 *Zval, op_2 *Zval, execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	if op_1.GetTypeInfo() == 0 {
		op_1 = _zvalUndefinedOp1(execute_data)
	}
	if op_2.GetTypeInfo() == 0 {
		op_2 = _zvalUndefinedOp2(execute_data)
	}
	CompareFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op_1, op_2)
	if (opline.GetOp1Type() & (1<<1 | 1<<2)) != 0 {
		ZvalPtrDtorNogc(op_1)
	}
	if (opline.GetOp2Type() & (1<<1 | 1<<2)) != 0 {
		ZvalPtrDtorNogc(op_2)
	}
	if EG.GetException() != nil {
		return 0
	}
	if (*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())).GetValue().GetLval() <= 0 {
		if (opline + 1).GetOpcode() == 44 {
			execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
			if EG.GetVmInterrupt() != 0 {
				return zend_interrupt_helper_SPEC(execute_data)
			}
			return 0
		} else if (opline + 1).GetOpcode() == 43 {
			execute_data.SetOpline(opline + 2)
			return 0
		}
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	} else {
		if (opline + 1).GetOpcode() == 44 {
			execute_data.SetOpline(opline + 2)
			return 0
		} else if (opline + 1).GetOpcode() == 43 {
			execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
			if EG.GetVmInterrupt() != 0 {
				return zend_interrupt_helper_SPEC(execute_data)
			}
			return 0
		}
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
}
func zend_bw_or_helper_SPEC(op_1 *Zval, op_2 *Zval, execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	if op_1.GetTypeInfo() == 0 {
		op_1 = _zvalUndefinedOp1(execute_data)
	}
	if op_2.GetTypeInfo() == 0 {
		op_2 = _zvalUndefinedOp2(execute_data)
	}
	BitwiseOrFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op_1, op_2)
	if (opline.GetOp1Type() & (1<<1 | 1<<2)) != 0 {
		ZvalPtrDtorNogc(op_1)
	}
	if (opline.GetOp2Type() & (1<<1 | 1<<2)) != 0 {
		ZvalPtrDtorNogc(op_2)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func zend_bw_and_helper_SPEC(op_1 *Zval, op_2 *Zval, execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	if op_1.GetTypeInfo() == 0 {
		op_1 = _zvalUndefinedOp1(execute_data)
	}
	if op_2.GetTypeInfo() == 0 {
		op_2 = _zvalUndefinedOp2(execute_data)
	}
	BitwiseAndFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op_1, op_2)
	if (opline.GetOp1Type() & (1<<1 | 1<<2)) != 0 {
		ZvalPtrDtorNogc(op_1)
	}
	if (opline.GetOp2Type() & (1<<1 | 1<<2)) != 0 {
		ZvalPtrDtorNogc(op_2)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func zend_bw_xor_helper_SPEC(op_1 *Zval, op_2 *Zval, execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	if op_1.GetTypeInfo() == 0 {
		op_1 = _zvalUndefinedOp1(execute_data)
	}
	if op_2.GetTypeInfo() == 0 {
		op_2 = _zvalUndefinedOp2(execute_data)
	}
	BitwiseXorFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op_1, op_2)
	if (opline.GetOp1Type() & (1<<1 | 1<<2)) != 0 {
		ZvalPtrDtorNogc(op_1)
	}
	if (opline.GetOp2Type() & (1<<1 | 1<<2)) != 0 {
		ZvalPtrDtorNogc(op_2)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func zend_this_not_in_object_context_helper_SPEC(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	ZendThrowError(nil, "Using $this when not in object context")
	if (opline + 1).GetOpcode() == 137 {
		if ((opline + 1).GetOp1Type() & (1<<1 | 1<<2)) != 0 {
			ZvalPtrDtorNogc((*Zval)((*byte)(execute_data) + int((opline + 1).GetOp1().GetVar())))
		}
	}
	if (opline.GetOp2Type() & (1<<1 | 1<<2)) != 0 {
		ZvalPtrDtorNogc((*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar())))
	}
	if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
	}
	return 0
}
func zend_undefined_function_helper_SPEC(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var function_name *Zval
	function_name = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	ZendThrowError(nil, "Call to undefined function %s()", function_name.GetValue().GetStr().GetVal())
	return 0
}
func ZEND_ASSIGN_STATIC_PROP_OP_SPEC_HANDLER(execute_data *ZendExecuteData) int {
	/* This helper actually never will receive IS_VAR as second op, and has the same handling for VAR and TMP in the first op, but for interoperability with the other binary_assign_op helpers, it is necessary to "include" it */

	var opline *ZendOp = execute_data.GetOpline()
	var free_op_data ZendFreeOp
	var prop *Zval
	var value *Zval
	var prop_info *ZendPropertyInfo
	var ref *ZendReference
	if ZendFetchStaticPropertyAddress(&prop, &prop_info, (opline+1).GetExtendedValue(), 2, 0, opline, execute_data) != SUCCESS {
		assert(EG.GetException() != nil)
		if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
		}
		if ((opline + 1).GetOp1Type() & (1<<1 | 1<<2)) != 0 {
			ZvalPtrDtorNogc((*Zval)((*byte)(execute_data) + int((opline + 1).GetOp1().GetVar())))
		}
		return 0
	}
	value = _getOpDataZvalPtrR((opline + 1).GetOp1Type(), (opline + 1).GetOp1(), &free_op_data, execute_data, opline)
	for {
		if prop.GetType() == 10 {
			ref = prop.GetValue().GetRef()
			prop = &(*prop).value.GetRef().GetVal()
			if ref.GetSources().GetPtr() != nil {
				ZendBinaryAssignOpTypedRef(ref, value, opline, execute_data)
				break
			}
		}
		if prop_info.GetType() != 0 {

			/* special case for typed properties */

			ZendBinaryAssignOpTypedProp(prop_info, prop, value, opline, execute_data)

			/* special case for typed properties */

		} else {
			ZendBinaryOp(prop, prop, value, opline)
		}
		break
	}
	if opline.GetResultType() != 0 {
		var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var _z2 *Zval = prop
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if (_t & 0xff00) != 0 {
			ZendGcAddref(&_gc.gc)
		}
	}
	if free_op_data != nil {
		ZvalPtrDtorNogc(free_op_data)
	}

	/* assign_static_prop has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}
func ZEND_PRE_INC_STATIC_PROP_SPEC_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var prop *Zval
	var prop_info *ZendPropertyInfo
	if ZendFetchStaticPropertyAddress(&prop, &prop_info, opline.GetExtendedValue(), 2, 0, opline, execute_data) != SUCCESS {
		if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
		}
		return 0
	}
	ZendPreIncdecPropertyZval(prop, g.Cond(prop_info.GetType() != 0, prop_info, nil), opline, execute_data)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}

/* No specialization for op_types (CONST|TMPVAR|CV, UNUSED|CONST|VAR) */

func ZEND_POST_INC_STATIC_PROP_SPEC_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var prop *Zval
	var prop_info *ZendPropertyInfo
	if ZendFetchStaticPropertyAddress(&prop, &prop_info, opline.GetExtendedValue(), 2, 0, opline, execute_data) != SUCCESS {
		if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
		}
		return 0
	}
	ZendPostIncdecPropertyZval(prop, g.Cond(prop_info.GetType() != 0, prop_info, nil), opline, execute_data)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}

/* No specialization for op_types (CONST|TMPVAR|CV, UNUSED|CONST|VAR) */

func zend_fetch_static_prop_helper_SPEC(type_ int, execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var prop *Zval
	if ZendFetchStaticPropertyAddress(&prop, nil, opline.GetExtendedValue() & ^3, type_, opline.GetExtendedValue()&3, opline, execute_data) != SUCCESS {
		assert(EG.GetException() != nil || type_ == 3)
		prop = &EG.uninitialized_zval
	}
	if type_ == 0 || type_ == 3 {
		var _z3 *Zval = prop
		if (_z3.GetTypeInfo() & 0xff00) != 0 {
			if (_z3.GetTypeInfo() & 0xff) == 10 {
				_z3 = &(*_z3).value.GetRef().GetVal()
				if (_z3.GetTypeInfo() & 0xff00) != 0 {
					ZvalAddrefP(_z3)
				}
			} else {
				ZvalAddrefP(_z3)
			}
		}
		var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var _z2 *Zval = _z3
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).GetValue().SetZv(prop)
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(13)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}

/* No specialization for op_types (CONST|TMPVAR|CV, UNUSED|CLASS_FETCH|CONST|VAR) */

func ZEND_FETCH_STATIC_PROP_R_SPEC_HANDLER(execute_data *ZendExecuteData) int {
	return zend_fetch_static_prop_helper_SPEC(0, execute_data)
}

/* No specialization for op_types (CONST|TMPVAR|CV, UNUSED|CLASS_FETCH|CONST|VAR) */

func ZEND_FETCH_STATIC_PROP_W_SPEC_HANDLER(execute_data *ZendExecuteData) int {
	return zend_fetch_static_prop_helper_SPEC(1, execute_data)
}

/* No specialization for op_types (CONST|TMPVAR|CV, UNUSED|CLASS_FETCH|CONST|VAR) */

func ZEND_FETCH_STATIC_PROP_RW_SPEC_HANDLER(execute_data *ZendExecuteData) int {
	return zend_fetch_static_prop_helper_SPEC(2, execute_data)
}

/* No specialization for op_types (CONST|TMPVAR|CV, UNUSED|CLASS_FETCH|CONST|VAR) */

func ZEND_FETCH_STATIC_PROP_FUNC_ARG_SPEC_HANDLER(execute_data *ZendExecuteData) int {
	var fetch_type int = g.Cond((execute_data.GetCall().GetThis().GetTypeInfo()&1<<31) != 0, 1, 0)
	return zend_fetch_static_prop_helper_SPEC(fetch_type, execute_data)
}

/* No specialization for op_types (CONST|TMPVAR|CV, UNUSED|CLASS_FETCH|CONST|VAR) */

func ZEND_FETCH_STATIC_PROP_UNSET_SPEC_HANDLER(execute_data *ZendExecuteData) int {
	return zend_fetch_static_prop_helper_SPEC(5, execute_data)
}

/* No specialization for op_types (CONST|TMPVAR|CV, UNUSED|CLASS_FETCH|CONST|VAR) */

func ZEND_FETCH_STATIC_PROP_IS_SPEC_HANDLER(execute_data *ZendExecuteData) int {
	return zend_fetch_static_prop_helper_SPEC(3, execute_data)
}
func zend_use_tmp_in_write_context_helper_SPEC(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	ZendThrowError(nil, "Cannot use temporary expression in write context")
	if (opline.GetOp2Type() & (1<<1 | 1<<2)) != 0 {
		ZvalPtrDtorNogc((*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar())))
	}
	if (opline.GetOp1Type() & (1<<1 | 1<<2)) != 0 {
		ZvalPtrDtorNogc((*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar())))
	}
	(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
	return 0
}
func zend_use_undef_in_read_context_helper_SPEC(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	ZendThrowError(nil, "Cannot use [] for reading")
	if (opline.GetOp2Type() & (1<<1 | 1<<2)) != 0 {
		ZvalPtrDtorNogc((*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar())))
	}
	if (opline.GetOp1Type() & (1<<1 | 1<<2)) != 0 {
		ZvalPtrDtorNogc((*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar())))
	}
	(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
	return 0
}
func ZEND_ASSIGN_STATIC_PROP_SPEC_OP_DATA_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var prop *Zval
	var value *Zval
	var prop_info *ZendPropertyInfo
	if ZendFetchStaticPropertyAddress(&prop, &prop_info, opline.GetExtendedValue(), 1, 0, opline, execute_data) != SUCCESS {
		if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
		}
		return 0
	}
	value = (*Zval)((*byte)(opline+1) + int32((opline + 1).GetOp1()).constant)
	if prop_info.GetType() != 0 {
		value = ZendAssignToTypedProp(prop_info, prop, value, execute_data)
	} else {
		value = ZendAssignToVariable(prop, value, 1<<0, (execute_data.GetFunc().GetFnFlags()&1<<31) != 0)
	}
	if opline.GetResultType() != 0 {
		var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var _z2 *Zval = value
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if (_t & 0xff00) != 0 {
			ZendGcAddref(&_gc.gc)
		}
	}

	/* assign_static_prop has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}
func ZEND_ASSIGN_STATIC_PROP_SPEC_OP_DATA_TMP_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op_data ZendFreeOp
	var prop *Zval
	var value *Zval
	var prop_info *ZendPropertyInfo
	if ZendFetchStaticPropertyAddress(&prop, &prop_info, opline.GetExtendedValue(), 1, 0, opline, execute_data) != SUCCESS {
		ZvalPtrDtorNogc((*Zval)((*byte)(execute_data) + int((opline + 1).GetOp1().GetVar())))
		if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
		}
		return 0
	}
	value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, execute_data)
	if prop_info.GetType() != 0 {
		value = ZendAssignToTypedProp(prop_info, prop, value, execute_data)
		ZvalPtrDtorNogc(free_op_data)
	} else {
		value = ZendAssignToVariable(prop, value, 1<<1, (execute_data.GetFunc().GetFnFlags()&1<<31) != 0)
	}
	if opline.GetResultType() != 0 {
		var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var _z2 *Zval = value
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if (_t & 0xff00) != 0 {
			ZendGcAddref(&_gc.gc)
		}
	}

	/* assign_static_prop has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}
func ZEND_ASSIGN_STATIC_PROP_SPEC_OP_DATA_VAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op_data ZendFreeOp
	var prop *Zval
	var value *Zval
	var prop_info *ZendPropertyInfo
	if ZendFetchStaticPropertyAddress(&prop, &prop_info, opline.GetExtendedValue(), 1, 0, opline, execute_data) != SUCCESS {
		ZvalPtrDtorNogc((*Zval)((*byte)(execute_data) + int((opline + 1).GetOp1().GetVar())))
		if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
		}
		return 0
	}
	value = _getZvalPtrVar((opline + 1).GetOp1().GetVar(), &free_op_data, execute_data)
	if prop_info.GetType() != 0 {
		value = ZendAssignToTypedProp(prop_info, prop, value, execute_data)
		ZvalPtrDtorNogc(free_op_data)
	} else {
		value = ZendAssignToVariable(prop, value, 1<<2, (execute_data.GetFunc().GetFnFlags()&1<<31) != 0)
	}
	if opline.GetResultType() != 0 {
		var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var _z2 *Zval = value
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if (_t & 0xff00) != 0 {
			ZendGcAddref(&_gc.gc)
		}
	}

	/* assign_static_prop has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}
func ZEND_ASSIGN_STATIC_PROP_SPEC_OP_DATA_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var prop *Zval
	var value *Zval
	var prop_info *ZendPropertyInfo
	if ZendFetchStaticPropertyAddress(&prop, &prop_info, opline.GetExtendedValue(), 1, 0, opline, execute_data) != SUCCESS {
		if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
		}
		return 0
	}
	value = _get_zval_ptr_cv_BP_VAR_R((opline + 1).GetOp1().GetVar(), execute_data)
	if prop_info.GetType() != 0 {
		value = ZendAssignToTypedProp(prop_info, prop, value, execute_data)
	} else {
		value = ZendAssignToVariable(prop, value, 1<<3, (execute_data.GetFunc().GetFnFlags()&1<<31) != 0)
	}
	if opline.GetResultType() != 0 {
		var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var _z2 *Zval = value
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if (_t & 0xff00) != 0 {
			ZendGcAddref(&_gc.gc)
		}
	}

	/* assign_static_prop has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}
func ZEND_ASSIGN_STATIC_PROP_REF_SPEC_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op_data ZendFreeOp
	var prop *Zval
	var value_ptr *Zval
	var prop_info *ZendPropertyInfo
	if ZendFetchStaticPropertyAddress(&prop, &prop_info, opline.GetExtendedValue() & ^(1<<0), 1, 0, opline, execute_data) != SUCCESS {
		if ((opline + 1).GetOp1Type() & (1<<1 | 1<<2)) != 0 {
			ZvalPtrDtorNogc((*Zval)((*byte)(execute_data) + int((opline + 1).GetOp1().GetVar())))
		}
		if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
		}
		return 0
	}
	value_ptr = _getZvalPtrPtr((opline + 1).GetOp1Type(), (opline + 1).GetOp1(), &free_op_data, 1, execute_data)
	if (opline+1).GetOp1Type() == 1<<2 && value_ptr.GetType() == 15 {
		prop = &EG.uninitialized_zval
	} else if (opline+1).GetOp1Type() == 1<<2 && (opline.GetExtendedValue()&1<<0) != 0 && value_ptr.GetType() != 10 {
		if ZendWrongAssignToVariableReference(prop, value_ptr, opline, execute_data) == nil {
			prop = &EG.uninitialized_zval
		}
	} else if prop_info.GetType() != 0 {
		prop = ZendAssignToTypedPropertyReference(prop_info, prop, value_ptr, execute_data)
	} else {
		ZendAssignToVariableReference(prop, value_ptr)
	}
	if opline.GetResultType() != 0 {
		var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var _z2 *Zval = prop
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if (_t & 0xff00) != 0 {
			ZendGcAddref(&_gc.gc)
		}
	}
	if free_op_data != nil {
		ZvalPtrDtorNogc(free_op_data)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}
func zend_leave_helper_SPEC(execute_data *ZendExecuteData) int {
	var old_execute_data *ZendExecuteData
	var call_info uint32 = execute_data.GetThis().GetTypeInfo()
	if (call_info & (1<<16 | 1<<17 | 1<<20 | 1<<19 | 1<<18)) == 0 {
		EG.SetCurrentExecuteData(execute_data.GetPrevExecuteData())
		IFreeCompiledVariables(execute_data)
		if (call_info & 1 << 21) != 0 {
			ZendObjectRelease(execute_data.GetThis().GetValue().GetObj())
		} else if (call_info & 1 << 22) != 0 {
			ZendObjectRelease((*ZendObject)((*byte)(execute_data.GetFunc() - g.SizeOf("zend_object"))))
		}
		EG.SetVmStackTop((*Zval)(execute_data))
		execute_data = execute_data.GetPrevExecuteData()
		if EG.GetException() != nil {
			ZendRethrowException(execute_data)
			return 2
		}
		execute_data.GetOpline()++
		return 2
	} else if (call_info & (1<<16 | 1<<17)) == 0 {
		EG.SetCurrentExecuteData(execute_data.GetPrevExecuteData())
		IFreeCompiledVariables(execute_data)
		if (call_info & 1 << 20) != 0 {
			ZendCleanAndCacheSymbolTable(execute_data.GetSymbolTable())
		}

		/* Free extra args before releasing the closure,
		 * as that may free the op_array. */

		ZendVmStackFreeExtraArgsEx(call_info, execute_data)
		if (call_info & 1 << 21) != 0 {
			ZendObjectRelease(execute_data.GetThis().GetValue().GetObj())
		} else if (call_info & 1 << 22) != 0 {
			ZendObjectRelease((*ZendObject)((*byte)(execute_data.GetFunc() - g.SizeOf("zend_object"))))
		}
		old_execute_data = execute_data
		execute_data = execute_data.GetPrevExecuteData()
		ZendVmStackFreeCallFrameEx(call_info, old_execute_data)
		if EG.GetException() != nil {
			ZendRethrowException(execute_data)
			return 2
		}
		execute_data.GetOpline()++
		return 2
	} else if (call_info & 1 << 17) == 0 {
		ZendDetachSymbolTable(execute_data)
		DestroyOpArray(&(execute_data.GetFunc()).op_array)
		_efree(execute_data.GetFunc())
		old_execute_data = execute_data
		EG.SetCurrentExecuteData(execute_data.GetPrevExecuteData())
		execute_data = EG.GetCurrentExecuteData()
		ZendVmStackFreeCallFrameEx(call_info, old_execute_data)
		ZendAttachSymbolTable(execute_data)
		if EG.GetException() != nil {
			ZendRethrowException(execute_data)
			return 2
		}
		execute_data.GetOpline()++
		return 2
	} else {
		if (call_info & 1 << 16) == 0 {
			EG.SetCurrentExecuteData(execute_data.GetPrevExecuteData())
			IFreeCompiledVariables(execute_data)
			if (call_info & (1<<20 | 1<<19)) != 0 {
				if (call_info & 1 << 20) != 0 {
					ZendCleanAndCacheSymbolTable(execute_data.GetSymbolTable())
				}
				ZendVmStackFreeExtraArgsEx(call_info, execute_data)
			}
			if (call_info & 1 << 22) != 0 {
				ZendObjectRelease((*ZendObject)((*byte)(execute_data.GetFunc() - g.SizeOf("zend_object"))))
			}
			return -1
		} else {
			var symbol_table *ZendArray = execute_data.GetSymbolTable()
			ZendDetachSymbolTable(execute_data)
			old_execute_data = execute_data.GetPrevExecuteData()
			for old_execute_data != nil {
				if old_execute_data.GetFunc() != nil && (old_execute_data.GetThis().GetTypeInfo()&1<<20) != 0 {
					if old_execute_data.GetSymbolTable() == symbol_table {
						ZendAttachSymbolTable(old_execute_data)
					}
					break
				}
				old_execute_data = old_execute_data.GetPrevExecuteData()
			}
			EG.SetCurrentExecuteData(execute_data.GetPrevExecuteData())
			return -1
		}
	}
}
func ZEND_JMP_SPEC_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()

	execute_data.SetOpline((*ZendOp)((*byte)(opline) + int(opline.GetOp1().GetJmpOffset())))
	if EG.GetVmInterrupt() != 0 {
		return zend_interrupt_helper_SPEC(execute_data)
	}
	return 0
}
func ZEND_DO_ICALL_SPEC_RETVAL_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var call *ZendExecuteData = execute_data.GetCall()
	var fbc *ZendFunction = call.GetFunc()
	var ret *Zval
	var retval Zval
	execute_data.SetCall(call.GetPrevExecuteData())
	call.SetPrevExecuteData(execute_data)
	EG.SetCurrentExecuteData(call)
	ret = &retval
	ret.SetTypeInfo(1)
	fbc.GetInternalFunction().GetHandler()(call, ret)
	EG.SetCurrentExecuteData(execute_data)
	ZendVmStackFreeArgs(call)
	ZendVmStackFreeCallFrame(call)
	IZvalPtrDtor(ret)
	if EG.GetException() != nil {
		ZendRethrowException(execute_data)
		return 0
	}
	execute_data.SetOpline(opline + 1)
	if EG.GetVmInterrupt() != 0 {
		return zend_interrupt_helper_SPEC(execute_data)
	}
	return 0
}
func ZEND_DO_ICALL_SPEC_RETVAL_USED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var call *ZendExecuteData = execute_data.GetCall()
	var fbc *ZendFunction = call.GetFunc()
	var ret *Zval
	var retval Zval
	execute_data.SetCall(call.GetPrevExecuteData())
	call.SetPrevExecuteData(execute_data)
	EG.SetCurrentExecuteData(call)
	ret = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	ret.SetTypeInfo(1)
	fbc.GetInternalFunction().GetHandler()(call, ret)
	EG.SetCurrentExecuteData(execute_data)
	ZendVmStackFreeArgs(call)
	ZendVmStackFreeCallFrame(call)

	if EG.GetException() != nil {
		ZendRethrowException(execute_data)
		return 0
	}
	execute_data.SetOpline(opline + 1)
	if EG.GetVmInterrupt() != 0 {
		return zend_interrupt_helper_SPEC(execute_data)
	}
	return 0
}
func ZEND_DO_UCALL_SPEC_RETVAL_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var call *ZendExecuteData = execute_data.GetCall()
	var fbc *ZendFunction = call.GetFunc()
	var ret *Zval
	execute_data.SetCall(call.GetPrevExecuteData())
	ret = nil

	call.SetPrevExecuteData(execute_data)
	execute_data = call
	IInitFuncExecuteData(&fbc.op_array, ret, 0, execute_data)
	return 1
}
func ZEND_DO_UCALL_SPEC_RETVAL_USED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var call *ZendExecuteData = execute_data.GetCall()
	var fbc *ZendFunction = call.GetFunc()
	var ret *Zval
	execute_data.SetCall(call.GetPrevExecuteData())
	ret = nil
	ret = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	call.SetPrevExecuteData(execute_data)
	execute_data = call
	IInitFuncExecuteData(&fbc.op_array, ret, 0, execute_data)
	return 1
}
func ZEND_DO_FCALL_BY_NAME_SPEC_RETVAL_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var call *ZendExecuteData = execute_data.GetCall()
	var fbc *ZendFunction = call.GetFunc()
	var ret *Zval
	execute_data.SetCall(call.GetPrevExecuteData())
	if fbc.GetType() == 2 {
		ret = nil

		call.SetPrevExecuteData(execute_data)
		execute_data = call
		IInitFuncExecuteData(&fbc.op_array, ret, 0, execute_data)
		return 1
	} else {
		var retval Zval
		assert(fbc.GetType() == 1)
		if (fbc.GetFnFlags() & 1 << 11) != 0 {
			ZendDeprecatedFunction(fbc)
			if EG.GetException() != nil {
				if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
					(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
				}
				ret = &retval
				ret.SetTypeInfo(0)
				goto fcall_by_name_end
			}
		}
		call.SetPrevExecuteData(execute_data)
		EG.SetCurrentExecuteData(call)
		if (fbc.GetFnFlags()&1<<8) != 0 && ZendVerifyInternalArgTypes(fbc, call) == 0 {
			if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
			}
			ret = &retval
			ret.SetTypeInfo(0)
			goto fcall_by_name_end
		}
		ret = &retval
		ret.SetTypeInfo(1)
		fbc.GetInternalFunction().GetHandler()(call, ret)
		EG.SetCurrentExecuteData(execute_data)
	fcall_by_name_end:
		ZendVmStackFreeArgs(call)
		ZendVmStackFreeCallFrame(call)
		IZvalPtrDtor(ret)
	}
	if EG.GetException() != nil {
		ZendRethrowException(execute_data)
		return 0
	}
	execute_data.SetOpline(opline + 1)
	if EG.GetVmInterrupt() != 0 {
		return zend_interrupt_helper_SPEC(execute_data)
	}
	return 0
}
func ZEND_DO_FCALL_BY_NAME_SPEC_RETVAL_USED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var call *ZendExecuteData = execute_data.GetCall()
	var fbc *ZendFunction = call.GetFunc()
	var ret *Zval
	execute_data.SetCall(call.GetPrevExecuteData())
	if fbc.GetType() == 2 {
		ret = nil
		ret = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		call.SetPrevExecuteData(execute_data)
		execute_data = call
		IInitFuncExecuteData(&fbc.op_array, ret, 0, execute_data)
		return 1
	} else {
		var retval Zval
		assert(fbc.GetType() == 1)
		if (fbc.GetFnFlags() & 1 << 11) != 0 {
			ZendDeprecatedFunction(fbc)
			if EG.GetException() != nil {
				if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
					(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
				}

				goto fcall_by_name_end
			}
		}
		call.SetPrevExecuteData(execute_data)
		EG.SetCurrentExecuteData(call)
		if (fbc.GetFnFlags()&1<<8) != 0 && ZendVerifyInternalArgTypes(fbc, call) == 0 {
			if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
			}

			goto fcall_by_name_end
		}
		ret = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		ret.SetTypeInfo(1)
		fbc.GetInternalFunction().GetHandler()(call, ret)
		EG.SetCurrentExecuteData(execute_data)
	fcall_by_name_end:
		ZendVmStackFreeArgs(call)
		ZendVmStackFreeCallFrame(call)

	}
	if EG.GetException() != nil {
		ZendRethrowException(execute_data)
		return 0
	}
	execute_data.SetOpline(opline + 1)
	if EG.GetVmInterrupt() != 0 {
		return zend_interrupt_helper_SPEC(execute_data)
	}
	return 0
}
func ZEND_DO_FCALL_SPEC_RETVAL_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var call *ZendExecuteData = execute_data.GetCall()
	var fbc *ZendFunction = call.GetFunc()
	var ret *Zval
	var retval Zval
	execute_data.SetCall(call.GetPrevExecuteData())
	if (fbc.GetFnFlags() & (1<<6 | 1<<11)) != 0 {
		if (fbc.GetFnFlags() & 1 << 6) != 0 {
			ZendAbstractMethod(fbc)
		fcall_except:
			if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
			}
			ret = &retval
			ret.SetTypeInfo(0)
			goto fcall_end
		} else {
			ZendDeprecatedFunction(fbc)
			if EG.GetException() != nil {
				goto fcall_except
			}
		}
	}
	if fbc.GetType() == 2 {
		ret = nil

		call.SetPrevExecuteData(execute_data)
		execute_data = call
		IInitFuncExecuteData(&fbc.op_array, ret, 1, execute_data)
		if ZendExecuteEx == ExecuteEx {
			return 1
		} else {
			execute_data = execute_data.GetPrevExecuteData()
			call.GetThis().SetTypeInfo(call.GetThis().GetTypeInfo() | 1<<17)
			ZendExecuteEx(call)
		}
	} else if fbc.GetType() < 2 {
		call.SetPrevExecuteData(execute_data)
		EG.SetCurrentExecuteData(call)
		if (fbc.GetFnFlags()&1<<8) != 0 && ZendVerifyInternalArgTypes(fbc, call) == 0 {
			goto fcall_except
		}
		ret = &retval
		ret.SetTypeInfo(1)
		if ZendExecuteInternal == nil {

			/* saves one function call if zend_execute_internal is not used */

			fbc.GetInternalFunction().GetHandler()(call, ret)

			/* saves one function call if zend_execute_internal is not used */

		} else {
			ZendExecuteInternal(call, ret)
		}
		EG.SetCurrentExecuteData(execute_data)
	fcall_end:
		ZendVmStackFreeArgs(call)
		IZvalPtrDtor(ret)
	} else {
		ret = &retval
		call.SetPrevExecuteData(execute_data)
		if ZendDoFcallOverloaded(call, ret) == 0 {
			if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
			}
			return 0
		}
		ZvalPtrDtor(ret)
	}
	if (call.GetThis().GetTypeInfo() & 1 << 21) != 0 {
		ZendObjectRelease(call.GetThis().GetValue().GetObj())
	}
	ZendVmStackFreeCallFrame(call)
	if EG.GetException() != nil {
		ZendRethrowException(execute_data)
		return 0
	}
	execute_data.SetOpline(opline + 1)
	if EG.GetVmInterrupt() != 0 {
		return zend_interrupt_helper_SPEC(execute_data)
	}
	return 0
}
func ZEND_DO_FCALL_SPEC_RETVAL_USED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var call *ZendExecuteData = execute_data.GetCall()
	var fbc *ZendFunction = call.GetFunc()
	var ret *Zval
	var retval Zval
	execute_data.SetCall(call.GetPrevExecuteData())
	if (fbc.GetFnFlags() & (1<<6 | 1<<11)) != 0 {
		if (fbc.GetFnFlags() & 1 << 6) != 0 {
			ZendAbstractMethod(fbc)
		fcall_except:
			if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
			}

			goto fcall_end
		} else {
			ZendDeprecatedFunction(fbc)
			if EG.GetException() != nil {
				goto fcall_except
			}
		}
	}
	if fbc.GetType() == 2 {
		ret = nil
		ret = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		call.SetPrevExecuteData(execute_data)
		execute_data = call
		IInitFuncExecuteData(&fbc.op_array, ret, 1, execute_data)
		if ZendExecuteEx == ExecuteEx {
			return 1
		} else {
			execute_data = execute_data.GetPrevExecuteData()
			call.GetThis().SetTypeInfo(call.GetThis().GetTypeInfo() | 1<<17)
			ZendExecuteEx(call)
		}
	} else if fbc.GetType() < 2 {
		call.SetPrevExecuteData(execute_data)
		EG.SetCurrentExecuteData(call)
		if (fbc.GetFnFlags()&1<<8) != 0 && ZendVerifyInternalArgTypes(fbc, call) == 0 {
			goto fcall_except
		}
		ret = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		ret.SetTypeInfo(1)
		if ZendExecuteInternal == nil {

			/* saves one function call if zend_execute_internal is not used */

			fbc.GetInternalFunction().GetHandler()(call, ret)

			/* saves one function call if zend_execute_internal is not used */

		} else {
			ZendExecuteInternal(call, ret)
		}
		EG.SetCurrentExecuteData(execute_data)
	fcall_end:
		ZendVmStackFreeArgs(call)

	} else {
		ret = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		call.SetPrevExecuteData(execute_data)
		if ZendDoFcallOverloaded(call, ret) == 0 {
			if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
			}
			return 0
		}

	}
	if (call.GetThis().GetTypeInfo() & 1 << 21) != 0 {
		ZendObjectRelease(call.GetThis().GetValue().GetObj())
	}
	ZendVmStackFreeCallFrame(call)
	if EG.GetException() != nil {
		ZendRethrowException(execute_data)
		return 0
	}
	execute_data.SetOpline(opline + 1)
	if EG.GetVmInterrupt() != 0 {
		return zend_interrupt_helper_SPEC(execute_data)
	}
	return 0
}
func ZEND_GENERATOR_CREATE_SPEC_HANDLER(execute_data *ZendExecuteData) int {
	var return_value *Zval = execute_data.GetReturnValue()
	if return_value != nil {
		var opline *ZendOp = execute_data.GetOpline()
		var generator *ZendGenerator
		var gen_execute_data *ZendExecuteData
		var num_args uint32
		var used_stack uint32
		var call_info uint32
		ObjectInitEx(return_value, ZendCeGenerator)

		/*
		 * Normally the execute_data is allocated on the VM stack (because it does
		 * not actually do any allocation and thus is faster). For generators
		 * though this behavior would be suboptimal, because the (rather large)
		 * structure would have to be copied back and forth every time execution is
		 * suspended or resumed. That's why for generators the execution context
		 * is allocated on heap.
		 */

		num_args = execute_data.GetThis().GetNumArgs()
		if num_args <= execute_data.GetFunc().GetOpArray().GetNumArgs() {
			used_stack = (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + execute_data.GetFunc().GetOpArray().GetLastVar() + execute_data.GetFunc().GetOpArray().GetT()) * g.SizeOf("zval")
			gen_execute_data = (*ZendExecuteData)(_emalloc(used_stack))
			used_stack = (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + execute_data.GetFunc().GetOpArray().GetLastVar()) * g.SizeOf("zval")
		} else {
			used_stack = (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + num_args + execute_data.GetFunc().GetOpArray().GetLastVar() + execute_data.GetFunc().GetOpArray().GetT() - execute_data.GetFunc().GetOpArray().GetNumArgs()) * g.SizeOf("zval")
			gen_execute_data = (*ZendExecuteData)(_emalloc(used_stack))
		}
		memcpy(gen_execute_data, execute_data, used_stack)

		/* Save execution context in generator object. */

		generator = (*ZendGenerator)(execute_data.GetReturnValue().GetValue().GetObj())
		generator.SetExecuteData(gen_execute_data)
		generator.SetFrozenCallStack(nil)
		generator.GetExecuteFake().SetOpline(nil)
		generator.GetExecuteFake().SetFunc(nil)
		generator.GetExecuteFake().SetPrevExecuteData(nil)
		var __z *Zval = &generator.execute_fake.GetThis()
		__z.GetValue().SetObj((*ZendObject)(generator))
		__z.SetTypeInfo(8 | 1<<0<<8 | 1<<1<<8)
		gen_execute_data.SetOpline(opline + 1)

		/* EX(return_value) keeps pointer to zend_object (not a real zval) */

		gen_execute_data.SetReturnValue((*Zval)(generator))
		call_info = execute_data.GetThis().GetTypeInfo()
		if (call_info&0xff) == 8 && ((call_info&(1<<22|1<<21)) == 0 || ZendExecuteEx != ExecuteEx) {
			call_info |= 1 << 21
			ZvalAddrefP(&(gen_execute_data.GetThis()))
		}
		call_info |= 1<<17 | 0<<16 | 1<<18 | 1<<24
		gen_execute_data.GetThis().SetTypeInfo(call_info)
		gen_execute_data.SetPrevExecuteData(nil)
		call_info = execute_data.GetThis().GetTypeInfo()
		EG.SetCurrentExecuteData(execute_data.GetPrevExecuteData())
		if (call_info & (1<<17 | 1<<18)) == 0 {
			EG.SetVmStackTop((*Zval)(execute_data))
			execute_data = execute_data.GetPrevExecuteData()
			execute_data.GetOpline()++
			return 2
		} else if (call_info & 1 << 17) == 0 {
			var old_execute_data *ZendExecuteData = execute_data
			execute_data = execute_data.GetPrevExecuteData()
			ZendVmStackFreeCallFrameEx(call_info, old_execute_data)
			execute_data.GetOpline()++
			return 2
		} else {
			return -1
		}
	} else {
		return zend_leave_helper_SPEC(execute_data)
	}
}
func zend_cannot_pass_by_ref_helper_SPEC(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var arg *Zval
	var arg_num uint32 = opline.GetOp2().GetNum()
	ZendThrowError(nil, "Cannot pass parameter %d by reference", arg_num)
	if (opline.GetOp1Type() & (1<<1 | 1<<2)) != 0 {
		ZvalPtrDtorNogc((*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar())))
	}
	arg = (*Zval)((*byte)(execute_data.GetCall()) + int(opline.GetResult().GetVar()))
	arg.SetTypeInfo(0)
	return 0
}
func ZEND_SEND_UNPACK_SPEC_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var args *Zval
	var arg_num int
	args = _getZvalPtrUndef(opline.GetOp1Type(), opline.GetOp1(), &free_op1, 0, execute_data, opline)
	arg_num = execute_data.GetCall().GetThis().GetNumArgs() + 1
send_again:
	if args.GetType() == 7 {
		var ht *HashTable = args.GetValue().GetArr()
		var arg *Zval
		var top *Zval
		var name *ZendString
		ZendVmStackExtendCallFrame(&(execute_data.GetCall()), arg_num-1, ht.GetNNumOfElements())
		if (opline.GetOp1Type()&(1<<2|1<<3)) != 0 && ZvalRefcountP(args) > 1 {
			var i uint32
			var separate int = 0

			/* check if any of arguments are going to be passed by reference */

			for i = 0; i < ht.GetNNumOfElements(); i++ {
				if ZendCheckArgSendType(execute_data.GetCall().GetFunc(), arg_num+i, 1|2) != 0 {
					separate = 1
					break
				}
			}
			if separate != 0 {
				var _zv *Zval = args
				var _arr *ZendArray = _zv.GetValue().GetArr()
				if ZendGcRefcount(&_arr.gc) > 1 {
					if _zv.GetTypeFlags() != 0 {
						ZendGcDelref(&_arr.gc)
					}
					var __arr *ZendArray = ZendArrayDup(_arr)
					var __z *Zval = _zv
					__z.GetValue().SetArr(__arr)
					__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
				}
				ht = args.GetValue().GetArr()
			}
		}
		for {
			var __ht *HashTable = ht
			var _p *Bucket = __ht.GetArData()
			var _end *Bucket = _p + __ht.GetNNumUsed()
			for ; _p != _end; _p++ {
				var _z *Zval = &_p.val

				if _z.GetType() == 0 {
					continue
				}
				name = _p.GetKey()
				arg = _z
				if name != nil {
					ZendThrowError(nil, "Cannot unpack array with string keys")
					if free_op1 != nil {
						ZvalPtrDtorNogc(free_op1)
					}
					return 0
				}
				top = (*Zval)(execute_data.GetCall()) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(arg_num)-1))
				if ZendCheckArgSendType(execute_data.GetCall().GetFunc(), arg_num, 1|2) != 0 {
					if arg.GetType() == 10 {
						ZvalAddrefP(arg)
						var __z *Zval = top
						__z.GetValue().SetRef(arg.GetValue().GetRef())
						__z.SetTypeInfo(10 | 1<<0<<8)
					} else if (opline.GetOp1Type() & (1<<2 | 1<<3)) != 0 {

						/* array is already separated above */

						var _z *Zval = arg
						var _ref *ZendReference = (*ZendReference)(_emalloc(g.SizeOf("zend_reference")))
						ZendGcSetRefcount(&_ref.gc, 2)
						_ref.GetGc().SetTypeInfo(10)
						var _z1 *Zval = &_ref.val
						var _z2 *Zval = _z
						var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
						var _t uint32 = _z2.GetTypeInfo()
						_z1.GetValue().SetCounted(_gc)
						_z1.SetTypeInfo(_t)
						_ref.GetSources().SetPtr(nil)
						_z.GetValue().SetRef(_ref)
						_z.SetTypeInfo(10 | 1<<0<<8)
						var __z *Zval = top
						__z.GetValue().SetRef(arg.GetValue().GetRef())
						__z.SetTypeInfo(10 | 1<<0<<8)
					} else {
						if arg.GetTypeFlags() != 0 {
							ZvalAddrefP(arg)
						}
						var _ref *ZendReference = (*ZendReference)(_emalloc(g.SizeOf("zend_reference")))
						ZendGcSetRefcount(&_ref.gc, 1)
						_ref.GetGc().SetTypeInfo(10)
						var _z1 *Zval = &_ref.val
						var _z2 *Zval = arg
						var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
						var _t uint32 = _z2.GetTypeInfo()
						_z1.GetValue().SetCounted(_gc)
						_z1.SetTypeInfo(_t)
						_ref.GetSources().SetPtr(nil)
						top.GetValue().SetRef(_ref)
						top.SetTypeInfo(10 | 1<<0<<8)
					}
				} else {
					var _z3 *Zval = arg
					if (_z3.GetTypeInfo() & 0xff00) != 0 {
						if (_z3.GetTypeInfo() & 0xff) == 10 {
							_z3 = &(*_z3).value.GetRef().GetVal()
							if (_z3.GetTypeInfo() & 0xff00) != 0 {
								ZvalAddrefP(_z3)
							}
						} else {
							ZvalAddrefP(_z3)
						}
					}
					var _z1 *Zval = top
					var _z2 *Zval = _z3
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
				}
				execute_data.GetCall().GetThis().GetNumArgs()++
				arg_num++
			}
			break
		}
	} else if args.GetType() == 8 {
		var ce *ZendClassEntry = args.GetValue().GetObj().GetCe()
		var iter *ZendObjectIterator
		if ce == nil || ce.GetGetIterator() == nil {
			ZendError(1<<1, "Only arrays and Traversables can be unpacked")
		} else {
			iter = ce.GetGetIterator()(ce, args, 0)
			if iter == nil {
				if free_op1 != nil {
					ZvalPtrDtorNogc(free_op1)
				}
				if EG.GetException() == nil {
					ZendThrowExceptionEx(nil, 0, "Object of type %s did not create an Iterator", ce.GetName().GetVal())
				}
				return 0
			}
			if iter.GetFuncs().GetRewind() != nil {
				iter.GetFuncs().GetRewind()(iter)
			}
			for ; iter.GetFuncs().GetValid()(iter) == SUCCESS; arg_num++ {
				var arg *Zval
				var top *Zval
				if EG.GetException() != nil {
					break
				}
				arg = iter.GetFuncs().GetGetCurrentData()(iter)
				if EG.GetException() != nil {
					break
				}
				if iter.GetFuncs().GetGetCurrentKey() != nil {
					var key Zval
					iter.GetFuncs().GetGetCurrentKey()(iter, &key)
					if EG.GetException() != nil {
						break
					}
					if key.GetType() != 4 {
						ZendThrowError(nil, g.Cond(key.GetType() == 6, "Cannot unpack Traversable with string keys", "Cannot unpack Traversable with non-integer keys"))
						ZvalPtrDtor(&key)
						break
					}
				}
				if ZendCheckArgSendType(execute_data.GetCall().GetFunc(), arg_num, 1) != 0 {
					ZendError(1<<1, "Cannot pass by-reference argument %d of %s%s%s()"+" by unpacking a Traversable, passing by-value instead", arg_num, g.CondF1(execute_data.GetCall().GetFunc().GetScope() != nil, func() []byte { return execute_data.GetCall().GetFunc().GetScope().GetName().GetVal() }, ""), g.Cond(execute_data.GetCall().GetFunc().GetScope() != nil, "::", ""), execute_data.GetCall().GetFunc().GetFunctionName().GetVal())
				}
				if arg.GetType() == 10 {
					arg = &(*arg).value.GetRef().GetVal()
				}
				if arg.GetTypeFlags() != 0 {
					ZvalAddrefP(arg)
				}
				ZendVmStackExtendCallFrame(&(execute_data.GetCall()), arg_num-1, 1)
				top = (*Zval)(execute_data.GetCall()) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(arg_num)-1))
				var _z1 *Zval = top
				var _z2 *Zval = arg
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				execute_data.GetCall().GetThis().GetNumArgs()++
				iter.GetFuncs().GetMoveForward()(iter)
			}
			ZendIteratorDtor(iter)
		}
	} else if args.GetType() == 10 {
		args = &(*args).value.GetRef().GetVal()
		goto send_again
	} else {
		if opline.GetOp1Type() == 1<<3 && args.GetType() == 0 {
			_zvalUndefinedOp1(execute_data)
		}
		ZendError(1<<1, "Only arrays and Traversables can be unpacked")
	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_SEND_ARRAY_SPEC_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var args *Zval
	args = _getZvalPtr(opline.GetOp1Type(), opline.GetOp1(), &free_op1, 0, execute_data, opline)
	if args.GetType() != 7 {
		if (opline.GetOp1Type()&(1<<2|1<<3)) != 0 && args.GetType() == 10 {
			args = &(*args).value.GetRef().GetVal()
			if args.GetType() == 7 {
				goto send_array
			}
		}
		ZendInternalTypeError((execute_data.GetFunc().GetFnFlags()&1<<31) != 0, "call_user_func_array() expects parameter 2 to be array, %s given", ZendGetTypeByConst(args.GetType()))
		if (execute_data.GetCall().GetThis().GetTypeInfo() & 1 << 22) != 0 {
			ZendObjectRelease((*ZendObject)((*byte)(execute_data.GetCall().GetFunc() - g.SizeOf("zend_object"))))
		} else if (execute_data.GetCall().GetThis().GetTypeInfo() & 1 << 21) != 0 {
			ZendObjectRelease(execute_data.GetCall().GetThis().GetValue().GetObj())
		}
		execute_data.GetCall().SetFunc((*ZendFunction)(&ZendPassFunction))
		execute_data.GetCall().GetThis().GetValue().SetObj(nil)
		execute_data.GetCall().GetThis().SetTypeInfo(execute_data.GetCall().GetThis().GetTypeInfo() &^ (1<<21 | (8 | 1<<0<<8 | 1<<1<<8)))
		if (opline.GetOp2Type() & (1<<1 | 1<<2)) != 0 {
			ZvalPtrDtorNogc((*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar())))
		}
	} else {
		var arg_num uint32
		var ht *HashTable
		var arg *Zval
		var param *Zval
	send_array:
		ht = args.GetValue().GetArr()
		if opline.GetOp2Type() != 0 {
			var free_op2 ZendFreeOp
			var op2 *Zval = _getZvalPtr(opline.GetOp2Type(), opline.GetOp2(), &free_op2, 0, execute_data, opline)
			var skip uint32 = opline.GetExtendedValue()
			var count uint32 = ht.GetNNumOfElements()
			var len_ ZendLong = ZvalGetLong(op2)
			if len_ < 0 {
				len_ += zend_long(count - skip)
			}
			if skip < count && len_ > 0 {
				if len_ > zend_long(count-skip) {
					len_ = zend_long(count - skip)
				}
				ZendVmStackExtendCallFrame(&(execute_data.GetCall()), 0, len_)
				arg_num = 1
				param = (*Zval)(execute_data.GetCall()) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(1)-1))
				for {
					var __ht *HashTable = ht
					var _p *Bucket = __ht.GetArData()
					var _end *Bucket = _p + __ht.GetNNumUsed()
					for ; _p != _end; _p++ {
						var _z *Zval = &_p.val

						if _z.GetType() == 0 {
							continue
						}
						arg = _z
						var must_wrap ZendBool = 0
						if skip > 0 {
							skip--
							continue
						} else if zend_long(arg_num-1) >= len_ {
							break
						} else if ZendCheckArgSendType(execute_data.GetCall().GetFunc(), arg_num, 1|2) != 0 {
							if arg.GetType() != 10 {
								if ZendCheckArgSendType(execute_data.GetCall().GetFunc(), arg_num, 2) == 0 {

									/* By-value send is not allowed -- emit a warning,
									 * but still perform the call. */

									ZendParamMustBeRef(execute_data.GetCall().GetFunc(), arg_num)
									must_wrap = 1
								}
							}
						} else {
							if arg.GetType() == 10 && (execute_data.GetCall().GetFunc().GetFnFlags()&1<<18) == 0 {

								/* don't separate references for __call */

								arg = &(*arg).value.GetRef().GetVal()

								/* don't separate references for __call */

							}
						}
						if must_wrap == 0 {
							var _z1 *Zval = param
							var _z2 *Zval = arg
							var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
							var _t uint32 = _z2.GetTypeInfo()
							_z1.GetValue().SetCounted(_gc)
							_z1.SetTypeInfo(_t)
							if (_t & 0xff00) != 0 {
								ZendGcAddref(&_gc.gc)
							}
						} else {
							if arg.GetTypeFlags() != 0 {
								ZvalAddrefP(arg)
							}
							var _ref *ZendReference = (*ZendReference)(_emalloc(g.SizeOf("zend_reference")))
							ZendGcSetRefcount(&_ref.gc, 1)
							_ref.GetGc().SetTypeInfo(10)
							var _z1 *Zval = &_ref.val
							var _z2 *Zval = arg
							var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
							var _t uint32 = _z2.GetTypeInfo()
							_z1.GetValue().SetCounted(_gc)
							_z1.SetTypeInfo(_t)
							_ref.GetSources().SetPtr(nil)
							param.GetValue().SetRef(_ref)
							param.SetTypeInfo(10 | 1<<0<<8)
						}
						execute_data.GetCall().GetThis().GetNumArgs()++
						arg_num++
						param++
					}
					break
				}
			}
			if free_op2 != nil {
				ZvalPtrDtorNogc(free_op2)
			}
		} else {
			ZendVmStackExtendCallFrame(&(execute_data.GetCall()), 0, ht.GetNNumOfElements())
			arg_num = 1
			param = (*Zval)(execute_data.GetCall()) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(1)-1))
			for {
				var __ht *HashTable = ht
				var _p *Bucket = __ht.GetArData()
				var _end *Bucket = _p + __ht.GetNNumUsed()
				for ; _p != _end; _p++ {
					var _z *Zval = &_p.val

					if _z.GetType() == 0 {
						continue
					}
					arg = _z
					var must_wrap ZendBool = 0
					if ZendCheckArgSendType(execute_data.GetCall().GetFunc(), arg_num, 1|2) != 0 {
						if arg.GetType() != 10 {
							if ZendCheckArgSendType(execute_data.GetCall().GetFunc(), arg_num, 2) == 0 {

								/* By-value send is not allowed -- emit a warning,
								 * but still perform the call. */

								ZendParamMustBeRef(execute_data.GetCall().GetFunc(), arg_num)
								must_wrap = 1
							}
						}
					} else {
						if arg.GetType() == 10 && (execute_data.GetCall().GetFunc().GetFnFlags()&1<<18) == 0 {

							/* don't separate references for __call */

							arg = &(*arg).value.GetRef().GetVal()

							/* don't separate references for __call */

						}
					}
					if must_wrap == 0 {
						var _z1 *Zval = param
						var _z2 *Zval = arg
						var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
						var _t uint32 = _z2.GetTypeInfo()
						_z1.GetValue().SetCounted(_gc)
						_z1.SetTypeInfo(_t)
						if (_t & 0xff00) != 0 {
							ZendGcAddref(&_gc.gc)
						}
					} else {
						if arg.GetTypeFlags() != 0 {
							ZvalAddrefP(arg)
						}
						var _ref *ZendReference = (*ZendReference)(_emalloc(g.SizeOf("zend_reference")))
						ZendGcSetRefcount(&_ref.gc, 1)
						_ref.GetGc().SetTypeInfo(10)
						var _z1 *Zval = &_ref.val
						var _z2 *Zval = arg
						var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
						var _t uint32 = _z2.GetTypeInfo()
						_z1.GetValue().SetCounted(_gc)
						_z1.SetTypeInfo(_t)
						_ref.GetSources().SetPtr(nil)
						param.GetValue().SetRef(_ref)
						param.SetTypeInfo(10 | 1<<0<<8)
					}
					execute_data.GetCall().GetThis().GetNumArgs()++
					arg_num++
					param++
				}
				break
			}
		}
	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func zend_case_helper_SPEC(op_1 *Zval, op_2 *Zval, execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	if op_1.GetTypeInfo() == 0 {
		op_1 = _zvalUndefinedOp1(execute_data)
	}
	if op_2.GetTypeInfo() == 0 {
		op_2 = _zvalUndefinedOp2(execute_data)
	}
	CompareFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op_1, op_2)
	if (opline.GetOp2Type() & (1<<1 | 1<<2)) != 0 {
		ZvalPtrDtorNogc(op_2)
	}
	if EG.GetException() != nil {
		return 0
	}
	if (*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())).GetValue().GetLval() == 0 {
		if (opline + 1).GetOpcode() == 44 {
			execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
			if EG.GetVmInterrupt() != 0 {
				return zend_interrupt_helper_SPEC(execute_data)
			}
			return 0
		} else if (opline + 1).GetOpcode() == 43 {
			execute_data.SetOpline(opline + 2)
			return 0
		}
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	} else {
		if (opline + 1).GetOpcode() == 44 {
			execute_data.SetOpline(opline + 2)
			return 0
		} else if (opline + 1).GetOpcode() == 43 {
			execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
			if EG.GetVmInterrupt() != 0 {
				return zend_interrupt_helper_SPEC(execute_data)
			}
			return 0
		}
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
}
func ZEND_ADD_ARRAY_UNPACK_SPEC_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *Zval
	op1 = _getZvalPtr(opline.GetOp1Type(), opline.GetOp1(), &free_op1, 0, execute_data, opline)
add_unpack_again:
	if op1.GetType() == 7 {
		var ht *HashTable = op1.GetValue().GetArr()
		var val *Zval
		var key *ZendString
		for {
			var __ht *HashTable = ht
			var _p *Bucket = __ht.GetArData()
			var _end *Bucket = _p + __ht.GetNNumUsed()
			for ; _p != _end; _p++ {
				var _z *Zval = &_p.val

				if _z.GetType() == 0 {
					continue
				}
				key = _p.GetKey()
				val = _z
				if key != nil {
					ZendThrowError(nil, "Cannot unpack array with string keys")
					if free_op1 != nil {
						ZvalPtrDtorNogc(free_op1)
					}
					return 0
				} else {
					if val.GetType() == 10 && ZvalRefcountP(val) == 1 {
						val = &(*val).value.GetRef().GetVal()
					}
					if val.GetTypeFlags() != 0 {
						ZvalAddrefP(val)
					}
					if ZendHashNextIndexInsert((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())).GetValue().GetArr(), val) == nil {
						ZendCannotAddElement()
						ZvalPtrDtorNogc(val)
						break
					}
				}
			}
			break
		}
	} else if op1.GetType() == 8 {
		var ce *ZendClassEntry = op1.GetValue().GetObj().GetCe()
		var iter *ZendObjectIterator
		if ce == nil || ce.GetGetIterator() == nil {
			ZendThrowError(nil, "Only arrays and Traversables can be unpacked")
		} else {
			iter = ce.GetGetIterator()(ce, op1, 0)
			if iter == nil {
				if free_op1 != nil {
					ZvalPtrDtorNogc(free_op1)
				}
				if EG.GetException() == nil {
					ZendThrowExceptionEx(nil, 0, "Object of type %s did not create an Iterator", ce.GetName().GetVal())
				}
				return 0
			}
			if iter.GetFuncs().GetRewind() != nil {
				iter.GetFuncs().GetRewind()(iter)
			}
			for iter.GetFuncs().GetValid()(iter) == SUCCESS {
				var val *Zval
				if EG.GetException() != nil {
					break
				}
				val = iter.GetFuncs().GetGetCurrentData()(iter)
				if EG.GetException() != nil {
					break
				}
				if iter.GetFuncs().GetGetCurrentKey() != nil {
					var key Zval
					iter.GetFuncs().GetGetCurrentKey()(iter, &key)
					if EG.GetException() != nil {
						break
					}
					if key.GetType() != 4 {
						ZendThrowError(nil, g.Cond(key.GetType() == 6, "Cannot unpack Traversable with string keys", "Cannot unpack Traversable with non-integer keys"))
						ZvalPtrDtor(&key)
						break
					}
				}
				if val.GetType() == 10 {
					val = &(*val).value.GetRef().GetVal()
				}
				if val.GetTypeFlags() != 0 {
					ZvalAddrefP(val)
				}
				if ZendHashNextIndexInsert((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())).GetValue().GetArr(), val) == nil {
					ZendCannotAddElement()
					ZvalPtrDtorNogc(val)
				}
				iter.GetFuncs().GetMoveForward()(iter)
			}
			ZendIteratorDtor(iter)
		}
	} else if op1.GetType() == 10 {
		op1 = &(*op1).value.GetRef().GetVal()
		goto add_unpack_again
	} else {
		ZendThrowError(nil, "Only arrays and Traversables can be unpacked")
	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_UNSET_STATIC_PROP_SPEC_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var varname *Zval
	var name *ZendString
	var tmp_name *ZendString = nil
	var ce *ZendClassEntry
	var free_op1 ZendFreeOp
	if opline.GetOp2Type() == 1<<0 {
		ce = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))[0]
		if ce == nil {
			ce = ZendFetchClassByName((*Zval)((*byte)(opline)+int32(opline.GetOp2()).constant).GetValue().GetStr(), ((*Zval)((*byte)(opline)+int32(opline.GetOp2()).constant) + 1).GetValue().GetStr(), 0|0x200)
			if ce == nil {
				assert(EG.GetException() != nil)
				if (opline.GetOp1Type() & (1<<1 | 1<<2)) != 0 {
					ZvalPtrDtorNogc((*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar())))
				}
				return 0
			}
		}
	} else if opline.GetOp2Type() == 0 {
		ce = ZendFetchClass(nil, opline.GetOp2().GetNum())
		if ce == nil {
			assert(EG.GetException() != nil)
			if (opline.GetOp1Type() & (1<<1 | 1<<2)) != 0 {
				ZvalPtrDtorNogc((*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar())))
			}
			return 0
		}
	} else {
		ce = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar())).GetValue().GetCe()
	}
	varname = _getZvalPtrUndef(opline.GetOp1Type(), opline.GetOp1(), &free_op1, 0, execute_data, opline)
	if opline.GetOp1Type() == 1<<0 {
		name = varname.GetValue().GetStr()
	} else if varname.GetType() == 6 {
		name = varname.GetValue().GetStr()
	} else {
		if opline.GetOp1Type() == 1<<3 && varname.GetType() == 0 {
			varname = _zvalUndefinedOp1(execute_data)
		}
		name = ZvalGetTmpString(varname, &tmp_name)
	}
	ZendStdUnsetStaticProperty(ce, name)
	ZendTmpStringRelease(tmp_name)
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_ISSET_ISEMPTY_STATIC_PROP_SPEC_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var value *Zval
	var result int
	result = ZendFetchStaticPropertyAddress(&value, nil, opline.GetExtendedValue() & ^(1<<0), 3, 0, opline, execute_data)
	if (opline.GetExtendedValue() & 1 << 0) == 0 {
		result = result == SUCCESS && value.GetType() > 1 && (value.GetType() != 10 || &(*value).value.GetRef().GetVal().u1.v.type_ != 1)
	} else {
		result = result != SUCCESS || IZendIsTrue(value) == 0
	}
	for {
		if EG.GetException() != nil {
			break
		}
		if (opline + 1).GetOpcode() == 43 {
			if result != 0 {
				execute_data.SetOpline(opline + 2)
			} else {
				execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
			}
		} else if (opline + 1).GetOpcode() == 44 {
			if result == 0 {
				execute_data.SetOpline(opline + 2)
			} else {
				execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
			}
		} else {
			break
		}
		return 0
		break
	}
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_EXIT_SPEC_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	if opline.GetOp1Type() != 0 {
		var free_op1 ZendFreeOp
		var ptr *Zval = _getZvalPtr(opline.GetOp1Type(), opline.GetOp1(), &free_op1, 0, execute_data, opline)
		for {
			if ptr.GetType() == 4 {
				EG.SetExitStatus(ptr.GetValue().GetLval())
			} else {
				if (opline.GetOp1Type()&(1<<2|1<<3)) != 0 && ptr.GetType() == 10 {
					ptr = &(*ptr).value.GetRef().GetVal()
					if ptr.GetType() == 4 {
						EG.SetExitStatus(ptr.GetValue().GetLval())
						break
					}
				}
				ZendPrintZval(ptr, 0)
			}
			break
		}
		if free_op1 != nil {
			ZvalPtrDtorNogc(free_op1)
		}
	}
	_zendBailout(__FILE__, __LINE__)
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_BEGIN_SILENCE_SPEC_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	__z.GetValue().SetLval(EG.GetErrorReporting())
	__z.SetTypeInfo(4)
	if EG.GetErrorReporting() != 0 {
		for {
			EG.SetErrorReporting(0)
			if EG.GetErrorReportingIniEntry() == nil {
				var zv *Zval = ZendHashFindEx(EG.GetIniDirectives(), ZendKnownStrings[ZEND_STR_ERROR_REPORTING], 1)
				if zv != nil {
					EG.SetErrorReportingIniEntry((*ZendIniEntry)(zv.GetValue().GetPtr()))
				} else {
					break
				}
			}
			if EG.GetErrorReportingIniEntry().GetModified() == 0 {
				if EG.GetModifiedIniDirectives() == nil {
					EG.SetModifiedIniDirectives((*HashTable)(_emalloc(g.SizeOf("HashTable"))))
					_zendHashInit(EG.GetModifiedIniDirectives(), 8, nil, 0)
				}
				if ZendHashAddPtr(EG.GetModifiedIniDirectives(), ZendKnownStrings[ZEND_STR_ERROR_REPORTING], EG.GetErrorReportingIniEntry()) != nil {
					EG.GetErrorReportingIniEntry().SetOrigValue(EG.GetErrorReportingIniEntry().GetValue())
					EG.GetErrorReportingIniEntry().SetOrigModifiable(EG.GetErrorReportingIniEntry().GetModifiable())
					EG.GetErrorReportingIniEntry().SetModified(1)
				}
			}
			break
		}
	}
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_EXT_STMT_SPEC_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	if EG.GetNoExtensions() == 0 {
		ZendLlistApplyWithArgument(&ZendExtensions, LlistApplyWithArgFuncT(ZendExtensionStatementHandler), execute_data)
		execute_data.SetOpline(execute_data.GetOpline() + 1)
		return 0
	}
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_EXT_FCALL_BEGIN_SPEC_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	if EG.GetNoExtensions() == 0 {
		ZendLlistApplyWithArgument(&ZendExtensions, LlistApplyWithArgFuncT(ZendExtensionFcallBeginHandler), execute_data)
		execute_data.SetOpline(execute_data.GetOpline() + 1)
		return 0
	}
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_EXT_FCALL_END_SPEC_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	if EG.GetNoExtensions() == 0 {
		ZendLlistApplyWithArgument(&ZendExtensions, LlistApplyWithArgFuncT(ZendExtensionFcallEndHandler), execute_data)
		execute_data.SetOpline(execute_data.GetOpline() + 1)
		return 0
	}
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_DECLARE_ANON_CLASS_SPEC_HANDLER(execute_data *ZendExecuteData) int {
	var zv *Zval
	var ce *ZendClassEntry
	var opline *ZendOp = execute_data.GetOpline()
	ce = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))[0]
	if ce == nil {
		var rtd_key *ZendString = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant).GetValue().GetStr()
		zv = ZendHashFindEx(EG.GetClassTable(), rtd_key, 1)
		if zv == nil {
			for {
				assert((execute_data.GetFunc().GetOpArray().GetFnFlags() & 1 << 10) != 0)
				if ZendPreloadAutoload != nil && ZendPreloadAutoload(execute_data.GetFunc().GetOpArray().GetFilename()) == SUCCESS {
					zv = ZendHashFindEx(EG.GetClassTable(), rtd_key, 1)
					if zv != nil {
						break
					}
				}
				ZendErrorNoreturn(1<<0, "Anonymous class wasn't preloaded")
				break
			}
		}
		assert(zv != nil)
		ce = zv.GetValue().GetCe()
		if (ce.GetCeFlags() & 1 << 3) == 0 {
			if ZendDoLinkClass(ce, g.CondF1(opline.GetOp2Type() == 1<<0, func() *ZendString {
				return (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant).GetValue().GetStr()
			}, nil)) == FAILURE {
				return 0
			}
		}
		(*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))[0] = ce
	}
	(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).GetValue().SetCe(ce)
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_DECLARE_FUNCTION_SPEC_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	DoBindFunction((*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant))
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_TICKS_SPEC_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	if uint32(g.PreInc(&(EG.GetTicksCount())) >= opline.GetExtendedValue()) != 0 {
		EG.SetTicksCount(0)
		if ZendTicksFunction != nil {
			ZendTicksFunction(opline.GetExtendedValue())
			execute_data.SetOpline(execute_data.GetOpline() + 1)
			return 0
		}
	}
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_EXT_NOP_SPEC_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_NOP_SPEC_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func zend_dispatch_try_catch_finally_helper_SPEC(try_catch_offset uint32, op_num uint32, execute_data *ZendExecuteData) int {
	/* May be NULL during generator closing (only finally blocks are executed) */

	var ex *ZendObject = EG.GetException()

	/* Walk try/catch/finally structures upwards, performing the necessary actions */

	for try_catch_offset != uint32-1 {
		var try_catch *ZendTryCatchElement = &(execute_data.GetFunc()).op_array.GetTryCatchArray()[try_catch_offset]
		if op_num < try_catch.GetCatchOp() && ex != nil {

			/* Go to catch block */

			CleanupLiveVars(execute_data, op_num, try_catch.GetCatchOp())

			execute_data.SetOpline(&(execute_data.GetFunc()).op_array.GetOpcodes()[try_catch.GetCatchOp()])
			if EG.GetVmInterrupt() != 0 {
				return zend_interrupt_helper_SPEC(execute_data)
			}
			return 0
		} else if op_num < try_catch.GetFinallyOp() {

			/* Go to finally block */

			var fast_call *Zval = (*Zval)((*byte)(execute_data) + int(execute_data.GetFunc().GetOpArray().GetOpcodes()[try_catch.GetFinallyEnd()].GetOp1().GetVar()))
			CleanupLiveVars(execute_data, op_num, try_catch.GetFinallyOp())
			fast_call.GetValue().SetObj(EG.GetException())
			EG.SetException(nil)
			fast_call.SetOplineNum(uint32 - 1)

			execute_data.SetOpline(&(execute_data.GetFunc()).op_array.GetOpcodes()[try_catch.GetFinallyOp()])
			if EG.GetVmInterrupt() != 0 {
				return zend_interrupt_helper_SPEC(execute_data)
			}
			return 0
		} else if op_num < try_catch.GetFinallyEnd() {
			var fast_call *Zval = (*Zval)((*byte)(execute_data) + int(execute_data.GetFunc().GetOpArray().GetOpcodes()[try_catch.GetFinallyEnd()].GetOp1().GetVar()))

			/* cleanup incomplete RETURN statement */

			if fast_call.GetOplineNum() != uint32-1 && (execute_data.GetFunc().GetOpArray().GetOpcodes()[fast_call.GetOplineNum()].GetOp2Type()&(1<<1|1<<2)) != 0 {
				var return_value *Zval = (*Zval)((*byte)(execute_data) + int(execute_data.GetFunc().GetOpArray().GetOpcodes()[fast_call.GetOplineNum()].GetOp2().GetVar()))
				ZvalPtrDtor(return_value)
			}

			/* Chain potential exception from wrapping finally block */

			if fast_call.GetValue().GetObj() != nil {
				if ex != nil {
					ZendExceptionSetPrevious(ex, fast_call.GetValue().GetObj())
				} else {
					EG.SetException(fast_call.GetValue().GetObj())
				}
				ex = fast_call.GetValue().GetObj()
			}

			/* Chain potential exception from wrapping finally block */

		}
		try_catch_offset--
	}

	/* Uncaught exception */

	CleanupLiveVars(execute_data, op_num, 0)
	if (execute_data.GetThis().GetTypeInfo() & 1 << 24) != 0 {
		var generator *ZendGenerator = ZendGetRunningGenerator(execute_data)
		ZendGeneratorClose(generator, 1)
		return -1
	} else {

		/* We didn't execute RETURN, and have to initialize return_value */

		if execute_data.GetReturnValue() != nil {
			execute_data.GetReturnValue().SetTypeInfo(0)
		}
		return zend_leave_helper_SPEC(execute_data)
	}
}
func ZEND_HANDLE_EXCEPTION_SPEC_HANDLER(execute_data *ZendExecuteData) int {
	var throw_op *ZendOp = EG.GetOplineBeforeException()
	var throw_op_num uint32 = throw_op - execute_data.GetFunc().GetOpArray().GetOpcodes()
	var i int
	var current_try_catch_offset int = -1
	if (throw_op.GetOpcode() == 70 || throw_op.GetOpcode() == 127) && (throw_op.GetExtendedValue()&1<<0) != 0 {

		/* exceptions thrown because of loop var destruction on return/break/...
		 * are logically thrown at the end of the foreach loop, so adjust the
		 * throw_op_num.
		 */

		var range_ *ZendLiveRange = FindLiveRange(&(execute_data.GetFunc()).op_array, throw_op_num, throw_op.GetOp1().GetVar())
		throw_op_num = range_.GetEnd()
	}

	/* Find the innermost try/catch/finally the exception was thrown in */

	for i = 0; i < execute_data.GetFunc().GetOpArray().GetLastTryCatch(); i++ {
		var try_catch *ZendTryCatchElement = &(execute_data.GetFunc()).op_array.GetTryCatchArray()[i]
		if try_catch.GetTryOp() > throw_op_num {

			/* further blocks will not be relevant... */

			break

			/* further blocks will not be relevant... */

		}
		if throw_op_num < try_catch.GetCatchOp() || throw_op_num < try_catch.GetFinallyEnd() {
			current_try_catch_offset = i
		}
	}
	CleanupUnfinishedCalls(execute_data, throw_op_num)
	if (throw_op.GetResultType() & (1<<2 | 1<<1)) != 0 {
		switch throw_op.GetOpcode() {
		case 72:

		case 147:

		case 54:

		case 55:
			break
		case 109:

		case 146:
			break
		default:
			ZvalPtrDtorNogc((*Zval)((*byte)(execute_data) + int(throw_op.GetResult().GetVar())))
		}
	}
	return zend_dispatch_try_catch_finally_helper_SPEC(current_try_catch_offset, throw_op_num, execute_data)
}
func ZEND_USER_OPCODE_SPEC_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var ret int
	ret = ZendUserOpcodeHandlers[opline.GetOpcode()](execute_data)
	opline = execute_data.GetOpline()
	switch ret {
	case 0:
		return 0
	case 1:
		if (execute_data.GetThis().GetTypeInfo() & 1 << 24) != 0 {
			var generator *ZendGenerator = ZendGetRunningGenerator(execute_data)
			ZendGeneratorClose(generator, 1)
			return -1
		} else {
			return zend_leave_helper_SPEC(execute_data)
		}
	case 3:
		return 1
	case 4:
		return 2
	case 2:
		return OpcodeHandlerT(ZendVmGetOpcodeHandler(opline.GetOpcode(), opline))(execute_data)
	default:
		return OpcodeHandlerT(ZendVmGetOpcodeHandler(zend_uchar(ret&0xff), opline))(execute_data)
	}
}
func zend_yield_in_closed_generator_helper_SPEC(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	ZendThrowError(nil, "Cannot yield from finally in a force-closed generator")
	if (opline.GetOp2Type() & (1<<1 | 1<<2)) != 0 {
		ZvalPtrDtorNogc((*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar())))
	}
	if (opline.GetOp1Type() & (1<<1 | 1<<2)) != 0 {
		ZvalPtrDtorNogc((*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar())))
	}
	if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
	}
	return 0
}
func ZEND_DISCARD_EXCEPTION_SPEC_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var fast_call *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))

	/* cleanup incomplete RETURN statement */

	if fast_call.GetOplineNum() != uint32-1 && (execute_data.GetFunc().GetOpArray().GetOpcodes()[fast_call.GetOplineNum()].GetOp2Type()&(1<<1|1<<2)) != 0 {
		var return_value *Zval = (*Zval)((*byte)(execute_data) + int(execute_data.GetFunc().GetOpArray().GetOpcodes()[fast_call.GetOplineNum()].GetOp2().GetVar()))
		ZvalPtrDtor(return_value)
	}

	/* cleanup delayed exception */

	if fast_call.GetValue().GetObj() != nil {

		/* discard the previously thrown exception */

		ZendObjectRelease(fast_call.GetValue().GetObj())
		fast_call.GetValue().SetObj(nil)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FAST_CALL_SPEC_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var fast_call *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	fast_call.GetValue().SetObj(nil)

	/* set return address */

	fast_call.SetOplineNum(opline - execute_data.GetFunc().GetOpArray().GetOpcodes())

	execute_data.SetOpline((*ZendOp)((*byte)(opline) + int(opline.GetOp1().GetJmpOffset())))
	if EG.GetVmInterrupt() != 0 {
		return zend_interrupt_helper_SPEC(execute_data)
	}
	return 0
}
func ZEND_FAST_RET_SPEC_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var fast_call *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	var current_try_catch_offset uint32
	var current_op_num uint32
	if fast_call.GetOplineNum() != uint32-1 {
		var fast_ret *ZendOp = execute_data.GetFunc().GetOpArray().GetOpcodes() + fast_call.GetOplineNum()

		execute_data.SetOpline(fast_ret + 1)
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
		return 0
	}

	/* special case for unhandled exceptions */

	EG.SetException(fast_call.GetValue().GetObj())
	fast_call.GetValue().SetObj(nil)
	current_try_catch_offset = opline.GetOp2().GetNum()
	current_op_num = opline - execute_data.GetFunc().GetOpArray().GetOpcodes()
	return zend_dispatch_try_catch_finally_helper_SPEC(current_try_catch_offset, current_op_num, execute_data)
}
func ZEND_ASSERT_CHECK_SPEC_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	if EG.GetAssertions() <= 0 {
		var target *ZendOp = (*ZendOp)((*byte)(opline) + int(opline.GetOp2().GetJmpOffset()))
		if opline.GetResultType() != 0 {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
		}

		execute_data.SetOpline(target)
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
		return 0
	} else {
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
}
func ZEND_CALL_TRAMPOLINE_SPEC_HANDLER(execute_data *ZendExecuteData) int {
	var args *ZendArray = nil
	var fbc *ZendFunction = execute_data.GetFunc()
	var ret *Zval = execute_data.GetReturnValue()
	var call_info uint32 = execute_data.GetThis().GetTypeInfo() & (0<<17 | 1<<17 | 1<<21)
	var num_args uint32 = execute_data.GetThis().GetNumArgs()
	var call *ZendExecuteData
	if num_args != 0 {
		var p *Zval = (*Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(1)-1))
		var end *Zval = p + num_args
		args = _zendNewArray(num_args)
		ZendHashRealInitPacked(args)
		for {
			var __fill_ht *HashTable = args
			var __fill_bkt *Bucket = __fill_ht.GetArData() + __fill_ht.GetNNumUsed()
			var __fill_idx uint32 = __fill_ht.GetNNumUsed()
			assert((__fill_ht.GetUFlags() & 1 << 2) != 0)
			for {
				var _z1 *Zval = &__fill_bkt.val
				var _z2 *Zval = p
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				__fill_bkt.SetH(__fill_idx)
				__fill_bkt.SetKey(nil)
				__fill_bkt++
				__fill_idx++
				p++
				if p == end {
					break
				}
			}
			__fill_ht.SetNNumUsed(__fill_idx)
			__fill_ht.SetNNumOfElements(__fill_idx)
			__fill_ht.SetNNextFreeElement(__fill_idx)
			__fill_ht.SetNInternalPointer(0)
			break
		}
	}
	call = execute_data
	EG.SetCurrentExecuteData(execute_data.GetPrevExecuteData())
	execute_data = EG.GetCurrentExecuteData()
	if (fbc.GetOpArray().GetFnFlags() & 1 << 4) != 0 {
		call.SetFunc(fbc.GetOpArray().GetScope().GetCallstatic())
	} else {
		call.SetFunc(fbc.GetOpArray().GetScope().GetCall())
	}
	assert(ZendVmCalcUsedStack(2, call.GetFunc()) <= size_t((*byte)(EG.GetVmStackEnd())-(*byte)(call)))
	call.GetThis().SetNumArgs(2)
	var __z *Zval = (*Zval)(call) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(1)-1))
	var __s *ZendString = fbc.GetFunctionName()
	__z.GetValue().SetStr(__s)
	if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
		__z.SetTypeInfo(6)
	} else {
		__z.SetTypeInfo(6 | 1<<0<<8)
	}
	if args != nil {
		var __arr *ZendArray = args
		var __z *Zval = (*Zval)(call) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(2)-1))
		__z.GetValue().SetArr(__arr)
		__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
	} else {
		var __z *Zval = (*Zval)(call) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(2)-1))
		__z.GetValue().SetArr((*ZendArray)(&ZendEmptyArray))
		__z.SetTypeInfo(7)
	}
	if fbc == &EG.trampoline {
		EG.GetTrampoline().SetFunctionName(nil)
	} else {
		_efree(fbc)
	}
	fbc = call.GetFunc()
	if fbc.GetType() == 2 {
		if !(g.CondF((uintptr_t(&fbc.op_array).run_time_cache__ptr&1) != 0, func() any {
			return *((*any)((*byte)(CG.GetMapPtrBase() + uintptr_t(&fbc.op_array).run_time_cache__ptr - 1)))
		}, func() any { return any(*(&fbc.op_array.run_time_cache__ptr)) })) {
			InitFuncRunTimeCache(&fbc.op_array)
		}
		execute_data = call
		IInitFuncExecuteData(&fbc.op_array, ret, 0, execute_data)
		if ZendExecuteEx == ExecuteEx {
			return 1
		} else {
			execute_data = execute_data.GetPrevExecuteData()
			call.GetThis().SetTypeInfo(call.GetThis().GetTypeInfo() | 1<<17)
			ZendExecuteEx(call)
		}
	} else {
		var retval Zval
		assert(fbc.GetType() == 1)
		EG.SetCurrentExecuteData(call)
		if (fbc.GetFnFlags()&1<<8) != 0 && ZendVerifyInternalArgTypes(fbc, call) == 0 {
			ZendVmStackFreeCallFrame(call)
			if ret != nil {
				ret.SetTypeInfo(0)
			}
			goto call_trampoline_end
		}
		if ret == nil {
			ret = &retval
		}
		ret.SetTypeInfo(1)
		if ZendExecuteInternal == nil {

			/* saves one function call if zend_execute_internal is not used */

			fbc.GetInternalFunction().GetHandler()(call, ret)

			/* saves one function call if zend_execute_internal is not used */

		} else {
			ZendExecuteInternal(call, ret)
		}
		EG.SetCurrentExecuteData(call.GetPrevExecuteData())
	call_trampoline_end:
		ZendVmStackFreeArgs(call)
		if ret == &retval {
			ZvalPtrDtor(ret)
		}
	}
	execute_data = EG.GetCurrentExecuteData()
	if execute_data.GetFunc() == nil || (execute_data.GetFunc().GetType()&1) != 0 || (call_info&1<<17) != 0 {
		return -1
	}
	if (call_info & 1 << 21) != 0 {
		var object *ZendObject = call.GetThis().GetValue().GetObj()
		ZendObjectRelease(object)
	}
	ZendVmStackFreeCallFrame(call)
	if EG.GetException() != nil {
		ZendRethrowException(execute_data)
		return 2
	}
	execute_data.GetOpline()++
	return 2
}
func ZEND_JMP_FORWARD_SPEC_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	execute_data.SetOpline((*ZendOp)((*byte)(opline) + int(opline.GetOp1().GetJmpOffset())))
	return 0
}
func zend_interrupt_helper_SPEC(execute_data *ZendExecuteData) int {
	EG.SetVmInterrupt(0)
	if EG.GetTimedOut() != 0 {
		ZendTimeout(0)
	} else if ZendInterruptFunction != nil {
		ZendInterruptFunction(execute_data)
		return 1
	}
	return 0
}
func ZEND_INIT_FCALL_BY_NAME_SPEC_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var fbc *ZendFunction
	var function_name *Zval
	var func_ *Zval
	var call *ZendExecuteData
	fbc = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetResult().GetNum()))[0]
	if fbc == nil {
		function_name = (*Zval)((*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant))
		func_ = ZendHashFindEx(EG.GetFunctionTable(), (function_name + 1).GetValue().GetStr(), 1)
		if func_ == nil {
			return zend_undefined_function_helper_SPEC(execute_data)
		}
		fbc = func_.GetValue().GetFunc()
		if fbc.GetType() == 2 && !(g.CondF((uintptr_t(&fbc.op_array).run_time_cache__ptr&1) != 0, func() any {
			return *((*any)((*byte)(CG.GetMapPtrBase() + uintptr_t(&fbc.op_array).run_time_cache__ptr - 1)))
		}, func() any { return any(*(&fbc.op_array.run_time_cache__ptr)) })) {
			InitFuncRunTimeCache(&fbc.op_array)
		}
		(*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetResult().GetNum()))[0] = fbc
	}
	call = ZendVmStackPushCallFrame(0<<16|0<<17, fbc, opline.GetExtendedValue(), nil)
	call.SetPrevExecuteData(execute_data.GetCall())
	execute_data.SetCall(call)
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_INIT_DYNAMIC_CALL_SPEC_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var function_name *Zval
	var call *ZendExecuteData
	function_name = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
try_function_name:
	if 1<<0 != 1<<0 && function_name.GetType() == 6 {
		call = ZendInitDynamicCallString(function_name.GetValue().GetStr(), opline.GetExtendedValue())
	} else if 1<<0 != 1<<0 && function_name.GetType() == 8 {
		call = ZendInitDynamicCallObject(function_name, opline.GetExtendedValue())
	} else if function_name.GetType() == 7 {
		call = ZendInitDynamicCallArray(function_name.GetValue().GetArr(), opline.GetExtendedValue())
	} else if (1<<0&(1<<2|1<<3)) != 0 && function_name.GetType() == 10 {
		function_name = &(*function_name).value.GetRef().GetVal()
		goto try_function_name
	} else {
		if 1<<0 == 1<<3 && function_name.GetType() == 0 {
			_zvalUndefinedOp2(execute_data)
			if EG.GetException() != nil {
				return 0
			}
		}
		ZendThrowError(nil, "Function name must be a string")
		call = nil
	}
	if call == nil {
		return 0
	}
	if (1 << 0 & (1<<2 | 1<<1)) != 0 {
		if EG.GetException() != nil {
			if call != nil {
				if (call.GetFunc().GetFnFlags() & 1 << 18) != 0 {
					ZendStringReleaseEx(call.GetFunc().GetFunctionName(), 0)
					if call.GetFunc() == &EG.trampoline {
						EG.GetTrampoline().SetFunctionName(nil)
					} else {
						_efree(call.GetFunc())
					}
				}
				ZendVmStackFreeCallFrame(call)
			}
			return 0
		}
	}
	call.SetPrevExecuteData(execute_data.GetCall())
	execute_data.SetCall(call)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_INIT_NS_FCALL_BY_NAME_SPEC_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var func_name *Zval
	var func_ *Zval
	var fbc *ZendFunction
	var call *ZendExecuteData
	fbc = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetResult().GetNum()))[0]
	if fbc == nil {
		func_name = (*Zval)((*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant))
		func_ = ZendHashFindEx(EG.GetFunctionTable(), (func_name + 1).GetValue().GetStr(), 1)
		if func_ == nil {
			func_ = ZendHashFindEx(EG.GetFunctionTable(), (func_name + 2).GetValue().GetStr(), 1)
			if func_ == nil {
				return zend_undefined_function_helper_SPEC(execute_data)
			}
		}
		fbc = func_.GetValue().GetFunc()
		if fbc.GetType() == 2 && !(g.CondF((uintptr_t(&fbc.op_array).run_time_cache__ptr&1) != 0, func() any {
			return *((*any)((*byte)(CG.GetMapPtrBase() + uintptr_t(&fbc.op_array).run_time_cache__ptr - 1)))
		}, func() any { return any(*(&fbc.op_array.run_time_cache__ptr)) })) {
			InitFuncRunTimeCache(&fbc.op_array)
		}
		(*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetResult().GetNum()))[0] = fbc
	}
	call = ZendVmStackPushCallFrame(0<<16|0<<17, fbc, opline.GetExtendedValue(), nil)
	call.SetPrevExecuteData(execute_data.GetCall())
	execute_data.SetCall(call)
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_INIT_FCALL_SPEC_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var fname *Zval
	var func_ *Zval
	var fbc *ZendFunction
	var call *ZendExecuteData
	fbc = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetResult().GetNum()))[0]
	if fbc == nil {
		fname = (*Zval)((*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant))
		func_ = ZendHashFindEx(EG.GetFunctionTable(), fname.GetValue().GetStr(), 1)
		if func_ == nil {
			return zend_undefined_function_helper_SPEC(execute_data)
		}
		fbc = func_.GetValue().GetFunc()
		if fbc.GetType() == 2 && !(g.CondF((uintptr_t(&fbc.op_array).run_time_cache__ptr&1) != 0, func() any {
			return *((*any)((*byte)(CG.GetMapPtrBase() + uintptr_t(&fbc.op_array).run_time_cache__ptr - 1)))
		}, func() any { return any(*(&fbc.op_array.run_time_cache__ptr)) })) {
			InitFuncRunTimeCache(&fbc.op_array)
		}
		(*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetResult().GetNum()))[0] = fbc
	}
	call = ZendVmStackPushCallFrameEx(opline.GetOp1().GetNum(), 0<<16|0<<17, fbc, opline.GetExtendedValue(), nil)
	call.SetPrevExecuteData(execute_data.GetCall())
	execute_data.SetCall(call)
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_RECV_INIT_SPEC_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var arg_num uint32
	var param *Zval
	for {
		arg_num = opline.GetOp1().GetNum()
		param = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		if arg_num > execute_data.GetThis().GetNumArgs() {
			var default_value *Zval = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
			if (default_value.GetTypeInfo() & 0xff) == 11 {
				var cache_val *Zval = (*Zval)((*any)((*byte)(execute_data.GetRunTimeCache() + default_value.GetCacheSlot())))

				/* we keep in cache only not refcounted values */

				if cache_val.GetType() != 0 {
					var _z1 *Zval = param
					var _z2 *Zval = cache_val
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
				} else {
					var _z1 *Zval = param
					var _z2 *Zval = default_value
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					if (_t & 0xff00) != 0 {
						ZendGcAddref(&_gc.gc)
					}
					if ZvalUpdateConstantEx(param, execute_data.GetFunc().GetOpArray().GetScope()) != SUCCESS {
						ZvalPtrDtorNogc(param)
						param.SetTypeInfo(0)
						return 0
					}
					if param.GetTypeFlags() == 0 {
						var _z1 *Zval = cache_val
						var _z2 *Zval = param
						var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
						var _t uint32 = _z2.GetTypeInfo()
						_z1.GetValue().SetCounted(_gc)
						_z1.SetTypeInfo(_t)
					}
				}
				goto recv_init_check_type
			} else {
				var _z1 *Zval = param
				var _z2 *Zval = default_value
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if (_t & 0xff00) != 0 {
					ZendGcAddref(&_gc.gc)
				}
			}
		} else {
		recv_init_check_type:
			if (execute_data.GetFunc().GetOpArray().GetFnFlags() & 1 << 8) != 0 {
				var default_value *Zval = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
				if ZendVerifyRecvArgType(execute_data.GetFunc(), arg_num, param, default_value, (*any)((*byte)(execute_data.GetRunTimeCache()+opline.GetExtendedValue()))) == 0 {
					return 0
				}
			}
		}
		if g.PreInc(&opline).opcode != 64 {
			break
		}
	}
	execute_data.SetOpline(opline)
	return 0
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_INIT_DYNAMIC_CALL_SPEC_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var function_name *Zval
	var call *ZendExecuteData
	function_name = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
try_function_name:
	if (1<<1|1<<2) != 1<<0 && function_name.GetType() == 6 {
		call = ZendInitDynamicCallString(function_name.GetValue().GetStr(), opline.GetExtendedValue())
	} else if (1<<1|1<<2) != 1<<0 && function_name.GetType() == 8 {
		call = ZendInitDynamicCallObject(function_name, opline.GetExtendedValue())
	} else if function_name.GetType() == 7 {
		call = ZendInitDynamicCallArray(function_name.GetValue().GetArr(), opline.GetExtendedValue())
	} else if ((1<<1|1<<2)&(1<<2|1<<3)) != 0 && function_name.GetType() == 10 {
		function_name = &(*function_name).value.GetRef().GetVal()
		goto try_function_name
	} else {
		if (1<<1|1<<2) == 1<<3 && function_name.GetType() == 0 {
			_zvalUndefinedOp2(execute_data)
			if EG.GetException() != nil {
				return 0
			}
		}
		ZendThrowError(nil, "Function name must be a string")
		call = nil
	}
	ZvalPtrDtorNogc(free_op2)
	if call == nil {
		return 0
	}
	if ((1<<1 | 1<<2) & (1<<2 | 1<<1)) != 0 {
		if EG.GetException() != nil {
			if call != nil {
				if (call.GetFunc().GetFnFlags() & 1 << 18) != 0 {
					ZendStringReleaseEx(call.GetFunc().GetFunctionName(), 0)
					if call.GetFunc() == &EG.trampoline {
						EG.GetTrampoline().SetFunctionName(nil)
					} else {
						_efree(call.GetFunc())
					}
				}
				ZendVmStackFreeCallFrame(call)
			}
			return 0
		}
	}
	call.SetPrevExecuteData(execute_data.GetCall())
	execute_data.SetCall(call)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_RECV_SPEC_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var arg_num uint32 = opline.GetOp1().GetNum()
	if arg_num > execute_data.GetThis().GetNumArgs() {
		ZendMissingArgError(execute_data)
		return 0
	} else {
		var param *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		if ZendVerifyRecvArgType(execute_data.GetFunc(), arg_num, param, nil, (*any)((*byte)(execute_data.GetRunTimeCache()+opline.GetOp2().GetNum()))) == 0 {
			return 0
		}
	}
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_RECV_VARIADIC_SPEC_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var arg_num uint32 = opline.GetOp1().GetNum()
	var arg_count uint32 = execute_data.GetThis().GetNumArgs()
	var params *Zval
	params = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	if arg_num <= arg_count {
		var param *Zval
		var __arr *ZendArray = _zendNewArray(arg_count - arg_num + 1)
		var __z *Zval = params
		__z.GetValue().SetArr(__arr)
		__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
		ZendHashRealInitPacked(params.GetValue().GetArr())
		for {
			var __fill_ht *HashTable = params.GetValue().GetArr()
			var __fill_bkt *Bucket = __fill_ht.GetArData() + __fill_ht.GetNNumUsed()
			var __fill_idx uint32 = __fill_ht.GetNNumUsed()
			assert((__fill_ht.GetUFlags() & 1 << 2) != 0)
			param = (*Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(execute_data.GetFunc().GetOpArray().GetLastVar()+execute_data.GetFunc().GetOpArray().GetT()))
			if (execute_data.GetFunc().GetOpArray().GetFnFlags() & 1 << 8) != 0 {
				execute_data.GetThis().SetTypeInfo(execute_data.GetThis().GetTypeInfo() | 1<<19)
				for {
					ZendVerifyVariadicArgType(execute_data.GetFunc(), arg_num, param, nil, (*any)((*byte)(execute_data.GetRunTimeCache()+opline.GetOp2().GetNum())))
					if (param.GetTypeInfo() & 0xff00) != 0 {
						ZvalAddrefP(param)
					}
					var _z1 *Zval = &__fill_bkt.val
					var _z2 *Zval = param
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					__fill_bkt.SetH(__fill_idx)
					__fill_bkt.SetKey(nil)
					__fill_bkt++
					__fill_idx++
					param++
					if g.PreInc(&arg_num) > arg_count {
						break
					}
				}
			} else {
				for {
					if (param.GetTypeInfo() & 0xff00) != 0 {
						ZvalAddrefP(param)
					}
					var _z1 *Zval = &__fill_bkt.val
					var _z2 *Zval = param
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					__fill_bkt.SetH(__fill_idx)
					__fill_bkt.SetKey(nil)
					__fill_bkt++
					__fill_idx++
					param++
					if g.PreInc(&arg_num) > arg_count {
						break
					}
				}
			}
			__fill_ht.SetNNumUsed(__fill_idx)
			__fill_ht.SetNNumOfElements(__fill_idx)
			__fill_ht.SetNNextFreeElement(__fill_idx)
			__fill_ht.SetNInternalPointer(0)
			break
		}
	} else {
		var __z *Zval = params
		__z.GetValue().SetArr((*ZendArray)(&ZendEmptyArray))
		__z.SetTypeInfo(7)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_INIT_DYNAMIC_CALL_SPEC_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var function_name *Zval
	var call *ZendExecuteData
	function_name = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
try_function_name:
	if 1<<3 != 1<<0 && function_name.GetType() == 6 {
		call = ZendInitDynamicCallString(function_name.GetValue().GetStr(), opline.GetExtendedValue())
	} else if 1<<3 != 1<<0 && function_name.GetType() == 8 {
		call = ZendInitDynamicCallObject(function_name, opline.GetExtendedValue())
	} else if function_name.GetType() == 7 {
		call = ZendInitDynamicCallArray(function_name.GetValue().GetArr(), opline.GetExtendedValue())
	} else if (1<<3&(1<<2|1<<3)) != 0 && function_name.GetType() == 10 {
		function_name = &(*function_name).value.GetRef().GetVal()
		goto try_function_name
	} else {
		if 1<<3 == 1<<3 && function_name.GetType() == 0 {
			_zvalUndefinedOp2(execute_data)
			if EG.GetException() != nil {
				return 0
			}
		}
		ZendThrowError(nil, "Function name must be a string")
		call = nil
	}
	if call == nil {
		return 0
	}
	if (1 << 3 & (1<<2 | 1<<1)) != 0 {
		if EG.GetException() != nil {
			if call != nil {
				if (call.GetFunc().GetFnFlags() & 1 << 18) != 0 {
					ZendStringReleaseEx(call.GetFunc().GetFunctionName(), 0)
					if call.GetFunc() == &EG.trampoline {
						EG.GetTrampoline().SetFunctionName(nil)
					} else {
						_efree(call.GetFunc())
					}
				}
				ZendVmStackFreeCallFrame(call)
			}
			return 0
		}
	}
	call.SetPrevExecuteData(execute_data.GetCall())
	execute_data.SetCall(call)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_BW_NOT_SPEC_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	if op1.GetTypeInfo() == 4 {
		var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		__z.GetValue().SetLval(^(op1.GetValue().GetLval()))
		__z.SetTypeInfo(4)
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
	if 1<<0 == 1<<3 && op1.GetType() == 0 {
		op1 = _zvalUndefinedOp1(execute_data)
	}
	BitwiseNotFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op1)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_BOOL_NOT_SPEC_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var val *Zval
	val = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	if val.GetTypeInfo() == 3 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	} else if val.GetTypeInfo() <= 3 {

		/* The result and op1 can be the same cv zval */

		var orig_val_type uint32 = val.GetTypeInfo()
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
		if 1<<0 == 1<<3 && orig_val_type == 0 {
			_zvalUndefinedOp1(execute_data)
			execute_data.SetOpline(execute_data.GetOpline() + 1)
			return 0
		}
	} else {
		if IZendIsTrue(val) == 0 {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
		} else {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
		}
		execute_data.SetOpline(execute_data.GetOpline() + 1)
		return 0
	}
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_ECHO_SPEC_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var z *Zval
	z = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	if z.GetType() == 6 {
		var str *ZendString = z.GetValue().GetStr()
		if str.GetLen() != 0 {
			ZendWrite(str.GetVal(), str.GetLen())
		}
	} else {
		var str *ZendString = ZvalGetStringFunc(z)
		if str.GetLen() != 0 {
			ZendWrite(str.GetVal(), str.GetLen())
		} else if 1<<0 == 1<<3 && z.GetType() == 0 {
			_zvalUndefinedOp1(execute_data)
		}
		ZendStringReleaseEx(str, 0)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_JMPZ_SPEC_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var val *Zval
	val = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	if val.GetTypeInfo() == 3 {
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	} else if val.GetTypeInfo() <= 3 {
		if 1<<0 == 1<<3 && val.GetTypeInfo() == 0 {
			_zvalUndefinedOp1(execute_data)
			if EG.GetException() != nil {
				return 0
			}
		}

		execute_data.SetOpline((*ZendOp)((*byte)(opline) + int(opline.GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
		return 0
	}
	if IZendIsTrue(val) != 0 {
		opline++
	} else {
		opline = (*ZendOp)((*byte)(opline) + int(opline.GetOp2().GetJmpOffset()))
	}
	if EG.GetException() != nil {
		return 0
	}
	execute_data.SetOpline(opline)
	if EG.GetVmInterrupt() != 0 {
		return zend_interrupt_helper_SPEC(execute_data)
	}
	return 0
}
func ZEND_JMPNZ_SPEC_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var val *Zval
	val = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	if val.GetTypeInfo() == 3 {

		execute_data.SetOpline((*ZendOp)((*byte)(opline) + int(opline.GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
		return 0
	} else if val.GetTypeInfo() <= 3 {
		if 1<<0 == 1<<3 && val.GetTypeInfo() == 0 {
			_zvalUndefinedOp1(execute_data)
			if EG.GetException() != nil {
				return 0
			}
		}
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
	if IZendIsTrue(val) != 0 {
		opline = (*ZendOp)((*byte)(opline) + int(opline.GetOp2().GetJmpOffset()))
	} else {
		opline++
	}
	if EG.GetException() != nil {
		return 0
	}
	execute_data.SetOpline(opline)
	if EG.GetVmInterrupt() != 0 {
		return zend_interrupt_helper_SPEC(execute_data)
	}
	return 0
}
func ZEND_JMPZNZ_SPEC_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var val *Zval
	val = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	if val.GetTypeInfo() == 3 {
		execute_data.SetOpline((*ZendOp)((*byte)(opline) + int(opline.GetExtendedValue())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
		return 0
	} else if val.GetTypeInfo() <= 3 {
		if 1<<0 == 1<<3 && val.GetTypeInfo() == 0 {
			_zvalUndefinedOp1(execute_data)
			if EG.GetException() != nil {
				return 0
			}
		}

		execute_data.SetOpline((*ZendOp)((*byte)(opline) + int(opline.GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
		return 0
	}
	if IZendIsTrue(val) != 0 {
		opline = (*ZendOp)((*byte)(opline) + int(opline.GetExtendedValue()))
	} else {
		opline = (*ZendOp)((*byte)(opline) + int(opline.GetOp2().GetJmpOffset()))
	}
	if EG.GetException() != nil {
		return 0
	}
	execute_data.SetOpline(opline)
	if EG.GetVmInterrupt() != 0 {
		return zend_interrupt_helper_SPEC(execute_data)
	}
	return 0
}
func ZEND_JMPZ_EX_SPEC_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var val *Zval
	var ret int
	val = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	if val.GetTypeInfo() == 3 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	} else if val.GetTypeInfo() <= 3 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
		if 1<<0 == 1<<3 && val.GetTypeInfo() == 0 {
			_zvalUndefinedOp1(execute_data)
			if EG.GetException() != nil {
				return 0
			}
		}

		execute_data.SetOpline((*ZendOp)((*byte)(opline) + int(opline.GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
		return 0
	}
	ret = IZendIsTrue(val)
	if ret != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
		opline++
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
		opline = (*ZendOp)((*byte)(opline) + int(opline.GetOp2().GetJmpOffset()))
	}
	if EG.GetException() != nil {
		return 0
	}
	execute_data.SetOpline(opline)
	if EG.GetVmInterrupt() != 0 {
		return zend_interrupt_helper_SPEC(execute_data)
	}
	return 0
}
func ZEND_JMPNZ_EX_SPEC_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var val *Zval
	var ret int
	val = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	if val.GetTypeInfo() == 3 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)

		execute_data.SetOpline((*ZendOp)((*byte)(opline) + int(opline.GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
		return 0
	} else if val.GetTypeInfo() <= 3 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
		if 1<<0 == 1<<3 && val.GetTypeInfo() == 0 {
			_zvalUndefinedOp1(execute_data)
			execute_data.SetOpline(execute_data.GetOpline() + 1)
			return 0
		} else {
			assert(EG.GetException() == nil)
			execute_data.SetOpline(opline + 1)
			return 0
		}
	}
	ret = IZendIsTrue(val)
	if ret != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
		opline = (*ZendOp)((*byte)(opline) + int(opline.GetOp2().GetJmpOffset()))
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
		opline++
	}
	if EG.GetException() != nil {
		return 0
	}
	execute_data.SetOpline(opline)
	if EG.GetVmInterrupt() != 0 {
		return zend_interrupt_helper_SPEC(execute_data)
	}
	return 0
}
func ZEND_RETURN_SPEC_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var retval_ptr *Zval
	var return_value *Zval
	var free_op1 ZendFreeOp
	retval_ptr = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	return_value = execute_data.GetReturnValue()
	if 1<<0 == 1<<3 && retval_ptr.GetTypeInfo() == 0 {
		retval_ptr = _zvalUndefinedOp1(execute_data)
		if return_value != nil {
			return_value.SetTypeInfo(1)
		}
	} else if return_value == nil {
		if (1 << 0 & (1<<2 | 1<<1)) != 0 {
			if free_op1.GetTypeFlags() != 0 && ZvalDelrefP(free_op1) == 0 {
				RcDtorFunc(free_op1.GetValue().GetCounted())
			}
		}
	} else {
		if (1 << 0 & (1<<0 | 1<<1)) != 0 {
			var _z1 *Zval = return_value
			var _z2 *Zval = retval_ptr
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if 1<<0 == 1<<0 {
				if (return_value.GetTypeInfo() & 0xff00) != 0 {
					ZvalAddrefP(return_value)
				}
			}
		} else if 1<<0 == 1<<3 {
			for {
				if (retval_ptr.GetTypeInfo() & 0xff00) != 0 {
					if (retval_ptr.GetTypeInfo() & 0xff) != 10 {
						if (execute_data.GetThis().GetTypeInfo() & 1 << 16) == 0 {
							var ref *ZendRefcounted = retval_ptr.GetValue().GetCounted()
							var _z1 *Zval = return_value
							var _z2 *Zval = retval_ptr
							var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
							var _t uint32 = _z2.GetTypeInfo()
							_z1.GetValue().SetCounted(_gc)
							_z1.SetTypeInfo(_t)
							if (ref.GetGc().GetTypeInfo() & (0xfffffc00 | 1<<4<<0)) == 1<<4<<0 {
								GcPossibleRoot(ref)
							}
							retval_ptr.SetTypeInfo(1)
							break
						} else {
							ZvalAddrefP(retval_ptr)
						}
					} else {
						retval_ptr = &(*retval_ptr).value.GetRef().GetVal()
						if (retval_ptr.GetTypeInfo() & 0xff00) != 0 {
							ZvalAddrefP(retval_ptr)
						}
					}
				}
				var _z1 *Zval = return_value
				var _z2 *Zval = retval_ptr
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				break
			}
		} else {
			if retval_ptr.GetType() == 10 {
				var ref *ZendRefcounted = retval_ptr.GetValue().GetCounted()
				retval_ptr = &(*retval_ptr).value.GetRef().GetVal()
				var _z1 *Zval = return_value
				var _z2 *Zval = retval_ptr
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if ZendGcDelref(&ref.gc) == 0 {
					_efree(ref)
				} else if (retval_ptr.GetTypeInfo() & 0xff00) != 0 {
					ZvalAddrefP(retval_ptr)
				}
			} else {
				var _z1 *Zval = return_value
				var _z2 *Zval = retval_ptr
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
			}
		}
	}
	return zend_leave_helper_SPEC(execute_data)
}
func ZEND_RETURN_BY_REF_SPEC_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var retval_ptr *Zval
	for {
		if (1<<0&(1<<0|1<<1)) != 0 || 1<<0 == 1<<2 && opline.GetExtendedValue() == 1<<1 {

			/* Not supposed to happen, but we'll allow it */

			ZendError(1<<3, "Only variable references should be returned by reference")
			retval_ptr = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
			if execute_data.GetReturnValue() == nil {

			} else {
				if 1<<0 == 1<<2 && retval_ptr.GetType() == 10 {
					var _z1 *Zval = execute_data.GetReturnValue()
					var _z2 *Zval = retval_ptr
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					break
				}
				var _ref *ZendReference = (*ZendReference)(_emalloc(g.SizeOf("zend_reference")))
				ZendGcSetRefcount(&_ref.gc, 1)
				_ref.GetGc().SetTypeInfo(10)
				var _z1 *Zval = &_ref.val
				var _z2 *Zval = retval_ptr
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				_ref.GetSources().SetPtr(nil)
				execute_data.GetReturnValue().GetValue().SetRef(_ref)
				execute_data.GetReturnValue().SetTypeInfo(10 | 1<<0<<8)
				if 1<<0 == 1<<0 {
					if retval_ptr.GetTypeFlags() != 0 {
						ZvalAddrefP(retval_ptr)
					}
				}
			}
			break
		}
		retval_ptr = nil
		if 1<<0 == 1<<2 {
			assert(retval_ptr != &EG.uninitialized_zval)
			if opline.GetExtendedValue() == 1<<0 && retval_ptr.GetType() != 10 {
				ZendError(1<<3, "Only variable references should be returned by reference")
				if execute_data.GetReturnValue() != nil {
					var _ref *ZendReference = (*ZendReference)(_emalloc(g.SizeOf("zend_reference")))
					ZendGcSetRefcount(&_ref.gc, 1)
					_ref.GetGc().SetTypeInfo(10)
					var _z1 *Zval = &_ref.val
					var _z2 *Zval = retval_ptr
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					_ref.GetSources().SetPtr(nil)
					execute_data.GetReturnValue().GetValue().SetRef(_ref)
					execute_data.GetReturnValue().SetTypeInfo(10 | 1<<0<<8)
				}
				break
			}
		}
		if execute_data.GetReturnValue() != nil {
			if retval_ptr.GetType() == 10 {
				ZvalAddrefP(retval_ptr)
			} else {
				var _z *Zval = retval_ptr
				var _ref *ZendReference = (*ZendReference)(_emalloc(g.SizeOf("zend_reference")))
				ZendGcSetRefcount(&_ref.gc, 2)
				_ref.GetGc().SetTypeInfo(10)
				var _z1 *Zval = &_ref.val
				var _z2 *Zval = _z
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				_ref.GetSources().SetPtr(nil)
				_z.GetValue().SetRef(_ref)
				_z.SetTypeInfo(10 | 1<<0<<8)
			}
			var __z *Zval = execute_data.GetReturnValue()
			__z.GetValue().SetRef(retval_ptr.GetValue().GetRef())
			__z.SetTypeInfo(10 | 1<<0<<8)
		}
		break
	}
	return zend_leave_helper_SPEC(execute_data)
}
func ZEND_GENERATOR_RETURN_SPEC_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var retval *Zval
	var generator *ZendGenerator = ZendGetRunningGenerator(execute_data)
	retval = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)

	/* Copy return value into generator->retval */

	if (1 << 0 & (1<<0 | 1<<1)) != 0 {
		var _z1 *Zval = &generator.retval
		var _z2 *Zval = retval
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if 1<<0 == 1<<0 {
			if (generator.GetRetval().GetTypeInfo() & 0xff00) != 0 {
				ZvalAddrefP(&(generator.GetRetval()))
			}
		}
	} else if 1<<0 == 1<<3 {
		var _z3 *Zval = retval
		if (_z3.GetTypeInfo() & 0xff00) != 0 {
			if (_z3.GetTypeInfo() & 0xff) == 10 {
				_z3 = &(*_z3).value.GetRef().GetVal()
				if (_z3.GetTypeInfo() & 0xff00) != 0 {
					ZvalAddrefP(_z3)
				}
			} else {
				ZvalAddrefP(_z3)
			}
		}
		var _z1 *Zval = &generator.retval
		var _z2 *Zval = _z3
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
	} else {
		if retval.GetType() == 10 {
			var ref *ZendRefcounted = retval.GetValue().GetCounted()
			retval = &(*retval).value.GetRef().GetVal()
			var _z1 *Zval = &generator.retval
			var _z2 *Zval = retval
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if ZendGcDelref(&ref.gc) == 0 {
				_efree(ref)
			} else if (retval.GetTypeInfo() & 0xff00) != 0 {
				ZvalAddrefP(retval)
			}
		} else {
			var _z1 *Zval = &generator.retval
			var _z2 *Zval = retval
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
		}
	}

	/* Close the generator to free up resources */

	ZendGeneratorClose(generator, 1)

	/* Pass execution back to handling code */

	return -1

	/* Pass execution back to handling code */
}
func ZEND_THROW_SPEC_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var value *Zval
	value = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	for {
		if 1<<0 == 1<<0 || value.GetType() != 8 {
			if (1<<0&(1<<2|1<<3)) != 0 && value.GetType() == 10 {
				value = &(*value).value.GetRef().GetVal()
				if value.GetType() == 8 {
					break
				}
			}
			if 1<<0 == 1<<3 && value.GetType() == 0 {
				_zvalUndefinedOp1(execute_data)
				if EG.GetException() != nil {
					return 0
				}
			}
			ZendThrowError(nil, "Can only throw objects")
			return 0
		}
		break
	}
	ZendExceptionSave()
	if 1<<0 != 1<<1 {
		if value.GetTypeFlags() != 0 {
			ZvalAddrefP(value)
		}
	}
	ZendThrowExceptionObject(value)
	ZendExceptionRestore()
	return 0
}
func ZEND_CATCH_SPEC_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var ce *ZendClassEntry
	var catch_ce *ZendClassEntry
	var exception *ZendObject
	var ex *Zval

	/* Check whether an exception has been thrown, if not, jump over code */

	ZendExceptionRestore()
	if EG.GetException() == nil {

		execute_data.SetOpline((*ZendOp)((*byte)(opline) + int(opline.GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
		return 0
	}
	catch_ce = (*any)((*byte)(execute_data.GetRunTimeCache() + (opline.GetExtendedValue() & ^(1 << 0))))[0]
	if catch_ce == nil {
		catch_ce = ZendFetchClassByName((*Zval)((*byte)(opline)+int32(opline.GetOp1()).constant).GetValue().GetStr(), ((*Zval)((*byte)(opline)+int32(opline.GetOp1()).constant) + 1).GetValue().GetStr(), 0x80)
		(*any)((*byte)(execute_data.GetRunTimeCache() + (opline.GetExtendedValue() & ^(1 << 0))))[0] = catch_ce
	}
	ce = EG.GetException().GetCe()
	if ce != catch_ce {
		if catch_ce == nil || InstanceofFunction(ce, catch_ce) == 0 {
			if (opline.GetExtendedValue() & 1 << 0) != 0 {
				ZendRethrowException(execute_data)
				return 0
			}

			execute_data.SetOpline((*ZendOp)((*byte)(opline) + int(opline.GetOp2().GetJmpOffset())))
			if EG.GetVmInterrupt() != 0 {
				return zend_interrupt_helper_SPEC(execute_data)
			}
			return 0
		}
	}
	exception = EG.GetException()
	ex = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))

	/* Always perform a strict assignment. There is a reasonable expectation that if you
	 * write "catch (Exception $e)" then $e will actually be instanceof Exception. As such,
	 * we should not permit coercion to string here. */

	var tmp Zval
	var __z *Zval = &tmp
	__z.GetValue().SetObj(exception)
	__z.SetTypeInfo(8 | 1<<0<<8 | 1<<1<<8)
	EG.SetException(nil)
	ZendAssignToVariable(ex, &tmp, 1<<1, 1)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_SEND_VAL_SPEC_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var value *Zval
	var arg *Zval
	value = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	arg = (*Zval)((*byte)(execute_data.GetCall()) + int(opline.GetResult().GetVar()))
	var _z1 *Zval = arg
	var _z2 *Zval = value
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	if 1<<0 == 1<<0 {
		if (arg.GetTypeInfo() & 0xff00) != 0 {
			ZvalAddrefP(arg)
		}
	}
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_SEND_VAL_EX_SPEC_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var value *Zval
	var arg *Zval
	var arg_num uint32 = opline.GetOp2().GetNum()
	if ZendCheckArgSendType(execute_data.GetCall().GetFunc(), arg_num, 1) != 0 {
	send_val_by_ref:
		return zend_cannot_pass_by_ref_helper_SPEC(execute_data)
	}
	value = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	arg = (*Zval)((*byte)(execute_data.GetCall()) + int(opline.GetResult().GetVar()))
	var _z1 *Zval = arg
	var _z2 *Zval = value
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	if 1<<0 == 1<<0 {
		if (arg.GetTypeInfo() & 0xff00) != 0 {
			ZvalAddrefP(arg)
		}
	}
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_SEND_VAL_EX_SPEC_CONST_QUICK_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var value *Zval
	var arg *Zval
	var arg_num uint32 = opline.GetOp2().GetNum()
	if (execute_data.GetCall().GetFunc().GetQuickArgFlags() >> (arg_num + 3) * 2 & 1) != 0 {
		goto send_val_by_ref
	}
	value = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	arg = (*Zval)((*byte)(execute_data.GetCall()) + int(opline.GetResult().GetVar()))
	var _z1 *Zval = arg
	var _z2 *Zval = value
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	if 1<<0 == 1<<0 {
		if (arg.GetTypeInfo() & 0xff00) != 0 {
			ZvalAddrefP(arg)
		}
	}
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_SEND_USER_SPEC_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var arg *Zval
	var param *Zval
	if ZendCheckArgSendType(execute_data.GetCall().GetFunc(), opline.GetOp2().GetNum(), 1) != 0 {
		ZendParamMustBeRef(execute_data.GetCall().GetFunc(), opline.GetOp2().GetNum())
	}
	arg = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	param = (*Zval)((*byte)(execute_data.GetCall()) + int(opline.GetResult().GetVar()))
	var _z1 *Zval = param
	var _z2 *Zval = arg
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	if (_t & 0xff00) != 0 {
		ZendGcAddref(&_gc.gc)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_BOOL_SPEC_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var val *Zval
	val = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	if val.GetTypeInfo() == 3 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else if val.GetTypeInfo() <= 3 {

		/* The result and op1 can be the same cv zval */

		var orig_val_type uint32 = val.GetTypeInfo()
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
		if 1<<0 == 1<<3 && orig_val_type == 0 {
			_zvalUndefinedOp1(execute_data)
			execute_data.SetOpline(execute_data.GetOpline() + 1)
			return 0
		}
	} else {
		if IZendIsTrue(val) != 0 {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
		} else {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
		}
		execute_data.SetOpline(execute_data.GetOpline() + 1)
		return 0
	}
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_CLONE_SPEC_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var obj *Zval
	var ce *ZendClassEntry
	var scope *ZendClassEntry
	var clone *ZendFunction
	var clone_call ZendObjectCloneObjT
	obj = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	if 1<<0 == 0 && obj.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	for {
		if 1<<0 == 1<<0 || 1<<0 != 0 && obj.GetType() != 8 {
			if (1<<0&(1<<2|1<<3)) != 0 && obj.GetType() == 10 {
				obj = &(*obj).value.GetRef().GetVal()
				if obj.GetType() == 8 {
					break
				}
			}
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
			if 1<<0 == 1<<3 && obj.GetType() == 0 {
				_zvalUndefinedOp1(execute_data)
				if EG.GetException() != nil {
					return 0
				}
			}
			ZendThrowError(nil, "__clone method called on non-object")
			return 0
		}
		break
	}
	ce = obj.GetValue().GetObj().GetCe()
	clone = ce.GetClone()
	clone_call = obj.GetValue().GetObj().GetHandlers().GetCloneObj()
	if clone_call == nil {
		ZendThrowError(nil, "Trying to clone an uncloneable object of class %s", ce.GetName().GetVal())
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
		return 0
	}
	if clone != nil && (clone.GetFnFlags()&1<<0) == 0 {
		scope = execute_data.GetFunc().GetOpArray().GetScope()
		if clone.GetScope() != scope {
			if (clone.GetFnFlags()&1<<2) != 0 || ZendCheckProtected(g.CondF(clone.GetPrototype() != nil, func() *ZendClassEntry { return clone.GetPrototype().GetScope() }, func() *ZendClassEntry { return clone.GetScope() }), scope) == 0 {
				ZendWrongCloneCall(clone, scope)
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
				return 0
			}
		}
	}
	var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	__z.GetValue().SetObj(clone_call(obj))
	__z.SetTypeInfo(8 | 1<<0<<8 | 1<<1<<8)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_CAST_SPEC_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var expr *Zval
	var result *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	var ht *HashTable
	expr = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	switch opline.GetExtendedValue() {
	case 1:
		result.SetTypeInfo(1)
		break
	case 16:
		if ZendIsTrue(expr) != 0 {
			result.SetTypeInfo(3)
		} else {
			result.SetTypeInfo(2)
		}
		break
	case 4:
		var __z *Zval = result
		__z.GetValue().SetLval(ZvalGetLong(expr))
		__z.SetTypeInfo(4)
		break
	case 5:
		var __z *Zval = result
		__z.GetValue().SetDval(ZvalGetDouble(expr))
		__z.SetTypeInfo(5)
		break
	case 6:
		var __z *Zval = result
		var __s *ZendString = ZvalGetString(expr)
		__z.GetValue().SetStr(__s)
		if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
			__z.SetTypeInfo(6)
		} else {
			__z.SetTypeInfo(6 | 1<<0<<8)
		}
		break
	default:
		if (1 << 0 & (1<<2 | 1<<3)) != 0 {
			if expr.GetType() == 10 {
				expr = &(*expr).value.GetRef().GetVal()
			}
		}

		/* If value is already of correct type, return it directly */

		if expr.GetType() == opline.GetExtendedValue() {
			var _z1 *Zval = result
			var _z2 *Zval = expr
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if 1<<0 == 1<<0 {
				if (result.GetTypeInfo() & 0xff00) != 0 {
					ZvalAddrefP(result)
				}
			} else if 1<<0 != 1<<1 {
				if (result.GetTypeInfo() & 0xff00) != 0 {
					ZvalAddrefP(result)
				}
			}
			execute_data.SetOpline(execute_data.GetOpline() + 1)
			return 0
		}
		if opline.GetExtendedValue() == 7 {
			if 1<<0 == 1<<0 || expr.GetType() != 8 || expr.GetValue().GetObj().GetCe() == ZendCeClosure {
				if expr.GetType() != 1 {
					var __arr *ZendArray = _zendNewArray(1)
					var __z *Zval = result
					__z.GetValue().SetArr(__arr)
					__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
					expr = ZendHashIndexAddNew(result.GetValue().GetArr(), 0, expr)
					if 1<<0 == 1<<0 {
						if (expr.GetTypeInfo() & 0xff00) != 0 {
							ZvalAddrefP(expr)
						}
					} else {
						if (expr.GetTypeInfo() & 0xff00) != 0 {
							ZvalAddrefP(expr)
						}
					}
				} else {
					var __z *Zval = result
					__z.GetValue().SetArr((*ZendArray)(&ZendEmptyArray))
					__z.SetTypeInfo(7)
				}
			} else {
				var obj_ht *HashTable = ZendGetPropertiesFor(expr, ZEND_PROP_PURPOSE_ARRAY_CAST)
				if obj_ht != nil {

					/* fast copy */

					var __arr *ZendArray = ZendProptableToSymtable(obj_ht, expr.GetValue().GetObj().GetCe().GetDefaultPropertiesCount() != 0 || expr.GetValue().GetObj().GetHandlers() != &StdObjectHandlers || (ZvalGcFlags(obj_ht.GetGc().GetTypeInfo())&1<<5) != 0)
					var __z *Zval = result
					__z.GetValue().SetArr(__arr)
					__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
					if obj_ht != nil && (ZvalGcFlags(obj_ht.GetGc().GetTypeInfo())&1<<6) == 0 && ZendGcDelref(&obj_ht.gc) == 0 {
						ZendArrayDestroy(obj_ht)
					}
				} else {
					var __z *Zval = result
					__z.GetValue().SetArr((*ZendArray)(&ZendEmptyArray))
					__z.SetTypeInfo(7)
				}
			}
		} else {
			var __z *Zval = result
			__z.GetValue().SetObj(ZendObjectsNew(ZendStandardClassDef))
			__z.SetTypeInfo(8 | 1<<0<<8 | 1<<1<<8)
			if expr.GetType() == 7 {
				ht = ZendSymtableToProptable(expr.GetValue().GetArr())
				if (ZvalGcFlags(ht.GetGc().GetTypeInfo()) & 1 << 6) != 0 {

					/* TODO: try not to duplicate immutable arrays as well ??? */

					ht = ZendArrayDup(ht)

					/* TODO: try not to duplicate immutable arrays as well ??? */

				}
				result.GetValue().GetObj().SetProperties(ht)
			} else if expr.GetType() != 1 {
				ht = _zendNewArray(1)
				result.GetValue().GetObj().SetProperties(ht)
				expr = ZendHashAddNew(ht, ZendKnownStrings[ZEND_STR_SCALAR], expr)
				if 1<<0 == 1<<0 {
					if (expr.GetTypeInfo() & 0xff00) != 0 {
						ZvalAddrefP(expr)
					}
				} else {
					if (expr.GetTypeInfo() & 0xff00) != 0 {
						ZvalAddrefP(expr)
					}
				}
			}
		}
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_INCLUDE_OR_EVAL_SPEC_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var new_op_array *ZendOpArray
	var inc_filename *Zval
	inc_filename = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	new_op_array = ZendIncludeOrEval(inc_filename, opline.GetExtendedValue())
	if EG.GetException() != nil {
		if new_op_array != (*ZendOpArray)(zend_intptr_t-1) && new_op_array != nil {
			DestroyOpArray(new_op_array)
			_efree(new_op_array)
		}
		if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
		}
		return 0
	} else if new_op_array == (*ZendOpArray)(zend_intptr_t-1) {
		if opline.GetResultType() != 0 {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
		}
	} else if new_op_array != nil {
		var return_value *Zval = nil
		var call *ZendExecuteData
		if opline.GetResultType() != 0 {
			return_value = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		}
		new_op_array.SetScope(execute_data.GetFunc().GetOpArray().GetScope())
		call = ZendVmStackPushCallFrame(execute_data.GetThis().GetTypeInfo()&(8|1<<0<<8|1<<1<<8)|(1<<16|0<<17)|1<<20, (*ZendFunction)(new_op_array), 0, execute_data.GetThis().GetValue().GetPtr())
		if (execute_data.GetThis().GetTypeInfo() & 1 << 20) != 0 {
			call.SetSymbolTable(execute_data.GetSymbolTable())
		} else {
			call.SetSymbolTable(ZendRebuildSymbolTable())
		}
		call.SetPrevExecuteData(execute_data)
		IInitCodeExecuteData(call, new_op_array, return_value)
		if ZendExecuteEx == ExecuteEx {
			return 1
		} else {
			call.GetThis().SetTypeInfo(call.GetThis().GetTypeInfo() | 1<<17)
			ZendExecuteEx(call)
			ZendVmStackFreeCallFrame(call)
		}
		DestroyOpArray(new_op_array)
		_efree(new_op_array)
		if EG.GetException() != nil {
			ZendRethrowException(execute_data)
			if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
			}
			return 0
		}
	} else if opline.GetResultType() != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_FE_RESET_R_SPEC_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var array_ptr *Zval
	var result *Zval
	array_ptr = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	if array_ptr.GetType() == 7 {
		result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var _z1 *Zval = result
		var _z2 *Zval = array_ptr
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if 1<<0 != 1<<1 && (result.GetTypeInfo()&0xff00) != 0 {
			ZvalAddrefP(array_ptr)
		}
		result.SetFePos(0)
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	} else if 1<<0 != 1<<0 && array_ptr.GetType() == 8 {
		if array_ptr.GetValue().GetObj().GetCe().GetGetIterator() == nil {
			var properties *HashTable
			if array_ptr.GetValue().GetObj().GetProperties() != nil && ZendGcRefcount(&(array_ptr.GetValue().GetObj().GetProperties()).gc) > 1 {
				if (ZvalGcFlags(array_ptr.GetValue().GetObj().GetProperties().GetGc().GetTypeInfo()) & 1 << 6) == 0 {
					ZendGcDelref(&(array_ptr.GetValue().GetObj().GetProperties()).gc)
				}
				array_ptr.GetValue().GetObj().SetProperties(ZendArrayDup(array_ptr.GetValue().GetObj().GetProperties()))
			}
			properties = array_ptr.GetValue().GetObj().GetHandlers().GetGetProperties()(&(*array_ptr))
			result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var _z1 *Zval = result
			var _z2 *Zval = array_ptr
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if 1<<0 != 1<<1 {
				ZvalAddrefP(array_ptr)
			}
			if properties.GetNNumOfElements() == 0 {
				result.SetFeIterIdx(uint32 - 1)
				if EG.GetException() != nil {
					return 0
				}
				execute_data.SetOpline((*ZendOp)((*byte)(opline) + int(opline.GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
				return 0
			}
			result.SetFeIterIdx(ZendHashIteratorAdd(properties, 0))
			execute_data.SetOpline(execute_data.GetOpline() + 1)
			return 0
		} else {
			var is_empty ZendBool = ZendFeResetIterator(array_ptr, 0, opline, execute_data)
			if EG.GetException() != nil {
				return 0
			} else if is_empty != 0 {

				execute_data.SetOpline((*ZendOp)((*byte)(opline) + int(opline.GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
				return 0
			} else {
				assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			}
		}
	} else {
		ZendError(1<<1, "Invalid argument supplied for foreach()")
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetFeIterIdx(uint32 - 1)
		if EG.GetException() != nil {
			return 0
		}
		execute_data.SetOpline((*ZendOp)((*byte)(opline) + int(opline.GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
		return 0
	}
}
func ZEND_FE_RESET_RW_SPEC_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var array_ptr *Zval
	var array_ref *Zval
	if 1<<0 == 1<<2 || 1<<0 == 1<<3 {
		array_ptr = nil
		array_ref = array_ptr
		if array_ref.GetType() == 10 {
			array_ptr = &(*array_ref).value.GetRef().GetVal()
		}
	} else {
		array_ptr = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
		array_ref = array_ptr
	}
	if array_ptr.GetType() == 7 {
		if 1<<0 == 1<<2 || 1<<0 == 1<<3 {
			if array_ptr == array_ref {
				var _ref *ZendReference = (*ZendReference)(_emalloc(g.SizeOf("zend_reference")))
				ZendGcSetRefcount(&_ref.gc, 1)
				_ref.GetGc().SetTypeInfo(10)
				var _z1 *Zval = &_ref.val
				var _z2 *Zval = array_ref
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				_ref.GetSources().SetPtr(nil)
				array_ref.GetValue().SetRef(_ref)
				array_ref.SetTypeInfo(10 | 1<<0<<8)
				array_ptr = &(*array_ref).value.GetRef().GetVal()
			}
			ZvalAddrefP(array_ref)
			var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var _z2 *Zval = array_ref
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
		} else {
			array_ref = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var _ref *ZendReference = (*ZendReference)(_emalloc(g.SizeOf("zend_reference")))
			ZendGcSetRefcount(&_ref.gc, 1)
			_ref.GetGc().SetTypeInfo(10)
			var _z1 *Zval = &_ref.val
			var _z2 *Zval = array_ptr
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			_ref.GetSources().SetPtr(nil)
			array_ref.GetValue().SetRef(_ref)
			array_ref.SetTypeInfo(10 | 1<<0<<8)
			array_ptr = &(*array_ref).value.GetRef().GetVal()
		}
		if 1<<0 == 1<<0 {
			var __arr *ZendArray = ZendArrayDup(array_ptr.GetValue().GetArr())
			var __z *Zval = array_ptr
			__z.GetValue().SetArr(__arr)
			__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
		} else {
			var _zv *Zval = array_ptr
			var _arr *ZendArray = _zv.GetValue().GetArr()
			if ZendGcRefcount(&_arr.gc) > 1 {
				if _zv.GetTypeFlags() != 0 {
					ZendGcDelref(&_arr.gc)
				}
				var __arr *ZendArray = ZendArrayDup(_arr)
				var __z *Zval = _zv
				__z.GetValue().SetArr(__arr)
				__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
			}
		}
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetFeIterIdx(ZendHashIteratorAdd(array_ptr.GetValue().GetArr(), 0))
		if 1<<0 == 1<<2 {

		}
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	} else if 1<<0 != 1<<0 && array_ptr.GetType() == 8 {
		if array_ptr.GetValue().GetObj().GetCe().GetGetIterator() == nil {
			var properties *HashTable
			if 1<<0 == 1<<2 || 1<<0 == 1<<3 {
				if array_ptr == array_ref {
					var _ref *ZendReference = (*ZendReference)(_emalloc(g.SizeOf("zend_reference")))
					ZendGcSetRefcount(&_ref.gc, 1)
					_ref.GetGc().SetTypeInfo(10)
					var _z1 *Zval = &_ref.val
					var _z2 *Zval = array_ref
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					_ref.GetSources().SetPtr(nil)
					array_ref.GetValue().SetRef(_ref)
					array_ref.SetTypeInfo(10 | 1<<0<<8)
					array_ptr = &(*array_ref).value.GetRef().GetVal()
				}
				ZvalAddrefP(array_ref)
				var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				var _z2 *Zval = array_ref
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
			} else {
				array_ptr = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				var _z1 *Zval = array_ptr
				var _z2 *Zval = array_ref
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
			}
			if array_ptr.GetValue().GetObj().GetProperties() != nil && ZendGcRefcount(&(array_ptr.GetValue().GetObj().GetProperties()).gc) > 1 {
				if (ZvalGcFlags(array_ptr.GetValue().GetObj().GetProperties().GetGc().GetTypeInfo()) & 1 << 6) == 0 {
					ZendGcDelref(&(array_ptr.GetValue().GetObj().GetProperties()).gc)
				}
				array_ptr.GetValue().GetObj().SetProperties(ZendArrayDup(array_ptr.GetValue().GetObj().GetProperties()))
			}
			properties = array_ptr.GetValue().GetObj().GetHandlers().GetGetProperties()(&(*array_ptr))
			if properties.GetNNumOfElements() == 0 {
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetFeIterIdx(uint32 - 1)
				if EG.GetException() != nil {
					return 0
				}
				execute_data.SetOpline((*ZendOp)((*byte)(opline) + int(opline.GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
				return 0
			}
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetFeIterIdx(ZendHashIteratorAdd(properties, 0))
			execute_data.SetOpline(execute_data.GetOpline() + 1)
			return 0
		} else {
			var is_empty ZendBool = ZendFeResetIterator(array_ptr, 1, opline, execute_data)
			if 1<<0 == 1<<2 {

			}
			if EG.GetException() != nil {
				return 0
			} else if is_empty != 0 {

				execute_data.SetOpline((*ZendOp)((*byte)(opline) + int(opline.GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
				return 0
			} else {
				assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			}
		}
	} else {
		ZendError(1<<1, "Invalid argument supplied for foreach()")
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetFeIterIdx(uint32 - 1)
		if 1<<0 == 1<<2 {

		}
		if EG.GetException() != nil {
			return 0
		}
		execute_data.SetOpline((*ZendOp)((*byte)(opline) + int(opline.GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
		return 0
	}
}
func ZEND_JMP_SET_SPEC_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var value *Zval
	var ref *Zval = nil
	var ret int
	value = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	if (1<<0 == 1<<2 || 1<<0 == 1<<3) && value.GetType() == 10 {
		if 1<<0 == 1<<2 {
			ref = value
		}
		value = &(*value).value.GetRef().GetVal()
	}
	ret = IZendIsTrue(value)
	if EG.GetException() != nil {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
		return 0
	}
	if ret != 0 {
		var result *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var _z1 *Zval = result
		var _z2 *Zval = value
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if 1<<0 == 1<<0 {
			if (result.GetTypeInfo() & 0xff00) != 0 {
				ZvalAddrefP(result)
			}
		} else if 1<<0 == 1<<3 {
			if (result.GetTypeInfo() & 0xff00) != 0 {
				ZvalAddrefP(result)
			}
		} else if 1<<0 == 1<<2 && ref != nil {
			var r *ZendReference = ref.GetValue().GetRef()
			if ZendGcDelref(&r.gc) == 0 {
				_efree(r)
			} else if (result.GetTypeInfo() & 0xff00) != 0 {
				ZvalAddrefP(result)
			}
		}

		execute_data.SetOpline((*ZendOp)((*byte)(opline) + int(opline.GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
		return 0
	}
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_COALESCE_SPEC_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var value *Zval
	var ref *Zval = nil
	value = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	if (1<<0&(1<<2|1<<3)) != 0 && value.GetType() == 10 {
		if (1 << 0 & 1 << 2) != 0 {
			ref = value
		}
		value = &(*value).value.GetRef().GetVal()
	}
	if value.GetType() > 1 {
		var result *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var _z1 *Zval = result
		var _z2 *Zval = value
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if 1<<0 == 1<<0 {
			if (result.GetTypeInfo() & 0xff00) != 0 {
				ZvalAddrefP(result)
			}
		} else if 1<<0 == 1<<3 {
			if (result.GetTypeInfo() & 0xff00) != 0 {
				ZvalAddrefP(result)
			}
		} else if (1<<0&1<<2) != 0 && ref != nil {
			var r *ZendReference = ref.GetValue().GetRef()
			if ZendGcDelref(&r.gc) == 0 {
				_efree(r)
			} else if (result.GetTypeInfo() & 0xff00) != 0 {
				ZvalAddrefP(result)
			}
		}

		execute_data.SetOpline((*ZendOp)((*byte)(opline) + int(opline.GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
		return 0
	}
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_QM_ASSIGN_SPEC_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var value *Zval
	var result *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	value = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	if 1<<0 == 1<<3 && value.GetType() == 0 {
		_zvalUndefinedOp1(execute_data)
		result.SetTypeInfo(1)
		execute_data.SetOpline(execute_data.GetOpline() + 1)
		return 0
	}
	if 1<<0 == 1<<3 {
		var _z3 *Zval = value
		if (_z3.GetTypeInfo() & 0xff00) != 0 {
			if (_z3.GetTypeInfo() & 0xff) == 10 {
				_z3 = &(*_z3).value.GetRef().GetVal()
				if (_z3.GetTypeInfo() & 0xff00) != 0 {
					ZvalAddrefP(_z3)
				}
			} else {
				ZvalAddrefP(_z3)
			}
		}
		var _z1 *Zval = result
		var _z2 *Zval = _z3
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
	} else if 1<<0 == 1<<2 {
		if value.GetType() == 10 {
			var _z1 *Zval = result
			var _z2 *Zval = &(*value).value.GetRef().GetVal()
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if ZvalDelrefP(value) == 0 {
				_efree(value.GetValue().GetRef())
			} else if (result.GetTypeInfo() & 0xff00) != 0 {
				ZvalAddrefP(result)
			}
		} else {
			var _z1 *Zval = result
			var _z2 *Zval = value
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
		}
	} else {
		var _z1 *Zval = result
		var _z2 *Zval = value
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if 1<<0 == 1<<0 {
			if (result.GetTypeInfo() & 0xff00) != 0 {
				ZvalAddrefP(result)
			}
		}
	}
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_DECLARE_CLASS_SPEC_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	DoBindClass((*Zval)((*byte)(opline)+int32(opline.GetOp1()).constant), g.CondF1(opline.GetOp2Type() == 1<<0, func() *ZendString {
		return (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant).GetValue().GetStr()
	}, nil))
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_YIELD_FROM_SPEC_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var generator *ZendGenerator = ZendGetRunningGenerator(execute_data)
	var val *Zval
	val = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	if (generator.GetFlags() & ZEND_GENERATOR_FORCED_CLOSE) != 0 {
		ZendThrowError(nil, "Cannot use \"yield from\" in a force-closed generator")
		if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
		}
		return 0
	}
	if val.GetType() == 7 {
		var _z1 *Zval = &generator.values
		var _z2 *Zval = val
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if 1<<0 != 1<<1 && (val.GetTypeInfo()&0xff00) != 0 {
			ZvalAddrefP(val)
		}
		generator.GetValues().SetFePos(0)
	} else if 1<<0 != 1<<0 && val.GetType() == 8 && val.GetValue().GetObj().GetCe().GetGetIterator() != nil {
		var ce *ZendClassEntry = val.GetValue().GetObj().GetCe()
		if ce == ZendCeGenerator {
			var new_gen *ZendGenerator = (*ZendGenerator)(val.GetValue().GetObj())
			if 1<<0 != 1<<1 {
				ZvalAddrefP(val)
			}
			if new_gen.GetRetval().GetType() == 0 {
				if ZendGeneratorGetCurrent(new_gen) == generator {
					ZendThrowError(nil, "Impossible to yield from the Generator being currently run")
					ZvalPtrDtor(val)
					if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
						(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
					}
					return 0
				} else {
					ZendGeneratorYieldFrom(generator, new_gen)
				}
			} else if new_gen.GetExecuteData() == nil {
				ZendThrowError(nil, "Generator passed to yield from was aborted without proper return and is unable to continue")
				ZvalPtrDtor(val)
				if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
					(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
				}
				return 0
			} else {
				if opline.GetResultType() != 0 {
					var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
					var _z2 *Zval = &new_gen.retval
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					if (_t & 0xff00) != 0 {
						ZendGcAddref(&_gc.gc)
					}
				}
				assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			}
		} else {
			var iter *ZendObjectIterator = ce.GetGetIterator()(ce, val, 0)
			if iter == nil || EG.GetException() != nil {
				if EG.GetException() == nil {
					ZendThrowError(nil, "Object of type %s did not create an Iterator", ce.GetName().GetVal())
				}
				if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
					(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
				}
				return 0
			}
			iter.SetIndex(0)
			if iter.GetFuncs().GetRewind() != nil {
				iter.GetFuncs().GetRewind()(iter)
				if EG.GetException() != nil {
					ZendObjectRelease(&iter.std)
					if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
						(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
					}
					return 0
				}
			}
			var __z *Zval = &generator.values
			__z.GetValue().SetObj(&iter.std)
			__z.SetTypeInfo(8 | 1<<0<<8 | 1<<1<<8)
		}
	} else {
		ZendThrowError(nil, "Can use \"yield from\" only with arrays and Traversables")
		if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
		}
		return 0
	}

	/* This is the default return value
	 * when the expression is a Generator, it will be overwritten in zend_generator_resume() */

	if opline.GetResultType() != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
	}

	/* This generator has no send target (though the generator we delegate to might have one) */

	generator.SetSendTarget(nil)

	/* We increment to the next op, so we are at the correct position when the
	 * generator is resumed. */

	execute_data.GetOpline()++

	/* The GOTO VM uses a local opline variable. We need to set the opline
	 * variable in execute_data so we don't resume at an old position. */

	return -1
}
func ZEND_STRLEN_SPEC_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var value *Zval
	value = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	if value.GetType() == 6 {
		var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		__z.GetValue().SetLval(value.GetValue().GetStr().GetLen())
		__z.SetTypeInfo(4)
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	} else {
		var strict ZendBool
		if (1<<0&(1<<2|1<<3)) != 0 && value.GetType() == 10 {
			value = &(*value).value.GetRef().GetVal()
			if value.GetType() == 6 {
				var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				__z.GetValue().SetLval(value.GetValue().GetStr().GetLen())
				__z.SetTypeInfo(4)
				assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			}
		}
		if 1<<0 == 1<<3 && value.GetType() == 0 {
			value = _zvalUndefinedOp1(execute_data)
		}
		strict = (execute_data.GetFunc().GetFnFlags() & 1 << 31) != 0
		for {
			if strict == 0 {
				var str *ZendString
				var tmp Zval
				var _z1 *Zval = &tmp
				var _z2 *Zval = value
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if (_t & 0xff00) != 0 {
					ZendGcAddref(&_gc.gc)
				}
				if ZendParseArgStrWeak(&tmp, &str) != 0 {
					var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
					__z.GetValue().SetLval(str.GetLen())
					__z.SetTypeInfo(4)
					ZvalPtrDtor(&tmp)
					break
				}
				ZvalPtrDtor(&tmp)
			}
			if EG.GetException() == nil {
				ZendInternalTypeError(strict, "strlen() expects parameter 1 to be string, %s given", ZendGetTypeByConst(value.GetType()))
			}
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
			break
		}
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_TYPE_CHECK_SPEC_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var value *Zval
	var result int = 0
	value = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	if (opline.GetExtendedValue() >> uint32(*value).u1.v.type_ & 1) != 0 {
	type_check_resource:
		if value.GetType() != 9 || nil != ZendRsrcListGetRsrcType(value.GetValue().GetRes()) {
			result = 1
		}
	} else if (1<<0&(1<<3|1<<2)) != 0 && value.GetType() == 10 {
		value = &(*value).value.GetRef().GetVal()
		if (opline.GetExtendedValue() >> uint32(*value).u1.v.type_ & 1) != 0 {
			goto type_check_resource
		}
	} else if 1<<0 == 1<<3 && value.GetType() == 0 {
		result = (1 << 1 & opline.GetExtendedValue()) != 0
		_zvalUndefinedOp1(execute_data)
		if EG.GetException() != nil {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
			return 0
		}
	}
	if (1 << 0 & (1<<1 | 1<<2)) != 0 {
		for {
			if EG.GetException() != nil {
				break
			}
			if (opline + 1).GetOpcode() == 43 {
				if result != 0 {
					execute_data.SetOpline(opline + 2)
				} else {
					execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
					if EG.GetVmInterrupt() != 0 {
						return zend_interrupt_helper_SPEC(execute_data)
					}
				}
			} else if (opline + 1).GetOpcode() == 44 {
				if result == 0 {
					execute_data.SetOpline(opline + 2)
				} else {
					execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
					if EG.GetVmInterrupt() != 0 {
						return zend_interrupt_helper_SPEC(execute_data)
					}
				}
			} else {
				break
			}
			return 0
			break
		}
		if result != 0 {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
		} else {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
		}
		execute_data.SetOpline(execute_data.GetOpline() + 1)
		return 0
	} else {
		for {

			if (opline + 1).GetOpcode() == 43 {
				if result != 0 {
					execute_data.SetOpline(opline + 2)
				} else {
					execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
					if EG.GetVmInterrupt() != 0 {
						return zend_interrupt_helper_SPEC(execute_data)
					}
				}
			} else if (opline + 1).GetOpcode() == 44 {
				if result == 0 {
					execute_data.SetOpline(opline + 2)
				} else {
					execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
					if EG.GetVmInterrupt() != 0 {
						return zend_interrupt_helper_SPEC(execute_data)
					}
				}
			} else {
				break
			}
			return 0
			break
		}
		if result != 0 {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
		} else {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
		}
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
}
func ZEND_DEFINED_SPEC_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var c *ZendConstant
	c = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))[0]
	if c != nil {
		if (uintptr_t(c) & 1 << 0) == 0 {
		defined_true:
			if (opline + 1).GetOpcode() == 44 {
				execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
				return 0
			} else if (opline + 1).GetOpcode() == 43 {
				execute_data.SetOpline(opline + 2)
				return 0
			}
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
			assert(EG.GetException() == nil)
			execute_data.SetOpline(opline + 1)
			return 0
		} else if EG.GetZendConstants().GetNNumOfElements() == uintptr_t(c)>>1 {
		defined_false:
			if (opline + 1).GetOpcode() == 44 {
				execute_data.SetOpline(opline + 2)
				return 0
			} else if (opline + 1).GetOpcode() == 43 {
				execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
				return 0
			}
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
			assert(EG.GetException() == nil)
			execute_data.SetOpline(opline + 1)
			return 0
		}
	}
	if ZendQuickCheckConstant((*Zval)((*byte)(opline)+int32(opline.GetOp1()).constant), opline, execute_data) != SUCCESS {
		(*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))[0] = any(uintptr_t(EG.GetZendConstants().GetNNumOfElements())<<1 | 1<<0)
		goto defined_false
	} else {
		goto defined_true
	}
}
func ZEND_QM_ASSIGN_LONG_SPEC_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var value *Zval
	value = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	__z.GetValue().SetLval(value.GetValue().GetLval())
	__z.SetTypeInfo(4)
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_QM_ASSIGN_DOUBLE_SPEC_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var value *Zval
	value = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	__z.GetValue().SetDval(value.GetValue().GetDval())
	__z.SetTypeInfo(5)
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_QM_ASSIGN_NOREF_SPEC_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var value *Zval
	value = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	var _z2 *Zval = value
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_SEND_VAL_SIMPLE_SPEC_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var value *Zval
	var arg *Zval
	value = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	arg = (*Zval)((*byte)(execute_data.GetCall()) + int(opline.GetResult().GetVar()))
	var _z1 *Zval = arg
	var _z2 *Zval = value
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_SEND_VAL_EX_SIMPLE_SPEC_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var value *Zval
	var arg *Zval
	var arg_num uint32 = opline.GetOp2().GetNum()
	if (execute_data.GetCall().GetFunc().GetQuickArgFlags() >> (arg_num + 3) * 2 & 1) != 0 {
		return zend_cannot_pass_by_ref_helper_SPEC(execute_data)
	}
	value = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	arg = (*Zval)((*byte)(execute_data.GetCall()) + int(opline.GetResult().GetVar()))
	var _z1 *Zval = arg
	var _z2 *Zval = value
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_ADD_SPEC_CONST_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result *Zval
	var d1 float64
	var d2 float64
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if 1<<0 == 1<<0 && 1<<0 == 1<<0 {

	} else if op1.GetTypeInfo() == 4 {
		if op2.GetTypeInfo() == 4 {
			result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			FastLongAddFunction(result, op1, op2)
			assert(EG.GetException() == nil)
			execute_data.SetOpline(opline + 1)
			return 0
		} else if op2.GetTypeInfo() == 5 {
			d1 = float64(op1.GetValue().GetLval())
			d2 = op2.GetValue().GetDval()
			goto add_double
		}
	} else if op1.GetTypeInfo() == 5 {
		if op2.GetTypeInfo() == 5 {
			d1 = op1.GetValue().GetDval()
			d2 = op2.GetValue().GetDval()
		add_double:
			result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var __z *Zval = result
			__z.GetValue().SetDval(d1 + d2)
			__z.SetTypeInfo(5)
			assert(EG.GetException() == nil)
			execute_data.SetOpline(opline + 1)
			return 0
		} else if op2.GetTypeInfo() == 4 {
			d1 = op1.GetValue().GetDval()
			d2 = float64(op2.GetValue().GetLval())
			goto add_double
		}
	}
	return zend_add_helper_SPEC(op1, op2, execute_data)
}
func ZEND_SUB_SPEC_CONST_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result *Zval
	var d1 float64
	var d2 float64
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if 1<<0 == 1<<0 && 1<<0 == 1<<0 {

	} else if op1.GetTypeInfo() == 4 {
		if op2.GetTypeInfo() == 4 {
			result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			FastLongSubFunction(result, op1, op2)
			assert(EG.GetException() == nil)
			execute_data.SetOpline(opline + 1)
			return 0
		} else if op2.GetTypeInfo() == 5 {
			d1 = float64(op1.GetValue().GetLval())
			d2 = op2.GetValue().GetDval()
			goto sub_double
		}
	} else if op1.GetTypeInfo() == 5 {
		if op2.GetTypeInfo() == 5 {
			d1 = op1.GetValue().GetDval()
			d2 = op2.GetValue().GetDval()
		sub_double:
			result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var __z *Zval = result
			__z.GetValue().SetDval(d1 - d2)
			__z.SetTypeInfo(5)
			assert(EG.GetException() == nil)
			execute_data.SetOpline(opline + 1)
			return 0
		} else if op2.GetTypeInfo() == 4 {
			d1 = op1.GetValue().GetDval()
			d2 = float64(op2.GetValue().GetLval())
			goto sub_double
		}
	}
	return zend_sub_helper_SPEC(op1, op2, execute_data)
}
func ZEND_MUL_SPEC_CONST_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result *Zval
	var d1 float64
	var d2 float64
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if 1<<0 == 1<<0 && 1<<0 == 1<<0 {

	} else if op1.GetTypeInfo() == 4 {
		if op2.GetTypeInfo() == 4 {
			var overflow ZendLong
			result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var __lres long = op1.GetValue().GetLval() * op2.GetValue().GetLval()
			var __dres long__double = long__double(op1.GetValue().GetLval() * long__double(op2.GetValue().GetLval()))
			var __delta long__double = long__double(__lres - __dres)
			if g.Assign(&overflow, __dres+__delta != __dres) {
				result.GetValue().SetDval(__dres)
			} else {
				result.GetValue().SetLval(__lres)
			}
			if overflow != 0 {
				result.SetTypeInfo(5)
			} else {
				result.SetTypeInfo(4)
			}
			assert(EG.GetException() == nil)
			execute_data.SetOpline(opline + 1)
			return 0
		} else if op2.GetTypeInfo() == 5 {
			d1 = float64(op1.GetValue().GetLval())
			d2 = op2.GetValue().GetDval()
			goto mul_double
		}
	} else if op1.GetTypeInfo() == 5 {
		if op2.GetTypeInfo() == 5 {
			d1 = op1.GetValue().GetDval()
			d2 = op2.GetValue().GetDval()
		mul_double:
			result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var __z *Zval = result
			__z.GetValue().SetDval(d1 * d2)
			__z.SetTypeInfo(5)
			assert(EG.GetException() == nil)
			execute_data.SetOpline(opline + 1)
			return 0
		} else if op2.GetTypeInfo() == 4 {
			d1 = op1.GetValue().GetDval()
			d2 = float64(op2.GetValue().GetLval())
			goto mul_double
		}
	}
	return zend_mul_helper_SPEC(op1, op2, execute_data)
}
func ZEND_DIV_SPEC_CONST_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	FastDivFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op1, op2)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_MOD_SPEC_CONST_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result *Zval
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if 1<<0 == 1<<0 && 1<<0 == 1<<0 {

	} else if op1.GetTypeInfo() == 4 {
		if op2.GetTypeInfo() == 4 {
			result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			if op2.GetValue().GetLval() == 0 {
				return zend_mod_by_zero_helper_SPEC(execute_data)
			} else if op2.GetValue().GetLval() == -1 {

				/* Prevent overflow error/crash if op1==ZEND_LONG_MIN */

				var __z *Zval = result
				__z.GetValue().SetLval(0)
				__z.SetTypeInfo(4)
			} else {
				var __z *Zval = result
				__z.GetValue().SetLval(op1.GetValue().GetLval() % op2.GetValue().GetLval())
				__z.SetTypeInfo(4)
			}
			assert(EG.GetException() == nil)
			execute_data.SetOpline(opline + 1)
			return 0
		}
	}
	return zend_mod_helper_SPEC(op1, op2, execute_data)
}
func ZEND_SL_SPEC_CONST_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if 1<<0 == 1<<0 && 1<<0 == 1<<0 {

	} else if op1.GetTypeInfo() == 4 && op2.GetTypeInfo() == 4 && zend_ulong(*op2).value.lval < 8*8 {

		/* Perform shift on unsigned numbers to get well-defined wrap behavior. */

		var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		__z.GetValue().SetLval(zend_long(zend_ulong(*op1).value.lval << op2.GetValue().GetLval()))
		__z.SetTypeInfo(4)
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
	return zend_shift_left_helper_SPEC(op1, op2, execute_data)
}
func ZEND_SR_SPEC_CONST_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if 1<<0 == 1<<0 && 1<<0 == 1<<0 {

	} else if op1.GetTypeInfo() == 4 && op2.GetTypeInfo() == 4 && zend_ulong(*op2).value.lval < 8*8 {
		var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		__z.GetValue().SetLval(op1.GetValue().GetLval() >> op2.GetValue().GetLval())
		__z.SetTypeInfo(4)
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
	return zend_shift_right_helper_SPEC(op1, op2, execute_data)
}
func ZEND_POW_SPEC_CONST_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	PowFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op1, op2)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_IS_IDENTICAL_SPEC_CONST_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result ZendBool
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	result = FastIsIdenticalFunction(op1, op2)
	for {
		if EG.GetException() != nil {
			break
		}
		if (opline + 1).GetOpcode() == 43 {
			if result != 0 {
				execute_data.SetOpline(opline + 2)
			} else {
				execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
			}
		} else if (opline + 1).GetOpcode() == 44 {
			if result == 0 {
				execute_data.SetOpline(opline + 2)
			} else {
				execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
			}
		} else {
			break
		}
		return 0
		break
	}
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_IS_NOT_IDENTICAL_SPEC_CONST_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result ZendBool
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	result = FastIsNotIdenticalFunction(op1, op2)
	for {
		if EG.GetException() != nil {
			break
		}
		if (opline + 1).GetOpcode() == 43 {
			if result != 0 {
				execute_data.SetOpline(opline + 2)
			} else {
				execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
			}
		} else if (opline + 1).GetOpcode() == 44 {
			if result == 0 {
				execute_data.SetOpline(opline + 2)
			} else {
				execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
			}
		} else {
			break
		}
		return 0
		break
	}
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_IS_EQUAL_SPEC_CONST_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if 1<<0 == 1<<0 && 1<<0 == 1<<0 {

	} else if op1.GetType() == 4 {
		if op2.GetType() == 4 {
			if op1.GetValue().GetLval() == op2.GetValue().GetLval() {
			is_equal_true:
				if (opline + 1).GetOpcode() == 44 {
					execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
					if EG.GetVmInterrupt() != 0 {
						return zend_interrupt_helper_SPEC(execute_data)
					}
					return 0
				} else if (opline + 1).GetOpcode() == 43 {
					execute_data.SetOpline(opline + 2)
					return 0
				}
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
				assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			} else {
			is_equal_false:
				if (opline + 1).GetOpcode() == 44 {
					execute_data.SetOpline(opline + 2)
					return 0
				} else if (opline + 1).GetOpcode() == 43 {
					execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
					if EG.GetVmInterrupt() != 0 {
						return zend_interrupt_helper_SPEC(execute_data)
					}
					return 0
				}
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
				assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			}
		} else if op2.GetType() == 5 {
			d1 = float64(op1.GetValue().GetLval())
			d2 = op2.GetValue().GetDval()
			goto is_equal_double
		}
	} else if op1.GetType() == 5 {
		if op2.GetType() == 5 {
			d1 = op1.GetValue().GetDval()
			d2 = op2.GetValue().GetDval()
		is_equal_double:
			if d1 == d2 {
				goto is_equal_true
			} else {
				goto is_equal_false
			}
		} else if op2.GetType() == 4 {
			d1 = op1.GetValue().GetDval()
			d2 = float64(op2.GetValue().GetLval())
			goto is_equal_double
		}
	} else if op1.GetType() == 6 {
		if op2.GetType() == 6 {
			var result int = ZendFastEqualStrings(op1.GetValue().GetStr(), op2.GetValue().GetStr())
			if (1 << 0 & (1<<1 | 1<<2)) != 0 {
				ZvalPtrDtorStr(op1)
			}
			if (1 << 0 & (1<<1 | 1<<2)) != 0 {
				ZvalPtrDtorStr(op2)
			}
			if result != 0 {
				goto is_equal_true
			} else {
				goto is_equal_false
			}
		}
	}
	return zend_is_equal_helper_SPEC(op1, op2, execute_data)
}
func ZEND_IS_NOT_EQUAL_SPEC_CONST_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if 1<<0 == 1<<0 && 1<<0 == 1<<0 {

	} else if op1.GetType() == 4 {
		if op2.GetType() == 4 {
			if op1.GetValue().GetLval() != op2.GetValue().GetLval() {
			is_not_equal_true:
				if (opline + 1).GetOpcode() == 44 {
					execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
					if EG.GetVmInterrupt() != 0 {
						return zend_interrupt_helper_SPEC(execute_data)
					}
					return 0
				} else if (opline + 1).GetOpcode() == 43 {
					execute_data.SetOpline(opline + 2)
					return 0
				}
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
				assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			} else {
			is_not_equal_false:
				if (opline + 1).GetOpcode() == 44 {
					execute_data.SetOpline(opline + 2)
					return 0
				} else if (opline + 1).GetOpcode() == 43 {
					execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
					if EG.GetVmInterrupt() != 0 {
						return zend_interrupt_helper_SPEC(execute_data)
					}
					return 0
				}
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
				assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			}
		} else if op2.GetType() == 5 {
			d1 = float64(op1.GetValue().GetLval())
			d2 = op2.GetValue().GetDval()
			goto is_not_equal_double
		}
	} else if op1.GetType() == 5 {
		if op2.GetType() == 5 {
			d1 = op1.GetValue().GetDval()
			d2 = op2.GetValue().GetDval()
		is_not_equal_double:
			if d1 != d2 {
				goto is_not_equal_true
			} else {
				goto is_not_equal_false
			}
		} else if op2.GetType() == 4 {
			d1 = op1.GetValue().GetDval()
			d2 = float64(op2.GetValue().GetLval())
			goto is_not_equal_double
		}
	} else if op1.GetType() == 6 {
		if op2.GetType() == 6 {
			var result int = ZendFastEqualStrings(op1.GetValue().GetStr(), op2.GetValue().GetStr())
			if (1 << 0 & (1<<1 | 1<<2)) != 0 {
				ZvalPtrDtorStr(op1)
			}
			if (1 << 0 & (1<<1 | 1<<2)) != 0 {
				ZvalPtrDtorStr(op2)
			}
			if result == 0 {
				goto is_not_equal_true
			} else {
				goto is_not_equal_false
			}
		}
	}
	return zend_is_not_equal_helper_SPEC(op1, op2, execute_data)
}
func ZEND_IS_SMALLER_SPEC_CONST_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if 1<<0 == 1<<0 && 1<<0 == 1<<0 {

	} else if op1.GetTypeInfo() == 4 {
		if op2.GetTypeInfo() == 4 {
			if op1.GetValue().GetLval() < op2.GetValue().GetLval() {
			is_smaller_true:
				if (opline + 1).GetOpcode() == 44 {
					execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
					if EG.GetVmInterrupt() != 0 {
						return zend_interrupt_helper_SPEC(execute_data)
					}
					return 0
				} else if (opline + 1).GetOpcode() == 43 {
					execute_data.SetOpline(opline + 2)
					return 0
				}
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
				assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			} else {
			is_smaller_false:
				if (opline + 1).GetOpcode() == 44 {
					execute_data.SetOpline(opline + 2)
					return 0
				} else if (opline + 1).GetOpcode() == 43 {
					execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
					if EG.GetVmInterrupt() != 0 {
						return zend_interrupt_helper_SPEC(execute_data)
					}
					return 0
				}
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
				assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			}
		} else if op2.GetTypeInfo() == 5 {
			d1 = float64(op1.GetValue().GetLval())
			d2 = op2.GetValue().GetDval()
			goto is_smaller_double
		}
	} else if op1.GetTypeInfo() == 5 {
		if op2.GetTypeInfo() == 5 {
			d1 = op1.GetValue().GetDval()
			d2 = op2.GetValue().GetDval()
		is_smaller_double:
			if d1 < d2 {
				goto is_smaller_true
			} else {
				goto is_smaller_false
			}
		} else if op2.GetTypeInfo() == 4 {
			d1 = op1.GetValue().GetDval()
			d2 = float64(op2.GetValue().GetLval())
			goto is_smaller_double
		}
	}
	return zend_is_smaller_helper_SPEC(op1, op2, execute_data)
}
func ZEND_IS_SMALLER_OR_EQUAL_SPEC_CONST_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if 1<<0 == 1<<0 && 1<<0 == 1<<0 {

	} else if op1.GetTypeInfo() == 4 {
		if op2.GetTypeInfo() == 4 {
			if op1.GetValue().GetLval() <= op2.GetValue().GetLval() {
			is_smaller_or_equal_true:
				if (opline + 1).GetOpcode() == 44 {
					execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
					if EG.GetVmInterrupt() != 0 {
						return zend_interrupt_helper_SPEC(execute_data)
					}
					return 0
				} else if (opline + 1).GetOpcode() == 43 {
					execute_data.SetOpline(opline + 2)
					return 0
				}
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
				assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			} else {
			is_smaller_or_equal_false:
				if (opline + 1).GetOpcode() == 44 {
					execute_data.SetOpline(opline + 2)
					return 0
				} else if (opline + 1).GetOpcode() == 43 {
					execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
					if EG.GetVmInterrupt() != 0 {
						return zend_interrupt_helper_SPEC(execute_data)
					}
					return 0
				}
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
				assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			}
		} else if op2.GetTypeInfo() == 5 {
			d1 = float64(op1.GetValue().GetLval())
			d2 = op2.GetValue().GetDval()
			goto is_smaller_or_equal_double
		}
	} else if op1.GetTypeInfo() == 5 {
		if op2.GetTypeInfo() == 5 {
			d1 = op1.GetValue().GetDval()
			d2 = op2.GetValue().GetDval()
		is_smaller_or_equal_double:
			if d1 <= d2 {
				goto is_smaller_or_equal_true
			} else {
				goto is_smaller_or_equal_false
			}
		} else if op2.GetTypeInfo() == 4 {
			d1 = op1.GetValue().GetDval()
			d2 = float64(op2.GetValue().GetLval())
			goto is_smaller_or_equal_double
		}
	}
	return zend_is_smaller_or_equal_helper_SPEC(op1, op2, execute_data)
}
func ZEND_SPACESHIP_SPEC_CONST_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	CompareFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op1, op2)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_BW_OR_SPEC_CONST_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if 1<<0 == 1<<0 && 1<<0 == 1<<0 {

	} else if op1.GetTypeInfo() == 4 && op2.GetTypeInfo() == 4 {
		var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		__z.GetValue().SetLval(op1.GetValue().GetLval() | op2.GetValue().GetLval())
		__z.SetTypeInfo(4)
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
	return zend_bw_or_helper_SPEC(op1, op2, execute_data)
}
func ZEND_BW_AND_SPEC_CONST_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if 1<<0 == 1<<0 && 1<<0 == 1<<0 {

	} else if op1.GetTypeInfo() == 4 && op2.GetTypeInfo() == 4 {
		var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		__z.GetValue().SetLval(op1.GetValue().GetLval() & op2.GetValue().GetLval())
		__z.SetTypeInfo(4)
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
	return zend_bw_and_helper_SPEC(op1, op2, execute_data)
}
func ZEND_BW_XOR_SPEC_CONST_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if 1<<0 == 1<<0 && 1<<0 == 1<<0 {

	} else if op1.GetTypeInfo() == 4 && op2.GetTypeInfo() == 4 {
		var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		__z.GetValue().SetLval(op1.GetValue().GetLval() ^ op2.GetValue().GetLval())
		__z.SetTypeInfo(4)
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
	return zend_bw_xor_helper_SPEC(op1, op2, execute_data)
}
func ZEND_BOOL_XOR_SPEC_CONST_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	BooleanXorFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op1, op2)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_DIM_R_SPEC_CONST_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var container *Zval
	var dim *Zval
	var value *Zval
	container = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	dim = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if 1<<0 != 1<<0 {
		if container.GetType() == 7 {
		fetch_dim_r_array:
			value = ZendFetchDimensionAddressInner(container.GetValue().GetArr(), dim, 1<<0, 0, execute_data)
			var _z3 *Zval = value
			if (_z3.GetTypeInfo() & 0xff00) != 0 {
				if (_z3.GetTypeInfo() & 0xff) == 10 {
					_z3 = &(*_z3).value.GetRef().GetVal()
					if (_z3.GetTypeInfo() & 0xff00) != 0 {
						ZvalAddrefP(_z3)
					}
				} else {
					ZvalAddrefP(_z3)
				}
			}
			var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var _z2 *Zval = _z3
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
		} else if container.GetType() == 10 {
			container = &(*container).value.GetRef().GetVal()
			if container.GetType() == 7 {
				goto fetch_dim_r_array
			} else {
				goto fetch_dim_r_slow
			}
		} else {
		fetch_dim_r_slow:
			if 1<<0 == 1<<0 && dim.GetU2Extra() == 1 {
				dim++
			}
			zend_fetch_dimension_address_read_R_slow(container, dim, opline, execute_data)
		}
	} else {
		zend_fetch_dimension_address_read_R(container, dim, 1<<0, opline, execute_data)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_DIM_IS_SPEC_CONST_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var container *Zval
	container = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	zend_fetch_dimension_address_read_IS(container, (*Zval)((*byte)(opline)+int32(opline.GetOp2()).constant), 1<<0, opline, execute_data)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_DIM_FUNC_ARG_SPEC_CONST_CONST_HANDLER(execute_data *ZendExecuteData) int {
	if (execute_data.GetCall().GetThis().GetTypeInfo() & 1 << 31) != 0 {
		if (1 << 0 & (1<<0 | 1<<1)) != 0 {
			return zend_use_tmp_in_write_context_helper_SPEC(execute_data)
		}
		return ZEND_NULL_HANDLER(execute_data)
	} else {
		if 1<<0 == 0 {
			return zend_use_undef_in_read_context_helper_SPEC(execute_data)
		}
		return ZEND_FETCH_DIM_R_SPEC_CONST_CONST_HANDLER(execute_data)
	}
}
func ZEND_FETCH_OBJ_R_SPEC_CONST_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var container *Zval
	var offset *Zval
	var cache_slot *any = nil
	container = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	if 1<<0 == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	offset = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if 1<<0 == 1<<0 || 1<<0 != 0 && container.GetType() != 8 {
		for {
			if (1<<0&(1<<2|1<<3)) != 0 && container.GetType() == 10 {
				container = &(*container).value.GetRef().GetVal()
				if container.GetType() == 8 {
					break
				}
			}
			if 1<<0 == 1<<3 && container.GetType() == 0 {
				_zvalUndefinedOp1(execute_data)
			}
			if 1<<0 == 1<<3 && offset.GetType() == 0 {
				_zvalUndefinedOp2(execute_data)
			}
			ZendWrongPropertyRead(offset)
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
			goto fetch_obj_r_finish
			break
		}
	}

	/* here we are sure we are dealing with an object */

	var zobj *ZendObject = container.GetValue().GetObj()
	var retval *Zval
	if 1<<0 == 1<<0 {
		cache_slot = (*any)((*byte)(execute_data.GetRunTimeCache() + (opline.GetExtendedValue() & ^1)))
		if zobj.GetCe() == cache_slot[0] {
			var prop_offset uintPtr = uintptr_t(cache_slot + 1)[0]
			if intptr_t(prop_offset) > 0 {
				retval = (*Zval)((*byte)(zobj + prop_offset))
				if retval.GetTypeInfo() != 0 {
					if (1 << 0 & (1<<1 | 1<<2)) != 0 {
						goto fetch_obj_r_copy
					} else {
					fetch_obj_r_fast_copy:
						var _z3 *Zval = retval
						if (_z3.GetTypeInfo() & 0xff00) != 0 {
							if (_z3.GetTypeInfo() & 0xff) == 10 {
								_z3 = &(*_z3).value.GetRef().GetVal()
								if (_z3.GetTypeInfo() & 0xff00) != 0 {
									ZvalAddrefP(_z3)
								}
							} else {
								ZvalAddrefP(_z3)
							}
						}
						var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
						var _z2 *Zval = _z3
						var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
						var _t uint32 = _z2.GetTypeInfo()
						_z1.GetValue().SetCounted(_gc)
						_z1.SetTypeInfo(_t)
						assert(EG.GetException() == nil)
						execute_data.SetOpline(opline + 1)
						return 0
					}
				}
			} else if zobj.GetProperties() != nil {
				if prop_offset != uintptr_t(intptr_t)(-1) {
					var idx uintPtr = uintptr_t(-(intptr_t(prop_offset)) - 2)
					if idx < zobj.GetProperties().GetNNumUsed()*g.SizeOf("Bucket") {
						var p *Bucket = (*Bucket)((*byte)(zobj.GetProperties().GetArData() + idx))
						if p.GetVal().GetType() != 0 && (p.GetKey() == offset.GetValue().GetStr() || p.GetH() == offset.GetValue().GetStr().GetH() && p.GetKey() != nil && ZendStringEqualContent(p.GetKey(), offset.GetValue().GetStr()) != 0) {
							retval = &p.val
							if (1 << 0 & (1<<1 | 1<<2)) != 0 {
								goto fetch_obj_r_copy
							} else {
								goto fetch_obj_r_fast_copy
							}
						}
					}
					(cache_slot + 1)[0] = any(uintptr_t(intptr_t)(-1))
				}
				retval = ZendHashFindEx(zobj.GetProperties(), offset.GetValue().GetStr(), 1)
				if retval != nil {
					var idx uintPtr = (*byte)(retval - (*byte)(zobj.GetProperties().GetArData()))
					(cache_slot + 1)[0] = any(uintptr_t(-(intptr_t(idx) + 2)))
					if (1 << 0 & (1<<1 | 1<<2)) != 0 {
						goto fetch_obj_r_copy
					} else {
						goto fetch_obj_r_fast_copy
					}
				}
			}
		}
	} else if 1<<0 == 1<<3 && offset.GetTypeInfo() == 0 {
		_zvalUndefinedOp2(execute_data)
	}
	retval = zobj.GetHandlers().GetReadProperty()(container, offset, 0, cache_slot, (*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())))
	if retval != (*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())) {
	fetch_obj_r_copy:
		var _z3 *Zval = retval
		if (_z3.GetTypeInfo() & 0xff00) != 0 {
			if (_z3.GetTypeInfo() & 0xff) == 10 {
				_z3 = &(*_z3).value.GetRef().GetVal()
				if (_z3.GetTypeInfo() & 0xff00) != 0 {
					ZvalAddrefP(_z3)
				}
			} else {
				ZvalAddrefP(_z3)
			}
		}
		var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var _z2 *Zval = _z3
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
	} else if retval.GetType() == 10 {
		ZendUnwrapReference(retval)
	}
fetch_obj_r_finish:
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_OBJ_IS_SPEC_CONST_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var container *Zval
	var offset *Zval
	var cache_slot *any = nil
	container = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	if 1<<0 == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	offset = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if 1<<0 == 1<<0 || 1<<0 != 0 && container.GetType() != 8 {
		for {
			if (1<<0&(1<<2|1<<3)) != 0 && container.GetType() == 10 {
				container = &(*container).value.GetRef().GetVal()
				if container.GetType() == 8 {
					break
				}
			}
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
			goto fetch_obj_is_finish
			break
		}
	}

	/* here we are sure we are dealing with an object */

	var zobj *ZendObject = container.GetValue().GetObj()
	var retval *Zval
	if 1<<0 == 1<<0 {
		cache_slot = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))
		if zobj.GetCe() == cache_slot[0] {
			var prop_offset uintPtr = uintptr_t(cache_slot + 1)[0]
			if intptr_t(prop_offset) > 0 {
				retval = (*Zval)((*byte)(zobj + prop_offset))
				if retval.GetType() != 0 {
					if (1 << 0 & (1<<1 | 1<<2)) != 0 {
						goto fetch_obj_is_copy
					} else {
					fetch_obj_is_fast_copy:
						var _z3 *Zval = retval
						if (_z3.GetTypeInfo() & 0xff00) != 0 {
							if (_z3.GetTypeInfo() & 0xff) == 10 {
								_z3 = &(*_z3).value.GetRef().GetVal()
								if (_z3.GetTypeInfo() & 0xff00) != 0 {
									ZvalAddrefP(_z3)
								}
							} else {
								ZvalAddrefP(_z3)
							}
						}
						var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
						var _z2 *Zval = _z3
						var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
						var _t uint32 = _z2.GetTypeInfo()
						_z1.GetValue().SetCounted(_gc)
						_z1.SetTypeInfo(_t)
						assert(EG.GetException() == nil)
						execute_data.SetOpline(opline + 1)
						return 0
					}
				}
			} else if zobj.GetProperties() != nil {
				if prop_offset != uintptr_t(intptr_t)(-1) {
					var idx uintPtr = uintptr_t(-(intptr_t(prop_offset)) - 2)
					if idx < zobj.GetProperties().GetNNumUsed()*g.SizeOf("Bucket") {
						var p *Bucket = (*Bucket)((*byte)(zobj.GetProperties().GetArData() + idx))
						if p.GetVal().GetType() != 0 && (p.GetKey() == offset.GetValue().GetStr() || p.GetH() == offset.GetValue().GetStr().GetH() && p.GetKey() != nil && ZendStringEqualContent(p.GetKey(), offset.GetValue().GetStr()) != 0) {
							retval = &p.val
							if (1 << 0 & (1<<1 | 1<<2)) != 0 {
								goto fetch_obj_is_copy
							} else {
								goto fetch_obj_is_fast_copy
							}
						}
					}
					(cache_slot + 1)[0] = any(uintptr_t(intptr_t)(-1))
				}
				retval = ZendHashFindEx(zobj.GetProperties(), offset.GetValue().GetStr(), 1)
				if retval != nil {
					var idx uintPtr = (*byte)(retval - (*byte)(zobj.GetProperties().GetArData()))
					(cache_slot + 1)[0] = any(uintptr_t(-(intptr_t(idx) + 2)))
					if (1 << 0 & (1<<1 | 1<<2)) != 0 {
						goto fetch_obj_is_copy
					} else {
						goto fetch_obj_is_fast_copy
					}
				}
			}
		}
	}
	retval = zobj.GetHandlers().GetReadProperty()(container, offset, 3, cache_slot, (*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())))
	if retval != (*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())) {
	fetch_obj_is_copy:
		var _z3 *Zval = retval
		if (_z3.GetTypeInfo() & 0xff00) != 0 {
			if (_z3.GetTypeInfo() & 0xff) == 10 {
				_z3 = &(*_z3).value.GetRef().GetVal()
				if (_z3.GetTypeInfo() & 0xff00) != 0 {
					ZvalAddrefP(_z3)
				}
			} else {
				ZvalAddrefP(_z3)
			}
		}
		var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var _z2 *Zval = _z3
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
	} else if retval.GetType() == 10 {
		ZendUnwrapReference(retval)
	}
fetch_obj_is_finish:
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_OBJ_FUNC_ARG_SPEC_CONST_CONST_HANDLER(execute_data *ZendExecuteData) int {
	if (execute_data.GetCall().GetThis().GetTypeInfo() & 1 << 31) != 0 {

		/* Behave like FETCH_OBJ_W */

		if (1 << 0 & (1<<0 | 1<<1)) != 0 {
			return zend_use_tmp_in_write_context_helper_SPEC(execute_data)
		}
		return ZEND_NULL_HANDLER(execute_data)
	} else {
		return ZEND_FETCH_OBJ_R_SPEC_CONST_CONST_HANDLER(execute_data)
	}
}
func ZEND_FETCH_LIST_R_SPEC_CONST_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var container *Zval
	container = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	zend_fetch_dimension_address_LIST_r(container, (*Zval)((*byte)(opline)+int32(opline.GetOp2()).constant), 1<<0, opline, execute_data)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FAST_CONCAT_SPEC_CONST_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var op1_str *ZendString
	var op2_str *ZendString
	var str *ZendString
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if (1<<0 == 1<<0 || op1.GetType() == 6) && (1<<0 == 1<<0 || op2.GetType() == 6) {
		var op1_str *ZendString = op1.GetValue().GetStr()
		var op2_str *ZendString = op2.GetValue().GetStr()
		var str *ZendString
		if 1<<0 != 1<<0 && op1_str.GetLen() == 0 {
			if 1<<0 == 1<<0 || 1<<0 == 1<<3 {
				var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				var __s *ZendString = op2_str
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					ZendGcAddref(&__s.gc)
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
			} else {
				var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				var __s *ZendString = op2_str
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
			}
			if (1 << 0 & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op1_str, 0)
			}
		} else if 1<<0 != 1<<0 && op2_str.GetLen() == 0 {
			if 1<<0 == 1<<0 || 1<<0 == 1<<3 {
				var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				var __s *ZendString = op1_str
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					ZendGcAddref(&__s.gc)
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
			} else {
				var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				var __s *ZendString = op1_str
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
			}
			if (1 << 0 & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op2_str, 0)
			}
		} else if 1<<0 != 1<<0 && 1<<0 != 1<<3 && (ZvalGcFlags(op1_str.GetGc().GetTypeInfo())&1<<6) == 0 && ZendGcRefcount(&op1_str.gc) == 1 {
			var len_ int = op1_str.GetLen()
			str = ZendStringExtend(op1_str, len_+op2_str.GetLen(), 0)
			memcpy(str.GetVal()+len_, op2_str.GetVal(), op2_str.GetLen()+1)
			var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var __s *ZendString = str
			__z.GetValue().SetStr(__s)
			__z.SetTypeInfo(6 | 1<<0<<8)
			if (1 << 0 & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op2_str, 0)
			}
		} else {
			str = ZendStringAlloc(op1_str.GetLen()+op2_str.GetLen(), 0)
			memcpy(str.GetVal(), op1_str.GetVal(), op1_str.GetLen())
			memcpy(str.GetVal()+op1_str.GetLen(), op2_str.GetVal(), op2_str.GetLen()+1)
			var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var __s *ZendString = str
			__z.GetValue().SetStr(__s)
			__z.SetTypeInfo(6 | 1<<0<<8)
			if (1 << 0 & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op1_str, 0)
			}
			if (1 << 0 & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op2_str, 0)
			}
		}
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
	if 1<<0 == 1<<0 {
		op1_str = op1.GetValue().GetStr()
	} else if op1.GetType() == 6 {
		op1_str = ZendStringCopy(op1.GetValue().GetStr())
	} else {
		if 1<<0 == 1<<3 && op1.GetType() == 0 {
			_zvalUndefinedOp1(execute_data)
		}
		op1_str = ZvalGetStringFunc(op1)
	}
	if 1<<0 == 1<<0 {
		op2_str = op2.GetValue().GetStr()
	} else if op2.GetType() == 6 {
		op2_str = ZendStringCopy(op2.GetValue().GetStr())
	} else {
		if 1<<0 == 1<<3 && op2.GetType() == 0 {
			_zvalUndefinedOp2(execute_data)
		}
		op2_str = ZvalGetStringFunc(op2)
	}
	for {
		if 1<<0 != 1<<0 {
			if op1_str.GetLen() == 0 {
				if 1<<0 == 1<<0 {
					if op2.GetTypeFlags() != 0 {
						ZendGcAddref(&op2_str.gc)
					}
				}
				var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				var __s *ZendString = op2_str
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
				ZendStringReleaseEx(op1_str, 0)
				break
			}
		}
		if 1<<0 != 1<<0 {
			if op2_str.GetLen() == 0 {
				if 1<<0 == 1<<0 {
					if op1.GetTypeFlags() != 0 {
						ZendGcAddref(&op1_str.gc)
					}
				}
				var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				var __s *ZendString = op1_str
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
				ZendStringReleaseEx(op2_str, 0)
				break
			}
		}
		str = ZendStringAlloc(op1_str.GetLen()+op2_str.GetLen(), 0)
		memcpy(str.GetVal(), op1_str.GetVal(), op1_str.GetLen())
		memcpy(str.GetVal()+op1_str.GetLen(), op2_str.GetVal(), op2_str.GetLen()+1)
		var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var __s *ZendString = str
		__z.GetValue().SetStr(__s)
		__z.SetTypeInfo(6 | 1<<0<<8)
		if 1<<0 != 1<<0 {
			ZendStringReleaseEx(op1_str, 0)
		}
		if 1<<0 != 1<<0 {
			ZendStringReleaseEx(op2_str, 0)
		}
		break
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_INIT_METHOD_CALL_SPEC_CONST_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var function_name *Zval
	var free_op1 ZendFreeOp
	var object *Zval
	var fbc *ZendFunction
	var called_scope *ZendClassEntry
	var obj *ZendObject
	var call *ZendExecuteData
	var call_info uint32
	object = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	if 1<<0 == 0 && object.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	if 1<<0 != 1<<0 {
		function_name = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	}
	if 1<<0 != 1<<0 && function_name.GetType() != 6 {
		for {
			if (1<<0&(1<<2|1<<3)) != 0 && function_name.GetType() == 10 {
				function_name = &(*function_name).value.GetRef().GetVal()
				if function_name.GetType() == 6 {
					break
				}
			} else if 1<<0 == 1<<3 && function_name.GetType() == 0 {
				_zvalUndefinedOp2(execute_data)
				if EG.GetException() != nil {
					return 0
				}
			}
			ZendThrowError(nil, "Method name must be a string")
			return 0
			break
		}
	}
	if 1<<0 != 0 {
		for {
			if 1<<0 == 1<<0 || object.GetType() != 8 {
				if (1<<0&(1<<2|1<<3)) != 0 && object.GetType() == 10 {
					object = &(*object).value.GetRef().GetVal()
					if object.GetType() == 8 {
						break
					}
				}
				if 1<<0 == 1<<3 && object.GetType() == 0 {
					object = _zvalUndefinedOp1(execute_data)
					if EG.GetException() != nil {
						if 1<<0 != 1<<0 {

						}
						return 0
					}
				}
				if 1<<0 == 1<<0 {
					function_name = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
				}
				ZendInvalidMethodCall(object, function_name)
				return 0
			}
			break
		}
	}
	obj = object.GetValue().GetObj()
	called_scope = obj.GetCe()
	if 1<<0 == 1<<0 && (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetResult().GetNum()))[0] == called_scope {
		fbc = (*any)((*byte)(execute_data.GetRunTimeCache() + (opline.GetResult().GetNum() + g.SizeOf("void *"))))[0]
	} else {
		var orig_obj *ZendObject = obj
		if 1<<0 == 1<<0 {
			function_name = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
		}

		/* First, locate the function. */

		fbc = obj.GetHandlers().GetGetMethod()(&obj, function_name.GetValue().GetStr(), g.CondF1(1<<0 == 1<<0, func() *Zval { return (*Zval)((*byte)(opline)+int32(opline.GetOp2()).constant) + 1 }, nil))
		if fbc == nil {
			if EG.GetException() == nil {
				ZendUndefinedMethod(obj.GetCe(), function_name.GetValue().GetStr())
			}
			return 0
		}
		if 1<<0 == 1<<0 && fbc.GetType() <= 2 && (fbc.GetFnFlags()&(1<<18|1<<19)) == 0 && obj == orig_obj {
			var slot *any = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetResult().GetNum()))
			slot[0] = called_scope
			slot[1] = fbc
		}
		if (1<<0&(1<<2|1<<1)) != 0 && obj != orig_obj {

			/* Reset "object" to trigger reference counting */

			object = nil

			/* Reset "object" to trigger reference counting */

		}
		if fbc.GetType() == 2 && !(g.CondF((uintptr_t(&fbc.op_array).run_time_cache__ptr&1) != 0, func() any {
			return *((*any)((*byte)(CG.GetMapPtrBase() + uintptr_t(&fbc.op_array).run_time_cache__ptr - 1)))
		}, func() any { return any(*(&fbc.op_array.run_time_cache__ptr)) })) {
			InitFuncRunTimeCache(&fbc.op_array)
		}
	}
	if 1<<0 != 1<<0 {

	}
	call_info = 0<<16 | 0<<17 | (8 | 1<<0<<8 | 1<<1<<8)
	if (fbc.GetFnFlags() & 1 << 4) != 0 {
		if (1<<0&(1<<2|1<<1)) != 0 && EG.GetException() != nil {
			return 0
		}

		/* call static method */

		obj = (*ZendObject)(called_scope)
		call_info = 0<<16 | 0<<17
	} else if (1 << 0 & (1<<2 | 1<<1 | 1<<3)) != 0 {
		if 1<<0 == 1<<3 {
			ZendGcAddref(&obj.gc)
		} else if free_op1 != object {
			ZendGcAddref(&obj.gc)
		}

		/* CV may be changed indirectly (e.g. when it's a reference) */

		call_info = 0<<16 | 0<<17 | (8 | 1<<0<<8 | 1<<1<<8) | 1<<21

		/* CV may be changed indirectly (e.g. when it's a reference) */

	}
	call = ZendVmStackPushCallFrame(call_info, fbc, opline.GetExtendedValue(), obj)
	call.SetPrevExecuteData(execute_data.GetCall())
	execute_data.SetCall(call)
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_INIT_STATIC_METHOD_CALL_SPEC_CONST_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var function_name *Zval
	var ce *ZendClassEntry
	var call_info uint32
	var fbc *ZendFunction
	var call *ZendExecuteData
	if 1<<0 == 1<<0 {

		/* no function found. try a static method in class */

		ce = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetResult().GetNum()))[0]
		if ce == nil {
			ce = ZendFetchClassByName((*Zval)((*byte)(opline)+int32(opline.GetOp1()).constant).GetValue().GetStr(), ((*Zval)((*byte)(opline)+int32(opline.GetOp1()).constant) + 1).GetValue().GetStr(), 0|0x200)
			if ce == nil {
				assert(EG.GetException() != nil)
				return 0
			}
			if 1<<0 != 1<<0 {
				(*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetResult().GetNum()))[0] = ce
			}
		}
	} else if 1<<0 == 0 {
		ce = ZendFetchClass(nil, opline.GetOp1().GetNum())
		if ce == nil {
			assert(EG.GetException() != nil)
			return 0
		}
	} else {
		ce = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar())).GetValue().GetCe()
	}
	if 1<<0 == 1<<0 && 1<<0 == 1<<0 && g.Assign(&fbc, (*any)((*byte)(execute_data.GetRunTimeCache() + (opline.GetResult().GetNum() + g.SizeOf("void *"))))[0]) != nil {

	} else if 1<<0 != 1<<0 && 1<<0 == 1<<0 && (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetResult().GetNum()))[0] == ce {
		fbc = (*any)((*byte)(execute_data.GetRunTimeCache() + (opline.GetResult().GetNum() + g.SizeOf("void *"))))[0]
	} else if 1<<0 != 0 {
		function_name = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
		if 1<<0 != 1<<0 {
			if function_name.GetType() != 6 {
				for {
					if (1<<0&(1<<2|1<<3)) != 0 && function_name.GetType() == 10 {
						function_name = &(*function_name).value.GetRef().GetVal()
						if function_name.GetType() == 6 {
							break
						}
					} else if 1<<0 == 1<<3 && function_name.GetType() == 0 {
						_zvalUndefinedOp2(execute_data)
						if EG.GetException() != nil {
							return 0
						}
					}
					ZendThrowError(nil, "Function name must be a string")
					return 0
					break
				}
			}
		}
		if ce.GetGetStaticMethod() != nil {
			fbc = ce.GetGetStaticMethod()(ce, function_name.GetValue().GetStr())
		} else {
			fbc = ZendStdGetStaticMethod(ce, function_name.GetValue().GetStr(), g.CondF1(1<<0 == 1<<0, func() *Zval { return (*Zval)((*byte)(opline)+int32(opline.GetOp2()).constant) + 1 }, nil))
		}
		if fbc == nil {
			if EG.GetException() == nil {
				ZendUndefinedMethod(ce, function_name.GetValue().GetStr())
			}
			return 0
		}
		if 1<<0 == 1<<0 && fbc.GetType() <= 2 && (fbc.GetFnFlags()&(1<<18|1<<19)) == 0 {
			var slot *any = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetResult().GetNum()))
			slot[0] = ce
			slot[1] = fbc
		}
		if fbc.GetType() == 2 && !(g.CondF((uintptr_t(&fbc.op_array).run_time_cache__ptr&1) != 0, func() any {
			return *((*any)((*byte)(CG.GetMapPtrBase() + uintptr_t(&fbc.op_array).run_time_cache__ptr - 1)))
		}, func() any { return any(*(&fbc.op_array.run_time_cache__ptr)) })) {
			InitFuncRunTimeCache(&fbc.op_array)
		}
		if 1<<0 != 1<<0 {

		}
	} else {
		if ce.GetConstructor() == nil {
			ZendThrowError(nil, "Cannot call constructor")
			return 0
		}
		if execute_data.GetThis().GetType() == 8 && execute_data.GetThis().GetValue().GetObj().GetCe() != ce.GetConstructor().GetScope() && (ce.GetConstructor().GetFnFlags()&1<<2) != 0 {
			ZendThrowError(nil, "Cannot call private %s::__construct()", ce.GetName().GetVal())
			return 0
		}
		fbc = ce.GetConstructor()
		if fbc.GetType() == 2 && !(g.CondF((uintptr_t(&fbc.op_array).run_time_cache__ptr&1) != 0, func() any {
			return *((*any)((*byte)(CG.GetMapPtrBase() + uintptr_t(&fbc.op_array).run_time_cache__ptr - 1)))
		}, func() any { return any(*(&fbc.op_array.run_time_cache__ptr)) })) {
			InitFuncRunTimeCache(&fbc.op_array)
		}
	}
	if (fbc.GetFnFlags() & 1 << 4) == 0 {
		if execute_data.GetThis().GetType() == 8 && InstanceofFunction(execute_data.GetThis().GetValue().GetObj().GetCe(), ce) != 0 {
			ce = (*ZendClassEntry)(execute_data.GetThis().GetValue().GetObj())
			call_info = 0<<16 | 0<<17 | (8 | 1<<0<<8 | 1<<1<<8)
		} else {
			ZendNonStaticMethodCall(fbc)
			if EG.GetException() != nil {
				return 0
			}
			goto check_parent_and_self
		}
	} else {
	check_parent_and_self:

		/* previous opcode is ZEND_FETCH_CLASS */

		if 1<<0 == 0 && ((opline.GetOp1().GetNum()&0xf) == 2 || (opline.GetOp1().GetNum()&0xf) == 1) {
			if execute_data.GetThis().GetType() == 8 {
				ce = execute_data.GetThis().GetValue().GetObj().GetCe()
			} else {
				ce = execute_data.GetThis().GetValue().GetCe()
			}
		}
		call_info = 0<<16 | 0<<17
	}
	call = ZendVmStackPushCallFrame(call_info, fbc, opline.GetExtendedValue(), ce)
	call.SetPrevExecuteData(execute_data.GetCall())
	execute_data.SetCall(call)
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_INIT_USER_CALL_SPEC_CONST_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var function_name *Zval
	var fcc ZendFcallInfoCache
	var error *byte = nil
	var func_ *ZendFunction
	var object_or_called_scope any
	var call *ZendExecuteData
	var call_info uint32 = 0<<16 | 0<<17 | 1<<25
	function_name = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if ZendIsCallableEx(function_name, nil, 0, nil, &fcc, &error) != 0 {
		func_ = fcc.GetFunctionHandler()
		if error != nil {
			_efree(error)

			/* This is the only soft error is_callable() can generate */

			ZendNonStaticMethodCall(func_)
			if EG.GetException() != nil {
				return 0
			}
		}
		object_or_called_scope = fcc.GetCalledScope()
		if (func_.GetFnFlags() & 1 << 20) != 0 {

			/* Delay closure destruction until its invocation */

			ZendGcAddref(&((*ZendObject)((*byte)(func_ - g.SizeOf("zend_object")))).gc)
			call_info |= 1 << 22
			if (func_.GetFnFlags() & 1 << 21) != 0 {
				call_info |= 1 << 23
			}
			if fcc.GetObject() != nil {
				object_or_called_scope = fcc.GetObject()
				call_info |= 8 | 1<<0<<8 | 1<<1<<8
			}
		} else if fcc.GetObject() != nil {
			ZendGcAddref(&(fcc.GetObject()).gc)
			object_or_called_scope = fcc.GetObject()
			call_info |= 1<<21 | (8 | 1<<0<<8 | 1<<1<<8)
		}
		if (1<<0&(1<<1|1<<2)) != 0 && EG.GetException() != nil {
			if (call_info & 1 << 22) != 0 {
				ZendObjectRelease((*ZendObject)((*byte)(func_ - g.SizeOf("zend_object"))))
			} else if (call_info & 1 << 21) != 0 {
				ZendObjectRelease(fcc.GetObject())
			}
			return 0
		}
		if func_.GetType() == 2 && !(g.CondF((uintptr_t(&func_.op_array).run_time_cache__ptr&1) != 0, func() any {
			return *((*any)((*byte)(CG.GetMapPtrBase() + uintptr_t(&func_.op_array).run_time_cache__ptr - 1)))
		}, func() any { return any(*(&func_.op_array.run_time_cache__ptr)) })) {
			InitFuncRunTimeCache(&func_.op_array)
		}
	} else {
		ZendInternalTypeError((execute_data.GetFunc().GetFnFlags()&1<<31) != 0, "%s() expects parameter 1 to be a valid callback, %s", (*Zval)((*byte)(opline)+int32(opline.GetOp1()).constant).GetValue().GetStr().GetVal(), error)
		_efree(error)
		if EG.GetException() != nil {
			return 0
		}
		func_ = (*ZendFunction)(&ZendPassFunction)
		object_or_called_scope = nil
	}
	call = ZendVmStackPushCallFrame(call_info, func_, opline.GetExtendedValue(), object_or_called_scope)
	call.SetPrevExecuteData(execute_data.GetCall())
	execute_data.SetCall(call)
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_FETCH_CLASS_CONSTANT_SPEC_CONST_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var ce *ZendClassEntry
	var scope *ZendClassEntry
	var c *ZendClassConstant
	var value *Zval
	var zv *Zval
	var opline *ZendOp = execute_data.GetOpline()
	for {
		if 1<<0 == 1<<0 {
			if (*any)((*byte)(execute_data.GetRunTimeCache() + (opline.GetExtendedValue() + g.SizeOf("void *"))))[0] {
				value = (*any)((*byte)(execute_data.GetRunTimeCache() + (opline.GetExtendedValue() + g.SizeOf("void *"))))[0]
				break
			} else if (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))[0] {
				ce = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))[0]
			} else {
				ce = ZendFetchClassByName((*Zval)((*byte)(opline)+int32(opline.GetOp1()).constant).GetValue().GetStr(), ((*Zval)((*byte)(opline)+int32(opline.GetOp1()).constant) + 1).GetValue().GetStr(), 0|0x200)
				if ce == nil {
					assert(EG.GetException() != nil)
					(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
					return 0
				}
			}
		} else {
			if 1<<0 == 0 {
				ce = ZendFetchClass(nil, opline.GetOp1().GetNum())
				if ce == nil {
					assert(EG.GetException() != nil)
					(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
					return 0
				}
			} else {
				ce = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar())).GetValue().GetCe()
			}
			if (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))[0] == ce {
				value = (*any)((*byte)(execute_data.GetRunTimeCache() + (opline.GetExtendedValue() + g.SizeOf("void *"))))[0]
				break
			}
		}
		zv = ZendHashFindEx(&ce.constants_table, (*Zval)((*byte)(opline)+int32(opline.GetOp2()).constant).GetValue().GetStr(), 1)
		if zv != nil {
			c = zv.GetValue().GetPtr()
			scope = execute_data.GetFunc().GetOpArray().GetScope()
			if ZendVerifyConstAccess(c, scope) == 0 {
				ZendThrowError(nil, "Cannot access %s const %s::%s", ZendVisibilityString(c.GetValue().GetAccessFlags()), ce.GetName().GetVal(), (*Zval)((*byte)(opline)+int32(opline.GetOp2()).constant).GetValue().GetStr().GetVal())
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
				return 0
			}
			value = &c.value
			if value.GetType() == 11 {
				ZvalUpdateConstantEx(value, c.GetCe())
				if EG.GetException() != nil {
					(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
					return 0
				}
			}
			var slot *any = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))
			slot[0] = ce
			slot[1] = value
		} else {
			ZendThrowError(nil, "Undefined class constant '%s'", (*Zval)((*byte)(opline)+int32(opline.GetOp2()).constant).GetValue().GetStr().GetVal())
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
			return 0
		}
		break
	}
	var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	var _z2 *Zval = value
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	if (_t & 0xff00) != 0 {
		if (ZvalGcFlags(_gc.GetGc().GetTypeInfo()) & 1 << 7) == 0 {
			ZendGcAddref(&_gc.gc)
		} else {
			ZvalCopyCtorFunc(_z1)
		}
	}
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_ADD_ARRAY_ELEMENT_SPEC_CONST_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var expr_ptr *Zval
	var new_expr Zval
	if (1<<0 == 1<<2 || 1<<0 == 1<<3) && (opline.GetExtendedValue()&1<<0) != 0 {
		expr_ptr = nil
		if expr_ptr.GetType() == 10 {
			ZvalAddrefP(expr_ptr)
		} else {
			var _z *Zval = expr_ptr
			var _ref *ZendReference = (*ZendReference)(_emalloc(g.SizeOf("zend_reference")))
			ZendGcSetRefcount(&_ref.gc, 2)
			_ref.GetGc().SetTypeInfo(10)
			var _z1 *Zval = &_ref.val
			var _z2 *Zval = _z
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			_ref.GetSources().SetPtr(nil)
			_z.GetValue().SetRef(_ref)
			_z.SetTypeInfo(10 | 1<<0<<8)
		}
	} else {
		expr_ptr = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
		if 1<<0 == 1<<1 {

		} else if 1<<0 == 1<<0 {
			if expr_ptr.GetTypeFlags() != 0 {
				ZvalAddrefP(expr_ptr)
			}
		} else if 1<<0 == 1<<3 {
			if expr_ptr.GetType() == 10 {
				expr_ptr = &(*expr_ptr).value.GetRef().GetVal()
			}
			if expr_ptr.GetTypeFlags() != 0 {
				ZvalAddrefP(expr_ptr)
			}
		} else {
			if expr_ptr.GetType() == 10 {
				var ref *ZendRefcounted = expr_ptr.GetValue().GetCounted()
				expr_ptr = &(*expr_ptr).value.GetRef().GetVal()
				if ZendGcDelref(&ref.gc) == 0 {
					var _z1 *Zval = &new_expr
					var _z2 *Zval = expr_ptr
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					expr_ptr = &new_expr
					_efree(ref)
				} else if (expr_ptr.GetTypeInfo() & 0xff00) != 0 {
					ZvalAddrefP(expr_ptr)
				}
			}
		}
	}
	if 1<<0 != 0 {
		var offset *Zval = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
		var str *ZendString
		var hval ZendUlong
	add_again:
		if offset.GetType() == 6 {
			str = offset.GetValue().GetStr()
			if 1<<0 != 1<<0 {
				if _zendHandleNumericStr(str.GetVal(), str.GetLen(), &hval) != 0 {
					goto num_index
				}
			}
		str_index:
			ZendHashUpdate((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())).GetValue().GetArr(), str, expr_ptr)
		} else if offset.GetType() == 4 {
			hval = offset.GetValue().GetLval()
		num_index:
			ZendHashIndexUpdate((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())).GetValue().GetArr(), hval, expr_ptr)
		} else if (1<<0&(1<<2|1<<3)) != 0 && offset.GetType() == 10 {
			offset = &(*offset).value.GetRef().GetVal()
			goto add_again
		} else if offset.GetType() == 1 {
			str = ZendEmptyString
			goto str_index
		} else if offset.GetType() == 5 {
			hval = ZendDvalToLval(offset.GetValue().GetDval())
			goto num_index
		} else if offset.GetType() == 2 {
			hval = 0
			goto num_index
		} else if offset.GetType() == 3 {
			hval = 1
			goto num_index
		} else if offset.GetType() == 9 {
			ZendUseResourceAsOffset(offset)
			hval = offset.GetValue().GetRes().GetHandle()
			goto num_index
		} else if 1<<0 == 1<<3 && offset.GetType() == 0 {
			_zvalUndefinedOp2(execute_data)
			str = ZendEmptyString
			goto str_index
		} else {
			ZendIllegalOffset()
			ZvalPtrDtorNogc(expr_ptr)
		}
	} else {
		if ZendHashNextIndexInsert((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())).GetValue().GetArr(), expr_ptr) == nil {
			ZendCannotAddElement()
			ZvalPtrDtorNogc(expr_ptr)
		}
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_INIT_ARRAY_SPEC_CONST_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var array *Zval
	var size uint32
	var opline *ZendOp = execute_data.GetOpline()
	array = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	if 1<<0 != 0 {
		size = opline.GetExtendedValue() >> 2
		var __arr *ZendArray = _zendNewArray(size)
		var __z *Zval = array
		__z.GetValue().SetArr(__arr)
		__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)

		/* Explicitly initialize array as not-packed if flag is set */

		if (opline.GetExtendedValue() & 1 << 1) != 0 {
			ZendHashRealInitMixed(array.GetValue().GetArr())
		}
		return ZEND_ADD_ARRAY_ELEMENT_SPEC_CONST_CONST_HANDLER(execute_data)
	} else {
		var __arr *ZendArray = _zendNewArray(0)
		var __z *Zval = array
		__z.GetValue().SetArr(__arr)
		__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
}
func ZEND_ISSET_ISEMPTY_DIM_OBJ_SPEC_CONST_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var container *Zval
	var result int
	var hval ZendUlong
	var offset *Zval
	container = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	offset = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if container.GetType() == 7 {
		var ht *HashTable
		var value *Zval
		var str *ZendString
	isset_dim_obj_array:
		ht = container.GetValue().GetArr()
	isset_again:
		if offset.GetType() == 6 {
			str = offset.GetValue().GetStr()
			if 1<<0 != 1<<0 {
				if _zendHandleNumericStr(str.GetVal(), str.GetLen(), &hval) != 0 {
					goto num_index_prop
				}
			}
			value = ZendHashFindExInd(ht, str, 1<<0 == 1<<0)
		} else if offset.GetType() == 4 {
			hval = offset.GetValue().GetLval()
		num_index_prop:
			value = ZendHashIndexFind(ht, hval)
		} else if (1<<0&(1<<2|1<<3)) != 0 && offset.GetType() == 10 {
			offset = &(*offset).value.GetRef().GetVal()
			goto isset_again
		} else {
			value = ZendFindArrayDimSlow(ht, offset, execute_data)
			if EG.GetException() != nil {
				result = 0
				goto isset_dim_obj_exit
			}
		}
		if (opline.GetExtendedValue() & 1 << 0) == 0 {

			/* > IS_NULL means not IS_UNDEF and not IS_NULL */

			result = value != nil && value.GetType() > 1 && (value.GetType() != 10 || &(*value).value.GetRef().GetVal().u1.v.type_ != 1)
			if (1 << 0 & (1<<0 | 1<<3)) != 0 {

				/* avoid exception check */

				for {

					if (opline + 1).GetOpcode() == 43 {
						if result != 0 {
							execute_data.SetOpline(opline + 2)
						} else {
							execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
							if EG.GetVmInterrupt() != 0 {
								return zend_interrupt_helper_SPEC(execute_data)
							}
						}
					} else if (opline + 1).GetOpcode() == 44 {
						if result == 0 {
							execute_data.SetOpline(opline + 2)
						} else {
							execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
							if EG.GetVmInterrupt() != 0 {
								return zend_interrupt_helper_SPEC(execute_data)
							}
						}
					} else {
						break
					}
					return 0
					break
				}
				if result != 0 {
					(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
				} else {
					(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
				}
				assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			}
		} else {
			result = value == nil || IZendIsTrue(value) == 0
		}
		goto isset_dim_obj_exit
	} else if (1<<0&(1<<2|1<<3)) != 0 && container.GetType() == 10 {
		container = &(*container).value.GetRef().GetVal()
		if container.GetType() == 7 {
			goto isset_dim_obj_array
		}
	}
	if 1<<0 == 1<<0 && offset.GetU2Extra() == 1 {
		offset++
	}
	if (opline.GetExtendedValue() & 1 << 0) == 0 {
		result = ZendIssetDimSlow(container, offset, execute_data)
	} else {
		result = ZendIsemptyDimSlow(container, offset, execute_data)
	}
isset_dim_obj_exit:
	for {
		if EG.GetException() != nil {
			break
		}
		if (opline + 1).GetOpcode() == 43 {
			if result != 0 {
				execute_data.SetOpline(opline + 2)
			} else {
				execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
			}
		} else if (opline + 1).GetOpcode() == 44 {
			if result == 0 {
				execute_data.SetOpline(opline + 2)
			} else {
				execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
			}
		} else {
			break
		}
		return 0
		break
	}
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_CONST_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var container *Zval
	var result int
	var offset *Zval
	container = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	if 1<<0 == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	offset = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if 1<<0 == 1<<0 || 1<<0 != 0 && container.GetType() != 8 {
		if (1<<0&(1<<2|1<<3)) != 0 && container.GetType() == 10 {
			container = &(*container).value.GetRef().GetVal()
			if container.GetType() != 8 {
				result = opline.GetExtendedValue() & 1 << 0
				goto isset_object_finish
			}
		} else {
			result = opline.GetExtendedValue() & 1 << 0
			goto isset_object_finish
		}
	}
	result = opline.GetExtendedValue()&1<<0 ^ container.GetValue().GetObj().GetHandlers().GetHasProperty()(container, offset, opline.GetExtendedValue()&1<<0, g.CondF1(1<<0 == 1<<0, func() *any {
		return (*any)((*byte)(execute_data.GetRunTimeCache() + (opline.GetExtendedValue() & ^(1 << 0))))
	}, nil))
isset_object_finish:
	for {
		if EG.GetException() != nil {
			break
		}
		if (opline + 1).GetOpcode() == 43 {
			if result != 0 {
				execute_data.SetOpline(opline + 2)
			} else {
				execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
			}
		} else if (opline + 1).GetOpcode() == 44 {
			if result == 0 {
				execute_data.SetOpline(opline + 2)
			} else {
				execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
			}
		} else {
			break
		}
		return 0
		break
	}
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_ARRAY_KEY_EXISTS_SPEC_CONST_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var key *Zval
	var subject *Zval
	var ht *HashTable
	var result uint32
	key = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	subject = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if subject.GetType() == 7 {
	array_key_exists_array:
		ht = subject.GetValue().GetArr()
		result = ZendArrayKeyExistsFast(ht, key, opline, execute_data)
	} else {
		if (1<<0&(1<<2|1<<3)) != 0 && subject.GetType() == 10 {
			subject = &(*subject).value.GetRef().GetVal()
			if subject.GetType() == 7 {
				goto array_key_exists_array
			}
		}
		result = ZendArrayKeyExistsSlow(subject, key, opline, execute_data)
	}
	for {
		if EG.GetException() != nil {
			break
		}
		if (opline + 1).GetOpcode() == 43 {
			if result == 3 {
				execute_data.SetOpline(opline + 2)
			} else {
				execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
			}
		} else if (opline + 1).GetOpcode() == 44 {
			if result != 3 {
				execute_data.SetOpline(opline + 2)
			} else {
				execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
			}
		} else {
			break
		}
		return 0
		break
	}
	(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(result)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}

/* No specialization for op_types (CONST|TMPVAR|UNUSED|CV, ANY) */

func ZEND_DECLARE_CLASS_DELAYED_SPEC_CONST_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var lcname *Zval
	var zv *Zval
	var ce *ZendClassEntry
	ce = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))[0]
	if ce == nil {
		lcname = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
		zv = ZendHashFindEx(EG.GetClassTable(), (lcname + 1).GetValue().GetStr(), 1)
		if zv != nil {
			ce = zv.GetValue().GetCe()
			zv = ZendHashSetBucketKey(EG.GetClassTable(), (*Bucket)(zv), lcname.GetValue().GetStr())
			if zv == nil {
				ZendErrorNoreturn(1<<6, "Cannot declare %s %s, because the name is already in use", ZendGetObjectType(ce), ce.GetName().GetVal())
			} else {
				if ZendDoLinkClass(ce, (*Zval)((*byte)(opline)+int32(opline.GetOp2()).constant).GetValue().GetStr()) == FAILURE {

					/* Reload bucket pointer, the hash table may have been reallocated */

					zv = ZendHashFind(EG.GetClassTable(), lcname.GetValue().GetStr())
					ZendHashSetBucketKey(EG.GetClassTable(), (*Bucket)(zv), (lcname + 1).GetValue().GetStr())
					return 0
				}
			}
		}
		(*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))[0] = ce
	}
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_DECLARE_CONST_SPEC_CONST_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var name *Zval
	var val *Zval
	var c ZendConstant
	name = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	val = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	var _z1 *Zval = &c.value
	var _z2 *Zval = val
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	if (_t & 0xff00) != 0 {
		ZendGcAddref(&_gc.gc)
	}
	if (c.GetValue().GetTypeInfo() & 0xff) == 11 {
		if ZvalUpdateConstantEx(&c.value, execute_data.GetFunc().GetOpArray().GetScope()) != SUCCESS {
			ZvalPtrDtorNogc(&c.value)
			return 0
		}
	}

	/* non persistent, case sensitive */

	&c.GetValue().SetConstantFlags(1<<0&0xff | 0x7fffff<<8)
	c.SetName(ZendStringCopy(name.GetValue().GetStr()))
	if ZendRegisterConstant(&c) == FAILURE {

	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_YIELD_SPEC_CONST_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var generator *ZendGenerator = ZendGetRunningGenerator(execute_data)
	if (generator.GetFlags() & ZEND_GENERATOR_FORCED_CLOSE) != 0 {
		return zend_yield_in_closed_generator_helper_SPEC(execute_data)
	}

	/* Destroy the previously yielded value */

	ZvalPtrDtor(&generator.value)

	/* Destroy the previously yielded key */

	ZvalPtrDtor(&generator.key)

	/* Set the new yielded value */

	if 1<<0 != 0 {
		if (execute_data.GetFunc().GetOpArray().GetFnFlags() & 1 << 12) != 0 {

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

			if (1 << 0 & (1<<0 | 1<<1)) != 0 {
				var value *Zval
				ZendError(1<<3, "Only variable references should be yielded by reference")
				value = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
				var _z1 *Zval = &generator.value
				var _z2 *Zval = value
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if 1<<0 == 1<<0 {
					if (generator.GetValue().GetTypeInfo() & 0xff00) != 0 {
						ZvalAddrefP(&(generator.GetValue()))
					}
				}
			} else {
				var value_ptr *Zval = nil

				/* If a function call result is yielded and the function did
				 * not return by reference we throw a notice. */

				for {
					if 1<<0 == 1<<2 {
						assert(value_ptr != &EG.uninitialized_zval)
						if opline.GetExtendedValue() == 1<<0 && value_ptr.GetType() != 10 {
							ZendError(1<<3, "Only variable references should be yielded by reference")
							var _z1 *Zval = &generator.value
							var _z2 *Zval = value_ptr
							var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
							var _t uint32 = _z2.GetTypeInfo()
							_z1.GetValue().SetCounted(_gc)
							_z1.SetTypeInfo(_t)
							if (_t & 0xff00) != 0 {
								ZendGcAddref(&_gc.gc)
							}
							break
						}
					}
					if value_ptr.GetType() == 10 {
						ZvalAddrefP(value_ptr)
					} else {
						var _z *Zval = value_ptr
						var _ref *ZendReference = (*ZendReference)(_emalloc(g.SizeOf("zend_reference")))
						ZendGcSetRefcount(&_ref.gc, 2)
						_ref.GetGc().SetTypeInfo(10)
						var _z1 *Zval = &_ref.val
						var _z2 *Zval = _z
						var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
						var _t uint32 = _z2.GetTypeInfo()
						_z1.GetValue().SetCounted(_gc)
						_z1.SetTypeInfo(_t)
						_ref.GetSources().SetPtr(nil)
						_z.GetValue().SetRef(_ref)
						_z.SetTypeInfo(10 | 1<<0<<8)
					}
					var __z *Zval = &generator.value
					__z.GetValue().SetRef(value_ptr.GetValue().GetRef())
					__z.SetTypeInfo(10 | 1<<0<<8)
					break
				}

				/* If a function call result is yielded and the function did
				 * not return by reference we throw a notice. */

			}

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

		} else {
			var value *Zval = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)

			/* Consts, temporary variables and references need copying */

			if 1<<0 == 1<<0 {
				var _z1 *Zval = &generator.value
				var _z2 *Zval = value
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if (generator.GetValue().GetTypeInfo() & 0xff00) != 0 {
					ZvalAddrefP(&(generator.GetValue()))
				}
			} else if 1<<0 == 1<<1 {
				var _z1 *Zval = &generator.value
				var _z2 *Zval = value
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
			} else if (1<<0&(1<<2|1<<3)) != 0 && value.GetType() == 10 {
				var _z1 *Zval = &generator.value
				var _z2 *Zval = &(*value).value.GetRef().GetVal()
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if (_t & 0xff00) != 0 {
					ZendGcAddref(&_gc.gc)
				}
			} else {
				var _z1 *Zval = &generator.value
				var _z2 *Zval = value
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if 1<<0 == 1<<3 {
					if (value.GetTypeInfo() & 0xff00) != 0 {
						ZvalAddrefP(value)
					}
				}
			}

			/* Consts, temporary variables and references need copying */

		}
	} else {

		/* If no value was specified yield null */

		&generator.value.u1.type_info = 1

		/* If no value was specified yield null */

	}

	/* Set the new yielded key */

	if 1<<0 != 0 {
		var key *Zval = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)

		/* Consts, temporary variables and references need copying */

		if 1<<0 == 1<<0 {
			var _z1 *Zval = &generator.key
			var _z2 *Zval = key
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if (generator.GetKey().GetTypeInfo() & 0xff00) != 0 {
				ZvalAddrefP(&(generator.GetKey()))
			}
		} else if 1<<0 == 1<<1 {
			var _z1 *Zval = &generator.key
			var _z2 *Zval = key
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
		} else if (1<<0&(1<<2|1<<3)) != 0 && key.GetType() == 10 {
			var _z1 *Zval = &generator.key
			var _z2 *Zval = &(*key).value.GetRef().GetVal()
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if (_t & 0xff00) != 0 {
				ZendGcAddref(&_gc.gc)
			}
		} else {
			var _z1 *Zval = &generator.key
			var _z2 *Zval = key
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if 1<<0 == 1<<3 {
				if (key.GetTypeInfo() & 0xff00) != 0 {
					ZvalAddrefP(key)
				}
			}
		}
		if generator.GetKey().GetType() == 4 && generator.GetKey().GetValue().GetLval() > generator.GetLargestUsedIntegerKey() {
			generator.SetLargestUsedIntegerKey(generator.GetKey().GetValue().GetLval())
		}
	} else {

		/* If no key was specified we use auto-increment keys */

		generator.GetLargestUsedIntegerKey()++
		var __z *Zval = &generator.key
		__z.GetValue().SetLval(generator.GetLargestUsedIntegerKey())
		__z.SetTypeInfo(4)
	}
	if opline.GetResultType() != 0 {

		/* If the return value of yield is used set the send
		 * target and initialize it to NULL */

		generator.SetSendTarget((*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())))
		generator.GetSendTarget().SetTypeInfo(1)
	} else {
		generator.SetSendTarget(nil)
	}

	/* We increment to the next op, so we are at the correct position when the
	 * generator is resumed. */

	execute_data.GetOpline()++

	/* The GOTO VM uses a local opline variable. We need to set the opline
	 * variable in execute_data so we don't resume at an old position. */

	return -1
}
func ZEND_SWITCH_LONG_SPEC_CONST_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op *Zval
	var jump_zv *Zval
	var jumptable *HashTable
	op = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	jumptable = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant).GetValue().GetArr()
	if op.GetType() != 4 {
		if op.GetType() == 10 {
			op = &(*op).value.GetRef().GetVal()
		}
		if op.GetType() != 4 {

			/* Wrong type, fall back to ZEND_CASE chain */

			assert(EG.GetException() == nil)
			execute_data.SetOpline(opline + 1)
			return 0
		}
	}
	jump_zv = ZendHashIndexFind(jumptable, op.GetValue().GetLval())
	if jump_zv != nil {
		execute_data.SetOpline((*ZendOp)((*byte)(opline) + int(jump_zv.GetValue().GetLval())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
		return 0
	} else {

		/* default */

		execute_data.SetOpline((*ZendOp)((*byte)(opline) + int(opline.GetExtendedValue())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
		return 0
	}
}
func ZEND_SWITCH_STRING_SPEC_CONST_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op *Zval
	var jump_zv *Zval
	var jumptable *HashTable
	op = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	jumptable = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant).GetValue().GetArr()
	if op.GetType() != 6 {
		if 1<<0 == 1<<0 {

			/* Wrong type, fall back to ZEND_CASE chain */

			assert(EG.GetException() == nil)
			execute_data.SetOpline(opline + 1)
			return 0
		} else {
			if op.GetType() == 10 {
				op = &(*op).value.GetRef().GetVal()
			}
			if op.GetType() != 6 {

				/* Wrong type, fall back to ZEND_CASE chain */

				assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			}
		}
	}
	jump_zv = ZendHashFindEx(jumptable, op.GetValue().GetStr(), 1<<0 == 1<<0)
	if jump_zv != nil {
		execute_data.SetOpline((*ZendOp)((*byte)(opline) + int(jump_zv.GetValue().GetLval())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
		return 0
	} else {

		/* default */

		execute_data.SetOpline((*ZendOp)((*byte)(opline) + int(opline.GetExtendedValue())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
		return 0
	}
}
func ZEND_IN_ARRAY_SPEC_CONST_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var ht *HashTable = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant).GetValue().GetArr()
	var result *Zval
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	if op1.GetType() == 6 {
		result = ZendHashFindEx(ht, op1.GetValue().GetStr(), 1<<0 == 1<<0)
	} else if opline.GetExtendedValue() != 0 {
		if op1.GetType() == 4 {
			result = ZendHashIndexFind(ht, op1.GetValue().GetLval())
		} else {
			result = nil
		}
	} else if op1.GetType() <= 2 {
		result = ZendHashFindEx(ht, ZendEmptyString, 1)
	} else {
		var key *ZendString
		var key_tmp Zval
		var result_tmp Zval
		var val *Zval
		result = nil
		for {
			var __ht *HashTable = ht
			var _p *Bucket = __ht.GetArData()
			var _end *Bucket = _p + __ht.GetNNumUsed()
			for ; _p != _end; _p++ {
				var _z *Zval = &_p.val

				if _z.GetType() == 0 {
					continue
				}
				key = _p.GetKey()
				val = _z
				var __z *Zval = &key_tmp
				var __s *ZendString = key
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
				CompareFunction(&result_tmp, op1, &key_tmp)
				if result_tmp.GetValue().GetLval() == 0 {
					result = val
					break
				}
			}
			break
		}
	}
	for {
		if EG.GetException() != nil {
			break
		}
		if (opline + 1).GetOpcode() == 43 {
			if result != nil {
				execute_data.SetOpline(opline + 2)
			} else {
				execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
			}
		} else if (opline + 1).GetOpcode() == 44 {
			if result == nil {
				execute_data.SetOpline(opline + 2)
			} else {
				execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
			}
		} else {
			break
		}
		return 0
		break
	}
	if result != nil {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_ADD_SPEC_CONST_TMPVARCV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result *Zval
	var d1 float64
	var d2 float64
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	if 1<<0 == 1<<0 && (1<<1|1<<2|1<<3) == 1<<0 {

	} else if op1.GetTypeInfo() == 4 {
		if op2.GetTypeInfo() == 4 {
			result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			FastLongAddFunction(result, op1, op2)
			assert(EG.GetException() == nil)
			execute_data.SetOpline(opline + 1)
			return 0
		} else if op2.GetTypeInfo() == 5 {
			d1 = float64(op1.GetValue().GetLval())
			d2 = op2.GetValue().GetDval()
			goto add_double
		}
	} else if op1.GetTypeInfo() == 5 {
		if op2.GetTypeInfo() == 5 {
			d1 = op1.GetValue().GetDval()
			d2 = op2.GetValue().GetDval()
		add_double:
			result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var __z *Zval = result
			__z.GetValue().SetDval(d1 + d2)
			__z.SetTypeInfo(5)
			assert(EG.GetException() == nil)
			execute_data.SetOpline(opline + 1)
			return 0
		} else if op2.GetTypeInfo() == 4 {
			d1 = op1.GetValue().GetDval()
			d2 = float64(op2.GetValue().GetLval())
			goto add_double
		}
	}
	return zend_add_helper_SPEC(op1, op2, execute_data)
}
func ZEND_SUB_SPEC_CONST_TMPVARCV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result *Zval
	var d1 float64
	var d2 float64
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	if 1<<0 == 1<<0 && (1<<1|1<<2|1<<3) == 1<<0 {

	} else if op1.GetTypeInfo() == 4 {
		if op2.GetTypeInfo() == 4 {
			result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			FastLongSubFunction(result, op1, op2)
			assert(EG.GetException() == nil)
			execute_data.SetOpline(opline + 1)
			return 0
		} else if op2.GetTypeInfo() == 5 {
			d1 = float64(op1.GetValue().GetLval())
			d2 = op2.GetValue().GetDval()
			goto sub_double
		}
	} else if op1.GetTypeInfo() == 5 {
		if op2.GetTypeInfo() == 5 {
			d1 = op1.GetValue().GetDval()
			d2 = op2.GetValue().GetDval()
		sub_double:
			result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var __z *Zval = result
			__z.GetValue().SetDval(d1 - d2)
			__z.SetTypeInfo(5)
			assert(EG.GetException() == nil)
			execute_data.SetOpline(opline + 1)
			return 0
		} else if op2.GetTypeInfo() == 4 {
			d1 = op1.GetValue().GetDval()
			d2 = float64(op2.GetValue().GetLval())
			goto sub_double
		}
	}
	return zend_sub_helper_SPEC(op1, op2, execute_data)
}
func ZEND_MOD_SPEC_CONST_TMPVARCV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result *Zval
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	if 1<<0 == 1<<0 && (1<<1|1<<2|1<<3) == 1<<0 {

	} else if op1.GetTypeInfo() == 4 {
		if op2.GetTypeInfo() == 4 {
			result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			if op2.GetValue().GetLval() == 0 {
				return zend_mod_by_zero_helper_SPEC(execute_data)
			} else if op2.GetValue().GetLval() == -1 {

				/* Prevent overflow error/crash if op1==ZEND_LONG_MIN */

				var __z *Zval = result
				__z.GetValue().SetLval(0)
				__z.SetTypeInfo(4)
			} else {
				var __z *Zval = result
				__z.GetValue().SetLval(op1.GetValue().GetLval() % op2.GetValue().GetLval())
				__z.SetTypeInfo(4)
			}
			assert(EG.GetException() == nil)
			execute_data.SetOpline(opline + 1)
			return 0
		}
	}
	return zend_mod_helper_SPEC(op1, op2, execute_data)
}
func ZEND_SL_SPEC_CONST_TMPVARCV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	if 1<<0 == 1<<0 && (1<<1|1<<2|1<<3) == 1<<0 {

	} else if op1.GetTypeInfo() == 4 && op2.GetTypeInfo() == 4 && zend_ulong(*op2).value.lval < 8*8 {

		/* Perform shift on unsigned numbers to get well-defined wrap behavior. */

		var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		__z.GetValue().SetLval(zend_long(zend_ulong(*op1).value.lval << op2.GetValue().GetLval()))
		__z.SetTypeInfo(4)
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
	return zend_shift_left_helper_SPEC(op1, op2, execute_data)
}
func ZEND_SR_SPEC_CONST_TMPVARCV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	if 1<<0 == 1<<0 && (1<<1|1<<2|1<<3) == 1<<0 {

	} else if op1.GetTypeInfo() == 4 && op2.GetTypeInfo() == 4 && zend_ulong(*op2).value.lval < 8*8 {
		var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		__z.GetValue().SetLval(op1.GetValue().GetLval() >> op2.GetValue().GetLval())
		__z.SetTypeInfo(4)
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
	return zend_shift_right_helper_SPEC(op1, op2, execute_data)
}
func ZEND_IS_SMALLER_SPEC_CONST_TMPVARCV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	if 1<<0 == 1<<0 && (1<<1|1<<2|1<<3) == 1<<0 {

	} else if op1.GetTypeInfo() == 4 {
		if op2.GetTypeInfo() == 4 {
			if op1.GetValue().GetLval() < op2.GetValue().GetLval() {
			is_smaller_true:
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
				assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			} else {
			is_smaller_false:
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
				assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			}
		} else if op2.GetTypeInfo() == 5 {
			d1 = float64(op1.GetValue().GetLval())
			d2 = op2.GetValue().GetDval()
			goto is_smaller_double
		}
	} else if op1.GetTypeInfo() == 5 {
		if op2.GetTypeInfo() == 5 {
			d1 = op1.GetValue().GetDval()
			d2 = op2.GetValue().GetDval()
		is_smaller_double:
			if d1 < d2 {
				goto is_smaller_true
			} else {
				goto is_smaller_false
			}
		} else if op2.GetTypeInfo() == 4 {
			d1 = op1.GetValue().GetDval()
			d2 = float64(op2.GetValue().GetLval())
			goto is_smaller_double
		}
	}
	return zend_is_smaller_helper_SPEC(op1, op2, execute_data)
}
func ZEND_IS_SMALLER_SPEC_CONST_TMPVARCV_JMPZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	if 1<<0 == 1<<0 && (1<<1|1<<2|1<<3) == 1<<0 {

	} else if op1.GetTypeInfo() == 4 {
		if op2.GetTypeInfo() == 4 {
			if op1.GetValue().GetLval() < op2.GetValue().GetLval() {
			is_smaller_true:
				execute_data.SetOpline(opline + 2)
				return 0
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
				assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			} else {
			is_smaller_false:
				execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
				return 0
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
				assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			}
		} else if op2.GetTypeInfo() == 5 {
			d1 = float64(op1.GetValue().GetLval())
			d2 = op2.GetValue().GetDval()
			goto is_smaller_double
		}
	} else if op1.GetTypeInfo() == 5 {
		if op2.GetTypeInfo() == 5 {
			d1 = op1.GetValue().GetDval()
			d2 = op2.GetValue().GetDval()
		is_smaller_double:
			if d1 < d2 {
				goto is_smaller_true
			} else {
				goto is_smaller_false
			}
		} else if op2.GetTypeInfo() == 4 {
			d1 = op1.GetValue().GetDval()
			d2 = float64(op2.GetValue().GetLval())
			goto is_smaller_double
		}
	}
	return zend_is_smaller_helper_SPEC(op1, op2, execute_data)
}
func ZEND_IS_SMALLER_SPEC_CONST_TMPVARCV_JMPNZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	if 1<<0 == 1<<0 && (1<<1|1<<2|1<<3) == 1<<0 {

	} else if op1.GetTypeInfo() == 4 {
		if op2.GetTypeInfo() == 4 {
			if op1.GetValue().GetLval() < op2.GetValue().GetLval() {
			is_smaller_true:
				execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
				return 0
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
				assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			} else {
			is_smaller_false:
				execute_data.SetOpline(opline + 2)
				return 0
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
				assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			}
		} else if op2.GetTypeInfo() == 5 {
			d1 = float64(op1.GetValue().GetLval())
			d2 = op2.GetValue().GetDval()
			goto is_smaller_double
		}
	} else if op1.GetTypeInfo() == 5 {
		if op2.GetTypeInfo() == 5 {
			d1 = op1.GetValue().GetDval()
			d2 = op2.GetValue().GetDval()
		is_smaller_double:
			if d1 < d2 {
				goto is_smaller_true
			} else {
				goto is_smaller_false
			}
		} else if op2.GetTypeInfo() == 4 {
			d1 = op1.GetValue().GetDval()
			d2 = float64(op2.GetValue().GetLval())
			goto is_smaller_double
		}
	}
	return zend_is_smaller_helper_SPEC(op1, op2, execute_data)
}
func ZEND_IS_SMALLER_OR_EQUAL_SPEC_CONST_TMPVARCV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	if 1<<0 == 1<<0 && (1<<1|1<<2|1<<3) == 1<<0 {

	} else if op1.GetTypeInfo() == 4 {
		if op2.GetTypeInfo() == 4 {
			if op1.GetValue().GetLval() <= op2.GetValue().GetLval() {
			is_smaller_or_equal_true:
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
				assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			} else {
			is_smaller_or_equal_false:
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
				assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			}
		} else if op2.GetTypeInfo() == 5 {
			d1 = float64(op1.GetValue().GetLval())
			d2 = op2.GetValue().GetDval()
			goto is_smaller_or_equal_double
		}
	} else if op1.GetTypeInfo() == 5 {
		if op2.GetTypeInfo() == 5 {
			d1 = op1.GetValue().GetDval()
			d2 = op2.GetValue().GetDval()
		is_smaller_or_equal_double:
			if d1 <= d2 {
				goto is_smaller_or_equal_true
			} else {
				goto is_smaller_or_equal_false
			}
		} else if op2.GetTypeInfo() == 4 {
			d1 = op1.GetValue().GetDval()
			d2 = float64(op2.GetValue().GetLval())
			goto is_smaller_or_equal_double
		}
	}
	return zend_is_smaller_or_equal_helper_SPEC(op1, op2, execute_data)
}
func ZEND_IS_SMALLER_OR_EQUAL_SPEC_CONST_TMPVARCV_JMPZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	if 1<<0 == 1<<0 && (1<<1|1<<2|1<<3) == 1<<0 {

	} else if op1.GetTypeInfo() == 4 {
		if op2.GetTypeInfo() == 4 {
			if op1.GetValue().GetLval() <= op2.GetValue().GetLval() {
			is_smaller_or_equal_true:
				execute_data.SetOpline(opline + 2)
				return 0
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
				assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			} else {
			is_smaller_or_equal_false:
				execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
				return 0
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
				assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			}
		} else if op2.GetTypeInfo() == 5 {
			d1 = float64(op1.GetValue().GetLval())
			d2 = op2.GetValue().GetDval()
			goto is_smaller_or_equal_double
		}
	} else if op1.GetTypeInfo() == 5 {
		if op2.GetTypeInfo() == 5 {
			d1 = op1.GetValue().GetDval()
			d2 = op2.GetValue().GetDval()
		is_smaller_or_equal_double:
			if d1 <= d2 {
				goto is_smaller_or_equal_true
			} else {
				goto is_smaller_or_equal_false
			}
		} else if op2.GetTypeInfo() == 4 {
			d1 = op1.GetValue().GetDval()
			d2 = float64(op2.GetValue().GetLval())
			goto is_smaller_or_equal_double
		}
	}
	return zend_is_smaller_or_equal_helper_SPEC(op1, op2, execute_data)
}
func ZEND_IS_SMALLER_OR_EQUAL_SPEC_CONST_TMPVARCV_JMPNZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	if 1<<0 == 1<<0 && (1<<1|1<<2|1<<3) == 1<<0 {

	} else if op1.GetTypeInfo() == 4 {
		if op2.GetTypeInfo() == 4 {
			if op1.GetValue().GetLval() <= op2.GetValue().GetLval() {
			is_smaller_or_equal_true:
				execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
				return 0
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
				assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			} else {
			is_smaller_or_equal_false:
				execute_data.SetOpline(opline + 2)
				return 0
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
				assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			}
		} else if op2.GetTypeInfo() == 5 {
			d1 = float64(op1.GetValue().GetLval())
			d2 = op2.GetValue().GetDval()
			goto is_smaller_or_equal_double
		}
	} else if op1.GetTypeInfo() == 5 {
		if op2.GetTypeInfo() == 5 {
			d1 = op1.GetValue().GetDval()
			d2 = op2.GetValue().GetDval()
		is_smaller_or_equal_double:
			if d1 <= d2 {
				goto is_smaller_or_equal_true
			} else {
				goto is_smaller_or_equal_false
			}
		} else if op2.GetTypeInfo() == 4 {
			d1 = op1.GetValue().GetDval()
			d2 = float64(op2.GetValue().GetLval())
			goto is_smaller_or_equal_double
		}
	}
	return zend_is_smaller_or_equal_helper_SPEC(op1, op2, execute_data)
}
func ZEND_SUB_LONG_NO_OVERFLOW_SPEC_CONST_TMPVARCV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result *Zval
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	var __z *Zval = result
	__z.GetValue().SetLval(op1.GetValue().GetLval() - op2.GetValue().GetLval())
	__z.SetTypeInfo(4)
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_SUB_LONG_SPEC_CONST_TMPVARCV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result *Zval
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	FastLongSubFunction(result, op1, op2)
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_SUB_DOUBLE_SPEC_CONST_TMPVARCV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result *Zval
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	var __z *Zval = result
	__z.GetValue().SetDval(op1.GetValue().GetDval() - op2.GetValue().GetDval())
	__z.SetTypeInfo(5)
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_SMALLER_LONG_SPEC_CONST_TMPVARCV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	result = op1.GetValue().GetLval() < op2.GetValue().GetLval()
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_SMALLER_LONG_SPEC_CONST_TMPVARCV_JMPZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	result = op1.GetValue().GetLval() < op2.GetValue().GetLval()

	if result != 0 {
		execute_data.SetOpline(opline + 2)
	} else {
		execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
	}
	return 0
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_SMALLER_LONG_SPEC_CONST_TMPVARCV_JMPNZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	result = op1.GetValue().GetLval() < op2.GetValue().GetLval()

	if result == 0 {
		execute_data.SetOpline(opline + 2)
	} else {
		execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
	}
	return 0
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_SMALLER_DOUBLE_SPEC_CONST_TMPVARCV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	result = op1.GetValue().GetDval() < op2.GetValue().GetDval()
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_SMALLER_DOUBLE_SPEC_CONST_TMPVARCV_JMPZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	result = op1.GetValue().GetDval() < op2.GetValue().GetDval()

	if result != 0 {
		execute_data.SetOpline(opline + 2)
	} else {
		execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
	}
	return 0
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_SMALLER_DOUBLE_SPEC_CONST_TMPVARCV_JMPNZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	result = op1.GetValue().GetDval() < op2.GetValue().GetDval()

	if result == 0 {
		execute_data.SetOpline(opline + 2)
	} else {
		execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
	}
	return 0
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_CONST_TMPVARCV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	result = op1.GetValue().GetLval() <= op2.GetValue().GetLval()
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_CONST_TMPVARCV_JMPZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	result = op1.GetValue().GetLval() <= op2.GetValue().GetLval()

	if result != 0 {
		execute_data.SetOpline(opline + 2)
	} else {
		execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
	}
	return 0
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_CONST_TMPVARCV_JMPNZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	result = op1.GetValue().GetLval() <= op2.GetValue().GetLval()

	if result == 0 {
		execute_data.SetOpline(opline + 2)
	} else {
		execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
	}
	return 0
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_CONST_TMPVARCV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	result = op1.GetValue().GetDval() <= op2.GetValue().GetDval()
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_CONST_TMPVARCV_JMPZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	result = op1.GetValue().GetDval() <= op2.GetValue().GetDval()

	if result != 0 {
		execute_data.SetOpline(opline + 2)
	} else {
		execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
	}
	return 0
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_CONST_TMPVARCV_JMPNZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	result = op1.GetValue().GetDval() <= op2.GetValue().GetDval()

	if result == 0 {
		execute_data.SetOpline(opline + 2)
	} else {
		execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
	}
	return 0
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_FETCH_DIM_R_INDEX_SPEC_CONST_TMPVARCV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var container *Zval
	var dim *Zval
	var value *Zval
	var offset ZendLong
	var ht *HashTable
	container = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	dim = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	if container.GetType() == 7 {
	fetch_dim_r_index_array:
		if dim.GetType() == 4 {
			offset = dim.GetValue().GetLval()
		} else {
			offset = ZvalGetLong(dim)
		}
		ht = container.GetValue().GetArr()
		if (ht.GetUFlags() & 1 << 2) != 0 {
			if zend_ulong(offset) < zend_ulong(ht).nNumUsed {
				value = &ht.arData[offset].GetVal()
				if value.GetType() == 0 {
					goto fetch_dim_r_index_undef
				}
			} else {
				goto fetch_dim_r_index_undef
			}
		} else {
			value = _zendHashIndexFind(ht, offset)
			if value == nil {
				goto fetch_dim_r_index_undef
			}
		}
		var _z3 *Zval = value
		if (_z3.GetTypeInfo() & 0xff00) != 0 {
			if (_z3.GetTypeInfo() & 0xff) == 10 {
				_z3 = &(*_z3).value.GetRef().GetVal()
				if (_z3.GetTypeInfo() & 0xff00) != 0 {
					ZvalAddrefP(_z3)
				}
			} else {
				ZvalAddrefP(_z3)
			}
		}
		var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var _z2 *Zval = _z3
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if (1 << 0 & (1<<1 | 1<<2)) != 0 {
			execute_data.SetOpline(execute_data.GetOpline() + 1)
			return 0
		} else {
			assert(EG.GetException() == nil)
			execute_data.SetOpline(opline + 1)
			return 0
		}
	} else if 1<<0 != 1<<0 && container.GetType() == 10 {
		container = &(*container).value.GetRef().GetVal()
		if container.GetType() == 7 {
			goto fetch_dim_r_index_array
		} else {
			goto fetch_dim_r_index_slow
		}
	} else {
	fetch_dim_r_index_slow:
		if (1<<1|1<<2|1<<3) == 1<<0 && dim.GetU2Extra() == 1 {
			dim++
		}
		zend_fetch_dimension_address_read_R_slow(container, dim, opline, execute_data)
		execute_data.SetOpline(execute_data.GetOpline() + 1)
		return 0
	}
fetch_dim_r_index_undef:
	(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
	ZendUndefinedOffset(offset)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_DIV_SPEC_CONST_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var op1 *Zval
	var op2 *Zval
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	FastDivFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op1, op2)
	ZvalPtrDtorNogc(free_op2)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_POW_SPEC_CONST_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var op1 *Zval
	var op2 *Zval
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	PowFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op1, op2)
	ZvalPtrDtorNogc(free_op2)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_CONCAT_SPEC_CONST_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var op1 *Zval
	var op2 *Zval
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	if (1<<0 == 1<<0 || op1.GetType() == 6) && ((1<<1|1<<2) == 1<<0 || op2.GetType() == 6) {
		var op1_str *ZendString = op1.GetValue().GetStr()
		var op2_str *ZendString = op2.GetValue().GetStr()
		var str *ZendString
		if 1<<0 != 1<<0 && op1_str.GetLen() == 0 {
			if (1<<1|1<<2) == 1<<0 || (1<<1|1<<2) == 1<<3 {
				var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				var __s *ZendString = op2_str
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					ZendGcAddref(&__s.gc)
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
			} else {
				var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				var __s *ZendString = op2_str
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
			}
			if (1 << 0 & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op1_str, 0)
			}
		} else if (1<<1|1<<2) != 1<<0 && op2_str.GetLen() == 0 {
			if 1<<0 == 1<<0 || 1<<0 == 1<<3 {
				var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				var __s *ZendString = op1_str
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					ZendGcAddref(&__s.gc)
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
			} else {
				var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				var __s *ZendString = op1_str
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
			}
			if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op2_str, 0)
			}
		} else if 1<<0 != 1<<0 && 1<<0 != 1<<3 && (ZvalGcFlags(op1_str.GetGc().GetTypeInfo())&1<<6) == 0 && ZendGcRefcount(&op1_str.gc) == 1 {
			var len_ int = op1_str.GetLen()
			str = ZendStringExtend(op1_str, len_+op2_str.GetLen(), 0)
			memcpy(str.GetVal()+len_, op2_str.GetVal(), op2_str.GetLen()+1)
			var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var __s *ZendString = str
			__z.GetValue().SetStr(__s)
			__z.SetTypeInfo(6 | 1<<0<<8)
			if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op2_str, 0)
			}
		} else {
			str = ZendStringAlloc(op1_str.GetLen()+op2_str.GetLen(), 0)
			memcpy(str.GetVal(), op1_str.GetVal(), op1_str.GetLen())
			memcpy(str.GetVal()+op1_str.GetLen(), op2_str.GetVal(), op2_str.GetLen()+1)
			var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var __s *ZendString = str
			__z.GetValue().SetStr(__s)
			__z.SetTypeInfo(6 | 1<<0<<8)
			if (1 << 0 & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op1_str, 0)
			}
			if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op2_str, 0)
			}
		}
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	} else {
		if 1<<0 == 1<<3 && op1.GetType() == 0 {
			op1 = _zvalUndefinedOp1(execute_data)
		}
		if (1<<1|1<<2) == 1<<3 && op2.GetType() == 0 {
			op2 = _zvalUndefinedOp2(execute_data)
		}
		ConcatFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op1, op2)
		ZvalPtrDtorNogc(free_op2)
		execute_data.SetOpline(execute_data.GetOpline() + 1)
		return 0
	}
}
func ZEND_SPACESHIP_SPEC_CONST_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var op1 *Zval
	var op2 *Zval
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	CompareFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op1, op2)
	ZvalPtrDtorNogc(free_op2)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_DIM_R_SPEC_CONST_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var container *Zval
	var dim *Zval
	var value *Zval
	container = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	dim = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	if 1<<0 != 1<<0 {
		if container.GetType() == 7 {
		fetch_dim_r_array:
			value = ZendFetchDimensionAddressInner(container.GetValue().GetArr(), dim, 1<<1|1<<2, 0, execute_data)
			var _z3 *Zval = value
			if (_z3.GetTypeInfo() & 0xff00) != 0 {
				if (_z3.GetTypeInfo() & 0xff) == 10 {
					_z3 = &(*_z3).value.GetRef().GetVal()
					if (_z3.GetTypeInfo() & 0xff00) != 0 {
						ZvalAddrefP(_z3)
					}
				} else {
					ZvalAddrefP(_z3)
				}
			}
			var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var _z2 *Zval = _z3
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
		} else if container.GetType() == 10 {
			container = &(*container).value.GetRef().GetVal()
			if container.GetType() == 7 {
				goto fetch_dim_r_array
			} else {
				goto fetch_dim_r_slow
			}
		} else {
		fetch_dim_r_slow:
			if (1<<1|1<<2) == 1<<0 && dim.GetU2Extra() == 1 {
				dim++
			}
			zend_fetch_dimension_address_read_R_slow(container, dim, opline, execute_data)
		}
	} else {
		zend_fetch_dimension_address_read_R(container, dim, 1<<1|1<<2, opline, execute_data)
	}
	ZvalPtrDtorNogc(free_op2)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_DIM_IS_SPEC_CONST_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var container *Zval
	container = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	zend_fetch_dimension_address_read_IS(container, _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data), 1<<1|1<<2, opline, execute_data)
	ZvalPtrDtorNogc(free_op2)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_DIM_FUNC_ARG_SPEC_CONST_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	if (execute_data.GetCall().GetThis().GetTypeInfo() & 1 << 31) != 0 {
		if (1 << 0 & (1<<0 | 1<<1)) != 0 {
			return zend_use_tmp_in_write_context_helper_SPEC(execute_data)
		}
		return ZEND_NULL_HANDLER(execute_data)
	} else {
		if (1<<1 | 1<<2) == 0 {
			return zend_use_undef_in_read_context_helper_SPEC(execute_data)
		}
		return ZEND_FETCH_DIM_R_SPEC_CONST_TMPVAR_HANDLER(execute_data)
	}
}
func ZEND_FETCH_OBJ_R_SPEC_CONST_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var container *Zval
	var free_op2 ZendFreeOp
	var offset *Zval
	var cache_slot *any = nil
	container = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	if 1<<0 == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	offset = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	if 1<<0 == 1<<0 || 1<<0 != 0 && container.GetType() != 8 {
		for {
			if (1<<0&(1<<2|1<<3)) != 0 && container.GetType() == 10 {
				container = &(*container).value.GetRef().GetVal()
				if container.GetType() == 8 {
					break
				}
			}
			if 1<<0 == 1<<3 && container.GetType() == 0 {
				_zvalUndefinedOp1(execute_data)
			}
			if (1<<1|1<<2) == 1<<3 && offset.GetType() == 0 {
				_zvalUndefinedOp2(execute_data)
			}
			ZendWrongPropertyRead(offset)
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
			goto fetch_obj_r_finish
			break
		}
	}

	/* here we are sure we are dealing with an object */

	var zobj *ZendObject = container.GetValue().GetObj()
	var retval *Zval
	if (1<<1 | 1<<2) == 1<<0 {
		cache_slot = (*any)((*byte)(execute_data.GetRunTimeCache() + (opline.GetExtendedValue() & ^1)))
		if zobj.GetCe() == cache_slot[0] {
			var prop_offset uintPtr = uintptr_t(cache_slot + 1)[0]
			if intptr_t(prop_offset) > 0 {
				retval = (*Zval)((*byte)(zobj + prop_offset))
				if retval.GetTypeInfo() != 0 {
					if (1 << 0 & (1<<1 | 1<<2)) != 0 {
						goto fetch_obj_r_copy
					} else {
					fetch_obj_r_fast_copy:
						var _z3 *Zval = retval
						if (_z3.GetTypeInfo() & 0xff00) != 0 {
							if (_z3.GetTypeInfo() & 0xff) == 10 {
								_z3 = &(*_z3).value.GetRef().GetVal()
								if (_z3.GetTypeInfo() & 0xff00) != 0 {
									ZvalAddrefP(_z3)
								}
							} else {
								ZvalAddrefP(_z3)
							}
						}
						var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
						var _z2 *Zval = _z3
						var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
						var _t uint32 = _z2.GetTypeInfo()
						_z1.GetValue().SetCounted(_gc)
						_z1.SetTypeInfo(_t)
						assert(EG.GetException() == nil)
						execute_data.SetOpline(opline + 1)
						return 0
					}
				}
			} else if zobj.GetProperties() != nil {
				if prop_offset != uintptr_t(intptr_t)(-1) {
					var idx uintPtr = uintptr_t(-(intptr_t(prop_offset)) - 2)
					if idx < zobj.GetProperties().GetNNumUsed()*g.SizeOf("Bucket") {
						var p *Bucket = (*Bucket)((*byte)(zobj.GetProperties().GetArData() + idx))
						if p.GetVal().GetType() != 0 && (p.GetKey() == offset.GetValue().GetStr() || p.GetH() == offset.GetValue().GetStr().GetH() && p.GetKey() != nil && ZendStringEqualContent(p.GetKey(), offset.GetValue().GetStr()) != 0) {
							retval = &p.val
							if (1 << 0 & (1<<1 | 1<<2)) != 0 {
								goto fetch_obj_r_copy
							} else {
								goto fetch_obj_r_fast_copy
							}
						}
					}
					(cache_slot + 1)[0] = any(uintptr_t(intptr_t)(-1))
				}
				retval = ZendHashFindEx(zobj.GetProperties(), offset.GetValue().GetStr(), 1)
				if retval != nil {
					var idx uintPtr = (*byte)(retval - (*byte)(zobj.GetProperties().GetArData()))
					(cache_slot + 1)[0] = any(uintptr_t(-(intptr_t(idx) + 2)))
					if (1 << 0 & (1<<1 | 1<<2)) != 0 {
						goto fetch_obj_r_copy
					} else {
						goto fetch_obj_r_fast_copy
					}
				}
			}
		}
	} else if (1<<1|1<<2) == 1<<3 && offset.GetTypeInfo() == 0 {
		_zvalUndefinedOp2(execute_data)
	}
	retval = zobj.GetHandlers().GetReadProperty()(container, offset, 0, cache_slot, (*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())))
	if retval != (*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())) {
	fetch_obj_r_copy:
		var _z3 *Zval = retval
		if (_z3.GetTypeInfo() & 0xff00) != 0 {
			if (_z3.GetTypeInfo() & 0xff) == 10 {
				_z3 = &(*_z3).value.GetRef().GetVal()
				if (_z3.GetTypeInfo() & 0xff00) != 0 {
					ZvalAddrefP(_z3)
				}
			} else {
				ZvalAddrefP(_z3)
			}
		}
		var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var _z2 *Zval = _z3
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
	} else if retval.GetType() == 10 {
		ZendUnwrapReference(retval)
	}
fetch_obj_r_finish:
	ZvalPtrDtorNogc(free_op2)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_OBJ_IS_SPEC_CONST_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var container *Zval
	var free_op2 ZendFreeOp
	var offset *Zval
	var cache_slot *any = nil
	container = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	if 1<<0 == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	offset = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	if 1<<0 == 1<<0 || 1<<0 != 0 && container.GetType() != 8 {
		for {
			if (1<<0&(1<<2|1<<3)) != 0 && container.GetType() == 10 {
				container = &(*container).value.GetRef().GetVal()
				if container.GetType() == 8 {
					break
				}
			}
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
			goto fetch_obj_is_finish
			break
		}
	}

	/* here we are sure we are dealing with an object */

	var zobj *ZendObject = container.GetValue().GetObj()
	var retval *Zval
	if (1<<1 | 1<<2) == 1<<0 {
		cache_slot = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))
		if zobj.GetCe() == cache_slot[0] {
			var prop_offset uintPtr = uintptr_t(cache_slot + 1)[0]
			if intptr_t(prop_offset) > 0 {
				retval = (*Zval)((*byte)(zobj + prop_offset))
				if retval.GetType() != 0 {
					if (1 << 0 & (1<<1 | 1<<2)) != 0 {
						goto fetch_obj_is_copy
					} else {
					fetch_obj_is_fast_copy:
						var _z3 *Zval = retval
						if (_z3.GetTypeInfo() & 0xff00) != 0 {
							if (_z3.GetTypeInfo() & 0xff) == 10 {
								_z3 = &(*_z3).value.GetRef().GetVal()
								if (_z3.GetTypeInfo() & 0xff00) != 0 {
									ZvalAddrefP(_z3)
								}
							} else {
								ZvalAddrefP(_z3)
							}
						}
						var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
						var _z2 *Zval = _z3
						var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
						var _t uint32 = _z2.GetTypeInfo()
						_z1.GetValue().SetCounted(_gc)
						_z1.SetTypeInfo(_t)
						assert(EG.GetException() == nil)
						execute_data.SetOpline(opline + 1)
						return 0
					}
				}
			} else if zobj.GetProperties() != nil {
				if prop_offset != uintptr_t(intptr_t)(-1) {
					var idx uintPtr = uintptr_t(-(intptr_t(prop_offset)) - 2)
					if idx < zobj.GetProperties().GetNNumUsed()*g.SizeOf("Bucket") {
						var p *Bucket = (*Bucket)((*byte)(zobj.GetProperties().GetArData() + idx))
						if p.GetVal().GetType() != 0 && (p.GetKey() == offset.GetValue().GetStr() || p.GetH() == offset.GetValue().GetStr().GetH() && p.GetKey() != nil && ZendStringEqualContent(p.GetKey(), offset.GetValue().GetStr()) != 0) {
							retval = &p.val
							if (1 << 0 & (1<<1 | 1<<2)) != 0 {
								goto fetch_obj_is_copy
							} else {
								goto fetch_obj_is_fast_copy
							}
						}
					}
					(cache_slot + 1)[0] = any(uintptr_t(intptr_t)(-1))
				}
				retval = ZendHashFindEx(zobj.GetProperties(), offset.GetValue().GetStr(), 1)
				if retval != nil {
					var idx uintPtr = (*byte)(retval - (*byte)(zobj.GetProperties().GetArData()))
					(cache_slot + 1)[0] = any(uintptr_t(-(intptr_t(idx) + 2)))
					if (1 << 0 & (1<<1 | 1<<2)) != 0 {
						goto fetch_obj_is_copy
					} else {
						goto fetch_obj_is_fast_copy
					}
				}
			}
		}
	}
	retval = zobj.GetHandlers().GetReadProperty()(container, offset, 3, cache_slot, (*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())))
	if retval != (*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())) {
	fetch_obj_is_copy:
		var _z3 *Zval = retval
		if (_z3.GetTypeInfo() & 0xff00) != 0 {
			if (_z3.GetTypeInfo() & 0xff) == 10 {
				_z3 = &(*_z3).value.GetRef().GetVal()
				if (_z3.GetTypeInfo() & 0xff00) != 0 {
					ZvalAddrefP(_z3)
				}
			} else {
				ZvalAddrefP(_z3)
			}
		}
		var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var _z2 *Zval = _z3
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
	} else if retval.GetType() == 10 {
		ZendUnwrapReference(retval)
	}
fetch_obj_is_finish:
	ZvalPtrDtorNogc(free_op2)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_OBJ_FUNC_ARG_SPEC_CONST_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	if (execute_data.GetCall().GetThis().GetTypeInfo() & 1 << 31) != 0 {

		/* Behave like FETCH_OBJ_W */

		if (1 << 0 & (1<<0 | 1<<1)) != 0 {
			return zend_use_tmp_in_write_context_helper_SPEC(execute_data)
		}
		return ZEND_NULL_HANDLER(execute_data)
	} else {
		return ZEND_FETCH_OBJ_R_SPEC_CONST_TMPVAR_HANDLER(execute_data)
	}
}
func ZEND_FETCH_LIST_R_SPEC_CONST_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var container *Zval
	container = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	zend_fetch_dimension_address_LIST_r(container, _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data), 1<<1|1<<2, opline, execute_data)
	ZvalPtrDtorNogc(free_op2)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FAST_CONCAT_SPEC_CONST_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var op1 *Zval
	var op2 *Zval
	var op1_str *ZendString
	var op2_str *ZendString
	var str *ZendString
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	if (1<<0 == 1<<0 || op1.GetType() == 6) && ((1<<1|1<<2) == 1<<0 || op2.GetType() == 6) {
		var op1_str *ZendString = op1.GetValue().GetStr()
		var op2_str *ZendString = op2.GetValue().GetStr()
		var str *ZendString
		if 1<<0 != 1<<0 && op1_str.GetLen() == 0 {
			if (1<<1|1<<2) == 1<<0 || (1<<1|1<<2) == 1<<3 {
				var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				var __s *ZendString = op2_str
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					ZendGcAddref(&__s.gc)
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
			} else {
				var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				var __s *ZendString = op2_str
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
			}
			if (1 << 0 & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op1_str, 0)
			}
		} else if (1<<1|1<<2) != 1<<0 && op2_str.GetLen() == 0 {
			if 1<<0 == 1<<0 || 1<<0 == 1<<3 {
				var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				var __s *ZendString = op1_str
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					ZendGcAddref(&__s.gc)
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
			} else {
				var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				var __s *ZendString = op1_str
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
			}
			if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op2_str, 0)
			}
		} else if 1<<0 != 1<<0 && 1<<0 != 1<<3 && (ZvalGcFlags(op1_str.GetGc().GetTypeInfo())&1<<6) == 0 && ZendGcRefcount(&op1_str.gc) == 1 {
			var len_ int = op1_str.GetLen()
			if len_ > SIZE_MAX-(zend_long((*byte)(&((*ZendString)(nil).GetVal()))-(*byte)(nil))+1+8 - 1 & ^(8-1))-op2_str.GetLen() {
				ZendErrorNoreturn(1<<0, "Integer overflow in memory allocation")
			}
			str = ZendStringExtend(op1_str, len_+op2_str.GetLen(), 0)
			memcpy(str.GetVal()+len_, op2_str.GetVal(), op2_str.GetLen()+1)
			var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var __s *ZendString = str
			__z.GetValue().SetStr(__s)
			__z.SetTypeInfo(6 | 1<<0<<8)
			if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op2_str, 0)
			}
		} else {
			str = ZendStringAlloc(op1_str.GetLen()+op2_str.GetLen(), 0)
			memcpy(str.GetVal(), op1_str.GetVal(), op1_str.GetLen())
			memcpy(str.GetVal()+op1_str.GetLen(), op2_str.GetVal(), op2_str.GetLen()+1)
			var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var __s *ZendString = str
			__z.GetValue().SetStr(__s)
			__z.SetTypeInfo(6 | 1<<0<<8)
			if (1 << 0 & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op1_str, 0)
			}
			if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op2_str, 0)
			}
		}
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
	if 1<<0 == 1<<0 {
		op1_str = op1.GetValue().GetStr()
	} else if op1.GetType() == 6 {
		op1_str = ZendStringCopy(op1.GetValue().GetStr())
	} else {
		if 1<<0 == 1<<3 && op1.GetType() == 0 {
			_zvalUndefinedOp1(execute_data)
		}
		op1_str = ZvalGetStringFunc(op1)
	}
	if (1<<1 | 1<<2) == 1<<0 {
		op2_str = op2.GetValue().GetStr()
	} else if op2.GetType() == 6 {
		op2_str = ZendStringCopy(op2.GetValue().GetStr())
	} else {
		if (1<<1|1<<2) == 1<<3 && op2.GetType() == 0 {
			_zvalUndefinedOp2(execute_data)
		}
		op2_str = ZvalGetStringFunc(op2)
	}
	for {
		if 1<<0 != 1<<0 {
			if op1_str.GetLen() == 0 {
				if (1<<1 | 1<<2) == 1<<0 {
					if op2.GetTypeFlags() != 0 {
						ZendGcAddref(&op2_str.gc)
					}
				}
				var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				var __s *ZendString = op2_str
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
				ZendStringReleaseEx(op1_str, 0)
				break
			}
		}
		if (1<<1 | 1<<2) != 1<<0 {
			if op2_str.GetLen() == 0 {
				if 1<<0 == 1<<0 {
					if op1.GetTypeFlags() != 0 {
						ZendGcAddref(&op1_str.gc)
					}
				}
				var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				var __s *ZendString = op1_str
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
				ZendStringReleaseEx(op2_str, 0)
				break
			}
		}
		str = ZendStringAlloc(op1_str.GetLen()+op2_str.GetLen(), 0)
		memcpy(str.GetVal(), op1_str.GetVal(), op1_str.GetLen())
		memcpy(str.GetVal()+op1_str.GetLen(), op2_str.GetVal(), op2_str.GetLen()+1)
		var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var __s *ZendString = str
		__z.GetValue().SetStr(__s)
		__z.SetTypeInfo(6 | 1<<0<<8)
		if 1<<0 != 1<<0 {
			ZendStringReleaseEx(op1_str, 0)
		}
		if (1<<1 | 1<<2) != 1<<0 {
			ZendStringReleaseEx(op2_str, 0)
		}
		break
	}
	ZvalPtrDtorNogc(free_op2)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_INIT_METHOD_CALL_SPEC_CONST_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var function_name *Zval
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var object *Zval
	var fbc *ZendFunction
	var called_scope *ZendClassEntry
	var obj *ZendObject
	var call *ZendExecuteData
	var call_info uint32
	object = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	if 1<<0 == 0 && object.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	if (1<<1 | 1<<2) != 1<<0 {
		function_name = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	}
	if (1<<1|1<<2) != 1<<0 && function_name.GetType() != 6 {
		for {
			if ((1<<1|1<<2)&(1<<2|1<<3)) != 0 && function_name.GetType() == 10 {
				function_name = &(*function_name).value.GetRef().GetVal()
				if function_name.GetType() == 6 {
					break
				}
			} else if (1<<1|1<<2) == 1<<3 && function_name.GetType() == 0 {
				_zvalUndefinedOp2(execute_data)
				if EG.GetException() != nil {
					return 0
				}
			}
			ZendThrowError(nil, "Method name must be a string")
			ZvalPtrDtorNogc(free_op2)
			return 0
			break
		}
	}
	if 1<<0 != 0 {
		for {
			if 1<<0 == 1<<0 || object.GetType() != 8 {
				if (1<<0&(1<<2|1<<3)) != 0 && object.GetType() == 10 {
					object = &(*object).value.GetRef().GetVal()
					if object.GetType() == 8 {
						break
					}
				}
				if 1<<0 == 1<<3 && object.GetType() == 0 {
					object = _zvalUndefinedOp1(execute_data)
					if EG.GetException() != nil {
						if (1<<1 | 1<<2) != 1<<0 {
							ZvalPtrDtorNogc(free_op2)
						}
						return 0
					}
				}
				if (1<<1 | 1<<2) == 1<<0 {
					function_name = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
				}
				ZendInvalidMethodCall(object, function_name)
				ZvalPtrDtorNogc(free_op2)
				return 0
			}
			break
		}
	}
	obj = object.GetValue().GetObj()
	called_scope = obj.GetCe()
	if (1<<1|1<<2) == 1<<0 && (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetResult().GetNum()))[0] == called_scope {
		fbc = (*any)((*byte)(execute_data.GetRunTimeCache() + (opline.GetResult().GetNum() + g.SizeOf("void *"))))[0]
	} else {
		var orig_obj *ZendObject = obj
		if (1<<1 | 1<<2) == 1<<0 {
			function_name = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
		}

		/* First, locate the function. */

		fbc = obj.GetHandlers().GetGetMethod()(&obj, function_name.GetValue().GetStr(), g.CondF1((1<<1|1<<2) == 1<<0, func() *Zval { return (*Zval)((*byte)(opline)+int32(opline.GetOp2()).constant) + 1 }, nil))
		if fbc == nil {
			if EG.GetException() == nil {
				ZendUndefinedMethod(obj.GetCe(), function_name.GetValue().GetStr())
			}
			ZvalPtrDtorNogc(free_op2)
			return 0
		}
		if (1<<1|1<<2) == 1<<0 && fbc.GetType() <= 2 && (fbc.GetFnFlags()&(1<<18|1<<19)) == 0 && obj == orig_obj {
			var slot *any = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetResult().GetNum()))
			slot[0] = called_scope
			slot[1] = fbc
		}
		if (1<<0&(1<<2|1<<1)) != 0 && obj != orig_obj {

			/* Reset "object" to trigger reference counting */

			object = nil

			/* Reset "object" to trigger reference counting */

		}
		if fbc.GetType() == 2 && !(g.CondF((uintptr_t(&fbc.op_array).run_time_cache__ptr&1) != 0, func() any {
			return *((*any)((*byte)(CG.GetMapPtrBase() + uintptr_t(&fbc.op_array).run_time_cache__ptr - 1)))
		}, func() any { return any(*(&fbc.op_array.run_time_cache__ptr)) })) {
			InitFuncRunTimeCache(&fbc.op_array)
		}
	}
	if (1<<1 | 1<<2) != 1<<0 {
		ZvalPtrDtorNogc(free_op2)
	}
	call_info = 0<<16 | 0<<17 | (8 | 1<<0<<8 | 1<<1<<8)
	if (fbc.GetFnFlags() & 1 << 4) != 0 {
		if (1<<0&(1<<2|1<<1)) != 0 && EG.GetException() != nil {
			return 0
		}

		/* call static method */

		obj = (*ZendObject)(called_scope)
		call_info = 0<<16 | 0<<17
	} else if (1 << 0 & (1<<2 | 1<<1 | 1<<3)) != 0 {
		if 1<<0 == 1<<3 {
			ZendGcAddref(&obj.gc)
		} else if free_op1 != object {
			ZendGcAddref(&obj.gc)
		}

		/* CV may be changed indirectly (e.g. when it's a reference) */

		call_info = 0<<16 | 0<<17 | (8 | 1<<0<<8 | 1<<1<<8) | 1<<21

		/* CV may be changed indirectly (e.g. when it's a reference) */

	}
	call = ZendVmStackPushCallFrame(call_info, fbc, opline.GetExtendedValue(), obj)
	call.SetPrevExecuteData(execute_data.GetCall())
	execute_data.SetCall(call)
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_INIT_STATIC_METHOD_CALL_SPEC_CONST_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var function_name *Zval
	var ce *ZendClassEntry
	var call_info uint32
	var fbc *ZendFunction
	var call *ZendExecuteData
	if 1<<0 == 1<<0 {

		/* no function found. try a static method in class */

		ce = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetResult().GetNum()))[0]
		if ce == nil {
			ce = ZendFetchClassByName((*Zval)((*byte)(opline)+int32(opline.GetOp1()).constant).GetValue().GetStr(), ((*Zval)((*byte)(opline)+int32(opline.GetOp1()).constant) + 1).GetValue().GetStr(), 0|0x200)
			if ce == nil {
				assert(EG.GetException() != nil)
				ZvalPtrDtorNogc((*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar())))
				return 0
			}
			if (1<<1 | 1<<2) != 1<<0 {
				(*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetResult().GetNum()))[0] = ce
			}
		}
	} else if 1<<0 == 0 {
		ce = ZendFetchClass(nil, opline.GetOp1().GetNum())
		if ce == nil {
			assert(EG.GetException() != nil)
			ZvalPtrDtorNogc((*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar())))
			return 0
		}
	} else {
		ce = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar())).GetValue().GetCe()
	}
	if 1<<0 == 1<<0 && (1<<1|1<<2) == 1<<0 && g.Assign(&fbc, (*any)((*byte)(execute_data.GetRunTimeCache() + (opline.GetResult().GetNum() + g.SizeOf("void *"))))[0]) != nil {

	} else if 1<<0 != 1<<0 && (1<<1|1<<2) == 1<<0 && (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetResult().GetNum()))[0] == ce {
		fbc = (*any)((*byte)(execute_data.GetRunTimeCache() + (opline.GetResult().GetNum() + g.SizeOf("void *"))))[0]
	} else if (1<<1 | 1<<2) != 0 {
		var free_op2 ZendFreeOp
		function_name = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
		if (1<<1 | 1<<2) != 1<<0 {
			if function_name.GetType() != 6 {
				for {
					if ((1<<1|1<<2)&(1<<2|1<<3)) != 0 && function_name.GetType() == 10 {
						function_name = &(*function_name).value.GetRef().GetVal()
						if function_name.GetType() == 6 {
							break
						}
					} else if (1<<1|1<<2) == 1<<3 && function_name.GetType() == 0 {
						_zvalUndefinedOp2(execute_data)
						if EG.GetException() != nil {
							return 0
						}
					}
					ZendThrowError(nil, "Function name must be a string")
					ZvalPtrDtorNogc(free_op2)
					return 0
					break
				}
			}
		}
		if ce.GetGetStaticMethod() != nil {
			fbc = ce.GetGetStaticMethod()(ce, function_name.GetValue().GetStr())
		} else {
			fbc = ZendStdGetStaticMethod(ce, function_name.GetValue().GetStr(), g.CondF1((1<<1|1<<2) == 1<<0, func() *Zval { return (*Zval)((*byte)(opline)+int32(opline.GetOp2()).constant) + 1 }, nil))
		}
		if fbc == nil {
			if EG.GetException() == nil {
				ZendUndefinedMethod(ce, function_name.GetValue().GetStr())
			}
			ZvalPtrDtorNogc(free_op2)
			return 0
		}
		if (1<<1|1<<2) == 1<<0 && fbc.GetType() <= 2 && (fbc.GetFnFlags()&(1<<18|1<<19)) == 0 {
			var slot *any = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetResult().GetNum()))
			slot[0] = ce
			slot[1] = fbc
		}
		if fbc.GetType() == 2 && !(g.CondF((uintptr_t(&fbc.op_array).run_time_cache__ptr&1) != 0, func() any {
			return *((*any)((*byte)(CG.GetMapPtrBase() + uintptr_t(&fbc.op_array).run_time_cache__ptr - 1)))
		}, func() any { return any(*(&fbc.op_array.run_time_cache__ptr)) })) {
			InitFuncRunTimeCache(&fbc.op_array)
		}
		if (1<<1 | 1<<2) != 1<<0 {
			ZvalPtrDtorNogc(free_op2)
		}
	} else {
		if ce.GetConstructor() == nil {
			ZendThrowError(nil, "Cannot call constructor")
			return 0
		}
		if execute_data.GetThis().GetType() == 8 && execute_data.GetThis().GetValue().GetObj().GetCe() != ce.GetConstructor().GetScope() && (ce.GetConstructor().GetFnFlags()&1<<2) != 0 {
			ZendThrowError(nil, "Cannot call private %s::__construct()", ce.GetName().GetVal())
			return 0
		}
		fbc = ce.GetConstructor()
		if fbc.GetType() == 2 && !(g.CondF((uintptr_t(&fbc.op_array).run_time_cache__ptr&1) != 0, func() any {
			return *((*any)((*byte)(CG.GetMapPtrBase() + uintptr_t(&fbc.op_array).run_time_cache__ptr - 1)))
		}, func() any { return any(*(&fbc.op_array.run_time_cache__ptr)) })) {
			InitFuncRunTimeCache(&fbc.op_array)
		}
	}
	if (fbc.GetFnFlags() & 1 << 4) == 0 {
		if execute_data.GetThis().GetType() == 8 && InstanceofFunction(execute_data.GetThis().GetValue().GetObj().GetCe(), ce) != 0 {
			ce = (*ZendClassEntry)(execute_data.GetThis().GetValue().GetObj())
			call_info = 0<<16 | 0<<17 | (8 | 1<<0<<8 | 1<<1<<8)
		} else {
			ZendNonStaticMethodCall(fbc)
			if EG.GetException() != nil {
				return 0
			}
			goto check_parent_and_self
		}
	} else {
	check_parent_and_self:

		/* previous opcode is ZEND_FETCH_CLASS */

		if 1<<0 == 0 && ((opline.GetOp1().GetNum()&0xf) == 2 || (opline.GetOp1().GetNum()&0xf) == 1) {
			if execute_data.GetThis().GetType() == 8 {
				ce = execute_data.GetThis().GetValue().GetObj().GetCe()
			} else {
				ce = execute_data.GetThis().GetValue().GetCe()
			}
		}
		call_info = 0<<16 | 0<<17
	}
	call = ZendVmStackPushCallFrame(call_info, fbc, opline.GetExtendedValue(), ce)
	call.SetPrevExecuteData(execute_data.GetCall())
	execute_data.SetCall(call)
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_INIT_USER_CALL_SPEC_CONST_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var function_name *Zval
	var fcc ZendFcallInfoCache
	var error *byte = nil
	var func_ *ZendFunction
	var object_or_called_scope any
	var call *ZendExecuteData
	var call_info uint32 = 0<<16 | 0<<17 | 1<<25
	function_name = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	if ZendIsCallableEx(function_name, nil, 0, nil, &fcc, &error) != 0 {
		func_ = fcc.GetFunctionHandler()
		if error != nil {
			_efree(error)

			/* This is the only soft error is_callable() can generate */

			ZendNonStaticMethodCall(func_)
			if EG.GetException() != nil {
				ZvalPtrDtorNogc(free_op2)
				return 0
			}
		}
		object_or_called_scope = fcc.GetCalledScope()
		if (func_.GetFnFlags() & 1 << 20) != 0 {

			/* Delay closure destruction until its invocation */

			ZendGcAddref(&((*ZendObject)((*byte)(func_ - g.SizeOf("zend_object")))).gc)
			call_info |= 1 << 22
			if (func_.GetFnFlags() & 1 << 21) != 0 {
				call_info |= 1 << 23
			}
			if fcc.GetObject() != nil {
				object_or_called_scope = fcc.GetObject()
				call_info |= 8 | 1<<0<<8 | 1<<1<<8
			}
		} else if fcc.GetObject() != nil {
			ZendGcAddref(&(fcc.GetObject()).gc)
			object_or_called_scope = fcc.GetObject()
			call_info |= 1<<21 | (8 | 1<<0<<8 | 1<<1<<8)
		}
		ZvalPtrDtorNogc(free_op2)
		if ((1<<1|1<<2)&(1<<1|1<<2)) != 0 && EG.GetException() != nil {
			if (call_info & 1 << 22) != 0 {
				ZendObjectRelease((*ZendObject)((*byte)(func_ - g.SizeOf("zend_object"))))
			} else if (call_info & 1 << 21) != 0 {
				ZendObjectRelease(fcc.GetObject())
			}
			return 0
		}
		if func_.GetType() == 2 && !(g.CondF((uintptr_t(&func_.op_array).run_time_cache__ptr&1) != 0, func() any {
			return *((*any)((*byte)(CG.GetMapPtrBase() + uintptr_t(&func_.op_array).run_time_cache__ptr - 1)))
		}, func() any { return any(*(&func_.op_array.run_time_cache__ptr)) })) {
			InitFuncRunTimeCache(&func_.op_array)
		}
	} else {
		ZendInternalTypeError((execute_data.GetFunc().GetFnFlags()&1<<31) != 0, "%s() expects parameter 1 to be a valid callback, %s", (*Zval)((*byte)(opline)+int32(opline.GetOp1()).constant).GetValue().GetStr().GetVal(), error)
		_efree(error)
		ZvalPtrDtorNogc(free_op2)
		if EG.GetException() != nil {
			return 0
		}
		func_ = (*ZendFunction)(&ZendPassFunction)
		object_or_called_scope = nil
	}
	call = ZendVmStackPushCallFrame(call_info, func_, opline.GetExtendedValue(), object_or_called_scope)
	call.SetPrevExecuteData(execute_data.GetCall())
	execute_data.SetCall(call)
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_ADD_ARRAY_ELEMENT_SPEC_CONST_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var expr_ptr *Zval
	var new_expr Zval
	if (1<<0 == 1<<2 || 1<<0 == 1<<3) && (opline.GetExtendedValue()&1<<0) != 0 {
		expr_ptr = nil
		if expr_ptr.GetType() == 10 {
			ZvalAddrefP(expr_ptr)
		} else {
			var _z *Zval = expr_ptr
			var _ref *ZendReference = (*ZendReference)(_emalloc(g.SizeOf("zend_reference")))
			ZendGcSetRefcount(&_ref.gc, 2)
			_ref.GetGc().SetTypeInfo(10)
			var _z1 *Zval = &_ref.val
			var _z2 *Zval = _z
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			_ref.GetSources().SetPtr(nil)
			_z.GetValue().SetRef(_ref)
			_z.SetTypeInfo(10 | 1<<0<<8)
		}
	} else {
		expr_ptr = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
		if 1<<0 == 1<<1 {

		} else if 1<<0 == 1<<0 {
			if expr_ptr.GetTypeFlags() != 0 {
				ZvalAddrefP(expr_ptr)
			}
		} else if 1<<0 == 1<<3 {
			if expr_ptr.GetType() == 10 {
				expr_ptr = &(*expr_ptr).value.GetRef().GetVal()
			}
			if expr_ptr.GetTypeFlags() != 0 {
				ZvalAddrefP(expr_ptr)
			}
		} else {
			if expr_ptr.GetType() == 10 {
				var ref *ZendRefcounted = expr_ptr.GetValue().GetCounted()
				expr_ptr = &(*expr_ptr).value.GetRef().GetVal()
				if ZendGcDelref(&ref.gc) == 0 {
					var _z1 *Zval = &new_expr
					var _z2 *Zval = expr_ptr
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					expr_ptr = &new_expr
					_efree(ref)
				} else if (expr_ptr.GetTypeInfo() & 0xff00) != 0 {
					ZvalAddrefP(expr_ptr)
				}
			}
		}
	}
	if (1<<1 | 1<<2) != 0 {
		var free_op2 ZendFreeOp
		var offset *Zval = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
		var str *ZendString
		var hval ZendUlong
	add_again:
		if offset.GetType() == 6 {
			str = offset.GetValue().GetStr()
			if (1<<1 | 1<<2) != 1<<0 {
				if _zendHandleNumericStr(str.GetVal(), str.GetLen(), &hval) != 0 {
					goto num_index
				}
			}
		str_index:
			ZendHashUpdate((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())).GetValue().GetArr(), str, expr_ptr)
		} else if offset.GetType() == 4 {
			hval = offset.GetValue().GetLval()
		num_index:
			ZendHashIndexUpdate((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())).GetValue().GetArr(), hval, expr_ptr)
		} else if ((1<<1|1<<2)&(1<<2|1<<3)) != 0 && offset.GetType() == 10 {
			offset = &(*offset).value.GetRef().GetVal()
			goto add_again
		} else if offset.GetType() == 1 {
			str = ZendEmptyString
			goto str_index
		} else if offset.GetType() == 5 {
			hval = ZendDvalToLval(offset.GetValue().GetDval())
			goto num_index
		} else if offset.GetType() == 2 {
			hval = 0
			goto num_index
		} else if offset.GetType() == 3 {
			hval = 1
			goto num_index
		} else if offset.GetType() == 9 {
			ZendUseResourceAsOffset(offset)
			hval = offset.GetValue().GetRes().GetHandle()
			goto num_index
		} else if (1<<1|1<<2) == 1<<3 && offset.GetType() == 0 {
			_zvalUndefinedOp2(execute_data)
			str = ZendEmptyString
			goto str_index
		} else {
			ZendIllegalOffset()
			ZvalPtrDtorNogc(expr_ptr)
		}
		ZvalPtrDtorNogc(free_op2)
	} else {
		if ZendHashNextIndexInsert((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())).GetValue().GetArr(), expr_ptr) == nil {
			ZendCannotAddElement()
			ZvalPtrDtorNogc(expr_ptr)
		}
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
