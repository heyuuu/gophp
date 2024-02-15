package builtin

import (
	"fmt"
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/php/zpp"
	"strings"
)

// func ZifZendVersion() string  { return php.ZEND_VERSION }
func ZifGcMemCaches() int     { return 0 }
func ZifGcCollectCycles() int { return 0 }
func ZifGcEnabled() bool      { return true }

//func ZifGcEnable(ctx *php.Context) {
//	php.ZendAlterIniEntryChars(ctx, "zend.enable_gc", "1", php.ZEND_INI_USER, php.ZEND_INI_STAGE_RUNTIME)
//}
//func ZifGcDisable(ctx *php.Context) {
//	php.ZendAlterIniEntryChars(ctx, "zend.enable_gc", "0", php.ZEND_INI_USER, php.ZEND_INI_STAGE_RUNTIME)
//}

func ZifGcStatus() *types.Array {
	arr := types.NewArrayCap(4)
	arr.AddAssocLong("runs", 0)
	arr.AddAssocLong("collected", 0)
	arr.AddAssocLong("threshold", 0)
	arr.AddAssocLong("roots", 0)
	return arr
}

func ZifFuncNumArgs(ctx *php.Context, executeData *php.ExecuteData) int {
	var ex = executeData.Prev()
	//if (php.ZEND_CALL_INFO(ex) & php.ZEND_CALL_CODE) != 0 {
	//	php.Error(ctx, perr.E_WARNING, "func_num_args():  Called from the global scope - no function context")
	//	return -1
	//}
	if !php.ForbidDynamicCall(ctx, "func_num_args()") {
		return -1
	}
	return ex.NumArgs()
}

// @zif(oldMode="l")
func ZifFuncGetArg(ctx *php.Context, executeData *php.ExecuteData, returnValue zpp.Ret, argNum int) types.Zval {
	if argNum < 0 {
		php.Error(ctx, perr.E_WARNING, "func_get_arg():  The argument number should be >= 0")
		return types.False
	}

	ex := executeData.Prev()
	//if (php.ZEND_CALL_INFO(ex) & php.ZEND_CALL_CODE) != 0 {
	//	php.Error(ctx, perr.E_WARNING, "func_get_arg():  Called from the global scope - no function context")
	//	return types.False
	//}
	if !php.ForbidDynamicCall(ctx, "func_get_arg()") {
		return types.False
	}

	argCount := ex.NumArgs()
	if argNum >= argCount {
		php.Error(ctx, perr.E_WARNING, fmt.Sprintf("func_get_arg():  Argument %d not passed to function", argNum))
		return types.False
	}

	return ex.Arg(argNum + 1)
}
func ZifFuncGetArgs(ctx *php.Context, executeData *php.ExecuteData) (*types.Array, bool) {
	ex := executeData.Prev()
	//if (php.ZEND_CALL_INFO(ex) & php.ZEND_CALL_CODE) != 0 {
	//	php.Error(ctx, perr.E_WARNING, "func_get_args():  Called from the global scope - no function context")
	//	return nil, false
	//}
	if !php.ForbidDynamicCall(ctx, "func_get_args()") {
		return nil, false
	}

	args := ex.Args()
	arr := types.NewArrayCap(len(args))
	for _, arg := range args {
		arg = arg.DeRef()
		if arg.IsUndef() {
			arg = types.Null
		}
		arr.Append(arg)
	}
	return arr, true
}
func ZifStrlen(str string) int { return len(str) }
func ZifStrcmp(str1 string, str2 string) int {
	return php.StringCompare(str1, str2)
}
func ZifStrncmp(ctx *php.Context, str1 string, str2 string, len_ int) (int, bool) {
	if len_ < 0 {
		php.Error(ctx, perr.E_WARNING, "Length must be greater than or equal to 0")
		return 0, false
	}
	if len(str1) > len_ {
		str1 = str1[:len_]
	}
	if len(str2) > len_ {
		str2 = str2[:len_]
	}
	return php.StringCompare(str1, str2), true
}
func ZifStrcasecmp(str1 string, str2 string) int {
	return php.StringCaseCompare(str1, str2)
}
func ZifStrncasecmp(ctx *php.Context, str1 string, str2 string, len_ int) (int, bool) {
	if len_ < 0 {
		php.Error(ctx, perr.E_WARNING, "Length must be greater than or equal to 0")
		return 0, false
	}
	return php.StringNCaseCompare(str1, str2, len_), true
}

// @zif(oldMode="z/")
func ZifEach(ctx *php.Context, array zpp.RefZval) (*types.Array, bool) {
	//if !ctx.EG().EachDeprecationThrown() {
	//	php.Error(ctx, perr.E_DEPRECATED, "The each() function is deprecated. This message will be suppressed on further calls")
	//	ctx.EG().SetEachDeprecationThrown(true)
	//}
	targetHash := php.HashOf(array.Val())
	if targetHash == nil {
		php.Error(ctx, perr.E_WARNING, "Variable passed to each() is not an array or object")
		return nil, false
	}

	pair := targetHash.Current()
	if !pair.IsValid() {
		return nil, false
	}
	key := pair.Key
	val := pair.Val.DeRef()

	result := types.NewArrayCap(4)

	/* add value elements */
	result.IndexAdd(1, val)
	result.KeyAdd(types.STR_VALUE, val)

	/* add the key elements */
	var tmp types.Zval
	if key.IsStrKey() {
		tmp.SetString(key.StrKey())
	} else {
		tmp.SetLong(key.IdxKey())
	}

	result.IndexAdd(0, tmp)
	result.KeyAdd(types.STR_KEY, tmp)

	result.MoveNext()

	return result, true
}

