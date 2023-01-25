// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/zend"
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
	return zend.SUCCESS
}
func ZmShutdownArray(type_ int, module_number int) int { return zend.SUCCESS }
func PhpArrayKeyCompare(a any, b any) int {
	var f *zend.Bucket = (*zend.Bucket)(a)
	var s *zend.Bucket = (*zend.Bucket)(b)
	var t zend.ZendUchar
	var l1 zend.ZendLong
	var l2 zend.ZendLong
	var d float64
	if f.key == nil {
		if s.key == nil {
			return zend.ZendLong(f.h > zend.ZendLong(b.Cond(s.h != 0, 1, -1)))
		} else {
			l1 = zend.ZendLong(f.h)
			t = zend.IsNumericString(s.key.val, s.key.len_, &l2, &d, 1)
			if t == zend.IS_LONG {

			} else if t == zend.IS_DOUBLE {
				return zend.ZEND_NORMALIZE_BOOL(float64(l1 - d))
			} else {
				l2 = 0
			}
		}
	} else {
		if s.key != nil {
			return zend.ZendiSmartStrcmp(f.key, s.key)
		} else {
			l2 = zend.ZendLong(s.h)
			t = zend.IsNumericString(f.key.val, f.key.len_, &l1, &d, 1)
			if t == zend.IS_LONG {

			} else if t == zend.IS_DOUBLE {
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
	var f *zend.Bucket = (*zend.Bucket)(a)
	var s *zend.Bucket = (*zend.Bucket)(b)
	if f.key == nil && s.key == nil {
		return zend.ZendLong(f.h > zend.ZendLong(b.Cond(s.h != 0, 1, -1)))
	} else {
		var d1 float64
		var d2 float64
		if f.key != nil {
			d1 = zend.ZendStrtod(f.key.val, nil)
		} else {
			d1 = float64(zend.ZendLong(f.h))
		}
		if s.key != nil {
			d2 = zend.ZendStrtod(s.key.val, nil)
		} else {
			d2 = float64(zend.ZendLong(s.h))
		}
		return zend.ZEND_NORMALIZE_BOOL(d1 - d2)
	}
}
func PhpArrayReverseKeyCompareNumeric(a any, b any) int { return PhpArrayKeyCompareNumeric(b, a) }
func PhpArrayKeyCompareStringCase(a any, b any) int {
	var f *zend.Bucket = (*zend.Bucket)(a)
	var s *zend.Bucket = (*zend.Bucket)(b)
	var s1 *byte
	var s2 *byte
	var l1 int
	var l2 int
	var buf1 []byte
	var buf2 []byte
	if f.key != nil {
		s1 = f.key.val
		l1 = f.key.len_
	} else {
		s1 = zend.ZendPrintLongToBuf(buf1+b.SizeOf("buf1")-1, f.h)
		l1 = buf1 + b.SizeOf("buf1") - 1 - s1
	}
	if s.key != nil {
		s2 = s.key.val
		l2 = s.key.len_
	} else {
		s2 = zend.ZendPrintLongToBuf(buf2+b.SizeOf("buf2")-1, s.h)
		l2 = buf2 + b.SizeOf("buf2") - 1 - s1
	}
	return zend.ZendBinaryStrcasecmpL(s1, l1, s2, l2)
}
func PhpArrayReverseKeyCompareStringCase(a any, b any) int {
	return PhpArrayKeyCompareStringCase(b, a)
}
func PhpArrayKeyCompareString(a any, b any) int {
	var f *zend.Bucket = (*zend.Bucket)(a)
	var s *zend.Bucket = (*zend.Bucket)(b)
	var s1 *byte
	var s2 *byte
	var l1 int
	var l2 int
	var buf1 []byte
	var buf2 []byte
	if f.key != nil {
		s1 = f.key.val
		l1 = f.key.len_
	} else {
		s1 = zend.ZendPrintLongToBuf(buf1+b.SizeOf("buf1")-1, f.h)
		l1 = buf1 + b.SizeOf("buf1") - 1 - s1
	}
	if s.key != nil {
		s2 = s.key.val
		l2 = s.key.len_
	} else {
		s2 = zend.ZendPrintLongToBuf(buf2+b.SizeOf("buf2")-1, s.h)
		l2 = buf2 + b.SizeOf("buf2") - 1 - s2
	}
	return zend.ZendBinaryStrcmp(s1, l1, s2, l2)
}
func PhpArrayReverseKeyCompareString(a any, b any) int { return PhpArrayKeyCompareString(b, a) }
func PhpArrayKeyCompareStringNaturalGeneral(a any, b any, fold_case int) int {
	var f *zend.Bucket = (*zend.Bucket)(a)
	var s *zend.Bucket = (*zend.Bucket)(b)
	var s1 *byte
	var s2 *byte
	var l1 int
	var l2 int
	var buf1 []byte
	var buf2 []byte
	if f.key != nil {
		s1 = f.key.val
		l1 = f.key.len_
	} else {
		s1 = zend.ZendPrintLongToBuf(buf1+b.SizeOf("buf1")-1, f.h)
		l1 = buf1 + b.SizeOf("buf1") - 1 - s1
	}
	if s.key != nil {
		s2 = s.key.val
		l2 = s.key.len_
	} else {
		s2 = zend.ZendPrintLongToBuf(buf2+b.SizeOf("buf2")-1, s.h)
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
	var f *zend.Bucket = (*zend.Bucket)(a)
	var s *zend.Bucket = (*zend.Bucket)(b)
	var s1 *byte
	var s2 *byte
	var buf1 []byte
	var buf2 []byte
	if f.key != nil {
		s1 = f.key.val
	} else {
		s1 = zend.ZendPrintLongToBuf(buf1+b.SizeOf("buf1")-1, f.h)
	}
	if s.key != nil {
		s2 = s.key.val
	} else {
		s2 = zend.ZendPrintLongToBuf(buf2+b.SizeOf("buf2")-1, s.h)
	}
	return strcoll(s1, s2)
}
func PhpArrayReverseKeyCompareStringLocale(a any, b any) int {
	return PhpArrayKeyCompareStringLocale(b, a)
}
func PhpArrayDataCompare(a any, b any) int {
	var f *zend.Bucket
	var s *zend.Bucket
	var result zend.Zval
	var first *zend.Zval
	var second *zend.Zval
	f = (*zend.Bucket)(a)
	s = (*zend.Bucket)(b)
	first = &f.val
	second = &s.val
	if zend.UNEXPECTED(zend.Z_TYPE_P(first) == zend.IS_INDIRECT) {
		first = zend.Z_INDIRECT_P(first)
	}
	if zend.UNEXPECTED(zend.Z_TYPE_P(second) == zend.IS_INDIRECT) {
		second = zend.Z_INDIRECT_P(second)
	}
	if zend.CompareFunction(&result, first, second) == zend.FAILURE {
		return 0
	}
	zend.ZEND_ASSERT(zend.Z_TYPE(result) == zend.IS_LONG)
	return zend.ZEND_NORMALIZE_BOOL(zend.Z_LVAL(result))
}
func PhpArrayReverseDataCompare(a any, b any) int { return PhpArrayDataCompare(a, b) * -1 }
func PhpArrayDataCompareNumeric(a any, b any) int {
	var f *zend.Bucket
	var s *zend.Bucket
	var first *zend.Zval
	var second *zend.Zval
	f = (*zend.Bucket)(a)
	s = (*zend.Bucket)(b)
	first = &f.val
	second = &s.val
	if zend.UNEXPECTED(zend.Z_TYPE_P(first) == zend.IS_INDIRECT) {
		first = zend.Z_INDIRECT_P(first)
	}
	if zend.UNEXPECTED(zend.Z_TYPE_P(second) == zend.IS_INDIRECT) {
		second = zend.Z_INDIRECT_P(second)
	}
	return zend.NumericCompareFunction(first, second)
}
func PhpArrayReverseDataCompareNumeric(a any, b any) int { return PhpArrayDataCompareNumeric(b, a) }
func PhpArrayDataCompareStringCase(a any, b any) int {
	var f *zend.Bucket
	var s *zend.Bucket
	var first *zend.Zval
	var second *zend.Zval
	f = (*zend.Bucket)(a)
	s = (*zend.Bucket)(b)
	first = &f.val
	second = &s.val
	if zend.UNEXPECTED(zend.Z_TYPE_P(first) == zend.IS_INDIRECT) {
		first = zend.Z_INDIRECT_P(first)
	}
	if zend.UNEXPECTED(zend.Z_TYPE_P(second) == zend.IS_INDIRECT) {
		second = zend.Z_INDIRECT_P(second)
	}
	return zend.StringCaseCompareFunction(first, second)
}
func PhpArrayReverseDataCompareStringCase(a any, b any) int {
	return PhpArrayDataCompareStringCase(b, a)
}
func PhpArrayDataCompareString(a any, b any) int {
	var f *zend.Bucket
	var s *zend.Bucket
	var first *zend.Zval
	var second *zend.Zval
	f = (*zend.Bucket)(a)
	s = (*zend.Bucket)(b)
	first = &f.val
	second = &s.val
	if zend.UNEXPECTED(zend.Z_TYPE_P(first) == zend.IS_INDIRECT) {
		first = zend.Z_INDIRECT_P(first)
	}
	if zend.UNEXPECTED(zend.Z_TYPE_P(second) == zend.IS_INDIRECT) {
		second = zend.Z_INDIRECT_P(second)
	}
	return zend.StringCompareFunction(first, second)
}
func PhpArrayReverseDataCompareString(a any, b any) int { return PhpArrayDataCompareString(b, a) }
func PhpArrayNaturalGeneralCompare(a any, b any, fold_case int) int {
	var f *zend.Bucket = (*zend.Bucket)(a)
	var s *zend.Bucket = (*zend.Bucket)(b)
	var tmp_str1 *zend.ZendString
	var tmp_str2 *zend.ZendString
	var str1 *zend.ZendString = zend.ZvalGetTmpString(&f.val, &tmp_str1)
	var str2 *zend.ZendString = zend.ZvalGetTmpString(&s.val, &tmp_str2)
	var result int = StrnatcmpEx(zend.ZSTR_VAL(str1), zend.ZSTR_LEN(str1), zend.ZSTR_VAL(str2), zend.ZSTR_LEN(str2), fold_case)
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
	var f *zend.Bucket
	var s *zend.Bucket
	var first *zend.Zval
	var second *zend.Zval
	f = (*zend.Bucket)(a)
	s = (*zend.Bucket)(b)
	first = &f.val
	second = &s.val
	if zend.UNEXPECTED(zend.Z_TYPE_P(first) == zend.IS_INDIRECT) {
		first = zend.Z_INDIRECT_P(first)
	}
	if zend.UNEXPECTED(zend.Z_TYPE_P(second) == zend.IS_INDIRECT) {
		second = zend.Z_INDIRECT_P(second)
	}
	return zend.StringLocaleCompareFunction(first, second)
}
func PhpArrayReverseDataCompareStringLocale(a any, b any) int {
	return PhpArrayDataCompareStringLocale(b, a)
}
func PhpGetKeyCompareFunc(sort_type zend.ZendLong, reverse int) zend.CompareFuncT {
	switch sort_type & ^PHP_SORT_FLAG_CASE {
	case PHP_SORT_NUMERIC:
		if reverse != 0 {
			return PhpArrayReverseKeyCompareNumeric
		} else {
			return PhpArrayKeyCompareNumeric
		}
		break
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
		break
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
		break
	case PHP_SORT_LOCALE_STRING:
		if reverse != 0 {
			return PhpArrayReverseKeyCompareStringLocale
		} else {
			return PhpArrayKeyCompareStringLocale
		}
		break
	case PHP_SORT_REGULAR:

	default:
		if reverse != 0 {
			return PhpArrayReverseKeyCompare
		} else {
			return PhpArrayKeyCompare
		}
		break
	}
	return nil
}
func PhpGetDataCompareFunc(sort_type zend.ZendLong, reverse int) zend.CompareFuncT {
	switch sort_type & ^PHP_SORT_FLAG_CASE {
	case PHP_SORT_NUMERIC:
		if reverse != 0 {
			return PhpArrayReverseDataCompareNumeric
		} else {
			return PhpArrayDataCompareNumeric
		}
		break
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
		break
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
		break
	case PHP_SORT_LOCALE_STRING:
		if reverse != 0 {
			return PhpArrayReverseDataCompareStringLocale
		} else {
			return PhpArrayDataCompareStringLocale
		}
		break
	case PHP_SORT_REGULAR:

	default:
		if reverse != 0 {
			return PhpArrayReverseDataCompare
		} else {
			return PhpArrayDataCompare
		}
		break
	}
	return nil
}
func ZifKrsort(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var array *zend.Zval
	var sort_type zend.ZendLong = PHP_SORT_REGULAR
	var cmp zend.CompareFuncT
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(1, 1)
			if zend.UNEXPECTED(zend.ZendParseArgArray(_arg, &array, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &sort_type, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			zend.RETVAL_FALSE
			return
		}
		break
	}
	cmp = PhpGetKeyCompareFunc(sort_type, 1)
	if zend.ZendHashSort(zend.Z_ARRVAL_P(array), cmp, 0) == zend.FAILURE {
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_TRUE
	return
}
func ZifKsort(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var array *zend.Zval
	var sort_type zend.ZendLong = PHP_SORT_REGULAR
	var cmp zend.CompareFuncT
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(1, 1)
			if zend.UNEXPECTED(zend.ZendParseArgArray(_arg, &array, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &sort_type, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			zend.RETVAL_FALSE
			return
		}
		break
	}
	cmp = PhpGetKeyCompareFunc(sort_type, 0)
	if zend.ZendHashSort(zend.Z_ARRVAL_P(array), cmp, 0) == zend.FAILURE {
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_TRUE
	return
}
func PhpCountRecursive(ht *zend.HashTable) zend.ZendLong {
	var cnt zend.ZendLong = 0
	var element *zend.Zval
	if (zend.GC_FLAGS(ht) & zend.GC_IMMUTABLE) == 0 {
		if zend.GC_IS_RECURSIVE(ht) != 0 {
			core.PhpErrorDocref(nil, zend.E_WARNING, "recursion detected")
			return 0
		}
		zend.GC_PROTECT_RECURSION(ht)
	}
	cnt = zend.ZendArrayCount(ht)
	for {
		var __ht *zend.HashTable = ht
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
				continue
			}
			element = _z
			zend.ZVAL_DEREF(element)
			if zend.Z_TYPE_P(element) == zend.IS_ARRAY {
				cnt += PhpCountRecursive(zend.Z_ARRVAL_P(element))
			}
		}
		break
	}
	if (zend.GC_FLAGS(ht) & zend.GC_IMMUTABLE) == 0 {
		zend.GC_UNPROTECT_RECURSION(ht)
	}
	return cnt
}
func ZifCount(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var array *zend.Zval
	var mode zend.ZendLong = COUNT_NORMAL
	var cnt zend.ZendLong
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &array, 0)
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &mode, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	switch zend.Z_TYPE_P(array) {
	case zend.IS_NULL:
		core.PhpErrorDocref(nil, zend.E_WARNING, "Parameter must be an array or an object that implements Countable")
		zend.RETVAL_LONG(0)
		return
		break
	case zend.IS_ARRAY:
		if mode != COUNT_RECURSIVE {
			cnt = zend.ZendArrayCount(zend.Z_ARRVAL_P(array))
		} else {
			cnt = PhpCountRecursive(zend.Z_ARRVAL_P(array))
		}
		zend.RETVAL_LONG(cnt)
		return
		break
	case zend.IS_OBJECT:
		var retval zend.Zval

		/* first, we check if the handler is defined */

		if zend.Z_OBJ_HT_P(array).count_elements != nil {
			zend.RETVAL_LONG(1)
			if zend.SUCCESS == zend.Z_OBJ_HT(*array).count_elements(array, &zend.Z_LVAL_P(return_value)) {
				return
			}
			if zend.ExecutorGlobals.exception != nil {
				return
			}
		}

		/* if not and the object implements Countable we call its count() method */

		if zend.InstanceofFunction(zend.Z_OBJCE_P(array), zend.ZendCeCountable) != 0 {
			zend.ZendCallMethodWith0Params(array, nil, nil, "count", &retval)
			if zend.Z_TYPE(retval) != zend.IS_UNDEF {
				zend.RETVAL_LONG(zend.ZvalGetLong(&retval))
				zend.ZvalPtrDtor(&retval)
			}
			return
		}

		/* If There's no handler and it doesn't implement Countable then add a warning */

		core.PhpErrorDocref(nil, zend.E_WARNING, "Parameter must be an array or an object that implements Countable")
		zend.RETVAL_LONG(1)
		return
		break
	default:
		core.PhpErrorDocref(nil, zend.E_WARNING, "Parameter must be an array or an object that implements Countable")
		zend.RETVAL_LONG(1)
		return
		break
	}
}
func PhpNatsort(execute_data *zend.ZendExecuteData, return_value *zend.Zval, fold_case int) {
	var array *zend.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(1, 1)
			if zend.UNEXPECTED(zend.ZendParseArgArray(_arg, &array, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	if fold_case != 0 {
		if zend.ZendHashSort(zend.Z_ARRVAL_P(array), PhpArrayNaturalCaseCompare, 0) == zend.FAILURE {
			return
		}
	} else {
		if zend.ZendHashSort(zend.Z_ARRVAL_P(array), PhpArrayNaturalCompare, 0) == zend.FAILURE {
			return
		}
	}
	zend.RETVAL_TRUE
	return
}
func ZifNatsort(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpNatsort(execute_data, return_value, 0)
}
func ZifNatcasesort(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpNatsort(execute_data, return_value, 1)
}
func ZifAsort(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var array *zend.Zval
	var sort_type zend.ZendLong = PHP_SORT_REGULAR
	var cmp zend.CompareFuncT
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(1, 1)
			if zend.UNEXPECTED(zend.ZendParseArgArray(_arg, &array, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &sort_type, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			zend.RETVAL_FALSE
			return
		}
		break
	}
	cmp = PhpGetDataCompareFunc(sort_type, 0)
	if zend.ZendHashSort(zend.Z_ARRVAL_P(array), cmp, 0) == zend.FAILURE {
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_TRUE
	return
}
func ZifArsort(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var array *zend.Zval
	var sort_type zend.ZendLong = PHP_SORT_REGULAR
	var cmp zend.CompareFuncT
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(1, 1)
			if zend.UNEXPECTED(zend.ZendParseArgArray(_arg, &array, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &sort_type, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			zend.RETVAL_FALSE
			return
		}
		break
	}
	cmp = PhpGetDataCompareFunc(sort_type, 1)
	if zend.ZendHashSort(zend.Z_ARRVAL_P(array), cmp, 0) == zend.FAILURE {
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_TRUE
	return
}
func ZifSort(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var array *zend.Zval
	var sort_type zend.ZendLong = PHP_SORT_REGULAR
	var cmp zend.CompareFuncT
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(1, 1)
			if zend.UNEXPECTED(zend.ZendParseArgArray(_arg, &array, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &sort_type, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			zend.RETVAL_FALSE
			return
		}
		break
	}
	cmp = PhpGetDataCompareFunc(sort_type, 0)
	if zend.ZendHashSort(zend.Z_ARRVAL_P(array), cmp, 1) == zend.FAILURE {
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_TRUE
	return
}
func ZifRsort(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var array *zend.Zval
	var sort_type zend.ZendLong = PHP_SORT_REGULAR
	var cmp zend.CompareFuncT
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(1, 1)
			if zend.UNEXPECTED(zend.ZendParseArgArray(_arg, &array, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &sort_type, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			zend.RETVAL_FALSE
			return
		}
		break
	}
	cmp = PhpGetDataCompareFunc(sort_type, 1)
	if zend.ZendHashSort(zend.Z_ARRVAL_P(array), cmp, 1) == zend.FAILURE {
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_TRUE
	return
}
func PhpArrayUserCompare(a any, b any) int {
	var f *zend.Bucket
	var s *zend.Bucket
	var args []zend.Zval
	var retval zend.Zval
	f = (*zend.Bucket)(a)
	s = (*zend.Bucket)(b)
	zend.ZVAL_COPY(&args[0], &f.val)
	zend.ZVAL_COPY(&args[1], &s.val)
	BG(user_compare_fci).param_count = 2
	BG(user_compare_fci).params = args
	BG(user_compare_fci).retval = &retval
	BG(user_compare_fci).no_separation = 0
	if zend.ZendCallFunction(&BG(user_compare_fci), &BG(user_compare_fci_cache)) == zend.SUCCESS && zend.Z_TYPE(retval) != zend.IS_UNDEF {
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
func PHP_ARRAY_CMP_FUNC_CHECK(func_name **zend.Zval) {
	if zend.ZendIsCallable(*func_name, 0, nil) == 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Invalid comparison function")
		BG(user_compare_fci) = old_user_compare_fci
		BG(user_compare_fci_cache) = old_user_compare_fci_cache
		zend.RETVAL_FALSE
		return
	}
}
func PHP_ARRAY_CMP_FUNC_BACKUP() {
	old_user_compare_fci = BG(user_compare_fci)
	old_user_compare_fci_cache = BG(user_compare_fci_cache)
	BG(user_compare_fci_cache) = zend.EmptyFcallInfoCache
}
func PHP_ARRAY_CMP_FUNC_RESTORE() {
	zend.ZendReleaseFcallInfoCache(&BG(user_compare_fci_cache))
	BG(user_compare_fci) = old_user_compare_fci
	BG(user_compare_fci_cache) = old_user_compare_fci_cache
}
func PhpUsort(execute_data *zend.ZendExecuteData, return_value *zend.Zval, compare_func zend.CompareFuncT, renumber zend.ZendBool) {
	var array *zend.Zval
	var arr *zend.ZendArray
	var retval zend.ZendBool
	var old_user_compare_fci zend.ZendFcallInfo
	var old_user_compare_fci_cache zend.ZendFcallInfoCache
	PHP_ARRAY_CMP_FUNC_BACKUP()
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 2
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(1, 0)
			if zend.UNEXPECTED(zend.ZendParseArgArray(_arg, &array, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgFunc(_arg, &BG(user_compare_fci), &BG(user_compare_fci_cache), 0, &_error) == 0) {
				if _error == nil {
					_expected_type = zend.Z_EXPECTED_FUNC
					_error_code = zend.ZPP_ERROR_WRONG_ARG
					break
				} else {
					_error_code = zend.ZPP_ERROR_WRONG_CALLBACK
					break
				}
			} else if zend.UNEXPECTED(_error != nil) {
				zend.ZendWrongCallbackDeprecated(_i, _error)
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			PHP_ARRAY_CMP_FUNC_RESTORE()
			return
		}
		break
	}
	arr = zend.Z_ARR_P(array)
	if zend.ZendHashNumElements(arr) == 0 {
		PHP_ARRAY_CMP_FUNC_RESTORE()
		zend.RETVAL_TRUE
		return
	}

	/* Copy array, so the in-place modifications will not be visible to the callback function */

	arr = zend.ZendArrayDup(arr)
	retval = zend.ZendHashSort(arr, compare_func, renumber) != zend.FAILURE
	var garbage zend.Zval
	zend.ZVAL_COPY_VALUE(&garbage, array)
	zend.ZVAL_ARR(array, arr)
	zend.ZvalPtrDtor(&garbage)
	PHP_ARRAY_CMP_FUNC_RESTORE()
	zend.RETVAL_BOOL(retval != 0)
	return
}
func ZifUsort(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpUsort(execute_data, return_value, PhpArrayUserCompare, 1)
}
func ZifUasort(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpUsort(execute_data, return_value, PhpArrayUserCompare, 0)
}
func PhpArrayUserKeyCompare(a any, b any) int {
	var f *zend.Bucket
	var s *zend.Bucket
	var args []zend.Zval
	var retval zend.Zval
	var result zend.ZendLong
	f = (*zend.Bucket)(a)
	s = (*zend.Bucket)(b)
	if f.key == nil {
		zend.ZVAL_LONG(&args[0], f.h)
	} else {
		zend.ZVAL_STR_COPY(&args[0], f.key)
	}
	if s.key == nil {
		zend.ZVAL_LONG(&args[1], s.h)
	} else {
		zend.ZVAL_STR_COPY(&args[1], s.key)
	}
	BG(user_compare_fci).param_count = 2
	BG(user_compare_fci).params = args
	BG(user_compare_fci).retval = &retval
	BG(user_compare_fci).no_separation = 0
	if zend.ZendCallFunction(&BG(user_compare_fci), &BG(user_compare_fci_cache)) == zend.SUCCESS && zend.Z_TYPE(retval) != zend.IS_UNDEF {
		result = zend.ZvalGetLong(&retval)
		zend.ZvalPtrDtor(&retval)
	} else {
		result = 0
	}
	zend.ZvalPtrDtor(&args[0])
	zend.ZvalPtrDtor(&args[1])
	return zend.ZEND_NORMALIZE_BOOL(result)
}
func ZifUksort(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpUsort(execute_data, return_value, PhpArrayUserKeyCompare, 0)
}
func ZifEnd(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var array *zend.HashTable
	var entry *zend.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(1, 1)
			if zend.UNEXPECTED(zend.ZendParseArgArrayHt(_arg, &array, 0, 1, 1) == 0) {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	zend.ZendHashInternalPointerEnd(array)
	if zend.USED_RET() {
		if b.Assign(&entry, zend.ZendHashGetCurrentData(array)) == nil {
			zend.RETVAL_FALSE
			return
		}
		if zend.Z_TYPE_P(entry) == zend.IS_INDIRECT {
			entry = zend.Z_INDIRECT_P(entry)
		}
		zend.ZVAL_COPY_DEREF(return_value, entry)
	}
}
func ZifPrev(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var array *zend.HashTable
	var entry *zend.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(1, 1)
			if zend.UNEXPECTED(zend.ZendParseArgArrayHt(_arg, &array, 0, 1, 1) == 0) {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	zend.ZendHashMoveBackwards(array)
	if zend.USED_RET() {
		if b.Assign(&entry, zend.ZendHashGetCurrentData(array)) == nil {
			zend.RETVAL_FALSE
			return
		}
		if zend.Z_TYPE_P(entry) == zend.IS_INDIRECT {
			entry = zend.Z_INDIRECT_P(entry)
		}
		zend.ZVAL_COPY_DEREF(return_value, entry)
	}
}
func ZifNext(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var array *zend.HashTable
	var entry *zend.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(1, 1)
			if zend.UNEXPECTED(zend.ZendParseArgArrayHt(_arg, &array, 0, 1, 1) == 0) {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	zend.ZendHashMoveForward(array)
	if zend.USED_RET() {
		if b.Assign(&entry, zend.ZendHashGetCurrentData(array)) == nil {
			zend.RETVAL_FALSE
			return
		}
		if zend.Z_TYPE_P(entry) == zend.IS_INDIRECT {
			entry = zend.Z_INDIRECT_P(entry)
		}
		zend.ZVAL_COPY_DEREF(return_value, entry)
	}
}
func ZifReset(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var array *zend.HashTable
	var entry *zend.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(1, 1)
			if zend.UNEXPECTED(zend.ZendParseArgArrayHt(_arg, &array, 0, 1, 1) == 0) {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	zend.ZendHashInternalPointerReset(array)
	if zend.USED_RET() {
		if b.Assign(&entry, zend.ZendHashGetCurrentData(array)) == nil {
			zend.RETVAL_FALSE
			return
		}
		if zend.Z_TYPE_P(entry) == zend.IS_INDIRECT {
			entry = zend.Z_INDIRECT_P(entry)
		}
		zend.ZVAL_COPY_DEREF(return_value, entry)
	}
}
func ZifCurrent(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var array *zend.HashTable
	var entry *zend.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgArrayHt(_arg, &array, 0, 1, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	if b.Assign(&entry, zend.ZendHashGetCurrentData(array)) == nil {
		zend.RETVAL_FALSE
		return
	}
	if zend.Z_TYPE_P(entry) == zend.IS_INDIRECT {
		entry = zend.Z_INDIRECT_P(entry)
	}
	zend.ZVAL_COPY_DEREF(return_value, entry)
}
func ZifKey(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var array *zend.HashTable
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgArrayHt(_arg, &array, 0, 1, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	zend.ZendHashGetCurrentKeyZval(array, return_value)
}
func ZifMin(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var argc int
	var args *zend.Zval = nil
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = -1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			var _num_varargs int = _num_args - _i - 0
			if zend.EXPECTED(_num_varargs > 0) {
				args = _real_arg + 1
				argc = _num_varargs
				_i += _num_varargs
				_real_arg += _num_varargs
			} else {
				args = nil
				argc = 0
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}

	/* mixed min ( array $values ) */

	if argc == 1 {
		var result *zend.Zval
		if zend.Z_TYPE(args[0]) != zend.IS_ARRAY {
			core.PhpErrorDocref(nil, zend.E_WARNING, "When only one parameter is given, it must be an array")
			zend.RETVAL_NULL()
		} else {
			if b.Assign(&result, zend.ZendHashMinmax(zend.Z_ARRVAL(args[0]), PhpArrayDataCompare, 0)) != nil {
				zend.ZVAL_COPY_DEREF(return_value, result)
			} else {
				core.PhpErrorDocref(nil, zend.E_WARNING, "Array must contain at least one element")
				zend.RETVAL_FALSE
			}
		}
	} else {

		/* mixed min ( mixed $value1 , mixed $value2 [, mixed $value3... ] ) */

		var min *zend.Zval
		var result zend.Zval
		var i int
		min = &args[0]
		for i = 1; i < argc; i++ {
			zend.IsSmallerFunction(&result, &args[i], min)
			if zend.Z_TYPE(result) == zend.IS_TRUE {
				min = &args[i]
			}
		}
		zend.ZVAL_COPY(return_value, min)
	}

	/* mixed min ( array $values ) */
}
func ZifMax(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var args *zend.Zval = nil
	var argc int
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = -1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			var _num_varargs int = _num_args - _i - 0
			if zend.EXPECTED(_num_varargs > 0) {
				args = _real_arg + 1
				argc = _num_varargs
				_i += _num_varargs
				_real_arg += _num_varargs
			} else {
				args = nil
				argc = 0
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}

	/* mixed max ( array $values ) */

	if argc == 1 {
		var result *zend.Zval
		if zend.Z_TYPE(args[0]) != zend.IS_ARRAY {
			core.PhpErrorDocref(nil, zend.E_WARNING, "When only one parameter is given, it must be an array")
			zend.RETVAL_NULL()
		} else {
			if b.Assign(&result, zend.ZendHashMinmax(zend.Z_ARRVAL(args[0]), PhpArrayDataCompare, 1)) != nil {
				zend.ZVAL_COPY_DEREF(return_value, result)
			} else {
				core.PhpErrorDocref(nil, zend.E_WARNING, "Array must contain at least one element")
				zend.RETVAL_FALSE
			}
		}
	} else {

		/* mixed max ( mixed $value1 , mixed $value2 [, mixed $value3... ] ) */

		var max *zend.Zval
		var result zend.Zval
		var i int
		max = &args[0]
		for i = 1; i < argc; i++ {
			zend.IsSmallerOrEqualFunction(&result, &args[i], max)
			if zend.Z_TYPE(result) == zend.IS_FALSE {
				max = &args[i]
			}
		}
		zend.ZVAL_COPY(return_value, max)
	}

	/* mixed max ( array $values ) */
}
func PhpArrayWalk(array *zend.Zval, userdata *zend.Zval, recursive int) int {
	var args []zend.Zval
	var retval zend.Zval
	var zv *zend.Zval
	var target_hash *zend.HashTable = zend.HASH_OF(array)
	var pos zend.HashPosition
	var ht_iter uint32
	var result int = zend.SUCCESS

	/* Set up known arguments */

	zend.ZVAL_UNDEF(&args[1])
	if userdata != nil {
		zend.ZVAL_COPY(&args[2], userdata)
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

		if zend.Z_TYPE_P(zv) == zend.IS_INDIRECT {
			zv = zend.Z_INDIRECT_P(zv)
			if zend.Z_TYPE_P(zv) == zend.IS_UNDEF {
				zend.ZendHashMoveForwardEx(target_hash, &pos)
				continue
			}

			/* Add type source for property references. */

			if zend.Z_TYPE_P(zv) != zend.IS_REFERENCE && zend.Z_TYPE_P(array) == zend.IS_OBJECT {
				var prop_info *zend.ZendPropertyInfo = zend.ZendGetTypedPropertyInfoForSlot(zend.Z_OBJ_P(array), zv)
				if prop_info != nil {
					zend.ZVAL_NEW_REF(zv, zv)
					zend.ZEND_REF_ADD_TYPE_SOURCE(zend.Z_REF_P(zv), prop_info)
				}
			}

			/* Add type source for property references. */

		}

		/* Ensure the value is a reference. Otherwise the location of the value may be freed. */

		zend.ZVAL_MAKE_REF(zv)

		/* Retrieve key */

		zend.ZendHashGetCurrentKeyZvalEx(target_hash, &args[1], &pos)

		/* Move to next element already now -- this mirrors the approach used by foreach
		 * and ensures proper behavior with regard to modifications. */

		zend.ZendHashMoveForwardEx(target_hash, &pos)

		/* Back up hash position, as it may change */

		zend.ExecutorGlobals.ht_iterators[ht_iter].pos = pos
		if recursive != 0 && zend.Z_TYPE_P(zend.Z_REFVAL_P(zv)) == zend.IS_ARRAY {
			var thash *zend.HashTable
			var orig_array_walk_fci zend.ZendFcallInfo
			var orig_array_walk_fci_cache zend.ZendFcallInfoCache
			var ref zend.Zval
			zend.ZVAL_COPY_VALUE(&ref, zv)
			zend.ZVAL_DEREF(zv)
			zend.SEPARATE_ARRAY(zv)
			thash = zend.Z_ARRVAL_P(zv)
			if zend.GC_IS_RECURSIVE(thash) != 0 {
				core.PhpErrorDocref(nil, zend.E_WARNING, "recursion detected")
				result = zend.FAILURE
				break
			}

			/* backup the fcall info and cache */

			orig_array_walk_fci = BG(array_walk_fci)
			orig_array_walk_fci_cache = BG(array_walk_fci_cache)
			zend.Z_ADDREF(ref)
			zend.GC_PROTECT_RECURSION(thash)
			result = PhpArrayWalk(zv, userdata, recursive)
			if zend.Z_TYPE_P(zend.Z_REFVAL(ref)) == zend.IS_ARRAY && thash == zend.Z_ARRVAL_P(zend.Z_REFVAL(ref)) {

				/* If the hashtable changed in the meantime, we'll "leak" this apply count
				 * increment -- our reference to thash is no longer valid. */

				zend.GC_UNPROTECT_RECURSION(thash)

				/* If the hashtable changed in the meantime, we'll "leak" this apply count
				 * increment -- our reference to thash is no longer valid. */

			}
			zend.ZvalPtrDtor(&ref)

			/* restore the fcall info and cache */

			BG(array_walk_fci) = orig_array_walk_fci
			BG(array_walk_fci_cache) = orig_array_walk_fci_cache
		} else {
			zend.ZVAL_COPY(&args[0], zv)

			/* Call the userland function */

			result = zend.ZendCallFunction(&BG(array_walk_fci), &BG(array_walk_fci_cache))
			if result == zend.SUCCESS {
				zend.ZvalPtrDtor(&retval)
			}
			zend.ZvalPtrDtor(&args[0])
		}
		if zend.Z_TYPE(args[1]) != zend.IS_UNDEF {
			zend.ZvalPtrDtor(&args[1])
			zend.ZVAL_UNDEF(&args[1])
		}
		if result == zend.FAILURE {
			break
		}

		/* Reload array and position -- both may have changed */

		if zend.Z_TYPE_P(array) == zend.IS_ARRAY {
			pos = zend.ZendHashIteratorPosEx(ht_iter, array)
			target_hash = zend.Z_ARRVAL_P(array)
		} else if zend.Z_TYPE_P(array) == zend.IS_OBJECT {
			target_hash = zend.Z_OBJPROP_P(array)
			pos = zend.ZendHashIteratorPos(ht_iter, target_hash)
		} else {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Iterated value is no longer an array or object")
			result = zend.FAILURE
			break
		}

		/* Reload array and position -- both may have changed */

		if zend.ExecutorGlobals.exception != nil {
			break
		}
	}
	if userdata != nil {
		zend.ZvalPtrDtor(&args[2])
	}
	zend.ZendHashIteratorDel(ht_iter)
	return result
}
func ZifArrayWalk(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var array *zend.Zval
	var userdata *zend.Zval = nil
	var orig_array_walk_fci zend.ZendFcallInfo
	var orig_array_walk_fci_cache zend.ZendFcallInfoCache
	orig_array_walk_fci = BG(array_walk_fci)
	orig_array_walk_fci_cache = BG(array_walk_fci_cache)
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 3
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(1, 1)
			if zend.UNEXPECTED(zend.ZendParseArgArray(_arg, &array, 0, 1) == 0) {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgFunc(_arg, &BG(array_walk_fci), &BG(array_walk_fci_cache), 0, &_error) == 0) {
				if _error == nil {
					_expected_type = zend.Z_EXPECTED_FUNC
					_error_code = zend.ZPP_ERROR_WRONG_ARG
					break
				} else {
					_error_code = zend.ZPP_ERROR_WRONG_CALLBACK
					break
				}
			} else if zend.UNEXPECTED(_error != nil) {
				zend.ZendWrongCallbackDeprecated(_i, _error)
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &userdata, 0)
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			BG(array_walk_fci) = orig_array_walk_fci
			BG(array_walk_fci_cache) = orig_array_walk_fci_cache
			return
		}
		break
	}
	PhpArrayWalk(array, userdata, 0)
	zend.ZendReleaseFcallInfoCache(&BG(array_walk_fci_cache))
	BG(array_walk_fci) = orig_array_walk_fci
	BG(array_walk_fci_cache) = orig_array_walk_fci_cache
	zend.RETVAL_TRUE
	return
}
func ZifArrayWalkRecursive(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var array *zend.Zval
	var userdata *zend.Zval = nil
	var orig_array_walk_fci zend.ZendFcallInfo
	var orig_array_walk_fci_cache zend.ZendFcallInfoCache
	orig_array_walk_fci = BG(array_walk_fci)
	orig_array_walk_fci_cache = BG(array_walk_fci_cache)
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 3
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(1, 1)
			if zend.UNEXPECTED(zend.ZendParseArgArray(_arg, &array, 0, 1) == 0) {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgFunc(_arg, &BG(array_walk_fci), &BG(array_walk_fci_cache), 0, &_error) == 0) {
				if _error == nil {
					_expected_type = zend.Z_EXPECTED_FUNC
					_error_code = zend.ZPP_ERROR_WRONG_ARG
					break
				} else {
					_error_code = zend.ZPP_ERROR_WRONG_CALLBACK
					break
				}
			} else if zend.UNEXPECTED(_error != nil) {
				zend.ZendWrongCallbackDeprecated(_i, _error)
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &userdata, 0)
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			BG(array_walk_fci) = orig_array_walk_fci
			BG(array_walk_fci_cache) = orig_array_walk_fci_cache
			return
		}
		break
	}
	PhpArrayWalk(array, userdata, 1)
	zend.ZendReleaseFcallInfoCache(&BG(array_walk_fci_cache))
	BG(array_walk_fci) = orig_array_walk_fci
	BG(array_walk_fci_cache) = orig_array_walk_fci_cache
	zend.RETVAL_TRUE
	return
}
func PhpSearchArray(execute_data *zend.ZendExecuteData, return_value *zend.Zval, behavior int) {
	var value *zend.Zval
	var array *zend.Zval
	var entry *zend.Zval
	var num_idx zend.ZendUlong
	var str_idx *zend.ZendString
	var strict zend.ZendBool = 0
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 3
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &value, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgArray(_arg, &array, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgBool(_arg, &strict, &_dummy, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_BOOL
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	if strict != 0 {
		if zend.Z_TYPE_P(value) == zend.IS_LONG {
			for {
				var __ht *zend.HashTable = zend.Z_ARRVAL_P(array)
				var _p *zend.Bucket = __ht.arData
				var _end *zend.Bucket = _p + __ht.nNumUsed
				for ; _p != _end; _p++ {
					var _z *zend.Zval = &_p.val
					if zend.Z_TYPE_P(_z) == zend.IS_INDIRECT {
						_z = zend.Z_INDIRECT_P(_z)
					}
					if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
						continue
					}
					num_idx = _p.h
					str_idx = _p.key
					entry = _z
					zend.ZVAL_DEREF(entry)
					if zend.Z_TYPE_P(entry) == zend.IS_LONG && zend.Z_LVAL_P(entry) == zend.Z_LVAL_P(value) {
						if behavior == 0 {
							zend.RETVAL_TRUE
							return
						} else {
							if str_idx != nil {
								zend.RETVAL_STR_COPY(str_idx)
							} else {
								zend.RETVAL_LONG(num_idx)
							}
							return
						}
					}
				}
				break
			}
		} else {
			for {
				var __ht *zend.HashTable = zend.Z_ARRVAL_P(array)
				var _p *zend.Bucket = __ht.arData
				var _end *zend.Bucket = _p + __ht.nNumUsed
				for ; _p != _end; _p++ {
					var _z *zend.Zval = &_p.val
					if zend.Z_TYPE_P(_z) == zend.IS_INDIRECT {
						_z = zend.Z_INDIRECT_P(_z)
					}
					if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
						continue
					}
					num_idx = _p.h
					str_idx = _p.key
					entry = _z
					zend.ZVAL_DEREF(entry)
					if zend.FastIsIdenticalFunction(value, entry) != 0 {
						if behavior == 0 {
							zend.RETVAL_TRUE
							return
						} else {
							if str_idx != nil {
								zend.RETVAL_STR_COPY(str_idx)
							} else {
								zend.RETVAL_LONG(num_idx)
							}
							return
						}
					}
				}
				break
			}
		}
	} else {
		if zend.Z_TYPE_P(value) == zend.IS_LONG {
			for {
				var __ht *zend.HashTable = zend.Z_ARRVAL_P(array)
				var _p *zend.Bucket = __ht.arData
				var _end *zend.Bucket = _p + __ht.nNumUsed
				for ; _p != _end; _p++ {
					var _z *zend.Zval = &_p.val
					if zend.Z_TYPE_P(_z) == zend.IS_INDIRECT {
						_z = zend.Z_INDIRECT_P(_z)
					}
					if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
						continue
					}
					num_idx = _p.h
					str_idx = _p.key
					entry = _z
					if zend.FastEqualCheckLong(value, entry) != 0 {
						if behavior == 0 {
							zend.RETVAL_TRUE
							return
						} else {
							if str_idx != nil {
								zend.RETVAL_STR_COPY(str_idx)
							} else {
								zend.RETVAL_LONG(num_idx)
							}
							return
						}
					}
				}
				break
			}
		} else if zend.Z_TYPE_P(value) == zend.IS_STRING {
			for {
				var __ht *zend.HashTable = zend.Z_ARRVAL_P(array)
				var _p *zend.Bucket = __ht.arData
				var _end *zend.Bucket = _p + __ht.nNumUsed
				for ; _p != _end; _p++ {
					var _z *zend.Zval = &_p.val
					if zend.Z_TYPE_P(_z) == zend.IS_INDIRECT {
						_z = zend.Z_INDIRECT_P(_z)
					}
					if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
						continue
					}
					num_idx = _p.h
					str_idx = _p.key
					entry = _z
					if zend.FastEqualCheckString(value, entry) != 0 {
						if behavior == 0 {
							zend.RETVAL_TRUE
							return
						} else {
							if str_idx != nil {
								zend.RETVAL_STR_COPY(str_idx)
							} else {
								zend.RETVAL_LONG(num_idx)
							}
							return
						}
					}
				}
				break
			}
		} else {
			for {
				var __ht *zend.HashTable = zend.Z_ARRVAL_P(array)
				var _p *zend.Bucket = __ht.arData
				var _end *zend.Bucket = _p + __ht.nNumUsed
				for ; _p != _end; _p++ {
					var _z *zend.Zval = &_p.val
					if zend.Z_TYPE_P(_z) == zend.IS_INDIRECT {
						_z = zend.Z_INDIRECT_P(_z)
					}
					if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
						continue
					}
					num_idx = _p.h
					str_idx = _p.key
					entry = _z
					if zend.FastEqualCheckFunction(value, entry) != 0 {
						if behavior == 0 {
							zend.RETVAL_TRUE
							return
						} else {
							if str_idx != nil {
								zend.RETVAL_STR_COPY(str_idx)
							} else {
								zend.RETVAL_LONG(num_idx)
							}
							return
						}
					}
				}
				break
			}
		}
	}
	zend.RETVAL_FALSE
	return
}
func ZifInArray(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpSearchArray(execute_data, return_value, 0)
}
func ZifArraySearch(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpSearchArray(execute_data, return_value, 1)
}
func PhpValidVarName(var_name *byte, var_name_len int) int {
	/* first 256 bits for first character, and second 256 bits for the next */

	var charset []uint32 = []uint32{0x0, 0x0, 0x87fffffe, 0x7fffffe, 0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff}
	var charset2 []uint32 = []uint32{0x0, 0x3ff0000, 0x87fffffe, 0x7fffffe, 0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff}
	var i int
	var ch uint32
	if zend.UNEXPECTED(var_name_len == 0) {
		return 0
	}

	/* These are allowed as first char: [a-zA-Z_\x7f-\xff] */

	ch = uint32_t((*uint8)(var_name))[0]
	if zend.UNEXPECTED(zend.ZEND_BIT_TEST(charset, ch) == 0) {
		return 0
	}

	/* And these as the rest: [a-zA-Z0-9_\x7f-\xff] */

	if var_name_len > 1 {
		i = 1
		for {
			ch = uint32_t((*uint8)(var_name))[i]
			if zend.UNEXPECTED(zend.ZEND_BIT_TEST(charset2, ch) == 0) {
				return 0
			}
			if b.PreInc(&i) >= var_name_len {
				break
			}
		}
	}
	return 1
}
func PhpPrefixVarname(result *zend.Zval, prefix *zend.Zval, var_name *byte, var_name_len int, add_underscore zend.ZendBool) int {
	zend.ZVAL_NEW_STR(result, zend.ZendStringAlloc(zend.Z_STRLEN_P(prefix)+b.Cond(add_underscore != 0, 1, 0)+var_name_len, 0))
	memcpy(zend.Z_STRVAL_P(result), zend.Z_STRVAL_P(prefix), zend.Z_STRLEN_P(prefix))
	if add_underscore != 0 {
		zend.Z_STRVAL_P(result)[zend.Z_STRLEN_P(prefix)] = '_'
	}
	memcpy(zend.Z_STRVAL_P(result)+zend.Z_STRLEN_P(prefix)+b.Cond(add_underscore != 0, 1, 0), var_name, var_name_len+1)
	return zend.SUCCESS
}
func PhpExtractRefIfExists(arr *zend.ZendArray, symbol_table *zend.ZendArray) zend.ZendLong {
	var count zend.ZendLong = 0
	var var_name *zend.ZendString
	var entry *zend.Zval
	var orig_var *zend.Zval
	for {
		var __ht *zend.HashTable = arr
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val
			if zend.Z_TYPE_P(_z) == zend.IS_INDIRECT {
				_z = zend.Z_INDIRECT_P(_z)
			}
			if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
				continue
			}
			var_name = _p.key
			entry = _z
			if var_name == nil {
				continue
			}
			orig_var = zend.ZendHashFindEx(symbol_table, var_name, 1)
			if orig_var != nil {
				if zend.Z_TYPE_P(orig_var) == zend.IS_INDIRECT {
					orig_var = zend.Z_INDIRECT_P(orig_var)
					if zend.Z_TYPE_P(orig_var) == zend.IS_UNDEF {
						continue
					}
				}
				if PhpValidVarName(zend.ZSTR_VAL(var_name), zend.ZSTR_LEN(var_name)) == 0 {
					continue
				}
				if zend.ZendStringEqualsLiteral(var_name, "GLOBALS") {
					continue
				}
				if zend.ZendStringEqualsLiteral(var_name, "this") {
					zend.ZendThrowError(nil, "Cannot re-assign $this")
					return -1
				}
				if zend.Z_ISREF_P(entry) {
					zend.Z_ADDREF_P(entry)
				} else {
					zend.ZVAL_MAKE_REF_EX(entry, 2)
				}
				zend.ZvalPtrDtor(orig_var)
				zend.ZVAL_REF(orig_var, zend.Z_REF_P(entry))
				count++
			}
		}
		break
	}
	return count
}
func PhpExtractIfExists(arr *zend.ZendArray, symbol_table *zend.ZendArray) zend.ZendLong {
	var count zend.ZendLong = 0
	var var_name *zend.ZendString
	var entry *zend.Zval
	var orig_var *zend.Zval
	for {
		var __ht *zend.HashTable = arr
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val
			if zend.Z_TYPE_P(_z) == zend.IS_INDIRECT {
				_z = zend.Z_INDIRECT_P(_z)
			}
			if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
				continue
			}
			var_name = _p.key
			entry = _z
			if var_name == nil {
				continue
			}
			orig_var = zend.ZendHashFindEx(symbol_table, var_name, 1)
			if orig_var != nil {
				if zend.Z_TYPE_P(orig_var) == zend.IS_INDIRECT {
					orig_var = zend.Z_INDIRECT_P(orig_var)
					if zend.Z_TYPE_P(orig_var) == zend.IS_UNDEF {
						continue
					}
				}
				if PhpValidVarName(zend.ZSTR_VAL(var_name), zend.ZSTR_LEN(var_name)) == 0 {
					continue
				}
				if zend.ZendStringEqualsLiteral(var_name, "GLOBALS") {
					continue
				}
				if zend.ZendStringEqualsLiteral(var_name, "this") {
					zend.ZendThrowError(nil, "Cannot re-assign $this")
					return -1
				}
				zend.ZVAL_DEREF(entry)
				zend.ZEND_TRY_ASSIGN_COPY_EX(orig_var, entry, 0)
				if zend.UNEXPECTED(zend.ExecutorGlobals.exception != nil) {
					return -1
				}
				count++
			}
		}
		break
	}
	return count
}
func PhpExtractRefOverwrite(arr *zend.ZendArray, symbol_table *zend.ZendArray) zend.ZendLong {
	var count zend.ZendLong = 0
	var var_name *zend.ZendString
	var entry *zend.Zval
	var orig_var *zend.Zval
	for {
		var __ht *zend.HashTable = arr
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val
			if zend.Z_TYPE_P(_z) == zend.IS_INDIRECT {
				_z = zend.Z_INDIRECT_P(_z)
			}
			if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
				continue
			}
			var_name = _p.key
			entry = _z
			if var_name == nil {
				continue
			}
			if PhpValidVarName(zend.ZSTR_VAL(var_name), zend.ZSTR_LEN(var_name)) == 0 {
				continue
			}
			if zend.ZendStringEqualsLiteral(var_name, "this") {
				zend.ZendThrowError(nil, "Cannot re-assign $this")
				return -1
			}
			orig_var = zend.ZendHashFindEx(symbol_table, var_name, 1)
			if orig_var != nil {
				if zend.Z_TYPE_P(orig_var) == zend.IS_INDIRECT {
					orig_var = zend.Z_INDIRECT_P(orig_var)
				}
				if zend.ZendStringEqualsLiteral(var_name, "GLOBALS") {
					continue
				}
				if zend.Z_ISREF_P(entry) {
					zend.Z_ADDREF_P(entry)
				} else {
					zend.ZVAL_MAKE_REF_EX(entry, 2)
				}
				zend.ZvalPtrDtor(orig_var)
				zend.ZVAL_REF(orig_var, zend.Z_REF_P(entry))
			} else {
				if zend.Z_ISREF_P(entry) {
					zend.Z_ADDREF_P(entry)
				} else {
					zend.ZVAL_MAKE_REF_EX(entry, 2)
				}
				zend.ZendHashAddNew(symbol_table, var_name, entry)
			}
			count++
		}
		break
	}
	return count
}
func PhpExtractOverwrite(arr *zend.ZendArray, symbol_table *zend.ZendArray) zend.ZendLong {
	var count zend.ZendLong = 0
	var var_name *zend.ZendString
	var entry *zend.Zval
	var orig_var *zend.Zval
	for {
		var __ht *zend.HashTable = arr
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val
			if zend.Z_TYPE_P(_z) == zend.IS_INDIRECT {
				_z = zend.Z_INDIRECT_P(_z)
			}
			if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
				continue
			}
			var_name = _p.key
			entry = _z
			if var_name == nil {
				continue
			}
			if PhpValidVarName(zend.ZSTR_VAL(var_name), zend.ZSTR_LEN(var_name)) == 0 {
				continue
			}
			if zend.ZendStringEqualsLiteral(var_name, "this") {
				zend.ZendThrowError(nil, "Cannot re-assign $this")
				return -1
			}
			orig_var = zend.ZendHashFindEx(symbol_table, var_name, 1)
			if orig_var != nil {
				if zend.Z_TYPE_P(orig_var) == zend.IS_INDIRECT {
					orig_var = zend.Z_INDIRECT_P(orig_var)
				}
				if zend.ZendStringEqualsLiteral(var_name, "GLOBALS") {
					continue
				}
				zend.ZVAL_DEREF(entry)
				zend.ZEND_TRY_ASSIGN_COPY_EX(orig_var, entry, 0)
				if zend.UNEXPECTED(zend.ExecutorGlobals.exception != nil) {
					return -1
				}
			} else {
				zend.ZVAL_DEREF(entry)
				zend.Z_TRY_ADDREF_P(entry)
				zend.ZendHashAddNew(symbol_table, var_name, entry)
			}
			count++
		}
		break
	}
	return count
}
func PhpExtractRefPrefixIfExists(arr *zend.ZendArray, symbol_table *zend.ZendArray, prefix *zend.Zval) zend.ZendLong {
	var count zend.ZendLong = 0
	var var_name *zend.ZendString
	var entry *zend.Zval
	var orig_var *zend.Zval
	var final_name zend.Zval
	for {
		var __ht *zend.HashTable = arr
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val
			if zend.Z_TYPE_P(_z) == zend.IS_INDIRECT {
				_z = zend.Z_INDIRECT_P(_z)
			}
			if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
				continue
			}
			var_name = _p.key
			entry = _z
			if var_name == nil {
				continue
			}
			orig_var = zend.ZendHashFindEx(symbol_table, var_name, 1)
			if orig_var != nil {
				if zend.Z_TYPE_P(orig_var) == zend.IS_INDIRECT {
					orig_var = zend.Z_INDIRECT_P(orig_var)
					if zend.Z_TYPE_P(orig_var) == zend.IS_UNDEF {
						if zend.Z_ISREF_P(entry) {
							zend.Z_ADDREF_P(entry)
						} else {
							zend.ZVAL_MAKE_REF_EX(entry, 2)
						}
						zend.ZVAL_REF(orig_var, zend.Z_REF_P(entry))
						count++
						continue
					}
				}
				PhpPrefixVarname(&final_name, prefix, zend.ZSTR_VAL(var_name), zend.ZSTR_LEN(var_name), 1)
				if PhpValidVarName(zend.Z_STRVAL(final_name), zend.Z_STRLEN(final_name)) != 0 {
					if zend.ZendStringEqualsLiteral(zend.Z_STR(final_name), "this") {
						zend.ZendThrowError(nil, "Cannot re-assign $this")
						return -1
					} else {
						if zend.Z_ISREF_P(entry) {
							zend.Z_ADDREF_P(entry)
						} else {
							zend.ZVAL_MAKE_REF_EX(entry, 2)
						}
						if b.Assign(&orig_var, zend.ZendHashFind(symbol_table, zend.Z_STR(final_name))) != nil {
							if zend.Z_TYPE_P(orig_var) == zend.IS_INDIRECT {
								orig_var = zend.Z_INDIRECT_P(orig_var)
							}
							zend.ZvalPtrDtor(orig_var)
							zend.ZVAL_REF(orig_var, zend.Z_REF_P(entry))
						} else {
							zend.ZendHashAddNew(symbol_table, zend.Z_STR(final_name), entry)
						}
						count++
					}
				}
				zend.ZvalPtrDtorStr(&final_name)
			}
		}
		break
	}
	return count
}
func PhpExtractPrefixIfExists(arr *zend.ZendArray, symbol_table *zend.ZendArray, prefix *zend.Zval) zend.ZendLong {
	var count zend.ZendLong = 0
	var var_name *zend.ZendString
	var entry *zend.Zval
	var orig_var *zend.Zval
	var final_name zend.Zval
	for {
		var __ht *zend.HashTable = arr
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val
			if zend.Z_TYPE_P(_z) == zend.IS_INDIRECT {
				_z = zend.Z_INDIRECT_P(_z)
			}
			if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
				continue
			}
			var_name = _p.key
			entry = _z
			if var_name == nil {
				continue
			}
			orig_var = zend.ZendHashFindEx(symbol_table, var_name, 1)
			if orig_var != nil {
				if zend.Z_TYPE_P(orig_var) == zend.IS_INDIRECT {
					orig_var = zend.Z_INDIRECT_P(orig_var)
					if zend.Z_TYPE_P(orig_var) == zend.IS_UNDEF {
						zend.ZVAL_COPY_DEREF(orig_var, entry)
						count++
						continue
					}
				}
				PhpPrefixVarname(&final_name, prefix, zend.ZSTR_VAL(var_name), zend.ZSTR_LEN(var_name), 1)
				if PhpValidVarName(zend.Z_STRVAL(final_name), zend.Z_STRLEN(final_name)) != 0 {
					if zend.ZendStringEqualsLiteral(zend.Z_STR(final_name), "this") {
						zend.ZendThrowError(nil, "Cannot re-assign $this")
						return -1
					} else {
						zend.ZVAL_DEREF(entry)
						if b.Assign(&orig_var, zend.ZendHashFind(symbol_table, zend.Z_STR(final_name))) != nil {
							if zend.Z_TYPE_P(orig_var) == zend.IS_INDIRECT {
								orig_var = zend.Z_INDIRECT_P(orig_var)
							}
							zend.ZEND_TRY_ASSIGN_COPY_EX(orig_var, entry, 0)
							if zend.UNEXPECTED(zend.ExecutorGlobals.exception != nil) {
								zend.ZendStringReleaseEx(zend.Z_STR(final_name), 0)
								return -1
							}
						} else {
							zend.Z_TRY_ADDREF_P(entry)
							zend.ZendHashAddNew(symbol_table, zend.Z_STR(final_name), entry)
						}
						count++
					}
				}
				zend.ZvalPtrDtorStr(&final_name)
			}
		}
		break
	}
	return count
}
func PhpExtractRefPrefixSame(arr *zend.ZendArray, symbol_table *zend.ZendArray, prefix *zend.Zval) zend.ZendLong {
	var count zend.ZendLong = 0
	var var_name *zend.ZendString
	var entry *zend.Zval
	var orig_var *zend.Zval
	var final_name zend.Zval
	for {
		var __ht *zend.HashTable = arr
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val
			if zend.Z_TYPE_P(_z) == zend.IS_INDIRECT {
				_z = zend.Z_INDIRECT_P(_z)
			}
			if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
				continue
			}
			var_name = _p.key
			entry = _z
			if var_name == nil {
				continue
			}
			if zend.ZSTR_LEN(var_name) == 0 {
				continue
			}
			orig_var = zend.ZendHashFindEx(symbol_table, var_name, 1)
			if orig_var != nil {
				if zend.Z_TYPE_P(orig_var) == zend.IS_INDIRECT {
					orig_var = zend.Z_INDIRECT_P(orig_var)
					if zend.Z_TYPE_P(orig_var) == zend.IS_UNDEF {
						if zend.Z_ISREF_P(entry) {
							zend.Z_ADDREF_P(entry)
						} else {
							zend.ZVAL_MAKE_REF_EX(entry, 2)
						}
						zend.ZVAL_REF(orig_var, zend.Z_REF_P(entry))
						count++
						continue
					}
				}
			prefix:
				PhpPrefixVarname(&final_name, prefix, zend.ZSTR_VAL(var_name), zend.ZSTR_LEN(var_name), 1)
				if PhpValidVarName(zend.Z_STRVAL(final_name), zend.Z_STRLEN(final_name)) != 0 {
					if zend.ZendStringEqualsLiteral(zend.Z_STR(final_name), "this") {
						zend.ZendThrowError(nil, "Cannot re-assign $this")
						return -1
					} else {
						if zend.Z_ISREF_P(entry) {
							zend.Z_ADDREF_P(entry)
						} else {
							zend.ZVAL_MAKE_REF_EX(entry, 2)
						}
						if b.Assign(&orig_var, zend.ZendHashFind(symbol_table, zend.Z_STR(final_name))) != nil {
							if zend.Z_TYPE_P(orig_var) == zend.IS_INDIRECT {
								orig_var = zend.Z_INDIRECT_P(orig_var)
							}
							zend.ZvalPtrDtor(orig_var)
							zend.ZVAL_REF(orig_var, zend.Z_REF_P(entry))
						} else {
							zend.ZendHashAddNew(symbol_table, zend.Z_STR(final_name), entry)
						}
						count++
					}
				}
				zend.ZvalPtrDtorStr(&final_name)
			} else {
				if PhpValidVarName(zend.ZSTR_VAL(var_name), zend.ZSTR_LEN(var_name)) == 0 {
					continue
				}
				if zend.ZendStringEqualsLiteral(var_name, "this") {
					goto prefix
				}
				if zend.Z_ISREF_P(entry) {
					zend.Z_ADDREF_P(entry)
				} else {
					zend.ZVAL_MAKE_REF_EX(entry, 2)
				}
				zend.ZendHashAddNew(symbol_table, var_name, entry)
				count++
			}
		}
		break
	}
	return count
}
func PhpExtractPrefixSame(arr *zend.ZendArray, symbol_table *zend.ZendArray, prefix *zend.Zval) zend.ZendLong {
	var count zend.ZendLong = 0
	var var_name *zend.ZendString
	var entry *zend.Zval
	var orig_var *zend.Zval
	var final_name zend.Zval
	for {
		var __ht *zend.HashTable = arr
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val
			if zend.Z_TYPE_P(_z) == zend.IS_INDIRECT {
				_z = zend.Z_INDIRECT_P(_z)
			}
			if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
				continue
			}
			var_name = _p.key
			entry = _z
			if var_name == nil {
				continue
			}
			if zend.ZSTR_LEN(var_name) == 0 {
				continue
			}
			orig_var = zend.ZendHashFindEx(symbol_table, var_name, 1)
			if orig_var != nil {
				if zend.Z_TYPE_P(orig_var) == zend.IS_INDIRECT {
					orig_var = zend.Z_INDIRECT_P(orig_var)
					if zend.Z_TYPE_P(orig_var) == zend.IS_UNDEF {
						zend.ZVAL_COPY_DEREF(orig_var, entry)
						count++
						continue
					}
				}
			prefix:
				PhpPrefixVarname(&final_name, prefix, zend.ZSTR_VAL(var_name), zend.ZSTR_LEN(var_name), 1)
				if PhpValidVarName(zend.Z_STRVAL(final_name), zend.Z_STRLEN(final_name)) != 0 {
					if zend.ZendStringEqualsLiteral(zend.Z_STR(final_name), "this") {
						zend.ZendThrowError(nil, "Cannot re-assign $this")
						return -1
					} else {
						zend.ZVAL_DEREF(entry)
						if b.Assign(&orig_var, zend.ZendHashFind(symbol_table, zend.Z_STR(final_name))) != nil {
							if zend.Z_TYPE_P(orig_var) == zend.IS_INDIRECT {
								orig_var = zend.Z_INDIRECT_P(orig_var)
							}
							zend.ZEND_TRY_ASSIGN_COPY_EX(orig_var, entry, 0)
							if zend.UNEXPECTED(zend.ExecutorGlobals.exception != nil) {
								zend.ZendStringReleaseEx(zend.Z_STR(final_name), 0)
								return -1
							}
						} else {
							zend.Z_TRY_ADDREF_P(entry)
							zend.ZendHashAddNew(symbol_table, zend.Z_STR(final_name), entry)
						}
						count++
					}
				}
				zend.ZvalPtrDtorStr(&final_name)
			} else {
				if PhpValidVarName(zend.ZSTR_VAL(var_name), zend.ZSTR_LEN(var_name)) == 0 {
					continue
				}
				if zend.ZendStringEqualsLiteral(var_name, "this") {
					goto prefix
				}
				zend.ZVAL_DEREF(entry)
				zend.Z_TRY_ADDREF_P(entry)
				zend.ZendHashAddNew(symbol_table, var_name, entry)
				count++
			}
		}
		break
	}
	return count
}
func PhpExtractRefPrefixAll(arr *zend.ZendArray, symbol_table *zend.ZendArray, prefix *zend.Zval) zend.ZendLong {
	var count zend.ZendLong = 0
	var var_name *zend.ZendString
	var num_key zend.ZendUlong
	var entry *zend.Zval
	var orig_var *zend.Zval
	var final_name zend.Zval
	for {
		var __ht *zend.HashTable = arr
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val
			if zend.Z_TYPE_P(_z) == zend.IS_INDIRECT {
				_z = zend.Z_INDIRECT_P(_z)
			}
			if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
				continue
			}
			num_key = _p.h
			var_name = _p.key
			entry = _z
			if var_name != nil {
				if zend.ZSTR_LEN(var_name) == 0 {
					continue
				}
				PhpPrefixVarname(&final_name, prefix, zend.ZSTR_VAL(var_name), zend.ZSTR_LEN(var_name), 1)
			} else {
				var str *zend.ZendString = zend.ZendLongToStr(num_key)
				PhpPrefixVarname(&final_name, prefix, zend.ZSTR_VAL(str), zend.ZSTR_LEN(str), 1)
				zend.ZendStringReleaseEx(str, 0)
			}
			if PhpValidVarName(zend.Z_STRVAL(final_name), zend.Z_STRLEN(final_name)) != 0 {
				if zend.ZendStringEqualsLiteral(zend.Z_STR(final_name), "this") {
					zend.ZendThrowError(nil, "Cannot re-assign $this")
					return -1
				} else {
					if zend.Z_ISREF_P(entry) {
						zend.Z_ADDREF_P(entry)
					} else {
						zend.ZVAL_MAKE_REF_EX(entry, 2)
					}
					if b.Assign(&orig_var, zend.ZendHashFind(symbol_table, zend.Z_STR(final_name))) != nil {
						if zend.Z_TYPE_P(orig_var) == zend.IS_INDIRECT {
							orig_var = zend.Z_INDIRECT_P(orig_var)
						}
						zend.ZvalPtrDtor(orig_var)
						zend.ZVAL_REF(orig_var, zend.Z_REF_P(entry))
					} else {
						zend.ZendHashAddNew(symbol_table, zend.Z_STR(final_name), entry)
					}
					count++
				}
			}
			zend.ZvalPtrDtorStr(&final_name)
		}
		break
	}
	return count
}
func PhpExtractPrefixAll(arr *zend.ZendArray, symbol_table *zend.ZendArray, prefix *zend.Zval) zend.ZendLong {
	var count zend.ZendLong = 0
	var var_name *zend.ZendString
	var num_key zend.ZendUlong
	var entry *zend.Zval
	var orig_var *zend.Zval
	var final_name zend.Zval
	for {
		var __ht *zend.HashTable = arr
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val
			if zend.Z_TYPE_P(_z) == zend.IS_INDIRECT {
				_z = zend.Z_INDIRECT_P(_z)
			}
			if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
				continue
			}
			num_key = _p.h
			var_name = _p.key
			entry = _z
			if var_name != nil {
				if zend.ZSTR_LEN(var_name) == 0 {
					continue
				}
				PhpPrefixVarname(&final_name, prefix, zend.ZSTR_VAL(var_name), zend.ZSTR_LEN(var_name), 1)
			} else {
				var str *zend.ZendString = zend.ZendLongToStr(num_key)
				PhpPrefixVarname(&final_name, prefix, zend.ZSTR_VAL(str), zend.ZSTR_LEN(str), 1)
				zend.ZendStringReleaseEx(str, 0)
			}
			if PhpValidVarName(zend.Z_STRVAL(final_name), zend.Z_STRLEN(final_name)) != 0 {
				if zend.ZendStringEqualsLiteral(zend.Z_STR(final_name), "this") {
					zend.ZendThrowError(nil, "Cannot re-assign $this")
					return -1
				} else {
					zend.ZVAL_DEREF(entry)
					if b.Assign(&orig_var, zend.ZendHashFind(symbol_table, zend.Z_STR(final_name))) != nil {
						if zend.Z_TYPE_P(orig_var) == zend.IS_INDIRECT {
							orig_var = zend.Z_INDIRECT_P(orig_var)
						}
						zend.ZEND_TRY_ASSIGN_COPY_EX(orig_var, entry, 0)
						if zend.UNEXPECTED(zend.ExecutorGlobals.exception != nil) {
							zend.ZendStringReleaseEx(zend.Z_STR(final_name), 0)
							return -1
						}
					} else {
						zend.Z_TRY_ADDREF_P(entry)
						zend.ZendHashAddNew(symbol_table, zend.Z_STR(final_name), entry)
					}
					count++
				}
			}
			zend.ZvalPtrDtorStr(&final_name)
		}
		break
	}
	return count
}
func PhpExtractRefPrefixInvalid(arr *zend.ZendArray, symbol_table *zend.ZendArray, prefix *zend.Zval) zend.ZendLong {
	var count zend.ZendLong = 0
	var var_name *zend.ZendString
	var num_key zend.ZendUlong
	var entry *zend.Zval
	var orig_var *zend.Zval
	var final_name zend.Zval
	for {
		var __ht *zend.HashTable = arr
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val
			if zend.Z_TYPE_P(_z) == zend.IS_INDIRECT {
				_z = zend.Z_INDIRECT_P(_z)
			}
			if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
				continue
			}
			num_key = _p.h
			var_name = _p.key
			entry = _z
			if var_name != nil {
				if PhpValidVarName(zend.ZSTR_VAL(var_name), zend.ZSTR_LEN(var_name)) == 0 || zend.ZendStringEqualsLiteral(var_name, "this") {
					PhpPrefixVarname(&final_name, prefix, zend.ZSTR_VAL(var_name), zend.ZSTR_LEN(var_name), 1)
					if PhpValidVarName(zend.Z_STRVAL(final_name), zend.Z_STRLEN(final_name)) == 0 {
						zend.ZvalPtrDtorStr(&final_name)
						continue
					}
				} else {
					zend.ZVAL_STR_COPY(&final_name, var_name)
				}
			} else {
				var str *zend.ZendString = zend.ZendLongToStr(num_key)
				PhpPrefixVarname(&final_name, prefix, zend.ZSTR_VAL(str), zend.ZSTR_LEN(str), 1)
				zend.ZendStringReleaseEx(str, 0)
				if PhpValidVarName(zend.Z_STRVAL(final_name), zend.Z_STRLEN(final_name)) == 0 {
					zend.ZvalPtrDtorStr(&final_name)
					continue
				}
			}
			if zend.ZendStringEqualsLiteral(zend.Z_STR(final_name), "this") {
				zend.ZendThrowError(nil, "Cannot re-assign $this")
				return -1
			} else {
				if zend.Z_ISREF_P(entry) {
					zend.Z_ADDREF_P(entry)
				} else {
					zend.ZVAL_MAKE_REF_EX(entry, 2)
				}
				if b.Assign(&orig_var, zend.ZendHashFind(symbol_table, zend.Z_STR(final_name))) != nil {
					if zend.Z_TYPE_P(orig_var) == zend.IS_INDIRECT {
						orig_var = zend.Z_INDIRECT_P(orig_var)
					}
					zend.ZvalPtrDtor(orig_var)
					zend.ZVAL_REF(orig_var, zend.Z_REF_P(entry))
				} else {
					zend.ZendHashAddNew(symbol_table, zend.Z_STR(final_name), entry)
				}
				count++
			}
			zend.ZvalPtrDtorStr(&final_name)
		}
		break
	}
	return count
}
func PhpExtractPrefixInvalid(arr *zend.ZendArray, symbol_table *zend.ZendArray, prefix *zend.Zval) zend.ZendLong {
	var count zend.ZendLong = 0
	var var_name *zend.ZendString
	var num_key zend.ZendUlong
	var entry *zend.Zval
	var orig_var *zend.Zval
	var final_name zend.Zval
	for {
		var __ht *zend.HashTable = arr
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val
			if zend.Z_TYPE_P(_z) == zend.IS_INDIRECT {
				_z = zend.Z_INDIRECT_P(_z)
			}
			if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
				continue
			}
			num_key = _p.h
			var_name = _p.key
			entry = _z
			if var_name != nil {
				if PhpValidVarName(zend.ZSTR_VAL(var_name), zend.ZSTR_LEN(var_name)) == 0 || zend.ZendStringEqualsLiteral(var_name, "this") {
					PhpPrefixVarname(&final_name, prefix, zend.ZSTR_VAL(var_name), zend.ZSTR_LEN(var_name), 1)
					if PhpValidVarName(zend.Z_STRVAL(final_name), zend.Z_STRLEN(final_name)) == 0 {
						zend.ZvalPtrDtorStr(&final_name)
						continue
					}
				} else {
					zend.ZVAL_STR_COPY(&final_name, var_name)
				}
			} else {
				var str *zend.ZendString = zend.ZendLongToStr(num_key)
				PhpPrefixVarname(&final_name, prefix, zend.ZSTR_VAL(str), zend.ZSTR_LEN(str), 1)
				zend.ZendStringReleaseEx(str, 0)
				if PhpValidVarName(zend.Z_STRVAL(final_name), zend.Z_STRLEN(final_name)) == 0 {
					zend.ZvalPtrDtorStr(&final_name)
					continue
				}
			}
			if zend.ZendStringEqualsLiteral(zend.Z_STR(final_name), "this") {
				zend.ZendThrowError(nil, "Cannot re-assign $this")
				return -1
			} else {
				zend.ZVAL_DEREF(entry)
				if b.Assign(&orig_var, zend.ZendHashFind(symbol_table, zend.Z_STR(final_name))) != nil {
					if zend.Z_TYPE_P(orig_var) == zend.IS_INDIRECT {
						orig_var = zend.Z_INDIRECT_P(orig_var)
					}
					zend.ZEND_TRY_ASSIGN_COPY_EX(orig_var, entry, 0)
					if zend.UNEXPECTED(zend.ExecutorGlobals.exception != nil) {
						zend.ZendStringReleaseEx(zend.Z_STR(final_name), 0)
						return -1
					}
				} else {
					zend.Z_TRY_ADDREF_P(entry)
					zend.ZendHashAddNew(symbol_table, zend.Z_STR(final_name), entry)
				}
				count++
			}
			zend.ZvalPtrDtorStr(&final_name)
		}
		break
	}
	return count
}
func PhpExtractRefSkip(arr *zend.ZendArray, symbol_table *zend.ZendArray) zend.ZendLong {
	var count zend.ZendLong = 0
	var var_name *zend.ZendString
	var entry *zend.Zval
	var orig_var *zend.Zval
	for {
		var __ht *zend.HashTable = arr
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val
			if zend.Z_TYPE_P(_z) == zend.IS_INDIRECT {
				_z = zend.Z_INDIRECT_P(_z)
			}
			if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
				continue
			}
			var_name = _p.key
			entry = _z
			if var_name == nil {
				continue
			}
			if PhpValidVarName(zend.ZSTR_VAL(var_name), zend.ZSTR_LEN(var_name)) == 0 {
				continue
			}
			if zend.ZendStringEqualsLiteral(var_name, "this") {
				continue
			}
			orig_var = zend.ZendHashFindEx(symbol_table, var_name, 1)
			if orig_var != nil {
				if zend.Z_TYPE_P(orig_var) == zend.IS_INDIRECT {
					orig_var = zend.Z_INDIRECT_P(orig_var)
					if zend.Z_TYPE_P(orig_var) == zend.IS_UNDEF {
						if zend.Z_ISREF_P(entry) {
							zend.Z_ADDREF_P(entry)
						} else {
							zend.ZVAL_MAKE_REF_EX(entry, 2)
						}
						zend.ZVAL_REF(orig_var, zend.Z_REF_P(entry))
						count++
					}
				}
			} else {
				if zend.Z_ISREF_P(entry) {
					zend.Z_ADDREF_P(entry)
				} else {
					zend.ZVAL_MAKE_REF_EX(entry, 2)
				}
				zend.ZendHashAddNew(symbol_table, var_name, entry)
				count++
			}
		}
		break
	}
	return count
}
func PhpExtractSkip(arr *zend.ZendArray, symbol_table *zend.ZendArray) zend.ZendLong {
	var count zend.ZendLong = 0
	var var_name *zend.ZendString
	var entry *zend.Zval
	var orig_var *zend.Zval
	for {
		var __ht *zend.HashTable = arr
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val
			if zend.Z_TYPE_P(_z) == zend.IS_INDIRECT {
				_z = zend.Z_INDIRECT_P(_z)
			}
			if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
				continue
			}
			var_name = _p.key
			entry = _z
			if var_name == nil {
				continue
			}
			if PhpValidVarName(zend.ZSTR_VAL(var_name), zend.ZSTR_LEN(var_name)) == 0 {
				continue
			}
			if zend.ZendStringEqualsLiteral(var_name, "this") {
				continue
			}
			orig_var = zend.ZendHashFindEx(symbol_table, var_name, 1)
			if orig_var != nil {
				if zend.Z_TYPE_P(orig_var) == zend.IS_INDIRECT {
					orig_var = zend.Z_INDIRECT_P(orig_var)
					if zend.Z_TYPE_P(orig_var) == zend.IS_UNDEF {
						zend.ZVAL_COPY_DEREF(orig_var, entry)
						count++
					}
				}
			} else {
				zend.ZVAL_DEREF(entry)
				zend.Z_TRY_ADDREF_P(entry)
				zend.ZendHashAddNew(symbol_table, var_name, entry)
				count++
			}
		}
		break
	}
	return count
}
func ZifExtract(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var var_array_param *zend.Zval
	var prefix *zend.Zval = nil
	var extract_refs zend.ZendLong
	var extract_type zend.ZendLong = EXTR_OVERWRITE
	var count zend.ZendLong
	var symbol_table *zend.ZendArray
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 3
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(1, 0)
			if zend.UNEXPECTED(zend.ZendParseArgArray(_arg, &var_array_param, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &extract_type, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &prefix, 0)
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	extract_refs = extract_type & EXTR_REFS
	if extract_refs != 0 {
		zend.SEPARATE_ARRAY(var_array_param)
	}
	extract_type &= 0xff
	if extract_type < EXTR_OVERWRITE || extract_type > EXTR_IF_EXISTS {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Invalid extract type")
		return
	}
	if extract_type > EXTR_SKIP && extract_type <= EXTR_PREFIX_IF_EXISTS && zend.ZEND_NUM_ARGS() < 3 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "specified extract type requires the prefix parameter")
		return
	}
	if prefix != nil {
		if zend.TryConvertToString(prefix) == 0 {
			return
		}
		if zend.Z_STRLEN_P(prefix) != 0 && PhpValidVarName(zend.Z_STRVAL_P(prefix), zend.Z_STRLEN_P(prefix)) == 0 {
			core.PhpErrorDocref(nil, zend.E_WARNING, "prefix is not a valid identifier")
			return
		}
	}
	if zend.ZendForbidDynamicCall("extract()") == zend.FAILURE {
		return
	}
	symbol_table = zend.ZendRebuildSymbolTable()
	if extract_refs != 0 {
		switch extract_type {
		case EXTR_IF_EXISTS:
			count = PhpExtractRefIfExists(zend.Z_ARRVAL_P(var_array_param), symbol_table)
			break
		case EXTR_OVERWRITE:
			count = PhpExtractRefOverwrite(zend.Z_ARRVAL_P(var_array_param), symbol_table)
			break
		case EXTR_PREFIX_IF_EXISTS:
			count = PhpExtractRefPrefixIfExists(zend.Z_ARRVAL_P(var_array_param), symbol_table, prefix)
			break
		case EXTR_PREFIX_SAME:
			count = PhpExtractRefPrefixSame(zend.Z_ARRVAL_P(var_array_param), symbol_table, prefix)
			break
		case EXTR_PREFIX_ALL:
			count = PhpExtractRefPrefixAll(zend.Z_ARRVAL_P(var_array_param), symbol_table, prefix)
			break
		case EXTR_PREFIX_INVALID:
			count = PhpExtractRefPrefixInvalid(zend.Z_ARRVAL_P(var_array_param), symbol_table, prefix)
			break
		default:
			count = PhpExtractRefSkip(zend.Z_ARRVAL_P(var_array_param), symbol_table)
			break
		}
	} else {

		/* The array might be stored in a local variable that will be overwritten */

		var array_copy zend.Zval
		zend.ZVAL_COPY(&array_copy, var_array_param)
		switch extract_type {
		case EXTR_IF_EXISTS:
			count = PhpExtractIfExists(zend.Z_ARRVAL(array_copy), symbol_table)
			break
		case EXTR_OVERWRITE:
			count = PhpExtractOverwrite(zend.Z_ARRVAL(array_copy), symbol_table)
			break
		case EXTR_PREFIX_IF_EXISTS:
			count = PhpExtractPrefixIfExists(zend.Z_ARRVAL(array_copy), symbol_table, prefix)
			break
		case EXTR_PREFIX_SAME:
			count = PhpExtractPrefixSame(zend.Z_ARRVAL(array_copy), symbol_table, prefix)
			break
		case EXTR_PREFIX_ALL:
			count = PhpExtractPrefixAll(zend.Z_ARRVAL(array_copy), symbol_table, prefix)
			break
		case EXTR_PREFIX_INVALID:
			count = PhpExtractPrefixInvalid(zend.Z_ARRVAL(array_copy), symbol_table, prefix)
			break
		default:
			count = PhpExtractSkip(zend.Z_ARRVAL(array_copy), symbol_table)
			break
		}
		zend.ZvalPtrDtor(&array_copy)
	}
	zend.RETVAL_LONG(count)
	return
}
func PhpCompactVar(eg_active_symbol_table *zend.HashTable, return_value *zend.Zval, entry *zend.Zval) {
	var value_ptr *zend.Zval
	var data zend.Zval
	zend.ZVAL_DEREF(entry)
	if zend.Z_TYPE_P(entry) == zend.IS_STRING {
		if b.Assign(&value_ptr, zend.ZendHashFindInd(eg_active_symbol_table, zend.Z_STR_P(entry))) != nil {
			zend.ZVAL_DEREF(value_ptr)
			zend.Z_TRY_ADDREF_P(value_ptr)
			zend.ZendHashUpdate(zend.Z_ARRVAL_P(return_value), zend.Z_STR_P(entry), value_ptr)
		} else if zend.ZendStringEqualsLiteral(zend.Z_STR_P(entry), "this") {
			var object *zend.ZendObject = zend.ZendGetThisObject(zend.ExecutorGlobals.current_execute_data)
			if object != nil {
				zend.GC_ADDREF(object)
				zend.ZVAL_OBJ(&data, object)
				zend.ZendHashUpdate(zend.Z_ARRVAL_P(return_value), zend.Z_STR_P(entry), &data)
			}
		} else {
			core.PhpErrorDocref(nil, zend.E_NOTICE, "Undefined variable: %s", zend.ZSTR_VAL(zend.Z_STR_P(entry)))
		}
	} else if zend.Z_TYPE_P(entry) == zend.IS_ARRAY {
		if zend.Z_REFCOUNTED_P(entry) {
			if zend.Z_IS_RECURSIVE_P(entry) != 0 {
				core.PhpErrorDocref(nil, zend.E_WARNING, "recursion detected")
				return
			}
			zend.Z_PROTECT_RECURSION_P(entry)
		}
		for {
			var __ht *zend.HashTable = zend.Z_ARRVAL_P(entry)
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val
				if zend.Z_TYPE_P(_z) == zend.IS_INDIRECT {
					_z = zend.Z_INDIRECT_P(_z)
				}
				if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
					continue
				}
				value_ptr = _z
				PhpCompactVar(eg_active_symbol_table, return_value, value_ptr)
			}
			break
		}
		if zend.Z_REFCOUNTED_P(entry) {
			zend.Z_UNPROTECT_RECURSION_P(entry)
		}
	}
}
func ZifCompact(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var args *zend.Zval = nil
	var num_args uint32
	var i uint32
	var symbol_table *zend.ZendArray
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = -1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			var _num_varargs int = _num_args - _i - 0
			if zend.EXPECTED(_num_varargs > 0) {
				args = _real_arg + 1
				num_args = _num_varargs
				_i += _num_varargs
				_real_arg += _num_varargs
			} else {
				args = nil
				num_args = 0
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	if zend.ZendForbidDynamicCall("compact()") == zend.FAILURE {
		return
	}
	symbol_table = zend.ZendRebuildSymbolTable()
	if zend.UNEXPECTED(symbol_table == nil) {
		return
	}

	/* compact() is probably most used with a single array of var_names
	   or multiple string names, rather than a combination of both.
	   So quickly guess a minimum result size based on that */

	if num_args != 0 && zend.Z_TYPE(args[0]) == zend.IS_ARRAY {
		zend.ArrayInitSize(return_value, zend.ZendHashNumElements(zend.Z_ARRVAL(args[0])))
	} else {
		zend.ArrayInitSize(return_value, num_args)
	}
	for i = 0; i < num_args; i++ {
		PhpCompactVar(symbol_table, return_value, &args[i])
	}
}
func ZifArrayFill(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var val *zend.Zval
	var start_key zend.ZendLong
	var num zend.ZendLong
	for {
		var _flags int = 0
		var _min_num_args int = 3
		var _max_num_args int = 3
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &start_key, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &num, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &val, 0)
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	if zend.EXPECTED(num > 0) {
		if b.SizeOf("num") > 4 && zend.UNEXPECTED(zend.EXPECTED(num > 0x7fffffff)) {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Too many elements")
			zend.RETVAL_FALSE
			return
		} else if zend.UNEXPECTED(start_key > zend.ZEND_LONG_MAX-num+1) {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Cannot add element to the array as the next element is already occupied")
			zend.RETVAL_FALSE
			return
		} else if zend.EXPECTED(start_key >= 0) && zend.EXPECTED(start_key < num) {

			/* create packed array */

			var p *zend.Bucket
			var n zend.ZendLong
			zend.ArrayInitSize(return_value, uint32_t(start_key+num))
			zend.ZendHashRealInitPacked(zend.Z_ARRVAL_P(return_value))
			zend.Z_ARRVAL_P(return_value).nNumUsed = uint32_t(start_key + num)
			zend.Z_ARRVAL_P(return_value).nNumOfElements = uint32(num)
			zend.Z_ARRVAL_P(return_value).nNextFreeElement = zend_long(start_key + num)
			if zend.Z_REFCOUNTED_P(val) {
				zend.GC_ADDREF_EX(zend.Z_COUNTED_P(val), uint32(num))
			}
			p = zend.Z_ARRVAL_P(return_value).arData
			n = start_key
			for b.PostDec(&start_key) {
				zend.ZVAL_UNDEF(&p.val)
				p++
			}
			for b.PostDec(&num) {
				zend.ZVAL_COPY_VALUE(&p.val, val)
				n++
				p.h = n - 1
				p.key = nil
				p++
			}
		} else {

			/* create hash */

			zend.ArrayInitSize(return_value, uint32(num))
			zend.ZendHashRealInitMixed(zend.Z_ARRVAL_P(return_value))
			if zend.Z_REFCOUNTED_P(val) {
				zend.GC_ADDREF_EX(zend.Z_COUNTED_P(val), uint32(num))
			}
			zend.ZendHashIndexAddNew(zend.Z_ARRVAL_P(return_value), start_key, val)
			for b.PreDec(&num) {
				zend.ZendHashNextIndexInsertNew(zend.Z_ARRVAL_P(return_value), val)
				start_key++
			}
		}
	} else if zend.EXPECTED(num == 0) {
		zend.RETVAL_EMPTY_ARRAY()
		return
	} else {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Number of elements can't be negative")
		zend.RETVAL_FALSE
		return
	}
}
func ZifArrayFillKeys(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var keys *zend.Zval
	var val *zend.Zval
	var entry *zend.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 2
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgArray(_arg, &keys, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &val, 0)
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}

	/* Initialize return array */

	zend.ArrayInitSize(return_value, zend.ZendHashNumElements(zend.Z_ARRVAL_P(keys)))
	for {
		var __ht *zend.HashTable = zend.Z_ARRVAL_P(keys)
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
				continue
			}
			entry = _z
			zend.ZVAL_DEREF(entry)
			zend.Z_TRY_ADDREF_P(val)
			if zend.Z_TYPE_P(entry) == zend.IS_LONG {
				zend.ZendHashIndexUpdate(zend.Z_ARRVAL_P(return_value), zend.Z_LVAL_P(entry), val)
			} else {
				var tmp_key *zend.ZendString
				var key *zend.ZendString = zend.ZvalGetTmpString(entry, &tmp_key)
				zend.ZendSymtableUpdate(zend.Z_ARRVAL_P(return_value), key, val)
				zend.ZendTmpStringRelease(tmp_key)
			}
		}
		break
	}
}
func ZifRange(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var zlow *zend.Zval
	var zhigh *zend.Zval
	var zstep *zend.Zval = nil
	var tmp zend.Zval
	var err int = 0
	var is_step_double int = 0
	var step float64 = 1.0
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 3
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &zlow, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &zhigh, 0)
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &zstep, 0)
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			zend.RETVAL_FALSE
			return
		}
		break
	}
	if zstep != nil {
		if zend.Z_TYPE_P(zstep) == zend.IS_DOUBLE {
			is_step_double = 1
		} else if zend.Z_TYPE_P(zstep) == zend.IS_STRING {
			var type_ int = zend.IsNumericString(zend.Z_STRVAL_P(zstep), zend.Z_STRLEN_P(zstep), nil, nil, 0)
			if type_ == zend.IS_DOUBLE {
				is_step_double = 1
			}
			if type_ == 0 {

				/* bad number */

				core.PhpErrorDocref(nil, zend.E_WARNING, "Invalid range string - must be numeric")
				zend.RETVAL_FALSE
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

	if zend.Z_TYPE_P(zlow) == zend.IS_STRING && zend.Z_TYPE_P(zhigh) == zend.IS_STRING && zend.Z_STRLEN_P(zlow) >= 1 && zend.Z_STRLEN_P(zhigh) >= 1 {
		var type1 int
		var type2 int
		var low uint8
		var high uint8
		var lstep zend.ZendLong = zend.ZendLong(step)
		type1 = zend.IsNumericString(zend.Z_STRVAL_P(zlow), zend.Z_STRLEN_P(zlow), nil, nil, 0)
		type2 = zend.IsNumericString(zend.Z_STRVAL_P(zhigh), zend.Z_STRLEN_P(zhigh), nil, nil, 0)
		if type1 == zend.IS_DOUBLE || type2 == zend.IS_DOUBLE || is_step_double != 0 {
			goto double_str
		} else if type1 == zend.IS_LONG || type2 == zend.IS_LONG {
			goto long_str
		}
		low = uint8(zend.Z_STRVAL_P(zlow)[0])
		high = uint8(zend.Z_STRVAL_P(zhigh)[0])
		if low > high {
			if lstep <= 0 {
				err = 1
				goto err
			}

			/* Initialize the return_value as an array. */

			zend.ArrayInitSize(return_value, uint32_t((low-high)/lstep+1))
			zend.ZendHashRealInitPacked(zend.Z_ARRVAL_P(return_value))
			for {
				var __fill_ht *zend.HashTable = zend.Z_ARRVAL_P(return_value)
				var __fill_bkt *zend.Bucket = __fill_ht.arData + __fill_ht.nNumUsed
				var __fill_idx uint32 = __fill_ht.nNumUsed
				zend.ZEND_ASSERT((zend.HT_FLAGS(__fill_ht) & zend.HASH_FLAG_PACKED) != 0)
				for ; low >= high; low -= uint(lstep) {
					zend.ZEND_HASH_FILL_SET_INTERNED_STR(zend.ZSTR_CHAR(low))
					zend.ZEND_HASH_FILL_NEXT()
					if signed__int(low-lstep) < 0 {
						break
					}
				}
				__fill_ht.nNumUsed = __fill_idx
				__fill_ht.nNumOfElements = __fill_idx
				__fill_ht.nNextFreeElement = __fill_idx
				__fill_ht.nInternalPointer = 0
				break
			}
		} else if high > low {
			if lstep <= 0 {
				err = 1
				goto err
			}
			zend.ArrayInitSize(return_value, uint32_t((high-low)/lstep+1))
			zend.ZendHashRealInitPacked(zend.Z_ARRVAL_P(return_value))
			for {
				var __fill_ht *zend.HashTable = zend.Z_ARRVAL_P(return_value)
				var __fill_bkt *zend.Bucket = __fill_ht.arData + __fill_ht.nNumUsed
				var __fill_idx uint32 = __fill_ht.nNumUsed
				zend.ZEND_ASSERT((zend.HT_FLAGS(__fill_ht) & zend.HASH_FLAG_PACKED) != 0)
				for ; low <= high; low += uint(lstep) {
					zend.ZEND_HASH_FILL_SET_INTERNED_STR(zend.ZSTR_CHAR(low))
					zend.ZEND_HASH_FILL_NEXT()
					if signed__int(low+lstep) > 255 {
						break
					}
				}
				__fill_ht.nNumUsed = __fill_idx
				__fill_ht.nNumOfElements = __fill_idx
				__fill_ht.nNextFreeElement = __fill_idx
				__fill_ht.nInternalPointer = 0
				break
			}
		} else {
			zend.ArrayInit(return_value)
			zend.ZVAL_INTERNED_STR(&tmp, zend.ZSTR_CHAR(low))
			zend.ZendHashNextIndexInsertNew(zend.Z_ARRVAL_P(return_value), &tmp)
		}
	} else if zend.Z_TYPE_P(zlow) == zend.IS_DOUBLE || zend.Z_TYPE_P(zhigh) == zend.IS_DOUBLE || is_step_double != 0 {
		var low float64
		var high float64
		var element float64
		var i uint32
		var size uint32
	double_str:
		low = zend.ZvalGetDouble(zlow)
		high = zend.ZvalGetDouble(zhigh)
		if core.ZendIsinf(high) || core.ZendIsinf(low) {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Invalid range supplied: start=%0.0f end=%0.0f", low, high)
			zend.RETVAL_FALSE
			return
		}
		if low > high {
			if low-high < step || step <= 0 {
				err = 1
				goto err
			}
			var __calc_size float64 = (low-high)/step + 1
			if __calc_size >= float64(zend.HT_MAX_SIZE) {
				core.PhpErrorDocref(nil, zend.E_WARNING, "The supplied range exceeds the maximum array size: start=%0.0f end=%0.0f", high, low)
				zend.RETVAL_FALSE
				return
			}
			size = uint32(_phpMathRound(__calc_size, 0, PHP_ROUND_HALF_UP))
			zend.ArrayInitSize(return_value, size)
			zend.ZendHashRealInitPacked(zend.Z_ARRVAL_P(return_value))
			var __fill_ht *zend.HashTable = zend.Z_ARRVAL_P(return_value)
			var __fill_bkt *zend.Bucket = __fill_ht.arData + __fill_ht.nNumUsed
			var __fill_idx uint32 = __fill_ht.nNumUsed
			zend.ZEND_ASSERT((zend.HT_FLAGS(__fill_ht) & zend.HASH_FLAG_PACKED) != 0)
			i = 0
			element = low
			for i < size && element >= high {
				zend.ZEND_HASH_FILL_SET_DOUBLE(element)
				zend.ZEND_HASH_FILL_NEXT()
				i++
				element = low - i*step
			}
			__fill_ht.nNumUsed = __fill_idx
			__fill_ht.nNumOfElements = __fill_idx
			__fill_ht.nNextFreeElement = __fill_idx
			__fill_ht.nInternalPointer = 0
		} else if high > low {
			if high-low < step || step <= 0 {
				err = 1
				goto err
			}
			var __calc_size float64 = (high-low)/step + 1
			if __calc_size >= float64(zend.HT_MAX_SIZE) {
				core.PhpErrorDocref(nil, zend.E_WARNING, "The supplied range exceeds the maximum array size: start=%0.0f end=%0.0f", low, high)
				zend.RETVAL_FALSE
				return
			}
			size = uint32(_phpMathRound(__calc_size, 0, PHP_ROUND_HALF_UP))
			zend.ArrayInitSize(return_value, size)
			zend.ZendHashRealInitPacked(zend.Z_ARRVAL_P(return_value))
			var __fill_ht *zend.HashTable = zend.Z_ARRVAL_P(return_value)
			var __fill_bkt *zend.Bucket = __fill_ht.arData + __fill_ht.nNumUsed
			var __fill_idx uint32 = __fill_ht.nNumUsed
			zend.ZEND_ASSERT((zend.HT_FLAGS(__fill_ht) & zend.HASH_FLAG_PACKED) != 0)
			i = 0
			element = low
			for i < size && element <= high {
				zend.ZEND_HASH_FILL_SET_DOUBLE(element)
				zend.ZEND_HASH_FILL_NEXT()
				i++
				element = low + i*step
			}
			__fill_ht.nNumUsed = __fill_idx
			__fill_ht.nNumOfElements = __fill_idx
			__fill_ht.nNextFreeElement = __fill_idx
			__fill_ht.nInternalPointer = 0
		} else {
			zend.ArrayInit(return_value)
			zend.ZVAL_DOUBLE(&tmp, low)
			zend.ZendHashNextIndexInsertNew(zend.Z_ARRVAL_P(return_value), &tmp)
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
			if __calc_size >= zend.HT_MAX_SIZE-1 {
				core.PhpErrorDocref(nil, zend.E_WARNING, "The supplied range exceeds the maximum array size: start="+zend.ZEND_LONG_FMT+" end="+zend.ZEND_LONG_FMT, high, low)
				zend.RETVAL_FALSE
				return
			}
			size = uint32_t(__calc_size + 1)
			zend.ArrayInitSize(return_value, size)
			zend.ZendHashRealInitPacked(zend.Z_ARRVAL_P(return_value))
			var __fill_ht *zend.HashTable = zend.Z_ARRVAL_P(return_value)
			var __fill_bkt *zend.Bucket = __fill_ht.arData + __fill_ht.nNumUsed
			var __fill_idx uint32 = __fill_ht.nNumUsed
			zend.ZEND_ASSERT((zend.HT_FLAGS(__fill_ht) & zend.HASH_FLAG_PACKED) != 0)
			for i = 0; i < size; i++ {
				zend.ZEND_HASH_FILL_SET_LONG(low - i*lstep)
				zend.ZEND_HASH_FILL_NEXT()
			}
			__fill_ht.nNumUsed = __fill_idx
			__fill_ht.nNumOfElements = __fill_idx
			__fill_ht.nNextFreeElement = __fill_idx
			__fill_ht.nInternalPointer = 0
		} else if high > low {
			if zend.ZendUlong(high-low < lstep) != 0 {
				err = 1
				goto err
			}
			var __calc_size zend.ZendUlong = zend.ZendUlong(high-low) / lstep
			if __calc_size >= zend.HT_MAX_SIZE-1 {
				core.PhpErrorDocref(nil, zend.E_WARNING, "The supplied range exceeds the maximum array size: start="+zend.ZEND_LONG_FMT+" end="+zend.ZEND_LONG_FMT, low, high)
				zend.RETVAL_FALSE
				return
			}
			size = uint32_t(__calc_size + 1)
			zend.ArrayInitSize(return_value, size)
			zend.ZendHashRealInitPacked(zend.Z_ARRVAL_P(return_value))
			var __fill_ht *zend.HashTable = zend.Z_ARRVAL_P(return_value)
			var __fill_bkt *zend.Bucket = __fill_ht.arData + __fill_ht.nNumUsed
			var __fill_idx uint32 = __fill_ht.nNumUsed
			zend.ZEND_ASSERT((zend.HT_FLAGS(__fill_ht) & zend.HASH_FLAG_PACKED) != 0)
			for i = 0; i < size; i++ {
				zend.ZEND_HASH_FILL_SET_LONG(low + i*lstep)
				zend.ZEND_HASH_FILL_NEXT()
			}
			__fill_ht.nNumUsed = __fill_idx
			__fill_ht.nNumOfElements = __fill_idx
			__fill_ht.nNextFreeElement = __fill_idx
			__fill_ht.nInternalPointer = 0
		} else {
			zend.ArrayInit(return_value)
			zend.ZVAL_LONG(&tmp, low)
			zend.ZendHashNextIndexInsertNew(zend.Z_ARRVAL_P(return_value), &tmp)
		}
	}
err:
	if err != 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "step exceeds the specified range")
		zend.RETVAL_FALSE
		return
	}
}
func PhpArrayDataShuffle(array *zend.Zval) {
	var idx uint32
	var j uint32
	var n_elems uint32
	var p *zend.Bucket
	var temp zend.Bucket
	var hash *zend.HashTable
	var rnd_idx zend.ZendLong
	var n_left uint32
	n_elems = zend.ZendHashNumElements(zend.Z_ARRVAL_P(array))
	if n_elems < 1 {
		return
	}
	hash = zend.Z_ARRVAL_P(array)
	n_left = n_elems
	if zend.EXPECTED(!(zend.HT_HAS_ITERATORS(hash))) {
		if hash.nNumUsed != hash.nNumOfElements {
			j = 0
			idx = 0
			for ; idx < hash.nNumUsed; idx++ {
				p = hash.arData + idx
				if zend.Z_TYPE(p.val) == zend.IS_UNDEF {
					continue
				}
				if j != idx {
					hash.arData[j] = *p
				}
				j++
			}
		}
		for b.PreDec(&n_left) {
			rnd_idx = PhpMtRandRange(0, n_left)
			if rnd_idx != n_left {
				temp = hash.arData[n_left]
				hash.arData[n_left] = hash.arData[rnd_idx]
				hash.arData[rnd_idx] = temp
			}
		}
	} else {
		var iter_pos uint32 = zend.ZendHashIteratorsLowerPos(hash, 0)
		if hash.nNumUsed != hash.nNumOfElements {
			j = 0
			idx = 0
			for ; idx < hash.nNumUsed; idx++ {
				p = hash.arData + idx
				if zend.Z_TYPE(p.val) == zend.IS_UNDEF {
					continue
				}
				if j != idx {
					hash.arData[j] = *p
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
				temp = hash.arData[n_left]
				hash.arData[n_left] = hash.arData[rnd_idx]
				hash.arData[rnd_idx] = temp
				zend.ZendHashIteratorsUpdate(hash, uint32(rnd_idx), n_left)
			}
		}
	}
	hash.nNumUsed = n_elems
	hash.nInternalPointer = 0
	for j = 0; j < n_elems; j++ {
		p = hash.arData + j
		if p.key != nil {
			zend.ZendStringReleaseEx(p.key, 0)
		}
		p.h = j
		p.key = nil
	}
	hash.nNextFreeElement = n_elems
	if (zend.HT_FLAGS(hash) & zend.HASH_FLAG_PACKED) == 0 {
		zend.ZendHashToPacked(hash)
	}
}
func ZifShuffle(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var array *zend.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(1, 1)
			if zend.UNEXPECTED(zend.ZendParseArgArray(_arg, &array, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			zend.RETVAL_FALSE
			return
		}
		break
	}
	PhpArrayDataShuffle(array)
	zend.RETVAL_TRUE
	return
}
func PhpSplice(in_hash *zend.HashTable, offset zend.ZendLong, length zend.ZendLong, replace *zend.HashTable, removed *zend.HashTable) {
	var out_hash zend.HashTable
	var num_in zend.ZendLong
	var pos zend.ZendLong
	var idx uint32
	var p *zend.Bucket
	var entry *zend.Zval
	var iter_pos uint32 = zend.ZendHashIteratorsLowerPos(in_hash, 0)

	/* Get number of entries in the input hash */

	num_in = zend.ZendHashNumElements(in_hash)

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

	zend.ZendHashInit(&out_hash, b.Cond(length > 0, num_in-length, 0)+b.CondF1(replace != nil, func() __auto__ { return zend.ZendHashNumElements(replace) }, 0), nil, zend.ZVAL_PTR_DTOR, 0)

	/* Start at the beginning of the input hash and copy entries to output hash until offset is reached */

	pos = 0
	idx = 0
	for ; pos < offset && idx < in_hash.nNumUsed; idx++ {
		p = in_hash.arData + idx
		if zend.Z_TYPE(p.val) == zend.IS_UNDEF {
			continue
		}

		/* Get entry and increase reference count */

		entry = &p.val

		/* Update output hash depending on key type */

		if p.key == nil {
			zend.ZendHashNextIndexInsertNew(&out_hash, entry)
		} else {
			zend.ZendHashAddNew(&out_hash, p.key, entry)
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
		for ; pos < offset+length && idx < in_hash.nNumUsed; idx++ {
			p = in_hash.arData + idx
			if zend.Z_TYPE(p.val) == zend.IS_UNDEF {
				continue
			}
			pos++
			entry = &p.val
			zend.Z_TRY_ADDREF_P(entry)
			if p.key == nil {
				zend.ZendHashNextIndexInsertNew(removed, entry)
				zend.ZendHashDelBucket(in_hash, p)
			} else {
				zend.ZendHashAddNew(removed, p.key, entry)
				if in_hash == &(zend.ExecutorGlobals.symbol_table) {
					zend.ZendDeleteGlobalVariable(p.key)
				} else {
					zend.ZendHashDelBucket(in_hash, p)
				}
			}
		}
	} else {
		var pos2 int = pos
		for ; pos2 < offset+length && idx < in_hash.nNumUsed; idx++ {
			p = in_hash.arData + idx
			if zend.Z_TYPE(p.val) == zend.IS_UNDEF {
				continue
			}
			pos2++
			if p.key != nil && in_hash == &(zend.ExecutorGlobals.symbol_table) {
				zend.ZendDeleteGlobalVariable(p.key)
			} else {
				zend.ZendHashDelBucket(in_hash, p)
			}
		}
	}
	iter_pos = zend.ZendHashIteratorsLowerPos(in_hash, iter_pos)

	/* If there are entries to insert.. */

	if replace != nil {
		for {
			var __ht *zend.HashTable = replace
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val
				if zend.Z_TYPE_P(_z) == zend.IS_INDIRECT {
					_z = zend.Z_INDIRECT_P(_z)
				}
				if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
					continue
				}
				entry = _z
				zend.Z_TRY_ADDREF_P(entry)
				zend.ZendHashNextIndexInsertNew(&out_hash, entry)
				pos++
			}
			break
		}
	}

	/* Copy the remaining input hash entries to the output hash */

	for ; idx < in_hash.nNumUsed; idx++ {
		p = in_hash.arData + idx
		if zend.Z_TYPE(p.val) == zend.IS_UNDEF {
			continue
		}
		entry = &p.val
		if p.key == nil {
			zend.ZendHashNextIndexInsertNew(&out_hash, entry)
		} else {
			zend.ZendHashAddNew(&out_hash, p.key, entry)
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

	zend.HT_SET_ITERATORS_COUNT(&out_hash, zend.HT_ITERATORS_COUNT(in_hash))
	zend.HT_SET_ITERATORS_COUNT(in_hash, 0)
	in_hash.pDestructor = nil
	zend.ZendHashDestroy(in_hash)
	zend.HT_FLAGS(in_hash) = zend.HT_FLAGS(&out_hash)
	in_hash.nTableSize = out_hash.nTableSize
	in_hash.nTableMask = out_hash.nTableMask
	in_hash.nNumUsed = out_hash.nNumUsed
	in_hash.nNumOfElements = out_hash.nNumOfElements
	in_hash.nNextFreeElement = out_hash.nNextFreeElement
	in_hash.arData = out_hash.arData
	in_hash.pDestructor = out_hash.pDestructor
	zend.ZendHashInternalPointerReset(in_hash)
}
func ZifArrayPush(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var args *zend.Zval
	var stack *zend.Zval
	var new_var zend.Zval
	var i int
	var argc int
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = -1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(1, 1)
			if zend.UNEXPECTED(zend.ZendParseArgArray(_arg, &stack, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			var _num_varargs int = _num_args - _i - 0
			if zend.EXPECTED(_num_varargs > 0) {
				args = _real_arg + 1
				argc = _num_varargs
				_i += _num_varargs
				_real_arg += _num_varargs
			} else {
				args = nil
				argc = 0
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}

	/* For each subsequent argument, make it a reference, increase refcount, and add it to the end of the array */

	for i = 0; i < argc; i++ {
		zend.ZVAL_COPY(&new_var, &args[i])
		if zend.ZendHashNextIndexInsert(zend.Z_ARRVAL_P(stack), &new_var) == nil {
			zend.Z_TRY_DELREF(new_var)
			core.PhpErrorDocref(nil, zend.E_WARNING, "Cannot add element to the array as the next element is already occupied")
			zend.RETVAL_FALSE
			return
		}
	}

	/* Clean up and return the number of values in the stack */

	zend.RETVAL_LONG(zend.ZendHashNumElements(zend.Z_ARRVAL_P(stack)))

	/* Clean up and return the number of values in the stack */
}
func ZifArrayPop(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var stack *zend.Zval
	var val *zend.Zval
	var idx uint32
	var p *zend.Bucket
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(1, 1)
			if zend.UNEXPECTED(zend.ZendParseArgArray(_arg, &stack, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	if zend.ZendHashNumElements(zend.Z_ARRVAL_P(stack)) == 0 {
		return
	}

	/* Get the last value and copy it into the return value */

	idx = zend.Z_ARRVAL_P(stack).nNumUsed
	for true {
		if idx == 0 {
			return
		}
		idx--
		p = zend.Z_ARRVAL_P(stack).arData + idx
		val = &p.val
		if zend.Z_TYPE_P(val) == zend.IS_INDIRECT {
			val = zend.Z_INDIRECT_P(val)
		}
		if zend.Z_TYPE_P(val) != zend.IS_UNDEF {
			break
		}
	}
	zend.ZVAL_COPY_DEREF(return_value, val)
	if p.key == nil && zend.Z_ARRVAL_P(stack).nNextFreeElement > 0 && p.h >= zend_ulong(zend.Z_ARRVAL_P(stack).nNextFreeElement-1) {
		zend.Z_ARRVAL_P(stack).nNextFreeElement = zend.Z_ARRVAL_P(stack).nNextFreeElement - 1
	}

	/* Delete the last value */

	if p.key != nil && zend.Z_ARRVAL_P(stack) == &(zend.ExecutorGlobals.symbol_table) {
		zend.ZendDeleteGlobalVariable(p.key)
	} else {
		zend.ZendHashDelBucket(zend.Z_ARRVAL_P(stack), p)
	}
	zend.ZendHashInternalPointerReset(zend.Z_ARRVAL_P(stack))
}
func ZifArrayShift(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var stack *zend.Zval
	var val *zend.Zval
	var idx uint32
	var p *zend.Bucket
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(1, 1)
			if zend.UNEXPECTED(zend.ZendParseArgArray(_arg, &stack, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	if zend.ZendHashNumElements(zend.Z_ARRVAL_P(stack)) == 0 {
		return
	}

	/* Get the first value and copy it into the return value */

	idx = 0
	for true {
		if idx == zend.Z_ARRVAL_P(stack).nNumUsed {
			return
		}
		p = zend.Z_ARRVAL_P(stack).arData + idx
		val = &p.val
		if zend.Z_TYPE_P(val) == zend.IS_INDIRECT {
			val = zend.Z_INDIRECT_P(val)
		}
		if zend.Z_TYPE_P(val) != zend.IS_UNDEF {
			break
		}
		idx++
	}
	zend.ZVAL_COPY_DEREF(return_value, val)

	/* Delete the first value */

	if p.key != nil && zend.Z_ARRVAL_P(stack) == &(zend.ExecutorGlobals.symbol_table) {
		zend.ZendDeleteGlobalVariable(p.key)
	} else {
		zend.ZendHashDelBucket(zend.Z_ARRVAL_P(stack), p)
	}

	/* re-index like it did before */

	if (zend.HT_FLAGS(zend.Z_ARRVAL_P(stack)) & zend.HASH_FLAG_PACKED) != 0 {
		var k uint32 = 0
		if zend.EXPECTED(!(zend.HT_HAS_ITERATORS(zend.Z_ARRVAL_P(stack)))) {
			for idx = 0; idx < zend.Z_ARRVAL_P(stack).nNumUsed; idx++ {
				p = zend.Z_ARRVAL_P(stack).arData + idx
				if zend.Z_TYPE(p.val) == zend.IS_UNDEF {
					continue
				}
				if idx != k {
					var q *zend.Bucket = zend.Z_ARRVAL_P(stack).arData + k
					q.h = k
					q.key = nil
					zend.ZVAL_COPY_VALUE(&q.val, &p.val)
					zend.ZVAL_UNDEF(&p.val)
				}
				k++
			}
		} else {
			var iter_pos uint32 = zend.ZendHashIteratorsLowerPos(zend.Z_ARRVAL_P(stack), 0)
			for idx = 0; idx < zend.Z_ARRVAL_P(stack).nNumUsed; idx++ {
				p = zend.Z_ARRVAL_P(stack).arData + idx
				if zend.Z_TYPE(p.val) == zend.IS_UNDEF {
					continue
				}
				if idx != k {
					var q *zend.Bucket = zend.Z_ARRVAL_P(stack).arData + k
					q.h = k
					q.key = nil
					zend.ZVAL_COPY_VALUE(&q.val, &p.val)
					zend.ZVAL_UNDEF(&p.val)
					if idx == iter_pos {
						zend.ZendHashIteratorsUpdate(zend.Z_ARRVAL_P(stack), idx, k)
						iter_pos = zend.ZendHashIteratorsLowerPos(zend.Z_ARRVAL_P(stack), iter_pos+1)
					}
				}
				k++
			}
		}
		zend.Z_ARRVAL_P(stack).nNumUsed = k
		zend.Z_ARRVAL_P(stack).nNextFreeElement = k
	} else {
		var k uint32 = 0
		var should_rehash int = 0
		for idx = 0; idx < zend.Z_ARRVAL_P(stack).nNumUsed; idx++ {
			p = zend.Z_ARRVAL_P(stack).arData + idx
			if zend.Z_TYPE(p.val) == zend.IS_UNDEF {
				continue
			}
			if p.key == nil {
				if p.h != k {
					k++
					p.h = k - 1
					should_rehash = 1
				} else {
					k++
				}
			}
		}
		zend.Z_ARRVAL_P(stack).nNextFreeElement = k
		if should_rehash != 0 {
			zend.ZendHashRehash(zend.Z_ARRVAL_P(stack))
		}
	}
	zend.ZendHashInternalPointerReset(zend.Z_ARRVAL_P(stack))
}
func ZifArrayUnshift(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var args *zend.Zval
	var stack *zend.Zval
	var new_hash zend.HashTable
	var argc int
	var i int
	var key *zend.ZendString
	var value *zend.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = -1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(1, 1)
			if zend.UNEXPECTED(zend.ZendParseArgArray(_arg, &stack, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			var _num_varargs int = _num_args - _i - 0
			if zend.EXPECTED(_num_varargs > 0) {
				args = _real_arg + 1
				argc = _num_varargs
				_i += _num_varargs
				_real_arg += _num_varargs
			} else {
				args = nil
				argc = 0
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	zend.ZendHashInit(&new_hash, zend.ZendHashNumElements(zend.Z_ARRVAL_P(stack))+argc, nil, zend.ZVAL_PTR_DTOR, 0)
	for i = 0; i < argc; i++ {
		zend.Z_TRY_ADDREF(args[i])
		zend.ZendHashNextIndexInsertNew(&new_hash, &args[i])
	}
	for {
		var __ht *zend.HashTable = zend.Z_ARRVAL_P(stack)
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
				continue
			}
			key = _p.key
			value = _z
			if key != nil {
				zend.ZendHashAddNew(&new_hash, key, value)
			} else {
				zend.ZendHashNextIndexInsertNew(&new_hash, value)
			}
		}
		break
	}
	if zend.UNEXPECTED(zend.HT_HAS_ITERATORS(zend.Z_ARRVAL_P(stack))) {
		zend.ZendHashIteratorsAdvance(zend.Z_ARRVAL_P(stack), argc)
		zend.HT_SET_ITERATORS_COUNT(&new_hash, zend.HT_ITERATORS_COUNT(zend.Z_ARRVAL_P(stack)))
		zend.HT_SET_ITERATORS_COUNT(zend.Z_ARRVAL_P(stack), 0)
	}

	/* replace HashTable data */

	zend.Z_ARRVAL_P(stack).pDestructor = nil
	zend.ZendHashDestroy(zend.Z_ARRVAL_P(stack))
	zend.HT_FLAGS(zend.Z_ARRVAL_P(stack)) = zend.HT_FLAGS(&new_hash)
	zend.Z_ARRVAL_P(stack).nTableSize = new_hash.nTableSize
	zend.Z_ARRVAL_P(stack).nTableMask = new_hash.nTableMask
	zend.Z_ARRVAL_P(stack).nNumUsed = new_hash.nNumUsed
	zend.Z_ARRVAL_P(stack).nNumOfElements = new_hash.nNumOfElements
	zend.Z_ARRVAL_P(stack).nNextFreeElement = new_hash.nNextFreeElement
	zend.Z_ARRVAL_P(stack).arData = new_hash.arData
	zend.Z_ARRVAL_P(stack).pDestructor = new_hash.pDestructor
	zend.ZendHashInternalPointerReset(zend.Z_ARRVAL_P(stack))

	/* Clean up and return the number of elements in the stack */

	zend.RETVAL_LONG(zend.ZendHashNumElements(zend.Z_ARRVAL_P(stack)))

	/* Clean up and return the number of elements in the stack */
}
func ZifArraySplice(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var array *zend.Zval
	var repl_array *zend.Zval = nil
	var rem_hash *zend.HashTable = nil
	var offset zend.ZendLong
	var length zend.ZendLong = 0
	var num_in int
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 4
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(1, 1)
			if zend.UNEXPECTED(zend.ZendParseArgArray(_arg, &array, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &offset, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &length, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &repl_array, 0)
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	num_in = zend.ZendHashNumElements(zend.Z_ARRVAL_P(array))
	if zend.ZEND_NUM_ARGS() < 3 {
		length = num_in
	}
	if zend.ZEND_NUM_ARGS() == 4 {

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
		rem_hash = zend.Z_ARRVAL_P(return_value)
	}

	/* Perform splice */

	PhpSplice(zend.Z_ARRVAL_P(array), offset, length, b.CondF1(repl_array != nil, func() *zend.ZendArray { return zend.Z_ARRVAL_P(repl_array) }, nil), rem_hash)

	/* Perform splice */
}
func ZifArraySlice(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var input *zend.Zval
	var z_length *zend.Zval = nil
	var entry *zend.Zval
	var offset zend.ZendLong
	var length zend.ZendLong = 0
	var preserve_keys zend.ZendBool = 0
	var num_in int
	var pos int
	var string_key *zend.ZendString
	var num_key zend.ZendUlong
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 4
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgArray(_arg, &input, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &offset, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &z_length, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgBool(_arg, &preserve_keys, &_dummy, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_BOOL
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}

	/* Get number of entries in the input hash */

	num_in = zend.ZendHashNumElements(zend.Z_ARRVAL_P(input))

	/* We want all entries from offset to the end if length is not passed or is null */

	if zend.ZEND_NUM_ARGS() < 3 || zend.Z_TYPE_P(z_length) == zend.IS_NULL {
		length = num_in
	} else {
		length = zend.ZvalGetLong(z_length)
	}

	/* Clamp the offset.. */

	if offset > num_in {
		zend.RETVAL_EMPTY_ARRAY()
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
		zend.RETVAL_EMPTY_ARRAY()
		return
	}

	/* Initialize returned array */

	zend.ArrayInitSize(return_value, uint32(length))

	/* Start at the beginning and go until we hit offset */

	pos = 0
	if zend.HT_IS_PACKED(zend.Z_ARRVAL_P(input)) && (preserve_keys == 0 || offset == 0 && zend.HT_IS_WITHOUT_HOLES(zend.Z_ARRVAL_P(input))) {
		zend.ZendHashRealInitPacked(zend.Z_ARRVAL_P(return_value))
		for {
			var __fill_ht *zend.HashTable = zend.Z_ARRVAL_P(return_value)
			var __fill_bkt *zend.Bucket = __fill_ht.arData + __fill_ht.nNumUsed
			var __fill_idx uint32 = __fill_ht.nNumUsed
			zend.ZEND_ASSERT((zend.HT_FLAGS(__fill_ht) & zend.HASH_FLAG_PACKED) != 0)
			for {
				var __ht *zend.HashTable = zend.Z_ARRVAL_P(input)
				var _p *zend.Bucket = __ht.arData
				var _end *zend.Bucket = _p + __ht.nNumUsed
				for ; _p != _end; _p++ {
					var _z *zend.Zval = &_p.val

					if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
						continue
					}
					entry = _z
					pos++
					if pos <= offset {
						continue
					}
					if pos > offset+length {
						break
					}
					if zend.UNEXPECTED(zend.Z_ISREF_P(entry)) && zend.UNEXPECTED(zend.Z_REFCOUNT_P(entry) == 1) {
						entry = zend.Z_REFVAL_P(entry)
					}
					zend.Z_TRY_ADDREF_P(entry)
					zend.ZEND_HASH_FILL_ADD(entry)
				}
				break
			}
			__fill_ht.nNumUsed = __fill_idx
			__fill_ht.nNumOfElements = __fill_idx
			__fill_ht.nNextFreeElement = __fill_idx
			__fill_ht.nInternalPointer = 0
			break
		}
	} else {
		for {
			var __ht *zend.HashTable = zend.Z_ARRVAL_P(input)
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val

				if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
					continue
				}
				num_key = _p.h
				string_key = _p.key
				entry = _z
				pos++
				if pos <= offset {
					continue
				}
				if pos > offset+length {
					break
				}
				if string_key != nil {
					entry = zend.ZendHashAddNew(zend.Z_ARRVAL_P(return_value), string_key, entry)
				} else {
					if preserve_keys != 0 {
						entry = zend.ZendHashIndexAddNew(zend.Z_ARRVAL_P(return_value), num_key, entry)
					} else {
						entry = zend.ZendHashNextIndexInsertNew(zend.Z_ARRVAL_P(return_value), entry)
					}
				}
				zend.ZvalAddRef(entry)
			}
			break
		}
	}
}
func PhpArrayMergeRecursive(dest *zend.HashTable, src *zend.HashTable) int {
	var src_entry *zend.Zval
	var dest_entry *zend.Zval
	var string_key *zend.ZendString
	for {
		var __ht *zend.HashTable = src
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
				continue
			}
			string_key = _p.key
			src_entry = _z
			if string_key != nil {
				if b.Assign(&dest_entry, zend.ZendHashFindEx(dest, string_key, 1)) != nil {
					var src_zval *zend.Zval = src_entry
					var dest_zval *zend.Zval = dest_entry
					var thash *zend.HashTable
					var tmp zend.Zval
					var ret int
					zend.ZVAL_DEREF(src_zval)
					zend.ZVAL_DEREF(dest_zval)
					if zend.Z_TYPE_P(dest_zval) == zend.IS_ARRAY {
						thash = zend.Z_ARRVAL_P(dest_zval)
					} else {
						thash = nil
					}
					if thash != nil && zend.GC_IS_RECURSIVE(thash) != 0 || src_entry == dest_entry && zend.Z_ISREF_P(dest_entry) && zend.Z_REFCOUNT_P(dest_entry)%2 != 0 {
						core.PhpErrorDocref(nil, zend.E_WARNING, "recursion detected")
						return 0
					}
					zend.ZEND_ASSERT(!(zend.Z_ISREF_P(dest_entry)) || zend.Z_REFCOUNT_P(dest_entry) > 1)
					zend.SEPARATE_ZVAL(dest_entry)
					dest_zval = dest_entry
					if zend.Z_TYPE_P(dest_zval) == zend.IS_NULL {
						zend.ConvertToArrayEx(dest_zval)
						zend.AddNextIndexNull(dest_zval)
					} else {
						zend.ConvertToArrayEx(dest_zval)
					}
					zend.ZVAL_UNDEF(&tmp)
					if zend.Z_TYPE_P(src_zval) == zend.IS_OBJECT {
						zend.ZVAL_COPY(&tmp, src_zval)
						zend.ConvertToArray(&tmp)
						src_zval = &tmp
					}
					if zend.Z_TYPE_P(src_zval) == zend.IS_ARRAY {
						if thash != nil && (zend.GC_FLAGS(thash)&zend.GC_IMMUTABLE) == 0 {
							zend.GC_PROTECT_RECURSION(thash)
						}
						ret = PhpArrayMergeRecursive(zend.Z_ARRVAL_P(dest_zval), zend.Z_ARRVAL_P(src_zval))
						if thash != nil && (zend.GC_FLAGS(thash)&zend.GC_IMMUTABLE) == 0 {
							zend.GC_UNPROTECT_RECURSION(thash)
						}
						if ret == 0 {
							return 0
						}
					} else {
						zend.Z_TRY_ADDREF_P(src_zval)
						zend.ZendHashNextIndexInsert(zend.Z_ARRVAL_P(dest_zval), src_zval)
					}
					zend.ZvalPtrDtor(&tmp)
				} else {
					var zv *zend.Zval = zend.ZendHashAddNew(dest, string_key, src_entry)
					zend.ZvalAddRef(zv)
				}
			} else {
				var zv *zend.Zval = zend.ZendHashNextIndexInsert(dest, src_entry)
				zend.ZvalAddRef(zv)
			}
		}
		break
	}
	return 1
}
func PhpArrayMerge(dest *zend.HashTable, src *zend.HashTable) int {
	var src_entry *zend.Zval
	var string_key *zend.ZendString
	if (zend.HT_FLAGS(dest)&zend.HASH_FLAG_PACKED) != 0 && (zend.HT_FLAGS(src)&zend.HASH_FLAG_PACKED) != 0 {
		zend.ZendHashExtend(dest, zend.ZendHashNumElements(dest)+zend.ZendHashNumElements(src), 1)
		for {
			var __fill_ht *zend.HashTable = dest
			var __fill_bkt *zend.Bucket = __fill_ht.arData + __fill_ht.nNumUsed
			var __fill_idx uint32 = __fill_ht.nNumUsed
			zend.ZEND_ASSERT((zend.HT_FLAGS(__fill_ht) & zend.HASH_FLAG_PACKED) != 0)
			for {
				var __ht *zend.HashTable = src
				var _p *zend.Bucket = __ht.arData
				var _end *zend.Bucket = _p + __ht.nNumUsed
				for ; _p != _end; _p++ {
					var _z *zend.Zval = &_p.val

					if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
						continue
					}
					src_entry = _z
					if zend.UNEXPECTED(zend.Z_ISREF_P(src_entry)) && zend.UNEXPECTED(zend.Z_REFCOUNT_P(src_entry) == 1) {
						src_entry = zend.Z_REFVAL_P(src_entry)
					}
					zend.Z_TRY_ADDREF_P(src_entry)
					zend.ZEND_HASH_FILL_ADD(src_entry)
				}
				break
			}
			__fill_ht.nNumUsed = __fill_idx
			__fill_ht.nNumOfElements = __fill_idx
			__fill_ht.nNextFreeElement = __fill_idx
			__fill_ht.nInternalPointer = 0
			break
		}
	} else {
		for {
			var __ht *zend.HashTable = src
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val

				if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
					continue
				}
				string_key = _p.key
				src_entry = _z
				if zend.UNEXPECTED(zend.Z_ISREF_P(src_entry) && zend.Z_REFCOUNT_P(src_entry) == 1) {
					src_entry = zend.Z_REFVAL_P(src_entry)
				}
				zend.Z_TRY_ADDREF_P(src_entry)
				if string_key != nil {
					zend.ZendHashUpdate(dest, string_key, src_entry)
				} else {
					zend.ZendHashNextIndexInsertNew(dest, src_entry)
				}
			}
			break
		}
	}
	return 1
}
func PhpArrayReplaceRecursive(dest *zend.HashTable, src *zend.HashTable) int {
	var src_entry *zend.Zval
	var dest_entry *zend.Zval
	var src_zval *zend.Zval
	var dest_zval *zend.Zval
	var string_key *zend.ZendString
	var num_key zend.ZendUlong
	var ret int
	for {
		var __ht *zend.HashTable = src
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
				continue
			}
			num_key = _p.h
			string_key = _p.key
			src_entry = _z
			src_zval = src_entry
			zend.ZVAL_DEREF(src_zval)
			if string_key != nil {
				if zend.Z_TYPE_P(src_zval) != zend.IS_ARRAY || b.Assign(&dest_entry, zend.ZendHashFindEx(dest, string_key, 1)) == nil || zend.Z_TYPE_P(dest_entry) != zend.IS_ARRAY && (!(zend.Z_ISREF_P(dest_entry)) || zend.Z_TYPE_P(zend.Z_REFVAL_P(dest_entry)) != zend.IS_ARRAY) {
					var zv *zend.Zval = zend.ZendHashUpdate(dest, string_key, src_entry)
					zend.ZvalAddRef(zv)
					continue
				}
			} else {
				if zend.Z_TYPE_P(src_zval) != zend.IS_ARRAY || b.Assign(&dest_entry, zend.ZendHashIndexFind(dest, num_key)) == nil || zend.Z_TYPE_P(dest_entry) != zend.IS_ARRAY && (!(zend.Z_ISREF_P(dest_entry)) || zend.Z_TYPE_P(zend.Z_REFVAL_P(dest_entry)) != zend.IS_ARRAY) {
					var zv *zend.Zval = zend.ZendHashIndexUpdate(dest, num_key, src_entry)
					zend.ZvalAddRef(zv)
					continue
				}
			}
			dest_zval = dest_entry
			zend.ZVAL_DEREF(dest_zval)
			if zend.Z_IS_RECURSIVE_P(dest_zval) != 0 || zend.Z_IS_RECURSIVE_P(src_zval) != 0 || zend.Z_ISREF_P(src_entry) && zend.Z_ISREF_P(dest_entry) && zend.Z_REF_P(src_entry) == zend.Z_REF_P(dest_entry) && zend.Z_REFCOUNT_P(dest_entry)%2 != 0 {
				core.PhpErrorDocref(nil, zend.E_WARNING, "recursion detected")
				return 0
			}
			zend.ZEND_ASSERT(!(zend.Z_ISREF_P(dest_entry)) || zend.Z_REFCOUNT_P(dest_entry) > 1)
			zend.SEPARATE_ZVAL(dest_entry)
			dest_zval = dest_entry
			if zend.Z_REFCOUNTED_P(dest_zval) {
				zend.Z_PROTECT_RECURSION_P(dest_zval)
			}
			if zend.Z_REFCOUNTED_P(src_zval) {
				zend.Z_PROTECT_RECURSION_P(src_zval)
			}
			ret = PhpArrayReplaceRecursive(zend.Z_ARRVAL_P(dest_zval), zend.Z_ARRVAL_P(src_zval))
			if zend.Z_REFCOUNTED_P(dest_zval) {
				zend.Z_UNPROTECT_RECURSION_P(dest_zval)
			}
			if zend.Z_REFCOUNTED_P(src_zval) {
				zend.Z_UNPROTECT_RECURSION_P(src_zval)
			}
			if ret == 0 {
				return 0
			}
		}
		break
	}
	return 1
}
func PhpArrayReplaceWrapper(execute_data *zend.ZendExecuteData, return_value *zend.Zval, recursive int) {
	var args *zend.Zval = nil
	var arg *zend.Zval
	var argc int
	var i int
	var dest *zend.HashTable
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = -1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			var _num_varargs int = _num_args - _i - 0
			if zend.EXPECTED(_num_varargs > 0) {
				args = _real_arg + 1
				argc = _num_varargs
				_i += _num_varargs
				_real_arg += _num_varargs
			} else {
				args = nil
				argc = 0
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	for i = 0; i < argc; i++ {
		var arg *zend.Zval = args + i
		if zend.Z_TYPE_P(arg) != zend.IS_ARRAY {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Expected parameter %d to be an array, %s given", i+1, zend.ZendZvalTypeName(arg))
			zend.RETVAL_NULL()
			return
		}
	}

	/* copy first array */

	arg = args
	dest = zend.ZendArrayDup(zend.Z_ARRVAL_P(arg))
	zend.ZVAL_ARR(return_value, dest)
	if recursive != 0 {
		for i = 1; i < argc; i++ {
			arg = args + i
			PhpArrayReplaceRecursive(dest, zend.Z_ARRVAL_P(arg))
		}
	} else {
		for i = 1; i < argc; i++ {
			arg = args + i
			zend.ZendHashMerge(dest, zend.Z_ARRVAL_P(arg), zend.ZvalAddRef, 1)
		}
	}
}
func PhpArrayMergeWrapper(execute_data *zend.ZendExecuteData, return_value *zend.Zval, recursive int) {
	var args *zend.Zval = nil
	var arg *zend.Zval
	var argc int
	var i int
	var src_entry *zend.Zval
	var src *zend.HashTable
	var dest *zend.HashTable
	var count uint32 = 0
	for {
		var _flags int = 0
		var _min_num_args int = 0
		var _max_num_args int = -1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			var _num_varargs int = _num_args - _i - 0
			if zend.EXPECTED(_num_varargs > 0) {
				args = _real_arg + 1
				argc = _num_varargs
				_i += _num_varargs
				_real_arg += _num_varargs
			} else {
				args = nil
				argc = 0
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	if argc == 0 {
		zend.RETVAL_EMPTY_ARRAY()
		return
	}
	for i = 0; i < argc; i++ {
		var arg *zend.Zval = args + i
		if zend.Z_TYPE_P(arg) != zend.IS_ARRAY {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Expected parameter %d to be an array, %s given", i+1, zend.ZendZvalTypeName(arg))
			zend.RETVAL_NULL()
			return
		}
		count += zend.ZendHashNumElements(zend.Z_ARRVAL_P(arg))
	}
	if argc == 2 {
		var ret *zend.Zval = nil
		if zend.ZendHashNumElements(zend.Z_ARRVAL(args[0])) == 0 {
			ret = &args[1]
		} else if zend.ZendHashNumElements(zend.Z_ARRVAL(args[1])) == 0 {
			ret = &args[0]
		}
		if ret != nil {
			if (zend.HT_FLAGS(zend.Z_ARRVAL_P(ret)) & zend.HASH_FLAG_PACKED) != 0 {
				if zend.HT_IS_WITHOUT_HOLES(zend.Z_ARRVAL_P(ret)) {
					zend.ZVAL_COPY(return_value, ret)
					return
				}
			} else {
				var copy zend.ZendBool = 1
				var string_key *zend.ZendString
				for {
					var __ht *zend.HashTable = zend.Z_ARRVAL_P(ret)
					var _p *zend.Bucket = __ht.arData
					var _end *zend.Bucket = _p + __ht.nNumUsed
					for ; _p != _end; _p++ {
						var _z *zend.Zval = &_p.val

						if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
							continue
						}
						string_key = _p.key
						if string_key == nil {
							copy = 0
							break
						}
					}
					break
				}
				if copy != 0 {
					zend.ZVAL_COPY(return_value, ret)
					return
				}
			}
		}
	}
	arg = args
	src = zend.Z_ARRVAL_P(arg)

	/* copy first array */

	zend.ArrayInitSize(return_value, count)
	dest = zend.Z_ARRVAL_P(return_value)
	if (zend.HT_FLAGS(src) & zend.HASH_FLAG_PACKED) != 0 {
		zend.ZendHashRealInitPacked(dest)
		for {
			var __fill_ht *zend.HashTable = dest
			var __fill_bkt *zend.Bucket = __fill_ht.arData + __fill_ht.nNumUsed
			var __fill_idx uint32 = __fill_ht.nNumUsed
			zend.ZEND_ASSERT((zend.HT_FLAGS(__fill_ht) & zend.HASH_FLAG_PACKED) != 0)
			for {
				var __ht *zend.HashTable = src
				var _p *zend.Bucket = __ht.arData
				var _end *zend.Bucket = _p + __ht.nNumUsed
				for ; _p != _end; _p++ {
					var _z *zend.Zval = &_p.val

					if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
						continue
					}
					src_entry = _z
					if zend.UNEXPECTED(zend.Z_ISREF_P(src_entry) && zend.Z_REFCOUNT_P(src_entry) == 1) {
						src_entry = zend.Z_REFVAL_P(src_entry)
					}
					zend.Z_TRY_ADDREF_P(src_entry)
					zend.ZEND_HASH_FILL_ADD(src_entry)
				}
				break
			}
			__fill_ht.nNumUsed = __fill_idx
			__fill_ht.nNumOfElements = __fill_idx
			__fill_ht.nNextFreeElement = __fill_idx
			__fill_ht.nInternalPointer = 0
			break
		}
	} else {
		var string_key *zend.ZendString
		zend.ZendHashRealInitMixed(dest)
		for {
			var __ht *zend.HashTable = src
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val

				if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
					continue
				}
				string_key = _p.key
				src_entry = _z
				if zend.UNEXPECTED(zend.Z_ISREF_P(src_entry) && zend.Z_REFCOUNT_P(src_entry) == 1) {
					src_entry = zend.Z_REFVAL_P(src_entry)
				}
				zend.Z_TRY_ADDREF_P(src_entry)
				if zend.EXPECTED(string_key != nil) {
					zend._zendHashAppend(dest, string_key, src_entry)
				} else {
					zend.ZendHashNextIndexInsertNew(dest, src_entry)
				}
			}
			break
		}
	}
	if recursive != 0 {
		for i = 1; i < argc; i++ {
			arg = args + i
			PhpArrayMergeRecursive(dest, zend.Z_ARRVAL_P(arg))
		}
	} else {
		for i = 1; i < argc; i++ {
			arg = args + i
			PhpArrayMerge(dest, zend.Z_ARRVAL_P(arg))
		}
	}
}
func ZifArrayMerge(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpArrayMergeWrapper(execute_data, return_value, 0)
}
func ZifArrayMergeRecursive(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpArrayMergeWrapper(execute_data, return_value, 1)
}
func ZifArrayReplace(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpArrayReplaceWrapper(execute_data, return_value, 0)
}
func ZifArrayReplaceRecursive(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpArrayReplaceWrapper(execute_data, return_value, 1)
}
func ZifArrayKeys(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var input *zend.Zval
	var search_value *zend.Zval = nil
	var entry *zend.Zval
	var new_val zend.Zval
	var strict zend.ZendBool = 0
	var num_idx zend.ZendUlong
	var str_idx *zend.ZendString
	var arrval *zend.ZendArray
	var elem_count zend.ZendUlong
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 3
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgArray(_arg, &input, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &search_value, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgBool(_arg, &strict, &_dummy, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_BOOL
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	arrval = zend.Z_ARRVAL_P(input)
	elem_count = zend.ZendHashNumElements(arrval)

	/* Base case: empty input */

	if elem_count == 0 {
		zend.RETVAL_ZVAL(input, 1, 0)
		return
	}

	/* Initialize return array */

	if search_value != nil {
		zend.ArrayInit(return_value)
		if strict != 0 {
			for {
				var __ht *zend.HashTable = arrval
				var _p *zend.Bucket = __ht.arData
				var _end *zend.Bucket = _p + __ht.nNumUsed
				for ; _p != _end; _p++ {
					var _z *zend.Zval = &_p.val
					if zend.Z_TYPE_P(_z) == zend.IS_INDIRECT {
						_z = zend.Z_INDIRECT_P(_z)
					}
					if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
						continue
					}
					num_idx = _p.h
					str_idx = _p.key
					entry = _z
					zend.ZVAL_DEREF(entry)
					if zend.FastIsIdenticalFunction(search_value, entry) != 0 {
						if str_idx != nil {
							zend.ZVAL_STR_COPY(&new_val, str_idx)
						} else {
							zend.ZVAL_LONG(&new_val, num_idx)
						}
						zend.ZendHashNextIndexInsertNew(zend.Z_ARRVAL_P(return_value), &new_val)
					}
				}
				break
			}
		} else {
			for {
				var __ht *zend.HashTable = arrval
				var _p *zend.Bucket = __ht.arData
				var _end *zend.Bucket = _p + __ht.nNumUsed
				for ; _p != _end; _p++ {
					var _z *zend.Zval = &_p.val
					if zend.Z_TYPE_P(_z) == zend.IS_INDIRECT {
						_z = zend.Z_INDIRECT_P(_z)
					}
					if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
						continue
					}
					num_idx = _p.h
					str_idx = _p.key
					entry = _z
					if zend.FastEqualCheckFunction(search_value, entry) != 0 {
						if str_idx != nil {
							zend.ZVAL_STR_COPY(&new_val, str_idx)
						} else {
							zend.ZVAL_LONG(&new_val, num_idx)
						}
						zend.ZendHashNextIndexInsertNew(zend.Z_ARRVAL_P(return_value), &new_val)
					}
				}
				break
			}
		}
	} else {
		zend.ArrayInitSize(return_value, elem_count)
		zend.ZendHashRealInitPacked(zend.Z_ARRVAL_P(return_value))
		for {
			var __fill_ht *zend.HashTable = zend.Z_ARRVAL_P(return_value)
			var __fill_bkt *zend.Bucket = __fill_ht.arData + __fill_ht.nNumUsed
			var __fill_idx uint32 = __fill_ht.nNumUsed
			zend.ZEND_ASSERT((zend.HT_FLAGS(__fill_ht) & zend.HASH_FLAG_PACKED) != 0)
			if zend.HT_IS_PACKED(arrval) && zend.HT_IS_WITHOUT_HOLES(arrval) {

				/* Optimistic case: range(0..n-1) for vector-like packed array */

				var lval zend.ZendUlong = 0
				for ; lval < elem_count; lval++ {
					zend.ZEND_HASH_FILL_SET_LONG(lval)
					zend.ZEND_HASH_FILL_NEXT()
				}
			} else {

				/* Go through input array and add keys to the return array */

				for {
					var __ht *zend.HashTable = zend.Z_ARRVAL_P(input)
					var _p *zend.Bucket = __ht.arData
					var _end *zend.Bucket = _p + __ht.nNumUsed
					for ; _p != _end; _p++ {
						var _z *zend.Zval = &_p.val
						if zend.Z_TYPE_P(_z) == zend.IS_INDIRECT {
							_z = zend.Z_INDIRECT_P(_z)
						}
						if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
							continue
						}
						num_idx = _p.h
						str_idx = _p.key
						entry = _z
						if str_idx != nil {
							zend.ZEND_HASH_FILL_SET_STR_COPY(str_idx)
						} else {
							zend.ZEND_HASH_FILL_SET_LONG(num_idx)
						}
						zend.ZEND_HASH_FILL_NEXT()
					}
					break
				}

				/* Go through input array and add keys to the return array */

			}
			__fill_ht.nNumUsed = __fill_idx
			__fill_ht.nNumOfElements = __fill_idx
			__fill_ht.nNextFreeElement = __fill_idx
			__fill_ht.nInternalPointer = 0
			break
		}
	}

	/* Initialize return array */
}
func ZifArrayKeyFirst(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var stack *zend.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgArray(_arg, &stack, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	var target_hash *zend.HashTable = zend.Z_ARRVAL_P(stack)
	var pos zend.HashPosition = 0
	zend.ZendHashGetCurrentKeyZvalEx(target_hash, return_value, &pos)
}
func ZifArrayKeyLast(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var stack *zend.Zval
	var pos zend.HashPosition
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgArray(_arg, &stack, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	var target_hash *zend.HashTable = zend.Z_ARRVAL_P(stack)
	zend.ZendHashInternalPointerEndEx(target_hash, &pos)
	zend.ZendHashGetCurrentKeyZvalEx(target_hash, return_value, &pos)
}
func ZifArrayValues(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var input *zend.Zval
	var entry *zend.Zval
	var arrval *zend.ZendArray
	var arrlen zend.ZendLong
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgArray(_arg, &input, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	arrval = zend.Z_ARRVAL_P(input)

	/* Return empty input as is */

	arrlen = zend.ZendHashNumElements(arrval)
	if arrlen == 0 {
		zend.RETVAL_EMPTY_ARRAY()
		return
	}

	/* Return vector-like packed arrays as-is */

	if zend.HT_IS_PACKED(arrval) && zend.HT_IS_WITHOUT_HOLES(arrval) && arrval.nNextFreeElement == arrlen {
		zend.RETVAL_ZVAL(input, 1, 0)
		return
	}

	/* Initialize return array */

	zend.ArrayInitSize(return_value, zend.ZendHashNumElements(arrval))
	zend.ZendHashRealInitPacked(zend.Z_ARRVAL_P(return_value))

	/* Go through input array and add values to the return array */

	for {
		var __fill_ht *zend.HashTable = zend.Z_ARRVAL_P(return_value)
		var __fill_bkt *zend.Bucket = __fill_ht.arData + __fill_ht.nNumUsed
		var __fill_idx uint32 = __fill_ht.nNumUsed
		zend.ZEND_ASSERT((zend.HT_FLAGS(__fill_ht) & zend.HASH_FLAG_PACKED) != 0)
		for {
			var __ht *zend.HashTable = arrval
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val

				if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
					continue
				}
				entry = _z
				if zend.UNEXPECTED(zend.Z_ISREF_P(entry) && zend.Z_REFCOUNT_P(entry) == 1) {
					entry = zend.Z_REFVAL_P(entry)
				}
				zend.Z_TRY_ADDREF_P(entry)
				zend.ZEND_HASH_FILL_ADD(entry)
			}
			break
		}
		__fill_ht.nNumUsed = __fill_idx
		__fill_ht.nNumOfElements = __fill_idx
		__fill_ht.nNextFreeElement = __fill_idx
		__fill_ht.nInternalPointer = 0
		break
	}

	/* Go through input array and add values to the return array */
}
func ZifArrayCountValues(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var input *zend.Zval
	var entry *zend.Zval
	var tmp *zend.Zval
	var myht *zend.HashTable
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgArray(_arg, &input, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}

	/* Initialize return array */

	zend.ArrayInit(return_value)

	/* Go through input array and add values to the return array */

	myht = zend.Z_ARRVAL_P(input)
	for {
		var __ht *zend.HashTable = myht
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
				continue
			}
			entry = _z
			zend.ZVAL_DEREF(entry)
			if zend.Z_TYPE_P(entry) == zend.IS_LONG {
				if b.Assign(&tmp, zend.ZendHashIndexFind(zend.Z_ARRVAL_P(return_value), zend.Z_LVAL_P(entry))) == nil {
					var data zend.Zval
					zend.ZVAL_LONG(&data, 1)
					zend.ZendHashIndexUpdate(zend.Z_ARRVAL_P(return_value), zend.Z_LVAL_P(entry), &data)
				} else {
					zend.Z_LVAL_P(tmp)++
				}
			} else if zend.Z_TYPE_P(entry) == zend.IS_STRING {
				if b.Assign(&tmp, zend.ZendSymtableFind(zend.Z_ARRVAL_P(return_value), zend.Z_STR_P(entry))) == nil {
					var data zend.Zval
					zend.ZVAL_LONG(&data, 1)
					zend.ZendSymtableUpdate(zend.Z_ARRVAL_P(return_value), zend.Z_STR_P(entry), &data)
				} else {
					zend.Z_LVAL_P(tmp)++
				}
			} else {
				core.PhpErrorDocref(nil, zend.E_WARNING, "Can only count STRING and INTEGER values!")
			}
		}
		break
	}
}
func ArrayColumnParamHelper(param *zend.Zval, name string) zend.ZendBool {
	switch zend.Z_TYPE_P(param) {
	case zend.IS_DOUBLE:
		if zend.Z_TYPE_P(param) != zend.IS_LONG {
			zend.ConvertToLong(param)
		}
	case zend.IS_LONG:
		return 1
	case zend.IS_OBJECT:
		if zend.TryConvertToString(param) == 0 {
			return 0
		}
	case zend.IS_STRING:
		return 1
	default:
		core.PhpErrorDocref(nil, zend.E_WARNING, "The %s key should be either a string or an integer", name)
		return 0
	}
}
func ArrayColumnFetchProp(data *zend.Zval, name *zend.Zval, rv *zend.Zval) *zend.Zval {
	var prop *zend.Zval = nil
	if zend.Z_TYPE_P(data) == zend.IS_OBJECT {

		/* The has_property check is first performed in "exists" mode (which returns true for
		 * properties that are null but exist) and then in "has" mode to handle objects that
		 * implement __isset (which is not called in "exists" mode). */

		if zend.Z_OBJ_HANDLER_P(data, has_property)(data, name, zend.ZEND_PROPERTY_EXISTS, nil) || zend.Z_OBJ_HANDLER_P(data, has_property)(data, name, zend.ZEND_PROPERTY_ISSET, nil) {
			prop = zend.Z_OBJ_HANDLER_P(data, read_property)(data, name, zend.BP_VAR_R, nil, rv)
			if prop != nil {
				zend.ZVAL_DEREF(prop)
				if prop != rv {
					zend.Z_TRY_ADDREF_P(prop)
				}
			}
		}

		/* The has_property check is first performed in "exists" mode (which returns true for
		 * properties that are null but exist) and then in "has" mode to handle objects that
		 * implement __isset (which is not called in "exists" mode). */

	} else if zend.Z_TYPE_P(data) == zend.IS_ARRAY {
		if zend.Z_TYPE_P(name) == zend.IS_STRING {
			prop = zend.ZendSymtableFind(zend.Z_ARRVAL_P(data), zend.Z_STR_P(name))
		} else if zend.Z_TYPE_P(name) == zend.IS_LONG {
			prop = zend.ZendHashIndexFind(zend.Z_ARRVAL_P(data), zend.Z_LVAL_P(name))
		}
		if prop != nil {
			zend.ZVAL_DEREF(prop)
			zend.Z_TRY_ADDREF_P(prop)
		}
	}
	return prop
}
func ZifArrayColumn(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var input *zend.HashTable
	var colval *zend.Zval
	var data *zend.Zval
	var rv zend.Zval
	var column *zend.Zval = nil
	var index *zend.Zval = nil
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 3
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgArrayHt(_arg, &input, 0, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &column, 1)
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &index, 1)
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	if column != nil && ArrayColumnParamHelper(column, "column") == 0 || index != nil && ArrayColumnParamHelper(index, "index") == 0 {
		zend.RETVAL_FALSE
		return
	}
	zend.ArrayInitSize(return_value, zend.ZendHashNumElements(input))
	if index == nil {
		zend.ZendHashRealInitPacked(zend.Z_ARRVAL_P(return_value))
		for {
			var __fill_ht *zend.HashTable = zend.Z_ARRVAL_P(return_value)
			var __fill_bkt *zend.Bucket = __fill_ht.arData + __fill_ht.nNumUsed
			var __fill_idx uint32 = __fill_ht.nNumUsed
			zend.ZEND_ASSERT((zend.HT_FLAGS(__fill_ht) & zend.HASH_FLAG_PACKED) != 0)
			for {
				var __ht *zend.HashTable = input
				var _p *zend.Bucket = __ht.arData
				var _end *zend.Bucket = _p + __ht.nNumUsed
				for ; _p != _end; _p++ {
					var _z *zend.Zval = &_p.val

					if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
						continue
					}
					data = _z
					zend.ZVAL_DEREF(data)
					if column == nil {
						zend.Z_TRY_ADDREF_P(data)
						colval = data
					} else if b.Assign(&colval, ArrayColumnFetchProp(data, column, &rv)) == nil {
						continue
					}
					zend.ZEND_HASH_FILL_ADD(colval)
				}
				break
			}
			__fill_ht.nNumUsed = __fill_idx
			__fill_ht.nNumOfElements = __fill_idx
			__fill_ht.nNextFreeElement = __fill_idx
			__fill_ht.nInternalPointer = 0
			break
		}
	} else {
		for {
			var __ht *zend.HashTable = input
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val

				if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
					continue
				}
				data = _z
				zend.ZVAL_DEREF(data)
				if column == nil {
					zend.Z_TRY_ADDREF_P(data)
					colval = data
				} else if b.Assign(&colval, ArrayColumnFetchProp(data, column, &rv)) == nil {
					continue
				}

				/* Failure will leave keyval alone which will land us on the final else block below
				 * which is to append the value as next_index
				 */

				if index != nil {
					var rv zend.Zval
					var keyval *zend.Zval = ArrayColumnFetchProp(data, index, &rv)
					if keyval != nil {
						switch zend.Z_TYPE_P(keyval) {
						case zend.IS_STRING:
							zend.ZendSymtableUpdate(zend.Z_ARRVAL_P(return_value), zend.Z_STR_P(keyval), colval)
							break
						case zend.IS_LONG:
							zend.ZendHashIndexUpdate(zend.Z_ARRVAL_P(return_value), zend.Z_LVAL_P(keyval), colval)
							break
						case zend.IS_OBJECT:
							var tmp_key *zend.ZendString
							var key *zend.ZendString = zend.ZvalGetTmpString(keyval, &tmp_key)
							zend.ZendSymtableUpdate(zend.Z_ARRVAL_P(return_value), key, colval)
							zend.ZendTmpStringRelease(tmp_key)
							break
						case zend.IS_NULL:
							zend.ZendHashUpdate(zend.Z_ARRVAL_P(return_value), zend.ZSTR_EMPTY_ALLOC(), colval)
							break
						case zend.IS_DOUBLE:
							zend.ZendHashIndexUpdate(zend.Z_ARRVAL_P(return_value), zend.ZendDvalToLval(zend.Z_DVAL_P(keyval)), colval)
							break
						case zend.IS_TRUE:
							zend.ZendHashIndexUpdate(zend.Z_ARRVAL_P(return_value), 1, colval)
							break
						case zend.IS_FALSE:
							zend.ZendHashIndexUpdate(zend.Z_ARRVAL_P(return_value), 0, colval)
							break
						case zend.IS_RESOURCE:
							zend.ZendHashIndexUpdate(zend.Z_ARRVAL_P(return_value), zend.Z_RES_HANDLE_P(keyval), colval)
							break
						default:
							zend.ZendHashNextIndexInsert(zend.Z_ARRVAL_P(return_value), colval)
							break
						}
						zend.ZvalPtrDtor(keyval)
					} else {
						zend.ZendHashNextIndexInsert(zend.Z_ARRVAL_P(return_value), colval)
					}
				} else {
					zend.ZendHashNextIndexInsert(zend.Z_ARRVAL_P(return_value), colval)
				}

				/* Failure will leave keyval alone which will land us on the final else block below
				 * which is to append the value as next_index
				 */

			}
			break
		}
	}
}
func ZifArrayReverse(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var input *zend.Zval
	var entry *zend.Zval
	var string_key *zend.ZendString
	var num_key zend.ZendUlong
	var preserve_keys zend.ZendBool = 0
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgArray(_arg, &input, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgBool(_arg, &preserve_keys, &_dummy, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_BOOL
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}

	/* Initialize return array */

	zend.ArrayInitSize(return_value, zend.ZendHashNumElements(zend.Z_ARRVAL_P(input)))
	if (zend.HT_FLAGS(zend.Z_ARRVAL_P(input))&zend.HASH_FLAG_PACKED) != 0 && preserve_keys == 0 {
		zend.ZendHashRealInitPacked(zend.Z_ARRVAL_P(return_value))
		for {
			var __fill_ht *zend.HashTable = zend.Z_ARRVAL_P(return_value)
			var __fill_bkt *zend.Bucket = __fill_ht.arData + __fill_ht.nNumUsed
			var __fill_idx uint32 = __fill_ht.nNumUsed
			zend.ZEND_ASSERT((zend.HT_FLAGS(__fill_ht) & zend.HASH_FLAG_PACKED) != 0)
			for {
				var __ht *zend.HashTable = zend.Z_ARRVAL_P(input)
				var _idx uint32 = __ht.nNumUsed
				var _p *zend.Bucket = __ht.arData + _idx
				var _z *zend.Zval
				for _idx = __ht.nNumUsed; _idx > 0; _idx-- {
					_p--
					_z = &_p.val

					if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
						continue
					}
					entry = _z
					if zend.UNEXPECTED(zend.Z_ISREF_P(entry) && zend.Z_REFCOUNT_P(entry) == 1) {
						entry = zend.Z_REFVAL_P(entry)
					}
					zend.Z_TRY_ADDREF_P(entry)
					zend.ZEND_HASH_FILL_ADD(entry)
				}
				break
			}
			__fill_ht.nNumUsed = __fill_idx
			__fill_ht.nNumOfElements = __fill_idx
			__fill_ht.nNextFreeElement = __fill_idx
			__fill_ht.nInternalPointer = 0
			break
		}
	} else {
		for {
			var __ht *zend.HashTable = zend.Z_ARRVAL_P(input)
			var _idx uint32 = __ht.nNumUsed
			var _p *zend.Bucket = __ht.arData + _idx
			var _z *zend.Zval
			for _idx = __ht.nNumUsed; _idx > 0; _idx-- {
				_p--
				_z = &_p.val

				if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
					continue
				}
				num_key = _p.h
				string_key = _p.key
				entry = _z
				if string_key != nil {
					entry = zend.ZendHashAddNew(zend.Z_ARRVAL_P(return_value), string_key, entry)
				} else {
					if preserve_keys != 0 {
						entry = zend.ZendHashIndexAddNew(zend.Z_ARRVAL_P(return_value), num_key, entry)
					} else {
						entry = zend.ZendHashNextIndexInsertNew(zend.Z_ARRVAL_P(return_value), entry)
					}
				}
				zend.ZvalAddRef(entry)
			}
			break
		}
	}
}
func ZifArrayPad(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var input *zend.Zval
	var pad_value *zend.Zval
	var pad_size zend.ZendLong
	var pad_size_abs zend.ZendLong
	var input_size zend.ZendLong
	var num_pads zend.ZendLong
	var i zend.ZendLong
	var key *zend.ZendString
	var value *zend.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 3
		var _max_num_args int = 3
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgArray(_arg, &input, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &pad_size, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &pad_value, 0)
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}

	/* Do some initial calculations */

	input_size = zend.ZendHashNumElements(zend.Z_ARRVAL_P(input))
	pad_size_abs = zend.ZEND_ABS(pad_size)
	if pad_size_abs < 0 || pad_size_abs-input_size > int64(1048576) {
		core.PhpErrorDocref(nil, zend.E_WARNING, "You may only pad up to 1048576 elements at a time")
		zend.RETVAL_FALSE
		return
	}
	if input_size >= pad_size_abs {

		/* Copy the original array */

		zend.ZVAL_COPY(return_value, input)
		return
	}
	num_pads = pad_size_abs - input_size
	if zend.Z_REFCOUNTED_P(pad_value) {
		zend.GC_ADDREF_EX(zend.Z_COUNTED_P(pad_value), num_pads)
	}
	zend.ArrayInitSize(return_value, pad_size_abs)
	if (zend.HT_FLAGS(zend.Z_ARRVAL_P(input)) & zend.HASH_FLAG_PACKED) != 0 {
		zend.ZendHashRealInitPacked(zend.Z_ARRVAL_P(return_value))
		if pad_size < 0 {
			var __fill_ht *zend.HashTable = zend.Z_ARRVAL_P(return_value)
			var __fill_bkt *zend.Bucket = __fill_ht.arData + __fill_ht.nNumUsed
			var __fill_idx uint32 = __fill_ht.nNumUsed
			zend.ZEND_ASSERT((zend.HT_FLAGS(__fill_ht) & zend.HASH_FLAG_PACKED) != 0)
			for i = 0; i < num_pads; i++ {
				zend.ZEND_HASH_FILL_ADD(pad_value)
			}
			__fill_ht.nNumUsed = __fill_idx
			__fill_ht.nNumOfElements = __fill_idx
			__fill_ht.nNextFreeElement = __fill_idx
			__fill_ht.nInternalPointer = 0
		}
		for {
			var __fill_ht *zend.HashTable = zend.Z_ARRVAL_P(return_value)
			var __fill_bkt *zend.Bucket = __fill_ht.arData + __fill_ht.nNumUsed
			var __fill_idx uint32 = __fill_ht.nNumUsed
			zend.ZEND_ASSERT((zend.HT_FLAGS(__fill_ht) & zend.HASH_FLAG_PACKED) != 0)
			for {
				var __ht *zend.HashTable = zend.Z_ARRVAL_P(input)
				var _p *zend.Bucket = __ht.arData
				var _end *zend.Bucket = _p + __ht.nNumUsed
				for ; _p != _end; _p++ {
					var _z *zend.Zval = &_p.val

					if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
						continue
					}
					value = _z
					zend.Z_TRY_ADDREF_P(value)
					zend.ZEND_HASH_FILL_ADD(value)
				}
				break
			}
			__fill_ht.nNumUsed = __fill_idx
			__fill_ht.nNumOfElements = __fill_idx
			__fill_ht.nNextFreeElement = __fill_idx
			__fill_ht.nInternalPointer = 0
			break
		}
		if pad_size > 0 {
			var __fill_ht *zend.HashTable = zend.Z_ARRVAL_P(return_value)
			var __fill_bkt *zend.Bucket = __fill_ht.arData + __fill_ht.nNumUsed
			var __fill_idx uint32 = __fill_ht.nNumUsed
			zend.ZEND_ASSERT((zend.HT_FLAGS(__fill_ht) & zend.HASH_FLAG_PACKED) != 0)
			for i = 0; i < num_pads; i++ {
				zend.ZEND_HASH_FILL_ADD(pad_value)
			}
			__fill_ht.nNumUsed = __fill_idx
			__fill_ht.nNumOfElements = __fill_idx
			__fill_ht.nNextFreeElement = __fill_idx
			__fill_ht.nInternalPointer = 0
		}
	} else {
		if pad_size < 0 {
			for i = 0; i < num_pads; i++ {
				zend.ZendHashNextIndexInsertNew(zend.Z_ARRVAL_P(return_value), pad_value)
			}
		}
		for {
			var __ht *zend.HashTable = zend.Z_ARRVAL_P(input)
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val
				if zend.Z_TYPE_P(_z) == zend.IS_INDIRECT {
					_z = zend.Z_INDIRECT_P(_z)
				}
				if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
					continue
				}
				key = _p.key
				value = _z
				zend.Z_TRY_ADDREF_P(value)
				if key != nil {
					zend.ZendHashAddNew(zend.Z_ARRVAL_P(return_value), key, value)
				} else {
					zend.ZendHashNextIndexInsertNew(zend.Z_ARRVAL_P(return_value), value)
				}
			}
			break
		}
		if pad_size > 0 {
			for i = 0; i < num_pads; i++ {
				zend.ZendHashNextIndexInsertNew(zend.Z_ARRVAL_P(return_value), pad_value)
			}
		}
	}
}
func ZifArrayFlip(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var array *zend.Zval
	var entry *zend.Zval
	var data zend.Zval
	var num_idx zend.ZendUlong
	var str_idx *zend.ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgArray(_arg, &array, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	zend.ArrayInitSize(return_value, zend.ZendHashNumElements(zend.Z_ARRVAL_P(array)))
	for {
		var __ht *zend.HashTable = zend.Z_ARRVAL_P(array)
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
				continue
			}
			num_idx = _p.h
			str_idx = _p.key
			entry = _z
			zend.ZVAL_DEREF(entry)
			if zend.Z_TYPE_P(entry) == zend.IS_LONG {
				if str_idx != nil {
					zend.ZVAL_STR_COPY(&data, str_idx)
				} else {
					zend.ZVAL_LONG(&data, num_idx)
				}
				zend.ZendHashIndexUpdate(zend.Z_ARRVAL_P(return_value), zend.Z_LVAL_P(entry), &data)
			} else if zend.Z_TYPE_P(entry) == zend.IS_STRING {
				if str_idx != nil {
					zend.ZVAL_STR_COPY(&data, str_idx)
				} else {
					zend.ZVAL_LONG(&data, num_idx)
				}
				zend.ZendSymtableUpdate(zend.Z_ARRVAL_P(return_value), zend.Z_STR_P(entry), &data)
			} else {
				core.PhpErrorDocref(nil, zend.E_WARNING, "Can only flip STRING and INTEGER values!")
			}
		}
		break
	}
}
func ZifArrayChangeKeyCase(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var array *zend.Zval
	var entry *zend.Zval
	var string_key *zend.ZendString
	var new_key *zend.ZendString
	var num_key zend.ZendUlong
	var change_to_upper zend.ZendLong = 0
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgArray(_arg, &array, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &change_to_upper, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	zend.ArrayInitSize(return_value, zend.ZendHashNumElements(zend.Z_ARRVAL_P(array)))
	for {
		var __ht *zend.HashTable = zend.Z_ARRVAL_P(array)
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
				continue
			}
			num_key = _p.h
			string_key = _p.key
			entry = _z
			if string_key == nil {
				entry = zend.ZendHashIndexUpdate(zend.Z_ARRVAL_P(return_value), num_key, entry)
			} else {
				if change_to_upper != 0 {
					new_key = PhpStringToupper(string_key)
				} else {
					new_key = PhpStringTolower(string_key)
				}
				entry = zend.ZendHashUpdate(zend.Z_ARRVAL_P(return_value), new_key, entry)
				zend.ZendStringReleaseEx(new_key, 0)
			}
			zend.ZvalAddRef(entry)
		}
		break
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
func ZifArrayUnique(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var array *zend.Zval
	var idx uint32
	var p *zend.Bucket
	var arTmp *Bucketindex
	var cmpdata *Bucketindex
	var lastkept *Bucketindex
	var i uint
	var sort_type zend.ZendLong = PHP_SORT_STRING
	var cmp zend.CompareFuncT
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgArray(_arg, &array, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &sort_type, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	if zend.Z_ARRVAL_P(array).nNumOfElements <= 1 {
		zend.ZVAL_COPY(return_value, array)
		return
	}
	if sort_type == PHP_SORT_STRING {
		var seen zend.HashTable
		var num_key zend.ZendLong
		var str_key *zend.ZendString
		var val *zend.Zval
		zend.ZendHashInit(&seen, zend.ZendHashNumElements(zend.Z_ARRVAL_P(array)), nil, nil, 0)
		zend.ArrayInit(return_value)
		for {
			var __ht *zend.HashTable = zend.Z_ARRVAL_P(array)
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val
				if zend.Z_TYPE_P(_z) == zend.IS_INDIRECT {
					_z = zend.Z_INDIRECT_P(_z)
				}
				if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
					continue
				}
				num_key = _p.h
				str_key = _p.key
				val = _z
				var retval *zend.Zval
				if zend.Z_TYPE_P(val) == zend.IS_STRING {
					retval = zend.ZendHashAddEmptyElement(&seen, zend.Z_STR_P(val))
				} else {
					var tmp_str_val *zend.ZendString
					var str_val *zend.ZendString = zend.ZvalGetTmpString(val, &tmp_str_val)
					retval = zend.ZendHashAddEmptyElement(&seen, str_val)
					zend.ZendTmpStringRelease(tmp_str_val)
				}
				if retval != nil {

					/* First occurrence of the value */

					if zend.UNEXPECTED(zend.Z_ISREF_P(val) && zend.Z_REFCOUNT_P(val) == 1) {
						zend.ZVAL_DEREF(val)
					}
					zend.Z_TRY_ADDREF_P(val)
					if str_key != nil {
						zend.ZendHashAddNew(zend.Z_ARRVAL_P(return_value), str_key, val)
					} else {
						zend.ZendHashIndexAddNew(zend.Z_ARRVAL_P(return_value), num_key, val)
					}
				}
			}
			break
		}
		zend.ZendHashDestroy(&seen)
		return
	}
	cmp = PhpGetDataCompareFunc(sort_type, 0)
	zend.RETVAL_ARR(zend.ZendArrayDup(zend.Z_ARRVAL_P(array)))

	/* create and sort array with pointers to the target_hash buckets */

	arTmp = (*Bucketindex)(zend.Pemalloc((zend.Z_ARRVAL_P(array).nNumOfElements+1)*b.SizeOf("struct bucketindex"), zend.GC_FLAGS(zend.Z_ARRVAL_P(array))&zend.IS_ARRAY_PERSISTENT))
	i = 0
	idx = 0
	for ; idx < zend.Z_ARRVAL_P(array).nNumUsed; idx++ {
		p = zend.Z_ARRVAL_P(array).arData + idx
		if zend.Z_TYPE(p.val) == zend.IS_UNDEF {
			continue
		}
		if zend.Z_TYPE(p.val) == zend.IS_INDIRECT && zend.Z_TYPE_P(zend.Z_INDIRECT(p.val)) == zend.IS_UNDEF {
			continue
		}
		arTmp[i].SetB(*p)
		arTmp[i].SetI(i)
		i++
	}
	zend.ZVAL_UNDEF(&arTmp[i].b.val)
	zend.ZendSort(any(arTmp), i, b.SizeOf("struct bucketindex"), cmp, zend.SwapFuncT(ArrayBucketindexSwap))

	/* go through the sorted array and delete duplicates from the copy */

	lastkept = arTmp
	for cmpdata = arTmp + 1; zend.Z_TYPE(cmpdata.b.val) != zend.IS_UNDEF; cmpdata++ {
		if cmp(lastkept, cmpdata) != 0 {
			lastkept = cmpdata
		} else {
			if lastkept.GetI() > cmpdata.GetI() {
				p = &lastkept.b
				lastkept = cmpdata
			} else {
				p = &cmpdata.b
			}
			if p.key == nil {
				zend.ZendHashIndexDel(zend.Z_ARRVAL_P(return_value), p.h)
			} else {
				if zend.Z_ARRVAL_P(return_value) == &(zend.ExecutorGlobals.symbol_table) {
					zend.ZendDeleteGlobalVariable(p.key)
				} else {
					zend.ZendHashDel(zend.Z_ARRVAL_P(return_value), p.key)
				}
			}
		}
	}
	zend.Pefree(arTmp, zend.GC_FLAGS(zend.Z_ARRVAL_P(array))&zend.IS_ARRAY_PERSISTENT)
}
func ZvalCompare(first *zend.Zval, second *zend.Zval) int {
	return zend.StringCompareFunction(first, second)
}
func ZvalUserCompare(a *zend.Zval, b *zend.Zval) int {
	var args []zend.Zval
	var retval zend.Zval
	zend.ZVAL_COPY_VALUE(&args[0], a)
	zend.ZVAL_COPY_VALUE(&args[1], b)
	BG(user_compare_fci).param_count = 2
	BG(user_compare_fci).params = args
	BG(user_compare_fci).retval = &retval
	BG(user_compare_fci).no_separation = 0
	if zend.ZendCallFunction(&BG(user_compare_fci), &BG(user_compare_fci_cache)) == zend.SUCCESS && zend.Z_TYPE(retval) != zend.IS_UNDEF {
		var ret zend.ZendLong = zend.ZvalGetLong(&retval)
		zend.ZvalPtrDtor(&retval)
		return zend.ZEND_NORMALIZE_BOOL(ret)
	} else {
		return 0
	}
}
func PhpArrayIntersectKey(execute_data *zend.ZendExecuteData, return_value *zend.Zval, data_compare_type int) {
	var idx uint32
	var p *zend.Bucket
	var argc int
	var i int
	var args *zend.Zval
	var intersect_data_compare_func func(*zend.Zval, *zend.Zval) int = nil
	var ok zend.ZendBool
	var val *zend.Zval
	var data *zend.Zval
	var req_args int
	var param_spec *byte

	/* Get the argument count */

	argc = zend.ZEND_NUM_ARGS()
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
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), param_spec, &args, &argc, &BG(user_compare_fci), &BG(user_compare_fci_cache)) == zend.FAILURE {
		return
	}
	for i = 0; i < argc; i++ {
		if zend.Z_TYPE(args[i]) != zend.IS_ARRAY {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Expected parameter %d to be an array, %s given", i+1, zend.ZendZvalTypeName(&args[i]))
			zend.RETVAL_NULL()
			return
		}
	}
	zend.ArrayInit(return_value)
	for idx = 0; idx < zend.Z_ARRVAL(args[0]).nNumUsed; idx++ {
		p = zend.Z_ARRVAL(args[0]).arData + idx
		val = &p.val
		if zend.Z_TYPE_P(val) == zend.IS_UNDEF {
			continue
		}
		if zend.UNEXPECTED(zend.Z_TYPE_P(val) == zend.IS_INDIRECT) {
			val = zend.Z_INDIRECT_P(val)
			if zend.Z_TYPE_P(val) == zend.IS_UNDEF {
				continue
			}
		}
		if zend.Z_ISREF_P(val) && zend.Z_REFCOUNT_P(val) == 1 {
			val = zend.Z_REFVAL_P(val)
		}
		if p.key == nil {
			ok = 1
			for i = 1; i < argc; i++ {
				if b.Assign(&data, zend.ZendHashIndexFind(zend.Z_ARRVAL(args[i]), p.h)) == nil || intersect_data_compare_func != nil && intersect_data_compare_func(val, data) != 0 {
					ok = 0
					break
				}
			}
			if ok != 0 {
				zend.Z_TRY_ADDREF_P(val)
				zend.ZendHashIndexUpdate(zend.Z_ARRVAL_P(return_value), p.h, val)
			}
		} else {
			ok = 1
			for i = 1; i < argc; i++ {
				if b.Assign(&data, zend.ZendHashFindExInd(zend.Z_ARRVAL(args[i]), p.key, 1)) == nil || intersect_data_compare_func != nil && intersect_data_compare_func(val, data) != 0 {
					ok = 0
					break
				}
			}
			if ok != 0 {
				zend.Z_TRY_ADDREF_P(val)
				zend.ZendHashUpdate(zend.Z_ARRVAL_P(return_value), p.key, val)
			}
		}
	}
}
func PhpArrayIntersect(execute_data *zend.ZendExecuteData, return_value *zend.Zval, behavior int, data_compare_type int, key_compare_type int) {
	var args *zend.Zval = nil
	var hash *zend.HashTable
	var arr_argc int
	var i int
	var c int = 0
	var idx uint32
	var lists **zend.Bucket
	var list **zend.Bucket
	var ptrs ***zend.Bucket
	var p **zend.Bucket
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
		if zend.ZEND_NUM_ARGS() < req_args {
			core.PhpErrorDocref(nil, zend.E_WARNING, "at least %d parameters are required, %d given", req_args, zend.ZEND_NUM_ARGS())
			return
		}
		if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), param_spec, &args, &arr_argc, &fci1, &fci1_cache) == zend.FAILURE {
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
		if zend.ZEND_NUM_ARGS() < req_args {
			core.PhpErrorDocref(nil, zend.E_WARNING, "at least %d parameters are required, %d given", req_args, zend.ZEND_NUM_ARGS())
			return
		}
		if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), param_spec, &args, &arr_argc, &fci1, &fci1_cache, &fci2, &fci2_cache) == zend.FAILURE {
			return
		}
	} else {
		core.PhpErrorDocref(nil, zend.E_WARNING, "behavior is %d. This should never happen. Please report as a bug", behavior)
		return
	}
	PHP_ARRAY_CMP_FUNC_BACKUP()

	/* for each argument, create and sort list with pointers to the hash buckets */

	lists = (**zend.Bucket)(zend.SafeEmalloc(arr_argc, b.SizeOf("Bucket *"), 0))
	ptrs = (**zend.Bucket)(zend.SafeEmalloc(arr_argc, b.SizeOf("Bucket *"), 0))
	if behavior == INTERSECT_NORMAL && data_compare_type == INTERSECT_COMP_DATA_USER {
		BG(user_compare_fci) = *fci_data
		BG(user_compare_fci_cache) = *fci_data_cache
	} else if (behavior&INTERSECT_ASSOC) != 0 && key_compare_type == INTERSECT_COMP_KEY_USER {
		BG(user_compare_fci) = *fci_key
		BG(user_compare_fci_cache) = *fci_key_cache
	}
	for i = 0; i < arr_argc; i++ {
		if zend.Z_TYPE(args[i]) != zend.IS_ARRAY {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Expected parameter %d to be an array, %s given", i+1, zend.ZendZvalTypeName(&args[i]))
			arr_argc = i
			goto out
		}
		hash = zend.Z_ARRVAL(args[i])
		list = (*zend.Bucket)(zend.Pemalloc((hash.nNumOfElements+1)*b.SizeOf("Bucket"), zend.GC_FLAGS(hash)&zend.IS_ARRAY_PERSISTENT))
		lists[i] = list
		ptrs[i] = list
		for idx = 0; idx < hash.nNumUsed; idx++ {
			p = hash.arData + idx
			if zend.Z_TYPE(p.val) == zend.IS_UNDEF {
				continue
			}
			b.PostInc(&(*list)) = *p
		}
		zend.ZVAL_UNDEF(&list.val)
		if hash.nNumOfElements > 1 {
			if behavior == INTERSECT_NORMAL {
				zend.ZendSort(any(lists[i]), hash.nNumOfElements, b.SizeOf("Bucket"), intersect_data_compare_func, zend.SwapFuncT(zend.ZendHashBucketSwap))
			} else if (behavior & INTERSECT_ASSOC) != 0 {
				zend.ZendSort(any(lists[i]), hash.nNumOfElements, b.SizeOf("Bucket"), intersect_key_compare_func, zend.SwapFuncT(zend.ZendHashBucketSwap))
			}
		}
	}

	/* copy the argument array */

	zend.RETVAL_ARR(zend.ZendArrayDup(zend.Z_ARRVAL(args[0])))

	/* go through the lists and look for common values */

	for zend.Z_TYPE(ptrs[0].val) != zend.IS_UNDEF {
		if (behavior&INTERSECT_ASSOC) != 0 && key_compare_type == INTERSECT_COMP_KEY_USER {
			BG(user_compare_fci) = *fci_key
			BG(user_compare_fci_cache) = *fci_key_cache
		}
		for i = 1; i < arr_argc; i++ {
			if (behavior & INTERSECT_NORMAL) != 0 {
				for zend.Z_TYPE(ptrs[i].val) != zend.IS_UNDEF && 0 < b.Assign(&c, intersect_data_compare_func(ptrs[0], ptrs[i])) {
					ptrs[i]++
				}
			} else if (behavior & INTERSECT_ASSOC) != 0 {
				for zend.Z_TYPE(ptrs[i].val) != zend.IS_UNDEF && 0 < b.Assign(&c, intersect_key_compare_func(ptrs[0], ptrs[i])) {
					ptrs[i]++
				}
				if c == 0 && zend.Z_TYPE(ptrs[i].val) != zend.IS_UNDEF && behavior == INTERSECT_ASSOC {

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
			if zend.Z_TYPE(ptrs[i].val) == zend.IS_UNDEF {

				/* delete any values corresponding to remains of ptrs[0] */

				for {
					ptrs[0]++
					p = ptrs[0] - 1
					if zend.Z_TYPE(p.val) == zend.IS_UNDEF {
						goto out
					}
					if p.key == nil {
						zend.ZendHashIndexDel(zend.Z_ARRVAL_P(return_value), p.h)
					} else {
						zend.ZendHashDel(zend.Z_ARRVAL_P(return_value), p.key)
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
				if p.key == nil {
					zend.ZendHashIndexDel(zend.Z_ARRVAL_P(return_value), p.h)
				} else {
					zend.ZendHashDel(zend.Z_ARRVAL_P(return_value), p.key)
				}
				if zend.Z_TYPE(b.PreInc(&ptrs[0]).val) == zend.IS_UNDEF {
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
				if zend.Z_TYPE(b.PreInc(&ptrs[0]).val) == zend.IS_UNDEF {
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
		hash = zend.Z_ARRVAL(args[i])
		zend.Pefree(lists[i], zend.GC_FLAGS(hash)&zend.IS_ARRAY_PERSISTENT)
	}
	PHP_ARRAY_CMP_FUNC_RESTORE()
	zend.Efree(ptrs)
	zend.Efree(lists)
}
func ZifArrayIntersectKey(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpArrayIntersectKey(execute_data, return_value, INTERSECT_COMP_DATA_NONE)
}
func ZifArrayIntersectUkey(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpArrayIntersect(execute_data, return_value, INTERSECT_KEY, INTERSECT_COMP_DATA_INTERNAL, INTERSECT_COMP_KEY_USER)
}
func ZifArrayIntersect(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpArrayIntersect(execute_data, return_value, INTERSECT_NORMAL, INTERSECT_COMP_DATA_INTERNAL, INTERSECT_COMP_KEY_INTERNAL)
}
func ZifArrayUintersect(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpArrayIntersect(execute_data, return_value, INTERSECT_NORMAL, INTERSECT_COMP_DATA_USER, INTERSECT_COMP_KEY_INTERNAL)
}
func ZifArrayIntersectAssoc(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpArrayIntersectKey(execute_data, return_value, INTERSECT_COMP_DATA_INTERNAL)
}
func ZifArrayIntersectUassoc(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpArrayIntersect(execute_data, return_value, INTERSECT_ASSOC, INTERSECT_COMP_DATA_INTERNAL, INTERSECT_COMP_KEY_USER)
}
func ZifArrayUintersectAssoc(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpArrayIntersectKey(execute_data, return_value, INTERSECT_COMP_DATA_USER)
}
func ZifArrayUintersectUassoc(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpArrayIntersect(execute_data, return_value, INTERSECT_ASSOC, INTERSECT_COMP_DATA_USER, INTERSECT_COMP_KEY_USER)
}
func PhpArrayDiffKey(execute_data *zend.ZendExecuteData, return_value *zend.Zval, data_compare_type int) {
	var idx uint32
	var p *zend.Bucket
	var argc int
	var i int
	var args *zend.Zval
	var diff_data_compare_func func(*zend.Zval, *zend.Zval) int = nil
	var ok zend.ZendBool
	var val *zend.Zval
	var data *zend.Zval

	/* Get the argument count */

	argc = zend.ZEND_NUM_ARGS()
	if data_compare_type == DIFF_COMP_DATA_USER {
		if argc < 3 {
			core.PhpErrorDocref(nil, zend.E_WARNING, "at least 3 parameters are required, %d given", zend.ZEND_NUM_ARGS())
			return
		}
		if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "+f", &args, &argc, &BG(user_compare_fci), &BG(user_compare_fci_cache)) == zend.FAILURE {
			return
		}
		diff_data_compare_func = ZvalUserCompare
	} else {
		if argc < 2 {
			core.PhpErrorDocref(nil, zend.E_WARNING, "at least 2 parameters are required, %d given", zend.ZEND_NUM_ARGS())
			return
		}
		if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "+", &args, &argc) == zend.FAILURE {
			return
		}
		if data_compare_type == DIFF_COMP_DATA_INTERNAL {
			diff_data_compare_func = ZvalCompare
		}
	}
	for i = 0; i < argc; i++ {
		if zend.Z_TYPE(args[i]) != zend.IS_ARRAY {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Expected parameter %d to be an array, %s given", i+1, zend.ZendZvalTypeName(&args[i]))
			zend.RETVAL_NULL()
			return
		}
	}
	zend.ArrayInit(return_value)
	for idx = 0; idx < zend.Z_ARRVAL(args[0]).nNumUsed; idx++ {
		p = zend.Z_ARRVAL(args[0]).arData + idx
		val = &p.val
		if zend.Z_TYPE_P(val) == zend.IS_UNDEF {
			continue
		}
		if zend.UNEXPECTED(zend.Z_TYPE_P(val) == zend.IS_INDIRECT) {
			val = zend.Z_INDIRECT_P(val)
			if zend.Z_TYPE_P(val) == zend.IS_UNDEF {
				continue
			}
		}
		if zend.Z_ISREF_P(val) && zend.Z_REFCOUNT_P(val) == 1 {
			val = zend.Z_REFVAL_P(val)
		}
		if p.key == nil {
			ok = 1
			for i = 1; i < argc; i++ {
				if b.Assign(&data, zend.ZendHashIndexFind(zend.Z_ARRVAL(args[i]), p.h)) != nil && (diff_data_compare_func == nil || diff_data_compare_func(val, data) == 0) {
					ok = 0
					break
				}
			}
			if ok != 0 {
				zend.Z_TRY_ADDREF_P(val)
				zend.ZendHashIndexUpdate(zend.Z_ARRVAL_P(return_value), p.h, val)
			}
		} else {
			ok = 1
			for i = 1; i < argc; i++ {
				if b.Assign(&data, zend.ZendHashFindExInd(zend.Z_ARRVAL(args[i]), p.key, 1)) != nil && (diff_data_compare_func == nil || diff_data_compare_func(val, data) == 0) {
					ok = 0
					break
				}
			}
			if ok != 0 {
				zend.Z_TRY_ADDREF_P(val)
				zend.ZendHashUpdate(zend.Z_ARRVAL_P(return_value), p.key, val)
			}
		}
	}
}
func PhpArrayDiff(execute_data *zend.ZendExecuteData, return_value *zend.Zval, behavior int, data_compare_type int, key_compare_type int) {
	var args *zend.Zval = nil
	var hash *zend.HashTable
	var arr_argc int
	var i int
	var c int
	var idx uint32
	var lists **zend.Bucket
	var list **zend.Bucket
	var ptrs ***zend.Bucket
	var p **zend.Bucket
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
		if zend.ZEND_NUM_ARGS() < req_args {
			core.PhpErrorDocref(nil, zend.E_WARNING, "at least %d parameters are required, %d given", req_args, zend.ZEND_NUM_ARGS())
			return
		}
		if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), param_spec, &args, &arr_argc, &fci1, &fci1_cache) == zend.FAILURE {
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
		if zend.ZEND_NUM_ARGS() < req_args {
			core.PhpErrorDocref(nil, zend.E_WARNING, "at least %d parameters are required, %d given", req_args, zend.ZEND_NUM_ARGS())
			return
		}
		if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), param_spec, &args, &arr_argc, &fci1, &fci1_cache, &fci2, &fci2_cache) == zend.FAILURE {
			return
		}
	} else {
		core.PhpErrorDocref(nil, zend.E_WARNING, "behavior is %d. This should never happen. Please report as a bug", behavior)
		return
	}
	PHP_ARRAY_CMP_FUNC_BACKUP()

	/* for each argument, create and sort list with pointers to the hash buckets */

	lists = (**zend.Bucket)(zend.SafeEmalloc(arr_argc, b.SizeOf("Bucket *"), 0))
	ptrs = (**zend.Bucket)(zend.SafeEmalloc(arr_argc, b.SizeOf("Bucket *"), 0))
	if behavior == DIFF_NORMAL && data_compare_type == DIFF_COMP_DATA_USER {
		BG(user_compare_fci) = *fci_data
		BG(user_compare_fci_cache) = *fci_data_cache
	} else if (behavior&DIFF_ASSOC) != 0 && key_compare_type == DIFF_COMP_KEY_USER {
		BG(user_compare_fci) = *fci_key
		BG(user_compare_fci_cache) = *fci_key_cache
	}
	for i = 0; i < arr_argc; i++ {
		if zend.Z_TYPE(args[i]) != zend.IS_ARRAY {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Expected parameter %d to be an array, %s given", i+1, zend.ZendZvalTypeName(&args[i]))
			arr_argc = i
			goto out
		}
		hash = zend.Z_ARRVAL(args[i])
		list = (*zend.Bucket)(zend.Pemalloc((hash.nNumOfElements+1)*b.SizeOf("Bucket"), zend.GC_FLAGS(hash)&zend.IS_ARRAY_PERSISTENT))
		lists[i] = list
		ptrs[i] = list
		for idx = 0; idx < hash.nNumUsed; idx++ {
			p = hash.arData + idx
			if zend.Z_TYPE(p.val) == zend.IS_UNDEF {
				continue
			}
			b.PostInc(&(*list)) = *p
		}
		zend.ZVAL_UNDEF(&list.val)
		if hash.nNumOfElements > 1 {
			if behavior == DIFF_NORMAL {
				zend.ZendSort(any(lists[i]), hash.nNumOfElements, b.SizeOf("Bucket"), diff_data_compare_func, zend.SwapFuncT(zend.ZendHashBucketSwap))
			} else if (behavior & DIFF_ASSOC) != 0 {
				zend.ZendSort(any(lists[i]), hash.nNumOfElements, b.SizeOf("Bucket"), diff_key_compare_func, zend.SwapFuncT(zend.ZendHashBucketSwap))
			}
		}
	}

	/* copy the argument array */

	zend.RETVAL_ARR(zend.ZendArrayDup(zend.Z_ARRVAL(args[0])))

	/* go through the lists and look for values of ptr[0] that are not in the others */

	for zend.Z_TYPE(ptrs[0].val) != zend.IS_UNDEF {
		if (behavior&DIFF_ASSOC) != 0 && key_compare_type == DIFF_COMP_KEY_USER {
			BG(user_compare_fci) = *fci_key
			BG(user_compare_fci_cache) = *fci_key_cache
		}
		c = 1
		for i = 1; i < arr_argc; i++ {
			var ptr *zend.Bucket = ptrs[i]
			if behavior == DIFF_NORMAL {
				for zend.Z_TYPE(ptrs[i].val) != zend.IS_UNDEF && 0 < b.Assign(&c, diff_data_compare_func(ptrs[0], ptrs[i])) {
					ptrs[i]++
				}
			} else if (behavior & DIFF_ASSOC) != 0 {
				for zend.Z_TYPE(ptr.val) != zend.IS_UNDEF && 0 != b.Assign(&c, diff_key_compare_func(ptrs[0], ptr)) {
					ptr++
				}
			}
			if c == 0 {
				if behavior == DIFF_NORMAL {
					if zend.Z_TYPE(ptrs[i].val) != zend.IS_UNDEF {
						ptrs[i]++
					}
					break
				} else if behavior == DIFF_ASSOC {

					/* In this branch is execute only when DIFF_ASSOC. If behavior == DIFF_KEY
					 * data comparison is not needed - skipped. */

					if zend.Z_TYPE(ptr.val) != zend.IS_UNDEF {
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
				if p.key == nil {
					zend.ZendHashIndexDel(zend.Z_ARRVAL_P(return_value), p.h)
				} else {
					zend.ZendHashDel(zend.Z_ARRVAL_P(return_value), p.key)
				}
				if zend.Z_TYPE(b.PreInc(&ptrs[0]).val) == zend.IS_UNDEF {
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
				if zend.Z_TYPE(b.PreInc(&ptrs[0]).val) == zend.IS_UNDEF {
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
		hash = zend.Z_ARRVAL(args[i])
		zend.Pefree(lists[i], zend.GC_FLAGS(hash)&zend.IS_ARRAY_PERSISTENT)
	}
	PHP_ARRAY_CMP_FUNC_RESTORE()
	zend.Efree(ptrs)
	zend.Efree(lists)
}
func ZifArrayDiffKey(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpArrayDiffKey(execute_data, return_value, DIFF_COMP_DATA_NONE)
}
func ZifArrayDiffUkey(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpArrayDiff(execute_data, return_value, DIFF_KEY, DIFF_COMP_DATA_INTERNAL, DIFF_COMP_KEY_USER)
}
func ZifArrayDiff(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var args *zend.Zval
	var argc int
	var i int
	var num uint32
	var exclude zend.HashTable
	var value *zend.Zval
	var str *zend.ZendString
	var tmp_str *zend.ZendString
	var key *zend.ZendString
	var idx zend.ZendLong
	var dummy zend.Zval
	if zend.ZEND_NUM_ARGS() < 2 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "at least 2 parameters are required, %d given", zend.ZEND_NUM_ARGS())
		return
	}
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = -1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			var _num_varargs int = _num_args - _i - 0
			if zend.EXPECTED(_num_varargs > 0) {
				args = _real_arg + 1
				argc = _num_varargs
				_i += _num_varargs
				_real_arg += _num_varargs
			} else {
				args = nil
				argc = 0
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	if zend.Z_TYPE(args[0]) != zend.IS_ARRAY {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Expected parameter 1 to be an array, %s given", zend.ZendZvalTypeName(&args[0]))
		zend.RETVAL_NULL()
		return
	}
	num = zend.ZendHashNumElements(zend.Z_ARRVAL(args[0]))
	if num == 0 {
		for i = 1; i < argc; i++ {
			if zend.Z_TYPE(args[i]) != zend.IS_ARRAY {
				core.PhpErrorDocref(nil, zend.E_WARNING, "Expected parameter %d to be an array, %s given", i+1, zend.ZendZvalTypeName(&args[i]))
				zend.RETVAL_NULL()
				return
			}
		}
		zend.RETVAL_EMPTY_ARRAY()
		return
	} else if num == 1 {
		var found int = 0
		var search_str *zend.ZendString
		var tmp_search_str *zend.ZendString
		value = nil
		for {
			var __ht *zend.HashTable = zend.Z_ARRVAL(args[0])
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val
				if zend.Z_TYPE_P(_z) == zend.IS_INDIRECT {
					_z = zend.Z_INDIRECT_P(_z)
				}
				if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
					continue
				}
				value = _z
				break
			}
			break
		}
		if value == nil {
			for i = 1; i < argc; i++ {
				if zend.Z_TYPE(args[i]) != zend.IS_ARRAY {
					core.PhpErrorDocref(nil, zend.E_WARNING, "Expected parameter %d to be an array, %s given", i+1, zend.ZendZvalTypeName(&args[i]))
					zend.RETVAL_NULL()
					return
				}
			}
			zend.RETVAL_EMPTY_ARRAY()
			return
		}
		search_str = zend.ZvalGetTmpString(value, &tmp_search_str)
		for i = 1; i < argc; i++ {
			if zend.Z_TYPE(args[i]) != zend.IS_ARRAY {
				core.PhpErrorDocref(nil, zend.E_WARNING, "Expected parameter %d to be an array, %s given", i+1, zend.ZendZvalTypeName(&args[i]))
				zend.RETVAL_NULL()
				return
			}
			if found == 0 {
				for {
					var __ht *zend.HashTable = zend.Z_ARRVAL(args[i])
					var _p *zend.Bucket = __ht.arData
					var _end *zend.Bucket = _p + __ht.nNumUsed
					for ; _p != _end; _p++ {
						var _z *zend.Zval = &_p.val
						if zend.Z_TYPE_P(_z) == zend.IS_INDIRECT {
							_z = zend.Z_INDIRECT_P(_z)
						}
						if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
							continue
						}
						value = _z
						str = zend.ZvalGetTmpString(value, &tmp_str)
						if zend.ZendStringEquals(search_str, str) != 0 {
							zend.ZendTmpStringRelease(tmp_str)
							found = 1
							break
						}
						zend.ZendTmpStringRelease(tmp_str)
					}
					break
				}
			}
		}
		zend.ZendTmpStringRelease(tmp_search_str)
		if found != 0 {
			zend.RETVAL_EMPTY_ARRAY()
		} else {
			zend.ZVAL_COPY(return_value, &args[0])
		}
		return
	}

	/* count number of elements */

	num = 0
	for i = 1; i < argc; i++ {
		if zend.Z_TYPE(args[i]) != zend.IS_ARRAY {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Expected parameter %d to be an array, %s given", i+1, zend.ZendZvalTypeName(&args[i]))
			zend.RETVAL_NULL()
			return
		}
		num += zend.ZendHashNumElements(zend.Z_ARRVAL(args[i]))
	}
	if num == 0 {
		zend.ZVAL_COPY(return_value, &args[0])
		return
	}
	zend.ZVAL_NULL(&dummy)

	/* create exclude map */

	zend.ZendHashInit(&exclude, num, nil, nil, 0)
	for i = 1; i < argc; i++ {
		for {
			var __ht *zend.HashTable = zend.Z_ARRVAL(args[i])
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val
				if zend.Z_TYPE_P(_z) == zend.IS_INDIRECT {
					_z = zend.Z_INDIRECT_P(_z)
				}
				if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
					continue
				}
				value = _z
				str = zend.ZvalGetTmpString(value, &tmp_str)
				zend.ZendHashAdd(&exclude, str, &dummy)
				zend.ZendTmpStringRelease(tmp_str)
			}
			break
		}
	}

	/* copy all elements of first array that are not in exclude set */

	zend.ArrayInitSize(return_value, zend.ZendHashNumElements(zend.Z_ARRVAL(args[0])))
	for {
		var __ht *zend.HashTable = zend.Z_ARRVAL(args[0])
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val
			if zend.Z_TYPE_P(_z) == zend.IS_INDIRECT {
				_z = zend.Z_INDIRECT_P(_z)
			}
			if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
				continue
			}
			idx = _p.h
			key = _p.key
			value = _z
			str = zend.ZvalGetTmpString(value, &tmp_str)
			if zend.ZendHashExists(&exclude, str) == 0 {
				if key != nil {
					value = zend.ZendHashAddNew(zend.Z_ARRVAL_P(return_value), key, value)
				} else {
					value = zend.ZendHashIndexAddNew(zend.Z_ARRVAL_P(return_value), idx, value)
				}
				zend.ZvalAddRef(value)
			}
			zend.ZendTmpStringRelease(tmp_str)
		}
		break
	}
	zend.ZendHashDestroy(&exclude)
}
func ZifArrayUdiff(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpArrayDiff(execute_data, return_value, DIFF_NORMAL, DIFF_COMP_DATA_USER, DIFF_COMP_KEY_INTERNAL)
}
func ZifArrayDiffAssoc(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpArrayDiffKey(execute_data, return_value, DIFF_COMP_DATA_INTERNAL)
}
func ZifArrayDiffUassoc(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpArrayDiff(execute_data, return_value, DIFF_ASSOC, DIFF_COMP_DATA_INTERNAL, DIFF_COMP_KEY_USER)
}
func ZifArrayUdiffAssoc(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpArrayDiffKey(execute_data, return_value, DIFF_COMP_DATA_USER)
}
func ZifArrayUdiffUassoc(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpArrayDiff(execute_data, return_value, DIFF_ASSOC, DIFF_COMP_DATA_USER, DIFF_COMP_KEY_USER)
}
func PhpMultisortCompare(a any, b any) int {
	var ab *zend.Bucket = *((**zend.Bucket)(a))
	var bb *zend.Bucket = *((**zend.Bucket)(b))
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
		if zend.Z_TYPE(ab[r].val) == zend.IS_UNDEF {
			break
		}
	}
	return 0
}
func ArrayBucketPSawp(p any, q any) {
	var t *zend.Bucket
	var f **zend.Bucket = (**zend.Bucket)(p)
	var g **zend.Bucket = (**zend.Bucket)(q)
	t = *f
	*f = *g
	*g = t
}
func ZifArrayMultisort(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var args *zend.Zval
	var arrays **zend.Zval
	var indirect **zend.Bucket
	var idx uint32
	var p *zend.Bucket
	var hash *zend.HashTable
	var argc int
	var array_size int
	var num_arrays int = 0
	var parse_state []int
	var sort_order int = PHP_SORT_ASC
	var sort_type int = PHP_SORT_REGULAR
	var i int
	var k int
	var n int
	var func_ *zend.CompareFuncT
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = -1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			var _num_varargs int = _num_args - _i - 0
			if zend.EXPECTED(_num_varargs > 0) {
				args = _real_arg + 1
				argc = _num_varargs
				_i += _num_varargs
				_real_arg += _num_varargs
			} else {
				args = nil
				argc = 0
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}

	/* Allocate space for storing pointers to input arrays and sort flags. */

	arrays = (**zend.Zval)(zend.Ecalloc(argc, b.SizeOf("zval *")))
	for i = 0; i < MULTISORT_LAST; i++ {
		parse_state[i] = 0
	}
	ARRAYG(multisort_func) = (*zend.CompareFuncT)(zend.Ecalloc(argc, b.SizeOf("compare_func_t")))
	func_ = ARRAYG(multisort_func)

	/* Here we go through the input arguments and parse them. Each one can
	 * be either an array or a sort flag which follows an array. If not
	 * specified, the sort flags defaults to PHP_SORT_ASC and PHP_SORT_REGULAR
	 * accordingly. There can't be two sort flags of the same type after an
	 * array, and the very first argument has to be an array. */

	for i = 0; i < argc; i++ {
		var arg *zend.Zval = &args[i]
		zend.ZVAL_DEREF(arg)
		if zend.Z_TYPE_P(arg) == zend.IS_ARRAY {
			zend.SEPARATE_ARRAY(arg)

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

		} else if zend.Z_TYPE_P(arg) == zend.IS_LONG {
			switch zend.Z_LVAL_P(arg) & ^PHP_SORT_FLAG_CASE {
			case PHP_SORT_ASC:

			case PHP_SORT_DESC:

				/* flag allowed here */

				if parse_state[MULTISORT_ORDER] == 1 {

					/* Save the flag and make sure then next arg is not the current flag. */

					if zend.Z_LVAL_P(arg) == PHP_SORT_DESC {
						sort_order = PHP_SORT_DESC
					} else {
						sort_order = PHP_SORT_ASC
					}
					parse_state[MULTISORT_ORDER] = 0
				} else {
					core.PhpErrorDocref(nil, zend.E_WARNING, "Argument #%d is expected to be an array or sorting flag that has not already been specified", i+1)
					zend.Efree(func_)
					zend.Efree(arrays)
					zend.RETVAL_FALSE
					return
				}
				break
			case PHP_SORT_REGULAR:

			case PHP_SORT_NUMERIC:

			case PHP_SORT_STRING:

			case PHP_SORT_NATURAL:

			case PHP_SORT_LOCALE_STRING:

				/* flag allowed here */

				if parse_state[MULTISORT_TYPE] == 1 {

					/* Save the flag and make sure then next arg is not the current flag. */

					sort_type = int(zend.Z_LVAL_P(arg))
					parse_state[MULTISORT_TYPE] = 0
				} else {
					core.PhpErrorDocref(nil, zend.E_WARNING, "Argument #%d is expected to be an array or sorting flag that has not already been specified", i+1)
					zend.Efree(func_)
					zend.Efree(arrays)
					zend.RETVAL_FALSE
					return
				}
				break
			default:
				core.PhpErrorDocref(nil, zend.E_WARNING, "Argument #%d is an unknown sort flag", i+1)
				zend.Efree(func_)
				zend.Efree(arrays)
				zend.RETVAL_FALSE
				return
				break
			}
		} else {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Argument #%d is expected to be an array or a sort flag", i+1)
			zend.Efree(func_)
			zend.Efree(arrays)
			zend.RETVAL_FALSE
			return
		}
	}

	/* Take care of the last array sort flags. */

	ARRAYG(multisort_func)[num_arrays-1] = PhpGetDataCompareFunc(sort_type, sort_order != PHP_SORT_ASC)

	/* Make sure the arrays are of the same size. */

	array_size = zend.ZendHashNumElements(zend.Z_ARRVAL_P(arrays[0]))
	for i = 0; i < num_arrays; i++ {
		if zend.ZendHashNumElements(zend.Z_ARRVAL_P(arrays[i])) != uint32(array_size) {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Array sizes are inconsistent")
			zend.Efree(func_)
			zend.Efree(arrays)
			zend.RETVAL_FALSE
			return
		}
	}

	/* If all arrays are empty we don't need to do anything. */

	if array_size < 1 {
		zend.Efree(func_)
		zend.Efree(arrays)
		zend.RETVAL_TRUE
		return
	}

	/* Create the indirection array. This array is of size MxN, where
	 * M is the number of entries in each input array and N is the number
	 * of the input arrays + 1. The last column is NULL to indicate the end
	 * of the row. */

	indirect = (**zend.Bucket)(zend.SafeEmalloc(array_size, b.SizeOf("Bucket *"), 0))
	for i = 0; i < array_size; i++ {
		indirect[i] = (*zend.Bucket)(zend.SafeEmalloc(num_arrays+1, b.SizeOf("Bucket"), 0))
	}
	for i = 0; i < num_arrays; i++ {
		k = 0
		for idx = 0; idx < zend.Z_ARRVAL_P(arrays[i]).nNumUsed; idx++ {
			p = zend.Z_ARRVAL_P(arrays[i]).arData + idx
			if zend.Z_TYPE(p.val) == zend.IS_UNDEF {
				continue
			}
			indirect[k][i] = *p
			k++
		}
	}
	for k = 0; k < array_size; k++ {
		zend.ZVAL_UNDEF(&indirect[k][num_arrays].val)
	}

	/* Do the actual sort magic - bada-bim, bada-boom. */

	zend.ZendSort(indirect, array_size, b.SizeOf("Bucket *"), PhpMultisortCompare, zend.SwapFuncT(ArrayBucketPSawp))

	/* Restructure the arrays based on sorted indirect - this is mostly taken from zend_hash_sort() function. */

	for i = 0; i < num_arrays; i++ {
		var repack int
		hash = zend.Z_ARRVAL_P(arrays[i])
		hash.nNumUsed = array_size
		hash.nInternalPointer = 0
		repack = !(zend.HT_FLAGS(hash) & zend.HASH_FLAG_PACKED)
		n = 0
		k = 0
		for ; k < array_size; k++ {
			hash.arData[k] = indirect[k][i]
			if hash.arData[k].key == nil {
				n++
				hash.arData[k].h = n - 1
			} else {
				repack = 0
			}
		}
		hash.nNextFreeElement = array_size
		if repack != 0 {
			zend.ZendHashToPacked(hash)
		} else if (zend.HT_FLAGS(hash) & zend.HASH_FLAG_PACKED) == 0 {
			zend.ZendHashRehash(hash)
		}
	}

	/* Clean up. */

	for i = 0; i < array_size; i++ {
		zend.Efree(indirect[i])
	}
	zend.Efree(indirect)
	zend.Efree(func_)
	zend.Efree(arrays)
	zend.RETVAL_TRUE
	return
}
func ZifArrayRand(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var input *zend.Zval
	var num_req zend.ZendLong = 1
	var string_key *zend.ZendString
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
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgArray(_arg, &input, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &num_req, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	num_avail = zend.ZendHashNumElements(zend.Z_ARRVAL_P(input))
	if num_avail == 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Array is empty")
		return
	}
	if num_req == 1 {
		var ht *zend.HashTable = zend.Z_ARRVAL_P(input)
		if uint32(num_avail < ht.nNumUsed-(ht.nNumUsed>>1)) != 0 {

			/* If less than 1/2 of elements are used, don't sample. Instead search for a
			 * specific offset using linear scan. */

			var i zend.ZendLong = 0
			var randval zend.ZendLong = PhpMtRandRange(0, num_avail-1)
			for {
				var __ht *zend.HashTable = zend.Z_ARRVAL_P(input)
				var _p *zend.Bucket = __ht.arData
				var _end *zend.Bucket = _p + __ht.nNumUsed
				for ; _p != _end; _p++ {
					var _z *zend.Zval = &_p.val

					if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
						continue
					}
					num_key = _p.h
					string_key = _p.key
					if i == randval {
						if string_key != nil {
							zend.RETVAL_STR_COPY(string_key)
							return
						} else {
							zend.RETVAL_LONG(num_key)
							return
						}
					}
					i++
				}
				break
			}
		}

		/* Sample random buckets until we hit one that is not empty.
		 * The worst case probability of hitting an empty element is 1-1/2. The worst case
		 * probability of hitting N empty elements in a row is (1-1/2)**N.
		 * For N=10 this becomes smaller than 0.1%. */

		for {
			var randval zend.ZendLong = PhpMtRandRange(0, ht.nNumUsed-1)
			var bucket *zend.Bucket = &ht.arData[randval]
			if !(zend.Z_ISUNDEF(bucket.val)) {
				if bucket.key != nil {
					zend.RETVAL_STR_COPY(bucket.key)
					return
				} else {
					zend.RETVAL_LONG(bucket.h)
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

	zend.ZendHashRealInitPacked(zend.Z_ARRVAL_P(return_value))
	for {
		var __fill_ht *zend.HashTable = zend.Z_ARRVAL_P(return_value)
		var __fill_bkt *zend.Bucket = __fill_ht.arData + __fill_ht.nNumUsed
		var __fill_idx uint32 = __fill_ht.nNumUsed
		zend.ZEND_ASSERT((zend.HT_FLAGS(__fill_ht) & zend.HASH_FLAG_PACKED) != 0)

		/* We can't use zend_hash_index_find()
		 * because the array may have string keys or gaps. */

		for {
			var __ht *zend.HashTable = zend.Z_ARRVAL_P(input)
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val

				if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
					continue
				}
				num_key = _p.h
				string_key = _p.key
				if (zend.ZendBitsetIn(bitset, i) ^ negative_bitset) != 0 {
					if string_key != nil {
						zend.ZEND_HASH_FILL_SET_STR_COPY(string_key)
					} else {
						zend.ZEND_HASH_FILL_SET_LONG(num_key)
					}
					zend.ZEND_HASH_FILL_NEXT()
				}
				i++
			}
			break
		}

		/* We can't use zend_hash_index_find()
		 * because the array may have string keys or gaps. */

		__fill_ht.nNumUsed = __fill_idx
		__fill_ht.nNumOfElements = __fill_idx
		__fill_ht.nNextFreeElement = __fill_idx
		__fill_ht.nInternalPointer = 0
		break
	}
	zend.FreeAlloca(bitset, use_heap)
}
func ZifArraySum(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var input *zend.Zval
	var entry *zend.Zval
	var entry_n zend.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgArray(_arg, &input, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	zend.ZVAL_LONG(return_value, 0)
	for {
		var __ht *zend.HashTable = zend.Z_ARRVAL_P(input)
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
				continue
			}
			entry = _z
			if zend.Z_TYPE_P(entry) == zend.IS_ARRAY || zend.Z_TYPE_P(entry) == zend.IS_OBJECT {
				continue
			}
			zend.ZVAL_COPY(&entry_n, entry)
			zend.ConvertScalarToNumber(&entry_n)
			zend.FastAddFunction(return_value, return_value, &entry_n)
		}
		break
	}
}
func ZifArrayProduct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var input *zend.Zval
	var entry *zend.Zval
	var entry_n zend.Zval
	var dval float64
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgArray(_arg, &input, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	zend.ZVAL_LONG(return_value, 1)
	if !(zend.ZendHashNumElements(zend.Z_ARRVAL_P(input))) {
		return
	}
	for {
		var __ht *zend.HashTable = zend.Z_ARRVAL_P(input)
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
				continue
			}
			entry = _z
			if zend.Z_TYPE_P(entry) == zend.IS_ARRAY || zend.Z_TYPE_P(entry) == zend.IS_OBJECT {
				continue
			}
			zend.ZVAL_COPY(&entry_n, entry)
			zend.ConvertScalarToNumber(&entry_n)
			if zend.Z_TYPE(entry_n) == zend.IS_LONG && zend.Z_TYPE_P(return_value) == zend.IS_LONG {
				dval = float64(zend.Z_LVAL_P(return_value) * float64(zend.Z_LVAL(entry_n)))
				if float64(zend.ZEND_LONG_MIN <= dval && dval <= float64(zend.ZEND_LONG_MAX)) {
					zend.Z_LVAL_P(return_value) *= zend.Z_LVAL(entry_n)
					continue
				}
			}
			zend.ConvertToDouble(return_value)
			zend.ConvertToDouble(&entry_n)
			zend.Z_DVAL_P(return_value) *= zend.Z_DVAL(entry_n)
		}
		break
	}
}
func ZifArrayReduce(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var input *zend.Zval
	var args []zend.Zval
	var operand *zend.Zval
	var result zend.Zval
	var retval zend.Zval
	var fci zend.ZendFcallInfo
	var fci_cache zend.ZendFcallInfoCache = zend.EmptyFcallInfoCache
	var initial *zend.Zval = nil
	var htbl *zend.HashTable
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 3
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgArray(_arg, &input, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgFunc(_arg, &fci, &fci_cache, 0, &_error) == 0) {
				if _error == nil {
					_expected_type = zend.Z_EXPECTED_FUNC
					_error_code = zend.ZPP_ERROR_WRONG_ARG
					break
				} else {
					_error_code = zend.ZPP_ERROR_WRONG_CALLBACK
					break
				}
			} else if zend.UNEXPECTED(_error != nil) {
				zend.ZendWrongCallbackDeprecated(_i, _error)
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &initial, 0)
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	if zend.ZEND_NUM_ARGS() > 2 {
		zend.ZVAL_COPY(&result, initial)
	} else {
		zend.ZVAL_NULL(&result)
	}

	/* (zval **)input points to an element of argument stack
	 * the base pointer of which is subject to change.
	 * thus we need to keep the pointer to the hashtable for safety */

	htbl = zend.Z_ARRVAL_P(input)
	if zend.ZendHashNumElements(htbl) == 0 {
		zend.ZVAL_COPY_VALUE(return_value, &result)
		zend.ZendReleaseFcallInfoCache(&fci_cache)
		return
	}
	fci.retval = &retval
	fci.param_count = 2
	fci.no_separation = 0
	for {
		var __ht *zend.HashTable = htbl
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
				continue
			}
			operand = _z
			zend.ZVAL_COPY_VALUE(&args[0], &result)
			zend.ZVAL_COPY(&args[1], operand)
			fci.params = args
			if zend.ZendCallFunction(&fci, &fci_cache) == zend.SUCCESS && zend.Z_TYPE(retval) != zend.IS_UNDEF {
				zend.ZvalPtrDtor(&args[1])
				zend.ZvalPtrDtor(&args[0])
				zend.ZVAL_COPY_VALUE(&result, &retval)
			} else {
				zend.ZvalPtrDtor(&args[1])
				zend.ZvalPtrDtor(&args[0])
				return
			}
		}
		break
	}
	zend.ZendReleaseFcallInfoCache(&fci_cache)
	zend.RETVAL_ZVAL(&result, 1, 1)
}
func ZifArrayFilter(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var array *zend.Zval
	var operand *zend.Zval
	var key *zend.Zval
	var args []zend.Zval
	var retval zend.Zval
	var have_callback zend.ZendBool = 0
	var use_type zend.ZendLong = 0
	var string_key *zend.ZendString
	var fci zend.ZendFcallInfo = zend.EmptyFcallInfo
	var fci_cache zend.ZendFcallInfoCache = zend.EmptyFcallInfoCache
	var num_key zend.ZendUlong
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 3
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgArray(_arg, &array, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgFunc(_arg, &fci, &fci_cache, 0, &_error) == 0) {
				if _error == nil {
					_expected_type = zend.Z_EXPECTED_FUNC
					_error_code = zend.ZPP_ERROR_WRONG_ARG
					break
				} else {
					_error_code = zend.ZPP_ERROR_WRONG_CALLBACK
					break
				}
			} else if zend.UNEXPECTED(_error != nil) {
				zend.ZendWrongCallbackDeprecated(_i, _error)
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &use_type, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	zend.ArrayInit(return_value)
	if zend.ZendHashNumElements(zend.Z_ARRVAL_P(array)) == 0 {
		zend.ZendReleaseFcallInfoCache(&fci_cache)
		return
	}
	if zend.ZEND_NUM_ARGS() > 1 {
		have_callback = 1
		fci.no_separation = 0
		fci.retval = &retval
		if use_type == ARRAY_FILTER_USE_BOTH {
			fci.param_count = 2
			key = &args[1]
		} else {
			fci.param_count = 1
			key = &args[0]
		}
	}
	for {
		var __ht *zend.HashTable = zend.Z_ARRVAL_P(array)
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val
			if zend.Z_TYPE_P(_z) == zend.IS_INDIRECT {
				_z = zend.Z_INDIRECT_P(_z)
			}
			if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
				continue
			}
			num_key = _p.h
			string_key = _p.key
			operand = _z
			if have_callback != 0 {
				if use_type != 0 {

					/* Set up the key */

					if string_key == nil {
						zend.ZVAL_LONG(key, num_key)
					} else {
						zend.ZVAL_STR_COPY(key, string_key)
					}

					/* Set up the key */

				}
				if use_type != ARRAY_FILTER_USE_KEY {
					zend.ZVAL_COPY(&args[0], operand)
				}
				fci.params = args
				if zend.ZendCallFunction(&fci, &fci_cache) == zend.SUCCESS {
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
				operand = zend.ZendHashUpdate(zend.Z_ARRVAL_P(return_value), string_key, operand)
			} else {
				operand = zend.ZendHashIndexUpdate(zend.Z_ARRVAL_P(return_value), num_key, operand)
			}
			zend.ZvalAddRef(operand)
		}
		break
	}
	zend.ZendReleaseFcallInfoCache(&fci_cache)
}
func ZifArrayMap(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var arrays *zend.Zval = nil
	var n_arrays int = 0
	var result zend.Zval
	var fci zend.ZendFcallInfo = zend.EmptyFcallInfo
	var fci_cache zend.ZendFcallInfoCache = zend.EmptyFcallInfoCache
	var i int
	var k uint32
	var maxlen uint32 = 0
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = -1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgFunc(_arg, &fci, &fci_cache, 1, &_error) == 0) {
				if _error == nil {
					_expected_type = zend.Z_EXPECTED_FUNC
					_error_code = zend.ZPP_ERROR_WRONG_ARG
					break
				} else {
					_error_code = zend.ZPP_ERROR_WRONG_CALLBACK
					break
				}
			} else if zend.UNEXPECTED(_error != nil) {
				zend.ZendWrongCallbackDeprecated(_i, _error)
			}
			var _num_varargs int = _num_args - _i - 0
			if zend.EXPECTED(_num_varargs > 0) {
				arrays = _real_arg + 1
				n_arrays = _num_varargs
				_i += _num_varargs
				_real_arg += _num_varargs
			} else {
				arrays = nil
				n_arrays = 0
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	zend.RETVAL_NULL()
	if n_arrays == 1 {
		var num_key zend.ZendUlong
		var str_key *zend.ZendString
		var zv *zend.Zval
		var arg zend.Zval
		var ret int
		if zend.Z_TYPE(arrays[0]) != zend.IS_ARRAY {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Expected parameter 2 to be an array, %s given", zend.ZendZvalTypeName(&arrays[0]))
			return
		}
		maxlen = zend.ZendHashNumElements(zend.Z_ARRVAL(arrays[0]))

		/* Short-circuit: if no callback and only one array, just return it. */

		if !(zend.ZEND_FCI_INITIALIZED(fci)) || maxlen == 0 {
			zend.ZVAL_COPY(return_value, &arrays[0])
			zend.ZendReleaseFcallInfoCache(&fci_cache)
			return
		}
		zend.ArrayInitSize(return_value, maxlen)
		zend.ZendHashRealInit(zend.Z_ARRVAL_P(return_value), zend.HT_FLAGS(zend.Z_ARRVAL(arrays[0]))&zend.HASH_FLAG_PACKED)
		for {
			var __ht *zend.HashTable = zend.Z_ARRVAL(arrays[0])
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val
				if zend.Z_TYPE_P(_z) == zend.IS_INDIRECT {
					_z = zend.Z_INDIRECT_P(_z)
				}
				if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
					continue
				}
				num_key = _p.h
				str_key = _p.key
				zv = _z
				fci.retval = &result
				fci.param_count = 1
				fci.params = &arg
				fci.no_separation = 0
				zend.ZVAL_COPY(&arg, zv)
				ret = zend.ZendCallFunction(&fci, &fci_cache)
				zend.IZvalPtrDtor(&arg)
				if ret != zend.SUCCESS || zend.Z_TYPE(result) == zend.IS_UNDEF {
					zend.ZendArrayDestroy(zend.Z_ARR_P(return_value))
					zend.RETVAL_NULL()
					return
				}
				if str_key != nil {
					zend._zendHashAppend(zend.Z_ARRVAL_P(return_value), str_key, &result)
				} else {
					zend.ZendHashIndexAddNew(zend.Z_ARRVAL_P(return_value), num_key, &result)
				}
			}
			break
		}
		zend.ZendReleaseFcallInfoCache(&fci_cache)
	} else {
		var array_pos *uint32 = (*zend.HashPosition)(zend.Ecalloc(n_arrays, b.SizeOf("HashPosition")))
		for i = 0; i < n_arrays; i++ {
			if zend.Z_TYPE(arrays[i]) != zend.IS_ARRAY {
				core.PhpErrorDocref(nil, zend.E_WARNING, "Expected parameter %d to be an array, %s given", i+2, zend.ZendZvalTypeName(&arrays[i]))
				zend.Efree(array_pos)
				return
			}
			if zend.ZendHashNumElements(zend.Z_ARRVAL(arrays[i])) > maxlen {
				maxlen = zend.ZendHashNumElements(zend.Z_ARRVAL(arrays[i]))
			}
		}
		zend.ArrayInitSize(return_value, maxlen)
		if !(zend.ZEND_FCI_INITIALIZED(fci)) {
			var zv zend.Zval

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
						if pos >= zend.Z_ARRVAL(arrays[i]).nNumUsed {
							zend.ZVAL_NULL(&zv)
							break
						} else if zend.Z_TYPE(zend.Z_ARRVAL(arrays[i]).arData[pos].val) != zend.IS_UNDEF {
							zend.ZVAL_COPY(&zv, &zend.Z_ARRVAL(arrays[i]).arData[pos].val)
							array_pos[i] = pos + 1
							break
						}
						pos++
					}
					zend.ZendHashNextIndexInsertNew(zend.Z_ARRVAL(result), &zv)
				}
				zend.ZendHashNextIndexInsertNew(zend.Z_ARRVAL_P(return_value), &result)
			}

			/* We iterate through all the arrays at once. */

		} else {
			var params *zend.Zval = (*zend.Zval)(zend.SafeEmalloc(n_arrays, b.SizeOf("zval"), 0))

			/* We iterate through all the arrays at once. */

			for k = 0; k < maxlen; k++ {
				for i = 0; i < n_arrays; i++ {

					/* If this array still has elements, add the current one to the
					 * parameter list, otherwise use null value. */

					var pos uint32 = array_pos[i]
					for true {
						if pos >= zend.Z_ARRVAL(arrays[i]).nNumUsed {
							zend.ZVAL_NULL(&params[i])
							break
						} else if zend.Z_TYPE(zend.Z_ARRVAL(arrays[i]).arData[pos].val) != zend.IS_UNDEF {
							zend.ZVAL_COPY(&params[i], &zend.Z_ARRVAL(arrays[i]).arData[pos].val)
							array_pos[i] = pos + 1
							break
						}
						pos++
					}
				}
				fci.retval = &result
				fci.param_count = n_arrays
				fci.params = params
				fci.no_separation = 0
				if zend.ZendCallFunction(&fci, &fci_cache) != zend.SUCCESS || zend.Z_TYPE(result) == zend.IS_UNDEF {
					zend.Efree(array_pos)
					zend.ZendArrayDestroy(zend.Z_ARR_P(return_value))
					for i = 0; i < n_arrays; i++ {
						zend.ZvalPtrDtor(&params[i])
					}
					zend.Efree(params)
					zend.RETVAL_NULL()
					return
				} else {
					for i = 0; i < n_arrays; i++ {
						zend.ZvalPtrDtor(&params[i])
					}
				}
				zend.ZendHashNextIndexInsertNew(zend.Z_ARRVAL_P(return_value), &result)
			}
			zend.Efree(params)
			zend.ZendReleaseFcallInfoCache(&fci_cache)
		}
		zend.Efree(array_pos)
	}
}
func ZifArrayKeyExists(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var key *zend.Zval
	var array *zend.Zval
	var ht *zend.HashTable
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 2
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &key, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgArray(_arg, &array, 0, 1) == 0) {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	if zend.EXPECTED(zend.Z_TYPE_P(array) == zend.IS_ARRAY) {
		ht = zend.Z_ARRVAL_P(array)
	} else {
		ht = zend.ZendGetPropertiesFor(array, zend.ZEND_PROP_PURPOSE_ARRAY_CAST)
		core.PhpErrorDocref(nil, zend.E_DEPRECATED, "Using array_key_exists() on objects is deprecated. "+"Use isset() or property_exists() instead")
	}
	switch zend.Z_TYPE_P(key) {
	case zend.IS_STRING:
		zend.RETVAL_BOOL(zend.ZendSymtableExistsInd(ht, zend.Z_STR_P(key)) != 0)
		break
	case zend.IS_LONG:
		zend.RETVAL_BOOL(zend.ZendHashIndexExists(ht, zend.Z_LVAL_P(key)) != 0)
		break
	case zend.IS_NULL:
		zend.RETVAL_BOOL(zend.ZendHashExistsInd(ht, zend.ZSTR_EMPTY_ALLOC()) != 0)
		break
	default:
		core.PhpErrorDocref(nil, zend.E_WARNING, "The first argument should be either a string or an integer")
		zend.RETVAL_FALSE
	}
	if zend.Z_TYPE_P(array) != zend.IS_ARRAY {
		zend.ZendReleaseProperties(ht)
	}
}
func ZifArrayChunk(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var num_in int
	var size zend.ZendLong
	var current zend.ZendLong = 0
	var str_key *zend.ZendString
	var num_key zend.ZendUlong
	var preserve_keys zend.ZendBool = 0
	var input *zend.Zval = nil
	var chunk zend.Zval
	var entry *zend.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 3
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgArray(_arg, &input, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &size, &_dummy, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgBool(_arg, &preserve_keys, &_dummy, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_BOOL
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}

	/* Do bounds checking for size parameter. */

	if size < 1 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Size parameter expected to be greater than 0")
		return
	}
	num_in = zend.ZendHashNumElements(zend.Z_ARRVAL_P(input))
	if size > num_in {
		if num_in > 0 {
			size = num_in
		} else {
			size = 1
		}
	}
	zend.ArrayInitSize(return_value, uint32_t((num_in-1)/size+1))
	zend.ZVAL_UNDEF(&chunk)
	for {
		var __ht *zend.HashTable = zend.Z_ARRVAL_P(input)
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
				continue
			}
			num_key = _p.h
			str_key = _p.key
			entry = _z

			/* If new chunk, create and initialize it. */

			if zend.Z_TYPE(chunk) == zend.IS_UNDEF {
				zend.ArrayInitSize(&chunk, uint32(size))
			}

			/* Add entry to the chunk, preserving keys if necessary. */

			if preserve_keys != 0 {
				if str_key != nil {
					entry = zend.ZendHashUpdate(zend.Z_ARRVAL(chunk), str_key, entry)
				} else {
					entry = zend.ZendHashIndexUpdate(zend.Z_ARRVAL(chunk), num_key, entry)
				}
			} else {
				entry = zend.ZendHashNextIndexInsert(zend.Z_ARRVAL(chunk), entry)
			}
			zend.ZvalAddRef(entry)

			/* If reached the chunk size, add it to the result array, and reset the
			 * pointer. */

			if b.PreInc(&current)%size == 0 {
				zend.AddNextIndexZval(return_value, &chunk)
				zend.ZVAL_UNDEF(&chunk)
			}

			/* If reached the chunk size, add it to the result array, and reset the
			 * pointer. */

		}
		break
	}

	/* Add the final chunk if there is one. */

	if zend.Z_TYPE(chunk) != zend.IS_UNDEF {
		zend.AddNextIndexZval(return_value, &chunk)
	}

	/* Add the final chunk if there is one. */
}
func ZifArrayCombine(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var values *zend.HashTable
	var keys *zend.HashTable
	var pos_values uint32 = 0
	var entry_keys *zend.Zval
	var entry_values *zend.Zval
	var num_keys int
	var num_values int
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 2
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgArrayHt(_arg, &keys, 0, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgArrayHt(_arg, &values, 0, 0, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	num_keys = zend.ZendHashNumElements(keys)
	num_values = zend.ZendHashNumElements(values)
	if num_keys != num_values {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Both parameters should have an equal number of elements")
		zend.RETVAL_FALSE
		return
	}
	if num_keys == 0 {
		zend.RETVAL_EMPTY_ARRAY()
		return
	}
	zend.ArrayInitSize(return_value, num_keys)
	for {
		var __ht *zend.HashTable = keys
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
				continue
			}
			entry_keys = _z
			for true {
				if pos_values >= values.nNumUsed {
					break
				} else if zend.Z_TYPE(values.arData[pos_values].val) != zend.IS_UNDEF {
					entry_values = &values.arData[pos_values].val
					if zend.Z_TYPE_P(entry_keys) == zend.IS_LONG {
						entry_values = zend.ZendHashIndexUpdate(zend.Z_ARRVAL_P(return_value), zend.Z_LVAL_P(entry_keys), entry_values)
					} else {
						var tmp_key *zend.ZendString
						var key *zend.ZendString = zend.ZvalGetTmpString(entry_keys, &tmp_key)
						entry_values = zend.ZendSymtableUpdate(zend.Z_ARRVAL_P(return_value), key, entry_values)
						zend.ZendTmpStringRelease(tmp_key)
					}
					zend.ZvalAddRef(entry_values)
					pos_values++
					break
				}
				pos_values++
			}
		}
		break
	}
}
