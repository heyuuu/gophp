package zpp

import (
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/types"
)

/**
 * 此处的类型用于 sikgen 脚本生成代码
 */
type (
	/* special */
	Ex  = *php.ExecuteData
	Ret = *types.Zval
	Ctx = *php.Context

	// FAST_ZPP: Z_PARAM_OPTIONAL
	Opt any

	/* base type */

	// FAST_ZPP: Z_PARAM_BOOL, Type: 'b'，直接使用 bool (略有差异，使用 bool 类型而非 ZendBool 类型)
	//Bool = bool

	// FAST_ZPP: Z_PARAM_LONG, Type: 'l'，直接使用 int
	//Long = int

	// FAST_ZPP: Z_PARAM_STRICT_LONG, Type: 'L'
	StrictLong = int

	// FAST_ZPP: Z_PARAM_DOUBLE, Type: 'd'，直接使用 double
	//Double = float64

	// FAST_ZPP: Z_PARAM_STRING, Type: 's', 直接使用 string (略有差异，使用 string 类型而非 *byte+len 类型爽字段)
	//String = string

	// FAST_ZPP: Z_PARAM_PATH, Type: 'p' (略有差异，使用 string 类型而非 *byte+len 类型爽字段)
	Path = string

	// FAST_ZPP: Z_PARAM_ARRAY_HT, Type: 'h'，直接使用	 *types.Array
	//ArrayHt = *types.Array

	// FAST_ZPP: Z_PARAM_ARRAY_OR_OBJECT_HT, Type: 'H'
	ArrayOrObjectHt = *types.Array

	// FAST_ZPP: Z_PARAM_ARRAY, Type: 'a'
	Array = *types.Zval

	// FAST_ZPP: Z_PARAM_ARRAY_OR_OBJECT, Type: 'A'
	ArrayOrObject = *types.Zval

	// FAST_ZPP: Z_PARAM_CLASS, Type: 'C'
	Class = *types.Class

	// FAST_ZPP: Z_PARAM_OBJECT, Type: 'o'
	Object = *types.Zval

	// FAST_ZPP: Z_PARAM_RESOURCE, Type: 'r'
	Resource = *types.Zval

	// FAST_ZPP: Z_PARAM_FUNC，Type: 'f' (参数略有差异，使用封装的结构体代替 fci + fcc 双指针)
	//Callable = *types.UserCallable

	// FAST_ZPP: Z_PARAM_ZVAL, Type: 'z', 直接使用 *types.Zval
	//Zval = *types.Zval

	// FAST_ZPP: Z_PARAM_ZVAL_DEREF, Type: ''
	ZvalDeref = *types.Zval

	// FAST_ZPP: Z_PARAM_VARIADIC, Type: '*' or '+', 直接使用 []*types.Zval
	//Variadic = []*types.Zval

	/* CheckNull */
	ZvalNullable     = *types.Zval // fp.ParseZvalEx(true, false)
	ArrayNullable    = *types.Zval // fp.ParseArrayEx(true, false)
	ResourceNullable = *types.Zval // fp.ParseResourceEx(true, false)

	/* ref type */
	RefZval          = *types.Zval //
	RefArray         = *types.Zval // fp.ParseArrayEx(false, true)
	DerefArray       = *types.Zval // fp.ParseArrayEx2(false, true, false)
	RefArrayOrObject = *types.Zval // fp.ParseArrayOrObjectEx(false, true)

	RefArrayHt = *types.Array // fp.ParseArrayHtEx(false, true)
)
