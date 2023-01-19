// <<generate>>

package zend

import (
	g "sik/runtime/grammar"
)

// Source: <Zend/zend_operators.h>

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

// #define ZEND_OPERATORS_H

// # include < errno . h >

// # include < math . h >

// # include < assert . h >

// # include < stddef . h >

// # include "zend_portability.h"

// # include "zend_strtod.h"

// # include "zend_multiply.h"

// # include "zend_object_handlers.h"

// #define LONG_SIGN_MASK       ZEND_LONG_MIN

/**
 * Checks whether the string "str" with length "length" is numeric. The value
 * of allow_errors determines whether it's required to be entirely numeric, or
 * just its prefix. Leading whitespace is allowed.
 *
 * The function returns 0 if the string did not contain a valid number; IS_LONG
 * if it contained a number that fits within the range of a long; or IS_DOUBLE
 * if the number was out of long range or contained a decimal point/exponent.
 * The number's value is returned into the respective pointer, *lval or *dval,
 * if that pointer is not NULL.
 *
 * This variant also gives information if a string that represents an integer
 * could not be represented as such due to overflow. It writes 1 to oflow_info
 * if the integer is larger than ZEND_LONG_MAX and -1 if it's smaller than ZEND_LONG_MIN.
 */

/* >= as (double)ZEND_LONG_MAX is outside signed range */

// #define ZEND_DOUBLE_FITS_LONG(d) ( ! ( ( d ) >= ( double ) ZEND_LONG_MAX || ( d ) < ( double ) ZEND_LONG_MIN ) )

func ZendDvalToLval(d float64) ZendLong {
	if !(isfinite(d)) || isnan(d) {
		return 0
	} else if d >= float64(INT64_MAX || d < float64(INT64_MIN)) {
		return ZendDvalToLvalSlow(d)
	}
	return ZendLong(d)
}
func ZendDvalToLvalCap(d float64) ZendLong {
	if !(isfinite(d)) || isnan(d) {
		return 0
	} else if d >= float64(INT64_MAX || d < float64(INT64_MIN)) {
		if d > 0 {
			return INT64_MAX
		} else {
			return INT64_MIN
		}
	}
	return ZendLong(d)
}

/* }}} */

// #define ZEND_IS_DIGIT(c) ( ( c ) >= '0' && ( c ) <= '9' )

// #define ZEND_IS_XDIGIT(c) ( ( ( c ) >= 'A' && ( c ) <= 'F' ) || ( ( c ) >= 'a' && ( c ) <= 'f' ) )

func IsNumericStringEx(str *byte, length int, lval *ZendLong, dval *float64, allow_errors int, oflow_info *int) ZendUchar {
	if (*str) > '9' {
		return 0
	}
	return _isNumericStringEx(str, length, lval, dval, allow_errors, oflow_info)
}
func IsNumericString(str *byte, length int, lval *ZendLong, dval *float64, allow_errors int) ZendUchar {
	return IsNumericStringEx(str, length, lval, dval, allow_errors, nil)
}
func ZendMemnstr(haystack *byte, needle string, needle_len int, end *byte) *byte {
	var p *byte = haystack
	var ne byte = needle[needle_len-1]
	var off_p ptrdiff_t
	var off_s int
	if needle_len == 1 {
		return (*byte)(memchr(p, *needle, end-p))
	}
	off_p = end - haystack
	if off_p > 0 {
		off_s = int(off_p)
	} else {
		off_s = 0
	}
	if needle_len > off_s {
		return nil
	}
	if off_s < 1024 || needle_len < 9 {
		end -= needle_len
		for p <= end {
			if g.Assign(&p, (*byte)(memchr(p, *needle, end-p+1))) && ne == p[needle_len-1] {
				if !(memcmp(needle+1, p+1, needle_len-2)) {
					return p
				}
			}
			if p == nil {
				return nil
			}
			p++
		}
		return nil
	} else {
		return ZendMemnstrEx(haystack, needle, needle_len, end)
	}
}
func ZendMemrchr(s any, c int, n int) any {
	var e *uint8
	if 0 == n {
		return nil
	}
	for e = (*uint8)(s + n - 1); e >= (*uint8)(s); e-- {
		if (*e) == uint8(c) {
			return any(e)
		}
	}
	return nil
}
func ZendMemnrstr(haystack *byte, needle *byte, needle_len int, end *byte) *byte {
	var p *byte = end
	var ne byte = needle[needle_len-1]
	var off_p ptrdiff_t
	var off_s int
	if needle_len == 1 {
		return (*byte)(ZendMemrchr(haystack, *needle, p-haystack))
	}
	off_p = end - haystack
	if off_p > 0 {
		off_s = int(off_p)
	} else {
		off_s = 0
	}
	if needle_len > off_s {
		return nil
	}
	if off_s < 1024 || needle_len < 3 {
		p -= needle_len
		for {
			p = (*byte)(ZendMemrchr(haystack, *needle, p-haystack+1))
			if p == nil {
				return nil
			}
			if ne == p[needle_len-1] && !(memcmp(needle+1, p+1, needle_len-2)) {
				return p
			}
			if g.PostDec(&p) < haystack {
				break
			}
		}
		return nil
	} else {
		return ZendMemnrstrEx(haystack, needle, needle_len, end)
	}
}
func ZvalGetLong(op *Zval) ZendLong {
	if op.GetType() == 4 {
		return op.GetValue().GetLval()
	} else {
		return ZvalGetLongFunc(op)
	}
}
func ZvalGetDouble(op *Zval) float64 {
	if op.GetType() == 5 {
		return op.GetValue().GetDval()
	} else {
		return ZvalGetDoubleFunc(op)
	}
}
func ZvalGetString(op *Zval) *ZendString {
	if op.GetType() == 6 {
		return ZendStringCopy(op.GetValue().GetStr())
	} else {
		return ZvalGetStringFunc(op)
	}
}
func ZvalGetTmpString(op *Zval, tmp **ZendString) *ZendString {
	if op.GetType() == 6 {
		*tmp = nil
		return op.GetValue().GetStr()
	} else {
		*tmp = ZvalGetStringFunc(op)
		return *tmp
	}
}
func ZendTmpStringRelease(tmp *ZendString) {
	if tmp != nil {
		ZendStringReleaseEx(tmp, 0)
	}
}

/* Like zval_get_string, but returns NULL if the conversion fails with an exception. */

func ZvalTryGetString(op *Zval) *ZendString {
	if op.GetType() == 6 {
		var ret *ZendString = ZendStringCopy(op.GetValue().GetStr())
		return ret
	} else {
		return ZvalTryGetStringFunc(op)
	}
}

/* Like zval_get_tmp_string, but returns NULL if the conversion fails with an exception. */

func ZvalTryGetTmpString(op *Zval, tmp **ZendString) *ZendString {
	if op.GetType() == 6 {
		var ret *ZendString = op.GetValue().GetStr()
		*tmp = nil
		return ret
	} else {
		*tmp = ZvalTryGetStringFunc(op)
		return *tmp
	}
}

/* Like convert_to_string(), but returns whether the conversion succeeded and does not modify the
 * zval in-place if it fails. */

func TryConvertToString(op *Zval) ZendBool {
	if op.GetType() == 6 {
		return 1
	}
	return _tryConvertToString(op)
}

/* Compatibility macros for 7.2 and below */

// #define _zval_get_long(op) zval_get_long ( op )

// #define _zval_get_double(op) zval_get_double ( op )

// #define _zval_get_string(op) zval_get_string ( op )

// #define _zval_get_long_func(op) zval_get_long_func ( op )

// #define _zval_get_double_func(op) zval_get_double_func ( op )

// #define _zval_get_string_func(op) zval_get_string_func ( op )

// #define convert_to_cstring(op) if ( Z_TYPE_P ( op ) != IS_STRING ) { _convert_to_cstring ( ( op ) ) ; }

// #define convert_to_string(op) if ( Z_TYPE_P ( op ) != IS_STRING ) { _convert_to_string ( ( op ) ) ; }

// #define zval_is_true(op) zend_is_true ( op )

func IZendIsTrue(op *Zval) int {
	var result int = 0
again:
	switch op.GetType() {
	case 3:
		result = 1
		break
	case 4:
		if op.GetValue().GetLval() != 0 {
			result = 1
		}
		break
	case 5:
		if op.GetValue().GetDval() {
			result = 1
		}
		break
	case 6:
		if op.GetValue().GetStr().GetLen() > 1 || op.GetValue().GetStr().GetLen() != 0 && op.GetValue().GetStr().GetVal()[0] != '0' {
			result = 1
		}
		break
	case 7:
		if op.GetValue().GetArr().GetNNumOfElements() != 0 {
			result = 1
		}
		break
	case 8:
		if op.GetValue().GetObj().GetHandlers().GetCastObject() == ZendStdCastObjectTostring {
			result = 1
		} else {
			result = ZendObjectIsTrue(op)
		}
		break
	case 9:
		if op.GetValue().GetRes().GetHandle() != 0 {
			result = 1
		}
		break
	case 10:
		op = &(*op).value.GetRef().GetVal()
		goto again
		break
	default:
		break
	}
	return result
}

// #define zend_string_tolower(str) zend_string_tolower_ex ( str , 0 )

// #define convert_to_ex_master(pzv,lower_type,upper_type) if ( Z_TYPE_P ( pzv ) != upper_type ) { convert_to_ ## lower_type ( pzv ) ; }

// #define convert_to_explicit_type(pzv,type) do { switch ( type ) { case IS_NULL : convert_to_null ( pzv ) ; break ; case IS_LONG : convert_to_long ( pzv ) ; break ; case IS_DOUBLE : convert_to_double ( pzv ) ; break ; case _IS_BOOL : convert_to_boolean ( pzv ) ; break ; case IS_ARRAY : convert_to_array ( pzv ) ; break ; case IS_OBJECT : convert_to_object ( pzv ) ; break ; case IS_STRING : convert_to_string ( pzv ) ; break ; default : assert ( 0 ) ; break ; } } while ( 0 ) ;

// #define convert_to_explicit_type_ex(pzv,str_type) if ( Z_TYPE_P ( pzv ) != str_type ) { convert_to_explicit_type ( pzv , str_type ) ; }

// #define convert_to_boolean_ex(pzv) do { if ( Z_TYPE_INFO_P ( pzv ) > IS_TRUE ) { convert_to_boolean ( pzv ) ; } else if ( Z_TYPE_INFO_P ( pzv ) < IS_FALSE ) { ZVAL_FALSE ( pzv ) ; } } while ( 0 )

// #define convert_to_long_ex(pzv) convert_to_ex_master ( pzv , long , IS_LONG )

// #define convert_to_double_ex(pzv) convert_to_ex_master ( pzv , double , IS_DOUBLE )

// #define convert_to_string_ex(pzv) convert_to_ex_master ( pzv , string , IS_STRING )

// #define convert_to_array_ex(pzv) convert_to_ex_master ( pzv , array , IS_ARRAY )

// #define convert_to_object_ex(pzv) convert_to_ex_master ( pzv , object , IS_OBJECT )

// #define convert_to_null_ex(pzv) convert_to_ex_master ( pzv , null , IS_NULL )

// #define convert_scalar_to_number_ex(pzv) if ( Z_TYPE_P ( pzv ) != IS_LONG && Z_TYPE_P ( pzv ) != IS_DOUBLE ) { convert_scalar_to_number ( pzv ) ; }

// #define zend_update_current_locale()

/* The offset in bytes between the value and type fields of a zval */

// #define ZVAL_OFFSETOF_TYPE       ( offsetof ( zval , u1 . type_info ) - offsetof ( zval , value ) )

// #define ZEND_USE_ASM_ARITHMETIC       1

func FastLongIncrementFunction(op1 *Zval) {
	if op1.GetValue().GetLval() == INT64_MAX {

		/* switch to double */

		var __z *Zval = op1
		__z.GetValue().SetDval(float64(INT64_MAX + 1.0))
		__z.SetTypeInfo(5)
	} else {
		op1.GetValue().GetLval()++
	}
}
func FastLongDecrementFunction(op1 *Zval) {
	if op1.GetValue().GetLval() == INT64_MIN {

		/* switch to double */

		var __z *Zval = op1
		__z.GetValue().SetDval(float64(INT64_MIN - 1.0))
		__z.SetTypeInfo(5)
	} else {
		op1.GetValue().GetLval()--
	}
}
func FastLongAddFunction(result *Zval, op1 *Zval, op2 *Zval) {
	/*
	 * 'result' may alias with op1 or op2, so we need to
	 * ensure that 'result' is not updated until after we
	 * have read the values of op1 and op2.
	 */

	if (op1.GetValue().GetLval()&INT64_MIN) == (op2.GetValue().GetLval()&INT64_MIN) && (op1.GetValue().GetLval()&INT64_MIN) != (op1.GetValue().GetLval()+op2.GetValue().GetLval()&INT64_MIN) {
		var __z *Zval = result
		__z.GetValue().SetDval(float64(op1.GetValue().GetLval() + float64(op2.GetValue().GetLval())))
		__z.SetTypeInfo(5)
	} else {
		var __z *Zval = result
		__z.GetValue().SetLval(op1.GetValue().GetLval() + op2.GetValue().GetLval())
		__z.SetTypeInfo(4)
	}

	/*
	 * 'result' may alias with op1 or op2, so we need to
	 * ensure that 'result' is not updated until after we
	 * have read the values of op1 and op2.
	 */
}
func FastAddFunction(result *Zval, op1 *Zval, op2 *Zval) int {
	if op1.GetType() == 4 {
		if op2.GetType() == 4 {
			FastLongAddFunction(result, op1, op2)
			return SUCCESS
		} else if op2.GetType() == 5 {
			var __z *Zval = result
			__z.GetValue().SetDval(float64(op1.GetValue().GetLval()) + op2.GetValue().GetDval())
			__z.SetTypeInfo(5)
			return SUCCESS
		}
	} else if op1.GetType() == 5 {
		if op2.GetType() == 5 {
			var __z *Zval = result
			__z.GetValue().SetDval(op1.GetValue().GetDval() + op2.GetValue().GetDval())
			__z.SetTypeInfo(5)
			return SUCCESS
		} else if op2.GetType() == 4 {
			var __z *Zval = result
			__z.GetValue().SetDval(op1.GetValue().GetDval() + float64(op2.GetValue().GetLval()))
			__z.SetTypeInfo(5)
			return SUCCESS
		}
	}
	return AddFunction(result, op1, op2)
}
func FastLongSubFunction(result *Zval, op1 *Zval, op2 *Zval) {
	var __z *Zval = result
	__z.GetValue().SetLval(op1.GetValue().GetLval() - op2.GetValue().GetLval())
	__z.SetTypeInfo(4)
	if (op1.GetValue().GetLval()&INT64_MIN) != (op2.GetValue().GetLval()&INT64_MIN) && (op1.GetValue().GetLval()&INT64_MIN) != (result.GetValue().GetLval()&INT64_MIN) {
		var __z *Zval = result
		__z.GetValue().SetDval(float64(op1.GetValue().GetLval() - float64(op2.GetValue().GetLval())))
		__z.SetTypeInfo(5)
	}
}
func FastDivFunction(result *Zval, op1 *Zval, op2 *Zval) int { return DivFunction(result, op1, op2) }
func ZendFastEqualStrings(s1 *ZendString, s2 *ZendString) int {
	if s1 == s2 {
		return 1
	} else if s1.GetVal()[0] > '9' || s2.GetVal()[0] > '9' {
		return ZendStringEqualContent(s1, s2)
	} else {
		return ZendiSmartStreq(s1, s2)
	}
}
func FastEqualCheckFunction(op1 *Zval, op2 *Zval) int {
	var result Zval
	if op1.GetType() == 4 {
		if op2.GetType() == 4 {
			return op1.GetValue().GetLval() == op2.GetValue().GetLval()
		} else if op2.GetType() == 5 {
			return float64(op1.GetValue().GetLval()) == op2.GetValue().GetDval()
		}
	} else if op1.GetType() == 5 {
		if op2.GetType() == 5 {
			return op1.GetValue().GetDval() == op2.GetValue().GetDval()
		} else if op2.GetType() == 4 {
			return op1.GetValue().GetDval() == float64(op2.GetValue().GetLval())
		}
	} else if op1.GetType() == 6 {
		if op2.GetType() == 6 {
			return ZendFastEqualStrings(op1.GetValue().GetStr(), op2.GetValue().GetStr())
		}
	}
	CompareFunction(&result, op1, op2)
	return result.GetValue().GetLval() == 0
}
func FastEqualCheckLong(op1 *Zval, op2 *Zval) int {
	var result Zval
	if op2.GetType() == 4 {
		return op1.GetValue().GetLval() == op2.GetValue().GetLval()
	}
	CompareFunction(&result, op1, op2)
	return result.GetValue().GetLval() == 0
}
func FastEqualCheckString(op1 *Zval, op2 *Zval) int {
	var result Zval
	if op2.GetType() == 6 {
		return ZendFastEqualStrings(op1.GetValue().GetStr(), op2.GetValue().GetStr())
	}
	CompareFunction(&result, op1, op2)
	return result.GetValue().GetLval() == 0
}
func FastIsIdenticalFunction(op1 *Zval, op2 *Zval) ZendBool {
	if op1.GetType() != op2.GetType() {
		return 0
	} else if op1.GetType() <= 3 {
		return 1
	}
	return ZendIsIdentical(op1, op2)
}
func FastIsNotIdenticalFunction(op1 *Zval, op2 *Zval) ZendBool {
	if op1.GetType() != op2.GetType() {
		return 1
	} else if op1.GetType() <= 3 {
		return 0
	}
	return !(ZendIsIdentical(op1, op2))
}

// #define ZEND_TRY_BINARY_OP1_OBJECT_OPERATION(opcode,binary_op) if ( UNEXPECTED ( Z_TYPE_P ( op1 ) == IS_OBJECT ) && op1 == result && UNEXPECTED ( Z_OBJ_HANDLER_P ( op1 , get ) ) && EXPECTED ( Z_OBJ_HANDLER_P ( op1 , set ) ) ) { int ret ; zval rv ; zval * objval = Z_OBJ_HANDLER_P ( op1 , get ) ( op1 , & rv ) ; Z_TRY_ADDREF_P ( objval ) ; ret = binary_op ( objval , objval , op2 ) ; Z_OBJ_HANDLER_P ( op1 , set ) ( op1 , objval ) ; zval_ptr_dtor ( objval ) ; return ret ; } else if ( UNEXPECTED ( Z_TYPE_P ( op1 ) == IS_OBJECT ) && UNEXPECTED ( Z_OBJ_HANDLER_P ( op1 , do_operation ) ) ) { if ( EXPECTED ( SUCCESS == Z_OBJ_HANDLER_P ( op1 , do_operation ) ( opcode , result , op1 , op2 ) ) ) { return SUCCESS ; } }

// #define ZEND_TRY_BINARY_OP2_OBJECT_OPERATION(opcode) if ( UNEXPECTED ( Z_TYPE_P ( op2 ) == IS_OBJECT ) && UNEXPECTED ( Z_OBJ_HANDLER_P ( op2 , do_operation ) ) && EXPECTED ( SUCCESS == Z_OBJ_HANDLER_P ( op2 , do_operation ) ( opcode , result , op1 , op2 ) ) ) { return SUCCESS ; }

// #define ZEND_TRY_BINARY_OBJECT_OPERATION(opcode,binary_op) ZEND_TRY_BINARY_OP1_OBJECT_OPERATION ( opcode , binary_op ) else ZEND_TRY_BINARY_OP2_OBJECT_OPERATION ( opcode )

// #define ZEND_TRY_UNARY_OBJECT_OPERATION(opcode) if ( UNEXPECTED ( Z_TYPE_P ( op1 ) == IS_OBJECT ) && UNEXPECTED ( Z_OBJ_HANDLER_P ( op1 , do_operation ) ) && EXPECTED ( SUCCESS == Z_OBJ_HANDLER_P ( op1 , do_operation ) ( opcode , result , op1 , NULL ) ) ) { return SUCCESS ; }

/* buf points to the END of the buffer */

func ZendPrintUlongToBuf(buf *byte, num ZendUlong) *byte {
	*buf = '0'
	for {
		*(g.PreDec(&buf)) = byte(num%10 + '0')
		num /= 10
		if num <= 0 {
			break
		}
	}
	return buf
}

/* buf points to the END of the buffer */

