package argparse

import (
	b "sik/builtin"
	"sik/zend"
	"sik/zend/types"
)

func ParseVaArgs(num_args int, type_spec string, va []any, flags int) int {
	var spec_walk *byte
	var c int
	var i int
	var min_num_args int = -1
	var max_num_args int = 0
	var post_varargs int = 0
	var arg *types.Zval
	var arg_count int
	var have_varargs types.ZendBool = 0
	var varargs **types.Zval = nil
	var n_varargs *int = nil
	for spec_walk = type_spec; *spec_walk; spec_walk++ {
		c = *spec_walk
		switch c {
		case 'l':
		case 'd':
		case 's':
		case 'b':
		case 'r':
		case 'a':
		case 'o':
		case 'O':
		case 'z':
		case 'Z':
		case 'C':
		case 'h':
		case 'f':
		case 'A':
		case 'H':
		case 'p':
		case 'S':
		case 'P':
		case 'L':
			max_num_args++
			break
		case '|':
			min_num_args = max_num_args
			break
		case '/':

		case '!':

			/* Pass */

			break
		case '*':

		case '+':
			if have_varargs != 0 {
				zend.ZendParseParametersDebugError("only one varargs specifier (* or +) is permitted")
				return types.FAILURE
			}
			have_varargs = 1

			/* we expect at least one parameter in varargs */

			if c == '+' {
				max_num_args++
			}

			/* mark the beginning of varargs */

			post_varargs = max_num_args
			break
		default:
			zend.ZendParseParametersDebugError("bad type specifier while parsing parameters")
			return types.FAILURE
		}
	}
	if min_num_args < 0 {
		min_num_args = max_num_args
	}
	if have_varargs != 0 {

		/* calculate how many required args are at the end of the specifier list */

		post_varargs = max_num_args - post_varargs
		max_num_args = -1
	}
	if num_args < min_num_args || num_args > max_num_args && max_num_args >= 0 {
		if (flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
			var active_function *zend.ZendFunction = zend.CurrEX().GetFunc()
			var class_name *byte = b.CondF1(active_function.GetScope() != nil, func() []byte { return active_function.GetScope().GetName().GetVal() }, "")
			var throw_exception = zend.CurrEX().IsArgUseStrictTypes() || (flags&zend.ZEND_PARSE_PARAMS_THROW) != 0
			zend.ZendInternalArgumentCountError(throw_exception, "%s%s%s() expects %s %d parameter%s, %d given", class_name, b.Cond(class_name[0], "::", ""), active_function.GetFunctionName().GetVal(), b.Cond(b.Cond(min_num_args == max_num_args, "exactly", num_args < min_num_args), "at least", "at most"), b.Cond(num_args < min_num_args, min_num_args, max_num_args), b.Cond(b.Cond(num_args < min_num_args, min_num_args, max_num_args) == 1, "", "s"), num_args)
		}
		return types.FAILURE
	}
	arg_count = zend.CurrEX().NumArgs()
	if num_args > arg_count {
		zend.ZendParseParametersDebugError("could not obtain parameters for parsing")
		return types.FAILURE
	}
	i = 0
	for b.PostDec(&num_args) > 0 {
		if (*type_spec) == '|' {
			type_spec++
		}
		if (*type_spec) == '*' || (*type_spec) == '+' {
			var num_varargs int = num_args + 1 - post_varargs

			/* eat up the passed in storage even if it won't be filled in with varargs */

			varargs = __va_arg(*va, (**types.Zval)(_))
			n_varargs = __va_arg(*va, (*int)(_))
			type_spec++
			if num_varargs > 0 {
				*n_varargs = num_varargs
				*varargs = zend.CurrEX().Arg(i + 1)

				/* adjust how many args we have left and restart loop */

				num_args += 1 - num_varargs
				i += num_varargs
				continue
			} else {
				*varargs = nil
				*n_varargs = 0
			}
		}
		arg = zend.CurrEX().Arg(i + 1)
		if ZendParseArg(i+1, arg, va, &type_spec, flags) == types.FAILURE {

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
