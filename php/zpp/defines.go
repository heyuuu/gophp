package zpp

import (
	"github.com/heyuuu/gophp/php/types"
)

/**
 * 此处的类型用于 sikgen 脚本生成代码
 */
type (
	/* special */
	Ret = *types.Zval

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

	// FAST_ZPP: Z_PARAM_ARRAY_HT, Type: 'h'，直接使用 *types.Array
	//Array = *types.Array
	ArrayNullable = *types.Array

	// FAST_ZPP: Z_PARAM_ARRAY_OR_OBJECT_HT, Type: 'H'
	ArrayOrObjectHt = *types.Array

	// FAST_ZPP: Z_PARAM_ARRAY_OR_OBJECT, Type: 'A'
	ArrayOrObjectZval = types.Zval

	// FAST_ZPP: Z_PARAM_CLASS, Type: 'C'
	Class = *types.Class

	// *types.Object 直接使用 *types.Object
	// Object = *types.Object
	ObjectNullable = *types.Object

	// FAST_ZPP: Z_PARAM_RESOURCE, Type: 'r'
	Resource         = types.Zval
	ResourceNullable = *types.Zval

	// FAST_ZPP: Z_PARAM_FUNC，Type: 'f' (参数略有差异，使用封装的结构体代替 fci + fcc 双指针)
	Callable = *types.UserCallable

	// FAST_ZPP: Z_PARAM_ZVAL, Type: 'z', 直接使用 *types.Zval
	//Zval = *types.Zval
	ZvalNullable = *types.Zval

	// FAST_ZPP: Z_PARAM_ZVAL_DEREF, Type: ''
	ZvalDeref         = *types.Zval
	ZvalDerefNullable = *types.Zval

	// FAST_ZPP: Z_PARAM_VARIADIC, Type: '*' or '+', 直接使用 []*types.Zval
	//Variadic = []*types.Zval

	/* ref type */
	RefZval          = types.RefZval
	RefZvalNullable  = types.RefZval
	RefArrayOrObject = types.RefZval
	RefArray         = *types.Array
	RefArrayNullable = *types.Array
)
