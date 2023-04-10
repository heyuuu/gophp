package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/builtin/ascii"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
	"math"
	"strconv"
	"strings"
)

func IsFinite(f float64) bool {
	if math.IsNaN(f) || math.IsInf(f, 1) || math.IsInf(f, -1) {
		return false
	}
	return true
}

func DoubleFitsLong(d float64) bool {
	return !(d >= ZEND_LONG_MAX || d < ZEND_LONG_MIN)
}
func DvalToLval(d float64) ZendLong {
	if !IsFinite(d) {
		return 0
	} else if !(DoubleFitsLong(d)) {
		return ZendDvalToLvalSlow(d)
	}
	return ZendLong(d)
}
func DvalToLvalCap(d float64) ZendLong {
	if !IsFinite(d) {
		return 0
	} else if !(DoubleFitsLong(d)) {
		if d > 0 {
			return ZEND_LONG_MAX
		} else {
			return ZEND_LONG_MIN
		}
	}
	return ZendLong(d)
}
func IsNumericStringEx(
	str string,
	lval *ZendLong,
	dval *float64,
	allow_errors int,
	oflow_info *int,
) types.ZendUchar {
	r := ConvertNumericStr(str, allow_errors)

	*lval = r.Lval
	if dval != nil {
		*dval = r.Dval
	}
	if oflow_info != nil {
		*oflow_info = r.Overflow
	}
	return r.Type
}
func IsNumericString(str string, lval *ZendLong, dval *float64, allow_errors int) types.ZendUchar {
	r := ConvertNumericStr(str, allow_errors)

	*lval = r.Lval
	if dval != nil {
		*dval = r.Dval
	}
	return r.Type
}
func ZendMemnstr(haystack *byte, needle string, needle_len int, end *byte) *byte {
	// todo 替换 - 查找haystack中needle首次出现的位置，没出现则返回nil
	pos := strings.Index(b.CastStr(haystack, end-haystack), b.CastStr(needle))
	if pos < 0 {
		return nil
	}
	return haystack + pos
}
func ZendMemrchr(s *byte, c byte, n int) *byte {
	str := b.CastStr(s, n)

	if pos := strings.LastIndexByte(str, c); pos >= 0 {
		return s + pos
	} else {
		return nil
	}
}

func ZvalGetLong(op *types.Zval) ZendLong {
	if op.IsLong() {
		return op.GetLval()
	} else {
		return ZvalGetLongFunc(op)
	}
}
func ZvalGetDouble(op *types.Zval) float64 {
	if op.IsDouble() {
		return op.GetDval()
	} else {
		return ZvalGetDoubleFunc(op)
	}
}

func ZvalGetStrVal(op *types.Zval) string {
	if op.IsString() {
		return op.GetStrVal()
	} else {
		zstr := ZvalGetStringFunc(op)
		if zstr == nil {
			return ""
		}
		return zstr.GetStr()
	}
}

