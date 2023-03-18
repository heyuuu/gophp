// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/zend"
	"sik/zend/argparse"
	"sik/zend/types"
)

func PhpArrayInitGlobals(array_globals *ZendArrayGlobals) {
	memset(array_globals, 0, b.SizeOf("zend_array_globals"))
}
func ZmStartupArray(type_ int, module_number int) int {
	PhpArrayInitGlobals(&ArrayGlobals)
	zend.REGISTER_LONG_CONSTANT("EXTR_OVERWRITE", EXTR_OVERWRITE, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("EXTR_SKIP", EXTR_SKIP, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("EXTR_PREFIX_SAME", EXTR_PREFIX_SAME, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("EXTR_PREFIX_ALL", EXTR_PREFIX_ALL, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("EXTR_PREFIX_INVALID", EXTR_PREFIX_INVALID, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("EXTR_PREFIX_IF_EXISTS", EXTR_PREFIX_IF_EXISTS, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("EXTR_IF_EXISTS", EXTR_IF_EXISTS, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("EXTR_REFS", EXTR_REFS, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("SORT_ASC", PHP_SORT_ASC, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("SORT_DESC", PHP_SORT_DESC, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("SORT_REGULAR", PHP_SORT_REGULAR, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("SORT_NUMERIC", PHP_SORT_NUMERIC, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("SORT_STRING", PHP_SORT_STRING, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("SORT_LOCALE_STRING", PHP_SORT_LOCALE_STRING, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("SORT_NATURAL", PHP_SORT_NATURAL, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("SORT_FLAG_CASE", PHP_SORT_FLAG_CASE, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("CASE_LOWER", CASE_LOWER, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("CASE_UPPER", CASE_UPPER, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("COUNT_NORMAL", COUNT_NORMAL, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("COUNT_RECURSIVE", COUNT_RECURSIVE, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("ARRAY_FILTER_USE_BOTH", ARRAY_FILTER_USE_BOTH, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("ARRAY_FILTER_USE_KEY", ARRAY_FILTER_USE_KEY, zend.CONST_CS|zend.CONST_PERSISTENT)
	return types.SUCCESS
}
func ZmShutdownArray(type_ int, module_number int) int { return types.SUCCESS }
func PhpArrayKeyCompare(a any, b any) int {
	var f *types.Bucket = (*types.Bucket)(a)
	var s *types.Bucket = (*types.Bucket)(b)
	var t types.ZendUchar
	var l1 zend.ZendLong
	var l2 zend.ZendLong
	var d float64
	if f.GetKey() == nil {
		if s.GetKey() == nil {
			return zend.ZendLong(f.GetH() > zend.ZendLong(b.Cond(s.GetH() != 0, 1, -1)))
		} else {
			l1 = zend.ZendLong(f.GetH())
			t = zend.IsNumericString(s.GetKey().GetStr(), &l2, &d, 1)
			if t == types.IS_LONG {

			} else if t == types.IS_DOUBLE {
				return zend.ZEND_NORMALIZE_BOOL(float64(l1 - d))
			} else {
				l2 = 0
			}
		}
	} else {
		if s.GetKey() != nil {
			return zend.ZendiSmartStrcmp(f.GetKey(), s.GetKey())
		} else {
			l2 = zend.ZendLong(s.GetH())
			t = zend.IsNumericString(f.GetKey().GetStr(), &l1, &d, 1)
			if t == types.IS_LONG {

			} else if t == types.IS_DOUBLE {
				return zend.ZEND_NORMALIZE_BOOL(d - float64(l2))
			} else {
				l1 = 0
			}
		}
	}
	return zend.ZEND_NORMALIZE_BOOL(l1 - l2)
}
func PhpArrayReverseKeyCompare(a any, b any) int { return PhpArrayKeyCompare(b, a) }
func PhpArrayKeyCompareNumeric(a any, b any) int {
	var f *types.Bucket = (*types.Bucket)(a)
	var s *types.Bucket = (*types.Bucket)(b)
	if f.GetKey() == nil && s.GetKey() == nil {
		return zend.ZendLong(f.GetH() > zend.ZendLong(b.Cond(s.GetH() != 0, 1, -1)))
	} else {
		var d1 float64
		var d2 float64
		if f.GetKey() != nil {
			d1 = zend.ZendStrtod(f.GetKey().GetVal(), nil)
		} else {
			d1 = float64(zend.ZendLong(f.GetH()))
		}
		if s.GetKey() != nil {
			d2 = zend.ZendStrtod(s.GetKey().GetVal(), nil)
		} else {
			d2 = float64(zend.ZendLong(s.GetH()))
		}
		return zend.ZEND_NORMALIZE_BOOL(d1 - d2)
	}
}
func PhpArrayReverseKeyCompareNumeric(a any, b any) int { return PhpArrayKeyCompareNumeric(b, a) }
func PhpArrayKeyCompareStringCase(a any, b any) int {
	var f *types.Bucket = (*types.Bucket)(a)
	var s *types.Bucket = (*types.Bucket)(b)
	var s1 *byte
	var s2 *byte
	var l1 int
	var l2 int
	var buf1 []byte
	var buf2 []byte
	if f.GetKey() != nil {
		s1 = f.GetKey().GetVal()
		l1 = f.GetKey().GetLen()
	} else {
		s1 = zend.ZendPrintLongToBuf(buf1+b.SizeOf("buf1")-1, f.GetH())
		l1 = buf1 + b.SizeOf("buf1") - 1 - s1
	}
	if s.GetKey() != nil {
		s2 = s.GetKey().GetVal()
		l2 = s.GetKey().GetLen()
	} else {
		s2 = zend.ZendPrintLongToBuf(buf2+b.SizeOf("buf2")-1, s.GetH())
		l2 = buf2 + b.SizeOf("buf2") - 1 - s1
	}
	return zend.ZendBinaryStrcasecmpL(s1, l1, s2, l2)
}
func PhpArrayReverseKeyCompareStringCase(a any, b any) int {
	return PhpArrayKeyCompareStringCase(b, a)
}
func PhpArrayKeyCompareString(a any, b any) int {
	var f *types.Bucket = (*types.Bucket)(a)
	var s *types.Bucket = (*types.Bucket)(b)
	var s1 *byte
	var s2 *byte
	var l1 int
	var l2 int
	var buf1 []byte
	var buf2 []byte
	if f.GetKey() != nil {
		s1 = f.GetKey().GetVal()
		l1 = f.GetKey().GetLen()
	} else {
		s1 = zend.ZendPrintLongToBuf(buf1+b.SizeOf("buf1")-1, f.GetH())
		l1 = buf1 + b.SizeOf("buf1") - 1 - s1
	}
	if s.GetKey() != nil {
		s2 = s.GetKey().GetVal()
		l2 = s.GetKey().GetLen()
	} else {
		s2 = zend.ZendPrintLongToBuf(buf2+b.SizeOf("buf2")-1, s.GetH())
		l2 = buf2 + b.SizeOf("buf2") - 1 - s2
	}
	return zend.ZendBinaryStrcmp(s1, l1, s2, l2)
}
func PhpArrayReverseKeyCompareString(a any, b any) int { return PhpArrayKeyCompareString(b, a) }
func PhpArrayKeyCompareStringNaturalGeneral(a any, b any, fold_case int) int {
	var f *types.Bucket = (*types.Bucket)(a)
	var s *types.Bucket = (*types.Bucket)(b)
	var s1 *byte
	var s2 *byte
	var l1 int
	var l2 int
	var buf1 []byte
	var buf2 []byte
	if f.GetKey() != nil {
		s1 = f.GetKey().GetVal()
		l1 = f.GetKey().GetLen()
	} else {
		s1 = zend.ZendPrintLongToBuf(buf1+b.SizeOf("buf1")-1, f.GetH())
		l1 = buf1 + b.SizeOf("buf1") - 1 - s1
	}
	if s.GetKey() != nil {
		s2 = s.GetKey().GetVal()
		l2 = s.GetKey().GetLen()
	} else {
		s2 = zend.ZendPrintLongToBuf(buf2+b.SizeOf("buf2")-1, s.GetH())
		l2 = buf2 + b.SizeOf("buf2") - 1 - s1
	}
	return StrnatcmpEx(s1, l1, s2, l2, fold_case)
}
func PhpArrayKeyCompareStringNaturalCase(a any, b any) int {
	return PhpArrayKeyCompareStringNaturalGeneral(a, b, 1)
}
func PhpArrayReverseKeyCompareStringNaturalCase(a any, b any) int {
	return PhpArrayKeyCompareStringNaturalGeneral(b, a, 1)
}
func PhpArrayKeyCompareStringNatural(a any, b any) int {
	return PhpArrayKeyCompareStringNaturalGeneral(a, b, 0)
}
func PhpArrayReverseKeyCompareStringNatural(a any, b any) int {
	return PhpArrayKeyCompareStringNaturalGeneral(b, a, 0)
}
func PhpArrayKeyCompareStringLocale(a any, b any) int {
	var f *types.Bucket = (*types.Bucket)(a)
	var s *types.Bucket = (*types.Bucket)(b)
	var s1 *byte
	var s2 *byte
	var buf1 []byte
	var buf2 []byte
	if f.GetKey() != nil {
		s1 = f.GetKey().GetVal()
	} else {
		s1 = zend.ZendPrintLongToBuf(buf1+b.SizeOf("buf1")-1, f.GetH())
	}
	if s.GetKey() != nil {
		s2 = s.GetKey().GetVal()
	} else {
		s2 = zend.ZendPrintLongToBuf(buf2+b.SizeOf("buf2")-1, s.GetH())
	}
	return strcoll(s1, s2)
}
func PhpArrayReverseKeyCompareStringLocale(a any, b any) int {
	return PhpArrayKeyCompareStringLocale(b, a)
}
func PhpArrayDataCompare(a any, b any) int {
	var f *types.Bucket
	var s *types.Bucket
	var result types.Zval
	var first *types.Zval
	var second *types.Zval
	f = (*types.Bucket)(a)
	s = (*types.Bucket)(b)
	first = f.GetVal()
	second = s.GetVal()
	if first.IsType(types.IS_INDIRECT) {
		first = first.GetZv()
	}
	if second.IsType(types.IS_INDIRECT) {
		second = second.GetZv()
	}
	if zend.CompareFunction(&result, first, second) == types.FAILURE {
		return 0
	}
	b.Assert(result.IsType(types.IS_LONG))
	return zend.ZEND_NORMALIZE_BOOL(result.GetLval())
}
func PhpArrayReverseDataCompare(a any, b any) int { return PhpArrayDataCompare(a, b) * -1 }
func PhpArrayDataCompareNumeric(a any, b any) int {
	var f *types.Bucket
	var s *types.Bucket
	var first *types.Zval
	var second *types.Zval
	f = (*types.Bucket)(a)
	s = (*types.Bucket)(b)
	first = f.GetVal()
	second = s.GetVal()
	if first.IsType(types.IS_INDIRECT) {
		first = first.GetZv()
	}
	if second.IsType(types.IS_INDIRECT) {
		second = second.GetZv()
	}
	return zend.NumericCompareFunction(first, second)
}
func PhpArrayReverseDataCompareNumeric(a any, b any) int { return PhpArrayDataCompareNumeric(b, a) }
func PhpArrayDataCompareStringCase(a any, b any) int {
	var f *types.Bucket
	var s *types.Bucket
	var first *types.Zval
	var second *types.Zval
	f = (*types.Bucket)(a)
	s = (*types.Bucket)(b)
	first = f.GetVal()
	second = s.GetVal()
	if first.IsType(types.IS_INDIRECT) {
		first = first.GetZv()
	}
	if second.IsType(types.IS_INDIRECT) {
		second = second.GetZv()
	}
	return zend.StringCaseCompareFunction(first, second)
}
func PhpArrayReverseDataCompareStringCase(a any, b any) int {
	return PhpArrayDataCompareStringCase(b, a)
}
func PhpArrayDataCompareString(a any, b any) int {
	var f *types.Bucket
	var s *types.Bucket
	var first *types.Zval
	var second *types.Zval
	f = (*types.Bucket)(a)
	s = (*types.Bucket)(b)
	first = f.GetVal()
	second = s.GetVal()
	if first.IsType(types.IS_INDIRECT) {
		first = first.GetZv()
	}
	if second.IsType(types.IS_INDIRECT) {
		second = second.GetZv()
	}
	return zend.StringCompareFunction(first, second)
}
func PhpArrayReverseDataCompareString(a any, b any) int { return PhpArrayDataCompareString(b, a) }
func PhpArrayNaturalGeneralCompare(a any, b any, fold_case int) int {
	var f *types.Bucket = (*types.Bucket)(a)
	var s *types.Bucket = (*types.Bucket)(b)
	var tmp_str1 *types.ZendString
	var tmp_str2 *types.ZendString
	var str1 *types.ZendString = zend.ZvalGetTmpString(f.GetVal(), &tmp_str1)
	var str2 *types.ZendString = zend.ZvalGetTmpString(s.GetVal(), &tmp_str2)
	var result int = StrnatcmpEx(str1.GetVal(), str1.GetLen(), str2.GetVal(), str2.GetLen(), fold_case)
	zend.ZendTmpStringRelease(tmp_str1)
	zend.ZendTmpStringRelease(tmp_str2)
	return result
}
func PhpArrayNaturalCompare(a any, b any) int {
	return PhpArrayNaturalGeneralCompare(a, b, 0)
}
func PhpArrayReverseNaturalCompare(a any, b any) int {
	return PhpArrayNaturalGeneralCompare(b, a, 0)
}
func PhpArrayNaturalCaseCompare(a any, b any) int {
	return PhpArrayNaturalGeneralCompare(a, b, 1)
}
func PhpArrayReverseNaturalCaseCompare(a any, b any) int {
	return PhpArrayNaturalGeneralCompare(b, a, 1)
}
func PhpArrayDataCompareStringLocale(a any, b any) int {
	var f *types.Bucket
	var s *types.Bucket
	var first *types.Zval
	var second *types.Zval
	f = (*types.Bucket)(a)
	s = (*types.Bucket)(b)
	first = f.GetVal()
	second = s.GetVal()
	if first.IsType(types.IS_INDIRECT) {
		first = first.GetZv()
	}
	if second.IsType(types.IS_INDIRECT) {
		second = second.GetZv()
	}
	return zend.StringLocaleCompareFunction(first, second)
}
func PhpArrayReverseDataCompareStringLocale(a any, b any) int {
	return PhpArrayDataCompareStringLocale(b, a)
}
func PhpGetKeyCompareFunc(sort_type zend.ZendLong, reverse int) types.CompareFuncT {
	switch sort_type & ^PHP_SORT_FLAG_CASE {
	case PHP_SORT_NUMERIC:
		if reverse != 0 {
			return PhpArrayReverseKeyCompareNumeric
		} else {
			return PhpArrayKeyCompareNumeric
		}
	case PHP_SORT_STRING:
		if (sort_type & PHP_SORT_FLAG_CASE) != 0 {
			if reverse != 0 {
				return PhpArrayReverseKeyCompareStringCase
			} else {
				return PhpArrayKeyCompareStringCase
			}
		} else {
			if reverse != 0 {
				return PhpArrayReverseKeyCompareString
			} else {
				return PhpArrayKeyCompareString
			}
		}
	case PHP_SORT_NATURAL:
		if (sort_type & PHP_SORT_FLAG_CASE) != 0 {
			if reverse != 0 {
				return PhpArrayReverseKeyCompareStringNaturalCase
			} else {
				return PhpArrayKeyCompareStringNaturalCase
			}
		} else {
			if reverse != 0 {
				return PhpArrayReverseKeyCompareStringNatural
			} else {
				return PhpArrayKeyCompareStringNatural
			}
		}
	case PHP_SORT_LOCALE_STRING:
		if reverse != 0 {
			return PhpArrayReverseKeyCompareStringLocale
		} else {
			return PhpArrayKeyCompareStringLocale
		}
	case PHP_SORT_REGULAR:
		fallthrough
	default:
		if reverse != 0 {
			return PhpArrayReverseKeyCompare
		} else {
			return PhpArrayKeyCompare
		}
	}
	return nil
}
func PhpGetDataCompareFunc(sort_type zend.ZendLong, reverse int) types.CompareFuncT {
	switch sort_type & ^PHP_SORT_FLAG_CASE {
	case PHP_SORT_NUMERIC:
		if reverse != 0 {
			return PhpArrayReverseDataCompareNumeric
		} else {
			return PhpArrayDataCompareNumeric
		}
	case PHP_SORT_STRING:
		if (sort_type & PHP_SORT_FLAG_CASE) != 0 {
			if reverse != 0 {
				return PhpArrayReverseDataCompareStringCase
			} else {
				return PhpArrayDataCompareStringCase
			}
		} else {
			if reverse != 0 {
				return PhpArrayReverseDataCompareString
			} else {
				return PhpArrayDataCompareString
			}
		}
	case PHP_SORT_NATURAL:
		if (sort_type & PHP_SORT_FLAG_CASE) != 0 {
			if reverse != 0 {
				return PhpArrayReverseNaturalCaseCompare
			} else {
				return PhpArrayNaturalCaseCompare
			}
		} else {
			if reverse != 0 {
				return PhpArrayReverseNaturalCompare
			} else {
				return PhpArrayNaturalCompare
			}
		}
	case PHP_SORT_LOCALE_STRING:
		if reverse != 0 {
			return PhpArrayReverseDataCompareStringLocale
		} else {
			return PhpArrayDataCompareStringLocale
		}
	case PHP_SORT_REGULAR:
		fallthrough
	default:
		if reverse != 0 {
			return PhpArrayReverseDataCompare
		} else {
			return PhpArrayDataCompare
		}
	}
	return nil
}
func ZifKrsort(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var array *types.Zval
	var sort_type zend.ZendLong = PHP_SORT_REGULAR
	var cmp types.CompareFuncT
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			array = fp.ParseArrayEx(false, true)
			fp.StartOptional()
			sort_type = fp.ParseLong()
			if fp.HasError() {
				fp.HandleError()
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	cmp = PhpGetKeyCompareFunc(sort_type, 1)
	if array.GetArr().SortCompatible(cmp, 0) == types.FAILURE {
		return_value.SetFalse()
		return
	}
	return_value.SetTrue()
	return
}
func ZifKsort(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var array *types.Zval
	var sort_type zend.ZendLong = PHP_SORT_REGULAR
	var cmp types.CompareFuncT
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			array = fp.ParseArrayEx(false, true)
			fp.StartOptional()
			sort_type = fp.ParseLong()
			if fp.HasError() {
				fp.HandleError()
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	cmp = PhpGetKeyCompareFunc(sort_type, 0)
	if array.GetArr().SortCompatible(cmp, 0) == types.FAILURE {
		return_value.SetFalse()
		return
	}
	return_value.SetTrue()
	return
}
func PhpCountRecursive(ht *types.HashTable) zend.ZendLong {
	var cnt zend.ZendLong = 0
	var element *types.Zval
	if (ht.GetGcFlags() & types.GC_IMMUTABLE) == 0 {
		if ht.IsRecursive() {
			core.PhpErrorDocref(nil, zend.E_WARNING, "recursion detected")
			return 0
		}
		ht.ProtectRecursive()
	}
	cnt = ht.Count()
	var __ht *types.HashTable = ht
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()

		element = _z
		element = types.ZVAL_DEREF(element)
		if element.IsType(types.IS_ARRAY) {
			cnt += PhpCountRecursive(element.GetArr())
		}
	}
	if (ht.GetGcFlags() & types.GC_IMMUTABLE) == 0 {
		ht.UnprotectRecursive()
	}
	return cnt
}
func ZifCount(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var array *types.Zval
	var mode zend.ZendLong = COUNT_NORMAL
	var cnt zend.ZendLong
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			argparse.Z_PARAM_PROLOGUE(0, 0)
			argparse.ZendParseArgZvalDeref(_arg, &array, 0)
			fp.StartOptional()
			mode = fp.ParseLong()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	switch array.GetType() {
	case types.IS_NULL:
		core.PhpErrorDocref(nil, zend.E_WARNING, "Parameter must be an array or an object that implements Countable")
		return_value.SetLong(0)
		return
	case types.IS_ARRAY:
		if mode != COUNT_RECURSIVE {
			cnt = array.GetArr().Count()
		} else {
			cnt = PhpCountRecursive(array.GetArr())
		}
		return_value.SetLong(cnt)
		return
	case types.IS_OBJECT:
		var retval types.Zval

		/* first, we check if the handler is defined */

		if types.Z_OBJ_HT_P(array).GetCountElements() != nil {
			return_value.SetLong(1)
			if types.SUCCESS == types.Z_OBJ_HT(*array).GetCountElements()(array, &(return_value.GetLval())) {
				return
			}
			if zend.EG__().GetException() != nil {
				return
			}
		}

		/* if not and the object implements Countable we call its count() method */

		if zend.InstanceofFunction(types.Z_OBJCE_P(array), zend.ZendCeCountable) != 0 {
			zend.ZendCallMethodWith0Params(array, nil, nil, "count", &retval)
			if retval.GetType() != types.IS_UNDEF {
				return_value.SetLong(zend.ZvalGetLong(&retval))
				zend.ZvalPtrDtor(&retval)
			}
			return
		}

		/* If There's no handler and it doesn't implement Countable then add a warning */

		core.PhpErrorDocref(nil, zend.E_WARNING, "Parameter must be an array or an object that implements Countable")
		return_value.SetLong(1)
		return
	default:
		core.PhpErrorDocref(nil, zend.E_WARNING, "Parameter must be an array or an object that implements Countable")
		return_value.SetLong(1)
		return
	}
}
func PhpNatsort(executeData *zend.ZendExecuteData, return_value *types.Zval, fold_case int) {
	var array *types.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			array = fp.ParseArrayEx(false, true)
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	if fold_case != 0 {
		if array.GetArr().SortCompatible(PhpArrayNaturalCaseCompare, 0) == types.FAILURE {
			return
		}
	} else {
		if array.GetArr().SortCompatible(PhpArrayNaturalCompare, 0) == types.FAILURE {
			return
		}
	}
	return_value.SetTrue()
	return
}
func ZifNatsort(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	PhpNatsort(executeData, return_value, 0)
}
func ZifNatcasesort(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	PhpNatsort(executeData, return_value, 1)
}
func ZifAsort(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var array *types.Zval
	var sort_type zend.ZendLong = PHP_SORT_REGULAR
	var cmp types.CompareFuncT
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			array = fp.ParseArrayEx(false, true)
			fp.StartOptional()
			sort_type = fp.ParseLong()
			if fp.HasError() {
				fp.HandleError()
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	cmp = PhpGetDataCompareFunc(sort_type, 0)
	if array.GetArr().SortCompatible(cmp, 0) == types.FAILURE {
		return_value.SetFalse()
		return
	}
	return_value.SetTrue()
	return
}
func ZifArsort(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var array *types.Zval
	var sort_type zend.ZendLong = PHP_SORT_REGULAR
	var cmp types.CompareFuncT
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			array = fp.ParseArrayEx(false, true)
			fp.StartOptional()
			sort_type = fp.ParseLong()
			if fp.HasError() {
				fp.HandleError()
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	cmp = PhpGetDataCompareFunc(sort_type, 1)
	if array.GetArr().SortCompatible(cmp, 0) == types.FAILURE {
		return_value.SetFalse()
		return
	}
	return_value.SetTrue()
	return
}
func ZifSort(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var array *types.Zval
	var sort_type zend.ZendLong = PHP_SORT_REGULAR
	var cmp types.CompareFuncT
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			array = fp.ParseArrayEx(false, true)
			fp.StartOptional()
			sort_type = fp.ParseLong()
			if fp.HasError() {
				fp.HandleError()
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	cmp = PhpGetDataCompareFunc(sort_type, 0)
	if array.GetArr().SortCompatible(cmp, 1) == types.FAILURE {
		return_value.SetFalse()
		return
	}
	return_value.SetTrue()
	return
}
func ZifRsort(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var array *types.Zval
	var sort_type zend.ZendLong = PHP_SORT_REGULAR
	var cmp types.CompareFuncT
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			array = fp.ParseArrayEx(false, true)
			fp.StartOptional()
			sort_type = fp.ParseLong()
			if fp.HasError() {
				fp.HandleError()
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	cmp = PhpGetDataCompareFunc(sort_type, 1)
	if array.GetArr().SortCompatible(cmp, 1) == types.FAILURE {
		return_value.SetFalse()
		return
	}
	return_value.SetTrue()
	return
}
func PhpArrayUserCompare(a any, b any) int {
	var f *types.Bucket
	var s *types.Bucket
	var args []types.Zval
	var retval types.Zval
	f = (*types.Bucket)(a)
	s = (*types.Bucket)(b)
	types.ZVAL_COPY(&args[0], f.GetVal())
	types.ZVAL_COPY(&args[1], s.GetVal())
	BG(user_compare_fci).param_count = 2
	BG(user_compare_fci).params = args
	BG(user_compare_fci).retval = &retval
	BG(user_compare_fci).no_separation = 0
	if zend.ZendCallFunction(&(BG(user_compare_fci)), &(BG(user_compare_fci_cache))) == types.SUCCESS && retval.GetType() != types.IS_UNDEF {
		var ret zend.ZendLong = zend.ZvalGetLong(&retval)
		zend.ZvalPtrDtor(&retval)
		zend.ZvalPtrDtor(&args[1])
		zend.ZvalPtrDtor(&args[0])
		return zend.ZEND_NORMALIZE_BOOL(ret)
	} else {
		zend.ZvalPtrDtor(&args[1])
		zend.ZvalPtrDtor(&args[0])
		return 0
	}
}
func PHP_ARRAY_CMP_FUNC_CHECK(func_name **types.Zval) {
	if zend.ZendIsCallable(*func_name, 0, nil) == 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Invalid comparison function")
		BG(user_compare_fci) = old_user_compare_fci
		BG(user_compare_fci_cache) = old_user_compare_fci_cache
		return_value.SetFalse()
		return
	}
}
func PHP_ARRAY_CMP_FUNC_BACKUP() {
	old_user_compare_fci = BG(user_compare_fci)
	old_user_compare_fci_cache = BG(user_compare_fci_cache)
	BG(user_compare_fci_cache) = zend.EmptyFcallInfoCache
}
func PHP_ARRAY_CMP_FUNC_RESTORE() {
	zend.ZendReleaseFcallInfoCache(&(BG(user_compare_fci_cache)))
	BG(user_compare_fci) = old_user_compare_fci
	BG(user_compare_fci_cache) = old_user_compare_fci_cache
}
func PhpUsort(executeData *zend.ZendExecuteData, return_value *types.Zval, compare_func types.CompareFuncT, renumber types.ZendBool) {
	var array *types.Zval
	var arr *types.ZendArray
	var retval types.ZendBool
	var old_user_compare_fci zend.ZendFcallInfo
	var old_user_compare_fci_cache zend.ZendFcallInfoCache
	PHP_ARRAY_CMP_FUNC_BACKUP()
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 2
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			array = fp.ParseArrayEx2(false, true, false)
			argparse.Z_PARAM_PROLOGUE(0, 0)
			if argparse.ZendParseArgFunc(_arg, &(BG(user_compare_fci)), &(BG(user_compare_fci_cache)), 0, &_error) == 0 {
				if _error == nil {
					_expected_type = argparse.Z_EXPECTED_FUNC
					_error_code = argparse.ZPP_ERROR_WRONG_ARG
					break
				} else {
					_error_code = argparse.ZPP_ERROR_WRONG_CALLBACK
					break
				}
			} else if _error != nil {
				zend.ZendWrongCallbackDeprecated(_i, _error)
			}
			break
		}
		if _error_code != argparse.ZPP_ERROR_OK {
			fp.HandleError()
			PHP_ARRAY_CMP_FUNC_RESTORE()
			return
		}
		break
	}
	arr = array.GetArr()
	if arr.GetNNumOfElements() == 0 {
		PHP_ARRAY_CMP_FUNC_RESTORE()
		return_value.SetTrue()
		return
	}

	/* Copy array, so the in-place modifications will not be visible to the callback function */

	arr = zend.ZendArrayDup(arr)
	retval = arr.SortCompatible(compare_func, renumber) != types.FAILURE
	var garbage types.Zval
	types.ZVAL_COPY_VALUE(&garbage, array)
	array.SetArray(arr)
	zend.ZvalPtrDtor(&garbage)
	PHP_ARRAY_CMP_FUNC_RESTORE()
	types.ZVAL_BOOL(return_value, retval != 0)
	return
}
func ZifUsort(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	PhpUsort(executeData, return_value, PhpArrayUserCompare, 1)
}
func ZifUasort(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	PhpUsort(executeData, return_value, PhpArrayUserCompare, 0)
}
func PhpArrayUserKeyCompare(a any, b any) int {
	var f *types.Bucket
	var s *types.Bucket
	var args []types.Zval
	var retval types.Zval
	var result zend.ZendLong
	f = (*types.Bucket)(a)
	s = (*types.Bucket)(b)
	if f.GetKey() == nil {
		args[0].SetLong(f.GetH())
	} else {
		args[0].SetStringCopy(f.GetKey())
	}
	if s.GetKey() == nil {
		args[1].SetLong(s.GetH())
	} else {
		args[1].SetStringCopy(s.GetKey())
	}
	BG(user_compare_fci).param_count = 2
	BG(user_compare_fci).params = args
	BG(user_compare_fci).retval = &retval
	BG(user_compare_fci).no_separation = 0
	if zend.ZendCallFunction(&(BG(user_compare_fci)), &(BG(user_compare_fci_cache))) == types.SUCCESS && retval.GetType() != types.IS_UNDEF {
		result = zend.ZvalGetLong(&retval)
		zend.ZvalPtrDtor(&retval)
	} else {
		result = 0
	}
	zend.ZvalPtrDtor(&args[0])
	zend.ZvalPtrDtor(&args[1])
	return zend.ZEND_NORMALIZE_BOOL(result)
}
func ZifUksort(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	PhpUsort(executeData, return_value, PhpArrayUserKeyCompare, 0)
}
func ZifEnd(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var array *types.HashTable
	var entry *types.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			array = fp.ParseArrayHtEx(false, true)
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	zend.ZendHashInternalPointerEnd(array)
	if zend.USED_RET() {
		if b.Assign(&entry, zend.ZendHashGetCurrentData(array)) == nil {
			return_value.SetFalse()
			return
		}
		if entry.IsType(types.IS_INDIRECT) {
			entry = entry.GetZv()
		}
		types.ZVAL_COPY_DEREF(return_value, entry)
	}
}
func ZifPrev(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var array *types.HashTable
	var entry *types.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			array = fp.ParseArrayHtEx(false, true)
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	zend.ZendHashMoveBackwards(array)
	if zend.USED_RET() {
		if b.Assign(&entry, zend.ZendHashGetCurrentData(array)) == nil {
			return_value.SetFalse()
			return
		}
		if entry.IsType(types.IS_INDIRECT) {
			entry = entry.GetZv()
		}
		types.ZVAL_COPY_DEREF(return_value, entry)
	}
}
func ZifNext(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var array *types.HashTable
	var entry *types.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			array = fp.ParseArrayHtEx(false, true)
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	zend.ZendHashMoveForward(array)
	if zend.USED_RET() {
		if b.Assign(&entry, zend.ZendHashGetCurrentData(array)) == nil {
			return_value.SetFalse()
			return
		}
		if entry.IsType(types.IS_INDIRECT) {
			entry = entry.GetZv()
		}
		types.ZVAL_COPY_DEREF(return_value, entry)
	}
}
func ZifReset(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var array *types.HashTable
	var entry *types.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			array = fp.ParseArrayHtEx(false, true)
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	zend.ZendHashInternalPointerReset(array)
	if zend.USED_RET() {
		if b.Assign(&entry, zend.ZendHashGetCurrentData(array)) == nil {
			return_value.SetFalse()
			return
		}
		if entry.IsType(types.IS_INDIRECT) {
			entry = entry.GetZv()
		}
		types.ZVAL_COPY_DEREF(return_value, entry)
	}
}
func ZifCurrent(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var array *types.HashTable
	var entry *types.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			array = fp.ParseArrayOrObjectHt()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	if b.Assign(&entry, zend.ZendHashGetCurrentData(array)) == nil {
		return_value.SetFalse()
		return
	}
	if entry.IsType(types.IS_INDIRECT) {
		entry = entry.GetZv()
	}
	types.ZVAL_COPY_DEREF(return_value, entry)
}
func ZifKey(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var array *types.HashTable
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			array = fp.ParseArrayOrObjectHt()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	zend.ZendHashGetCurrentKeyZval(array, return_value)
}
func ZifMin(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var argc int
	var args *types.Zval = nil
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = -1
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			var _num_varargs int = _num_args - _i - 0
			if _num_varargs > 0 {
				args = _real_arg + 1
				argc = _num_varargs
				_i += _num_varargs
				_real_arg += _num_varargs
			} else {
				args = nil
				argc = 0
			}
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}

	/* mixed min ( array $values ) */

	if argc == 1 {
		var result *types.Zval
		if args[0].GetType() != types.IS_ARRAY {
			core.PhpErrorDocref(nil, zend.E_WARNING, "When only one parameter is given, it must be an array")
			return_value.SetNull()
		} else {
			if b.Assign(&result, zend.ZendHashMinmax(args[0].GetArr(), PhpArrayDataCompare, 0)) != nil {
				types.ZVAL_COPY_DEREF(return_value, result)
			} else {
				core.PhpErrorDocref(nil, zend.E_WARNING, "Array must contain at least one element")
				return_value.SetFalse()
			}
		}
	} else {

		/* mixed min ( mixed $value1 , mixed $value2 [, mixed $value3... ] ) */

		var min *types.Zval
		var result types.Zval
		var i int
		min = &args[0]
		for i = 1; i < argc; i++ {
			zend.IsSmallerFunction(&result, &args[i], min)
			if result.IsType(types.IS_TRUE) {
				min = &args[i]
			}
		}
		types.ZVAL_COPY(return_value, min)
	}

	/* mixed min ( array $values ) */
}
func ZifMax(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var args *types.Zval = nil
	var argc int
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = -1
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			var _num_varargs int = _num_args - _i - 0
			if _num_varargs > 0 {
				args = _real_arg + 1
				argc = _num_varargs
				_i += _num_varargs
				_real_arg += _num_varargs
			} else {
				args = nil
				argc = 0
			}
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}

	/* mixed max ( array $values ) */

	if argc == 1 {
		var result *types.Zval
		if args[0].GetType() != types.IS_ARRAY {
			core.PhpErrorDocref(nil, zend.E_WARNING, "When only one parameter is given, it must be an array")
			return_value.SetNull()
		} else {
			if b.Assign(&result, zend.ZendHashMinmax(args[0].GetArr(), PhpArrayDataCompare, 1)) != nil {
				types.ZVAL_COPY_DEREF(return_value, result)
			} else {
				core.PhpErrorDocref(nil, zend.E_WARNING, "Array must contain at least one element")
				return_value.SetFalse()
			}
		}
	} else {

		/* mixed max ( mixed $value1 , mixed $value2 [, mixed $value3... ] ) */

		var max *types.Zval
		var result types.Zval
		var i int
		max = &args[0]
		for i = 1; i < argc; i++ {
			zend.IsSmallerOrEqualFunction(&result, &args[i], max)
			if result.IsType(types.IS_FALSE) {
				max = &args[i]
			}
		}
		types.ZVAL_COPY(return_value, max)
	}

	/* mixed max ( array $values ) */
}
func PhpArrayWalk(array *types.Zval, userdata *types.Zval, recursive int) int {
	var args []types.Zval
	var retval types.Zval
	var zv *types.Zval
	var target_hash *types.HashTable = zend.HASH_OF(array)
	var pos types.HashPosition
	var ht_iter uint32
	var result int = types.SUCCESS

	/* Set up known arguments */

	args[1].SetUndef()
	if userdata != nil {
		types.ZVAL_COPY(&args[2], userdata)
	}
	BG(array_walk_fci).retval = &retval
	if userdata != nil {
		BG(array_walk_fci).param_count = 3
	} else {
		BG(array_walk_fci).param_count = 2
	}
	BG(array_walk_fci).params = args
	BG(array_walk_fci).no_separation = 0
	zend.ZendHashInternalPointerResetEx(target_hash, &pos)
	ht_iter = zend.ZendHashIteratorAdd(target_hash, pos)

	/* Iterate through hash */

	for {

		/* Retrieve value */

		zv = zend.ZendHashGetCurrentDataEx(target_hash, &pos)
		if zv == nil {
			break
		}

		/* Skip undefined indirect elements */

		if zv.IsType(types.IS_INDIRECT) {
			zv = zv.GetZv()
			if zv.IsType(types.IS_UNDEF) {
				zend.ZendHashMoveForwardEx(target_hash, &pos)
				continue
			}

			/* Add type source for property references. */

			if zv.GetType() != types.IS_REFERENCE && array.IsType(types.IS_OBJECT) {
				var prop_info *zend.ZendPropertyInfo = zend.ZendGetTypedPropertyInfoForSlot(array.GetObj(), zv)
				if prop_info != nil {
					zv.SetNewRef(zv)
					zend.ZEND_REF_ADD_TYPE_SOURCE(zv.GetRef(), prop_info)
				}
			}

			/* Add type source for property references. */

		}

		/* Ensure the value is a reference. Otherwise the location of the value may be freed. */

		types.ZVAL_MAKE_REF(zv)

		/* Retrieve key */

		zend.ZendHashGetCurrentKeyZvalEx(target_hash, &args[1], &pos)

		/* Move to next element already now -- this mirrors the approach used by foreach
		 * and ensures proper behavior with regard to modifications. */

		zend.ZendHashMoveForwardEx(target_hash, &pos)

		/* Back up hash position, as it may change */

		zend.EG__().GetHtIterators()[ht_iter].SetPos(pos)
		if recursive != 0 && types.Z_REFVAL_P(zv).IsType(types.IS_ARRAY) {
			var thash *types.HashTable
			var orig_array_walk_fci zend.ZendFcallInfo
			var orig_array_walk_fci_cache zend.ZendFcallInfoCache
			var ref types.Zval
			types.ZVAL_COPY_VALUE(&ref, zv)
			zv = types.ZVAL_DEREF(zv)
			types.SEPARATE_ARRAY(zv)
			thash = zv.GetArr()
			if thash.IsRecursive() {
				core.PhpErrorDocref(nil, zend.E_WARNING, "recursion detected")
				result = types.FAILURE
				break
			}

			/* backup the fcall info and cache */

			orig_array_walk_fci = BG(array_walk_fci)
			orig_array_walk_fci_cache = BG(array_walk_fci_cache)
			ref.AddRefcount()
			thash.ProtectRecursive()
			result = PhpArrayWalk(zv, userdata, recursive)
			if types.Z_REFVAL(ref).IsType(types.IS_ARRAY) && thash == types.Z_REFVAL(ref).GetArr() {

				/* If the hashtable changed in the meantime, we'll "leak" this apply count
				 * increment -- our reference to thash is no longer valid. */

				thash.UnprotectRecursive()

				/* If the hashtable changed in the meantime, we'll "leak" this apply count
				 * increment -- our reference to thash is no longer valid. */

			}
			zend.ZvalPtrDtor(&ref)

			/* restore the fcall info and cache */

			BG(array_walk_fci) = orig_array_walk_fci
			BG(array_walk_fci_cache) = orig_array_walk_fci_cache
		} else {
			types.ZVAL_COPY(&args[0], zv)

			/* Call the userland function */

			result = zend.ZendCallFunction(&(BG(array_walk_fci)), &(BG(array_walk_fci_cache)))
			if result == types.SUCCESS {
				zend.ZvalPtrDtor(&retval)
			}
			zend.ZvalPtrDtor(&args[0])
		}
		if args[1].GetType() != types.IS_UNDEF {
			zend.ZvalPtrDtor(&args[1])
			args[1].SetUndef()
		}
		if result == types.FAILURE {
			break
		}

		/* Reload array and position -- both may have changed */

		if array.IsType(types.IS_ARRAY) {
			pos = zend.ZendHashIteratorPosEx(ht_iter, array)
			target_hash = array.GetArr()
		} else if array.IsType(types.IS_OBJECT) {
			target_hash = types.Z_OBJPROP_P(array)
			pos = zend.ZendHashIteratorPos(ht_iter, target_hash)
		} else {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Iterated value is no longer an array or object")
			result = types.FAILURE
			break
		}

		/* Reload array and position -- both may have changed */

		if zend.EG__().GetException() != nil {
			break
		}
	}
	if userdata != nil {
		zend.ZvalPtrDtor(&args[2])
	}
	zend.ZendHashIteratorDel(ht_iter)
	return result
}
func ZifArrayWalk(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var array *types.Zval
	var userdata *types.Zval = nil
	var orig_array_walk_fci zend.ZendFcallInfo
	var orig_array_walk_fci_cache zend.ZendFcallInfoCache
	orig_array_walk_fci = BG(array_walk_fci)
	orig_array_walk_fci_cache = BG(array_walk_fci_cache)
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 3
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			array = fp.ParseArrayOrObjectEx(false, true)
			argparse.Z_PARAM_PROLOGUE(0, 0)
			if argparse.ZendParseArgFunc(_arg, &(BG(array_walk_fci)), &(BG(array_walk_fci_cache)), 0, &_error) == 0 {
				if _error == nil {
					_expected_type = argparse.Z_EXPECTED_FUNC
					_error_code = argparse.ZPP_ERROR_WRONG_ARG
					break
				} else {
					_error_code = argparse.ZPP_ERROR_WRONG_CALLBACK
					break
				}
			} else if _error != nil {
				zend.ZendWrongCallbackDeprecated(_i, _error)
			}
			fp.StartOptional()
			argparse.Z_PARAM_PROLOGUE(0, 0)
			argparse.ZendParseArgZvalDeref(_arg, &userdata, 0)
			break
		}
		if _error_code != argparse.ZPP_ERROR_OK {
			fp.HandleError()
			BG(array_walk_fci) = orig_array_walk_fci
			BG(array_walk_fci_cache) = orig_array_walk_fci_cache
			return
		}
		break
	}
	PhpArrayWalk(array, userdata, 0)
	zend.ZendReleaseFcallInfoCache(&(BG(array_walk_fci_cache)))
	BG(array_walk_fci) = orig_array_walk_fci
	BG(array_walk_fci_cache) = orig_array_walk_fci_cache
	return_value.SetTrue()
	return
}
func ZifArrayWalkRecursive(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var array *types.Zval
	var userdata *types.Zval = nil
	var orig_array_walk_fci zend.ZendFcallInfo
	var orig_array_walk_fci_cache zend.ZendFcallInfoCache
	orig_array_walk_fci = BG(array_walk_fci)
	orig_array_walk_fci_cache = BG(array_walk_fci_cache)
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 3
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			array = fp.ParseArrayOrObjectEx(false, true)
			argparse.Z_PARAM_PROLOGUE(0, 0)
			if argparse.ZendParseArgFunc(_arg, &(BG(array_walk_fci)), &(BG(array_walk_fci_cache)), 0, &_error) == 0 {
				if _error == nil {
					_expected_type = argparse.Z_EXPECTED_FUNC
					_error_code = argparse.ZPP_ERROR_WRONG_ARG
					break
				} else {
					_error_code = argparse.ZPP_ERROR_WRONG_CALLBACK
					break
				}
			} else if _error != nil {
				zend.ZendWrongCallbackDeprecated(_i, _error)
			}
			fp.StartOptional()
			argparse.Z_PARAM_PROLOGUE(0, 0)
			argparse.ZendParseArgZvalDeref(_arg, &userdata, 0)
			break
		}
		if _error_code != argparse.ZPP_ERROR_OK {
			fp.HandleError()
			BG(array_walk_fci) = orig_array_walk_fci
			BG(array_walk_fci_cache) = orig_array_walk_fci_cache
			return
		}
		break
	}
	PhpArrayWalk(array, userdata, 1)
	zend.ZendReleaseFcallInfoCache(&(BG(array_walk_fci_cache)))
	BG(array_walk_fci) = orig_array_walk_fci
	BG(array_walk_fci_cache) = orig_array_walk_fci_cache
	return_value.SetTrue()
	return
}
func PhpSearchArray(executeData *zend.ZendExecuteData, return_value *types.Zval, behavior int) {
	var value *types.Zval
	var array *types.Zval
	var entry *types.Zval
	var num_idx zend.ZendUlong
	var str_idx *types.ZendString
	var strict types.ZendBool = 0
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 3
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			argparse.Z_PARAM_PROLOGUE(0, 0)
			argparse.ZendParseArgZvalDeref(_arg, &value, 0)
			array = fp.ParseArray()
			fp.StartOptional()
			strict = fp.ParseBool()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	if strict != 0 {
		if value.IsType(types.IS_LONG) {
			var __ht *types.HashTable = array.GetArr()
			for _, _p := range __ht.foreachData() {
				var _z *types.Zval = _p.GetVal()
				if _z.IsType(types.IS_INDIRECT) {
					_z = _z.GetZv()
					if _z.IsType(types.IS_UNDEF) {
						continue
					}
				}
				num_idx = _p.GetH()
				str_idx = _p.GetKey()
				entry = _z
				entry = types.ZVAL_DEREF(entry)
				if entry.IsType(types.IS_LONG) && entry.GetLval() == value.GetLval() {
					if behavior == 0 {
						return_value.SetTrue()
						return
					} else {
						if str_idx != nil {
							return_value.SetStringCopy(str_idx)
						} else {
							return_value.SetLong(num_idx)
						}
						return
					}
				}
			}
		} else {
			var __ht *types.HashTable = array.GetArr()
			for _, _p := range __ht.foreachData() {
				var _z *types.Zval = _p.GetVal()
				if _z.IsType(types.IS_INDIRECT) {
					_z = _z.GetZv()
					if _z.IsType(types.IS_UNDEF) {
						continue
					}
				}
				num_idx = _p.GetH()
				str_idx = _p.GetKey()
				entry = _z
				entry = types.ZVAL_DEREF(entry)
				if zend.FastIsIdenticalFunction(value, entry) != 0 {
					if behavior == 0 {
						return_value.SetTrue()
						return
					} else {
						if str_idx != nil {
							return_value.SetStringCopy(str_idx)
						} else {
							return_value.SetLong(num_idx)
						}
						return
					}
				}
			}
		}
	} else {
		if value.IsType(types.IS_LONG) {
			var __ht *types.HashTable = array.GetArr()
			for _, _p := range __ht.foreachData() {
				var _z *types.Zval = _p.GetVal()
				if _z.IsType(types.IS_INDIRECT) {
					_z = _z.GetZv()
					if _z.IsType(types.IS_UNDEF) {
						continue
					}
				}
				num_idx = _p.GetH()
				str_idx = _p.GetKey()
				entry = _z
				if zend.FastEqualCheckLong(value, entry) != 0 {
					if behavior == 0 {
						return_value.SetTrue()
						return
					} else {
						if str_idx != nil {
							return_value.SetStringCopy(str_idx)
						} else {
							return_value.SetLong(num_idx)
						}
						return
					}
				}
			}
		} else if value.IsType(types.IS_STRING) {
			var __ht *types.HashTable = array.GetArr()
			for _, _p := range __ht.foreachData() {
				var _z *types.Zval = _p.GetVal()
				if _z.IsType(types.IS_INDIRECT) {
					_z = _z.GetZv()
					if _z.IsType(types.IS_UNDEF) {
						continue
					}
				}
				num_idx = _p.GetH()
				str_idx = _p.GetKey()
				entry = _z
				if zend.FastEqualCheckString(value, entry) != 0 {
					if behavior == 0 {
						return_value.SetTrue()
						return
					} else {
						if str_idx != nil {
							return_value.SetStringCopy(str_idx)
						} else {
							return_value.SetLong(num_idx)
						}
						return
					}
				}
			}
		} else {
			var __ht *types.HashTable = array.GetArr()
			for _, _p := range __ht.foreachData() {
				var _z *types.Zval = _p.GetVal()
				if _z.IsType(types.IS_INDIRECT) {
					_z = _z.GetZv()
					if _z.IsType(types.IS_UNDEF) {
						continue
					}
				}
				num_idx = _p.GetH()
				str_idx = _p.GetKey()
				entry = _z
				if zend.FastEqualCheckFunction(value, entry) != 0 {
					if behavior == 0 {
						return_value.SetTrue()
						return
					} else {
						if str_idx != nil {
							return_value.SetStringCopy(str_idx)
						} else {
							return_value.SetLong(num_idx)
						}
						return
					}
				}
			}
		}
	}
	return_value.SetFalse()
	return
}
func ZifInArray(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	PhpSearchArray(executeData, return_value, 0)
}
func ZifArraySearch(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	PhpSearchArray(executeData, return_value, 1)
}
func PhpValidVarName(var_name *byte, var_name_len int) int {
	/* first 256 bits for first character, and second 256 bits for the next */

	var charset []uint32 = []uint32{0x0, 0x0, 0x87fffffe, 0x7fffffe, 0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff}
	var charset2 []uint32 = []uint32{0x0, 0x3ff0000, 0x87fffffe, 0x7fffffe, 0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff}
	var i int
	var ch uint32
	if var_name_len == 0 {
		return 0
	}

	/* These are allowed as first char: [a-zA-Z_\x7f-\xff] */

	ch = uint32((*uint8)(var_name))[0]
	if zend.ZEND_BIT_TEST(charset, ch) == 0 {
		return 0
	}

	/* And these as the rest: [a-zA-Z0-9_\x7f-\xff] */

	if var_name_len > 1 {
		i = 1
		for {
			ch = uint32((*uint8)(var_name))[i]
			if zend.ZEND_BIT_TEST(charset2, ch) == 0 {
				return 0
			}
			if b.PreInc(&i) >= var_name_len {
				break
			}
		}
	}
	return 1
}
func PhpPrefixVarname(result *types.Zval, prefix *types.Zval, var_name *byte, var_name_len int, add_underscore types.ZendBool) int {
	result.SetString(types.ZendStringAlloc(prefix.GetStr().GetLen()+b.Cond(add_underscore != 0, 1, 0)+var_name_len, 0))
	memcpy(result.GetStr().GetVal(), prefix.GetStr().GetVal(), prefix.GetStr().GetLen())
	if add_underscore != 0 {
		result.GetStr().GetVal()[prefix.GetStr().GetLen()] = '_'
	}
	memcpy(result.GetStr().GetVal()+prefix.GetStr().GetLen()+b.Cond(add_underscore != 0, 1, 0), var_name, var_name_len+1)
	return types.SUCCESS
}
func PhpExtractRefIfExists(arr *types.ZendArray, symbol_table *types.ZendArray) zend.ZendLong {
	var count zend.ZendLong = 0
	var var_name *types.ZendString
	var entry *types.Zval
	var orig_var *types.Zval
	var __ht *types.HashTable = arr
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()
		if _z.IsType(types.IS_INDIRECT) {
			_z = _z.GetZv()
			if _z.IsType(types.IS_UNDEF) {
				continue
			}
		}
		var_name = _p.GetKey()
		entry = _z
		if var_name == nil {
			continue
		}
		orig_var = symbol_table.KeyFind(var_name.GetStr())
		if orig_var != nil {
			if orig_var.IsType(types.IS_INDIRECT) {
				orig_var = orig_var.GetZv()
				if orig_var.IsType(types.IS_UNDEF) {
					continue
				}
			}
			if PhpValidVarName(var_name.GetVal(), var_name.GetLen()) == 0 {
				continue
			}
			if types.ZendStringEqualsLiteral(var_name, "GLOBALS") {
				continue
			}
			if types.ZendStringEqualsLiteral(var_name, "this") {
				zend.ZendThrowError(nil, "Cannot re-assign $this")
				return -1
			}
			if entry.IsReference() {
				entry.AddRefcount()
			} else {
				types.ZVAL_MAKE_REF_EX(entry, 2)
			}
			zend.ZvalPtrDtor(orig_var)
			orig_var.SetReference(entry.GetRef())
			count++
		}
	}
	return count
}
func PhpExtractIfExists(arr *types.ZendArray, symbol_table *types.ZendArray) zend.ZendLong {
	var count zend.ZendLong = 0
	var var_name *types.ZendString
	var entry *types.Zval
	var orig_var *types.Zval
	var __ht *types.HashTable = arr
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()
		if _z.IsType(types.IS_INDIRECT) {
			_z = _z.GetZv()
			if _z.IsType(types.IS_UNDEF) {
				continue
			}
		}
		var_name = _p.GetKey()
		entry = _z
		if var_name == nil {
			continue
		}
		orig_var = symbol_table.KeyFind(var_name.GetStr())
		if orig_var != nil {
			if orig_var.IsType(types.IS_INDIRECT) {
				orig_var = orig_var.GetZv()
				if orig_var.IsType(types.IS_UNDEF) {
					continue
				}
			}
			if PhpValidVarName(var_name.GetVal(), var_name.GetLen()) == 0 {
				continue
			}
			if types.ZendStringEqualsLiteral(var_name, "GLOBALS") {
				continue
			}
			if types.ZendStringEqualsLiteral(var_name, "this") {
				zend.ZendThrowError(nil, "Cannot re-assign $this")
				return -1
			}
			entry = types.ZVAL_DEREF(entry)
			zend.ZEND_TRY_ASSIGN_COPY_EX(orig_var, entry, 0)
			if zend.EG__().GetException() != nil {
				return -1
			}
			count++
		}
	}
	return count
}
func PhpExtractRefOverwrite(arr *types.ZendArray, symbol_table *types.ZendArray) zend.ZendLong {
	var count zend.ZendLong = 0
	var var_name *types.ZendString
	var entry *types.Zval
	var orig_var *types.Zval
	var __ht *types.HashTable = arr
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()
		if _z.IsType(types.IS_INDIRECT) {
			_z = _z.GetZv()
			if _z.IsType(types.IS_UNDEF) {
				continue
			}
		}
		var_name = _p.GetKey()
		entry = _z
		if var_name == nil {
			continue
		}
		if PhpValidVarName(var_name.GetVal(), var_name.GetLen()) == 0 {
			continue
		}
		if types.ZendStringEqualsLiteral(var_name, "this") {
			zend.ZendThrowError(nil, "Cannot re-assign $this")
			return -1
		}
		orig_var = symbol_table.KeyFind(var_name.GetStr())
		if orig_var != nil {
			if orig_var.IsType(types.IS_INDIRECT) {
				orig_var = orig_var.GetZv()
			}
			if types.ZendStringEqualsLiteral(var_name, "GLOBALS") {
				continue
			}
			if entry.IsReference() {
				entry.AddRefcount()
			} else {
				types.ZVAL_MAKE_REF_EX(entry, 2)
			}
			zend.ZvalPtrDtor(orig_var)
			orig_var.SetReference(entry.GetRef())
		} else {
			if entry.IsReference() {
				entry.AddRefcount()
			} else {
				types.ZVAL_MAKE_REF_EX(entry, 2)
			}
			symbol_table.KeyAddNew(var_name.GetStr(), entry)
		}
		count++
	}
	return count
}
func PhpExtractOverwrite(arr *types.ZendArray, symbol_table *types.ZendArray) zend.ZendLong {
	var count zend.ZendLong = 0
	var var_name *types.ZendString
	var entry *types.Zval
	var orig_var *types.Zval
	var __ht *types.HashTable = arr
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()
		if _z.IsType(types.IS_INDIRECT) {
			_z = _z.GetZv()
			if _z.IsType(types.IS_UNDEF) {
				continue
			}
		}
		var_name = _p.GetKey()
		entry = _z
		if var_name == nil {
			continue
		}
		if PhpValidVarName(var_name.GetVal(), var_name.GetLen()) == 0 {
			continue
		}
		if types.ZendStringEqualsLiteral(var_name, "this") {
			zend.ZendThrowError(nil, "Cannot re-assign $this")
			return -1
		}
		orig_var = symbol_table.KeyFind(var_name.GetStr())
		if orig_var != nil {
			if orig_var.IsType(types.IS_INDIRECT) {
				orig_var = orig_var.GetZv()
			}
			if types.ZendStringEqualsLiteral(var_name, "GLOBALS") {
				continue
			}
			entry = types.ZVAL_DEREF(entry)
			zend.ZEND_TRY_ASSIGN_COPY_EX(orig_var, entry, 0)
			if zend.EG__().GetException() != nil {
				return -1
			}
		} else {
			entry = types.ZVAL_DEREF(entry)
			entry.TryAddRefcount()
			symbol_table.KeyAddNew(var_name.GetStr(), entry)
		}
		count++
	}
	return count
}
func PhpExtractRefPrefixIfExists(arr *types.ZendArray, symbol_table *types.ZendArray, prefix *types.Zval) zend.ZendLong {
	var count zend.ZendLong = 0
	var var_name *types.ZendString
	var entry *types.Zval
	var orig_var *types.Zval
	var final_name types.Zval
	var __ht *types.HashTable = arr
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()
		if _z.IsType(types.IS_INDIRECT) {
			_z = _z.GetZv()
			if _z.IsType(types.IS_UNDEF) {
				continue
			}
		}
		var_name = _p.GetKey()
		entry = _z
		if var_name == nil {
			continue
		}
		orig_var = symbol_table.KeyFind(var_name.GetStr())
		if orig_var != nil {
			if orig_var.IsType(types.IS_INDIRECT) {
				orig_var = orig_var.GetZv()
				if orig_var.IsType(types.IS_UNDEF) {
					if entry.IsReference() {
						entry.AddRefcount()
					} else {
						types.ZVAL_MAKE_REF_EX(entry, 2)
					}
					orig_var.SetReference(entry.GetRef())
					count++
					continue
				}
			}
			PhpPrefixVarname(&final_name, prefix, var_name.GetVal(), var_name.GetLen(), 1)
			if PhpValidVarName(final_name.GetStr().GetVal(), final_name.GetStr().GetLen()) != 0 {
				if types.ZendStringEqualsLiteral(final_name.GetStr(), "this") {
					zend.ZendThrowError(nil, "Cannot re-assign $this")
					return -1
				} else {
					if entry.IsReference() {
						entry.AddRefcount()
					} else {
						types.ZVAL_MAKE_REF_EX(entry, 2)
					}
					if b.Assign(&orig_var, symbol_table.KeyFind(final_name.GetStr().GetStr())) != nil {
						if orig_var.IsType(types.IS_INDIRECT) {
							orig_var = orig_var.GetZv()
						}
						zend.ZvalPtrDtor(orig_var)
						orig_var.SetReference(entry.GetRef())
					} else {
						symbol_table.KeyAddNew(final_name.GetStr().GetStr(), entry)
					}
					count++
				}
			}
			zend.ZvalPtrDtorStr(&final_name)
		}
	}
	return count
}
func PhpExtractPrefixIfExists(arr *types.ZendArray, symbol_table *types.ZendArray, prefix *types.Zval) zend.ZendLong {
	var count zend.ZendLong = 0
	var var_name *types.ZendString
	var entry *types.Zval
	var orig_var *types.Zval
	var final_name types.Zval
	var __ht *types.HashTable = arr
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()
		if _z.IsType(types.IS_INDIRECT) {
			_z = _z.GetZv()
			if _z.IsType(types.IS_UNDEF) {
				continue
			}
		}
		var_name = _p.GetKey()
		entry = _z
		if var_name == nil {
			continue
		}
		orig_var = symbol_table.KeyFind(var_name.GetStr())
		if orig_var != nil {
			if orig_var.IsType(types.IS_INDIRECT) {
				orig_var = orig_var.GetZv()
				if orig_var.IsType(types.IS_UNDEF) {
					types.ZVAL_COPY_DEREF(orig_var, entry)
					count++
					continue
				}
			}
			PhpPrefixVarname(&final_name, prefix, var_name.GetVal(), var_name.GetLen(), 1)
			if PhpValidVarName(final_name.GetStr().GetVal(), final_name.GetStr().GetLen()) != 0 {
				if types.ZendStringEqualsLiteral(final_name.GetStr(), "this") {
					zend.ZendThrowError(nil, "Cannot re-assign $this")
					return -1
				} else {
					entry = types.ZVAL_DEREF(entry)
					if b.Assign(&orig_var, symbol_table.KeyFind(final_name.GetStr().GetStr())) != nil {
						if orig_var.IsType(types.IS_INDIRECT) {
							orig_var = orig_var.GetZv()
						}
						zend.ZEND_TRY_ASSIGN_COPY_EX(orig_var, entry, 0)
						if zend.EG__().GetException() != nil {
							types.ZendStringReleaseEx(final_name.GetStr(), 0)
							return -1
						}
					} else {
						entry.TryAddRefcount()
						symbol_table.KeyAddNew(final_name.GetStr().GetStr(), entry)
					}
					count++
				}
			}
			zend.ZvalPtrDtorStr(&final_name)
		}
	}
	return count
}
func PhpExtractRefPrefixSame(arr *types.ZendArray, symbol_table *types.ZendArray, prefix *types.Zval) zend.ZendLong {
	var count zend.ZendLong = 0
	var var_name *types.ZendString
	var entry *types.Zval
	var orig_var *types.Zval
	var final_name types.Zval
	var __ht *types.HashTable = arr
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()
		if _z.IsType(types.IS_INDIRECT) {
			_z = _z.GetZv()
			if _z.IsType(types.IS_UNDEF) {
				continue
			}
		}
		var_name = _p.GetKey()
		entry = _z
		if var_name == nil {
			continue
		}
		if var_name.GetLen() == 0 {
			continue
		}
		orig_var = symbol_table.KeyFind(var_name.GetStr())
		if orig_var != nil {
			if orig_var.IsType(types.IS_INDIRECT) {
				orig_var = orig_var.GetZv()
				if orig_var.IsType(types.IS_UNDEF) {
					if entry.IsReference() {
						entry.AddRefcount()
					} else {
						types.ZVAL_MAKE_REF_EX(entry, 2)
					}
					orig_var.SetReference(entry.GetRef())
					count++
					continue
				}
			}
		prefix:
			PhpPrefixVarname(&final_name, prefix, var_name.GetVal(), var_name.GetLen(), 1)
			if PhpValidVarName(final_name.GetStr().GetVal(), final_name.GetStr().GetLen()) != 0 {
				if types.ZendStringEqualsLiteral(final_name.GetStr(), "this") {
					zend.ZendThrowError(nil, "Cannot re-assign $this")
					return -1
				} else {
					if entry.IsReference() {
						entry.AddRefcount()
					} else {
						types.ZVAL_MAKE_REF_EX(entry, 2)
					}
					if b.Assign(&orig_var, symbol_table.KeyFind(final_name.GetStr().GetStr())) != nil {
						if orig_var.IsType(types.IS_INDIRECT) {
							orig_var = orig_var.GetZv()
						}
						zend.ZvalPtrDtor(orig_var)
						orig_var.SetReference(entry.GetRef())
					} else {
						symbol_table.KeyAddNew(final_name.GetStr().GetStr(), entry)
					}
					count++
				}
			}
			zend.ZvalPtrDtorStr(&final_name)
		} else {
			if PhpValidVarName(var_name.GetVal(), var_name.GetLen()) == 0 {
				continue
			}
			if types.ZendStringEqualsLiteral(var_name, "this") {
				goto prefix
			}
			if entry.IsReference() {
				entry.AddRefcount()
			} else {
				types.ZVAL_MAKE_REF_EX(entry, 2)
			}
			symbol_table.KeyAddNew(var_name.GetStr(), entry)
			count++
		}
	}
	return count
}
func PhpExtractPrefixSame(arr *types.ZendArray, symbol_table *types.ZendArray, prefix *types.Zval) zend.ZendLong {
	var count zend.ZendLong = 0
	var var_name *types.ZendString
	var entry *types.Zval
	var orig_var *types.Zval
	var final_name types.Zval
	var __ht *types.HashTable = arr
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()
		if _z.IsType(types.IS_INDIRECT) {
			_z = _z.GetZv()
			if _z.IsType(types.IS_UNDEF) {
				continue
			}
		}
		var_name = _p.GetKey()
		entry = _z
		if var_name == nil {
			continue
		}
		if var_name.GetLen() == 0 {
			continue
		}
		orig_var = symbol_table.KeyFind(var_name.GetStr())
		if orig_var != nil {
			if orig_var.IsType(types.IS_INDIRECT) {
				orig_var = orig_var.GetZv()
				if orig_var.IsType(types.IS_UNDEF) {
					types.ZVAL_COPY_DEREF(orig_var, entry)
					count++
					continue
				}
			}
		prefix:
			PhpPrefixVarname(&final_name, prefix, var_name.GetVal(), var_name.GetLen(), 1)
			if PhpValidVarName(final_name.GetStr().GetVal(), final_name.GetStr().GetLen()) != 0 {
				if types.ZendStringEqualsLiteral(final_name.GetStr(), "this") {
					zend.ZendThrowError(nil, "Cannot re-assign $this")
					return -1
				} else {
					entry = types.ZVAL_DEREF(entry)
					if b.Assign(&orig_var, symbol_table.KeyFind(final_name.GetStr().GetStr())) != nil {
						if orig_var.IsType(types.IS_INDIRECT) {
							orig_var = orig_var.GetZv()
						}
						zend.ZEND_TRY_ASSIGN_COPY_EX(orig_var, entry, 0)
						if zend.EG__().GetException() != nil {
							types.ZendStringReleaseEx(final_name.GetStr(), 0)
							return -1
						}
					} else {
						entry.TryAddRefcount()
						symbol_table.KeyAddNew(final_name.GetStr().GetStr(), entry)
					}
					count++
				}
			}
			zend.ZvalPtrDtorStr(&final_name)
		} else {
			if PhpValidVarName(var_name.GetVal(), var_name.GetLen()) == 0 {
				continue
			}
			if types.ZendStringEqualsLiteral(var_name, "this") {
				goto prefix
			}
			entry = types.ZVAL_DEREF(entry)
			entry.TryAddRefcount()
			symbol_table.KeyAddNew(var_name.GetStr(), entry)
			count++
		}
	}
	return count
}
func PhpExtractRefPrefixAll(arr *types.ZendArray, symbol_table *types.ZendArray, prefix *types.Zval) zend.ZendLong {
	var count zend.ZendLong = 0
	var var_name *types.ZendString
	var num_key zend.ZendUlong
	var entry *types.Zval
	var orig_var *types.Zval
	var final_name types.Zval
	var __ht *types.HashTable = arr
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()
		if _z.IsType(types.IS_INDIRECT) {
			_z = _z.GetZv()
			if _z.IsType(types.IS_UNDEF) {
				continue
			}
		}
		num_key = _p.GetH()
		var_name = _p.GetKey()
		entry = _z
		if var_name != nil {
			if var_name.GetLen() == 0 {
				continue
			}
			PhpPrefixVarname(&final_name, prefix, var_name.GetVal(), var_name.GetLen(), 1)
		} else {
			var str *types.ZendString = zend.ZendLongToStr(num_key)
			PhpPrefixVarname(&final_name, prefix, str.GetVal(), str.GetLen(), 1)
			types.ZendStringReleaseEx(str, 0)
		}
		if PhpValidVarName(final_name.GetStr().GetVal(), final_name.GetStr().GetLen()) != 0 {
			if types.ZendStringEqualsLiteral(final_name.GetStr(), "this") {
				zend.ZendThrowError(nil, "Cannot re-assign $this")
				return -1
			} else {
				if entry.IsReference() {
					entry.AddRefcount()
				} else {
					types.ZVAL_MAKE_REF_EX(entry, 2)
				}
				if b.Assign(&orig_var, symbol_table.KeyFind(final_name.GetStr().GetStr())) != nil {
					if orig_var.IsType(types.IS_INDIRECT) {
						orig_var = orig_var.GetZv()
					}
					zend.ZvalPtrDtor(orig_var)
					orig_var.SetReference(entry.GetRef())
				} else {
					symbol_table.KeyAddNew(final_name.GetStr().GetStr(), entry)
				}
				count++
			}
		}
		zend.ZvalPtrDtorStr(&final_name)
	}
	return count
}
func PhpExtractPrefixAll(arr *types.ZendArray, symbol_table *types.ZendArray, prefix *types.Zval) zend.ZendLong {
	var count zend.ZendLong = 0
	var var_name *types.ZendString
	var num_key zend.ZendUlong
	var entry *types.Zval
	var orig_var *types.Zval
	var final_name types.Zval
	var __ht *types.HashTable = arr
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()
		if _z.IsType(types.IS_INDIRECT) {
			_z = _z.GetZv()
			if _z.IsType(types.IS_UNDEF) {
				continue
			}
		}
		num_key = _p.GetH()
		var_name = _p.GetKey()
		entry = _z
		if var_name != nil {
			if var_name.GetLen() == 0 {
				continue
			}
			PhpPrefixVarname(&final_name, prefix, var_name.GetVal(), var_name.GetLen(), 1)
		} else {
			var str *types.ZendString = zend.ZendLongToStr(num_key)
			PhpPrefixVarname(&final_name, prefix, str.GetVal(), str.GetLen(), 1)
			types.ZendStringReleaseEx(str, 0)
		}
		if PhpValidVarName(final_name.GetStr().GetVal(), final_name.GetStr().GetLen()) != 0 {
			if types.ZendStringEqualsLiteral(final_name.GetStr(), "this") {
				zend.ZendThrowError(nil, "Cannot re-assign $this")
				return -1
			} else {
				entry = types.ZVAL_DEREF(entry)
				if b.Assign(&orig_var, symbol_table.KeyFind(final_name.GetStr().GetStr())) != nil {
					if orig_var.IsType(types.IS_INDIRECT) {
						orig_var = orig_var.GetZv()
					}
					zend.ZEND_TRY_ASSIGN_COPY_EX(orig_var, entry, 0)
					if zend.EG__().GetException() != nil {
						types.ZendStringReleaseEx(final_name.GetStr(), 0)
						return -1
					}
				} else {
					entry.TryAddRefcount()
					symbol_table.KeyAddNew(final_name.GetStr().GetStr(), entry)
				}
				count++
			}
		}
		zend.ZvalPtrDtorStr(&final_name)
	}
	return count
}
func PhpExtractRefPrefixInvalid(arr *types.ZendArray, symbol_table *types.ZendArray, prefix *types.Zval) zend.ZendLong {
	var count zend.ZendLong = 0
	var var_name *types.ZendString
	var num_key zend.ZendUlong
	var entry *types.Zval
	var orig_var *types.Zval
	var final_name types.Zval
	var __ht *types.HashTable = arr
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()
		if _z.IsType(types.IS_INDIRECT) {
			_z = _z.GetZv()
			if _z.IsType(types.IS_UNDEF) {
				continue
			}
		}
		num_key = _p.GetH()
		var_name = _p.GetKey()
		entry = _z
		if var_name != nil {
			if PhpValidVarName(var_name.GetVal(), var_name.GetLen()) == 0 || types.ZendStringEqualsLiteral(var_name, "this") {
				PhpPrefixVarname(&final_name, prefix, var_name.GetVal(), var_name.GetLen(), 1)
				if PhpValidVarName(final_name.GetStr().GetVal(), final_name.GetStr().GetLen()) == 0 {
					zend.ZvalPtrDtorStr(&final_name)
					continue
				}
			} else {
				final_name.SetStringCopy(var_name)
			}
		} else {
			var str *types.ZendString = zend.ZendLongToStr(num_key)
			PhpPrefixVarname(&final_name, prefix, str.GetVal(), str.GetLen(), 1)
			types.ZendStringReleaseEx(str, 0)
			if PhpValidVarName(final_name.GetStr().GetVal(), final_name.GetStr().GetLen()) == 0 {
				zend.ZvalPtrDtorStr(&final_name)
				continue
			}
		}
		if types.ZendStringEqualsLiteral(final_name.GetStr(), "this") {
			zend.ZendThrowError(nil, "Cannot re-assign $this")
			return -1
		} else {
			if entry.IsReference() {
				entry.AddRefcount()
			} else {
				types.ZVAL_MAKE_REF_EX(entry, 2)
			}
			if b.Assign(&orig_var, symbol_table.KeyFind(final_name.GetStr().GetStr())) != nil {
				if orig_var.IsType(types.IS_INDIRECT) {
					orig_var = orig_var.GetZv()
				}
				zend.ZvalPtrDtor(orig_var)
				orig_var.SetReference(entry.GetRef())
			} else {
				symbol_table.KeyAddNew(final_name.GetStr().GetStr(), entry)
			}
			count++
		}
		zend.ZvalPtrDtorStr(&final_name)
	}
	return count
}
func PhpExtractPrefixInvalid(arr *types.ZendArray, symbol_table *types.ZendArray, prefix *types.Zval) zend.ZendLong {
	var count zend.ZendLong = 0
	var var_name *types.ZendString
	var num_key zend.ZendUlong
	var entry *types.Zval
	var orig_var *types.Zval
	var final_name types.Zval
	var __ht *types.HashTable = arr
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()
		if _z.IsType(types.IS_INDIRECT) {
			_z = _z.GetZv()
			if _z.IsType(types.IS_UNDEF) {
				continue
			}
		}
		num_key = _p.GetH()
		var_name = _p.GetKey()
		entry = _z
		if var_name != nil {
			if PhpValidVarName(var_name.GetVal(), var_name.GetLen()) == 0 || types.ZendStringEqualsLiteral(var_name, "this") {
				PhpPrefixVarname(&final_name, prefix, var_name.GetVal(), var_name.GetLen(), 1)
				if PhpValidVarName(final_name.GetStr().GetVal(), final_name.GetStr().GetLen()) == 0 {
					zend.ZvalPtrDtorStr(&final_name)
					continue
				}
			} else {
				final_name.SetStringCopy(var_name)
			}
		} else {
			var str *types.ZendString = zend.ZendLongToStr(num_key)
			PhpPrefixVarname(&final_name, prefix, str.GetVal(), str.GetLen(), 1)
			types.ZendStringReleaseEx(str, 0)
			if PhpValidVarName(final_name.GetStr().GetVal(), final_name.GetStr().GetLen()) == 0 {
				zend.ZvalPtrDtorStr(&final_name)
				continue
			}
		}
		if types.ZendStringEqualsLiteral(final_name.GetStr(), "this") {
			zend.ZendThrowError(nil, "Cannot re-assign $this")
			return -1
		} else {
			entry = types.ZVAL_DEREF(entry)
			if b.Assign(&orig_var, symbol_table.KeyFind(final_name.GetStr().GetStr())) != nil {
				if orig_var.IsType(types.IS_INDIRECT) {
					orig_var = orig_var.GetZv()
				}
				zend.ZEND_TRY_ASSIGN_COPY_EX(orig_var, entry, 0)
				if zend.EG__().GetException() != nil {
					types.ZendStringReleaseEx(final_name.GetStr(), 0)
					return -1
				}
			} else {
				entry.TryAddRefcount()
				symbol_table.KeyAddNew(final_name.GetStr().GetStr(), entry)
			}
			count++
		}
		zend.ZvalPtrDtorStr(&final_name)
	}
	return count
}
func PhpExtractRefSkip(arr *types.ZendArray, symbol_table *types.ZendArray) zend.ZendLong {
	var count zend.ZendLong = 0
	var var_name *types.ZendString
	var entry *types.Zval
	var orig_var *types.Zval
	var __ht *types.HashTable = arr
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()
		if _z.IsType(types.IS_INDIRECT) {
			_z = _z.GetZv()
			if _z.IsType(types.IS_UNDEF) {
				continue
			}
		}
		var_name = _p.GetKey()
		entry = _z
		if var_name == nil {
			continue
		}
		if PhpValidVarName(var_name.GetVal(), var_name.GetLen()) == 0 {
			continue
		}
		if types.ZendStringEqualsLiteral(var_name, "this") {
			continue
		}
		orig_var = symbol_table.KeyFind(var_name.GetStr())
		if orig_var != nil {
			if orig_var.IsType(types.IS_INDIRECT) {
				orig_var = orig_var.GetZv()
				if orig_var.IsType(types.IS_UNDEF) {
					if entry.IsReference() {
						entry.AddRefcount()
					} else {
						types.ZVAL_MAKE_REF_EX(entry, 2)
					}
					orig_var.SetReference(entry.GetRef())
					count++
				}
			}
		} else {
			if entry.IsReference() {
				entry.AddRefcount()
			} else {
				types.ZVAL_MAKE_REF_EX(entry, 2)
			}
			symbol_table.KeyAddNew(var_name.GetStr(), entry)
			count++
		}
	}
	return count
}
func PhpExtractSkip(arr *types.ZendArray, symbol_table *types.ZendArray) zend.ZendLong {
	var count zend.ZendLong = 0
	var var_name *types.ZendString
	var entry *types.Zval
	var orig_var *types.Zval
	var __ht *types.HashTable = arr
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()
		if _z.IsType(types.IS_INDIRECT) {
			_z = _z.GetZv()
			if _z.IsType(types.IS_UNDEF) {
				continue
			}
		}
		var_name = _p.GetKey()
		entry = _z
		if var_name == nil {
			continue
		}
		if PhpValidVarName(var_name.GetVal(), var_name.GetLen()) == 0 {
			continue
		}
		if types.ZendStringEqualsLiteral(var_name, "this") {
			continue
		}
		orig_var = symbol_table.KeyFind(var_name.GetStr())
		if orig_var != nil {
			if orig_var.IsType(types.IS_INDIRECT) {
				orig_var = orig_var.GetZv()
				if orig_var.IsType(types.IS_UNDEF) {
					types.ZVAL_COPY_DEREF(orig_var, entry)
					count++
				}
			}
		} else {
			entry = types.ZVAL_DEREF(entry)
			entry.TryAddRefcount()
			symbol_table.KeyAddNew(var_name.GetStr(), entry)
			count++
		}
	}
	return count
}
func ZifExtract(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var var_array_param *types.Zval
	var prefix *types.Zval = nil
	var extract_refs zend.ZendLong
	var extract_type zend.ZendLong = EXTR_OVERWRITE
	var count zend.ZendLong
	var symbol_table *types.ZendArray
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 3
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			var_array_param = fp.ParseArrayEx2(false, true, false)
			fp.StartOptional()
			extract_type = fp.ParseLong()
			argparse.Z_PARAM_PROLOGUE(0, 0)
			argparse.ZendParseArgZvalDeref(_arg, &prefix, 0)
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	extract_refs = extract_type & EXTR_REFS
	if extract_refs != 0 {
		types.SEPARATE_ARRAY(var_array_param)
	}
	extract_type &= 0xff
	if extract_type < EXTR_OVERWRITE || extract_type > EXTR_IF_EXISTS {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Invalid extract type")
		return
	}
	if extract_type > EXTR_SKIP && extract_type <= EXTR_PREFIX_IF_EXISTS && executeData.NumArgs() < 3 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "specified extract type requires the prefix parameter")
		return
	}
	if prefix != nil {
		if zend.TryConvertToString(prefix) == 0 {
			return
		}
		if prefix.GetStr().GetLen() != 0 && PhpValidVarName(prefix.GetStr().GetVal(), prefix.GetStr().GetLen()) == 0 {
			core.PhpErrorDocref(nil, zend.E_WARNING, "prefix is not a valid identifier")
			return
		}
	}
	if zend.ZendForbidDynamicCall("extract()") == types.FAILURE {
		return
	}
	symbol_table = zend.ZendRebuildSymbolTable()
	if extract_refs != 0 {
		switch extract_type {
		case EXTR_IF_EXISTS:
			count = PhpExtractRefIfExists(var_array_param.GetArr(), symbol_table)
		case EXTR_OVERWRITE:
			count = PhpExtractRefOverwrite(var_array_param.GetArr(), symbol_table)
		case EXTR_PREFIX_IF_EXISTS:
			count = PhpExtractRefPrefixIfExists(var_array_param.GetArr(), symbol_table, prefix)
		case EXTR_PREFIX_SAME:
			count = PhpExtractRefPrefixSame(var_array_param.GetArr(), symbol_table, prefix)
		case EXTR_PREFIX_ALL:
			count = PhpExtractRefPrefixAll(var_array_param.GetArr(), symbol_table, prefix)
		case EXTR_PREFIX_INVALID:
			count = PhpExtractRefPrefixInvalid(var_array_param.GetArr(), symbol_table, prefix)
		default:
			count = PhpExtractRefSkip(var_array_param.GetArr(), symbol_table)
		}
	} else {

		/* The array might be stored in a local variable that will be overwritten */

		var array_copy types.Zval
		types.ZVAL_COPY(&array_copy, var_array_param)
		switch extract_type {
		case EXTR_IF_EXISTS:
			count = PhpExtractIfExists(array_copy.GetArr(), symbol_table)
		case EXTR_OVERWRITE:
			count = PhpExtractOverwrite(array_copy.GetArr(), symbol_table)
		case EXTR_PREFIX_IF_EXISTS:
			count = PhpExtractPrefixIfExists(array_copy.GetArr(), symbol_table, prefix)
		case EXTR_PREFIX_SAME:
			count = PhpExtractPrefixSame(array_copy.GetArr(), symbol_table, prefix)
		case EXTR_PREFIX_ALL:
			count = PhpExtractPrefixAll(array_copy.GetArr(), symbol_table, prefix)
		case EXTR_PREFIX_INVALID:
			count = PhpExtractPrefixInvalid(array_copy.GetArr(), symbol_table, prefix)
		default:
			count = PhpExtractSkip(array_copy.GetArr(), symbol_table)
		}
		zend.ZvalPtrDtor(&array_copy)
	}
	return_value.SetLong(count)
	return
}
func PhpCompactVar(eg_active_symbol_table *types.HashTable, return_value *types.Zval, entry *types.Zval) {
	var value_ptr *types.Zval
	var data types.Zval
	entry = types.ZVAL_DEREF(entry)
	if entry.IsType(types.IS_STRING) {
		if b.Assign(&value_ptr, zend.ZendHashFindInd(eg_active_symbol_table, entry.GetStr())) != nil {
			value_ptr = types.ZVAL_DEREF(value_ptr)
			value_ptr.TryAddRefcount()
			return_value.GetArr().KeyUpdate(entry.GetStr().GetStr(), value_ptr)
		} else if types.ZendStringEqualsLiteral(entry.GetStr(), "this") {
			var object *types.ZendObject = zend.ZendGetThisObject(zend.CurrEX())
			if object != nil {
				object.AddRefcount()
				data.SetObject(object)
				return_value.GetArr().KeyUpdate(entry.GetStr().GetStr(), &data)
			}
		} else {
			core.PhpErrorDocref(nil, zend.E_NOTICE, "Undefined variable: %s", types.Z_STR_P(entry).GetVal())
		}
	} else if entry.IsType(types.IS_ARRAY) {
		if entry.IsRefcounted() {
			if entry.IsRecursive() {
				core.PhpErrorDocref(nil, zend.E_WARNING, "recursion detected")
				return
			}
			entry.ProtectRecursive()
		}
		var __ht *types.HashTable = entry.GetArr()
		for _, _p := range __ht.foreachData() {
			var _z *types.Zval = _p.GetVal()
			if _z.IsType(types.IS_INDIRECT) {
				_z = _z.GetZv()
				if _z.IsType(types.IS_UNDEF) {
					continue
				}
			}
			value_ptr = _z
			PhpCompactVar(eg_active_symbol_table, return_value, value_ptr)
		}
		if entry.IsRefcounted() {
			entry.UnprotectRecursive()
		}
	}
}
func ZifCompact(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var args *types.Zval = nil
	var num_args uint32
	var i uint32
	var symbol_table *types.ZendArray
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = -1
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			var _num_varargs int = _num_args - _i - 0
			if _num_varargs > 0 {
				args = _real_arg + 1
				num_args = _num_varargs
				_i += _num_varargs
				_real_arg += _num_varargs
			} else {
				args = nil
				num_args = 0
			}
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	if zend.ZendForbidDynamicCall("compact()") == types.FAILURE {
		return
	}
	symbol_table = zend.ZendRebuildSymbolTable()
	if symbol_table == nil {
		return
	}

	/* compact() is probably most used with a single array of var_names
	   or multiple string names, rather than a combination of both.
	   So quickly guess a minimum result size based on that */

	if num_args != 0 && args[0].IsType(types.IS_ARRAY) {
		zend.ArrayInitSize(return_value, types.Z_ARRVAL(args[0]).GetNNumOfElements())
	} else {
		zend.ArrayInitSize(return_value, num_args)
	}
	for i = 0; i < num_args; i++ {
		PhpCompactVar(symbol_table, return_value, &args[i])
	}
}
func ZifArrayFill(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var val *types.Zval
	var start_key zend.ZendLong
	var num zend.ZendLong
	for {
		var _flags int = 0
		var _min_num_args int = 3
		var _max_num_args int = 3
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			start_key = fp.ParseLong()
			num = fp.ParseLong()
			argparse.Z_PARAM_PROLOGUE(0, 0)
			argparse.ZendParseArgZvalDeref(_arg, &val, 0)
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	if num > 0 {
		if b.SizeOf("num") > 4 && num > 0x7fffffff {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Too many elements")
			return_value.SetFalse()
			return
		} else if start_key > zend.ZEND_LONG_MAX-num+1 {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Cannot add element to the array as the next element is already occupied")
			return_value.SetFalse()
			return
		} else if start_key >= 0 && start_key < num {

			/* create packed array */

			var p *types.Bucket
			var n zend.ZendLong
			zend.ArrayInitSize(return_value, uint32(start_key+num))
			zend.ZendHashRealInitPacked(return_value.GetArr())
			types.Z_ARRVAL_P(return_value).SetNNumUsed(uint32(start_key + num))
			types.Z_ARRVAL_P(return_value).SetNNumOfElements(uint32(num))
			types.Z_ARRVAL_P(return_value).SetNNextFreeElement(zend_long(start_key + num))
			if val.IsRefcounted() {
				val.GetCounted().AddRefcountEx(uint32(num))
			}
			p = types.Z_ARRVAL_P(return_value).GetArData()
			n = start_key
			for b.PostDec(&start_key) {
				p.GetVal().SetUndef()
				p++
			}
			for b.PostDec(&num) {
				types.ZVAL_COPY_VALUE(p.GetVal(), val)
				n++
				p.SetH(n - 1)
				p.SetKey(nil)
				p++
			}
		} else {

			/* create hash */

			zend.ArrayInitSize(return_value, uint32(num))
			zend.ZendHashRealInitMixed(return_value.GetArr())
			if val.IsRefcounted() {
				val.GetCounted().AddRefcountEx(uint32(num))
			}
			return_value.GetArr().IndexAddNewH(start_key, val)
			for b.PreDec(&num) {
				return_value.GetArr().NextIndexInsertNew(val)
				start_key++
			}
		}
	} else if num == 0 {
		zend.ZVAL_EMPTY_ARRAY(return_value)
		return
	} else {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Number of elements can't be negative")
		return_value.SetFalse()
		return
	}
}
func ZifArrayFillKeys(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var keys *types.Zval
	var val *types.Zval
	var entry *types.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 2
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			keys = fp.ParseArray()
			argparse.Z_PARAM_PROLOGUE(0, 0)
			argparse.ZendParseArgZvalDeref(_arg, &val, 0)
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}

	/* Initialize return array */

	zend.ArrayInitSize(return_value, types.Z_ARRVAL_P(keys).GetNNumOfElements())
	var __ht *types.HashTable = keys.GetArr()
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()

		entry = _z
		entry = types.ZVAL_DEREF(entry)
		val.TryAddRefcount()
		if entry.IsType(types.IS_LONG) {
			return_value.GetArr().IndexUpdateH(entry.GetLval(), val)
		} else {
			var tmp_key *types.ZendString
			var key *types.ZendString = zend.ZvalGetTmpString(entry, &tmp_key)
			return_value.GetArr().SymtableUpdate(key.GetStr(), val)
			zend.ZendTmpStringRelease(tmp_key)
		}
	}
}
func ZifRange(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var zlow *types.Zval
	var zhigh *types.Zval
	var zstep *types.Zval = nil
	var tmp types.Zval
	var err int = 0
	var is_step_double int = 0
	var step float64 = 1.0
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 3
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			argparse.Z_PARAM_PROLOGUE(0, 0)
			argparse.ZendParseArgZvalDeref(_arg, &zlow, 0)
			argparse.Z_PARAM_PROLOGUE(0, 0)
			argparse.ZendParseArgZvalDeref(_arg, &zhigh, 0)
			fp.StartOptional()
			argparse.Z_PARAM_PROLOGUE(0, 0)
			argparse.ZendParseArgZvalDeref(_arg, &zstep, 0)
			if fp.HasError() {
				fp.HandleError()
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	if zstep != nil {
		if zstep.IsType(types.IS_DOUBLE) {
			is_step_double = 1
		} else if zstep.IsType(types.IS_STRING) {
			var type_ int = zend.IsNumericString(zstep.GetStr().GetStr(), nil, nil, 0)
			if type_ == types.IS_DOUBLE {
				is_step_double = 1
			}
			if type_ == 0 {

				/* bad number */

				core.PhpErrorDocref(nil, zend.E_WARNING, "Invalid range string - must be numeric")
				return_value.SetFalse()
				return
			}
		}
		step = zend.ZvalGetDouble(zstep)

		/* We only want positive step values. */

		if step < 0.0 {
			step *= -1
		}

		/* We only want positive step values. */

	}

	/* If the range is given as strings, generate an array of characters. */

	if zlow.IsType(types.IS_STRING) && zhigh.IsType(types.IS_STRING) && zlow.GetStr().GetLen() >= 1 && zhigh.GetStr().GetLen() >= 1 {
		var type1 int
		var type2 int
		var low uint8
		var high uint8
		var lstep zend.ZendLong = zend.ZendLong(step)
		type1 = zend.IsNumericString(zlow.GetStr().GetStr(), nil, nil, 0)
		type2 = zend.IsNumericString(zhigh.GetStr().GetStr(), nil, nil, 0)
		if type1 == types.IS_DOUBLE || type2 == types.IS_DOUBLE || is_step_double != 0 {
			goto double_str
		} else if type1 == types.IS_LONG || type2 == types.IS_LONG {
			goto long_str
		}
		low = uint8(zlow.GetStr().GetVal()[0])
		high = uint8(zhigh.GetStr().GetVal()[0])
		if low > high {
			if lstep <= 0 {
				err = 1
				goto err
			}

			/* Initialize the return_value as an array. */

			zend.ArrayInitSize(return_value, uint32((low-high)/lstep+1))
			zend.ZendHashRealInitPacked(return_value.GetArr())
			for {
				var __fill_ht *types.HashTable = return_value.GetArr()
				var __fill_bkt *types.Bucket = __fill_ht.GetArData() + __fill_ht.GetNNumUsed()
				var __fill_idx uint32 = __fill_ht.GetNNumUsed()
				b.Assert(__fill_ht.HasUFlags(types.HASH_FLAG_PACKED))
				for ; low >= high; low -= uint(lstep) {
					__fill_bkt.GetVal().SetInternedString(types.ZSTR_CHAR(low))
					__fill_bkt.SetH(__fill_idx)
					__fill_bkt.SetKey(nil)
					__fill_bkt++
					__fill_idx++
					if signed__int(low-lstep) < 0 {
						break
					}
				}
				__fill_ht.SetNNumUsed(__fill_idx)
				__fill_ht.SetNNumOfElements(__fill_idx)
				__fill_ht.SetNNextFreeElement(__fill_idx)
				__fill_ht.SetNInternalPointer(0)
				break
			}
		} else if high > low {
			if lstep <= 0 {
				err = 1
				goto err
			}
			zend.ArrayInitSize(return_value, uint32((high-low)/lstep+1))
			zend.ZendHashRealInitPacked(return_value.GetArr())
			for {
				var __fill_ht *types.HashTable = return_value.GetArr()
				var __fill_bkt *types.Bucket = __fill_ht.GetArData() + __fill_ht.GetNNumUsed()
				var __fill_idx uint32 = __fill_ht.GetNNumUsed()
				b.Assert(__fill_ht.HasUFlags(types.HASH_FLAG_PACKED))
				for ; low <= high; low += uint(lstep) {
					__fill_bkt.GetVal().SetInternedString(types.ZSTR_CHAR(low))
					__fill_bkt.SetH(__fill_idx)
					__fill_bkt.SetKey(nil)
					__fill_bkt++
					__fill_idx++
					if signed__int(low+lstep) > 255 {
						break
					}
				}
				__fill_ht.SetNNumUsed(__fill_idx)
				__fill_ht.SetNNumOfElements(__fill_idx)
				__fill_ht.SetNNextFreeElement(__fill_idx)
				__fill_ht.SetNInternalPointer(0)
				break
			}
		} else {
			zend.ArrayInit(return_value)
			tmp.SetInternedString(types.ZSTR_CHAR(low))
			return_value.GetArr().NextIndexInsertNew(&tmp)
		}
	} else if zlow.IsType(types.IS_DOUBLE) || zhigh.IsType(types.IS_DOUBLE) || is_step_double != 0 {
		var low float64
		var high float64
		var element float64
		var i uint32
		var size uint32
	double_str:
		low = zend.ZvalGetDouble(zlow)
		high = zend.ZvalGetDouble(zhigh)
		if core.ZendIsInf(high) || core.ZendIsInf(low) {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Invalid range supplied: start=%0.0f end=%0.0f", low, high)
			return_value.SetFalse()
			return
		}
		if low > high {
			if low-high < step || step <= 0 {
				err = 1
				goto err
			}
			var __calc_size float64 = (low-high)/step + 1
			if __calc_size >= float64(types.HT_MAX_SIZE) {
				core.PhpErrorDocref(nil, zend.E_WARNING, "The supplied range exceeds the maximum array size: start=%0.0f end=%0.0f", high, low)
				return_value.SetFalse()
				return
			}
			size = uint32(_phpMathRound(__calc_size, 0, PHP_ROUND_HALF_UP))
			zend.ArrayInitSize(return_value, size)
			zend.ZendHashRealInitPacked(return_value.GetArr())
			var __fill_ht *types.HashTable = return_value.GetArr()
			var __fill_bkt *types.Bucket = __fill_ht.GetArData() + __fill_ht.GetNNumUsed()
			var __fill_idx uint32 = __fill_ht.GetNNumUsed()
			b.Assert(__fill_ht.HasUFlags(types.HASH_FLAG_PACKED))
			i = 0
			element = low
			for i < size && element >= high {
				__fill_bkt.GetVal().SetDouble(element)
				__fill_bkt.SetH(__fill_idx)
				__fill_bkt.SetKey(nil)
				__fill_bkt++
				__fill_idx++
				i++
				element = low - i*step
			}
			__fill_ht.SetNNumUsed(__fill_idx)
			__fill_ht.SetNNumOfElements(__fill_idx)
			__fill_ht.SetNNextFreeElement(__fill_idx)
			__fill_ht.SetNInternalPointer(0)
		} else if high > low {
			if high-low < step || step <= 0 {
				err = 1
				goto err
			}
			var __calc_size float64 = (high-low)/step + 1
			if __calc_size >= float64(types.HT_MAX_SIZE) {
				core.PhpErrorDocref(nil, zend.E_WARNING, "The supplied range exceeds the maximum array size: start=%0.0f end=%0.0f", low, high)
				return_value.SetFalse()
				return
			}
			size = uint32(_phpMathRound(__calc_size, 0, PHP_ROUND_HALF_UP))
			zend.ArrayInitSize(return_value, size)
			zend.ZendHashRealInitPacked(return_value.GetArr())
			var __fill_ht *types.HashTable = return_value.GetArr()
			var __fill_bkt *types.Bucket = __fill_ht.GetArData() + __fill_ht.GetNNumUsed()
			var __fill_idx uint32 = __fill_ht.GetNNumUsed()
			b.Assert(__fill_ht.HasUFlags(types.HASH_FLAG_PACKED))
			i = 0
			element = low
			for i < size && element <= high {
				__fill_bkt.GetVal().SetDouble(element)
				__fill_bkt.SetH(__fill_idx)
				__fill_bkt.SetKey(nil)
				__fill_bkt++
				__fill_idx++
				i++
				element = low + i*step
			}
			__fill_ht.SetNNumUsed(__fill_idx)
			__fill_ht.SetNNumOfElements(__fill_idx)
			__fill_ht.SetNNextFreeElement(__fill_idx)
			__fill_ht.SetNInternalPointer(0)
		} else {
			zend.ArrayInit(return_value)
			tmp.SetDouble(low)
			return_value.GetArr().NextIndexInsertNew(&tmp)
		}
	} else {
		var low zend.ZendLong
		var high zend.ZendLong

		/* lstep is a zend_ulong so that comparisons to it don't overflow, i.e. low - high < lstep */

		var lstep zend.ZendUlong
		var i uint32
		var size uint32
	long_str:
		low = zend.ZvalGetLong(zlow)
		high = zend.ZvalGetLong(zhigh)
		if step <= 0 {
			err = 1
			goto err
		}
		lstep = zend.ZendUlong(step)
		if step <= 0 {
			err = 1
			goto err
		}
		if low > high {
			if zend.ZendUlong(low-high < lstep) != 0 {
				err = 1
				goto err
			}
			var __calc_size zend.ZendUlong = zend.ZendUlong(low-high) / lstep
			if __calc_size >= types.HT_MAX_SIZE-1 {
				core.PhpErrorDocref(nil, zend.E_WARNING, "The supplied range exceeds the maximum array size: start="+zend.ZEND_LONG_FMT+" end="+zend.ZEND_LONG_FMT, high, low)
				return_value.SetFalse()
				return
			}
			size = uint32(__calc_size + 1)
			zend.ArrayInitSize(return_value, size)
			zend.ZendHashRealInitPacked(return_value.GetArr())
			var __fill_ht *types.HashTable = return_value.GetArr()
			var __fill_bkt *types.Bucket = __fill_ht.GetArData() + __fill_ht.GetNNumUsed()
			var __fill_idx uint32 = __fill_ht.GetNNumUsed()
			b.Assert(__fill_ht.HasUFlags(types.HASH_FLAG_PACKED))
			for i = 0; i < size; i++ {
				__fill_bkt.GetVal().SetLong(low - i*lstep)
				__fill_bkt.SetH(__fill_idx)
				__fill_bkt.SetKey(nil)
				__fill_bkt++
				__fill_idx++
			}
			__fill_ht.SetNNumUsed(__fill_idx)
			__fill_ht.SetNNumOfElements(__fill_idx)
			__fill_ht.SetNNextFreeElement(__fill_idx)
			__fill_ht.SetNInternalPointer(0)
		} else if high > low {
			if zend.ZendUlong(high-low < lstep) != 0 {
				err = 1
				goto err
			}
			var __calc_size zend.ZendUlong = zend.ZendUlong(high-low) / lstep
			if __calc_size >= types.HT_MAX_SIZE-1 {
				core.PhpErrorDocref(nil, zend.E_WARNING, "The supplied range exceeds the maximum array size: start="+zend.ZEND_LONG_FMT+" end="+zend.ZEND_LONG_FMT, low, high)
				return_value.SetFalse()
				return
			}
			size = uint32(__calc_size + 1)
			zend.ArrayInitSize(return_value, size)
			zend.ZendHashRealInitPacked(return_value.GetArr())
			var __fill_ht *types.HashTable = return_value.GetArr()
			var __fill_bkt *types.Bucket = __fill_ht.GetArData() + __fill_ht.GetNNumUsed()
			var __fill_idx uint32 = __fill_ht.GetNNumUsed()
			b.Assert(__fill_ht.HasUFlags(types.HASH_FLAG_PACKED))
			for i = 0; i < size; i++ {
				__fill_bkt.GetVal().SetLong(low + i*lstep)
				__fill_bkt.SetH(__fill_idx)
				__fill_bkt.SetKey(nil)
				__fill_bkt++
				__fill_idx++
			}
			__fill_ht.SetNNumUsed(__fill_idx)
			__fill_ht.SetNNumOfElements(__fill_idx)
			__fill_ht.SetNNextFreeElement(__fill_idx)
			__fill_ht.SetNInternalPointer(0)
		} else {
			zend.ArrayInit(return_value)
			tmp.SetLong(low)
			return_value.GetArr().NextIndexInsertNew(&tmp)
		}
	}
err:
	if err != 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "step exceeds the specified range")
		return_value.SetFalse()
		return
	}
}
func PhpArrayDataShuffle(array *types.Zval) {
	var idx uint32
	var j uint32
	var n_elems uint32
	var p *types.Bucket
	var temp types.Bucket
	var hash *types.HashTable
	var rnd_idx zend.ZendLong
	var n_left uint32
	n_elems = types.Z_ARRVAL_P(array).GetNNumOfElements()
	if n_elems < 1 {
		return
	}
	hash = array.GetArr()
	n_left = n_elems
	if !(hash.HasIterators()) {
		if hash.GetNNumUsed() != hash.GetNNumOfElements() {
			j = 0
			idx = 0
			for ; idx < hash.GetNNumUsed(); idx++ {
				p = hash.GetArData() + idx
				if p.GetVal().IsType(types.IS_UNDEF) {
					continue
				}
				if j != idx {
					hash.GetArData()[j] = *p
				}
				j++
			}
		}
		for b.PreDec(&n_left) {
			rnd_idx = PhpMtRandRange(0, n_left)
			if rnd_idx != n_left {
				temp = hash.GetArData()[n_left]
				hash.GetArData()[n_left] = hash.GetArData()[rnd_idx]
				hash.GetArData()[rnd_idx] = temp
			}
		}
	} else {
		var iter_pos uint32 = zend.ZendHashIteratorsLowerPos(hash, 0)
		if hash.GetNNumUsed() != hash.GetNNumOfElements() {
			j = 0
			idx = 0
			for ; idx < hash.GetNNumUsed(); idx++ {
				p = hash.GetArData() + idx
				if p.GetVal().IsType(types.IS_UNDEF) {
					continue
				}
				if j != idx {
					hash.GetArData()[j] = *p
					if idx == iter_pos {
						zend.ZendHashIteratorsUpdate(hash, idx, j)
						iter_pos = zend.ZendHashIteratorsLowerPos(hash, iter_pos+1)
					}
				}
				j++
			}
		}
		for b.PreDec(&n_left) {
			rnd_idx = PhpMtRandRange(0, n_left)
			if rnd_idx != n_left {
				temp = hash.GetArData()[n_left]
				hash.GetArData()[n_left] = hash.GetArData()[rnd_idx]
				hash.GetArData()[rnd_idx] = temp
				zend.ZendHashIteratorsUpdate(hash, uint32(rnd_idx), n_left)
			}
		}
	}
	hash.SetNNumUsed(n_elems)
	hash.SetNInternalPointer(0)
	for j = 0; j < n_elems; j++ {
		p = hash.GetArData() + j
		if p.GetKey() != nil {
			types.ZendStringReleaseEx(p.GetKey(), 0)
		}
		p.SetH(j)
		p.SetKey(nil)
	}
	hash.SetNNextFreeElement(n_elems)
	if !hash.HasUFlags(types.HASH_FLAG_PACKED) {
		zend.ZendHashToPacked(hash)
	}
}
func ZifShuffle(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var array *types.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			array = fp.ParseArrayEx(false, true)
			if fp.HasError() {
				fp.HandleError()
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	PhpArrayDataShuffle(array)
	return_value.SetTrue()
	return
}
func PhpSplice(in_hash *types.HashTable, offset zend.ZendLong, length zend.ZendLong, replace *types.HashTable, removed *types.HashTable) {
	var out_hash types.HashTable
	var num_in zend.ZendLong
	var pos zend.ZendLong
	var idx uint32
	var p *types.Bucket
	var entry *types.Zval
	var iter_pos uint32 = zend.ZendHashIteratorsLowerPos(in_hash, 0)

	/* Get number of entries in the input hash */

	num_in = in_hash.GetNNumOfElements()

	/* Clamp the offset.. */

	if offset > num_in {
		offset = num_in
	} else if offset < 0 && b.Assign(&offset, num_in+offset) < 0 {
		offset = 0
	}

	/* ..and the length */

	if length < 0 {
		length = num_in - offset + length
	} else if unsigned(offset+unsigned(length)) > unsigned(num_in) {
		length = num_in - offset
	}

	/* Create and initialize output hash */

	zend.ZendHashInit(&out_hash, b.Cond(length > 0, num_in-length, 0)+b.CondF1(replace != nil, func() __auto__ { return replace.GetNNumOfElements() }, 0), nil, zend.ZVAL_PTR_DTOR, 0)

	/* Start at the beginning of the input hash and copy entries to output hash until offset is reached */

	pos = 0
	idx = 0
	for ; pos < offset && idx < in_hash.GetNNumUsed(); idx++ {
		p = in_hash.GetArData() + idx
		if p.GetVal().IsType(types.IS_UNDEF) {
			continue
		}

		/* Get entry and increase reference count */

		entry = p.GetVal()

		/* Update output hash depending on key type */

		if p.GetKey() == nil {
			out_hash.NextIndexInsertNew(entry)
		} else {
			out_hash.KeyAddNew(p.GetKey().GetStr(), entry)
		}
		if idx == iter_pos {
			if zend.ZendLong(idx != pos) != 0 {
				zend.ZendHashIteratorsUpdate(in_hash, idx, pos)
			}
			iter_pos = zend.ZendHashIteratorsLowerPos(in_hash, iter_pos+1)
		}
		pos++
	}

	/* If hash for removed entries exists, go until offset+length and copy the entries to it */

	if removed != nil {
		for ; pos < offset+length && idx < in_hash.GetNNumUsed(); idx++ {
			p = in_hash.GetArData() + idx
			if p.GetVal().IsType(types.IS_UNDEF) {
				continue
			}
			pos++
			entry = p.GetVal()
			entry.TryAddRefcount()
			if p.GetKey() == nil {
				removed.NextIndexInsertNew(entry)
				zend.ZendHashDelBucket(in_hash, p)
			} else {
				removed.KeyAddNew(p.GetKey().GetStr(), entry)
				if in_hash == zend.EG__().GetSymbolTable() {
					zend.ZendDeleteGlobalVariable(p.GetKey())
				} else {
					zend.ZendHashDelBucket(in_hash, p)
				}
			}
		}
	} else {
		var pos2 int = pos
		for ; pos2 < offset+length && idx < in_hash.GetNNumUsed(); idx++ {
			p = in_hash.GetArData() + idx
			if p.GetVal().IsType(types.IS_UNDEF) {
				continue
			}
			pos2++
			if p.GetKey() != nil && in_hash == zend.EG__().GetSymbolTable() {
				zend.ZendDeleteGlobalVariable(p.GetKey())
			} else {
				zend.ZendHashDelBucket(in_hash, p)
			}
		}
	}
	iter_pos = zend.ZendHashIteratorsLowerPos(in_hash, iter_pos)

	/* If there are entries to insert.. */

	if replace != nil {
		var __ht *types.HashTable = replace
		for _, _p := range __ht.foreachData() {
			var _z *types.Zval = _p.GetVal()
			if _z.IsType(types.IS_INDIRECT) {
				_z = _z.GetZv()
				if _z.IsType(types.IS_UNDEF) {
					continue
				}
			}
			entry = _z
			entry.TryAddRefcount()
			out_hash.NextIndexInsertNew(entry)
			pos++
		}
	}

	/* Copy the remaining input hash entries to the output hash */

	for ; idx < in_hash.GetNNumUsed(); idx++ {
		p = in_hash.GetArData() + idx
		if p.GetVal().IsType(types.IS_UNDEF) {
			continue
		}
		entry = p.GetVal()
		if p.GetKey() == nil {
			out_hash.NextIndexInsertNew(entry)
		} else {
			out_hash.KeyAddNew(p.GetKey().GetStr(), entry)
		}
		if idx == iter_pos {
			if zend.ZendLong(idx != pos) != 0 {
				zend.ZendHashIteratorsUpdate(in_hash, idx, pos)
			}
			iter_pos = zend.ZendHashIteratorsLowerPos(in_hash, iter_pos+1)
		}
		pos++
	}

	/* replace HashTable data */

	out_hash.SetNIteratorsCount(in_hash.GetNIteratorsCount())
	in_hash.SetNIteratorsCount(0)
	in_hash.SetPDestructor(nil)
	in_hash.Destroy()
	in_hash.SetUFlags(out_hash.GetUFlags())
	in_hash.SetNTableSize(out_hash.GetNTableSize())
	in_hash.SetNTableMask(out_hash.GetNTableMask())
	in_hash.SetNNumUsed(out_hash.GetNNumUsed())
	in_hash.SetNNumOfElements(out_hash.GetNNumOfElements())
	in_hash.SetNNextFreeElement(out_hash.GetNNextFreeElement())
	in_hash.SetArData(out_hash.GetArData())
	in_hash.SetPDestructor(out_hash.GetPDestructor())
	zend.ZendHashInternalPointerReset(in_hash)
}
func ZifArrayPush(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var args *types.Zval
	var stack *types.Zval
	var new_var types.Zval
	var i int
	var argc int
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = -1
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			stack = fp.ParseArrayEx(false, true)
			var _num_varargs int = _num_args - _i - 0
			if _num_varargs > 0 {
				args = _real_arg + 1
				argc = _num_varargs
				_i += _num_varargs
				_real_arg += _num_varargs
			} else {
				args = nil
				argc = 0
			}
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}

	/* For each subsequent argument, make it a reference, increase refcount, and add it to the end of the array */

	for i = 0; i < argc; i++ {
		types.ZVAL_COPY(&new_var, &args[i])
		if stack.GetArr().NextIndexInsert(&new_var) == nil {
			new_var.TryDelRefcount()
			core.PhpErrorDocref(nil, zend.E_WARNING, "Cannot add element to the array as the next element is already occupied")
			return_value.SetFalse()
			return
		}
	}

	/* Clean up and return the number of values in the stack */

	return_value.SetLong(types.Z_ARRVAL_P(stack).GetNNumOfElements())

	/* Clean up and return the number of values in the stack */
}
func ZifArrayPop(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var stack *types.Zval
	var val *types.Zval
	var idx uint32
	var p *types.Bucket
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			stack = fp.ParseArrayEx(false, true)
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	if types.Z_ARRVAL_P(stack).GetNNumOfElements() == 0 {
		return
	}

	/* Get the last value and copy it into the return value */

	idx = types.Z_ARRVAL_P(stack).GetNNumUsed()
	for true {
		if idx == 0 {
			return
		}
		idx--
		p = types.Z_ARRVAL_P(stack).GetArData() + idx
		val = p.GetVal()
		if val.IsType(types.IS_INDIRECT) {
			val = val.GetZv()
		}
		if val.GetType() != types.IS_UNDEF {
			break
		}
	}
	types.ZVAL_COPY_DEREF(return_value, val)
	if p.GetKey() == nil && types.Z_ARRVAL_P(stack).GetNNextFreeElement() > 0 && p.GetH() >= zend_ulong(types.Z_ARRVAL_P(stack).GetNNextFreeElement()-1) {
		types.Z_ARRVAL_P(stack).SetNNextFreeElement(types.Z_ARRVAL_P(stack).GetNNextFreeElement() - 1)
	}

	/* Delete the last value */

	if p.GetKey() != nil && stack.GetArr() == zend.EG__().GetSymbolTable() {
		zend.ZendDeleteGlobalVariable(p.GetKey())
	} else {
		zend.ZendHashDelBucket(stack.GetArr(), p)
	}
	zend.ZendHashInternalPointerReset(stack.GetArr())
}
func ZifArrayShift(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var stack *types.Zval
	var val *types.Zval
	var idx uint32
	var p *types.Bucket
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			stack = fp.ParseArrayEx(false, true)
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	if types.Z_ARRVAL_P(stack).GetNNumOfElements() == 0 {
		return
	}

	/* Get the first value and copy it into the return value */

	idx = 0
	for true {
		if idx == types.Z_ARRVAL_P(stack).GetNNumUsed() {
			return
		}
		p = types.Z_ARRVAL_P(stack).GetArData() + idx
		val = p.GetVal()
		if val.IsType(types.IS_INDIRECT) {
			val = val.GetZv()
		}
		if val.GetType() != types.IS_UNDEF {
			break
		}
		idx++
	}
	types.ZVAL_COPY_DEREF(return_value, val)

	/* Delete the first value */

	if p.GetKey() != nil && stack.GetArr() == zend.EG__().GetSymbolTable() {
		zend.ZendDeleteGlobalVariable(p.GetKey())
	} else {
		zend.ZendHashDelBucket(stack.GetArr(), p)
	}

	/* re-index like it did before */

	if types.Z_ARRVAL_P(stack).HasUFlags(types.HASH_FLAG_PACKED) {
		var k uint32 = 0
		if !(stack.GetArr().HasIterators()) {
			for idx = 0; idx < types.Z_ARRVAL_P(stack).GetNNumUsed(); idx++ {
				p = types.Z_ARRVAL_P(stack).GetArData() + idx
				if p.GetVal().IsType(types.IS_UNDEF) {
					continue
				}
				if idx != k {
					var q *types.Bucket = types.Z_ARRVAL_P(stack).GetArData() + k
					q.SetH(k)
					q.SetKey(nil)
					types.ZVAL_COPY_VALUE(q.GetVal(), p.GetVal())
					p.GetVal().SetUndef()
				}
				k++
			}
		} else {
			var iter_pos uint32 = zend.ZendHashIteratorsLowerPos(stack.GetArr(), 0)
			for idx = 0; idx < types.Z_ARRVAL_P(stack).GetNNumUsed(); idx++ {
				p = types.Z_ARRVAL_P(stack).GetArData() + idx
				if p.GetVal().IsType(types.IS_UNDEF) {
					continue
				}
				if idx != k {
					var q *types.Bucket = types.Z_ARRVAL_P(stack).GetArData() + k
					q.SetH(k)
					q.SetKey(nil)
					types.ZVAL_COPY_VALUE(q.GetVal(), p.GetVal())
					p.GetVal().SetUndef()
					if idx == iter_pos {
						zend.ZendHashIteratorsUpdate(stack.GetArr(), idx, k)
						iter_pos = zend.ZendHashIteratorsLowerPos(stack.GetArr(), iter_pos+1)
					}
				}
				k++
			}
		}
		types.Z_ARRVAL_P(stack).SetNNumUsed(k)
		types.Z_ARRVAL_P(stack).SetNNextFreeElement(k)
	} else {
		var k uint32 = 0
		var should_rehash int = 0
		for idx = 0; idx < types.Z_ARRVAL_P(stack).GetNNumUsed(); idx++ {
			p = types.Z_ARRVAL_P(stack).GetArData() + idx
			if p.GetVal().IsType(types.IS_UNDEF) {
				continue
			}
			if p.GetKey() == nil {
				if p.GetH() != k {
					k++
					p.SetH(k - 1)
					should_rehash = 1
				} else {
					k++
				}
			}
		}
		types.Z_ARRVAL_P(stack).SetNNextFreeElement(k)
		if should_rehash != 0 {
			stack.GetArr().Rehash()
		}
	}
	zend.ZendHashInternalPointerReset(stack.GetArr())
}
func ZifArrayUnshift(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var args *types.Zval
	var stack *types.Zval
	var new_hash types.HashTable
	var argc int
	var i int
	var key *types.ZendString
	var value *types.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = -1
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			stack = fp.ParseArrayEx(false, true)
			var _num_varargs int = _num_args - _i - 0
			if _num_varargs > 0 {
				args = _real_arg + 1
				argc = _num_varargs
				_i += _num_varargs
				_real_arg += _num_varargs
			} else {
				args = nil
				argc = 0
			}
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	zend.ZendHashInit(&new_hash, types.Z_ARRVAL_P(stack).GetNNumOfElements()+argc, nil, zend.ZVAL_PTR_DTOR, 0)
	for i = 0; i < argc; i++ {
		args[i].TryAddRefcount()
		new_hash.NextIndexInsertNew(&args[i])
	}
	var __ht *types.HashTable = stack.GetArr()
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()

		key = _p.GetKey()
		value = _z
		if key != nil {
			new_hash.KeyAddNew(key.GetStr(), value)
		} else {
			new_hash.NextIndexInsertNew(value)
		}
	}
	if stack.GetArr().HasIterators() {
		zend.ZendHashIteratorsAdvance(stack.GetArr(), argc)
		new_hash.SetNIteratorsCount(types.Z_ARRVAL_P(stack).GetNIteratorsCount())
		stack.GetArr().SetNIteratorsCount(0)
	}

	/* replace HashTable data */

	types.Z_ARRVAL_P(stack).SetPDestructor(nil)
	stack.GetArr().Destroy()
	types.Z_ARRVAL_P(stack).SetUFlags(new_hash.GetUFlags())
	types.Z_ARRVAL_P(stack).SetNTableSize(new_hash.GetNTableSize())
	types.Z_ARRVAL_P(stack).SetNTableMask(new_hash.GetNTableMask())
	types.Z_ARRVAL_P(stack).SetNNumUsed(new_hash.GetNNumUsed())
	types.Z_ARRVAL_P(stack).SetNNumOfElements(new_hash.GetNNumOfElements())
	types.Z_ARRVAL_P(stack).SetNNextFreeElement(new_hash.GetNNextFreeElement())
	types.Z_ARRVAL_P(stack).SetArData(new_hash.GetArData())
	types.Z_ARRVAL_P(stack).SetPDestructor(new_hash.GetPDestructor())
	zend.ZendHashInternalPointerReset(stack.GetArr())

	/* Clean up and return the number of elements in the stack */

	return_value.SetLong(types.Z_ARRVAL_P(stack).GetNNumOfElements())

	/* Clean up and return the number of elements in the stack */
}
func ZifArraySplice(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var array *types.Zval
	var repl_array *types.Zval = nil
	var rem_hash *types.HashTable = nil
	var offset zend.ZendLong
	var length zend.ZendLong = 0
	var num_in int
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 4
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			array = fp.ParseArrayEx(false, true)
			offset = fp.ParseLong()
			fp.StartOptional()
			length = fp.ParseLong()
			argparse.Z_PARAM_PROLOGUE(0, 0)
			argparse.ZendParseArgZvalDeref(_arg, &repl_array, 0)
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	num_in = types.Z_ARRVAL_P(array).GetNNumOfElements()
	if executeData.NumArgs() < 3 {
		length = num_in
	}
	if executeData.NumArgs() == 4 {

		/* Make sure the last argument, if passed, is an array */

		zend.ConvertToArrayEx(repl_array)

		/* Make sure the last argument, if passed, is an array */

	}

	/* Don't create the array of removed elements if it's not going
	 * to be used; e.g. only removing and/or replacing elements */

	if zend.USED_RET() {
		var size zend.ZendLong = length

		/* Clamp the offset.. */

		if offset > num_in {
			offset = num_in
		} else if offset < 0 && b.Assign(&offset, num_in+offset) < 0 {
			offset = 0
		}

		/* ..and the length */

		if length < 0 {
			size = num_in - offset + length
		} else if zend.ZendUlong(offset+zend.ZendUlong(length)) > uint32(num_in) {
			size = num_in - offset
		}

		/* Initialize return value */

		zend.ArrayInitSize(return_value, b.CondF1(size > 0, func() uint32 { return uint32(size) }, 0))
		rem_hash = return_value.GetArr()
	}

	/* Perform splice */

	PhpSplice(array.GetArr(), offset, length, b.CondF1(repl_array != nil, func() *types.ZendArray { return repl_array.GetArr() }, nil), rem_hash)

	/* Perform splice */
}
func ZifArraySlice(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var input *types.Zval
	var z_length *types.Zval = nil
	var entry *types.Zval
	var offset zend.ZendLong
	var length zend.ZendLong = 0
	var preserve_keys types.ZendBool = 0
	var num_in int
	var pos int
	var string_key *types.ZendString
	var num_key zend.ZendUlong
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 4
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			input = fp.ParseArray()
			offset = fp.ParseLong()
			fp.StartOptional()
			argparse.Z_PARAM_PROLOGUE(0, 0)
			argparse.ZendParseArgZvalDeref(_arg, &z_length, 0)
			preserve_keys = fp.ParseBool()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}

	/* Get number of entries in the input hash */

	num_in = types.Z_ARRVAL_P(input).GetNNumOfElements()

	/* We want all entries from offset to the end if length is not passed or is null */

	if executeData.NumArgs() < 3 || z_length.IsType(types.IS_NULL) {
		length = num_in
	} else {
		length = zend.ZvalGetLong(z_length)
	}

	/* Clamp the offset.. */

	if offset > num_in {
		zend.ZVAL_EMPTY_ARRAY(return_value)
		return
	} else if offset < 0 && b.Assign(&offset, num_in+offset) < 0 {
		offset = 0
	}

	/* ..and the length */

	if length < 0 {
		length = num_in - offset + length
	} else if zend.ZendUlong(offset+zend.ZendUlong(length)) > unsigned(num_in) {
		length = num_in - offset
	}
	if length <= 0 {
		zend.ZVAL_EMPTY_ARRAY(return_value)
		return
	}

	/* Initialize returned array */

	zend.ArrayInitSize(return_value, uint32(length))

	/* Start at the beginning and go until we hit offset */

	pos = 0
	var __ht *types.HashTable = types.Z_ARRVAL_P(input)
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()

		num_key = _p.GetH()
		string_key = _p.GetKey()
		entry = _z
		pos++
		if pos <= offset {
			continue
		}
		if pos > offset+length {
			break
		}
		if string_key != nil {
			entry = return_value.GetArr().KeyAddNew(string_key.GetStr(), entry)
		} else {
			if preserve_keys != 0 {
				entry = return_value.GetArr().IndexAddNewH(num_key, entry)
			} else {
				entry = return_value.GetArr().NextIndexInsertNew(entry)
			}
		}
		zend.ZvalAddRef(entry)
	}
}
func PhpArrayMergeRecursive(dest *types.HashTable, src *types.HashTable) int {
	var src_entry *types.Zval
	var dest_entry *types.Zval
	var string_key *types.ZendString
	var __ht *types.HashTable = src
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()

		string_key = _p.GetKey()
		src_entry = _z
		if string_key != nil {
			if b.Assign(&dest_entry, dest.KeyFind(string_key.GetStr())) != nil {
				var src_zval *types.Zval = src_entry
				var dest_zval *types.Zval = dest_entry
				var thash *types.HashTable
				var tmp types.Zval
				var ret int
				src_zval = types.ZVAL_DEREF(src_zval)
				dest_zval = types.ZVAL_DEREF(dest_zval)
				if dest_zval.IsType(types.IS_ARRAY) {
					thash = dest_zval.GetArr()
				} else {
					thash = nil
				}
				if thash != nil && thash.IsRecursive() || src_entry == dest_entry && dest_entry.IsReference() && dest_entry.GetRefcount()%2 != 0 {
					core.PhpErrorDocref(nil, zend.E_WARNING, "recursion detected")
					return 0
				}
				b.Assert(!(dest_entry.IsReference()) || dest_entry.GetRefcount() > 1)
				types.SEPARATE_ZVAL(dest_entry)
				dest_zval = dest_entry
				if dest_zval.IsType(types.IS_NULL) {
					zend.ConvertToArrayEx(dest_zval)
					zend.AddNextIndexNull(dest_zval)
				} else {
					zend.ConvertToArrayEx(dest_zval)
				}
				tmp.SetUndef()
				if src_zval.IsType(types.IS_OBJECT) {
					types.ZVAL_COPY(&tmp, src_zval)
					zend.ConvertToArray(&tmp)
					src_zval = &tmp
				}
				if src_zval.IsType(types.IS_ARRAY) {
					if thash != nil && (thash.GetGcFlags()&types.GC_IMMUTABLE) == 0 {
						thash.ProtectRecursive()
					}
					ret = PhpArrayMergeRecursive(dest_zval.GetArr(), src_zval.GetArr())
					if thash != nil && (thash.GetGcFlags()&types.GC_IMMUTABLE) == 0 {
						thash.UnprotectRecursive()
					}
					if ret == 0 {
						return 0
					}
				} else {
					src_zval.TryAddRefcount()
					dest_zval.GetArr().NextIndexInsert(src_zval)
				}
				zend.ZvalPtrDtor(&tmp)
			} else {
				var zv *types.Zval = dest.KeyAddNew(string_key.GetStr(), src_entry)
				zend.ZvalAddRef(zv)
			}
		} else {
			var zv *types.Zval = dest.NextIndexInsert(src_entry)
			zend.ZvalAddRef(zv)
		}
	}
	return 1
}
func PhpArrayMerge(dest *types.HashTable, src *types.HashTable) int {
	var src_entry *types.Zval
	var string_key *types.ZendString
	if dest.HasUFlags(types.HASH_FLAG_PACKED) && src.HasUFlags(types.HASH_FLAG_PACKED) {
		dest.Extend(dest.GetNNumOfElements() + src.GetNNumOfElements())
		var __fill_ht *types.HashTable = dest
		var __fill_bkt *types.Bucket = __fill_ht.GetArData() + __fill_ht.GetNNumUsed()
		var __fill_idx uint32 = __fill_ht.GetNNumUsed()
		b.Assert(__fill_ht.HasUFlags(types.HASH_FLAG_PACKED))
		var __ht *types.HashTable = src
		for _, _p := range __ht.foreachData() {
			var _z *types.Zval = _p.GetVal()

			src_entry = _z
			if src_entry.IsReference() && src_entry.GetRefcount() == 1 {
				src_entry = types.Z_REFVAL_P(src_entry)
			}
			src_entry.TryAddRefcount()
			types.ZVAL_COPY_VALUE(__fill_bkt.GetVal(), src_entry)
			__fill_bkt.SetH(__fill_idx)
			__fill_bkt.SetKey(nil)
			__fill_bkt++
			__fill_idx++
		}
		__fill_ht.SetNNumUsed(__fill_idx)
		__fill_ht.SetNNumOfElements(__fill_idx)
		__fill_ht.SetNNextFreeElement(__fill_idx)
		__fill_ht.SetNInternalPointer(0)
	} else {
		var __ht *types.HashTable = src
		for _, _p := range __ht.foreachData() {
			var _z *types.Zval = _p.GetVal()

			string_key = _p.GetKey()
			src_entry = _z
			if src_entry.IsReference() && src_entry.GetRefcount() == 1 {
				src_entry = types.Z_REFVAL_P(src_entry)
			}
			src_entry.TryAddRefcount()
			if string_key != nil {
				dest.KeyUpdate(string_key.GetStr(), src_entry)
			} else {
				dest.NextIndexInsertNew(src_entry)
			}
		}
	}
	return 1
}
func PhpArrayReplaceRecursive(dest *types.HashTable, src *types.HashTable) int {
	var src_entry *types.Zval
	var dest_entry *types.Zval
	var src_zval *types.Zval
	var dest_zval *types.Zval
	var string_key *types.ZendString
	var num_key zend.ZendUlong
	var ret int
	var __ht *types.HashTable = src
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()

		num_key = _p.GetH()
		string_key = _p.GetKey()
		src_entry = _z
		src_zval = src_entry
		src_zval = types.ZVAL_DEREF(src_zval)
		if string_key != nil {
			if src_zval.GetType() != types.IS_ARRAY || b.Assign(&dest_entry, dest.KeyFind(string_key.GetStr())) == nil || dest_entry.GetType() != types.IS_ARRAY && (!(dest_entry.IsReference()) || types.Z_REFVAL_P(dest_entry).GetType() != types.IS_ARRAY) {
				var zv *types.Zval = dest.KeyUpdate(string_key.GetStr(), src_entry)
				zend.ZvalAddRef(zv)
				continue
			}
		} else {
			if src_zval.GetType() != types.IS_ARRAY || b.Assign(&dest_entry, dest.IndexFindH(num_key)) == nil || dest_entry.GetType() != types.IS_ARRAY && (!(dest_entry.IsReference()) || types.Z_REFVAL_P(dest_entry).GetType() != types.IS_ARRAY) {
				var zv *types.Zval = dest.IndexUpdateH(num_key, src_entry)
				zend.ZvalAddRef(zv)
				continue
			}
		}
		dest_zval = dest_entry
		dest_zval = types.ZVAL_DEREF(dest_zval)
		if dest_zval.IsRecursive() || src_zval.IsRecursive() || src_entry.IsReference() && dest_entry.IsReference() && src_entry.GetRef() == dest_entry.GetRef() && dest_entry.GetRefcount()%2 != 0 {
			core.PhpErrorDocref(nil, zend.E_WARNING, "recursion detected")
			return 0
		}
		b.Assert(!(dest_entry.IsReference()) || dest_entry.GetRefcount() > 1)
		types.SEPARATE_ZVAL(dest_entry)
		dest_zval = dest_entry
		if dest_zval.IsRefcounted() {
			dest_zval.ProtectRecursive()
		}
		if src_zval.IsRefcounted() {
			src_zval.ProtectRecursive()
		}
		ret = PhpArrayReplaceRecursive(dest_zval.GetArr(), src_zval.GetArr())
		if dest_zval.IsRefcounted() {
			dest_zval.UnprotectRecursive()
		}
		if src_zval.IsRefcounted() {
			src_zval.UnprotectRecursive()
		}
		if ret == 0 {
			return 0
		}
	}
	return 1
}
func PhpArrayReplaceWrapper(executeData *zend.ZendExecuteData, return_value *types.Zval, recursive int) {
	var args *types.Zval = nil
	var arg *types.Zval
	var argc int
	var i int
	var dest *types.HashTable
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = -1
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			var _num_varargs int = _num_args - _i - 0
			if _num_varargs > 0 {
				args = _real_arg + 1
				argc = _num_varargs
				_i += _num_varargs
				_real_arg += _num_varargs
			} else {
				args = nil
				argc = 0
			}
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	for i = 0; i < argc; i++ {
		var arg *types.Zval = args + i
		if arg.GetType() != types.IS_ARRAY {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Expected parameter %d to be an array, %s given", i+1, zend.ZendZvalTypeName(arg))
			return_value.SetNull()
			return
		}
	}

	/* copy first array */

	arg = args
	dest = zend.ZendArrayDup(arg.GetArr())
	return_value.SetArray(dest)
	if recursive != 0 {
		for i = 1; i < argc; i++ {
			arg = args + i
			PhpArrayReplaceRecursive(dest, arg.GetArr())
		}
	} else {
		for i = 1; i < argc; i++ {
			arg = args + i
			zend.ZendHashMerge(dest, arg.GetArr(), zend.ZvalAddRef, 1)
		}
	}
}
func PhpArrayMergeWrapper(executeData *zend.ZendExecuteData, return_value *types.Zval, recursive int) {
	var args *types.Zval = nil
	var arg *types.Zval
	var argc int
	var i int
	var src_entry *types.Zval
	var src *types.HashTable
	var dest *types.HashTable
	var count uint32 = 0
	for {
		var _flags int = 0
		var _min_num_args int = 0
		var _max_num_args int = -1
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			var _num_varargs int = _num_args - _i - 0
			if _num_varargs > 0 {
				args = _real_arg + 1
				argc = _num_varargs
				_i += _num_varargs
				_real_arg += _num_varargs
			} else {
				args = nil
				argc = 0
			}
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	if argc == 0 {
		zend.ZVAL_EMPTY_ARRAY(return_value)
		return
	}
	for i = 0; i < argc; i++ {
		var arg *types.Zval = args + i
		if arg.GetType() != types.IS_ARRAY {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Expected parameter %d to be an array, %s given", i+1, zend.ZendZvalTypeName(arg))
			return_value.SetNull()
			return
		}
		count += types.Z_ARRVAL_P(arg).GetNNumOfElements()
	}
	if argc == 2 {
		var ret *types.Zval = nil
		if types.Z_ARRVAL(args[0]).GetNNumOfElements() == 0 {
			ret = &args[1]
		} else if types.Z_ARRVAL(args[1]).GetNNumOfElements() == 0 {
			ret = &args[0]
		}
		if ret != nil {
			if types.Z_ARRVAL_P(ret).HasUFlags(types.HASH_FLAG_PACKED) {
				if ret.GetArr().IsWithoutHoles() {
					types.ZVAL_COPY(return_value, ret)
					return
				}
			} else {
				var copy types.ZendBool = 1
				var string_key *types.ZendString
				var __ht *types.HashTable = ret.GetArr()
				for _, _p := range __ht.foreachData() {
					var _z *types.Zval = _p.GetVal()

					string_key = _p.GetKey()
					if string_key == nil {
						copy = 0
						break
					}
				}
				if copy != 0 {
					types.ZVAL_COPY(return_value, ret)
					return
				}
			}
		}
	}
	arg = args
	src = arg.GetArr()

	/* copy first array */

	zend.ArrayInitSize(return_value, count)
	dest = return_value.GetArr()
	if src.HasUFlags(types.HASH_FLAG_PACKED) {
		zend.ZendHashRealInitPacked(dest)
		var __fill_ht *types.HashTable = dest
		var __fill_bkt *types.Bucket = __fill_ht.GetArData() + __fill_ht.GetNNumUsed()
		var __fill_idx uint32 = __fill_ht.GetNNumUsed()
		b.Assert(__fill_ht.HasUFlags(types.HASH_FLAG_PACKED))
		var __ht *types.HashTable = src
		for _, _p := range __ht.foreachData() {
			var _z *types.Zval = _p.GetVal()

			src_entry = _z
			if src_entry.IsReference() && src_entry.GetRefcount() == 1 {
				src_entry = types.Z_REFVAL_P(src_entry)
			}
			src_entry.TryAddRefcount()
			types.ZVAL_COPY_VALUE(__fill_bkt.GetVal(), src_entry)
			__fill_bkt.SetH(__fill_idx)
			__fill_bkt.SetKey(nil)
			__fill_bkt++
			__fill_idx++
		}
		__fill_ht.SetNNumUsed(__fill_idx)
		__fill_ht.SetNNumOfElements(__fill_idx)
		__fill_ht.SetNNextFreeElement(__fill_idx)
		__fill_ht.SetNInternalPointer(0)
	} else {
		var string_key *types.ZendString
		zend.ZendHashRealInitMixed(dest)
		var __ht *types.HashTable = src
		for _, _p := range __ht.foreachData() {
			var _z *types.Zval = _p.GetVal()

			string_key = _p.GetKey()
			src_entry = _z
			if src_entry.IsReference() && src_entry.GetRefcount() == 1 {
				src_entry = types.Z_REFVAL_P(src_entry)
			}
			src_entry.TryAddRefcount()
			if string_key != nil {
				zend._zendHashAppend(dest, string_key, src_entry)
			} else {
				dest.NextIndexInsertNew(src_entry)
			}
		}
	}
	if recursive != 0 {
		for i = 1; i < argc; i++ {
			arg = args + i
			PhpArrayMergeRecursive(dest, arg.GetArr())
		}
	} else {
		for i = 1; i < argc; i++ {
			arg = args + i
			PhpArrayMerge(dest, arg.GetArr())
		}
	}
}
func ZifArrayMerge(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	PhpArrayMergeWrapper(executeData, return_value, 0)
}
func ZifArrayMergeRecursive(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	PhpArrayMergeWrapper(executeData, return_value, 1)
}
func ZifArrayReplace(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	PhpArrayReplaceWrapper(executeData, return_value, 0)
}
func ZifArrayReplaceRecursive(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	PhpArrayReplaceWrapper(executeData, return_value, 1)
}
func ZifArrayKeys(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var input *types.Zval
	var search_value *types.Zval = nil
	var entry *types.Zval
	var new_val types.Zval
	var strict types.ZendBool = 0
	var num_idx zend.ZendUlong
	var str_idx *types.ZendString
	var arrval *types.ZendArray
	var elem_count zend.ZendUlong
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 3
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			input = fp.ParseArray()
			fp.StartOptional()
			argparse.Z_PARAM_PROLOGUE(0, 0)
			argparse.ZendParseArgZvalDeref(_arg, &search_value, 0)
			strict = fp.ParseBool()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	arrval = input.GetArr()
	elem_count = arrval.GetNNumOfElements()

	/* Base case: empty input */

	if elem_count == 0 {
		zend.ZVAL_ZVAL(return_value, input, 1, 0)
		return
	}

	/* Initialize return array */

	if search_value != nil {
		zend.ArrayInit(return_value)
		if strict != 0 {
			var __ht *types.HashTable = arrval
			for _, _p := range __ht.foreachData() {
				var _z *types.Zval = _p.GetVal()
				if _z.IsType(types.IS_INDIRECT) {
					_z = _z.GetZv()
					if _z.IsType(types.IS_UNDEF) {
						continue
					}
				}
				num_idx = _p.GetH()
				str_idx = _p.GetKey()
				entry = _z
				entry = types.ZVAL_DEREF(entry)
				if zend.FastIsIdenticalFunction(search_value, entry) != 0 {
					if str_idx != nil {
						new_val.SetStringCopy(str_idx)
					} else {
						new_val.SetLong(num_idx)
					}
					return_value.GetArr().NextIndexInsertNew(&new_val)
				}
			}
		} else {
			var __ht *types.HashTable = arrval
			for _, _p := range __ht.foreachData() {
				var _z *types.Zval = _p.GetVal()
				if _z.IsType(types.IS_INDIRECT) {
					_z = _z.GetZv()
					if _z.IsType(types.IS_UNDEF) {
						continue
					}
				}
				num_idx = _p.GetH()
				str_idx = _p.GetKey()
				entry = _z
				if zend.FastEqualCheckFunction(search_value, entry) != 0 {
					if str_idx != nil {
						new_val.SetStringCopy(str_idx)
					} else {
						new_val.SetLong(num_idx)
					}
					return_value.GetArr().NextIndexInsertNew(&new_val)
				}
			}
		}
	} else {
		zend.ArrayInitSize(return_value, elem_count)
		zend.ZendHashRealInitPacked(return_value.GetArr())
		var __fill_ht *types.HashTable = return_value.GetArr()
		var __fill_bkt *types.Bucket = __fill_ht.GetArData() + __fill_ht.GetNNumUsed()
		var __fill_idx uint32 = __fill_ht.GetNNumUsed()
		b.Assert(__fill_ht.HasUFlags(types.HASH_FLAG_PACKED))

		/* Go through input array and add keys to the return array */

		var __ht *types.HashTable = input.GetArr()
		for _, _p := range __ht.foreachData() {
			var _z *types.Zval = _p.GetVal()
			if _z.IsType(types.IS_INDIRECT) {
				_z = _z.GetZv()
				if _z.IsType(types.IS_UNDEF) {
					continue
				}
			}
			num_idx = _p.GetH()
			str_idx = _p.GetKey()
			entry = _z
			if str_idx != nil {
				__fill_bkt.GetVal().SetStringCopy(str_idx)
			} else {
				__fill_bkt.GetVal().SetLong(num_idx)
			}
			__fill_bkt.SetH(__fill_idx)
			__fill_bkt.SetKey(nil)
			__fill_bkt++
			__fill_idx++
		}

		/* Go through input array and add keys to the return array */

		__fill_ht.SetNNumUsed(__fill_idx)
		__fill_ht.SetNNumOfElements(__fill_idx)
		__fill_ht.SetNNextFreeElement(__fill_idx)
		__fill_ht.SetNInternalPointer(0)
	}

	/* Initialize return array */
}
func ZifArrayKeyFirst(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var stack *types.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			stack = fp.ParseArray()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	var target_hash *types.HashTable = stack.GetArr()
	var pos types.HashPosition = 0
	zend.ZendHashGetCurrentKeyZvalEx(target_hash, return_value, &pos)
}
func ZifArrayKeyLast(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var stack *types.Zval
	var pos types.HashPosition
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			stack = fp.ParseArray()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	var target_hash *types.HashTable = stack.GetArr()
	zend.ZendHashInternalPointerEndEx(target_hash, &pos)
	zend.ZendHashGetCurrentKeyZvalEx(target_hash, return_value, &pos)
}
func ZifArrayValues(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var input *types.Zval
	var entry *types.Zval
	var arrval *types.ZendArray
	var arrlen zend.ZendLong
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			input = fp.ParseArray()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	arrval = input.GetArr()

	/* Return empty input as is */

	arrlen = arrval.GetNNumOfElements()
	if arrlen == 0 {
		zend.ZVAL_EMPTY_ARRAY(return_value)
		return
	}

	/* Return vector-like packed arrays as-is */

	/* Initialize return array */

	zend.ArrayInitSize(return_value, arrval.GetNNumOfElements())
	zend.ZendHashRealInitPacked(return_value.GetArr())

	/* Go through input array and add values to the return array */

	var __fill_ht *types.HashTable = return_value.GetArr()
	var __fill_bkt *types.Bucket = __fill_ht.GetArData() + __fill_ht.GetNNumUsed()
	var __fill_idx uint32 = __fill_ht.GetNNumUsed()
	b.Assert(__fill_ht.HasUFlags(types.HASH_FLAG_PACKED))
	var __ht *types.HashTable = arrval
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()

		entry = _z
		if entry.IsReference() && entry.GetRefcount() == 1 {
			entry = types.Z_REFVAL_P(entry)
		}
		entry.TryAddRefcount()
		types.ZVAL_COPY_VALUE(__fill_bkt.GetVal(), entry)
		__fill_bkt.SetH(__fill_idx)
		__fill_bkt.SetKey(nil)
		__fill_bkt++
		__fill_idx++
	}
	__fill_ht.SetNNumUsed(__fill_idx)
	__fill_ht.SetNNumOfElements(__fill_idx)
	__fill_ht.SetNNextFreeElement(__fill_idx)
	__fill_ht.SetNInternalPointer(0)

	/* Go through input array and add values to the return array */
}
func ZifArrayCountValues(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var input *types.Zval
	var entry *types.Zval
	var tmp *types.Zval
	var myht *types.HashTable
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			input = fp.ParseArray()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}

	/* Initialize return array */

	zend.ArrayInit(return_value)

	/* Go through input array and add values to the return array */

	myht = input.GetArr()
	var __ht *types.HashTable = myht
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()

		entry = _z
		entry = types.ZVAL_DEREF(entry)
		if entry.IsType(types.IS_LONG) {
			if b.Assign(&tmp, return_value.GetArr().IndexFindH(entry.GetLval())) == nil {
				var data types.Zval
				data.SetLong(1)
				return_value.GetArr().IndexUpdateH(entry.GetLval(), &data)
			} else {
				tmp.GetLval()++
			}
		} else if entry.IsType(types.IS_STRING) {
			if b.Assign(&tmp, return_value.GetArr().SymtableFind(entry.GetStr().GetStr())) == nil {
				var data types.Zval
				data.SetLong(1)
				return_value.GetArr().SymtableUpdate(entry.GetStr().GetStr(), &data)
			} else {
				tmp.GetLval()++
			}
		} else {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Can only count STRING and INTEGER values!")
		}
	}
}
func ArrayColumnParamHelper(param *types.Zval, name string) types.ZendBool {
	switch param.GetType() {
	case types.IS_DOUBLE:
		if param.GetType() != types.IS_LONG {
			zend.ConvertToLong(param)
		}
		fallthrough
	case types.IS_LONG:
		return 1
	case types.IS_OBJECT:
		if zend.TryConvertToString(param) == 0 {
			return 0
		}
		fallthrough
	case types.IS_STRING:
		return 1
	default:
		core.PhpErrorDocref(nil, zend.E_WARNING, "The %s key should be either a string or an integer", name)
		return 0
	}
}
func ArrayColumnFetchProp(data *types.Zval, name *types.Zval, rv *types.Zval) *types.Zval {
	var prop *types.Zval = nil
	if data.IsType(types.IS_OBJECT) {

		/* The has_property check is first performed in "exists" mode (which returns true for
		 * properties that are null but exist) and then in "has" mode to handle objects that
		 * implement __isset (which is not called in "exists" mode). */

		if types.Z_OBJ_HT(*data).GetHasProperty()(data, name, zend.ZEND_PROPERTY_EXISTS, nil) != 0 || types.Z_OBJ_HT(*data).GetHasProperty()(data, name, zend.ZEND_PROPERTY_ISSET, nil) != 0 {
			prop = types.Z_OBJ_HT(*data).GetReadProperty()(data, name, zend.BP_VAR_R, nil, rv)
			if prop != nil {
				prop = types.ZVAL_DEREF(prop)
				if prop != rv {
					prop.TryAddRefcount()
				}
			}
		}

		/* The has_property check is first performed in "exists" mode (which returns true for
		 * properties that are null but exist) and then in "has" mode to handle objects that
		 * implement __isset (which is not called in "exists" mode). */

	} else if data.IsType(types.IS_ARRAY) {
		if name.IsType(types.IS_STRING) {
			prop = data.GetArr().SymtableFind(name.GetStr().GetStr())
		} else if name.IsType(types.IS_LONG) {
			prop = data.GetArr().IndexFindH(name.GetLval())
		}
		if prop != nil {
			prop = types.ZVAL_DEREF(prop)
			prop.TryAddRefcount()
		}
	}
	return prop
}
func ZifArrayColumn(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var input *types.HashTable
	var colval *types.Zval
	var data *types.Zval
	var rv types.Zval
	var column *types.Zval = nil
	var index *types.Zval = nil
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 3
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			input = fp.ParseArrayHt()
			argparse.Z_PARAM_PROLOGUE(0, 0)
			argparse.ZendParseArgZvalDeref(_arg, &column, 1)
			fp.StartOptional()
			argparse.Z_PARAM_PROLOGUE(0, 0)
			argparse.ZendParseArgZvalDeref(_arg, &index, 1)
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	if column != nil && ArrayColumnParamHelper(column, "column") == 0 || index != nil && ArrayColumnParamHelper(index, "index") == 0 {
		return_value.SetFalse()
		return
	}
	zend.ArrayInitSize(return_value, input.GetNNumOfElements())
	if index == nil {
		zend.ZendHashRealInitPacked(return_value.GetArr())
		var __fill_ht *types.HashTable = return_value.GetArr()
		var __fill_bkt *types.Bucket = __fill_ht.GetArData() + __fill_ht.GetNNumUsed()
		var __fill_idx uint32 = __fill_ht.GetNNumUsed()
		b.Assert(__fill_ht.HasUFlags(types.HASH_FLAG_PACKED))
		var __ht *types.HashTable = input
		for _, _p := range __ht.foreachData() {
			var _z *types.Zval = _p.GetVal()

			data = _z
			data = types.ZVAL_DEREF(data)
			if column == nil {
				data.TryAddRefcount()
				colval = data
			} else if b.Assign(&colval, ArrayColumnFetchProp(data, column, &rv)) == nil {
				continue
			}
			types.ZVAL_COPY_VALUE(__fill_bkt.GetVal(), colval)
			__fill_bkt.SetH(__fill_idx)
			__fill_bkt.SetKey(nil)
			__fill_bkt++
			__fill_idx++
		}
		__fill_ht.SetNNumUsed(__fill_idx)
		__fill_ht.SetNNumOfElements(__fill_idx)
		__fill_ht.SetNNextFreeElement(__fill_idx)
		__fill_ht.SetNInternalPointer(0)
	} else {
		var __ht *types.HashTable = input
		for _, _p := range __ht.foreachData() {
			var _z *types.Zval = _p.GetVal()

			data = _z
			data = types.ZVAL_DEREF(data)
			if column == nil {
				data.TryAddRefcount()
				colval = data
			} else if b.Assign(&colval, ArrayColumnFetchProp(data, column, &rv)) == nil {
				continue
			}

			/* Failure will leave keyval alone which will land us on the final else block below
			 * which is to append the value as next_index
			 */

			if index != nil {
				var rv types.Zval
				var keyval *types.Zval = ArrayColumnFetchProp(data, index, &rv)
				if keyval != nil {
					switch keyval.GetType() {
					case types.IS_STRING:
						return_value.GetArr().SymtableUpdate(keyval.GetStr().GetStr(), colval)
					case types.IS_LONG:
						return_value.GetArr().IndexUpdateH(keyval.GetLval(), colval)
					case types.IS_OBJECT:
						var tmp_key *types.ZendString
						var key *types.ZendString = zend.ZvalGetTmpString(keyval, &tmp_key)
						return_value.GetArr().SymtableUpdate(key.GetStr(), colval)
						zend.ZendTmpStringRelease(tmp_key)
					case types.IS_NULL:
						return_value.GetArr().KeyUpdate(types.ZSTR_EMPTY_ALLOC().GetStr(), colval)
					case types.IS_DOUBLE:
						return_value.GetArr().IndexUpdateH(zend.ZendDvalToLval(keyval.GetDval()), colval)
					case types.IS_TRUE:
						return_value.GetArr().IndexUpdateH(1, colval)
					case types.IS_FALSE:
						return_value.GetArr().IndexUpdateH(0, colval)
					case types.IS_RESOURCE:
						return_value.GetArr().IndexUpdateH(types.Z_RES_HANDLE_P(keyval), colval)
					default:
						return_value.GetArr().NextIndexInsert(colval)
					}
					zend.ZvalPtrDtor(keyval)
				} else {
					return_value.GetArr().NextIndexInsert(colval)
				}
			} else {
				return_value.GetArr().NextIndexInsert(colval)
			}

			/* Failure will leave keyval alone which will land us on the final else block below
			 * which is to append the value as next_index
			 */

		}
	}
}
func ZifArrayReverse(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var input *types.Zval
	var entry *types.Zval
	var string_key *types.ZendString
	var num_key zend.ZendUlong
	var preserve_keys types.ZendBool = 0
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			input = fp.ParseArray()
			fp.StartOptional()
			preserve_keys = fp.ParseBool()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}

	/* Initialize return array */

	zend.ArrayInitSize(return_value, types.Z_ARRVAL_P(input).GetNNumOfElements())
	if types.Z_ARRVAL_P(input).HasUFlags(types.HASH_FLAG_PACKED) && preserve_keys == 0 {
		zend.ZendHashRealInitPacked(return_value.GetArr())
		var __fill_ht *types.HashTable = return_value.GetArr()
		var __fill_bkt *types.Bucket = __fill_ht.GetArData() + __fill_ht.GetNNumUsed()
		var __fill_idx uint32 = __fill_ht.GetNNumUsed()
		b.Assert(__fill_ht.HasUFlags(types.HASH_FLAG_PACKED))
		var __ht *types.HashTable = input.GetArr()
		for _, _p := range __ht.foreachDataReserve() {
			var _z types.Zval = _p.GetVal()

			entry = _z
			if entry.IsReference() && entry.GetRefcount() == 1 {
				entry = types.Z_REFVAL_P(entry)
			}
			entry.TryAddRefcount()
			types.ZVAL_COPY_VALUE(__fill_bkt.GetVal(), entry)
			__fill_bkt.SetH(__fill_idx)
			__fill_bkt.SetKey(nil)
			__fill_bkt++
			__fill_idx++
		}
		__fill_ht.SetNNumUsed(__fill_idx)
		__fill_ht.SetNNumOfElements(__fill_idx)
		__fill_ht.SetNNextFreeElement(__fill_idx)
		__fill_ht.SetNInternalPointer(0)
	} else {
		var __ht *types.HashTable = input.GetArr()
		for _, _p := range __ht.foreachDataReserve() {
			var _z types.Zval = _p.GetVal()

			num_key = _p.GetH()
			string_key = _p.GetKey()
			entry = _z
			if string_key != nil {
				entry = return_value.GetArr().KeyAddNew(string_key.GetStr(), entry)
			} else {
				if preserve_keys != 0 {
					entry = return_value.GetArr().IndexAddNewH(num_key, entry)
				} else {
					entry = return_value.GetArr().NextIndexInsertNew(entry)
				}
			}
			zend.ZvalAddRef(entry)
		}
	}
}
func ZifArrayPad(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var input *types.Zval
	var pad_value *types.Zval
	var pad_size zend.ZendLong
	var pad_size_abs zend.ZendLong
	var input_size zend.ZendLong
	var num_pads zend.ZendLong
	var i zend.ZendLong
	var key *types.ZendString
	var value *types.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 3
		var _max_num_args int = 3
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			input = fp.ParseArray()
			pad_size = fp.ParseLong()
			argparse.Z_PARAM_PROLOGUE(0, 0)
			argparse.ZendParseArgZvalDeref(_arg, &pad_value, 0)
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}

	/* Do some initial calculations */

	input_size = types.Z_ARRVAL_P(input).GetNNumOfElements()
	pad_size_abs = zend.ZEND_ABS(pad_size)
	if pad_size_abs < 0 || pad_size_abs-input_size > int64(1048576) {
		core.PhpErrorDocref(nil, zend.E_WARNING, "You may only pad up to 1048576 elements at a time")
		return_value.SetFalse()
		return
	}
	if input_size >= pad_size_abs {

		/* Copy the original array */

		types.ZVAL_COPY(return_value, input)
		return
	}
	num_pads = pad_size_abs - input_size
	if pad_value.IsRefcounted() {
		pad_value.GetCounted().AddRefcountEx(num_pads)
	}
	zend.ArrayInitSize(return_value, pad_size_abs)
	if types.Z_ARRVAL_P(input).HasUFlags(types.HASH_FLAG_PACKED) {
		zend.ZendHashRealInitPacked(return_value.GetArr())
		if pad_size < 0 {
			var __fill_ht *types.HashTable = return_value.GetArr()
			var __fill_bkt *types.Bucket = __fill_ht.GetArData() + __fill_ht.GetNNumUsed()
			var __fill_idx uint32 = __fill_ht.GetNNumUsed()
			b.Assert(__fill_ht.HasUFlags(types.HASH_FLAG_PACKED))
			for i = 0; i < num_pads; i++ {
				types.ZVAL_COPY_VALUE(__fill_bkt.GetVal(), pad_value)
				__fill_bkt.SetH(__fill_idx)
				__fill_bkt.SetKey(nil)
				__fill_bkt++
				__fill_idx++
			}
			__fill_ht.SetNNumUsed(__fill_idx)
			__fill_ht.SetNNumOfElements(__fill_idx)
			__fill_ht.SetNNextFreeElement(__fill_idx)
			__fill_ht.SetNInternalPointer(0)
		}
		var __fill_ht *types.HashTable = return_value.GetArr()
		var __fill_bkt *types.Bucket = __fill_ht.GetArData() + __fill_ht.GetNNumUsed()
		var __fill_idx uint32 = __fill_ht.GetNNumUsed()
		b.Assert(__fill_ht.HasUFlags(types.HASH_FLAG_PACKED))
		var __ht *types.HashTable = input.GetArr()
		for _, _p := range __ht.foreachData() {
			var _z *types.Zval = _p.GetVal()

			value = _z
			value.TryAddRefcount()
			types.ZVAL_COPY_VALUE(__fill_bkt.GetVal(), value)
			__fill_bkt.SetH(__fill_idx)
			__fill_bkt.SetKey(nil)
			__fill_bkt++
			__fill_idx++
		}
		__fill_ht.SetNNumUsed(__fill_idx)
		__fill_ht.SetNNumOfElements(__fill_idx)
		__fill_ht.SetNNextFreeElement(__fill_idx)
		__fill_ht.SetNInternalPointer(0)
		if pad_size > 0 {
			var __fill_ht *types.HashTable = return_value.GetArr()
			var __fill_bkt *types.Bucket = __fill_ht.GetArData() + __fill_ht.GetNNumUsed()
			var __fill_idx uint32 = __fill_ht.GetNNumUsed()
			b.Assert(__fill_ht.HasUFlags(types.HASH_FLAG_PACKED))
			for i = 0; i < num_pads; i++ {
				types.ZVAL_COPY_VALUE(__fill_bkt.GetVal(), pad_value)
				__fill_bkt.SetH(__fill_idx)
				__fill_bkt.SetKey(nil)
				__fill_bkt++
				__fill_idx++
			}
			__fill_ht.SetNNumUsed(__fill_idx)
			__fill_ht.SetNNumOfElements(__fill_idx)
			__fill_ht.SetNNextFreeElement(__fill_idx)
			__fill_ht.SetNInternalPointer(0)
		}
	} else {
		if pad_size < 0 {
			for i = 0; i < num_pads; i++ {
				return_value.GetArr().NextIndexInsertNew(pad_value)
			}
		}
		var __ht *types.HashTable = input.GetArr()
		for _, _p := range __ht.foreachData() {
			var _z *types.Zval = _p.GetVal()
			if _z.IsType(types.IS_INDIRECT) {
				_z = _z.GetZv()
				if _z.IsType(types.IS_UNDEF) {
					continue
				}
			}
			key = _p.GetKey()
			value = _z
			value.TryAddRefcount()
			if key != nil {
				return_value.GetArr().KeyAddNew(key.GetStr(), value)
			} else {
				return_value.GetArr().NextIndexInsertNew(value)
			}
		}
		if pad_size > 0 {
			for i = 0; i < num_pads; i++ {
				return_value.GetArr().NextIndexInsertNew(pad_value)
			}
		}
	}
}
func ZifArrayFlip(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var array *types.Zval
	var entry *types.Zval
	var data types.Zval
	var num_idx zend.ZendUlong
	var str_idx *types.ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			array = fp.ParseArray()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	zend.ArrayInitSize(return_value, types.Z_ARRVAL_P(array).GetNNumOfElements())
	var __ht *types.HashTable = array.GetArr()
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()

		num_idx = _p.GetH()
		str_idx = _p.GetKey()
		entry = _z
		entry = types.ZVAL_DEREF(entry)
		if entry.IsType(types.IS_LONG) {
			if str_idx != nil {
				data.SetStringCopy(str_idx)
			} else {
				data.SetLong(num_idx)
			}
			return_value.GetArr().IndexUpdateH(entry.GetLval(), &data)
		} else if entry.IsType(types.IS_STRING) {
			if str_idx != nil {
				data.SetStringCopy(str_idx)
			} else {
				data.SetLong(num_idx)
			}
			return_value.GetArr().SymtableUpdate(entry.GetStr().GetStr(), &data)
		} else {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Can only flip STRING and INTEGER values!")
		}
	}
}
func ZifArrayChangeKeyCase(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var array *types.Zval
	var entry *types.Zval
	var string_key *types.ZendString
	var new_key *types.ZendString
	var num_key zend.ZendUlong
	var change_to_upper zend.ZendLong = 0
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			array = fp.ParseArray()
			fp.StartOptional()
			change_to_upper = fp.ParseLong()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	zend.ArrayInitSize(return_value, types.Z_ARRVAL_P(array).GetNNumOfElements())
	var __ht *types.HashTable = array.GetArr()
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()

		num_key = _p.GetH()
		string_key = _p.GetKey()
		entry = _z
		if string_key == nil {
			entry = return_value.GetArr().IndexUpdateH(num_key, entry)
		} else {
			if change_to_upper != 0 {
				new_key = PhpStringToupper(string_key)
			} else {
				new_key = PhpStringTolower(string_key)
			}
			entry = return_value.GetArr().KeyUpdate(new_key.GetStr(), entry)
			types.ZendStringReleaseEx(new_key, 0)
		}
		zend.ZvalAddRef(entry)
	}
}
func ArrayBucketindexSwap(p any, q any) {
	var f *Bucketindex = (*Bucketindex)(p)
	var g *Bucketindex = (*Bucketindex)(q)
	var t Bucketindex
	t = *f
	*f = *g
	*g = t
}
func ZifArrayUnique(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var array *types.Zval
	var idx uint32
	var p *types.Bucket
	var arTmp *Bucketindex
	var cmpdata *Bucketindex
	var lastkept *Bucketindex
	var i uint
	var sort_type zend.ZendLong = PHP_SORT_STRING
	var cmp types.CompareFuncT
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			array = fp.ParseArray()
			fp.StartOptional()
			sort_type = fp.ParseLong()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	if types.Z_ARRVAL_P(array).GetNNumOfElements() <= 1 {
		types.ZVAL_COPY(return_value, array)
		return
	}
	if sort_type == PHP_SORT_STRING {
		var seen types.HashTable
		var num_key zend.ZendLong
		var str_key *types.ZendString
		var val *types.Zval
		zend.ZendHashInit(&seen, types.Z_ARRVAL_P(array).GetNNumOfElements(), nil, nil, 0)
		zend.ArrayInit(return_value)
		var __ht *types.HashTable = array.GetArr()
		for _, _p := range __ht.foreachData() {
			var _z *types.Zval = _p.GetVal()
			if _z.IsType(types.IS_INDIRECT) {
				_z = _z.GetZv()
				if _z.IsType(types.IS_UNDEF) {
					continue
				}
			}
			num_key = _p.GetH()
			str_key = _p.GetKey()
			val = _z
			var retval *types.Zval
			if val.IsType(types.IS_STRING) {
				retval = zend.ZendHashAddEmptyElement(&seen, val.GetStr())
			} else {
				var tmp_str_val *types.ZendString
				var str_val *types.ZendString = zend.ZvalGetTmpString(val, &tmp_str_val)
				retval = zend.ZendHashAddEmptyElement(&seen, str_val)
				zend.ZendTmpStringRelease(tmp_str_val)
			}
			if retval != nil {

				/* First occurrence of the value */

				if val.IsReference() && val.GetRefcount() == 1 {
					val = types.ZVAL_DEREF(val)
				}
				val.TryAddRefcount()
				if str_key != nil {
					return_value.GetArr().KeyAddNew(str_key.GetStr(), val)
				} else {
					return_value.GetArr().IndexAddNewH(num_key, val)
				}
			}
		}
		seen.Destroy()
		return
	}
	cmp = PhpGetDataCompareFunc(sort_type, 0)
	return_value.SetArray(zend.ZendArrayDup(array.GetArr()))

	/* create and sort array with pointers to the target_hash buckets */

	arTmp = (*Bucketindex)(zend.Pemalloc((types.Z_ARRVAL_P(array).GetNNumOfElements()+1)*b.SizeOf("struct bucketindex"), array.GetArr().GetGcFlags()&types.IS_ARRAY_PERSISTENT))
	i = 0
	idx = 0
	for ; idx < types.Z_ARRVAL_P(array).GetNNumUsed(); idx++ {
		p = types.Z_ARRVAL_P(array).GetArData() + idx
		if p.GetVal().IsType(types.IS_UNDEF) {
			continue
		}
		if p.GetVal().IsType(types.IS_INDIRECT) && types.Z_INDIRECT(p.GetVal()).IsType(types.IS_UNDEF) {
			continue
		}
		arTmp[i].SetB(*p)
		arTmp[i].SetI(i)
		i++
	}
	arTmp[i].GetB().GetVal().SetUndef()
	zend.ZendSort(any(arTmp), i, b.SizeOf("struct bucketindex"), cmp, types.SwapFuncT(ArrayBucketindexSwap))

	/* go through the sorted array and delete duplicates from the copy */

	lastkept = arTmp
	for cmpdata = arTmp + 1; cmpdata.GetB().GetVal().GetType() != types.IS_UNDEF; cmpdata++ {
		if cmp(lastkept, cmpdata) != 0 {
			lastkept = cmpdata
		} else {
			if lastkept.GetI() > cmpdata.GetI() {
				p = lastkept.GetB()
				lastkept = cmpdata
			} else {
				p = cmpdata.GetB()
			}
			if p.GetKey() == nil {
				zend.ZendHashIndexDel(return_value.GetArr(), p.GetH())
			} else {
				if return_value.GetArr() == zend.EG__().GetSymbolTable() {
					zend.ZendDeleteGlobalVariable(p.GetKey())
				} else {
					zend.ZendHashDel(return_value.GetArr(), p.GetKey())
				}
			}
		}
	}
	zend.Pefree(arTmp, array.GetArr().GetGcFlags()&types.IS_ARRAY_PERSISTENT)
}
func ZvalCompare(first *types.Zval, second *types.Zval) int {
	return zend.StringCompareFunction(first, second)
}
func ZvalUserCompare(a *types.Zval, b *types.Zval) int {
	var args []types.Zval
	var retval types.Zval
	types.ZVAL_COPY_VALUE(&args[0], a)
	types.ZVAL_COPY_VALUE(&args[1], b)
	BG(user_compare_fci).param_count = 2
	BG(user_compare_fci).params = args
	BG(user_compare_fci).retval = &retval
	BG(user_compare_fci).no_separation = 0
	if zend.ZendCallFunction(&(BG(user_compare_fci)), &(BG(user_compare_fci_cache))) == types.SUCCESS && retval.GetType() != types.IS_UNDEF {
		var ret zend.ZendLong = zend.ZvalGetLong(&retval)
		zend.ZvalPtrDtor(&retval)
		return zend.ZEND_NORMALIZE_BOOL(ret)
	} else {
		return 0
	}
}
func PhpArrayIntersectKey(executeData *zend.ZendExecuteData, return_value *types.Zval, data_compare_type int) {
	var idx uint32
	var p *types.Bucket
	var argc int
	var i int
	var args *types.Zval
	var intersect_data_compare_func func(*types.Zval, *types.Zval) int = nil
	var ok types.ZendBool
	var val *types.Zval
	var data *types.Zval
	var req_args int
	var param_spec *byte

	/* Get the argument count */

	argc = executeData.NumArgs()
	if data_compare_type == INTERSECT_COMP_DATA_USER {

		/* INTERSECT_COMP_DATA_USER - array_uintersect_assoc() */

		req_args = 3
		param_spec = "+f"
		intersect_data_compare_func = ZvalUserCompare
	} else {

		/*     INTERSECT_COMP_DATA_NONE - array_intersect_key()
		       INTERSECT_COMP_DATA_INTERNAL - array_intersect_assoc() */

		req_args = 2
		param_spec = "+"
		if data_compare_type == INTERSECT_COMP_DATA_INTERNAL {
			intersect_data_compare_func = ZvalCompare
		}
	}
	if argc < req_args {
		core.PhpErrorDocref(nil, zend.E_WARNING, "at least %d parameters are required, %d given", req_args, argc)
		return
	}
	if zend.ZendParseParameters(executeData.NumArgs(), param_spec, &args, &argc, &(BG(user_compare_fci)), &(BG(user_compare_fci_cache))) == types.FAILURE {
		return
	}
	for i = 0; i < argc; i++ {
		if args[i].GetType() != types.IS_ARRAY {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Expected parameter %d to be an array, %s given", i+1, zend.ZendZvalTypeName(&args[i]))
			return_value.SetNull()
			return
		}
	}
	zend.ArrayInit(return_value)
	for idx = 0; idx < types.Z_ARRVAL(args[0]).GetNNumUsed(); idx++ {
		p = types.Z_ARRVAL(args[0]).GetArData() + idx
		val = p.GetVal()
		if val.IsType(types.IS_UNDEF) {
			continue
		}
		if val.IsType(types.IS_INDIRECT) {
			val = val.GetZv()
			if val.IsType(types.IS_UNDEF) {
				continue
			}
		}
		if val.IsReference() && val.GetRefcount() == 1 {
			val = types.Z_REFVAL_P(val)
		}
		if p.GetKey() == nil {
			ok = 1
			for i = 1; i < argc; i++ {
				if b.Assign(&data, args[i].GetArr().IndexFindH(p.GetH())) == nil || intersect_data_compare_func != nil && intersect_data_compare_func(val, data) != 0 {
					ok = 0
					break
				}
			}
			if ok != 0 {
				val.TryAddRefcount()
				return_value.GetArr().IndexUpdateH(p.GetH(), val)
			}
		} else {
			ok = 1
			for i = 1; i < argc; i++ {
				if b.Assign(&data, zend.ZendHashFindExInd(args[i].GetArr(), p.GetKey(), 1)) == nil || intersect_data_compare_func != nil && intersect_data_compare_func(val, data) != 0 {
					ok = 0
					break
				}
			}
			if ok != 0 {
				val.TryAddRefcount()
				return_value.GetArr().KeyUpdate(p.GetKey().GetStr(), val)
			}
		}
	}
}
func PhpArrayIntersect(executeData *zend.ZendExecuteData, return_value *types.Zval, behavior int, data_compare_type int, key_compare_type int) {
	var args *types.Zval = nil
	var hash *types.HashTable
	var arr_argc int
	var i int
	var c int = 0
	var idx uint32
	var lists **types.Bucket
	var list **types.Bucket
	var ptrs ***types.Bucket
	var p **types.Bucket
	var req_args uint32
	var param_spec *byte
	var fci1 zend.ZendFcallInfo
	var fci2 zend.ZendFcallInfo
	var fci1_cache zend.ZendFcallInfoCache = zend.EmptyFcallInfoCache
	var fci2_cache zend.ZendFcallInfoCache = zend.EmptyFcallInfoCache
	var fci_key *zend.ZendFcallInfo = nil
	var fci_data *zend.ZendFcallInfo
	var fci_key_cache *zend.ZendFcallInfoCache = nil
	var fci_data_cache *zend.ZendFcallInfoCache
	var old_user_compare_fci zend.ZendFcallInfo
	var old_user_compare_fci_cache zend.ZendFcallInfoCache
	var intersect_key_compare_func func(any, any) int
	var intersect_data_compare_func func(any, any) int
	if behavior == INTERSECT_NORMAL {
		intersect_key_compare_func = PhpArrayKeyCompareString
		if data_compare_type == INTERSECT_COMP_DATA_INTERNAL {

			/* array_intersect() */

			req_args = 2
			param_spec = "+"
			intersect_data_compare_func = PhpArrayDataCompareString
		} else if data_compare_type == INTERSECT_COMP_DATA_USER {

			/* array_uintersect() */

			req_args = 3
			param_spec = "+f"
			intersect_data_compare_func = PhpArrayUserCompare
		} else {
			core.PhpErrorDocref(nil, zend.E_WARNING, "data_compare_type is %d. This should never happen. Please report as a bug", data_compare_type)
			return
		}
		if executeData.NumArgs() < req_args {
			core.PhpErrorDocref(nil, zend.E_WARNING, "at least %d parameters are required, %d given", req_args, executeData.NumArgs())
			return
		}
		if zend.ZendParseParameters(executeData.NumArgs(), param_spec, &args, &arr_argc, &fci1, &fci1_cache) == types.FAILURE {
			return
		}
		fci_data = &fci1
		fci_data_cache = &fci1_cache
	} else if (behavior & INTERSECT_ASSOC) != 0 {

		/* INTERSECT_KEY is subset of INTERSECT_ASSOC. When having the former
		 * no comparison of the data is done (part of INTERSECT_ASSOC) */

		if data_compare_type == INTERSECT_COMP_DATA_INTERNAL && key_compare_type == INTERSECT_COMP_KEY_INTERNAL {

			/* array_intersect_assoc() or array_intersect_key() */

			req_args = 2
			param_spec = "+"
			intersect_key_compare_func = PhpArrayKeyCompareString
			intersect_data_compare_func = PhpArrayDataCompareString
		} else if data_compare_type == INTERSECT_COMP_DATA_USER && key_compare_type == INTERSECT_COMP_KEY_INTERNAL {

			/* array_uintersect_assoc() */

			req_args = 3
			param_spec = "+f"
			intersect_key_compare_func = PhpArrayKeyCompareString
			intersect_data_compare_func = PhpArrayUserCompare
			fci_data = &fci1
			fci_data_cache = &fci1_cache
		} else if data_compare_type == INTERSECT_COMP_DATA_INTERNAL && key_compare_type == INTERSECT_COMP_KEY_USER {

			/* array_intersect_uassoc() or array_intersect_ukey() */

			req_args = 3
			param_spec = "+f"
			intersect_key_compare_func = PhpArrayUserKeyCompare
			intersect_data_compare_func = PhpArrayDataCompareString
			fci_key = &fci1
			fci_key_cache = &fci1_cache
		} else if data_compare_type == INTERSECT_COMP_DATA_USER && key_compare_type == INTERSECT_COMP_KEY_USER {

			/* array_uintersect_uassoc() */

			req_args = 4
			param_spec = "+ff"
			intersect_key_compare_func = PhpArrayUserKeyCompare
			intersect_data_compare_func = PhpArrayUserCompare
			fci_data = &fci1
			fci_data_cache = &fci1_cache
			fci_key = &fci2
			fci_key_cache = &fci2_cache
		} else {
			core.PhpErrorDocref(nil, zend.E_WARNING, "data_compare_type is %d. key_compare_type is %d. This should never happen. Please report as a bug", data_compare_type, key_compare_type)
			return
		}
		if executeData.NumArgs() < req_args {
			core.PhpErrorDocref(nil, zend.E_WARNING, "at least %d parameters are required, %d given", req_args, executeData.NumArgs())
			return
		}
		if zend.ZendParseParameters(executeData.NumArgs(), param_spec, &args, &arr_argc, &fci1, &fci1_cache, &fci2, &fci2_cache) == types.FAILURE {
			return
		}
	} else {
		core.PhpErrorDocref(nil, zend.E_WARNING, "behavior is %d. This should never happen. Please report as a bug", behavior)
		return
	}
	PHP_ARRAY_CMP_FUNC_BACKUP()

	/* for each argument, create and sort list with pointers to the hash buckets */

	lists = (**types.Bucket)(zend.SafeEmalloc(arr_argc, b.SizeOf("Bucket *"), 0))
	ptrs = (**types.Bucket)(zend.SafeEmalloc(arr_argc, b.SizeOf("Bucket *"), 0))
	if behavior == INTERSECT_NORMAL && data_compare_type == INTERSECT_COMP_DATA_USER {
		BG(user_compare_fci) = *fci_data
		BG(user_compare_fci_cache) = *fci_data_cache
	} else if (behavior&INTERSECT_ASSOC) != 0 && key_compare_type == INTERSECT_COMP_KEY_USER {
		BG(user_compare_fci) = *fci_key
		BG(user_compare_fci_cache) = *fci_key_cache
	}
	for i = 0; i < arr_argc; i++ {
		if args[i].GetType() != types.IS_ARRAY {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Expected parameter %d to be an array, %s given", i+1, zend.ZendZvalTypeName(&args[i]))
			arr_argc = i
			goto out
		}
		hash = args[i].GetArr()
		list = (*types.Bucket)(zend.Pemalloc((hash.GetNNumOfElements()+1)*b.SizeOf("Bucket"), hash.GetGcFlags()&types.IS_ARRAY_PERSISTENT))
		lists[i] = list
		ptrs[i] = list
		for idx = 0; idx < hash.GetNNumUsed(); idx++ {
			p = hash.GetArData() + idx
			if p.GetVal().IsType(types.IS_UNDEF) {
				continue
			}
			b.PostInc(&(*list)) = *p
		}
		list.GetVal().SetUndef()
		if hash.GetNNumOfElements() > 1 {
			if behavior == INTERSECT_NORMAL {
				zend.ZendSort(any(lists[i]), hash.GetNNumOfElements(), b.SizeOf("Bucket"), intersect_data_compare_func, types.SwapFuncT(zend.ZendHashBucketSwap))
			} else if (behavior & INTERSECT_ASSOC) != 0 {
				zend.ZendSort(any(lists[i]), hash.GetNNumOfElements(), b.SizeOf("Bucket"), intersect_key_compare_func, types.SwapFuncT(zend.ZendHashBucketSwap))
			}
		}
	}

	/* copy the argument array */

	return_value.SetArray(zend.ZendArrayDup(args[0].GetArr()))

	/* go through the lists and look for common values */

	for ptrs[0].GetVal().GetType() != types.IS_UNDEF {
		if (behavior&INTERSECT_ASSOC) != 0 && key_compare_type == INTERSECT_COMP_KEY_USER {
			BG(user_compare_fci) = *fci_key
			BG(user_compare_fci_cache) = *fci_key_cache
		}
		for i = 1; i < arr_argc; i++ {
			if (behavior & INTERSECT_NORMAL) != 0 {
				for ptrs[i].GetVal().GetType() != types.IS_UNDEF && 0 < b.Assign(&c, intersect_data_compare_func(ptrs[0], ptrs[i])) {
					ptrs[i]++
				}
			} else if (behavior & INTERSECT_ASSOC) != 0 {
				for ptrs[i].GetVal().GetType() != types.IS_UNDEF && 0 < b.Assign(&c, intersect_key_compare_func(ptrs[0], ptrs[i])) {
					ptrs[i]++
				}
				if c == 0 && ptrs[i].GetVal().GetType() != types.IS_UNDEF && behavior == INTERSECT_ASSOC {

					/* this means that ptrs[i] is not NULL so we can compare
					 * and "c==0" is from last operation
					 * in this branch of code we enter only when INTERSECT_ASSOC
					 * since when we have INTERSECT_KEY compare of data is not wanted. */

					if data_compare_type == INTERSECT_COMP_DATA_USER {
						BG(user_compare_fci) = *fci_data
						BG(user_compare_fci_cache) = *fci_data_cache
					}
					if intersect_data_compare_func(ptrs[0], ptrs[i]) != 0 {
						c = 1
						if key_compare_type == INTERSECT_COMP_KEY_USER {
							BG(user_compare_fci) = *fci_key
							BG(user_compare_fci_cache) = *fci_key_cache
						}
					}
				}
			}
			if ptrs[i].GetVal().IsType(types.IS_UNDEF) {

				/* delete any values corresponding to remains of ptrs[0] */

				for {
					ptrs[0]++
					p = ptrs[0] - 1
					if p.GetVal().IsType(types.IS_UNDEF) {
						goto out
					}
					if p.GetKey() == nil {
						zend.ZendHashIndexDel(return_value.GetArr(), p.GetH())
					} else {
						zend.ZendHashDel(return_value.GetArr(), p.GetKey())
					}
				}

				/* delete any values corresponding to remains of ptrs[0] */

			}
			if c != 0 {
				break
			}
			ptrs[i]++
		}
		if c != 0 {

			/* Value of ptrs[0] not in all arguments, delete all entries */

			for {
				p = ptrs[0]
				if p.GetKey() == nil {
					zend.ZendHashIndexDel(return_value.GetArr(), p.GetH())
				} else {
					zend.ZendHashDel(return_value.GetArr(), p.GetKey())
				}
				if b.PreInc(&ptrs[0]).val.u1.v.type_ == types.IS_UNDEF {
					goto out
				}
				if behavior == INTERSECT_NORMAL {
					if 0 <= intersect_data_compare_func(ptrs[0], ptrs[i]) {
						break
					}
				} else if (behavior & INTERSECT_ASSOC) != 0 {

					/* no need of looping because indexes are unique */

					break

					/* no need of looping because indexes are unique */

				}
			}

			/* Value of ptrs[0] not in all arguments, delete all entries */

		} else {

			/* ptrs[0] is present in all the arguments */

			for {
				if b.PreInc(&ptrs[0]).val.u1.v.type_ == types.IS_UNDEF {
					goto out
				}
				if behavior == INTERSECT_NORMAL {
					if intersect_data_compare_func(ptrs[0]-1, ptrs[0]) != 0 {
						break
					}
				} else if (behavior & INTERSECT_ASSOC) != 0 {

					/* no need of looping because indexes are unique */

					break

					/* no need of looping because indexes are unique */

				}
			}

			/* ptrs[0] is present in all the arguments */

		}
	}
out:
	for i = 0; i < arr_argc; i++ {
		hash = args[i].GetArr()
		zend.Pefree(lists[i], hash.GetGcFlags()&types.IS_ARRAY_PERSISTENT)
	}
	PHP_ARRAY_CMP_FUNC_RESTORE()
	zend.Efree(ptrs)
	zend.Efree(lists)
}
func ZifArrayIntersectKey(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	PhpArrayIntersectKey(executeData, return_value, INTERSECT_COMP_DATA_NONE)
}
func ZifArrayIntersectUkey(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	PhpArrayIntersect(executeData, return_value, INTERSECT_KEY, INTERSECT_COMP_DATA_INTERNAL, INTERSECT_COMP_KEY_USER)
}
func ZifArrayIntersect(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	PhpArrayIntersect(executeData, return_value, INTERSECT_NORMAL, INTERSECT_COMP_DATA_INTERNAL, INTERSECT_COMP_KEY_INTERNAL)
}
func ZifArrayUintersect(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	PhpArrayIntersect(executeData, return_value, INTERSECT_NORMAL, INTERSECT_COMP_DATA_USER, INTERSECT_COMP_KEY_INTERNAL)
}
func ZifArrayIntersectAssoc(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	PhpArrayIntersectKey(executeData, return_value, INTERSECT_COMP_DATA_INTERNAL)
}
func ZifArrayIntersectUassoc(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	PhpArrayIntersect(executeData, return_value, INTERSECT_ASSOC, INTERSECT_COMP_DATA_INTERNAL, INTERSECT_COMP_KEY_USER)
}
func ZifArrayUintersectAssoc(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	PhpArrayIntersectKey(executeData, return_value, INTERSECT_COMP_DATA_USER)
}
func ZifArrayUintersectUassoc(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	PhpArrayIntersect(executeData, return_value, INTERSECT_ASSOC, INTERSECT_COMP_DATA_USER, INTERSECT_COMP_KEY_USER)
}
func PhpArrayDiffKey(executeData *zend.ZendExecuteData, return_value *types.Zval, data_compare_type int) {
	var idx uint32
	var p *types.Bucket
	var argc int
	var i int
	var args *types.Zval
	var diff_data_compare_func func(*types.Zval, *types.Zval) int = nil
	var ok types.ZendBool
	var val *types.Zval
	var data *types.Zval

	/* Get the argument count */

	argc = executeData.NumArgs()
	if data_compare_type == DIFF_COMP_DATA_USER {
		if argc < 3 {
			core.PhpErrorDocref(nil, zend.E_WARNING, "at least 3 parameters are required, %d given", executeData.NumArgs())
			return
		}
		if zend.ZendParseParameters(executeData.NumArgs(), "+f", &args, &argc, &(BG(user_compare_fci)), &(BG(user_compare_fci_cache))) == types.FAILURE {
			return
		}
		diff_data_compare_func = ZvalUserCompare
	} else {
		if argc < 2 {
			core.PhpErrorDocref(nil, zend.E_WARNING, "at least 2 parameters are required, %d given", executeData.NumArgs())
			return
		}
		if zend.ZendParseParameters(executeData.NumArgs(), "+", &args, &argc) == types.FAILURE {
			return
		}
		if data_compare_type == DIFF_COMP_DATA_INTERNAL {
			diff_data_compare_func = ZvalCompare
		}
	}
	for i = 0; i < argc; i++ {
		if args[i].GetType() != types.IS_ARRAY {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Expected parameter %d to be an array, %s given", i+1, zend.ZendZvalTypeName(&args[i]))
			return_value.SetNull()
			return
		}
	}
	zend.ArrayInit(return_value)
	for idx = 0; idx < types.Z_ARRVAL(args[0]).GetNNumUsed(); idx++ {
		p = types.Z_ARRVAL(args[0]).GetArData() + idx
		val = p.GetVal()
		if val.IsType(types.IS_UNDEF) {
			continue
		}
		if val.IsType(types.IS_INDIRECT) {
			val = val.GetZv()
			if val.IsType(types.IS_UNDEF) {
				continue
			}
		}
		if val.IsReference() && val.GetRefcount() == 1 {
			val = types.Z_REFVAL_P(val)
		}
		if p.GetKey() == nil {
			ok = 1
			for i = 1; i < argc; i++ {
				if b.Assign(&data, args[i].GetArr().IndexFindH(p.GetH())) != nil && (diff_data_compare_func == nil || diff_data_compare_func(val, data) == 0) {
					ok = 0
					break
				}
			}
			if ok != 0 {
				val.TryAddRefcount()
				return_value.GetArr().IndexUpdateH(p.GetH(), val)
			}
		} else {
			ok = 1
			for i = 1; i < argc; i++ {
				if b.Assign(&data, zend.ZendHashFindExInd(args[i].GetArr(), p.GetKey(), 1)) != nil && (diff_data_compare_func == nil || diff_data_compare_func(val, data) == 0) {
					ok = 0
					break
				}
			}
			if ok != 0 {
				val.TryAddRefcount()
				return_value.GetArr().KeyUpdate(p.GetKey().GetStr(), val)
			}
		}
	}
}
func PhpArrayDiff(executeData *zend.ZendExecuteData, return_value *types.Zval, behavior int, data_compare_type int, key_compare_type int) {
	var args *types.Zval = nil
	var hash *types.HashTable
	var arr_argc int
	var i int
	var c int
	var idx uint32
	var lists **types.Bucket
	var list **types.Bucket
	var ptrs ***types.Bucket
	var p **types.Bucket
	var req_args uint32
	var param_spec *byte
	var fci1 zend.ZendFcallInfo
	var fci2 zend.ZendFcallInfo
	var fci1_cache zend.ZendFcallInfoCache = zend.EmptyFcallInfoCache
	var fci2_cache zend.ZendFcallInfoCache = zend.EmptyFcallInfoCache
	var fci_key *zend.ZendFcallInfo = nil
	var fci_data *zend.ZendFcallInfo
	var fci_key_cache *zend.ZendFcallInfoCache = nil
	var fci_data_cache *zend.ZendFcallInfoCache
	var old_user_compare_fci zend.ZendFcallInfo
	var old_user_compare_fci_cache zend.ZendFcallInfoCache
	var diff_key_compare_func func(any, any) int
	var diff_data_compare_func func(any, any) int
	if behavior == DIFF_NORMAL {
		diff_key_compare_func = PhpArrayKeyCompareString
		if data_compare_type == DIFF_COMP_DATA_INTERNAL {

			/* array_diff */

			req_args = 2
			param_spec = "+"
			diff_data_compare_func = PhpArrayDataCompareString
		} else if data_compare_type == DIFF_COMP_DATA_USER {

			/* array_udiff */

			req_args = 3
			param_spec = "+f"
			diff_data_compare_func = PhpArrayUserCompare
		} else {
			core.PhpErrorDocref(nil, zend.E_WARNING, "data_compare_type is %d. This should never happen. Please report as a bug", data_compare_type)
			return
		}
		if executeData.NumArgs() < req_args {
			core.PhpErrorDocref(nil, zend.E_WARNING, "at least %d parameters are required, %d given", req_args, executeData.NumArgs())
			return
		}
		if zend.ZendParseParameters(executeData.NumArgs(), param_spec, &args, &arr_argc, &fci1, &fci1_cache) == types.FAILURE {
			return
		}
		fci_data = &fci1
		fci_data_cache = &fci1_cache
	} else if (behavior & DIFF_ASSOC) != 0 {

		/* DIFF_KEY is subset of DIFF_ASSOC. When having the former
		 * no comparison of the data is done (part of DIFF_ASSOC) */

		if data_compare_type == DIFF_COMP_DATA_INTERNAL && key_compare_type == DIFF_COMP_KEY_INTERNAL {

			/* array_diff_assoc() or array_diff_key() */

			req_args = 2
			param_spec = "+"
			diff_key_compare_func = PhpArrayKeyCompareString
			diff_data_compare_func = PhpArrayDataCompareString
		} else if data_compare_type == DIFF_COMP_DATA_USER && key_compare_type == DIFF_COMP_KEY_INTERNAL {

			/* array_udiff_assoc() */

			req_args = 3
			param_spec = "+f"
			diff_key_compare_func = PhpArrayKeyCompareString
			diff_data_compare_func = PhpArrayUserCompare
			fci_data = &fci1
			fci_data_cache = &fci1_cache
		} else if data_compare_type == DIFF_COMP_DATA_INTERNAL && key_compare_type == DIFF_COMP_KEY_USER {

			/* array_diff_uassoc() or array_diff_ukey() */

			req_args = 3
			param_spec = "+f"
			diff_key_compare_func = PhpArrayUserKeyCompare
			diff_data_compare_func = PhpArrayDataCompareString
			fci_key = &fci1
			fci_key_cache = &fci1_cache
		} else if data_compare_type == DIFF_COMP_DATA_USER && key_compare_type == DIFF_COMP_KEY_USER {

			/* array_udiff_uassoc() */

			req_args = 4
			param_spec = "+ff"
			diff_key_compare_func = PhpArrayUserKeyCompare
			diff_data_compare_func = PhpArrayUserCompare
			fci_data = &fci1
			fci_data_cache = &fci1_cache
			fci_key = &fci2
			fci_key_cache = &fci2_cache
		} else {
			core.PhpErrorDocref(nil, zend.E_WARNING, "data_compare_type is %d. key_compare_type is %d. This should never happen. Please report as a bug", data_compare_type, key_compare_type)
			return
		}
		if executeData.NumArgs() < req_args {
			core.PhpErrorDocref(nil, zend.E_WARNING, "at least %d parameters are required, %d given", req_args, executeData.NumArgs())
			return
		}
		if zend.ZendParseParameters(executeData.NumArgs(), param_spec, &args, &arr_argc, &fci1, &fci1_cache, &fci2, &fci2_cache) == types.FAILURE {
			return
		}
	} else {
		core.PhpErrorDocref(nil, zend.E_WARNING, "behavior is %d. This should never happen. Please report as a bug", behavior)
		return
	}
	PHP_ARRAY_CMP_FUNC_BACKUP()

	/* for each argument, create and sort list with pointers to the hash buckets */

	lists = (**types.Bucket)(zend.SafeEmalloc(arr_argc, b.SizeOf("Bucket *"), 0))
	ptrs = (**types.Bucket)(zend.SafeEmalloc(arr_argc, b.SizeOf("Bucket *"), 0))
	if behavior == DIFF_NORMAL && data_compare_type == DIFF_COMP_DATA_USER {
		BG(user_compare_fci) = *fci_data
		BG(user_compare_fci_cache) = *fci_data_cache
	} else if (behavior&DIFF_ASSOC) != 0 && key_compare_type == DIFF_COMP_KEY_USER {
		BG(user_compare_fci) = *fci_key
		BG(user_compare_fci_cache) = *fci_key_cache
	}
	for i = 0; i < arr_argc; i++ {
		if args[i].GetType() != types.IS_ARRAY {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Expected parameter %d to be an array, %s given", i+1, zend.ZendZvalTypeName(&args[i]))
			arr_argc = i
			goto out
		}
		hash = args[i].GetArr()
		list = (*types.Bucket)(zend.Pemalloc((hash.GetNNumOfElements()+1)*b.SizeOf("Bucket"), hash.GetGcFlags()&types.IS_ARRAY_PERSISTENT))
		lists[i] = list
		ptrs[i] = list
		for idx = 0; idx < hash.GetNNumUsed(); idx++ {
			p = hash.GetArData() + idx
			if p.GetVal().IsType(types.IS_UNDEF) {
				continue
			}
			b.PostInc(&(*list)) = *p
		}
		list.GetVal().SetUndef()
		if hash.GetNNumOfElements() > 1 {
			if behavior == DIFF_NORMAL {
				zend.ZendSort(any(lists[i]), hash.GetNNumOfElements(), b.SizeOf("Bucket"), diff_data_compare_func, types.SwapFuncT(zend.ZendHashBucketSwap))
			} else if (behavior & DIFF_ASSOC) != 0 {
				zend.ZendSort(any(lists[i]), hash.GetNNumOfElements(), b.SizeOf("Bucket"), diff_key_compare_func, types.SwapFuncT(zend.ZendHashBucketSwap))
			}
		}
	}

	/* copy the argument array */

	return_value.SetArray(zend.ZendArrayDup(args[0].GetArr()))

	/* go through the lists and look for values of ptr[0] that are not in the others */

	for ptrs[0].GetVal().GetType() != types.IS_UNDEF {
		if (behavior&DIFF_ASSOC) != 0 && key_compare_type == DIFF_COMP_KEY_USER {
			BG(user_compare_fci) = *fci_key
			BG(user_compare_fci_cache) = *fci_key_cache
		}
		c = 1
		for i = 1; i < arr_argc; i++ {
			var ptr *types.Bucket = ptrs[i]
			if behavior == DIFF_NORMAL {
				for ptrs[i].GetVal().GetType() != types.IS_UNDEF && 0 < b.Assign(&c, diff_data_compare_func(ptrs[0], ptrs[i])) {
					ptrs[i]++
				}
			} else if (behavior & DIFF_ASSOC) != 0 {
				for ptr.GetVal().GetType() != types.IS_UNDEF && 0 != b.Assign(&c, diff_key_compare_func(ptrs[0], ptr)) {
					ptr++
				}
			}
			if c == 0 {
				if behavior == DIFF_NORMAL {
					if ptrs[i].GetVal().GetType() != types.IS_UNDEF {
						ptrs[i]++
					}
					break
				} else if behavior == DIFF_ASSOC {

					/* In this branch is execute only when DIFF_ASSOC. If behavior == DIFF_KEY
					 * data comparison is not needed - skipped. */

					if ptr.GetVal().GetType() != types.IS_UNDEF {
						if data_compare_type == DIFF_COMP_DATA_USER {
							BG(user_compare_fci) = *fci_data
							BG(user_compare_fci_cache) = *fci_data_cache
						}
						if diff_data_compare_func(ptrs[0], ptr) != 0 {

							/* the data is not the same */

							c = -1
							if key_compare_type == DIFF_COMP_KEY_USER {
								BG(user_compare_fci) = *fci_key
								BG(user_compare_fci_cache) = *fci_key_cache
							}
						} else {
							break
						}
					}

					/* In this branch is execute only when DIFF_ASSOC. If behavior == DIFF_KEY
					 * data comparison is not needed - skipped. */

				} else if behavior == DIFF_KEY {

					/* the behavior here differs from INTERSECT_KEY in php_intersect
					 * since in the "diff" case we have to remove the entry from
					 * return_value while when doing intersection the entry must not
					 * be deleted. */

					break

					/* the behavior here differs from INTERSECT_KEY in php_intersect
					 * since in the "diff" case we have to remove the entry from
					 * return_value while when doing intersection the entry must not
					 * be deleted. */

				}
			}
		}
		if c == 0 {

			/* ptrs[0] in one of the other arguments */

			for {
				p = ptrs[0]
				if p.GetKey() == nil {
					zend.ZendHashIndexDel(return_value.GetArr(), p.GetH())
				} else {
					zend.ZendHashDel(return_value.GetArr(), p.GetKey())
				}
				if b.PreInc(&ptrs[0]).val.u1.v.type_ == types.IS_UNDEF {
					goto out
				}
				if behavior == DIFF_NORMAL {
					if diff_data_compare_func(ptrs[0]-1, ptrs[0]) != 0 {
						break
					}
				} else if (behavior & DIFF_ASSOC) != 0 {

					/* in this case no array_key_compare is needed */

					break

					/* in this case no array_key_compare is needed */

				}
			}

			/* ptrs[0] in one of the other arguments */

		} else {

			/* ptrs[0] in none of the other arguments */

			for {
				if b.PreInc(&ptrs[0]).val.u1.v.type_ == types.IS_UNDEF {
					goto out
				}
				if behavior == DIFF_NORMAL {
					if diff_data_compare_func(ptrs[0]-1, ptrs[0]) != 0 {
						break
					}
				} else if (behavior & DIFF_ASSOC) != 0 {

					/* in this case no array_key_compare is needed */

					break

					/* in this case no array_key_compare is needed */

				}
			}

			/* ptrs[0] in none of the other arguments */

		}
	}
out:
	for i = 0; i < arr_argc; i++ {
		hash = args[i].GetArr()
		zend.Pefree(lists[i], hash.GetGcFlags()&types.IS_ARRAY_PERSISTENT)
	}
	PHP_ARRAY_CMP_FUNC_RESTORE()
	zend.Efree(ptrs)
	zend.Efree(lists)
}
func ZifArrayDiffKey(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	PhpArrayDiffKey(executeData, return_value, DIFF_COMP_DATA_NONE)
}
func ZifArrayDiffUkey(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	PhpArrayDiff(executeData, return_value, DIFF_KEY, DIFF_COMP_DATA_INTERNAL, DIFF_COMP_KEY_USER)
}
func ZifArrayDiff(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var args *types.Zval
	var argc int
	var i int
	var num uint32
	var exclude types.HashTable
	var value *types.Zval
	var str *types.ZendString
	var tmp_str *types.ZendString
	var key *types.ZendString
	var idx zend.ZendLong
	var dummy types.Zval
	if executeData.NumArgs() < 2 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "at least 2 parameters are required, %d given", executeData.NumArgs())
		return
	}
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = -1
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			var _num_varargs int = _num_args - _i - 0
			if _num_varargs > 0 {
				args = _real_arg + 1
				argc = _num_varargs
				_i += _num_varargs
				_real_arg += _num_varargs
			} else {
				args = nil
				argc = 0
			}
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	if args[0].GetType() != types.IS_ARRAY {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Expected parameter 1 to be an array, %s given", zend.ZendZvalTypeName(&args[0]))
		return_value.SetNull()
		return
	}
	num = types.Z_ARRVAL(args[0]).GetNNumOfElements()
	if num == 0 {
		for i = 1; i < argc; i++ {
			if args[i].GetType() != types.IS_ARRAY {
				core.PhpErrorDocref(nil, zend.E_WARNING, "Expected parameter %d to be an array, %s given", i+1, zend.ZendZvalTypeName(&args[i]))
				return_value.SetNull()
				return
			}
		}
		zend.ZVAL_EMPTY_ARRAY(return_value)
		return
	} else if num == 1 {
		var found int = 0
		var search_str *types.ZendString
		var tmp_search_str *types.ZendString
		value = nil
		var __ht *types.HashTable = args[0].GetArr()
		for _, _p := range __ht.foreachData() {
			var _z *types.Zval = _p.GetVal()
			if _z.IsType(types.IS_INDIRECT) {
				_z = _z.GetZv()
				if _z.IsType(types.IS_UNDEF) {
					continue
				}
			}
			value = _z
			break
		}
		if value == nil {
			for i = 1; i < argc; i++ {
				if args[i].GetType() != types.IS_ARRAY {
					core.PhpErrorDocref(nil, zend.E_WARNING, "Expected parameter %d to be an array, %s given", i+1, zend.ZendZvalTypeName(&args[i]))
					return_value.SetNull()
					return
				}
			}
			zend.ZVAL_EMPTY_ARRAY(return_value)
			return
		}
		search_str = zend.ZvalGetTmpString(value, &tmp_search_str)
		for i = 1; i < argc; i++ {
			if args[i].GetType() != types.IS_ARRAY {
				core.PhpErrorDocref(nil, zend.E_WARNING, "Expected parameter %d to be an array, %s given", i+1, zend.ZendZvalTypeName(&args[i]))
				return_value.SetNull()
				return
			}
			if found == 0 {
				var __ht *types.HashTable = args[i].GetArr()
				for _, _p := range __ht.foreachData() {
					var _z *types.Zval = _p.GetVal()
					if _z.IsType(types.IS_INDIRECT) {
						_z = _z.GetZv()
						if _z.IsType(types.IS_UNDEF) {
							continue
						}
					}
					value = _z
					str = zend.ZvalGetTmpString(value, &tmp_str)
					if types.ZendStringEquals(search_str, str) != 0 {
						zend.ZendTmpStringRelease(tmp_str)
						found = 1
						break
					}
					zend.ZendTmpStringRelease(tmp_str)
				}
			}
		}
		zend.ZendTmpStringRelease(tmp_search_str)
		if found != 0 {
			zend.ZVAL_EMPTY_ARRAY(return_value)
		} else {
			types.ZVAL_COPY(return_value, &args[0])
		}
		return
	}

	/* count number of elements */

	num = 0
	for i = 1; i < argc; i++ {
		if args[i].GetType() != types.IS_ARRAY {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Expected parameter %d to be an array, %s given", i+1, zend.ZendZvalTypeName(&args[i]))
			return_value.SetNull()
			return
		}
		num += types.Z_ARRVAL(args[i]).GetNNumOfElements()
	}
	if num == 0 {
		types.ZVAL_COPY(return_value, &args[0])
		return
	}
	dummy.SetNull()

	/* create exclude map */

	zend.ZendHashInit(&exclude, num, nil, nil, 0)
	for i = 1; i < argc; i++ {
		var __ht *types.HashTable = args[i].GetArr()
		for _, _p := range __ht.foreachData() {
			var _z *types.Zval = _p.GetVal()
			if _z.IsType(types.IS_INDIRECT) {
				_z = _z.GetZv()
				if _z.IsType(types.IS_UNDEF) {
					continue
				}
			}
			value = _z
			str = zend.ZvalGetTmpString(value, &tmp_str)
			exclude.KeyAdd(str.GetStr(), &dummy)
			zend.ZendTmpStringRelease(tmp_str)
		}
	}

	/* copy all elements of first array that are not in exclude set */

	zend.ArrayInitSize(return_value, types.Z_ARRVAL(args[0]).GetNNumOfElements())
	var __ht *types.HashTable = args[0].GetArr()
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()
		if _z.IsType(types.IS_INDIRECT) {
			_z = _z.GetZv()
			if _z.IsType(types.IS_UNDEF) {
				continue
			}
		}
		idx = _p.GetH()
		key = _p.GetKey()
		value = _z
		str = zend.ZvalGetTmpString(value, &tmp_str)
		if zend.ZendHashExists(&exclude, str) == 0 {
			if key != nil {
				value = return_value.GetArr().KeyAddNew(key.GetStr(), value)
			} else {
				value = return_value.GetArr().IndexAddNewH(idx, value)
			}
			zend.ZvalAddRef(value)
		}
		zend.ZendTmpStringRelease(tmp_str)
	}
	exclude.Destroy()
}
func ZifArrayUdiff(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	PhpArrayDiff(executeData, return_value, DIFF_NORMAL, DIFF_COMP_DATA_USER, DIFF_COMP_KEY_INTERNAL)
}
func ZifArrayDiffAssoc(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	PhpArrayDiffKey(executeData, return_value, DIFF_COMP_DATA_INTERNAL)
}
func ZifArrayDiffUassoc(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	PhpArrayDiff(executeData, return_value, DIFF_ASSOC, DIFF_COMP_DATA_INTERNAL, DIFF_COMP_KEY_USER)
}
func ZifArrayUdiffAssoc(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	PhpArrayDiffKey(executeData, return_value, DIFF_COMP_DATA_USER)
}
func ZifArrayUdiffUassoc(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	PhpArrayDiff(executeData, return_value, DIFF_ASSOC, DIFF_COMP_DATA_USER, DIFF_COMP_KEY_USER)
}
func PhpMultisortCompare(a any, b any) int {
	var ab *types.Bucket = *((**types.Bucket)(a))
	var bb *types.Bucket = *((**types.Bucket)(b))
	var r int
	var result zend.ZendLong
	r = 0
	for {
		result = ARRAYG(multisort_func)[r](&ab[r], &bb[r])
		if result != 0 {
			if result > 0 {
				return 1
			} else {
				return -1
			}
		}
		r++
		if ab[r].GetVal().IsType(types.IS_UNDEF) {
			break
		}
	}
	return 0
}
func ArrayBucketPSawp(p any, q any) {
	var t *types.Bucket
	var f **types.Bucket = (**types.Bucket)(p)
	var g **types.Bucket = (**types.Bucket)(q)
	t = *f
	*f = *g
	*g = t
}
func ZifArrayMultisort(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var args *types.Zval
	var arrays **types.Zval
	var indirect **types.Bucket
	var idx uint32
	var p *types.Bucket
	var hash *types.HashTable
	var argc int
	var array_size int
	var num_arrays int = 0
	var parse_state []int
	var sort_order int = PHP_SORT_ASC
	var sort_type int = PHP_SORT_REGULAR
	var i int
	var k int
	var n int
	var func_ *types.CompareFuncT
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = -1
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			var _num_varargs int = _num_args - _i - 0
			if _num_varargs > 0 {
				args = _real_arg + 1
				argc = _num_varargs
				_i += _num_varargs
				_real_arg += _num_varargs
			} else {
				args = nil
				argc = 0
			}
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}

	/* Allocate space for storing pointers to input arrays and sort flags. */

	arrays = (**types.Zval)(zend.Ecalloc(argc, b.SizeOf("zval *")))
	for i = 0; i < MULTISORT_LAST; i++ {
		parse_state[i] = 0
	}
	ARRAYG(multisort_func) = (*types.CompareFuncT)(zend.Ecalloc(argc, b.SizeOf("compare_func_t")))
	func_ = ARRAYG(multisort_func)

	/* Here we go through the input arguments and parse them. Each one can
	 * be either an array or a sort flag which follows an array. If not
	 * specified, the sort flags defaults to PHP_SORT_ASC and PHP_SORT_REGULAR
	 * accordingly. There can't be two sort flags of the same type after an
	 * array, and the very first argument has to be an array. */

	for i = 0; i < argc; i++ {
		var arg *types.Zval = &args[i]
		arg = types.ZVAL_DEREF(arg)
		if arg.IsType(types.IS_ARRAY) {
			types.SEPARATE_ARRAY(arg)

			/* We see the next array, so we update the sort flags of
			 * the previous array and reset the sort flags. */

			if i > 0 {
				ARRAYG(multisort_func)[num_arrays-1] = PhpGetDataCompareFunc(sort_type, sort_order != PHP_SORT_ASC)
				sort_order = PHP_SORT_ASC
				sort_type = PHP_SORT_REGULAR
			}
			arrays[b.PostInc(&num_arrays)] = arg

			/* Next one may be an array or a list of sort flags. */

			for k = 0; k < MULTISORT_LAST; k++ {
				parse_state[k] = 1
			}

			/* Next one may be an array or a list of sort flags. */

		} else if arg.IsType(types.IS_LONG) {
			switch arg.GetLval() & ^PHP_SORT_FLAG_CASE {
			case PHP_SORT_ASC:
				fallthrough
			case PHP_SORT_DESC:

				/* flag allowed here */

				if parse_state[MULTISORT_ORDER] == 1 {

					/* Save the flag and make sure then next arg is not the current flag. */

					if arg.GetLval() == PHP_SORT_DESC {
						sort_order = PHP_SORT_DESC
					} else {
						sort_order = PHP_SORT_ASC
					}
					parse_state[MULTISORT_ORDER] = 0
				} else {
					core.PhpErrorDocref(nil, zend.E_WARNING, "Argument #%d is expected to be an array or sorting flag that has not already been specified", i+1)
					zend.Efree(func_)
					zend.Efree(arrays)
					return_value.SetFalse()
					return
				}
			case PHP_SORT_REGULAR:
				fallthrough
			case PHP_SORT_NUMERIC:
				fallthrough
			case PHP_SORT_STRING:
				fallthrough
			case PHP_SORT_NATURAL:
				fallthrough
			case PHP_SORT_LOCALE_STRING:

				/* flag allowed here */

				if parse_state[MULTISORT_TYPE] == 1 {

					/* Save the flag and make sure then next arg is not the current flag. */

					sort_type = int(arg.GetLval())
					parse_state[MULTISORT_TYPE] = 0
				} else {
					core.PhpErrorDocref(nil, zend.E_WARNING, "Argument #%d is expected to be an array or sorting flag that has not already been specified", i+1)
					zend.Efree(func_)
					zend.Efree(arrays)
					return_value.SetFalse()
					return
				}
			default:
				core.PhpErrorDocref(nil, zend.E_WARNING, "Argument #%d is an unknown sort flag", i+1)
				zend.Efree(func_)
				zend.Efree(arrays)
				return_value.SetFalse()
				return
			}
		} else {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Argument #%d is expected to be an array or a sort flag", i+1)
			zend.Efree(func_)
			zend.Efree(arrays)
			return_value.SetFalse()
			return
		}
	}

	/* Take care of the last array sort flags. */

	ARRAYG(multisort_func)[num_arrays-1] = PhpGetDataCompareFunc(sort_type, sort_order != PHP_SORT_ASC)

	/* Make sure the arrays are of the same size. */

	array_size = types.Z_ARRVAL_P(arrays[0]).GetNNumOfElements()
	for i = 0; i < num_arrays; i++ {
		if types.Z_ARRVAL_P(arrays[i]).GetNNumOfElements() != uint32(array_size) {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Array sizes are inconsistent")
			zend.Efree(func_)
			zend.Efree(arrays)
			return_value.SetFalse()
			return
		}
	}

	/* If all arrays are empty we don't need to do anything. */

	if array_size < 1 {
		zend.Efree(func_)
		zend.Efree(arrays)
		return_value.SetTrue()
		return
	}

	/* Create the indirection array. This array is of size MxN, where
	 * M is the number of entries in each input array and N is the number
	 * of the input arrays + 1. The last column is NULL to indicate the end
	 * of the row. */

	indirect = (**types.Bucket)(zend.SafeEmalloc(array_size, b.SizeOf("Bucket *"), 0))
	for i = 0; i < array_size; i++ {
		indirect[i] = (*types.Bucket)(zend.SafeEmalloc(num_arrays+1, b.SizeOf("Bucket"), 0))
	}
	for i = 0; i < num_arrays; i++ {
		k = 0
		for idx = 0; idx < types.Z_ARRVAL_P(arrays[i]).GetNNumUsed(); idx++ {
			p = types.Z_ARRVAL_P(arrays[i]).GetArData() + idx
			if p.GetVal().IsType(types.IS_UNDEF) {
				continue
			}
			indirect[k][i] = *p
			k++
		}
	}
	for k = 0; k < array_size; k++ {
		indirect[k][num_arrays].GetVal().SetUndef()
	}

	/* Do the actual sort magic - bada-bim, bada-boom. */

	zend.ZendSort(indirect, array_size, b.SizeOf("Bucket *"), PhpMultisortCompare, types.SwapFuncT(ArrayBucketPSawp))

	/* Restructure the arrays based on sorted indirect - this is mostly taken from zend_hash_sort() function. */

	for i = 0; i < num_arrays; i++ {
		var repack int
		hash = arrays[i].GetArr()
		hash.SetNNumUsed(array_size)
		hash.SetNInternalPointer(0)
		repack = !(hash.GetUFlags() & types.HASH_FLAG_PACKED)
		n = 0
		k = 0
		for ; k < array_size; k++ {
			hash.GetArData()[k] = indirect[k][i]
			if hash.GetArData()[k].GetKey() == nil {
				n++
				hash.GetArData()[k].SetH(n - 1)
			} else {
				repack = 0
			}
		}
		hash.SetNNextFreeElement(array_size)
		if repack != 0 {
			zend.ZendHashToPacked(hash)
		} else if !hash.HasUFlags(types.HASH_FLAG_PACKED) {
			hash.Rehash()
		}
	}

	/* Clean up. */

	for i = 0; i < array_size; i++ {
		zend.Efree(indirect[i])
	}
	zend.Efree(indirect)
	zend.Efree(func_)
	zend.Efree(arrays)
	return_value.SetTrue()
	return
}
func ZifArrayRand(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var input *types.Zval
	var num_req zend.ZendLong = 1
	var string_key *types.ZendString
	var num_key zend.ZendUlong
	var i int
	var num_avail int
	var bitset zend.ZendBitset
	var negative_bitset int = 0
	var bitset_len uint32
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			input = fp.ParseArray()
			fp.StartOptional()
			num_req = fp.ParseLong()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	num_avail = types.Z_ARRVAL_P(input).GetNNumOfElements()
	if num_avail == 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Array is empty")
		return
	}
	if num_req == 1 {
		var ht *types.HashTable = input.GetArr()
		if uint32(num_avail < ht.GetNNumUsed()-(ht.GetNNumUsed()>>1)) != 0 {

			/* If less than 1/2 of elements are used, don't sample. Instead search for a
			 * specific offset using linear scan. */

			var i zend.ZendLong = 0
			var randval zend.ZendLong = PhpMtRandRange(0, num_avail-1)
			var __ht *types.HashTable = input.GetArr()
			for _, _p := range __ht.foreachData() {
				var _z *types.Zval = _p.GetVal()

				num_key = _p.GetH()
				string_key = _p.GetKey()
				if i == randval {
					if string_key != nil {
						return_value.SetStringCopy(string_key)
						return
					} else {
						return_value.SetLong(num_key)
						return
					}
				}
				i++
			}
		}

		/* Sample random buckets until we hit one that is not empty.
		 * The worst case probability of hitting an empty element is 1-1/2. The worst case
		 * probability of hitting N empty elements in a row is (1-1/2)**N.
		 * For N=10 this becomes smaller than 0.1%. */

		for {
			var randval zend.ZendLong = PhpMtRandRange(0, ht.GetNNumUsed()-1)
			var bucket *types.Bucket = ht.GetArData()[randval]
			if !(bucket.GetVal().IsUndef()) {
				if bucket.GetKey() != nil {
					return_value.SetStringCopy(bucket.GetKey())
					return
				} else {
					return_value.SetLong(bucket.GetH())
					return
				}
			}

		}

		/* Sample random buckets until we hit one that is not empty.
		 * The worst case probability of hitting an empty element is 1-1/2. The worst case
		 * probability of hitting N empty elements in a row is (1-1/2)**N.
		 * For N=10 this becomes smaller than 0.1%. */

	}
	if num_req <= 0 || num_req > num_avail {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Second argument has to be between 1 and the number of elements in the array")
		return
	}

	/* Make the return value an array only if we need to pass back more than one result. */

	zend.ArrayInitSize(return_value, uint32(num_req))
	if num_req > num_avail>>1 {
		negative_bitset = 1
		num_req = num_avail - num_req
	}
	bitset_len = zend.ZendBitsetLen(num_avail)
	bitset = zend.ZEND_BITSET_ALLOCA(bitset_len, use_heap)
	zend.ZendBitsetClear(bitset, bitset_len)
	i = num_req
	for i != 0 {
		var randval zend.ZendLong = PhpMtRandRange(0, num_avail-1)
		if zend.ZendBitsetIn(bitset, randval) == 0 {
			zend.ZendBitsetIncl(bitset, randval)
			i--
		}
	}

	/* i = 0; */

	zend.ZendHashRealInitPacked(return_value.GetArr())
	var __fill_ht *types.HashTable = return_value.GetArr()
	var __fill_bkt *types.Bucket = __fill_ht.GetArData() + __fill_ht.GetNNumUsed()
	var __fill_idx uint32 = __fill_ht.GetNNumUsed()
	b.Assert(__fill_ht.HasUFlags(types.HASH_FLAG_PACKED))

	/* We can't use zend_hash_index_find()
	 * because the array may have string keys or gaps. */

	var __ht *types.HashTable = input.GetArr()
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()

		num_key = _p.GetH()
		string_key = _p.GetKey()
		if (zend.ZendBitsetIn(bitset, i) ^ negative_bitset) != 0 {
			if string_key != nil {
				__fill_bkt.GetVal().SetStringCopy(string_key)
			} else {
				__fill_bkt.GetVal().SetLong(num_key)
			}
			__fill_bkt.SetH(__fill_idx)
			__fill_bkt.SetKey(nil)
			__fill_bkt++
			__fill_idx++
		}
		i++
	}

	/* We can't use zend_hash_index_find()
	 * because the array may have string keys or gaps. */

	__fill_ht.SetNNumUsed(__fill_idx)
	__fill_ht.SetNNumOfElements(__fill_idx)
	__fill_ht.SetNNextFreeElement(__fill_idx)
	__fill_ht.SetNInternalPointer(0)
	zend.FreeAlloca(bitset, use_heap)
}
func ZifArraySum(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var input *types.Zval
	var entry *types.Zval
	var entry_n types.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			input = fp.ParseArray()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	return_value.SetLong(0)
	var __ht *types.HashTable = input.GetArr()
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()

		entry = _z
		if entry.IsType(types.IS_ARRAY) || entry.IsType(types.IS_OBJECT) {
			continue
		}
		types.ZVAL_COPY(&entry_n, entry)
		zend.ConvertScalarToNumber(&entry_n)
		zend.FastAddFunction(return_value, return_value, &entry_n)
	}
}
func ZifArrayProduct(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var input *types.Zval
	var entry *types.Zval
	var entry_n types.Zval
	var dval float64
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			input = fp.ParseArray()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	return_value.SetLong(1)
	if !(types.Z_ARRVAL_P(input).GetNNumOfElements()) {
		return
	}
	var __ht *types.HashTable = input.GetArr()
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()

		entry = _z
		if entry.IsType(types.IS_ARRAY) || entry.IsType(types.IS_OBJECT) {
			continue
		}
		types.ZVAL_COPY(&entry_n, entry)
		zend.ConvertScalarToNumber(&entry_n)
		if entry_n.IsType(types.IS_LONG) && return_value.IsType(types.IS_LONG) {
			dval = float64(return_value.GetLval() * float64(entry_n.GetLval()))
			if float64(zend.ZEND_LONG_MIN <= dval && dval <= float64(zend.ZEND_LONG_MAX)) {
				return_value.SetLval(return_value.GetLval() * entry_n.GetLval())
				continue
			}
		}
		zend.ConvertToDouble(return_value)
		zend.ConvertToDouble(&entry_n)
		return_value.SetDval(return_value.GetDval() * entry_n.GetDval())
	}
}
func ZifArrayReduce(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var input *types.Zval
	var args []types.Zval
	var operand *types.Zval
	var result types.Zval
	var retval types.Zval
	var fci zend.ZendFcallInfo
	var fci_cache zend.ZendFcallInfoCache = zend.EmptyFcallInfoCache
	var initial *types.Zval = nil
	var htbl *types.HashTable
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 3
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			input = fp.ParseArray()
			argparse.Z_PARAM_PROLOGUE(0, 0)
			if argparse.ZendParseArgFunc(_arg, &fci, &fci_cache, 0, &_error) == 0 {
				if _error == nil {
					_expected_type = argparse.Z_EXPECTED_FUNC
					_error_code = argparse.ZPP_ERROR_WRONG_ARG
					break
				} else {
					_error_code = argparse.ZPP_ERROR_WRONG_CALLBACK
					break
				}
			} else if _error != nil {
				zend.ZendWrongCallbackDeprecated(_i, _error)
			}
			fp.StartOptional()
			argparse.Z_PARAM_PROLOGUE(0, 0)
			argparse.ZendParseArgZvalDeref(_arg, &initial, 0)
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	if executeData.NumArgs() > 2 {
		types.ZVAL_COPY(&result, initial)
	} else {
		result.SetNull()
	}

	/* (zval **)input points to an element of argument stack
	 * the base pointer of which is subject to change.
	 * thus we need to keep the pointer to the hashtable for safety */

	htbl = input.GetArr()
	if htbl.GetNNumOfElements() == 0 {
		types.ZVAL_COPY_VALUE(return_value, &result)
		zend.ZendReleaseFcallInfoCache(&fci_cache)
		return
	}
	fci.SetRetval(&retval)
	fci.SetParamCount(2)
	fci.SetNoSeparation(0)
	var __ht *types.HashTable = htbl
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()

		operand = _z
		types.ZVAL_COPY_VALUE(&args[0], &result)
		types.ZVAL_COPY(&args[1], operand)
		fci.SetParams(args)
		if zend.ZendCallFunction(&fci, &fci_cache) == types.SUCCESS && retval.GetType() != types.IS_UNDEF {
			zend.ZvalPtrDtor(&args[1])
			zend.ZvalPtrDtor(&args[0])
			types.ZVAL_COPY_VALUE(&result, &retval)
		} else {
			zend.ZvalPtrDtor(&args[1])
			zend.ZvalPtrDtor(&args[0])
			return
		}
	}
	zend.ZendReleaseFcallInfoCache(&fci_cache)
	zend.ZVAL_ZVAL(return_value, &result, 1, 1)
}
func ZifArrayFilter(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var array *types.Zval
	var operand *types.Zval
	var key *types.Zval
	var args []types.Zval
	var retval types.Zval
	var have_callback types.ZendBool = 0
	var use_type zend.ZendLong = 0
	var string_key *types.ZendString
	var fci zend.ZendFcallInfo = zend.EmptyFcallInfo
	var fci_cache zend.ZendFcallInfoCache = zend.EmptyFcallInfoCache
	var num_key zend.ZendUlong
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 3
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			array = fp.ParseArray()
			fp.StartOptional()
			argparse.Z_PARAM_PROLOGUE(0, 0)
			if argparse.ZendParseArgFunc(_arg, &fci, &fci_cache, 0, &_error) == 0 {
				if _error == nil {
					_expected_type = argparse.Z_EXPECTED_FUNC
					_error_code = argparse.ZPP_ERROR_WRONG_ARG
					break
				} else {
					_error_code = argparse.ZPP_ERROR_WRONG_CALLBACK
					break
				}
			} else if _error != nil {
				zend.ZendWrongCallbackDeprecated(_i, _error)
			}
			use_type = fp.ParseLong()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	zend.ArrayInit(return_value)
	if types.Z_ARRVAL_P(array).GetNNumOfElements() == 0 {
		zend.ZendReleaseFcallInfoCache(&fci_cache)
		return
	}
	if executeData.NumArgs() > 1 {
		have_callback = 1
		fci.SetNoSeparation(0)
		fci.SetRetval(&retval)
		if use_type == ARRAY_FILTER_USE_BOTH {
			fci.SetParamCount(2)
			key = &args[1]
		} else {
			fci.SetParamCount(1)
			key = &args[0]
		}
	}
	var __ht *types.HashTable = array.GetArr()
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()
		if _z.IsType(types.IS_INDIRECT) {
			_z = _z.GetZv()
			if _z.IsType(types.IS_UNDEF) {
				continue
			}
		}
		num_key = _p.GetH()
		string_key = _p.GetKey()
		operand = _z
		if have_callback != 0 {
			if use_type != 0 {

				/* Set up the key */

				if string_key == nil {
					key.SetLong(num_key)
				} else {
					key.SetStringCopy(string_key)
				}

				/* Set up the key */

			}
			if use_type != ARRAY_FILTER_USE_KEY {
				types.ZVAL_COPY(&args[0], operand)
			}
			fci.SetParams(args)
			if zend.ZendCallFunction(&fci, &fci_cache) == types.SUCCESS {
				var retval_true int
				zend.ZvalPtrDtor(&args[0])
				if use_type == ARRAY_FILTER_USE_BOTH {
					zend.ZvalPtrDtor(&args[1])
				}
				retval_true = zend.ZendIsTrue(&retval)
				zend.ZvalPtrDtor(&retval)
				if retval_true == 0 {
					continue
				}
			} else {
				zend.ZvalPtrDtor(&args[0])
				if use_type == ARRAY_FILTER_USE_BOTH {
					zend.ZvalPtrDtor(&args[1])
				}
				return
			}
		} else if zend.ZendIsTrue(operand) == 0 {
			continue
		}
		if string_key != nil {
			operand = return_value.GetArr().KeyUpdate(string_key.GetStr(), operand)
		} else {
			operand = return_value.GetArr().IndexUpdateH(num_key, operand)
		}
		zend.ZvalAddRef(operand)
	}
	zend.ZendReleaseFcallInfoCache(&fci_cache)
}
func ZifArrayMap(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var arrays *types.Zval = nil
	var n_arrays int = 0
	var result types.Zval
	var fci zend.ZendFcallInfo = zend.EmptyFcallInfo
	var fci_cache zend.ZendFcallInfoCache = zend.EmptyFcallInfoCache
	var i int
	var k uint32
	var maxlen uint32 = 0
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = -1
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			argparse.Z_PARAM_PROLOGUE(0, 0)
			if argparse.ZendParseArgFunc(_arg, &fci, &fci_cache, 1, &_error) == 0 {
				if _error == nil {
					_expected_type = argparse.Z_EXPECTED_FUNC
					_error_code = argparse.ZPP_ERROR_WRONG_ARG
					break
				} else {
					_error_code = argparse.ZPP_ERROR_WRONG_CALLBACK
					break
				}
			} else if _error != nil {
				zend.ZendWrongCallbackDeprecated(_i, _error)
			}
			var _num_varargs int = _num_args - _i - 0
			if _num_varargs > 0 {
				arrays = _real_arg + 1
				n_arrays = _num_varargs
				_i += _num_varargs
				_real_arg += _num_varargs
			} else {
				arrays = nil
				n_arrays = 0
			}
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	return_value.SetNull()
	if n_arrays == 1 {
		var num_key zend.ZendUlong
		var str_key *types.ZendString
		var zv *types.Zval
		var arg types.Zval
		var ret int
		if arrays[0].GetType() != types.IS_ARRAY {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Expected parameter 2 to be an array, %s given", zend.ZendZvalTypeName(&arrays[0]))
			return
		}
		maxlen = types.Z_ARRVAL(arrays[0]).GetNNumOfElements()

		/* Short-circuit: if no callback and only one array, just return it. */

		if !(zend.ZEND_FCI_INITIALIZED(fci)) || maxlen == 0 {
			types.ZVAL_COPY(return_value, &arrays[0])
			zend.ZendReleaseFcallInfoCache(&fci_cache)
			return
		}
		zend.ArrayInitSize(return_value, maxlen)
		zend.ZendHashRealInit(return_value.GetArr(), types.Z_ARRVAL(arrays[0]).GetUFlags()&types.HASH_FLAG_PACKED)
		var __ht *types.HashTable = arrays[0].GetArr()
		for _, _p := range __ht.foreachData() {
			var _z *types.Zval = _p.GetVal()
			if _z.IsType(types.IS_INDIRECT) {
				_z = _z.GetZv()
				if _z.IsType(types.IS_UNDEF) {
					continue
				}
			}
			num_key = _p.GetH()
			str_key = _p.GetKey()
			zv = _z
			fci.SetRetval(&result)
			fci.SetParamCount(1)
			fci.SetParams(&arg)
			fci.SetNoSeparation(0)
			types.ZVAL_COPY(&arg, zv)
			ret = zend.ZendCallFunction(&fci, &fci_cache)
			zend.IZvalPtrDtor(&arg)
			if ret != types.SUCCESS || result.IsType(types.IS_UNDEF) {
				return_value.GetArr().DestroyEx()
				return_value.SetNull()
				return
			}
			if str_key != nil {
				zend._zendHashAppend(return_value.GetArr(), str_key, &result)
			} else {
				return_value.GetArr().IndexAddNewH(num_key, &result)
			}
		}
		zend.ZendReleaseFcallInfoCache(&fci_cache)
	} else {
		var array_pos *uint32 = (*types.HashPosition)(zend.Ecalloc(n_arrays, b.SizeOf("HashPosition")))
		for i = 0; i < n_arrays; i++ {
			if arrays[i].GetType() != types.IS_ARRAY {
				core.PhpErrorDocref(nil, zend.E_WARNING, "Expected parameter %d to be an array, %s given", i+2, zend.ZendZvalTypeName(&arrays[i]))
				zend.Efree(array_pos)
				return
			}
			if types.Z_ARRVAL(arrays[i]).GetNNumOfElements() > maxlen {
				maxlen = types.Z_ARRVAL(arrays[i]).GetNNumOfElements()
			}
		}
		zend.ArrayInitSize(return_value, maxlen)
		if !(zend.ZEND_FCI_INITIALIZED(fci)) {
			var zv types.Zval

			/* We iterate through all the arrays at once. */

			for k = 0; k < maxlen; k++ {

				/* If no callback, the result will be an array, consisting of current
				 * entries from all arrays. */

				zend.ArrayInitSize(&result, n_arrays)
				for i = 0; i < n_arrays; i++ {

					/* If this array still has elements, add the current one to the
					 * parameter list, otherwise use null value. */

					var pos uint32 = array_pos[i]
					for true {
						if pos >= types.Z_ARRVAL(arrays[i]).GetNNumUsed() {
							zv.SetNull()
							break
						} else if types.Z_ARRVAL(arrays[i]).GetArData()[pos].GetVal().GetType() != types.IS_UNDEF {
							types.ZVAL_COPY(&zv, types.Z_ARRVAL(arrays[i]).GetArData()[pos].GetVal())
							array_pos[i] = pos + 1
							break
						}
						pos++
					}
					result.GetArr().NextIndexInsertNew(&zv)
				}
				return_value.GetArr().NextIndexInsertNew(&result)
			}

			/* We iterate through all the arrays at once. */

		} else {
			var params *types.Zval = (*types.Zval)(zend.SafeEmalloc(n_arrays, b.SizeOf("zval"), 0))

			/* We iterate through all the arrays at once. */

			for k = 0; k < maxlen; k++ {
				for i = 0; i < n_arrays; i++ {

					/* If this array still has elements, add the current one to the
					 * parameter list, otherwise use null value. */

					var pos uint32 = array_pos[i]
					for true {
						if pos >= types.Z_ARRVAL(arrays[i]).GetNNumUsed() {
							params[i].SetNull()
							break
						} else if types.Z_ARRVAL(arrays[i]).GetArData()[pos].GetVal().GetType() != types.IS_UNDEF {
							types.ZVAL_COPY(&params[i], types.Z_ARRVAL(arrays[i]).GetArData()[pos].GetVal())
							array_pos[i] = pos + 1
							break
						}
						pos++
					}
				}
				fci.SetRetval(&result)
				fci.SetParamCount(n_arrays)
				fci.SetParams(params)
				fci.SetNoSeparation(0)
				if zend.ZendCallFunction(&fci, &fci_cache) != types.SUCCESS || result.IsType(types.IS_UNDEF) {
					zend.Efree(array_pos)
					return_value.GetArr().DestroyEx()
					for i = 0; i < n_arrays; i++ {
						zend.ZvalPtrDtor(&params[i])
					}
					zend.Efree(params)
					return_value.SetNull()
					return
				} else {
					for i = 0; i < n_arrays; i++ {
						zend.ZvalPtrDtor(&params[i])
					}
				}
				return_value.GetArr().NextIndexInsertNew(&result)
			}
			zend.Efree(params)
			zend.ZendReleaseFcallInfoCache(&fci_cache)
		}
		zend.Efree(array_pos)
	}
}
func ZifArrayKeyExists(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var key *types.Zval
	var array *types.Zval
	var ht *types.HashTable
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 2
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			argparse.Z_PARAM_PROLOGUE(0, 0)
			argparse.ZendParseArgZvalDeref(_arg, &key, 0)
			array = fp.ParseArrayOrObject()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	if array.IsType(types.IS_ARRAY) {
		ht = array.GetArr()
	} else {
		ht = zend.ZendGetPropertiesFor(array, zend.ZEND_PROP_PURPOSE_ARRAY_CAST)
		core.PhpErrorDocref(nil, zend.E_DEPRECATED, "Using array_key_exists() on objects is deprecated. "+"Use isset() or property_exists() instead")
	}
	switch key.GetType() {
	case types.IS_STRING:
		types.ZVAL_BOOL(return_value, ht.SymtableExistsInd(key.GetStr().GetStr()))
	case types.IS_LONG:
		types.ZVAL_BOOL(return_value, zend.ZendHashIndexExists(ht, key.GetLval()) != 0)
	case types.IS_NULL:
		types.ZVAL_BOOL(return_value, ht.KeyExistsInd(types.ZSTR_EMPTY_ALLOC().GetStr()))
	default:
		core.PhpErrorDocref(nil, zend.E_WARNING, "The first argument should be either a string or an integer")
		return_value.SetFalse()
	}
	if array.GetType() != types.IS_ARRAY {
		zend.ZendReleaseProperties(ht)
	}
}
func ZifArrayChunk(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var num_in int
	var size zend.ZendLong
	var current zend.ZendLong = 0
	var str_key *types.ZendString
	var num_key zend.ZendUlong
	var preserve_keys types.ZendBool = 0
	var input *types.Zval = nil
	var chunk types.Zval
	var entry *types.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 3
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			input = fp.ParseArray()
			size = fp.ParseLong()
			fp.StartOptional()
			preserve_keys = fp.ParseBool()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}

	/* Do bounds checking for size parameter. */

	if size < 1 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Size parameter expected to be greater than 0")
		return
	}
	num_in = types.Z_ARRVAL_P(input).GetNNumOfElements()
	if size > num_in {
		if num_in > 0 {
			size = num_in
		} else {
			size = 1
		}
	}
	zend.ArrayInitSize(return_value, uint32((num_in-1)/size+1))
	chunk.SetUndef()
	var __ht *types.HashTable = input.GetArr()
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()

		num_key = _p.GetH()
		str_key = _p.GetKey()
		entry = _z

		/* If new chunk, create and initialize it. */

		if chunk.IsType(types.IS_UNDEF) {
			zend.ArrayInitSize(&chunk, uint32(size))
		}

		/* Add entry to the chunk, preserving keys if necessary. */

		if preserve_keys != 0 {
			if str_key != nil {
				entry = chunk.GetArr().KeyUpdate(str_key.GetStr(), entry)
			} else {
				entry = chunk.GetArr().IndexUpdateH(num_key, entry)
			}
		} else {
			entry = chunk.GetArr().NextIndexInsert(entry)
		}
		zend.ZvalAddRef(entry)

		/* If reached the chunk size, add it to the result array, and reset the
		 * pointer. */

		if b.PreInc(&current)%size == 0 {
			zend.AddNextIndexZval(return_value, &chunk)
			chunk.SetUndef()
		}

		/* If reached the chunk size, add it to the result array, and reset the
		 * pointer. */

	}

	/* Add the final chunk if there is one. */

	if chunk.GetType() != types.IS_UNDEF {
		zend.AddNextIndexZval(return_value, &chunk)
	}

	/* Add the final chunk if there is one. */
}
func ZifArrayCombine(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var values *types.HashTable
	var keys *types.HashTable
	var pos_values uint32 = 0
	var entry_keys *types.Zval
	var entry_values *types.Zval
	var num_keys int
	var num_values int
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 2
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			keys = fp.ParseArrayHt()
			values = fp.ParseArrayHt()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	num_keys = keys.GetNNumOfElements()
	num_values = values.GetNNumOfElements()
	if num_keys != num_values {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Both parameters should have an equal number of elements")
		return_value.SetFalse()
		return
	}
	if num_keys == 0 {
		zend.ZVAL_EMPTY_ARRAY(return_value)
		return
	}
	zend.ArrayInitSize(return_value, num_keys)
	var __ht *types.HashTable = keys
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()

		entry_keys = _z
		for true {
			if pos_values >= values.GetNNumUsed() {
				break
			} else if values.GetArData()[pos_values].GetVal().GetType() != types.IS_UNDEF {
				entry_values = values.GetArData()[pos_values].GetVal()
				if entry_keys.IsType(types.IS_LONG) {
					entry_values = return_value.GetArr().IndexUpdateH(entry_keys.GetLval(), entry_values)
				} else {
					var tmp_key *types.ZendString
					var key *types.ZendString = zend.ZvalGetTmpString(entry_keys, &tmp_key)
					entry_values = return_value.GetArr().SymtableUpdate(key.GetStr(), entry_values)
					zend.ZendTmpStringRelease(tmp_key)
				}
				zend.ZvalAddRef(entry_values)
				pos_values++
				break
			}
			pos_values++
		}
	}
}
