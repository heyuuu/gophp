package zpp

import (
	b "github.com/heyuuu/gophp/builtin"
	types2 "github.com/heyuuu/gophp/php/types"
)

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

	parser := FastParseStartEx(numArgs, executeData, minNumArgs, maxNumArgs, flags|FlagOldMode)
	if parser.HasError() {
		return false
	}

	return parseVaArgs(parser, typeSpec, postVarargs, va)
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

func parseVaArgs(p *FastParser, typeSpec string, postVarargs int, args []any) bool {
	var varargs **types2.Zval
	var nVarargs *int
	r := typeSpecReader{typeSpec}
	va := newVaList(args)
	for p.idx = 1; p.idx <= p.numArgs; {
		/* scan through modifiers */
		typ, checkNull, separate := r.Next()
		if typ == '*' || typ == '+' {
			var numVarargs = p.numArgs - (p.idx - 1) - postVarargs

			/* eat up the passed in storage even if it won't be filled in with varargs */
			varargs = va.Pop().(**types2.Zval)
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
			parseArg(p, va, typ, checkNull, separate)
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

// 解析下一位参数，若有错误记录在 p.errorCode 上
func parseArg(p *FastParser, va *vaList, typ byte, checkNull bool, separate bool) {
	switch typ {
	case 'l':
		val, isNull := p.ParseLongEx(checkNull, separate)
		va.Long(val)
		if checkNull {
			va.IntBool(isNull)
		}
	case 'L':
		val, isNull := p.ParseStrictLongEx(checkNull, separate)
		va.Long(val)
		if checkNull {
			va.IntBool(isNull)
		}
	case 'd':
		val, isNull := p.ParseDoubleEx(checkNull, separate)
		va.Double(val)
		if checkNull {
			va.IntBool(isNull)
		}
	case 's':
		strPtr, strLen := p.ParseStringEx(checkNull, separate)
		va.StrPtr(strPtr, strLen)
	case 'p':
		strPtr, strLen := p.ParsePathEx(checkNull, separate)
		va.StrPtr(strPtr, strLen)
	case 'P':
		str := p.ParsePathStrEx(checkNull, separate)
		va.ZStr(str)
	case 'S':
		str := p.ParseStrEx(checkNull, separate)
		va.ZStr(str)
	case 'b':
		val, isNull := p.ParseBoolEx(checkNull, separate)
		va.IntBool(val)
		if checkNull {
			va.IntBool(isNull)
		}
	case 'r':
		res := p.ParseResourceEx(checkNull, separate)
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
		obj := p.ParseObjectEx(checkNull, separate)
		va.Zval(obj)
	case 'O':
		objPtr := va.Pop().(**types2.Zval)
		ce := va.Pop().(*types2.ClassEntry)
		*objPtr = p.ParseObjectOfClassEx(ce, checkNull, separate)
	case 'C':
		pce := va.Pop().(**types2.ClassEntry)
		*pce = p.ParseClassEx(*pce, checkNull, separate)
	case 'f':
		fci := va.Pop().(*types2.ZendFcallInfo)
		fcc := va.Pop().(*types2.ZendFcallInfoCache)
		p.ParseFuncEx(fci, fcc, checkNull, false)
	case 'z':
		zv := p.ParseZvalDerefEx(checkNull, separate)
		va.Zval(zv)
	case 'Z':
		/* 'Z' iz not supported anymore and should be replaced with 'z' */
		b.Assert(typ != 'Z')
	default:
		p.triggerError(ZPP_ERROR_WRONG_ARG, "unknown")
	}
}