func ZendPrintLongToBuf(buf *byte, num ZendLong) *byte {
	if num < 0 {
		var result *byte = ZendPrintUlongToBuf(buf, ^ZendUlong(num)+1)
		*(g.PreDec(&result)) = '-'
		return result
	} else {
		return ZendPrintUlongToBuf(buf, num)
	}
}
func ZendUnwrapReference(op *Zval) {
	if ZvalRefcountP(op) == 1 {
		var _z *Zval = op
		var ref *ZendReference
		assert(_z.GetType() == 10)
		ref = _z.GetValue().GetRef()
		var _z1 *Zval = _z
		var _z2 *Zval = &ref.val
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		_efree(ref)
	} else {
		ZvalDelrefP(op)
		var _z1 *Zval = op
		var _z2 *Zval = &(*op).value.GetRef().GetVal()
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if (_t & 0xff00) != 0 {
			ZendGcAddref(&_gc.gc)
		}
	}
}

/* }}} */

// Source: <Zend/zend_operators.c>

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

// # include < ctype . h >

// # include "zend.h"

// # include "zend_operators.h"

// # include "zend_variables.h"

// # include "zend_globals.h"

// # include "zend_list.h"

// # include "zend_API.h"

// # include "zend_strtod.h"

// # include "zend_exceptions.h"

// # include "zend_closures.h"

// #define zend_tolower(c) tolower ( c )

// #define TYPE_PAIR(t1,t2) ( ( ( t1 ) << 4 ) | ( t2 ) )

var TolowerMap []uint8 = []uint8{0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xf, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f, 0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f, 0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, 0x3e, 0x3f, 0x40, 0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68, 0x69, 0x6a, 0x6b, 0x6c, 0x6d, 0x6e, 0x6f, 0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x76, 0x77, 0x78, 0x79, 0x7a, 0x5b, 0x5c, 0x5d, 0x5e, 0x5f, 0x60, 0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68, 0x69, 0x6a, 0x6b, 0x6c, 0x6d, 0x6e, 0x6f, 0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x76, 0x77, 0x78, 0x79, 0x7a, 0x7b, 0x7c, 0x7d, 0x7e, 0x7f, 0x80, 0x81, 0x82, 0x83, 0x84, 0x85, 0x86, 0x87, 0x88, 0x89, 0x8a, 0x8b, 0x8c, 0x8d, 0x8e, 0x8f, 0x90, 0x91, 0x92, 0x93, 0x94, 0x95, 0x96, 0x97, 0x98, 0x99, 0x9a, 0x9b, 0x9c, 0x9d, 0x9e, 0x9f, 0xa0, 0xa1, 0xa2, 0xa3, 0xa4, 0xa5, 0xa6, 0xa7, 0xa8, 0xa9, 0xaa, 0xab, 0xac, 0xad, 0xae, 0xaf, 0xb0, 0xb1, 0xb2, 0xb3, 0xb4, 0xb5, 0xb6, 0xb7, 0xb8, 0xb9, 0xba, 0xbb, 0xbc, 0xbd, 0xbe, 0xbf, 0xc0, 0xc1, 0xc2, 0xc3, 0xc4, 0xc5, 0xc6, 0xc7, 0xc8, 0xc9, 0xca, 0xcb, 0xcc, 0xcd, 0xce, 0xcf, 0xd0, 0xd1, 0xd2, 0xd3, 0xd4, 0xd5, 0xd6, 0xd7, 0xd8, 0xd9, 0xda, 0xdb, 0xdc, 0xdd, 0xde, 0xdf, 0xe0, 0xe1, 0xe2, 0xe3, 0xe4, 0xe5, 0xe6, 0xe7, 0xe8, 0xe9, 0xea, 0xeb, 0xec, 0xed, 0xee, 0xef, 0xf0, 0xf1, 0xf2, 0xf3, 0xf4, 0xf5, 0xf6, 0xf7, 0xf8, 0xf9, 0xfa, 0xfb, 0xfc, 0xfd, 0xfe, 0xff}

// #define zend_tolower_ascii(c) ( tolower_map [ ( unsigned char ) ( c ) ] )

/**
* Functions using locale lowercase:
         zend_binary_strncasecmp_l
         zend_binary_strcasecmp_l
       zend_binary_zval_strcasecmp
       zend_binary_zval_strncasecmp
       string_compare_function_ex
       string_case_compare_function
* Functions using ascii lowercase:
         zend_str_tolower_copy
       zend_str_tolower_dup
       zend_str_tolower
       zend_binary_strcasecmp
       zend_binary_strncasecmp
*/

func ZendAtoi(str *byte, str_len int) int {
	var retval int
	if str_len == 0 {
		str_len = strlen(str)
	}
	retval = strtoll(str, nil, 0)
	if str_len > 0 {
		switch str[str_len-1] {
		case 'g':

		case 'G':
			retval *= 1024
		case 'm':

		case 'M':
			retval *= 1024
		case 'k':

		case 'K':
			retval *= 1024
			break
		}
	}
	return retval
}

/* }}} */

func ZendAtol(str *byte, str_len int) ZendLong {
	var retval ZendLong
	if str_len == 0 {
		str_len = strlen(str)
	}
	retval = strtoll(str, nil, 0)
	if str_len > 0 {
		switch str[str_len-1] {
		case 'g':

		case 'G':
			retval *= 1024
		case 'm':

		case 'M':
			retval *= 1024
		case 'k':

		case 'K':
			retval *= 1024
			break
		}
	}
	return retval
}

/* }}} */

// #define convert_object_to_type(op,dst,ctype,conv_func) ZVAL_UNDEF ( dst ) ; if ( Z_OBJ_HT_P ( op ) -> cast_object ) { if ( Z_OBJ_HT_P ( op ) -> cast_object ( op , dst , ctype ) == FAILURE ) { zend_error ( E_RECOVERABLE_ERROR , "Object of class %s could not be converted to %s" , ZSTR_VAL ( Z_OBJCE_P ( op ) -> name ) , zend_get_type_by_const ( ctype ) ) ; } } else if ( Z_OBJ_HT_P ( op ) -> get ) { zval * newop = Z_OBJ_HT_P ( op ) -> get ( op , dst ) ; if ( Z_TYPE_P ( newop ) != IS_OBJECT ) { ZVAL_COPY_VALUE ( dst , newop ) ; conv_func ( dst ) ; } }

/* }}} */

func _convertScalarToNumber(op *Zval, silent ZendBool, check ZendBool) {
try_again:
	switch op.GetType() {
	case 10:
		ZendUnwrapReference(op)
		goto try_again
	case 6:
		var str *ZendString
		str = op.GetValue().GetStr()
		if g.Assign(&(op.GetTypeInfo()), IsNumericString(str.GetVal(), str.GetLen(), &(*op).value.GetLval(), &(*op).value.GetDval(), g.Cond(silent != 0, 1, -1))) == 0 {
			var __z *Zval = op
			__z.GetValue().SetLval(0)
			__z.SetTypeInfo(4)
			if silent == 0 {
				ZendError(1<<1, "A non-numeric value encountered")
			}
		}
		ZendStringReleaseEx(str, 0)
		break
	case 1:

	case 2:
		var __z *Zval = op
		__z.GetValue().SetLval(0)
		__z.SetTypeInfo(4)
		break
	case 3:
		var __z *Zval = op
		__z.GetValue().SetLval(1)
		__z.SetTypeInfo(4)
		break
	case 9:
		var l ZendLong = op.GetValue().GetRes().GetHandle()
		ZvalPtrDtor(op)
		var __z *Zval = op
		__z.GetValue().SetLval(l)
		__z.SetTypeInfo(4)
		break
	case 8:
		var dst Zval
		&dst.SetTypeInfo(0)
		if op.GetValue().GetObj().GetHandlers().GetCastObject() != nil {
			if op.GetValue().GetObj().GetHandlers().GetCastObject()(op, &dst, 20) == FAILURE {
				ZendError(1<<12, "Object of class %s could not be converted to %s", op.GetValue().GetObj().GetCe().GetName().GetVal(), ZendGetTypeByConst(20))
			}
		} else if op.GetValue().GetObj().GetHandlers().GetGet() != nil {
			var newop *Zval = op.GetValue().GetObj().GetHandlers().GetGet()(op, &dst)
			if newop.GetType() != 8 {
				var _z1 *Zval = &dst
				var _z2 *Zval = newop
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				ConvertScalarToNumber(&dst)
			}
		}
		if check != 0 && EG.GetException() != nil {
			return
		}
		ZvalPtrDtor(op)
		if dst.GetType() == 4 || dst.GetType() == 5 {
			var _z1 *Zval = op
			var _z2 *Zval = &dst
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
		} else {
			var __z *Zval = op
			__z.GetValue().SetLval(1)
			__z.SetTypeInfo(4)
		}
		break
	}
}

/* }}} */

func ConvertScalarToNumber(op *Zval) { _convertScalarToNumber(op, 1, 0) }

/* }}} */

func _zendiConvertScalarToNumberEx(op *Zval, holder *Zval, silent ZendBool) *Zval {
	switch op.GetType() {
	case 1:

	case 2:
		var __z *Zval = holder
		__z.GetValue().SetLval(0)
		__z.SetTypeInfo(4)
		return holder
	case 3:
		var __z *Zval = holder
		__z.GetValue().SetLval(1)
		__z.SetTypeInfo(4)
		return holder
	case 6:
		if g.Assign(&(holder.GetTypeInfo()), IsNumericString(op.GetValue().GetStr().GetVal(), op.GetValue().GetStr().GetLen(), &(*holder).value.GetLval(), &(*holder).value.GetDval(), g.Cond(silent != 0, 1, -1))) == 0 {
			var __z *Zval = holder
			__z.GetValue().SetLval(0)
			__z.SetTypeInfo(4)
			if silent == 0 {
				ZendError(1<<1, "A non-numeric value encountered")
			}
		}
		return holder
	case 9:
		var __z *Zval = holder
		__z.GetValue().SetLval(op.GetValue().GetRes().GetHandle())
		__z.SetTypeInfo(4)
		return holder
	case 8:
		holder.SetTypeInfo(0)
		if op.GetValue().GetObj().GetHandlers().GetCastObject() != nil {
			if op.GetValue().GetObj().GetHandlers().GetCastObject()(op, holder, 20) == FAILURE {
				ZendError(1<<12, "Object of class %s could not be converted to %s", op.GetValue().GetObj().GetCe().GetName().GetVal(), ZendGetTypeByConst(20))
			}
		} else if op.GetValue().GetObj().GetHandlers().GetGet() != nil {
			var newop *Zval = op.GetValue().GetObj().GetHandlers().GetGet()(op, holder)
			if newop.GetType() != 8 {
				var _z1 *Zval = holder
				var _z2 *Zval = newop
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				ConvertScalarToNumber(holder)
			}
		}
		if EG.GetException() != nil || holder.GetType() != 4 && holder.GetType() != 5 {
			var __z *Zval = holder
			__z.GetValue().SetLval(1)
			__z.SetTypeInfo(4)
		}
		return holder
	case 4:

	case 5:

	default:
		return op
	}
}

/* }}} */

func _zendiConvertScalarToNumber(op *Zval, holder *Zval) *Zval {
	return _zendiConvertScalarToNumberEx(op, holder, 1)
}

/* }}} */

func _zendiConvertScalarToNumberNoisy(op *Zval, holder *Zval) *Zval {
	return _zendiConvertScalarToNumberEx(op, holder, 0)
}

/* }}} */

// #define zendi_convert_scalar_to_number(op,holder,result,silent) ( ( Z_TYPE_P ( op ) == IS_LONG || Z_TYPE_P ( op ) == IS_DOUBLE ) ? ( op ) : ( ( ( op ) == result ) ? ( _convert_scalar_to_number ( ( op ) , silent , 1 ) , ( op ) ) : ( silent ? _zendi_convert_scalar_to_number ( ( op ) , holder ) : _zendi_convert_scalar_to_number_noisy ( ( op ) , holder ) ) ) )

// #define convert_op1_op2_long(op1,op1_lval,op2,op2_lval,result,op,op_func) do { if ( UNEXPECTED ( Z_TYPE_P ( op1 ) != IS_LONG ) ) { if ( Z_ISREF_P ( op1 ) ) { op1 = Z_REFVAL_P ( op1 ) ; if ( Z_TYPE_P ( op1 ) == IS_LONG ) { op1_lval = Z_LVAL_P ( op1 ) ; break ; } } ZEND_TRY_BINARY_OP1_OBJECT_OPERATION ( op , op_func ) ; op1_lval = _zval_get_long_func_noisy ( op1 ) ; if ( UNEXPECTED ( EG ( exception ) ) ) { if ( result != op1 ) { ZVAL_UNDEF ( result ) ; } return FAILURE ; } } else { op1_lval = Z_LVAL_P ( op1 ) ; } } while ( 0 ) ; do { if ( UNEXPECTED ( Z_TYPE_P ( op2 ) != IS_LONG ) ) { if ( Z_ISREF_P ( op2 ) ) { op2 = Z_REFVAL_P ( op2 ) ; if ( Z_TYPE_P ( op2 ) == IS_LONG ) { op2_lval = Z_LVAL_P ( op2 ) ; break ; } } ZEND_TRY_BINARY_OP2_OBJECT_OPERATION ( op ) ; op2_lval = _zval_get_long_func_noisy ( op2 ) ; if ( UNEXPECTED ( EG ( exception ) ) ) { if ( result != op1 ) { ZVAL_UNDEF ( result ) ; } return FAILURE ; } } else { op2_lval = Z_LVAL_P ( op2 ) ; } } while ( 0 ) ;

func ConvertToLong(op *Zval) {
	if op.GetType() != 4 {
		ConvertToLongBase(op, 10)
	}
}

/* }}} */

func ConvertToLongBase(op *Zval, base int) {
	var tmp ZendLong
try_again:
	switch op.GetType() {
	case 1:

	case 2:
		var __z *Zval = op
		__z.GetValue().SetLval(0)
		__z.SetTypeInfo(4)
		break
	case 3:
		var __z *Zval = op
		__z.GetValue().SetLval(1)
		__z.SetTypeInfo(4)
		break
	case 9:
		tmp = op.GetValue().GetRes().GetHandle()
		ZvalPtrDtor(op)
		var __z *Zval = op
		__z.GetValue().SetLval(tmp)
		__z.SetTypeInfo(4)
		break
	case 4:
		break
	case 5:
		var __z *Zval = op
		__z.GetValue().SetLval(ZendDvalToLval(op.GetValue().GetDval()))
		__z.SetTypeInfo(4)
		break
	case 6:
		var str *ZendString = op.GetValue().GetStr()
		if base == 10 {
			var __z *Zval = op
			__z.GetValue().SetLval(ZvalGetLong(op))
			__z.SetTypeInfo(4)
		} else {
			var __z *Zval = op
			__z.GetValue().SetLval(strtoll(str.GetVal(), nil, base))
			__z.SetTypeInfo(4)
		}
		ZendStringReleaseEx(str, 0)
		break
	case 7:
		if op.GetValue().GetArr().GetNNumOfElements() != 0 {
			tmp = 1
		} else {
			tmp = 0
		}
		ZvalPtrDtor(op)
		var __z *Zval = op
		__z.GetValue().SetLval(tmp)
		__z.SetTypeInfo(4)
		break
	case 8:
		var dst Zval
		&dst.SetTypeInfo(0)
		if op.GetValue().GetObj().GetHandlers().GetCastObject() != nil {
			if op.GetValue().GetObj().GetHandlers().GetCastObject()(op, &dst, 4) == FAILURE {
				ZendError(1<<12, "Object of class %s could not be converted to %s", op.GetValue().GetObj().GetCe().GetName().GetVal(), ZendGetTypeByConst(4))
			}
		} else if op.GetValue().GetObj().GetHandlers().GetGet() != nil {
			var newop *Zval = op.GetValue().GetObj().GetHandlers().GetGet()(op, &dst)
			if newop.GetType() != 8 {
				var _z1 *Zval = &dst
				var _z2 *Zval = newop
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				ConvertToLong(&dst)
			}
		}
		ZvalPtrDtor(op)
		if dst.GetType() == 4 {
			var __z *Zval = op
			__z.GetValue().SetLval(dst.GetValue().GetLval())
			__z.SetTypeInfo(4)
		} else {
			var __z *Zval = op
			__z.GetValue().SetLval(1)
			__z.SetTypeInfo(4)
		}
		return
	case 10:
		ZendUnwrapReference(op)
		goto try_again
	default:
		break
	}
}

/* }}} */

func ConvertToDouble(op *Zval) {
	var tmp float64
try_again:
	switch op.GetType() {
	case 1:

	case 2:
		var __z *Zval = op
		__z.GetValue().SetDval(0.0)
		__z.SetTypeInfo(5)
		break
	case 3:
		var __z *Zval = op
		__z.GetValue().SetDval(1.0)
		__z.SetTypeInfo(5)
		break
	case 9:
		var d float64 = float64(op.GetValue().GetRes().GetHandle())
		ZvalPtrDtor(op)
		var __z *Zval = op
		__z.GetValue().SetDval(d)
		__z.SetTypeInfo(5)
		break
	case 4:
		var __z *Zval = op
		__z.GetValue().SetDval(float64(op.GetValue().GetLval()))
		__z.SetTypeInfo(5)
		break
	case 5:
		break
	case 6:
		var str *ZendString = op.GetValue().GetStr()
		var __z *Zval = op
		__z.GetValue().SetDval(ZendStrtod(str.GetVal(), nil))
		__z.SetTypeInfo(5)
		ZendStringReleaseEx(str, 0)
		break
	case 7:
		if op.GetValue().GetArr().GetNNumOfElements() != 0 {
			tmp = 1
		} else {
			tmp = 0
		}
		ZvalPtrDtor(op)
		var __z *Zval = op
		__z.GetValue().SetDval(tmp)
		__z.SetTypeInfo(5)
		break
	case 8:
		var dst Zval
		&dst.SetTypeInfo(0)
		if op.GetValue().GetObj().GetHandlers().GetCastObject() != nil {
			if op.GetValue().GetObj().GetHandlers().GetCastObject()(op, &dst, 5) == FAILURE {
				ZendError(1<<12, "Object of class %s could not be converted to %s", op.GetValue().GetObj().GetCe().GetName().GetVal(), ZendGetTypeByConst(5))
			}
		} else if op.GetValue().GetObj().GetHandlers().GetGet() != nil {
			var newop *Zval = op.GetValue().GetObj().GetHandlers().GetGet()(op, &dst)
			if newop.GetType() != 8 {
				var _z1 *Zval = &dst
				var _z2 *Zval = newop
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				ConvertToDouble(&dst)
			}
		}
		ZvalPtrDtor(op)
		if dst.GetType() == 5 {
			var __z *Zval = op
			__z.GetValue().SetDval(dst.GetValue().GetDval())
			__z.SetTypeInfo(5)
		} else {
			var __z *Zval = op
			__z.GetValue().SetDval(1.0)
			__z.SetTypeInfo(5)
		}
		break
	case 10:
		ZendUnwrapReference(op)
		goto try_again
	default:
		break
	}
}

/* }}} */

func ConvertToNull(op *Zval) {
	ZvalPtrDtor(op)
	op.SetTypeInfo(1)
}

/* }}} */

func ConvertToBoolean(op *Zval) {
	var tmp int
try_again:
	switch op.GetType() {
	case 2:

	case 3:
		break
	case 1:
		op.SetTypeInfo(2)
		break
	case 9:
		var l ZendLong = g.Cond(op.GetValue().GetRes().GetHandle() != 0, 1, 0)
		ZvalPtrDtor(op)
		if l != 0 {
			op.SetTypeInfo(3)
		} else {
			op.SetTypeInfo(2)
		}
		break
	case 4:
		if g.Cond(op.GetValue().GetLval() != 0, 1, 0) {
			op.SetTypeInfo(3)
		} else {
			op.SetTypeInfo(2)
		}
		break
	case 5:
		if g.Cond(op.GetValue().GetDval(), 1, 0) {
			op.SetTypeInfo(3)
		} else {
			op.SetTypeInfo(2)
		}
		break
	case 6:
		var str *ZendString = op.GetValue().GetStr()
		if str.GetLen() == 0 || str.GetLen() == 1 && str.GetVal()[0] == '0' {
			op.SetTypeInfo(2)
		} else {
			op.SetTypeInfo(3)
		}
		ZendStringReleaseEx(str, 0)
		break
	case 7:
		if op.GetValue().GetArr().GetNNumOfElements() != 0 {
			tmp = 1
		} else {
			tmp = 0
		}
		ZvalPtrDtor(op)
		if tmp != 0 {
			op.SetTypeInfo(3)
		} else {
			op.SetTypeInfo(2)
		}
		break
	case 8:
		var dst Zval
		&dst.SetTypeInfo(0)
		if op.GetValue().GetObj().GetHandlers().GetCastObject() != nil {
			if op.GetValue().GetObj().GetHandlers().GetCastObject()(op, &dst, 16) == FAILURE {
				ZendError(1<<12, "Object of class %s could not be converted to %s", op.GetValue().GetObj().GetCe().GetName().GetVal(), ZendGetTypeByConst(16))
			}
		} else if op.GetValue().GetObj().GetHandlers().GetGet() != nil {
			var newop *Zval = op.GetValue().GetObj().GetHandlers().GetGet()(op, &dst)
			if newop.GetType() != 8 {
				var _z1 *Zval = &dst
				var _z2 *Zval = newop
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				ConvertToBoolean(&dst)
			}
		}
		ZvalPtrDtor(op)
		if dst.GetTypeInfo() == 2 || dst.GetTypeInfo() == 3 {
			op.SetTypeInfo(dst.GetTypeInfo())
		} else {
			op.SetTypeInfo(3)
		}
		break
	case 10:
		ZendUnwrapReference(op)
		goto try_again
	default:
		break
	}
}

