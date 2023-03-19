package argparse

import (
	"fmt"
	b "sik/builtin"
	"sik/zend"
	"sik/zend/faults"
	"sik/zend/types"
)

type TypeSpecParser struct {
	baseParser
	typeSpec    string
	va          []any
	postVarargs int
	err         *parseArgError
	vaReceiver  *VaArgsReceiver
	spec        *typeSpecReader
}

func TypeSpecParseStart(numArgs int, typeSpec string, va []any, flags int) *TypeSpecParser {
	minNumArgs, maxNumArgs, postVarargs, ok := checkTypeSpec(typeSpec)
	if !ok {
		return nil
	}
	executeData := currExecuteData()
	if numArgs > executeData.NumArgs() {
		ZendParseParametersDebugError("could not obtain parameters for parsing")
		return nil
	}

	p := &TypeSpecParser{
		baseParser:  makeBaseParser(executeData, numArgs, minNumArgs, maxNumArgs, flags),
		typeSpec:    typeSpec,
		va:          va,
		postVarargs: postVarargs,
	}

	p.start()

	return p
}

func (p *TypeSpecParser) parseVaArgs(va []any) bool {
	numArgs := p.numArgs
	typeSpec := p.typeSpec

	var i int
	var varargs **types.Zval = nil
	var nVarargs *int = nil
	i = 0
	r := strReader{typeSpec}
	for b.PostDec(&numArgs) > 0 {
		if r.curr() == '|' {
			r.inc()
		}
		if r.curr() == '*' || r.curr() == '+' {
			var num_varargs int = numArgs + 1 - p.postVarargs

			/* eat up the passed in storage even if it won't be filled in with varargs */
			varargs = vaArg[*types.Zval](&va)
			nVarargs = vaArg[int](&va)
			r.inc()

			if num_varargs > 0 {
				*nVarargs = num_varargs
				*varargs = zend.CurrEX().Arg(i + 1)

				/* adjust how many args we have left and restart loop */
				numArgs += 1 - num_varargs
				i += num_varargs
				continue
			} else {
				*varargs = nil
				*nVarargs = 0
			}
		}

		p.idx = i + 1
		p.arg = p.currArg()
		if !p.parseArg() {
			/* clean up varargs array if it was used */
			if varargs != nil && (*varargs) != nil {
				*varargs = nil
			}
			return false
		}
		i++
	}
	return true
}

func (p *TypeSpecParser) parseArg() bool {
	p.ZendParseArgImpl()
	err := p.err
	if err != nil {
		if zend.EG__().GetException() != nil {
			return false
		}
		if !p.isQuiet() {
			message := fmt.Sprintf("%s() expects parameter %d %s", zend.GetActiveCalleeName(), p.idx, err.Error())
			faults.InternalTypeErrorEx(p.isThrowEx(), message)
		}
		if err.severity != faults.E_DEPRECATED {
			return false
		}
	}
	if p.HasError() {
		return false
	}
	return true
}

func (p *TypeSpecParser) ZendParseArgImpl() {
	arg := p.arg
	va := p.vaReceiver
	spec := p.spec
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
			p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_STRING)
		}
	case 'p':
		if s, l, ok := ParsePathStrPtr(arg, checkNull); ok {
			va.StrPtr(s, l)
		} else {
			p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_PATH)
		}
	case 'P':
		if val, ok := ParsePathStr(arg, checkNull); ok {
			va.ZStr(val)
		} else {
			p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_PATH)
		}
	case 'S':
		if val, ok := ParseZStr(arg, checkNull); ok {
			va.ZStr(val)
		} else {
			p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_STRING)
		}
	case 'b':
		if val, isNull, ok := ParseBool(arg, checkNull); ok {
			va.Bool(val)
			if checkNull {
				va.Bool(isNull)
			}
		} else {
			p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_BOOL)
		}
	case 'r':
		if res, ok := ParseResource(p.arg, checkNull); ok {
			va.Zval(res)
		} else {
			p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_RESOURCE)
		}
	case 'A', 'a':
		if val, ok := ParseArray(arg, checkNull, typ == 'A'); ok {
			va.Zval(val)
		} else {
			p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_ARRAY)
		}
	case 'H', 'h':
		if ht, ok := ParseArrayHt(arg, checkNull, typ == 'H', separate); ok {
			va.Array(ht)
		} else {
			p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_ARRAY)
		}
	case 'o':
		if obj, ok := ParseObject(arg, nil, check_null != 0); ok {
			va.Zval(obj)
		} else {
			p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_OBJECT)
		}
	case 'O':
		objPtr := va.Pop().(**types.Zval)
		ce := va.Pop().(*types.ClassEntry)
		if obj, ok := ParseObject(arg, ce, check_null != 0); ok {
			*objPtr = obj
		} else {
			if ce != nil {
				p.triggerError(ZPP_ERROR_WRONG_CLASS, ce.Name())
			} else {
				p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_OBJECT)
			}
		}
	case 'C':
		pce := va.Pop().(**types.ClassEntry)
		baseCe := *pce

		// todo 待替换为 ParseClass

		var lookup *types.ClassEntry
		if checkNull && arg.IsNull() {
			*pce = nil
			break
		}
		if zend.TryConvertToString(arg) == 0 {
			*pce = nil
			p.triggerError(ZPP_ERROR_WRONG_ARG, "valid class name")
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
				p.parseError(faults.E_ERROR, "to be a valid callback, %s", err)
			} else {
				p.triggerError(ZPP_ERROR_WRONG_ARG, "valid callback")
			}
		} else {
			if err != nil {
				p.parseError(faults.E_DEPRECATED, "to be a valid callback, %s", err)
			}
		}
	case 'z':
		zv := ParseZvalDeref(real_arg, checkNull)
		va.Zval(zv)
	case 'Z':
		/* 'Z' iz not supported anymore and should be replaced with 'z' */
		b.Assert(typ != 'Z')
	default:
		p.triggerError(ZPP_ERROR_WRONG_ARG, "unknown")
	}
}
