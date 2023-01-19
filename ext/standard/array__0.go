// <<generate>>

package standard

import (
	"sik/core"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <ext/standard/array.c>

/*
   +----------------------------------------------------------------------+
   | PHP Version 7                                                        |
   +----------------------------------------------------------------------+
   | Copyright (c) The PHP Group                                          |
   +----------------------------------------------------------------------+
   | This source file is subject to version 3.01 of the PHP license,      |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.php.net/license/3_01.txt                                  |
   | If you did not receive a copy of the PHP license and are unable to   |
   | obtain it through the world-wide-web, please send a note to          |
   | license@php.net so we can mail you a copy immediately.               |
   +----------------------------------------------------------------------+
   | Authors: Andi Gutmans <andi@php.net>                                 |
   |          Zeev Suraski <zeev@php.net>                                 |
   |          Rasmus Lerdorf <rasmus@php.net>                             |
   |          Andrei Zmievski <andrei@php.net>                            |
   |          Stig Venaas <venaas@php.net>                                |
   |          Jason Greene <jason@php.net>                                |
   +----------------------------------------------------------------------+
*/

// # include "php.h"

// # include "php_ini.h"

// # include < stdarg . h >

// # include < stdlib . h >

// # include < math . h >

// # include < time . h >

// # include < stdio . h >

// # include < string . h >

// # include "zend_globals.h"

// # include "zend_interfaces.h"

// # include "php_globals.h"

// # include "php_array.h"

// # include "basic_functions.h"

// # include "php_string.h"

// # include "php_rand.h"

// # include "php_math.h"

// # include "zend_smart_str.h"

// # include "zend_bitset.h"

// # include "ext/spl/spl_array.h"

/* {{{ defines */

// #define EXTR_OVERWRITE       0

// #define EXTR_SKIP       1

// #define EXTR_PREFIX_SAME       2

// #define EXTR_PREFIX_ALL       3

// #define EXTR_PREFIX_INVALID       4

// #define EXTR_PREFIX_IF_EXISTS       5

// #define EXTR_IF_EXISTS       6

// #define EXTR_REFS       0x100

// #define CASE_LOWER       0

// #define CASE_UPPER       1

// #define DIFF_NORMAL       1

// #define DIFF_KEY       2

// #define DIFF_ASSOC       6

// #define DIFF_COMP_DATA_NONE       - 1

// #define DIFF_COMP_DATA_INTERNAL       0

// #define DIFF_COMP_DATA_USER       1

// #define DIFF_COMP_KEY_INTERNAL       0

// #define DIFF_COMP_KEY_USER       1

// #define INTERSECT_NORMAL       1

// #define INTERSECT_KEY       2

// #define INTERSECT_ASSOC       6

// #define INTERSECT_COMP_DATA_NONE       - 1

// #define INTERSECT_COMP_DATA_INTERNAL       0

// #define INTERSECT_COMP_DATA_USER       1

// #define INTERSECT_COMP_KEY_INTERNAL       0

// #define INTERSECT_COMP_KEY_USER       1

/* }}} */

var ArrayGlobals ZendArrayGlobals

/* {{{ php_array_init_globals
 */

func PhpArrayInitGlobals(array_globals *ZendArrayGlobals) {
	memset(array_globals, 0, g.SizeOf("zend_array_globals"))
}

/* }}} */

func ZmStartupArray(type_ int, module_number int) int {
	PhpArrayInitGlobals(&ArrayGlobals)
	zend.ZendRegisterLongConstant("EXTR_OVERWRITE", g.SizeOf("\"EXTR_OVERWRITE\"")-1, 0, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("EXTR_SKIP", g.SizeOf("\"EXTR_SKIP\"")-1, 1, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("EXTR_PREFIX_SAME", g.SizeOf("\"EXTR_PREFIX_SAME\"")-1, 2, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("EXTR_PREFIX_ALL", g.SizeOf("\"EXTR_PREFIX_ALL\"")-1, 3, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("EXTR_PREFIX_INVALID", g.SizeOf("\"EXTR_PREFIX_INVALID\"")-1, 4, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("EXTR_PREFIX_IF_EXISTS", g.SizeOf("\"EXTR_PREFIX_IF_EXISTS\"")-1, 5, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("EXTR_IF_EXISTS", g.SizeOf("\"EXTR_IF_EXISTS\"")-1, 6, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("EXTR_REFS", g.SizeOf("\"EXTR_REFS\"")-1, 0x100, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("SORT_ASC", g.SizeOf("\"SORT_ASC\"")-1, 4, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("SORT_DESC", g.SizeOf("\"SORT_DESC\"")-1, 3, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("SORT_REGULAR", g.SizeOf("\"SORT_REGULAR\"")-1, 0, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("SORT_NUMERIC", g.SizeOf("\"SORT_NUMERIC\"")-1, 1, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("SORT_STRING", g.SizeOf("\"SORT_STRING\"")-1, 2, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("SORT_LOCALE_STRING", g.SizeOf("\"SORT_LOCALE_STRING\"")-1, 5, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("SORT_NATURAL", g.SizeOf("\"SORT_NATURAL\"")-1, 6, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("SORT_FLAG_CASE", g.SizeOf("\"SORT_FLAG_CASE\"")-1, 8, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("CASE_LOWER", g.SizeOf("\"CASE_LOWER\"")-1, 0, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("CASE_UPPER", g.SizeOf("\"CASE_UPPER\"")-1, 1, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("COUNT_NORMAL", g.SizeOf("\"COUNT_NORMAL\"")-1, 0, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("COUNT_RECURSIVE", g.SizeOf("\"COUNT_RECURSIVE\"")-1, 1, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("ARRAY_FILTER_USE_BOTH", g.SizeOf("\"ARRAY_FILTER_USE_BOTH\"")-1, 1, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("ARRAY_FILTER_USE_KEY", g.SizeOf("\"ARRAY_FILTER_USE_KEY\"")-1, 2, 1<<0|1<<1, module_number)
	return zend.SUCCESS
}

/* }}} */

func ZmShutdownArray(type_ int, module_number int) int { return zend.SUCCESS }

/* }}} */

func PhpArrayKeyCompare(a any, b any) int {
	var f *zend.Bucket = (*zend.Bucket)(a)
	var s *zend.Bucket = (*zend.Bucket)(b)
	var t zend.ZendUchar
	var l1 zend.ZendLong
	var l2 zend.ZendLong
	var d float64
	if f.key == nil {
		if s.key == nil {
			return zend.ZendLong(f.h > zend.ZendLong(g.Cond(s.h != 0, 1, -1)))
		} else {
			l1 = zend.ZendLong(f.h)
			t = zend.IsNumericString(s.key.val, s.key.len_, &l2, &d, 1)
			if t == 4 {

			} else if t == 5 {
				if float64(l1 - d) {
					if float64(l1-d) < 0 {
						return -1
					} else {
						return 1
					}
				} else {
					return 0
				}
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
			if t == 4 {

			} else if t == 5 {
				if d-float64(l2) != 0 {
					if d-float64(l2) < 0 {
						return -1
					} else {
						return 1
					}
				} else {
					return 0
				}
			} else {
				l1 = 0
			}
		}
	}
	if l1-l2 != 0 {
		if l1-l2 < 0 {
			return -1
		} else {
			return 1
		}
	} else {
		return 0
	}
}

/* }}} */

func PhpArrayReverseKeyCompare(a any, b any) int { return PhpArrayKeyCompare(b, a) }

/* }}} */

func PhpArrayKeyCompareNumeric(a any, b any) int {
	var f *zend.Bucket = (*zend.Bucket)(a)
	var s *zend.Bucket = (*zend.Bucket)(b)
	if f.key == nil && s.key == nil {
		return zend.ZendLong(f.h > zend.ZendLong(g.Cond(s.h != 0, 1, -1)))
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
		if d1-d2 != 0 {
			if d1-d2 < 0 {
				return -1
			} else {
				return 1
			}
		} else {
			return 0
		}
	}
}

/* }}} */

func PhpArrayReverseKeyCompareNumeric(a any, b any) int { return PhpArrayKeyCompareNumeric(b, a) }

/* }}} */

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
		s1 = zend.ZendPrintLongToBuf(buf1+g.SizeOf("buf1")-1, f.h)
		l1 = buf1 + g.SizeOf("buf1") - 1 - s1
	}
	if s.key != nil {
		s2 = s.key.val
		l2 = s.key.len_
	} else {
		s2 = zend.ZendPrintLongToBuf(buf2+g.SizeOf("buf2")-1, s.h)
		l2 = buf2 + g.SizeOf("buf2") - 1 - s1
	}
	return zend.ZendBinaryStrcasecmpL(s1, l1, s2, l2)
}

/* }}} */

func PhpArrayReverseKeyCompareStringCase(a any, b any) int {
	return PhpArrayKeyCompareStringCase(b, a)
}

/* }}} */

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
		s1 = zend.ZendPrintLongToBuf(buf1+g.SizeOf("buf1")-1, f.h)
		l1 = buf1 + g.SizeOf("buf1") - 1 - s1
	}
	if s.key != nil {
		s2 = s.key.val
		l2 = s.key.len_
	} else {
		s2 = zend.ZendPrintLongToBuf(buf2+g.SizeOf("buf2")-1, s.h)
		l2 = buf2 + g.SizeOf("buf2") - 1 - s2
	}
	return zend.ZendBinaryStrcmp(s1, l1, s2, l2)
}

/* }}} */

func PhpArrayReverseKeyCompareString(a any, b any) int { return PhpArrayKeyCompareString(b, a) }

/* }}} */

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
		s1 = zend.ZendPrintLongToBuf(buf1+g.SizeOf("buf1")-1, f.h)
		l1 = buf1 + g.SizeOf("buf1") - 1 - s1
	}
	if s.key != nil {
		s2 = s.key.val
		l2 = s.key.len_
	} else {
		s2 = zend.ZendPrintLongToBuf(buf2+g.SizeOf("buf2")-1, s.h)
		l2 = buf2 + g.SizeOf("buf2") - 1 - s1
	}
	return StrnatcmpEx(s1, l1, s2, l2, fold_case)
}

/* }}} */

func PhpArrayKeyCompareStringNaturalCase(a any, b any) int {
	return PhpArrayKeyCompareStringNaturalGeneral(a, b, 1)
}

/* }}} */

func PhpArrayReverseKeyCompareStringNaturalCase(a any, b any) int {
	return PhpArrayKeyCompareStringNaturalGeneral(b, a, 1)
}

/* }}} */

func PhpArrayKeyCompareStringNatural(a any, b any) int {
	return PhpArrayKeyCompareStringNaturalGeneral(a, b, 0)
}

/* }}} */

func PhpArrayReverseKeyCompareStringNatural(a any, b any) int {
	return PhpArrayKeyCompareStringNaturalGeneral(b, a, 0)
}

/* }}} */

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
		s1 = zend.ZendPrintLongToBuf(buf1+g.SizeOf("buf1")-1, f.h)
	}
	if s.key != nil {
		s2 = s.key.val
	} else {
		s2 = zend.ZendPrintLongToBuf(buf2+g.SizeOf("buf2")-1, s.h)
	}
	return strcoll(s1, s2)
}

/* }}} */

func PhpArrayReverseKeyCompareStringLocale(a any, b any) int {
	return PhpArrayKeyCompareStringLocale(b, a)
}

/* }}} */

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
	if first.u1.v.type_ == 13 {
		first = first.value.zv
	}
	if second.u1.v.type_ == 13 {
		second = second.value.zv
	}
	if zend.CompareFunction(&result, first, second) == zend.FAILURE {
		return 0
	}
	assert(result.u1.v.type_ == 4)
	if result.value.lval != 0 {
		if result.value.lval < 0 {
			return -1
		} else {
			return 1
		}
	} else {
		return 0
	}
}

/* }}} */

func PhpArrayReverseDataCompare(a any, b any) int { return PhpArrayDataCompare(a, b) * -1 }

/* }}} */

func PhpArrayDataCompareNumeric(a any, b any) int {
	var f *zend.Bucket
	var s *zend.Bucket
	var first *zend.Zval
	var second *zend.Zval
	f = (*zend.Bucket)(a)
	s = (*zend.Bucket)(b)
	first = &f.val
	second = &s.val
	if first.u1.v.type_ == 13 {
		first = first.value.zv
	}
	if second.u1.v.type_ == 13 {
		second = second.value.zv
	}
	return zend.NumericCompareFunction(first, second)
}

/* }}} */

func PhpArrayReverseDataCompareNumeric(a any, b any) int { return PhpArrayDataCompareNumeric(b, a) }

/* }}} */

func PhpArrayDataCompareStringCase(a any, b any) int {
	var f *zend.Bucket
	var s *zend.Bucket
	var first *zend.Zval
	var second *zend.Zval
	f = (*zend.Bucket)(a)
	s = (*zend.Bucket)(b)
	first = &f.val
	second = &s.val
	if first.u1.v.type_ == 13 {
		first = first.value.zv
	}
	if second.u1.v.type_ == 13 {
		second = second.value.zv
	}
	return zend.StringCaseCompareFunction(first, second)
}

/* }}} */

func PhpArrayReverseDataCompareStringCase(a any, b any) int {
	return PhpArrayDataCompareStringCase(b, a)
}

/* }}} */

func PhpArrayDataCompareString(a any, b any) int {
	var f *zend.Bucket
	var s *zend.Bucket
	var first *zend.Zval
	var second *zend.Zval
	f = (*zend.Bucket)(a)
	s = (*zend.Bucket)(b)
	first = &f.val
	second = &s.val
	if first.u1.v.type_ == 13 {
		first = first.value.zv
	}
	if second.u1.v.type_ == 13 {
		second = second.value.zv
	}
	return zend.StringCompareFunction(first, second)
}

/* }}} */

func PhpArrayReverseDataCompareString(a any, b any) int { return PhpArrayDataCompareString(b, a) }

/* }}} */

func PhpArrayNaturalGeneralCompare(a any, b any, fold_case int) int {
	var f *zend.Bucket = (*zend.Bucket)(a)
	var s *zend.Bucket = (*zend.Bucket)(b)
	var tmp_str1 *zend.ZendString
	var tmp_str2 *zend.ZendString
	var str1 *zend.ZendString = zend.ZvalGetTmpString(&f.val, &tmp_str1)
	var str2 *zend.ZendString = zend.ZvalGetTmpString(&s.val, &tmp_str2)
	var result int = StrnatcmpEx(str1.val, str1.len_, str2.val, str2.len_, fold_case)
	zend.ZendTmpStringRelease(tmp_str1)
	zend.ZendTmpStringRelease(tmp_str2)
	return result
}

/* }}} */

func PhpArrayNaturalCompare(a any, b any) int {
	return PhpArrayNaturalGeneralCompare(a, b, 0)
}

/* }}} */

func PhpArrayReverseNaturalCompare(a any, b any) int {
	return PhpArrayNaturalGeneralCompare(b, a, 0)
}

/* }}} */

func PhpArrayNaturalCaseCompare(a any, b any) int {
	return PhpArrayNaturalGeneralCompare(a, b, 1)
}

/* }}} */

func PhpArrayReverseNaturalCaseCompare(a any, b any) int {
	return PhpArrayNaturalGeneralCompare(b, a, 1)
}

/* }}} */

func PhpArrayDataCompareStringLocale(a any, b any) int {
	var f *zend.Bucket
	var s *zend.Bucket
	var first *zend.Zval
	var second *zend.Zval
	f = (*zend.Bucket)(a)
	s = (*zend.Bucket)(b)
	first = &f.val
	second = &s.val
	if first.u1.v.type_ == 13 {
		first = first.value.zv
	}
	if second.u1.v.type_ == 13 {
		second = second.value.zv
	}
	return zend.StringLocaleCompareFunction(first, second)
}

/* }}} */

func PhpArrayReverseDataCompareStringLocale(a any, b any) int {
	return PhpArrayDataCompareStringLocale(b, a)
}

/* }}} */

func PhpGetKeyCompareFunc(sort_type zend.ZendLong, reverse int) zend.CompareFuncT {
	switch sort_type & ^8 {
	case 1:
		if reverse != 0 {
			return PhpArrayReverseKeyCompareNumeric
		} else {
			return PhpArrayKeyCompareNumeric
		}
		break
	case 2:
		if (sort_type & 8) != 0 {
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
	case 6:
		if (sort_type & 8) != 0 {
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
	case 5:
		if reverse != 0 {
			return PhpArrayReverseKeyCompareStringLocale
		} else {
			return PhpArrayKeyCompareStringLocale
		}
		break
	case 0:

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

/* }}} */

func PhpGetDataCompareFunc(sort_type zend.ZendLong, reverse int) zend.CompareFuncT {
	switch sort_type & ^8 {
	case 1:
		if reverse != 0 {
			return PhpArrayReverseDataCompareNumeric
		} else {
			return PhpArrayDataCompareNumeric
		}
		break
	case 2:
		if (sort_type & 8) != 0 {
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
	case 6:
		if (sort_type & 8) != 0 {
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
	case 5:
		if reverse != 0 {
			return PhpArrayReverseDataCompareStringLocale
		} else {
			return PhpArrayDataCompareStringLocale
		}
		break
	case 0:

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

/* }}} */

func ZifKrsort(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var array *zend.Zval
	var sort_type zend.ZendLong = 0
	var cmp zend.CompareFuncT
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg
			if _arg.u1.v.type_ == 10 {
				_arg = &(*_arg).value.ref.val
			}
			var _zv *zend.Zval = _arg
			assert(_zv.u1.v.type_ != 10)
			var __zv *zend.Zval = _zv
			if __zv.u1.v.type_ == 7 {
				if zend.ZvalRefcountP(__zv) > 1 {
					if __zv.u1.v.type_flags != 0 {
						zend.ZvalDelrefP(__zv)
					}
					var __arr *zend.ZendArray = zend.ZendArrayDup(__zv.value.arr)
					var __z *zend.Zval = __zv
					__z.value.arr = __arr
					__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
				}
			}
			if zend.ZendParseArgArray(_arg, &array, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = 4
				break
			}
			_optional = 1
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgLong(_arg, &sort_type, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return_value.u1.type_info = 2
			return
		}
		break
	}
	cmp = PhpGetKeyCompareFunc(sort_type, 1)
	if zend.ZendHashSortEx(array.value.arr, zend.ZendSort, cmp, 0) == zend.FAILURE {
		return_value.u1.type_info = 2
		return
	}
	return_value.u1.type_info = 3
	return
}

/* }}} */

func ZifKsort(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var array *zend.Zval
	var sort_type zend.ZendLong = 0
	var cmp zend.CompareFuncT
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg
			if _arg.u1.v.type_ == 10 {
				_arg = &(*_arg).value.ref.val
			}
			var _zv *zend.Zval = _arg
			assert(_zv.u1.v.type_ != 10)
			var __zv *zend.Zval = _zv
			if __zv.u1.v.type_ == 7 {
				if zend.ZvalRefcountP(__zv) > 1 {
					if __zv.u1.v.type_flags != 0 {
						zend.ZvalDelrefP(__zv)
					}
					var __arr *zend.ZendArray = zend.ZendArrayDup(__zv.value.arr)
					var __z *zend.Zval = __zv
					__z.value.arr = __arr
					__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
				}
			}
			if zend.ZendParseArgArray(_arg, &array, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = 4
				break
			}
			_optional = 1
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgLong(_arg, &sort_type, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return_value.u1.type_info = 2
			return
		}
		break
	}
	cmp = PhpGetKeyCompareFunc(sort_type, 0)
	if zend.ZendHashSortEx(array.value.arr, zend.ZendSort, cmp, 0) == zend.FAILURE {
		return_value.u1.type_info = 2
		return
	}
	return_value.u1.type_info = 3
	return
}

/* }}} */

func PhpCountRecursive(ht *zend.HashTable) zend.ZendLong {
	var cnt zend.ZendLong = 0
	var element *zend.Zval
	if (zend.ZvalGcFlags(ht.gc.u.type_info) & 1 << 6) == 0 {
		if (zend.ZvalGcFlags(ht.gc.u.type_info) & 1 << 5) != 0 {
			core.PhpErrorDocref(nil, 1<<1, "recursion detected")
			return 0
		}
		ht.gc.u.type_info |= 1 << 5 << 0
	}
	cnt = zend.ZendArrayCount(ht)
	for {
		var __ht *zend.HashTable = ht
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if _z.u1.v.type_ == 0 {
				continue
			}
			element = _z
			if element.u1.v.type_ == 10 {
				element = &(*element).value.ref.val
			}
			if element.u1.v.type_ == 7 {
				cnt += PhpCountRecursive(element.value.arr)
			}
		}
		break
	}
	if (zend.ZvalGcFlags(ht.gc.u.type_info) & 1 << 6) == 0 {
		ht.gc.u.type_info &= ^(1 << 5 << 0)
	}
	return cnt
}

/* }}} */

func ZifCount(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var array *zend.Zval
	var mode zend.ZendLong = 0
	var cnt zend.ZendLong
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			zend.ZendParseArgZvalDeref(_arg, &array, 0)
			_optional = 1
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgLong(_arg, &mode, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
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
	switch array.u1.v.type_ {
	case 1:
		core.PhpErrorDocref(nil, 1<<1, "Parameter must be an array or an object that implements Countable")
		var __z *zend.Zval = return_value
		__z.value.lval = 0
		__z.u1.type_info = 4
		return
		break
	case 7:
		if mode != 1 {
			cnt = zend.ZendArrayCount(array.value.arr)
		} else {
			cnt = PhpCountRecursive(array.value.arr)
		}
		var __z *zend.Zval = return_value
		__z.value.lval = cnt
		__z.u1.type_info = 4
		return
		break
	case 8:
		var retval zend.Zval

		/* first, we check if the handler is defined */

		if array.value.obj.handlers.count_elements != nil {
			var __z *zend.Zval = return_value
			__z.value.lval = 1
			__z.u1.type_info = 4
			if zend.SUCCESS == array.value.obj.handlers.count_elements(array, &(*return_value).value.lval) {
				return
			}
			if zend.EG.exception != nil {
				return
			}
		}

		/* if not and the object implements Countable we call its count() method */

		if zend.InstanceofFunction(array.value.obj.ce, zend.ZendCeCountable) != 0 {
			zend.ZendCallMethod(array, nil, nil, "count", g.SizeOf("\"count\"")-1, &retval, 0, nil, nil)
			if retval.u1.v.type_ != 0 {
				var __z *zend.Zval = return_value
				__z.value.lval = zend.ZvalGetLong(&retval)
				__z.u1.type_info = 4
				zend.ZvalPtrDtor(&retval)
			}
			return
		}

		/* If There's no handler and it doesn't implement Countable then add a warning */

		core.PhpErrorDocref(nil, 1<<1, "Parameter must be an array or an object that implements Countable")
		var __z *zend.Zval = return_value
		__z.value.lval = 1
		__z.u1.type_info = 4
		return
		break
	default:
		core.PhpErrorDocref(nil, 1<<1, "Parameter must be an array or an object that implements Countable")
		var __z *zend.Zval = return_value
		__z.value.lval = 1
		__z.u1.type_info = 4
		return
		break
	}
}

/* }}} */

func PhpNatsort(execute_data *zend.ZendExecuteData, return_value *zend.Zval, fold_case int) {
	var array *zend.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg
			if _arg.u1.v.type_ == 10 {
				_arg = &(*_arg).value.ref.val
			}
			var _zv *zend.Zval = _arg
			assert(_zv.u1.v.type_ != 10)
			var __zv *zend.Zval = _zv
			if __zv.u1.v.type_ == 7 {
				if zend.ZvalRefcountP(__zv) > 1 {
					if __zv.u1.v.type_flags != 0 {
						zend.ZvalDelrefP(__zv)
					}
					var __arr *zend.ZendArray = zend.ZendArrayDup(__zv.value.arr)
					var __z *zend.Zval = __zv
					__z.value.arr = __arr
					__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
				}
			}
			if zend.ZendParseArgArray(_arg, &array, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
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
		if zend.ZendHashSortEx(array.value.arr, zend.ZendSort, PhpArrayNaturalCaseCompare, 0) == zend.FAILURE {
			return
		}
	} else {
		if zend.ZendHashSortEx(array.value.arr, zend.ZendSort, PhpArrayNaturalCompare, 0) == zend.FAILURE {
			return
		}
	}
	return_value.u1.type_info = 3
	return
}

/* }}} */

func ZifNatsort(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpNatsort(execute_data, return_value, 0)
}

/* }}} */

func ZifNatcasesort(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpNatsort(execute_data, return_value, 1)
}

/* }}} */

func ZifAsort(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var array *zend.Zval
	var sort_type zend.ZendLong = 0
	var cmp zend.CompareFuncT
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg
			if _arg.u1.v.type_ == 10 {
				_arg = &(*_arg).value.ref.val
			}
			var _zv *zend.Zval = _arg
			assert(_zv.u1.v.type_ != 10)
			var __zv *zend.Zval = _zv
			if __zv.u1.v.type_ == 7 {
				if zend.ZvalRefcountP(__zv) > 1 {
					if __zv.u1.v.type_flags != 0 {
						zend.ZvalDelrefP(__zv)
					}
					var __arr *zend.ZendArray = zend.ZendArrayDup(__zv.value.arr)
					var __z *zend.Zval = __zv
					__z.value.arr = __arr
					__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
				}
			}
			if zend.ZendParseArgArray(_arg, &array, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = 4
				break
			}
			_optional = 1
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgLong(_arg, &sort_type, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return_value.u1.type_info = 2
			return
		}
		break
	}
	cmp = PhpGetDataCompareFunc(sort_type, 0)
	if zend.ZendHashSortEx(array.value.arr, zend.ZendSort, cmp, 0) == zend.FAILURE {
		return_value.u1.type_info = 2
		return
	}
	return_value.u1.type_info = 3
	return
}

/* }}} */

func ZifArsort(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var array *zend.Zval
	var sort_type zend.ZendLong = 0
	var cmp zend.CompareFuncT
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg
			if _arg.u1.v.type_ == 10 {
				_arg = &(*_arg).value.ref.val
			}
			var _zv *zend.Zval = _arg
			assert(_zv.u1.v.type_ != 10)
			var __zv *zend.Zval = _zv
			if __zv.u1.v.type_ == 7 {
				if zend.ZvalRefcountP(__zv) > 1 {
					if __zv.u1.v.type_flags != 0 {
						zend.ZvalDelrefP(__zv)
					}
					var __arr *zend.ZendArray = zend.ZendArrayDup(__zv.value.arr)
					var __z *zend.Zval = __zv
					__z.value.arr = __arr
					__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
				}
			}
			if zend.ZendParseArgArray(_arg, &array, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = 4
				break
			}
			_optional = 1
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgLong(_arg, &sort_type, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return_value.u1.type_info = 2
			return
		}
		break
	}
	cmp = PhpGetDataCompareFunc(sort_type, 1)
	if zend.ZendHashSortEx(array.value.arr, zend.ZendSort, cmp, 0) == zend.FAILURE {
		return_value.u1.type_info = 2
		return
	}
	return_value.u1.type_info = 3
	return
}

/* }}} */

func ZifSort(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var array *zend.Zval
	var sort_type zend.ZendLong = 0
	var cmp zend.CompareFuncT
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg
			if _arg.u1.v.type_ == 10 {
				_arg = &(*_arg).value.ref.val
			}
			var _zv *zend.Zval = _arg
			assert(_zv.u1.v.type_ != 10)
			var __zv *zend.Zval = _zv
			if __zv.u1.v.type_ == 7 {
				if zend.ZvalRefcountP(__zv) > 1 {
					if __zv.u1.v.type_flags != 0 {
						zend.ZvalDelrefP(__zv)
					}
					var __arr *zend.ZendArray = zend.ZendArrayDup(__zv.value.arr)
					var __z *zend.Zval = __zv
					__z.value.arr = __arr
					__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
				}
			}
			if zend.ZendParseArgArray(_arg, &array, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = 4
				break
			}
			_optional = 1
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgLong(_arg, &sort_type, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return_value.u1.type_info = 2
			return
		}
		break
	}
	cmp = PhpGetDataCompareFunc(sort_type, 0)
	if zend.ZendHashSortEx(array.value.arr, zend.ZendSort, cmp, 1) == zend.FAILURE {
		return_value.u1.type_info = 2
		return
	}
	return_value.u1.type_info = 3
	return
}

/* }}} */

func ZifRsort(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var array *zend.Zval
	var sort_type zend.ZendLong = 0
	var cmp zend.CompareFuncT
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg
			if _arg.u1.v.type_ == 10 {
				_arg = &(*_arg).value.ref.val
			}
			var _zv *zend.Zval = _arg
			assert(_zv.u1.v.type_ != 10)
			var __zv *zend.Zval = _zv
			if __zv.u1.v.type_ == 7 {
				if zend.ZvalRefcountP(__zv) > 1 {
					if __zv.u1.v.type_flags != 0 {
						zend.ZvalDelrefP(__zv)
					}
					var __arr *zend.ZendArray = zend.ZendArrayDup(__zv.value.arr)
					var __z *zend.Zval = __zv
					__z.value.arr = __arr
					__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
				}
			}
			if zend.ZendParseArgArray(_arg, &array, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = 4
				break
			}
			_optional = 1
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgLong(_arg, &sort_type, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return_value.u1.type_info = 2
			return
		}
		break
	}
	cmp = PhpGetDataCompareFunc(sort_type, 1)
	if zend.ZendHashSortEx(array.value.arr, zend.ZendSort, cmp, 1) == zend.FAILURE {
		return_value.u1.type_info = 2
		return
	}
	return_value.u1.type_info = 3
	return
}

/* }}} */

func PhpArrayUserCompare(a any, b any) int {
	var f *zend.Bucket
	var s *zend.Bucket
	var args []zend.Zval
	var retval zend.Zval
	f = (*zend.Bucket)(a)
	s = (*zend.Bucket)(b)
	var _z1 *zend.Zval = &args[0]
	var _z2 *zend.Zval = &f.val
	var _gc *zend.ZendRefcounted = _z2.value.counted
	var _t uint32 = _z2.u1.type_info
	_z1.value.counted = _gc
	_z1.u1.type_info = _t
	if (_t & 0xff00) != 0 {
		zend.ZendGcAddref(&_gc.gc)
	}
	var _z1 *zend.Zval = &args[1]
	var _z2 *zend.Zval = &s.val
	var _gc *zend.ZendRefcounted = _z2.value.counted
	var _t uint32 = _z2.u1.type_info
	_z1.value.counted = _gc
	_z1.u1.type_info = _t
	if (_t & 0xff00) != 0 {
		zend.ZendGcAddref(&_gc.gc)
	}
	BasicGlobals.user_compare_fci.param_count = 2
	BasicGlobals.user_compare_fci.params = args
	BasicGlobals.user_compare_fci.retval = &retval
	BasicGlobals.user_compare_fci.no_separation = 0
	if zend.ZendCallFunction(&(BasicGlobals.GetUserCompareFci()), &(BasicGlobals.GetUserCompareFciCache())) == zend.SUCCESS && retval.u1.v.type_ != 0 {
		var ret zend.ZendLong = zend.ZvalGetLong(&retval)
		zend.ZvalPtrDtor(&retval)
		zend.ZvalPtrDtor(&args[1])
		zend.ZvalPtrDtor(&args[0])
		if ret != 0 {
			if ret < 0 {
				return -1
			} else {
				return 1
			}
		} else {
			return 0
		}
	} else {
		zend.ZvalPtrDtor(&args[1])
		zend.ZvalPtrDtor(&args[0])
		return 0
	}
}

/* }}} */

// #define PHP_ARRAY_CMP_FUNC_CHECK(func_name) if ( ! zend_is_callable ( * func_name , 0 , NULL ) ) { php_error_docref ( NULL , E_WARNING , "Invalid comparison function" ) ; BG ( user_compare_fci ) = old_user_compare_fci ; BG ( user_compare_fci_cache ) = old_user_compare_fci_cache ; RETURN_FALSE ; }

/* Clear FCI cache otherwise : for example the same or other array with
 * (partly) the same key values has been sorted with uasort() or
 * other sorting function the comparison is cached, however the name
 * of the function for comparison is not respected. see bug #28739 AND #33295
 *
 * Following defines will assist in backup / restore values. */

// #define PHP_ARRAY_CMP_FUNC_VARS       zend_fcall_info old_user_compare_fci ; zend_fcall_info_cache old_user_compare_fci_cache

// #define PHP_ARRAY_CMP_FUNC_BACKUP() old_user_compare_fci = BG ( user_compare_fci ) ; old_user_compare_fci_cache = BG ( user_compare_fci_cache ) ; BG ( user_compare_fci_cache ) = empty_fcall_info_cache ;

// #define PHP_ARRAY_CMP_FUNC_RESTORE() zend_release_fcall_info_cache ( & BG ( user_compare_fci_cache ) ) ; BG ( user_compare_fci ) = old_user_compare_fci ; BG ( user_compare_fci_cache ) = old_user_compare_fci_cache ;

func PhpUsort(execute_data *zend.ZendExecuteData, return_value *zend.Zval, compare_func zend.CompareFuncT, renumber zend.ZendBool) {
	var array *zend.Zval
	var arr *zend.ZendArray
	var retval zend.ZendBool
	var old_user_compare_fci zend.ZendFcallInfo
	var old_user_compare_fci_cache zend.ZendFcallInfoCache
	old_user_compare_fci = BasicGlobals.GetUserCompareFci()
	old_user_compare_fci_cache = BasicGlobals.GetUserCompareFciCache()
	BasicGlobals.SetUserCompareFciCache(zend.EmptyFcallInfoCache)
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 2
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg
			if _arg.u1.v.type_ == 10 {
				_arg = &(*_arg).value.ref.val
			}

			if zend.ZendParseArgArray(_arg, &array, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = 4
				break
			}
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgFunc(_arg, &(BasicGlobals.GetUserCompareFci()), &(BasicGlobals.GetUserCompareFciCache()), 0, &_error) == 0 {
				if _error == nil {
					_expected_type = zend.Z_EXPECTED_FUNC
					_error_code = 4
					break
				} else {
					_error_code = 2
					break
				}
			} else if _error != nil {
				zend.ZendWrongCallbackDeprecated(_i, _error)
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			zend.ZendReleaseFcallInfoCache(&(BasicGlobals.GetUserCompareFciCache()))
			BasicGlobals.SetUserCompareFci(old_user_compare_fci)
			BasicGlobals.SetUserCompareFciCache(old_user_compare_fci_cache)
			return
		}
		break
	}
	arr = array.value.arr
	if arr.nNumOfElements == 0 {
		zend.ZendReleaseFcallInfoCache(&(BasicGlobals.GetUserCompareFciCache()))
		BasicGlobals.SetUserCompareFci(old_user_compare_fci)
		BasicGlobals.SetUserCompareFciCache(old_user_compare_fci_cache)
		return_value.u1.type_info = 3
		return
	}

	/* Copy array, so the in-place modifications will not be visible to the callback function */

	arr = zend.ZendArrayDup(arr)
	retval = zend.ZendHashSortEx(arr, zend.ZendSort, compare_func, renumber) != zend.FAILURE
	var garbage zend.Zval
	var _z1 *zend.Zval = &garbage
	var _z2 *zend.Zval = array
	var _gc *zend.ZendRefcounted = _z2.value.counted
	var _t uint32 = _z2.u1.type_info
	_z1.value.counted = _gc
	_z1.u1.type_info = _t
	var __arr *zend.ZendArray = arr
	var __z *zend.Zval = array
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	zend.ZvalPtrDtor(&garbage)
	zend.ZendReleaseFcallInfoCache(&(BasicGlobals.GetUserCompareFciCache()))
	BasicGlobals.SetUserCompareFci(old_user_compare_fci)
	BasicGlobals.SetUserCompareFciCache(old_user_compare_fci_cache)
	if retval != 0 {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
	return
}

/* }}} */

func ZifUsort(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpUsort(execute_data, return_value, PhpArrayUserCompare, 1)
}

/* }}} */

func ZifUasort(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpUsort(execute_data, return_value, PhpArrayUserCompare, 0)
}

/* }}} */

func PhpArrayUserKeyCompare(a any, b any) int {
	var f *zend.Bucket
	var s *zend.Bucket
	var args []zend.Zval
	var retval zend.Zval
	var result zend.ZendLong
	f = (*zend.Bucket)(a)
	s = (*zend.Bucket)(b)
	if f.key == nil {
		var __z *zend.Zval = &args[0]
		__z.value.lval = f.h
		__z.u1.type_info = 4
	} else {
		var __z *zend.Zval = &args[0]
		var __s *zend.ZendString = f.key
		__z.value.str = __s
		if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
			__z.u1.type_info = 6
		} else {
			zend.ZendGcAddref(&__s.gc)
			__z.u1.type_info = 6 | 1<<0<<8
		}
	}
	if s.key == nil {
		var __z *zend.Zval = &args[1]
		__z.value.lval = s.h
		__z.u1.type_info = 4
	} else {
		var __z *zend.Zval = &args[1]
		var __s *zend.ZendString = s.key
		__z.value.str = __s
		if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
			__z.u1.type_info = 6
		} else {
			zend.ZendGcAddref(&__s.gc)
			__z.u1.type_info = 6 | 1<<0<<8
		}
	}
	BasicGlobals.user_compare_fci.param_count = 2
	BasicGlobals.user_compare_fci.params = args
	BasicGlobals.user_compare_fci.retval = &retval
	BasicGlobals.user_compare_fci.no_separation = 0
	if zend.ZendCallFunction(&(BasicGlobals.GetUserCompareFci()), &(BasicGlobals.GetUserCompareFciCache())) == zend.SUCCESS && retval.u1.v.type_ != 0 {
		result = zend.ZvalGetLong(&retval)
		zend.ZvalPtrDtor(&retval)
	} else {
		result = 0
	}
	zend.ZvalPtrDtor(&args[0])
	zend.ZvalPtrDtor(&args[1])
	if result != 0 {
		if result < 0 {
			return -1
		} else {
			return 1
		}
	} else {
		return 0
	}
}

/* }}} */

func ZifUksort(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpUsort(execute_data, return_value, PhpArrayUserKeyCompare, 0)
}

/* }}} */

func ZifEnd(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var array *zend.HashTable
	var entry *zend.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg
			if _arg.u1.v.type_ == 10 {
				_arg = &(*_arg).value.ref.val
			}
			var _zv *zend.Zval = _arg
			assert(_zv.u1.v.type_ != 10)
			var __zv *zend.Zval = _zv
			if __zv.u1.v.type_ == 7 {
				if zend.ZvalRefcountP(__zv) > 1 {
					if __zv.u1.v.type_flags != 0 {
						zend.ZvalDelrefP(__zv)
					}
					var __arr *zend.ZendArray = zend.ZendArrayDup(__zv.value.arr)
					var __z *zend.Zval = __zv
					__z.value.arr = __arr
					__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
				}
			}
			if zend.ZendParseArgArrayHt(_arg, &array, 0, 1, 1) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
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
	zend.ZendHashInternalPointerEndEx(array, &array.nInternalPointer)
	if execute_data.prev_execute_data == nil || (execute_data.prev_execute_data.func_.common.type_&1) != 0 || execute_data.prev_execute_data.opline.result_type != 0 {
		if g.Assign(&entry, zend.ZendHashGetCurrentDataEx(array, &array.nInternalPointer)) == nil {
			return_value.u1.type_info = 2
			return
		}
		if entry.u1.v.type_ == 13 {
			entry = entry.value.zv
		}
		var _z3 *zend.Zval = entry
		if (_z3.u1.type_info & 0xff00) != 0 {
			if (_z3.u1.type_info & 0xff) == 10 {
				_z3 = &(*_z3).value.ref.val
				if (_z3.u1.type_info & 0xff00) != 0 {
					zend.ZvalAddrefP(_z3)
				}
			} else {
				zend.ZvalAddrefP(_z3)
			}
		}
		var _z1 *zend.Zval = return_value
		var _z2 *zend.Zval = _z3
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
	}
}

/* }}} */

func ZifPrev(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var array *zend.HashTable
	var entry *zend.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg
			if _arg.u1.v.type_ == 10 {
				_arg = &(*_arg).value.ref.val
			}
			var _zv *zend.Zval = _arg
			assert(_zv.u1.v.type_ != 10)
			var __zv *zend.Zval = _zv
			if __zv.u1.v.type_ == 7 {
				if zend.ZvalRefcountP(__zv) > 1 {
					if __zv.u1.v.type_flags != 0 {
						zend.ZvalDelrefP(__zv)
					}
					var __arr *zend.ZendArray = zend.ZendArrayDup(__zv.value.arr)
					var __z *zend.Zval = __zv
					__z.value.arr = __arr
					__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
				}
			}
			if zend.ZendParseArgArrayHt(_arg, &array, 0, 1, 1) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
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
	zend.ZendHashMoveBackwardsEx(array, &array.nInternalPointer)
	if execute_data.prev_execute_data == nil || (execute_data.prev_execute_data.func_.common.type_&1) != 0 || execute_data.prev_execute_data.opline.result_type != 0 {
		if g.Assign(&entry, zend.ZendHashGetCurrentDataEx(array, &array.nInternalPointer)) == nil {
			return_value.u1.type_info = 2
			return
		}
		if entry.u1.v.type_ == 13 {
			entry = entry.value.zv
		}
		var _z3 *zend.Zval = entry
		if (_z3.u1.type_info & 0xff00) != 0 {
			if (_z3.u1.type_info & 0xff) == 10 {
				_z3 = &(*_z3).value.ref.val
				if (_z3.u1.type_info & 0xff00) != 0 {
					zend.ZvalAddrefP(_z3)
				}
			} else {
				zend.ZvalAddrefP(_z3)
			}
		}
		var _z1 *zend.Zval = return_value
		var _z2 *zend.Zval = _z3
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
	}
}

/* }}} */

func ZifNext(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var array *zend.HashTable
	var entry *zend.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg
			if _arg.u1.v.type_ == 10 {
				_arg = &(*_arg).value.ref.val
			}
			var _zv *zend.Zval = _arg
			assert(_zv.u1.v.type_ != 10)
			var __zv *zend.Zval = _zv
			if __zv.u1.v.type_ == 7 {
				if zend.ZvalRefcountP(__zv) > 1 {
					if __zv.u1.v.type_flags != 0 {
						zend.ZvalDelrefP(__zv)
					}
					var __arr *zend.ZendArray = zend.ZendArrayDup(__zv.value.arr)
					var __z *zend.Zval = __zv
					__z.value.arr = __arr
					__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
				}
			}
			if zend.ZendParseArgArrayHt(_arg, &array, 0, 1, 1) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
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
	zend.ZendHashMoveForwardEx(array, &array.nInternalPointer)
	if execute_data.prev_execute_data == nil || (execute_data.prev_execute_data.func_.common.type_&1) != 0 || execute_data.prev_execute_data.opline.result_type != 0 {
		if g.Assign(&entry, zend.ZendHashGetCurrentDataEx(array, &array.nInternalPointer)) == nil {
			return_value.u1.type_info = 2
			return
		}
		if entry.u1.v.type_ == 13 {
			entry = entry.value.zv
		}
		var _z3 *zend.Zval = entry
		if (_z3.u1.type_info & 0xff00) != 0 {
			if (_z3.u1.type_info & 0xff) == 10 {
				_z3 = &(*_z3).value.ref.val
				if (_z3.u1.type_info & 0xff00) != 0 {
					zend.ZvalAddrefP(_z3)
				}
			} else {
				zend.ZvalAddrefP(_z3)
			}
		}
		var _z1 *zend.Zval = return_value
		var _z2 *zend.Zval = _z3
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
	}
}

/* }}} */

func ZifReset(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var array *zend.HashTable
	var entry *zend.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg
			if _arg.u1.v.type_ == 10 {
				_arg = &(*_arg).value.ref.val
			}
			var _zv *zend.Zval = _arg
			assert(_zv.u1.v.type_ != 10)
			var __zv *zend.Zval = _zv
			if __zv.u1.v.type_ == 7 {
				if zend.ZvalRefcountP(__zv) > 1 {
					if __zv.u1.v.type_flags != 0 {
						zend.ZvalDelrefP(__zv)
					}
					var __arr *zend.ZendArray = zend.ZendArrayDup(__zv.value.arr)
					var __z *zend.Zval = __zv
					__z.value.arr = __arr
					__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
				}
			}
			if zend.ZendParseArgArrayHt(_arg, &array, 0, 1, 1) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
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
	zend.ZendHashInternalPointerResetEx(array, &array.nInternalPointer)
	if execute_data.prev_execute_data == nil || (execute_data.prev_execute_data.func_.common.type_&1) != 0 || execute_data.prev_execute_data.opline.result_type != 0 {
		if g.Assign(&entry, zend.ZendHashGetCurrentDataEx(array, &array.nInternalPointer)) == nil {
			return_value.u1.type_info = 2
			return
		}
		if entry.u1.v.type_ == 13 {
			entry = entry.value.zv
		}
		var _z3 *zend.Zval = entry
		if (_z3.u1.type_info & 0xff00) != 0 {
			if (_z3.u1.type_info & 0xff) == 10 {
				_z3 = &(*_z3).value.ref.val
				if (_z3.u1.type_info & 0xff00) != 0 {
					zend.ZvalAddrefP(_z3)
				}
			} else {
				zend.ZvalAddrefP(_z3)
			}
		}
		var _z1 *zend.Zval = return_value
		var _z2 *zend.Zval = _z3
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
	}
}

/* }}} */

func ZifCurrent(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var array *zend.HashTable
	var entry *zend.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgArrayHt(_arg, &array, 0, 1, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
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
	if g.Assign(&entry, zend.ZendHashGetCurrentDataEx(array, &array.nInternalPointer)) == nil {
		return_value.u1.type_info = 2
		return
	}
	if entry.u1.v.type_ == 13 {
		entry = entry.value.zv
	}
	var _z3 *zend.Zval = entry
	if (_z3.u1.type_info & 0xff00) != 0 {
		if (_z3.u1.type_info & 0xff) == 10 {
			_z3 = &(*_z3).value.ref.val
			if (_z3.u1.type_info & 0xff00) != 0 {
				zend.ZvalAddrefP(_z3)
			}
		} else {
			zend.ZvalAddrefP(_z3)
		}
	}
	var _z1 *zend.Zval = return_value
	var _z2 *zend.Zval = _z3
	var _gc *zend.ZendRefcounted = _z2.value.counted
	var _t uint32 = _z2.u1.type_info
	_z1.value.counted = _gc
	_z1.u1.type_info = _t
}

/* }}} */

func ZifKey(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var array *zend.HashTable
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgArrayHt(_arg, &array, 0, 1, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
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
	zend.ZendHashGetCurrentKeyZvalEx(array, return_value, &array.nInternalPointer)
}

/* }}} */

func ZifMin(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var argc int
	var args *zend.Zval = nil
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = -1
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
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
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
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
		if args[0].u1.v.type_ != 7 {
			core.PhpErrorDocref(nil, 1<<1, "When only one parameter is given, it must be an array")
			return_value.u1.type_info = 1
		} else {
			if g.Assign(&result, zend.ZendHashMinmax(args[0].value.arr, PhpArrayDataCompare, 0)) != nil {
				var _z3 *zend.Zval = result
				if (_z3.u1.type_info & 0xff00) != 0 {
					if (_z3.u1.type_info & 0xff) == 10 {
						_z3 = &(*_z3).value.ref.val
						if (_z3.u1.type_info & 0xff00) != 0 {
							zend.ZvalAddrefP(_z3)
						}
					} else {
						zend.ZvalAddrefP(_z3)
					}
				}
				var _z1 *zend.Zval = return_value
				var _z2 *zend.Zval = _z3
				var _gc *zend.ZendRefcounted = _z2.value.counted
				var _t uint32 = _z2.u1.type_info
				_z1.value.counted = _gc
				_z1.u1.type_info = _t
			} else {
				core.PhpErrorDocref(nil, 1<<1, "Array must contain at least one element")
				return_value.u1.type_info = 2
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
			if result.u1.v.type_ == 3 {
				min = &args[i]
			}
		}
		var _z1 *zend.Zval = return_value
		var _z2 *zend.Zval = min
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
		if (_t & 0xff00) != 0 {
			zend.ZendGcAddref(&_gc.gc)
		}
	}

	/* mixed min ( array $values ) */
}

/* }}} */

func ZifMax(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var args *zend.Zval = nil
	var argc int
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = -1
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
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
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
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
		if args[0].u1.v.type_ != 7 {
			core.PhpErrorDocref(nil, 1<<1, "When only one parameter is given, it must be an array")
			return_value.u1.type_info = 1
		} else {
			if g.Assign(&result, zend.ZendHashMinmax(args[0].value.arr, PhpArrayDataCompare, 1)) != nil {
				var _z3 *zend.Zval = result
				if (_z3.u1.type_info & 0xff00) != 0 {
					if (_z3.u1.type_info & 0xff) == 10 {
						_z3 = &(*_z3).value.ref.val
						if (_z3.u1.type_info & 0xff00) != 0 {
							zend.ZvalAddrefP(_z3)
						}
					} else {
						zend.ZvalAddrefP(_z3)
					}
				}
				var _z1 *zend.Zval = return_value
				var _z2 *zend.Zval = _z3
				var _gc *zend.ZendRefcounted = _z2.value.counted
				var _t uint32 = _z2.u1.type_info
				_z1.value.counted = _gc
				_z1.u1.type_info = _t
			} else {
				core.PhpErrorDocref(nil, 1<<1, "Array must contain at least one element")
				return_value.u1.type_info = 2
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
			if result.u1.v.type_ == 2 {
				max = &args[i]
			}
		}
		var _z1 *zend.Zval = return_value
		var _z2 *zend.Zval = max
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
		if (_t & 0xff00) != 0 {
			zend.ZendGcAddref(&_gc.gc)
		}
	}

	/* mixed max ( array $values ) */
}

/* }}} */

func PhpArrayWalk(array *zend.Zval, userdata *zend.Zval, recursive int) int {
	var args []zend.Zval
	var retval zend.Zval
	var zv *zend.Zval
	var target_hash *zend.HashTable = g.CondF(array.u1.v.type_ == 7, func() *zend.ZendArray { return array.value.arr }, func() __auto__ {
		if array.u1.v.type_ == 8 {
			return array.value.obj.handlers.get_properties(array)
		} else {
			return nil
		}
	})
	var pos zend.HashPosition
	var ht_iter uint32
	var result int = zend.SUCCESS

	/* Set up known arguments */

	&args[1].u1.type_info = 0
	if userdata != nil {
		var _z1 *zend.Zval = &args[2]
		var _z2 *zend.Zval = userdata
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
		if (_t & 0xff00) != 0 {
			zend.ZendGcAddref(&_gc.gc)
		}
	}
	BasicGlobals.array_walk_fci.retval = &retval
	if userdata != nil {
		BasicGlobals.array_walk_fci.param_count = 3
	} else {
		BasicGlobals.array_walk_fci.param_count = 2
	}
	BasicGlobals.array_walk_fci.params = args
	BasicGlobals.array_walk_fci.no_separation = 0
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

		if zv.u1.v.type_ == 13 {
			zv = zv.value.zv
			if zv.u1.v.type_ == 0 {
				zend.ZendHashMoveForwardEx(target_hash, &pos)
				continue
			}

			/* Add type source for property references. */

			if zv.u1.v.type_ != 10 && array.u1.v.type_ == 8 {
				var prop_info *zend.ZendPropertyInfo = zend.ZendGetTypedPropertyInfoForSlot(array.value.obj, zv)
				if prop_info != nil {
					var _ref *zend.ZendReference = (*zend.ZendReference)(zend._emalloc(g.SizeOf("zend_reference")))
					zend.ZendGcSetRefcount(&_ref.gc, 1)
					_ref.gc.u.type_info = 10
					var _z1 *zend.Zval = &_ref.val
					var _z2 *zend.Zval = zv
					var _gc *zend.ZendRefcounted = _z2.value.counted
					var _t uint32 = _z2.u1.type_info
					_z1.value.counted = _gc
					_z1.u1.type_info = _t
					_ref.sources.ptr = nil
					zv.value.ref = _ref
					zv.u1.type_info = 10 | 1<<0<<8
					zend.ZendRefAddTypeSource(&(zv.value.ref).sources, prop_info)
				}
			}

			/* Add type source for property references. */

		}

		/* Ensure the value is a reference. Otherwise the location of the value may be freed. */

		var __zv *zend.Zval = zv
		if __zv.u1.v.type_ != 10 {
			var _ref *zend.ZendReference = (*zend.ZendReference)(zend._emalloc(g.SizeOf("zend_reference")))
			zend.ZendGcSetRefcount(&_ref.gc, 1)
			_ref.gc.u.type_info = 10
			var _z1 *zend.Zval = &_ref.val
			var _z2 *zend.Zval = __zv
			var _gc *zend.ZendRefcounted = _z2.value.counted
			var _t uint32 = _z2.u1.type_info
			_z1.value.counted = _gc
			_z1.u1.type_info = _t
			_ref.sources.ptr = nil
			__zv.value.ref = _ref
			__zv.u1.type_info = 10 | 1<<0<<8
		}

		/* Retrieve key */

		zend.ZendHashGetCurrentKeyZvalEx(target_hash, &args[1], &pos)

		/* Move to next element already now -- this mirrors the approach used by foreach
		 * and ensures proper behavior with regard to modifications. */

		zend.ZendHashMoveForwardEx(target_hash, &pos)

		/* Back up hash position, as it may change */

		zend.EG.ht_iterators[ht_iter].pos = pos
		if recursive != 0 && &(*zv).value.ref.val.u1.v.type_ == 7 {
			var thash *zend.HashTable
			var orig_array_walk_fci zend.ZendFcallInfo
			var orig_array_walk_fci_cache zend.ZendFcallInfoCache
			var ref zend.Zval
			var _z1 *zend.Zval = &ref
			var _z2 *zend.Zval = zv
			var _gc *zend.ZendRefcounted = _z2.value.counted
			var _t uint32 = _z2.u1.type_info
			_z1.value.counted = _gc
			_z1.u1.type_info = _t
			if zv.u1.v.type_ == 10 {
				zv = &(*zv).value.ref.val
			}
			var _zv *zend.Zval = zv
			var _arr *zend.ZendArray = _zv.value.arr
			if zend.ZendGcRefcount(&_arr.gc) > 1 {
				if _zv.u1.v.type_flags != 0 {
					zend.ZendGcDelref(&_arr.gc)
				}
				var __arr *zend.ZendArray = zend.ZendArrayDup(_arr)
				var __z *zend.Zval = _zv
				__z.value.arr = __arr
				__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
			}
			thash = zv.value.arr
			if (zend.ZvalGcFlags(thash.gc.u.type_info) & 1 << 5) != 0 {
				core.PhpErrorDocref(nil, 1<<1, "recursion detected")
				result = zend.FAILURE
				break
			}

			/* backup the fcall info and cache */

			orig_array_walk_fci = BasicGlobals.GetArrayWalkFci()
			orig_array_walk_fci_cache = BasicGlobals.GetArrayWalkFciCache()
			zend.ZvalAddrefP(&ref)
			thash.gc.u.type_info |= 1 << 5 << 0
			result = PhpArrayWalk(zv, userdata, recursive)
			if &ref.value.ref.val.u1.v.type_ == 7 && thash == &ref.value.ref.val.value.arr {

				/* If the hashtable changed in the meantime, we'll "leak" this apply count
				 * increment -- our reference to thash is no longer valid. */

				thash.gc.u.type_info &= ^(1 << 5 << 0)

				/* If the hashtable changed in the meantime, we'll "leak" this apply count
				 * increment -- our reference to thash is no longer valid. */

			}
			zend.ZvalPtrDtor(&ref)

			/* restore the fcall info and cache */

			BasicGlobals.SetArrayWalkFci(orig_array_walk_fci)
			BasicGlobals.SetArrayWalkFciCache(orig_array_walk_fci_cache)
		} else {
			var _z1 *zend.Zval = &args[0]
			var _z2 *zend.Zval = zv
			var _gc *zend.ZendRefcounted = _z2.value.counted
			var _t uint32 = _z2.u1.type_info
			_z1.value.counted = _gc
			_z1.u1.type_info = _t
			if (_t & 0xff00) != 0 {
				zend.ZendGcAddref(&_gc.gc)
			}

			/* Call the userland function */

			result = zend.ZendCallFunction(&(BasicGlobals.GetArrayWalkFci()), &(BasicGlobals.GetArrayWalkFciCache()))
			if result == zend.SUCCESS {
				zend.ZvalPtrDtor(&retval)
			}
			zend.ZvalPtrDtor(&args[0])
		}
		if args[1].u1.v.type_ != 0 {
			zend.ZvalPtrDtor(&args[1])
			&args[1].u1.type_info = 0
		}
		if result == zend.FAILURE {
			break
		}

		/* Reload array and position -- both may have changed */

		if array.u1.v.type_ == 7 {
			pos = zend.ZendHashIteratorPosEx(ht_iter, array)
			target_hash = array.value.arr
		} else if array.u1.v.type_ == 8 {
			target_hash = array.value.obj.handlers.get_properties(&(*array))
			pos = zend.ZendHashIteratorPos(ht_iter, target_hash)
		} else {
			core.PhpErrorDocref(nil, 1<<1, "Iterated value is no longer an array or object")
			result = zend.FAILURE
			break
		}

		/* Reload array and position -- both may have changed */

		if zend.EG.exception != nil {
			break
		}
	}
	if userdata != nil {
		zend.ZvalPtrDtor(&args[2])
	}
	zend.ZendHashIteratorDel(ht_iter)
	return result
}

/* }}} */

func ZifArrayWalk(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var array *zend.Zval
	var userdata *zend.Zval = nil
	var orig_array_walk_fci zend.ZendFcallInfo
	var orig_array_walk_fci_cache zend.ZendFcallInfoCache
	orig_array_walk_fci = BasicGlobals.GetArrayWalkFci()
	orig_array_walk_fci_cache = BasicGlobals.GetArrayWalkFciCache()
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 3
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg
			if _arg.u1.v.type_ == 10 {
				_arg = &(*_arg).value.ref.val
			}
			var _zv *zend.Zval = _arg
			assert(_zv.u1.v.type_ != 10)
			var __zv *zend.Zval = _zv
			if __zv.u1.v.type_ == 7 {
				if zend.ZvalRefcountP(__zv) > 1 {
					if __zv.u1.v.type_flags != 0 {
						zend.ZvalDelrefP(__zv)
					}
					var __arr *zend.ZendArray = zend.ZendArrayDup(__zv.value.arr)
					var __z *zend.Zval = __zv
					__z.value.arr = __arr
					__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
				}
			}
			if zend.ZendParseArgArray(_arg, &array, 0, 1) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = 4
				break
			}
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgFunc(_arg, &(BasicGlobals.GetArrayWalkFci()), &(BasicGlobals.GetArrayWalkFciCache()), 0, &_error) == 0 {
				if _error == nil {
					_expected_type = zend.Z_EXPECTED_FUNC
					_error_code = 4
					break
				} else {
					_error_code = 2
					break
				}
			} else if _error != nil {
				zend.ZendWrongCallbackDeprecated(_i, _error)
			}
			_optional = 1
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			zend.ZendParseArgZvalDeref(_arg, &userdata, 0)
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			BasicGlobals.SetArrayWalkFci(orig_array_walk_fci)
			BasicGlobals.SetArrayWalkFciCache(orig_array_walk_fci_cache)
			return
		}
		break
	}
	PhpArrayWalk(array, userdata, 0)
	zend.ZendReleaseFcallInfoCache(&(BasicGlobals.GetArrayWalkFciCache()))
	BasicGlobals.SetArrayWalkFci(orig_array_walk_fci)
	BasicGlobals.SetArrayWalkFciCache(orig_array_walk_fci_cache)
	return_value.u1.type_info = 3
	return
}

/* }}} */

func ZifArrayWalkRecursive(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var array *zend.Zval
	var userdata *zend.Zval = nil
	var orig_array_walk_fci zend.ZendFcallInfo
	var orig_array_walk_fci_cache zend.ZendFcallInfoCache
	orig_array_walk_fci = BasicGlobals.GetArrayWalkFci()
	orig_array_walk_fci_cache = BasicGlobals.GetArrayWalkFciCache()
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 3
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg
			if _arg.u1.v.type_ == 10 {
				_arg = &(*_arg).value.ref.val
			}
			var _zv *zend.Zval = _arg
			assert(_zv.u1.v.type_ != 10)
			var __zv *zend.Zval = _zv
			if __zv.u1.v.type_ == 7 {
				if zend.ZvalRefcountP(__zv) > 1 {
					if __zv.u1.v.type_flags != 0 {
						zend.ZvalDelrefP(__zv)
					}
					var __arr *zend.ZendArray = zend.ZendArrayDup(__zv.value.arr)
					var __z *zend.Zval = __zv
					__z.value.arr = __arr
					__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
				}
			}
			if zend.ZendParseArgArray(_arg, &array, 0, 1) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = 4
				break
			}
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgFunc(_arg, &(BasicGlobals.GetArrayWalkFci()), &(BasicGlobals.GetArrayWalkFciCache()), 0, &_error) == 0 {
				if _error == nil {
					_expected_type = zend.Z_EXPECTED_FUNC
					_error_code = 4
					break
				} else {
					_error_code = 2
					break
				}
			} else if _error != nil {
				zend.ZendWrongCallbackDeprecated(_i, _error)
			}
			_optional = 1
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			zend.ZendParseArgZvalDeref(_arg, &userdata, 0)
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			BasicGlobals.SetArrayWalkFci(orig_array_walk_fci)
			BasicGlobals.SetArrayWalkFciCache(orig_array_walk_fci_cache)
			return
		}
		break
	}
	PhpArrayWalk(array, userdata, 1)
	zend.ZendReleaseFcallInfoCache(&(BasicGlobals.GetArrayWalkFciCache()))
	BasicGlobals.SetArrayWalkFci(orig_array_walk_fci)
	BasicGlobals.SetArrayWalkFciCache(orig_array_walk_fci_cache)
	return_value.u1.type_info = 3
	return
}

/* }}} */

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
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			zend.ZendParseArgZvalDeref(_arg, &value, 0)
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgArray(_arg, &array, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = 4
				break
			}
			_optional = 1
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgBool(_arg, &strict, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
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
		if value.u1.v.type_ == 4 {
			for {
				var __ht *zend.HashTable = array.value.arr
				var _p *zend.Bucket = __ht.arData
				var _end *zend.Bucket = _p + __ht.nNumUsed
				for ; _p != _end; _p++ {
					var _z *zend.Zval = &_p.val
					if _z.u1.v.type_ == 13 {
						_z = _z.value.zv
					}
					if _z.u1.v.type_ == 0 {
						continue
					}
					num_idx = _p.h
					str_idx = _p.key
					entry = _z
					if entry.u1.v.type_ == 10 {
						entry = &(*entry).value.ref.val
					}
					if entry.u1.v.type_ == 4 && entry.value.lval == value.value.lval {
						if behavior == 0 {
							return_value.u1.type_info = 3
							return
						} else {
							if str_idx != nil {
								var __z *zend.Zval = return_value
								var __s *zend.ZendString = str_idx
								__z.value.str = __s
								if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
									__z.u1.type_info = 6
								} else {
									zend.ZendGcAddref(&__s.gc)
									__z.u1.type_info = 6 | 1<<0<<8
								}
							} else {
								var __z *zend.Zval = return_value
								__z.value.lval = num_idx
								__z.u1.type_info = 4
							}
							return
						}
					}
				}
				break
			}
		} else {
			for {
				var __ht *zend.HashTable = array.value.arr
				var _p *zend.Bucket = __ht.arData
				var _end *zend.Bucket = _p + __ht.nNumUsed
				for ; _p != _end; _p++ {
					var _z *zend.Zval = &_p.val
					if _z.u1.v.type_ == 13 {
						_z = _z.value.zv
					}
					if _z.u1.v.type_ == 0 {
						continue
					}
					num_idx = _p.h
					str_idx = _p.key
					entry = _z
					if entry.u1.v.type_ == 10 {
						entry = &(*entry).value.ref.val
					}
					if zend.FastIsIdenticalFunction(value, entry) != 0 {
						if behavior == 0 {
							return_value.u1.type_info = 3
							return
						} else {
							if str_idx != nil {
								var __z *zend.Zval = return_value
								var __s *zend.ZendString = str_idx
								__z.value.str = __s
								if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
									__z.u1.type_info = 6
								} else {
									zend.ZendGcAddref(&__s.gc)
									__z.u1.type_info = 6 | 1<<0<<8
								}
							} else {
								var __z *zend.Zval = return_value
								__z.value.lval = num_idx
								__z.u1.type_info = 4
							}
							return
						}
					}
				}
				break
			}
		}
	} else {
		if value.u1.v.type_ == 4 {
			for {
				var __ht *zend.HashTable = array.value.arr
				var _p *zend.Bucket = __ht.arData
				var _end *zend.Bucket = _p + __ht.nNumUsed
				for ; _p != _end; _p++ {
					var _z *zend.Zval = &_p.val
					if _z.u1.v.type_ == 13 {
						_z = _z.value.zv
					}
					if _z.u1.v.type_ == 0 {
						continue
					}
					num_idx = _p.h
					str_idx = _p.key
					entry = _z
					if zend.FastEqualCheckLong(value, entry) != 0 {
						if behavior == 0 {
							return_value.u1.type_info = 3
							return
						} else {
							if str_idx != nil {
								var __z *zend.Zval = return_value
								var __s *zend.ZendString = str_idx
								__z.value.str = __s
								if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
									__z.u1.type_info = 6
								} else {
									zend.ZendGcAddref(&__s.gc)
									__z.u1.type_info = 6 | 1<<0<<8
								}
							} else {
								var __z *zend.Zval = return_value
								__z.value.lval = num_idx
								__z.u1.type_info = 4
							}
							return
						}
					}
				}
				break
			}
		} else if value.u1.v.type_ == 6 {
			for {
				var __ht *zend.HashTable = array.value.arr
				var _p *zend.Bucket = __ht.arData
				var _end *zend.Bucket = _p + __ht.nNumUsed
				for ; _p != _end; _p++ {
					var _z *zend.Zval = &_p.val
					if _z.u1.v.type_ == 13 {
						_z = _z.value.zv
					}
					if _z.u1.v.type_ == 0 {
						continue
					}
					num_idx = _p.h
					str_idx = _p.key
					entry = _z
					if zend.FastEqualCheckString(value, entry) != 0 {
						if behavior == 0 {
							return_value.u1.type_info = 3
							return
						} else {
							if str_idx != nil {
								var __z *zend.Zval = return_value
								var __s *zend.ZendString = str_idx
								__z.value.str = __s
								if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
									__z.u1.type_info = 6
								} else {
									zend.ZendGcAddref(&__s.gc)
									__z.u1.type_info = 6 | 1<<0<<8
								}
							} else {
								var __z *zend.Zval = return_value
								__z.value.lval = num_idx
								__z.u1.type_info = 4
							}
							return
						}
					}
				}
				break
			}
		} else {
			for {
				var __ht *zend.HashTable = array.value.arr
				var _p *zend.Bucket = __ht.arData
				var _end *zend.Bucket = _p + __ht.nNumUsed
				for ; _p != _end; _p++ {
					var _z *zend.Zval = &_p.val
					if _z.u1.v.type_ == 13 {
						_z = _z.value.zv
					}
					if _z.u1.v.type_ == 0 {
						continue
					}
					num_idx = _p.h
					str_idx = _p.key
					entry = _z
					if zend.FastEqualCheckFunction(value, entry) != 0 {
						if behavior == 0 {
							return_value.u1.type_info = 3
							return
						} else {
							if str_idx != nil {
								var __z *zend.Zval = return_value
								var __s *zend.ZendString = str_idx
								__z.value.str = __s
								if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
									__z.u1.type_info = 6
								} else {
									zend.ZendGcAddref(&__s.gc)
									__z.u1.type_info = 6 | 1<<0<<8
								}
							} else {
								var __z *zend.Zval = return_value
								__z.value.lval = num_idx
								__z.u1.type_info = 4
							}
							return
						}
					}
				}
				break
			}
		}
	}
	return_value.u1.type_info = 2
	return
}

/* }}} */

func ZifInArray(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpSearchArray(execute_data, return_value, 0)
}

/* }}} */

func ZifArraySearch(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpSearchArray(execute_data, return_value, 1)
}

/* }}} */

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
	if (charset[ch/(g.SizeOf("( charset ) [ 0 ]")*8)] >> (ch&g.SizeOf("( charset ) [ 0 ]")*8 - 1) & 1) == 0 {
		return 0
	}

	/* And these as the rest: [a-zA-Z0-9_\x7f-\xff] */

	if var_name_len > 1 {
		i = 1
		for {
			ch = uint32((*uint8)(var_name))[i]
			if (charset2[ch/(g.SizeOf("( charset2 ) [ 0 ]")*8)] >> (ch&g.SizeOf("( charset2 ) [ 0 ]")*8 - 1) & 1) == 0 {
				return 0
			}
			if g.PreInc(&i) >= var_name_len {
				break
			}
		}
	}
	return 1
}

/* }}} */

func PhpPrefixVarname(result *zend.Zval, prefix *zend.Zval, var_name *byte, var_name_len int, add_underscore zend.ZendBool) int {
	var __z *zend.Zval = result
	var __s *zend.ZendString = zend.ZendStringAlloc(prefix.value.str.len_+g.Cond(add_underscore != 0, 1, 0)+var_name_len, 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	memcpy(result.value.str.val, prefix.value.str.val, prefix.value.str.len_)
	if add_underscore != 0 {
		result.value.str.val[prefix.value.str.len_] = '_'
	}
	memcpy(result.value.str.val+prefix.value.str.len_+g.Cond(add_underscore != 0, 1, 0), var_name, var_name_len+1)
	return zend.SUCCESS
}

/* }}} */

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
			if _z.u1.v.type_ == 13 {
				_z = _z.value.zv
			}
			if _z.u1.v.type_ == 0 {
				continue
			}
			var_name = _p.key
			entry = _z
			if var_name == nil {
				continue
			}
			orig_var = zend.ZendHashFindEx(symbol_table, var_name, 1)
			if orig_var != nil {
				if orig_var.u1.v.type_ == 13 {
					orig_var = orig_var.value.zv
					if orig_var.u1.v.type_ == 0 {
						continue
					}
				}
				if PhpValidVarName(var_name.val, var_name.len_) == 0 {
					continue
				}
				if var_name.len_ == g.SizeOf("\"GLOBALS\"")-1 && !(memcmp(var_name.val, "GLOBALS", g.SizeOf("\"GLOBALS\"")-1)) {
					continue
				}
				if var_name.len_ == g.SizeOf("\"this\"")-1 && !(memcmp(var_name.val, "this", g.SizeOf("\"this\"")-1)) {
					zend.ZendThrowError(nil, "Cannot re-assign $this")
					return -1
				}
				if entry.u1.v.type_ == 10 {
					zend.ZvalAddrefP(entry)
				} else {
					var _z *zend.Zval = entry
					var _ref *zend.ZendReference = (*zend.ZendReference)(zend._emalloc(g.SizeOf("zend_reference")))
					zend.ZendGcSetRefcount(&_ref.gc, 2)
					_ref.gc.u.type_info = 10
					var _z1 *zend.Zval = &_ref.val
					var _z2 *zend.Zval = _z
					var _gc *zend.ZendRefcounted = _z2.value.counted
					var _t uint32 = _z2.u1.type_info
					_z1.value.counted = _gc
					_z1.u1.type_info = _t
					_ref.sources.ptr = nil
					_z.value.ref = _ref
					_z.u1.type_info = 10 | 1<<0<<8
				}
				zend.ZvalPtrDtor(orig_var)
				var __z *zend.Zval = orig_var
				__z.value.ref = entry.value.ref
				__z.u1.type_info = 10 | 1<<0<<8
				count++
			}
		}
		break
	}
	return count
}

