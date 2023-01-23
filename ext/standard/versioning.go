// <<generate>>

package standard

import (
	r "sik/runtime"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <ext/standard/versioning.c>

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
   | Author: Stig Sæther Bakken <ssb@php.net>                             |
   +----------------------------------------------------------------------+
*/

// # include < stdio . h >

// # include < sys / types . h >

// # include < ctype . h >

// # include < stdlib . h >

// # include < string . h >

// # include "php.h"

// # include "php_versioning.h"

/* {{{ php_canonicalize_version() */

func PhpCanonicalizeVersion(version *byte) *byte {
	var len_ int = strlen(version)
	var buf *byte = zend._safeEmalloc(len_, 2, 1)
	var q *byte
	var lp byte
	var lq byte
	var p *byte
	if len_ == 0 {
		*buf = '0'
		return buf
	}
	p = version
	q = buf
	*p++
	lp = (*p) - 1
	g.PostInc(&(*q)) = lp
	for *p {

		/*  s/[-_+]/./g;
		 *  s/([^\d\.])([^\D\.])/$1.$2/g;
		 *  s/([^\D\.])([^\d\.])/$1.$2/g;
		 */

		// #define isdig(x) ( isdigit ( x ) && ( x ) != '.' )

		// #define isndig(x) ( ! isdigit ( x ) && ( x ) != '.' )

		// #define isspecialver(x) ( ( x ) == '-' || ( x ) == '_' || ( x ) == '+' )

		lq = *(q - 1)
		if (*p) == '-' || (*p) == '_' || (*p) == '+' {
			if lq != '.' {
				g.PostInc(&(*q)) = '.'
			}
		} else if !(isdigit(lp)) && lp != '.' && (isdigit(*p) && (*p) != '.') || isdigit(lp) && lp != '.' && (!(isdigit(*p)) && (*p) != '.') {
			if lq != '.' {
				g.PostInc(&(*q)) = '.'
			}
			g.PostInc(&(*q)) = *p
		} else if !(isalnum(*p)) {
			if lq != '.' {
				g.PostInc(&(*q)) = '.'
			}
		} else {
			g.PostInc(&(*q)) = *p
		}
		*p++
		lp = (*p) - 1
	}
	g.PostInc(&(*q)) = '0'
	return buf
}

/* }}} */

// @type SpecialFormsT struct
func CompareSpecialVersionForms(form1 *byte, form2 *byte) int {
	var found1 int = -1
	var found2 int = -1
	var special_forms []SpecialFormsT = []SpecialFormsT{{"dev", 0}, {"alpha", 1}, {"a", 1}, {"beta", 2}, {"b", 2}, {"RC", 3}, {"rc", 3}, {"#", 4}, {"pl", 5}, {"p", 5}, {nil, 0}}
	var pp *SpecialFormsT
	for pp = special_forms; pp != nil && pp.GetName() != nil; pp++ {
		if strncmp(form1, pp.GetName(), strlen(pp.GetName())) == 0 {
			found1 = pp.GetOrder()
			break
		}
	}
	for pp = special_forms; pp != nil && pp.GetName() != nil; pp++ {
		if strncmp(form2, pp.GetName(), strlen(pp.GetName())) == 0 {
			found2 = pp.GetOrder()
			break
		}
	}
	if found1-found2 != 0 {
		if found1-found2 < 0 {
			return -1
		} else {
			return 1
		}
	} else {
		return 0
	}
}

/* }}} */

func PhpVersionCompare(orig_ver1 *byte, orig_ver2 *byte) int {
	var ver1 *byte
	var ver2 *byte
	var p1 *byte
	var p2 *byte
	var n1 *byte
	var n2 *byte
	var l1 long
	var l2 long
	var compare int = 0
	if !(*orig_ver1) || !(*orig_ver2) {
		if !(*orig_ver1) && !(*orig_ver2) {
			return 0
		} else {
			if *orig_ver1 {
				return 1
			} else {
				return -1
			}
		}
	}
	if orig_ver1[0] == '#' {
		ver1 = zend._estrdup(orig_ver1)
	} else {
		ver1 = PhpCanonicalizeVersion(orig_ver1)
	}
	if orig_ver2[0] == '#' {
		ver2 = zend._estrdup(orig_ver2)
	} else {
		ver2 = PhpCanonicalizeVersion(orig_ver2)
	}
	n1 = ver1
	p1 = n1
	n2 = ver2
	p2 = n2
	for (*p1) && (*p2) && n1 != nil && n2 != nil {
		if g.Assign(&n1, strchr(p1, '.')) != nil {
			*n1 = '0'
		}
		if g.Assign(&n2, strchr(p2, '.')) != nil {
			*n2 = '0'
		}
		if isdigit(*p1) && isdigit(*p2) {

			/* compare element numerically */

			l1 = strtol(p1, nil, 10)
			l2 = strtol(p2, nil, 10)
			if l1-l2 != 0 {
				if l1-l2 < 0 {
					compare = -1
				} else {
					compare = 1
				}
			} else {
				compare = 0
			}
		} else if !(isdigit(*p1)) && !(isdigit(*p2)) {

			/* compare element names */

			compare = CompareSpecialVersionForms(p1, p2)

			/* compare element names */

		} else {

			/* mix of names and numbers */

			if isdigit(*p1) {
				compare = CompareSpecialVersionForms("#N#", p2)
			} else {
				compare = CompareSpecialVersionForms(p1, "#N#")
			}

			/* mix of names and numbers */

		}
		if compare != 0 {
			break
		}
		if n1 != nil {
			p1 = n1 + 1
		}
		if n2 != nil {
			p2 = n2 + 1
		}
	}
	if compare == 0 {
		if n1 != nil {
			if isdigit(*p1) {
				compare = 1
			} else {
				compare = PhpVersionCompare(p1, "#N#")
			}
		} else if n2 != nil {
			if isdigit(*p2) {
				compare = -1
			} else {
				compare = PhpVersionCompare("#N#", p2)
			}
		}
	}
	zend._efree(ver1)
	zend._efree(ver2)
	return compare
}

/* }}} */

func ZifVersionCompare(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var v1 *byte
	var v2 *byte
	var op *byte = nil
	var v1_len int
	var v2_len int
	var op_len int = 0
	var compare int
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgString(_arg, &v1, &v1_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = 4
				break
			}
			_i++
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgString(_arg, &v2, &v2_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = 4
				break
			}
			_optional = 1
			_i++
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgString(_arg, &op, &op_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
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
	compare = PhpVersionCompare(v1, v2)
	if op == nil {
		var __z *zend.Zval = return_value
		__z.value.lval = compare
		__z.u1.type_info = 4
		return
	}
	if !(strncmp(op, "<", op_len)) || !(strncmp(op, "lt", op_len)) {
		if compare == -1 {
			return_value.u1.type_info = 3
		} else {
			return_value.u1.type_info = 2
		}
		return
	}
	if !(strncmp(op, "<=", op_len)) || !(strncmp(op, "le", op_len)) {
		if compare != 1 {
			return_value.u1.type_info = 3
		} else {
			return_value.u1.type_info = 2
		}
		return
	}
	if !(strncmp(op, ">", op_len)) || !(strncmp(op, "gt", op_len)) {
		if compare == 1 {
			return_value.u1.type_info = 3
		} else {
			return_value.u1.type_info = 2
		}
		return
	}
	if !(strncmp(op, ">=", op_len)) || !(strncmp(op, "ge", op_len)) {
		if compare != -1 {
			return_value.u1.type_info = 3
		} else {
			return_value.u1.type_info = 2
		}
		return
	}
	if !(strncmp(op, "==", op_len)) || !(strncmp(op, "=", op_len)) || !(strncmp(op, "eq", op_len)) {
		if compare == 0 {
			return_value.u1.type_info = 3
		} else {
			return_value.u1.type_info = 2
		}
		return
	}
	if !(strncmp(op, "!=", op_len)) || !(strncmp(op, "<>", op_len)) || !(strncmp(op, "ne", op_len)) {
		if compare != 0 {
			return_value.u1.type_info = 3
		} else {
			return_value.u1.type_info = 2
		}
		return
	}
	return_value.u1.type_info = 1
	return
}

/* }}} */