func ZifErrorReporting(ctx *php.Context, ret zpp.Ret, _ zpp.Opt, newErrorLevel *types.Zval) int {
	oldVal := ctx.EG().ErrorReporting()
	if newErrorLevel != nil {
		newVal := php.ZvalGetLong(ctx, *newErrorLevel)
		ctx.EG().SetErrorReporting(newVal)
	}
	return oldVal
}

func validateConstantArray(ctx *php.Context, ht *types.Array) bool {
	ht.ProtectRecursive()
	err := ht.EachEx(func(_ types.ArrayKey, val types.Zval) error {
		val = val.DeRef()
		if val.IsArray() {
			if val.Array().IsRecursive() {
				php.Error(ctx, perr.E_WARNING, "Constants cannot be recursive arrays")
				return lang.BreakErr
			} else if !validateConstantArray(ctx, val.Array()) {
				return lang.BreakErr
			}
		} else if val.Type() > types.IsArray && val.Type() != types.IsResource {
			php.Error(ctx, perr.E_WARNING, "Constants may only evaluate to scalar values, arrays or resources")
			return lang.BreakErr
		}
		return nil
	})
	ht.UnprotectRecursive()
	return err == nil
}

func copyConstantZval(src types.Zval) types.Zval {
	src = src.DeRef()
	if !src.IsArray() {
		return src
	}

	srcArr := src.Array()
	dstArr := types.NewArrayCap(srcArr.Len())
	srcArr.Each(func(key types.ArrayKey, val types.Zval) {
		/* constant arrays can't contain references */
		val = copyConstantZval(val)
		dstArr.Add(key, val)
	})
	return types.ZvalArray(dstArr)
}

func getConstantZval(ctx *php.Context, value types.Zval) (types.Zval, bool) {
	switch value.Type() {
	case types.IsNull,
		types.IsFalse,
		types.IsTrue,
		types.IsLong,
		types.IsDouble,
		types.IsString,
		types.IsResource:
		return value, true
	case types.IsArray:
		if !validateConstantArray(ctx, value.Array()) {
			return types.Undef, false
		} else {
			return copyConstantZval(value), true
		}
	case types.IsObject:
		if retval, ok := value.Object().Cast(types.IsString); ok {
			return retval, true
		}
		fallthrough
	default:
		php.Error(ctx, perr.E_WARNING, "Constants may only evaluate to scalar values, arrays or resources")
		return types.Undef, false
	}
}

func ZifDefine(ctx *php.Context, constantName string, value types.Zval, _ zpp.Opt, caseInsensitive bool) bool {
	var caseSensitive = types.ConstCs
	if caseInsensitive {
		caseSensitive = 0
	}
	if strings.Contains(constantName, "::") {
		php.Error(ctx, perr.E_WARNING, "Class constants cannot be defined or redefined")
		return false
	}

	constValue, ok := getConstantZval(ctx, value)
	if !ok {
		return false
	}

	newValue := copyConstantZval(constValue)
	if caseInsensitive {
		php.Error(ctx, perr.E_DEPRECATED, "define(): Declaration of case-insensitive constants is deprecated")
	}

	return php.RegisterUserConstant(ctx, constantName, newValue, caseSensitive)
}

func ZifDefined(ctx *php.Context, constantName string) bool {
	c := php.GetConstantEx(ctx, constantName, nil, 0)
	return c.IsNotUndef()
}

// @zif(oldMode="|o")
func ZifGetClass(ctx *php.Context, _ zpp.Opt, object *types.Zval) (string, bool) {
	if object == nil {
		var scope = php.GetExecutedScope(ctx)
		if scope != nil {
			return scope.Name(), true
		} else {
			php.Error(ctx, perr.E_WARNING, "get_class() called without object from outside a class")
			return "", false
		}
	}
	return object.Object().ClassName(), true
}

func ZifFunctionExists(ctx *php.Context, functionName string) bool {
	var func_ *types.Function
	var lcname string
	if functionName[0] == '\\' {
		lcname = ascii.StrToLower(functionName[1:])
	} else {
		lcname = ascii.StrToLower(functionName)
	}

	func_ = ctx.EG().FunctionTable().Get(lcname)

	/*
	 * A bit of a hack, but not a bad one: we see if the handler of the function
	 * is actually one that displays "function is disabled" message.
	 */
	return func_ != nil
	//return func_ != nil && !php.IsDisabledFunction(func_)
}

// @zif(oldMode="")
func ZifSetTimeLimit(ctx *php.Context, seconds int) bool {
	return true
}