/* }}} */

func _convertToCstring(op *Zval) {
	if op.GetType() == 5 {
		var str *ZendString
		var dval float64 = op.GetValue().GetDval()
		str = ZendStrpprintfUnchecked(0, "%.*H", int(EG.GetPrecision()), dval)
		var __z *Zval = op
		var __s *ZendString = str
		__z.GetValue().SetStr(__s)
		__z.SetTypeInfo(6 | 1<<0<<8)
	} else {
		_convertToString(op)
	}
}

/* }}} */

func _convertToString(op *Zval) {
try_again:
	switch op.GetType() {
	case 0:

	case 1:

	case 2:
		var __z *Zval = op
		var __s *ZendString = ZendEmptyString
		__z.GetValue().SetStr(__s)
		__z.SetTypeInfo(6)
		break
	case 3:
		var __z *Zval = op
		var __s *ZendString = ZendOneCharString['1']
		__z.GetValue().SetStr(__s)
		__z.SetTypeInfo(6)
		break
	case 6:
		break
	case 9:
		var str *ZendString = ZendStrpprintf(0, "Resource id #"+"%"+"lld", zend_long(*op).value.res.handle)
		ZvalPtrDtor(op)
		var __z *Zval = op
		var __s *ZendString = str
		__z.GetValue().SetStr(__s)
		__z.SetTypeInfo(6 | 1<<0<<8)
		break
	case 4:
		var __z *Zval = op
		var __s *ZendString = ZendLongToStr(op.GetValue().GetLval())
		__z.GetValue().SetStr(__s)
		if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
			__z.SetTypeInfo(6)
		} else {
			__z.SetTypeInfo(6 | 1<<0<<8)
		}
		break
	case 5:
		var str *ZendString
		var dval float64 = op.GetValue().GetDval()
		str = ZendStrpprintf(0, "%.*G", int(EG.GetPrecision()), dval)

		/* %G already handles removing trailing zeros from the fractional part, yay */

		var __z *Zval = op
		var __s *ZendString = str
		__z.GetValue().SetStr(__s)
		__z.SetTypeInfo(6 | 1<<0<<8)
		break
	case 7:
		ZendError(1<<3, "Array to string conversion")
		ZvalPtrDtor(op)
		var __z *Zval = op
		var __s *ZendString = ZendKnownStrings[ZEND_STR_ARRAY_CAPITALIZED]
		__z.GetValue().SetStr(__s)
		__z.SetTypeInfo(6)
		break
	case 8:
		var tmp Zval
		if op.GetValue().GetObj().GetHandlers().GetCastObject() != nil {
			if op.GetValue().GetObj().GetHandlers().GetCastObject()(op, &tmp, 6) == SUCCESS {
				ZvalPtrDtor(op)
				var _z1 *Zval = op
				var _z2 *Zval = &tmp
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				return
			}
		} else if op.GetValue().GetObj().GetHandlers().GetGet() != nil {
			var z *Zval = op.GetValue().GetObj().GetHandlers().GetGet()(op, &tmp)
			if z.GetType() != 8 {
				var str *ZendString = ZvalGetString(z)
				ZvalPtrDtor(z)
				ZvalPtrDtor(op)
				var __z *Zval = op
				var __s *ZendString = str
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
				return
			}
			ZvalPtrDtor(z)
		}
		if EG.GetException() == nil {
			ZendThrowError(nil, "Object of class %s could not be converted to string", op.GetValue().GetObj().GetCe().GetName().GetVal())
		}
		ZvalPtrDtor(op)
		var __z *Zval = op
		var __s *ZendString = ZendEmptyString
		__z.GetValue().SetStr(__s)
		__z.SetTypeInfo(6)
		break
	case 10:
		ZendUnwrapReference(op)
		goto try_again
	default:
		break
	}
}

/* }}} */

func _tryConvertToString(op *Zval) ZendBool {
	var str *ZendString
	assert(op.GetType() != 6)
	str = ZvalTryGetStringFunc(op)
	if str == nil {
		return 0
	}
	ZvalPtrDtor(op)
	var __z *Zval = op
	var __s *ZendString = str
	__z.GetValue().SetStr(__s)
	if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
		__z.SetTypeInfo(6)
	} else {
		__z.SetTypeInfo(6 | 1<<0<<8)
	}
	return 1
}
func ConvertScalarToArray(op *Zval) {
	var ht *HashTable = _zendNewArray(1)
	ZendHashIndexAddNew(ht, 0, op)
	var __arr *ZendArray = ht
	var __z *Zval = op
	__z.GetValue().SetArr(__arr)
	__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
}

/* }}} */

func ConvertToArray(op *Zval) {
try_again:
	switch op.GetType() {
	case 7:
		break
	case 8:
		if op.GetValue().GetObj().GetCe() == ZendCeClosure {
			ConvertScalarToArray(op)
		} else {
			var obj_ht *HashTable = ZendGetPropertiesFor(op, ZEND_PROP_PURPOSE_ARRAY_CAST)
			if obj_ht != nil {
				var new_obj_ht *HashTable = ZendProptableToSymtable(obj_ht, op.GetValue().GetObj().GetCe().GetDefaultPropertiesCount() != 0 || op.GetValue().GetObj().GetHandlers() != &StdObjectHandlers || (ZvalGcFlags(obj_ht.GetGc().GetTypeInfo())&1<<5) != 0)
				ZvalPtrDtor(op)
				var __arr *ZendArray = new_obj_ht
				var __z *Zval = op
				__z.GetValue().SetArr(__arr)
				__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
				if obj_ht != nil && (ZvalGcFlags(obj_ht.GetGc().GetTypeInfo())&1<<6) == 0 && ZendGcDelref(&obj_ht.gc) == 0 {
					ZendArrayDestroy(obj_ht)
				}
			} else {
				ZvalPtrDtor(op)

				/*ZVAL_EMPTY_ARRAY(op);*/

				var __arr *ZendArray = _zendNewArray(0)
				var __z *Zval = op
				__z.GetValue().SetArr(__arr)
				__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)

				/*ZVAL_EMPTY_ARRAY(op);*/

			}
		}
		break
	case 1:

		/*ZVAL_EMPTY_ARRAY(op);*/

		var __arr *ZendArray = _zendNewArray(0)
		var __z *Zval = op
		__z.GetValue().SetArr(__arr)
		__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
		break
	case 10:
		ZendUnwrapReference(op)
		goto try_again
	default:
		ConvertScalarToArray(op)
		break
	}
}

/* }}} */

func ConvertToObject(op *Zval) {
try_again:
	switch op.GetType() {
	case 7:
		var ht *HashTable = ZendSymtableToProptable(op.GetValue().GetArr())
		var obj *ZendObject
		if (ZvalGcFlags(ht.GetGc().GetTypeInfo()) & 1 << 6) != 0 {

			/* TODO: try not to duplicate immutable arrays as well ??? */

			ht = ZendArrayDup(ht)

			/* TODO: try not to duplicate immutable arrays as well ??? */

		} else if ht != op.GetValue().GetArr() {
			ZvalPtrDtor(op)
		} else {
			ZendGcDelref(&ht.gc)
		}
		obj = ZendObjectsNew(ZendStandardClassDef)
		obj.SetProperties(ht)
		var __z *Zval = op
		__z.GetValue().SetObj(obj)
		__z.SetTypeInfo(8 | 1<<0<<8 | 1<<1<<8)
		break
	case 8:
		break
	case 1:
		ObjectInit(op)
		break
	case 10:
		ZendUnwrapReference(op)
		goto try_again
	default:
		var tmp Zval
		var _z1 *Zval = &tmp
		var _z2 *Zval = op
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		ObjectInit(op)
		ZendHashAddNew(op.GetValue().GetObj().GetHandlers().GetGetProperties()(&(*op)), ZendKnownStrings[ZEND_STR_SCALAR], &tmp)
		break
	}
}

/* }}} */

func MultiConvertToLongEx(argc int, _ ...any) {
	var arg *Zval
	var ap va_list
	va_start(ap, argc)
	for g.PostDec(&argc) {
		arg = __va_arg(ap, (*Zval)(_))
		if arg.GetType() != 4 {
			ConvertToLong(arg)
		}
	}
	va_end(ap)
}

/* }}} */

func MultiConvertToDoubleEx(argc int, _ ...any) {
	var arg *Zval
	var ap va_list
	va_start(ap, argc)
	for g.PostDec(&argc) {
		arg = __va_arg(ap, (*Zval)(_))
		if arg.GetType() != 5 {
			ConvertToDouble(arg)
		}
	}
	va_end(ap)
}

/* }}} */

func MultiConvertToStringEx(argc int, _ ...any) {
	var arg *Zval
	var ap va_list
	va_start(ap, argc)
	for g.PostDec(&argc) {
		arg = __va_arg(ap, (*Zval)(_))
		if arg.GetType() != 6 {
			if arg.GetType() != 6 {
				_convertToString(arg)
			}
		}
	}
	va_end(ap)
}

/* }}} */

func _zvalGetLongFuncEx(op *Zval, silent ZendBool) ZendLong {
try_again:
	switch op.GetType() {
	case 0:

	case 1:

	case 2:
		return 0
	case 3:
		return 1
	case 9:
		return op.GetValue().GetRes().GetHandle()
	case 4:
		return op.GetValue().GetLval()
	case 5:
		return ZendDvalToLval(op.GetValue().GetDval())
	case 6:
		var type_ ZendUchar
		var lval ZendLong
		var dval float64
		if 0 == g.Assign(&type_, IsNumericString(op.GetValue().GetStr().GetVal(), op.GetValue().GetStr().GetLen(), &lval, &dval, g.Cond(silent != 0, 1, -1))) {
			if silent == 0 {
				ZendError(1<<1, "A non-numeric value encountered")
			}
			return 0
		} else if type_ == 4 {
			return lval
		} else {

			/* Previously we used strtol here, not is_numeric_string,
			 * and strtol gives you LONG_MAX/_MIN on overflow.
			 * We use use saturating conversion to emulate strtol()'s
			 * behaviour.
			 */

			return ZendDvalToLvalCap(dval)

			/* Previously we used strtol here, not is_numeric_string,
			 * and strtol gives you LONG_MAX/_MIN on overflow.
			 * We use use saturating conversion to emulate strtol()'s
			 * behaviour.
			 */

		}
	case 7:
		if op.GetValue().GetArr().GetNNumOfElements() != 0 {
			return 1
		} else {
			return 0
		}
	case 8:
		var dst Zval
		&dst.SetTypeInfo(0)
		if op.GetValue().GetObj().GetHandlers().GetCastObject() != nil {
			if op.GetValue().GetObj().GetHandlers().GetCastObject()(op, &dst, 4) == FAILURE {
				ZendError(1<<12, "Object of class %s could not be converted to %s", op.GetValue().GetObj().GetCe().GetName().GetVal(), ZendGetTypeByConst(4))
			}
		} else if op.GetValue().GetObj().GetHandlers().GetGet() != nil {
			var newop *Zval = op.GetValue().GetObj().GetHandlers().GetGet()(op, &dst)
			if newop.GetType() != 8 {
				var _z1 *Zval = &dst
				var _z2 *Zval = newop
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				ConvertToLong(&dst)
			}
		}
		if dst.GetType() == 4 {
			return dst.GetValue().GetLval()
		} else {
			return 1
		}
	case 10:
		op = &(*op).value.GetRef().GetVal()
		goto try_again
	default:
		break
	}
	return 0
}

/* }}} */

func ZvalGetLongFunc(op *Zval) ZendLong { return _zvalGetLongFuncEx(op, 1) }

/* }}} */

func _zvalGetLongFuncNoisy(op *Zval) ZendLong { return _zvalGetLongFuncEx(op, 0) }

/* }}} */

func ZvalGetDoubleFunc(op *Zval) float64 {
try_again:
	switch op.GetType() {
	case 1:

	case 2:
		return 0.0
	case 3:
		return 1.0
	case 9:
		return float64(op.GetValue().GetRes().GetHandle())
	case 4:
		return float64(op.GetValue().GetLval())
	case 5:
		return op.GetValue().GetDval()
	case 6:
		return ZendStrtod(op.GetValue().GetStr().GetVal(), nil)
	case 7:
		if op.GetValue().GetArr().GetNNumOfElements() != 0 {
			return 1.0
		} else {
			return 0.0
		}
	case 8:
		var dst Zval
		&dst.SetTypeInfo(0)
		if op.GetValue().GetObj().GetHandlers().GetCastObject() != nil {
			if op.GetValue().GetObj().GetHandlers().GetCastObject()(op, &dst, 5) == FAILURE {
				ZendError(1<<12, "Object of class %s could not be converted to %s", op.GetValue().GetObj().GetCe().GetName().GetVal(), ZendGetTypeByConst(5))
			}
		} else if op.GetValue().GetObj().GetHandlers().GetGet() != nil {
			var newop *Zval = op.GetValue().GetObj().GetHandlers().GetGet()(op, &dst)
			if newop.GetType() != 8 {
				var _z1 *Zval = &dst
				var _z2 *Zval = newop
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				ConvertToDouble(&dst)
			}
		}
		if dst.GetType() == 5 {
			return dst.GetValue().GetDval()
		} else {
			return 1.0
		}
	case 10:
		op = &(*op).value.GetRef().GetVal()
		goto try_again
	default:
		break
	}
	return 0.0
}

/* }}} */

func __zvalGetStringFunc(op *Zval, try ZendBool) *ZendString {
try_again:
	switch op.GetType() {
	case 0:

	case 1:

	case 2:
		return ZendEmptyString
	case 3:
		return ZendOneCharString['1']
	case 9:
		return ZendStrpprintf(0, "Resource id #"+"%"+"lld", zend_long(*op).value.res.handle)
	case 4:
		return ZendLongToStr(op.GetValue().GetLval())
	case 5:
		return ZendStrpprintf(0, "%.*G", int(EG.GetPrecision()), op.GetValue().GetDval())
	case 7:
		ZendError(1<<3, "Array to string conversion")
		if try != 0 && EG.GetException() != nil {
			return nil
		} else {
			return ZendKnownStrings[ZEND_STR_ARRAY_CAPITALIZED]
		}
	case 8:
		var tmp Zval
		if op.GetValue().GetObj().GetHandlers().GetCastObject() != nil {
			if op.GetValue().GetObj().GetHandlers().GetCastObject()(op, &tmp, 6) == SUCCESS {
				return tmp.GetValue().GetStr()
			}
		} else if op.GetValue().GetObj().GetHandlers().GetGet() != nil {
			var z *Zval = op.GetValue().GetObj().GetHandlers().GetGet()(op, &tmp)
			if z.GetType() != 8 {
				var str *ZendString = g.CondF(try != 0, func() *ZendString { return ZvalTryGetString(z) }, func() *ZendString { return ZvalGetString(z) })
				ZvalPtrDtor(z)
				return str
			}
			ZvalPtrDtor(z)
		}
		if EG.GetException() == nil {
			ZendThrowError(nil, "Object of class %s could not be converted to string", op.GetValue().GetObj().GetCe().GetName().GetVal())
		}
		if try != 0 {
			return nil
		} else {
			return ZendEmptyString
		}
	case 10:
		op = &(*op).value.GetRef().GetVal()
		goto try_again
	case 6:
		return ZendStringCopy(op.GetValue().GetStr())
	default:
		break
	}
	return nil
}

/* }}} */

func ZvalGetStringFunc(op *Zval) *ZendString { return __zvalGetStringFunc(op, 0) }

/* }}} */

func ZvalTryGetStringFunc(op *Zval) *ZendString { return __zvalGetStringFunc(op, 1) }

/* }}} */

