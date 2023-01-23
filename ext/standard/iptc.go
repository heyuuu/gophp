// <<generate>>

package standard

import (
	"sik/core"
	r "sik/runtime"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <ext/standard/iptc.c>

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
   | Author: Thies C. Arntzen <thies@thieso.net>                          |
   +----------------------------------------------------------------------+
*/

// # include "php.h"

// # include "php_iptc.h"

// # include "ext/standard/head.h"

// # include < sys / stat . h >

// # include < inttypes . h >

/* some defines for the different JPEG block types */

// #define M_SOF0       0xC0

// #define M_SOF1       0xC1

// #define M_SOF2       0xC2

// #define M_SOF3       0xC3

// #define M_SOF5       0xC5

// #define M_SOF6       0xC6

// #define M_SOF7       0xC7

// #define M_SOF9       0xC9

// #define M_SOF10       0xCA

// #define M_SOF11       0xCB

// #define M_SOF13       0xCD

// #define M_SOF14       0xCE

// #define M_SOF15       0xCF

// #define M_SOI       0xD8

// #define M_EOI       0xD9

// #define M_SOS       0xDA

// #define M_APP0       0xe0

// #define M_APP1       0xe1

// #define M_APP2       0xe2

// #define M_APP3       0xe3

// #define M_APP4       0xe4

// #define M_APP5       0xe5

// #define M_APP6       0xe6

// #define M_APP7       0xe7

// #define M_APP8       0xe8

// #define M_APP9       0xe9

// #define M_APP10       0xea

// #define M_APP11       0xeb

// #define M_APP12       0xec

// #define M_APP13       0xed

// #define M_APP14       0xee

// #define M_APP15       0xef

/* {{{ php_iptc_put1
 */

func PhpIptcPut1(fp *r.FILE, spool int, c uint8, spoolbuf **uint8) int {
	if spool > 0 {
		core.PhpOutputWrite((*byte)(&c), 1)
	}
	if spoolbuf != nil {
		g.PostInc(&(*(*spoolbuf))) = c
	}
	return c
}

/* }}} */

func PhpIptcGet1(fp *r.FILE, spool int, spoolbuf **uint8) int {
	var c int
	var cc byte
	c = r.Getc(fp)
	if c == -1 {
		return -1
	}
	if spool > 0 {
		cc = c
		core.PhpOutputWrite((*byte)(&cc), 1)
	}
	if spoolbuf != nil {
		g.PostInc(&(*(*spoolbuf))) = c
	}
	return c
}

/* }}} */

func PhpIptcReadRemaining(fp *r.FILE, spool int, spoolbuf **uint8) int {
	for PhpIptcGet1(fp, spool, spoolbuf) != -1 {
		continue
	}
	return 0xd9
}

/* }}} */

func PhpIptcSkipVariable(fp *r.FILE, spool int, spoolbuf **uint8) int {
	var length uint
	var c1 int
	var c2 int
	if g.Assign(&c1, PhpIptcGet1(fp, spool, spoolbuf)) == -1 {
		return 0xd9
	}
	if g.Assign(&c2, PhpIptcGet1(fp, spool, spoolbuf)) == -1 {
		return 0xd9
	}
	length = (uint8(c1) << 8) + uint8(c2)
	length -= 2
	for g.PostDec(&length) {
		if PhpIptcGet1(fp, spool, spoolbuf) == -1 {
			return 0xd9
		}
	}
	return 0
}

/* }}} */

func PhpIptcNextMarker(fp *r.FILE, spool int, spoolbuf **uint8) int {
	var c int

	/* skip unimportant stuff */

	c = PhpIptcGet1(fp, spool, spoolbuf)
	if c == -1 {
		return 0xd9
	}
	for c != 0xff {
		if g.Assign(&c, PhpIptcGet1(fp, spool, spoolbuf)) == -1 {
			return 0xd9
		}
	}

	/* get marker byte, swallowing possible padding */

	for {
		c = PhpIptcGet1(fp, 0, 0)
		if c == -1 {
			return 0xd9
		} else if c == 0xff {
			PhpIptcPut1(fp, spool, uint8(c), spoolbuf)
		}
		if c != 0xff {
			break
		}
	}
	return uint(c)
}

/* }}} */

var Psheader []byte = "xFFxED00Photoshop 3.008BIMx04x040000"

/* {{{ proto array iptcembed(string iptcdata, string jpeg_file_name [, int spool])
   Embed binary IPTC data into a JPEG image. */

func ZifIptcembed(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var iptcdata *byte
	var jpeg_file *byte
	var iptcdata_len int
	var jpeg_file_len int
	var spool zend.ZendLong = 0
	var fp *r.FILE
	var marker uint
	var done uint = 0
	var inx int
	var spoolbuf *zend.ZendString = nil
	var poi *uint8 = nil
	var sb zend.ZendStatT
	var written zend.ZendBool = 0
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

			if zend.ZendParseArgString(_arg, &iptcdata, &iptcdata_len, 0) == 0 {
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

			if zend.ZendParseArgPath(_arg, &jpeg_file, &jpeg_file_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
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

			if zend.ZendParseArgLong(_arg, &spool, &_dummy, 0, 0) == 0 {
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
	if core.PhpCheckOpenBasedir(jpeg_file) != 0 {
		return_value.u1.type_info = 2
		return
	}
	if iptcdata_len >= SIZE_MAX-g.SizeOf("psheader")-1025 {
		core.PhpErrorDocref(nil, 1<<1, "IPTC data too large")
		return_value.u1.type_info = 2
		return
	}
	if g.Assign(&fp, r.Fopen(jpeg_file, "rb")) == 0 {
		core.PhpErrorDocref(nil, 1<<1, "Unable to open %s", jpeg_file)
		return_value.u1.type_info = 2
		return
	}
	if spool < 2 {
		if fstat(fileno(fp), &sb) != 0 {
			return_value.u1.type_info = 2
			return
		}
		spoolbuf = zend.ZendStringSafeAlloc(1, iptcdata_len+g.SizeOf("psheader")+1024+1, sb.st_size, 0)
		poi = (*uint8)(spoolbuf.val)
		memset(poi, 0, iptcdata_len+g.SizeOf("psheader")+sb.st_size+1024+1)
	}
	if PhpIptcGet1(fp, spool, g.Cond(poi != nil, &poi, 0)) != 0xff {
		r.Fclose(fp)
		if spoolbuf != nil {
			zend.ZendStringEfree(spoolbuf)
		}
		return_value.u1.type_info = 2
		return
	}
	if PhpIptcGet1(fp, spool, g.Cond(poi != nil, &poi, 0)) != 0xd8 {
		r.Fclose(fp)
		if spoolbuf != nil {
			zend.ZendStringEfree(spoolbuf)
		}
		return_value.u1.type_info = 2
		return
	}
	for done == 0 {
		marker = PhpIptcNextMarker(fp, spool, g.Cond(poi != nil, &poi, 0))
		if marker == 0xd9 {
			break
		} else if marker != 0xed {
			PhpIptcPut1(fp, spool, uint8(marker), g.Cond(poi != nil, &poi, 0))
		}
		switch marker {
		case 0xed:

			/* we are going to write a new APP13 marker, so don't output the old one */

			PhpIptcSkipVariable(fp, 0, 0)
			r.Fgetc(fp)
			PhpIptcReadRemaining(fp, spool, g.Cond(poi != nil, &poi, 0))
			done = 1
			break
		case 0xe0:

		case 0xe1:
			if written != 0 {

				/* don't try to write the data twice */

				break

				/* don't try to write the data twice */

			}
			written = 1
			PhpIptcSkipVariable(fp, spool, g.Cond(poi != nil, &poi, 0))
			if (iptcdata_len & 1) != 0 {
				iptcdata_len++
			}
			Psheader[2] = byte(iptcdata_len + 28>>8)
			Psheader[3] = iptcdata_len + 28&0xff
			for inx = 0; inx < 28; inx++ {
				PhpIptcPut1(fp, spool, Psheader[inx], g.Cond(poi != nil, &poi, 0))
			}
			PhpIptcPut1(fp, spool, uint8(iptcdata_len>>8), g.Cond(poi != nil, &poi, 0))
			PhpIptcPut1(fp, spool, uint8(iptcdata_len&0xff), g.Cond(poi != nil, &poi, 0))
			for inx = 0; inx < iptcdata_len; inx++ {
				PhpIptcPut1(fp, spool, iptcdata[inx], g.Cond(poi != nil, &poi, 0))
			}
			break
		case 0xda:

			/* we hit data, no more marker-inserting can be done! */

			PhpIptcReadRemaining(fp, spool, g.Cond(poi != nil, &poi, 0))
			done = 1
			break
		default:
			PhpIptcSkipVariable(fp, spool, g.Cond(poi != nil, &poi, 0))
			break
		}
	}
	r.Fclose(fp)
	if spool < 2 {
		spoolbuf = zend.ZendStringTruncate(spoolbuf, poi-(*uint8)(spoolbuf.val), 0)
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = spoolbuf
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		return
	} else {
		return_value.u1.type_info = 3
		return
	}
}

/* }}} */

func ZifIptcparse(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var inx int = 0
	var len_ int
	var tagsfound uint = 0
	var buffer *uint8
	var recnum uint8
	var dataset uint8
	var str *byte
	var key []*byte
	var str_len int
	var values zend.Zval
	var element *zend.Zval
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgString(_arg, &str, &str_len, 0) == 0 {
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
	buffer = (*uint8)(str)
	for inx < str_len {
		if buffer[inx] == 0x1c && (buffer[inx+1] == 0x1 || buffer[inx+1] == 0x2) {
			break
		} else {
			inx++
		}
	}
	for inx < str_len {
		if buffer[g.PostInc(&inx)] != 0x1c {
			break
		}
		if inx+4 >= str_len {
			break
		}
		dataset = buffer[g.PostInc(&inx)]
		recnum = buffer[g.PostInc(&inx)]
		if (buffer[inx] & uint8(0x80)) != 0 {
			if inx+6 >= str_len {
				break
			}
			len_ = (zend.ZendLong(buffer[inx+2]) << 24) + (zend.ZendLong(buffer[inx+3]) << 16) + (zend.ZendLong(buffer[inx+4]) << 8) + zend.ZendLong(buffer[inx+5])
			inx += 6
		} else {
			len_ = uint16(buffer[inx])<<8 | uint16(buffer[inx+1])
			inx += 2
		}
		if len_ > str_len || inx+len_ > str_len {
			break
		}
		core.ApPhpSnprintf(key, g.SizeOf("key"), "%d#%03d", uint(dataset), uint(recnum))
		if tagsfound == 0 {
			var __arr *zend.ZendArray = zend._zendNewArray(0)
			var __z *zend.Zval = return_value
			__z.value.arr = __arr
			__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
		}
		if g.Assign(&element, zend.ZendHashStrFind(return_value.value.arr, key, strlen(key))) == nil {
			var __arr *zend.ZendArray = zend._zendNewArray(0)
			var __z *zend.Zval = &values
			__z.value.arr = __arr
			__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
			element = zend.ZendHashStrUpdate(return_value.value.arr, key, strlen(key), &values)
		}
		zend.AddNextIndexStringl(element, (*byte)(buffer+inx), len_)
		inx += len_
		tagsfound++
	}
	if tagsfound == 0 {
		return_value.u1.type_info = 2
		return
	}
}

/* }}} */
