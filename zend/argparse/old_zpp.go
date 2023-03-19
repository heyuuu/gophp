package argparse

import (
	b "sik/builtin"
	"sik/zend"
	"sik/zend/types"
)

/**
 * 基于类型标识符字符串(Type Spec)的参数解析器
 */
type TypeSpecParser struct {
	baseParser
	typeSpec    string
	va          []any
	postVarargs int
}

func ParseVaArgs(numArgs int, typeSpec string, va []any, flags int) int {
	parser := TypeSpecParseStart(numArgs, typeSpec, flags)
	if parser == nil || parser.HasError() {
		return types.FAILURE
	}

	ok := parser.parseVaArgs(va)
	return types.IntBool(ok)
}

func TypeSpecParseStart(numArgs int, typeSpec string, flags int) *TypeSpecParser {
	minNumArgs, maxNumArgs, postVarargs, ok := checkTypeSpec(typeSpec)
	if !ok {
		return nil
	}
	executeData := currExecuteData()

	p := &TypeSpecParser{
		baseParser:  makeBaseParser(executeData, numArgs, minNumArgs, maxNumArgs, flags),
		typeSpec:    typeSpec,
		postVarargs: postVarargs,
	}

	p.start()
	if numArgs > executeData.NumArgs() {
		ZendParseParametersDebugError("could not obtain parameters for parsing")
		return nil
	}

	return p
}

func checkTypeSpec(typeSpec string) (minNumArgs int, maxNumArgs int, postVarargs int, ok bool) {
	minNumArgs = -1
	maxNumArgs = 0
	postVarargs = 0
	haveVarargs := false
	for _, c := range []byte(typeSpec) {
		switch c {
		case 'l', 'd', 's', 'b', 'r', 'a', 'o', 'O', 'z', 'Z', 'C', 'h', 'f', 'A', 'H', 'p', 'S', 'P', 'L':
			maxNumArgs++
		case '|':
			minNumArgs = maxNumArgs
		case '/', '!':
			/* Pass */
		case '*', '+':
			if haveVarargs {
				ZendParseParametersDebugError("only one varargs specifier (* or +) is permitted")
				return
			}
			haveVarargs = true

			/* we expect at least one parameter in varargs */
			if c == '+' {
				maxNumArgs++
			}

			/* mark the beginning of varargs */
			postVarargs = maxNumArgs
		default:
			ZendParseParametersDebugError("bad type specifier while parsing parameters")
			return
		}
	}
	if minNumArgs < 0 {
		minNumArgs = maxNumArgs
	}
	if haveVarargs {
		/* calculate how many required args are at the end of the specifier list */
		postVarargs = maxNumArgs - postVarargs
		maxNumArgs = -1
	}

	ok = true
	return
}

func (p *TypeSpecParser) parseVaArgs(args []any) bool {
	var varargs **types.Zval
	var nVarargs *int
	r := typeSpecReader{p.typeSpec}
	va := newVaList(args)
	for p.idx = 1; p.idx <= p.numArgs; {
		/* scan through modifiers */
		typ, checkNull, separate := r.Next()
		if typ == '*' || typ == '+' {
			var numVarargs = p.numArgs - (p.idx - 1) - p.postVarargs

			/* eat up the passed in storage even if it won't be filled in with varargs */
			varargs = va.Pop().(**types.Zval)
			nVarargs = va.Pop().(*int)

			if numVarargs > 0 {
				*nVarargs = numVarargs
				*varargs = p.currArg()

				/* adjust how many args we have left and restart loop */
				p.idx += numVarargs
			} else {
				*varargs = nil
				*nVarargs = 0
			}
		} else {
			p.parseArg(va, typ, checkNull, separate)
			p.idx++

			if p.HasError() {
				/* clean up varargs array if it was used */
				if varargs != nil && (*varargs) != nil {
					*varargs = nil
				}
				return false
			}
		}
	}
	return true
}

func (p *TypeSpecParser) parsePrologue(deref bool, separate bool) {
	p.arg = p.currArg()
	if deref {
		p.arg = types.ZVAL_DEREF(p.arg)
	}
	if separate {
		types.SEPARATE_ZVAL_NOREF(p.arg)
	}
}

// 解析下一位参数，若有错误记录在 p.errorCode 上
func (p *TypeSpecParser) parseArg(va *vaList, typ byte, checkNull bool, separate bool) {
	deref := typ != 'z' || separate
	p.parsePrologue(deref, separate)

	switch typ {
	case 'l', 'L':
		if val, isNull, ok := ParseLong(p.arg, checkNull, typ == 'L'); ok {
			va.Long(val)
			if checkNull {
				va.Bool(isNull)
			}
		} else {
			p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_LONG)
		}
	case 'd':
		if val, isNull, ok := ParseDouble(p.arg, checkNull); ok {
			va.Double(val)
			if checkNull {
				va.Bool(isNull)
			}
		} else {
			p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_DOUBLE)
		}
	case 's':
		if s, l, ok := ParseStrPtr(p.arg, checkNull); ok {
			va.StrPtr(s, l)
		} else {
			p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_STRING)
		}
	case 'p':
		if s, l, ok := ParsePathStrPtr(p.arg, checkNull); ok {
			va.StrPtr(s, l)
		} else {
			p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_PATH)
		}
	case 'P':
		if val, ok := ParsePathStr(p.arg, checkNull); ok {
			va.ZStr(val)
		} else {
			p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_PATH)
		}
	case 'S':
		if val, ok := ParseZStr(p.arg, checkNull); ok {
			va.ZStr(val)
		} else {
			p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_STRING)
		}
	case 'b':
		if val, isNull, ok := ParseBool(p.arg, checkNull); ok {
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
		if val, ok := ParseArray(p.arg, checkNull, typ == 'A'); ok {
			va.Zval(val)
		} else {
			p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_ARRAY)
		}
	case 'H', 'h':
		if ht, ok := ParseArrayHt(p.arg, checkNull, typ == 'H', separate); ok {
			va.Array(ht)
		} else {
			p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_ARRAY)
		}
	case 'o':
		if obj, ok := ParseObject(p.arg, nil, checkNull); ok {
			va.Zval(obj)
		} else {
			p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_OBJECT)
		}
	case 'O':
		objPtr := va.Pop().(**types.Zval)
		ce := va.Pop().(*types.ClassEntry)
		if obj, ok := ParseObject(p.arg, ce, checkNull); ok {
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
		if ce, ok := ParseClass(p.arg, baseCe, p.idx, checkNull); ok {
			*pce = ce
		} else {
			p.triggerError(ZPP_ERROR_FAILURE, "")
		}
	case 'f':
		fci := va.Pop().(*zend.ZendFcallInfo)
		fcc := va.Pop().(*zend.ZendFcallInfoCache)
		err, ok := ParseFunc(p.arg, fci, fcc, checkNull)
		if !ok {
			if err != nil {
				p.triggerError(ZPP_ERROR_WRONG_CALLBACK, *err)
			} else {
				p.triggerError(ZPP_ERROR_WRONG_ARG, "valid callback")
			}
		} else if err != nil {
			p.triggerDeprecated(ZPP_ERROR_WRONG_CALLBACK, *err)
		}
	case 'z':
		zv := ParseZvalDeref(p.arg, checkNull)
		va.Zval(zv)
	case 'Z':
		/* 'Z' iz not supported anymore and should be replaced with 'z' */
		b.Assert(typ != 'Z')
	default:
		p.triggerError(ZPP_ERROR_WRONG_ARG, "unknown")
	}
}
