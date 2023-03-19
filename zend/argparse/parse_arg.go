package argparse

import (
	"fmt"
	b "sik/builtin"
	"sik/zend"
	"sik/zend/types"
)

type parseArgError struct {
	message  string
	severity int
}

func (p parseArgError) Error() string { return p.message }

func parseTypeError(arg *types.Zval, expectedType string) *parseArgError {
	return parseError(0, "to be %s, %s given", expectedType, zend.ZendZvalTypeName(arg))
}
func parseError(severity int, format string, args ...any) *parseArgError {
	return &parseArgError{severity: severity, message: fmt.Sprintf(format, args...)}
}

func ZendParseArg(arg_num int, arg *types.Zval, va *VaArgsReceiver, spec *b.StrReader, flags int) int {
	err := ZendParseArgImpl(arg, va, spec)
	if err != nil {
		if zend.EG__().GetException() != nil {
			return types.FAILURE
		}
		if (flags & ZEND_PARSE_PARAMS_QUIET) == 0 {
			var throwException = zend.CurrEX().IsArgUseStrictTypes() || (flags&ZEND_PARSE_PARAMS_THROW) != 0
			zend.ZendInternalTypeError(throwException, "%s() expects parameter %d %s", zend.GetActiveCalleeName(), arg_num, err.Error())
		}
		if err.severity != zend.E_DEPRECATED {
			return types.FAILURE
		}
	}
	return types.SUCCESS
}

func ZendParseArgImpl(arg *types.Zval, va *VaArgsReceiver, spec *b.StrReader) *parseArgError {
	specWalk := spec.Copy()
	c := specWalk.Read()
	var check_null int = 0
	var separate int = 0
	var real_arg *types.Zval = arg

	/* scan through modifiers */

	arg = types.ZVAL_DEREF(arg)
	for true {
		if specWalk.Curr() == '/' {
			types.SEPARATE_ZVAL_NOREF(arg)
			real_arg = arg
			separate = 1
		} else if specWalk.Curr() == '!' {
			check_null = 1
		} else {
			break
		}
		specWalk.Next()
	}
	switch c {
	case 'l', 'L':
		if val, isNull, ok := ParseLong(arg, check_null != 0, c == 'L'); ok {
			va.Long(val)
			if check_null != 0 {
				va.Bool(isNull)
			}
		} else {
			return parseTypeError(arg, "int")
		}
	case 'd':
		if val, isNull, ok := ParseDouble(arg, check_null != 0); ok {
			va.Double(val)
			if check_null != 0 {
				va.Bool(isNull)
			}
		} else {
			return parseTypeError(arg, "float")
		}
	case 's':
		var p **byte = __va_arg(*va, (**byte)(_))
		var pl *int = __va_arg(*va, (*int)(_))
		if ZendParseArgString(arg, p, pl, check_null) == 0 {
			return parseTypeError(arg, "string")
		}
	case 'p':
		var p **byte = __va_arg(*va, (**byte)(_))
		var pl *int = __va_arg(*va, (*int)(_))
		if ZendParseArgPath(arg, p, pl, check_null) == 0 {
			return parseTypeError(arg, "a valid path")
		}
	case 'P':
		var str **types.ZendString = __va_arg(*va, (**types.ZendString)(_))
		if ZendParseArgPathStr(arg, str, check_null) == 0 {
			return parseTypeError(arg, "a valid path")
		}
	case 'S':
		if val, ok := ParseZStr(arg, check_null != 0); ok {
			va.ZStr(val)
		} else {
			return parseTypeError(arg, "string")
		}
	case 'b':
		if val, isNull, ok := ParseBool(arg, check_null != 0); ok {
			va.Bool(val)
			if check_null != 0 {
				va.Bool(isNull)
			}
		} else {
			return parseTypeError(arg, "bool")
		}
	case 'r':
		var p **types.Zval = __va_arg(*va, (**types.Zval)(_))
		if ZendParseArgResource(arg, p, check_null) == 0 {
			return parseTypeError(arg, "resource")
		}
	case 'A', 'a':
		var p **types.Zval = __va_arg(*va, (**types.Zval)(_))
		if ZendParseArgArray(arg, p, check_null, c == 'A') == 0 {
			return parseTypeError(arg, "array")
		}
	case 'H', 'h':
		var p **types.HashTable = __va_arg(*va, (**types.HashTable)(_))
		if ZendParseArgArrayHt(arg, p, check_null, c == 'H', separate) == 0 {
			return parseTypeError(arg, "array")
		}
	case 'o':
		var p **types.Zval = __va_arg(*va, (**types.Zval)(_))
		if ZendParseArgObject(arg, p, nil, check_null) == 0 {
			return parseTypeError(arg, "object")
		}
	case 'O':
		var p **types.Zval = __va_arg(*va, (**types.Zval)(_))
		var ce *zend.ZendClassEntry = __va_arg(*va, (*zend.ZendClassEntry)(_))
		if ZendParseArgObject(arg, p, ce, check_null) == 0 {
			if ce != nil {
				return ce.GetName().GetVal()
			} else {
				return parseTypeError(arg, "object")
			}
		}
	case 'C':
		var lookup *zend.ZendClassEntry
		var pce **zend.ZendClassEntry = __va_arg(*va, (**zend.ZendClassEntry)(_))
		var ce_base *zend.ZendClassEntry = *pce
		if check_null != 0 && arg.IsNull() {
			*pce = nil
			break
		}
		if zend.TryConvertToString(arg) == 0 {
			*pce = nil
			return parseTypeError(arg, "valid class name")
		}
		if b.Assign(&lookup, zend.ZendLookupClass(arg.GetStr())) == nil {
			*pce = nil
		} else {
			*pce = lookup
		}
		if ce_base != nil {
			if (*pce) == nil || zend.InstanceofFunction(*pce, ce_base) == 0 {
				return parseError(0, "to be a class name derived from %s, '%s' given", ce_base.Name(), arg.GetRawStr())
			}
		}
		if (*pce) == nil {
			return parseError(0, "to be a valid class name, '%s' given", arg.GetRawStr())
		}
	case 'f':
		var fci *zend.ZendFcallInfo = __va_arg(*va, (*zend.ZendFcallInfo)(_))
		var fcc *zend.ZendFcallInfoCache = __va_arg(*va, (*zend.ZendFcallInfoCache)(_))
		var is_callable_error *byte = nil
		if check_null != 0 && arg.IsNull() {
			fci.SetSize(0)
			fcc.SetFunctionHandler(0)
			break
		}
		if zend.ZendFcallInfoInit(arg, 0, fci, fcc, nil, &is_callable_error) == types.SUCCESS {
			if is_callable_error != nil {
				*spec = *specWalk
				return parseError(zend.E_DEPRECATED, "to be a valid callback, %s", is_callable_error)
			}
			break
		} else {
			if is_callable_error != nil {
				return parseError(zend.E_ERROR, "to be a valid callback, %s", is_callable_error)
			} else {
				return parseTypeError(arg, "valid callback")
			}
		}
	case 'z':
		var p **types.Zval = __va_arg(*va, (**types.Zval)(_))
		ZendParseArgZvalDeref(real_arg, p, check_null)
	case 'Z':
		/* 'Z' iz not supported anymore and should be replaced with 'z' */
		b.Assert(c != 'Z')
	default:
		return parseTypeError(arg, "unknown")
	}
	*spec = *specWalk
	return nil
}
