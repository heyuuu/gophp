// <<generate>>

package zend

import (
	b "sik/builtin"
	"sik/core"
	r "sik/runtime"
)

func ZEND_DOUBLE_FITS_LONG(d float64) bool {
	return !(d >= float64(ZEND_LONG_MAX != 0 || d < float64(ZEND_LONG_MIN)))
}
func ZendDvalToLval(d float64) ZendLong {
	if !(core.ZendFinite(d)) || core.ZendIsnan(d) {
		return 0
	} else if !(ZEND_DOUBLE_FITS_LONG(d)) {
		return ZendDvalToLvalSlow(d)
	}
	return ZendLong(d)
}
func ZendDvalToLvalCap(d float64) ZendLong {
	if !(core.ZendFinite(d)) || core.ZendIsnan(d) {
		return 0
	} else if !(ZEND_DOUBLE_FITS_LONG(d)) {
		if d > 0 {
			return ZEND_LONG_MAX
		} else {
			return ZEND_LONG_MIN
		}
	}
	return ZendLong(d)
}
func ZEND_IS_DIGIT(c byte) bool { return c >= '0' && c <= '9' }
func ZEND_IS_XDIGIT(c __auto__) bool {
	return c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f'
}
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
			if b.Assign(&p, (*byte)(memchr(p, *needle, end-p+1))) && ne == p[needle_len-1] {
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
			if b.PostDec(&p) < haystack {
				break
			}
		}
		return nil
	} else {
		return ZendMemnrstrEx(haystack, needle, needle_len, end)
	}
}
func ZvalGetLong(op *Zval) ZendLong {
	if op.IsType(IS_LONG) {
		return op.GetLval()
	} else {
		return ZvalGetLongFunc(op)
	}
}
func ZvalGetDouble(op *Zval) float64 {
	if op.IsType(IS_DOUBLE) {
		return op.GetDval()
	} else {
		return ZvalGetDoubleFunc(op)
	}
}
func ZvalGetString(op *Zval) *ZendString {
	if op.IsType(IS_STRING) {
		return ZendStringCopy(op.GetStr())
	} else {
		return ZvalGetStringFunc(op)
	}
}
func ZvalGetTmpString(op *Zval, tmp **ZendString) *ZendString {
	if op.IsType(IS_STRING) {
		*tmp = nil
		return op.GetStr()
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
func ZvalTryGetString(op *Zval) *ZendString {
	if op.IsType(IS_STRING) {
		var ret *ZendString = ZendStringCopy(op.GetStr())
		return ret
	} else {
		return ZvalTryGetStringFunc(op)
	}
}
func ZvalTryGetTmpString(op *Zval, tmp **ZendString) *ZendString {
	if op.IsType(IS_STRING) {
		var ret *ZendString = op.GetStr()
		*tmp = nil
		return ret
	} else {
		*tmp = ZvalTryGetStringFunc(op)
		return *tmp
	}
}
func TryConvertToString(op *Zval) ZendBool {
	if op.IsType(IS_STRING) {
		return 1
	}
	return _tryConvertToString(op)
}
func _zvalGetLong(op *Zval) ZendLong          { return ZvalGetLong(op) }
func _zvalGetDouble(op *Zval) float64         { return ZvalGetDouble(op) }
func _zvalGetString(op *Zval) *ZendString     { return ZvalGetString(op) }
func _zvalGetLongFunc(op *Zval) ZendLong      { return ZvalGetLongFunc(op) }
func _zvalGetDoubleFunc(op *Zval) float64     { return ZvalGetDoubleFunc(op) }
func _zvalGetStringFunc(op *Zval) *ZendString { return ZvalGetStringFunc(op) }
func ConvertToCstring(op *Zval) {
	if op.GetType() != IS_STRING {
		_convertToCstring(op)
	}
}
func ConvertToString(op *Zval) {
	if op.GetType() != IS_STRING {
		_convertToString(op)
	}
}
func ZvalIsTrue(op *Zval) int { return ZendIsTrue(op) }
func IZendIsTrue(op *Zval) int {
	var result int = 0
again:
	switch op.GetType() {
	case IS_TRUE:
		result = 1
		break
	case IS_LONG:
		if op.GetLval() != 0 {
			result = 1
		}
		break
	case IS_DOUBLE:
		if op.GetDval() {
			result = 1
		}
		break
	case IS_STRING:
		if Z_STRLEN_P(op) > 1 || Z_STRLEN_P(op) != 0 && Z_STRVAL_P(op)[0] != '0' {
			result = 1
		}
		break
	case IS_ARRAY:
		if Z_ARRVAL_P(op).GetNNumOfElements() {
			result = 1
		}
		break
	case IS_OBJECT:
		if Z_OBJ_HT_P(op).GetCastObject() == ZendStdCastObjectTostring {
			result = 1
		} else {
			result = ZendObjectIsTrue(op)
		}
		break
	case IS_RESOURCE:
		if Z_RES_HANDLE_P(op) != 0 {
			result = 1
		}
		break
	case IS_REFERENCE:
		op = Z_REFVAL_P(op)
		goto again
		break
	default:
		break
	}
	return result
}
func ZendStringTolower(str *ZendString) *ZendString { return ZendStringTolowerEx(str, 0) }
func ConvertToExplicitType(pzv *Zval, type_ __auto__) {
	for {
		switch type_ {
		case IS_NULL:
			ConvertToNull(pzv)
			break
		case IS_LONG:
			ConvertToLong(pzv)
			break
		case IS_DOUBLE:
			ConvertToDouble(pzv)
			break
		case _IS_BOOL:
			ConvertToBoolean(pzv)
			break
		case IS_ARRAY:
			ConvertToArray(pzv)
			break
		case IS_OBJECT:
			ConvertToObject(pzv)
			break
		case IS_STRING:
			ConvertToString(pzv)
			break
		default:
			r.Assert(false)
			break
		}
		break
	}
}
func ConvertToExplicitTypeEx(pzv *Zval, str_type __auto__) {
	if pzv.GetType() != str_type {
		ConvertToExplicitType(pzv, str_type)
	}
}
func ConvertToBooleanEx(pzv *Zval) {
	if pzv.GetTypeInfo() > IS_TRUE {
		ConvertToBoolean(pzv)
	} else if pzv.GetTypeInfo() < IS_FALSE {
		ZVAL_FALSE(pzv)
	}
}
func ConvertToStringEx(pzv *Zval) {
	if pzv.GetType() != IS_STRING {
		ConvertToString(pzv)
	}
}
func ConvertToArrayEx(pzv *Zval) {
	if pzv.GetType() != IS_ARRAY {
		ConvertToArray(pzv)
	}
}
func ConvertToObjectEx(pzv *Zval) {
	if pzv.GetType() != IS_OBJECT {
		ConvertToObject(pzv)
	}
}
func ConvertToNullEx(pzv *Zval) {
	if pzv.GetType() != IS_NULL {
		ConvertToNull(pzv)
	}
}
func ConvertScalarToNumberEx(pzv *Zval) {
	if pzv.GetType() != IS_LONG && pzv.GetType() != IS_DOUBLE {
		ConvertScalarToNumber(pzv)
	}
}
func FastLongIncrementFunction(op1 *Zval) {
	if op1.GetLval() == ZEND_LONG_MAX {

		/* switch to double */

		ZVAL_DOUBLE(op1, float64(ZEND_LONG_MAX+1.0))

		/* switch to double */

	} else {
		op1.GetLval()++
	}
}
func FastLongDecrementFunction(op1 *Zval) {
	if op1.GetLval() == ZEND_LONG_MIN {

		/* switch to double */

		ZVAL_DOUBLE(op1, float64(ZEND_LONG_MIN-1.0))

		/* switch to double */

	} else {
		op1.GetLval()--
	}
}
func FastLongAddFunction(result *Zval, op1 *Zval, op2 *Zval) {
	/*
	 * 'result' may alias with op1 or op2, so we need to
	 * ensure that 'result' is not updated until after we
	 * have read the values of op1 and op2.
	 */

	if (op1.GetLval()&LONG_SIGN_MASK) == (op2.GetLval()&LONG_SIGN_MASK) && (op1.GetLval()&LONG_SIGN_MASK) != (op1.GetLval()+op2.GetLval()&LONG_SIGN_MASK) {
		ZVAL_DOUBLE(result, float64(op1.GetLval()+float64(op2.GetLval())))
	} else {
		ZVAL_LONG(result, op1.GetLval()+op2.GetLval())
	}

	/*
	 * 'result' may alias with op1 or op2, so we need to
	 * ensure that 'result' is not updated until after we
	 * have read the values of op1 and op2.
	 */
}
func FastAddFunction(result *Zval, op1 *Zval, op2 *Zval) int {
	if op1.IsType(IS_LONG) {
		if op2.IsType(IS_LONG) {
			FastLongAddFunction(result, op1, op2)
			return SUCCESS
		} else if op2.IsType(IS_DOUBLE) {
			ZVAL_DOUBLE(result, float64(op1.GetLval())+op2.GetDval())
			return SUCCESS
		}
	} else if op1.IsType(IS_DOUBLE) {
		if op2.IsType(IS_DOUBLE) {
			ZVAL_DOUBLE(result, op1.GetDval()+op2.GetDval())
			return SUCCESS
		} else if op2.IsType(IS_LONG) {
			ZVAL_DOUBLE(result, op1.GetDval()+float64(op2.GetLval()))
			return SUCCESS
		}
	}
	return AddFunction(result, op1, op2)
}
func FastLongSubFunction(result *Zval, op1 *Zval, op2 *Zval) {
	ZVAL_LONG(result, op1.GetLval()-op2.GetLval())
	if (op1.GetLval()&LONG_SIGN_MASK) != (op2.GetLval()&LONG_SIGN_MASK) && (op1.GetLval()&LONG_SIGN_MASK) != (result.GetLval()&LONG_SIGN_MASK) {
		ZVAL_DOUBLE(result, float64(op1.GetLval()-float64(op2.GetLval())))
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
	if op1.IsType(IS_LONG) {
		if op2.IsType(IS_LONG) {
			return op1.GetLval() == op2.GetLval()
		} else if op2.IsType(IS_DOUBLE) {
			return float64(op1.GetLval()) == op2.GetDval()
		}
	} else if op1.IsType(IS_DOUBLE) {
		if op2.IsType(IS_DOUBLE) {
			return op1.GetDval() == op2.GetDval()
		} else if op2.IsType(IS_LONG) {
			return op1.GetDval() == float64(op2.GetLval())
		}
	} else if op1.IsType(IS_STRING) {
		if op2.IsType(IS_STRING) {
			return ZendFastEqualStrings(op1.GetStr(), op2.GetStr())
		}
	}
	CompareFunction(&result, op1, op2)
	return result.GetLval() == 0
}
func FastEqualCheckLong(op1 *Zval, op2 *Zval) int {
	var result Zval
	if op2.IsType(IS_LONG) {
		return op1.GetLval() == op2.GetLval()
	}
	CompareFunction(&result, op1, op2)
	return result.GetLval() == 0
}
func FastEqualCheckString(op1 *Zval, op2 *Zval) int {
	var result Zval
	if op2.IsType(IS_STRING) {
		return ZendFastEqualStrings(op1.GetStr(), op2.GetStr())
	}
	CompareFunction(&result, op1, op2)
	return result.GetLval() == 0
}
func FastIsIdenticalFunction(op1 *Zval, op2 *Zval) ZendBool {
	if op1.GetType() != op2.GetType() {
		return 0
	} else if op1.GetType() <= IS_TRUE {
		return 1
	}
	return ZendIsIdentical(op1, op2)
}
func FastIsNotIdenticalFunction(op1 *Zval, op2 *Zval) ZendBool {
	if op1.GetType() != op2.GetType() {
		return 1
	} else if op1.GetType() <= IS_TRUE {
		return 0
	}
	return !(ZendIsIdentical(op1, op2))
}
func ZendPrintUlongToBuf(buf *byte, num ZendUlong) *byte {
	*buf = '0'
	for {
		*(b.PreDec(&buf)) = byte(num%10 + '0')
		num /= 10
		if num <= 0 {
			break
		}
	}
	return buf
}
func ZendPrintLongToBuf(buf *byte, num ZendLong) *byte {
	if num < 0 {
		var result *byte = ZendPrintUlongToBuf(buf, ^ZendUlong(num)+1)
		*(b.PreDec(&result)) = '-'
		return result
	} else {
		return ZendPrintUlongToBuf(buf, num)
	}
}
func ZendUnwrapReference(op *Zval) {
	if Z_REFCOUNT_P(op) == 1 {
		ZVAL_UNREF(op)
	} else {
		Z_DELREF_P(op)
		ZVAL_COPY(op, Z_REFVAL_P(op))
	}
}
func ZendTolower(c int) __auto__         { return tolower(c) }
func TYPE_PAIR(t1 uint32, t2 uint32) int { return t1<<4 | t2 }
func ZendTolowerAscii(c uint8) uint8     { return TolowerMap[uint8(c)] }
func ZendAtoi(str *byte, str_len int) int {
	var retval int
	if str_len == 0 {
		str_len = strlen(str)
	}
	retval = ZEND_STRTOL(str, nil, 0)
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
func ZendAtol(str *byte, str_len int) ZendLong {
	var retval ZendLong
	if str_len == 0 {
		str_len = strlen(str)
	}
	retval = ZEND_STRTOL(str, nil, 0)
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
func ConvertObjectToType(op *Zval, dst *Zval, ctype int, conv_func func(op *Zval)) {
	ZVAL_UNDEF(dst)
	if Z_OBJ_HT_P(op).GetCastObject() != nil {
		if Z_OBJ_HT_P(op).GetCastObject()(op, dst, ctype) == FAILURE {
			ZendError(E_RECOVERABLE_ERROR, "Object of class %s could not be converted to %s", Z_OBJCE_P(op).GetName().GetVal(), ZendGetTypeByConst(ctype))
		}
	} else if Z_OBJ_HT_P(op).GetGet() != nil {
		var newop *Zval = Z_OBJ_HT_P(op).GetGet()(op, dst)
		if newop.GetType() != IS_OBJECT {
			ZVAL_COPY_VALUE(dst, newop)
			conv_func(dst)
		}
	}
}
func _convertScalarToNumber(op *Zval, silent ZendBool, check ZendBool) {
try_again:
	switch op.GetType() {
	case IS_REFERENCE:
		ZendUnwrapReference(op)
		goto try_again
	case IS_STRING:
		var str *ZendString
		str = op.GetStr()
		if b.Assign(&(op.GetTypeInfo()), IsNumericString(str.GetVal(), str.GetLen(), &(op.GetLval()), &(op.GetDval()), b.Cond(silent != 0, 1, -1))) == 0 {
			ZVAL_LONG(op, 0)
			if silent == 0 {
				ZendError(E_WARNING, "A non-numeric value encountered")
			}
		}
		ZendStringReleaseEx(str, 0)
		break
	case IS_NULL:

	case IS_FALSE:
		ZVAL_LONG(op, 0)
		break
	case IS_TRUE:
		ZVAL_LONG(op, 1)
		break
	case IS_RESOURCE:
		var l ZendLong = Z_RES_HANDLE_P(op)
		ZvalPtrDtor(op)
		ZVAL_LONG(op, l)
		break
	case IS_OBJECT:
		var dst Zval
		ConvertObjectToType(op, &dst, _IS_NUMBER, ConvertScalarToNumber)
		if check != 0 && ExecutorGlobals.GetException() != nil {
			return
		}
		ZvalPtrDtor(op)
		if dst.IsType(IS_LONG) || dst.IsType(IS_DOUBLE) {
			ZVAL_COPY_VALUE(op, &dst)
		} else {
			ZVAL_LONG(op, 1)
		}
		break
	}
}
func ConvertScalarToNumber(op *Zval) { _convertScalarToNumber(op, 1, 0) }
func _zendiConvertScalarToNumberEx(op *Zval, holder *Zval, silent ZendBool) *Zval {
	switch op.GetType() {
	case IS_NULL:

	case IS_FALSE:
		ZVAL_LONG(holder, 0)
		return holder
	case IS_TRUE:
		ZVAL_LONG(holder, 1)
		return holder
	case IS_STRING:
		if b.Assign(&(holder.GetTypeInfo()), IsNumericString(Z_STRVAL_P(op), Z_STRLEN_P(op), &(holder.GetLval()), &(holder.GetDval()), b.Cond(silent != 0, 1, -1))) == 0 {
			ZVAL_LONG(holder, 0)
			if silent == 0 {
				ZendError(E_WARNING, "A non-numeric value encountered")
			}
		}
		return holder
	case IS_RESOURCE:
		ZVAL_LONG(holder, Z_RES_HANDLE_P(op))
		return holder
	case IS_OBJECT:
		ConvertObjectToType(op, holder, _IS_NUMBER, ConvertScalarToNumber)
		if ExecutorGlobals.GetException() != nil || holder.GetType() != IS_LONG && holder.GetType() != IS_DOUBLE {
			ZVAL_LONG(holder, 1)
		}
		return holder
	case IS_LONG:

	case IS_DOUBLE:

	default:
		return op
	}
}
func _zendiConvertScalarToNumber(op *Zval, holder *Zval) *Zval {
	return _zendiConvertScalarToNumberEx(op, holder, 1)
}
func _zendiConvertScalarToNumberNoisy(op *Zval, holder *Zval) *Zval {
	return _zendiConvertScalarToNumberEx(op, holder, 0)
}
func ZendiConvertScalarToNumber(op *Zval, holder *Zval, result *Zval, silent ZendBool) *Zval {
	if op.IsType(IS_LONG) || op.IsType(IS_DOUBLE) {
		return op
	} else {
		if op == result {
			_convertScalarToNumber(op, silent, 1)
			return op
		} else {
			if silent != 0 {
				return _zendiConvertScalarToNumber(op, holder)
			} else {
				return _zendiConvertScalarToNumberNoisy(op, holder)
			}
		}
	}
}
func ConvertToLong(op *Zval) {
	if op.GetType() != IS_LONG {
		ConvertToLongBase(op, 10)
	}
}
func ConvertToLongBase(op *Zval, base int) {
	var tmp ZendLong
try_again:
	switch op.GetType() {
	case IS_NULL:

	case IS_FALSE:
		ZVAL_LONG(op, 0)
		break
	case IS_TRUE:
		ZVAL_LONG(op, 1)
		break
	case IS_RESOURCE:
		tmp = Z_RES_HANDLE_P(op)
		ZvalPtrDtor(op)
		ZVAL_LONG(op, tmp)
		break
	case IS_LONG:
		break
	case IS_DOUBLE:
		ZVAL_LONG(op, ZendDvalToLval(op.GetDval()))
		break
	case IS_STRING:
		var str *ZendString = op.GetStr()
		if base == 10 {
			ZVAL_LONG(op, ZvalGetLong(op))
		} else {
			ZVAL_LONG(op, ZEND_STRTOL(str.GetVal(), nil, base))
		}
		ZendStringReleaseEx(str, 0)
		break
	case IS_ARRAY:
		if Z_ARRVAL_P(op).GetNNumOfElements() {
			tmp = 1
		} else {
			tmp = 0
		}
		ZvalPtrDtor(op)
		ZVAL_LONG(op, tmp)
		break
	case IS_OBJECT:
		var dst Zval
		ConvertObjectToType(op, &dst, IS_LONG, ConvertToLong)
		ZvalPtrDtor(op)
		if dst.IsType(IS_LONG) {
			ZVAL_LONG(op, dst.GetLval())
		} else {
			ZVAL_LONG(op, 1)
		}
		return
	case IS_REFERENCE:
		ZendUnwrapReference(op)
		goto try_again
	default:
		break
	}
}
func ConvertToDouble(op *Zval) {
	var tmp float64
try_again:
	switch op.GetType() {
	case IS_NULL:

	case IS_FALSE:
		ZVAL_DOUBLE(op, 0.0)
		break
	case IS_TRUE:
		ZVAL_DOUBLE(op, 1.0)
		break
	case IS_RESOURCE:
		var d float64 = float64(Z_RES_HANDLE_P(op))
		ZvalPtrDtor(op)
		ZVAL_DOUBLE(op, d)
		break
	case IS_LONG:
		ZVAL_DOUBLE(op, float64(op.GetLval()))
		break
	case IS_DOUBLE:
		break
	case IS_STRING:
		var str *ZendString = op.GetStr()
		ZVAL_DOUBLE(op, ZendStrtod(str.GetVal(), nil))
		ZendStringReleaseEx(str, 0)
		break
	case IS_ARRAY:
		if Z_ARRVAL_P(op).GetNNumOfElements() {
			tmp = 1
		} else {
			tmp = 0
		}
		ZvalPtrDtor(op)
		ZVAL_DOUBLE(op, tmp)
		break
	case IS_OBJECT:
		var dst Zval
		ConvertObjectToType(op, &dst, IS_DOUBLE, ConvertToDouble)
		ZvalPtrDtor(op)
		if dst.IsType(IS_DOUBLE) {
			ZVAL_DOUBLE(op, dst.GetDval())
		} else {
			ZVAL_DOUBLE(op, 1.0)
		}
		break
	case IS_REFERENCE:
		ZendUnwrapReference(op)
		goto try_again
	default:
		break
	}
}
func ConvertToNull(op *Zval) {
	ZvalPtrDtor(op)
	ZVAL_NULL(op)
}
func ConvertToBoolean(op *Zval) {
	var tmp int
try_again:
	switch op.GetType() {
	case IS_FALSE:

	case IS_TRUE:
		break
	case IS_NULL:
		ZVAL_FALSE(op)
		break
	case IS_RESOURCE:
		var l ZendLong = b.Cond(Z_RES_HANDLE_P(op) != 0, 1, 0)
		ZvalPtrDtor(op)
		ZVAL_BOOL(op, l)
		break
	case IS_LONG:
		ZVAL_BOOL(op, b.Cond(op.GetLval() != 0, 1, 0))
		break
	case IS_DOUBLE:
		ZVAL_BOOL(op, b.Cond(op.GetDval(), 1, 0))
		break
	case IS_STRING:
		var str *ZendString = op.GetStr()
		if str.GetLen() == 0 || str.GetLen() == 1 && str.GetVal()[0] == '0' {
			ZVAL_FALSE(op)
		} else {
			ZVAL_TRUE(op)
		}
		ZendStringReleaseEx(str, 0)
		break
	case IS_ARRAY:
		if Z_ARRVAL_P(op).GetNNumOfElements() {
			tmp = 1
		} else {
			tmp = 0
		}
		ZvalPtrDtor(op)
		ZVAL_BOOL(op, tmp)
		break
	case IS_OBJECT:
		var dst Zval
		ConvertObjectToType(op, &dst, _IS_BOOL, ConvertToBoolean)
		ZvalPtrDtor(op)
		if dst.GetTypeInfo() == IS_FALSE || dst.GetTypeInfo() == IS_TRUE {
			op.SetTypeInfo(dst.GetTypeInfo())
		} else {
			ZVAL_TRUE(op)
		}
		break
	case IS_REFERENCE:
		ZendUnwrapReference(op)
		goto try_again
	default:
		break
	}
}
func _convertToCstring(op *Zval) {
	if op.IsType(IS_DOUBLE) {
		var str *ZendString
		var dval float64 = op.GetDval()
		str = ZendStrpprintfUnchecked(0, "%.*H", int(ExecutorGlobals.GetPrecision()), dval)
		ZVAL_NEW_STR(op, str)
	} else {
		_convertToString(op)
	}
}
func _convertToString(op *Zval) {
try_again:
	switch op.GetType() {
	case IS_UNDEF:

	case IS_NULL:

	case IS_FALSE:
		ZVAL_EMPTY_STRING(op)
		break
	case IS_TRUE:
		ZVAL_INTERNED_STR(op, ZSTR_CHAR('1'))
		break
	case IS_STRING:
		break
	case IS_RESOURCE:
		var str *ZendString = ZendStrpprintf(0, "Resource id #"+ZEND_LONG_FMT, ZendLong(Z_RES_HANDLE_P(op)))
		ZvalPtrDtor(op)
		ZVAL_NEW_STR(op, str)
		break
	case IS_LONG:
		ZVAL_STR(op, ZendLongToStr(op.GetLval()))
		break
	case IS_DOUBLE:
		var str *ZendString
		var dval float64 = op.GetDval()
		str = ZendStrpprintf(0, "%.*G", int(ExecutorGlobals.GetPrecision()), dval)

		/* %G already handles removing trailing zeros from the fractional part, yay */

		ZVAL_NEW_STR(op, str)
		break
	case IS_ARRAY:
		ZendError(E_NOTICE, "Array to string conversion")
		ZvalPtrDtor(op)
		ZVAL_INTERNED_STR(op, ZSTR_KNOWN(ZEND_STR_ARRAY_CAPITALIZED))
		break
	case IS_OBJECT:
		var tmp Zval
		if Z_OBJ_HT_P(op).GetCastObject() != nil {
			if Z_OBJ_HT_P(op).GetCastObject()(op, &tmp, IS_STRING) == SUCCESS {
				ZvalPtrDtor(op)
				ZVAL_COPY_VALUE(op, &tmp)
				return
			}
		} else if Z_OBJ_HT_P(op).GetGet() != nil {
			var z *Zval = Z_OBJ_HT_P(op).GetGet()(op, &tmp)
			if z.GetType() != IS_OBJECT {
				var str *ZendString = ZvalGetString(z)
				ZvalPtrDtor(z)
				ZvalPtrDtor(op)
				ZVAL_STR(op, str)
				return
			}
			ZvalPtrDtor(z)
		}
		if ExecutorGlobals.GetException() == nil {
			ZendThrowError(nil, "Object of class %s could not be converted to string", Z_OBJCE_P(op).GetName().GetVal())
		}
		ZvalPtrDtor(op)
		ZVAL_EMPTY_STRING(op)
		break
	case IS_REFERENCE:
		ZendUnwrapReference(op)
		goto try_again
	default:
		break
	}
}
func _tryConvertToString(op *Zval) ZendBool {
	var str *ZendString
	ZEND_ASSERT(op.GetType() != IS_STRING)
	str = ZvalTryGetStringFunc(op)
	if str == nil {
		return 0
	}
	ZvalPtrDtor(op)
	ZVAL_STR(op, str)
	return 1
}
func ConvertScalarToArray(op *Zval) {
	var ht *HashTable = ZendNewArray(1)
	ZendHashIndexAddNew(ht, 0, op)
	ZVAL_ARR(op, ht)
}
func ConvertToArray(op *Zval) {
try_again:
	switch op.GetType() {
	case IS_ARRAY:
		break
	case IS_OBJECT:
		if Z_OBJCE_P(op) == ZendCeClosure {
			ConvertScalarToArray(op)
		} else {
			var obj_ht *HashTable = ZendGetPropertiesFor(op, ZEND_PROP_PURPOSE_ARRAY_CAST)
			if obj_ht != nil {
				var new_obj_ht *HashTable = ZendProptableToSymtable(obj_ht, Z_OBJCE_P(op).GetDefaultPropertiesCount() != 0 || Z_OBJ_P(op).GetHandlers() != &StdObjectHandlers || GC_IS_RECURSIVE(obj_ht) != 0)
				ZvalPtrDtor(op)
				ZVAL_ARR(op, new_obj_ht)
				ZendReleaseProperties(obj_ht)
			} else {
				ZvalPtrDtor(op)

				/*ZVAL_EMPTY_ARRAY(op);*/

				ArrayInit(op)

				/*ZVAL_EMPTY_ARRAY(op);*/

			}
		}
		break
	case IS_NULL:

		/*ZVAL_EMPTY_ARRAY(op);*/

		ArrayInit(op)
		break
	case IS_REFERENCE:
		ZendUnwrapReference(op)
		goto try_again
	default:
		ConvertScalarToArray(op)
		break
	}
}
func ConvertToObject(op *Zval) {
try_again:
	switch op.GetType() {
	case IS_ARRAY:
		var ht *HashTable = ZendSymtableToProptable(op.GetArr())
		var obj *ZendObject
		if (ht.GetGcFlags() & IS_ARRAY_IMMUTABLE) != 0 {

			/* TODO: try not to duplicate immutable arrays as well ??? */

			ht = ZendArrayDup(ht)

			/* TODO: try not to duplicate immutable arrays as well ??? */

		} else if ht != op.GetArr() {
			ZvalPtrDtor(op)
		} else {
			ht.DecGcRefcount()
		}
		obj = ZendObjectsNew(ZendStandardClassDef)
		obj.SetProperties(ht)
		ZVAL_OBJ(op, obj)
		break
	case IS_OBJECT:
		break
	case IS_NULL:
		ObjectInit(op)
		break
	case IS_REFERENCE:
		ZendUnwrapReference(op)
		goto try_again
	default:
		var tmp Zval
		ZVAL_COPY_VALUE(&tmp, op)
		ObjectInit(op)
		ZendHashAddNew(Z_OBJPROP_P(op), ZSTR_KNOWN(ZEND_STR_SCALAR), &tmp)
		break
	}
}
func MultiConvertToLongEx(argc int, _ ...any) {
	var arg *Zval
	var ap va_list
	va_start(ap, argc)
	for b.PostDec(&argc) {
		arg = __va_arg(ap, (*Zval)(_))
		if arg.GetType() != IS_LONG {
			ConvertToLong(arg)
		}
	}
	va_end(ap)
}
func MultiConvertToDoubleEx(argc int, _ ...any) {
	var arg *Zval
	var ap va_list
	va_start(ap, argc)
	for b.PostDec(&argc) {
		arg = __va_arg(ap, (*Zval)(_))
		if arg.GetType() != IS_DOUBLE {
			ConvertToDouble(arg)
		}
	}
	va_end(ap)
}
func MultiConvertToStringEx(argc int, _ ...any) {
	var arg *Zval
	var ap va_list
	va_start(ap, argc)
	for b.PostDec(&argc) {
		arg = __va_arg(ap, (*Zval)(_))
		ConvertToStringEx(arg)
	}
	va_end(ap)
}
func _zvalGetLongFuncEx(op *Zval, silent ZendBool) ZendLong {
try_again:
	switch op.GetType() {
	case IS_UNDEF:

	case IS_NULL:

	case IS_FALSE:
		return 0
	case IS_TRUE:
		return 1
	case IS_RESOURCE:
		return Z_RES_HANDLE_P(op)
	case IS_LONG:
		return op.GetLval()
	case IS_DOUBLE:
		return ZendDvalToLval(op.GetDval())
	case IS_STRING:
		var type_ ZendUchar
		var lval ZendLong
		var dval float64
		if 0 == b.Assign(&type_, IsNumericString(Z_STRVAL_P(op), Z_STRLEN_P(op), &lval, &dval, b.Cond(silent != 0, 1, -1))) {
			if silent == 0 {
				ZendError(E_WARNING, "A non-numeric value encountered")
			}
			return 0
		} else if type_ == IS_LONG {
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
	case IS_ARRAY:
		if Z_ARRVAL_P(op).GetNNumOfElements() {
			return 1
		} else {
			return 0
		}
	case IS_OBJECT:
		var dst Zval
		ConvertObjectToType(op, &dst, IS_LONG, ConvertToLong)
		if dst.IsType(IS_LONG) {
			return dst.GetLval()
		} else {
			return 1
		}
	case IS_REFERENCE:
		op = Z_REFVAL_P(op)
		goto try_again
	default:
		break
	}
	return 0
}
func ZvalGetLongFunc(op *Zval) ZendLong       { return _zvalGetLongFuncEx(op, 1) }
func _zvalGetLongFuncNoisy(op *Zval) ZendLong { return _zvalGetLongFuncEx(op, 0) }
func ZvalGetDoubleFunc(op *Zval) float64 {
try_again:
	switch op.GetType() {
	case IS_NULL:

	case IS_FALSE:
		return 0.0
	case IS_TRUE:
		return 1.0
	case IS_RESOURCE:
		return float64(Z_RES_HANDLE_P(op))
	case IS_LONG:
		return float64(op.GetLval())
	case IS_DOUBLE:
		return op.GetDval()
	case IS_STRING:
		return ZendStrtod(Z_STRVAL_P(op), nil)
	case IS_ARRAY:
		if Z_ARRVAL_P(op).GetNNumOfElements() {
			return 1.0
		} else {
			return 0.0
		}
	case IS_OBJECT:
		var dst Zval
		ConvertObjectToType(op, &dst, IS_DOUBLE, ConvertToDouble)
		if dst.IsType(IS_DOUBLE) {
			return dst.GetDval()
		} else {
			return 1.0
		}
	case IS_REFERENCE:
		op = Z_REFVAL_P(op)
		goto try_again
	default:
		break
	}
	return 0.0
}
func __zvalGetStringFunc(op *Zval, try ZendBool) *ZendString {
try_again:
	switch op.GetType() {
	case IS_UNDEF:

	case IS_NULL:

	case IS_FALSE:
		return ZSTR_EMPTY_ALLOC()
	case IS_TRUE:
		return ZSTR_CHAR('1')
	case IS_RESOURCE:
		return ZendStrpprintf(0, "Resource id #"+ZEND_LONG_FMT, ZendLong(Z_RES_HANDLE_P(op)))
	case IS_LONG:
		return ZendLongToStr(op.GetLval())
	case IS_DOUBLE:
		return ZendStrpprintf(0, "%.*G", int(ExecutorGlobals.GetPrecision()), op.GetDval())
	case IS_ARRAY:
		ZendError(E_NOTICE, "Array to string conversion")
		if try != 0 && ExecutorGlobals.GetException() != nil {
			return nil
		} else {
			return ZSTR_KNOWN(ZEND_STR_ARRAY_CAPITALIZED)
		}
	case IS_OBJECT:
		var tmp Zval
		if Z_OBJ_HT_P(op).GetCastObject() != nil {
			if Z_OBJ_HT_P(op).GetCastObject()(op, &tmp, IS_STRING) == SUCCESS {
				return tmp.GetStr()
			}
		} else if Z_OBJ_HT_P(op).GetGet() != nil {
			var z *Zval = Z_OBJ_HT_P(op).GetGet()(op, &tmp)
			if z.GetType() != IS_OBJECT {
				var str *ZendString = b.CondF(try != 0, func() *ZendString { return ZvalTryGetString(z) }, func() *ZendString { return ZvalGetString(z) })
				ZvalPtrDtor(z)
				return str
			}
			ZvalPtrDtor(z)
		}
		if ExecutorGlobals.GetException() == nil {
			ZendThrowError(nil, "Object of class %s could not be converted to string", Z_OBJCE_P(op).GetName().GetVal())
		}
		if try != 0 {
			return nil
		} else {
			return ZSTR_EMPTY_ALLOC()
		}
	case IS_REFERENCE:
		op = Z_REFVAL_P(op)
		goto try_again
	case IS_STRING:
		return ZendStringCopy(op.GetStr())
	default:
		break
	}
	return nil
}
func ZvalGetStringFunc(op *Zval) *ZendString    { return __zvalGetStringFunc(op, 0) }
func ZvalTryGetStringFunc(op *Zval) *ZendString { return __zvalGetStringFunc(op, 1) }
func AddFunctionArray(result *Zval, op1 *Zval, op2 *Zval) {
	if result == op1 && op1.GetArr() == op2.GetArr() {

		/* $a += $a */

		return

		/* $a += $a */

	}
	if result != op1 {
		ZVAL_ARR(result, ZendArrayDup(op1.GetArr()))
	} else {
		SEPARATE_ARRAY(result)
	}
	ZendHashMerge(result.GetArr(), op2.GetArr(), ZvalAddRef, 0)
}
func AddFunctionFast(result *Zval, op1 *Zval, op2 *Zval) int {
	var type_pair ZendUchar = TYPE_PAIR(op1.GetType(), op2.GetType())
	if type_pair == TYPE_PAIR(IS_LONG, IS_LONG) {
		FastLongAddFunction(result, op1, op2)
		return SUCCESS
	} else if type_pair == TYPE_PAIR(IS_DOUBLE, IS_DOUBLE) {
		ZVAL_DOUBLE(result, op1.GetDval()+op2.GetDval())
		return SUCCESS
	} else if type_pair == TYPE_PAIR(IS_LONG, IS_DOUBLE) {
		ZVAL_DOUBLE(result, float64(op1.GetLval())+op2.GetDval())
		return SUCCESS
	} else if type_pair == TYPE_PAIR(IS_DOUBLE, IS_LONG) {
		ZVAL_DOUBLE(result, op1.GetDval()+float64(op2.GetLval()))
		return SUCCESS
	} else if type_pair == TYPE_PAIR(IS_ARRAY, IS_ARRAY) {
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
		if Z_ISREF_P(op1) {
			op1 = Z_REFVAL_P(op1)
		} else if Z_ISREF_P(op2) {
			op2 = Z_REFVAL_P(op2)
		} else if converted == 0 {
			if op1.IsType(IS_OBJECT) && op1 == result && Z_OBJ_HT(*op1).GetGet() != nil && Z_OBJ_HT(*op1).GetSet() != nil {
				var ret int
				var rv Zval
				var objval *Zval = Z_OBJ_HT(*op1).GetGet()(op1, &rv)
				Z_TRY_ADDREF_P(objval)
				ret = AddFunction(objval, objval, op2)
				Z_OBJ_HT(*op1).GetSet()(op1, objval)
				ZvalPtrDtor(objval)
				return ret
			} else if op1.IsType(IS_OBJECT) && Z_OBJ_HT(*op1).GetDoOperation() != nil {
				if SUCCESS == Z_OBJ_HT(*op1).GetDoOperation()(ZEND_ADD, result, op1, op2) {
					return SUCCESS
				}
			} else if op2.IsType(IS_OBJECT) && Z_OBJ_HT(*op2).GetDoOperation() != nil && SUCCESS == Z_OBJ_HT(*op2).GetDoOperation()(ZEND_ADD, result, op1, op2) {
				return SUCCESS
			}
			if op1 != op2 {
				op1 = ZendiConvertScalarToNumber(op1, &op1_copy, result, 0)
				op2 = ZendiConvertScalarToNumber(op2, &op2_copy, result, 0)
			} else {
				op1 = ZendiConvertScalarToNumber(op1, &op1_copy, result, 0)
				op2 = op1
			}
			if ExecutorGlobals.GetException() != nil {
				if result != op1 {
					ZVAL_UNDEF(result)
				}
				return FAILURE
			}
			converted = 1
		} else {
			if result != op1 {
				ZVAL_UNDEF(result)
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
func SubFunctionFast(result *Zval, op1 *Zval, op2 *Zval) int {
	var type_pair ZendUchar = TYPE_PAIR(op1.GetType(), op2.GetType())
	if type_pair == TYPE_PAIR(IS_LONG, IS_LONG) {
		FastLongSubFunction(result, op1, op2)
		return SUCCESS
	} else if type_pair == TYPE_PAIR(IS_DOUBLE, IS_DOUBLE) {
		ZVAL_DOUBLE(result, op1.GetDval()-op2.GetDval())
		return SUCCESS
	} else if type_pair == TYPE_PAIR(IS_LONG, IS_DOUBLE) {
		ZVAL_DOUBLE(result, float64(op1.GetLval())-op2.GetDval())
		return SUCCESS
	} else if type_pair == TYPE_PAIR(IS_DOUBLE, IS_LONG) {
		ZVAL_DOUBLE(result, op1.GetDval()-float64(op2.GetLval()))
		return SUCCESS
	} else {
		return FAILURE
	}
}
func SubFunctionSlow(result *Zval, op1 *Zval, op2 *Zval) int {
	var op1_copy Zval
	var op2_copy Zval
	var converted int = 0
	for true {
		if Z_ISREF_P(op1) {
			op1 = Z_REFVAL_P(op1)
		} else if Z_ISREF_P(op2) {
			op2 = Z_REFVAL_P(op2)
		} else if converted == 0 {
			if op1.IsType(IS_OBJECT) && op1 == result && Z_OBJ_HT(*op1).GetGet() != nil && Z_OBJ_HT(*op1).GetSet() != nil {
				var ret int
				var rv Zval
				var objval *Zval = Z_OBJ_HT(*op1).GetGet()(op1, &rv)
				Z_TRY_ADDREF_P(objval)
				ret = SubFunction(objval, objval, op2)
				Z_OBJ_HT(*op1).GetSet()(op1, objval)
				ZvalPtrDtor(objval)
				return ret
			} else if op1.IsType(IS_OBJECT) && Z_OBJ_HT(*op1).GetDoOperation() != nil {
				if SUCCESS == Z_OBJ_HT(*op1).GetDoOperation()(ZEND_SUB, result, op1, op2) {
					return SUCCESS
				}
			} else if op2.IsType(IS_OBJECT) && Z_OBJ_HT(*op2).GetDoOperation() != nil && SUCCESS == Z_OBJ_HT(*op2).GetDoOperation()(ZEND_SUB, result, op1, op2) {
				return SUCCESS
			}
			if op1 != op2 {
				op1 = ZendiConvertScalarToNumber(op1, &op1_copy, result, 0)
				op2 = ZendiConvertScalarToNumber(op2, &op2_copy, result, 0)
			} else {
				op1 = ZendiConvertScalarToNumber(op1, &op1_copy, result, 0)
				op2 = op1
			}
			if ExecutorGlobals.GetException() != nil {
				if result != op1 {
					ZVAL_UNDEF(result)
				}
				return FAILURE
			}
			converted = 1
		} else {
			if result != op1 {
				ZVAL_UNDEF(result)
			}
			ZendThrowError(nil, "Unsupported operand types")
			return FAILURE
		}
		if SubFunctionFast(result, op1, op2) == SUCCESS {
			return SUCCESS
		}
	}
}
func SubFunction(result *Zval, op1 *Zval, op2 *Zval) int {
	if SubFunctionFast(result, op1, op2) == SUCCESS {
		return SUCCESS
	} else {
		return SubFunctionSlow(result, op1, op2)
	}
}
func MulFunction(result *Zval, op1 *Zval, op2 *Zval) int {
	var op1_copy Zval
	var op2_copy Zval
	var converted int = 0
	for true {
		var type_pair ZendUchar = TYPE_PAIR(op1.GetType(), op2.GetType())
		if type_pair == TYPE_PAIR(IS_LONG, IS_LONG) {
			var overflow ZendLong
			ZEND_SIGNED_MULTIPLY_LONG(op1.GetLval(), op2.GetLval(), result.GetLval(), result.GetDval(), overflow)
			if overflow != 0 {
				result.SetTypeInfo(IS_DOUBLE)
			} else {
				result.SetTypeInfo(IS_LONG)
			}
			return SUCCESS
		} else if type_pair == TYPE_PAIR(IS_DOUBLE, IS_DOUBLE) {
			ZVAL_DOUBLE(result, op1.GetDval()*op2.GetDval())
			return SUCCESS
		} else if type_pair == TYPE_PAIR(IS_LONG, IS_DOUBLE) {
			ZVAL_DOUBLE(result, float64(op1.GetLval())*op2.GetDval())
			return SUCCESS
		} else if type_pair == TYPE_PAIR(IS_DOUBLE, IS_LONG) {
			ZVAL_DOUBLE(result, op1.GetDval()*float64(op2.GetLval()))
			return SUCCESS
		} else {
			if Z_ISREF_P(op1) {
				op1 = Z_REFVAL_P(op1)
			} else if Z_ISREF_P(op2) {
				op2 = Z_REFVAL_P(op2)
			} else if converted == 0 {
				if op1.IsType(IS_OBJECT) && op1 == result && Z_OBJ_HT(*op1).GetGet() != nil && Z_OBJ_HT(*op1).GetSet() != nil {
					var ret int
					var rv Zval
					var objval *Zval = Z_OBJ_HT(*op1).GetGet()(op1, &rv)
					Z_TRY_ADDREF_P(objval)
					ret = MulFunction(objval, objval, op2)
					Z_OBJ_HT(*op1).GetSet()(op1, objval)
					ZvalPtrDtor(objval)
					return ret
				} else if op1.IsType(IS_OBJECT) && Z_OBJ_HT(*op1).GetDoOperation() != nil {
					if SUCCESS == Z_OBJ_HT(*op1).GetDoOperation()(ZEND_MUL, result, op1, op2) {
						return SUCCESS
					}
				} else if op2.IsType(IS_OBJECT) && Z_OBJ_HT(*op2).GetDoOperation() != nil && SUCCESS == Z_OBJ_HT(*op2).GetDoOperation()(ZEND_MUL, result, op1, op2) {
					return SUCCESS
				}
				if op1 != op2 {
					op1 = ZendiConvertScalarToNumber(op1, &op1_copy, result, 0)
					op2 = ZendiConvertScalarToNumber(op2, &op2_copy, result, 0)
				} else {
					op1 = ZendiConvertScalarToNumber(op1, &op1_copy, result, 0)
					op2 = op1
				}
				if ExecutorGlobals.GetException() != nil {
					if result != op1 {
						ZVAL_UNDEF(result)
					}
					return FAILURE
				}
				converted = 1
			} else {
				if result != op1 {
					ZVAL_UNDEF(result)
				}
				ZendThrowError(nil, "Unsupported operand types")
				return FAILURE
			}
		}
	}
}
func PowFunction(result *Zval, op1 *Zval, op2 *Zval) int {
	var op1_copy Zval
	var op2_copy Zval
	var converted int = 0
	for true {
		var type_pair ZendUchar = TYPE_PAIR(op1.GetType(), op2.GetType())
		if type_pair == TYPE_PAIR(IS_LONG, IS_LONG) {
			if op2.GetLval() >= 0 {
				var l1 ZendLong = 1
				var l2 ZendLong = op1.GetLval()
				var i ZendLong = op2.GetLval()
				if i == 0 {
					ZVAL_LONG(result, 1)
					return SUCCESS
				} else if l2 == 0 {
					ZVAL_LONG(result, 0)
					return SUCCESS
				}
				for i >= 1 {
					var overflow ZendLong
					var dval float64 = 0.0
					if i%2 != 0 {
						i--
						ZEND_SIGNED_MULTIPLY_LONG(l1, l2, l1, dval, overflow)
						if overflow != 0 {
							ZVAL_DOUBLE(result, dval*pow(l2, i))
							return SUCCESS
						}
					} else {
						i /= 2
						ZEND_SIGNED_MULTIPLY_LONG(l2, l2, l2, dval, overflow)
						if overflow != 0 {
							ZVAL_DOUBLE(result, float64(l1*pow(dval, i)))
							return SUCCESS
						}
					}
				}

				/* i == 0 */

				ZVAL_LONG(result, l1)

				/* i == 0 */

			} else {
				ZVAL_DOUBLE(result, pow(float64(op1.GetLval()), float64(op2.GetLval())))
			}
			return SUCCESS
		} else if type_pair == TYPE_PAIR(IS_DOUBLE, IS_DOUBLE) {
			ZVAL_DOUBLE(result, pow(op1.GetDval(), op2.GetDval()))
			return SUCCESS
		} else if type_pair == TYPE_PAIR(IS_LONG, IS_DOUBLE) {
			ZVAL_DOUBLE(result, pow(float64(op1.GetLval()), op2.GetDval()))
			return SUCCESS
		} else if type_pair == TYPE_PAIR(IS_DOUBLE, IS_LONG) {
			ZVAL_DOUBLE(result, pow(op1.GetDval(), float64(op2.GetLval())))
			return SUCCESS
		} else {
			if Z_ISREF_P(op1) {
				op1 = Z_REFVAL_P(op1)
			} else if Z_ISREF_P(op2) {
				op2 = Z_REFVAL_P(op2)
			} else if converted == 0 {
				if op1.IsType(IS_OBJECT) && op1 == result && Z_OBJ_HT(*op1).GetGet() != nil && Z_OBJ_HT(*op1).GetSet() != nil {
					var ret int
					var rv Zval
					var objval *Zval = Z_OBJ_HT(*op1).GetGet()(op1, &rv)
					Z_TRY_ADDREF_P(objval)
					ret = PowFunction(objval, objval, op2)
					Z_OBJ_HT(*op1).GetSet()(op1, objval)
					ZvalPtrDtor(objval)
					return ret
				} else if op1.IsType(IS_OBJECT) && Z_OBJ_HT(*op1).GetDoOperation() != nil {
					if SUCCESS == Z_OBJ_HT(*op1).GetDoOperation()(ZEND_POW, result, op1, op2) {
						return SUCCESS
					}
				} else if op2.IsType(IS_OBJECT) && Z_OBJ_HT(*op2).GetDoOperation() != nil && SUCCESS == Z_OBJ_HT(*op2).GetDoOperation()(ZEND_POW, result, op1, op2) {
					return SUCCESS
				}
				if op1 != op2 {
					if op1.IsType(IS_ARRAY) {
						if op1 == result {
							ZvalPtrDtor(result)
						}
						ZVAL_LONG(result, 0)
						return SUCCESS
					} else {
						op1 = ZendiConvertScalarToNumber(op1, &op1_copy, result, 0)
					}
					if op2.IsType(IS_ARRAY) {
						if op1 == result {
							ZvalPtrDtor(result)
						}
						ZVAL_LONG(result, 1)
						return SUCCESS
					} else {
						op2 = ZendiConvertScalarToNumber(op2, &op2_copy, result, 0)
					}
				} else {
					if op1.IsType(IS_ARRAY) {
						if op1 == result {
							ZvalPtrDtor(result)
						}
						ZVAL_LONG(result, 0)
						return SUCCESS
					} else {
						op1 = ZendiConvertScalarToNumber(op1, &op1_copy, result, 0)
					}
					op2 = op1
				}
				if ExecutorGlobals.GetException() != nil {
					if result != op1 {
						ZVAL_UNDEF(result)
					}
					return FAILURE
				}
				converted = 1
			} else {
				if result != op1 {
					ZVAL_UNDEF(result)
				}
				ZendThrowError(nil, "Unsupported operand types")
				return FAILURE
			}
		}
	}
}
func DivFunction(result *Zval, op1 *Zval, op2 *Zval) int {
	var op1_copy Zval
	var op2_copy Zval
	var converted int = 0
	for true {
		var type_pair ZendUchar = TYPE_PAIR(op1.GetType(), op2.GetType())
		if type_pair == TYPE_PAIR(IS_LONG, IS_LONG) {
			if op2.GetLval() == 0 {
				ZendError(E_WARNING, "Division by zero")
				ZVAL_DOUBLE(result, float64(op1.GetLval()/float64(op2.GetLval())))
				return SUCCESS
			} else if op2.GetLval() == -1 && op1.GetLval() == ZEND_LONG_MIN {

				/* Prevent overflow error/crash */

				ZVAL_DOUBLE(result, float64(ZEND_LONG_MIN/-1))
				return SUCCESS
			}
			if op1.GetLval()%op2.GetLval() == 0 {
				ZVAL_LONG(result, op1.GetLval()/op2.GetLval())
			} else {
				ZVAL_DOUBLE(result, float64(op1.GetLval())/op2.GetLval())
			}
			return SUCCESS
		} else if type_pair == TYPE_PAIR(IS_DOUBLE, IS_DOUBLE) {
			if op2.GetDval() == 0 {
				ZendError(E_WARNING, "Division by zero")
			}
			ZVAL_DOUBLE(result, op1.GetDval()/op2.GetDval())
			return SUCCESS
		} else if type_pair == TYPE_PAIR(IS_DOUBLE, IS_LONG) {
			if op2.GetLval() == 0 {
				ZendError(E_WARNING, "Division by zero")
			}
			ZVAL_DOUBLE(result, op1.GetDval()/float64(op2.GetLval()))
			return SUCCESS
		} else if type_pair == TYPE_PAIR(IS_LONG, IS_DOUBLE) {
			if op2.GetDval() == 0 {
				ZendError(E_WARNING, "Division by zero")
			}
			ZVAL_DOUBLE(result, float64(op1.GetLval()/op2.GetDval()))
			return SUCCESS
		} else {
			if Z_ISREF_P(op1) {
				op1 = Z_REFVAL_P(op1)
			} else if Z_ISREF_P(op2) {
				op2 = Z_REFVAL_P(op2)
			} else if converted == 0 {
				if op1.IsType(IS_OBJECT) && op1 == result && Z_OBJ_HT(*op1).GetGet() != nil && Z_OBJ_HT(*op1).GetSet() != nil {
					var ret int
					var rv Zval
					var objval *Zval = Z_OBJ_HT(*op1).GetGet()(op1, &rv)
					Z_TRY_ADDREF_P(objval)
					ret = DivFunction(objval, objval, op2)
					Z_OBJ_HT(*op1).GetSet()(op1, objval)
					ZvalPtrDtor(objval)
					return ret
				} else if op1.IsType(IS_OBJECT) && Z_OBJ_HT(*op1).GetDoOperation() != nil {
					if SUCCESS == Z_OBJ_HT(*op1).GetDoOperation()(ZEND_DIV, result, op1, op2) {
						return SUCCESS
					}
				} else if op2.IsType(IS_OBJECT) && Z_OBJ_HT(*op2).GetDoOperation() != nil && SUCCESS == Z_OBJ_HT(*op2).GetDoOperation()(ZEND_DIV, result, op1, op2) {
					return SUCCESS
				}
				if op1 != op2 {
					op1 = ZendiConvertScalarToNumber(op1, &op1_copy, result, 0)
					op2 = ZendiConvertScalarToNumber(op2, &op2_copy, result, 0)
				} else {
					op1 = ZendiConvertScalarToNumber(op1, &op1_copy, result, 0)
					op2 = op1
				}
				if ExecutorGlobals.GetException() != nil {
					if result != op1 {
						ZVAL_UNDEF(result)
					}
					return FAILURE
				}
				converted = 1
			} else {
				if result != op1 {
					ZVAL_UNDEF(result)
				}
				ZendThrowError(nil, "Unsupported operand types")
				return FAILURE
			}
		}
	}
}
func ModFunction(result *Zval, op1 *Zval, op2 *Zval) int {
	var op1_lval ZendLong
	var op2_lval ZendLong
	for {
		if op1.GetType() != IS_LONG {
			if Z_ISREF_P(op1) {
				op1 = Z_REFVAL_P(op1)
				if op1.IsType(IS_LONG) {
					op1_lval = op1.GetLval()
					break
				}
			}
			if op1.IsType(IS_OBJECT) && op1 == result && Z_OBJ_HT(*op1).GetGet() != nil && Z_OBJ_HT(*op1).GetSet() != nil {
				var ret int
				var rv Zval
				var objval *Zval = Z_OBJ_HT(*op1).GetGet()(op1, &rv)
				Z_TRY_ADDREF_P(objval)
				ret = ModFunction(objval, objval, op2)
				Z_OBJ_HT(*op1).GetSet()(op1, objval)
				ZvalPtrDtor(objval)
				return ret
			} else if op1.IsType(IS_OBJECT) && Z_OBJ_HT(*op1).GetDoOperation() != nil {
				if SUCCESS == Z_OBJ_HT(*op1).GetDoOperation()(ZEND_MOD, result, op1, op2) {
					return SUCCESS
				}
			}
			op1_lval = _zvalGetLongFuncNoisy(op1)
			if ExecutorGlobals.GetException() != nil {
				if result != op1 {
					ZVAL_UNDEF(result)
				}
				return FAILURE
			}
		} else {
			op1_lval = op1.GetLval()
		}
		break
	}
	for {
		if op2.GetType() != IS_LONG {
			if Z_ISREF_P(op2) {
				op2 = Z_REFVAL_P(op2)
				if op2.IsType(IS_LONG) {
					op2_lval = op2.GetLval()
					break
				}
			}
			if op2.IsType(IS_OBJECT) && Z_OBJ_HT(*op2).GetDoOperation() != nil && SUCCESS == Z_OBJ_HT(*op2).GetDoOperation()(ZEND_MOD, result, op1, op2) {
				return SUCCESS
			}
			op2_lval = _zvalGetLongFuncNoisy(op2)
			if ExecutorGlobals.GetException() != nil {
				if result != op1 {
					ZVAL_UNDEF(result)
				}
				return FAILURE
			}
		} else {
			op2_lval = op2.GetLval()
		}
		break
	}
	if op2_lval == 0 {

		/* modulus by zero */

		if ExecutorGlobals.GetCurrentExecuteData() != nil && CompilerGlobals.GetInCompilation() == 0 {
			ZendThrowExceptionEx(ZendCeDivisionByZeroError, 0, "Modulo by zero")
		} else {
			ZendErrorNoreturn(E_ERROR, "Modulo by zero")
		}
		if op1 != result {
			ZVAL_UNDEF(result)
		}
		return FAILURE
	}
	if op1 == result {
		ZvalPtrDtor(result)
	}
	if op2_lval == -1 {

		/* Prevent overflow error/crash if op1==LONG_MIN */

		ZVAL_LONG(result, 0)
		return SUCCESS
	}
	ZVAL_LONG(result, op1_lval%op2_lval)
	return SUCCESS
}
func BooleanXorFunction(result *Zval, op1 *Zval, op2 *Zval) int {
	var op1_val int
	var op2_val int
	for {
		if op1.IsType(IS_FALSE) {
			op1_val = 0
		} else if op1.IsType(IS_TRUE) {
			op1_val = 1
		} else {
			if Z_ISREF_P(op1) {
				op1 = Z_REFVAL_P(op1)
				if op1.IsType(IS_FALSE) {
					op1_val = 0
					break
				} else if op1.IsType(IS_TRUE) {
					op1_val = 1
					break
				}
			}
			if op1.IsType(IS_OBJECT) && op1 == result && Z_OBJ_HT(*op1).GetGet() != nil && Z_OBJ_HT(*op1).GetSet() != nil {
				var ret int
				var rv Zval
				var objval *Zval = Z_OBJ_HT(*op1).GetGet()(op1, &rv)
				Z_TRY_ADDREF_P(objval)
				ret = BooleanXorFunction(objval, objval, op2)
				Z_OBJ_HT(*op1).GetSet()(op1, objval)
				ZvalPtrDtor(objval)
				return ret
			} else if op1.IsType(IS_OBJECT) && Z_OBJ_HT(*op1).GetDoOperation() != nil {
				if SUCCESS == Z_OBJ_HT(*op1).GetDoOperation()(ZEND_BOOL_XOR, result, op1, op2) {
					return SUCCESS
				}
			}
			op1_val = ZvalIsTrue(op1)
		}
		break
	}
	for {
		if op2.IsType(IS_FALSE) {
			op2_val = 0
		} else if op2.IsType(IS_TRUE) {
			op2_val = 1
		} else {
			if Z_ISREF_P(op2) {
				op2 = Z_REFVAL_P(op2)
				if op2.IsType(IS_FALSE) {
					op2_val = 0
					break
				} else if op2.IsType(IS_TRUE) {
					op2_val = 1
					break
				}
			}
			if op2.IsType(IS_OBJECT) && Z_OBJ_HT(*op2).GetDoOperation() != nil && SUCCESS == Z_OBJ_HT(*op2).GetDoOperation()(ZEND_BOOL_XOR, result, op1, op2) {
				return SUCCESS
			}
			op2_val = ZvalIsTrue(op2)
		}
		break
	}
	ZVAL_BOOL(result, op1_val^op2_val)
	return SUCCESS
}
func BooleanNotFunction(result *Zval, op1 *Zval) int {
	if op1.GetType() < IS_TRUE {
		ZVAL_TRUE(result)
	} else if op1.IsType(IS_TRUE) {
		ZVAL_FALSE(result)
	} else {
		if Z_ISREF_P(op1) {
			op1 = Z_REFVAL_P(op1)
			if op1.GetType() < IS_TRUE {
				ZVAL_TRUE(result)
				return SUCCESS
			} else if op1.IsType(IS_TRUE) {
				ZVAL_FALSE(result)
				return SUCCESS
			}
		}
		if op1.IsType(IS_OBJECT) && Z_OBJ_HT(*op1).GetDoOperation() != nil && SUCCESS == Z_OBJ_HT(*op1).GetDoOperation()(ZEND_BOOL_NOT, result, op1, nil) {
			return SUCCESS
		}
		ZVAL_BOOL(result, !(ZvalIsTrue(op1)))
	}
	return SUCCESS
}
func BitwiseNotFunction(result *Zval, op1 *Zval) int {
try_again:
	switch op1.GetType() {
	case IS_LONG:
		ZVAL_LONG(result, ^(op1.GetLval()))
		return SUCCESS
	case IS_DOUBLE:
		ZVAL_LONG(result, ^(ZendDvalToLval(op1.GetDval())))
		return SUCCESS
	case IS_STRING:
		var i int
		if Z_STRLEN_P(op1) == 1 {
			var not ZendUchar = ZendUchar(^((*Z_STRVAL_P)(op1)))
			ZVAL_INTERNED_STR(result, ZSTR_CHAR(not))
		} else {
			ZVAL_NEW_STR(result, ZendStringAlloc(Z_STRLEN_P(op1), 0))
			for i = 0; i < Z_STRLEN_P(op1); i++ {
				Z_STRVAL_P(result)[i] = ^(Z_STRVAL_P(op1)[i])
			}
			Z_STRVAL_P(result)[i] = 0
		}
		return SUCCESS
	case IS_REFERENCE:
		op1 = Z_REFVAL_P(op1)
		goto try_again
	default:
		if op1.IsType(IS_OBJECT) && Z_OBJ_HT(*op1).GetDoOperation() != nil && SUCCESS == Z_OBJ_HT(*op1).GetDoOperation()(ZEND_BW_NOT, result, op1, nil) {
			return SUCCESS
		}
		if result != op1 {
			ZVAL_UNDEF(result)
		}
		ZendThrowError(nil, "Unsupported operand types")
		return FAILURE
	}
}
func BitwiseOrFunction(result *Zval, op1 *Zval, op2 *Zval) int {
	var op1_lval ZendLong
	var op2_lval ZendLong
	if op1.IsType(IS_LONG) && op2.IsType(IS_LONG) {
		ZVAL_LONG(result, op1.GetLval()|op2.GetLval())
		return SUCCESS
	}
	ZVAL_DEREF(op1)
	ZVAL_DEREF(op2)
	if op1.IsType(IS_STRING) && op2.IsType(IS_STRING) {
		var longer *Zval
		var shorter *Zval
		var str *ZendString
		var i int
		if Z_STRLEN_P(op1) >= Z_STRLEN_P(op2) {
			if Z_STRLEN_P(op1) == Z_STRLEN_P(op2) && Z_STRLEN_P(op1) == 1 {
				var or ZendUchar = zend_uchar((*Z_STRVAL_P)(op1) | (*Z_STRVAL_P)(op2))
				if result == op1 {
					ZvalPtrDtorStr(result)
				}
				ZVAL_INTERNED_STR(result, ZSTR_CHAR(or))
				return SUCCESS
			}
			longer = op1
			shorter = op2
		} else {
			longer = op2
			shorter = op1
		}
		str = ZendStringAlloc(Z_STRLEN_P(longer), 0)
		for i = 0; i < Z_STRLEN_P(shorter); i++ {
			str.GetVal()[i] = Z_STRVAL_P(longer)[i] | Z_STRVAL_P(shorter)[i]
		}
		memcpy(str.GetVal()+i, Z_STRVAL_P(longer)+i, Z_STRLEN_P(longer)-i+1)
		if result == op1 {
			ZvalPtrDtorStr(result)
		}
		ZVAL_NEW_STR(result, str)
		return SUCCESS
	}
	if op1.GetType() != IS_LONG {
		if op1.IsType(IS_OBJECT) && op1 == result && Z_OBJ_HT(*op1).GetGet() != nil && Z_OBJ_HT(*op1).GetSet() != nil {
			var ret int
			var rv Zval
			var objval *Zval = Z_OBJ_HT(*op1).GetGet()(op1, &rv)
			Z_TRY_ADDREF_P(objval)
			ret = BitwiseOrFunction(objval, objval, op2)
			Z_OBJ_HT(*op1).GetSet()(op1, objval)
			ZvalPtrDtor(objval)
			return ret
		} else if op1.IsType(IS_OBJECT) && Z_OBJ_HT(*op1).GetDoOperation() != nil {
			if SUCCESS == Z_OBJ_HT(*op1).GetDoOperation()(ZEND_BW_OR, result, op1, op2) {
				return SUCCESS
			}
		}
		op1_lval = _zvalGetLongFuncNoisy(op1)
		if ExecutorGlobals.GetException() != nil {
			if result != op1 {
				ZVAL_UNDEF(result)
			}
			return FAILURE
		}
	} else {
		op1_lval = op1.GetLval()
	}
	if op2.GetType() != IS_LONG {
		if op2.IsType(IS_OBJECT) && Z_OBJ_HT(*op2).GetDoOperation() != nil && SUCCESS == Z_OBJ_HT(*op2).GetDoOperation()(ZEND_BW_OR, result, op1, op2) {
			return SUCCESS
		}
		op2_lval = _zvalGetLongFuncNoisy(op2)
		if ExecutorGlobals.GetException() != nil {
			if result != op1 {
				ZVAL_UNDEF(result)
			}
			return FAILURE
		}
	} else {
		op2_lval = op2.GetLval()
	}
	if op1 == result {
		ZvalPtrDtor(result)
	}
	ZVAL_LONG(result, op1_lval|op2_lval)
	return SUCCESS
}
func BitwiseAndFunction(result *Zval, op1 *Zval, op2 *Zval) int {
	var op1_lval ZendLong
	var op2_lval ZendLong
	if op1.IsType(IS_LONG) && op2.IsType(IS_LONG) {
		ZVAL_LONG(result, op1.GetLval()&op2.GetLval())
		return SUCCESS
	}
	ZVAL_DEREF(op1)
	ZVAL_DEREF(op2)
	if op1.IsType(IS_STRING) && op2.IsType(IS_STRING) {
		var longer *Zval
		var shorter *Zval
		var str *ZendString
		var i int
		if Z_STRLEN_P(op1) >= Z_STRLEN_P(op2) {
			if Z_STRLEN_P(op1) == Z_STRLEN_P(op2) && Z_STRLEN_P(op1) == 1 {
				var and ZendUchar = zend_uchar((*Z_STRVAL_P)(op1) & (*Z_STRVAL_P)(op2))
				if result == op1 {
					ZvalPtrDtorStr(result)
				}
				ZVAL_INTERNED_STR(result, ZSTR_CHAR(and))
				return SUCCESS
			}
			longer = op1
			shorter = op2
		} else {
			longer = op2
			shorter = op1
		}
		str = ZendStringAlloc(Z_STRLEN_P(shorter), 0)
		for i = 0; i < Z_STRLEN_P(shorter); i++ {
			str.GetVal()[i] = Z_STRVAL_P(shorter)[i] & Z_STRVAL_P(longer)[i]
		}
		str.GetVal()[i] = 0
		if result == op1 {
			ZvalPtrDtorStr(result)
		}
		ZVAL_NEW_STR(result, str)
		return SUCCESS
	}
	if op1.GetType() != IS_LONG {
		if op1.IsType(IS_OBJECT) && op1 == result && Z_OBJ_HT(*op1).GetGet() != nil && Z_OBJ_HT(*op1).GetSet() != nil {
			var ret int
			var rv Zval
			var objval *Zval = Z_OBJ_HT(*op1).GetGet()(op1, &rv)
			Z_TRY_ADDREF_P(objval)
			ret = BitwiseAndFunction(objval, objval, op2)
			Z_OBJ_HT(*op1).GetSet()(op1, objval)
			ZvalPtrDtor(objval)
			return ret
		} else if op1.IsType(IS_OBJECT) && Z_OBJ_HT(*op1).GetDoOperation() != nil {
			if SUCCESS == Z_OBJ_HT(*op1).GetDoOperation()(ZEND_BW_AND, result, op1, op2) {
				return SUCCESS
			}
		}
		op1_lval = _zvalGetLongFuncNoisy(op1)
		if ExecutorGlobals.GetException() != nil {
			if result != op1 {
				ZVAL_UNDEF(result)
			}
			return FAILURE
		}
	} else {
		op1_lval = op1.GetLval()
	}
	if op2.GetType() != IS_LONG {
		if op2.IsType(IS_OBJECT) && Z_OBJ_HT(*op2).GetDoOperation() != nil && SUCCESS == Z_OBJ_HT(*op2).GetDoOperation()(ZEND_BW_AND, result, op1, op2) {
			return SUCCESS
		}
		op2_lval = _zvalGetLongFuncNoisy(op2)
		if ExecutorGlobals.GetException() != nil {
			if result != op1 {
				ZVAL_UNDEF(result)
			}
			return FAILURE
		}
	} else {
		op2_lval = op2.GetLval()
	}
	if op1 == result {
		ZvalPtrDtor(result)
	}
	ZVAL_LONG(result, op1_lval&op2_lval)
	return SUCCESS
}
func BitwiseXorFunction(result *Zval, op1 *Zval, op2 *Zval) int {
	var op1_lval ZendLong
	var op2_lval ZendLong
	if op1.IsType(IS_LONG) && op2.IsType(IS_LONG) {
		ZVAL_LONG(result, op1.GetLval()^op2.GetLval())
		return SUCCESS
	}
	ZVAL_DEREF(op1)
	ZVAL_DEREF(op2)
	if op1.IsType(IS_STRING) && op2.IsType(IS_STRING) {
		var longer *Zval
		var shorter *Zval
		var str *ZendString
		var i int
		if Z_STRLEN_P(op1) >= Z_STRLEN_P(op2) {
			if Z_STRLEN_P(op1) == Z_STRLEN_P(op2) && Z_STRLEN_P(op1) == 1 {
				var xor ZendUchar = zend_uchar((*Z_STRVAL_P)(op1) ^ (*Z_STRVAL_P)(op2))
				if result == op1 {
					ZvalPtrDtorStr(result)
				}
				ZVAL_INTERNED_STR(result, ZSTR_CHAR(xor))
				return SUCCESS
			}
			longer = op1
			shorter = op2
		} else {
			longer = op2
			shorter = op1
		}
		str = ZendStringAlloc(Z_STRLEN_P(shorter), 0)
		for i = 0; i < Z_STRLEN_P(shorter); i++ {
			str.GetVal()[i] = Z_STRVAL_P(shorter)[i] ^ Z_STRVAL_P(longer)[i]
		}
		str.GetVal()[i] = 0
		if result == op1 {
			ZvalPtrDtorStr(result)
		}
		ZVAL_NEW_STR(result, str)
		return SUCCESS
	}
	if op1.GetType() != IS_LONG {
		if op1.IsType(IS_OBJECT) && op1 == result && Z_OBJ_HT(*op1).GetGet() != nil && Z_OBJ_HT(*op1).GetSet() != nil {
			var ret int
			var rv Zval
			var objval *Zval = Z_OBJ_HT(*op1).GetGet()(op1, &rv)
			Z_TRY_ADDREF_P(objval)
			ret = BitwiseXorFunction(objval, objval, op2)
			Z_OBJ_HT(*op1).GetSet()(op1, objval)
			ZvalPtrDtor(objval)
			return ret
		} else if op1.IsType(IS_OBJECT) && Z_OBJ_HT(*op1).GetDoOperation() != nil {
			if SUCCESS == Z_OBJ_HT(*op1).GetDoOperation()(ZEND_BW_XOR, result, op1, op2) {
				return SUCCESS
			}
		}
		op1_lval = _zvalGetLongFuncNoisy(op1)
		if ExecutorGlobals.GetException() != nil {
			if result != op1 {
				ZVAL_UNDEF(result)
			}
			return FAILURE
		}
	} else {
		op1_lval = op1.GetLval()
	}
	if op2.GetType() != IS_LONG {
		if op2.IsType(IS_OBJECT) && Z_OBJ_HT(*op2).GetDoOperation() != nil && SUCCESS == Z_OBJ_HT(*op2).GetDoOperation()(ZEND_BW_XOR, result, op1, op2) {
			return SUCCESS
		}
		op2_lval = _zvalGetLongFuncNoisy(op2)
		if ExecutorGlobals.GetException() != nil {
			if result != op1 {
				ZVAL_UNDEF(result)
			}
			return FAILURE
		}
	} else {
		op2_lval = op2.GetLval()
	}
	if op1 == result {
		ZvalPtrDtor(result)
	}
	ZVAL_LONG(result, op1_lval^op2_lval)
	return SUCCESS
}
func ShiftLeftFunction(result *Zval, op1 *Zval, op2 *Zval) int {
	var op1_lval ZendLong
	var op2_lval ZendLong
	for {
		if op1.GetType() != IS_LONG {
			if Z_ISREF_P(op1) {
				op1 = Z_REFVAL_P(op1)
				if op1.IsType(IS_LONG) {
					op1_lval = op1.GetLval()
					break
				}
			}
			if op1.IsType(IS_OBJECT) && op1 == result && Z_OBJ_HT(*op1).GetGet() != nil && Z_OBJ_HT(*op1).GetSet() != nil {
				var ret int
				var rv Zval
				var objval *Zval = Z_OBJ_HT(*op1).GetGet()(op1, &rv)
				Z_TRY_ADDREF_P(objval)
				ret = ShiftLeftFunction(objval, objval, op2)
				Z_OBJ_HT(*op1).GetSet()(op1, objval)
				ZvalPtrDtor(objval)
				return ret
			} else if op1.IsType(IS_OBJECT) && Z_OBJ_HT(*op1).GetDoOperation() != nil {
				if SUCCESS == Z_OBJ_HT(*op1).GetDoOperation()(ZEND_SL, result, op1, op2) {
					return SUCCESS
				}
			}
			op1_lval = _zvalGetLongFuncNoisy(op1)
			if ExecutorGlobals.GetException() != nil {
				if result != op1 {
					ZVAL_UNDEF(result)
				}
				return FAILURE
			}
		} else {
			op1_lval = op1.GetLval()
		}
		break
	}
	for {
		if op2.GetType() != IS_LONG {
			if Z_ISREF_P(op2) {
				op2 = Z_REFVAL_P(op2)
				if op2.IsType(IS_LONG) {
					op2_lval = op2.GetLval()
					break
				}
			}
			if op2.IsType(IS_OBJECT) && Z_OBJ_HT(*op2).GetDoOperation() != nil && SUCCESS == Z_OBJ_HT(*op2).GetDoOperation()(ZEND_SL, result, op1, op2) {
				return SUCCESS
			}
			op2_lval = _zvalGetLongFuncNoisy(op2)
			if ExecutorGlobals.GetException() != nil {
				if result != op1 {
					ZVAL_UNDEF(result)
				}
				return FAILURE
			}
		} else {
			op2_lval = op2.GetLval()
		}
		break
	}

	/* prevent wrapping quirkiness on some processors where << 64 + x == << x */

	if ZendUlong(op2_lval >= SIZEOF_ZEND_LONG*8) != 0 {
		if op2_lval > 0 {
			if op1 == result {
				ZvalPtrDtor(result)
			}
			ZVAL_LONG(result, 0)
			return SUCCESS
		} else {
			if ExecutorGlobals.GetCurrentExecuteData() != nil && CompilerGlobals.GetInCompilation() == 0 {
				ZendThrowExceptionEx(ZendCeArithmeticError, 0, "Bit shift by negative number")
			} else {
				ZendErrorNoreturn(E_ERROR, "Bit shift by negative number")
			}
			if op1 != result {
				ZVAL_UNDEF(result)
			}
			return FAILURE
		}
	}
	if op1 == result {
		ZvalPtrDtor(result)
	}

	/* Perform shift on unsigned numbers to get well-defined wrap behavior. */

	ZVAL_LONG(result, zend_long(ZendUlong(op1_lval<<op2_lval)))
	return SUCCESS
}
func ShiftRightFunction(result *Zval, op1 *Zval, op2 *Zval) int {
	var op1_lval ZendLong
	var op2_lval ZendLong
	for {
		if op1.GetType() != IS_LONG {
			if Z_ISREF_P(op1) {
				op1 = Z_REFVAL_P(op1)
				if op1.IsType(IS_LONG) {
					op1_lval = op1.GetLval()
					break
				}
			}
			if op1.IsType(IS_OBJECT) && op1 == result && Z_OBJ_HT(*op1).GetGet() != nil && Z_OBJ_HT(*op1).GetSet() != nil {
				var ret int
				var rv Zval
				var objval *Zval = Z_OBJ_HT(*op1).GetGet()(op1, &rv)
				Z_TRY_ADDREF_P(objval)
				ret = ShiftRightFunction(objval, objval, op2)
				Z_OBJ_HT(*op1).GetSet()(op1, objval)
				ZvalPtrDtor(objval)
				return ret
			} else if op1.IsType(IS_OBJECT) && Z_OBJ_HT(*op1).GetDoOperation() != nil {
				if SUCCESS == Z_OBJ_HT(*op1).GetDoOperation()(ZEND_SR, result, op1, op2) {
					return SUCCESS
				}
			}
			op1_lval = _zvalGetLongFuncNoisy(op1)
			if ExecutorGlobals.GetException() != nil {
				if result != op1 {
					ZVAL_UNDEF(result)
				}
				return FAILURE
			}
		} else {
			op1_lval = op1.GetLval()
		}
		break
	}
	for {
		if op2.GetType() != IS_LONG {
			if Z_ISREF_P(op2) {
				op2 = Z_REFVAL_P(op2)
				if op2.IsType(IS_LONG) {
					op2_lval = op2.GetLval()
					break
				}
			}
			if op2.IsType(IS_OBJECT) && Z_OBJ_HT(*op2).GetDoOperation() != nil && SUCCESS == Z_OBJ_HT(*op2).GetDoOperation()(ZEND_SR, result, op1, op2) {
				return SUCCESS
			}
			op2_lval = _zvalGetLongFuncNoisy(op2)
			if ExecutorGlobals.GetException() != nil {
				if result != op1 {
					ZVAL_UNDEF(result)
				}
				return FAILURE
			}
		} else {
			op2_lval = op2.GetLval()
		}
		break
	}

	/* prevent wrapping quirkiness on some processors where >> 64 + x == >> x */

	if ZendUlong(op2_lval >= SIZEOF_ZEND_LONG*8) != 0 {
		if op2_lval > 0 {
			if op1 == result {
				ZvalPtrDtor(result)
			}
			ZVAL_LONG(result, b.Cond(op1_lval < 0, -1, 0))
			return SUCCESS
		} else {
			if ExecutorGlobals.GetCurrentExecuteData() != nil && CompilerGlobals.GetInCompilation() == 0 {
				ZendThrowExceptionEx(ZendCeArithmeticError, 0, "Bit shift by negative number")
			} else {
				ZendErrorNoreturn(E_ERROR, "Bit shift by negative number")
			}
			if op1 != result {
				ZVAL_UNDEF(result)
			}
			return FAILURE
		}
	}
	if op1 == result {
		ZvalPtrDtor(result)
	}
	ZVAL_LONG(result, op1_lval>>op2_lval)
	return SUCCESS
}
func ConcatFunction(result *Zval, op1 *Zval, op2 *Zval) int {
	var orig_op1 *Zval = op1
	var op1_copy Zval
	var op2_copy Zval
	ZVAL_UNDEF(&op1_copy)
	ZVAL_UNDEF(&op2_copy)
	for {
		if op1.GetType() != IS_STRING {
			if Z_ISREF_P(op1) {
				op1 = Z_REFVAL_P(op1)
				if op1.IsType(IS_STRING) {
					break
				}
			}
			if op1.IsType(IS_OBJECT) && op1 == result && Z_OBJ_HT(*op1).GetGet() != nil && Z_OBJ_HT(*op1).GetSet() != nil {
				var ret int
				var rv Zval
				var objval *Zval = Z_OBJ_HT(*op1).GetGet()(op1, &rv)
				Z_TRY_ADDREF_P(objval)
				ret = ConcatFunction(objval, objval, op2)
				Z_OBJ_HT(*op1).GetSet()(op1, objval)
				ZvalPtrDtor(objval)
				return ret
			} else if op1.IsType(IS_OBJECT) && Z_OBJ_HT(*op1).GetDoOperation() != nil {
				if SUCCESS == Z_OBJ_HT(*op1).GetDoOperation()(ZEND_CONCAT, result, op1, op2) {
					return SUCCESS
				}
			} else if op2.IsType(IS_OBJECT) && Z_OBJ_HT(*op2).GetDoOperation() != nil && SUCCESS == Z_OBJ_HT(*op2).GetDoOperation()(ZEND_CONCAT, result, op1, op2) {
				return SUCCESS
			}
			ZVAL_STR(&op1_copy, ZvalGetStringFunc(op1))
			if ExecutorGlobals.GetException() != nil {
				ZvalPtrDtorStr(&op1_copy)
				if orig_op1 != result {
					ZVAL_UNDEF(result)
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
		if op2.GetType() != IS_STRING {
			if Z_ISREF_P(op2) {
				op2 = Z_REFVAL_P(op2)
				if op2.IsType(IS_STRING) {
					break
				}
			}
			if op2.IsType(IS_OBJECT) && Z_OBJ_HT(*op2).GetDoOperation() != nil && SUCCESS == Z_OBJ_HT(*op2).GetDoOperation()(ZEND_CONCAT, result, op1, op2) {
				return SUCCESS
			}
			ZVAL_STR(&op2_copy, ZvalGetStringFunc(op2))
			if ExecutorGlobals.GetException() != nil {
				ZvalPtrDtorStr(&op1_copy)
				ZvalPtrDtorStr(&op2_copy)
				if orig_op1 != result {
					ZVAL_UNDEF(result)
				}
				return FAILURE
			}
			op2 = &op2_copy
		}
		break
	}
	if Z_STRLEN_P(op1) == 0 {
		if result != op2 {
			if result == orig_op1 {
				IZvalPtrDtor(result)
			}
			ZVAL_COPY(result, op2)
		}
	} else if Z_STRLEN_P(op2) == 0 {
		if result != op1 {
			if result == orig_op1 {
				IZvalPtrDtor(result)
			}
			ZVAL_COPY(result, op1)
		}
	} else {
		var op1_len int = Z_STRLEN_P(op1)
		var op2_len int = Z_STRLEN_P(op2)
		var result_len int = op1_len + op2_len
		var result_str *ZendString
		if op1_len > ZSTR_MAX_LEN-op2_len {
			ZendThrowError(nil, "String size overflow")
			ZvalPtrDtorStr(&op1_copy)
			ZvalPtrDtorStr(&op2_copy)
			if orig_op1 != result {
				ZVAL_UNDEF(result)
			}
			return FAILURE
		}
		if result == op1 && Z_REFCOUNTED_P(result) {

			/* special case, perform operations on result */

			result_str = ZendStringExtend(result.GetStr(), result_len, 0)

			/* special case, perform operations on result */

		} else {
			result_str = ZendStringAlloc(result_len, 0)
			memcpy(result_str.GetVal(), Z_STRVAL_P(op1), op1_len)
			if result == orig_op1 {
				IZvalPtrDtor(result)
			}
		}

		/* This has to happen first to account for the cases where result == op1 == op2 and
		 * the realloc is done. In this case this line will also update Z_STRVAL_P(op2) to
		 * point to the new string. The first op2_len bytes of result will still be the same. */

		ZVAL_NEW_STR(result, result_str)
		memcpy(result_str.GetVal()+op1_len, Z_STRVAL_P(op2), op2_len)
		result_str.GetVal()[result_len] = '0'
	}
	ZvalPtrDtorStr(&op1_copy)
	ZvalPtrDtorStr(&op2_copy)
	return SUCCESS
}
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
func StringCompareFunction(op1 *Zval, op2 *Zval) int {
	if op1.IsType(IS_STRING) && op2.IsType(IS_STRING) {
		if op1.GetStr() == op2.GetStr() {
			return 0
		} else {
			return ZendBinaryStrcmp(Z_STRVAL_P(op1), Z_STRLEN_P(op1), Z_STRVAL_P(op2), Z_STRLEN_P(op2))
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
func StringCaseCompareFunction(op1 *Zval, op2 *Zval) int {
	if op1.IsType(IS_STRING) && op2.IsType(IS_STRING) {
		if op1.GetStr() == op2.GetStr() {
			return 0
		} else {
			return ZendBinaryStrcasecmpL(Z_STRVAL_P(op1), Z_STRLEN_P(op1), Z_STRVAL_P(op2), Z_STRLEN_P(op2))
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
func NumericCompareFunction(op1 *Zval, op2 *Zval) int {
	var d1 float64
	var d2 float64
	d1 = ZvalGetDouble(op1)
	d2 = ZvalGetDouble(op2)
	return ZEND_NORMALIZE_BOOL(d1 - d2)
}
func ZendFreeObjGetResult(op *Zval) {
	ZEND_ASSERT(!(Z_REFCOUNTED_P(op)) || Z_REFCOUNT_P(op) != 0)
	ZvalPtrDtor(op)
}
func ConvertCompareResultToLong(result *Zval) {
	if result.IsType(IS_DOUBLE) {
		ZVAL_LONG(result, ZEND_NORMALIZE_BOOL(result.GetDval()))
	} else {
		ConvertToLong(result)
	}
}
func CompareFunction(result *Zval, op1 *Zval, op2 *Zval) int {
	var ret int
	var converted int = 0
	var op1_copy Zval
	var op2_copy Zval
	var op_free *Zval
	var tmp_free Zval
	for true {
		switch TYPE_PAIR(op1.GetType(), op2.GetType()) {
		case TYPE_PAIR(IS_LONG, IS_LONG):
			ZVAL_LONG(result, b.CondF2(op1.GetLval() > op2.GetLval(), 1, func() int {
				if op1.GetLval() < op2.GetLval() {
					return -1
				} else {
					return 0
				}
			}))
			return SUCCESS
		case TYPE_PAIR(IS_DOUBLE, IS_LONG):
			result.SetDval(op1.GetDval() - float64(op2.GetLval()))
			ZVAL_LONG(result, ZEND_NORMALIZE_BOOL(result.GetDval()))
			return SUCCESS
		case TYPE_PAIR(IS_LONG, IS_DOUBLE):
			result.SetDval(float64(op1.GetLval() - op2.GetDval()))
			ZVAL_LONG(result, ZEND_NORMALIZE_BOOL(result.GetDval()))
			return SUCCESS
		case TYPE_PAIR(IS_DOUBLE, IS_DOUBLE):
			if op1.GetDval() == op2.GetDval() {
				ZVAL_LONG(result, 0)
			} else {
				result.SetDval(op1.GetDval() - op2.GetDval())
				ZVAL_LONG(result, ZEND_NORMALIZE_BOOL(result.GetDval()))
			}
			return SUCCESS
		case TYPE_PAIR(IS_ARRAY, IS_ARRAY):
			ZVAL_LONG(result, ZendCompareArrays(op1, op2))
			return SUCCESS
		case TYPE_PAIR(IS_NULL, IS_NULL):

		case TYPE_PAIR(IS_NULL, IS_FALSE):

		case TYPE_PAIR(IS_FALSE, IS_NULL):

		case TYPE_PAIR(IS_FALSE, IS_FALSE):

		case TYPE_PAIR(IS_TRUE, IS_TRUE):
			ZVAL_LONG(result, 0)
			return SUCCESS
		case TYPE_PAIR(IS_NULL, IS_TRUE):
			ZVAL_LONG(result, -1)
			return SUCCESS
		case TYPE_PAIR(IS_TRUE, IS_NULL):
			ZVAL_LONG(result, 1)
			return SUCCESS
		case TYPE_PAIR(IS_STRING, IS_STRING):
			if op1.GetStr() == op2.GetStr() {
				ZVAL_LONG(result, 0)
				return SUCCESS
			}
			ZVAL_LONG(result, ZendiSmartStrcmp(op1.GetStr(), op2.GetStr()))
			return SUCCESS
		case TYPE_PAIR(IS_NULL, IS_STRING):
			ZVAL_LONG(result, b.Cond(Z_STRLEN_P(op2) == 0, 0, -1))
			return SUCCESS
		case TYPE_PAIR(IS_STRING, IS_NULL):
			ZVAL_LONG(result, b.Cond(Z_STRLEN_P(op1) == 0, 0, 1))
			return SUCCESS
		case TYPE_PAIR(IS_OBJECT, IS_NULL):
			ZVAL_LONG(result, 1)
			return SUCCESS
		case TYPE_PAIR(IS_NULL, IS_OBJECT):
			ZVAL_LONG(result, -1)
			return SUCCESS
		default:
			if Z_ISREF_P(op1) {
				op1 = Z_REFVAL_P(op1)
				continue
			} else if Z_ISREF_P(op2) {
				op2 = Z_REFVAL_P(op2)
				continue
			}
			if op1.IsType(IS_OBJECT) && Z_OBJ_HT(*op1).GetCompare() != nil {
				ret = Z_OBJ_HT(*op1).GetCompare()(result, op1, op2)
				if result.GetType() != IS_LONG {
					ConvertCompareResultToLong(result)
				}
				return ret
			} else if op2.IsType(IS_OBJECT) && Z_OBJ_HT(*op2).GetCompare() != nil {
				ret = Z_OBJ_HT(*op2).GetCompare()(result, op1, op2)
				if result.GetType() != IS_LONG {
					ConvertCompareResultToLong(result)
				}
				return ret
			}
			if op1.IsType(IS_OBJECT) && op2.IsType(IS_OBJECT) {
				if op1.GetObj() == op2.GetObj() {

					/* object handles are identical, apparently this is the same object */

					ZVAL_LONG(result, 0)
					return SUCCESS
				}
				if Z_OBJ_HT(*op1).GetCompareObjects() == Z_OBJ_HT(*op2).GetCompareObjects() {
					ZVAL_LONG(result, Z_OBJ_HT(*op1).GetCompareObjects()(op1, op2))
					return SUCCESS
				}
			}
			if op1.IsType(IS_OBJECT) {
				if Z_OBJ_HT_P(op1).GetGet() != nil {
					var rv Zval
					op_free = Z_OBJ_HT_P(op1).GetGet()(op1, &rv)
					ret = CompareFunction(result, op_free, op2)
					ZendFreeObjGetResult(op_free)
					return ret
				} else if op2.GetType() != IS_OBJECT && Z_OBJ_HT_P(op1).GetCastObject() != nil {
					ZVAL_UNDEF(&tmp_free)
					if Z_OBJ_HT_P(op1).GetCastObject()(op1, &tmp_free, b.CondF2(op2.IsType(IS_FALSE) || op2.IsType(IS_TRUE), _IS_BOOL, func() ZendUchar { return op2.GetType() })) == FAILURE {
						ZVAL_LONG(result, 1)
						ZendFreeObjGetResult(&tmp_free)
						return SUCCESS
					}
					ret = CompareFunction(result, &tmp_free, op2)
					ZendFreeObjGetResult(&tmp_free)
					return ret
				}
			}
			if op2.IsType(IS_OBJECT) {
				if Z_OBJ_HT_P(op2).GetGet() != nil {
					var rv Zval
					op_free = Z_OBJ_HT_P(op2).GetGet()(op2, &rv)
					ret = CompareFunction(result, op1, op_free)
					ZendFreeObjGetResult(op_free)
					return ret
				} else if op1.GetType() != IS_OBJECT && Z_OBJ_HT_P(op2).GetCastObject() != nil {
					ZVAL_UNDEF(&tmp_free)
					if Z_OBJ_HT_P(op2).GetCastObject()(op2, &tmp_free, b.CondF2(op1.IsType(IS_FALSE) || op1.IsType(IS_TRUE), _IS_BOOL, func() ZendUchar { return op1.GetType() })) == FAILURE {
						ZVAL_LONG(result, -1)
						ZendFreeObjGetResult(&tmp_free)
						return SUCCESS
					}
					ret = CompareFunction(result, op1, &tmp_free)
					ZendFreeObjGetResult(&tmp_free)
					return ret
				} else if op1.IsType(IS_OBJECT) {
					ZVAL_LONG(result, 1)
					return SUCCESS
				}
			}
			if converted == 0 {
				if op1.GetType() < IS_TRUE {
					ZVAL_LONG(result, b.Cond(ZvalIsTrue(op2) != 0, -1, 0))
					return SUCCESS
				} else if op1.IsType(IS_TRUE) {
					ZVAL_LONG(result, b.Cond(ZvalIsTrue(op2) != 0, 0, 1))
					return SUCCESS
				} else if op2.GetType() < IS_TRUE {
					ZVAL_LONG(result, b.Cond(ZvalIsTrue(op1) != 0, 1, 0))
					return SUCCESS
				} else if op2.IsType(IS_TRUE) {
					ZVAL_LONG(result, b.Cond(ZvalIsTrue(op1) != 0, 0, -1))
					return SUCCESS
				} else {
					op1 = ZendiConvertScalarToNumber(op1, &op1_copy, result, 1)
					op2 = ZendiConvertScalarToNumber(op2, &op2_copy, result, 1)
					if ExecutorGlobals.GetException() != nil {
						if result != op1 {
							ZVAL_UNDEF(result)
						}
						return FAILURE
					}
					converted = 1
				}
			} else if op1.IsType(IS_ARRAY) {
				ZVAL_LONG(result, 1)
				return SUCCESS
			} else if op2.IsType(IS_ARRAY) {
				ZVAL_LONG(result, -1)
				return SUCCESS
			} else {
				ZEND_ASSERT(false)
				ZendThrowError(nil, "Unsupported operand types")
				if result != op1 {
					ZVAL_UNDEF(result)
				}
				return FAILURE
			}
		}
	}
}
func HashZvalIdenticalFunction(z1 *Zval, z2 *Zval) int {
	/* is_identical_function() returns 1 in case of identity and 0 in case
	 * of a difference;
	 * whereas this comparison function is expected to return 0 on identity,
	 * and non zero otherwise.
	 */

	ZVAL_DEREF(z1)
	ZVAL_DEREF(z2)
	return FastIsNotIdenticalFunction(z1, z2)
}
func ZendIsIdentical(op1 *Zval, op2 *Zval) ZendBool {
	if op1.GetType() != op2.GetType() {
		return 0
	}
	switch op1.GetType() {
	case IS_NULL:

	case IS_FALSE:

	case IS_TRUE:
		return 1
	case IS_LONG:
		return op1.GetLval() == op2.GetLval()
	case IS_RESOURCE:
		return op1.GetRes() == op2.GetRes()
	case IS_DOUBLE:
		return op1.GetDval() == op2.GetDval()
	case IS_STRING:
		return ZendStringEquals(op1.GetStr(), op2.GetStr())
	case IS_ARRAY:
		return op1.GetArr() == op2.GetArr() || ZendHashCompare(op1.GetArr(), op2.GetArr(), CompareFuncT(HashZvalIdenticalFunction), 1) == 0
	case IS_OBJECT:
		return op1.GetObj() == op2.GetObj()
	default:
		return 0
	}
}
func IsIdenticalFunction(result *Zval, op1 *Zval, op2 *Zval) int {
	ZVAL_BOOL(result, ZendIsIdentical(op1, op2))
	return SUCCESS
}
func IsNotIdenticalFunction(result *Zval, op1 *Zval, op2 *Zval) int {
	ZVAL_BOOL(result, !(ZendIsIdentical(op1, op2)))
	return SUCCESS
}
func IsEqualFunction(result *Zval, op1 *Zval, op2 *Zval) int {
	if CompareFunction(result, op1, op2) == FAILURE {
		return FAILURE
	}
	ZVAL_BOOL(result, result.GetLval() == 0)
	return SUCCESS
}
func IsNotEqualFunction(result *Zval, op1 *Zval, op2 *Zval) int {
	if CompareFunction(result, op1, op2) == FAILURE {
		return FAILURE
	}
	ZVAL_BOOL(result, result.GetLval() != 0)
	return SUCCESS
}
func IsSmallerFunction(result *Zval, op1 *Zval, op2 *Zval) int {
	if CompareFunction(result, op1, op2) == FAILURE {
		return FAILURE
	}
	ZVAL_BOOL(result, result.GetLval() < 0)
	return SUCCESS
}
func IsSmallerOrEqualFunction(result *Zval, op1 *Zval, op2 *Zval) int {
	if CompareFunction(result, op1, op2) == FAILURE {
		return FAILURE
	}
	ZVAL_BOOL(result, result.GetLval() <= 0)
	return SUCCESS
}
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
func InstanceofInterface(instance_ce *ZendClassEntry, ce *ZendClassEntry) ZendBool {
	var i uint32
	if instance_ce.GetNumInterfaces() != 0 {
		ZEND_ASSERT(instance_ce.IsResolvedInterfaces())
		for i = 0; i < instance_ce.GetNumInterfaces(); i++ {
			if instance_ce.interfaces[i] == ce {
				return 1
			}
		}
	}
	return instance_ce == ce
}
func InstanceofFunctionEx(instance_ce *ZendClassEntry, ce *ZendClassEntry, is_interface ZendBool) ZendBool {
	if is_interface != 0 {
		ZEND_ASSERT(ce.IsInterface())
		return InstanceofInterface(instance_ce, ce)
	} else {
		ZEND_ASSERT(!ce.IsInterface())
		return InstanceofClass(instance_ce, ce)
	}
}
func InstanceofFunction(instance_ce *ZendClassEntry, ce *ZendClassEntry) ZendBool {
	if ce.IsInterface() {
		return InstanceofInterface(instance_ce, ce)
	} else {
		return InstanceofClass(instance_ce, ce)
	}
}
func IncrementString(str *Zval) {
	var carry int = 0
	var pos int = Z_STRLEN_P(str) - 1
	var s *byte
	var t *ZendString
	var last int = 0
	var ch int
	if Z_STRLEN_P(str) == 0 {
		ZvalPtrDtorStr(str)
		ZVAL_INTERNED_STR(str, ZSTR_CHAR('1'))
		return
	}
	if !(Z_REFCOUNTED_P(str)) {
		str.SetStr(ZendStringInit(Z_STRVAL_P(str), Z_STRLEN_P(str), 0))
		str.SetTypeInfo(IS_STRING_EX)
	} else if Z_REFCOUNT_P(str) > 1 {
		Z_DELREF_P(str)
		str.SetStr(ZendStringInit(Z_STRVAL_P(str), Z_STRLEN_P(str), 0))
	} else {
		ZendStringForgetHashVal(str.GetStr())
	}
	s = Z_STRVAL_P(str)
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
			last = LOWER_CASE
		} else if ch >= 'A' && ch <= 'Z' {
			if ch == 'Z' {
				s[pos] = 'A'
				carry = 1
			} else {
				s[pos]++
				carry = 0
			}
			last = UPPER_CASE
		} else if ch >= '0' && ch <= '9' {
			if ch == '9' {
				s[pos] = '0'
				carry = 1
			} else {
				s[pos]++
				carry = 0
			}
			last = NUMERIC
		} else {
			carry = 0
			break
		}
		if carry == 0 {
			break
		}
		if b.PostDec(&pos) <= 0 {
			break
		}
	}
	if carry != 0 {
		t = ZendStringAlloc(Z_STRLEN_P(str)+1, 0)
		memcpy(t.GetVal()+1, Z_STRVAL_P(str), Z_STRLEN_P(str))
		t.GetVal()[Z_STRLEN_P(str)+1] = '0'
		switch last {
		case NUMERIC:
			t.GetVal()[0] = '1'
			break
		case UPPER_CASE:
			t.GetVal()[0] = 'A'
			break
		case LOWER_CASE:
			t.GetVal()[0] = 'a'
			break
		}
		ZendStringFree(str.GetStr())
		ZVAL_NEW_STR(str, t)
	}
}
func IncrementFunction(op1 *Zval) int {
try_again:
	switch op1.GetType() {
	case IS_LONG:
		FastLongIncrementFunction(op1)
		break
	case IS_DOUBLE:
		op1.SetDval(op1.GetDval() + 1)
		break
	case IS_NULL:
		ZVAL_LONG(op1, 1)
		break
	case IS_STRING:
		var lval ZendLong
		var dval float64
		switch IsNumericString(Z_STRVAL_P(op1), Z_STRLEN_P(op1), &lval, &dval, 0) {
		case IS_LONG:
			ZvalPtrDtorStr(op1)
			if lval == ZEND_LONG_MAX {

				/* switch to double */

				var d float64 = float64(lval)
				ZVAL_DOUBLE(op1, d+1)
			} else {
				ZVAL_LONG(op1, lval+1)
			}
			break
		case IS_DOUBLE:
			ZvalPtrDtorStr(op1)
			ZVAL_DOUBLE(op1, dval+1)
			break
		default:

			/* Perl style string increment */

			IncrementString(op1)
			break
		}
		break
	case IS_OBJECT:
		if Z_OBJ_HT(*op1).GetGet() != nil && Z_OBJ_HT(*op1).GetSet() != nil {

			/* proxy object */

			var rv Zval
			var val *Zval
			val = Z_OBJ_HT(*op1).GetGet()(op1, &rv)
			Z_TRY_ADDREF_P(val)
			IncrementFunction(val)
			Z_OBJ_HT(*op1).GetSet()(op1, val)
			ZvalPtrDtor(val)
		} else if Z_OBJ_HT(*op1).GetDoOperation() != nil {
			var op2 Zval
			var res int
			ZVAL_LONG(&op2, 1)
			res = Z_OBJ_HT(*op1).GetDoOperation()(ZEND_ADD, op1, op1, &op2)
			return res
		}
		return FAILURE
	case IS_REFERENCE:
		op1 = Z_REFVAL_P(op1)
		goto try_again
	default:
		return FAILURE
	}
	return SUCCESS
}
func DecrementFunction(op1 *Zval) int {
	var lval ZendLong
	var dval float64
try_again:
	switch op1.GetType() {
	case IS_LONG:
		FastLongDecrementFunction(op1)
		break
	case IS_DOUBLE:
		op1.SetDval(op1.GetDval() - 1)
		break
	case IS_STRING:
		if Z_STRLEN_P(op1) == 0 {
			ZvalPtrDtorStr(op1)
			ZVAL_LONG(op1, -1)
			break
		}
		switch IsNumericString(Z_STRVAL_P(op1), Z_STRLEN_P(op1), &lval, &dval, 0) {
		case IS_LONG:
			ZvalPtrDtorStr(op1)
			if lval == ZEND_LONG_MIN {
				var d float64 = float64(lval)
				ZVAL_DOUBLE(op1, d-1)
			} else {
				ZVAL_LONG(op1, lval-1)
			}
			break
		case IS_DOUBLE:
			ZvalPtrDtorStr(op1)
			ZVAL_DOUBLE(op1, dval-1)
			break
		}
		break
	case IS_OBJECT:
		if Z_OBJ_HT(*op1).GetGet() != nil && Z_OBJ_HT(*op1).GetSet() != nil {

			/* proxy object */

			var rv Zval
			var val *Zval
			val = Z_OBJ_HT(*op1).GetGet()(op1, &rv)
			Z_TRY_ADDREF_P(val)
			DecrementFunction(val)
			Z_OBJ_HT(*op1).GetSet()(op1, val)
			ZvalPtrDtor(val)
		} else if Z_OBJ_HT(*op1).GetDoOperation() != nil {
			var op2 Zval
			var res int
			ZVAL_LONG(&op2, 1)
			res = Z_OBJ_HT(*op1).GetDoOperation()(ZEND_SUB, op1, op1, &op2)
			return res
		}
		return FAILURE
	case IS_REFERENCE:
		op1 = Z_REFVAL_P(op1)
		goto try_again
	default:
		return FAILURE
	}
	return SUCCESS
}
func ZendIsTrue(op *Zval) int { return IZendIsTrue(op) }
func ZendObjectIsTrue(op *Zval) int {
	if Z_OBJ_HT_P(op).GetCastObject() != nil {
		var tmp Zval
		if Z_OBJ_HT_P(op).GetCastObject()(op, &tmp, _IS_BOOL) == SUCCESS {
			return tmp.IsType(IS_TRUE)
		}
		ZendError(E_RECOVERABLE_ERROR, "Object of class %s could not be converted to bool", Z_OBJ_P(op).GetCe().GetName().GetVal())
	} else if Z_OBJ_HT_P(op).GetGet() != nil {
		var result int
		var rv Zval
		var tmp *Zval = Z_OBJ_HT_P(op).GetGet()(op, &rv)
		if tmp.GetType() != IS_OBJECT {

			/* for safety - avoid loop */

			result = IZendIsTrue(tmp)
			ZvalPtrDtor(tmp)
			return result
		}
	}
	return 1
}
func ZendStrTolowerCopy(dest *byte, source *byte, length int) *byte {
	var str *uint8 = (*uint8)(source)
	var result *uint8 = (*uint8)(dest)
	var end *uint8 = str + length
	for str < end {
		b.PostInc(&(*result)) = ZendTolowerAscii(b.PostInc(&(*str)))
	}
	*result = '0'
	return dest
}
func ZendStrTolowerDup(source *byte, length int) *byte {
	return ZendStrTolowerCopy((*byte)(Emalloc(length+1)), source, length)
}
func ZendStrTolower(str *byte, length int) {
	var p *uint8 = (*uint8)(str)
	var end *uint8 = p + length
	for p < end {
		*p = ZendTolowerAscii(*p)
		p++
	}
}
func ZendStrTolowerDupEx(source *byte, length int) *byte {
	var p *uint8 = (*uint8)(source)
	var end *uint8 = p + length
	for p < end {
		if (*p) != ZendTolowerAscii(*p) {
			var res *byte = (*byte)(Emalloc(length + 1))
			var r *uint8
			if p != (*uint8)(source) {
				memcpy(res, source, p-(*uint8)(source))
			}
			r = (*uint8)(p + (res - source))
			for p < end {
				*r = ZendTolowerAscii(*p)
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
func ZendStringTolowerEx(str *ZendString, persistent int) *ZendString {
	var p *uint8 = (*uint8)(str.GetVal())
	var end *uint8 = p + str.GetLen()
	for p < end {
		if (*p) != ZendTolowerAscii(*p) {
			var res *ZendString = ZendStringAlloc(str.GetLen(), persistent)
			var r *uint8
			if p != (*uint8)(str.GetVal()) {
				memcpy(res.GetVal(), str.GetVal(), p-(*uint8)(str.GetVal()))
			}
			r = p + (res.GetVal() - str.GetVal())
			for p < end {
				*r = ZendTolowerAscii(*p)
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
func ZendBinaryStrcmp(s1 *byte, len1 int, s2 *byte, len2 int) int {
	var retval int
	if s1 == s2 {
		return 0
	}
	retval = memcmp(s1, s2, MIN(len1, len2))
	if retval == 0 {
		return int(len1 - len2)
	} else {
		return retval
	}
}
func ZendBinaryStrncmp(s1 *byte, len1 int, s2 *byte, len2 int, length int) int {
	var retval int
	if s1 == s2 {
		return 0
	}
	retval = memcmp(s1, s2, MIN(length, MIN(len1, len2)))
	if retval == 0 {
		return int(MIN(length, len1) - MIN(length, len2))
	} else {
		return retval
	}
}
func ZendBinaryStrcasecmp(s1 *byte, len1 int, s2 *byte, len2 int) int {
	var len_ int
	var c1 int
	var c2 int
	if s1 == s2 {
		return 0
	}
	len_ = MIN(len1, len2)
	for b.PostDec(&len_) {
		c1 = ZendTolowerAscii(*((*uint8)(b.PostInc(&s1))))
		c2 = ZendTolowerAscii(*((*uint8)(b.PostInc(&s2))))
		if c1 != c2 {
			return c1 - c2
		}
	}
	return int(len1 - len2)
}
func ZendBinaryStrncasecmp(s1 *byte, len1 int, s2 string, len2 int, length int) int {
	var len_ int
	var c1 int
	var c2 int
	if s1 == s2 {
		return 0
	}
	len_ = MIN(length, MIN(len1, len2))
	for b.PostDec(&len_) {
		c1 = ZendTolowerAscii(*((*uint8)(b.PostInc(&s1))))
		c2 = ZendTolowerAscii(*((*uint8)(b.PostInc(&s2))))
		if c1 != c2 {
			return c1 - c2
		}
	}
	return int(MIN(length, len1) - MIN(length, len2))
}
func ZendBinaryStrcasecmpL(s1 *byte, len1 int, s2 *byte, len2 int) int {
	var len_ int
	var c1 int
	var c2 int
	if s1 == s2 {
		return 0
	}
	len_ = MIN(len1, len2)
	for b.PostDec(&len_) {
		c1 = ZendTolower(int(*((*uint8)(b.PostInc(&s1)))))
		c2 = ZendTolower(int(*((*uint8)(b.PostInc(&s2)))))
		if c1 != c2 {
			return c1 - c2
		}
	}
	return int(len1 - len2)
}
func ZendBinaryStrncasecmpL(s1 *byte, len1 int, s2 *byte, len2 int, length int) int {
	var len_ int
	var c1 int
	var c2 int
	if s1 == s2 {
		return 0
	}
	len_ = MIN(length, MIN(len1, len2))
	for b.PostDec(&len_) {
		c1 = ZendTolower(int(*((*uint8)(b.PostInc(&s1)))))
		c2 = ZendTolower(int(*((*uint8)(b.PostInc(&s2)))))
		if c1 != c2 {
			return c1 - c2
		}
	}
	return int(MIN(length, len1) - MIN(length, len2))
}
func ZendBinaryZvalStrcmp(s1 *Zval, s2 *Zval) int {
	return ZendBinaryStrcmp(Z_STRVAL_P(s1), Z_STRLEN_P(s1), Z_STRVAL_P(s2), Z_STRLEN_P(s2))
}
func ZendBinaryZvalStrncmp(s1 *Zval, s2 *Zval, s3 *Zval) int {
	return ZendBinaryStrncmp(Z_STRVAL_P(s1), Z_STRLEN_P(s1), Z_STRVAL_P(s2), Z_STRLEN_P(s2), s3.GetLval())
}
func ZendBinaryZvalStrcasecmp(s1 *Zval, s2 *Zval) int {
	return ZendBinaryStrcasecmpL(Z_STRVAL_P(s1), Z_STRLEN_P(s1), Z_STRVAL_P(s2), Z_STRLEN_P(s2))
}
func ZendBinaryZvalStrncasecmp(s1 *Zval, s2 *Zval, s3 *Zval) int {
	return ZendBinaryStrncasecmpL(Z_STRVAL_P(s1), Z_STRLEN_P(s1), Z_STRVAL_P(s2), Z_STRLEN_P(s2), s3.GetLval())
}
func ZendiSmartStreq(s1 *ZendString, s2 *ZendString) int {
	var ret1 int
	var ret2 int
	var oflow1 int
	var oflow2 int
	var lval1 ZendLong = 0
	var lval2 ZendLong = 0
	var dval1 float64 = 0.0
	var dval2 float64 = 0.0
	if b.Assign(&ret1, IsNumericStringEx(s1.GetVal(), s1.GetLen(), &lval1, &dval1, 0, &oflow1)) && b.Assign(&ret2, IsNumericStringEx(s2.GetVal(), s2.GetLen(), &lval2, &dval2, 0, &oflow2)) {
		if oflow1 != 0 && oflow1 == oflow2 && dval1-dval2 == 0.0 {

			/* both values are integers overflown to the same side, and the
			 * double comparison may have resulted in crucial accuracy lost */

			goto string_cmp

			/* both values are integers overflown to the same side, and the
			 * double comparison may have resulted in crucial accuracy lost */

		}
		if ret1 == IS_DOUBLE || ret2 == IS_DOUBLE {
			if ret1 != IS_DOUBLE {
				if oflow2 != 0 {

					/* 2nd operand is integer > LONG_MAX (oflow2==1) or < LONG_MIN (-1) */

					return 0

					/* 2nd operand is integer > LONG_MAX (oflow2==1) or < LONG_MIN (-1) */

				}
				dval1 = float64(lval1)
			} else if ret2 != IS_DOUBLE {
				if oflow1 != 0 {
					return 0
				}
				dval2 = float64(lval2)
			} else if dval1 == dval2 && !(core.ZendFinite(dval1)) {

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
func ZendiSmartStrcmp(s1 *ZendString, s2 *ZendString) int {
	var ret1 int
	var ret2 int
	var oflow1 int
	var oflow2 int
	var lval1 ZendLong = 0
	var lval2 ZendLong = 0
	var dval1 float64 = 0.0
	var dval2 float64 = 0.0
	if b.Assign(&ret1, IsNumericStringEx(s1.GetVal(), s1.GetLen(), &lval1, &dval1, 0, &oflow1)) && b.Assign(&ret2, IsNumericStringEx(s2.GetVal(), s2.GetLen(), &lval2, &dval2, 0, &oflow2)) {
		if oflow1 != 0 && oflow1 == oflow2 && dval1-dval2 == 0.0 {

			/* both values are integers overflown to the same side, and the
			 * double comparison may have resulted in crucial accuracy lost */

			goto string_cmp

			/* both values are integers overflown to the same side, and the
			 * double comparison may have resulted in crucial accuracy lost */

		}
		if ret1 == IS_DOUBLE || ret2 == IS_DOUBLE {
			if ret1 != IS_DOUBLE {
				if oflow2 != 0 {

					/* 2nd operand is integer > LONG_MAX (oflow2==1) or < LONG_MIN (-1) */

					return -1 * oflow2

					/* 2nd operand is integer > LONG_MAX (oflow2==1) or < LONG_MIN (-1) */

				}
				dval1 = float64(lval1)
			} else if ret2 != IS_DOUBLE {
				if oflow1 != 0 {
					return oflow1
				}
				dval2 = float64(lval2)
			} else if dval1 == dval2 && !(core.ZendFinite(dval1)) {

				/* Both values overflowed and have the same sign,
				 * so a numeric comparison would be inaccurate */

				goto string_cmp

				/* Both values overflowed and have the same sign,
				 * so a numeric comparison would be inaccurate */

			}
			dval1 = dval1 - dval2
			return ZEND_NORMALIZE_BOOL(dval1)
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
		return ZEND_NORMALIZE_BOOL(strcmp_ret)
	}
}
func HashZvalCompareFunction(z1 *Zval, z2 *Zval) int {
	var result Zval
	if CompareFunction(&result, z1, z2) == FAILURE {
		return 1
	}
	return result.GetLval()
}
func ZendCompareSymbolTables(ht1 *HashTable, ht2 *HashTable) int {
	if ht1 == ht2 {
		return 0
	} else {
		return ZendHashCompare(ht1, ht2, CompareFuncT(HashZvalCompareFunction), 0)
	}
}
func ZendCompareArrays(a1 *Zval, a2 *Zval) int {
	return ZendCompareSymbolTables(a1.GetArr(), a2.GetArr())
}
func ZendCompareObjects(o1 *Zval, o2 *Zval) int {
	if o1.GetObj() == o2.GetObj() {
		return 0
	}
	if Z_OBJ_HT_P(o1).GetCompareObjects() == nil {
		return 1
	} else {
		return Z_OBJ_HT_P(o1).GetCompareObjects()(o1, o2)
	}
}
func ZendLocaleSprintfDouble(op *Zval) {
	var str *ZendString
	str = ZendStrpprintf(0, "%.*G", int(ExecutorGlobals.GetPrecision()), float64(op.GetDval()))
	ZVAL_NEW_STR(op, str)
}
func ZendLongToStr(num ZendLong) *ZendString {
	if ZendUlong(num <= 9) != 0 {
		return ZSTR_CHAR(ZendUchar('0' + ZendUchar(num)))
	} else {
		var buf []byte
		var res *byte = ZendPrintLongToBuf(buf+b.SizeOf("buf")-1, num)
		return ZendStringInit(res, buf+b.SizeOf("buf")-1-res, 0)
	}
}
func IsNumericStrFunction(str *ZendString, lval *ZendLong, dval *float64) ZendUchar {
	return IsNumericStringEx(str.GetVal(), str.GetLen(), lval, dval, -1, nil)
}
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
	if ZEND_IS_DIGIT(*ptr) {

		/* Skip any leading 0s */

		for (*ptr) == '0' {
			ptr++
		}

		/* Count the number of digits. If a decimal point/exponent is found,
		 * it's a double. Otherwise, if there's a dval or no need to check for
		 * a full match, stop when there are too many digits for a long */

		for type_ = IS_LONG; !(digits >= MAX_LENGTH_OF_LONG && (dval != nil || allow_errors == 1)); {
		check_digits:
			if ZEND_IS_DIGIT(*ptr) {
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
				if ZEND_IS_DIGIT(*e) {
					goto process_double
				}
			}
			break
			digits++
			ptr++
		}
		if digits >= MAX_LENGTH_OF_LONG {
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
	} else if (*ptr) == '.' && ZEND_IS_DIGIT(ptr[1]) {
	process_double:
		type_ = IS_DOUBLE

		/* If there's a dval, do the conversion; else continue checking
		 * the digits if we need to check for a full match */

		if dval != nil {
			local_dval = ZendStrtod(str, &ptr)
		} else if allow_errors != 1 && dp_or_e != -1 {
			if b.PostInc(&(*ptr)) == '.' {
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
			ZendError(E_NOTICE, "A non well formed numeric value encountered")
			if ExecutorGlobals.GetException() != nil {
				return 0
			}
		}
	}
	if type_ == IS_LONG {
		if digits == MAX_LENGTH_OF_LONG-1 {
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
				return IS_DOUBLE
			}
		}
		if lval != nil {
			if neg != 0 {
				tmp_lval = -tmp_lval
			}
			*lval = ZendLong(tmp_lval)
		}
		return IS_LONG
	} else {
		if dval != nil {
			*dval = local_dval
		}
		return IS_DOUBLE
	}
}
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
