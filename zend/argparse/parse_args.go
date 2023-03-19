package argparse

import (
	b "sik/builtin"
	"sik/zend"
	"sik/zend/types"
)

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
				zend.ZendParseParametersDebugError("only one varargs specifier (* or +) is permitted")
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
			zend.ZendParseParametersDebugError("bad type specifier while parsing parameters")
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

func ParseVaArgs(numArgs int, typeSpec string, va []any, flags int) int {
	minNumArgs, maxNumArgs, postVarargs, ok := checkTypeSpec(typeSpec)
	if !ok {
		return types.FAILURE
	}
	if !CheckNumArgsEx(numArgs, zend.CurrEX(), minNumArgs, maxNumArgs, flags) {
		return types.FAILURE
	}
	argCount := zend.CurrEX().NumArgs()
	if numArgs > argCount {
		zend.ZendParseParametersDebugError("could not obtain parameters for parsing")
		return types.FAILURE
	}

	var i int
	var arg *types.Zval
	var varargs **types.Zval = nil
	var nVarargs *int = nil
	i = 0

	r := strReader{typeSpec}

	for b.PostDec(&numArgs) > 0 {
		if r.curr() == '|' {
			r.inc()
		}
		if r.curr() == '*' || r.curr() == '+' {
			var num_varargs int = numArgs + 1 - postVarargs

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
		arg = zend.CurrEX().Arg(i + 1)
		if ZendParseArg(i+1, arg, va, &typeSpec, flags) == types.FAILURE {

			/* clean up varargs array if it was used */

			if varargs != nil && (*varargs) != nil {
				*varargs = nil
			}
			return types.FAILURE
		}
		i++
	}
	return types.SUCCESS
}