func AddFunctionArray(result *Zval, op1 *Zval, op2 *Zval) {
	if result == op1 && op1.GetValue().GetArr() == op2.GetValue().GetArr() {

		/* $a += $a */

		return

		/* $a += $a */

	}
	if result != op1 {
		var __arr *ZendArray = ZendArrayDup(op1.GetValue().GetArr())
		var __z *Zval = result
		__z.GetValue().SetArr(__arr)
		__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
	} else {
		var _zv *Zval = result
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
	ZendHashMerge(result.GetValue().GetArr(), op2.GetValue().GetArr(), ZvalAddRef, 0)
}

/* }}} */

func AddFunctionFast(result *Zval, op1 *Zval, op2 *Zval) int {
	var type_pair ZendUchar = op1.GetType()<<4 | op2.GetType()
	if type_pair == (4<<4 | 4) {
		FastLongAddFunction(result, op1, op2)
		return SUCCESS
	} else if type_pair == (5<<4 | 5) {
		var __z *Zval = result
		__z.GetValue().SetDval(op1.GetValue().GetDval() + op2.GetValue().GetDval())
		__z.SetTypeInfo(5)
		return SUCCESS
	} else if type_pair == (4<<4 | 5) {
		var __z *Zval = result
		__z.GetValue().SetDval(float64(op1.GetValue().GetLval()) + op2.GetValue().GetDval())
		__z.SetTypeInfo(5)
		return SUCCESS
	} else if type_pair == (5<<4 | 4) {
		var __z *Zval = result
		__z.GetValue().SetDval(op1.GetValue().GetDval() + float64(op2.GetValue().GetLval()))
		__z.SetTypeInfo(5)
		return SUCCESS
	} else if type_pair == (7<<4 | 7) {
		AddFunctionArray(result, op1, op2)
		return SUCCESS
	} else {
		return FAILURE
	}
}
func AddFunctionSlow(result *Zval, op1 *Zval, op2 *Zval) int {
	var op1_copy Zval
	var op2_copy Zval
	var converted int = 0
	for true {
		if op1.GetType() == 10 {
			op1 = &(*op1).value.GetRef().GetVal()
		} else if op2.GetType() == 10 {
			op2 = &(*op2).value.GetRef().GetVal()
		} else if converted == 0 {
			if op1.GetType() == 8 && op1 == result && op1.GetValue().GetObj().GetHandlers().GetGet() != nil && op1.GetValue().GetObj().GetHandlers().GetSet() != nil {
				var ret int
				var rv Zval
				var objval *Zval = op1.GetValue().GetObj().GetHandlers().GetGet()(op1, &rv)
				if objval.GetTypeFlags() != 0 {
					ZvalAddrefP(objval)
				}
				ret = AddFunction(objval, objval, op2)
				op1.GetValue().GetObj().GetHandlers().GetSet()(op1, objval)
				ZvalPtrDtor(objval)
				return ret
			} else if op1.GetType() == 8 && op1.GetValue().GetObj().GetHandlers().GetDoOperation() != nil {
				if SUCCESS == op1.GetValue().GetObj().GetHandlers().GetDoOperation()(1, result, op1, op2) {
					return SUCCESS
				}
			} else if op2.GetType() == 8 && op2.GetValue().GetObj().GetHandlers().GetDoOperation() != nil && SUCCESS == op2.GetValue().GetObj().GetHandlers().GetDoOperation()(1, result, op1, op2) {
				return SUCCESS
			}
			if op1 != op2 {
				if op1.GetType() == 4 || op1.GetType() == 5 {
					op1 = op1
				} else {
					if op1 == result {
						_convertScalarToNumber(op1, 0, 1)
						op1 = op1
					} else {
						op1 = _zendiConvertScalarToNumberNoisy(op1, &op1_copy)
					}
				}
				if op2.GetType() == 4 || op2.GetType() == 5 {
					op2 = op2
				} else {
					if op2 == result {
						_convertScalarToNumber(op2, 0, 1)
						op2 = op2
					} else {
						op2 = _zendiConvertScalarToNumberNoisy(op2, &op2_copy)
					}
				}
			} else {
				if op1.GetType() == 4 || op1.GetType() == 5 {
					op1 = op1
				} else {
					if op1 == result {
						_convertScalarToNumber(op1, 0, 1)
						op1 = op1
					} else {
						op1 = _zendiConvertScalarToNumberNoisy(op1, &op1_copy)
					}
				}
				op2 = op1
			}
			if EG.GetException() != nil {
				if result != op1 {
					result.SetTypeInfo(0)
				}
				return FAILURE
			}
			converted = 1
		} else {
			if result != op1 {
				result.SetTypeInfo(0)
			}
			ZendThrowError(nil, "Unsupported operand types")
			return FAILURE
		}
		if AddFunctionFast(result, op1, op2) == SUCCESS {
			return SUCCESS
		}
	}
}
func AddFunction(result *Zval, op1 *Zval, op2 *Zval) int {
	if AddFunctionFast(result, op1, op2) == SUCCESS {
		return SUCCESS
	} else {
		return AddFunctionSlow(result, op1, op2)
	}
}

/* }}} */

func SubFunctionFast(result *Zval, op1 *Zval, op2 *Zval) int {
	var type_pair ZendUchar = op1.GetType()<<4 | op2.GetType()
	if type_pair == (4<<4 | 4) {
		FastLongSubFunction(result, op1, op2)
		return SUCCESS
	} else if type_pair == (5<<4 | 5) {
		var __z *Zval = result
		__z.GetValue().SetDval(op1.GetValue().GetDval() - op2.GetValue().GetDval())
		__z.SetTypeInfo(5)
		return SUCCESS
	} else if type_pair == (4<<4 | 5) {
		var __z *Zval = result
		__z.GetValue().SetDval(float64(op1.GetValue().GetLval()) - op2.GetValue().GetDval())
		__z.SetTypeInfo(5)
		return SUCCESS
	} else if type_pair == (5<<4 | 4) {
		var __z *Zval = result
		__z.GetValue().SetDval(op1.GetValue().GetDval() - float64(op2.GetValue().GetLval()))
		__z.SetTypeInfo(5)
		return SUCCESS
	} else {
		return FAILURE
	}
}

/* }}} */

func SubFunctionSlow(result *Zval, op1 *Zval, op2 *Zval) int {
	var op1_copy Zval
	var op2_copy Zval
	var converted int = 0
	for true {
		if op1.GetType() == 10 {
			op1 = &(*op1).value.GetRef().GetVal()
		} else if op2.GetType() == 10 {
			op2 = &(*op2).value.GetRef().GetVal()
		} else if converted == 0 {
			if op1.GetType() == 8 && op1 == result && op1.GetValue().GetObj().GetHandlers().GetGet() != nil && op1.GetValue().GetObj().GetHandlers().GetSet() != nil {
				var ret int
				var rv Zval
				var objval *Zval = op1.GetValue().GetObj().GetHandlers().GetGet()(op1, &rv)
				if objval.GetTypeFlags() != 0 {
					ZvalAddrefP(objval)
				}
				ret = SubFunction(objval, objval, op2)
				op1.GetValue().GetObj().GetHandlers().GetSet()(op1, objval)
				ZvalPtrDtor(objval)
				return ret
			} else if op1.GetType() == 8 && op1.GetValue().GetObj().GetHandlers().GetDoOperation() != nil {
				if SUCCESS == op1.GetValue().GetObj().GetHandlers().GetDoOperation()(2, result, op1, op2) {
					return SUCCESS
				}
			} else if op2.GetType() == 8 && op2.GetValue().GetObj().GetHandlers().GetDoOperation() != nil && SUCCESS == op2.GetValue().GetObj().GetHandlers().GetDoOperation()(2, result, op1, op2) {
				return SUCCESS
			}
			if op1 != op2 {
				if op1.GetType() == 4 || op1.GetType() == 5 {
					op1 = op1
				} else {
					if op1 == result {
						_convertScalarToNumber(op1, 0, 1)
						op1 = op1
					} else {
						op1 = _zendiConvertScalarToNumberNoisy(op1, &op1_copy)
					}
				}
				if op2.GetType() == 4 || op2.GetType() == 5 {
					op2 = op2
				} else {
					if op2 == result {
						_convertScalarToNumber(op2, 0, 1)
						op2 = op2
					} else {
						op2 = _zendiConvertScalarToNumberNoisy(op2, &op2_copy)
					}
				}
			} else {
				if op1.GetType() == 4 || op1.GetType() == 5 {
					op1 = op1
				} else {
					if op1 == result {
						_convertScalarToNumber(op1, 0, 1)
						op1 = op1
					} else {
						op1 = _zendiConvertScalarToNumberNoisy(op1, &op1_copy)
					}
				}
				op2 = op1
			}
			if EG.GetException() != nil {
				if result != op1 {
					result.SetTypeInfo(0)
				}
				return FAILURE
			}
			converted = 1
		} else {
			if result != op1 {
				result.SetTypeInfo(0)
			}
			ZendThrowError(nil, "Unsupported operand types")
			return FAILURE
		}
		if SubFunctionFast(result, op1, op2) == SUCCESS {
			return SUCCESS
		}
	}
}

/* }}} */

func SubFunction(result *Zval, op1 *Zval, op2 *Zval) int {
	if SubFunctionFast(result, op1, op2) == SUCCESS {
		return SUCCESS
	} else {
		return SubFunctionSlow(result, op1, op2)
	}
}

/* }}} */

func MulFunction(result *Zval, op1 *Zval, op2 *Zval) int {
	var op1_copy Zval
	var op2_copy Zval
	var converted int = 0
	for true {
		var type_pair ZendUchar = op1.GetType()<<4 | op2.GetType()
		if type_pair == (4<<4 | 4) {
			var overflow ZendLong
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
			return SUCCESS
		} else if type_pair == (5<<4 | 5) {
			var __z *Zval = result
			__z.GetValue().SetDval(op1.GetValue().GetDval() * op2.GetValue().GetDval())
			__z.SetTypeInfo(5)
			return SUCCESS
		} else if type_pair == (4<<4 | 5) {
			var __z *Zval = result
			__z.GetValue().SetDval(float64(op1.GetValue().GetLval()) * op2.GetValue().GetDval())
			__z.SetTypeInfo(5)
			return SUCCESS
		} else if type_pair == (5<<4 | 4) {
			var __z *Zval = result
			__z.GetValue().SetDval(op1.GetValue().GetDval() * float64(op2.GetValue().GetLval()))
			__z.SetTypeInfo(5)
			return SUCCESS
		} else {
			if op1.GetType() == 10 {
				op1 = &(*op1).value.GetRef().GetVal()
			} else if op2.GetType() == 10 {
				op2 = &(*op2).value.GetRef().GetVal()
			} else if converted == 0 {
				if op1.GetType() == 8 && op1 == result && op1.GetValue().GetObj().GetHandlers().GetGet() != nil && op1.GetValue().GetObj().GetHandlers().GetSet() != nil {
					var ret int
					var rv Zval
					var objval *Zval = op1.GetValue().GetObj().GetHandlers().GetGet()(op1, &rv)
					if objval.GetTypeFlags() != 0 {
						ZvalAddrefP(objval)
					}
					ret = MulFunction(objval, objval, op2)
					op1.GetValue().GetObj().GetHandlers().GetSet()(op1, objval)
					ZvalPtrDtor(objval)
					return ret
				} else if op1.GetType() == 8 && op1.GetValue().GetObj().GetHandlers().GetDoOperation() != nil {
					if SUCCESS == op1.GetValue().GetObj().GetHandlers().GetDoOperation()(3, result, op1, op2) {
						return SUCCESS
					}
				} else if op2.GetType() == 8 && op2.GetValue().GetObj().GetHandlers().GetDoOperation() != nil && SUCCESS == op2.GetValue().GetObj().GetHandlers().GetDoOperation()(3, result, op1, op2) {
					return SUCCESS
				}
				if op1 != op2 {
					if op1.GetType() == 4 || op1.GetType() == 5 {
						op1 = op1
					} else {
						if op1 == result {
							_convertScalarToNumber(op1, 0, 1)
							op1 = op1
						} else {
							op1 = _zendiConvertScalarToNumberNoisy(op1, &op1_copy)
						}
					}
					if op2.GetType() == 4 || op2.GetType() == 5 {
						op2 = op2
					} else {
						if op2 == result {
							_convertScalarToNumber(op2, 0, 1)
							op2 = op2
						} else {
							op2 = _zendiConvertScalarToNumberNoisy(op2, &op2_copy)
						}
					}
				} else {
					if op1.GetType() == 4 || op1.GetType() == 5 {
						op1 = op1
					} else {
						if op1 == result {
							_convertScalarToNumber(op1, 0, 1)
							op1 = op1
						} else {
							op1 = _zendiConvertScalarToNumberNoisy(op1, &op1_copy)
						}
					}
					op2 = op1
				}
				if EG.GetException() != nil {
					if result != op1 {
						result.SetTypeInfo(0)
					}
					return FAILURE
				}
				converted = 1
			} else {
				if result != op1 {
					result.SetTypeInfo(0)
				}
				ZendThrowError(nil, "Unsupported operand types")
				return FAILURE
			}
		}
	}
}

/* }}} */

func PowFunction(result *Zval, op1 *Zval, op2 *Zval) int {
	var op1_copy Zval
	var op2_copy Zval
	var converted int = 0
	for true {
		var type_pair ZendUchar = op1.GetType()<<4 | op2.GetType()
		if type_pair == (4<<4 | 4) {
			if op2.GetValue().GetLval() >= 0 {
				var l1 ZendLong = 1
				var l2 ZendLong = op1.GetValue().GetLval()
				var i ZendLong = op2.GetValue().GetLval()
				if i == 0 {
					var __z *Zval = result
					__z.GetValue().SetLval(1)
					__z.SetTypeInfo(4)
					return SUCCESS
				} else if l2 == 0 {
					var __z *Zval = result
					__z.GetValue().SetLval(0)
					__z.SetTypeInfo(4)
					return SUCCESS
				}
				for i >= 1 {
					var overflow ZendLong
					var dval float64 = 0.0
					if i%2 != 0 {
						i--
						var __lres long = l1 * l2
						var __dres long__double = long__double(l1 * long__double(l2))
						var __delta long__double = long__double(__lres - __dres)
						if g.Assign(&overflow, __dres+__delta != __dres) {
							dval = __dres
						} else {
							l1 = __lres
						}
						if overflow != 0 {
							var __z *Zval = result
							__z.GetValue().SetDval(dval * pow(l2, i))
							__z.SetTypeInfo(5)
							return SUCCESS
						}
					} else {
						i /= 2
						var __lres long = l2 * l2
						var __dres long__double = long__double(l2 * long__double(l2))
						var __delta long__double = long__double(__lres - __dres)
						if g.Assign(&overflow, __dres+__delta != __dres) {
							dval = __dres
						} else {
							l2 = __lres
						}
						if overflow != 0 {
							var __z *Zval = result
							__z.GetValue().SetDval(float64(l1 * pow(dval, i)))
							__z.SetTypeInfo(5)
							return SUCCESS
						}
					}
				}

				/* i == 0 */

				var __z *Zval = result
				__z.GetValue().SetLval(l1)
				__z.SetTypeInfo(4)
			} else {
				var __z *Zval = result
				__z.GetValue().SetDval(pow(float64(op1.GetValue().GetLval()), float64(op2.GetValue().GetLval())))
				__z.SetTypeInfo(5)
			}
			return SUCCESS
		} else if type_pair == (5<<4 | 5) {
			var __z *Zval = result
			__z.GetValue().SetDval(pow(op1.GetValue().GetDval(), op2.GetValue().GetDval()))
			__z.SetTypeInfo(5)
			return SUCCESS
		} else if type_pair == (4<<4 | 5) {
			var __z *Zval = result
			__z.GetValue().SetDval(pow(float64(op1.GetValue().GetLval()), op2.GetValue().GetDval()))
			__z.SetTypeInfo(5)
			return SUCCESS
		} else if type_pair == (5<<4 | 4) {
			var __z *Zval = result
			__z.GetValue().SetDval(pow(op1.GetValue().GetDval(), float64(op2.GetValue().GetLval())))
			__z.SetTypeInfo(5)
			return SUCCESS
		} else {
			if op1.GetType() == 10 {
				op1 = &(*op1).value.GetRef().GetVal()
			} else if op2.GetType() == 10 {
				op2 = &(*op2).value.GetRef().GetVal()
			} else if converted == 0 {
				if op1.GetType() == 8 && op1 == result && op1.GetValue().GetObj().GetHandlers().GetGet() != nil && op1.GetValue().GetObj().GetHandlers().GetSet() != nil {
					var ret int
					var rv Zval
					var objval *Zval = op1.GetValue().GetObj().GetHandlers().GetGet()(op1, &rv)
					if objval.GetTypeFlags() != 0 {
						ZvalAddrefP(objval)
					}
					ret = PowFunction(objval, objval, op2)
					op1.GetValue().GetObj().GetHandlers().GetSet()(op1, objval)
					ZvalPtrDtor(objval)
					return ret
				} else if op1.GetType() == 8 && op1.GetValue().GetObj().GetHandlers().GetDoOperation() != nil {
					if SUCCESS == op1.GetValue().GetObj().GetHandlers().GetDoOperation()(12, result, op1, op2) {
						return SUCCESS
					}
				} else if op2.GetType() == 8 && op2.GetValue().GetObj().GetHandlers().GetDoOperation() != nil && SUCCESS == op2.GetValue().GetObj().GetHandlers().GetDoOperation()(12, result, op1, op2) {
					return SUCCESS
				}
				if op1 != op2 {
					if op1.GetType() == 7 {
						if op1 == result {
							ZvalPtrDtor(result)
						}
						var __z *Zval = result
						__z.GetValue().SetLval(0)
						__z.SetTypeInfo(4)
						return SUCCESS
					} else {
						if op1.GetType() == 4 || op1.GetType() == 5 {
							op1 = op1
						} else {
							if op1 == result {
								_convertScalarToNumber(op1, 0, 1)
								op1 = op1
							} else {
								op1 = _zendiConvertScalarToNumberNoisy(op1, &op1_copy)
							}
						}
					}
					if op2.GetType() == 7 {
						if op1 == result {
							ZvalPtrDtor(result)
						}
						var __z *Zval = result
						__z.GetValue().SetLval(1)
						__z.SetTypeInfo(4)
						return SUCCESS
					} else {
						if op2.GetType() == 4 || op2.GetType() == 5 {
							op2 = op2
						} else {
							if op2 == result {
								_convertScalarToNumber(op2, 0, 1)
								op2 = op2
							} else {
								op2 = _zendiConvertScalarToNumberNoisy(op2, &op2_copy)
							}
						}
					}
				} else {
					if op1.GetType() == 7 {
						if op1 == result {
							ZvalPtrDtor(result)
						}
						var __z *Zval = result
						__z.GetValue().SetLval(0)
						__z.SetTypeInfo(4)
						return SUCCESS
					} else {
						if op1.GetType() == 4 || op1.GetType() == 5 {
							op1 = op1
						} else {
							if op1 == result {
								_convertScalarToNumber(op1, 0, 1)
								op1 = op1
							} else {
								op1 = _zendiConvertScalarToNumberNoisy(op1, &op1_copy)
							}
						}
					}
					op2 = op1
				}
				if EG.GetException() != nil {
					if result != op1 {
						result.SetTypeInfo(0)
					}
					return FAILURE
				}
				converted = 1
			} else {
				if result != op1 {
					result.SetTypeInfo(0)
				}
				ZendThrowError(nil, "Unsupported operand types")
				return FAILURE
			}
		}
	}
}

/* }}} */

func DivFunction(result *Zval, op1 *Zval, op2 *Zval) int {
	var op1_copy Zval
	var op2_copy Zval
	var converted int = 0
	for true {
		var type_pair ZendUchar = op1.GetType()<<4 | op2.GetType()
		if type_pair == (4<<4 | 4) {
			if op2.GetValue().GetLval() == 0 {
				ZendError(1<<1, "Division by zero")
				var __z *Zval = result
				__z.GetValue().SetDval(float64(op1.GetValue().GetLval() / float64(op2.GetValue().GetLval())))
				__z.SetTypeInfo(5)
				return SUCCESS
			} else if op2.GetValue().GetLval() == -1 && op1.GetValue().GetLval() == INT64_MIN {

				/* Prevent overflow error/crash */

				var __z *Zval = result
				__z.GetValue().SetDval(float64(INT64_MIN / -1))
				__z.SetTypeInfo(5)
				return SUCCESS
			}
			if op1.GetValue().GetLval()%op2.GetValue().GetLval() == 0 {
				var __z *Zval = result
				__z.GetValue().SetLval(op1.GetValue().GetLval() / op2.GetValue().GetLval())
				__z.SetTypeInfo(4)
			} else {
				var __z *Zval = result
				__z.GetValue().SetDval(float64(op1.GetValue().GetLval()) / op2.GetValue().GetLval())
				__z.SetTypeInfo(5)
			}
			return SUCCESS
		} else if type_pair == (5<<4 | 5) {
			if op2.GetValue().GetDval() == 0 {
				ZendError(1<<1, "Division by zero")
			}
			var __z *Zval = result
			__z.GetValue().SetDval(op1.GetValue().GetDval() / op2.GetValue().GetDval())
			__z.SetTypeInfo(5)
			return SUCCESS
		} else if type_pair == (5<<4 | 4) {
			if op2.GetValue().GetLval() == 0 {
				ZendError(1<<1, "Division by zero")
			}
			var __z *Zval = result
			__z.GetValue().SetDval(op1.GetValue().GetDval() / float64(op2.GetValue().GetLval()))
			__z.SetTypeInfo(5)
			return SUCCESS
		} else if type_pair == (4<<4 | 5) {
			if op2.GetValue().GetDval() == 0 {
				ZendError(1<<1, "Division by zero")
			}
			var __z *Zval = result
			__z.GetValue().SetDval(float64(op1.GetValue().GetLval() / op2.GetValue().GetDval()))
			__z.SetTypeInfo(5)
			return SUCCESS
		} else {
			if op1.GetType() == 10 {
				op1 = &(*op1).value.GetRef().GetVal()
			} else if op2.GetType() == 10 {
				op2 = &(*op2).value.GetRef().GetVal()
			} else if converted == 0 {
				if op1.GetType() == 8 && op1 == result && op1.GetValue().GetObj().GetHandlers().GetGet() != nil && op1.GetValue().GetObj().GetHandlers().GetSet() != nil {
					var ret int
					var rv Zval
					var objval *Zval = op1.GetValue().GetObj().GetHandlers().GetGet()(op1, &rv)
					if objval.GetTypeFlags() != 0 {
						ZvalAddrefP(objval)
					}
					ret = DivFunction(objval, objval, op2)
					op1.GetValue().GetObj().GetHandlers().GetSet()(op1, objval)
					ZvalPtrDtor(objval)
					return ret
				} else if op1.GetType() == 8 && op1.GetValue().GetObj().GetHandlers().GetDoOperation() != nil {
					if SUCCESS == op1.GetValue().GetObj().GetHandlers().GetDoOperation()(4, result, op1, op2) {
						return SUCCESS
					}
				} else if op2.GetType() == 8 && op2.GetValue().GetObj().GetHandlers().GetDoOperation() != nil && SUCCESS == op2.GetValue().GetObj().GetHandlers().GetDoOperation()(4, result, op1, op2) {
					return SUCCESS
				}
				if op1 != op2 {
					if op1.GetType() == 4 || op1.GetType() == 5 {
						op1 = op1
					} else {
						if op1 == result {
							_convertScalarToNumber(op1, 0, 1)
							op1 = op1
						} else {
							op1 = _zendiConvertScalarToNumberNoisy(op1, &op1_copy)
						}
					}
					if op2.GetType() == 4 || op2.GetType() == 5 {
						op2 = op2
					} else {
						if op2 == result {
							_convertScalarToNumber(op2, 0, 1)
							op2 = op2
						} else {
							op2 = _zendiConvertScalarToNumberNoisy(op2, &op2_copy)
						}
					}
				} else {
					if op1.GetType() == 4 || op1.GetType() == 5 {
						op1 = op1
					} else {
						if op1 == result {
							_convertScalarToNumber(op1, 0, 1)
							op1 = op1
						} else {
							op1 = _zendiConvertScalarToNumberNoisy(op1, &op1_copy)
						}
					}
					op2 = op1
				}
				if EG.GetException() != nil {
					if result != op1 {
						result.SetTypeInfo(0)
					}
					return FAILURE
				}
				converted = 1
			} else {
				if result != op1 {
					result.SetTypeInfo(0)
				}
				ZendThrowError(nil, "Unsupported operand types")
				return FAILURE
			}
		}
	}
}

/* }}} */

func ModFunction(result *Zval, op1 *Zval, op2 *Zval) int {
	var op1_lval ZendLong
	var op2_lval ZendLong
	for {
		if op1.GetType() != 4 {
			if op1.GetType() == 10 {
				op1 = &(*op1).value.GetRef().GetVal()
				if op1.GetType() == 4 {
					op1_lval = op1.GetValue().GetLval()
					break
				}
			}
			if op1.GetType() == 8 && op1 == result && op1.GetValue().GetObj().GetHandlers().GetGet() != nil && op1.GetValue().GetObj().GetHandlers().GetSet() != nil {
				var ret int
				var rv Zval
				var objval *Zval = op1.GetValue().GetObj().GetHandlers().GetGet()(op1, &rv)
				if objval.GetTypeFlags() != 0 {
					ZvalAddrefP(objval)
				}
				ret = ModFunction(objval, objval, op2)
				op1.GetValue().GetObj().GetHandlers().GetSet()(op1, objval)
				ZvalPtrDtor(objval)
				return ret
			} else if op1.GetType() == 8 && op1.GetValue().GetObj().GetHandlers().GetDoOperation() != nil {
				if SUCCESS == op1.GetValue().GetObj().GetHandlers().GetDoOperation()(5, result, op1, op2) {
					return SUCCESS
				}
			}
			op1_lval = _zvalGetLongFuncNoisy(op1)
			if EG.GetException() != nil {
				if result != op1 {
					result.SetTypeInfo(0)
				}
				return FAILURE
			}
		} else {
			op1_lval = op1.GetValue().GetLval()
		}
		break
	}
	for {
		if op2.GetType() != 4 {
			if op2.GetType() == 10 {
				op2 = &(*op2).value.GetRef().GetVal()
				if op2.GetType() == 4 {
					op2_lval = op2.GetValue().GetLval()
					break
				}
			}
			if op2.GetType() == 8 && op2.GetValue().GetObj().GetHandlers().GetDoOperation() != nil && SUCCESS == op2.GetValue().GetObj().GetHandlers().GetDoOperation()(5, result, op1, op2) {
				return SUCCESS
			}
			op2_lval = _zvalGetLongFuncNoisy(op2)
			if EG.GetException() != nil {
				if result != op1 {
					result.SetTypeInfo(0)
				}
				return FAILURE
			}
		} else {
			op2_lval = op2.GetValue().GetLval()
		}
		break
	}
	if op2_lval == 0 {

		/* modulus by zero */

		if EG.GetCurrentExecuteData() != nil && CG.GetInCompilation() == 0 {
			ZendThrowExceptionEx(ZendCeDivisionByZeroError, 0, "Modulo by zero")
		} else {
			ZendErrorNoreturn(1<<0, "Modulo by zero")
		}
		if op1 != result {
			result.SetTypeInfo(0)
		}
		return FAILURE
	}
	if op1 == result {
		ZvalPtrDtor(result)
	}
	if op2_lval == -1 {

		/* Prevent overflow error/crash if op1==LONG_MIN */

		var __z *Zval = result
		__z.GetValue().SetLval(0)
		__z.SetTypeInfo(4)
		return SUCCESS
	}
	var __z *Zval = result
	__z.GetValue().SetLval(op1_lval % op2_lval)
	__z.SetTypeInfo(4)
	return SUCCESS
}

/* }}} */

func BooleanXorFunction(result *Zval, op1 *Zval, op2 *Zval) int {
	var op1_val int
	var op2_val int
	for {
		if op1.GetType() == 2 {
			op1_val = 0
		} else if op1.GetType() == 3 {
			op1_val = 1
		} else {
			if op1.GetType() == 10 {
				op1 = &(*op1).value.GetRef().GetVal()
				if op1.GetType() == 2 {
					op1_val = 0
					break
				} else if op1.GetType() == 3 {
					op1_val = 1
					break
				}
			}
			if op1.GetType() == 8 && op1 == result && op1.GetValue().GetObj().GetHandlers().GetGet() != nil && op1.GetValue().GetObj().GetHandlers().GetSet() != nil {
				var ret int
				var rv Zval
				var objval *Zval = op1.GetValue().GetObj().GetHandlers().GetGet()(op1, &rv)
				if objval.GetTypeFlags() != 0 {
					ZvalAddrefP(objval)
				}
				ret = BooleanXorFunction(objval, objval, op2)
				op1.GetValue().GetObj().GetHandlers().GetSet()(op1, objval)
				ZvalPtrDtor(objval)
				return ret
			} else if op1.GetType() == 8 && op1.GetValue().GetObj().GetHandlers().GetDoOperation() != nil {
				if SUCCESS == op1.GetValue().GetObj().GetHandlers().GetDoOperation()(15, result, op1, op2) {
					return SUCCESS
				}
			}
			op1_val = ZendIsTrue(op1)
		}
		break
	}
	for {
		if op2.GetType() == 2 {
			op2_val = 0
		} else if op2.GetType() == 3 {
			op2_val = 1
		} else {
			if op2.GetType() == 10 {
				op2 = &(*op2).value.GetRef().GetVal()
				if op2.GetType() == 2 {
					op2_val = 0
					break
				} else if op2.GetType() == 3 {
					op2_val = 1
					break
				}
			}
			if op2.GetType() == 8 && op2.GetValue().GetObj().GetHandlers().GetDoOperation() != nil && SUCCESS == op2.GetValue().GetObj().GetHandlers().GetDoOperation()(15, result, op1, op2) {
				return SUCCESS
			}
			op2_val = ZendIsTrue(op2)
		}
		break
	}
	if (op1_val ^ op2_val) != 0 {
		result.SetTypeInfo(3)
	} else {
		result.SetTypeInfo(2)
	}
	return SUCCESS
}

/* }}} */

func BooleanNotFunction(result *Zval, op1 *Zval) int {
	if op1.GetType() < 3 {
		result.SetTypeInfo(3)
	} else if op1.GetType() == 3 {
		result.SetTypeInfo(2)
	} else {
		if op1.GetType() == 10 {
			op1 = &(*op1).value.GetRef().GetVal()
			if op1.GetType() < 3 {
				result.SetTypeInfo(3)
				return SUCCESS
			} else if op1.GetType() == 3 {
				result.SetTypeInfo(2)
				return SUCCESS
			}
		}
		if op1.GetType() == 8 && op1.GetValue().GetObj().GetHandlers().GetDoOperation() != nil && SUCCESS == op1.GetValue().GetObj().GetHandlers().GetDoOperation()(14, result, op1, nil) {
			return SUCCESS
		}
		if ZendIsTrue(op1) == 0 {
			result.SetTypeInfo(3)
		} else {
			result.SetTypeInfo(2)
		}
	}
	return SUCCESS
}

/* }}} */

func BitwiseNotFunction(result *Zval, op1 *Zval) int {
try_again:
	switch op1.GetType() {
	case 4:
		var __z *Zval = result
		__z.GetValue().SetLval(^(op1.GetValue().GetLval()))
		__z.SetTypeInfo(4)
		return SUCCESS
	case 5:
		var __z *Zval = result
		__z.GetValue().SetLval(^(ZendDvalToLval(op1.GetValue().GetDval())))
		__z.SetTypeInfo(4)
		return SUCCESS
	case 6:
		var i int
		if op1.GetValue().GetStr().GetLen() == 1 {
			var not ZendUchar = ZendUchar(^((*(op1.GetValue().GetStr())).val))
			var __z *Zval = result
			var __s *ZendString = ZendOneCharString[not]
			__z.GetValue().SetStr(__s)
			__z.SetTypeInfo(6)
		} else {
			var __z *Zval = result
			var __s *ZendString = ZendStringAlloc(op1.GetValue().GetStr().GetLen(), 0)
			__z.GetValue().SetStr(__s)
			__z.SetTypeInfo(6 | 1<<0<<8)
			for i = 0; i < op1.GetValue().GetStr().GetLen(); i++ {
				result.GetValue().GetStr().GetVal()[i] = ^(op1.GetValue().GetStr().GetVal()[i])
			}
			result.GetValue().GetStr().GetVal()[i] = 0
		}
		return SUCCESS
	case 10:
		op1 = &(*op1).value.GetRef().GetVal()
		goto try_again
	default:
		if op1.GetType() == 8 && op1.GetValue().GetObj().GetHandlers().GetDoOperation() != nil && SUCCESS == op1.GetValue().GetObj().GetHandlers().GetDoOperation()(13, result, op1, nil) {
			return SUCCESS
		}
		if result != op1 {
			result.SetTypeInfo(0)
		}
		ZendThrowError(nil, "Unsupported operand types")
		return FAILURE
	}
}

/* }}} */

func BitwiseOrFunction(result *Zval, op1 *Zval, op2 *Zval) int {
	var op1_lval ZendLong
	var op2_lval ZendLong
	if op1.GetType() == 4 && op2.GetType() == 4 {
		var __z *Zval = result
		__z.GetValue().SetLval(op1.GetValue().GetLval() | op2.GetValue().GetLval())
		__z.SetTypeInfo(4)
		return SUCCESS
	}
	if op1.GetType() == 10 {
		op1 = &(*op1).value.GetRef().GetVal()
	}
	if op2.GetType() == 10 {
		op2 = &(*op2).value.GetRef().GetVal()
	}
	if op1.GetType() == 6 && op2.GetType() == 6 {
		var longer *Zval
		var shorter *Zval
		var str *ZendString
		var i int
		if op1.GetValue().GetStr().GetLen() >= op2.GetValue().GetStr().GetLen() {
			if op1.GetValue().GetStr().GetLen() == op2.GetValue().GetStr().GetLen() && op1.GetValue().GetStr().GetLen() == 1 {
				var or ZendUchar = zend_uchar((*(op1.GetValue().GetStr())).val | (*(op2.GetValue().GetStr())).val)
				if result == op1 {
					ZvalPtrDtorStr(result)
				}
				var __z *Zval = result
				var __s *ZendString = ZendOneCharString[or]
				__z.GetValue().SetStr(__s)
				__z.SetTypeInfo(6)
				return SUCCESS
			}
			longer = op1
			shorter = op2
		} else {
			longer = op2
			shorter = op1
		}
		str = ZendStringAlloc(longer.GetValue().GetStr().GetLen(), 0)
		for i = 0; i < shorter.GetValue().GetStr().GetLen(); i++ {
			str.GetVal()[i] = longer.GetValue().GetStr().GetVal()[i] | shorter.GetValue().GetStr().GetVal()[i]
		}
		memcpy(str.GetVal()+i, longer.GetValue().GetStr().GetVal()+i, longer.GetValue().GetStr().GetLen()-i+1)
		if result == op1 {
			ZvalPtrDtorStr(result)
		}
		var __z *Zval = result
		var __s *ZendString = str
		__z.GetValue().SetStr(__s)
		__z.SetTypeInfo(6 | 1<<0<<8)
		return SUCCESS
	}
	if op1.GetType() != 4 {
		if op1.GetType() == 8 && op1 == result && op1.GetValue().GetObj().GetHandlers().GetGet() != nil && op1.GetValue().GetObj().GetHandlers().GetSet() != nil {
			var ret int
			var rv Zval
			var objval *Zval = op1.GetValue().GetObj().GetHandlers().GetGet()(op1, &rv)
			if objval.GetTypeFlags() != 0 {
				ZvalAddrefP(objval)
			}
			ret = BitwiseOrFunction(objval, objval, op2)
			op1.GetValue().GetObj().GetHandlers().GetSet()(op1, objval)
			ZvalPtrDtor(objval)
			return ret
		} else if op1.GetType() == 8 && op1.GetValue().GetObj().GetHandlers().GetDoOperation() != nil {
			if SUCCESS == op1.GetValue().GetObj().GetHandlers().GetDoOperation()(9, result, op1, op2) {
				return SUCCESS
			}
		}
		op1_lval = _zvalGetLongFuncNoisy(op1)
		if EG.GetException() != nil {
			if result != op1 {
				result.SetTypeInfo(0)
			}
			return FAILURE
		}
	} else {
		op1_lval = op1.GetValue().GetLval()
	}
	if op2.GetType() != 4 {
		if op2.GetType() == 8 && op2.GetValue().GetObj().GetHandlers().GetDoOperation() != nil && SUCCESS == op2.GetValue().GetObj().GetHandlers().GetDoOperation()(9, result, op1, op2) {
			return SUCCESS
		}
		op2_lval = _zvalGetLongFuncNoisy(op2)
		if EG.GetException() != nil {
			if result != op1 {
				result.SetTypeInfo(0)
			}
			return FAILURE
		}
	} else {
		op2_lval = op2.GetValue().GetLval()
	}
	if op1 == result {
		ZvalPtrDtor(result)
	}
	var __z *Zval = result
	__z.GetValue().SetLval(op1_lval | op2_lval)
	__z.SetTypeInfo(4)
	return SUCCESS
}

/* }}} */

func BitwiseAndFunction(result *Zval, op1 *Zval, op2 *Zval) int {
	var op1_lval ZendLong
	var op2_lval ZendLong
	if op1.GetType() == 4 && op2.GetType() == 4 {
		var __z *Zval = result
		__z.GetValue().SetLval(op1.GetValue().GetLval() & op2.GetValue().GetLval())
		__z.SetTypeInfo(4)
		return SUCCESS
	}
	if op1.GetType() == 10 {
		op1 = &(*op1).value.GetRef().GetVal()
	}
	if op2.GetType() == 10 {
		op2 = &(*op2).value.GetRef().GetVal()
	}
	if op1.GetType() == 6 && op2.GetType() == 6 {
		var longer *Zval
		var shorter *Zval
		var str *ZendString
		var i int
		if op1.GetValue().GetStr().GetLen() >= op2.GetValue().GetStr().GetLen() {
			if op1.GetValue().GetStr().GetLen() == op2.GetValue().GetStr().GetLen() && op1.GetValue().GetStr().GetLen() == 1 {
				var and ZendUchar = zend_uchar((*(op1.GetValue().GetStr())).val & (*(op2.GetValue().GetStr())).val)
				if result == op1 {
					ZvalPtrDtorStr(result)
				}
				var __z *Zval = result
				var __s *ZendString = ZendOneCharString[and]
				__z.GetValue().SetStr(__s)
				__z.SetTypeInfo(6)
				return SUCCESS
			}
			longer = op1
			shorter = op2
		} else {
			longer = op2
			shorter = op1
		}
		str = ZendStringAlloc(shorter.GetValue().GetStr().GetLen(), 0)
		for i = 0; i < shorter.GetValue().GetStr().GetLen(); i++ {
			str.GetVal()[i] = shorter.GetValue().GetStr().GetVal()[i] & longer.GetValue().GetStr().GetVal()[i]
		}
		str.GetVal()[i] = 0
		if result == op1 {
			ZvalPtrDtorStr(result)
		}
		var __z *Zval = result
		var __s *ZendString = str
		__z.GetValue().SetStr(__s)
		__z.SetTypeInfo(6 | 1<<0<<8)
		return SUCCESS
	}
	if op1.GetType() != 4 {
		if op1.GetType() == 8 && op1 == result && op1.GetValue().GetObj().GetHandlers().GetGet() != nil && op1.GetValue().GetObj().GetHandlers().GetSet() != nil {
			var ret int
			var rv Zval
			var objval *Zval = op1.GetValue().GetObj().GetHandlers().GetGet()(op1, &rv)
			if objval.GetTypeFlags() != 0 {
				ZvalAddrefP(objval)
			}
			ret = BitwiseAndFunction(objval, objval, op2)
			op1.GetValue().GetObj().GetHandlers().GetSet()(op1, objval)
			ZvalPtrDtor(objval)
			return ret
		} else if op1.GetType() == 8 && op1.GetValue().GetObj().GetHandlers().GetDoOperation() != nil {
			if SUCCESS == op1.GetValue().GetObj().GetHandlers().GetDoOperation()(10, result, op1, op2) {
				return SUCCESS
			}
		}
		op1_lval = _zvalGetLongFuncNoisy(op1)
		if EG.GetException() != nil {
			if result != op1 {
				result.SetTypeInfo(0)
			}
			return FAILURE
		}
	} else {
		op1_lval = op1.GetValue().GetLval()
	}
	if op2.GetType() != 4 {
		if op2.GetType() == 8 && op2.GetValue().GetObj().GetHandlers().GetDoOperation() != nil && SUCCESS == op2.GetValue().GetObj().GetHandlers().GetDoOperation()(10, result, op1, op2) {
			return SUCCESS
		}
		op2_lval = _zvalGetLongFuncNoisy(op2)
		if EG.GetException() != nil {
			if result != op1 {
				result.SetTypeInfo(0)
			}
			return FAILURE
		}
	} else {
		op2_lval = op2.GetValue().GetLval()
	}
	if op1 == result {
		ZvalPtrDtor(result)
	}
	var __z *Zval = result
	__z.GetValue().SetLval(op1_lval & op2_lval)
	__z.SetTypeInfo(4)
	return SUCCESS
}

/* }}} */

func BitwiseXorFunction(result *Zval, op1 *Zval, op2 *Zval) int {
	var op1_lval ZendLong
	var op2_lval ZendLong
	if op1.GetType() == 4 && op2.GetType() == 4 {
		var __z *Zval = result
		__z.GetValue().SetLval(op1.GetValue().GetLval() ^ op2.GetValue().GetLval())
		__z.SetTypeInfo(4)
		return SUCCESS
	}
	if op1.GetType() == 10 {
		op1 = &(*op1).value.GetRef().GetVal()
	}
	if op2.GetType() == 10 {
		op2 = &(*op2).value.GetRef().GetVal()
	}
	if op1.GetType() == 6 && op2.GetType() == 6 {
		var longer *Zval
		var shorter *Zval
		var str *ZendString
		var i int
		if op1.GetValue().GetStr().GetLen() >= op2.GetValue().GetStr().GetLen() {
			if op1.GetValue().GetStr().GetLen() == op2.GetValue().GetStr().GetLen() && op1.GetValue().GetStr().GetLen() == 1 {
				var xor ZendUchar = zend_uchar((*(op1.GetValue().GetStr())).val ^ (*(op2.GetValue().GetStr())).val)
				if result == op1 {
					ZvalPtrDtorStr(result)
				}
				var __z *Zval = result
				var __s *ZendString = ZendOneCharString[xor]
				__z.GetValue().SetStr(__s)
				__z.SetTypeInfo(6)
				return SUCCESS
			}
			longer = op1
			shorter = op2
		} else {
			longer = op2
			shorter = op1
		}
		str = ZendStringAlloc(shorter.GetValue().GetStr().GetLen(), 0)
		for i = 0; i < shorter.GetValue().GetStr().GetLen(); i++ {
			str.GetVal()[i] = shorter.GetValue().GetStr().GetVal()[i] ^ longer.GetValue().GetStr().GetVal()[i]
		}
		str.GetVal()[i] = 0
		if result == op1 {
			ZvalPtrDtorStr(result)
		}
		var __z *Zval = result
		var __s *ZendString = str
		__z.GetValue().SetStr(__s)
		__z.SetTypeInfo(6 | 1<<0<<8)
		return SUCCESS
	}
	if op1.GetType() != 4 {
		if op1.GetType() == 8 && op1 == result && op1.GetValue().GetObj().GetHandlers().GetGet() != nil && op1.GetValue().GetObj().GetHandlers().GetSet() != nil {
			var ret int
			var rv Zval
			var objval *Zval = op1.GetValue().GetObj().GetHandlers().GetGet()(op1, &rv)
			if objval.GetTypeFlags() != 0 {
				ZvalAddrefP(objval)
			}
			ret = BitwiseXorFunction(objval, objval, op2)
			op1.GetValue().GetObj().GetHandlers().GetSet()(op1, objval)
			ZvalPtrDtor(objval)
			return ret
		} else if op1.GetType() == 8 && op1.GetValue().GetObj().GetHandlers().GetDoOperation() != nil {
			if SUCCESS == op1.GetValue().GetObj().GetHandlers().GetDoOperation()(11, result, op1, op2) {
				return SUCCESS
			}
		}
		op1_lval = _zvalGetLongFuncNoisy(op1)
		if EG.GetException() != nil {
			if result != op1 {
				result.SetTypeInfo(0)
			}
			return FAILURE
		}
	} else {
		op1_lval = op1.GetValue().GetLval()
	}
	if op2.GetType() != 4 {
		if op2.GetType() == 8 && op2.GetValue().GetObj().GetHandlers().GetDoOperation() != nil && SUCCESS == op2.GetValue().GetObj().GetHandlers().GetDoOperation()(11, result, op1, op2) {
			return SUCCESS
		}
		op2_lval = _zvalGetLongFuncNoisy(op2)
		if EG.GetException() != nil {
			if result != op1 {
				result.SetTypeInfo(0)
			}
			return FAILURE
		}
	} else {
		op2_lval = op2.GetValue().GetLval()
	}
	if op1 == result {
		ZvalPtrDtor(result)
	}
	var __z *Zval = result
	__z.GetValue().SetLval(op1_lval ^ op2_lval)
	__z.SetTypeInfo(4)
	return SUCCESS
}

/* }}} */

func ShiftLeftFunction(result *Zval, op1 *Zval, op2 *Zval) int {
	var op1_lval ZendLong
	var op2_lval ZendLong
	for {
		if op1.GetType() != 4 {
			if op1.GetType() == 10 {
				op1 = &(*op1).value.GetRef().GetVal()
				if op1.GetType() == 4 {
					op1_lval = op1.GetValue().GetLval()
					break
				}
			}
			if op1.GetType() == 8 && op1 == result && op1.GetValue().GetObj().GetHandlers().GetGet() != nil && op1.GetValue().GetObj().GetHandlers().GetSet() != nil {
				var ret int
				var rv Zval
				var objval *Zval = op1.GetValue().GetObj().GetHandlers().GetGet()(op1, &rv)
				if objval.GetTypeFlags() != 0 {
					ZvalAddrefP(objval)
				}
				ret = ShiftLeftFunction(objval, objval, op2)
				op1.GetValue().GetObj().GetHandlers().GetSet()(op1, objval)
				ZvalPtrDtor(objval)
				return ret
			} else if op1.GetType() == 8 && op1.GetValue().GetObj().GetHandlers().GetDoOperation() != nil {
				if SUCCESS == op1.GetValue().GetObj().GetHandlers().GetDoOperation()(6, result, op1, op2) {
					return SUCCESS
				}
			}
			op1_lval = _zvalGetLongFuncNoisy(op1)
			if EG.GetException() != nil {
				if result != op1 {
					result.SetTypeInfo(0)
				}
				return FAILURE
			}
		} else {
			op1_lval = op1.GetValue().GetLval()
		}
		break
	}
	for {
		if op2.GetType() != 4 {
			if op2.GetType() == 10 {
				op2 = &(*op2).value.GetRef().GetVal()
				if op2.GetType() == 4 {
					op2_lval = op2.GetValue().GetLval()
					break
				}
			}
			if op2.GetType() == 8 && op2.GetValue().GetObj().GetHandlers().GetDoOperation() != nil && SUCCESS == op2.GetValue().GetObj().GetHandlers().GetDoOperation()(6, result, op1, op2) {
				return SUCCESS
			}
			op2_lval = _zvalGetLongFuncNoisy(op2)
			if EG.GetException() != nil {
				if result != op1 {
					result.SetTypeInfo(0)
				}
				return FAILURE
			}
		} else {
			op2_lval = op2.GetValue().GetLval()
		}
		break
	}

	/* prevent wrapping quirkiness on some processors where << 64 + x == << x */

	if ZendUlong(op2_lval >= 8*8) != 0 {
		if op2_lval > 0 {
			if op1 == result {
				ZvalPtrDtor(result)
			}
			var __z *Zval = result
			__z.GetValue().SetLval(0)
			__z.SetTypeInfo(4)
			return SUCCESS
		} else {
			if EG.GetCurrentExecuteData() != nil && CG.GetInCompilation() == 0 {
				ZendThrowExceptionEx(ZendCeArithmeticError, 0, "Bit shift by negative number")
			} else {
				ZendErrorNoreturn(1<<0, "Bit shift by negative number")
			}
			if op1 != result {
				result.SetTypeInfo(0)
			}
			return FAILURE
		}
	}
	if op1 == result {
		ZvalPtrDtor(result)
	}

	/* Perform shift on unsigned numbers to get well-defined wrap behavior. */

	var __z *Zval = result
	__z.GetValue().SetLval(zend_long(ZendUlong(op1_lval << op2_lval)))
	__z.SetTypeInfo(4)
	return SUCCESS
}

/* }}} */

func ShiftRightFunction(result *Zval, op1 *Zval, op2 *Zval) int {
	var op1_lval ZendLong
	var op2_lval ZendLong
	for {
		if op1.GetType() != 4 {
			if op1.GetType() == 10 {
				op1 = &(*op1).value.GetRef().GetVal()
				if op1.GetType() == 4 {
					op1_lval = op1.GetValue().GetLval()
					break
				}
			}
			if op1.GetType() == 8 && op1 == result && op1.GetValue().GetObj().GetHandlers().GetGet() != nil && op1.GetValue().GetObj().GetHandlers().GetSet() != nil {
				var ret int
				var rv Zval
				var objval *Zval = op1.GetValue().GetObj().GetHandlers().GetGet()(op1, &rv)
				if objval.GetTypeFlags() != 0 {
					ZvalAddrefP(objval)
				}
				ret = ShiftRightFunction(objval, objval, op2)
				op1.GetValue().GetObj().GetHandlers().GetSet()(op1, objval)
				ZvalPtrDtor(objval)
				return ret
			} else if op1.GetType() == 8 && op1.GetValue().GetObj().GetHandlers().GetDoOperation() != nil {
				if SUCCESS == op1.GetValue().GetObj().GetHandlers().GetDoOperation()(7, result, op1, op2) {
					return SUCCESS
				}
			}
			op1_lval = _zvalGetLongFuncNoisy(op1)
			if EG.GetException() != nil {
				if result != op1 {
					result.SetTypeInfo(0)
				}
				return FAILURE
			}
		} else {
			op1_lval = op1.GetValue().GetLval()
		}
		break
	}
	for {
		if op2.GetType() != 4 {
			if op2.GetType() == 10 {
				op2 = &(*op2).value.GetRef().GetVal()
				if op2.GetType() == 4 {
					op2_lval = op2.GetValue().GetLval()
					break
				}
			}
			if op2.GetType() == 8 && op2.GetValue().GetObj().GetHandlers().GetDoOperation() != nil && SUCCESS == op2.GetValue().GetObj().GetHandlers().GetDoOperation()(7, result, op1, op2) {
				return SUCCESS
			}
			op2_lval = _zvalGetLongFuncNoisy(op2)
			if EG.GetException() != nil {
				if result != op1 {
					result.SetTypeInfo(0)
				}
				return FAILURE
			}
		} else {
			op2_lval = op2.GetValue().GetLval()
		}
		break
	}

	/* prevent wrapping quirkiness on some processors where >> 64 + x == >> x */

	if ZendUlong(op2_lval >= 8*8) != 0 {
		if op2_lval > 0 {
			if op1 == result {
				ZvalPtrDtor(result)
			}
			var __z *Zval = result
			if op1_lval < 0 {
				__z.GetValue().SetLval(-1)
			} else {
				__z.GetValue().SetLval(0)
			}
			__z.SetTypeInfo(4)
			return SUCCESS
		} else {
			if EG.GetCurrentExecuteData() != nil && CG.GetInCompilation() == 0 {
				ZendThrowExceptionEx(ZendCeArithmeticError, 0, "Bit shift by negative number")
			} else {
				ZendErrorNoreturn(1<<0, "Bit shift by negative number")
			}
			if op1 != result {
				result.SetTypeInfo(0)
			}
			return FAILURE
		}
	}
	if op1 == result {
		ZvalPtrDtor(result)
	}
	var __z *Zval = result
	__z.GetValue().SetLval(op1_lval >> op2_lval)
	__z.SetTypeInfo(4)
	return SUCCESS
}

/* }}} */

func ConcatFunction(result *Zval, op1 *Zval, op2 *Zval) int {
	var orig_op1 *Zval = op1
	var op1_copy Zval
	var op2_copy Zval
	&op1_copy.SetTypeInfo(0)
	&op2_copy.SetTypeInfo(0)
	for {
		if op1.GetType() != 6 {
			if op1.GetType() == 10 {
				op1 = &(*op1).value.GetRef().GetVal()
				if op1.GetType() == 6 {
					break
				}
			}
			if op1.GetType() == 8 && op1 == result && op1.GetValue().GetObj().GetHandlers().GetGet() != nil && op1.GetValue().GetObj().GetHandlers().GetSet() != nil {
				var ret int
				var rv Zval
				var objval *Zval = op1.GetValue().GetObj().GetHandlers().GetGet()(op1, &rv)
				if objval.GetTypeFlags() != 0 {
					ZvalAddrefP(objval)
				}
				ret = ConcatFunction(objval, objval, op2)
				op1.GetValue().GetObj().GetHandlers().GetSet()(op1, objval)
				ZvalPtrDtor(objval)
				return ret
			} else if op1.GetType() == 8 && op1.GetValue().GetObj().GetHandlers().GetDoOperation() != nil {
				if SUCCESS == op1.GetValue().GetObj().GetHandlers().GetDoOperation()(8, result, op1, op2) {
					return SUCCESS
				}
			} else if op2.GetType() == 8 && op2.GetValue().GetObj().GetHandlers().GetDoOperation() != nil && SUCCESS == op2.GetValue().GetObj().GetHandlers().GetDoOperation()(8, result, op1, op2) {
				return SUCCESS
			}
			var __z *Zval = &op1_copy
			var __s *ZendString = ZvalGetStringFunc(op1)
			__z.GetValue().SetStr(__s)
			if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
				__z.SetTypeInfo(6)
			} else {
				__z.SetTypeInfo(6 | 1<<0<<8)
			}
			if EG.GetException() != nil {
				ZvalPtrDtorStr(&op1_copy)
				if orig_op1 != result {
					result.SetTypeInfo(0)
				}
				return FAILURE
			}
			if result == op1 {
				if op1 == op2 {
					op2 = &op1_copy
				}
			}
			op1 = &op1_copy
		}
		break
	}
	for {
		if op2.GetType() != 6 {
			if op2.GetType() == 10 {
				op2 = &(*op2).value.GetRef().GetVal()
				if op2.GetType() == 6 {
					break
				}
			}
			if op2.GetType() == 8 && op2.GetValue().GetObj().GetHandlers().GetDoOperation() != nil && SUCCESS == op2.GetValue().GetObj().GetHandlers().GetDoOperation()(8, result, op1, op2) {
				return SUCCESS
			}
			var __z *Zval = &op2_copy
			var __s *ZendString = ZvalGetStringFunc(op2)
			__z.GetValue().SetStr(__s)
			if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
				__z.SetTypeInfo(6)
			} else {
				__z.SetTypeInfo(6 | 1<<0<<8)
			}
			if EG.GetException() != nil {
				ZvalPtrDtorStr(&op1_copy)
				ZvalPtrDtorStr(&op2_copy)
				if orig_op1 != result {
					result.SetTypeInfo(0)
				}
				return FAILURE
			}
			op2 = &op2_copy
		}
		break
	}
	if op1.GetValue().GetStr().GetLen() == 0 {
		if result != op2 {
			if result == orig_op1 {
				IZvalPtrDtor(result)
			}
			var _z1 *Zval = result
			var _z2 *Zval = op2
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if (_t & 0xff00) != 0 {
				ZendGcAddref(&_gc.gc)
			}
		}
	} else if op2.GetValue().GetStr().GetLen() == 0 {
		if result != op1 {
			if result == orig_op1 {
				IZvalPtrDtor(result)
			}
			var _z1 *Zval = result
			var _z2 *Zval = op1
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if (_t & 0xff00) != 0 {
				ZendGcAddref(&_gc.gc)
			}
		}
	} else {
		var op1_len int = op1.GetValue().GetStr().GetLen()
		var op2_len int = op2.GetValue().GetStr().GetLen()
		var result_len int = op1_len + op2_len
		var result_str *ZendString
		if op1_len > SIZE_MAX-(zend_long((*byte)(&((*ZendString)(nil).GetVal()))-(*byte)(nil))+1+8 - 1 & ^(8-1))-op2_len {
			ZendThrowError(nil, "String size overflow")
			ZvalPtrDtorStr(&op1_copy)
			ZvalPtrDtorStr(&op2_copy)
			if orig_op1 != result {
				result.SetTypeInfo(0)
			}
			return FAILURE
		}
		if result == op1 && result.GetTypeFlags() != 0 {

			/* special case, perform operations on result */

			result_str = ZendStringExtend(result.GetValue().GetStr(), result_len, 0)

			/* special case, perform operations on result */

		} else {
			result_str = ZendStringAlloc(result_len, 0)
			memcpy(result_str.GetVal(), op1.GetValue().GetStr().GetVal(), op1_len)
			if result == orig_op1 {
				IZvalPtrDtor(result)
			}
		}

		/* This has to happen first to account for the cases where result == op1 == op2 and
		 * the realloc is done. In this case this line will also update Z_STRVAL_P(op2) to
		 * point to the new string. The first op2_len bytes of result will still be the same. */

		var __z *Zval = result
		var __s *ZendString = result_str
		__z.GetValue().SetStr(__s)
		__z.SetTypeInfo(6 | 1<<0<<8)
		memcpy(result_str.GetVal()+op1_len, op2.GetValue().GetStr().GetVal(), op2_len)
		result_str.GetVal()[result_len] = '0'
	}
	ZvalPtrDtorStr(&op1_copy)
	ZvalPtrDtorStr(&op2_copy)
	return SUCCESS
}

