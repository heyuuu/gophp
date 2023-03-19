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

func parseError(severity int, format string, args ...any) *parseArgError {
	return &parseArgError{severity: severity, message: fmt.Sprintf(format, args...)}
}

func (p *OldParser) ZendParseArg(arg_num int, arg *types.Zval, va *VaArgsReceiver, spec *typeSpecReader, flags int) int {
	p.ZendParseArgImpl(arg, va, spec)
	err := p.err
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

func (p *OldParser) ZendParseArgImpl(arg *types.Zval, va *VaArgsReceiver, spec *typeSpecReader) {
	/* scan through modifiers */
	typ, checkNull, separate := spec.Next()
	check_null := types.IntBool(checkNull)

	real_arg := arg
	arg = types.ZVAL_DEREF(arg)
	if separate {
		types.SEPARATE_ZVAL_NOREF(arg)
		real_arg = arg
	}
	p.arg = arg

	switch typ {
	case 'l', 'L':
		p.parseLong(checkNull, typ == 'L')
	case 'd':
		p.parseDouble(checkNull)
	case 's':
		if s, l, ok := ParseStrPtr(arg, checkNull); ok {
			va.StrPtr(s, l)
		} else {
			p.parseTypeError(arg, "string")
		}
	case 'p':
		if s, l, ok := ParsePathStrPtr(arg, checkNull); ok {
			va.StrPtr(s, l)
		} else {
			p.parseTypeError(arg, "a valid path")
		}
	case 'P':
		if val, ok := ParsePathStr(arg, checkNull); ok {
			va.ZStr(val)
		} else {
			p.parseTypeError(arg, "a valid path")
		}
	case 'S':
		if val, ok := ParseZStr(arg, checkNull); ok {
			va.ZStr(val)
		} else {
			p.parseTypeError(arg, "string")
		}
	case 'b':
		if val, isNull, ok := ParseBool(arg, checkNull); ok {
			va.Bool(val)
			if checkNull {
				va.Bool(isNull)
			}
		} else {
			p.parseTypeError(arg, "bool")
		}
	case 'r':
		if res, ok := ParseResource(p.arg, checkNull); ok {
			va.Zval(res)
		} else {
			p.parseTypeError(arg, "resource")
		}
	case 'A', 'a':
		if val, ok := ParseArray(arg, checkNull, typ == 'A'); ok {
			va.Zval(val)
		} else {
			p.parseTypeError(arg, "array")
		}
	case 'H', 'h':
		if ht, ok := ParseArrayHt(arg, checkNull, typ == 'H', separate); ok {
			va.Array(ht)
		} else {
			p.parseTypeError(arg, "array")
		}
	case 'o':
		if obj, ok := ParseObject(arg, nil, check_null != 0); ok {
			va.Zval(obj)
		} else {
			p.parseTypeError(arg, "object")
		}
	case 'O':
		objPtr := va.Pop().(**types.Zval)
		ce := va.Pop().(*zend.ZendClassEntry)
		if obj, ok := ParseObject(arg, ce, check_null != 0); ok {
			*objPtr = obj
		} else {
			if ce != nil {
				p.parseTypeError(arg, ce.Name())
			} else {
				p.parseTypeError(arg, "object")
			}
		}
	case 'C':
		pce := va.Pop().(**zend.ZendClassEntry)
		baseCe := *pce

		// todo 待替换为 ParseClass

		var lookup *zend.ZendClassEntry
		if checkNull && arg.IsNull() {
			*pce = nil
			break
		}
		if zend.TryConvertToString(arg) == 0 {
			*pce = nil
			p.parseTypeError(arg, "valid class name")
			break
		}
		if b.Assign(&lookup, zend.ZendLookupClass(arg.GetStr())) == nil {
			*pce = nil
		} else {
			*pce = lookup
		}
		if baseCe != nil {
			if (*pce) == nil || zend.InstanceofFunction(*pce, baseCe) == 0 {
				p.parseError(0, "to be a class name derived from %s, '%s' given", baseCe.Name(), arg.GetRawStr())
				break
			}
		}
		if (*pce) == nil {
			p.parseError(0, "to be a valid class name, '%s' given", arg.GetRawStr())
			break
		}
	case 'f':
		fci := va.Pop().(*zend.ZendFcallInfo)
		fcc := va.Pop().(*zend.ZendFcallInfoCache)

		err, ok := ParseFunc(arg, fci, fcc, checkNull)
		if !ok {
			if err != nil {
				p.parseError(zend.E_ERROR, "to be a valid callback, %s", err)
			} else {
				p.parseTypeError(arg, "valid callback")
			}
		} else {
			if err != nil {
				p.parseError(zend.E_DEPRECATED, "to be a valid callback, %s", err)
			}
		}
	case 'z':
		zv := ParseZvalDeref(real_arg, checkNull)
		va.Zval(zv)
	case 'Z':
		/* 'Z' iz not supported anymore and should be replaced with 'z' */
		b.Assert(typ != 'Z')
	default:
		p.parseTypeError(arg, "unknown")
	}
}
