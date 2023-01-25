// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	r "sik/runtime"
	"sik/zend"
)

func PhpIptcPut1(fp *r.FILE, spool int, c uint8, spoolbuf **uint8) int {
	if spool > 0 {
		core.PUTC(c)
	}
	if spoolbuf != nil {
		b.PostInc(&(*(*spoolbuf))) = c
	}
	return c
}
func PhpIptcGet1(fp *r.FILE, spool int, spoolbuf **uint8) int {
	var c int
	var cc byte
	c = r.Getc(fp)
	if c == r.EOF {
		return r.EOF
	}
	if spool > 0 {
		cc = c
		core.PUTC(cc)
	}
	if spoolbuf != nil {
		b.PostInc(&(*(*spoolbuf))) = c
	}
	return c
}
func PhpIptcReadRemaining(fp *r.FILE, spool int, spoolbuf **uint8) int {
	for PhpIptcGet1(fp, spool, spoolbuf) != r.EOF {
		continue
	}
	return M_EOI
}
func PhpIptcSkipVariable(fp *r.FILE, spool int, spoolbuf **uint8) int {
	var length uint
	var c1 int
	var c2 int
	if b.Assign(&c1, PhpIptcGet1(fp, spool, spoolbuf)) == r.EOF {
		return M_EOI
	}
	if b.Assign(&c2, PhpIptcGet1(fp, spool, spoolbuf)) == r.EOF {
		return M_EOI
	}
	length = (uint8(c1) << 8) + uint8(c2)
	length -= 2
	for b.PostDec(&length) {
		if PhpIptcGet1(fp, spool, spoolbuf) == r.EOF {
			return M_EOI
		}
	}
	return 0
}
func PhpIptcNextMarker(fp *r.FILE, spool int, spoolbuf **uint8) int {
	var c int

	/* skip unimportant stuff */

	c = PhpIptcGet1(fp, spool, spoolbuf)
	if c == r.EOF {
		return M_EOI
	}
	for c != 0xff {
		if b.Assign(&c, PhpIptcGet1(fp, spool, spoolbuf)) == r.EOF {
			return M_EOI
		}
	}

	/* get marker byte, swallowing possible padding */

	for {
		c = PhpIptcGet1(fp, 0, 0)
		if c == r.EOF {
			return M_EOI
		} else if c == 0xff {
			PhpIptcPut1(fp, spool, uint8(c), spoolbuf)
		}
		if c != 0xff {
			break
		}
	}
	return uint(c)
}
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
			if zend.UNEXPECTED(zend.ZendParseArgString(_arg, &iptcdata, &iptcdata_len, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgPath(_arg, &jpeg_file, &jpeg_file_len, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_PATH
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgLong(_arg, &spool, &_dummy, 0, 0) == 0) {
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
	if core.PhpCheckOpenBasedir(jpeg_file) != 0 {
		zend.RETVAL_FALSE
		return
	}
	if iptcdata_len >= SIZE_MAX-b.SizeOf("psheader")-1025 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "IPTC data too large")
		zend.RETVAL_FALSE
		return
	}
	if b.Assign(&fp, zend.VCWD_FOPEN(jpeg_file, "rb")) == 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Unable to open %s", jpeg_file)
		zend.RETVAL_FALSE
		return
	}
	if spool < 2 {
		if zend.ZendFstat(fileno(fp), &sb) != 0 {
			zend.RETVAL_FALSE
			return
		}
		spoolbuf = zend.ZendStringSafeAlloc(1, iptcdata_len+b.SizeOf("psheader")+1024+1, sb.st_size, 0)
		poi = (*uint8)(zend.ZSTR_VAL(spoolbuf))
		memset(poi, 0, iptcdata_len+b.SizeOf("psheader")+sb.st_size+1024+1)
	}
	if PhpIptcGet1(fp, spool, b.Cond(poi != nil, &poi, 0)) != 0xff {
		r.Fclose(fp)
		if spoolbuf != nil {
			zend.ZendStringEfree(spoolbuf)
		}
		zend.RETVAL_FALSE
		return
	}
	if PhpIptcGet1(fp, spool, b.Cond(poi != nil, &poi, 0)) != 0xd8 {
		r.Fclose(fp)
		if spoolbuf != nil {
			zend.ZendStringEfree(spoolbuf)
		}
		zend.RETVAL_FALSE
		return
	}
	for done == 0 {
		marker = PhpIptcNextMarker(fp, spool, b.Cond(poi != nil, &poi, 0))
		if marker == M_EOI {
			break
		} else if marker != M_APP13 {
			PhpIptcPut1(fp, spool, uint8(marker), b.Cond(poi != nil, &poi, 0))
		}
		switch marker {
		case M_APP13:

			/* we are going to write a new APP13 marker, so don't output the old one */

			PhpIptcSkipVariable(fp, 0, 0)
			r.Fgetc(fp)
			PhpIptcReadRemaining(fp, spool, b.Cond(poi != nil, &poi, 0))
			done = 1
			break
		case M_APP0:

		case M_APP1:
			if written != 0 {

				/* don't try to write the data twice */

				break

				/* don't try to write the data twice */

			}
			written = 1
			PhpIptcSkipVariable(fp, spool, b.Cond(poi != nil, &poi, 0))
			if (iptcdata_len & 1) != 0 {
				iptcdata_len++
			}
			Psheader[2] = byte(iptcdata_len + 28>>8)
			Psheader[3] = iptcdata_len + 28&0xff
			for inx = 0; inx < 28; inx++ {
				PhpIptcPut1(fp, spool, Psheader[inx], b.Cond(poi != nil, &poi, 0))
			}
			PhpIptcPut1(fp, spool, uint8(iptcdata_len>>8), b.Cond(poi != nil, &poi, 0))
			PhpIptcPut1(fp, spool, uint8(iptcdata_len&0xff), b.Cond(poi != nil, &poi, 0))
			for inx = 0; inx < iptcdata_len; inx++ {
				PhpIptcPut1(fp, spool, iptcdata[inx], b.Cond(poi != nil, &poi, 0))
			}
			break
		case M_SOS:

			/* we hit data, no more marker-inserting can be done! */

			PhpIptcReadRemaining(fp, spool, b.Cond(poi != nil, &poi, 0))
			done = 1
			break
		default:
			PhpIptcSkipVariable(fp, spool, b.Cond(poi != nil, &poi, 0))
			break
		}
	}
	r.Fclose(fp)
	if spool < 2 {
		spoolbuf = zend.ZendStringTruncate(spoolbuf, poi-(*uint8)(zend.ZSTR_VAL(spoolbuf)), 0)
		zend.RETVAL_NEW_STR(spoolbuf)
		return
	} else {
		zend.RETVAL_TRUE
		return
	}
}
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
			if zend.UNEXPECTED(zend.ZendParseArgString(_arg, &str, &str_len, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
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
	buffer = (*uint8)(str)
	for inx < str_len {
		if buffer[inx] == 0x1c && (buffer[inx+1] == 0x1 || buffer[inx+1] == 0x2) {
			break
		} else {
			inx++
		}
	}
	for inx < str_len {
		if buffer[b.PostInc(&inx)] != 0x1c {
			break
		}
		if inx+4 >= str_len {
			break
		}
		dataset = buffer[b.PostInc(&inx)]
		recnum = buffer[b.PostInc(&inx)]
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
		core.Snprintf(key, b.SizeOf("key"), "%d#%03d", uint(dataset), uint(recnum))
		if tagsfound == 0 {
			zend.ArrayInit(return_value)
		}
		if b.Assign(&element, zend.ZendHashStrFind(zend.Z_ARRVAL_P(return_value), key, strlen(key))) == nil {
			zend.ArrayInit(&values)
			element = zend.ZendHashStrUpdate(zend.Z_ARRVAL_P(return_value), key, strlen(key), &values)
		}
		zend.AddNextIndexStringl(element, (*byte)(buffer+inx), len_)
		inx += len_
		tagsfound++
	}
	if tagsfound == 0 {
		zend.RETVAL_FALSE
		return
	}
}