/* }}} */

func StringCompareFunctionEx(op1 *Zval, op2 *Zval, case_insensitive ZendBool) int {
	var tmp_str1 *ZendString
	var tmp_str2 *ZendString
	var str1 *ZendString = ZvalGetTmpString(op1, &tmp_str1)
	var str2 *ZendString = ZvalGetTmpString(op2, &tmp_str2)
	var ret int
	if case_insensitive != 0 {
		ret = ZendBinaryStrcasecmpL(str1.GetVal(), str1.GetLen(), str2.GetVal(), str1.GetLen())
	} else {
		ret = ZendBinaryStrcmp(str1.GetVal(), str1.GetLen(), str2.GetVal(), str2.GetLen())
	}
	ZendTmpStringRelease(tmp_str1)
	ZendTmpStringRelease(tmp_str2)
	return ret
}

/* }}} */

func StringCompareFunction(op1 *Zval, op2 *Zval) int {
	if op1.GetType() == 6 && op2.GetType() == 6 {
		if op1.GetValue().GetStr() == op2.GetValue().GetStr() {
			return 0
		} else {
			return ZendBinaryStrcmp(op1.GetValue().GetStr().GetVal(), op1.GetValue().GetStr().GetLen(), op2.GetValue().GetStr().GetVal(), op2.GetValue().GetStr().GetLen())
		}
	} else {
		var tmp_str1 *ZendString
		var tmp_str2 *ZendString
		var str1 *ZendString = ZvalGetTmpString(op1, &tmp_str1)
		var str2 *ZendString = ZvalGetTmpString(op2, &tmp_str2)
		var ret int = ZendBinaryStrcmp(str1.GetVal(), str1.GetLen(), str2.GetVal(), str2.GetLen())
		ZendTmpStringRelease(tmp_str1)
		ZendTmpStringRelease(tmp_str2)
		return ret
	}
}

