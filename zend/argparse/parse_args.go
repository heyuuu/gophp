package argparse

import (
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

func ParseVaArgs(numArgs int, typeSpec string, va []any, flags int) int {
	parser := TypeSpecParseStart(numArgs, typeSpec, va, flags)
	if parser == nil || parser.HasError() {
		return types.FAILURE
	}

	return parser.parseVaArgs(numArgs, typeSpec, va, flags)
}
