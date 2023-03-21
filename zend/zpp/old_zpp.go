package zpp

import (
	b "sik/builtin"
	"sik/zend/types"
)

/**
 * 基于类型标识符字符串(Type Spec)的参数解析器
 */
type TypeSpecParser struct {
	FastParser
	typeSpec    string
	va          []any
	postVarargs int
}

func ParseVaArgs(numArgs int, typeSpec string, va []any, flags int) bool {
	minNumArgs, maxNumArgs, postVarargs, ok := checkTypeSpec(typeSpec)
	if !ok {
		return false
	}

	executeData := currExecuteData()
	if numArgs > executeData.NumArgs() {
		ZendParseParametersDebugError("could not obtain parameters for parsing")
		return false
	}

	parser := TypeSpecParseStart(numArgs, executeData, minNumArgs, maxNumArgs, postVarargs, typeSpec, flags)
	if parser == nil || parser.HasError() {
		return false
	}

	return parser.parseVaArgs(va)
}

func TypeSpecParseStart(numArgs int, executeData ExecuteData, minNumArgs int, maxNumArgs int, postVarargs int, typeSpec string, flags int) *TypeSpecParser {
	p := &TypeSpecParser{
		FastParser: FastParser{
			baseParser: makeBaseParser(executeData, numArgs, minNumArgs, maxNumArgs, flags),
		},
		typeSpec:    typeSpec,
		postVarargs: postVarargs,
	}

	// check num args
	p.start()

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
	// todo
}

func (p *TypeSpecParser) realParsePrologue(deref bool, separate bool) {
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
	p.realParsePrologue(deref, separate)

	switch typ {
	case 'l':
		val, isNull := p.ParseLongEx(checkNull)
		va.Long(val)
		if checkNull {
			va.IntBool(isNull)
		}
	case 'L':
		val, isNull := p.ParseStrictLongEx(checkNull)
		va.Long(val)
		if checkNull {
			va.IntBool(isNull)
		}
	case 'd':
		val, isNull := p.ParseDoubleEx(checkNull)
		va.Double(val)
		if checkNull {
			va.IntBool(isNull)
		}
	case 's':
		strPtr, strLen := p.ParseStringEx(checkNull)
		va.StrPtr(strPtr, strLen)
	case 'p':
		strPtr, strLen := p.ParsePathEx(checkNull)
		va.StrPtr(strPtr, strLen)
	case 'P':
		str := p.ParsePathStrEx(checkNull)
		va.ZStr(str)
	case 'S':
		str := p.ParseStrEx(checkNull)
		va.ZStr(str)
	case 'b':
		val, isNull := p.ParseBoolEx(checkNull)
		va.IntBool(val)
		if checkNull {
			va.IntBool(isNull)
		}
	case 'r':
		res := p.ParseResourceEx(checkNull)
		va.Zval(res)
	case 'a':
		val := p.ParseArrayEx(checkNull, separate)
		va.Zval(val)
	case 'A':
		val := p.ParseArrayOrObjectEx(checkNull, separate)
		va.Zval(val)
	case 'h':
		ht := p.ParseArrayHtEx(checkNull, separate)
		va.Array(ht)
	case 'H':
		ht := p.ParseArrayOrObjectHtEx(checkNull, separate)
		va.Array(ht)
	case 'o':
		obj := p.ParseObjectEx(checkNull)
		va.Zval(obj)
	case 'O':
		objPtr := va.Pop().(**types.Zval)
		ce := va.Pop().(*types.ClassEntry)
		*objPtr = p.ParseObjectOfClassEx(ce, checkNull)
	case 'C':
		pce := va.Pop().(**types.ClassEntry)
		*pce = p.ParseClassEx(*pce, checkNull)
	case 'f':
		fci := va.Pop().(*types.ZendFcallInfo)
		fcc := va.Pop().(*types.ZendFcallInfoCache)
		p.ParseFuncEx(fci, fcc, checkNull)
	case 'z':
		zv := p.ParseZvalDerefEx(checkNull)
		va.Zval(zv)
	case 'Z':
		/* 'Z' iz not supported anymore and should be replaced with 'z' */
		b.Assert(typ != 'Z')
	default:
		p.triggerError(ZPP_ERROR_WRONG_ARG, "unknown")
	}
}

func (p *TypeSpecParser) ParseZvalDerefEx(checkNull bool) (dest *types.Zval) {
	p.parsePrologue(true, false)
	if p.IsFinish() {
		return
	}

	return ParseZvalDeref(p.arg, checkNull)
}