/* }}} */

func StringCaseCompareFunction(op1 *Zval, op2 *Zval) int {
	if op1.GetType() == 6 && op2.GetType() == 6 {
		if op1.GetValue().GetStr() == op2.GetValue().GetStr() {
			return 0
		} else {
			return ZendBinaryStrcasecmpL(op1.GetValue().GetStr().GetVal(), op1.GetValue().GetStr().GetLen(), op2.GetValue().GetStr().GetVal(), op2.GetValue().GetStr().GetLen())
		}
	} else {
		var tmp_str1 *ZendString
		var tmp_str2 *ZendString
		var str1 *ZendString = ZvalGetTmpString(op1, &tmp_str1)
		var str2 *ZendString = ZvalGetTmpString(op2, &tmp_str2)
		var ret int = ZendBinaryStrcasecmpL(str1.GetVal(), str1.GetLen(), str2.GetVal(), str1.GetLen())
		ZendTmpStringRelease(tmp_str1)
		ZendTmpStringRelease(tmp_str2)
		return ret
	}
}

/* }}} */

func StringLocaleCompareFunction(op1 *Zval, op2 *Zval) int {
	var tmp_str1 *ZendString
	var tmp_str2 *ZendString
	var str1 *ZendString = ZvalGetTmpString(op1, &tmp_str1)
	var str2 *ZendString = ZvalGetTmpString(op2, &tmp_str2)
	var ret int = strcoll(str1.GetVal(), str2.GetVal())
	ZendTmpStringRelease(tmp_str1)
	ZendTmpStringRelease(tmp_str2)
	return ret
}

/* }}} */

func NumericCompareFunction(op1 *Zval, op2 *Zval) int {
	var d1 float64
	var d2 float64
	d1 = ZvalGetDouble(op1)
	d2 = ZvalGetDouble(op2)
	if d1-d2 != 0 {
		if d1-d2 < 0 {
			return -1
		} else {
			return 1
		}
	} else {
		return 0
	}
}

/* }}} */

func ZendFreeObjGetResult(op *Zval) {
	assert(op.GetTypeFlags() == 0 || ZvalRefcountP(op) != 0)
	ZvalPtrDtor(op)
}

/* }}} */

func ConvertCompareResultToLong(result *Zval) {
	if result.GetType() == 5 {
		var __z *Zval = result
		if result.GetValue().GetDval() {
			if result.GetValue().GetDval() < 0 {
				__z.GetValue().SetLval(-1)
			} else {
				__z.GetValue().SetLval(1)
			}
		} else {
			__z.GetValue().SetLval(0)
		}
		__z.SetTypeInfo(4)
	} else {
		ConvertToLong(result)
	}
}

/* }}} */