func ZvalGetString(op *types.Zval) *types.String {
	if op.IsString() {
		return op.GetStr().Copy()
	} else {
		return ZvalGetStringFunc(op)
	}
}
func ZvalGetTmpString(op *types.Zval, tmp **types.String) *types.String {
	if op.IsString() {
		*tmp = nil
		return op.GetStr()
	} else {
		*tmp = ZvalGetStringFunc(op)
		return *tmp
	}
}
func ZendTmpStringRelease(tmp *types.String) {
	if tmp != nil {
		// types.ZendStringReleaseEx(tmp, 0)
	}
}
func ZvalTryGetString(op *types.Zval) *types.String {
	if op.IsString() {
		var ret *types.String = op.GetStr().Copy()
		return ret
	} else {
		return ZvalTryGetStringFunc(op)
	}
}
func ZvalTryGetTmpString(op *types.Zval, tmp **types.String) *types.String {
	if op.IsString() {
		var ret *types.String = op.GetStr()
		*tmp = nil
		return ret
	} else {
		*tmp = ZvalTryGetStringFunc(op)
		return *tmp
	}
}
func TryConvertToString(op *types.Zval) types.ZendBool {
	if op.IsString() {
		return 1
	}
	return _tryConvertToString(op)
}
func ConvertToString(op *types.Zval) {
	if op.GetType() != types.IS_STRING {
		_convertToString(op)
	}
}
func ZvalIsTrue(op *types.Zval) int { return ZendIsTrue(op) }
func IZendIsTrue(op *types.Zval) int {
	var result int = 0
again:
	switch op.GetType() {
	case types.IS_TRUE:
		result = 1
	case types.IS_LONG:
		if op.GetLval() != 0 {
			result = 1
		}
	case types.IS_DOUBLE:
		if op.GetDval() != 0 {
			result = 1
		}
	case types.IS_STRING:
		if op.GetStr().GetLen() > 1 || op.GetStr().GetLen() != 0 && op.GetStr().GetVal()[0] != '0' {
			result = 1
		}
	case types.IS_ARRAY:
		if types.Z_ARRVAL_P(op).Len() {
			result = 1
		}
	case types.IS_OBJECT:
		if types.Z_OBJ_HT_P(op).GetCastObject() == ZendStdCastObjectTostring {
			result = 1
		} else {
			result = ZendObjectIsTrue(op)
		}
	case types.IS_RESOURCE:
		if types.Z_RES_HANDLE_P(op) != 0 {
			result = 1
		}
	case types.IS_REFERENCE:
		op = types.Z_REFVAL_P(op)
		goto again
	default:

	}
	return result
}
func IZendIsTrueEx(op *types.Zval) bool {
again:
	switch op.GetType() {
	case types.IS_TRUE:
		return true
	case types.IS_LONG:
		return op.GetLval() != 0
	case types.IS_DOUBLE:
		return op.GetDval() != 0
	case types.IS_STRING:
		str := op.GetStrVal()
		return str != "" && str != "0"
	case types.IS_ARRAY:
		return op.GetArr().Len() != 0
	case types.IS_OBJECT:
		if types.Z_OBJ_HT_P(op).GetCastObject() == ZendStdCastObjectTostring {
			return true
		} else {
			return ZendObjectIsTrue(op)
		}
	case types.IS_RESOURCE:
		return types.Z_RES_HANDLE_P(op) != 0
	case types.IS_REFERENCE:
		op = types.Z_REFVAL_P(op)
		goto again
	}
	return false
}
func ZendStringTolower(str *types.String) *types.String { return ZendStringTolowerEx(str) }
func ConvertToStringEx(pzv *types.Zval) {
	if pzv.GetType() != types.IS_STRING {
		ConvertToString(pzv)
	}
}
func ConvertToArrayEx(pzv *types.Zval) {
	if pzv.GetType() != types.IS_ARRAY {
		ConvertToArray(pzv)
	}
}
func ConvertScalarToNumberEx(pzv *types.Zval) {
	if pzv.GetType() != types.IS_LONG && pzv.GetType() != types.IS_DOUBLE {
		ConvertScalarToNumber(pzv)
	}
}
func FastLongIncrementFunction(op1 *types.Zval) {
	if op1.GetLval() == ZEND_LONG_MAX {

		/* switch to double */

		op1.SetDouble(float64(ZEND_LONG_MAX + 1.0))

		/* switch to double */

	} else {
		op1.GetLval()++
	}
}
func FastLongDecrementFunction(op1 *types.Zval) {
	if op1.GetLval() == ZEND_LONG_MIN {

		/* switch to double */

		op1.SetDouble(float64(ZEND_LONG_MIN - 1.0))

		/* switch to double */

	} else {
		op1.GetLval()--
	}
}
func FastLongAddFunction(result *types.Zval, op1 *types.Zval, op2 *types.Zval) {
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
}
func FastAddFunction(result *types.Zval, op1 *types.Zval, op2 *types.Zval) int {
	if op1.IsLong() {
		if op2.IsLong() {
			FastLongAddFunction(result, op1, op2)
			return types.SUCCESS
		} else if op2.IsDouble() {
			result.SetDouble(float64(op1.GetLval()) + op2.GetDval())
			return types.SUCCESS
		}
	} else if op1.IsDouble() {
		if op2.IsDouble() {
			result.SetDouble(op1.GetDval() + op2.GetDval())
			return types.SUCCESS
		} else if op2.IsLong() {
			result.SetDouble(op1.GetDval() + float64(op2.GetLval()))
			return types.SUCCESS
		}
	}
	return AddFunction(result, op1, op2)
}
func FastLongSubFunction(result *types.Zval, op1 *types.Zval, op2 *types.Zval) {
	result.SetLong(op1.GetLval() - op2.GetLval())
	if (op1.GetLval()&LONG_SIGN_MASK) != (op2.GetLval()&LONG_SIGN_MASK) && (op1.GetLval()&LONG_SIGN_MASK) != (result.GetLval()&LONG_SIGN_MASK) {
		result.SetDouble(float64(op1.GetLval() - float64(op2.GetLval())))
	}
}
func FastDivFunction(result *types.Zval, op1 *types.Zval, op2 *types.Zval) int {
	return DivFunction(result, op1, op2)
}
func ZendFastEqualStrings(s1 *types.String, s2 *types.String) int {
	if s1 == s2 {
		return 1
	} else if s1.GetVal()[0] > '9' || s2.GetVal()[0] > '9' {
		return types.IntBool(s1.GetStr() == s2.GetStr())
	} else {
		return ZendiSmartStreq(s1, s2)
	}
}
func FastEqualCheckFunction(op1 *types.Zval, op2 *types.Zval) int {
	var result types.Zval
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
func FastEqualCheckLong(op1 *types.Zval, op2 *types.Zval) int {
	var result types.Zval
	if op2.IsLong() {
		return op1.GetLval() == op2.GetLval()
	}
	CompareFunction(&result, op1, op2)
	return result.GetLval() == 0
}
func FastEqualCheckString(op1 *types.Zval, op2 *types.Zval) int {
	var result types.Zval
	if op2.IsString() {
		return ZendFastEqualStrings(op1.GetStr(), op2.GetStr())
	}
	CompareFunction(&result, op1, op2)
	return result.GetLval() == 0
}
func FastIsIdenticalFunction(op1 *types.Zval, op2 *types.Zval) types.ZendBool {
	if op1.GetType() != op2.GetType() {
		return 0
	} else if op1.GetType() <= types.IS_TRUE {
		return 1
	}
	return types.IntBool(ZendIsIdentical(op1, op2))
}
func FastIsNotIdenticalFunction(op1 *types.Zval, op2 *types.Zval) types.ZendBool {
	if op1.GetType() != op2.GetType() {
		return 1
	} else if op1.GetType() <= types.IS_TRUE {
		return 0
	}
	return types.IntBool(!(ZendIsIdentical(op1, op2)))
}
func ZendUnwrapReference(op *types.Zval) {
	if op.GetRefcount() == 1 {
		types.ZVAL_UNREF(op)
	} else {
		op.DelRefcount()
		types.ZVAL_COPY(op, types.Z_REFVAL_P(op))
	}
}
func TYPE_PAIR(t1 types.ZvalType, t2 types.ZvalType) uint { return uint(t1)<<4 | uint(t2) }
func ConvertObjectToType(op *types.Zval, dst *types.Zval, ctype int, conv_func func(op *types.Zval)) {
	dst.SetUndef()
	if types.Z_OBJ_HT_P(op).GetCastObject() != nil {
		if types.Z_OBJ_HT_P(op).GetCastObject()(op, dst, ctype) == types.FAILURE {
			faults.Error(faults.E_RECOVERABLE_ERROR, "Object of class %s could not be converted to %s", types.Z_OBJCE_P(op).GetName().GetVal(), types.ZendGetTypeByConst(ctype))
		}
	} else if types.Z_OBJ_HT_P(op).GetGet() != nil {
		var newop *types.Zval = types.Z_OBJ_HT_P(op).GetGet()(op, dst)
		if newop.GetType() != types.IS_OBJECT {
			types.ZVAL_COPY_VALUE(dst, newop)
			conv_func(dst)
		}
	}
}
func _convertScalarToNumber(op *types.Zval, silent types.ZendBool, check types.ZendBool) {
try_again:
	switch op.GetType() {
	case types.IS_REFERENCE:
		ZendUnwrapReference(op)
		goto try_again
	case types.IS_STRING:
		var str *types.String
		str = op.GetStr()
		if b.Assign(&(op.GetTypeInfo()), IsNumericString(str.GetStr(), &(op.GetLval()), &(op.GetDval()), b.Cond(silent != 0, 1, -1))) == 0 {
			op.SetLong(0)
			if silent == 0 {
				faults.Error(faults.E_WARNING, "A non-numeric value encountered")
			}
		}
		// types.ZendStringReleaseEx(str, 0)
	case types.IS_NULL:
		fallthrough
	case types.IS_FALSE:
		op.SetLong(0)
	case types.IS_TRUE:
		op.SetLong(1)
	case types.IS_RESOURCE:
		var l ZendLong = types.Z_RES_HANDLE_P(op)
		ZvalPtrDtor(op)
		op.SetLong(l)
	case types.IS_OBJECT:
		var dst types.Zval
		ConvertObjectToType(op, &dst, types.IS_NUMBER, ConvertScalarToNumber)
		if check != 0 && EG__().GetException() != nil {
			return
		}
		ZvalPtrDtor(op)
		if dst.IsLong() || dst.IsDouble() {
			types.ZVAL_COPY_VALUE(op, &dst)
		} else {
			op.SetLong(1)
		}
	}
}
func ConvertScalarToNumber(op *types.Zval) { _convertScalarToNumber(op, 1, 0) }
func _zendiConvertScalarToNumberEx(op *types.Zval, holder *types.Zval, silent types.ZendBool) *types.Zval {
	switch op.GetType() {
	case types.IS_NULL:
		fallthrough
	case types.IS_FALSE:
		holder.SetLong(0)
		return holder
	case types.IS_TRUE:
		holder.SetLong(1)
		return holder
	case types.IS_STRING:
		var mode ConvertNumericMode
		if silent != 0 {
			mode = ConvertContinueOnErrors
		} else {
			mode = ConvertNoticeOnErrors
		}
		r := ConvertNumericStr(op.GetStrVal(), mode)
		switch r.Type {
		case types.IS_LONG:
			holder.SetLong(r.Lval)
		case types.IS_DOUBLE:
			holder.SetDouble(r.Dval)
		default:
			holder.SetLong(0)
			if silent == 0 {
				faults.Error(faults.E_WARNING, "A non-numeric value encountered")
			}
		}
		return holder
	case types.IS_RESOURCE:
		holder.SetLong(types.Z_RES_HANDLE_P(op))
		return holder
	case types.IS_OBJECT:
		ConvertObjectToType(op, holder, types.IS_NUMBER, ConvertScalarToNumber)
		if EG__().GetException() != nil || holder.GetType() != types.IS_LONG && holder.GetType() != types.IS_DOUBLE {
			holder.SetLong(1)
		}
		return holder
	case types.IS_LONG:
		fallthrough
	case types.IS_DOUBLE:
		fallthrough
	default:
		return op
	}
}
func _zendiConvertScalarToNumber(op *types.Zval, holder *types.Zval) *types.Zval {
	return _zendiConvertScalarToNumberEx(op, holder, 1)
}
func _zendiConvertScalarToNumberNoisy(op *types.Zval, holder *types.Zval) *types.Zval {
	return _zendiConvertScalarToNumberEx(op, holder, 0)
}
func ZendiConvertScalarToNumber(op *types.Zval, holder *types.Zval, result *types.Zval, silent types.ZendBool) *types.Zval {
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
func ConvertToLong(op *types.Zval) {
	if !op.IsLong() {
		ConvertToLongBase(op, 10)
	}
}
func ConvertToLongBase(op *types.Zval, base int) {
	var tmp ZendLong
try_again:
	switch op.GetType() {
	case types.IS_NULL:
		fallthrough
	case types.IS_FALSE:
		op.SetLong(0)
	case types.IS_TRUE:
		op.SetLong(1)
	case types.IS_RESOURCE:
		tmp = types.Z_RES_HANDLE_P(op)
		ZvalPtrDtor(op)
		op.SetLong(tmp)
	case types.IS_LONG:

	case types.IS_DOUBLE:
		op.SetLong(DvalToLval(op.GetDval()))
	case types.IS_STRING:
		var str *types.String = op.GetStr()
		if base == 10 {
			op.SetLong(ZvalGetLong(op))
		} else {
			op.SetLong(ZEND_STRTOL(str.GetVal(), nil, base))
		}
		// types.ZendStringReleaseEx(str, 0)
	case types.IS_ARRAY:
		if types.Z_ARRVAL_P(op).Len() {
			tmp = 1
		} else {
			tmp = 0
		}
		ZvalPtrDtor(op)
		op.SetLong(tmp)
	case types.IS_OBJECT:
		var dst types.Zval
		ConvertObjectToType(op, &dst, types.IS_LONG, ConvertToLong)
		ZvalPtrDtor(op)
		if dst.IsLong() {
			op.SetLong(dst.GetLval())
		} else {
			op.SetLong(1)
		}
		return
	case types.IS_REFERENCE:
		ZendUnwrapReference(op)
		goto try_again
	default:

	}
}
func ConvertToDouble(op *types.Zval) {
	var tmp float64
try_again:
	switch op.GetType() {
	case types.IS_NULL:
		fallthrough
	case types.IS_FALSE:
		op.SetDouble(0.0)
	case types.IS_TRUE:
		op.SetDouble(1.0)
	case types.IS_RESOURCE:
		var d float64 = float64(types.Z_RES_HANDLE_P(op))
		ZvalPtrDtor(op)
		op.SetDouble(d)
	case types.IS_LONG:
		op.SetDouble(float64(op.GetLval()))
	case types.IS_DOUBLE:

	case types.IS_STRING:
		var str *types.String = op.GetStr()
		op.SetDouble(ZendStrtod(str.GetVal(), nil))
		// types.ZendStringReleaseEx(str, 0)
	case types.IS_ARRAY:
		if types.Z_ARRVAL_P(op).Len() {
			tmp = 1
		} else {
			tmp = 0
		}
		ZvalPtrDtor(op)
		op.SetDouble(tmp)
	case types.IS_OBJECT:
		var dst types.Zval
		ConvertObjectToType(op, &dst, types.IS_DOUBLE, ConvertToDouble)
		ZvalPtrDtor(op)
		if dst.IsDouble() {
			op.SetDouble(dst.GetDval())
		} else {
			op.SetDouble(1.0)
		}
	case types.IS_REFERENCE:
		ZendUnwrapReference(op)
		goto try_again
	default:

	}
}
func ConvertToNull(op *types.Zval) {
	ZvalPtrDtor(op)
	op.SetNull()
}
func ConvertToBoolean(op *types.Zval) {
	var tmp int
try_again:
	switch op.GetType() {
	case types.IS_FALSE:
		fallthrough
	case types.IS_TRUE:

	case types.IS_NULL:
		op.SetFalse()
	case types.IS_RESOURCE:
		var l ZendLong = b.Cond(types.Z_RES_HANDLE_P(op) != 0, 1, 0)
		ZvalPtrDtor(op)
		types.ZVAL_BOOL(op, l != 0)
	case types.IS_LONG:
		types.ZVAL_BOOL(op, op.GetLval() != 0)
	case types.IS_DOUBLE:
		types.ZVAL_BOOL(op, op.GetDval() != 0)
	case types.IS_STRING:
		var str *types.String = op.GetStr()
		if str.GetLen() == 0 || str.GetLen() == 1 && str.GetVal()[0] == '0' {
			op.SetFalse()
		} else {
			op.SetTrue()
		}
		// types.ZendStringReleaseEx(str, 0)
	case types.IS_ARRAY:
		if types.Z_ARRVAL_P(op).Len() != 0 {
			tmp = 1
		} else {
			tmp = 0
		}
		ZvalPtrDtor(op)
		types.ZVAL_BOOL(op, tmp != 0)
	case types.IS_OBJECT:
		var dst types.Zval
		ConvertObjectToType(op, &dst, types.IS_BOOL, ConvertToBoolean)
		ZvalPtrDtor(op)
		if dst.IsFalse() || dst.IsTrue() {
			op.SetTypeInfo(dst.GetTypeInfo())
		} else {
			op.SetTrue()
		}
	case types.IS_REFERENCE:
		ZendUnwrapReference(op)
		goto try_again
	default:

	}
}

func _convertToString(op *types.Zval) {
try_again:
	switch op.GetType() {
	case types.IS_UNDEF:
		fallthrough
	case types.IS_NULL:
		fallthrough
	case types.IS_FALSE:
		op.SetStringVal("")
	case types.IS_TRUE:
		op.SetStringVal(string('1'))
	case types.IS_STRING:

	case types.IS_RESOURCE:
		var str *types.String = ZendStrpprintf(0, "Resource id #"+ZEND_LONG_FMT, ZendLong(types.Z_RES_HANDLE_P(op)))
		ZvalPtrDtor(op)
		op.SetString(str)
	case types.IS_LONG:
		op.SetString(ZendLongToStr(op.GetLval()))
	case types.IS_DOUBLE:
		var str *types.String
		var dval = op.GetDval()
		str = ZendStrpprintf(0, "%.*G", int(EG__().GetPrecision()), dval)

		/* %G already handles removing trailing zeros from the fractional part, yay */

		op.SetString(str)
	case types.IS_ARRAY:
		faults.Error(faults.E_NOTICE, "Array to string conversion")
		ZvalPtrDtor(op)
		op.SetStringVal(types.STR_ARRAY_CAPITALIZED)
	case types.IS_OBJECT:
		var tmp types.Zval
		if types.Z_OBJ_HT_P(op).GetCastObject() != nil {
			if types.Z_OBJ_HT_P(op).GetCastObject()(op, &tmp, types.IS_STRING) == types.SUCCESS {
				ZvalPtrDtor(op)
				types.ZVAL_COPY_VALUE(op, &tmp)
				return
			}
		} else if types.Z_OBJ_HT_P(op).GetGet() != nil {
			var z *types.Zval = types.Z_OBJ_HT_P(op).GetGet()(op, &tmp)
			if z.GetType() != types.IS_OBJECT {
				var str *types.String = ZvalGetString(z)
				ZvalPtrDtor(z)
				ZvalPtrDtor(op)
				op.SetString(str)
				return
			}
			ZvalPtrDtor(z)
		}
		if EG__().GetException() == nil {
			faults.ThrowError(nil, "Object of class %s could not be converted to string", types.Z_OBJCE_P(op).GetName().GetVal())
		}
		ZvalPtrDtor(op)
		op.SetStringVal("")
	case types.IS_REFERENCE:
		ZendUnwrapReference(op)
		goto try_again
	default:

	}
}
func _tryConvertToString(op *types.Zval) types.ZendBool {
	var str *types.String
	b.Assert(op.GetType() != types.IS_STRING)
	str = ZvalTryGetStringFunc(op)
	if str == nil {
		return 0
	}
	ZvalPtrDtor(op)
	op.SetString(str)
	return 1
}
func ConvertScalarToArray(op *types.Zval) {
	var ht *types.Array = types.NewArray(1)
	ht.IndexAddNew(0, op)
	op.SetArray(ht)
}
func ConvertToArray(op *types.Zval) {
try_again:
	switch op.GetType() {
	case types.IS_ARRAY:

	case types.IS_OBJECT:
		if types.Z_OBJCE_P(op) == ZendCeClosure {
			ConvertScalarToArray(op)
		} else {
			var obj_ht *types.Array = ZendGetPropertiesFor(op, ZEND_PROP_PURPOSE_ARRAY_CAST)
			if obj_ht != nil {
				var new_obj_ht *types.Array = types.ZendProptableToSymtable(obj_ht, types.Z_OBJCE_P(op).GetDefaultPropertiesCount() != 0 || types.Z_OBJ_P(op).GetHandlers() != &StdObjectHandlers || obj_ht.IsRecursive())
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
	case types.IS_NULL:

		/*ZVAL_EMPTY_ARRAY(op);*/

		ArrayInit(op)
	case types.IS_REFERENCE:
		ZendUnwrapReference(op)
		goto try_again
	default:
		ConvertScalarToArray(op)
	}
}
func ConvertToObject(op *types.Zval) {
try_again:
	switch op.GetType() {
	case types.IS_ARRAY:
		var ht = types.ZendSymtableToProptable(op.GetArr())
		var obj *types.ZendObject
		if (ht.GetGcFlags() & types.IS_ARRAY_IMMUTABLE) != 0 {

			/* TODO: try not to duplicate immutable arrays as well ??? */

			ht = types.ZendArrayDup(ht)

			/* TODO: try not to duplicate immutable arrays as well ??? */

		} else if ht != op.GetArr() {
			ZvalPtrDtor(op)
		} else {
			ht.DelRefcount()
		}
		obj = ZendObjectsNew(ZendStandardClassDef)
		obj.SetProperties(ht)
		op.SetObject(obj)
	case types.IS_OBJECT:

	case types.IS_NULL:
		ObjectInit(op)
	case types.IS_REFERENCE:
		ZendUnwrapReference(op)
		goto try_again
	default:
		var tmp types.Zval
		types.ZVAL_COPY_VALUE(&tmp, op)
		ObjectInit(op)
		types.Z_OBJPROP_P(op).KeyAddNew(types.STR_SCALAR, &tmp)
	}
}
func _zvalGetLongFuncEx(op *types.Zval, silent types.ZendBool) ZendLong {
try_again:
	switch op.GetType() {
	case types.IS_UNDEF:
		fallthrough
	case types.IS_NULL:
		fallthrough
	case types.IS_FALSE:
		return 0
	case types.IS_TRUE:
		return 1
	case types.IS_RESOURCE:
		return types.Z_RES_HANDLE_P(op)
	case types.IS_LONG:
		return op.GetLval()
	case types.IS_DOUBLE:
		return DvalToLval(op.GetDval())
	case types.IS_STRING:
		var type_ types.ZendUchar
		var lval ZendLong
		var dval float64
		if 0 == b.Assign(&type_, IsNumericString(op.GetStr().GetStr(), &lval, &dval, b.Cond(silent != 0, 1, -1))) {
			if silent == 0 {
				faults.Error(faults.E_WARNING, "A non-numeric value encountered")
			}
			return 0
		} else if type_ == types.IS_LONG {
			return lval
		} else {

			/* Previously we used strtol here, not is_numeric_string,
			 * and strtol gives you LONG_MAX/_MIN on overflow.
			 * We use use saturating conversion to emulate strtol()'s
			 * behaviour.
			 */

			return DvalToLvalCap(dval)
		}
	case types.IS_ARRAY:
		if types.Z_ARRVAL_P(op).Len() != 0 {
			return 1
		} else {
			return 0
		}
	case types.IS_OBJECT:
		var dst types.Zval
		ConvertObjectToType(op, &dst, types.IS_LONG, ConvertToLong)
		if dst.IsLong() {
			return dst.GetLval()
		} else {
			return 1
		}
	case types.IS_REFERENCE:
		op = types.Z_REFVAL_P(op)
		goto try_again
	default:

	}
	return 0
}
func ZvalGetLongFunc(op *types.Zval) ZendLong       { return _zvalGetLongFuncEx(op, 1) }
func _zvalGetLongFuncNoisy(op *types.Zval) ZendLong { return _zvalGetLongFuncEx(op, 0) }
func ZvalGetDoubleFunc(op *types.Zval) float64 {
try_again:
	switch op.GetType() {
	case types.IS_NULL:
		fallthrough
	case types.IS_FALSE:
		return 0.0
	case types.IS_TRUE:
		return 1.0
	case types.IS_RESOURCE:
		return float64(types.Z_RES_HANDLE_P(op))
	case types.IS_LONG:
		return float64(op.GetLval())
	case types.IS_DOUBLE:
		return op.GetDval()
	case types.IS_STRING:
		return ZendStrtod(op.GetStr().GetVal(), nil)
	case types.IS_ARRAY:
		if types.Z_ARRVAL_P(op).Len() {
			return 1.0
		} else {
			return 0.0
		}
		fallthrough
	case types.IS_OBJECT:
		var dst types.Zval
		ConvertObjectToType(op, &dst, types.IS_DOUBLE, ConvertToDouble)
		if dst.IsDouble() {
			return dst.GetDval()
		} else {
			return 1.0
		}
		fallthrough
	case types.IS_REFERENCE:
		op = types.Z_REFVAL_P(op)
		goto try_again
	default:

	}
	return 0.0
}
func __zvalGetStringFunc(op *types.Zval, try types.ZendBool) *types.String {
try_again:
	switch op.GetType() {
	case types.IS_UNDEF:
		fallthrough
	case types.IS_NULL:
		fallthrough
	case types.IS_FALSE:
		return types.NewString("")
	case types.IS_TRUE:
		return types.NewString("1")
	case types.IS_RESOURCE:
		return ZendStrpprintf(0, "Resource id #"+ZEND_LONG_FMT, ZendLong(types.Z_RES_HANDLE_P(op)))
	case types.IS_LONG:
		return ZendLongToStr(op.GetLval())
	case types.IS_DOUBLE:
		return ZendStrpprintf(0, "%.*G", int(EG__().GetPrecision()), op.GetDval())
	case types.IS_ARRAY:
		faults.Error(faults.E_NOTICE, "Array to string conversion")
		if try != 0 && EG__().GetException() != nil {
			return nil
		} else {
			return types.NewString(types.STR_ARRAY_CAPITALIZED)
		}
		fallthrough
	case types.IS_OBJECT:
		var tmp types.Zval
		if types.Z_OBJ_HT_P(op).GetCastObject() != nil {
			if types.Z_OBJ_HT_P(op).GetCastObject()(op, &tmp, types.IS_STRING) == types.SUCCESS {
				return tmp.GetStr()
			}
		} else if types.Z_OBJ_HT_P(op).GetGet() != nil {
			var z *types.Zval = types.Z_OBJ_HT_P(op).GetGet()(op, &tmp)
			if z.GetType() != types.IS_OBJECT {
				var str *types.String = b.CondF(try != 0, func() *types.String { return ZvalTryGetString(z) }, func() *types.String { return ZvalGetString(z) })
				ZvalPtrDtor(z)
				return str
			}
			ZvalPtrDtor(z)
		}
		if EG__().GetException() == nil {
			faults.ThrowError(nil, "Object of class %s could not be converted to string", types.Z_OBJCE_P(op).GetName().GetVal())
		}
		if try != 0 {
			return nil
		} else {
			return types.NewString("")
		}
		fallthrough
	case types.IS_REFERENCE:
		op = types.Z_REFVAL_P(op)
		goto try_again
	case types.IS_STRING:
		return op.GetStr().Copy()
	default:

	}
	return nil
}
func ZvalGetStringFunc(op *types.Zval) *types.String    { return __zvalGetStringFunc(op, 0) }
func ZvalTryGetStringFunc(op *types.Zval) *types.String { return __zvalGetStringFunc(op, 1) }
func AddFunctionArray(result *types.Zval, op1 *types.Zval, op2 *types.Zval) {
	if result == op1 && op1.GetArr() == op2.GetArr() {

		/* $a += $a */

		return

		/* $a += $a */

	}
	if result != op1 {
		result.SetArray(types.ZendArrayDup(op1.GetArr()))
	} else {
		types.SEPARATE_ARRAY(result)
	}
	types.ZendHashMerge(result.GetArr(), op2.GetArr(), ZvalAddRef, 0)
}
func AddFunctionFast(result *types.Zval, op1 *types.Zval, op2 *types.Zval) int {
	var type_pair types.ZendUchar = TYPE_PAIR(op1.GetType(), op2.GetType())
	if type_pair == TYPE_PAIR(types.IS_LONG, types.IS_LONG) {
		FastLongAddFunction(result, op1, op2)
		return types.SUCCESS
	} else if type_pair == TYPE_PAIR(types.IS_DOUBLE, types.IS_DOUBLE) {
		result.SetDouble(op1.GetDval() + op2.GetDval())
		return types.SUCCESS
	} else if type_pair == TYPE_PAIR(types.IS_LONG, types.IS_DOUBLE) {
		result.SetDouble(float64(op1.GetLval()) + op2.GetDval())
		return types.SUCCESS
	} else if type_pair == TYPE_PAIR(types.IS_DOUBLE, types.IS_LONG) {
		result.SetDouble(op1.GetDval() + float64(op2.GetLval()))
		return types.SUCCESS
	} else if type_pair == TYPE_PAIR(types.IS_ARRAY, types.IS_ARRAY) {
		AddFunctionArray(result, op1, op2)
		return types.SUCCESS
	} else {
		return types.FAILURE
	}
}
func AddFunctionSlow(result *types.Zval, op1 *types.Zval, op2 *types.Zval) int {
	var op1_copy types.Zval
	var op2_copy types.Zval
	var converted int = 0
	for true {
		if op1.IsReference() {
			op1 = types.Z_REFVAL_P(op1)
		} else if op2.IsReference() {
			op2 = types.Z_REFVAL_P(op2)
		} else if converted == 0 {
			if op1.IsObject() && op1 == result && types.Z_OBJ_HT(*op1).GetGet() != nil && types.Z_OBJ_HT(*op1).GetSet() != nil {
				var ret int
				var rv types.Zval
				var objval *types.Zval = types.Z_OBJ_HT(*op1).GetGet()(op1, &rv)
				objval.TryAddRefcount()
				ret = AddFunction(objval, objval, op2)
				types.Z_OBJ_HT(*op1).GetSet()(op1, objval)
				ZvalPtrDtor(objval)
				return ret
			} else if op1.IsObject() && types.Z_OBJ_HT(*op1).GetDoOperation() != nil {
				if types.SUCCESS == types.Z_OBJ_HT(*op1).GetDoOperation()(ZEND_ADD, result, op1, op2) {
					return types.SUCCESS
				}
			} else if op2.IsObject() && types.Z_OBJ_HT(*op2).GetDoOperation() != nil && types.SUCCESS == types.Z_OBJ_HT(*op2).GetDoOperation()(ZEND_ADD, result, op1, op2) {
				return types.SUCCESS
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
				return types.FAILURE
			}
			converted = 1
		} else {
			if result != op1 {
				result.SetUndef()
			}
			faults.ThrowError(nil, "Unsupported operand types")
			return types.FAILURE
		}
		if AddFunctionFast(result, op1, op2) == types.SUCCESS {
			return types.SUCCESS
		}
	}
}
func AddFunction(result *types.Zval, op1 *types.Zval, op2 *types.Zval) int {
	if AddFunctionFast(result, op1, op2) == types.SUCCESS {
		return types.SUCCESS
	} else {
		return AddFunctionSlow(result, op1, op2)
	}
}
func SubFunctionFast(result *types.Zval, op1 *types.Zval, op2 *types.Zval) int {
	var type_pair types.ZendUchar = TYPE_PAIR(op1.GetType(), op2.GetType())
	if type_pair == TYPE_PAIR(types.IS_LONG, types.IS_LONG) {
		FastLongSubFunction(result, op1, op2)
		return types.SUCCESS
	} else if type_pair == TYPE_PAIR(types.IS_DOUBLE, types.IS_DOUBLE) {
		result.SetDouble(op1.GetDval() - op2.GetDval())
		return types.SUCCESS
	} else if type_pair == TYPE_PAIR(types.IS_LONG, types.IS_DOUBLE) {
		result.SetDouble(float64(op1.GetLval()) - op2.GetDval())
		return types.SUCCESS
	} else if type_pair == TYPE_PAIR(types.IS_DOUBLE, types.IS_LONG) {
		result.SetDouble(op1.GetDval() - float64(op2.GetLval()))
		return types.SUCCESS
	} else {
		return types.FAILURE
	}
}
func SubFunctionSlow(result *types.Zval, op1 *types.Zval, op2 *types.Zval) int {
	var op1_copy types.Zval
	var op2_copy types.Zval
	var converted int = 0
	for true {
		if op1.IsReference() {
			op1 = types.Z_REFVAL_P(op1)
		} else if op2.IsReference() {
			op2 = types.Z_REFVAL_P(op2)
		} else if converted == 0 {
			if op1.IsObject() && op1 == result && types.Z_OBJ_HT(*op1).GetGet() != nil && types.Z_OBJ_HT(*op1).GetSet() != nil {
				var ret int
				var rv types.Zval
				var objval *types.Zval = types.Z_OBJ_HT(*op1).GetGet()(op1, &rv)
				objval.TryAddRefcount()
				ret = SubFunction(objval, objval, op2)
				types.Z_OBJ_HT(*op1).GetSet()(op1, objval)
				ZvalPtrDtor(objval)
				return ret
			} else if op1.IsObject() && types.Z_OBJ_HT(*op1).GetDoOperation() != nil {
				if types.SUCCESS == types.Z_OBJ_HT(*op1).GetDoOperation()(ZEND_SUB, result, op1, op2) {
					return types.SUCCESS
				}
			} else if op2.IsObject() && types.Z_OBJ_HT(*op2).GetDoOperation() != nil && types.SUCCESS == types.Z_OBJ_HT(*op2).GetDoOperation()(ZEND_SUB, result, op1, op2) {
				return types.SUCCESS
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
				return types.FAILURE
			}
			converted = 1
		} else {
			if result != op1 {
				result.SetUndef()
			}
			faults.ThrowError(nil, "Unsupported operand types")
			return types.FAILURE
		}
		if SubFunctionFast(result, op1, op2) == types.SUCCESS {
			return types.SUCCESS
		}
	}
}
func SubFunction(result *types.Zval, op1 *types.Zval, op2 *types.Zval) int {
	if SubFunctionFast(result, op1, op2) == types.SUCCESS {
		return types.SUCCESS
	} else {
		return SubFunctionSlow(result, op1, op2)
	}
}
func MulFunction(result *types.Zval, op1 *types.Zval, op2 *types.Zval) int {
	var op1_copy types.Zval
	var op2_copy types.Zval
	var converted int = 0
	for true {
		var type_pair types.ZendUchar = TYPE_PAIR(op1.GetType(), op2.GetType())
		if type_pair == TYPE_PAIR(types.IS_LONG, types.IS_LONG) {
			var overflow ZendLong
			ZEND_SIGNED_MULTIPLY_LONG(op1.GetLval(), op2.GetLval(), result.GetLval(), result.GetDval(), overflow)
			if overflow != 0 {
				result.SetTypeInfo(types.IS_DOUBLE)
			} else {
				result.SetTypeInfo(types.IS_LONG)
			}
			return types.SUCCESS
		} else if type_pair == TYPE_PAIR(types.IS_DOUBLE, types.IS_DOUBLE) {
			result.SetDouble(op1.GetDval() * op2.GetDval())
			return types.SUCCESS
		} else if type_pair == TYPE_PAIR(types.IS_LONG, types.IS_DOUBLE) {
			result.SetDouble(float64(op1.GetLval()) * op2.GetDval())
			return types.SUCCESS
		} else if type_pair == TYPE_PAIR(types.IS_DOUBLE, types.IS_LONG) {
			result.SetDouble(op1.GetDval() * float64(op2.GetLval()))
			return types.SUCCESS
		} else {
			if op1.IsReference() {
				op1 = types.Z_REFVAL_P(op1)
			} else if op2.IsReference() {
				op2 = types.Z_REFVAL_P(op2)
			} else if converted == 0 {
				if op1.IsObject() && op1 == result && types.Z_OBJ_HT(*op1).GetGet() != nil && types.Z_OBJ_HT(*op1).GetSet() != nil {
					var ret int
					var rv types.Zval
					var objval *types.Zval = types.Z_OBJ_HT(*op1).GetGet()(op1, &rv)
					objval.TryAddRefcount()
					ret = MulFunction(objval, objval, op2)
					types.Z_OBJ_HT(*op1).GetSet()(op1, objval)
					ZvalPtrDtor(objval)
					return ret
				} else if op1.IsObject() && types.Z_OBJ_HT(*op1).GetDoOperation() != nil {
					if types.SUCCESS == types.Z_OBJ_HT(*op1).GetDoOperation()(ZEND_MUL, result, op1, op2) {
						return types.SUCCESS
					}
				} else if op2.IsObject() && types.Z_OBJ_HT(*op2).GetDoOperation() != nil && types.SUCCESS == types.Z_OBJ_HT(*op2).GetDoOperation()(ZEND_MUL, result, op1, op2) {
					return types.SUCCESS
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
					return types.FAILURE
				}
				converted = 1
			} else {
				if result != op1 {
					result.SetUndef()
				}
				faults.ThrowError(nil, "Unsupported operand types")
				return types.FAILURE
			}
		}
	}
}
func PowFunction(result *types.Zval, op1 *types.Zval, op2 *types.Zval) int {
	var op1_copy types.Zval
	var op2_copy types.Zval
	var converted int = 0
	for true {
		var type_pair types.ZendUchar = TYPE_PAIR(op1.GetType(), op2.GetType())
		if type_pair == TYPE_PAIR(types.IS_LONG, types.IS_LONG) {
			if op2.GetLval() >= 0 {
				var l1 ZendLong = 1
				var l2 ZendLong = op1.GetLval()
				var i ZendLong = op2.GetLval()
				if i == 0 {
					result.SetLong(1)
					return types.SUCCESS
				} else if l2 == 0 {
					result.SetLong(0)
					return types.SUCCESS
				}
				for i >= 1 {
					var overflow ZendLong
					var dval float64 = 0.0
					if i%2 != 0 {
						i--
						ZEND_SIGNED_MULTIPLY_LONG(l1, l2, l1, dval, overflow)
						if overflow != 0 {
							result.SetDouble(dval * pow(l2, i))
							return types.SUCCESS
						}
					} else {
						i /= 2
						ZEND_SIGNED_MULTIPLY_LONG(l2, l2, l2, dval, overflow)
						if overflow != 0 {
							result.SetDouble(float64(l1 * pow(dval, i)))
							return types.SUCCESS
						}
					}
				}

				/* i == 0 */

				result.SetLong(l1)

				/* i == 0 */

			} else {
				result.SetDouble(pow(float64(op1.GetLval()), float64(op2.GetLval())))
			}
			return types.SUCCESS
		} else if type_pair == TYPE_PAIR(types.IS_DOUBLE, types.IS_DOUBLE) {
			result.SetDouble(pow(op1.GetDval(), op2.GetDval()))
			return types.SUCCESS
		} else if type_pair == TYPE_PAIR(types.IS_LONG, types.IS_DOUBLE) {
			result.SetDouble(pow(float64(op1.GetLval()), op2.GetDval()))
			return types.SUCCESS
		} else if type_pair == TYPE_PAIR(types.IS_DOUBLE, types.IS_LONG) {
			result.SetDouble(pow(op1.GetDval(), float64(op2.GetLval())))
			return types.SUCCESS
		} else {
			if op1.IsReference() {
				op1 = types.Z_REFVAL_P(op1)
			} else if op2.IsReference() {
				op2 = types.Z_REFVAL_P(op2)
			} else if converted == 0 {
				if op1.IsObject() && op1 == result && types.Z_OBJ_HT(*op1).GetGet() != nil && types.Z_OBJ_HT(*op1).GetSet() != nil {
					var ret int
					var rv types.Zval
					var objval *types.Zval = types.Z_OBJ_HT(*op1).GetGet()(op1, &rv)
					objval.TryAddRefcount()
					ret = PowFunction(objval, objval, op2)
					types.Z_OBJ_HT(*op1).GetSet()(op1, objval)
					ZvalPtrDtor(objval)
					return ret
				} else if op1.IsObject() && types.Z_OBJ_HT(*op1).GetDoOperation() != nil {
					if types.SUCCESS == types.Z_OBJ_HT(*op1).GetDoOperation()(ZEND_POW, result, op1, op2) {
						return types.SUCCESS
					}
				} else if op2.IsObject() && types.Z_OBJ_HT(*op2).GetDoOperation() != nil && types.SUCCESS == types.Z_OBJ_HT(*op2).GetDoOperation()(ZEND_POW, result, op1, op2) {
					return types.SUCCESS
				}
				if op1 != op2 {
					if op1.IsArray() {
						if op1 == result {
							ZvalPtrDtor(result)
						}
						result.SetLong(0)
						return types.SUCCESS
					} else {
						op1 = ZendiConvertScalarToNumber(op1, &op1_copy, result, 0)
					}
					if op2.IsArray() {
						if op1 == result {
							ZvalPtrDtor(result)
						}
						result.SetLong(1)
						return types.SUCCESS
					} else {
						op2 = ZendiConvertScalarToNumber(op2, &op2_copy, result, 0)
					}
				} else {
					if op1.IsArray() {
						if op1 == result {
							ZvalPtrDtor(result)
						}
						result.SetLong(0)
						return types.SUCCESS
					} else {
						op1 = ZendiConvertScalarToNumber(op1, &op1_copy, result, 0)
					}
					op2 = op1
				}
				if EG__().GetException() != nil {
					if result != op1 {
						result.SetUndef()
					}
					return types.FAILURE
				}
				converted = 1
			} else {
				if result != op1 {
					result.SetUndef()
				}
				faults.ThrowError(nil, "Unsupported operand types")
				return types.FAILURE
			}
		}
	}
}
func DivFunction(result *types.Zval, op1 *types.Zval, op2 *types.Zval) int {
	var op1_copy types.Zval
	var op2_copy types.Zval
	var converted int = 0
	for true {
		var type_pair types.ZendUchar = TYPE_PAIR(op1.GetType(), op2.GetType())
		if type_pair == TYPE_PAIR(types.IS_LONG, types.IS_LONG) {
			if op2.GetLval() == 0 {
				faults.Error(faults.E_WARNING, "Division by zero")
				result.SetDouble(float64(op1.GetLval() / float64(op2.GetLval())))
				return types.SUCCESS
			} else if op2.GetLval() == -1 && op1.GetLval() == ZEND_LONG_MIN {

				/* Prevent overflow error/crash */

				result.SetDouble(float64(ZEND_LONG_MIN / -1))
				return types.SUCCESS
			}
			if op1.GetLval()%op2.GetLval() == 0 {
				result.SetLong(op1.GetLval() / op2.GetLval())
			} else {
				result.SetDouble(float64(op1.GetLval()) / op2.GetLval())
			}
			return types.SUCCESS
		} else if type_pair == TYPE_PAIR(types.IS_DOUBLE, types.IS_DOUBLE) {
			if op2.GetDval() == 0 {
				faults.Error(faults.E_WARNING, "Division by zero")
			}
			result.SetDouble(op1.GetDval() / op2.GetDval())
			return types.SUCCESS
		} else if type_pair == TYPE_PAIR(types.IS_DOUBLE, types.IS_LONG) {
			if op2.GetLval() == 0 {
				faults.Error(faults.E_WARNING, "Division by zero")
			}
			result.SetDouble(op1.GetDval() / float64(op2.GetLval()))
			return types.SUCCESS
		} else if type_pair == TYPE_PAIR(types.IS_LONG, types.IS_DOUBLE) {
			if op2.GetDval() == 0 {
				faults.Error(faults.E_WARNING, "Division by zero")
			}
			result.SetDouble(float64(op1.GetLval() / op2.GetDval()))
			return types.SUCCESS
		} else {
			if op1.IsReference() {
				op1 = types.Z_REFVAL_P(op1)
			} else if op2.IsReference() {
				op2 = types.Z_REFVAL_P(op2)
			} else if converted == 0 {
				if op1.IsObject() && op1 == result && types.Z_OBJ_HT(*op1).GetGet() != nil && types.Z_OBJ_HT(*op1).GetSet() != nil {
					var ret int
					var rv types.Zval
					var objval *types.Zval = types.Z_OBJ_HT(*op1).GetGet()(op1, &rv)
					objval.TryAddRefcount()
					ret = DivFunction(objval, objval, op2)
					types.Z_OBJ_HT(*op1).GetSet()(op1, objval)
					ZvalPtrDtor(objval)
					return ret
				} else if op1.IsObject() && types.Z_OBJ_HT(*op1).GetDoOperation() != nil {
					if types.SUCCESS == types.Z_OBJ_HT(*op1).GetDoOperation()(ZEND_DIV, result, op1, op2) {
						return types.SUCCESS
					}
				} else if op2.IsObject() && types.Z_OBJ_HT(*op2).GetDoOperation() != nil && types.SUCCESS == types.Z_OBJ_HT(*op2).GetDoOperation()(ZEND_DIV, result, op1, op2) {
					return types.SUCCESS
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
					return types.FAILURE
				}
				converted = 1
			} else {
				if result != op1 {
					result.SetUndef()
				}
				faults.ThrowError(nil, "Unsupported operand types")
				return types.FAILURE
			}
		}
	}
}
func ModFunction(result *types.Zval, op1 *types.Zval, op2 *types.Zval) int {
	var op1_lval ZendLong
	var op2_lval ZendLong
	for {
		if op1.GetType() != types.IS_LONG {
			if op1.IsReference() {
				op1 = types.Z_REFVAL_P(op1)
				if op1.IsLong() {
					op1_lval = op1.GetLval()
					break
				}
			}
			if op1.IsObject() && op1 == result && types.Z_OBJ_HT(*op1).GetGet() != nil && types.Z_OBJ_HT(*op1).GetSet() != nil {
				var ret int
				var rv types.Zval
				var objval *types.Zval = types.Z_OBJ_HT(*op1).GetGet()(op1, &rv)
				objval.TryAddRefcount()
				ret = ModFunction(objval, objval, op2)
				types.Z_OBJ_HT(*op1).GetSet()(op1, objval)
				ZvalPtrDtor(objval)
				return ret
			} else if op1.IsObject() && types.Z_OBJ_HT(*op1).GetDoOperation() != nil {
				if types.SUCCESS == types.Z_OBJ_HT(*op1).GetDoOperation()(ZEND_MOD, result, op1, op2) {
					return types.SUCCESS
				}
			}
			op1_lval = _zvalGetLongFuncNoisy(op1)
			if EG__().GetException() != nil {
				if result != op1 {
					result.SetUndef()
				}
				return types.FAILURE
			}
		} else {
			op1_lval = op1.GetLval()
		}
		break
	}
	for {
		if op2.GetType() != types.IS_LONG {
			if op2.IsReference() {
				op2 = types.Z_REFVAL_P(op2)
				if op2.IsLong() {
					op2_lval = op2.GetLval()
					break
				}
			}
			if op2.IsObject() && types.Z_OBJ_HT(*op2).GetDoOperation() != nil && types.SUCCESS == types.Z_OBJ_HT(*op2).GetDoOperation()(ZEND_MOD, result, op1, op2) {
				return types.SUCCESS
			}
			op2_lval = _zvalGetLongFuncNoisy(op2)
			if EG__().GetException() != nil {
				if result != op1 {
					result.SetUndef()
				}
				return types.FAILURE
			}
		} else {
			op2_lval = op2.GetLval()
		}
		break
	}
	if op2_lval == 0 {

		/* modulus by zero */

		if CurrEX() != nil && CG__().GetInCompilation() == 0 {
			faults.ThrowExceptionEx(faults.ZendCeDivisionByZeroError, 0, "Modulo by zero")
		} else {
			faults.ErrorNoreturn(faults.E_ERROR, "Modulo by zero")
		}
		if op1 != result {
			result.SetUndef()
		}
		return types.FAILURE
	}
	if op1 == result {
		ZvalPtrDtor(result)
	}
	if op2_lval == -1 {

		/* Prevent overflow error/crash if op1==LONG_MIN */

		result.SetLong(0)
		return types.SUCCESS
	}
	result.SetLong(op1_lval % op2_lval)
	return types.SUCCESS
}
func BooleanXorFunction(result *types.Zval, op1 *types.Zval, op2 *types.Zval) int {
	var op1_val int
	var op2_val int
	for {
		if op1.IsFalse() {
			op1_val = 0
		} else if op1.IsTrue() {
			op1_val = 1
		} else {
			if op1.IsReference() {
				op1 = types.Z_REFVAL_P(op1)
				if op1.IsFalse() {
					op1_val = 0
					break
				} else if op1.IsTrue() {
					op1_val = 1
					break
				}
			}
			if op1.IsObject() && op1 == result && types.Z_OBJ_HT(*op1).GetGet() != nil && types.Z_OBJ_HT(*op1).GetSet() != nil {
				var ret int
				var rv types.Zval
				var objval *types.Zval = types.Z_OBJ_HT(*op1).GetGet()(op1, &rv)
				objval.TryAddRefcount()
				ret = BooleanXorFunction(objval, objval, op2)
				types.Z_OBJ_HT(*op1).GetSet()(op1, objval)
				ZvalPtrDtor(objval)
				return ret
			} else if op1.IsObject() && types.Z_OBJ_HT(*op1).GetDoOperation() != nil {
				if types.SUCCESS == types.Z_OBJ_HT(*op1).GetDoOperation()(ZEND_BOOL_XOR, result, op1, op2) {
					return types.SUCCESS
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
				op2 = types.Z_REFVAL_P(op2)
				if op2.IsFalse() {
					op2_val = 0
					break
				} else if op2.IsTrue() {
					op2_val = 1
					break
				}
			}
			if op2.IsObject() && types.Z_OBJ_HT(*op2).GetDoOperation() != nil && types.SUCCESS == types.Z_OBJ_HT(*op2).GetDoOperation()(ZEND_BOOL_XOR, result, op1, op2) {
				return types.SUCCESS
			}
			op2_val = ZvalIsTrue(op2)
		}
		break
	}
	types.ZVAL_BOOL(result, (op1_val^op2_val) != 0)
	return types.SUCCESS
}
func BooleanNotFunction(result *types.Zval, op1 *types.Zval) int {
	if op1.GetType() < types.IS_TRUE {
		result.SetTrue()
	} else if op1.IsTrue() {
		result.SetFalse()
	} else {
		if op1.IsReference() {
			op1 = types.Z_REFVAL_P(op1)
			if op1.GetType() < types.IS_TRUE {
				result.SetTrue()
				return types.SUCCESS
			} else if op1.IsTrue() {
				result.SetFalse()
				return types.SUCCESS
			}
		}
		if op1.IsObject() && types.Z_OBJ_HT(*op1).GetDoOperation() != nil && types.SUCCESS == types.Z_OBJ_HT(*op1).GetDoOperation()(ZEND_BOOL_NOT, result, op1, nil) {
			return types.SUCCESS
		}
		types.ZVAL_BOOL(result, ZvalIsTrue(op1) == 0)
	}
	return types.SUCCESS
}
func BitwiseNotFunction(result *types.Zval, op1 *types.Zval) int {
try_again:
	switch op1.GetType() {
	case types.IS_LONG:
		result.SetLong(^(op1.GetLval()))
		return types.SUCCESS
	case types.IS_DOUBLE:
		result.SetLong(^(DvalToLval(op1.GetDval())))
		return types.SUCCESS
	case types.IS_STRING:
		var i int
		if op1.GetStr().GetLen() == 1 {
			var not types.ZendUchar = types.ZendUchar(^((*types.Z_STRVAL_P)(op1)))
			result.SetStringVal(string(not))
		} else {
			result.SetString(types.ZendStringAlloc(op1.GetStr().GetLen(), 0))
			for i = 0; i < op1.GetStr().GetLen(); i++ {
				result.GetStr().GetVal()[i] = ^(op1.GetStr().GetVal()[i])
			}
			result.GetStr().GetVal()[i] = 0
		}
		return types.SUCCESS
	case types.IS_REFERENCE:
		op1 = types.Z_REFVAL_P(op1)
		goto try_again
	default:
		if op1.IsObject() && types.Z_OBJ_HT(*op1).GetDoOperation() != nil && types.SUCCESS == types.Z_OBJ_HT(*op1).GetDoOperation()(ZEND_BW_NOT, result, op1, nil) {
			return types.SUCCESS
		}
		if result != op1 {
			result.SetUndef()
		}
		faults.ThrowError(nil, "Unsupported operand types")
		return types.FAILURE
	}
}
func BitwiseOrFunction(result *types.Zval, op1 *types.Zval, op2 *types.Zval) int {
	var op1_lval ZendLong
	var op2_lval ZendLong
	if op1.IsLong() && op2.IsLong() {
		result.SetLong(op1.GetLval() | op2.GetLval())
		return types.SUCCESS
	}
	op1 = types.ZVAL_DEREF(op1)
	op2 = types.ZVAL_DEREF(op2)
	if op1.IsString() && op2.IsString() {
		var longer *types.Zval
		var shorter *types.Zval
		var str *types.String
		var i int
		if op1.GetStr().GetLen() >= op2.GetStr().GetLen() {
			if op1.GetStr().GetLen() == op2.GetStr().GetLen() && op1.GetStr().GetLen() == 1 {
				var or types.ZendUchar = zend_uchar((*types.Z_STRVAL_P)(op1) | (*types.Z_STRVAL_P)(op2))
				if result == op1 {

				}
				result.SetStringVal(string(or))
				return types.SUCCESS
			}
			longer = op1
			shorter = op2
		} else {
			longer = op2
			shorter = op1
		}
		str = types.ZendStringAlloc(longer.GetStr().GetLen(), 0)
		for i = 0; i < shorter.GetStr().GetLen(); i++ {
			str.GetVal()[i] = longer.GetStr().GetVal()[i] | shorter.GetStr().GetVal()[i]
		}
		memcpy(str.GetVal()+i, longer.GetStr().GetVal()+i, longer.GetStr().GetLen()-i+1)
		if result == op1 {

		}
		result.SetString(str)
		return types.SUCCESS
	}
	if op1.GetType() != types.IS_LONG {
		if op1.IsObject() && op1 == result && types.Z_OBJ_HT(*op1).GetGet() != nil && types.Z_OBJ_HT(*op1).GetSet() != nil {
			var ret int
			var rv types.Zval
			var objval *types.Zval = types.Z_OBJ_HT(*op1).GetGet()(op1, &rv)
			objval.TryAddRefcount()
			ret = BitwiseOrFunction(objval, objval, op2)
			types.Z_OBJ_HT(*op1).GetSet()(op1, objval)
			ZvalPtrDtor(objval)
			return ret
		} else if op1.IsObject() && types.Z_OBJ_HT(*op1).GetDoOperation() != nil {
			if types.SUCCESS == types.Z_OBJ_HT(*op1).GetDoOperation()(ZEND_BW_OR, result, op1, op2) {
				return types.SUCCESS
			}
		}
		op1_lval = _zvalGetLongFuncNoisy(op1)
		if EG__().GetException() != nil {
			if result != op1 {
				result.SetUndef()
			}
			return types.FAILURE
		}
	} else {
		op1_lval = op1.GetLval()
	}
	if op2.GetType() != types.IS_LONG {
		if op2.IsObject() && types.Z_OBJ_HT(*op2).GetDoOperation() != nil && types.SUCCESS == types.Z_OBJ_HT(*op2).GetDoOperation()(ZEND_BW_OR, result, op1, op2) {
			return types.SUCCESS
		}
		op2_lval = _zvalGetLongFuncNoisy(op2)
		if EG__().GetException() != nil {
			if result != op1 {
				result.SetUndef()
			}
			return types.FAILURE
		}
	} else {
		op2_lval = op2.GetLval()
	}
	if op1 == result {
		ZvalPtrDtor(result)
	}
	result.SetLong(op1_lval | op2_lval)
	return types.SUCCESS
}
func BitwiseAndFunction(result *types.Zval, op1 *types.Zval, op2 *types.Zval) int {
	var op1_lval ZendLong
	var op2_lval ZendLong
	if op1.IsLong() && op2.IsLong() {
		result.SetLong(op1.GetLval() & op2.GetLval())
		return types.SUCCESS
	}
	op1 = types.ZVAL_DEREF(op1)
	op2 = types.ZVAL_DEREF(op2)
	if op1.IsString() && op2.IsString() {
		var longer *types.Zval
		var shorter *types.Zval
		var str *types.String
		var i int
		if op1.GetStr().GetLen() >= op2.GetStr().GetLen() {
			if op1.GetStr().GetLen() == op2.GetStr().GetLen() && op1.GetStr().GetLen() == 1 {
				var and types.ZendUchar = byte((*types.Z_STRVAL_P)(op1) & (*types.Z_STRVAL_P)(op2))
				if result == op1 {

				}
				result.SetStringVal(string(and))
				return types.SUCCESS
			}
			longer = op1
			shorter = op2
		} else {
			longer = op2
			shorter = op1
		}
		str = types.ZendStringAlloc(shorter.GetStr().GetLen(), 0)
		for i = 0; i < shorter.GetStr().GetLen(); i++ {
			str.GetVal()[i] = shorter.GetStr().GetVal()[i] & longer.GetStr().GetVal()[i]
		}
		str.GetVal()[i] = 0
		if result == op1 {

		}
		result.SetString(str)
		return types.SUCCESS
	}
	if op1.GetType() != types.IS_LONG {
		if op1.IsObject() && op1 == result && types.Z_OBJ_HT(*op1).GetGet() != nil && types.Z_OBJ_HT(*op1).GetSet() != nil {
			var ret int
			var rv types.Zval
			var objval *types.Zval = types.Z_OBJ_HT(*op1).GetGet()(op1, &rv)
			objval.TryAddRefcount()
			ret = BitwiseAndFunction(objval, objval, op2)
			types.Z_OBJ_HT(*op1).GetSet()(op1, objval)
			ZvalPtrDtor(objval)
			return ret
		} else if op1.IsObject() && types.Z_OBJ_HT(*op1).GetDoOperation() != nil {
			if types.SUCCESS == types.Z_OBJ_HT(*op1).GetDoOperation()(ZEND_BW_AND, result, op1, op2) {
				return types.SUCCESS
			}
		}
		op1_lval = _zvalGetLongFuncNoisy(op1)
		if EG__().GetException() != nil {
			if result != op1 {
				result.SetUndef()
			}
			return types.FAILURE
		}
	} else {
		op1_lval = op1.GetLval()
	}
	if op2.GetType() != types.IS_LONG {
		if op2.IsObject() && types.Z_OBJ_HT(*op2).GetDoOperation() != nil && types.SUCCESS == types.Z_OBJ_HT(*op2).GetDoOperation()(ZEND_BW_AND, result, op1, op2) {
			return types.SUCCESS
		}
		op2_lval = _zvalGetLongFuncNoisy(op2)
		if EG__().GetException() != nil {
			if result != op1 {
				result.SetUndef()
			}
			return types.FAILURE
		}
	} else {
		op2_lval = op2.GetLval()
	}
	if op1 == result {
		ZvalPtrDtor(result)
	}
	result.SetLong(op1_lval & op2_lval)
	return types.SUCCESS
}
func BitwiseXorFunction(result *types.Zval, op1 *types.Zval, op2 *types.Zval) int {
	var op1_lval ZendLong
	var op2_lval ZendLong
	if op1.IsLong() && op2.IsLong() {
		result.SetLong(op1.GetLval() ^ op2.GetLval())
		return types.SUCCESS
	}
	op1 = types.ZVAL_DEREF(op1)
	op2 = types.ZVAL_DEREF(op2)
	if op1.IsString() && op2.IsString() {
		var longer *types.Zval
		var shorter *types.Zval
		var str *types.String
		var i int
		if op1.GetStr().GetLen() >= op2.GetStr().GetLen() {
			if op1.GetStr().GetLen() == op2.GetStr().GetLen() && op1.GetStr().GetLen() == 1 {
				var xor types.ZendUchar = zend_uchar((*types.Z_STRVAL_P)(op1) ^ (*types.Z_STRVAL_P)(op2))
				if result == op1 {

				}
				result.SetStringVal(string(xor))
				return types.SUCCESS
			}
			longer = op1
			shorter = op2
		} else {
			longer = op2
			shorter = op1
		}
		str = types.ZendStringAlloc(shorter.GetStr().GetLen(), 0)
		for i = 0; i < shorter.GetStr().GetLen(); i++ {
			str.GetVal()[i] = shorter.GetStr().GetVal()[i] ^ longer.GetStr().GetVal()[i]
		}
		str.GetVal()[i] = 0
		if result == op1 {

		}
		result.SetString(str)
		return types.SUCCESS
	}
	if op1.GetType() != types.IS_LONG {
		if op1.IsObject() && op1 == result && types.Z_OBJ_HT(*op1).GetGet() != nil && types.Z_OBJ_HT(*op1).GetSet() != nil {
			var ret int
			var rv types.Zval
			var objval *types.Zval = types.Z_OBJ_HT(*op1).GetGet()(op1, &rv)
			objval.TryAddRefcount()
			ret = BitwiseXorFunction(objval, objval, op2)
			types.Z_OBJ_HT(*op1).GetSet()(op1, objval)
			ZvalPtrDtor(objval)
			return ret
		} else if op1.IsObject() && types.Z_OBJ_HT(*op1).GetDoOperation() != nil {
			if types.SUCCESS == types.Z_OBJ_HT(*op1).GetDoOperation()(ZEND_BW_XOR, result, op1, op2) {
				return types.SUCCESS
			}
		}
		op1_lval = _zvalGetLongFuncNoisy(op1)
		if EG__().GetException() != nil {
			if result != op1 {
				result.SetUndef()
			}
			return types.FAILURE
		}
	} else {
		op1_lval = op1.GetLval()
	}
	if op2.GetType() != types.IS_LONG {
		if op2.IsObject() && types.Z_OBJ_HT(*op2).GetDoOperation() != nil && types.SUCCESS == types.Z_OBJ_HT(*op2).GetDoOperation()(ZEND_BW_XOR, result, op1, op2) {
			return types.SUCCESS
		}
		op2_lval = _zvalGetLongFuncNoisy(op2)
		if EG__().GetException() != nil {
			if result != op1 {
				result.SetUndef()
			}
			return types.FAILURE
		}
	} else {
		op2_lval = op2.GetLval()
	}
	if op1 == result {
		ZvalPtrDtor(result)
	}
	result.SetLong(op1_lval ^ op2_lval)
	return types.SUCCESS
}
func ShiftLeftFunction(result *types.Zval, op1 *types.Zval, op2 *types.Zval) int {
	var op1_lval ZendLong
	var op2_lval ZendLong
	for {
		if op1.GetType() != types.IS_LONG {
			if op1.IsReference() {
				op1 = types.Z_REFVAL_P(op1)
				if op1.IsLong() {
					op1_lval = op1.GetLval()
					break
				}
			}
			if op1.IsObject() && op1 == result && types.Z_OBJ_HT(*op1).GetGet() != nil && types.Z_OBJ_HT(*op1).GetSet() != nil {
				var ret int
				var rv types.Zval
				var objval *types.Zval = types.Z_OBJ_HT(*op1).GetGet()(op1, &rv)
				objval.TryAddRefcount()
				ret = ShiftLeftFunction(objval, objval, op2)
				types.Z_OBJ_HT(*op1).GetSet()(op1, objval)
				ZvalPtrDtor(objval)
				return ret
			} else if op1.IsObject() && types.Z_OBJ_HT(*op1).GetDoOperation() != nil {
				if types.SUCCESS == types.Z_OBJ_HT(*op1).GetDoOperation()(ZEND_SL, result, op1, op2) {
					return types.SUCCESS
				}
			}
			op1_lval = _zvalGetLongFuncNoisy(op1)
			if EG__().GetException() != nil {
				if result != op1 {
					result.SetUndef()
				}
				return types.FAILURE
			}
		} else {
			op1_lval = op1.GetLval()
		}
		break
	}
	for {
		if op2.GetType() != types.IS_LONG {
			if op2.IsReference() {
				op2 = types.Z_REFVAL_P(op2)
				if op2.IsLong() {
					op2_lval = op2.GetLval()
					break
				}
			}
			if op2.IsObject() && types.Z_OBJ_HT(*op2).GetDoOperation() != nil && types.SUCCESS == types.Z_OBJ_HT(*op2).GetDoOperation()(ZEND_SL, result, op1, op2) {
				return types.SUCCESS
			}
			op2_lval = _zvalGetLongFuncNoisy(op2)
			if EG__().GetException() != nil {
				if result != op1 {
					result.SetUndef()
				}
				return types.FAILURE
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
			return types.SUCCESS
		} else {
			if CurrEX() != nil && CG__().GetInCompilation() == 0 {
				faults.ThrowExceptionEx(faults.ZendCeArithmeticError, 0, "Bit shift by negative number")
			} else {
				faults.ErrorNoreturn(faults.E_ERROR, "Bit shift by negative number")
			}
			if op1 != result {
				result.SetUndef()
			}
			return types.FAILURE
		}
	}
	if op1 == result {
		ZvalPtrDtor(result)
	}

	/* Perform shift on unsigned numbers to get well-defined wrap behavior. */

	result.SetLong(zend_long(ZendUlong(op1_lval << op2_lval)))
	return types.SUCCESS
}
func ShiftRightFunction(result *types.Zval, op1 *types.Zval, op2 *types.Zval) int {
	var op1_lval ZendLong
	var op2_lval ZendLong
	for {
		if op1.GetType() != types.IS_LONG {
			if op1.IsReference() {
				op1 = types.Z_REFVAL_P(op1)
				if op1.IsLong() {
					op1_lval = op1.GetLval()
					break
				}
			}
			if op1.IsObject() && op1 == result && types.Z_OBJ_HT(*op1).GetGet() != nil && types.Z_OBJ_HT(*op1).GetSet() != nil {
				var ret int
				var rv types.Zval
				var objval *types.Zval = types.Z_OBJ_HT(*op1).GetGet()(op1, &rv)
				objval.TryAddRefcount()
				ret = ShiftRightFunction(objval, objval, op2)
				types.Z_OBJ_HT(*op1).GetSet()(op1, objval)
				ZvalPtrDtor(objval)
				return ret
			} else if op1.IsObject() && types.Z_OBJ_HT(*op1).GetDoOperation() != nil {
				if types.SUCCESS == types.Z_OBJ_HT(*op1).GetDoOperation()(ZEND_SR, result, op1, op2) {
					return types.SUCCESS
				}
			}
			op1_lval = _zvalGetLongFuncNoisy(op1)
			if EG__().GetException() != nil {
				if result != op1 {
					result.SetUndef()
				}
				return types.FAILURE
			}
		} else {
			op1_lval = op1.GetLval()
		}
		break
	}
	for {
		if op2.GetType() != types.IS_LONG {
			if op2.IsReference() {
				op2 = types.Z_REFVAL_P(op2)
				if op2.IsLong() {
					op2_lval = op2.GetLval()
					break
				}
			}
			if op2.IsObject() && types.Z_OBJ_HT(*op2).GetDoOperation() != nil && types.SUCCESS == types.Z_OBJ_HT(*op2).GetDoOperation()(ZEND_SR, result, op1, op2) {
				return types.SUCCESS
			}
			op2_lval = _zvalGetLongFuncNoisy(op2)
			if EG__().GetException() != nil {
				if result != op1 {
					result.SetUndef()
				}
				return types.FAILURE
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
			return types.SUCCESS
		} else {
			if CurrEX() != nil && CG__().GetInCompilation() == 0 {
				faults.ThrowExceptionEx(faults.ZendCeArithmeticError, 0, "Bit shift by negative number")
			} else {
				faults.ErrorNoreturn(faults.E_ERROR, "Bit shift by negative number")
			}
			if op1 != result {
				result.SetUndef()
			}
			return types.FAILURE
		}
	}
	if op1 == result {
		ZvalPtrDtor(result)
	}
	result.SetLong(op1_lval >> op2_lval)
	return types.SUCCESS
}
func ConcatFunction(result *types.Zval, op1 *types.Zval, op2 *types.Zval) int {
	var orig_op1 *types.Zval = op1
	var op1_copy types.Zval
	var op2_copy types.Zval
	op1_copy.SetUndef()
	op2_copy.SetUndef()
	for {
		if op1.GetType() != types.IS_STRING {
			if op1.IsReference() {
				op1 = types.Z_REFVAL_P(op1)
				if op1.IsString() {
					break
				}
			}
			if op1.IsObject() && op1 == result && types.Z_OBJ_HT(*op1).GetGet() != nil && types.Z_OBJ_HT(*op1).GetSet() != nil {
				var ret int
				var rv types.Zval
				var objval *types.Zval = types.Z_OBJ_HT(*op1).GetGet()(op1, &rv)
				objval.TryAddRefcount()
				ret = ConcatFunction(objval, objval, op2)
				types.Z_OBJ_HT(*op1).GetSet()(op1, objval)
				ZvalPtrDtor(objval)
				return ret
			} else if op1.IsObject() && types.Z_OBJ_HT(*op1).GetDoOperation() != nil {
				if types.SUCCESS == types.Z_OBJ_HT(*op1).GetDoOperation()(ZEND_CONCAT, result, op1, op2) {
					return types.SUCCESS
				}
			} else if op2.IsObject() && types.Z_OBJ_HT(*op2).GetDoOperation() != nil && types.SUCCESS == types.Z_OBJ_HT(*op2).GetDoOperation()(ZEND_CONCAT, result, op1, op2) {
				return types.SUCCESS
			}
			op1_copy.SetString(ZvalGetStringFunc(op1))
			if EG__().GetException() != nil {

				if orig_op1 != result {
					result.SetUndef()
				}
				return types.FAILURE
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
		if op2.GetType() != types.IS_STRING {
			if op2.IsReference() {
				op2 = types.Z_REFVAL_P(op2)
				if op2.IsString() {
					break
				}
			}
			if op2.IsObject() && types.Z_OBJ_HT(*op2).GetDoOperation() != nil && types.SUCCESS == types.Z_OBJ_HT(*op2).GetDoOperation()(ZEND_CONCAT, result, op1, op2) {
				return types.SUCCESS
			}
			op2_copy.SetString(ZvalGetStringFunc(op2))
			if EG__().GetException() != nil {

				if orig_op1 != result {
					result.SetUndef()
				}
				return types.FAILURE
			}
			op2 = &op2_copy
		}
		break
	}
	if op1.GetStr().GetLen() == 0 {
		if result != op2 {
			if result == orig_op1 {
				IZvalPtrDtor(result)
			}
			types.ZVAL_COPY(result, op2)
		}
	} else if op2.GetStr().GetLen() == 0 {
		if result != op1 {
			if result == orig_op1 {
				IZvalPtrDtor(result)
			}
			types.ZVAL_COPY(result, op1)
		}
	} else {
		var op1_len int = op1.GetStr().GetLen()
		var op2_len int = op2.GetStr().GetLen()
		var result_len int = op1_len + op2_len
		var result_str *types.String
		if op1_len > types.STR_MAX_LEN-op2_len {
			faults.ThrowError(nil, "String size overflow")

			if orig_op1 != result {
				result.SetUndef()
			}
			return types.FAILURE
		}
		if result == op1 && result.IsRefcounted() {

			/* special case, perform operations on result */

			result_str = types.ZendStringExtend(result.GetStr(), result_len)

			/* special case, perform operations on result */

		} else {
			result_str = types.ZendStringAlloc(result_len, 0)
			memcpy(result_str.GetVal(), op1.GetStr().GetVal(), op1_len)
			if result == orig_op1 {
				IZvalPtrDtor(result)
			}
		}

		/* This has to happen first to account for the cases where result == op1 == op2 and
		 * the realloc is done. In this case this line will also update Z_STRVAL_P(op2) to
		 * point to the new string. The first op2_len bytes of result will still be the same. */

		result.SetString(result_str)
		memcpy(result_str.GetVal()+op1_len, op2.GetStr().GetVal(), op2_len)
		result_str.GetVal()[result_len] = '0'
	}

	return types.SUCCESS
}
func StringCompareFunction(op1 *types.Zval, op2 *types.Zval) int {
	if op1.IsString() && op2.IsString() {
		if op1.GetStr() == op2.GetStr() {
			return 0
		} else {
			return ZendBinaryStrcmp(op1.GetStr().GetStr(), op2.GetStr().GetStr())
		}
	} else {
		var tmp_str1 *types.String
		var tmp_str2 *types.String
		var str1 *types.String = ZvalGetTmpString(op1, &tmp_str1)
		var str2 *types.String = ZvalGetTmpString(op2, &tmp_str2)
		var ret int = ZendBinaryStrcmp(str1.GetStr(), str2.GetStr())
		ZendTmpStringRelease(tmp_str1)
		ZendTmpStringRelease(tmp_str2)
		return ret
	}
}
func StringCaseCompareFunction(op1 *types.Zval, op2 *types.Zval) int {
	if op1.IsString() && op2.IsString() {
		if op1.GetStr() == op2.GetStr() {
			return 0
		} else {
			return ZendBinaryStrcasecmpL(b.CastStr(op1.GetStr().GetVal(), op1.GetStr().GetLen()), b.CastStr(op2.GetStr().GetVal(), op2.GetStr().GetLen()))
		}
	} else {
		var tmp_str1 *types.String
		var tmp_str2 *types.String
		var str1 *types.String = ZvalGetTmpString(op1, &tmp_str1)
		var str2 *types.String = ZvalGetTmpString(op2, &tmp_str2)
		var ret int = ZendBinaryStrcasecmpL(b.CastStr(str1.GetVal(), str1.GetLen()), b.CastStr(str2.GetVal(), str1.GetLen()))
		ZendTmpStringRelease(tmp_str1)
		ZendTmpStringRelease(tmp_str2)
		return ret
	}
}
func StringLocaleCompareFunction(op1 *types.Zval, op2 *types.Zval) int {
	var tmp_str1 *types.String
	var tmp_str2 *types.String
	var str1 *types.String = ZvalGetTmpString(op1, &tmp_str1)
	var str2 *types.String = ZvalGetTmpString(op2, &tmp_str2)
	var ret int = strcoll(str1.GetVal(), str2.GetVal())
	ZendTmpStringRelease(tmp_str1)
	ZendTmpStringRelease(tmp_str2)
	return ret
}
func NumericCompareFunction(op1 *types.Zval, op2 *types.Zval) int {
	var d1 float64
	var d2 float64
	d1 = ZvalGetDouble(op1)
	d2 = ZvalGetDouble(op2)
	return ZEND_NORMALIZE_BOOL(d1 - d2)
}
func ZendFreeObjGetResult(op *types.Zval) {
	b.Assert(!(op.IsRefcounted()) || op.GetRefcount() != 0)
	ZvalPtrDtor(op)
}
func ConvertCompareResultToLong(result *types.Zval) {
	if result.IsDouble() {
		result.SetLong(ZEND_NORMALIZE_BOOL(result.GetDval()))
	} else {
		ConvertToLong(result)
	}
}
func CompareFunction(result *types.Zval, op1 *types.Zval, op2 *types.Zval) int {
	var ret int
	var converted int = 0
	var op1_copy types.Zval
	var op2_copy types.Zval
	var op_free *types.Zval
	var tmp_free types.Zval
	for true {
		switch TYPE_PAIR(op1.GetType(), op2.GetType()) {
		case TYPE_PAIR(types.IS_LONG, types.IS_LONG):
			result.SetLong(b.CondF2(op1.GetLval() > op2.GetLval(), 1, func() int {
				if op1.GetLval() < op2.GetLval() {
					return -1
				} else {
					return 0
				}
			}))
			return types.SUCCESS
		case TYPE_PAIR(types.IS_DOUBLE, types.IS_LONG):
			result.SetDval(op1.GetDval() - float64(op2.GetLval()))
			result.SetLong(ZEND_NORMALIZE_BOOL(result.GetDval()))
			return types.SUCCESS
		case TYPE_PAIR(types.IS_LONG, types.IS_DOUBLE):
			result.SetDval(float64(op1.GetLval() - op2.GetDval()))
			result.SetLong(ZEND_NORMALIZE_BOOL(result.GetDval()))
			return types.SUCCESS
		case TYPE_PAIR(types.IS_DOUBLE, types.IS_DOUBLE):
			if op1.GetDval() == op2.GetDval() {
				result.SetLong(0)
			} else {
				result.SetDval(op1.GetDval() - op2.GetDval())
				result.SetLong(ZEND_NORMALIZE_BOOL(result.GetDval()))
			}
			return types.SUCCESS
		case TYPE_PAIR(types.IS_ARRAY, types.IS_ARRAY):
			result.SetLong(ZendCompareArrays(op1, op2))
			return types.SUCCESS
		case TYPE_PAIR(types.IS_NULL, types.IS_NULL):
			fallthrough
		case TYPE_PAIR(types.IS_NULL, types.IS_FALSE):
			fallthrough
		case TYPE_PAIR(types.IS_FALSE, types.IS_NULL):
			fallthrough
		case TYPE_PAIR(types.IS_FALSE, types.IS_FALSE):
			fallthrough
		case TYPE_PAIR(types.IS_TRUE, types.IS_TRUE):
			result.SetLong(0)
			return types.SUCCESS
		case TYPE_PAIR(types.IS_NULL, types.IS_TRUE):
			result.SetLong(-1)
			return types.SUCCESS
		case TYPE_PAIR(types.IS_TRUE, types.IS_NULL):
			result.SetLong(1)
			return types.SUCCESS
		case TYPE_PAIR(types.IS_STRING, types.IS_STRING):
			if op1.GetStr() == op2.GetStr() {
				result.SetLong(0)
				return types.SUCCESS
			}
			result.SetLong(ZendiSmartStrcmp(op1.GetStrVal(), op2.GetStrVal()))
			return types.SUCCESS
		case TYPE_PAIR(types.IS_NULL, types.IS_STRING):
			result.SetLong(b.Cond(op2.GetStr().GetLen() == 0, 0, -1))
			return types.SUCCESS
		case TYPE_PAIR(types.IS_STRING, types.IS_NULL):
			result.SetLong(b.Cond(op1.GetStr().GetLen() == 0, 0, 1))
			return types.SUCCESS
		case TYPE_PAIR(types.IS_OBJECT, types.IS_NULL):
			result.SetLong(1)
			return types.SUCCESS
		case TYPE_PAIR(types.IS_NULL, types.IS_OBJECT):
			result.SetLong(-1)
			return types.SUCCESS
		default:
			if op1.IsReference() {
				op1 = types.Z_REFVAL_P(op1)
				continue
			} else if op2.IsReference() {
				op2 = types.Z_REFVAL_P(op2)
				continue
			}
			if op1.IsObject() && types.Z_OBJ_HT(*op1).GetCompare() != nil {
				ret = types.Z_OBJ_HT(*op1).GetCompare()(result, op1, op2)
				if result.GetType() != types.IS_LONG {
					ConvertCompareResultToLong(result)
				}
				return ret
			} else if op2.IsObject() && types.Z_OBJ_HT(*op2).GetCompare() != nil {
				ret = types.Z_OBJ_HT(*op2).GetCompare()(result, op1, op2)
				if result.GetType() != types.IS_LONG {
					ConvertCompareResultToLong(result)
				}
				return ret
			}
			if op1.IsObject() && op2.IsObject() {
				if op1.GetObj() == op2.GetObj() {

					/* object handles are identical, apparently this is the same object */

					result.SetLong(0)
					return types.SUCCESS
				}
				if types.Z_OBJ_HT(*op1).GetCompareObjects() == types.Z_OBJ_HT(*op2).GetCompareObjects() {
					result.SetLong(types.Z_OBJ_HT(*op1).GetCompareObjects()(op1, op2))
					return types.SUCCESS
				}
			}
			if op1.IsObject() {
				if types.Z_OBJ_HT_P(op1).GetGet() != nil {
					var rv types.Zval
					op_free = types.Z_OBJ_HT_P(op1).GetGet()(op1, &rv)
					ret = CompareFunction(result, op_free, op2)
					ZendFreeObjGetResult(op_free)
					return ret
				} else if op2.GetType() != types.IS_OBJECT && types.Z_OBJ_HT_P(op1).GetCastObject() != nil {
					tmp_free.SetUndef()
					if types.Z_OBJ_HT_P(op1).GetCastObject()(op1, &tmp_free, b.CondF2(op2.IsFalse() || op2.IsTrue(), types.IS_BOOL, func() __auto__ { return op2.GetType() })) == types.FAILURE {
						result.SetLong(1)
						ZendFreeObjGetResult(&tmp_free)
						return types.SUCCESS
					}
					ret = CompareFunction(result, &tmp_free, op2)
					ZendFreeObjGetResult(&tmp_free)
					return ret
				}
			}
			if op2.IsObject() {
				if types.Z_OBJ_HT_P(op2).GetGet() != nil {
					var rv types.Zval
					op_free = types.Z_OBJ_HT_P(op2).GetGet()(op2, &rv)
					ret = CompareFunction(result, op1, op_free)
					ZendFreeObjGetResult(op_free)
					return ret
				} else if op1.GetType() != types.IS_OBJECT && types.Z_OBJ_HT_P(op2).GetCastObject() != nil {
					tmp_free.SetUndef()
					if types.Z_OBJ_HT_P(op2).GetCastObject()(op2, &tmp_free, b.CondF2(op1.IsFalse() || op1.IsTrue(), types.IS_BOOL, func() __auto__ { return op1.GetType() })) == types.FAILURE {
						result.SetLong(-1)
						ZendFreeObjGetResult(&tmp_free)
						return types.SUCCESS
					}
					ret = CompareFunction(result, op1, &tmp_free)
					ZendFreeObjGetResult(&tmp_free)
					return ret
				} else if op1.IsObject() {
					result.SetLong(1)
					return types.SUCCESS
				}
			}
			if converted == 0 {
				if op1.GetType() < types.IS_TRUE {
					result.SetLong(b.Cond(ZvalIsTrue(op2) != 0, -1, 0))
					return types.SUCCESS
				} else if op1.IsTrue() {
					result.SetLong(b.Cond(ZvalIsTrue(op2) != 0, 0, 1))
					return types.SUCCESS
				} else if op2.GetType() < types.IS_TRUE {
					result.SetLong(b.Cond(ZvalIsTrue(op1) != 0, 1, 0))
					return types.SUCCESS
				} else if op2.IsTrue() {
					result.SetLong(b.Cond(ZvalIsTrue(op1) != 0, 0, -1))
					return types.SUCCESS
				} else {
					op1 = ZendiConvertScalarToNumber(op1, &op1_copy, result, 1)
					op2 = ZendiConvertScalarToNumber(op2, &op2_copy, result, 1)
					if EG__().GetException() != nil {
						if result != op1 {
							result.SetUndef()
						}
						return types.FAILURE
					}
					converted = 1
				}
			} else if op1.IsArray() {
				result.SetLong(1)
				return types.SUCCESS
			} else if op2.IsArray() {
				result.SetLong(-1)
				return types.SUCCESS
			} else {
				b.Assert(false)
				faults.ThrowError(nil, "Unsupported operand types")
				if result != op1 {
					result.SetUndef()
				}
				return types.FAILURE
			}
		}
	}
}
func HashZvalIdenticalFunction(z1 *types.Zval, z2 *types.Zval) int {
	/* is_identical_function() returns 1 in case of identity and 0 in case
	 * of a difference;
	 * whereas this comparison function is expected to return 0 on identity,
	 * and non zero otherwise.
	 */

	z1 = types.ZVAL_DEREF(z1)
	z2 = types.ZVAL_DEREF(z2)
	return FastIsNotIdenticalFunction(z1, z2)
}
func ZendIsIdentical(op1 *types.Zval, op2 *types.Zval) bool {
	if op1.GetType() != op2.GetType() {
		return false
	}
	switch op1.GetType() {
	case types.IS_NULL, types.IS_FALSE, types.IS_TRUE:
		return true
	case types.IS_LONG:
		return op1.GetLval() == op2.GetLval()
	case types.IS_RESOURCE:
		return op1.GetRes() == op2.GetRes()
	case types.IS_DOUBLE:
		return op1.GetDval() == op2.GetDval()
	case types.IS_STRING:
		return op1.GetStrVal() == op2.GetStrVal()
	case types.IS_ARRAY:
		return op1.GetArr() == op2.GetArr() || types.ZendHashCompare(op1.GetArr(), op2.GetArr(), types.CompareFuncT(HashZvalIdenticalFunction), 1) == 0
	case types.IS_OBJECT:
		return op1.GetObj() == op2.GetObj()
	default:
		return false
	}
}
func IsIdenticalFunction(result *types.Zval, op1 *types.Zval, op2 *types.Zval) int {
	types.ZVAL_BOOL(result, ZendIsIdentical(op1, op2) != 0)
	return types.SUCCESS
}
func IsNotIdenticalFunction(result *types.Zval, op1 *types.Zval, op2 *types.Zval) int {
	types.ZVAL_BOOL(result, ZendIsIdentical(op1, op2) == 0)
	return types.SUCCESS
}
func IsEqualFunction(result *types.Zval, op1 *types.Zval, op2 *types.Zval) int {
	if CompareFunction(result, op1, op2) == types.FAILURE {
		return types.FAILURE
	}
	types.ZVAL_BOOL(result, result.GetLval() == 0)
	return types.SUCCESS
}
func IsNotEqualFunction(result *types.Zval, op1 *types.Zval, op2 *types.Zval) int {
	if CompareFunction(result, op1, op2) == types.FAILURE {
		return types.FAILURE
	}
	types.ZVAL_BOOL(result, result.GetLval() != 0)
	return types.SUCCESS
}
func IsSmallerFunction(result *types.Zval, op1 *types.Zval, op2 *types.Zval) int {
	if CompareFunction(result, op1, op2) == types.FAILURE {
		return types.FAILURE
	}
	types.ZVAL_BOOL(result, result.GetLval() < 0)
	return types.SUCCESS
}
func IsSmallerOrEqualFunction(result *types.Zval, op1 *types.Zval, op2 *types.Zval) int {
	if CompareFunction(result, op1, op2) == types.FAILURE {
		return types.FAILURE
	}
	types.ZVAL_BOOL(result, result.GetLval() <= 0)
	return types.SUCCESS
}
func InstanceofClass(instance_ce *types.ClassEntry, ce *types.ClassEntry) types.ZendBool {
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
func InstanceofInterface(instance_ce *types.ClassEntry, ce *types.ClassEntry) types.ZendBool {
	var i uint32
	if instance_ce.GetNumInterfaces() != 0 {
		b.Assert(instance_ce.IsResolvedInterfaces())
		for i = 0; i < instance_ce.GetNumInterfaces(); i++ {
			if instance_ce.GetInterfaces()[i] == ce {
				return 1
			}
		}
	}
	return instance_ce == ce
}
func InstanceofFunctionEx(instance_ce *types.ClassEntry, ce *types.ClassEntry, is_interface types.ZendBool) types.ZendBool {
	if is_interface != 0 {
		b.Assert(ce.IsInterface())
		return InstanceofInterface(instance_ce, ce)
	} else {
		b.Assert(!ce.IsInterface())
		return InstanceofClass(instance_ce, ce)
	}
}
func InstanceofFunction(instance_ce *types.ClassEntry, ce *types.ClassEntry) types.ZendBool {
	if ce.IsInterface() {
		return InstanceofInterface(instance_ce, ce)
	} else {
		return InstanceofClass(instance_ce, ce)
	}
}
func IncrementString(str *types.Zval) {
	// notice: 前置要求 str 必须是 IS_STRING 类型，且其值不为数字字符串
	str.SetStringVal(IncrementStringEx(str.GetStrVal()))
}

func IncrementStringEx(str string) string {
	// notice: 前置要求 str 其值不为数字字符串
	if str == "" {
		return "1"
	}

	s := []byte(str)
	last := 0
	carry := false
	for i := len(s) - 1; i >= 0; i-- {
		c := s[i]
		if ascii.IsLower(c) {
			last = LOWER_CASE
			carry = c == 'z'
			s[i] = b.Cond(carry, 'a', c+1)
		} else if ascii.IsUpper(c) {
			last = UPPER_CASE
			carry = c == 'Z'
			s[i] = b.Cond(carry, 'A', c+1)
		} else if ascii.IsDigit(c) {
			last = NUMERIC
			carry = c == '9'
			s[i] = b.Cond(carry, '0', c+1)
		} else {
			carry = false
			break
		}
		if !carry {
			break
		}
	}

	if carry {
		switch last {
		case NUMERIC:
			return "1" + string(s)
		case UPPER_CASE:
			return "A" + string(s)
		case LOWER_CASE:
			return "a" + string(s)
		}
	}

	return string(s)
}

func IncrementFunction(op1 *types.Zval) int {
try_again:
	switch op1.GetType() {
	case types.IS_LONG:
		FastLongIncrementFunction(op1)
	case types.IS_DOUBLE:
		op1.SetDval(op1.GetDval() + 1)
	case types.IS_NULL:
		op1.SetLong(1)
	case types.IS_STRING:
		var lval ZendLong
		var dval float64
		switch IsNumericString(op1.GetStr().GetStr(), &lval, &dval, 0) {
		case types.IS_LONG:

			if lval == ZEND_LONG_MAX {

				/* switch to double */

				var d float64 = float64(lval)
				op1.SetDouble(d + 1)
			} else {
				op1.SetLong(lval + 1)
			}
		case types.IS_DOUBLE:

			op1.SetDouble(dval + 1)
		default:
			/* Perl style string increment */
			IncrementString(op1)
		}
	case types.IS_OBJECT:
		if types.Z_OBJ_HT(*op1).GetGet() != nil && types.Z_OBJ_HT(*op1).GetSet() != nil {

			/* proxy object */

			var rv types.Zval
			var val *types.Zval
			val = types.Z_OBJ_HT(*op1).GetGet()(op1, &rv)
			val.TryAddRefcount()
			IncrementFunction(val)
			types.Z_OBJ_HT(*op1).GetSet()(op1, val)
			ZvalPtrDtor(val)
		} else if types.Z_OBJ_HT(*op1).GetDoOperation() != nil {
			var op2 types.Zval
			var res int
			op2.SetLong(1)
			res = types.Z_OBJ_HT(*op1).GetDoOperation()(ZEND_ADD, op1, op1, &op2)
			return res
		}
		return types.FAILURE
	case types.IS_REFERENCE:
		op1 = types.Z_REFVAL_P(op1)
		goto try_again
	default:
		return types.FAILURE
	}
	return types.SUCCESS
}
func DecrementFunction(op1 *types.Zval) int {
	var lval ZendLong
	var dval float64
try_again:
	switch op1.GetType() {
	case types.IS_LONG:
		FastLongDecrementFunction(op1)
	case types.IS_DOUBLE:
		op1.SetDval(op1.GetDval() - 1)
	case types.IS_STRING:
		if op1.GetStr().GetLen() == 0 {

			op1.SetLong(-1)
			break
		}
		switch IsNumericString(op1.GetStr().GetStr(), &lval, &dval, 0) {
		case types.IS_LONG:

			if lval == ZEND_LONG_MIN {
				var d float64 = float64(lval)
				op1.SetDouble(d - 1)
			} else {
				op1.SetLong(lval - 1)
			}
		case types.IS_DOUBLE:

			op1.SetDouble(dval - 1)
		}
	case types.IS_OBJECT:
		if types.Z_OBJ_HT(*op1).GetGet() != nil && types.Z_OBJ_HT(*op1).GetSet() != nil {

			/* proxy object */

			var rv types.Zval
			var val *types.Zval
			val = types.Z_OBJ_HT(*op1).GetGet()(op1, &rv)
			val.TryAddRefcount()
			DecrementFunction(val)
			types.Z_OBJ_HT(*op1).GetSet()(op1, val)
			ZvalPtrDtor(val)
		} else if types.Z_OBJ_HT(*op1).GetDoOperation() != nil {
			var op2 types.Zval
			var res int
			op2.SetLong(1)
			res = types.Z_OBJ_HT(*op1).GetDoOperation()(ZEND_SUB, op1, op1, &op2)
			return res
		}
		return types.FAILURE
	case types.IS_REFERENCE:
		op1 = types.Z_REFVAL_P(op1)
		goto try_again
	default:
		return types.FAILURE
	}
	return types.SUCCESS
}
func ZendIsTrueEx(op *types.Zval) bool { return IZendIsTrue(op) != 0 }
func ZendIsTrue(op *types.Zval) int    { return IZendIsTrue(op) }
func ZendObjectIsTrue(op *types.Zval) bool {
	if types.Z_OBJ_HT_P(op).GetCastObject() != nil {
		var tmp types.Zval
		if types.Z_OBJ_HT_P(op).GetCastObject()(op, &tmp, types.IS_BOOL) == types.SUCCESS {
			return tmp.IsTrue()
		}
		faults.Error(faults.E_RECOVERABLE_ERROR, "Object of class %s could not be converted to bool", types.Z_OBJ_P(op).GetCe().GetName().GetVal())
	} else if types.Z_OBJ_HT_P(op).GetGet() != nil {
		var result bool
		var rv types.Zval
		var tmp *types.Zval = types.Z_OBJ_HT_P(op).GetGet()(op, &rv)
		if tmp.GetType() != types.IS_OBJECT {

			/* for safety - avoid loop */

			result = IZendIsTrueEx(tmp)
			ZvalPtrDtor(tmp)
			return result
		}
	}
	return true
}
func ZendStrTolowerCopy(dest *byte, source *byte, length int) *byte {
	var str *uint8 = (*uint8)(source)
	var result *uint8 = (*uint8)(dest)
	var end *uint8 = str + length
	for str < end {
		b.PostInc(&(*result)) = ascii.ToLower(b.PostInc(&(*str)))
	}
	*result = '0'
	return dest
}
func ZendStrTolower(str *byte, length int) {
	var p *uint8 = (*uint8)(str)
	var end *uint8 = p + length
	for p < end {
		*p = ascii.ToLower(*p)
		p++
	}
}
func ZendStrTolowerDupEx(source *byte, length int) *byte {
	var p *uint8 = (*uint8)(source)
	var end *uint8 = p + length
	for p < end {
		if (*p) != ascii.ToLower(*p) {
			var res *byte = (*byte)(Emalloc(length + 1))
			var r *uint8
			if p != (*uint8)(source) {
				memcpy(res, source, p-(*uint8)(source))
			}
			r = (*uint8)(p + (res - source))
			for p < end {
				*r = ascii.ToLower(*p)
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
func ZendStringTolowerEx(str *types.String) *types.String {
	return types.NewString(ascii.StrToLower(str.GetStr()))
}
func ZendBinaryStrcmp(s1 string, s2 string) int {
	return strings.Compare(s1, s2)
}
func ZendBinaryStrncmp(s1 string, s2 string, length int) int {
	if len(s1) > length {
		s1 = s1[:length]
	}
	if len(s2) > length {
		s2 = s2[:length]
	}
	return strings.Compare(s1, s2)
}
func ZendBinaryStrcasecmp(s1 string, s2 string) int { return ascii.StrCaseCompare(s1, s2) }
func ZendBinaryStrncasecmp(s1 string, s2 string, length int) int {
	if len(s1) > length {
		s1 = s1[:length]
	}
	if len(s2) > length {
		s2 = s2[:length]
	}
	return ascii.StrCaseCompare(s1, s2)
}
func ZendBinaryStrcasecmpL(s1 string, s2 string) int {
	return ascii.StrCaseCompare(s1, s2)
}
func ZendBinaryStrncasecmpL(s1 string, s2 string, length int) int {
	if len(s1) > length {
		s1 = s1[:length]
	}
	if len(s2) > length {
		s2 = s2[:length]
	}
	return ascii.StrCaseCompare(s1, s2)
}
func ZendiSmartStreq(s1 *types.String, s2 *types.String) int {
	var ret1 int
	var ret2 int
	var oflow1 int
	var oflow2 int
	var lval1 ZendLong = 0
	var lval2 ZendLong = 0
	var dval1 float64 = 0.0
	var dval2 float64 = 0.0
	if b.Assign(&ret1, IsNumericStringEx(s1.GetStr(), &lval1, &dval1, 0, &oflow1)) && b.Assign(&ret2, IsNumericStringEx(s2.GetStr(), &lval2, &dval2, 0, &oflow2)) {
		if oflow1 != 0 && oflow1 == oflow2 && dval1-dval2 == 0.0 {

			/* both values are integers overflown to the same side, and the
			 * double comparison may have resulted in crucial accuracy lost */

			goto string_cmp

			/* both values are integers overflown to the same side, and the
			 * double comparison may have resulted in crucial accuracy lost */

		}
		if ret1 == types.IS_DOUBLE || ret2 == types.IS_DOUBLE {
			if ret1 != types.IS_DOUBLE {
				if oflow2 != 0 {

					/* 2nd operand is integer > LONG_MAX (oflow2==1) or < LONG_MIN (-1) */

					return 0

					/* 2nd operand is integer > LONG_MAX (oflow2==1) or < LONG_MIN (-1) */

				}
				dval1 = float64(lval1)
			} else if ret2 != types.IS_DOUBLE {
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
		return IntBool(s1.GetStr() == s2.GetStr())
	}
}
func ZendiSmartStrcmp(s1 string, s2 string) int {
	var r1, r2 NumericStrResult
	r1 = ConvertNumericStr(s1, 0)
	if r1.Type == 0 {
		goto string_cmp
	}
	r2 = ConvertNumericStr(s2, 0)
	if r2.Type == 0 {
		goto string_cmp
	}

	if r1.Overflow != 0 && r1.Overflow == r2.Overflow && r1.Dval-r2.Dval == 0.0 {
		/* both values are integers overflown to the same side, and the
		 * double comparison may have resulted in crucial accuracy lost */
		goto string_cmp
	}
	if r1.Type == types.IS_DOUBLE || r2.Type == types.IS_DOUBLE {
		dval1, dval2 := r1.Dval, r2.Dval
		if r1.Type != types.IS_DOUBLE {
			if r2.Overflow != 0 {
				/* 2nd operand is integer > LONG_MAX (oflow2==1) or < LONG_MIN (-1) */
				return -1 * r2.Overflow
			}
			dval1 = float64(r1.Lval)
		} else if r2.Type != types.IS_DOUBLE {
			if r1.Overflow != 0 {
				return r1.Overflow
			}
			dval2 = float64(r2.Lval)
		} else if r2.Dval == r2.Dval && !(core.ZendFinite(r1.Dval)) {
			/* Both values overflowed and have the same sign,
			 * so a numeric comparison would be inaccurate */
			goto string_cmp
		}
		return b.Compare(dval1, dval2)
	} else {
		return b.Compare(r1.Lval, r2.Lval)
	}

string_cmp:
	return ZendBinaryStrcmp(s1, s2)
}
func HashZvalCompareFunction(z1 *types.Zval, z2 *types.Zval) int {
	var result types.Zval
	if CompareFunction(&result, z1, z2) == types.FAILURE {
		return 1
	}
	return result.GetLval()
}
func ZendCompareSymbolTables(ht1 *types.Array, ht2 *types.Array) int {
	if ht1 == ht2 {
		return 0
	} else {
		return types.ZendHashCompare(ht1, ht2, types.CompareFuncT(HashZvalCompareFunction), 0)
	}
}
func ZendCompareArrays(a1 *types.Zval, a2 *types.Zval) int {
	return ZendCompareSymbolTables(a1.GetArr(), a2.GetArr())
}
func ZendCompareObjects(o1 *types.Zval, o2 *types.Zval) int {
	if o1.GetObj() == o2.GetObj() {
		return 0
	}
	if types.Z_OBJ_HT_P(o1).GetCompareObjects() == nil {
		return 1
	} else {
		return types.Z_OBJ_HT_P(o1).GetCompareObjects()(o1, o2)
	}
}
func ZendLongToStr(num ZendLong) *types.String {
	var res = strconv.Itoa(num)
	return types.NewString(res)
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
	dmod := math.Mod(d, 1<<64)
	if dmod > math.MaxInt {
		dmod -= 1 << 64
	} else if dmod < math.MinInt {
		dmod += 1 << 64
	}
	return ZendLong(dmod)
}