/* }}} */

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
			if _z.u1.v.type_ == 13 {
				_z = _z.value.zv
			}
			if _z.u1.v.type_ == 0 {
				continue
			}
			var_name = _p.key
			entry = _z
			if var_name == nil {
				continue
			}
			orig_var = zend.ZendHashFindEx(symbol_table, var_name, 1)
			if orig_var != nil {
				if orig_var.u1.v.type_ == 13 {
					orig_var = orig_var.value.zv
					if orig_var.u1.v.type_ == 0 {
						continue
					}
				}
				if PhpValidVarName(var_name.val, var_name.len_) == 0 {
					continue
				}
				if var_name.len_ == g.SizeOf("\"GLOBALS\"")-1 && !(memcmp(var_name.val, "GLOBALS", g.SizeOf("\"GLOBALS\"")-1)) {
					continue
				}
				if var_name.len_ == g.SizeOf("\"this\"")-1 && !(memcmp(var_name.val, "this", g.SizeOf("\"this\"")-1)) {
					zend.ZendThrowError(nil, "Cannot re-assign $this")
					return -1
				}
				if entry.u1.v.type_ == 10 {
					entry = &(*entry).value.ref.val
				}
				for {
					if entry.u1.v.type_flags != 0 {
						zend.ZvalAddrefP(entry)
					}
					for {
						var _zv *zend.Zval = orig_var
						if _zv.u1.v.type_ == 10 {
							var ref *zend.ZendReference = _zv.value.ref
							if ref.sources.ptr != nil {
								zend.ZendTryAssignTypedRefZvalEx(ref, entry, 0)
								break
							}
							_zv = &ref.val
						}
						zend.ZvalPtrDtor(_zv)
						var _z1 *zend.Zval = _zv
						var _z2 *zend.Zval = entry
						var _gc *zend.ZendRefcounted = _z2.value.counted
						var _t uint32 = _z2.u1.type_info
						_z1.value.counted = _gc
						_z1.u1.type_info = _t
						break
					}
					break
				}
				if zend.EG.exception != nil {
					return -1
				}
				count++
			}
		}
		break
	}
	return count
}