func CompareFunction(result *Zval, op1 *Zval, op2 *Zval) int {
	var ret int
	var converted int = 0
	var op1_copy Zval
	var op2_copy Zval
	var op_free *Zval
	var tmp_free Zval
	for true {
		switch op1.GetType()<<4 | op2.GetType() {
		case 4<<4 | 4:
			var __z *Zval = result
			if op1.GetValue().GetLval() > op2.GetValue().GetLval() {
				__z.GetValue().SetLval(1)
			} else {
				if op1.GetValue().GetLval() < op2.GetValue().GetLval() {
					__z.GetValue().SetLval(-1)
				} else {
					__z.GetValue().SetLval(0)
				}
			}
			__z.SetTypeInfo(4)
			return SUCCESS
		case 5<<4 | 4:
			result.GetValue().SetDval(op1.GetValue().GetDval() - float64(op2.GetValue().GetLval()))
			var __z *Zval = result
			if result.GetValue().GetDval() {
				if result.GetValue().GetDval() < 0 {
					__z.GetValue().SetLval(-1)
				} else {
					__z.GetValue().SetLval(1)
				}
			} else {
				__z.GetValue().SetLval(0)
			}
			__z.SetTypeInfo(4)
			return SUCCESS
		case 4<<4 | 5:
			result.GetValue().SetDval(float64(op1.GetValue().GetLval() - op2.GetValue().GetDval()))
			var __z *Zval = result
			if result.GetValue().GetDval() {
				if result.GetValue().GetDval() < 0 {
					__z.GetValue().SetLval(-1)
				} else {
					__z.GetValue().SetLval(1)
				}
			} else {
				__z.GetValue().SetLval(0)
			}
			__z.SetTypeInfo(4)
			return SUCCESS
		case 5<<4 | 5:
			if op1.GetValue().GetDval() == op2.GetValue().GetDval() {
				var __z *Zval = result
				__z.GetValue().SetLval(0)
				__z.SetTypeInfo(4)
			} else {
				result.GetValue().SetDval(op1.GetValue().GetDval() - op2.GetValue().GetDval())
				var __z *Zval = result
				if result.GetValue().GetDval() {
					if result.GetValue().GetDval() < 0 {
						__z.GetValue().SetLval(-1)
					} else {
						__z.GetValue().SetLval(1)
					}
				} else {
					__z.GetValue().SetLval(0)
				}
				__z.SetTypeInfo(4)
			}
			return SUCCESS
		case 7<<4 | 7:
			var __z *Zval = result
			__z.GetValue().SetLval(ZendCompareArrays(op1, op2))
			__z.SetTypeInfo(4)
			return SUCCESS
		case 1<<4 | 1:

		case 1<<4 | 2:

		case 2<<4 | 1:

		case 2<<4 | 2:

		case 3<<4 | 3:
			var __z *Zval = result
			__z.GetValue().SetLval(0)
			__z.SetTypeInfo(4)
			return SUCCESS
		case 1<<4 | 3:
			var __z *Zval = result
			__z.GetValue().SetLval(-1)
			__z.SetTypeInfo(4)
			return SUCCESS
		case 3<<4 | 1:
			var __z *Zval = result
			__z.GetValue().SetLval(1)
			__z.SetTypeInfo(4)
			return SUCCESS
		case 6<<4 | 6:
			if op1.GetValue().GetStr() == op2.GetValue().GetStr() {
				var __z *Zval = result
				__z.GetValue().SetLval(0)
				__z.SetTypeInfo(4)
				return SUCCESS
			}
			var __z *Zval = result
			__z.GetValue().SetLval(ZendiSmartStrcmp(op1.GetValue().GetStr(), op2.GetValue().GetStr()))
			__z.SetTypeInfo(4)
			return SUCCESS
		case 1<<4 | 6:
			var __z *Zval = result
			if op2.GetValue().GetStr().GetLen() == 0 {
				__z.GetValue().SetLval(0)
			} else {
				__z.GetValue().SetLval(-1)
			}
			__z.SetTypeInfo(4)
			return SUCCESS
		case 6<<4 | 1:
			var __z *Zval = result
			if op1.GetValue().GetStr().GetLen() == 0 {
				__z.GetValue().SetLval(0)
			} else {
				__z.GetValue().SetLval(1)
			}
			__z.SetTypeInfo(4)
			return SUCCESS
		case 8<<4 | 1:
			var __z *Zval = result
			__z.GetValue().SetLval(1)
			__z.SetTypeInfo(4)
			return SUCCESS
		case 1<<4 | 8:
			var __z *Zval = result
			__z.GetValue().SetLval(-1)
			__z.SetTypeInfo(4)
			return SUCCESS
		default:
			if op1.GetType() == 10 {
				op1 = &(*op1).value.GetRef().GetVal()
				continue
			} else if op2.GetType() == 10 {
				op2 = &(*op2).value.GetRef().GetVal()
				continue
			}
			if op1.GetType() == 8 && op1.GetValue().GetObj().GetHandlers().GetCompare() != nil {
				ret = op1.GetValue().GetObj().GetHandlers().GetCompare()(result, op1, op2)
				if result.GetType() != 4 {
					ConvertCompareResultToLong(result)
				}
				return ret
			} else if op2.GetType() == 8 && op2.GetValue().GetObj().GetHandlers().GetCompare() != nil {
				ret = op2.GetValue().GetObj().GetHandlers().GetCompare()(result, op1, op2)
				if result.GetType() != 4 {
					ConvertCompareResultToLong(result)
				}
				return ret
			}
			if op1.GetType() == 8 && op2.GetType() == 8 {
				if op1.GetValue().GetObj() == op2.GetValue().GetObj() {

					/* object handles are identical, apparently this is the same object */

					var __z *Zval = result
					__z.GetValue().SetLval(0)
					__z.SetTypeInfo(4)
					return SUCCESS
				}
				if op1.GetValue().GetObj().GetHandlers().GetCompareObjects() == op2.GetValue().GetObj().GetHandlers().GetCompareObjects() {
					var __z *Zval = result
					__z.GetValue().SetLval(op1.GetValue().GetObj().GetHandlers().GetCompareObjects()(op1, op2))
					__z.SetTypeInfo(4)
					return SUCCESS
				}
			}
			if op1.GetType() == 8 {
				if op1.GetValue().GetObj().GetHandlers().GetGet() != nil {
					var rv Zval
					op_free = op1.GetValue().GetObj().GetHandlers().GetGet()(op1, &rv)
					ret = CompareFunction(result, op_free, op2)
					ZendFreeObjGetResult(op_free)
					return ret
				} else if op2.GetType() != 8 && op1.GetValue().GetObj().GetHandlers().GetCastObject() != nil {
					&tmp_free.SetTypeInfo(0)
					if op1.GetValue().GetObj().GetHandlers().GetCastObject()(op1, &tmp_free, g.CondF2(op2.GetType() == 2 || op2.GetType() == 3, 16, func() ZendUchar { return op2.GetType() })) == FAILURE {
						var __z *Zval = result
						__z.GetValue().SetLval(1)
						__z.SetTypeInfo(4)
						ZendFreeObjGetResult(&tmp_free)
						return SUCCESS
					}
					ret = CompareFunction(result, &tmp_free, op2)
					ZendFreeObjGetResult(&tmp_free)
					return ret
				}
			}
			if op2.GetType() == 8 {
				if op2.GetValue().GetObj().GetHandlers().GetGet() != nil {
					var rv Zval
					op_free = op2.GetValue().GetObj().GetHandlers().GetGet()(op2, &rv)
					ret = CompareFunction(result, op1, op_free)
					ZendFreeObjGetResult(op_free)
					return ret
				} else if op1.GetType() != 8 && op2.GetValue().GetObj().GetHandlers().GetCastObject() != nil {
					&tmp_free.SetTypeInfo(0)
					if op2.GetValue().GetObj().GetHandlers().GetCastObject()(op2, &tmp_free, g.CondF2(op1.GetType() == 2 || op1.GetType() == 3, 16, func() ZendUchar { return op1.GetType() })) == FAILURE {
						var __z *Zval = result
						__z.GetValue().SetLval(-1)
						__z.SetTypeInfo(4)
						ZendFreeObjGetResult(&tmp_free)
						return SUCCESS
					}
					ret = CompareFunction(result, op1, &tmp_free)
					ZendFreeObjGetResult(&tmp_free)
					return ret
				} else if op1.GetType() == 8 {
					var __z *Zval = result
					__z.GetValue().SetLval(1)
					__z.SetTypeInfo(4)
					return SUCCESS
				}
			}
			if converted == 0 {
				if op1.GetType() < 3 {
					var __z *Zval = result
					if ZendIsTrue(op2) != 0 {
						__z.GetValue().SetLval(-1)
					} else {
						__z.GetValue().SetLval(0)
					}
					__z.SetTypeInfo(4)
					return SUCCESS
				} else if op1.GetType() == 3 {
					var __z *Zval = result
					if ZendIsTrue(op2) != 0 {
						__z.GetValue().SetLval(0)
					} else {
						__z.GetValue().SetLval(1)
					}
					__z.SetTypeInfo(4)
					return SUCCESS
				} else if op2.GetType() < 3 {
					var __z *Zval = result
					if ZendIsTrue(op1) != 0 {
						__z.GetValue().SetLval(1)
					} else {
						__z.GetValue().SetLval(0)
					}
					__z.SetTypeInfo(4)
					return SUCCESS
				} else if op2.GetType() == 3 {
					var __z *Zval = result
					if ZendIsTrue(op1) != 0 {
						__z.GetValue().SetLval(0)
					} else {
						__z.GetValue().SetLval(-1)
					}
					__z.SetTypeInfo(4)
					return SUCCESS
				} else {
					if op1.GetType() == 4 || op1.GetType() == 5 {
						op1 = op1
					} else {
						if op1 == result {
							_convertScalarToNumber(op1, 1, 1)
							op1 = op1
						} else {
							op1 = _zendiConvertScalarToNumber(op1, &op1_copy)
						}
					}
					if op2.GetType() == 4 || op2.GetType() == 5 {
						op2 = op2
					} else {
						if op2 == result {
							_convertScalarToNumber(op2, 1, 1)
							op2 = op2
						} else {
							op2 = _zendiConvertScalarToNumber(op2, &op2_copy)
						}
					}
					if EG.GetException() != nil {
						if result != op1 {
							result.SetTypeInfo(0)
						}
						return FAILURE
					}
					converted = 1
				}
			} else if op1.GetType() == 7 {
				var __z *Zval = result
				__z.GetValue().SetLval(1)
				__z.SetTypeInfo(4)
				return SUCCESS
			} else if op2.GetType() == 7 {
				var __z *Zval = result
				__z.GetValue().SetLval(-1)
				__z.SetTypeInfo(4)
				return SUCCESS
			} else {
				assert(false)
				ZendThrowError(nil, "Unsupported operand types")
				if result != op1 {
					result.SetTypeInfo(0)
				}
				return FAILURE
			}
		}
	}
}

/* }}} */

func HashZvalIdenticalFunction(z1 *Zval, z2 *Zval) int {
	/* is_identical_function() returns 1 in case of identity and 0 in case
	 * of a difference;
	 * whereas this comparison function is expected to return 0 on identity,
	 * and non zero otherwise.
	 */

	if z1.GetType() == 10 {
		z1 = &(*z1).value.GetRef().GetVal()
	}
	if z2.GetType() == 10 {
		z2 = &(*z2).value.GetRef().GetVal()
	}
	return FastIsNotIdenticalFunction(z1, z2)
}

/* }}} */

func ZendIsIdentical(op1 *Zval, op2 *Zval) ZendBool {
	if op1.GetType() != op2.GetType() {
		return 0
	}
	switch op1.GetType() {
	case 1:

	case 2:

	case 3:
		return 1
	case 4:
		return op1.GetValue().GetLval() == op2.GetValue().GetLval()
	case 9:
		return op1.GetValue().GetRes() == op2.GetValue().GetRes()
	case 5:
		return op1.GetValue().GetDval() == op2.GetValue().GetDval()
	case 6:
		return ZendStringEquals(op1.GetValue().GetStr(), op2.GetValue().GetStr())
	case 7:
		return op1.GetValue().GetArr() == op2.GetValue().GetArr() || ZendHashCompare(op1.GetValue().GetArr(), op2.GetValue().GetArr(), CompareFuncT(HashZvalIdenticalFunction), 1) == 0
	case 8:
		return op1.GetValue().GetObj() == op2.GetValue().GetObj()
	default:
		return 0
	}
}

/* }}} */

func IsIdenticalFunction(result *Zval, op1 *Zval, op2 *Zval) int {
	if ZendIsIdentical(op1, op2) != 0 {
		result.SetTypeInfo(3)
	} else {
		result.SetTypeInfo(2)
	}
	return SUCCESS
}

/* }}} */

func IsNotIdenticalFunction(result *Zval, op1 *Zval, op2 *Zval) int {
	if ZendIsIdentical(op1, op2) == 0 {
		result.SetTypeInfo(3)
	} else {
		result.SetTypeInfo(2)
	}
	return SUCCESS
}

/* }}} */

func IsEqualFunction(result *Zval, op1 *Zval, op2 *Zval) int {
	if CompareFunction(result, op1, op2) == FAILURE {
		return FAILURE
	}
	if result.GetValue().GetLval() == 0 {
		result.SetTypeInfo(3)
	} else {
		result.SetTypeInfo(2)
	}
	return SUCCESS
}

/* }}} */

func IsNotEqualFunction(result *Zval, op1 *Zval, op2 *Zval) int {
	if CompareFunction(result, op1, op2) == FAILURE {
		return FAILURE
	}
	if result.GetValue().GetLval() != 0 {
		result.SetTypeInfo(3)
	} else {
		result.SetTypeInfo(2)
	}
	return SUCCESS
}

/* }}} */

func IsSmallerFunction(result *Zval, op1 *Zval, op2 *Zval) int {
	if CompareFunction(result, op1, op2) == FAILURE {
		return FAILURE
	}
	if result.GetValue().GetLval() < 0 {
		result.SetTypeInfo(3)
	} else {
		result.SetTypeInfo(2)
	}
	return SUCCESS
}

/* }}} */

func IsSmallerOrEqualFunction(result *Zval, op1 *Zval, op2 *Zval) int {
	if CompareFunction(result, op1, op2) == FAILURE {
		return FAILURE
	}
	if result.GetValue().GetLval() <= 0 {
		result.SetTypeInfo(3)
	} else {
		result.SetTypeInfo(2)
	}
	return SUCCESS
}

/* }}} */

func InstanceofClass(instance_ce *ZendClassEntry, ce *ZendClassEntry) ZendBool {
	for {
		if instance_ce == ce {
			return 1
		}
		instance_ce = instance_ce.parent
		if instance_ce == nil {
			break
		}
	}
	return 0
}

/* }}} */

func InstanceofInterface(instance_ce *ZendClassEntry, ce *ZendClassEntry) ZendBool {
	var i uint32
	if instance_ce.GetNumInterfaces() != 0 {
		assert((instance_ce.GetCeFlags() & 1 << 20) != 0)
		for i = 0; i < instance_ce.GetNumInterfaces(); i++ {
			if instance_ce.interfaces[i] == ce {
				return 1
			}
		}
	}
	return instance_ce == ce
}

/* }}} */

func InstanceofFunctionEx(instance_ce *ZendClassEntry, ce *ZendClassEntry, is_interface ZendBool) ZendBool {
	if is_interface != 0 {
		assert((ce.GetCeFlags() & 1 << 0) != 0)
		return InstanceofInterface(instance_ce, ce)
	} else {
		assert((ce.GetCeFlags() & 1 << 0) == 0)
		return InstanceofClass(instance_ce, ce)
	}
}

/* }}} */

func InstanceofFunction(instance_ce *ZendClassEntry, ce *ZendClassEntry) ZendBool {
	if (ce.GetCeFlags() & 1 << 0) != 0 {
		return InstanceofInterface(instance_ce, ce)
	} else {
		return InstanceofClass(instance_ce, ce)
	}
}

/* }}} */

// #define LOWER_CASE       1

// #define UPPER_CASE       2

// #define NUMERIC       3

func IncrementString(str *Zval) {
	var carry int = 0
	var pos int = str.GetValue().GetStr().GetLen() - 1
	var s *byte
	var t *ZendString
	var last int = 0
	var ch int
	if str.GetValue().GetStr().GetLen() == 0 {
		ZvalPtrDtorStr(str)
		var __z *Zval = str
		var __s *ZendString = ZendOneCharString['1']
		__z.GetValue().SetStr(__s)
		__z.SetTypeInfo(6)
		return
	}
	if str.GetTypeFlags() == 0 {
		str.GetValue().SetStr(ZendStringInit(str.GetValue().GetStr().GetVal(), str.GetValue().GetStr().GetLen(), 0))
		str.SetTypeInfo(6 | 1<<0<<8)
	} else if ZvalRefcountP(str) > 1 {
		ZvalDelrefP(str)
		str.GetValue().SetStr(ZendStringInit(str.GetValue().GetStr().GetVal(), str.GetValue().GetStr().GetLen(), 0))
	} else {
		ZendStringForgetHashVal(str.GetValue().GetStr())
	}
	s = str.GetValue().GetStr().GetVal()
	for {
		ch = s[pos]
		if ch >= 'a' && ch <= 'z' {
			if ch == 'z' {
				s[pos] = 'a'
				carry = 1
			} else {
				s[pos]++
				carry = 0
			}
			last = 1
		} else if ch >= 'A' && ch <= 'Z' {
			if ch == 'Z' {
				s[pos] = 'A'
				carry = 1
			} else {
				s[pos]++
				carry = 0
			}
			last = 2
		} else if ch >= '0' && ch <= '9' {
			if ch == '9' {
				s[pos] = '0'
				carry = 1
			} else {
				s[pos]++
				carry = 0
			}
			last = 3
		} else {
			carry = 0
			break
		}
		if carry == 0 {
			break
		}
		if g.PostDec(&pos) <= 0 {
			break
		}
	}
	if carry != 0 {
		t = ZendStringAlloc(str.GetValue().GetStr().GetLen()+1, 0)
		memcpy(t.GetVal()+1, str.GetValue().GetStr().GetVal(), str.GetValue().GetStr().GetLen())
		t.GetVal()[str.GetValue().GetStr().GetLen()+1] = '0'
		switch last {
		case 3:
			t.GetVal()[0] = '1'
			break
		case 2:
			t.GetVal()[0] = 'A'
			break
		case 1:
			t.GetVal()[0] = 'a'
			break
		}
		ZendStringFree(str.GetValue().GetStr())
		var __z *Zval = str
		var __s *ZendString = t
		__z.GetValue().SetStr(__s)
		__z.SetTypeInfo(6 | 1<<0<<8)
	}
}

/* }}} */

func IncrementFunction(op1 *Zval) int {
try_again:
	switch op1.GetType() {
	case 4:
		FastLongIncrementFunction(op1)
		break
	case 5:
		op1.GetValue().SetDval(op1.GetValue().GetDval() + 1)
		break
	case 1:
		var __z *Zval = op1
		__z.GetValue().SetLval(1)
		__z.SetTypeInfo(4)
		break
	case 6:
		var lval ZendLong
		var dval float64
		switch IsNumericString(op1.GetValue().GetStr().GetVal(), op1.GetValue().GetStr().GetLen(), &lval, &dval, 0) {
		case 4:
			ZvalPtrDtorStr(op1)
			if lval == INT64_MAX {

				/* switch to double */

				var d float64 = float64(lval)
				var __z *Zval = op1
				__z.GetValue().SetDval(d + 1)
				__z.SetTypeInfo(5)
			} else {
				var __z *Zval = op1
				__z.GetValue().SetLval(lval + 1)
				__z.SetTypeInfo(4)
			}
			break
		case 5:
			ZvalPtrDtorStr(op1)
			var __z *Zval = op1
			__z.GetValue().SetDval(dval + 1)
			__z.SetTypeInfo(5)
			break
		default:

			/* Perl style string increment */

			IncrementString(op1)
			break
		}
		break
	case 8:
		if op1.GetValue().GetObj().GetHandlers().GetGet() != nil && op1.GetValue().GetObj().GetHandlers().GetSet() != nil {

			/* proxy object */

			var rv Zval
			var val *Zval
			val = op1.GetValue().GetObj().GetHandlers().GetGet()(op1, &rv)
			if val.GetTypeFlags() != 0 {
				ZvalAddrefP(val)
			}
			IncrementFunction(val)
			op1.GetValue().GetObj().GetHandlers().GetSet()(op1, val)
			ZvalPtrDtor(val)
		} else if op1.GetValue().GetObj().GetHandlers().GetDoOperation() != nil {
			var op2 Zval
			var res int
			var __z *Zval = &op2
			__z.GetValue().SetLval(1)
			__z.SetTypeInfo(4)
			res = op1.GetValue().GetObj().GetHandlers().GetDoOperation()(1, op1, op1, &op2)
			return res
		}
		return FAILURE
	case 10:
		op1 = &(*op1).value.GetRef().GetVal()
		goto try_again
	default:
		return FAILURE
	}
	return SUCCESS
}

/* }}} */

func DecrementFunction(op1 *Zval) int {
	var lval ZendLong
	var dval float64
try_again:
	switch op1.GetType() {
	case 4:
		FastLongDecrementFunction(op1)
		break
	case 5:
		op1.GetValue().SetDval(op1.GetValue().GetDval() - 1)
		break
	case 6:
		if op1.GetValue().GetStr().GetLen() == 0 {
			ZvalPtrDtorStr(op1)
			var __z *Zval = op1
			__z.GetValue().SetLval(-1)
			__z.SetTypeInfo(4)
			break
		}
		switch IsNumericString(op1.GetValue().GetStr().GetVal(), op1.GetValue().GetStr().GetLen(), &lval, &dval, 0) {
		case 4:
			ZvalPtrDtorStr(op1)
			if lval == INT64_MIN {
				var d float64 = float64(lval)
				var __z *Zval = op1
				__z.GetValue().SetDval(d - 1)
				__z.SetTypeInfo(5)
			} else {
				var __z *Zval = op1
				__z.GetValue().SetLval(lval - 1)
				__z.SetTypeInfo(4)
			}
			break
		case 5:
			ZvalPtrDtorStr(op1)
			var __z *Zval = op1
			__z.GetValue().SetDval(dval - 1)
			__z.SetTypeInfo(5)
			break
		}
		break
	case 8:
		if op1.GetValue().GetObj().GetHandlers().GetGet() != nil && op1.GetValue().GetObj().GetHandlers().GetSet() != nil {

			/* proxy object */

			var rv Zval
			var val *Zval
			val = op1.GetValue().GetObj().GetHandlers().GetGet()(op1, &rv)
			if val.GetTypeFlags() != 0 {
				ZvalAddrefP(val)
			}
			DecrementFunction(val)
			op1.GetValue().GetObj().GetHandlers().GetSet()(op1, val)
			ZvalPtrDtor(val)
		} else if op1.GetValue().GetObj().GetHandlers().GetDoOperation() != nil {
			var op2 Zval
			var res int
			var __z *Zval = &op2
			__z.GetValue().SetLval(1)
			__z.SetTypeInfo(4)
			res = op1.GetValue().GetObj().GetHandlers().GetDoOperation()(2, op1, op1, &op2)
			return res
		}
		return FAILURE
	case 10:
		op1 = &(*op1).value.GetRef().GetVal()
		goto try_again
	default:
		return FAILURE
	}
	return SUCCESS
}

/* }}} */

func ZendIsTrue(op *Zval) int { return IZendIsTrue(op) }

/* }}} */

