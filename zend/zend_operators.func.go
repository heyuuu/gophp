// <<generate>>

package zend

import (
	"log"
	b "sik/builtin"
	"sik/core"
	. "sik/runtime/ctype"
	"strconv"
	"strings"
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
func IsNumericStringEx(
	str *byte,
	length int,
	lval *ZendLong,
	dval *float64,
	allow_errors int,
	oflow_info *int,
) ZendUchar {
	typ, retLval, retDval, overflowInfo := ParseNumericStrEx(b.CastStr(str, length), allow_errors)

	*lval = retLval
	if dval != nil {
		*dval = retDval
	}
	if oflow_info != nil {
		*oflow_info = overflowInfo
	}
	return typ
}
func IsNumericString(str *byte, length int, lval *ZendLong, dval *float64, allow_errors int) ZendUchar {
	typ, retLval, retDval, _ := ParseNumericStrEx(b.CastStr(str, length), allow_errors)

	*lval = retLval
	if dval != nil {
		*dval = retDval
	}
	return typ
}
func ZendMemnstr(haystack *byte, needle string, needle_len int, end *byte) *byte {
	// todo 替换
	pos := strings.Index(b.CastStr(haystack, end-haystack), b.CastStr(needle))
	if pos < 0 {
		return nil
	}
	return haystack + pos
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
	if op.IsLong() {
		return op.GetLval()
	} else {
		return ZvalGetLongFunc(op)
	}
}
func ZvalGetDouble(op *Zval) float64 {
	if op.IsDouble() {
		return op.GetDval()
	} else {
		return ZvalGetDoubleFunc(op)
	}
}
func ZvalGetString(op *Zval) *ZendString {
	if op.IsString() {
		return op.GetStr().Copy()
	} else {
		return ZvalGetStringFunc(op)
	}
}
func ZvalGetTmpString(op *Zval, tmp **ZendString) *ZendString {
	if op.IsString() {
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
	if op.IsString() {
		var ret *ZendString = op.GetStr().Copy()
		return ret
	} else {
		return ZvalTryGetStringFunc(op)
	}
}
func ZvalTryGetTmpString(op *Zval, tmp **ZendString) *ZendString {
	if op.IsString() {
		var ret *ZendString = op.GetStr()
		*tmp = nil
		return ret
	} else {
		*tmp = ZvalTryGetStringFunc(op)
		return *tmp
	}
}
func TryConvertToString(op *Zval) ZendBool {
	if op.IsString() {
		return 1
	}
	return _tryConvertToString(op)
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
	case IS_LONG:
		if op.GetLval() != 0 {
			result = 1
		}
	case IS_DOUBLE:
		if op.GetDval() {
			result = 1
		}
	case IS_STRING:
		if Z_STRLEN_P(op) > 1 || Z_STRLEN_P(op) != 0 && Z_STRVAL_P(op)[0] != '0' {
			result = 1
		}
	case IS_ARRAY:
		if Z_ARRVAL_P(op).GetNNumOfElements() {
			result = 1
		}
	case IS_OBJECT:
		if Z_OBJ_HT_P(op).GetCastObject() == ZendStdCastObjectTostring {
			result = 1
		} else {
			result = ZendObjectIsTrue(op)
		}
	case IS_RESOURCE:
		if Z_RES_HANDLE_P(op) != 0 {
			result = 1
		}
	case IS_REFERENCE:
		op = Z_REFVAL_P(op)
		goto again
	default:

	}
	return result
}
func ZendStringTolower(str *ZendString) *ZendString { return ZendStringTolowerEx(str, 0) }
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
func ConvertScalarToNumberEx(pzv *Zval) {
	if pzv.GetType() != IS_LONG && pzv.GetType() != IS_DOUBLE {
		ConvertScalarToNumber(pzv)
	}
}
func FastLongIncrementFunction(op1 *Zval) {
	if op1.GetLval() == ZEND_LONG_MAX {

		/* switch to double */

		op1.SetDouble(float64(ZEND_LONG_MAX + 1.0))

		/* switch to double */

	} else {
		op1.GetLval()++
	}
}
func FastLongDecrementFunction(op1 *Zval) {
	if op1.GetLval() == ZEND_LONG_MIN {

		/* switch to double */

		op1.SetDouble(float64(ZEND_LONG_MIN - 1.0))

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
		result.SetDouble(float64(op1.GetLval() + float64(op2.GetLval())))
	} else {
		result.SetLong(op1.GetLval() + op2.GetLval())
	}

	/*
	 * 'result' may alias with op1 or op2, so we need to
	 * ensure that 'result' is not updated until after we
	 * have read the values of op1 and op2.
	 */
}
func FastAddFunction(result *Zval, op1 *Zval, op2 *Zval) int {
	if op1.IsLong() {
		if op2.IsLong() {
			FastLongAddFunction(result, op1, op2)
			return SUCCESS
		} else if op2.IsDouble() {
			result.SetDouble(float64(op1.GetLval()) + op2.GetDval())
			return SUCCESS
		}
	} else if op1.IsDouble() {
		if op2.IsDouble() {
			result.SetDouble(op1.GetDval() + op2.GetDval())
			return SUCCESS
		} else if op2.IsLong() {
			result.SetDouble(op1.GetDval() + float64(op2.GetLval()))
			return SUCCESS
		}
	}
	return AddFunction(result, op1, op2)
}
func FastLongSubFunction(result *Zval, op1 *Zval, op2 *Zval) {
	result.SetLong(op1.GetLval() - op2.GetLval())
	if (op1.GetLval()&LONG_SIGN_MASK) != (op2.GetLval()&LONG_SIGN_MASK) && (op1.GetLval()&LONG_SIGN_MASK) != (result.GetLval()&LONG_SIGN_MASK) {
		result.SetDouble(float64(op1.GetLval() - float64(op2.GetLval())))
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
	if op1.IsLong() {
		if op2.IsLong() {
			return op1.GetLval() == op2.GetLval()
		} else if op2.IsDouble() {
			return float64(op1.GetLval()) == op2.GetDval()
		}
	} else if op1.IsDouble() {
		if op2.IsDouble() {
			return op1.GetDval() == op2.GetDval()
		} else if op2.IsLong() {
			return op1.GetDval() == float64(op2.GetLval())
		}
	} else if op1.IsString() {
		if op2.IsString() {
			return ZendFastEqualStrings(op1.GetStr(), op2.GetStr())
		}
	}
	CompareFunction(&result, op1, op2)
	return result.GetLval() == 0
}
func FastEqualCheckLong(op1 *Zval, op2 *Zval) int {
	var result Zval
	if op2.IsLong() {
		return op1.GetLval() == op2.GetLval()
	}
	CompareFunction(&result, op1, op2)
	return result.GetLval() == 0
}
func FastEqualCheckString(op1 *Zval, op2 *Zval) int {
	var result Zval
	if op2.IsString() {
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
	if op.GetRefcount() == 1 {
		ZVAL_UNREF(op)
	} else {
		op.DelRefcount()
		ZVAL_COPY(op, Z_REFVAL_P(op))
	}
}
func ZendTolower(c int) __auto__         { return tolower(c) }
func TYPE_PAIR(t1 uint32, t2 uint32) int { return t1<<4 | t2 }
func ZendTolowerAscii(c uint8) uint8     { return TolowerMap[uint8(c)] }
func ZendAtolEx(str string) ZendLong {
	if len(str) == 0 {
		return 0
	}
	retval := ZEND_STRTOL_EX(str, 0)
	switch str[len(str)-1] {
	case 'g', 'G':
		retval *= 1024
		fallthrough
	case 'm', 'M':
		retval *= 1024
		fallthrough
	case 'k', 'K':
		retval *= 1024
	}

	return int(retval)
}
func ZendAtol(str *byte, str_len int) ZendLong {
	var retval ZendLong
	if str_len == 0 {
		str_len = strlen(str)
	}
	retval = ZEND_STRTOL(str, nil, 0)
	if str_len > 0 {
		switch str[str_len-1] {
		case 'g', 'G':
			retval *= 1024
			fallthrough
		case 'm', 'M':
			retval *= 1024
			fallthrough
		case 'k', 'K':
			retval *= 1024
		}
	}
	return retval
}
func ConvertObjectToType(op *Zval, dst *Zval, ctype int, conv_func func(op *Zval)) {
	dst.SetUndef()
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
			op.SetLong(0)
			if silent == 0 {
				ZendError(E_WARNING, "A non-numeric value encountered")
			}
		}
		ZendStringReleaseEx(str, 0)
	case IS_NULL:
		fallthrough
	case IS_FALSE:
		op.SetLong(0)
	case IS_TRUE:
		op.SetLong(1)
	case IS_RESOURCE:
		var l ZendLong = Z_RES_HANDLE_P(op)
		ZvalPtrDtor(op)
		op.SetLong(l)
	case IS_OBJECT:
		var dst Zval
		ConvertObjectToType(op, &dst, _IS_NUMBER, ConvertScalarToNumber)
		if check != 0 && EG__().GetException() != nil {
			return
		}
		ZvalPtrDtor(op)
		if dst.IsLong() || dst.IsDouble() {
			ZVAL_COPY_VALUE(op, &dst)
		} else {
			op.SetLong(1)
		}
	}
}
func ConvertScalarToNumber(op *Zval) { _convertScalarToNumber(op, 1, 0) }
func _zendiConvertScalarToNumberEx(op *Zval, holder *Zval, silent ZendBool) *Zval {
	switch op.GetType() {
	case IS_NULL:
		fallthrough
	case IS_FALSE:
		holder.SetLong(0)
		return holder
	case IS_TRUE:
		holder.SetLong(1)
		return holder
	case IS_STRING:
		if b.Assign(&(holder.GetTypeInfo()), IsNumericString(Z_STRVAL_P(op), Z_STRLEN_P(op), &(holder.GetLval()), &(holder.GetDval()), b.Cond(silent != 0, 1, -1))) == 0 {
			holder.SetLong(0)
			if silent == 0 {
				ZendError(E_WARNING, "A non-numeric value encountered")
			}
		}
		return holder
	case IS_RESOURCE:
		holder.SetLong(Z_RES_HANDLE_P(op))
		return holder
	case IS_OBJECT:
		ConvertObjectToType(op, holder, _IS_NUMBER, ConvertScalarToNumber)
		if EG__().GetException() != nil || holder.GetType() != IS_LONG && holder.GetType() != IS_DOUBLE {
			holder.SetLong(1)
		}
		return holder
	case IS_LONG:
		fallthrough
	case IS_DOUBLE:
		fallthrough
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
	if op.IsLong() || op.IsDouble() {
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
		fallthrough
	case IS_FALSE:
		op.SetLong(0)
	case IS_TRUE:
		op.SetLong(1)
	case IS_RESOURCE:
		tmp = Z_RES_HANDLE_P(op)
		ZvalPtrDtor(op)
		op.SetLong(tmp)
	case IS_LONG:

	case IS_DOUBLE:
		op.SetLong(ZendDvalToLval(op.GetDval()))
	case IS_STRING:
		var str *ZendString = Z_STR_P(op)
		if base == 10 {
			op.SetLong(ZvalGetLong(op))
		} else {
			op.SetLong(ZEND_STRTOL(str.GetVal(), nil, base))
		}
		ZendStringReleaseEx(str, 0)
	case IS_ARRAY:
		if Z_ARRVAL_P(op).GetNNumOfElements() {
			tmp = 1
		} else {
			tmp = 0
		}
		ZvalPtrDtor(op)
		op.SetLong(tmp)
	case IS_OBJECT:
		var dst Zval
		ConvertObjectToType(op, &dst, IS_LONG, ConvertToLong)
		ZvalPtrDtor(op)
		if dst.IsLong() {
			op.SetLong(dst.GetLval())
		} else {
			op.SetLong(1)
		}
		return
	case IS_REFERENCE:
		ZendUnwrapReference(op)
		goto try_again
	default:

	}
}
func ConvertToDouble(op *Zval) {
	var tmp float64
try_again:
	switch op.GetType() {
	case IS_NULL:
		fallthrough
	case IS_FALSE:
		op.SetDouble(0.0)
	case IS_TRUE:
		op.SetDouble(1.0)
	case IS_RESOURCE:
		var d float64 = float64(Z_RES_HANDLE_P(op))
		ZvalPtrDtor(op)
		op.SetDouble(d)
	case IS_LONG:
		op.SetDouble(float64(op.GetLval()))
	case IS_DOUBLE:

	case IS_STRING:
		var str *ZendString = Z_STR_P(op)
		op.SetDouble(ZendStrtod(str.GetVal(), nil))
		ZendStringReleaseEx(str, 0)
	case IS_ARRAY:
		if Z_ARRVAL_P(op).GetNNumOfElements() {
			tmp = 1
		} else {
			tmp = 0
		}
		ZvalPtrDtor(op)
		op.SetDouble(tmp)
	case IS_OBJECT:
		var dst Zval
		ConvertObjectToType(op, &dst, IS_DOUBLE, ConvertToDouble)
		ZvalPtrDtor(op)
		if dst.IsDouble() {
			op.SetDouble(dst.GetDval())
		} else {
			op.SetDouble(1.0)
		}
	case IS_REFERENCE:
		ZendUnwrapReference(op)
		goto try_again
	default:

	}
}
func ConvertToNull(op *Zval) {
	ZvalPtrDtor(op)
	op.SetNull()
}
func ConvertToBoolean(op *Zval) {
	var tmp int
try_again:
	switch op.GetType() {
	case IS_FALSE:
		fallthrough
	case IS_TRUE:

	case IS_NULL:
		op.SetFalse()
	case IS_RESOURCE:
		var l ZendLong = b.Cond(Z_RES_HANDLE_P(op) != 0, 1, 0)
		ZvalPtrDtor(op)
		ZVAL_BOOL(op, l != 0)
	case IS_LONG:
		ZVAL_BOOL(op, b.Cond(op.GetLval() != 0, 1, 0))
	case IS_DOUBLE:
		ZVAL_BOOL(op, b.Cond(op.GetDval(), 1, 0))
	case IS_STRING:
		var str *ZendString = Z_STR_P(op)
		if str.GetLen() == 0 || str.GetLen() == 1 && str.GetVal()[0] == '0' {
			op.SetFalse()
		} else {
			op.SetTrue()
		}
		ZendStringReleaseEx(str, 0)
	case IS_ARRAY:
		if Z_ARRVAL_P(op).GetNNumOfElements() {
			tmp = 1
		} else {
			tmp = 0
		}
		ZvalPtrDtor(op)
		ZVAL_BOOL(op, tmp != 0)
	case IS_OBJECT:
		var dst Zval
		ConvertObjectToType(op, &dst, _IS_BOOL, ConvertToBoolean)
		ZvalPtrDtor(op)
		if dst.GetTypeInfo() == IS_FALSE || dst.GetTypeInfo() == IS_TRUE {
			op.SetTypeInfo(dst.GetTypeInfo())
		} else {
			op.SetTrue()
		}
	case IS_REFERENCE:
		ZendUnwrapReference(op)
		goto try_again
	default:

	}
}
func _convertToString(op *Zval) {
try_again:
	switch op.GetType() {
	case IS_UNDEF:
		fallthrough
	case IS_NULL:
		fallthrough
	case IS_FALSE:
		ZVAL_EMPTY_STRING(op)
	case IS_TRUE:
		op.SetInternedString(ZSTR_CHAR('1'))
	case IS_STRING:

	case IS_RESOURCE:
		var str *ZendString = ZendStrpprintf(0, "Resource id #"+ZEND_LONG_FMT, ZendLong(Z_RES_HANDLE_P(op)))
		ZvalPtrDtor(op)
		op.SetString(str)
	case IS_LONG:
		op.SetString(ZendLongToStr(op.GetLval()))
	case IS_DOUBLE:
		var str *ZendString
		var dval float64 = Z_DVAL_P(op)
		str = ZendStrpprintf(0, "%.*G", int(EG__().GetPrecision()), dval)

		/* %G already handles removing trailing zeros from the fractional part, yay */

		op.SetString(str)
	case IS_ARRAY:
		ZendError(E_NOTICE, "Array to string conversion")
		ZvalPtrDtor(op)
		op.SetInternedString(ZSTR_KNOWN(ZEND_STR_ARRAY_CAPITALIZED))
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
				op.SetString(str)
				return
			}
			ZvalPtrDtor(z)
		}
		if EG__().GetException() == nil {
			ZendThrowError(nil, "Object of class %s could not be converted to string", Z_OBJCE_P(op).GetName().GetVal())
		}
		ZvalPtrDtor(op)
		ZVAL_EMPTY_STRING(op)
	case IS_REFERENCE:
		ZendUnwrapReference(op)
		goto try_again
	default:

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
	op.SetString(str)
	return 1
}
func ConvertScalarToArray(op *Zval) {
	var ht *HashTable = ZendNewArray(1)
	ht.IndexAddNewH(0, op)
	op.SetArray(ht)
}
func ConvertToArray(op *Zval) {
try_again:
	switch op.GetType() {
	case IS_ARRAY:

	case IS_OBJECT:
		if Z_OBJCE_P(op) == ZendCeClosure {
			ConvertScalarToArray(op)
		} else {
			var obj_ht *HashTable = ZendGetPropertiesFor(op, ZEND_PROP_PURPOSE_ARRAY_CAST)
			if obj_ht != nil {
				var new_obj_ht *HashTable = ZendProptableToSymtable(obj_ht, Z_OBJCE_P(op).GetDefaultPropertiesCount() != 0 || Z_OBJ_P(op).GetHandlers() != &StdObjectHandlers || obj_ht.IsRecursive())
				ZvalPtrDtor(op)
				op.SetArray(new_obj_ht)
				ZendReleaseProperties(obj_ht)
			} else {
				ZvalPtrDtor(op)

				/*ZVAL_EMPTY_ARRAY(op);*/

				ArrayInit(op)

				/*ZVAL_EMPTY_ARRAY(op);*/

			}
		}
	case IS_NULL:

		/*ZVAL_EMPTY_ARRAY(op);*/

		ArrayInit(op)
	case IS_REFERENCE:
		ZendUnwrapReference(op)
		goto try_again
	default:
		ConvertScalarToArray(op)
	}
}
func ConvertToObject(op *Zval) {
try_again:
	switch op.GetType() {
	case IS_ARRAY:
		var ht *HashTable = ZendSymtableToProptable(Z_ARR_P(op))
		var obj *ZendObject
		if (ht.GetGcFlags() & IS_ARRAY_IMMUTABLE) != 0 {

			/* TODO: try not to duplicate immutable arrays as well ??? */

			ht = ZendArrayDup(ht)

			/* TODO: try not to duplicate immutable arrays as well ??? */

		} else if ht != op.GetArr() {
			ZvalPtrDtor(op)
		} else {
			ht.DelRefcount()
		}
		obj = ZendObjectsNew(ZendStandardClassDef)
		obj.SetProperties(ht)
		op.SetObject(obj)
	case IS_OBJECT:

	case IS_NULL:
		ObjectInit(op)
	case IS_REFERENCE:
		ZendUnwrapReference(op)
		goto try_again
	default:
		var tmp Zval
		ZVAL_COPY_VALUE(&tmp, op)
		ObjectInit(op)
		Z_OBJPROP_P(op).KeyAddNew(ZSTR_KNOWN(ZEND_STR_SCALAR).GetStr(), &tmp)
	}
}
func _zvalGetLongFuncEx(op *Zval, silent ZendBool) ZendLong {
try_again:
	switch op.GetType() {
	case IS_UNDEF:
		fallthrough
	case IS_NULL:
		fallthrough
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
		fallthrough
	case IS_ARRAY:
		if Z_ARRVAL_P(op).GetNNumOfElements() {
			return 1
		} else {
			return 0
		}
		fallthrough
	case IS_OBJECT:
		var dst Zval
		ConvertObjectToType(op, &dst, IS_LONG, ConvertToLong)
		if dst.IsLong() {
			return dst.GetLval()
		} else {
			return 1
		}
		fallthrough
	case IS_REFERENCE:
		op = Z_REFVAL_P(op)
		goto try_again
	default:

	}
	return 0
}
func ZvalGetLongFunc(op *Zval) ZendLong       { return _zvalGetLongFuncEx(op, 1) }
func _zvalGetLongFuncNoisy(op *Zval) ZendLong { return _zvalGetLongFuncEx(op, 0) }
func ZvalGetDoubleFunc(op *Zval) float64 {
try_again:
	switch op.GetType() {
	case IS_NULL:
		fallthrough
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
		fallthrough
	case IS_OBJECT:
		var dst Zval
		ConvertObjectToType(op, &dst, IS_DOUBLE, ConvertToDouble)
		if dst.IsDouble() {
			return dst.GetDval()
		} else {
			return 1.0
		}
		fallthrough
	case IS_REFERENCE:
		op = Z_REFVAL_P(op)
		goto try_again
	default:

	}
	return 0.0
}
func __zvalGetStringFunc(op *Zval, try ZendBool) *ZendString {
try_again:
	switch op.GetType() {
	case IS_UNDEF:
		fallthrough
	case IS_NULL:
		fallthrough
	case IS_FALSE:
		return ZSTR_EMPTY_ALLOC()
	case IS_TRUE:
		return ZSTR_CHAR('1')
	case IS_RESOURCE:
		return ZendStrpprintf(0, "Resource id #"+ZEND_LONG_FMT, ZendLong(Z_RES_HANDLE_P(op)))
	case IS_LONG:
		return ZendLongToStr(op.GetLval())
	case IS_DOUBLE:
		return ZendStrpprintf(0, "%.*G", int(EG__().GetPrecision()), op.GetDval())
	case IS_ARRAY:
		ZendError(E_NOTICE, "Array to string conversion")
		if try != 0 && EG__().GetException() != nil {
			return nil
		} else {
			return ZSTR_KNOWN(ZEND_STR_ARRAY_CAPITALIZED)
		}
		fallthrough
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
		if EG__().GetException() == nil {
			ZendThrowError(nil, "Object of class %s could not be converted to string", Z_OBJCE_P(op).GetName().GetVal())
		}
		if try != 0 {
			return nil
		} else {
			return ZSTR_EMPTY_ALLOC()
		}
		fallthrough
	case IS_REFERENCE:
		op = Z_REFVAL_P(op)
		goto try_again
	case IS_STRING:
		return op.GetStr().Copy()
	default:

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
		result.SetArray(ZendArrayDup(op1.GetArr()))
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
		result.SetDouble(op1.GetDval() + op2.GetDval())
		return SUCCESS
	} else if type_pair == TYPE_PAIR(IS_LONG, IS_DOUBLE) {
		result.SetDouble(float64(op1.GetLval()) + op2.GetDval())
		return SUCCESS
	} else if type_pair == TYPE_PAIR(IS_DOUBLE, IS_LONG) {
		result.SetDouble(op1.GetDval() + float64(op2.GetLval()))
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
		if op1.IsReference() {
			op1 = Z_REFVAL_P(op1)
		} else if op2.IsReference() {
			op2 = Z_REFVAL_P(op2)
		} else if converted == 0 {
			if op1.IsObject() && op1 == result && Z_OBJ_HT(*op1).GetGet() != nil && Z_OBJ_HT(*op1).GetSet() != nil {
				var ret int
				var rv Zval
				var objval *Zval = Z_OBJ_HT(*op1).GetGet()(op1, &rv)
				objval.TryAddRefcount()
				ret = AddFunction(objval, objval, op2)
				Z_OBJ_HT(*op1).GetSet()(op1, objval)
				ZvalPtrDtor(objval)
				return ret
			} else if op1.IsObject() && Z_OBJ_HT(*op1).GetDoOperation() != nil {
				if SUCCESS == Z_OBJ_HT(*op1).GetDoOperation()(ZEND_ADD, result, op1, op2) {
					return SUCCESS
				}
			} else if op2.IsObject() && Z_OBJ_HT(*op2).GetDoOperation() != nil && SUCCESS == Z_OBJ_HT(*op2).GetDoOperation()(ZEND_ADD, result, op1, op2) {
				return SUCCESS
			}
			if op1 != op2 {
				op1 = ZendiConvertScalarToNumber(op1, &op1_copy, result, 0)
				op2 = ZendiConvertScalarToNumber(op2, &op2_copy, result, 0)
			} else {
				op1 = ZendiConvertScalarToNumber(op1, &op1_copy, result, 0)
				op2 = op1
			}
			if EG__().GetException() != nil {
				if result != op1 {
					result.SetUndef()
				}
				return FAILURE
			}
			converted = 1
		} else {
			if result != op1 {
				result.SetUndef()
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
		result.SetDouble(op1.GetDval() - op2.GetDval())
		return SUCCESS
	} else if type_pair == TYPE_PAIR(IS_LONG, IS_DOUBLE) {
		result.SetDouble(float64(op1.GetLval()) - op2.GetDval())
		return SUCCESS
	} else if type_pair == TYPE_PAIR(IS_DOUBLE, IS_LONG) {
		result.SetDouble(op1.GetDval() - float64(op2.GetLval()))
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
		if op1.IsReference() {
			op1 = Z_REFVAL_P(op1)
		} else if op2.IsReference() {
			op2 = Z_REFVAL_P(op2)
		} else if converted == 0 {
			if op1.IsObject() && op1 == result && Z_OBJ_HT(*op1).GetGet() != nil && Z_OBJ_HT(*op1).GetSet() != nil {
				var ret int
				var rv Zval
				var objval *Zval = Z_OBJ_HT(*op1).GetGet()(op1, &rv)
				objval.TryAddRefcount()
				ret = SubFunction(objval, objval, op2)
				Z_OBJ_HT(*op1).GetSet()(op1, objval)
				ZvalPtrDtor(objval)
				return ret
			} else if op1.IsObject() && Z_OBJ_HT(*op1).GetDoOperation() != nil {
				if SUCCESS == Z_OBJ_HT(*op1).GetDoOperation()(ZEND_SUB, result, op1, op2) {
					return SUCCESS
				}
			} else if op2.IsObject() && Z_OBJ_HT(*op2).GetDoOperation() != nil && SUCCESS == Z_OBJ_HT(*op2).GetDoOperation()(ZEND_SUB, result, op1, op2) {
				return SUCCESS
			}
			if op1 != op2 {
				op1 = ZendiConvertScalarToNumber(op1, &op1_copy, result, 0)
				op2 = ZendiConvertScalarToNumber(op2, &op2_copy, result, 0)
			} else {
				op1 = ZendiConvertScalarToNumber(op1, &op1_copy, result, 0)
				op2 = op1
			}
			if EG__().GetException() != nil {
				if result != op1 {
					result.SetUndef()
				}
				return FAILURE
			}
			converted = 1
		} else {
			if result != op1 {
				result.SetUndef()
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
			result.SetDouble(op1.GetDval() * op2.GetDval())
			return SUCCESS
		} else if type_pair == TYPE_PAIR(IS_LONG, IS_DOUBLE) {
			result.SetDouble(float64(op1.GetLval()) * op2.GetDval())
			return SUCCESS
		} else if type_pair == TYPE_PAIR(IS_DOUBLE, IS_LONG) {
			result.SetDouble(op1.GetDval() * float64(op2.GetLval()))
			return SUCCESS
		} else {
			if op1.IsReference() {
				op1 = Z_REFVAL_P(op1)
			} else if op2.IsReference() {
				op2 = Z_REFVAL_P(op2)
			} else if converted == 0 {
				if op1.IsObject() && op1 == result && Z_OBJ_HT(*op1).GetGet() != nil && Z_OBJ_HT(*op1).GetSet() != nil {
					var ret int
					var rv Zval
					var objval *Zval = Z_OBJ_HT(*op1).GetGet()(op1, &rv)
					objval.TryAddRefcount()
					ret = MulFunction(objval, objval, op2)
					Z_OBJ_HT(*op1).GetSet()(op1, objval)
					ZvalPtrDtor(objval)
					return ret
				} else if op1.IsObject() && Z_OBJ_HT(*op1).GetDoOperation() != nil {
					if SUCCESS == Z_OBJ_HT(*op1).GetDoOperation()(ZEND_MUL, result, op1, op2) {
						return SUCCESS
					}
				} else if op2.IsObject() && Z_OBJ_HT(*op2).GetDoOperation() != nil && SUCCESS == Z_OBJ_HT(*op2).GetDoOperation()(ZEND_MUL, result, op1, op2) {
					return SUCCESS
				}
				if op1 != op2 {
					op1 = ZendiConvertScalarToNumber(op1, &op1_copy, result, 0)
					op2 = ZendiConvertScalarToNumber(op2, &op2_copy, result, 0)
				} else {
					op1 = ZendiConvertScalarToNumber(op1, &op1_copy, result, 0)
					op2 = op1
				}
				if EG__().GetException() != nil {
					if result != op1 {
						result.SetUndef()
					}
					return FAILURE
				}
				converted = 1
			} else {
				if result != op1 {
					result.SetUndef()
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
					result.SetLong(1)
					return SUCCESS
				} else if l2 == 0 {
					result.SetLong(0)
					return SUCCESS
				}
				for i >= 1 {
					var overflow ZendLong
					var dval float64 = 0.0
					if i%2 != 0 {
						i--
						ZEND_SIGNED_MULTIPLY_LONG(l1, l2, l1, dval, overflow)
						if overflow != 0 {
							result.SetDouble(dval * pow(l2, i))
							return SUCCESS
						}
					} else {
						i /= 2
						ZEND_SIGNED_MULTIPLY_LONG(l2, l2, l2, dval, overflow)
						if overflow != 0 {
							result.SetDouble(float64(l1 * pow(dval, i)))
							return SUCCESS
						}
					}
				}

				/* i == 0 */

				result.SetLong(l1)

				/* i == 0 */

			} else {
				result.SetDouble(pow(float64(op1.GetLval()), float64(op2.GetLval())))
			}
			return SUCCESS
		} else if type_pair == TYPE_PAIR(IS_DOUBLE, IS_DOUBLE) {
			result.SetDouble(pow(op1.GetDval(), op2.GetDval()))
			return SUCCESS
		} else if type_pair == TYPE_PAIR(IS_LONG, IS_DOUBLE) {
			result.SetDouble(pow(float64(op1.GetLval()), op2.GetDval()))
			return SUCCESS
		} else if type_pair == TYPE_PAIR(IS_DOUBLE, IS_LONG) {
			result.SetDouble(pow(op1.GetDval(), float64(op2.GetLval())))
			return SUCCESS
		} else {
			if op1.IsReference() {
				op1 = Z_REFVAL_P(op1)
			} else if op2.IsReference() {
				op2 = Z_REFVAL_P(op2)
			} else if converted == 0 {
				if op1.IsObject() && op1 == result && Z_OBJ_HT(*op1).GetGet() != nil && Z_OBJ_HT(*op1).GetSet() != nil {
					var ret int
					var rv Zval
					var objval *Zval = Z_OBJ_HT(*op1).GetGet()(op1, &rv)
					objval.TryAddRefcount()
					ret = PowFunction(objval, objval, op2)
					Z_OBJ_HT(*op1).GetSet()(op1, objval)
					ZvalPtrDtor(objval)
					return ret
				} else if op1.IsObject() && Z_OBJ_HT(*op1).GetDoOperation() != nil {
					if SUCCESS == Z_OBJ_HT(*op1).GetDoOperation()(ZEND_POW, result, op1, op2) {
						return SUCCESS
					}
				} else if op2.IsObject() && Z_OBJ_HT(*op2).GetDoOperation() != nil && SUCCESS == Z_OBJ_HT(*op2).GetDoOperation()(ZEND_POW, result, op1, op2) {
					return SUCCESS
				}
				if op1 != op2 {
					if op1.IsArray() {
						if op1 == result {
							ZvalPtrDtor(result)
						}
						result.SetLong(0)
						return SUCCESS
					} else {
						op1 = ZendiConvertScalarToNumber(op1, &op1_copy, result, 0)
					}
					if op2.IsArray() {
						if op1 == result {
							ZvalPtrDtor(result)
						}
						result.SetLong(1)
						return SUCCESS
					} else {
						op2 = ZendiConvertScalarToNumber(op2, &op2_copy, result, 0)
					}
				} else {
					if op1.IsArray() {
						if op1 == result {
							ZvalPtrDtor(result)
						}
						result.SetLong(0)
						return SUCCESS
					} else {
						op1 = ZendiConvertScalarToNumber(op1, &op1_copy, result, 0)
					}
					op2 = op1
				}
				if EG__().GetException() != nil {
					if result != op1 {
						result.SetUndef()
					}
					return FAILURE
				}
				converted = 1
			} else {
				if result != op1 {
					result.SetUndef()
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
				result.SetDouble(float64(op1.GetLval() / float64(op2.GetLval())))
				return SUCCESS
			} else if op2.GetLval() == -1 && op1.GetLval() == ZEND_LONG_MIN {

				/* Prevent overflow error/crash */

				result.SetDouble(float64(ZEND_LONG_MIN / -1))
				return SUCCESS
			}
			if op1.GetLval()%op2.GetLval() == 0 {
				result.SetLong(op1.GetLval() / op2.GetLval())
			} else {
				result.SetDouble(float64(op1.GetLval()) / op2.GetLval())
			}
			return SUCCESS
		} else if type_pair == TYPE_PAIR(IS_DOUBLE, IS_DOUBLE) {
			if op2.GetDval() == 0 {
				ZendError(E_WARNING, "Division by zero")
			}
			result.SetDouble(op1.GetDval() / op2.GetDval())
			return SUCCESS
		} else if type_pair == TYPE_PAIR(IS_DOUBLE, IS_LONG) {
			if op2.GetLval() == 0 {
				ZendError(E_WARNING, "Division by zero")
			}
			result.SetDouble(op1.GetDval() / float64(op2.GetLval()))
			return SUCCESS
		} else if type_pair == TYPE_PAIR(IS_LONG, IS_DOUBLE) {
			if op2.GetDval() == 0 {
				ZendError(E_WARNING, "Division by zero")
			}
			result.SetDouble(float64(op1.GetLval() / op2.GetDval()))
			return SUCCESS
		} else {
			if op1.IsReference() {
				op1 = Z_REFVAL_P(op1)
			} else if op2.IsReference() {
				op2 = Z_REFVAL_P(op2)
			} else if converted == 0 {
				if op1.IsObject() && op1 == result && Z_OBJ_HT(*op1).GetGet() != nil && Z_OBJ_HT(*op1).GetSet() != nil {
					var ret int
					var rv Zval
					var objval *Zval = Z_OBJ_HT(*op1).GetGet()(op1, &rv)
					objval.TryAddRefcount()
					ret = DivFunction(objval, objval, op2)
					Z_OBJ_HT(*op1).GetSet()(op1, objval)
					ZvalPtrDtor(objval)
					return ret
				} else if op1.IsObject() && Z_OBJ_HT(*op1).GetDoOperation() != nil {
					if SUCCESS == Z_OBJ_HT(*op1).GetDoOperation()(ZEND_DIV, result, op1, op2) {
						return SUCCESS
					}
				} else if op2.IsObject() && Z_OBJ_HT(*op2).GetDoOperation() != nil && SUCCESS == Z_OBJ_HT(*op2).GetDoOperation()(ZEND_DIV, result, op1, op2) {
					return SUCCESS
				}
				if op1 != op2 {
					op1 = ZendiConvertScalarToNumber(op1, &op1_copy, result, 0)
					op2 = ZendiConvertScalarToNumber(op2, &op2_copy, result, 0)
				} else {
					op1 = ZendiConvertScalarToNumber(op1, &op1_copy, result, 0)
					op2 = op1
				}
				if EG__().GetException() != nil {
					if result != op1 {
						result.SetUndef()
					}
					return FAILURE
				}
				converted = 1
			} else {
				if result != op1 {
					result.SetUndef()
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
			if op1.IsReference() {
				op1 = Z_REFVAL_P(op1)
				if op1.IsLong() {
					op1_lval = op1.GetLval()
					break
				}
			}
			if op1.IsObject() && op1 == result && Z_OBJ_HT(*op1).GetGet() != nil && Z_OBJ_HT(*op1).GetSet() != nil {
				var ret int
				var rv Zval
				var objval *Zval = Z_OBJ_HT(*op1).GetGet()(op1, &rv)
				objval.TryAddRefcount()
				ret = ModFunction(objval, objval, op2)
				Z_OBJ_HT(*op1).GetSet()(op1, objval)
				ZvalPtrDtor(objval)
				return ret
			} else if op1.IsObject() && Z_OBJ_HT(*op1).GetDoOperation() != nil {
				if SUCCESS == Z_OBJ_HT(*op1).GetDoOperation()(ZEND_MOD, result, op1, op2) {
					return SUCCESS
				}
			}
			op1_lval = _zvalGetLongFuncNoisy(op1)
			if EG__().GetException() != nil {
				if result != op1 {
					result.SetUndef()
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
			if op2.IsReference() {
				op2 = Z_REFVAL_P(op2)
				if op2.IsLong() {
					op2_lval = op2.GetLval()
					break
				}
			}
			if op2.IsObject() && Z_OBJ_HT(*op2).GetDoOperation() != nil && SUCCESS == Z_OBJ_HT(*op2).GetDoOperation()(ZEND_MOD, result, op1, op2) {
				return SUCCESS
			}
			op2_lval = _zvalGetLongFuncNoisy(op2)
			if EG__().GetException() != nil {
				if result != op1 {
					result.SetUndef()
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

		if CurrEX() != nil && CG__().GetInCompilation() == 0 {
			ZendThrowExceptionEx(ZendCeDivisionByZeroError, 0, "Modulo by zero")
		} else {
			ZendErrorNoreturn(E_ERROR, "Modulo by zero")
		}
		if op1 != result {
			result.SetUndef()
		}
		return FAILURE
	}
	if op1 == result {
		ZvalPtrDtor(result)
	}
	if op2_lval == -1 {

		/* Prevent overflow error/crash if op1==LONG_MIN */

		result.SetLong(0)
		return SUCCESS
	}
	result.SetLong(op1_lval % op2_lval)
	return SUCCESS
}
func BooleanXorFunction(result *Zval, op1 *Zval, op2 *Zval) int {
	var op1_val int
	var op2_val int
	for {
		if op1.IsFalse() {
			op1_val = 0
		} else if op1.IsTrue() {
			op1_val = 1
		} else {
			if op1.IsReference() {
				op1 = Z_REFVAL_P(op1)
				if op1.IsFalse() {
					op1_val = 0
					break
				} else if op1.IsTrue() {
					op1_val = 1
					break
				}
			}
			if op1.IsObject() && op1 == result && Z_OBJ_HT(*op1).GetGet() != nil && Z_OBJ_HT(*op1).GetSet() != nil {
				var ret int
				var rv Zval
				var objval *Zval = Z_OBJ_HT(*op1).GetGet()(op1, &rv)
				objval.TryAddRefcount()
				ret = BooleanXorFunction(objval, objval, op2)
				Z_OBJ_HT(*op1).GetSet()(op1, objval)
				ZvalPtrDtor(objval)
				return ret
			} else if op1.IsObject() && Z_OBJ_HT(*op1).GetDoOperation() != nil {
				if SUCCESS == Z_OBJ_HT(*op1).GetDoOperation()(ZEND_BOOL_XOR, result, op1, op2) {
					return SUCCESS
				}
			}
			op1_val = ZvalIsTrue(op1)
		}
		break
	}
	for {
		if op2.IsFalse() {
			op2_val = 0
		} else if op2.IsTrue() {
			op2_val = 1
		} else {
			if op2.IsReference() {
				op2 = Z_REFVAL_P(op2)
				if op2.IsFalse() {
					op2_val = 0
					break
				} else if op2.IsTrue() {
					op2_val = 1
					break
				}
			}
			if op2.IsObject() && Z_OBJ_HT(*op2).GetDoOperation() != nil && SUCCESS == Z_OBJ_HT(*op2).GetDoOperation()(ZEND_BOOL_XOR, result, op1, op2) {
				return SUCCESS
			}
			op2_val = ZvalIsTrue(op2)
		}
		break
	}
	ZVAL_BOOL(result, (op1_val^op2_val) != 0)
	return SUCCESS
}
func BooleanNotFunction(result *Zval, op1 *Zval) int {
	if op1.GetType() < IS_TRUE {
		result.SetTrue()
	} else if op1.IsTrue() {
		result.SetFalse()
	} else {
		if op1.IsReference() {
			op1 = Z_REFVAL_P(op1)
			if op1.GetType() < IS_TRUE {
				result.SetTrue()
				return SUCCESS
			} else if op1.IsTrue() {
				result.SetFalse()
				return SUCCESS
			}
		}
		if op1.IsObject() && Z_OBJ_HT(*op1).GetDoOperation() != nil && SUCCESS == Z_OBJ_HT(*op1).GetDoOperation()(ZEND_BOOL_NOT, result, op1, nil) {
			return SUCCESS
		}
		ZVAL_BOOL(result, ZvalIsTrue(op1) == 0)
	}
	return SUCCESS
}
func BitwiseNotFunction(result *Zval, op1 *Zval) int {
try_again:
	switch op1.GetType() {
	case IS_LONG:
		result.SetLong(^(op1.GetLval()))
		return SUCCESS
	case IS_DOUBLE:
		result.SetLong(^(ZendDvalToLval(op1.GetDval())))
		return SUCCESS
	case IS_STRING:
		var i int
		if Z_STRLEN_P(op1) == 1 {
			var not ZendUchar = ZendUchar(^((*Z_STRVAL_P)(op1)))
			result.SetInternedString(ZSTR_CHAR(not))
		} else {
			result.SetString(ZendStringAlloc(Z_STRLEN_P(op1), 0))
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
		if op1.IsObject() && Z_OBJ_HT(*op1).GetDoOperation() != nil && SUCCESS == Z_OBJ_HT(*op1).GetDoOperation()(ZEND_BW_NOT, result, op1, nil) {
			return SUCCESS
		}
		if result != op1 {
			result.SetUndef()
		}
		ZendThrowError(nil, "Unsupported operand types")
		return FAILURE
	}
}
func BitwiseOrFunction(result *Zval, op1 *Zval, op2 *Zval) int {
	var op1_lval ZendLong
	var op2_lval ZendLong
	if op1.IsLong() && op2.IsLong() {
		result.SetLong(op1.GetLval() | op2.GetLval())
		return SUCCESS
	}
	op1 = ZVAL_DEREF(op1)
	op2 = ZVAL_DEREF(op2)
	if op1.IsString() && op2.IsString() {
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
				result.SetInternedString(ZSTR_CHAR(or))
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
		result.SetString(str)
		return SUCCESS
	}
	if op1.GetType() != IS_LONG {
		if op1.IsObject() && op1 == result && Z_OBJ_HT(*op1).GetGet() != nil && Z_OBJ_HT(*op1).GetSet() != nil {
			var ret int
			var rv Zval
			var objval *Zval = Z_OBJ_HT(*op1).GetGet()(op1, &rv)
			objval.TryAddRefcount()
			ret = BitwiseOrFunction(objval, objval, op2)
			Z_OBJ_HT(*op1).GetSet()(op1, objval)
			ZvalPtrDtor(objval)
			return ret
		} else if op1.IsObject() && Z_OBJ_HT(*op1).GetDoOperation() != nil {
			if SUCCESS == Z_OBJ_HT(*op1).GetDoOperation()(ZEND_BW_OR, result, op1, op2) {
				return SUCCESS
			}
		}
		op1_lval = _zvalGetLongFuncNoisy(op1)
		if EG__().GetException() != nil {
			if result != op1 {
				result.SetUndef()
			}
			return FAILURE
		}
	} else {
		op1_lval = op1.GetLval()
	}
	if op2.GetType() != IS_LONG {
		if op2.IsObject() && Z_OBJ_HT(*op2).GetDoOperation() != nil && SUCCESS == Z_OBJ_HT(*op2).GetDoOperation()(ZEND_BW_OR, result, op1, op2) {
			return SUCCESS
		}
		op2_lval = _zvalGetLongFuncNoisy(op2)
		if EG__().GetException() != nil {
			if result != op1 {
				result.SetUndef()
			}
			return FAILURE
		}
	} else {
		op2_lval = op2.GetLval()
	}
	if op1 == result {
		ZvalPtrDtor(result)
	}
	result.SetLong(op1_lval | op2_lval)
	return SUCCESS
}
func BitwiseAndFunction(result *Zval, op1 *Zval, op2 *Zval) int {
	var op1_lval ZendLong
	var op2_lval ZendLong
	if op1.IsLong() && op2.IsLong() {
		result.SetLong(op1.GetLval() & op2.GetLval())
		return SUCCESS
	}
	op1 = ZVAL_DEREF(op1)
	op2 = ZVAL_DEREF(op2)
	if op1.IsString() && op2.IsString() {
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
				result.SetInternedString(ZSTR_CHAR(and))
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
		result.SetString(str)
		return SUCCESS
	}
	if op1.GetType() != IS_LONG {
		if op1.IsObject() && op1 == result && Z_OBJ_HT(*op1).GetGet() != nil && Z_OBJ_HT(*op1).GetSet() != nil {
			var ret int
			var rv Zval
			var objval *Zval = Z_OBJ_HT(*op1).GetGet()(op1, &rv)
			objval.TryAddRefcount()
			ret = BitwiseAndFunction(objval, objval, op2)
			Z_OBJ_HT(*op1).GetSet()(op1, objval)
			ZvalPtrDtor(objval)
			return ret
		} else if op1.IsObject() && Z_OBJ_HT(*op1).GetDoOperation() != nil {
			if SUCCESS == Z_OBJ_HT(*op1).GetDoOperation()(ZEND_BW_AND, result, op1, op2) {
				return SUCCESS
			}
		}
		op1_lval = _zvalGetLongFuncNoisy(op1)
		if EG__().GetException() != nil {
			if result != op1 {
				result.SetUndef()
			}
			return FAILURE
		}
	} else {
		op1_lval = op1.GetLval()
	}
	if op2.GetType() != IS_LONG {
		if op2.IsObject() && Z_OBJ_HT(*op2).GetDoOperation() != nil && SUCCESS == Z_OBJ_HT(*op2).GetDoOperation()(ZEND_BW_AND, result, op1, op2) {
			return SUCCESS
		}
		op2_lval = _zvalGetLongFuncNoisy(op2)
		if EG__().GetException() != nil {
			if result != op1 {
				result.SetUndef()
			}
			return FAILURE
		}
	} else {
		op2_lval = op2.GetLval()
	}
	if op1 == result {
		ZvalPtrDtor(result)
	}
	result.SetLong(op1_lval & op2_lval)
	return SUCCESS
}
func BitwiseXorFunction(result *Zval, op1 *Zval, op2 *Zval) int {
	var op1_lval ZendLong
	var op2_lval ZendLong
	if op1.IsLong() && op2.IsLong() {
		result.SetLong(op1.GetLval() ^ op2.GetLval())
		return SUCCESS
	}
	op1 = ZVAL_DEREF(op1)
	op2 = ZVAL_DEREF(op2)
	if op1.IsString() && op2.IsString() {
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
				result.SetInternedString(ZSTR_CHAR(xor))
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
		result.SetString(str)
		return SUCCESS
	}
	if op1.GetType() != IS_LONG {
		if op1.IsObject() && op1 == result && Z_OBJ_HT(*op1).GetGet() != nil && Z_OBJ_HT(*op1).GetSet() != nil {
			var ret int
			var rv Zval
			var objval *Zval = Z_OBJ_HT(*op1).GetGet()(op1, &rv)
			objval.TryAddRefcount()
			ret = BitwiseXorFunction(objval, objval, op2)
			Z_OBJ_HT(*op1).GetSet()(op1, objval)
			ZvalPtrDtor(objval)
			return ret
		} else if op1.IsObject() && Z_OBJ_HT(*op1).GetDoOperation() != nil {
			if SUCCESS == Z_OBJ_HT(*op1).GetDoOperation()(ZEND_BW_XOR, result, op1, op2) {
				return SUCCESS
			}
		}
		op1_lval = _zvalGetLongFuncNoisy(op1)
		if EG__().GetException() != nil {
			if result != op1 {
				result.SetUndef()
			}
			return FAILURE
		}
	} else {
		op1_lval = op1.GetLval()
	}
	if op2.GetType() != IS_LONG {
		if op2.IsObject() && Z_OBJ_HT(*op2).GetDoOperation() != nil && SUCCESS == Z_OBJ_HT(*op2).GetDoOperation()(ZEND_BW_XOR, result, op1, op2) {
			return SUCCESS
		}
		op2_lval = _zvalGetLongFuncNoisy(op2)
		if EG__().GetException() != nil {
			if result != op1 {
				result.SetUndef()
			}
			return FAILURE
		}
	} else {
		op2_lval = op2.GetLval()
	}
	if op1 == result {
		ZvalPtrDtor(result)
	}
	result.SetLong(op1_lval ^ op2_lval)
	return SUCCESS
}
func ShiftLeftFunction(result *Zval, op1 *Zval, op2 *Zval) int {
	var op1_lval ZendLong
	var op2_lval ZendLong
	for {
		if op1.GetType() != IS_LONG {
			if op1.IsReference() {
				op1 = Z_REFVAL_P(op1)
				if op1.IsLong() {
					op1_lval = op1.GetLval()
					break
				}
			}
			if op1.IsObject() && op1 == result && Z_OBJ_HT(*op1).GetGet() != nil && Z_OBJ_HT(*op1).GetSet() != nil {
				var ret int
				var rv Zval
				var objval *Zval = Z_OBJ_HT(*op1).GetGet()(op1, &rv)
				objval.TryAddRefcount()
				ret = ShiftLeftFunction(objval, objval, op2)
				Z_OBJ_HT(*op1).GetSet()(op1, objval)
				ZvalPtrDtor(objval)
				return ret
			} else if op1.IsObject() && Z_OBJ_HT(*op1).GetDoOperation() != nil {
				if SUCCESS == Z_OBJ_HT(*op1).GetDoOperation()(ZEND_SL, result, op1, op2) {
					return SUCCESS
				}
			}
			op1_lval = _zvalGetLongFuncNoisy(op1)
			if EG__().GetException() != nil {
				if result != op1 {
					result.SetUndef()
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
			if op2.IsReference() {
				op2 = Z_REFVAL_P(op2)
				if op2.IsLong() {
					op2_lval = op2.GetLval()
					break
				}
			}
			if op2.IsObject() && Z_OBJ_HT(*op2).GetDoOperation() != nil && SUCCESS == Z_OBJ_HT(*op2).GetDoOperation()(ZEND_SL, result, op1, op2) {
				return SUCCESS
			}
			op2_lval = _zvalGetLongFuncNoisy(op2)
			if EG__().GetException() != nil {
				if result != op1 {
					result.SetUndef()
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
			result.SetLong(0)
			return SUCCESS
		} else {
			if CurrEX() != nil && CG__().GetInCompilation() == 0 {
				ZendThrowExceptionEx(ZendCeArithmeticError, 0, "Bit shift by negative number")
			} else {
				ZendErrorNoreturn(E_ERROR, "Bit shift by negative number")
			}
			if op1 != result {
				result.SetUndef()
			}
			return FAILURE
		}
	}
	if op1 == result {
		ZvalPtrDtor(result)
	}

	/* Perform shift on unsigned numbers to get well-defined wrap behavior. */

	result.SetLong(zend_long(ZendUlong(op1_lval << op2_lval)))
	return SUCCESS
}
func ShiftRightFunction(result *Zval, op1 *Zval, op2 *Zval) int {
	var op1_lval ZendLong
	var op2_lval ZendLong
	for {
		if op1.GetType() != IS_LONG {
			if op1.IsReference() {
				op1 = Z_REFVAL_P(op1)
				if op1.IsLong() {
					op1_lval = op1.GetLval()
					break
				}
			}
			if op1.IsObject() && op1 == result && Z_OBJ_HT(*op1).GetGet() != nil && Z_OBJ_HT(*op1).GetSet() != nil {
				var ret int
				var rv Zval
				var objval *Zval = Z_OBJ_HT(*op1).GetGet()(op1, &rv)
				objval.TryAddRefcount()
				ret = ShiftRightFunction(objval, objval, op2)
				Z_OBJ_HT(*op1).GetSet()(op1, objval)
				ZvalPtrDtor(objval)
				return ret
			} else if op1.IsObject() && Z_OBJ_HT(*op1).GetDoOperation() != nil {
				if SUCCESS == Z_OBJ_HT(*op1).GetDoOperation()(ZEND_SR, result, op1, op2) {
					return SUCCESS
				}
			}
			op1_lval = _zvalGetLongFuncNoisy(op1)
			if EG__().GetException() != nil {
				if result != op1 {
					result.SetUndef()
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
			if op2.IsReference() {
				op2 = Z_REFVAL_P(op2)
				if op2.IsLong() {
					op2_lval = op2.GetLval()
					break
				}
			}
			if op2.IsObject() && Z_OBJ_HT(*op2).GetDoOperation() != nil && SUCCESS == Z_OBJ_HT(*op2).GetDoOperation()(ZEND_SR, result, op1, op2) {
				return SUCCESS
			}
			op2_lval = _zvalGetLongFuncNoisy(op2)
			if EG__().GetException() != nil {
				if result != op1 {
					result.SetUndef()
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
			result.SetLong(b.Cond(op1_lval < 0, -1, 0))
			return SUCCESS
		} else {
			if CurrEX() != nil && CG__().GetInCompilation() == 0 {
				ZendThrowExceptionEx(ZendCeArithmeticError, 0, "Bit shift by negative number")
			} else {
				ZendErrorNoreturn(E_ERROR, "Bit shift by negative number")
			}
			if op1 != result {
				result.SetUndef()
			}
			return FAILURE
		}
	}
	if op1 == result {
		ZvalPtrDtor(result)
	}
	result.SetLong(op1_lval >> op2_lval)
	return SUCCESS
}
func ConcatFunction(result *Zval, op1 *Zval, op2 *Zval) int {
	var orig_op1 *Zval = op1
	var op1_copy Zval
	var op2_copy Zval
	op1_copy.SetUndef()
	op2_copy.SetUndef()
	for {
		if op1.GetType() != IS_STRING {
			if op1.IsReference() {
				op1 = Z_REFVAL_P(op1)
				if op1.IsString() {
					break
				}
			}
			if op1.IsObject() && op1 == result && Z_OBJ_HT(*op1).GetGet() != nil && Z_OBJ_HT(*op1).GetSet() != nil {
				var ret int
				var rv Zval
				var objval *Zval = Z_OBJ_HT(*op1).GetGet()(op1, &rv)
				objval.TryAddRefcount()
				ret = ConcatFunction(objval, objval, op2)
				Z_OBJ_HT(*op1).GetSet()(op1, objval)
				ZvalPtrDtor(objval)
				return ret
			} else if op1.IsObject() && Z_OBJ_HT(*op1).GetDoOperation() != nil {
				if SUCCESS == Z_OBJ_HT(*op1).GetDoOperation()(ZEND_CONCAT, result, op1, op2) {
					return SUCCESS
				}
			} else if op2.IsObject() && Z_OBJ_HT(*op2).GetDoOperation() != nil && SUCCESS == Z_OBJ_HT(*op2).GetDoOperation()(ZEND_CONCAT, result, op1, op2) {
				return SUCCESS
			}
			op1_copy.SetString(ZvalGetStringFunc(op1))
			if EG__().GetException() != nil {
				ZvalPtrDtorStr(&op1_copy)
				if orig_op1 != result {
					result.SetUndef()
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
			if op2.IsReference() {
				op2 = Z_REFVAL_P(op2)
				if op2.IsString() {
					break
				}
			}
			if op2.IsObject() && Z_OBJ_HT(*op2).GetDoOperation() != nil && SUCCESS == Z_OBJ_HT(*op2).GetDoOperation()(ZEND_CONCAT, result, op1, op2) {
				return SUCCESS
			}
			op2_copy.SetString(ZvalGetStringFunc(op2))
			if EG__().GetException() != nil {
				ZvalPtrDtorStr(&op1_copy)
				ZvalPtrDtorStr(&op2_copy)
				if orig_op1 != result {
					result.SetUndef()
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
				result.SetUndef()
			}
			return FAILURE
		}
		if result == op1 && result.IsRefcounted() {

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

		result.SetString(result_str)
		memcpy(result_str.GetVal()+op1_len, Z_STRVAL_P(op2), op2_len)
		result_str.GetVal()[result_len] = '0'
	}
	ZvalPtrDtorStr(&op1_copy)
	ZvalPtrDtorStr(&op2_copy)
	return SUCCESS
}
func StringCompareFunction(op1 *Zval, op2 *Zval) int {
	if op1.IsString() && op2.IsString() {
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
	if op1.IsString() && op2.IsString() {
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
	ZEND_ASSERT(!(op.IsRefcounted()) || op.GetRefcount() != 0)
	ZvalPtrDtor(op)
}
func ConvertCompareResultToLong(result *Zval) {
	if result.IsDouble() {
		result.SetLong(ZEND_NORMALIZE_BOOL(result.GetDval()))
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
			result.SetLong(b.CondF2(op1.GetLval() > op2.GetLval(), 1, func() int {
				if op1.GetLval() < op2.GetLval() {
					return -1
				} else {
					return 0
				}
			}))
			return SUCCESS
		case TYPE_PAIR(IS_DOUBLE, IS_LONG):
			result.SetDval(op1.GetDval() - float64(op2.GetLval()))
			result.SetLong(ZEND_NORMALIZE_BOOL(result.GetDval()))
			return SUCCESS
		case TYPE_PAIR(IS_LONG, IS_DOUBLE):
			result.SetDval(float64(op1.GetLval() - op2.GetDval()))
			result.SetLong(ZEND_NORMALIZE_BOOL(result.GetDval()))
			return SUCCESS
		case TYPE_PAIR(IS_DOUBLE, IS_DOUBLE):
			if op1.GetDval() == op2.GetDval() {
				result.SetLong(0)
			} else {
				result.SetDval(op1.GetDval() - op2.GetDval())
				result.SetLong(ZEND_NORMALIZE_BOOL(result.GetDval()))
			}
			return SUCCESS
		case TYPE_PAIR(IS_ARRAY, IS_ARRAY):
			result.SetLong(ZendCompareArrays(op1, op2))
			return SUCCESS
		case TYPE_PAIR(IS_NULL, IS_NULL):
			fallthrough
		case TYPE_PAIR(IS_NULL, IS_FALSE):
			fallthrough
		case TYPE_PAIR(IS_FALSE, IS_NULL):
			fallthrough
		case TYPE_PAIR(IS_FALSE, IS_FALSE):
			fallthrough
		case TYPE_PAIR(IS_TRUE, IS_TRUE):
			result.SetLong(0)
			return SUCCESS
		case TYPE_PAIR(IS_NULL, IS_TRUE):
			result.SetLong(-1)
			return SUCCESS
		case TYPE_PAIR(IS_TRUE, IS_NULL):
			result.SetLong(1)
			return SUCCESS
		case TYPE_PAIR(IS_STRING, IS_STRING):
			if op1.GetStr() == op2.GetStr() {
				result.SetLong(0)
				return SUCCESS
			}
			result.SetLong(ZendiSmartStrcmp(op1.GetStr(), op2.GetStr()))
			return SUCCESS
		case TYPE_PAIR(IS_NULL, IS_STRING):
			result.SetLong(b.Cond(Z_STRLEN_P(op2) == 0, 0, -1))
			return SUCCESS
		case TYPE_PAIR(IS_STRING, IS_NULL):
			result.SetLong(b.Cond(Z_STRLEN_P(op1) == 0, 0, 1))
			return SUCCESS
		case TYPE_PAIR(IS_OBJECT, IS_NULL):
			result.SetLong(1)
			return SUCCESS
		case TYPE_PAIR(IS_NULL, IS_OBJECT):
			result.SetLong(-1)
			return SUCCESS
		default:
			if op1.IsReference() {
				op1 = Z_REFVAL_P(op1)
				continue
			} else if op2.IsReference() {
				op2 = Z_REFVAL_P(op2)
				continue
			}
			if op1.IsObject() && Z_OBJ_HT(*op1).GetCompare() != nil {
				ret = Z_OBJ_HT(*op1).GetCompare()(result, op1, op2)
				if result.GetType() != IS_LONG {
					ConvertCompareResultToLong(result)
				}
				return ret
			} else if op2.IsObject() && Z_OBJ_HT(*op2).GetCompare() != nil {
				ret = Z_OBJ_HT(*op2).GetCompare()(result, op1, op2)
				if result.GetType() != IS_LONG {
					ConvertCompareResultToLong(result)
				}
				return ret
			}
			if op1.IsObject() && op2.IsObject() {
				if op1.GetObj() == op2.GetObj() {

					/* object handles are identical, apparently this is the same object */

					result.SetLong(0)
					return SUCCESS
				}
				if Z_OBJ_HT(*op1).GetCompareObjects() == Z_OBJ_HT(*op2).GetCompareObjects() {
					result.SetLong(Z_OBJ_HT(*op1).GetCompareObjects()(op1, op2))
					return SUCCESS
				}
			}
			if op1.IsObject() {
				if Z_OBJ_HT_P(op1).GetGet() != nil {
					var rv Zval
					op_free = Z_OBJ_HT_P(op1).GetGet()(op1, &rv)
					ret = CompareFunction(result, op_free, op2)
					ZendFreeObjGetResult(op_free)
					return ret
				} else if op2.GetType() != IS_OBJECT && Z_OBJ_HT_P(op1).GetCastObject() != nil {
					tmp_free.SetUndef()
					if Z_OBJ_HT_P(op1).GetCastObject()(op1, &tmp_free, b.CondF2(op2.IsFalse() || op2.IsTrue(), _IS_BOOL, func() __auto__ { return op2.GetType() })) == FAILURE {
						result.SetLong(1)
						ZendFreeObjGetResult(&tmp_free)
						return SUCCESS
					}
					ret = CompareFunction(result, &tmp_free, op2)
					ZendFreeObjGetResult(&tmp_free)
					return ret
				}
			}
			if op2.IsObject() {
				if Z_OBJ_HT_P(op2).GetGet() != nil {
					var rv Zval
					op_free = Z_OBJ_HT_P(op2).GetGet()(op2, &rv)
					ret = CompareFunction(result, op1, op_free)
					ZendFreeObjGetResult(op_free)
					return ret
				} else if op1.GetType() != IS_OBJECT && Z_OBJ_HT_P(op2).GetCastObject() != nil {
					tmp_free.SetUndef()
					if Z_OBJ_HT_P(op2).GetCastObject()(op2, &tmp_free, b.CondF2(op1.IsFalse() || op1.IsTrue(), _IS_BOOL, func() __auto__ { return op1.GetType() })) == FAILURE {
						result.SetLong(-1)
						ZendFreeObjGetResult(&tmp_free)
						return SUCCESS
					}
					ret = CompareFunction(result, op1, &tmp_free)
					ZendFreeObjGetResult(&tmp_free)
					return ret
				} else if op1.IsObject() {
					result.SetLong(1)
					return SUCCESS
				}
			}
			if converted == 0 {
				if op1.GetType() < IS_TRUE {
					result.SetLong(b.Cond(ZvalIsTrue(op2) != 0, -1, 0))
					return SUCCESS
				} else if op1.IsTrue() {
					result.SetLong(b.Cond(ZvalIsTrue(op2) != 0, 0, 1))
					return SUCCESS
				} else if op2.GetType() < IS_TRUE {
					result.SetLong(b.Cond(ZvalIsTrue(op1) != 0, 1, 0))
					return SUCCESS
				} else if op2.IsTrue() {
					result.SetLong(b.Cond(ZvalIsTrue(op1) != 0, 0, -1))
					return SUCCESS
				} else {
					op1 = ZendiConvertScalarToNumber(op1, &op1_copy, result, 1)
					op2 = ZendiConvertScalarToNumber(op2, &op2_copy, result, 1)
					if EG__().GetException() != nil {
						if result != op1 {
							result.SetUndef()
						}
						return FAILURE
					}
					converted = 1
				}
			} else if op1.IsArray() {
				result.SetLong(1)
				return SUCCESS
			} else if op2.IsArray() {
				result.SetLong(-1)
				return SUCCESS
			} else {
				ZEND_ASSERT(false)
				ZendThrowError(nil, "Unsupported operand types")
				if result != op1 {
					result.SetUndef()
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

	z1 = ZVAL_DEREF(z1)
	z2 = ZVAL_DEREF(z2)
	return FastIsNotIdenticalFunction(z1, z2)
}
func ZendIsIdentical(op1 *Zval, op2 *Zval) ZendBool {
	if op1.GetType() != op2.GetType() {
		return 0
	}
	switch op1.GetType() {
	case IS_NULL:
		fallthrough
	case IS_FALSE:
		fallthrough
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
	ZVAL_BOOL(result, ZendIsIdentical(op1, op2) != 0)
	return SUCCESS
}
func IsNotIdenticalFunction(result *Zval, op1 *Zval, op2 *Zval) int {
	ZVAL_BOOL(result, ZendIsIdentical(op1, op2) == 0)
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
		instance_ce = instance_ce.GetParent()
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
			if instance_ce.GetInterfaces()[i] == ce {
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
		str.SetInternedString(ZSTR_CHAR('1'))
		return
	}
	if !(str.IsRefcounted()) {
		str.SetStr(ZendStringInit(Z_STRVAL_P(str), Z_STRLEN_P(str), 0))
		str.SetTypeInfo(IS_STRING_EX)
	} else if str.GetRefcount() > 1 {
		str.DelRefcount()
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
		case UPPER_CASE:
			t.GetVal()[0] = 'A'
		case LOWER_CASE:
			t.GetVal()[0] = 'a'
		}
		ZendStringFree(str.GetStr())
		str.SetString(t)
	}
}
func IncrementFunction(op1 *Zval) int {
try_again:
	switch op1.GetType() {
	case IS_LONG:
		FastLongIncrementFunction(op1)
	case IS_DOUBLE:
		op1.SetDval(op1.GetDval() + 1)
	case IS_NULL:
		op1.SetLong(1)
	case IS_STRING:
		var lval ZendLong
		var dval float64
		switch IsNumericString(Z_STRVAL_P(op1), Z_STRLEN_P(op1), &lval, &dval, 0) {
		case IS_LONG:
			ZvalPtrDtorStr(op1)
			if lval == ZEND_LONG_MAX {

				/* switch to double */

				var d float64 = float64(lval)
				op1.SetDouble(d + 1)
			} else {
				op1.SetLong(lval + 1)
			}
		case IS_DOUBLE:
			ZvalPtrDtorStr(op1)
			op1.SetDouble(dval + 1)
		default:

			/* Perl style string increment */

			IncrementString(op1)
		}
	case IS_OBJECT:
		if Z_OBJ_HT(*op1).GetGet() != nil && Z_OBJ_HT(*op1).GetSet() != nil {

			/* proxy object */

			var rv Zval
			var val *Zval
			val = Z_OBJ_HT(*op1).GetGet()(op1, &rv)
			val.TryAddRefcount()
			IncrementFunction(val)
			Z_OBJ_HT(*op1).GetSet()(op1, val)
			ZvalPtrDtor(val)
		} else if Z_OBJ_HT(*op1).GetDoOperation() != nil {
			var op2 Zval
			var res int
			op2.SetLong(1)
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
	case IS_DOUBLE:
		op1.SetDval(op1.GetDval() - 1)
	case IS_STRING:
		if Z_STRLEN_P(op1) == 0 {
			ZvalPtrDtorStr(op1)
			op1.SetLong(-1)
			break
		}
		switch IsNumericString(Z_STRVAL_P(op1), Z_STRLEN_P(op1), &lval, &dval, 0) {
		case IS_LONG:
			ZvalPtrDtorStr(op1)
			if lval == ZEND_LONG_MIN {
				var d float64 = float64(lval)
				op1.SetDouble(d - 1)
			} else {
				op1.SetLong(lval - 1)
			}
		case IS_DOUBLE:
			ZvalPtrDtorStr(op1)
			op1.SetDouble(dval - 1)
		}
	case IS_OBJECT:
		if Z_OBJ_HT(*op1).GetGet() != nil && Z_OBJ_HT(*op1).GetSet() != nil {

			/* proxy object */

			var rv Zval
			var val *Zval
			val = Z_OBJ_HT(*op1).GetGet()(op1, &rv)
			val.TryAddRefcount()
			DecrementFunction(val)
			Z_OBJ_HT(*op1).GetSet()(op1, val)
			ZvalPtrDtor(val)
		} else if Z_OBJ_HT(*op1).GetDoOperation() != nil {
			var op2 Zval
			var res int
			op2.SetLong(1)
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
func ZendIsTrueEx(op *Zval) bool { return IZendIsTrue(op) != 0 }
func ZendIsTrue(op *Zval) int    { return IZendIsTrue(op) }
func ZendObjectIsTrue(op *Zval) int {
	if Z_OBJ_HT_P(op).GetCastObject() != nil {
		var tmp Zval
		if Z_OBJ_HT_P(op).GetCastObject()(op, &tmp, _IS_BOOL) == SUCCESS {
			return tmp.IsTrue()
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
	return str.Copy()
}
func ZendBinaryStrcmp(s1 *byte, len1 int, s2 *byte, len2 int) int {
	var retval int
	if s1 == s2 {
		return 0
	}
	retval = memcmp(s1, s2, b.Min(len1, len2))
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
	retval = memcmp(s1, s2, b.Min(length, b.Min(len1, len2)))
	if retval == 0 {
		return int(b.Min(length, len1) - b.Min(length, len2))
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
	len_ = b.Min(len1, len2)
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
	len_ = b.Min(length, b.Min(len1, len2))
	for b.PostDec(&len_) {
		c1 = ZendTolowerAscii(*((*uint8)(b.PostInc(&s1))))
		c2 = ZendTolowerAscii(*((*uint8)(b.PostInc(&s2))))
		if c1 != c2 {
			return c1 - c2
		}
	}
	return int(b.Min(length, len1) - b.Min(length, len2))
}
func ZendBinaryStrcasecmpL(s1 *byte, len1 int, s2 *byte, len2 int) int {
	var len_ int
	var c1 int
	var c2 int
	if s1 == s2 {
		return 0
	}
	len_ = b.Min(len1, len2)
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
	len_ = b.Min(length, b.Min(len1, len2))
	for b.PostDec(&len_) {
		c1 = ZendTolower(int(*((*uint8)(b.PostInc(&s1)))))
		c2 = ZendTolower(int(*((*uint8)(b.PostInc(&s2)))))
		if c1 != c2 {
			return c1 - c2
		}
	}
	return int(b.Min(length, len1) - b.Min(length, len2))
}
func ZendBinaryZvalStrcmp(s1 *Zval, s2 *Zval) int {
	return ZendBinaryStrcmp(Z_STRVAL_P(s1), Z_STRLEN_P(s1), Z_STRVAL_P(s2), Z_STRLEN_P(s2))
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

/**
 * ParseNumericStrEx 尝试转换字符串为数字
 * @param	str 		待转换的字符串
 * @param	allowErrors 是否允许错误。可选值为 0 不允许错误; 1 允许不完全匹配; 2 不完全匹配时触发 Notice
 * @return 	typ 		数字类型，可能值为 0, IS_LONG, IS_DOUBLE
 * @return 	lval		数字为整数时的值，默认为 0
 * @return 	dval		数字为浮点数时的值，默认为 0.0
 * @return 	overflowInfo	溢出信息。1 正数溢出，-1 负数溢出，0 无溢出或本身就是浮点数格式
 */
func ParseNumericStrEx(str string, allowErrors int) (typ ZendUchar, lval ZendLong, dval float64, overflowInfo int) {
	if len(str) == 0 {
		return
	} else if str[0] > '9' {
		// fast fail. 因为 digit | space | + | - 等都小于等于 '9'
		return
	}

	/* Skip any whitespace */
	str = strings.TrimLeft(str, " \t\n\r\v\f")

	// 扫描字符串，确认字符串为 整数|小数|非法字符串
	state := 0 // 状态机: 0 未开始, 1 整数部分; 2 小数部分; 3 指数部分
	i := 0
	for ; i < len(str); i++ {
		c := str[i]
		if IsDigit(c) {
			if state == 0 {
				state = 1
			}
			continue
		} else if c == '.' && (state == 0 || state == 1) { // 存在小数点，进入小数部分
			state = 2
			continue
		} else if (c == 'e' || c == 'E') && (state == 1 || state == 2) { // e|E + (+|-)? + 数字，进入指数部分
			ptr := i + 1
			// 跳过符号
			if ptr < len(str) && (str[ptr] == '+' || str[ptr] == '-') {
				ptr++
			}
			// 判断是否接数字，若是则进入指数部分
			if ptr < len(str) && IsDigit(str[ptr]) {
				state = 3
				i = ptr
				continue
			}
		}
		// 未匹配任何内容
		break
	}
	// 未匹配时
	if state == 0 {
		return
	}
	// 未完成匹配时
	if i != len(str) {
		if allowErrors == 0 {
			return
		}
		if allowErrors == -1 {
			ZendError(E_NOTICE, "A non well formed numeric value encountered")
			if EG__().GetException() != nil {
				return
			}
		}
	}
	// 转义匹配字符串
	matchStr := str[:i]
	if state == 1 {
		// 尝试转 int，若成功直接返回
		if len(matchStr) < MAX_LENGTH_OF_LONG {
			tmpVal, err := strconv.Atoi(matchStr)
			if err == nil {
				typ, lval = IS_LONG, tmpVal
				return
			}
		}
		// 整数溢出, 记录溢出信息
		if matchStr[0] == '-' {
			overflowInfo = -1
		} else {
			overflowInfo = 1
		}
	}

	tmpDVal, err := strconv.ParseFloat(matchStr, 64)
	if err != nil {
		log.Panicf("代码逻辑错误，预期为数字字符串，但转换失败了: s=%s ,err=%s", matchStr, err.Error())
	}
	typ, dval = IS_DOUBLE, tmpDVal
	return
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
