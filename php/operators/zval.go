package operators

import (
	"fmt"
	"github.com/heyuuu/gophp/php/faults"
	"github.com/heyuuu/gophp/php/types"
	"strconv"
)

// bool

func ZvalIsTrue(op Val) bool {
again:
	switch op.Type() {
	case types.IsTrue:
		return true
	case types.IsLong:
		return op.Long() != 0
	case types.IsDouble:
		return op.Double() != 0
	case types.IsString:
		str := op.String()
		return str != "" && str != "0"
	case types.IsArray:
		return op.Array().Len() != 0
	case types.IsObject:
		dst := ConvertObjectToType(op.Object(), types.IsBool)
		if dst != nil {
			return dst.IsTrue()
		}
		return true
	case types.IsResource:
		return op.ResourceHandle() != 0
	case types.IsRef:
		op = op.RefVal()
		goto again
	}
	return false
}

// long

func ZvalGetLong(op Val) int      { return zvalGetLongEx(op, true) }
func zvalGetLongNoisy(op Val) int { return zvalGetLongEx(op, false) }
func zvalGetLongEx(op Val, silent bool) int {
	// fast
	if op.IsLong() {
		return op.Long()
	}

	// common
again:
	switch op.Type() {
	case types.IsUndef, types.IsNull, types.IsFalse:
		return 0
	case types.IsTrue:
		return 1
	case types.IsResource:
		return op.ResourceHandle()
	case types.IsLong:
		return op.Long()
	case types.IsDouble:
		return DoubleToLong(op.Double())
	case types.IsString:
		var r Val = StrToNumberPrefix(op.String(), silent)
		if r == nil {
			if !silent {
				faults.Error(faults.E_WARNING, "A non-numeric value encountered")
			}
			return 0
		}
		if r.IsLong() {
			return r.Long()
		} else { //  r.IsDouble()
			/* Previously we used strtol here, not is_numeric_string,
			 * and strtol gives you LONG_MAX/_MIN on overflow.
			 * We use use saturating conversion to emulate strtol()'s
			 * behaviour.
			 */
			return DoubleToLongCap(r.Double())
		}
	case types.IsArray:
		if op.Array().Len() != 0 {
			return 1
		} else {
			return 0
		}
	case types.IsObject:
		dst := ConvertObjectToType(op.Object(), types.IsLong)
		if dst.IsLong() {
			return dst.Long()
		} else {
			return 1
		}
	case types.IsRef:
		op = op.DeRef()
		goto again
	default:
		return 0
	}
}

// ZvalTryGetLong。 相比 ZvalGetLong，不考虑 Array/Object/Resource 等复杂类型。
func ZvalTryGetLong(op Val) (int, bool) {
	op = op.DeRef()
	if op.Type() < types.IsString {
		return ZvalGetLong(op), true
	} else if op.IsString() {
		v, err := strconv.Atoi(op.String())
		if err == nil {
			return v, true
		}
	}
	return 0, false
}

// double

func ZvalGetDouble(op Val) float64 {
	if op.IsDouble() {
		return op.Double()
	}

	op = op.DeRef()
	switch op.Type() {
	case types.IsNull:
		fallthrough
	case types.IsFalse:
		return 0.0
	case types.IsTrue:
		return 1.0
	case types.IsResource:
		return float64(op.ResourceHandle())
	case types.IsLong:
		return float64(op.Long())
	case types.IsDouble:
		return op.Double()
	case types.IsString:
		return StrToDouble(op.String())
	case types.IsArray:
		if op.Array().Len() != 0 {
			return 1.0
		} else {
			return 0.0
		}
	case types.IsObject:
		dst := ConvertObjectToType(op.Object(), types.IsDouble)
		if dst.IsDouble() {
			return dst.Double()
		} else {
			return 1.0
		}
	default:
		return 0.0
	}
}

// scalar to number
func ScalarGetNumber(op Val, silent bool) Val {
	switch op.Type() {
	case types.IsNull, types.IsFalse:
		return Long(0)
	case types.IsTrue:
		return Long(1)
	case types.IsLong:
		return Long(op.Long())
	case types.IsDouble:
		return Double(op.Double())
	case types.IsString:
		r := StrToNumberPrefix(op.String(), silent)
		if r == nil {
			if !silent {
				faults.Error(faults.E_WARNING, "A non-numeric value encountered")
			}
			return Long(0)
		}
		return r
	case types.IsResource:
		var l = op.ResourceHandle()
		return Long(l)
	case types.IsObject:
		dst := ConvertObjectToType(op.Object(), types.IsNumber)
		if hasException() {
			return Long(1)
		}
		if dst.IsLong() || dst.IsDouble() {
			return dst
		} else {
			return Long(1)
		}
	default:
		// todo fail log
		return nil
	}
}

// string

func ZvalGetStrVal(op Val) string {
	str, _ := ZvalGetStr(op)
	return str
}
func ZvalGetStr(op Val) (string, bool) {
	return zvalGetStrEx(op, false)
}
func ZvalTryGetStrVal(op Val) string {
	str, _ := zvalGetStrEx(op, true)
	return str
}
func ZvalTryGetStr(op Val) (string, bool) {
	return zvalGetStrEx(op, true)
}

/**
 * 从 Zval 转字符串
 * @return string 返回的字符串值。
 * @return bool   是否成功。
 */
func zvalGetStrEx(op Val, try bool) (string, bool) {
	op = op.DeRef()
	switch op.Type() {
	case types.IsString:
		return op.String(), true
	case types.IsUndef, types.IsNull, types.IsFalse:
		return "", true
	case types.IsTrue:
		return "1", true
	case types.IsResource:
		return fmt.Sprintf("Resource id #%d", op.ResourceHandle()), true
	case types.IsLong:
		return strconv.Itoa(op.Long()), true
	case types.IsDouble:
		return fmt.Sprintf("%.*G", getPrecision(), op.Double()), true
	case types.IsArray:
		faults.Error(faults.E_NOTICE, "Array to string conversion")
		if try && hasException() {
			return "", false
		}
		return "Array", true
	case types.IsObject:
		if tmp, ok := op.Object().Cast(types.IsString); ok {
			return tmp.String(), true
		}
		if !hasException() {
			faults.ThrowError(nil, fmt.Sprintf("Object of class %s could not be converted to string", op.Object().CeName()))
		}
		if try {
			return "", false
		} else {
			return "", true
		}
	default:
		return "", false
	}
}