/* }}} */

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
			if _z.u1.v.type_ == 13 {
				_z = _z.value.zv
			}
			if _z.u1.v.type_ == 0 {
				continue
			}
			var_name = _p.key
			entry = _z
			if var_name == nil {
				continue
			}
			if PhpValidVarName(var_name.val, var_name.len_) == 0 {
				continue
			}
			if var_name.len_ == g.SizeOf("\"this\"")-1 && !(memcmp(var_name.val, "this", g.SizeOf("\"this\"")-1)) {
				zend.ZendThrowError(nil, "Cannot re-assign $this")
				return -1
			}
			orig_var = zend.ZendHashFindEx(symbol_table, var_name, 1)
			if orig_var != nil {
				if orig_var.u1.v.type_ == 13 {
					orig_var = orig_var.value.zv
				}
				if var_name.len_ == g.SizeOf("\"GLOBALS\"")-1 && !(memcmp(var_name.val, "GLOBALS", g.SizeOf("\"GLOBALS\"")-1)) {
					continue
				}
				if entry.u1.v.type_ == 10 {
					zend.ZvalAddrefP(entry)
				} else {
					var _z *zend.Zval = entry
					var _ref *zend.ZendReference = (*zend.ZendReference)(zend._emalloc(g.SizeOf("zend_reference")))
					zend.ZendGcSetRefcount(&_ref.gc, 2)
					_ref.gc.u.type_info = 10
					var _z1 *zend.Zval = &_ref.val
					var _z2 *zend.Zval = _z
					var _gc *zend.ZendRefcounted = _z2.value.counted
					var _t uint32 = _z2.u1.type_info
					_z1.value.counted = _gc
					_z1.u1.type_info = _t
					_ref.sources.ptr = nil
					_z.value.ref = _ref
					_z.u1.type_info = 10 | 1<<0<<8
				}
				zend.ZvalPtrDtor(orig_var)
				var __z *zend.Zval = orig_var
				__z.value.ref = entry.value.ref
				__z.u1.type_info = 10 | 1<<0<<8
			} else {
				if entry.u1.v.type_ == 10 {
					zend.ZvalAddrefP(entry)
				} else {
					var _z *zend.Zval = entry
					var _ref *zend.ZendReference = (*zend.ZendReference)(zend._emalloc(g.SizeOf("zend_reference")))
					zend.ZendGcSetRefcount(&_ref.gc, 2)
					_ref.gc.u.type_info = 10
					var _z1 *zend.Zval = &_ref.val
					var _z2 *zend.Zval = _z
					var _gc *zend.ZendRefcounted = _z2.value.counted
					var _t uint32 = _z2.u1.type_info
					_z1.value.counted = _gc
					_z1.u1.type_info = _t
					_ref.sources.ptr = nil
					_z.value.ref = _ref
					_z.u1.type_info = 10 | 1<<0<<8
				}
				zend.ZendHashAddNew(symbol_table, var_name, entry)
			}
			count++
		}
		break
	}
	return count
}

/* }}} */

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
			if _z.u1.v.type_ == 13 {
				_z = _z.value.zv
			}
			if _z.u1.v.type_ == 0 {
				continue
			}
			var_name = _p.key
			entry = _z
			if var_name == nil {
				continue
			}
			if PhpValidVarName(var_name.val, var_name.len_) == 0 {
				continue
			}
			if var_name.len_ == g.SizeOf("\"this\"")-1 && !(memcmp(var_name.val, "this", g.SizeOf("\"this\"")-1)) {
				zend.ZendThrowError(nil, "Cannot re-assign $this")
				return -1
			}
			orig_var = zend.ZendHashFindEx(symbol_table, var_name, 1)
			if orig_var != nil {
				if orig_var.u1.v.type_ == 13 {
					orig_var = orig_var.value.zv
				}
				if var_name.len_ == g.SizeOf("\"GLOBALS\"")-1 && !(memcmp(var_name.val, "GLOBALS", g.SizeOf("\"GLOBALS\"")-1)) {
					continue
				}
				if entry.u1.v.type_ == 10 {
					entry = &(*entry).value.ref.val
				}
				for {
					if entry.u1.v.type_flags != 0 {
						zend.ZvalAddrefP(entry)
					}
					for {
						var _zv *zend.Zval = orig_var
						if _zv.u1.v.type_ == 10 {
							var ref *zend.ZendReference = _zv.value.ref
							if ref.sources.ptr != nil {
								zend.ZendTryAssignTypedRefZvalEx(ref, entry, 0)
								break
							}
							_zv = &ref.val
						}
						zend.ZvalPtrDtor(_zv)
						var _z1 *zend.Zval = _zv
						var _z2 *zend.Zval = entry
						var _gc *zend.ZendRefcounted = _z2.value.counted
						var _t uint32 = _z2.u1.type_info
						_z1.value.counted = _gc
						_z1.u1.type_info = _t
						break
					}
					break
				}
				if zend.EG.exception != nil {
					return -1
				}
			} else {
				if entry.u1.v.type_ == 10 {
					entry = &(*entry).value.ref.val
				}
				if entry.u1.v.type_flags != 0 {
					zend.ZvalAddrefP(entry)
				}
				zend.ZendHashAddNew(symbol_table, var_name, entry)
			}
			count++
		}
		break
	}
	return count
}

