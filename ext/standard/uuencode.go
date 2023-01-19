// <<generate>>

package standard

import (
	"sik/core"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <ext/standard/uuencode.c>

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
   | Author: Ilia Alshanetsky <ilia@php.net>                              |
   +----------------------------------------------------------------------+
*/

// # include < math . h >

// # include "php.h"

// # include "php_uuencode.h"

// #define PHP_UU_ENC(c) ( ( c ) ? ( ( c ) & 077 ) + ' ' : '`' )

// #define PHP_UU_ENC_C2(c) PHP_UU_ENC ( ( ( * ( c ) << 4 ) & 060 ) | ( ( * ( ( c ) + 1 ) >> 4 ) & 017 ) )

// #define PHP_UU_ENC_C3(c) PHP_UU_ENC ( ( ( * ( c + 1 ) << 2 ) & 074 ) | ( ( * ( ( c ) + 2 ) >> 6 ) & 03 ) )

// #define PHP_UU_DEC(c) ( ( ( c ) - ' ' ) & 077 )

func PhpUuencode(src *byte, src_len int) *zend.ZendString {
	var len_ int = 45
	var p *uint8
	var s *uint8
	var e *uint8
	var ee *uint8
	var dest *zend.ZendString

	/* encoded length is ~ 38% greater than the original
	   Use 1.5 for easier calculation.
	*/

	dest = zend.ZendStringSafeAlloc(src_len/2, 3, 46, 0)
	p = (*uint8)(dest.val)
	s = (*uint8)(src)
	e = s + src_len
	for s+3 < e {
		ee = s + len_
		if ee > e {
			ee = e
			len_ = ee - s
			if len_%3 != 0 {
				ee = s + int(floor(float64(len_/3))*3)
			}
		}
		if len_ != 0 {
			g.PostInc(&(*p)) = (len_ & 077) + ' '
		} else {
			g.PostInc(&(*p)) = '`'
		}
		for s < ee {
			if (*s)>>2 != 0 {
				g.PostInc(&(*p)) = ((*s) >> 2 & 077) + ' '
			} else {
				g.PostInc(&(*p)) = '`'
			}
			if ((*s)<<4&060 | (*(s + 1))>>4&017) != 0 {
				g.PostInc(&(*p)) = (((*s)<<4&060 | (*(s + 1))>>4&017) & 077) + ' '
			} else {
				g.PostInc(&(*p)) = '`'
			}
			if ((*(s + 1))<<2&074 | (*(s + 2))>>6&3) != 0 {
				g.PostInc(&(*p)) = (((*(s + 1))<<2&074 | (*(s + 2))>>6&3) & 077) + ' '
			} else {
				g.PostInc(&(*p)) = '`'
			}
			if ((*(s + 2)) & 077) != 0 {
				g.PostInc(&(*p)) = ((*(s + 2)) & 077 & 077) + ' '
			} else {
				g.PostInc(&(*p)) = '`'
			}
			s += 3
		}
		if len_ == 45 {
			g.PostInc(&(*p)) = '\n'
		}
	}
	if s < e {
		if len_ == 45 {
			if e-s != 0 {
				g.PostInc(&(*p)) = (e - s&077) + ' '
			} else {
				g.PostInc(&(*p)) = '`'
			}
			len_ = 0
		}
		if (*s)>>2 != 0 {
			g.PostInc(&(*p)) = ((*s) >> 2 & 077) + ' '
		} else {
			g.PostInc(&(*p)) = '`'
		}
		if ((*s)<<4&060 | (*(s + 1))>>4&017) != 0 {
			g.PostInc(&(*p)) = (((*s)<<4&060 | (*(s + 1))>>4&017) & 077) + ' '
		} else {
			g.PostInc(&(*p)) = '`'
		}
		if e-s > 1 {
			if ((*(s + 1))<<2&074 | (*(s + 2))>>6&3) != 0 {
				g.PostInc(&(*p)) = (((*(s + 1))<<2&074 | (*(s + 2))>>6&3) & 077) + ' '
			} else {
				g.PostInc(&(*p)) = '`'
			}
		} else {
			if '0' {
				g.PostInc(&(*p)) = ('0' & 077) + ' '
			} else {
				g.PostInc(&(*p)) = '`'
			}
		}
		if e-s > 2 {
			if ((*(s + 2)) & 077) != 0 {
				g.PostInc(&(*p)) = ((*(s + 2)) & 077 & 077) + ' '
			} else {
				g.PostInc(&(*p)) = '`'
			}
		} else {
			if '0' {
				g.PostInc(&(*p)) = ('0' & 077) + ' '
			} else {
				g.PostInc(&(*p)) = '`'
			}
		}
	}
	if len_ < 45 {
		g.PostInc(&(*p)) = '\n'
	}
	if '0' {
		g.PostInc(&(*p)) = ('0' & 077) + ' '
	} else {
		g.PostInc(&(*p)) = '`'
	}
	g.PostInc(&(*p)) = '\n'
	*p = '0'
	dest = zend.ZendStringTruncate(dest, (*byte)(p-dest.val), 0)
	return dest
}

/* }}} */

func PhpUudecode(src *byte, src_len int) *zend.ZendString {
	var len_ int
	var total_len int = 0
	var s *byte
	var e *byte
	var p *byte
	var ee *byte
	var dest *zend.ZendString
	dest = zend.ZendStringAlloc(int(ceil(src_len*0.75)), 0)
	p = dest.val
	s = src
	e = src + src_len
	for s < e {
		if g.Assign(&len_, g.PostInc(&(*s))-' '&077) == 0 {
			break
		}

		/* sanity check */

		if len_ > src_len {
			goto err
		}
		total_len += len_
		ee = s + g.CondF2(len_ == 45, 60, func() int { return int(floor(len_ * 1.33)) })

		/* sanity check */

		if ee > e {
			goto err
		}
		for s < ee {
			if s+4 > e {
				goto err
			}
			g.PostInc(&(*p)) = ((*s)-' '&077)<<2 | ((*(s + 1))-' '&077)>>4
			g.PostInc(&(*p)) = ((*(s + 1))-' '&077)<<4 | ((*(s + 2))-' '&077)>>2
			g.PostInc(&(*p)) = ((*(s + 2))-' '&077)<<6 | (*(s + 3)) - ' '&077
			s += 4
		}
		if len_ < 45 {
			break
		}

		/* skip \n */

		s++

		/* skip \n */

	}
	assert(p >= dest.val)
	if g.Assign(&len_, total_len) > size_t(p-dest.val) {
		g.PostInc(&(*p)) = ((*s)-' '&077)<<2 | ((*(s + 1))-' '&077)>>4
		if len_ > 1 {
			g.PostInc(&(*p)) = ((*(s + 1))-' '&077)<<4 | ((*(s + 2))-' '&077)>>2
			if len_ > 2 {
				g.PostInc(&(*p)) = ((*(s + 2))-' '&077)<<6 | (*(s + 3)) - ' '&077
			}
		}
	}
	dest.len_ = total_len
	dest.val[dest.len_] = '0'
	return dest
err:
	zend.ZendStringEfree(dest)
	return nil
}

/* }}} */

func ZifConvertUuencode(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var src *zend.ZendString
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

			if zend.ZendParseArgStr(_arg, &src, 0) == 0 {
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
			return_value.u1.type_info = 2
			return
		}
		break
	}
	if src.len_ < 1 {
		return_value.u1.type_info = 2
		return
	}
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = PhpUuencode(src.val, src.len_)
	__z.value.str = __s
	if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
		__z.u1.type_info = 6
	} else {
		__z.u1.type_info = 6 | 1<<0<<8
	}
	return
}

/* }}} */

func ZifConvertUudecode(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var src *zend.ZendString
	var dest *zend.ZendString
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

			if zend.ZendParseArgStr(_arg, &src, 0) == 0 {
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
			return_value.u1.type_info = 2
			return
		}
		break
	}
	if src.len_ < 1 {
		return_value.u1.type_info = 2
		return
	}
	if g.Assign(&dest, PhpUudecode(src.val, src.len_)) == nil {
		core.PhpErrorDocref(nil, 1<<1, "The given parameter is not a valid uuencoded string")
		return_value.u1.type_info = 2
		return
	}
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = dest
	__z.value.str = __s
	if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
		__z.u1.type_info = 6
	} else {
		__z.u1.type_info = 6 | 1<<0<<8
	}
	return
}

/* }}} */
