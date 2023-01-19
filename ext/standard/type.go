// <<generate>>

package standard

import (
	"sik/core"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <ext/standard/type.c>

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
   | Author: Rasmus Lerdorf <rasmus@php.net>                              |
   +----------------------------------------------------------------------+
*/

// # include "php.h"

// # include "php_incomplete_class.h"

/* {{{ proto string gettype(mixed var)
   Returns the type of the variable */

func ZifGettype(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var arg *zend.Zval
	var type_ *zend.ZendString
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

			zend.ZendParseArgZvalDeref(_arg, &arg, 0)
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
	type_ = zend.ZendZvalGetType(arg)
	if type_ != nil {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = type_
		__z.value.str = __s
		__z.u1.type_info = 6
		return
	} else {
		var _s *byte = "unknown type"
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		return
	}
}

/* }}} */

func ZifSettype(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var var_ *zend.Zval
	var type_ *zend.ZendString
	var tmp zend.Zval
	var ptr *zend.Zval
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

			zend.ZendParseArgZvalDeref(_arg, &var_, 0)
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

			if zend.ZendParseArgStr(_arg, &type_, 0) == 0 {
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
	assert(var_.u1.v.type_ == 10)
	if var_.value.ref.sources.ptr != nil {
		var _z1 *zend.Zval = &tmp
		var _z2 *zend.Zval = &(*var_).value.ref.val
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
		if (_t & 0xff00) != 0 {
			zend.ZendGcAddref(&_gc.gc)
		}
		ptr = &tmp
	} else {
		ptr = &(*var_).value.ref.val
	}
	if type_.len_ == g.SizeOf("\"integer\"")-1 && zend.ZendBinaryStrcasecmp(type_.val, type_.len_, "integer", g.SizeOf("\"integer\"")-1) == 0 {
		zend.ConvertToLong(ptr)
	} else if type_.len_ == g.SizeOf("\"int\"")-1 && zend.ZendBinaryStrcasecmp(type_.val, type_.len_, "int", g.SizeOf("\"int\"")-1) == 0 {
		zend.ConvertToLong(ptr)
	} else if type_.len_ == g.SizeOf("\"float\"")-1 && zend.ZendBinaryStrcasecmp(type_.val, type_.len_, "float", g.SizeOf("\"float\"")-1) == 0 {
		zend.ConvertToDouble(ptr)
	} else if type_.len_ == g.SizeOf("\"double\"")-1 && zend.ZendBinaryStrcasecmp(type_.val, type_.len_, "double", g.SizeOf("\"double\"")-1) == 0 {
		zend.ConvertToDouble(ptr)
	} else if type_.len_ == g.SizeOf("\"string\"")-1 && zend.ZendBinaryStrcasecmp(type_.val, type_.len_, "string", g.SizeOf("\"string\"")-1) == 0 {
		if ptr.u1.v.type_ != 6 {
			zend._convertToString(ptr)
		}
	} else if type_.len_ == g.SizeOf("\"array\"")-1 && zend.ZendBinaryStrcasecmp(type_.val, type_.len_, "array", g.SizeOf("\"array\"")-1) == 0 {
		zend.ConvertToArray(ptr)
	} else if type_.len_ == g.SizeOf("\"object\"")-1 && zend.ZendBinaryStrcasecmp(type_.val, type_.len_, "object", g.SizeOf("\"object\"")-1) == 0 {
		zend.ConvertToObject(ptr)
	} else if type_.len_ == g.SizeOf("\"bool\"")-1 && zend.ZendBinaryStrcasecmp(type_.val, type_.len_, "bool", g.SizeOf("\"bool\"")-1) == 0 {
		zend.ConvertToBoolean(ptr)
	} else if type_.len_ == g.SizeOf("\"boolean\"")-1 && zend.ZendBinaryStrcasecmp(type_.val, type_.len_, "boolean", g.SizeOf("\"boolean\"")-1) == 0 {
		zend.ConvertToBoolean(ptr)
	} else if type_.len_ == g.SizeOf("\"null\"")-1 && zend.ZendBinaryStrcasecmp(type_.val, type_.len_, "null", g.SizeOf("\"null\"")-1) == 0 {
		zend.ConvertToNull(ptr)
	} else {
		if ptr == &tmp {
			zend.ZvalPtrDtor(&tmp)
		}
		if type_.len_ == g.SizeOf("\"resource\"")-1 && zend.ZendBinaryStrcasecmp(type_.val, type_.len_, "resource", g.SizeOf("\"resource\"")-1) == 0 {
			core.PhpErrorDocref(nil, 1<<1, "Cannot convert to resource type")
		} else {
			core.PhpErrorDocref(nil, 1<<1, "Invalid type")
		}
		return_value.u1.type_info = 2
		return
	}
	if ptr == &tmp {
		zend.ZendTryAssignTypedRef(var_.value.ref, &tmp)
	}
	return_value.u1.type_info = 3
}

/* }}} */

func ZifIntval(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var num *zend.Zval
	var base zend.ZendLong = 10
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

			zend.ZendParseArgZvalDeref(_arg, &num, 0)
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

			if zend.ZendParseArgLong(_arg, &base, &_dummy, 0, 0) == 0 {
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
	if num.u1.v.type_ != 6 || base == 10 {
		var __z *zend.Zval = return_value
		__z.value.lval = zend.ZvalGetLong(num)
		__z.u1.type_info = 4
		return
	}
	if base == 0 || base == 2 {
		var strval *byte = num.value.str.val
		var strlen int = num.value.str.len_
		for isspace(*strval) && strlen != 0 {
			strval++
			strlen--
		}

		/* Length of 3+ covers "0b#" and "-0b" (which results in 0) */

		if strlen > 2 {
			var offset int = 0
			if strval[0] == '-' || strval[0] == '+' {
				offset = 1
			}
			if strval[offset] == '0' && (strval[offset+1] == 'b' || strval[offset+1] == 'B') {
				var tmpval *byte
				strlen -= 2
				tmpval = zend._emalloc(strlen + 1)

				/* Place the unary symbol at pos 0 if there was one */

				if offset != 0 {
					tmpval[0] = strval[0]
				}

				/* Copy the data from after "0b" to the end of the buffer */

				memcpy(tmpval+offset, strval+offset+2, strlen-offset)
				tmpval[strlen] = 0
				var __z *zend.Zval = return_value
				__z.value.lval = strtoll(tmpval, nil, 2)
				__z.u1.type_info = 4
				zend._efree(tmpval)
				return
			}
		}

		/* Length of 3+ covers "0b#" and "-0b" (which results in 0) */

	}
	var __z *zend.Zval = return_value
	__z.value.lval = strtoll(num.value.str.val, nil, base)
	__z.u1.type_info = 4
}

/* }}} */

func ZifFloatval(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var num *zend.Zval
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

			zend.ZendParseArgZvalDeref(_arg, &num, 0)
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
	var __z *zend.Zval = return_value
	__z.value.dval = zend.ZvalGetDouble(num)
	__z.u1.type_info = 5
	return
}

/* }}} */

func ZifBoolval(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var val *zend.Zval
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
	if zend.ZendIsTrue(val) != 0 {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
	return
}

/* }}} */

func ZifStrval(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var num *zend.Zval
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

			zend.ZendParseArgZvalDeref(_arg, &num, 0)
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
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = zend.ZvalGetString(num)
	__z.value.str = __s
	if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
		__z.u1.type_info = 6
	} else {
		__z.u1.type_info = 6 | 1<<0<<8
	}
}

/* }}} */

func PhpIsType(execute_data *zend.ZendExecuteData, return_value *zend.Zval, type_ int) {
	var arg *zend.Zval
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

			zend.ZendParseArgZvalDeref(_arg, &arg, 0)
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
	if arg.u1.v.type_ == type_ {
		if type_ == 9 {
			var type_name *byte = zend.ZendRsrcListGetRsrcType(arg.value.res)
			if type_name == nil {
				return_value.u1.type_info = 2
				return
			}
		}
		return_value.u1.type_info = 3
		return
	} else {
		return_value.u1.type_info = 2
		return
	}
}

/* {{{ proto bool is_null(mixed var)
   Returns true if variable is null
   Warning: This function is special-cased by zend_compile.c and so is usually bypassed */

func ZifIsNull(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpIsType(execute_data, return_value, 1)
}

/* }}} */

func ZifIsResource(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpIsType(execute_data, return_value, 9)
}

/* }}} */

func ZifIsBool(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var arg *zend.Zval
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

			zend.ZendParseArgZvalDeref(_arg, &arg, 0)
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
	if arg.u1.v.type_ == 2 || arg.u1.v.type_ == 3 {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
	return
}

/* }}} */

func ZifIsInt(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpIsType(execute_data, return_value, 4)
}

/* }}} */

func ZifIsFloat(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpIsType(execute_data, return_value, 5)
}

/* }}} */

func ZifIsString(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpIsType(execute_data, return_value, 6)
}

/* }}} */

func ZifIsArray(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpIsType(execute_data, return_value, 7)
}

/* }}} */

func ZifIsObject(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpIsType(execute_data, return_value, 8)
}

/* }}} */

func ZifIsNumeric(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var arg *zend.Zval
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

			zend.ZendParseArgZvalDeref(_arg, &arg, 0)
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
	switch arg.u1.v.type_ {
	case 4:

	case 5:
		return_value.u1.type_info = 3
		return
		break
	case 6:
		if zend.IsNumericString(arg.value.str.val, arg.value.str.len_, nil, nil, 0) != 0 {
			return_value.u1.type_info = 3
			return
		} else {
			return_value.u1.type_info = 2
			return
		}
		break
	default:
		return_value.u1.type_info = 2
		return
		break
	}
}

/* }}} */

func ZifIsScalar(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var arg *zend.Zval
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

			zend.ZendParseArgZvalDeref(_arg, &arg, 0)
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
	switch arg.u1.v.type_ {
	case 2:

	case 3:

	case 5:

	case 4:

	case 6:
		return_value.u1.type_info = 3
		return
		break
	default:
		return_value.u1.type_info = 2
		return
		break
	}
}

/* }}} */

func ZifIsCallable(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var var_ *zend.Zval
	var callable_name *zend.Zval = nil
	var name *zend.ZendString
	var error *byte
	var retval zend.ZendBool
	var syntax_only zend.ZendBool = 0
	var check_flags int = 0
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

			zend.ZendParseArgZvalDeref(_arg, &var_, 0)
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

			if zend.ZendParseArgBool(_arg, &syntax_only, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
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

			zend.ZendParseArgZvalDeref(_arg, &callable_name, 0)
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
	if syntax_only != 0 {
		check_flags |= 1 << 0
	}
	if execute_data.This.u2.num_args > 2 {
		retval = zend.ZendIsCallableEx(var_, nil, check_flags, &name, nil, &error)
		for {
			assert(callable_name.u1.v.type_ == 10)
			for {
				var _zv *zend.Zval = callable_name
				var ref *zend.ZendReference = _zv.value.ref
				if ref.sources.ptr != nil {
					zend.ZendTryAssignTypedRefStr(ref, name)
					break
				}
				_zv = &ref.val
				zend.ZvalPtrDtor(_zv)
				var __z *zend.Zval = _zv
				var __s *zend.ZendString = name
				__z.value.str = __s
				if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
					__z.u1.type_info = 6
				} else {
					__z.u1.type_info = 6 | 1<<0<<8
				}
				break
			}
			break
		}
	} else {
		retval = zend.ZendIsCallableEx(var_, nil, check_flags, nil, nil, &error)
	}
	if error != nil {

		/* ignore errors */

		zend._efree(error)

		/* ignore errors */

	}
	if retval != 0 {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
	return
}

/* }}} */

func ZifIsIterable(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var var_ *zend.Zval
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

			zend.ZendParseArgZvalDeref(_arg, &var_, 0)
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
	if zend.ZendIsIterable(var_) != 0 {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
	return
}

/* }}} */

func ZifIsCountable(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var var_ *zend.Zval
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

			zend.ZendParseArgZvalDeref(_arg, &var_, 0)
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
	if zend.ZendIsCountable(var_) != 0 {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
	return
}

/* }}} */