/* }}} */

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
			if _z.u1.v.type_ == 13 {
				_z = _z.value.zv
			}
			if _z.u1.v.type_ == 0 {
				continue
			}
			var_name = _p.key
			entry = _z
			if var_name == nil {
				continue
			}
			orig_var = zend.ZendHashFindEx(symbol_table, var_name, 1)
			if orig_var != nil {
				if orig_var.u1.v.type_ == 13 {
					orig_var = orig_var.value.zv
					if orig_var.u1.v.type_ == 0 {
						if entry.u1.v.type_ == 10 {
							zend.ZvalAddrefP(entry)
						} else {
							var _z *zend.Zval = entry
							var _ref *zend.ZendReference = (*zend.ZendReference)(zend._emalloc(g.SizeOf("zend_reference")))
							zend.ZendGcSetRefcount(&_ref.gc, 2)
							_ref.gc.u.type_info = 10
							var _z1 *zend.Zval = &_ref.val
							var _z2 *zend.Zval = _z
							var _gc *zend.ZendRefcounted = _z2.value.counted
							var _t uint32 = _z2.u1.type_info
							_z1.value.counted = _gc
							_z1.u1.type_info = _t
							_ref.sources.ptr = nil
							_z.value.ref = _ref
							_z.u1.type_info = 10 | 1<<0<<8
						}
						var __z *zend.Zval = orig_var
						__z.value.ref = entry.value.ref
						__z.u1.type_info = 10 | 1<<0<<8
						count++
						continue
					}
				}
				PhpPrefixVarname(&final_name, prefix, var_name.val, var_name.len_, 1)
				if PhpValidVarName(final_name.value.str.val, final_name.value.str.len_) != 0 {
					if final_name.value.str.len_ == g.SizeOf("\"this\"")-1 && !(memcmp(final_name.value.str.val, "this", g.SizeOf("\"this\"")-1)) {
						zend.ZendThrowError(nil, "Cannot re-assign $this")
						return -1
					} else {
						if entry.u1.v.type_ == 10 {
							zend.ZvalAddrefP(entry)
						} else {
							var _z *zend.Zval = entry
							var _ref *zend.ZendReference = (*zend.ZendReference)(zend._emalloc(g.SizeOf("zend_reference")))
							zend.ZendGcSetRefcount(&_ref.gc, 2)
							_ref.gc.u.type_info = 10
							var _z1 *zend.Zval = &_ref.val
							var _z2 *zend.Zval = _z
							var _gc *zend.ZendRefcounted = _z2.value.counted
							var _t uint32 = _z2.u1.type_info
							_z1.value.counted = _gc
							_z1.u1.type_info = _t
							_ref.sources.ptr = nil
							_z.value.ref = _ref
							_z.u1.type_info = 10 | 1<<0<<8
						}
						if g.Assign(&orig_var, zend.ZendHashFind(symbol_table, final_name.value.str)) != nil {
							if orig_var.u1.v.type_ == 13 {
								orig_var = orig_var.value.zv
							}
							zend.ZvalPtrDtor(orig_var)
							var __z *zend.Zval = orig_var
							__z.value.ref = entry.value.ref
							__z.u1.type_info = 10 | 1<<0<<8
						} else {
							zend.ZendHashAddNew(symbol_table, final_name.value.str, entry)
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

/* }}} */

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
			if _z.u1.v.type_ == 13 {
				_z = _z.value.zv
			}
			if _z.u1.v.type_ == 0 {
				continue
			}
			var_name = _p.key
			entry = _z
			if var_name == nil {
				continue
			}
			orig_var = zend.ZendHashFindEx(symbol_table, var_name, 1)
			if orig_var != nil {
				if orig_var.u1.v.type_ == 13 {
					orig_var = orig_var.value.zv
					if orig_var.u1.v.type_ == 0 {
						var _z3 *zend.Zval = entry
						if (_z3.u1.type_info & 0xff00) != 0 {
							if (_z3.u1.type_info & 0xff) == 10 {
								_z3 = &(*_z3).value.ref.val
								if (_z3.u1.type_info & 0xff00) != 0 {
									zend.ZvalAddrefP(_z3)
								}
							} else {
								zend.ZvalAddrefP(_z3)
							}
						}
						var _z1 *zend.Zval = orig_var
						var _z2 *zend.Zval = _z3
						var _gc *zend.ZendRefcounted = _z2.value.counted
						var _t uint32 = _z2.u1.type_info
						_z1.value.counted = _gc
						_z1.u1.type_info = _t
						count++
						continue
					}
				}
				PhpPrefixVarname(&final_name, prefix, var_name.val, var_name.len_, 1)
				if PhpValidVarName(final_name.value.str.val, final_name.value.str.len_) != 0 {
					if final_name.value.str.len_ == g.SizeOf("\"this\"")-1 && !(memcmp(final_name.value.str.val, "this", g.SizeOf("\"this\"")-1)) {
						zend.ZendThrowError(nil, "Cannot re-assign $this")
						return -1
					} else {
						if entry.u1.v.type_ == 10 {
							entry = &(*entry).value.ref.val
						}
						if g.Assign(&orig_var, zend.ZendHashFind(symbol_table, final_name.value.str)) != nil {
							if orig_var.u1.v.type_ == 13 {
								orig_var = orig_var.value.zv
							}
							for {
								if entry.u1.v.type_flags != 0 {
									zend.ZvalAddrefP(entry)
								}
								for {
									var _zv *zend.Zval = orig_var
									if _zv.u1.v.type_ == 10 {
										var ref *zend.ZendReference = _zv.value.ref
										if ref.sources.ptr != nil {
											zend.ZendTryAssignTypedRefZvalEx(ref, entry, 0)
											break
										}
										_zv = &ref.val
									}
									zend.ZvalPtrDtor(_zv)
									var _z1 *zend.Zval = _zv
									var _z2 *zend.Zval = entry
									var _gc *zend.ZendRefcounted = _z2.value.counted
									var _t uint32 = _z2.u1.type_info
									_z1.value.counted = _gc
									_z1.u1.type_info = _t
									break
								}
								break
							}
							if zend.EG.exception != nil {
								zend.ZendStringReleaseEx(final_name.value.str, 0)
								return -1
							}
						} else {
							if entry.u1.v.type_flags != 0 {
								zend.ZvalAddrefP(entry)
							}
							zend.ZendHashAddNew(symbol_table, final_name.value.str, entry)
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

/* }}} */

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
			if _z.u1.v.type_ == 13 {
				_z = _z.value.zv
			}
			if _z.u1.v.type_ == 0 {
				continue
			}
			var_name = _p.key
			entry = _z
			if var_name == nil {
				continue
			}
			if var_name.len_ == 0 {
				continue
			}
			orig_var = zend.ZendHashFindEx(symbol_table, var_name, 1)
			if orig_var != nil {
				if orig_var.u1.v.type_ == 13 {
					orig_var = orig_var.value.zv
					if orig_var.u1.v.type_ == 0 {
						if entry.u1.v.type_ == 10 {
							zend.ZvalAddrefP(entry)
						} else {
							var _z *zend.Zval = entry
							var _ref *zend.ZendReference = (*zend.ZendReference)(zend._emalloc(g.SizeOf("zend_reference")))
							zend.ZendGcSetRefcount(&_ref.gc, 2)
							_ref.gc.u.type_info = 10
							var _z1 *zend.Zval = &_ref.val
							var _z2 *zend.Zval = _z
							var _gc *zend.ZendRefcounted = _z2.value.counted
							var _t uint32 = _z2.u1.type_info
							_z1.value.counted = _gc
							_z1.u1.type_info = _t
							_ref.sources.ptr = nil
							_z.value.ref = _ref
							_z.u1.type_info = 10 | 1<<0<<8
						}
						var __z *zend.Zval = orig_var
						__z.value.ref = entry.value.ref
						__z.u1.type_info = 10 | 1<<0<<8
						count++
						continue
					}
				}
			prefix:
				PhpPrefixVarname(&final_name, prefix, var_name.val, var_name.len_, 1)
				if PhpValidVarName(final_name.value.str.val, final_name.value.str.len_) != 0 {
					if final_name.value.str.len_ == g.SizeOf("\"this\"")-1 && !(memcmp(final_name.value.str.val, "this", g.SizeOf("\"this\"")-1)) {
						zend.ZendThrowError(nil, "Cannot re-assign $this")
						return -1
					} else {
						if entry.u1.v.type_ == 10 {
							zend.ZvalAddrefP(entry)
						} else {
							var _z *zend.Zval = entry
							var _ref *zend.ZendReference = (*zend.ZendReference)(zend._emalloc(g.SizeOf("zend_reference")))
							zend.ZendGcSetRefcount(&_ref.gc, 2)
							_ref.gc.u.type_info = 10
							var _z1 *zend.Zval = &_ref.val
							var _z2 *zend.Zval = _z
							var _gc *zend.ZendRefcounted = _z2.value.counted
							var _t uint32 = _z2.u1.type_info
							_z1.value.counted = _gc
							_z1.u1.type_info = _t
							_ref.sources.ptr = nil
							_z.value.ref = _ref
							_z.u1.type_info = 10 | 1<<0<<8
						}
						if g.Assign(&orig_var, zend.ZendHashFind(symbol_table, final_name.value.str)) != nil {
							if orig_var.u1.v.type_ == 13 {
								orig_var = orig_var.value.zv
							}
							zend.ZvalPtrDtor(orig_var)
							var __z *zend.Zval = orig_var
							__z.value.ref = entry.value.ref
							__z.u1.type_info = 10 | 1<<0<<8
						} else {
							zend.ZendHashAddNew(symbol_table, final_name.value.str, entry)
						}
						count++
					}
				}
				zend.ZvalPtrDtorStr(&final_name)
			} else {
				if PhpValidVarName(var_name.val, var_name.len_) == 0 {
					continue
				}
				if var_name.len_ == g.SizeOf("\"this\"")-1 && !(memcmp(var_name.val, "this", g.SizeOf("\"this\"")-1)) {
					goto prefix
				}
				if entry.u1.v.type_ == 10 {
					zend.ZvalAddrefP(entry)
				} else {
					var _z *zend.Zval = entry
					var _ref *zend.ZendReference = (*zend.ZendReference)(zend._emalloc(g.SizeOf("zend_reference")))
					zend.ZendGcSetRefcount(&_ref.gc, 2)
					_ref.gc.u.type_info = 10
					var _z1 *zend.Zval = &_ref.val
					var _z2 *zend.Zval = _z
					var _gc *zend.ZendRefcounted = _z2.value.counted
					var _t uint32 = _z2.u1.type_info
					_z1.value.counted = _gc
					_z1.u1.type_info = _t
					_ref.sources.ptr = nil
					_z.value.ref = _ref
					_z.u1.type_info = 10 | 1<<0<<8
				}
				zend.ZendHashAddNew(symbol_table, var_name, entry)
				count++
			}
		}
		break
	}
	return count
}

/* }}} */

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
			if _z.u1.v.type_ == 13 {
				_z = _z.value.zv
			}
			if _z.u1.v.type_ == 0 {
				continue
			}
			var_name = _p.key
			entry = _z
			if var_name == nil {
				continue
			}
			if var_name.len_ == 0 {
				continue
			}
			orig_var = zend.ZendHashFindEx(symbol_table, var_name, 1)
			if orig_var != nil {
				if orig_var.u1.v.type_ == 13 {
					orig_var = orig_var.value.zv
					if orig_var.u1.v.type_ == 0 {
						var _z3 *zend.Zval = entry
						if (_z3.u1.type_info & 0xff00) != 0 {
							if (_z3.u1.type_info & 0xff) == 10 {
								_z3 = &(*_z3).value.ref.val
								if (_z3.u1.type_info & 0xff00) != 0 {
									zend.ZvalAddrefP(_z3)
								}
							} else {
								zend.ZvalAddrefP(_z3)
							}
						}
						var _z1 *zend.Zval = orig_var
						var _z2 *zend.Zval = _z3
						var _gc *zend.ZendRefcounted = _z2.value.counted
						var _t uint32 = _z2.u1.type_info
						_z1.value.counted = _gc
						_z1.u1.type_info = _t
						count++
						continue
					}
				}
			prefix:
				PhpPrefixVarname(&final_name, prefix, var_name.val, var_name.len_, 1)
				if PhpValidVarName(final_name.value.str.val, final_name.value.str.len_) != 0 {
					if final_name.value.str.len_ == g.SizeOf("\"this\"")-1 && !(memcmp(final_name.value.str.val, "this", g.SizeOf("\"this\"")-1)) {
						zend.ZendThrowError(nil, "Cannot re-assign $this")
						return -1
					} else {
						if entry.u1.v.type_ == 10 {
							entry = &(*entry).value.ref.val
						}
						if g.Assign(&orig_var, zend.ZendHashFind(symbol_table, final_name.value.str)) != nil {
							if orig_var.u1.v.type_ == 13 {
								orig_var = orig_var.value.zv
							}
							for {
								if entry.u1.v.type_flags != 0 {
									zend.ZvalAddrefP(entry)
								}
								for {
									var _zv *zend.Zval = orig_var
									if _zv.u1.v.type_ == 10 {
										var ref *zend.ZendReference = _zv.value.ref
										if ref.sources.ptr != nil {
											zend.ZendTryAssignTypedRefZvalEx(ref, entry, 0)
											break
										}
										_zv = &ref.val
									}
									zend.ZvalPtrDtor(_zv)
									var _z1 *zend.Zval = _zv
									var _z2 *zend.Zval = entry
									var _gc *zend.ZendRefcounted = _z2.value.counted
									var _t uint32 = _z2.u1.type_info
									_z1.value.counted = _gc
									_z1.u1.type_info = _t
									break
								}
								break
							}
							if zend.EG.exception != nil {
								zend.ZendStringReleaseEx(final_name.value.str, 0)
								return -1
							}
						} else {
							if entry.u1.v.type_flags != 0 {
								zend.ZvalAddrefP(entry)
							}
							zend.ZendHashAddNew(symbol_table, final_name.value.str, entry)
						}
						count++
					}
				}
				zend.ZvalPtrDtorStr(&final_name)
			} else {
				if PhpValidVarName(var_name.val, var_name.len_) == 0 {
					continue
				}
				if var_name.len_ == g.SizeOf("\"this\"")-1 && !(memcmp(var_name.val, "this", g.SizeOf("\"this\"")-1)) {
					goto prefix
				}
				if entry.u1.v.type_ == 10 {
					entry = &(*entry).value.ref.val
				}
				if entry.u1.v.type_flags != 0 {
					zend.ZvalAddrefP(entry)
				}
				zend.ZendHashAddNew(symbol_table, var_name, entry)
				count++
			}
		}
		break
	}
	return count
}

/* }}} */

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
			if _z.u1.v.type_ == 13 {
				_z = _z.value.zv
			}
			if _z.u1.v.type_ == 0 {
				continue
			}
			num_key = _p.h
			var_name = _p.key
			entry = _z
			if var_name != nil {
				if var_name.len_ == 0 {
					continue
				}
				PhpPrefixVarname(&final_name, prefix, var_name.val, var_name.len_, 1)
			} else {
				var str *zend.ZendString = zend.ZendLongToStr(num_key)
				PhpPrefixVarname(&final_name, prefix, str.val, str.len_, 1)
				zend.ZendStringReleaseEx(str, 0)
			}
			if PhpValidVarName(final_name.value.str.val, final_name.value.str.len_) != 0 {
				if final_name.value.str.len_ == g.SizeOf("\"this\"")-1 && !(memcmp(final_name.value.str.val, "this", g.SizeOf("\"this\"")-1)) {
					zend.ZendThrowError(nil, "Cannot re-assign $this")
					return -1
				} else {
					if entry.u1.v.type_ == 10 {
						zend.ZvalAddrefP(entry)
					} else {
						var _z *zend.Zval = entry
						var _ref *zend.ZendReference = (*zend.ZendReference)(zend._emalloc(g.SizeOf("zend_reference")))
						zend.ZendGcSetRefcount(&_ref.gc, 2)
						_ref.gc.u.type_info = 10
						var _z1 *zend.Zval = &_ref.val
						var _z2 *zend.Zval = _z
						var _gc *zend.ZendRefcounted = _z2.value.counted
						var _t uint32 = _z2.u1.type_info
						_z1.value.counted = _gc
						_z1.u1.type_info = _t
						_ref.sources.ptr = nil
						_z.value.ref = _ref
						_z.u1.type_info = 10 | 1<<0<<8
					}
					if g.Assign(&orig_var, zend.ZendHashFind(symbol_table, final_name.value.str)) != nil {
						if orig_var.u1.v.type_ == 13 {
							orig_var = orig_var.value.zv
						}
						zend.ZvalPtrDtor(orig_var)
						var __z *zend.Zval = orig_var
						__z.value.ref = entry.value.ref
						__z.u1.type_info = 10 | 1<<0<<8
					} else {
						zend.ZendHashAddNew(symbol_table, final_name.value.str, entry)
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

/* }}} */

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
			if _z.u1.v.type_ == 13 {
				_z = _z.value.zv
			}
			if _z.u1.v.type_ == 0 {
				continue
			}
			num_key = _p.h
			var_name = _p.key
			entry = _z
			if var_name != nil {
				if var_name.len_ == 0 {
					continue
				}
				PhpPrefixVarname(&final_name, prefix, var_name.val, var_name.len_, 1)
			} else {
				var str *zend.ZendString = zend.ZendLongToStr(num_key)
				PhpPrefixVarname(&final_name, prefix, str.val, str.len_, 1)
				zend.ZendStringReleaseEx(str, 0)
			}
			if PhpValidVarName(final_name.value.str.val, final_name.value.str.len_) != 0 {
				if final_name.value.str.len_ == g.SizeOf("\"this\"")-1 && !(memcmp(final_name.value.str.val, "this", g.SizeOf("\"this\"")-1)) {
					zend.ZendThrowError(nil, "Cannot re-assign $this")
					return -1
				} else {
					if entry.u1.v.type_ == 10 {
						entry = &(*entry).value.ref.val
					}
					if g.Assign(&orig_var, zend.ZendHashFind(symbol_table, final_name.value.str)) != nil {
						if orig_var.u1.v.type_ == 13 {
							orig_var = orig_var.value.zv
						}
						for {
							if entry.u1.v.type_flags != 0 {
								zend.ZvalAddrefP(entry)
							}
							for {
								var _zv *zend.Zval = orig_var
								if _zv.u1.v.type_ == 10 {
									var ref *zend.ZendReference = _zv.value.ref
									if ref.sources.ptr != nil {
										zend.ZendTryAssignTypedRefZvalEx(ref, entry, 0)
										break
									}
									_zv = &ref.val
								}
								zend.ZvalPtrDtor(_zv)
								var _z1 *zend.Zval = _zv
								var _z2 *zend.Zval = entry
								var _gc *zend.ZendRefcounted = _z2.value.counted
								var _t uint32 = _z2.u1.type_info
								_z1.value.counted = _gc
								_z1.u1.type_info = _t
								break
							}
							break
						}
						if zend.EG.exception != nil {
							zend.ZendStringReleaseEx(final_name.value.str, 0)
							return -1
						}
					} else {
						if entry.u1.v.type_flags != 0 {
							zend.ZvalAddrefP(entry)
						}
						zend.ZendHashAddNew(symbol_table, final_name.value.str, entry)
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

/* }}} */

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
			if _z.u1.v.type_ == 13 {
				_z = _z.value.zv
			}
			if _z.u1.v.type_ == 0 {
				continue
			}
			num_key = _p.h
			var_name = _p.key
			entry = _z
			if var_name != nil {
				if PhpValidVarName(var_name.val, var_name.len_) == 0 || var_name.len_ == g.SizeOf("\"this\"")-1 && !(memcmp(var_name.val, "this", g.SizeOf("\"this\"")-1)) {
					PhpPrefixVarname(&final_name, prefix, var_name.val, var_name.len_, 1)
					if PhpValidVarName(final_name.value.str.val, final_name.value.str.len_) == 0 {
						zend.ZvalPtrDtorStr(&final_name)
						continue
					}
				} else {
					var __z *zend.Zval = &final_name
					var __s *zend.ZendString = var_name
					__z.value.str = __s
					if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
						__z.u1.type_info = 6
					} else {
						zend.ZendGcAddref(&__s.gc)
						__z.u1.type_info = 6 | 1<<0<<8
					}
				}
			} else {
				var str *zend.ZendString = zend.ZendLongToStr(num_key)
				PhpPrefixVarname(&final_name, prefix, str.val, str.len_, 1)
				zend.ZendStringReleaseEx(str, 0)
				if PhpValidVarName(final_name.value.str.val, final_name.value.str.len_) == 0 {
					zend.ZvalPtrDtorStr(&final_name)
					continue
				}
			}
			if final_name.value.str.len_ == g.SizeOf("\"this\"")-1 && !(memcmp(final_name.value.str.val, "this", g.SizeOf("\"this\"")-1)) {
				zend.ZendThrowError(nil, "Cannot re-assign $this")
				return -1
			} else {
				if entry.u1.v.type_ == 10 {
					zend.ZvalAddrefP(entry)
				} else {
					var _z *zend.Zval = entry
					var _ref *zend.ZendReference = (*zend.ZendReference)(zend._emalloc(g.SizeOf("zend_reference")))
					zend.ZendGcSetRefcount(&_ref.gc, 2)
					_ref.gc.u.type_info = 10
					var _z1 *zend.Zval = &_ref.val
					var _z2 *zend.Zval = _z
					var _gc *zend.ZendRefcounted = _z2.value.counted
					var _t uint32 = _z2.u1.type_info
					_z1.value.counted = _gc
					_z1.u1.type_info = _t
					_ref.sources.ptr = nil
					_z.value.ref = _ref
					_z.u1.type_info = 10 | 1<<0<<8
				}
				if g.Assign(&orig_var, zend.ZendHashFind(symbol_table, final_name.value.str)) != nil {
					if orig_var.u1.v.type_ == 13 {
						orig_var = orig_var.value.zv
					}
					zend.ZvalPtrDtor(orig_var)
					var __z *zend.Zval = orig_var
					__z.value.ref = entry.value.ref
					__z.u1.type_info = 10 | 1<<0<<8
				} else {
					zend.ZendHashAddNew(symbol_table, final_name.value.str, entry)
				}
				count++
			}
			zend.ZvalPtrDtorStr(&final_name)
		}
		break
	}
	return count
}

/* }}} */

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
			if _z.u1.v.type_ == 13 {
				_z = _z.value.zv
			}
			if _z.u1.v.type_ == 0 {
				continue
			}
			num_key = _p.h
			var_name = _p.key
			entry = _z
			if var_name != nil {
				if PhpValidVarName(var_name.val, var_name.len_) == 0 || var_name.len_ == g.SizeOf("\"this\"")-1 && !(memcmp(var_name.val, "this", g.SizeOf("\"this\"")-1)) {
					PhpPrefixVarname(&final_name, prefix, var_name.val, var_name.len_, 1)
					if PhpValidVarName(final_name.value.str.val, final_name.value.str.len_) == 0 {
						zend.ZvalPtrDtorStr(&final_name)
						continue
					}
				} else {
					var __z *zend.Zval = &final_name
					var __s *zend.ZendString = var_name
					__z.value.str = __s
					if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
						__z.u1.type_info = 6
					} else {
						zend.ZendGcAddref(&__s.gc)
						__z.u1.type_info = 6 | 1<<0<<8
					}
				}
			} else {
				var str *zend.ZendString = zend.ZendLongToStr(num_key)
				PhpPrefixVarname(&final_name, prefix, str.val, str.len_, 1)
				zend.ZendStringReleaseEx(str, 0)
				if PhpValidVarName(final_name.value.str.val, final_name.value.str.len_) == 0 {
					zend.ZvalPtrDtorStr(&final_name)
					continue
				}
			}
			if final_name.value.str.len_ == g.SizeOf("\"this\"")-1 && !(memcmp(final_name.value.str.val, "this", g.SizeOf("\"this\"")-1)) {
				zend.ZendThrowError(nil, "Cannot re-assign $this")
				return -1
			} else {
				if entry.u1.v.type_ == 10 {
					entry = &(*entry).value.ref.val
				}
				if g.Assign(&orig_var, zend.ZendHashFind(symbol_table, final_name.value.str)) != nil {
					if orig_var.u1.v.type_ == 13 {
						orig_var = orig_var.value.zv
					}
					for {
						if entry.u1.v.type_flags != 0 {
							zend.ZvalAddrefP(entry)
						}
						for {
							var _zv *zend.Zval = orig_var
							if _zv.u1.v.type_ == 10 {
								var ref *zend.ZendReference = _zv.value.ref
								if ref.sources.ptr != nil {
									zend.ZendTryAssignTypedRefZvalEx(ref, entry, 0)
									break
								}
								_zv = &ref.val
							}
							zend.ZvalPtrDtor(_zv)
							var _z1 *zend.Zval = _zv
							var _z2 *zend.Zval = entry
							var _gc *zend.ZendRefcounted = _z2.value.counted
							var _t uint32 = _z2.u1.type_info
							_z1.value.counted = _gc
							_z1.u1.type_info = _t
							break
						}
						break
					}
					if zend.EG.exception != nil {
						zend.ZendStringReleaseEx(final_name.value.str, 0)
						return -1
					}
				} else {
					if entry.u1.v.type_flags != 0 {
						zend.ZvalAddrefP(entry)
					}
					zend.ZendHashAddNew(symbol_table, final_name.value.str, entry)
				}
				count++
			}
			zend.ZvalPtrDtorStr(&final_name)
		}
		break
	}
	return count
}

/* }}} */

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
			if _z.u1.v.type_ == 13 {
				_z = _z.value.zv
			}
			if _z.u1.v.type_ == 0 {
				continue
			}
			var_name = _p.key
			entry = _z
			if var_name == nil {
				continue
			}
			if PhpValidVarName(var_name.val, var_name.len_) == 0 {
				continue
			}
			if var_name.len_ == g.SizeOf("\"this\"")-1 && !(memcmp(var_name.val, "this", g.SizeOf("\"this\"")-1)) {
				continue
			}
			orig_var = zend.ZendHashFindEx(symbol_table, var_name, 1)
			if orig_var != nil {
				if orig_var.u1.v.type_ == 13 {
					orig_var = orig_var.value.zv
					if orig_var.u1.v.type_ == 0 {
						if entry.u1.v.type_ == 10 {
							zend.ZvalAddrefP(entry)
						} else {
							var _z *zend.Zval = entry
							var _ref *zend.ZendReference = (*zend.ZendReference)(zend._emalloc(g.SizeOf("zend_reference")))
							zend.ZendGcSetRefcount(&_ref.gc, 2)
							_ref.gc.u.type_info = 10
							var _z1 *zend.Zval = &_ref.val
							var _z2 *zend.Zval = _z
							var _gc *zend.ZendRefcounted = _z2.value.counted
							var _t uint32 = _z2.u1.type_info
							_z1.value.counted = _gc
							_z1.u1.type_info = _t
							_ref.sources.ptr = nil
							_z.value.ref = _ref
							_z.u1.type_info = 10 | 1<<0<<8
						}
						var __z *zend.Zval = orig_var
						__z.value.ref = entry.value.ref
						__z.u1.type_info = 10 | 1<<0<<8
						count++
					}
				}
			} else {
				if entry.u1.v.type_ == 10 {
					zend.ZvalAddrefP(entry)
				} else {
					var _z *zend.Zval = entry
					var _ref *zend.ZendReference = (*zend.ZendReference)(zend._emalloc(g.SizeOf("zend_reference")))
					zend.ZendGcSetRefcount(&_ref.gc, 2)
					_ref.gc.u.type_info = 10
					var _z1 *zend.Zval = &_ref.val
					var _z2 *zend.Zval = _z
					var _gc *zend.ZendRefcounted = _z2.value.counted
					var _t uint32 = _z2.u1.type_info
					_z1.value.counted = _gc
					_z1.u1.type_info = _t
					_ref.sources.ptr = nil
					_z.value.ref = _ref
					_z.u1.type_info = 10 | 1<<0<<8
				}
				zend.ZendHashAddNew(symbol_table, var_name, entry)
				count++
			}
		}
		break
	}
	return count
}

/* }}} */

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
			if _z.u1.v.type_ == 13 {
				_z = _z.value.zv
			}
			if _z.u1.v.type_ == 0 {
				continue
			}
			var_name = _p.key
			entry = _z
			if var_name == nil {
				continue
			}
			if PhpValidVarName(var_name.val, var_name.len_) == 0 {
				continue
			}
			if var_name.len_ == g.SizeOf("\"this\"")-1 && !(memcmp(var_name.val, "this", g.SizeOf("\"this\"")-1)) {
				continue
			}
			orig_var = zend.ZendHashFindEx(symbol_table, var_name, 1)
			if orig_var != nil {
				if orig_var.u1.v.type_ == 13 {
					orig_var = orig_var.value.zv
					if orig_var.u1.v.type_ == 0 {
						var _z3 *zend.Zval = entry
						if (_z3.u1.type_info & 0xff00) != 0 {
							if (_z3.u1.type_info & 0xff) == 10 {
								_z3 = &(*_z3).value.ref.val
								if (_z3.u1.type_info & 0xff00) != 0 {
									zend.ZvalAddrefP(_z3)
								}
							} else {
								zend.ZvalAddrefP(_z3)
							}
						}
						var _z1 *zend.Zval = orig_var
						var _z2 *zend.Zval = _z3
						var _gc *zend.ZendRefcounted = _z2.value.counted
						var _t uint32 = _z2.u1.type_info
						_z1.value.counted = _gc
						_z1.u1.type_info = _t
						count++
					}
				}
			} else {
				if entry.u1.v.type_ == 10 {
					entry = &(*entry).value.ref.val
				}
				if entry.u1.v.type_flags != 0 {
					zend.ZvalAddrefP(entry)
				}
				zend.ZendHashAddNew(symbol_table, var_name, entry)
				count++
			}
		}
		break
	}
	return count
}

/* }}} */

func ZifExtract(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var var_array_param *zend.Zval
	var prefix *zend.Zval = nil
	var extract_refs zend.ZendLong
	var extract_type zend.ZendLong = 0
	var count zend.ZendLong
	var symbol_table *zend.ZendArray
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 3
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg
			if _arg.u1.v.type_ == 10 {
				_arg = &(*_arg).value.ref.val
			}

			if zend.ZendParseArgArray(_arg, &var_array_param, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = 4
				break
			}
			_optional = 1
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgLong(_arg, &extract_type, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = 4
				break
			}
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			zend.ZendParseArgZvalDeref(_arg, &prefix, 0)
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
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
	extract_refs = extract_type & 0x100
	if extract_refs != 0 {
		var _zv *zend.Zval = var_array_param
		var _arr *zend.ZendArray = _zv.value.arr
		if zend.ZendGcRefcount(&_arr.gc) > 1 {
			if _zv.u1.v.type_flags != 0 {
				zend.ZendGcDelref(&_arr.gc)
			}
			var __arr *zend.ZendArray = zend.ZendArrayDup(_arr)
			var __z *zend.Zval = _zv
			__z.value.arr = __arr
			__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
		}
	}
	extract_type &= 0xff
	if extract_type < 0 || extract_type > 6 {
		core.PhpErrorDocref(nil, 1<<1, "Invalid extract type")
		return
	}
	if extract_type > 1 && extract_type <= 5 && execute_data.This.u2.num_args < 3 {
		core.PhpErrorDocref(nil, 1<<1, "specified extract type requires the prefix parameter")
		return
	}
	if prefix != nil {
		if zend.TryConvertToString(prefix) == 0 {
			return
		}
		if prefix.value.str.len_ != 0 && PhpValidVarName(prefix.value.str.val, prefix.value.str.len_) == 0 {
			core.PhpErrorDocref(nil, 1<<1, "prefix is not a valid identifier")
			return
		}
	}
	if zend.ZendForbidDynamicCall("extract()") == zend.FAILURE {
		return
	}
	symbol_table = zend.ZendRebuildSymbolTable()
	if extract_refs != 0 {
		switch extract_type {
		case 6:
			count = PhpExtractRefIfExists(var_array_param.value.arr, symbol_table)
			break
		case 0:
			count = PhpExtractRefOverwrite(var_array_param.value.arr, symbol_table)
			break
		case 5:
			count = PhpExtractRefPrefixIfExists(var_array_param.value.arr, symbol_table, prefix)
			break
		case 2:
			count = PhpExtractRefPrefixSame(var_array_param.value.arr, symbol_table, prefix)
			break
		case 3:
			count = PhpExtractRefPrefixAll(var_array_param.value.arr, symbol_table, prefix)
			break
		case 4:
			count = PhpExtractRefPrefixInvalid(var_array_param.value.arr, symbol_table, prefix)
			break
		default:
			count = PhpExtractRefSkip(var_array_param.value.arr, symbol_table)
			break
		}
	} else {

		/* The array might be stored in a local variable that will be overwritten */

		var array_copy zend.Zval
		var _z1 *zend.Zval = &array_copy
		var _z2 *zend.Zval = var_array_param
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
		if (_t & 0xff00) != 0 {
			zend.ZendGcAddref(&_gc.gc)
		}
		switch extract_type {
		case 6:
			count = PhpExtractIfExists(array_copy.value.arr, symbol_table)
			break
		case 0:
			count = PhpExtractOverwrite(array_copy.value.arr, symbol_table)
			break
		case 5:
			count = PhpExtractPrefixIfExists(array_copy.value.arr, symbol_table, prefix)
			break
		case 2:
			count = PhpExtractPrefixSame(array_copy.value.arr, symbol_table, prefix)
			break
		case 3:
			count = PhpExtractPrefixAll(array_copy.value.arr, symbol_table, prefix)
			break
		case 4:
			count = PhpExtractPrefixInvalid(array_copy.value.arr, symbol_table, prefix)
			break
		default:
			count = PhpExtractSkip(array_copy.value.arr, symbol_table)
			break
		}
		zend.ZvalPtrDtor(&array_copy)
	}
	var __z *zend.Zval = return_value
	__z.value.lval = count
	__z.u1.type_info = 4
	return
}

/* }}} */

func PhpCompactVar(eg_active_symbol_table *zend.HashTable, return_value *zend.Zval, entry *zend.Zval) {
	var value_ptr *zend.Zval
	var data zend.Zval
	if entry.u1.v.type_ == 10 {
		entry = &(*entry).value.ref.val
	}
	if entry.u1.v.type_ == 6 {
		if g.Assign(&value_ptr, zend.ZendHashFindInd(eg_active_symbol_table, entry.value.str)) != nil {
			if value_ptr.u1.v.type_ == 10 {
				value_ptr = &(*value_ptr).value.ref.val
			}
			if value_ptr.u1.v.type_flags != 0 {
				zend.ZvalAddrefP(value_ptr)
			}
			zend.ZendHashUpdate(return_value.value.arr, entry.value.str, value_ptr)
		} else if entry.value.str.len_ == g.SizeOf("\"this\"")-1 && !(memcmp(entry.value.str.val, "this", g.SizeOf("\"this\"")-1)) {
			var object *zend.ZendObject = zend.ZendGetThisObject(zend.EG.current_execute_data)
			if object != nil {
				zend.ZendGcAddref(&object.gc)
				var __z *zend.Zval = &data
				__z.value.obj = object
				__z.u1.type_info = 8 | 1<<0<<8 | 1<<1<<8
				zend.ZendHashUpdate(return_value.value.arr, entry.value.str, &data)
			}
		} else {
			core.PhpErrorDocref(nil, 1<<3, "Undefined variable: %s", entry.value.str.val)
		}
	} else if entry.u1.v.type_ == 7 {
		if entry.u1.v.type_flags != 0 {
			if (zend.ZvalGcFlags(entry.value.counted.gc.u.type_info) & 1 << 5) != 0 {
				core.PhpErrorDocref(nil, 1<<1, "recursion detected")
				return
			}
			entry.value.counted.gc.u.type_info |= 1 << 5 << 0
		}
		for {
			var __ht *zend.HashTable = entry.value.arr
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val
				if _z.u1.v.type_ == 13 {
					_z = _z.value.zv
				}
				if _z.u1.v.type_ == 0 {
					continue
				}
				value_ptr = _z
				PhpCompactVar(eg_active_symbol_table, return_value, value_ptr)
			}
			break
		}
		if entry.u1.v.type_flags != 0 {
			entry.value.counted.gc.u.type_info &= ^(1 << 5 << 0)
		}
	}
}

/* }}} */

func ZifCompact(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var args *zend.Zval = nil
	var num_args uint32
	var i uint32
	var symbol_table *zend.ZendArray
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = -1
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
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
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
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
	if symbol_table == nil {
		return
	}

	/* compact() is probably most used with a single array of var_names
	   or multiple string names, rather than a combination of both.
	   So quickly guess a minimum result size based on that */

	if num_args != 0 && args[0].u1.v.type_ == 7 {
		var __arr *zend.ZendArray = zend._zendNewArray(args[0].value.arr.nNumOfElements)
		var __z *zend.Zval = return_value
		__z.value.arr = __arr
		__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	} else {
		var __arr *zend.ZendArray = zend._zendNewArray(num_args)
		var __z *zend.Zval = return_value
		__z.value.arr = __arr
		__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	}
	for i = 0; i < num_args; i++ {
		PhpCompactVar(symbol_table, return_value, &args[i])
	}
}

/* }}} */

func ZifArrayFill(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var val *zend.Zval
	var start_key zend.ZendLong
	var num zend.ZendLong
	for {
		var _flags int = 0
		var _min_num_args int = 3
		var _max_num_args int = 3
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgLong(_arg, &start_key, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = 4
				break
			}
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgLong(_arg, &num, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = 4
				break
			}
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			zend.ZendParseArgZvalDeref(_arg, &val, 0)
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
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
	if num > 0 {
		if g.SizeOf("num") > 4 && num > 0x7fffffff {
			core.PhpErrorDocref(nil, 1<<1, "Too many elements")
			return_value.u1.type_info = 2
			return
		} else if start_key > INT64_MAX-num+1 {
			core.PhpErrorDocref(nil, 1<<1, "Cannot add element to the array as the next element is already occupied")
			return_value.u1.type_info = 2
			return
		} else if start_key >= 0 && start_key < num {

			/* create packed array */

			var p *zend.Bucket
			var n zend.ZendLong
			var __arr *zend.ZendArray = zend._zendNewArray(uint32(start_key + num))
			var __z *zend.Zval = return_value
			__z.value.arr = __arr
			__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
			zend.ZendHashRealInitPacked(return_value.value.arr)
			return_value.value.arr.nNumUsed = uint32(start_key + num)
			return_value.value.arr.nNumOfElements = uint32(num)
			return_value.value.arr.nNextFreeElement = zend_long(start_key + num)
			if val.u1.v.type_flags != 0 {
				zend.ZendGcAddrefEx(&(val.value.counted).gc, uint32(num))
			}
			p = return_value.value.arr.arData
			n = start_key
			for g.PostDec(&start_key) {
				&p.val.u1.type_info = 0
				p++
			}
			for g.PostDec(&num) {
				var _z1 *zend.Zval = &p.val
				var _z2 *zend.Zval = val
				var _gc *zend.ZendRefcounted = _z2.value.counted
				var _t uint32 = _z2.u1.type_info
				_z1.value.counted = _gc
				_z1.u1.type_info = _t
				n++
				p.h = n - 1
				p.key = nil
				p++
			}
		} else {

			/* create hash */

			var __arr *zend.ZendArray = zend._zendNewArray(uint32(num))
			var __z *zend.Zval = return_value
			__z.value.arr = __arr
			__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
			zend.ZendHashRealInitMixed(return_value.value.arr)
			if val.u1.v.type_flags != 0 {
				zend.ZendGcAddrefEx(&(val.value.counted).gc, uint32(num))
			}
			zend.ZendHashIndexAddNew(return_value.value.arr, start_key, val)
			for g.PreDec(&num) {
				zend.ZendHashNextIndexInsertNew(return_value.value.arr, val)
				start_key++
			}
		}
	} else if num == 0 {
		var __z *zend.Zval = return_value
		__z.value.arr = (*zend.ZendArray)(&zend.ZendEmptyArray)
		__z.u1.type_info = 7
		return
	} else {
		core.PhpErrorDocref(nil, 1<<1, "Number of elements can't be negative")
		return_value.u1.type_info = 2
		return
	}
}

/* }}} */

func ZifArrayFillKeys(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var keys *zend.Zval
	var val *zend.Zval
	var entry *zend.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 2
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgArray(_arg, &keys, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = 4
				break
			}
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			zend.ZendParseArgZvalDeref(_arg, &val, 0)
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
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

	var __arr *zend.ZendArray = zend._zendNewArray(keys.value.arr.nNumOfElements)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	for {
		var __ht *zend.HashTable = keys.value.arr
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if _z.u1.v.type_ == 0 {
				continue
			}
			entry = _z
			if entry.u1.v.type_ == 10 {
				entry = &(*entry).value.ref.val
			}
			if val.u1.v.type_flags != 0 {
				zend.ZvalAddrefP(val)
			}
			if entry.u1.v.type_ == 4 {
				zend.ZendHashIndexUpdate(return_value.value.arr, entry.value.lval, val)
			} else {
				var tmp_key *zend.ZendString
				var key *zend.ZendString = zend.ZvalGetTmpString(entry, &tmp_key)
				zend.ZendSymtableUpdate(return_value.value.arr, key, val)
				zend.ZendTmpStringRelease(tmp_key)
			}
		}
		break
	}
}

/* }}} */

// #define RANGE_CHECK_DOUBLE_INIT_ARRAY(start,end) do { double __calc_size = ( ( start - end ) / step ) + 1 ; if ( __calc_size >= ( double ) HT_MAX_SIZE ) { php_error_docref ( NULL , E_WARNING , "The supplied range exceeds the maximum array size: start=%0.0f end=%0.0f" , end , start ) ; RETURN_FALSE ; } size = ( uint32_t ) _php_math_round ( __calc_size , 0 , PHP_ROUND_HALF_UP ) ; array_init_size ( return_value , size ) ; zend_hash_real_init_packed ( Z_ARRVAL_P ( return_value ) ) ; } while ( 0 )

// #define RANGE_CHECK_LONG_INIT_ARRAY(start,end) do { zend_ulong __calc_size = ( ( zend_ulong ) start - end ) / lstep ; if ( __calc_size >= HT_MAX_SIZE - 1 ) { php_error_docref ( NULL , E_WARNING , "The supplied range exceeds the maximum array size: start=" ZEND_LONG_FMT " end=" ZEND_LONG_FMT , end , start ) ; RETURN_FALSE ; } size = ( uint32_t ) ( __calc_size + 1 ) ; array_init_size ( return_value , size ) ; zend_hash_real_init_packed ( Z_ARRVAL_P ( return_value ) ) ; } while ( 0 )