func ZendObjectIsTrue(op *Zval) int {
	if op.GetValue().GetObj().GetHandlers().GetCastObject() != nil {
		var tmp Zval
		if op.GetValue().GetObj().GetHandlers().GetCastObject()(op, &tmp, 16) == SUCCESS {
			return tmp.GetType() == 3
		}
		ZendError(1<<12, "Object of class %s could not be converted to bool", op.GetValue().GetObj().GetCe().GetName().GetVal())
	} else if op.GetValue().GetObj().GetHandlers().GetGet() != nil {
		var result int
		var rv Zval
		var tmp *Zval = op.GetValue().GetObj().GetHandlers().GetGet()(op, &rv)
		if tmp.GetType() != 8 {

			/* for safety - avoid loop */

			result = IZendIsTrue(tmp)
			ZvalPtrDtor(tmp)
			return result
		}
	}
	return 1
}

/* }}} */

func ZendStrTolowerCopy(dest *byte, source *byte, length int) *byte {
	var str *uint8 = (*uint8)(source)
	var result *uint8 = (*uint8)(dest)
	var end *uint8 = str + length
	for str < end {
		g.PostInc(&(*result)) = TolowerMap[uint8(g.PostInc(&(*str)))]
	}
	*result = '0'
	return dest
}

/* }}} */

func ZendStrTolowerDup(source *byte, length int) *byte {
	return ZendStrTolowerCopy((*byte)(_emalloc(length+1)), source, length)
}

/* }}} */

func ZendStrTolower(str *byte, length int) {
	var p *uint8 = (*uint8)(str)
	var end *uint8 = p + length
	for p < end {
		*p = TolowerMap[uint8(*p)]
		p++
	}
}

/* }}} */

func ZendStrTolowerDupEx(source *byte, length int) *byte {
	var p *uint8 = (*uint8)(source)
	var end *uint8 = p + length
	for p < end {
		if (*p) != TolowerMap[uint8(*p)] {
			var res *byte = (*byte)(_emalloc(length + 1))
			var r *uint8
			if p != (*uint8)(source) {
				memcpy(res, source, p-(*uint8)(source))
			}
			r = (*uint8)(p + (res - source))
			for p < end {
				*r = TolowerMap[uint8(*p)]
				p++
				r++
			}
			*r = '0'
			return res
		}
		p++
	}
	return nil
}

/* }}} */

func ZendStringTolowerEx(str *ZendString, persistent int) *ZendString {
	var p *uint8 = (*uint8)(str.GetVal())
	var end *uint8 = p + str.GetLen()
	for p < end {
		if (*p) != TolowerMap[uint8(*p)] {
			var res *ZendString = ZendStringAlloc(str.GetLen(), persistent)
			var r *uint8
			if p != (*uint8)(str.GetVal()) {
				memcpy(res.GetVal(), str.GetVal(), p-(*uint8)(str.GetVal()))
			}
			r = p + (res.GetVal() - str.GetVal())
			for p < end {
				*r = TolowerMap[uint8(*p)]
				p++
				r++
			}
			*r = '0'
			return res
		}
		p++
	}
	return ZendStringCopy(str)
}

/* }}} */

func ZendBinaryStrcmp(s1 *byte, len1 int, s2 *byte, len2 int) int {
	var retval int
	if s1 == s2 {
		return 0
	}
	retval = memcmp(s1, s2, g.Cond(len1 < len2, len1, len2))
	if retval == 0 {
		return int(len1 - len2)
	} else {
		return retval
	}
}

/* }}} */

func ZendBinaryStrncmp(s1 *byte, len1 int, s2 *byte, len2 int, length int) int {
	var retval int
	if s1 == s2 {
		return 0
	}
	retval = memcmp(s1, s2, g.CondF2(length < g.Cond(len1 < len2, len1, len2), length, func() int {
		if len1 < len2 {
			return len1
		} else {
			return len2
		}
	}))
	if retval == 0 {
		return int(g.Cond(length < len1, length, len1) - g.Cond(length < len2, length, len2))
	} else {
		return retval
	}
}

/* }}} */

func ZendBinaryStrcasecmp(s1 *byte, len1 int, s2 *byte, len2 int) int {
	var len_ int
	var c1 int
	var c2 int
	if s1 == s2 {
		return 0
	}
	if len1 < len2 {
		len_ = len1
	} else {
		len_ = len2
	}
	for g.PostDec(&len_) {
		c1 = TolowerMap[uint8(*((*uint8)(g.PostInc(&s1))))]
		c2 = TolowerMap[uint8(*((*uint8)(g.PostInc(&s2))))]
		if c1 != c2 {
			return c1 - c2
		}
	}
	return int(len1 - len2)
}

/* }}} */

func ZendBinaryStrncasecmp(s1 *byte, len1 int, s2 string, len2 int, length int) int {
	var len_ int
	var c1 int
	var c2 int
	if s1 == s2 {
		return 0
	}
	if length < g.Cond(len1 < len2, len1, len2) {
		len_ = length
	} else {
		if len1 < len2 {
			len_ = len1
		} else {
			len_ = len2
		}
	}
	for g.PostDec(&len_) {
		c1 = TolowerMap[uint8(*((*uint8)(g.PostInc(&s1))))]
		c2 = TolowerMap[uint8(*((*uint8)(g.PostInc(&s2))))]
		if c1 != c2 {
			return c1 - c2
		}
	}
	return int(g.Cond(length < len1, length, len1) - g.Cond(length < len2, length, len2))
}

/* }}} */

func ZendBinaryStrcasecmpL(s1 *byte, len1 int, s2 *byte, len2 int) int {
	var len_ int
	var c1 int
	var c2 int
	if s1 == s2 {
		return 0
	}
	if len1 < len2 {
		len_ = len1
	} else {
		len_ = len2
	}
	for g.PostDec(&len_) {
		c1 = tolower(int(*((*uint8)(g.PostInc(&s1)))))
		c2 = tolower(int(*((*uint8)(g.PostInc(&s2)))))
		if c1 != c2 {
			return c1 - c2
		}
	}
	return int(len1 - len2)
}

/* }}} */

func ZendBinaryStrncasecmpL(s1 *byte, len1 int, s2 *byte, len2 int, length int) int {
	var len_ int
	var c1 int
	var c2 int
	if s1 == s2 {
		return 0
	}
	if length < g.Cond(len1 < len2, len1, len2) {
		len_ = length
	} else {
		if len1 < len2 {
			len_ = len1
		} else {
			len_ = len2
		}
	}
	for g.PostDec(&len_) {
		c1 = tolower(int(*((*uint8)(g.PostInc(&s1)))))
		c2 = tolower(int(*((*uint8)(g.PostInc(&s2)))))
		if c1 != c2 {
			return c1 - c2
		}
	}
	return int(g.Cond(length < len1, length, len1) - g.Cond(length < len2, length, len2))
}

/* }}} */

func ZendBinaryZvalStrcmp(s1 *Zval, s2 *Zval) int {
	return ZendBinaryStrcmp(s1.GetValue().GetStr().GetVal(), s1.GetValue().GetStr().GetLen(), s2.GetValue().GetStr().GetVal(), s2.GetValue().GetStr().GetLen())
}

/* }}} */

func ZendBinaryZvalStrncmp(s1 *Zval, s2 *Zval, s3 *Zval) int {
	return ZendBinaryStrncmp(s1.GetValue().GetStr().GetVal(), s1.GetValue().GetStr().GetLen(), s2.GetValue().GetStr().GetVal(), s2.GetValue().GetStr().GetLen(), s3.GetValue().GetLval())
}

/* }}} */

func ZendBinaryZvalStrcasecmp(s1 *Zval, s2 *Zval) int {
	return ZendBinaryStrcasecmpL(s1.GetValue().GetStr().GetVal(), s1.GetValue().GetStr().GetLen(), s2.GetValue().GetStr().GetVal(), s2.GetValue().GetStr().GetLen())
}

/* }}} */

func ZendBinaryZvalStrncasecmp(s1 *Zval, s2 *Zval, s3 *Zval) int {
	return ZendBinaryStrncasecmpL(s1.GetValue().GetStr().GetVal(), s1.GetValue().GetStr().GetLen(), s2.GetValue().GetStr().GetVal(), s2.GetValue().GetStr().GetLen(), s3.GetValue().GetLval())
}

/* }}} */

func ZendiSmartStreq(s1 *ZendString, s2 *ZendString) int {
	var ret1 int
	var ret2 int
	var oflow1 int
	var oflow2 int
	var lval1 ZendLong = 0
	var lval2 ZendLong = 0
	var dval1 float64 = 0.0
	var dval2 float64 = 0.0
	if g.Assign(&ret1, IsNumericStringEx(s1.GetVal(), s1.GetLen(), &lval1, &dval1, 0, &oflow1)) && g.Assign(&ret2, IsNumericStringEx(s2.GetVal(), s2.GetLen(), &lval2, &dval2, 0, &oflow2)) {
		if oflow1 != 0 && oflow1 == oflow2 && dval1-dval2 == 0.0 {

			/* both values are integers overflown to the same side, and the
			 * double comparison may have resulted in crucial accuracy lost */

			goto string_cmp

			/* both values are integers overflown to the same side, and the
			 * double comparison may have resulted in crucial accuracy lost */

		}
		if ret1 == 5 || ret2 == 5 {
			if ret1 != 5 {
				if oflow2 != 0 {

					/* 2nd operand is integer > LONG_MAX (oflow2==1) or < LONG_MIN (-1) */

					return 0

					/* 2nd operand is integer > LONG_MAX (oflow2==1) or < LONG_MIN (-1) */

				}
				dval1 = float64(lval1)
			} else if ret2 != 5 {
				if oflow1 != 0 {
					return 0
				}
				dval2 = float64(lval2)
			} else if dval1 == dval2 && !(isfinite(dval1)) {

				/* Both values overflowed and have the same sign,
				 * so a numeric comparison would be inaccurate */

				goto string_cmp

				/* Both values overflowed and have the same sign,
				 * so a numeric comparison would be inaccurate */

			}
			return dval1 == dval2
		} else {
			return lval1 == lval2
		}
	} else {
	string_cmp:
		return ZendStringEqualContent(s1, s2)
	}
}

/* }}} */

func ZendiSmartStrcmp(s1 *ZendString, s2 *ZendString) int {
	var ret1 int
	var ret2 int
	var oflow1 int
	var oflow2 int
	var lval1 ZendLong = 0
	var lval2 ZendLong = 0
	var dval1 float64 = 0.0
	var dval2 float64 = 0.0
	if g.Assign(&ret1, IsNumericStringEx(s1.GetVal(), s1.GetLen(), &lval1, &dval1, 0, &oflow1)) && g.Assign(&ret2, IsNumericStringEx(s2.GetVal(), s2.GetLen(), &lval2, &dval2, 0, &oflow2)) {
		if oflow1 != 0 && oflow1 == oflow2 && dval1-dval2 == 0.0 {

			/* both values are integers overflown to the same side, and the
			 * double comparison may have resulted in crucial accuracy lost */

			goto string_cmp

			/* both values are integers overflown to the same side, and the
			 * double comparison may have resulted in crucial accuracy lost */

		}
		if ret1 == 5 || ret2 == 5 {
			if ret1 != 5 {
				if oflow2 != 0 {

					/* 2nd operand is integer > LONG_MAX (oflow2==1) or < LONG_MIN (-1) */

					return -1 * oflow2

					/* 2nd operand is integer > LONG_MAX (oflow2==1) or < LONG_MIN (-1) */

				}
				dval1 = float64(lval1)
			} else if ret2 != 5 {
				if oflow1 != 0 {
					return oflow1
				}
				dval2 = float64(lval2)
			} else if dval1 == dval2 && !(isfinite(dval1)) {

				/* Both values overflowed and have the same sign,
				 * so a numeric comparison would be inaccurate */

				goto string_cmp

				/* Both values overflowed and have the same sign,
				 * so a numeric comparison would be inaccurate */

			}
			dval1 = dval1 - dval2
			if dval1 {
				if dval1 < 0 {
					return -1
				} else {
					return 1
				}
			} else {
				return 0
			}
		} else {
			if lval1 > lval2 {
				return 1
			} else {
				if lval1 < lval2 {
					return -1
				} else {
					return 0
				}
			}
		}
	} else {
		var strcmp_ret int
	string_cmp:
		strcmp_ret = ZendBinaryStrcmp(s1.GetVal(), s1.GetLen(), s2.GetVal(), s2.GetLen())
		if strcmp_ret != 0 {
			if strcmp_ret < 0 {
				return -1
			} else {
				return 1
			}
		} else {
			return 0
		}
	}
}

/* }}} */

func HashZvalCompareFunction(z1 *Zval, z2 *Zval) int {
	var result Zval
	if CompareFunction(&result, z1, z2) == FAILURE {
		return 1
	}
	return result.GetValue().GetLval()
}

/* }}} */

func ZendCompareSymbolTables(ht1 *HashTable, ht2 *HashTable) int {
	if ht1 == ht2 {
		return 0
	} else {
		return ZendHashCompare(ht1, ht2, CompareFuncT(HashZvalCompareFunction), 0)
	}
}

/* }}} */

func ZendCompareArrays(a1 *Zval, a2 *Zval) int {
	return ZendCompareSymbolTables(a1.GetValue().GetArr(), a2.GetValue().GetArr())
}

/* }}} */

func ZendCompareObjects(o1 *Zval, o2 *Zval) int {
	if o1.GetValue().GetObj() == o2.GetValue().GetObj() {
		return 0
	}
	if o1.GetValue().GetObj().GetHandlers().GetCompareObjects() == nil {
		return 1
	} else {
		return o1.GetValue().GetObj().GetHandlers().GetCompareObjects()(o1, o2)
	}
}

/* }}} */

func ZendLocaleSprintfDouble(op *Zval) {
	var str *ZendString
	str = ZendStrpprintf(0, "%.*G", int(EG.GetPrecision()), float64(op.GetValue().GetDval()))
	var __z *Zval = op
	var __s *ZendString = str
	__z.GetValue().SetStr(__s)
	__z.SetTypeInfo(6 | 1<<0<<8)
}

/* }}} */

func ZendLongToStr(num ZendLong) *ZendString {
	if ZendUlong(num <= 9) != 0 {
		return ZendOneCharString[ZendUchar('0'+ZendUchar(num))]
	} else {
		var buf []byte
		var res *byte = ZendPrintLongToBuf(buf+g.SizeOf("buf")-1, num)
		return ZendStringInit(res, buf+g.SizeOf("buf")-1-res, 0)
	}
}

/* }}} */

func IsNumericStrFunction(str *ZendString, lval *ZendLong, dval *float64) ZendUchar {
	return IsNumericStringEx(str.GetVal(), str.GetLen(), lval, dval, -1, nil)
}

/* }}} */

func _isNumericStringEx(str *byte, length int, lval *ZendLong, dval *float64, allow_errors int, oflow_info *int) ZendUchar {
	var ptr *byte
	var digits int = 0
	var dp_or_e int = 0
	var local_dval float64 = 0.0
	var type_ ZendUchar
	var tmp_lval ZendUlong = 0
	var neg int = 0
	if length == 0 {
		return 0
	}
	if oflow_info != nil {
		*oflow_info = 0
	}

	/* Skip any whitespace
	 * This is much faster than the isspace() function */

	for (*str) == ' ' || (*str) == '\t' || (*str) == '\n' || (*str) == '\r' || (*str) == 'v' || (*str) == 'f' {
		str++
		length--
	}
	ptr = str
	if (*ptr) == '-' {
		neg = 1
		ptr++
	} else if (*ptr) == '+' {
		ptr++
	}
	if (*ptr) >= '0' && (*ptr) <= '9' {

		/* Skip any leading 0s */

		for (*ptr) == '0' {
			ptr++
		}

		/* Count the number of digits. If a decimal point/exponent is found,
		 * it's a double. Otherwise, if there's a dval or no need to check for
		 * a full match, stop when there are too many digits for a long */

		for type_ = 4; !(digits >= 20 && (dval != nil || allow_errors == 1)); {
		check_digits:
			if (*ptr) >= '0' && (*ptr) <= '9' {
				tmp_lval = tmp_lval*10 + (*ptr) - '0'
				continue
			} else if (*ptr) == '.' && dp_or_e < 1 {
				goto process_double
			} else if ((*ptr) == 'e' || (*ptr) == 'E') && dp_or_e < 2 {
				var e *byte = ptr + 1
				if (*e) == '-' || (*e) == '+' {
					e++
					ptr = e - 1
				}
				if (*e) >= '0' && (*e) <= '9' {
					goto process_double
				}
			}
			break
			digits++
			ptr++
		}
		if digits >= 20 {
			if oflow_info != nil {
				if (*str) == '-' {
					*oflow_info = -1
				} else {
					*oflow_info = 1
				}
			}
			dp_or_e = -1
			goto process_double
		}
	} else if (*ptr) == '.' && (ptr[1] >= '0' && ptr[1] <= '9') {
	process_double:
		type_ = 5

		/* If there's a dval, do the conversion; else continue checking
		 * the digits if we need to check for a full match */

		if dval != nil {
			local_dval = ZendStrtod(str, &ptr)
		} else if allow_errors != 1 && dp_or_e != -1 {
			if g.PostInc(&(*ptr)) == '.' {
				dp_or_e = 1
			} else {
				dp_or_e = 2
			}
			goto check_digits
		}

		/* If there's a dval, do the conversion; else continue checking
		 * the digits if we need to check for a full match */

	} else {
		return 0
	}
	if ptr != str+length {
		if allow_errors == 0 {
			return 0
		}
		if allow_errors == -1 {
			ZendError(1<<3, "A non well formed numeric value encountered")
			if EG.GetException() != nil {
				return 0
			}
		}
	}
	if type_ == 4 {
		if digits == 20-1 {
			var cmp int = strcmp(&ptr[-digits], LongMinDigits)
			if !(cmp < 0 || cmp == 0 && (*str) == '-') {
				if dval != nil {
					*dval = ZendStrtod(str, nil)
				}
				if oflow_info != nil {
					if (*str) == '-' {
						*oflow_info = -1
					} else {
						*oflow_info = 1
					}
				}
				return 5
			}
		}
		if lval != nil {
			if neg != 0 {
				tmp_lval = -tmp_lval
			}
			*lval = ZendLong(tmp_lval)
		}
		return 4
	} else {
		if dval != nil {
			*dval = local_dval
		}
		return 5
	}
}

/* }}} */

func ZendMemnstrExPre(td []uint, needle *byte, needle_len int, reverse int) {
	var i int
	for i = 0; i < 256; i++ {
		td[i] = needle_len + 1
	}
	if reverse != 0 {
		for i = needle_len - 1; i >= 0; i-- {
			td[uint8(needle[i])] = i + 1
		}
	} else {
		var i int
		for i = 0; i < needle_len; i++ {
			td[uint8(needle[i])] = int(needle_len - i)
		}
	}
}

/* }}} */

func ZendMemnstrEx(haystack *byte, needle string, needle_len int, end *byte) *byte {
	var td []uint
	var i int
	var p *byte
	if needle_len == 0 || end-haystack < needle_len {
		return nil
	}
	ZendMemnstrExPre(td, needle, needle_len, 0)
	p = haystack
	end -= needle_len
	for p <= end {
		for i = 0; i < needle_len; i++ {
			if needle[i] != p[i] {
				break
			}
		}
		if i == needle_len {
			return p
		}
		if p == end {
			return nil
		}
		p += td[uint8(p[needle_len])]
	}
	return nil
}

/* }}} */

func ZendMemnrstrEx(haystack *byte, needle *byte, needle_len int, end *byte) *byte {
	var td []uint
	var i int
	var p *byte
	if needle_len == 0 || end-haystack < needle_len {
		return nil
	}
	ZendMemnstrExPre(td, needle, needle_len, 1)
	p = end
	p -= needle_len
	for p >= haystack {
		for i = 0; i < needle_len; i++ {
			if needle[i] != p[i] {
				break
			}
		}
		if i == needle_len {
			return (*byte)(p)
		}
		if p == haystack {
			return nil
		}
		p -= td[uint8(p[-1])]
	}
	return nil
}

/* }}} */

func ZendDvalToLvalSlow(d float64) ZendLong {
	var two_pow_64 float64 = pow(2.0, 64.0)
	var dmod float64
	dmod = fmod(d, two_pow_64)
	if dmod < 0 {

		/* no need to call ceil; original double must have had no
		 * fractional part, hence dmod does not have one either */

		dmod += two_pow_64

		/* no need to call ceil; original double must have had no
		 * fractional part, hence dmod does not have one either */

	}
	return ZendLong(ZendUlong(dmod))
}
