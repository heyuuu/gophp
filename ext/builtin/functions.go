package builtin

import (
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/php/zpp"
	"strings"
)

//func ZifZendVersion() string  { return php.ZEND_VERSION }
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

func ZifStrlen(str string) int               { return len(str) }
func ZifStrcmp(str1 string, str2 string) int { return strings.Compare(str1, str2) }
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
	return strings.Compare(str1, str2), true
}
func ZifStrcasecmp(str1 string, str2 string) int {
	return ascii.StrCaseCompare(str1, str2)
}

//@zif(oldMode="z/")
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