/* {{{ proto array range(mixed low, mixed high[, int step])
   Create an array containing the range of integers or characters from low to high (inclusive) */

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
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			zend.ZendParseArgZvalDeref(_arg, &zlow, 0)
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			zend.ZendParseArgZvalDeref(_arg, &zhigh, 0)
			_optional = 1
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			zend.ZendParseArgZvalDeref(_arg, &zstep, 0)
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return_value.u1.type_info = 2
			return
		}
		break
	}
	if zstep != nil {
		if zstep.u1.v.type_ == 5 {
			is_step_double = 1
		} else if zstep.u1.v.type_ == 6 {
			var type_ int = zend.IsNumericString(zstep.value.str.val, zstep.value.str.len_, nil, nil, 0)
			if type_ == 5 {
				is_step_double = 1
			}
			if type_ == 0 {

				/* bad number */

				core.PhpErrorDocref(nil, 1<<1, "Invalid range string - must be numeric")
				return_value.u1.type_info = 2
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

	if zlow.u1.v.type_ == 6 && zhigh.u1.v.type_ == 6 && zlow.value.str.len_ >= 1 && zhigh.value.str.len_ >= 1 {
		var type1 int
		var type2 int
		var low uint8
		var high uint8
		var lstep zend.ZendLong = zend.ZendLong(step)
		type1 = zend.IsNumericString(zlow.value.str.val, zlow.value.str.len_, nil, nil, 0)
		type2 = zend.IsNumericString(zhigh.value.str.val, zhigh.value.str.len_, nil, nil, 0)
		if type1 == 5 || type2 == 5 || is_step_double != 0 {
			goto double_str
		} else if type1 == 4 || type2 == 4 {
			goto long_str
		}
		low = uint8(zlow.value.str.val[0])
		high = uint8(zhigh.value.str.val[0])
		if low > high {
			if lstep <= 0 {
				err = 1
				goto err
			}

			/* Initialize the return_value as an array. */

			var __arr *zend.ZendArray = zend._zendNewArray(uint32((low-high)/lstep + 1))
			var __z *zend.Zval = return_value
			__z.value.arr = __arr
			__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
			zend.ZendHashRealInitPacked(return_value.value.arr)
			for {
				var __fill_ht *zend.HashTable = return_value.value.arr
				var __fill_bkt *zend.Bucket = __fill_ht.arData + __fill_ht.nNumUsed
				var __fill_idx uint32 = __fill_ht.nNumUsed
				assert((__fill_ht.u.flags & 1 << 2) != 0)
				for ; low >= high; low -= uint(lstep) {
					var __z *zend.Zval = &__fill_bkt.val
					var __s *zend.ZendString = zend.ZendOneCharString[low]
					__z.value.str = __s
					__z.u1.type_info = 6
					__fill_bkt.h = __fill_idx
					__fill_bkt.key = nil
					__fill_bkt++
					__fill_idx++
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
			var __arr *zend.ZendArray = zend._zendNewArray(uint32((high-low)/lstep + 1))
			var __z *zend.Zval = return_value
			__z.value.arr = __arr
			__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
			zend.ZendHashRealInitPacked(return_value.value.arr)
			for {
				var __fill_ht *zend.HashTable = return_value.value.arr
				var __fill_bkt *zend.Bucket = __fill_ht.arData + __fill_ht.nNumUsed
				var __fill_idx uint32 = __fill_ht.nNumUsed
				assert((__fill_ht.u.flags & 1 << 2) != 0)
				for ; low <= high; low += uint(lstep) {
					var __z *zend.Zval = &__fill_bkt.val
					var __s *zend.ZendString = zend.ZendOneCharString[low]
					__z.value.str = __s
					__z.u1.type_info = 6
					__fill_bkt.h = __fill_idx
					__fill_bkt.key = nil
					__fill_bkt++
					__fill_idx++
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
			var __arr *zend.ZendArray = zend._zendNewArray(0)
			var __z *zend.Zval = return_value
			__z.value.arr = __arr
			__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
			var __z *zend.Zval = &tmp
			var __s *zend.ZendString = zend.ZendOneCharString[low]
			__z.value.str = __s
			__z.u1.type_info = 6
			zend.ZendHashNextIndexInsertNew(return_value.value.arr, &tmp)
		}
	} else if zlow.u1.v.type_ == 5 || zhigh.u1.v.type_ == 5 || is_step_double != 0 {
		var low float64
		var high float64
		var element float64
		var i uint32
		var size uint32
	double_str:
		low = zend.ZvalGetDouble(zlow)
		high = zend.ZvalGetDouble(zhigh)
		if isinf(high) || isinf(low) {
			core.PhpErrorDocref(nil, 1<<1, "Invalid range supplied: start=%0.0f end=%0.0f", low, high)
			return_value.u1.type_info = 2
			return
		}
		if low > high {
			if low-high < step || step <= 0 {
				err = 1
				goto err
			}
			var __calc_size float64 = (low-high)/step + 1
			if __calc_size >= float64(0x80000000) {
				core.PhpErrorDocref(nil, 1<<1, "The supplied range exceeds the maximum array size: start=%0.0f end=%0.0f", high, low)
				return_value.u1.type_info = 2
				return
			}
			size = uint32(_phpMathRound(__calc_size, 0, 0x1))
			var __arr *zend.ZendArray = zend._zendNewArray(size)
			var __z *zend.Zval = return_value
			__z.value.arr = __arr
			__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
			zend.ZendHashRealInitPacked(return_value.value.arr)
			var __fill_ht *zend.HashTable = return_value.value.arr
			var __fill_bkt *zend.Bucket = __fill_ht.arData + __fill_ht.nNumUsed
			var __fill_idx uint32 = __fill_ht.nNumUsed
			assert((__fill_ht.u.flags & 1 << 2) != 0)
			i = 0
			element = low
			for i < size && element >= high {
				var __z *zend.Zval = &__fill_bkt.val
				__z.value.dval = element
				__z.u1.type_info = 5
				__fill_bkt.h = __fill_idx
				__fill_bkt.key = nil
				__fill_bkt++
				__fill_idx++
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
			if __calc_size >= float64(0x80000000) {
				core.PhpErrorDocref(nil, 1<<1, "The supplied range exceeds the maximum array size: start=%0.0f end=%0.0f", low, high)
				return_value.u1.type_info = 2
				return
			}
			size = uint32(_phpMathRound(__calc_size, 0, 0x1))
			var __arr *zend.ZendArray = zend._zendNewArray(size)
			var __z *zend.Zval = return_value
			__z.value.arr = __arr
			__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
			zend.ZendHashRealInitPacked(return_value.value.arr)
			var __fill_ht *zend.HashTable = return_value.value.arr
			var __fill_bkt *zend.Bucket = __fill_ht.arData + __fill_ht.nNumUsed
			var __fill_idx uint32 = __fill_ht.nNumUsed
			assert((__fill_ht.u.flags & 1 << 2) != 0)
			i = 0
			element = low
			for i < size && element <= high {
				var __z *zend.Zval = &__fill_bkt.val
				__z.value.dval = element
				__z.u1.type_info = 5
				__fill_bkt.h = __fill_idx
				__fill_bkt.key = nil
				__fill_bkt++
				__fill_idx++
				i++
				element = low + i*step
			}
			__fill_ht.nNumUsed = __fill_idx
			__fill_ht.nNumOfElements = __fill_idx
			__fill_ht.nNextFreeElement = __fill_idx
			__fill_ht.nInternalPointer = 0
		} else {
			var __arr *zend.ZendArray = zend._zendNewArray(0)
			var __z *zend.Zval = return_value
			__z.value.arr = __arr
			__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
			var __z *zend.Zval = &tmp
			__z.value.dval = low
			__z.u1.type_info = 5
			zend.ZendHashNextIndexInsertNew(return_value.value.arr, &tmp)
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
			if __calc_size >= 0x80000000-1 {
				core.PhpErrorDocref(nil, 1<<1, "The supplied range exceeds the maximum array size: start="+"%"+"lld"+" end="+"%"+"lld", high, low)
				return_value.u1.type_info = 2
				return
			}
			size = uint32(__calc_size + 1)
			var __arr *zend.ZendArray = zend._zendNewArray(size)
			var __z *zend.Zval = return_value
			__z.value.arr = __arr
			__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
			zend.ZendHashRealInitPacked(return_value.value.arr)
			var __fill_ht *zend.HashTable = return_value.value.arr
			var __fill_bkt *zend.Bucket = __fill_ht.arData + __fill_ht.nNumUsed
			var __fill_idx uint32 = __fill_ht.nNumUsed
			assert((__fill_ht.u.flags & 1 << 2) != 0)
			for i = 0; i < size; i++ {
				var __z *zend.Zval = &__fill_bkt.val
				__z.value.lval = low - i*lstep
				__z.u1.type_info = 4
				__fill_bkt.h = __fill_idx
				__fill_bkt.key = nil
				__fill_bkt++
				__fill_idx++
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
			if __calc_size >= 0x80000000-1 {
				core.PhpErrorDocref(nil, 1<<1, "The supplied range exceeds the maximum array size: start="+"%"+"lld"+" end="+"%"+"lld", low, high)
				return_value.u1.type_info = 2
				return
			}
			size = uint32(__calc_size + 1)
			var __arr *zend.ZendArray = zend._zendNewArray(size)
			var __z *zend.Zval = return_value
			__z.value.arr = __arr
			__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
			zend.ZendHashRealInitPacked(return_value.value.arr)
			var __fill_ht *zend.HashTable = return_value.value.arr
			var __fill_bkt *zend.Bucket = __fill_ht.arData + __fill_ht.nNumUsed
			var __fill_idx uint32 = __fill_ht.nNumUsed
			assert((__fill_ht.u.flags & 1 << 2) != 0)
			for i = 0; i < size; i++ {
				var __z *zend.Zval = &__fill_bkt.val
				__z.value.lval = low + i*lstep
				__z.u1.type_info = 4
				__fill_bkt.h = __fill_idx
				__fill_bkt.key = nil
				__fill_bkt++
				__fill_idx++
			}
			__fill_ht.nNumUsed = __fill_idx
			__fill_ht.nNumOfElements = __fill_idx
			__fill_ht.nNextFreeElement = __fill_idx
			__fill_ht.nInternalPointer = 0
		} else {
			var __arr *zend.ZendArray = zend._zendNewArray(0)
			var __z *zend.Zval = return_value
			__z.value.arr = __arr
			__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
			var __z *zend.Zval = &tmp
			__z.value.lval = low
			__z.u1.type_info = 4
			zend.ZendHashNextIndexInsertNew(return_value.value.arr, &tmp)
		}
	}
err:
	if err != 0 {
		core.PhpErrorDocref(nil, 1<<1, "step exceeds the specified range")
		return_value.u1.type_info = 2
		return
	}
}

/* }}} */

func PhpArrayDataShuffle(array *zend.Zval) {
	var idx uint32
	var j uint32
	var n_elems uint32
	var p *zend.Bucket
	var temp zend.Bucket
	var hash *zend.HashTable
	var rnd_idx zend.ZendLong
	var n_left uint32
	n_elems = array.value.arr.nNumOfElements
	if n_elems < 1 {
		return
	}
	hash = array.value.arr
	n_left = n_elems
	if hash.u.v.nIteratorsCount == 0 {
		if hash.nNumUsed != hash.nNumOfElements {
			j = 0
			idx = 0
			for ; idx < hash.nNumUsed; idx++ {
				p = hash.arData + idx
				if p.val.u1.v.type_ == 0 {
					continue
				}
				if j != idx {
					hash.arData[j] = *p
				}
				j++
			}
		}
		for g.PreDec(&n_left) {
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
				if p.val.u1.v.type_ == 0 {
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
		for g.PreDec(&n_left) {
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
	if (hash.u.flags & 1 << 2) == 0 {
		zend.ZendHashToPacked(hash)
	}
}

/* }}} */

func ZifShuffle(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var array *zend.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg
			if _arg.u1.v.type_ == 10 {
				_arg = &(*_arg).value.ref.val
			}
			var _zv *zend.Zval = _arg
			assert(_zv.u1.v.type_ != 10)
			var __zv *zend.Zval = _zv
			if __zv.u1.v.type_ == 7 {
				if zend.ZvalRefcountP(__zv) > 1 {
					if __zv.u1.v.type_flags != 0 {
						zend.ZvalDelrefP(__zv)
					}
					var __arr *zend.ZendArray = zend.ZendArrayDup(__zv.value.arr)
					var __z *zend.Zval = __zv
					__z.value.arr = __arr
					__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
				}
			}
			if zend.ZendParseArgArray(_arg, &array, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return_value.u1.type_info = 2
			return
		}
		break
	}
	PhpArrayDataShuffle(array)
	return_value.u1.type_info = 3
	return
}

/* }}} */

func PhpSplice(in_hash *zend.HashTable, offset zend.ZendLong, length zend.ZendLong, replace *zend.HashTable, removed *zend.HashTable) {
	var out_hash zend.HashTable
	var num_in zend.ZendLong
	var pos zend.ZendLong
	var idx uint32
	var p *zend.Bucket
	var entry *zend.Zval
	var iter_pos uint32 = zend.ZendHashIteratorsLowerPos(in_hash, 0)

	/* Get number of entries in the input hash */

	num_in = in_hash.nNumOfElements

	/* Clamp the offset.. */

	if offset > num_in {
		offset = num_in
	} else if offset < 0 && g.Assign(&offset, num_in+offset) < 0 {
		offset = 0
	}

	/* ..and the length */

	if length < 0 {
		length = num_in - offset + length
	} else if unsigned(offset+unsigned(length)) > unsigned(num_in) {
		length = num_in - offset
	}

	/* Create and initialize output hash */

	zend._zendHashInit(&out_hash, g.Cond(length > 0, num_in-length, 0)+g.CondF1(replace != nil, func() uint32 { return replace.nNumOfElements }, 0), zend.ZvalPtrDtor, 0)

	/* Start at the beginning of the input hash and copy entries to output hash until offset is reached */

	pos = 0
	idx = 0
	for ; pos < offset && idx < in_hash.nNumUsed; idx++ {
		p = in_hash.arData + idx
		if p.val.u1.v.type_ == 0 {
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
			if p.val.u1.v.type_ == 0 {
				continue
			}
			pos++
			entry = &p.val
			if entry.u1.v.type_flags != 0 {
				zend.ZvalAddrefP(entry)
			}
			if p.key == nil {
				zend.ZendHashNextIndexInsertNew(removed, entry)
				zend.ZendHashDelBucket(in_hash, p)
			} else {
				zend.ZendHashAddNew(removed, p.key, entry)
				if in_hash == &zend.EG.symbol_table {
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
			if p.val.u1.v.type_ == 0 {
				continue
			}
			pos2++
			if p.key != nil && in_hash == &zend.EG.symbol_table {
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
				if _z.u1.v.type_ == 13 {
					_z = _z.value.zv
				}
				if _z.u1.v.type_ == 0 {
					continue
				}
				entry = _z
				if entry.u1.v.type_flags != 0 {
					zend.ZvalAddrefP(entry)
				}
				zend.ZendHashNextIndexInsertNew(&out_hash, entry)
				pos++
			}
			break
		}
	}

	/* Copy the remaining input hash entries to the output hash */

	for ; idx < in_hash.nNumUsed; idx++ {
		p = in_hash.arData + idx
		if p.val.u1.v.type_ == 0 {
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

	&out_hash.u.v.nIteratorsCount = in_hash.u.v.nIteratorsCount
	in_hash.u.v.nIteratorsCount = 0
	in_hash.pDestructor = nil
	zend.ZendHashDestroy(in_hash)
	in_hash.u.flags = &out_hash.u.flags
	in_hash.nTableSize = out_hash.nTableSize
	in_hash.nTableMask = out_hash.nTableMask
	in_hash.nNumUsed = out_hash.nNumUsed
	in_hash.nNumOfElements = out_hash.nNumOfElements
	in_hash.nNextFreeElement = out_hash.nNextFreeElement
	in_hash.arData = out_hash.arData
	in_hash.pDestructor = out_hash.pDestructor
	zend.ZendHashInternalPointerResetEx(in_hash, &in_hash.nInternalPointer)
}

/* }}} */

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
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg
			if _arg.u1.v.type_ == 10 {
				_arg = &(*_arg).value.ref.val
			}
			var _zv *zend.Zval = _arg
			assert(_zv.u1.v.type_ != 10)
			var __zv *zend.Zval = _zv
			if __zv.u1.v.type_ == 7 {
				if zend.ZvalRefcountP(__zv) > 1 {
					if __zv.u1.v.type_flags != 0 {
						zend.ZvalDelrefP(__zv)
					}
					var __arr *zend.ZendArray = zend.ZendArrayDup(__zv.value.arr)
					var __z *zend.Zval = __zv
					__z.value.arr = __arr
					__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
				}
			}
			if zend.ZendParseArgArray(_arg, &stack, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = 4
				break
			}
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
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
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
		var _z1 *zend.Zval = &new_var
		var _z2 *zend.Zval = &args[i]
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
		if (_t & 0xff00) != 0 {
			zend.ZendGcAddref(&_gc.gc)
		}
		if zend.ZendHashNextIndexInsert(stack.value.arr, &new_var) == nil {
			if &new_var.u1.v.type_flags != 0 {
				zend.ZvalDelrefP(&new_var)
			}
			core.PhpErrorDocref(nil, 1<<1, "Cannot add element to the array as the next element is already occupied")
			return_value.u1.type_info = 2
			return
		}
	}

	/* Clean up and return the number of values in the stack */

	var __z *zend.Zval = return_value
	__z.value.lval = stack.value.arr.nNumOfElements
	__z.u1.type_info = 4
}

/* }}} */

func ZifArrayPop(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var stack *zend.Zval
	var val *zend.Zval
	var idx uint32
	var p *zend.Bucket
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg
			if _arg.u1.v.type_ == 10 {
				_arg = &(*_arg).value.ref.val
			}
			var _zv *zend.Zval = _arg
			assert(_zv.u1.v.type_ != 10)
			var __zv *zend.Zval = _zv
			if __zv.u1.v.type_ == 7 {
				if zend.ZvalRefcountP(__zv) > 1 {
					if __zv.u1.v.type_flags != 0 {
						zend.ZvalDelrefP(__zv)
					}
					var __arr *zend.ZendArray = zend.ZendArrayDup(__zv.value.arr)
					var __z *zend.Zval = __zv
					__z.value.arr = __arr
					__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
				}
			}
			if zend.ZendParseArgArray(_arg, &stack, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
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
	if stack.value.arr.nNumOfElements == 0 {
		return
	}

	/* Get the last value and copy it into the return value */

	idx = stack.value.arr.nNumUsed
	for true {
		if idx == 0 {
			return
		}
		idx--
		p = stack.value.arr.arData + idx
		val = &p.val
		if val.u1.v.type_ == 13 {
			val = val.value.zv
		}
		if val.u1.v.type_ != 0 {
			break
		}
	}
	var _z3 *zend.Zval = val
	if (_z3.u1.type_info & 0xff00) != 0 {
		if (_z3.u1.type_info & 0xff) == 10 {
			_z3 = &(*_z3).value.ref.val
			if (_z3.u1.type_info & 0xff00) != 0 {
				zend.ZvalAddrefP(_z3)
			}
		} else {
			zend.ZvalAddrefP(_z3)
		}
	}
	var _z1 *zend.Zval = return_value
	var _z2 *zend.Zval = _z3
	var _gc *zend.ZendRefcounted = _z2.value.counted
	var _t uint32 = _z2.u1.type_info
	_z1.value.counted = _gc
	_z1.u1.type_info = _t
	if p.key == nil && stack.value.arr.nNextFreeElement > 0 && p.h >= zend_ulong(stack.value.arr.nNextFreeElement-1) {
		stack.value.arr.nNextFreeElement = stack.value.arr.nNextFreeElement - 1
	}

	/* Delete the last value */

	if p.key != nil && stack.value.arr == &zend.EG.symbol_table {
		zend.ZendDeleteGlobalVariable(p.key)
	} else {
		zend.ZendHashDelBucket(stack.value.arr, p)
	}
	zend.ZendHashInternalPointerResetEx(stack.value.arr, &(stack.value.arr).nInternalPointer)
}

/* }}} */

func ZifArrayShift(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var stack *zend.Zval
	var val *zend.Zval
	var idx uint32
	var p *zend.Bucket
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg
			if _arg.u1.v.type_ == 10 {
				_arg = &(*_arg).value.ref.val
			}
			var _zv *zend.Zval = _arg
			assert(_zv.u1.v.type_ != 10)
			var __zv *zend.Zval = _zv
			if __zv.u1.v.type_ == 7 {
				if zend.ZvalRefcountP(__zv) > 1 {
					if __zv.u1.v.type_flags != 0 {
						zend.ZvalDelrefP(__zv)
					}
					var __arr *zend.ZendArray = zend.ZendArrayDup(__zv.value.arr)
					var __z *zend.Zval = __zv
					__z.value.arr = __arr
					__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
				}
			}
			if zend.ZendParseArgArray(_arg, &stack, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
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
	if stack.value.arr.nNumOfElements == 0 {
		return
	}

	/* Get the first value and copy it into the return value */

	idx = 0
	for true {
		if idx == stack.value.arr.nNumUsed {
			return
		}
		p = stack.value.arr.arData + idx
		val = &p.val
		if val.u1.v.type_ == 13 {
			val = val.value.zv
		}
		if val.u1.v.type_ != 0 {
			break
		}
		idx++
	}
	var _z3 *zend.Zval = val
	if (_z3.u1.type_info & 0xff00) != 0 {
		if (_z3.u1.type_info & 0xff) == 10 {
			_z3 = &(*_z3).value.ref.val
			if (_z3.u1.type_info & 0xff00) != 0 {
				zend.ZvalAddrefP(_z3)
			}
		} else {
			zend.ZvalAddrefP(_z3)
		}
	}
	var _z1 *zend.Zval = return_value
	var _z2 *zend.Zval = _z3
	var _gc *zend.ZendRefcounted = _z2.value.counted
	var _t uint32 = _z2.u1.type_info
	_z1.value.counted = _gc
	_z1.u1.type_info = _t

	/* Delete the first value */

	if p.key != nil && stack.value.arr == &zend.EG.symbol_table {
		zend.ZendDeleteGlobalVariable(p.key)
	} else {
		zend.ZendHashDelBucket(stack.value.arr, p)
	}

	/* re-index like it did before */

	if (stack.value.arr.u.flags & 1 << 2) != 0 {
		var k uint32 = 0
		if stack.value.arr.u.v.nIteratorsCount == 0 {
			for idx = 0; idx < stack.value.arr.nNumUsed; idx++ {
				p = stack.value.arr.arData + idx
				if p.val.u1.v.type_ == 0 {
					continue
				}
				if idx != k {
					var q *zend.Bucket = stack.value.arr.arData + k
					q.h = k
					q.key = nil
					var _z1 *zend.Zval = &q.val
					var _z2 *zend.Zval = &p.val
					var _gc *zend.ZendRefcounted = _z2.value.counted
					var _t uint32 = _z2.u1.type_info
					_z1.value.counted = _gc
					_z1.u1.type_info = _t
					&p.val.u1.type_info = 0
				}
				k++
			}
		} else {
			var iter_pos uint32 = zend.ZendHashIteratorsLowerPos(stack.value.arr, 0)
			for idx = 0; idx < stack.value.arr.nNumUsed; idx++ {
				p = stack.value.arr.arData + idx
				if p.val.u1.v.type_ == 0 {
					continue
				}
				if idx != k {
					var q *zend.Bucket = stack.value.arr.arData + k
					q.h = k
					q.key = nil
					var _z1 *zend.Zval = &q.val
					var _z2 *zend.Zval = &p.val
					var _gc *zend.ZendRefcounted = _z2.value.counted
					var _t uint32 = _z2.u1.type_info
					_z1.value.counted = _gc
					_z1.u1.type_info = _t
					&p.val.u1.type_info = 0
					if idx == iter_pos {
						zend.ZendHashIteratorsUpdate(stack.value.arr, idx, k)
						iter_pos = zend.ZendHashIteratorsLowerPos(stack.value.arr, iter_pos+1)
					}
				}
				k++
			}
		}
		stack.value.arr.nNumUsed = k
		stack.value.arr.nNextFreeElement = k
	} else {
		var k uint32 = 0
		var should_rehash int = 0
		for idx = 0; idx < stack.value.arr.nNumUsed; idx++ {
			p = stack.value.arr.arData + idx
			if p.val.u1.v.type_ == 0 {
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
		stack.value.arr.nNextFreeElement = k
		if should_rehash != 0 {
			zend.ZendHashRehash(stack.value.arr)
		}
	}
	zend.ZendHashInternalPointerResetEx(stack.value.arr, &(stack.value.arr).nInternalPointer)
}

/* }}} */

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
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg
			if _arg.u1.v.type_ == 10 {
				_arg = &(*_arg).value.ref.val
			}
			var _zv *zend.Zval = _arg
			assert(_zv.u1.v.type_ != 10)
			var __zv *zend.Zval = _zv
			if __zv.u1.v.type_ == 7 {
				if zend.ZvalRefcountP(__zv) > 1 {
					if __zv.u1.v.type_flags != 0 {
						zend.ZvalDelrefP(__zv)
					}
					var __arr *zend.ZendArray = zend.ZendArrayDup(__zv.value.arr)
					var __z *zend.Zval = __zv
					__z.value.arr = __arr
					__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
				}
			}
			if zend.ZendParseArgArray(_arg, &stack, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = 4
				break
			}
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
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
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
	zend._zendHashInit(&new_hash, stack.value.arr.nNumOfElements+argc, zend.ZvalPtrDtor, 0)
	for i = 0; i < argc; i++ {
		if &args[i].u1.v.type_flags != 0 {
			zend.ZvalAddrefP(&args[i])
		}
		zend.ZendHashNextIndexInsertNew(&new_hash, &args[i])
	}
	for {
		var __ht *zend.HashTable = stack.value.arr
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if _z.u1.v.type_ == 0 {
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
	if stack.value.arr.u.v.nIteratorsCount != 0 {
		zend.ZendHashIteratorsAdvance(stack.value.arr, argc)
		&new_hash.u.v.nIteratorsCount = stack.value.arr.u.v.nIteratorsCount
		stack.value.arr.u.v.nIteratorsCount = 0
	}

	/* replace HashTable data */

	stack.value.arr.pDestructor = nil
	zend.ZendHashDestroy(stack.value.arr)
	stack.value.arr.u.flags = &new_hash.u.flags
	stack.value.arr.nTableSize = new_hash.nTableSize
	stack.value.arr.nTableMask = new_hash.nTableMask
	stack.value.arr.nNumUsed = new_hash.nNumUsed
	stack.value.arr.nNumOfElements = new_hash.nNumOfElements
	stack.value.arr.nNextFreeElement = new_hash.nNextFreeElement
	stack.value.arr.arData = new_hash.arData
	stack.value.arr.pDestructor = new_hash.pDestructor
	zend.ZendHashInternalPointerResetEx(stack.value.arr, &(stack.value.arr).nInternalPointer)

	/* Clean up and return the number of elements in the stack */

	var __z *zend.Zval = return_value
	__z.value.lval = stack.value.arr.nNumOfElements
	__z.u1.type_info = 4
}

/* }}} */

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
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg
			if _arg.u1.v.type_ == 10 {
				_arg = &(*_arg).value.ref.val
			}
			var _zv *zend.Zval = _arg
			assert(_zv.u1.v.type_ != 10)
			var __zv *zend.Zval = _zv
			if __zv.u1.v.type_ == 7 {
				if zend.ZvalRefcountP(__zv) > 1 {
					if __zv.u1.v.type_flags != 0 {
						zend.ZvalDelrefP(__zv)
					}
					var __arr *zend.ZendArray = zend.ZendArrayDup(__zv.value.arr)
					var __z *zend.Zval = __zv
					__z.value.arr = __arr
					__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
				}
			}
			if zend.ZendParseArgArray(_arg, &array, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = 4
				break
			}
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgLong(_arg, &offset, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = 4
				break
			}
			_optional = 1
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgLong(_arg, &length, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = 4
				break
			}
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			zend.ZendParseArgZvalDeref(_arg, &repl_array, 0)
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
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
	num_in = array.value.arr.nNumOfElements
	if execute_data.This.u2.num_args < 3 {
		length = num_in
	}
	if execute_data.This.u2.num_args == 4 {

		/* Make sure the last argument, if passed, is an array */

		if repl_array.u1.v.type_ != 7 {
			zend.ConvertToArray(repl_array)
		}
	}

	/* Don't create the array of removed elements if it's not going
	 * to be used; e.g. only removing and/or replacing elements */

	if execute_data.prev_execute_data == nil || (execute_data.prev_execute_data.func_.common.type_&1) != 0 || execute_data.prev_execute_data.opline.result_type != 0 {
		var size zend.ZendLong = length

		/* Clamp the offset.. */

		if offset > num_in {
			offset = num_in
		} else if offset < 0 && g.Assign(&offset, num_in+offset) < 0 {
			offset = 0
		}

		/* ..and the length */

		if length < 0 {
			size = num_in - offset + length
		} else if zend.ZendUlong(offset+zend.ZendUlong(length)) > uint32(num_in) {
			size = num_in - offset
		}

		/* Initialize return value */

		var __arr *zend.ZendArray = zend._zendNewArray(g.CondF1(size > 0, func() uint32 { return uint32(size) }, 0))
		var __z *zend.Zval = return_value
		__z.value.arr = __arr
		__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
		rem_hash = return_value.value.arr
	}

	/* Perform splice */

	PhpSplice(array.value.arr, offset, length, g.CondF1(repl_array != nil, func() *zend.ZendArray { return repl_array.value.arr }, nil), rem_hash)

	/* Perform splice */
}

/* }}} */

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
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgArray(_arg, &input, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = 4
				break
			}
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgLong(_arg, &offset, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = 4
				break
			}
			_optional = 1
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			zend.ZendParseArgZvalDeref(_arg, &z_length, 0)
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgBool(_arg, &preserve_keys, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
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

	num_in = input.value.arr.nNumOfElements

	/* We want all entries from offset to the end if length is not passed or is null */

	if execute_data.This.u2.num_args < 3 || z_length.u1.v.type_ == 1 {
		length = num_in
	} else {
		length = zend.ZvalGetLong(z_length)
	}

	/* Clamp the offset.. */

	if offset > num_in {
		var __z *zend.Zval = return_value
		__z.value.arr = (*zend.ZendArray)(&zend.ZendEmptyArray)
		__z.u1.type_info = 7
		return
	} else if offset < 0 && g.Assign(&offset, num_in+offset) < 0 {
		offset = 0
	}

	/* ..and the length */

	if length < 0 {
		length = num_in - offset + length
	} else if zend.ZendUlong(offset+zend.ZendUlong(length)) > unsigned(num_in) {
		length = num_in - offset
	}
	if length <= 0 {
		var __z *zend.Zval = return_value
		__z.value.arr = (*zend.ZendArray)(&zend.ZendEmptyArray)
		__z.u1.type_info = 7
		return
	}

	/* Initialize returned array */

	var __arr *zend.ZendArray = zend._zendNewArray(uint32(length))
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8

	/* Start at the beginning and go until we hit offset */

	pos = 0
	if (input.value.arr.u.flags&1<<2) != 0 && (preserve_keys == 0 || offset == 0 && input.value.arr.nNumUsed == input.value.arr.nNumOfElements) {
		zend.ZendHashRealInitPacked(return_value.value.arr)
		for {
			var __fill_ht *zend.HashTable = return_value.value.arr
			var __fill_bkt *zend.Bucket = __fill_ht.arData + __fill_ht.nNumUsed
			var __fill_idx uint32 = __fill_ht.nNumUsed
			assert((__fill_ht.u.flags & 1 << 2) != 0)
			for {
				var __ht *zend.HashTable = input.value.arr
				var _p *zend.Bucket = __ht.arData
				var _end *zend.Bucket = _p + __ht.nNumUsed
				for ; _p != _end; _p++ {
					var _z *zend.Zval = &_p.val

					if _z.u1.v.type_ == 0 {
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
					if entry.u1.v.type_ == 10 && zend.ZvalRefcountP(entry) == 1 {
						entry = &(*entry).value.ref.val
					}
					if entry.u1.v.type_flags != 0 {
						zend.ZvalAddrefP(entry)
					}
					var _z1 *zend.Zval = &__fill_bkt.val
					var _z2 *zend.Zval = entry
					var _gc *zend.ZendRefcounted = _z2.value.counted
					var _t uint32 = _z2.u1.type_info
					_z1.value.counted = _gc
					_z1.u1.type_info = _t
					__fill_bkt.h = __fill_idx
					__fill_bkt.key = nil
					__fill_bkt++
					__fill_idx++
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
			var __ht *zend.HashTable = input.value.arr
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val

				if _z.u1.v.type_ == 0 {
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
					entry = zend.ZendHashAddNew(return_value.value.arr, string_key, entry)
				} else {
					if preserve_keys != 0 {
						entry = zend.ZendHashIndexAddNew(return_value.value.arr, num_key, entry)
					} else {
						entry = zend.ZendHashNextIndexInsertNew(return_value.value.arr, entry)
					}
				}
				zend.ZvalAddRef(entry)
			}
			break
		}
	}
}

/* }}} */

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

			if _z.u1.v.type_ == 0 {
				continue
			}
			string_key = _p.key
			src_entry = _z
			if string_key != nil {
				if g.Assign(&dest_entry, zend.ZendHashFindEx(dest, string_key, 1)) != nil {
					var src_zval *zend.Zval = src_entry
					var dest_zval *zend.Zval = dest_entry
					var thash *zend.HashTable
					var tmp zend.Zval
					var ret int
					if src_zval.u1.v.type_ == 10 {
						src_zval = &(*src_zval).value.ref.val
					}
					if dest_zval.u1.v.type_ == 10 {
						dest_zval = &(*dest_zval).value.ref.val
					}
					if dest_zval.u1.v.type_ == 7 {
						thash = dest_zval.value.arr
					} else {
						thash = nil
					}
					if thash != nil && (zend.ZvalGcFlags(thash.gc.u.type_info)&1<<5) != 0 || src_entry == dest_entry && dest_entry.u1.v.type_ == 10 && zend.ZvalRefcountP(dest_entry)%2 != 0 {
						core.PhpErrorDocref(nil, 1<<1, "recursion detected")
						return 0
					}
					assert(dest_entry.u1.v.type_ != 10 || zend.ZvalRefcountP(dest_entry) > 1)
					for {
						var _zv *zend.Zval = dest_entry
						if _zv.u1.v.type_ == 10 {
							var _r *zend.ZendReference = _zv.value.ref
							var _z1 *zend.Zval = _zv
							var _z2 *zend.Zval = &_r.val
							var _gc *zend.ZendRefcounted = _z2.value.counted
							var _t uint32 = _z2.u1.type_info
							_z1.value.counted = _gc
							_z1.u1.type_info = _t
							if zend.ZendGcDelref(&_r.gc) == 0 {
								zend._efree(_r)
							} else if (_zv.u1.type_info & 0xff) == 7 {
								var __arr *zend.ZendArray = zend.ZendArrayDup(_zv.value.arr)
								var __z *zend.Zval = _zv
								__z.value.arr = __arr
								__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
								break
							} else if (_zv.u1.type_info & 0xff00) != 0 {
								zend.ZvalAddrefP(_zv)
								break
							}
						}
						var __zv *zend.Zval = _zv
						if __zv.u1.v.type_ == 7 {
							if zend.ZvalRefcountP(__zv) > 1 {
								if __zv.u1.v.type_flags != 0 {
									zend.ZvalDelrefP(__zv)
								}
								var __arr *zend.ZendArray = zend.ZendArrayDup(__zv.value.arr)
								var __z *zend.Zval = __zv
								__z.value.arr = __arr
								__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
							}
						}
						break
					}
					dest_zval = dest_entry
					if dest_zval.u1.v.type_ == 1 {
						if dest_zval.u1.v.type_ != 7 {
							zend.ConvertToArray(dest_zval)
						}
						zend.AddNextIndexNull(dest_zval)
					} else {
						if dest_zval.u1.v.type_ != 7 {
							zend.ConvertToArray(dest_zval)
						}
					}
					&tmp.u1.type_info = 0
					if src_zval.u1.v.type_ == 8 {
						var _z1 *zend.Zval = &tmp
						var _z2 *zend.Zval = src_zval
						var _gc *zend.ZendRefcounted = _z2.value.counted
						var _t uint32 = _z2.u1.type_info
						_z1.value.counted = _gc
						_z1.u1.type_info = _t
						if (_t & 0xff00) != 0 {
							zend.ZendGcAddref(&_gc.gc)
						}
						zend.ConvertToArray(&tmp)
						src_zval = &tmp
					}
					if src_zval.u1.v.type_ == 7 {
						if thash != nil && (zend.ZvalGcFlags(thash.gc.u.type_info)&1<<6) == 0 {
							thash.gc.u.type_info |= 1 << 5 << 0
						}
						ret = PhpArrayMergeRecursive(dest_zval.value.arr, src_zval.value.arr)
						if thash != nil && (zend.ZvalGcFlags(thash.gc.u.type_info)&1<<6) == 0 {
							thash.gc.u.type_info &= ^(1 << 5 << 0)
						}
						if ret == 0 {
							return 0
						}
					} else {
						if src_zval.u1.v.type_flags != 0 {
							zend.ZvalAddrefP(src_zval)
						}
						zend.ZendHashNextIndexInsert(dest_zval.value.arr, src_zval)
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

/* }}} */

func PhpArrayMerge(dest *zend.HashTable, src *zend.HashTable) int {
	var src_entry *zend.Zval
	var string_key *zend.ZendString
	if (dest.u.flags&1<<2) != 0 && (src.u.flags&1<<2) != 0 {
		zend.ZendHashExtend(dest, dest.nNumOfElements+src.nNumOfElements, 1)
		for {
			var __fill_ht *zend.HashTable = dest
			var __fill_bkt *zend.Bucket = __fill_ht.arData + __fill_ht.nNumUsed
			var __fill_idx uint32 = __fill_ht.nNumUsed
			assert((__fill_ht.u.flags & 1 << 2) != 0)
			for {
				var __ht *zend.HashTable = src
				var _p *zend.Bucket = __ht.arData
				var _end *zend.Bucket = _p + __ht.nNumUsed
				for ; _p != _end; _p++ {
					var _z *zend.Zval = &_p.val

					if _z.u1.v.type_ == 0 {
						continue
					}
					src_entry = _z
					if src_entry.u1.v.type_ == 10 && zend.ZvalRefcountP(src_entry) == 1 {
						src_entry = &(*src_entry).value.ref.val
					}
					if src_entry.u1.v.type_flags != 0 {
						zend.ZvalAddrefP(src_entry)
					}
					var _z1 *zend.Zval = &__fill_bkt.val
					var _z2 *zend.Zval = src_entry
					var _gc *zend.ZendRefcounted = _z2.value.counted
					var _t uint32 = _z2.u1.type_info
					_z1.value.counted = _gc
					_z1.u1.type_info = _t
					__fill_bkt.h = __fill_idx
					__fill_bkt.key = nil
					__fill_bkt++
					__fill_idx++
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

				if _z.u1.v.type_ == 0 {
					continue
				}
				string_key = _p.key
				src_entry = _z
				if src_entry.u1.v.type_ == 10 && zend.ZvalRefcountP(src_entry) == 1 {
					src_entry = &(*src_entry).value.ref.val
				}
				if src_entry.u1.v.type_flags != 0 {
					zend.ZvalAddrefP(src_entry)
				}
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

/* }}} */

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

			if _z.u1.v.type_ == 0 {
				continue
			}
			num_key = _p.h
			string_key = _p.key
			src_entry = _z
			src_zval = src_entry
			if src_zval.u1.v.type_ == 10 {
				src_zval = &(*src_zval).value.ref.val
			}
			if string_key != nil {
				if src_zval.u1.v.type_ != 7 || g.Assign(&dest_entry, zend.ZendHashFindEx(dest, string_key, 1)) == nil || dest_entry.u1.v.type_ != 7 && (dest_entry.u1.v.type_ != 10 || &(*dest_entry).value.ref.val.u1.v.type_ != 7) {
					var zv *zend.Zval = zend.ZendHashUpdate(dest, string_key, src_entry)
					zend.ZvalAddRef(zv)
					continue
				}
			} else {
				if src_zval.u1.v.type_ != 7 || g.Assign(&dest_entry, zend.ZendHashIndexFind(dest, num_key)) == nil || dest_entry.u1.v.type_ != 7 && (dest_entry.u1.v.type_ != 10 || &(*dest_entry).value.ref.val.u1.v.type_ != 7) {
					var zv *zend.Zval = zend.ZendHashIndexUpdate(dest, num_key, src_entry)
					zend.ZvalAddRef(zv)
					continue
				}
			}
			dest_zval = dest_entry
			if dest_zval.u1.v.type_ == 10 {
				dest_zval = &(*dest_zval).value.ref.val
			}
			if (zend.ZvalGcFlags(dest_zval.value.counted.gc.u.type_info)&1<<5) != 0 || (zend.ZvalGcFlags(src_zval.value.counted.gc.u.type_info)&1<<5) != 0 || src_entry.u1.v.type_ == 10 && dest_entry.u1.v.type_ == 10 && src_entry.value.ref == dest_entry.value.ref && zend.ZvalRefcountP(dest_entry)%2 != 0 {
				core.PhpErrorDocref(nil, 1<<1, "recursion detected")
				return 0
			}
			assert(dest_entry.u1.v.type_ != 10 || zend.ZvalRefcountP(dest_entry) > 1)
			for {
				var _zv *zend.Zval = dest_entry
				if _zv.u1.v.type_ == 10 {
					var _r *zend.ZendReference = _zv.value.ref
					var _z1 *zend.Zval = _zv
					var _z2 *zend.Zval = &_r.val
					var _gc *zend.ZendRefcounted = _z2.value.counted
					var _t uint32 = _z2.u1.type_info
					_z1.value.counted = _gc
					_z1.u1.type_info = _t
					if zend.ZendGcDelref(&_r.gc) == 0 {
						zend._efree(_r)
					} else if (_zv.u1.type_info & 0xff) == 7 {
						var __arr *zend.ZendArray = zend.ZendArrayDup(_zv.value.arr)
						var __z *zend.Zval = _zv
						__z.value.arr = __arr
						__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
						break
					} else if (_zv.u1.type_info & 0xff00) != 0 {
						zend.ZvalAddrefP(_zv)
						break
					}
				}
				var __zv *zend.Zval = _zv
				if __zv.u1.v.type_ == 7 {
					if zend.ZvalRefcountP(__zv) > 1 {
						if __zv.u1.v.type_flags != 0 {
							zend.ZvalDelrefP(__zv)
						}
						var __arr *zend.ZendArray = zend.ZendArrayDup(__zv.value.arr)
						var __z *zend.Zval = __zv
						__z.value.arr = __arr
						__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
					}
				}
				break
			}
			dest_zval = dest_entry
			if dest_zval.u1.v.type_flags != 0 {
				dest_zval.value.counted.gc.u.type_info |= 1 << 5 << 0
			}
			if src_zval.u1.v.type_flags != 0 {
				src_zval.value.counted.gc.u.type_info |= 1 << 5 << 0
			}
			ret = PhpArrayReplaceRecursive(dest_zval.value.arr, src_zval.value.arr)
			if dest_zval.u1.v.type_flags != 0 {
				dest_zval.value.counted.gc.u.type_info &= ^(1 << 5 << 0)
			}
			if src_zval.u1.v.type_flags != 0 {
				src_zval.value.counted.gc.u.type_info &= ^(1 << 5 << 0)
			}
			if ret == 0 {
				return 0
			}
		}
		break
	}
	return 1
}

/* }}} */

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
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
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
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
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
		if arg.u1.v.type_ != 7 {
			core.PhpErrorDocref(nil, 1<<1, "Expected parameter %d to be an array, %s given", i+1, zend.ZendZvalTypeName(arg))
			return_value.u1.type_info = 1
			return
		}
	}

	/* copy first array */

	arg = args
	dest = zend.ZendArrayDup(arg.value.arr)
	var __arr *zend.ZendArray = dest
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	if recursive != 0 {
		for i = 1; i < argc; i++ {
			arg = args + i
			PhpArrayReplaceRecursive(dest, arg.value.arr)
		}
	} else {
		for i = 1; i < argc; i++ {
			arg = args + i
			zend.ZendHashMerge(dest, arg.value.arr, zend.ZvalAddRef, 1)
		}
	}
}

/* }}} */

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
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
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
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
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
		var __z *zend.Zval = return_value
		__z.value.arr = (*zend.ZendArray)(&zend.ZendEmptyArray)
		__z.u1.type_info = 7
		return
	}
	for i = 0; i < argc; i++ {
		var arg *zend.Zval = args + i
		if arg.u1.v.type_ != 7 {
			core.PhpErrorDocref(nil, 1<<1, "Expected parameter %d to be an array, %s given", i+1, zend.ZendZvalTypeName(arg))
			return_value.u1.type_info = 1
			return
		}
		count += arg.value.arr.nNumOfElements
	}
	if argc == 2 {
		var ret *zend.Zval = nil
		if args[0].value.arr.nNumOfElements == 0 {
			ret = &args[1]
		} else if args[1].value.arr.nNumOfElements == 0 {
			ret = &args[0]
		}
		if ret != nil {
			if (ret.value.arr.u.flags & 1 << 2) != 0 {
				if ret.value.arr.nNumUsed == ret.value.arr.nNumOfElements {
					var _z1 *zend.Zval = return_value
					var _z2 *zend.Zval = ret
					var _gc *zend.ZendRefcounted = _z2.value.counted
					var _t uint32 = _z2.u1.type_info
					_z1.value.counted = _gc
					_z1.u1.type_info = _t
					if (_t & 0xff00) != 0 {
						zend.ZendGcAddref(&_gc.gc)
					}
					return
				}
			} else {
				var copy zend.ZendBool = 1
				var string_key *zend.ZendString
				for {
					var __ht *zend.HashTable = ret.value.arr
					var _p *zend.Bucket = __ht.arData
					var _end *zend.Bucket = _p + __ht.nNumUsed
					for ; _p != _end; _p++ {
						var _z *zend.Zval = &_p.val

						if _z.u1.v.type_ == 0 {
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
					var _z1 *zend.Zval = return_value
					var _z2 *zend.Zval = ret
					var _gc *zend.ZendRefcounted = _z2.value.counted
					var _t uint32 = _z2.u1.type_info
					_z1.value.counted = _gc
					_z1.u1.type_info = _t
					if (_t & 0xff00) != 0 {
						zend.ZendGcAddref(&_gc.gc)
					}
					return
				}
			}
		}
	}
	arg = args
	src = arg.value.arr

	/* copy first array */

	var __arr *zend.ZendArray = zend._zendNewArray(count)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	dest = return_value.value.arr
	if (src.u.flags & 1 << 2) != 0 {
		zend.ZendHashRealInitPacked(dest)
		for {
			var __fill_ht *zend.HashTable = dest
			var __fill_bkt *zend.Bucket = __fill_ht.arData + __fill_ht.nNumUsed
			var __fill_idx uint32 = __fill_ht.nNumUsed
			assert((__fill_ht.u.flags & 1 << 2) != 0)
			for {
				var __ht *zend.HashTable = src
				var _p *zend.Bucket = __ht.arData
				var _end *zend.Bucket = _p + __ht.nNumUsed
				for ; _p != _end; _p++ {
					var _z *zend.Zval = &_p.val

					if _z.u1.v.type_ == 0 {
						continue
					}
					src_entry = _z
					if src_entry.u1.v.type_ == 10 && zend.ZvalRefcountP(src_entry) == 1 {
						src_entry = &(*src_entry).value.ref.val
					}
					if src_entry.u1.v.type_flags != 0 {
						zend.ZvalAddrefP(src_entry)
					}
					var _z1 *zend.Zval = &__fill_bkt.val
					var _z2 *zend.Zval = src_entry
					var _gc *zend.ZendRefcounted = _z2.value.counted
					var _t uint32 = _z2.u1.type_info
					_z1.value.counted = _gc
					_z1.u1.type_info = _t
					__fill_bkt.h = __fill_idx
					__fill_bkt.key = nil
					__fill_bkt++
					__fill_idx++
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

				if _z.u1.v.type_ == 0 {
					continue
				}
				string_key = _p.key
				src_entry = _z
				if src_entry.u1.v.type_ == 10 && zend.ZvalRefcountP(src_entry) == 1 {
					src_entry = &(*src_entry).value.ref.val
				}
				if src_entry.u1.v.type_flags != 0 {
					zend.ZvalAddrefP(src_entry)
				}
				if string_key != nil {
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
			PhpArrayMergeRecursive(dest, arg.value.arr)
		}
	} else {
		for i = 1; i < argc; i++ {
			arg = args + i
			PhpArrayMerge(dest, arg.value.arr)
		}
	}
}

/* }}} */

func ZifArrayMerge(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpArrayMergeWrapper(execute_data, return_value, 0)
}

/* }}} */

func ZifArrayMergeRecursive(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpArrayMergeWrapper(execute_data, return_value, 1)
}

/* }}} */

func ZifArrayReplace(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpArrayReplaceWrapper(execute_data, return_value, 0)
}

/* }}} */

func ZifArrayReplaceRecursive(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpArrayReplaceWrapper(execute_data, return_value, 1)
}

/* }}} */

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
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgArray(_arg, &input, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = 4
				break
			}
			_optional = 1
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			zend.ZendParseArgZvalDeref(_arg, &search_value, 0)
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgBool(_arg, &strict, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
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
	arrval = input.value.arr
	elem_count = arrval.nNumOfElements

	/* Base case: empty input */

	if elem_count == 0 {
		var __z *zend.Zval = return_value
		var __zv *zend.Zval = input
		if __zv.u1.v.type_ != 10 {
			var _z1 *zend.Zval = __z
			var _z2 *zend.Zval = __zv
			var _gc *zend.ZendRefcounted = _z2.value.counted
			var _t uint32 = _z2.u1.type_info
			_z1.value.counted = _gc
			_z1.u1.type_info = _t
			if (_t & 0xff00) != 0 {
				zend.ZendGcAddref(&_gc.gc)
			}
		} else {
			var _z1 *zend.Zval = __z
			var _z2 *zend.Zval = &(*__zv).value.ref.val
			var _gc *zend.ZendRefcounted = _z2.value.counted
			var _t uint32 = _z2.u1.type_info
			_z1.value.counted = _gc
			_z1.u1.type_info = _t
			if (_t & 0xff00) != 0 {
				zend.ZendGcAddref(&_gc.gc)
			}

		}
		return
	}

	/* Initialize return array */

	if search_value != nil {
		var __arr *zend.ZendArray = zend._zendNewArray(0)
		var __z *zend.Zval = return_value
		__z.value.arr = __arr
		__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
		if strict != 0 {
			for {
				var __ht *zend.HashTable = arrval
				var _p *zend.Bucket = __ht.arData
				var _end *zend.Bucket = _p + __ht.nNumUsed
				for ; _p != _end; _p++ {
					var _z *zend.Zval = &_p.val
					if _z.u1.v.type_ == 13 {
						_z = _z.value.zv
					}
					if _z.u1.v.type_ == 0 {
						continue
					}
					num_idx = _p.h
					str_idx = _p.key
					entry = _z
					if entry.u1.v.type_ == 10 {
						entry = &(*entry).value.ref.val
					}
					if zend.FastIsIdenticalFunction(search_value, entry) != 0 {
						if str_idx != nil {
							var __z *zend.Zval = &new_val
							var __s *zend.ZendString = str_idx
							__z.value.str = __s
							if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
								__z.u1.type_info = 6
							} else {
								zend.ZendGcAddref(&__s.gc)
								__z.u1.type_info = 6 | 1<<0<<8
							}
						} else {
							var __z *zend.Zval = &new_val
							__z.value.lval = num_idx
							__z.u1.type_info = 4
						}
						zend.ZendHashNextIndexInsertNew(return_value.value.arr, &new_val)
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
					if _z.u1.v.type_ == 13 {
						_z = _z.value.zv
					}
					if _z.u1.v.type_ == 0 {
						continue
					}
					num_idx = _p.h
					str_idx = _p.key
					entry = _z
					if zend.FastEqualCheckFunction(search_value, entry) != 0 {
						if str_idx != nil {
							var __z *zend.Zval = &new_val
							var __s *zend.ZendString = str_idx
							__z.value.str = __s
							if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
								__z.u1.type_info = 6
							} else {
								zend.ZendGcAddref(&__s.gc)
								__z.u1.type_info = 6 | 1<<0<<8
							}
						} else {
							var __z *zend.Zval = &new_val
							__z.value.lval = num_idx
							__z.u1.type_info = 4
						}
						zend.ZendHashNextIndexInsertNew(return_value.value.arr, &new_val)
					}
				}
				break
			}
		}
	} else {
		var __arr *zend.ZendArray = zend._zendNewArray(elem_count)
		var __z *zend.Zval = return_value
		__z.value.arr = __arr
		__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
		zend.ZendHashRealInitPacked(return_value.value.arr)
		for {
			var __fill_ht *zend.HashTable = return_value.value.arr
			var __fill_bkt *zend.Bucket = __fill_ht.arData + __fill_ht.nNumUsed
			var __fill_idx uint32 = __fill_ht.nNumUsed
			assert((__fill_ht.u.flags & 1 << 2) != 0)
			if (arrval.u.flags&1<<2) != 0 && arrval.nNumUsed == arrval.nNumOfElements {

				/* Optimistic case: range(0..n-1) for vector-like packed array */

				var lval zend.ZendUlong = 0
				for ; lval < elem_count; lval++ {
					var __z *zend.Zval = &__fill_bkt.val
					__z.value.lval = lval
					__z.u1.type_info = 4
					__fill_bkt.h = __fill_idx
					__fill_bkt.key = nil
					__fill_bkt++
					__fill_idx++
				}
			} else {

				/* Go through input array and add keys to the return array */

				for {
					var __ht *zend.HashTable = input.value.arr
					var _p *zend.Bucket = __ht.arData
					var _end *zend.Bucket = _p + __ht.nNumUsed
					for ; _p != _end; _p++ {
						var _z *zend.Zval = &_p.val
						if _z.u1.v.type_ == 13 {
							_z = _z.value.zv
						}
						if _z.u1.v.type_ == 0 {
							continue
						}
						num_idx = _p.h
						str_idx = _p.key
						entry = _z
						if str_idx != nil {
							var __z *zend.Zval = &__fill_bkt.val
							var __s *zend.ZendString = str_idx
							__z.value.str = __s
							if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
								__z.u1.type_info = 6
							} else {
								zend.ZendGcAddref(&__s.gc)
								__z.u1.type_info = 6 | 1<<0<<8
							}
						} else {
							var __z *zend.Zval = &__fill_bkt.val
							__z.value.lval = num_idx
							__z.u1.type_info = 4
						}
						__fill_bkt.h = __fill_idx
						__fill_bkt.key = nil
						__fill_bkt++
						__fill_idx++
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

/* }}} */

func ZifArrayKeyFirst(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var stack *zend.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgArray(_arg, &stack, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
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
	var target_hash *zend.HashTable = stack.value.arr
	var pos zend.HashPosition = 0
	zend.ZendHashGetCurrentKeyZvalEx(target_hash, return_value, &pos)
}

/* }}} */

func ZifArrayKeyLast(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var stack *zend.Zval
	var pos zend.HashPosition
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgArray(_arg, &stack, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
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
	var target_hash *zend.HashTable = stack.value.arr
	zend.ZendHashInternalPointerEndEx(target_hash, &pos)
	zend.ZendHashGetCurrentKeyZvalEx(target_hash, return_value, &pos)
}

/* }}} */

func ZifArrayValues(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var input *zend.Zval
	var entry *zend.Zval
	var arrval *zend.ZendArray
	var arrlen zend.ZendLong
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgArray(_arg, &input, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
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
	arrval = input.value.arr

	/* Return empty input as is */

	arrlen = arrval.nNumOfElements
	if arrlen == 0 {
		var __z *zend.Zval = return_value
		__z.value.arr = (*zend.ZendArray)(&zend.ZendEmptyArray)
		__z.u1.type_info = 7
		return
	}

	/* Return vector-like packed arrays as-is */

	if (arrval.u.flags&1<<2) != 0 && arrval.nNumUsed == arrval.nNumOfElements && arrval.nNextFreeElement == arrlen {
		var __z *zend.Zval = return_value
		var __zv *zend.Zval = input
		if __zv.u1.v.type_ != 10 {
			var _z1 *zend.Zval = __z
			var _z2 *zend.Zval = __zv
			var _gc *zend.ZendRefcounted = _z2.value.counted
			var _t uint32 = _z2.u1.type_info
			_z1.value.counted = _gc
			_z1.u1.type_info = _t
			if (_t & 0xff00) != 0 {
				zend.ZendGcAddref(&_gc.gc)
			}
		} else {
			var _z1 *zend.Zval = __z
			var _z2 *zend.Zval = &(*__zv).value.ref.val
			var _gc *zend.ZendRefcounted = _z2.value.counted
			var _t uint32 = _z2.u1.type_info
			_z1.value.counted = _gc
			_z1.u1.type_info = _t
			if (_t & 0xff00) != 0 {
				zend.ZendGcAddref(&_gc.gc)
			}

		}
		return
	}

	/* Initialize return array */

	var __arr *zend.ZendArray = zend._zendNewArray(arrval.nNumOfElements)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	zend.ZendHashRealInitPacked(return_value.value.arr)

	/* Go through input array and add values to the return array */

	for {
		var __fill_ht *zend.HashTable = return_value.value.arr
		var __fill_bkt *zend.Bucket = __fill_ht.arData + __fill_ht.nNumUsed
		var __fill_idx uint32 = __fill_ht.nNumUsed
		assert((__fill_ht.u.flags & 1 << 2) != 0)
		for {
			var __ht *zend.HashTable = arrval
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val

				if _z.u1.v.type_ == 0 {
					continue
				}
				entry = _z
				if entry.u1.v.type_ == 10 && zend.ZvalRefcountP(entry) == 1 {
					entry = &(*entry).value.ref.val
				}
				if entry.u1.v.type_flags != 0 {
					zend.ZvalAddrefP(entry)
				}
				var _z1 *zend.Zval = &__fill_bkt.val
				var _z2 *zend.Zval = entry
				var _gc *zend.ZendRefcounted = _z2.value.counted
				var _t uint32 = _z2.u1.type_info
				_z1.value.counted = _gc
				_z1.u1.type_info = _t
				__fill_bkt.h = __fill_idx
				__fill_bkt.key = nil
				__fill_bkt++
				__fill_idx++
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

/* }}} */

func ZifArrayCountValues(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var input *zend.Zval
	var entry *zend.Zval
	var tmp *zend.Zval
	var myht *zend.HashTable
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgArray(_arg, &input, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
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

	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8

	/* Go through input array and add values to the return array */

	myht = input.value.arr
	for {
		var __ht *zend.HashTable = myht
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if _z.u1.v.type_ == 0 {
				continue
			}
			entry = _z
			if entry.u1.v.type_ == 10 {
				entry = &(*entry).value.ref.val
			}
			if entry.u1.v.type_ == 4 {
				if g.Assign(&tmp, zend.ZendHashIndexFind(return_value.value.arr, entry.value.lval)) == nil {
					var data zend.Zval
					var __z *zend.Zval = &data
					__z.value.lval = 1
					__z.u1.type_info = 4
					zend.ZendHashIndexUpdate(return_value.value.arr, entry.value.lval, &data)
				} else {
					tmp.value.lval++
				}
			} else if entry.u1.v.type_ == 6 {
				if g.Assign(&tmp, zend.ZendSymtableFind(return_value.value.arr, entry.value.str)) == nil {
					var data zend.Zval
					var __z *zend.Zval = &data
					__z.value.lval = 1
					__z.u1.type_info = 4
					zend.ZendSymtableUpdate(return_value.value.arr, entry.value.str, &data)
				} else {
					tmp.value.lval++
				}
			} else {
				core.PhpErrorDocref(nil, 1<<1, "Can only count STRING and INTEGER values!")
			}
		}
		break
	}
}

/* }}} */

func ArrayColumnParamHelper(param *zend.Zval, name string) zend.ZendBool {
	switch param.u1.v.type_ {
	case 5:
		if param.u1.v.type_ != 4 {
			zend.ConvertToLong(param)
		}
	case 4:
		return 1
	case 8:
		if zend.TryConvertToString(param) == 0 {
			return 0
		}
	case 6:
		return 1
	default:
		core.PhpErrorDocref(nil, 1<<1, "The %s key should be either a string or an integer", name)
		return 0
	}
}

/* }}} */

func ArrayColumnFetchProp(data *zend.Zval, name *zend.Zval, rv *zend.Zval) *zend.Zval {
	var prop *zend.Zval = nil
	if data.u1.v.type_ == 8 {

		/* The has_property check is first performed in "exists" mode (which returns true for
		 * properties that are null but exist) and then in "has" mode to handle objects that
		 * implement __isset (which is not called in "exists" mode). */

		if data.value.obj.handlers.has_property(data, name, 0x2, nil) != 0 || data.value.obj.handlers.has_property(data, name, 0x0, nil) != 0 {
			prop = data.value.obj.handlers.read_property(data, name, 0, nil, rv)
			if prop != nil {
				if prop.u1.v.type_ == 10 {
					prop = &(*prop).value.ref.val
				}
				if prop != rv {
					if prop.u1.v.type_flags != 0 {
						zend.ZvalAddrefP(prop)
					}
				}
			}
		}

		/* The has_property check is first performed in "exists" mode (which returns true for
		 * properties that are null but exist) and then in "has" mode to handle objects that
		 * implement __isset (which is not called in "exists" mode). */

	} else if data.u1.v.type_ == 7 {
		if name.u1.v.type_ == 6 {
			prop = zend.ZendSymtableFind(data.value.arr, name.value.str)
		} else if name.u1.v.type_ == 4 {
			prop = zend.ZendHashIndexFind(data.value.arr, name.value.lval)
		}
		if prop != nil {
			if prop.u1.v.type_ == 10 {
				prop = &(*prop).value.ref.val
			}
			if prop.u1.v.type_flags != 0 {
				zend.ZvalAddrefP(prop)
			}
		}
	}
	return prop
}

/* }}} */

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
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgArrayHt(_arg, &input, 0, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = 4
				break
			}
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			zend.ZendParseArgZvalDeref(_arg, &column, 1)
			_optional = 1
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			zend.ZendParseArgZvalDeref(_arg, &index, 1)
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
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
		return_value.u1.type_info = 2
		return
	}
	var __arr *zend.ZendArray = zend._zendNewArray(input.nNumOfElements)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	if index == nil {
		zend.ZendHashRealInitPacked(return_value.value.arr)
		for {
			var __fill_ht *zend.HashTable = return_value.value.arr
			var __fill_bkt *zend.Bucket = __fill_ht.arData + __fill_ht.nNumUsed
			var __fill_idx uint32 = __fill_ht.nNumUsed
			assert((__fill_ht.u.flags & 1 << 2) != 0)
			for {
				var __ht *zend.HashTable = input
				var _p *zend.Bucket = __ht.arData
				var _end *zend.Bucket = _p + __ht.nNumUsed
				for ; _p != _end; _p++ {
					var _z *zend.Zval = &_p.val

					if _z.u1.v.type_ == 0 {
						continue
					}
					data = _z
					if data.u1.v.type_ == 10 {
						data = &(*data).value.ref.val
					}
					if column == nil {
						if data.u1.v.type_flags != 0 {
							zend.ZvalAddrefP(data)
						}
						colval = data
					} else if g.Assign(&colval, ArrayColumnFetchProp(data, column, &rv)) == nil {
						continue
					}
					var _z1 *zend.Zval = &__fill_bkt.val
					var _z2 *zend.Zval = colval
					var _gc *zend.ZendRefcounted = _z2.value.counted
					var _t uint32 = _z2.u1.type_info
					_z1.value.counted = _gc
					_z1.u1.type_info = _t
					__fill_bkt.h = __fill_idx
					__fill_bkt.key = nil
					__fill_bkt++
					__fill_idx++
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

				if _z.u1.v.type_ == 0 {
					continue
				}
				data = _z
				if data.u1.v.type_ == 10 {
					data = &(*data).value.ref.val
				}
				if column == nil {
					if data.u1.v.type_flags != 0 {
						zend.ZvalAddrefP(data)
					}
					colval = data
				} else if g.Assign(&colval, ArrayColumnFetchProp(data, column, &rv)) == nil {
					continue
				}

				/* Failure will leave keyval alone which will land us on the final else block below
				 * which is to append the value as next_index
				 */

				if index != nil {
					var rv zend.Zval
					var keyval *zend.Zval = ArrayColumnFetchProp(data, index, &rv)
					if keyval != nil {
						switch keyval.u1.v.type_ {
						case 6:
							zend.ZendSymtableUpdate(return_value.value.arr, keyval.value.str, colval)
							break
						case 4:
							zend.ZendHashIndexUpdate(return_value.value.arr, keyval.value.lval, colval)
							break
						case 8:
							var tmp_key *zend.ZendString
							var key *zend.ZendString = zend.ZvalGetTmpString(keyval, &tmp_key)
							zend.ZendSymtableUpdate(return_value.value.arr, key, colval)
							zend.ZendTmpStringRelease(tmp_key)
							break
						case 1:
							zend.ZendHashUpdate(return_value.value.arr, zend.ZendEmptyString, colval)
							break
						case 5:
							zend.ZendHashIndexUpdate(return_value.value.arr, zend.ZendDvalToLval(keyval.value.dval), colval)
							break
						case 3:
							zend.ZendHashIndexUpdate(return_value.value.arr, 1, colval)
							break
						case 2:
							zend.ZendHashIndexUpdate(return_value.value.arr, 0, colval)
							break
						case 9:
							zend.ZendHashIndexUpdate(return_value.value.arr, keyval.value.res.handle, colval)
							break
						default:
							zend.ZendHashNextIndexInsert(return_value.value.arr, colval)
							break
						}
						zend.ZvalPtrDtor(keyval)
					} else {
						zend.ZendHashNextIndexInsert(return_value.value.arr, colval)
					}
				} else {
					zend.ZendHashNextIndexInsert(return_value.value.arr, colval)
				}

				/* Failure will leave keyval alone which will land us on the final else block below
				 * which is to append the value as next_index
				 */

			}
			break
		}
	}
}

/* }}} */

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
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgArray(_arg, &input, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = 4
				break
			}
			_optional = 1
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgBool(_arg, &preserve_keys, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
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

	var __arr *zend.ZendArray = zend._zendNewArray(input.value.arr.nNumOfElements)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	if (input.value.arr.u.flags&1<<2) != 0 && preserve_keys == 0 {
		zend.ZendHashRealInitPacked(return_value.value.arr)
		for {
			var __fill_ht *zend.HashTable = return_value.value.arr
			var __fill_bkt *zend.Bucket = __fill_ht.arData + __fill_ht.nNumUsed
			var __fill_idx uint32 = __fill_ht.nNumUsed
			assert((__fill_ht.u.flags & 1 << 2) != 0)
			for {
				var __ht *zend.HashTable = input.value.arr
				var _idx uint32 = __ht.nNumUsed
				var _p *zend.Bucket = __ht.arData + _idx
				var _z *zend.Zval
				for _idx = __ht.nNumUsed; _idx > 0; _idx-- {
					_p--
					_z = &_p.val

					if _z.u1.v.type_ == 0 {
						continue
					}
					entry = _z
					if entry.u1.v.type_ == 10 && zend.ZvalRefcountP(entry) == 1 {
						entry = &(*entry).value.ref.val
					}
					if entry.u1.v.type_flags != 0 {
						zend.ZvalAddrefP(entry)
					}
					var _z1 *zend.Zval = &__fill_bkt.val
					var _z2 *zend.Zval = entry
					var _gc *zend.ZendRefcounted = _z2.value.counted
					var _t uint32 = _z2.u1.type_info
					_z1.value.counted = _gc
					_z1.u1.type_info = _t
					__fill_bkt.h = __fill_idx
					__fill_bkt.key = nil
					__fill_bkt++
					__fill_idx++
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
			var __ht *zend.HashTable = input.value.arr
			var _idx uint32 = __ht.nNumUsed
			var _p *zend.Bucket = __ht.arData + _idx
			var _z *zend.Zval
			for _idx = __ht.nNumUsed; _idx > 0; _idx-- {
				_p--
				_z = &_p.val

				if _z.u1.v.type_ == 0 {
					continue
				}
				num_key = _p.h
				string_key = _p.key
				entry = _z
				if string_key != nil {
					entry = zend.ZendHashAddNew(return_value.value.arr, string_key, entry)
				} else {
					if preserve_keys != 0 {
						entry = zend.ZendHashIndexAddNew(return_value.value.arr, num_key, entry)
					} else {
						entry = zend.ZendHashNextIndexInsertNew(return_value.value.arr, entry)
					}
				}
				zend.ZvalAddRef(entry)
			}
			break
		}
	}
}

/* }}} */

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
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgArray(_arg, &input, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = 4
				break
			}
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgLong(_arg, &pad_size, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = 4
				break
			}
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			zend.ZendParseArgZvalDeref(_arg, &pad_value, 0)
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
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

	input_size = input.value.arr.nNumOfElements
	pad_size_abs = imaxabs(pad_size)
	if pad_size_abs < 0 || pad_size_abs-input_size > 1048576 {
		core.PhpErrorDocref(nil, 1<<1, "You may only pad up to 1048576 elements at a time")
		return_value.u1.type_info = 2
		return
	}
	if input_size >= pad_size_abs {

		/* Copy the original array */

		var _z1 *zend.Zval = return_value
		var _z2 *zend.Zval = input
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
		if (_t & 0xff00) != 0 {
			zend.ZendGcAddref(&_gc.gc)
		}
		return
	}
	num_pads = pad_size_abs - input_size
	if pad_value.u1.v.type_flags != 0 {
		zend.ZendGcAddrefEx(&(pad_value.value.counted).gc, num_pads)
	}
	var __arr *zend.ZendArray = zend._zendNewArray(pad_size_abs)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	if (input.value.arr.u.flags & 1 << 2) != 0 {
		zend.ZendHashRealInitPacked(return_value.value.arr)
		if pad_size < 0 {
			var __fill_ht *zend.HashTable = return_value.value.arr
			var __fill_bkt *zend.Bucket = __fill_ht.arData + __fill_ht.nNumUsed
			var __fill_idx uint32 = __fill_ht.nNumUsed
			assert((__fill_ht.u.flags & 1 << 2) != 0)
			for i = 0; i < num_pads; i++ {
				var _z1 *zend.Zval = &__fill_bkt.val
				var _z2 *zend.Zval = pad_value
				var _gc *zend.ZendRefcounted = _z2.value.counted
				var _t uint32 = _z2.u1.type_info
				_z1.value.counted = _gc
				_z1.u1.type_info = _t
				__fill_bkt.h = __fill_idx
				__fill_bkt.key = nil
				__fill_bkt++
				__fill_idx++
			}
			__fill_ht.nNumUsed = __fill_idx
			__fill_ht.nNumOfElements = __fill_idx
			__fill_ht.nNextFreeElement = __fill_idx
			__fill_ht.nInternalPointer = 0
		}
		for {
			var __fill_ht *zend.HashTable = return_value.value.arr
			var __fill_bkt *zend.Bucket = __fill_ht.arData + __fill_ht.nNumUsed
			var __fill_idx uint32 = __fill_ht.nNumUsed
			assert((__fill_ht.u.flags & 1 << 2) != 0)
			for {
				var __ht *zend.HashTable = input.value.arr
				var _p *zend.Bucket = __ht.arData
				var _end *zend.Bucket = _p + __ht.nNumUsed
				for ; _p != _end; _p++ {
					var _z *zend.Zval = &_p.val

					if _z.u1.v.type_ == 0 {
						continue
					}
					value = _z
					if value.u1.v.type_flags != 0 {
						zend.ZvalAddrefP(value)
					}
					var _z1 *zend.Zval = &__fill_bkt.val
					var _z2 *zend.Zval = value
					var _gc *zend.ZendRefcounted = _z2.value.counted
					var _t uint32 = _z2.u1.type_info
					_z1.value.counted = _gc
					_z1.u1.type_info = _t
					__fill_bkt.h = __fill_idx
					__fill_bkt.key = nil
					__fill_bkt++
					__fill_idx++
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
			var __fill_ht *zend.HashTable = return_value.value.arr
			var __fill_bkt *zend.Bucket = __fill_ht.arData + __fill_ht.nNumUsed
			var __fill_idx uint32 = __fill_ht.nNumUsed
			assert((__fill_ht.u.flags & 1 << 2) != 0)
			for i = 0; i < num_pads; i++ {
				var _z1 *zend.Zval = &__fill_bkt.val
				var _z2 *zend.Zval = pad_value
				var _gc *zend.ZendRefcounted = _z2.value.counted
				var _t uint32 = _z2.u1.type_info
				_z1.value.counted = _gc
				_z1.u1.type_info = _t
				__fill_bkt.h = __fill_idx
				__fill_bkt.key = nil
				__fill_bkt++
				__fill_idx++
			}
			__fill_ht.nNumUsed = __fill_idx
			__fill_ht.nNumOfElements = __fill_idx
			__fill_ht.nNextFreeElement = __fill_idx
			__fill_ht.nInternalPointer = 0
		}
	} else {
		if pad_size < 0 {
			for i = 0; i < num_pads; i++ {
				zend.ZendHashNextIndexInsertNew(return_value.value.arr, pad_value)
			}
		}
		for {
			var __ht *zend.HashTable = input.value.arr
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val
				if _z.u1.v.type_ == 13 {
					_z = _z.value.zv
				}
				if _z.u1.v.type_ == 0 {
					continue
				}
				key = _p.key
				value = _z
				if value.u1.v.type_flags != 0 {
					zend.ZvalAddrefP(value)
				}
				if key != nil {
					zend.ZendHashAddNew(return_value.value.arr, key, value)
				} else {
					zend.ZendHashNextIndexInsertNew(return_value.value.arr, value)
				}
			}
			break
		}
		if pad_size > 0 {
			for i = 0; i < num_pads; i++ {
				zend.ZendHashNextIndexInsertNew(return_value.value.arr, pad_value)
			}
		}
	}
}

/* }}} */

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
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgArray(_arg, &array, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
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
	var __arr *zend.ZendArray = zend._zendNewArray(array.value.arr.nNumOfElements)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	for {
		var __ht *zend.HashTable = array.value.arr
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if _z.u1.v.type_ == 0 {
				continue
			}
			num_idx = _p.h
			str_idx = _p.key
			entry = _z
			if entry.u1.v.type_ == 10 {
				entry = &(*entry).value.ref.val
			}
			if entry.u1.v.type_ == 4 {
				if str_idx != nil {
					var __z *zend.Zval = &data
					var __s *zend.ZendString = str_idx
					__z.value.str = __s
					if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
						__z.u1.type_info = 6
					} else {
						zend.ZendGcAddref(&__s.gc)
						__z.u1.type_info = 6 | 1<<0<<8
					}
				} else {
					var __z *zend.Zval = &data
					__z.value.lval = num_idx
					__z.u1.type_info = 4
				}
				zend.ZendHashIndexUpdate(return_value.value.arr, entry.value.lval, &data)
			} else if entry.u1.v.type_ == 6 {
				if str_idx != nil {
					var __z *zend.Zval = &data
					var __s *zend.ZendString = str_idx
					__z.value.str = __s
					if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
						__z.u1.type_info = 6
					} else {
						zend.ZendGcAddref(&__s.gc)
						__z.u1.type_info = 6 | 1<<0<<8
					}
				} else {
					var __z *zend.Zval = &data
					__z.value.lval = num_idx
					__z.u1.type_info = 4
				}
				zend.ZendSymtableUpdate(return_value.value.arr, entry.value.str, &data)
			} else {
				core.PhpErrorDocref(nil, 1<<1, "Can only flip STRING and INTEGER values!")
			}
		}
		break
	}
}

/* }}} */

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
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgArray(_arg, &array, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = 4
				break
			}
			_optional = 1
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgLong(_arg, &change_to_upper, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
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
	var __arr *zend.ZendArray = zend._zendNewArray(array.value.arr.nNumOfElements)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	for {
		var __ht *zend.HashTable = array.value.arr
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if _z.u1.v.type_ == 0 {
				continue
			}
			num_key = _p.h
			string_key = _p.key
			entry = _z
			if string_key == nil {
				entry = zend.ZendHashIndexUpdate(return_value.value.arr, num_key, entry)
			} else {
				if change_to_upper != 0 {
					new_key = PhpStringToupper(string_key)
				} else {
					new_key = PhpStringTolower(string_key)
				}
				entry = zend.ZendHashUpdate(return_value.value.arr, new_key, entry)
				zend.ZendStringReleaseEx(new_key, 0)
			}
			zend.ZvalAddRef(entry)
		}
		break
	}
}

/* }}} */

// @type Bucketindex struct
func ArrayBucketindexSwap(p any, q any) {
	var f *Bucketindex = (*Bucketindex)(p)
	var g *Bucketindex = (*Bucketindex)(q)
	var t Bucketindex
	t = *f
	*f = *g
	*g = t
}

/* }}} */

func ZifArrayUnique(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var array *zend.Zval
	var idx uint32
	var p *zend.Bucket
	var arTmp *Bucketindex
	var cmpdata *Bucketindex
	var lastkept *Bucketindex
	var i uint
	var sort_type zend.ZendLong = 2
	var cmp zend.CompareFuncT
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgArray(_arg, &array, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = 4
				break
			}
			_optional = 1
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgLong(_arg, &sort_type, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
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
	if array.value.arr.nNumOfElements <= 1 {
		var _z1 *zend.Zval = return_value
		var _z2 *zend.Zval = array
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
		if (_t & 0xff00) != 0 {
			zend.ZendGcAddref(&_gc.gc)
		}
		return
	}
	if sort_type == 2 {
		var seen zend.HashTable
		var num_key zend.ZendLong
		var str_key *zend.ZendString
		var val *zend.Zval
		zend._zendHashInit(&seen, array.value.arr.nNumOfElements, nil, 0)
		var __arr *zend.ZendArray = zend._zendNewArray(0)
		var __z *zend.Zval = return_value
		__z.value.arr = __arr
		__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
		for {
			var __ht *zend.HashTable = array.value.arr
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val
				if _z.u1.v.type_ == 13 {
					_z = _z.value.zv
				}
				if _z.u1.v.type_ == 0 {
					continue
				}
				num_key = _p.h
				str_key = _p.key
				val = _z
				var retval *zend.Zval
				if val.u1.v.type_ == 6 {
					retval = zend.ZendHashAddEmptyElement(&seen, val.value.str)
				} else {
					var tmp_str_val *zend.ZendString
					var str_val *zend.ZendString = zend.ZvalGetTmpString(val, &tmp_str_val)
					retval = zend.ZendHashAddEmptyElement(&seen, str_val)
					zend.ZendTmpStringRelease(tmp_str_val)
				}
				if retval != nil {

					/* First occurrence of the value */

					if val.u1.v.type_ == 10 && zend.ZvalRefcountP(val) == 1 {
						if val.u1.v.type_ == 10 {
							val = &(*val).value.ref.val
						}
					}
					if val.u1.v.type_flags != 0 {
						zend.ZvalAddrefP(val)
					}
					if str_key != nil {
						zend.ZendHashAddNew(return_value.value.arr, str_key, val)
					} else {
						zend.ZendHashIndexAddNew(return_value.value.arr, num_key, val)
					}
				}
			}
			break
		}
		zend.ZendHashDestroy(&seen)
		return
	}
	cmp = PhpGetDataCompareFunc(sort_type, 0)
	var __arr *zend.ZendArray = zend.ZendArrayDup(array.value.arr)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8

	/* create and sort array with pointers to the target_hash buckets */

	arTmp = (*Bucketindex)(g.CondF((zend.ZvalGcFlags(array.value.arr.gc.u.type_info)&1<<7) != 0, func() any {
		return zend.__zendMalloc((array.value.arr.nNumOfElements + 1) * g.SizeOf("struct bucketindex"))
	}, func() any {
		return zend._emalloc((array.value.arr.nNumOfElements + 1) * g.SizeOf("struct bucketindex"))
	}))
	i = 0
	idx = 0
	for ; idx < array.value.arr.nNumUsed; idx++ {
		p = array.value.arr.arData + idx
		if p.val.u1.v.type_ == 0 {
			continue
		}
		if p.val.u1.v.type_ == 13 && p.val.value.zv.u1.v.type_ == 0 {
			continue
		}
		arTmp[i].SetB(*p)
		arTmp[i].SetI(i)
		i++
	}
	&arTmp[i].b.val.u1.type_info = 0
	zend.ZendSort(any(arTmp), i, g.SizeOf("struct bucketindex"), cmp, zend.SwapFuncT(ArrayBucketindexSwap))

	/* go through the sorted array and delete duplicates from the copy */

	lastkept = arTmp
	for cmpdata = arTmp + 1; cmpdata.b.val.u1.v.type_ != 0; cmpdata++ {
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
				zend.ZendHashIndexDel(return_value.value.arr, p.h)
			} else {
				if return_value.value.arr == &zend.EG.symbol_table {
					zend.ZendDeleteGlobalVariable(p.key)
				} else {
					zend.ZendHashDel(return_value.value.arr, p.key)
				}
			}
		}
	}
	g.CondF((zend.ZvalGcFlags(array.value.arr.gc.u.type_info)&1<<7) != 0, func() { return zend.Free(arTmp) }, func() { return zend._efree(arTmp) })
}

/* }}} */

func ZvalCompare(first *zend.Zval, second *zend.Zval) int {
	return zend.StringCompareFunction(first, second)
}

/* }}} */

func ZvalUserCompare(a *zend.Zval, b *zend.Zval) int {
	var args []zend.Zval
	var retval zend.Zval
	var _z1 *zend.Zval = &args[0]
	var _z2 *zend.Zval = a
	var _gc *zend.ZendRefcounted = _z2.value.counted
	var _t uint32 = _z2.u1.type_info
	_z1.value.counted = _gc
	_z1.u1.type_info = _t
	var _z1 *zend.Zval = &args[1]
	var _z2 *zend.Zval = b
	var _gc *zend.ZendRefcounted = _z2.value.counted
	var _t uint32 = _z2.u1.type_info
	_z1.value.counted = _gc
	_z1.u1.type_info = _t
	BasicGlobals.user_compare_fci.param_count = 2
	BasicGlobals.user_compare_fci.params = args
	BasicGlobals.user_compare_fci.retval = &retval
	BasicGlobals.user_compare_fci.no_separation = 0
	if zend.ZendCallFunction(&(BasicGlobals.GetUserCompareFci()), &(BasicGlobals.GetUserCompareFciCache())) == zend.SUCCESS && retval.u1.v.type_ != 0 {
		var ret zend.ZendLong = zend.ZvalGetLong(&retval)
		zend.ZvalPtrDtor(&retval)
		if ret != 0 {
			if ret < 0 {
				return -1
			} else {
				return 1
			}
		} else {
			return 0
		}
	} else {
		return 0
	}
}

/* }}} */

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

	argc = execute_data.This.u2.num_args
	if data_compare_type == 1 {

		/* INTERSECT_COMP_DATA_USER - array_uintersect_assoc() */

		req_args = 3
		param_spec = "+f"
		intersect_data_compare_func = ZvalUserCompare
	} else {

		/*     INTERSECT_COMP_DATA_NONE - array_intersect_key()
		       INTERSECT_COMP_DATA_INTERNAL - array_intersect_assoc() */

		req_args = 2
		param_spec = "+"
		if data_compare_type == 0 {
			intersect_data_compare_func = ZvalCompare
		}
	}
	if argc < req_args {
		core.PhpErrorDocref(nil, 1<<1, "at least %d parameters are required, %d given", req_args, argc)
		return
	}
	if zend.ZendParseParameters(execute_data.This.u2.num_args, param_spec, &args, &argc, &(BasicGlobals.GetUserCompareFci()), &(BasicGlobals.GetUserCompareFciCache())) == zend.FAILURE {
		return
	}
	for i = 0; i < argc; i++ {
		if args[i].u1.v.type_ != 7 {
			core.PhpErrorDocref(nil, 1<<1, "Expected parameter %d to be an array, %s given", i+1, zend.ZendZvalTypeName(&args[i]))
			return_value.u1.type_info = 1
			return
		}
	}
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	for idx = 0; idx < args[0].value.arr.nNumUsed; idx++ {
		p = args[0].value.arr.arData + idx
		val = &p.val
		if val.u1.v.type_ == 0 {
			continue
		}
		if val.u1.v.type_ == 13 {
			val = val.value.zv
			if val.u1.v.type_ == 0 {
				continue
			}
		}
		if val.u1.v.type_ == 10 && zend.ZvalRefcountP(val) == 1 {
			val = &(*val).value.ref.val
		}
		if p.key == nil {
			ok = 1
			for i = 1; i < argc; i++ {
				if g.Assign(&data, zend.ZendHashIndexFind(args[i].value.arr, p.h)) == nil || intersect_data_compare_func != nil && intersect_data_compare_func(val, data) != 0 {
					ok = 0
					break
				}
			}
			if ok != 0 {
				if val.u1.v.type_flags != 0 {
					zend.ZvalAddrefP(val)
				}
				zend.ZendHashIndexUpdate(return_value.value.arr, p.h, val)
			}
		} else {
			ok = 1
			for i = 1; i < argc; i++ {
				if g.Assign(&data, zend.ZendHashFindExInd(args[i].value.arr, p.key, 1)) == nil || intersect_data_compare_func != nil && intersect_data_compare_func(val, data) != 0 {
					ok = 0
					break
				}
			}
			if ok != 0 {
				if val.u1.v.type_flags != 0 {
					zend.ZvalAddrefP(val)
				}
				zend.ZendHashUpdate(return_value.value.arr, p.key, val)
			}
		}
	}
}

/* }}} */

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
	if behavior == 1 {
		intersect_key_compare_func = PhpArrayKeyCompareString
		if data_compare_type == 0 {

			/* array_intersect() */

			req_args = 2
			param_spec = "+"
			intersect_data_compare_func = PhpArrayDataCompareString
		} else if data_compare_type == 1 {

			/* array_uintersect() */

			req_args = 3
			param_spec = "+f"
			intersect_data_compare_func = PhpArrayUserCompare
		} else {
			core.PhpErrorDocref(nil, 1<<1, "data_compare_type is %d. This should never happen. Please report as a bug", data_compare_type)
			return
		}
		if execute_data.This.u2.num_args < req_args {
			core.PhpErrorDocref(nil, 1<<1, "at least %d parameters are required, %d given", req_args, execute_data.This.u2.num_args)
			return
		}
		if zend.ZendParseParameters(execute_data.This.u2.num_args, param_spec, &args, &arr_argc, &fci1, &fci1_cache) == zend.FAILURE {
			return
		}
		fci_data = &fci1
		fci_data_cache = &fci1_cache
	} else if (behavior & 6) != 0 {

		/* INTERSECT_KEY is subset of INTERSECT_ASSOC. When having the former
		 * no comparison of the data is done (part of INTERSECT_ASSOC) */

		if data_compare_type == 0 && key_compare_type == 0 {

			/* array_intersect_assoc() or array_intersect_key() */

			req_args = 2
			param_spec = "+"
			intersect_key_compare_func = PhpArrayKeyCompareString
			intersect_data_compare_func = PhpArrayDataCompareString
		} else if data_compare_type == 1 && key_compare_type == 0 {

			/* array_uintersect_assoc() */

			req_args = 3
			param_spec = "+f"
			intersect_key_compare_func = PhpArrayKeyCompareString
			intersect_data_compare_func = PhpArrayUserCompare
			fci_data = &fci1
			fci_data_cache = &fci1_cache
		} else if data_compare_type == 0 && key_compare_type == 1 {

			/* array_intersect_uassoc() or array_intersect_ukey() */

			req_args = 3
			param_spec = "+f"
			intersect_key_compare_func = PhpArrayUserKeyCompare
			intersect_data_compare_func = PhpArrayDataCompareString
			fci_key = &fci1
			fci_key_cache = &fci1_cache
		} else if data_compare_type == 1 && key_compare_type == 1 {

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
			core.PhpErrorDocref(nil, 1<<1, "data_compare_type is %d. key_compare_type is %d. This should never happen. Please report as a bug", data_compare_type, key_compare_type)
			return
		}
		if execute_data.This.u2.num_args < req_args {
			core.PhpErrorDocref(nil, 1<<1, "at least %d parameters are required, %d given", req_args, execute_data.This.u2.num_args)
			return
		}
		if zend.ZendParseParameters(execute_data.This.u2.num_args, param_spec, &args, &arr_argc, &fci1, &fci1_cache, &fci2, &fci2_cache) == zend.FAILURE {
			return
		}
	} else {
		core.PhpErrorDocref(nil, 1<<1, "behavior is %d. This should never happen. Please report as a bug", behavior)
		return
	}
	old_user_compare_fci = BasicGlobals.GetUserCompareFci()
	old_user_compare_fci_cache = BasicGlobals.GetUserCompareFciCache()
	BasicGlobals.SetUserCompareFciCache(zend.EmptyFcallInfoCache)

	/* for each argument, create and sort list with pointers to the hash buckets */

	lists = (**zend.Bucket)(zend._safeEmalloc(arr_argc, g.SizeOf("Bucket *"), 0))
	ptrs = (**zend.Bucket)(zend._safeEmalloc(arr_argc, g.SizeOf("Bucket *"), 0))
	if behavior == 1 && data_compare_type == 1 {
		BasicGlobals.SetUserCompareFci(*fci_data)
		BasicGlobals.SetUserCompareFciCache(*fci_data_cache)
	} else if (behavior&6) != 0 && key_compare_type == 1 {
		BasicGlobals.SetUserCompareFci(*fci_key)
		BasicGlobals.SetUserCompareFciCache(*fci_key_cache)
	}
	for i = 0; i < arr_argc; i++ {
		if args[i].u1.v.type_ != 7 {
			core.PhpErrorDocref(nil, 1<<1, "Expected parameter %d to be an array, %s given", i+1, zend.ZendZvalTypeName(&args[i]))
			arr_argc = i
			goto out
		}
		hash = args[i].value.arr
		list = (*zend.Bucket)(g.CondF((zend.ZvalGcFlags(hash.gc.u.type_info)&1<<7) != 0, func() any { return zend.__zendMalloc((hash.nNumOfElements + 1) * g.SizeOf("Bucket")) }, func() any { return zend._emalloc((hash.nNumOfElements + 1) * g.SizeOf("Bucket")) }))
		lists[i] = list
		ptrs[i] = list
		for idx = 0; idx < hash.nNumUsed; idx++ {
			p = hash.arData + idx
			if p.val.u1.v.type_ == 0 {
				continue
			}
			g.PostInc(&(*list)) = *p
		}
		&list.val.u1.type_info = 0
		if hash.nNumOfElements > 1 {
			if behavior == 1 {
				zend.ZendSort(any(lists[i]), hash.nNumOfElements, g.SizeOf("Bucket"), intersect_data_compare_func, zend.SwapFuncT(zend.ZendHashBucketSwap))
			} else if (behavior & 6) != 0 {
				zend.ZendSort(any(lists[i]), hash.nNumOfElements, g.SizeOf("Bucket"), intersect_key_compare_func, zend.SwapFuncT(zend.ZendHashBucketSwap))
			}
		}
	}

	/* copy the argument array */

	var __arr *zend.ZendArray = zend.ZendArrayDup(args[0].value.arr)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8

	/* go through the lists and look for common values */

	for ptrs[0].val.u1.v.type_ != 0 {
		if (behavior&6) != 0 && key_compare_type == 1 {
			BasicGlobals.SetUserCompareFci(*fci_key)
			BasicGlobals.SetUserCompareFciCache(*fci_key_cache)
		}
		for i = 1; i < arr_argc; i++ {
			if (behavior & 1) != 0 {
				for ptrs[i].val.u1.v.type_ != 0 && 0 < g.Assign(&c, intersect_data_compare_func(ptrs[0], ptrs[i])) {
					ptrs[i]++
				}
			} else if (behavior & 6) != 0 {
				for ptrs[i].val.u1.v.type_ != 0 && 0 < g.Assign(&c, intersect_key_compare_func(ptrs[0], ptrs[i])) {
					ptrs[i]++
				}
				if c == 0 && ptrs[i].val.u1.v.type_ != 0 && behavior == 6 {

					/* this means that ptrs[i] is not NULL so we can compare
					 * and "c==0" is from last operation
					 * in this branch of code we enter only when INTERSECT_ASSOC
					 * since when we have INTERSECT_KEY compare of data is not wanted. */

					if data_compare_type == 1 {
						BasicGlobals.SetUserCompareFci(*fci_data)
						BasicGlobals.SetUserCompareFciCache(*fci_data_cache)
					}
					if intersect_data_compare_func(ptrs[0], ptrs[i]) != 0 {
						c = 1
						if key_compare_type == 1 {
							BasicGlobals.SetUserCompareFci(*fci_key)
							BasicGlobals.SetUserCompareFciCache(*fci_key_cache)
						}
					}
				}
			}
			if ptrs[i].val.u1.v.type_ == 0 {

				/* delete any values corresponding to remains of ptrs[0] */

				for {
					ptrs[0]++
					p = ptrs[0] - 1
					if p.val.u1.v.type_ == 0 {
						goto out
					}
					if p.key == nil {
						zend.ZendHashIndexDel(return_value.value.arr, p.h)
					} else {
						zend.ZendHashDel(return_value.value.arr, p.key)
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
					zend.ZendHashIndexDel(return_value.value.arr, p.h)
				} else {
					zend.ZendHashDel(return_value.value.arr, p.key)
				}
				if g.PreInc(&ptrs[0]).val.u1.v.type_ == 0 {
					goto out
				}
				if behavior == 1 {
					if 0 <= intersect_data_compare_func(ptrs[0], ptrs[i]) {
						break
					}
				} else if (behavior & 6) != 0 {

					/* no need of looping because indexes are unique */

					break

					/* no need of looping because indexes are unique */

				}
			}

			/* Value of ptrs[0] not in all arguments, delete all entries */

		} else {

			/* ptrs[0] is present in all the arguments */

			for {
				if g.PreInc(&ptrs[0]).val.u1.v.type_ == 0 {
					goto out
				}
				if behavior == 1 {
					if intersect_data_compare_func(ptrs[0]-1, ptrs[0]) != 0 {
						break
					}
				} else if (behavior & 6) != 0 {

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
		hash = args[i].value.arr
		g.CondF((zend.ZvalGcFlags(hash.gc.u.type_info)&1<<7) != 0, func() { return zend.Free(lists[i]) }, func() { return zend._efree(lists[i]) })
	}
	zend.ZendReleaseFcallInfoCache(&(BasicGlobals.GetUserCompareFciCache()))
	BasicGlobals.SetUserCompareFci(old_user_compare_fci)
	BasicGlobals.SetUserCompareFciCache(old_user_compare_fci_cache)
	zend._efree(ptrs)
	zend._efree(lists)
}

/* }}} */

func ZifArrayIntersectKey(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpArrayIntersectKey(execute_data, return_value, -1)
}

/* }}} */

func ZifArrayIntersectUkey(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpArrayIntersect(execute_data, return_value, 2, 0, 1)
}

/* }}} */

func ZifArrayIntersect(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpArrayIntersect(execute_data, return_value, 1, 0, 0)
}

/* }}} */

func ZifArrayUintersect(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpArrayIntersect(execute_data, return_value, 1, 1, 0)
}

/* }}} */

func ZifArrayIntersectAssoc(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpArrayIntersectKey(execute_data, return_value, 0)
}

/* }}} */

func ZifArrayIntersectUassoc(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpArrayIntersect(execute_data, return_value, 6, 0, 1)
}

/* }}} */

func ZifArrayUintersectAssoc(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpArrayIntersectKey(execute_data, return_value, 1)
}

/* }}} */

func ZifArrayUintersectUassoc(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpArrayIntersect(execute_data, return_value, 6, 1, 1)
}

/* }}} */

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

	argc = execute_data.This.u2.num_args
	if data_compare_type == 1 {
		if argc < 3 {
			core.PhpErrorDocref(nil, 1<<1, "at least 3 parameters are required, %d given", execute_data.This.u2.num_args)
			return
		}
		if zend.ZendParseParameters(execute_data.This.u2.num_args, "+f", &args, &argc, &(BasicGlobals.GetUserCompareFci()), &(BasicGlobals.GetUserCompareFciCache())) == zend.FAILURE {
			return
		}
		diff_data_compare_func = ZvalUserCompare
	} else {
		if argc < 2 {
			core.PhpErrorDocref(nil, 1<<1, "at least 2 parameters are required, %d given", execute_data.This.u2.num_args)
			return
		}
		if zend.ZendParseParameters(execute_data.This.u2.num_args, "+", &args, &argc) == zend.FAILURE {
			return
		}
		if data_compare_type == 0 {
			diff_data_compare_func = ZvalCompare
		}
	}
	for i = 0; i < argc; i++ {
		if args[i].u1.v.type_ != 7 {
			core.PhpErrorDocref(nil, 1<<1, "Expected parameter %d to be an array, %s given", i+1, zend.ZendZvalTypeName(&args[i]))
			return_value.u1.type_info = 1
			return
		}
	}
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	for idx = 0; idx < args[0].value.arr.nNumUsed; idx++ {
		p = args[0].value.arr.arData + idx
		val = &p.val
		if val.u1.v.type_ == 0 {
			continue
		}
		if val.u1.v.type_ == 13 {
			val = val.value.zv
			if val.u1.v.type_ == 0 {
				continue
			}
		}
		if val.u1.v.type_ == 10 && zend.ZvalRefcountP(val) == 1 {
			val = &(*val).value.ref.val
		}
		if p.key == nil {
			ok = 1
			for i = 1; i < argc; i++ {
				if g.Assign(&data, zend.ZendHashIndexFind(args[i].value.arr, p.h)) != nil && (diff_data_compare_func == nil || diff_data_compare_func(val, data) == 0) {
					ok = 0
					break
				}
			}
			if ok != 0 {
				if val.u1.v.type_flags != 0 {
					zend.ZvalAddrefP(val)
				}
				zend.ZendHashIndexUpdate(return_value.value.arr, p.h, val)
			}
		} else {
			ok = 1
			for i = 1; i < argc; i++ {
				if g.Assign(&data, zend.ZendHashFindExInd(args[i].value.arr, p.key, 1)) != nil && (diff_data_compare_func == nil || diff_data_compare_func(val, data) == 0) {
					ok = 0
					break
				}
			}
			if ok != 0 {
				if val.u1.v.type_flags != 0 {
					zend.ZvalAddrefP(val)
				}
				zend.ZendHashUpdate(return_value.value.arr, p.key, val)
			}
		}
	}
}

/* }}} */

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
	if behavior == 1 {
		diff_key_compare_func = PhpArrayKeyCompareString
		if data_compare_type == 0 {

			/* array_diff */

			req_args = 2
			param_spec = "+"
			diff_data_compare_func = PhpArrayDataCompareString
		} else if data_compare_type == 1 {

			/* array_udiff */

			req_args = 3
			param_spec = "+f"
			diff_data_compare_func = PhpArrayUserCompare
		} else {
			core.PhpErrorDocref(nil, 1<<1, "data_compare_type is %d. This should never happen. Please report as a bug", data_compare_type)
			return
		}
		if execute_data.This.u2.num_args < req_args {
			core.PhpErrorDocref(nil, 1<<1, "at least %d parameters are required, %d given", req_args, execute_data.This.u2.num_args)
			return
		}
		if zend.ZendParseParameters(execute_data.This.u2.num_args, param_spec, &args, &arr_argc, &fci1, &fci1_cache) == zend.FAILURE {
			return
		}
		fci_data = &fci1
		fci_data_cache = &fci1_cache
	} else if (behavior & 6) != 0 {

		/* DIFF_KEY is subset of DIFF_ASSOC. When having the former
		 * no comparison of the data is done (part of DIFF_ASSOC) */

		if data_compare_type == 0 && key_compare_type == 0 {

			/* array_diff_assoc() or array_diff_key() */

			req_args = 2
			param_spec = "+"
			diff_key_compare_func = PhpArrayKeyCompareString
			diff_data_compare_func = PhpArrayDataCompareString
		} else if data_compare_type == 1 && key_compare_type == 0 {

			/* array_udiff_assoc() */

			req_args = 3
			param_spec = "+f"
			diff_key_compare_func = PhpArrayKeyCompareString
			diff_data_compare_func = PhpArrayUserCompare
			fci_data = &fci1
			fci_data_cache = &fci1_cache
		} else if data_compare_type == 0 && key_compare_type == 1 {

			/* array_diff_uassoc() or array_diff_ukey() */

			req_args = 3
			param_spec = "+f"
			diff_key_compare_func = PhpArrayUserKeyCompare
			diff_data_compare_func = PhpArrayDataCompareString
			fci_key = &fci1
			fci_key_cache = &fci1_cache
		} else if data_compare_type == 1 && key_compare_type == 1 {

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
			core.PhpErrorDocref(nil, 1<<1, "data_compare_type is %d. key_compare_type is %d. This should never happen. Please report as a bug", data_compare_type, key_compare_type)
			return
		}
		if execute_data.This.u2.num_args < req_args {
			core.PhpErrorDocref(nil, 1<<1, "at least %d parameters are required, %d given", req_args, execute_data.This.u2.num_args)
			return
		}
		if zend.ZendParseParameters(execute_data.This.u2.num_args, param_spec, &args, &arr_argc, &fci1, &fci1_cache, &fci2, &fci2_cache) == zend.FAILURE {
			return
		}
	} else {
		core.PhpErrorDocref(nil, 1<<1, "behavior is %d. This should never happen. Please report as a bug", behavior)
		return
	}
	old_user_compare_fci = BasicGlobals.GetUserCompareFci()
	old_user_compare_fci_cache = BasicGlobals.GetUserCompareFciCache()
	BasicGlobals.SetUserCompareFciCache(zend.EmptyFcallInfoCache)

	/* for each argument, create and sort list with pointers to the hash buckets */

	lists = (**zend.Bucket)(zend._safeEmalloc(arr_argc, g.SizeOf("Bucket *"), 0))
	ptrs = (**zend.Bucket)(zend._safeEmalloc(arr_argc, g.SizeOf("Bucket *"), 0))
	if behavior == 1 && data_compare_type == 1 {
		BasicGlobals.SetUserCompareFci(*fci_data)
		BasicGlobals.SetUserCompareFciCache(*fci_data_cache)
	} else if (behavior&6) != 0 && key_compare_type == 1 {
		BasicGlobals.SetUserCompareFci(*fci_key)
		BasicGlobals.SetUserCompareFciCache(*fci_key_cache)
	}
	for i = 0; i < arr_argc; i++ {
		if args[i].u1.v.type_ != 7 {
			core.PhpErrorDocref(nil, 1<<1, "Expected parameter %d to be an array, %s given", i+1, zend.ZendZvalTypeName(&args[i]))
			arr_argc = i
			goto out
		}
		hash = args[i].value.arr
		list = (*zend.Bucket)(g.CondF((zend.ZvalGcFlags(hash.gc.u.type_info)&1<<7) != 0, func() any { return zend.__zendMalloc((hash.nNumOfElements + 1) * g.SizeOf("Bucket")) }, func() any { return zend._emalloc((hash.nNumOfElements + 1) * g.SizeOf("Bucket")) }))
		lists[i] = list
		ptrs[i] = list
		for idx = 0; idx < hash.nNumUsed; idx++ {
			p = hash.arData + idx
			if p.val.u1.v.type_ == 0 {
				continue
			}
			g.PostInc(&(*list)) = *p
		}
		&list.val.u1.type_info = 0
		if hash.nNumOfElements > 1 {
			if behavior == 1 {
				zend.ZendSort(any(lists[i]), hash.nNumOfElements, g.SizeOf("Bucket"), diff_data_compare_func, zend.SwapFuncT(zend.ZendHashBucketSwap))
			} else if (behavior & 6) != 0 {
				zend.ZendSort(any(lists[i]), hash.nNumOfElements, g.SizeOf("Bucket"), diff_key_compare_func, zend.SwapFuncT(zend.ZendHashBucketSwap))
			}
		}
	}

	/* copy the argument array */

	var __arr *zend.ZendArray = zend.ZendArrayDup(args[0].value.arr)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8

	/* go through the lists and look for values of ptr[0] that are not in the others */

	for ptrs[0].val.u1.v.type_ != 0 {
		if (behavior&6) != 0 && key_compare_type == 1 {
			BasicGlobals.SetUserCompareFci(*fci_key)
			BasicGlobals.SetUserCompareFciCache(*fci_key_cache)
		}
		c = 1
		for i = 1; i < arr_argc; i++ {
			var ptr *zend.Bucket = ptrs[i]
			if behavior == 1 {
				for ptrs[i].val.u1.v.type_ != 0 && 0 < g.Assign(&c, diff_data_compare_func(ptrs[0], ptrs[i])) {
					ptrs[i]++
				}
			} else if (behavior & 6) != 0 {
				for ptr.val.u1.v.type_ != 0 && 0 != g.Assign(&c, diff_key_compare_func(ptrs[0], ptr)) {
					ptr++
				}
			}
			if c == 0 {
				if behavior == 1 {
					if ptrs[i].val.u1.v.type_ != 0 {
						ptrs[i]++
					}
					break
				} else if behavior == 6 {

					/* In this branch is execute only when DIFF_ASSOC. If behavior == DIFF_KEY
					 * data comparison is not needed - skipped. */

					if ptr.val.u1.v.type_ != 0 {
						if data_compare_type == 1 {
							BasicGlobals.SetUserCompareFci(*fci_data)
							BasicGlobals.SetUserCompareFciCache(*fci_data_cache)
						}
						if diff_data_compare_func(ptrs[0], ptr) != 0 {

							/* the data is not the same */

							c = -1
							if key_compare_type == 1 {
								BasicGlobals.SetUserCompareFci(*fci_key)
								BasicGlobals.SetUserCompareFciCache(*fci_key_cache)
							}
						} else {
							break
						}
					}

					/* In this branch is execute only when DIFF_ASSOC. If behavior == DIFF_KEY
					 * data comparison is not needed - skipped. */

				} else if behavior == 2 {

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
					zend.ZendHashIndexDel(return_value.value.arr, p.h)
				} else {
					zend.ZendHashDel(return_value.value.arr, p.key)
				}
				if g.PreInc(&ptrs[0]).val.u1.v.type_ == 0 {
					goto out
				}
				if behavior == 1 {
					if diff_data_compare_func(ptrs[0]-1, ptrs[0]) != 0 {
						break
					}
				} else if (behavior & 6) != 0 {

					/* in this case no array_key_compare is needed */

					break

					/* in this case no array_key_compare is needed */

				}
			}

			/* ptrs[0] in one of the other arguments */

		} else {

			/* ptrs[0] in none of the other arguments */

			for {
				if g.PreInc(&ptrs[0]).val.u1.v.type_ == 0 {
					goto out
				}
				if behavior == 1 {
					if diff_data_compare_func(ptrs[0]-1, ptrs[0]) != 0 {
						break
					}
				} else if (behavior & 6) != 0 {

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
		hash = args[i].value.arr
		g.CondF((zend.ZvalGcFlags(hash.gc.u.type_info)&1<<7) != 0, func() { return zend.Free(lists[i]) }, func() { return zend._efree(lists[i]) })
	}
	zend.ZendReleaseFcallInfoCache(&(BasicGlobals.GetUserCompareFciCache()))
	BasicGlobals.SetUserCompareFci(old_user_compare_fci)
	BasicGlobals.SetUserCompareFciCache(old_user_compare_fci_cache)
	zend._efree(ptrs)
	zend._efree(lists)
}

/* }}} */

func ZifArrayDiffKey(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpArrayDiffKey(execute_data, return_value, -1)
}

/* }}} */

func ZifArrayDiffUkey(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpArrayDiff(execute_data, return_value, 2, 0, 1)
}

/* }}} */

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
	if execute_data.This.u2.num_args < 2 {
		core.PhpErrorDocref(nil, 1<<1, "at least 2 parameters are required, %d given", execute_data.This.u2.num_args)
		return
	}
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = -1
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
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
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
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
	if args[0].u1.v.type_ != 7 {
		core.PhpErrorDocref(nil, 1<<1, "Expected parameter 1 to be an array, %s given", zend.ZendZvalTypeName(&args[0]))
		return_value.u1.type_info = 1
		return
	}
	num = args[0].value.arr.nNumOfElements
	if num == 0 {
		for i = 1; i < argc; i++ {
			if args[i].u1.v.type_ != 7 {
				core.PhpErrorDocref(nil, 1<<1, "Expected parameter %d to be an array, %s given", i+1, zend.ZendZvalTypeName(&args[i]))
				return_value.u1.type_info = 1
				return
			}
		}
		var __z *zend.Zval = return_value
		__z.value.arr = (*zend.ZendArray)(&zend.ZendEmptyArray)
		__z.u1.type_info = 7
		return
	} else if num == 1 {
		var found int = 0
		var search_str *zend.ZendString
		var tmp_search_str *zend.ZendString
		value = nil
		for {
			var __ht *zend.HashTable = args[0].value.arr
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val
				if _z.u1.v.type_ == 13 {
					_z = _z.value.zv
				}
				if _z.u1.v.type_ == 0 {
					continue
				}
				value = _z
				break
			}
			break
		}
		if value == nil {
			for i = 1; i < argc; i++ {
				if args[i].u1.v.type_ != 7 {
					core.PhpErrorDocref(nil, 1<<1, "Expected parameter %d to be an array, %s given", i+1, zend.ZendZvalTypeName(&args[i]))
					return_value.u1.type_info = 1
					return
				}
			}
			var __z *zend.Zval = return_value
			__z.value.arr = (*zend.ZendArray)(&zend.ZendEmptyArray)
			__z.u1.type_info = 7
			return
		}
		search_str = zend.ZvalGetTmpString(value, &tmp_search_str)
		for i = 1; i < argc; i++ {
			if args[i].u1.v.type_ != 7 {
				core.PhpErrorDocref(nil, 1<<1, "Expected parameter %d to be an array, %s given", i+1, zend.ZendZvalTypeName(&args[i]))
				return_value.u1.type_info = 1
				return
			}
			if found == 0 {
				for {
					var __ht *zend.HashTable = args[i].value.arr
					var _p *zend.Bucket = __ht.arData
					var _end *zend.Bucket = _p + __ht.nNumUsed
					for ; _p != _end; _p++ {
						var _z *zend.Zval = &_p.val
						if _z.u1.v.type_ == 13 {
							_z = _z.value.zv
						}
						if _z.u1.v.type_ == 0 {
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
			var __z *zend.Zval = return_value
			__z.value.arr = (*zend.ZendArray)(&zend.ZendEmptyArray)
			__z.u1.type_info = 7
		} else {
			var _z1 *zend.Zval = return_value
			var _z2 *zend.Zval = &args[0]
			var _gc *zend.ZendRefcounted = _z2.value.counted
			var _t uint32 = _z2.u1.type_info
			_z1.value.counted = _gc
			_z1.u1.type_info = _t
			if (_t & 0xff00) != 0 {
				zend.ZendGcAddref(&_gc.gc)
			}
		}
		return
	}

	/* count number of elements */

	num = 0
	for i = 1; i < argc; i++ {
		if args[i].u1.v.type_ != 7 {
			core.PhpErrorDocref(nil, 1<<1, "Expected parameter %d to be an array, %s given", i+1, zend.ZendZvalTypeName(&args[i]))
			return_value.u1.type_info = 1
			return
		}
		num += args[i].value.arr.nNumOfElements
	}
	if num == 0 {
		var _z1 *zend.Zval = return_value
		var _z2 *zend.Zval = &args[0]
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
		if (_t & 0xff00) != 0 {
			zend.ZendGcAddref(&_gc.gc)
		}
		return
	}
	&dummy.u1.type_info = 1

	/* create exclude map */

	zend._zendHashInit(&exclude, num, nil, 0)
	for i = 1; i < argc; i++ {
		for {
			var __ht *zend.HashTable = args[i].value.arr
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val
				if _z.u1.v.type_ == 13 {
					_z = _z.value.zv
				}
				if _z.u1.v.type_ == 0 {
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

	var __arr *zend.ZendArray = zend._zendNewArray(args[0].value.arr.nNumOfElements)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	for {
		var __ht *zend.HashTable = args[0].value.arr
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val
			if _z.u1.v.type_ == 13 {
				_z = _z.value.zv
			}
			if _z.u1.v.type_ == 0 {
				continue
			}
			idx = _p.h
			key = _p.key
			value = _z
			str = zend.ZvalGetTmpString(value, &tmp_str)
			if zend.ZendHashExists(&exclude, str) == 0 {
				if key != nil {
					value = zend.ZendHashAddNew(return_value.value.arr, key, value)
				} else {
					value = zend.ZendHashIndexAddNew(return_value.value.arr, idx, value)
				}
				zend.ZvalAddRef(value)
			}
			zend.ZendTmpStringRelease(tmp_str)
		}
		break
	}
	zend.ZendHashDestroy(&exclude)
}

/* }}} */

func ZifArrayUdiff(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpArrayDiff(execute_data, return_value, 1, 1, 0)
}

/* }}} */

func ZifArrayDiffAssoc(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpArrayDiffKey(execute_data, return_value, 0)
}

/* }}} */

func ZifArrayDiffUassoc(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpArrayDiff(execute_data, return_value, 6, 0, 1)
}

/* }}} */

func ZifArrayUdiffAssoc(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpArrayDiffKey(execute_data, return_value, 1)
}

/* }}} */

func ZifArrayUdiffUassoc(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpArrayDiff(execute_data, return_value, 6, 1, 1)
}

/* }}} */

// #define MULTISORT_ORDER       0

// #define MULTISORT_TYPE       1

// #define MULTISORT_LAST       2

func PhpMultisortCompare(a any, b any) int {
	var ab *zend.Bucket = *((**zend.Bucket)(a))
	var bb *zend.Bucket = *((**zend.Bucket)(b))
	var r int
	var result zend.ZendLong
	r = 0
	for {
		result = ArrayGlobals.GetMultisortFunc()[r](&ab[r], &bb[r])
		if result != 0 {
			if result > 0 {
				return 1
			} else {
				return -1
			}
		}
		r++
		if ab[r].val.u1.v.type_ == 0 {
			break
		}
	}
	return 0
}

/* }}} */

// #define MULTISORT_ABORT       efree ( func ) ; efree ( arrays ) ; RETURN_FALSE ;

func ArrayBucketPSawp(p any, q any) {
	var t *zend.Bucket
	var f **zend.Bucket = (**zend.Bucket)(p)
	var g **zend.Bucket = (**zend.Bucket)(q)
	t = *f
	*f = *g
	*g = t
}

/* }}} */

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
	var sort_order int = 4
	var sort_type int = 0
	var i int
	var k int
	var n int
	var func_ *zend.CompareFuncT
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = -1
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
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
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
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

	arrays = (**zend.Zval)(zend._ecalloc(argc, g.SizeOf("zval *")))
	for i = 0; i < 2; i++ {
		parse_state[i] = 0
	}
	ArrayGlobals.SetMultisortFunc((*zend.CompareFuncT)(zend._ecalloc(argc, g.SizeOf("compare_func_t"))))
	func_ = ArrayGlobals.GetMultisortFunc()

	/* Here we go through the input arguments and parse them. Each one can
	 * be either an array or a sort flag which follows an array. If not
	 * specified, the sort flags defaults to PHP_SORT_ASC and PHP_SORT_REGULAR
	 * accordingly. There can't be two sort flags of the same type after an
	 * array, and the very first argument has to be an array. */

	for i = 0; i < argc; i++ {
		var arg *zend.Zval = &args[i]
		if arg.u1.v.type_ == 10 {
			arg = &(*arg).value.ref.val
		}
		if arg.u1.v.type_ == 7 {
			var _zv *zend.Zval = arg
			var _arr *zend.ZendArray = _zv.value.arr
			if zend.ZendGcRefcount(&_arr.gc) > 1 {
				if _zv.u1.v.type_flags != 0 {
					zend.ZendGcDelref(&_arr.gc)
				}
				var __arr *zend.ZendArray = zend.ZendArrayDup(_arr)
				var __z *zend.Zval = _zv
				__z.value.arr = __arr
				__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
			}

			/* We see the next array, so we update the sort flags of
			 * the previous array and reset the sort flags. */

			if i > 0 {
				ArrayGlobals.GetMultisortFunc()[num_arrays-1] = PhpGetDataCompareFunc(sort_type, sort_order != 4)
				sort_order = 4
				sort_type = 0
			}
			arrays[g.PostInc(&num_arrays)] = arg

			/* Next one may be an array or a list of sort flags. */

			for k = 0; k < 2; k++ {
				parse_state[k] = 1
			}

			/* Next one may be an array or a list of sort flags. */

		} else if arg.u1.v.type_ == 4 {
			switch arg.value.lval & ^8 {
			case 4:

			case 3:

				/* flag allowed here */

				if parse_state[0] == 1 {

					/* Save the flag and make sure then next arg is not the current flag. */

					if arg.value.lval == 3 {
						sort_order = 3
					} else {
						sort_order = 4
					}
					parse_state[0] = 0
				} else {
					core.PhpErrorDocref(nil, 1<<1, "Argument #%d is expected to be an array or sorting flag that has not already been specified", i+1)
					zend._efree(func_)
					zend._efree(arrays)
					return_value.u1.type_info = 2
					return
				}
				break
			case 0:

			case 1:

			case 2:

			case 6:

			case 5:

				/* flag allowed here */

				if parse_state[1] == 1 {

					/* Save the flag and make sure then next arg is not the current flag. */

					sort_type = int(arg.value.lval)
					parse_state[1] = 0
				} else {
					core.PhpErrorDocref(nil, 1<<1, "Argument #%d is expected to be an array or sorting flag that has not already been specified", i+1)
					zend._efree(func_)
					zend._efree(arrays)
					return_value.u1.type_info = 2
					return
				}
				break
			default:
				core.PhpErrorDocref(nil, 1<<1, "Argument #%d is an unknown sort flag", i+1)
				zend._efree(func_)
				zend._efree(arrays)
				return_value.u1.type_info = 2
				return
				break
			}
		} else {
			core.PhpErrorDocref(nil, 1<<1, "Argument #%d is expected to be an array or a sort flag", i+1)
			zend._efree(func_)
			zend._efree(arrays)
			return_value.u1.type_info = 2
			return
		}
	}

	/* Take care of the last array sort flags. */

	ArrayGlobals.GetMultisortFunc()[num_arrays-1] = PhpGetDataCompareFunc(sort_type, sort_order != 4)

	/* Make sure the arrays are of the same size. */

	array_size = arrays[0].value.arr.nNumOfElements
	for i = 0; i < num_arrays; i++ {
		if arrays[i].value.arr.nNumOfElements != uint32(array_size) {
			core.PhpErrorDocref(nil, 1<<1, "Array sizes are inconsistent")
			zend._efree(func_)
			zend._efree(arrays)
			return_value.u1.type_info = 2
			return
		}
	}

	/* If all arrays are empty we don't need to do anything. */

	if array_size < 1 {
		zend._efree(func_)
		zend._efree(arrays)
		return_value.u1.type_info = 3
		return
	}

	/* Create the indirection array. This array is of size MxN, where
	 * M is the number of entries in each input array and N is the number
	 * of the input arrays + 1. The last column is NULL to indicate the end
	 * of the row. */

	indirect = (**zend.Bucket)(zend._safeEmalloc(array_size, g.SizeOf("Bucket *"), 0))
	for i = 0; i < array_size; i++ {
		indirect[i] = (*zend.Bucket)(zend._safeEmalloc(num_arrays+1, g.SizeOf("Bucket"), 0))
	}
	for i = 0; i < num_arrays; i++ {
		k = 0
		for idx = 0; idx < arrays[i].value.arr.nNumUsed; idx++ {
			p = arrays[i].value.arr.arData + idx
			if p.val.u1.v.type_ == 0 {
				continue
			}
			indirect[k][i] = *p
			k++
		}
	}
	for k = 0; k < array_size; k++ {
		&indirect[k][num_arrays].val.u1.type_info = 0
	}

	/* Do the actual sort magic - bada-bim, bada-boom. */

	zend.ZendSort(indirect, array_size, g.SizeOf("Bucket *"), PhpMultisortCompare, zend.SwapFuncT(ArrayBucketPSawp))

	/* Restructure the arrays based on sorted indirect - this is mostly taken from zend_hash_sort() function. */

	for i = 0; i < num_arrays; i++ {
		var repack int
		hash = arrays[i].value.arr
		hash.nNumUsed = array_size
		hash.nInternalPointer = 0
		repack = !(hash.u.flags & 1 << 2)
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
		} else if (hash.u.flags & 1 << 2) == 0 {
			zend.ZendHashRehash(hash)
		}
	}

	/* Clean up. */

	for i = 0; i < array_size; i++ {
		zend._efree(indirect[i])
	}
	zend._efree(indirect)
	zend._efree(func_)
	zend._efree(arrays)
	return_value.u1.type_info = 3
	return
}
