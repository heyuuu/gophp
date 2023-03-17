// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/zend"
)

func PhpUrlEncodeHashEx(
	ht *zend.HashTable,
	formstr *zend.SmartStr,
	num_prefix *byte,
	num_prefix_len int,
	key_prefix *byte,
	key_prefix_len int,
	key_suffix string,
	key_suffix_len int,
	type_ *zend.Zval,
	arg_sep *byte,
	enc_type int,
) int {
	var key *zend.ZendString = nil
	var newprefix *byte
	var p *byte
	var prop_name *byte
	var arg_sep_len int
	var newprefix_len int
	var prop_len int
	var idx zend.ZendUlong
	var zdata *zend.Zval = nil
	if ht == nil {
		return zend.FAILURE
	}
	if ht.IsRecursive() {

		/* Prevent recursion */

		return zend.SUCCESS

		/* Prevent recursion */

	}
	if arg_sep == nil {
		arg_sep = zend.INI_STR("arg_separator.output")
		if arg_sep == nil || !(strlen(arg_sep)) {
			arg_sep = URL_DEFAULT_ARG_SEP
		}
	}
	arg_sep_len = strlen(arg_sep)
	var __ht *zend.HashTable = ht
	for _, _p := range __ht.foreachData() {
		var _z *zend.Zval = _p.GetVal()

		idx = _p.GetH()
		key = _p.GetKey()
		zdata = _z
		var is_dynamic zend.ZendBool = 1
		if zdata.IsType(zend.IS_INDIRECT) {
			zdata = zdata.GetZv()
			if zdata.IsUndef() {
				continue
			}
			is_dynamic = 0
		}

		/* handling for private & protected object properties */

		if key != nil {
			prop_name = key.GetVal()
			prop_len = key.GetLen()
			if type_ != nil && zend.ZendCheckPropertyAccess(type_.GetObj(), key, is_dynamic) != zend.SUCCESS {

				/* property not visible in this scope */

				continue

				/* property not visible in this scope */

			}
			if key.GetVal()[0] == '0' && type_ != nil {
				var tmp *byte
				zend.ZendUnmanglePropertyNameEx(key, &tmp, &prop_name, &prop_len)
			} else {
				prop_name = key.GetVal()
				prop_len = key.GetLen()
			}
		} else {
			prop_name = nil
			prop_len = 0
		}
		zdata = zend.ZVAL_DEREF(zdata)
		if zdata.IsType(zend.IS_ARRAY) || zdata.IsType(zend.IS_OBJECT) {
			if key != nil {
				var ekey *zend.ZendString
				if enc_type == PHP_QUERY_RFC3986 {
					ekey = PhpRawUrlEncode(prop_name, prop_len)
				} else {
					ekey = PhpUrlEncode(prop_name, prop_len)
				}
				newprefix_len = key_suffix_len + ekey.GetLen() + key_prefix_len + 3
				newprefix = zend.Emalloc(newprefix_len + 1)
				p = newprefix
				if key_prefix != nil {
					memcpy(p, key_prefix, key_prefix_len)
					p += key_prefix_len
				}
				memcpy(p, ekey.GetVal(), ekey.GetLen())
				p += ekey.GetLen()
				zend.ZendStringFree(ekey)
				if key_suffix {
					memcpy(p, key_suffix, key_suffix_len)
					p += key_suffix_len
				}
				*(b.PostInc(&p)) = '%'
				*(b.PostInc(&p)) = '5'
				*(b.PostInc(&p)) = 'B'
				*p = '0'
			} else {
				var ekey *byte
				var ekey_len int

				/* Is an integer key */

				ekey_len = core.Spprintf(&ekey, 0, zend.ZEND_LONG_FMT, idx)
				newprefix_len = key_prefix_len + num_prefix_len + ekey_len + key_suffix_len + 3
				newprefix = zend.Emalloc(newprefix_len + 1)
				p = newprefix
				if key_prefix != nil {
					memcpy(p, key_prefix, key_prefix_len)
					p += key_prefix_len
				}
				if num_prefix != nil {
					memcpy(p, num_prefix, num_prefix_len)
					p += num_prefix_len
				}
				memcpy(p, ekey, ekey_len)
				p += ekey_len
				zend.Efree(ekey)
				if key_suffix {
					memcpy(p, key_suffix, key_suffix_len)
					p += key_suffix_len
				}
				*(b.PostInc(&p)) = '%'
				*(b.PostInc(&p)) = '5'
				*(b.PostInc(&p)) = 'B'
				*p = '0'
			}
			if (ht.GetGcFlags() & zend.GC_IMMUTABLE) == 0 {
				ht.ProtectRecursive()
			}
			PhpUrlEncodeHashEx(zend.HASH_OF(zdata), formstr, nil, 0, newprefix, newprefix_len, "%5D", 3, b.Cond(zdata.IsType(zend.IS_OBJECT), zdata, nil), arg_sep, enc_type)
			if (ht.GetGcFlags() & zend.GC_IMMUTABLE) == 0 {
				ht.UnprotectRecursive()
			}
			zend.Efree(newprefix)
		} else if zdata.IsType(zend.IS_NULL) || zdata.IsType(zend.IS_RESOURCE) {

			/* Skip these types */

			continue

			/* Skip these types */

		} else {
			if formstr.GetS() != nil {
				formstr.AppendString(b.CastStr(arg_sep, arg_sep_len))
			}

			/* Simple key=value */

			if key_prefix != nil {
				formstr.AppendString(b.CastStr(key_prefix, key_prefix_len))
			}
			if key != nil {
				var ekey *zend.ZendString
				if enc_type == PHP_QUERY_RFC3986 {
					ekey = PhpRawUrlEncode(prop_name, prop_len)
				} else {
					ekey = PhpUrlEncode(prop_name, prop_len)
				}
				formstr.AppendString(ekey.GetStr())
				zend.ZendStringFree(ekey)
			} else {

				/* Numeric key */

				if num_prefix != nil {
					formstr.AppendString(b.CastStr(num_prefix, num_prefix_len))
				}
				formstr.AppendLong(idx)
			}
			if key_suffix {
				formstr.AppendString(b.CastStr(key_suffix, key_suffix_len))
			}
			formstr.AppendString("=")
			switch zdata.GetType() {
			case zend.IS_STRING:
				var ekey *zend.ZendString
				if enc_type == PHP_QUERY_RFC3986 {
					ekey = PhpRawUrlEncode(zdata.GetStr().GetVal(), zdata.GetStr().GetLen())
				} else {
					ekey = PhpUrlEncode(zdata.GetStr().GetVal(), zdata.GetStr().GetLen())
				}
				formstr.AppendString(ekey.GetStr())
				zend.ZendStringFree(ekey)
			case zend.IS_LONG:
				formstr.AppendLong(zdata.GetLval())
			case zend.IS_FALSE:
				formstr.AppendString("0")
			case zend.IS_TRUE:
				formstr.AppendString("1")
			default:
				var ekey *zend.ZendString
				var tmp *zend.ZendString
				var str *zend.ZendString = zend.ZvalGetTmpString(zdata, &tmp)
				if enc_type == PHP_QUERY_RFC3986 {
					ekey = PhpRawUrlEncode(str.GetVal(), str.GetLen())
				} else {
					ekey = PhpUrlEncode(str.GetVal(), str.GetLen())
				}
				formstr.AppendString(ekey.GetStr())
				zend.ZendTmpStringRelease(tmp)
				zend.ZendStringFree(ekey)
			}
		}
	}
	return zend.SUCCESS
}
func ZifHttpBuildQuery(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var formdata *zend.Zval
	var prefix *byte = nil
	var arg_sep *byte = nil
	var arg_sep_len int = 0
	var prefix_len int = 0
	var formstr zend.SmartStr = zend.MakeSmartStr(0)
	var enc_type zend.ZendLong = PHP_QUERY_RFC1738
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 4
		var _num_args int = executeData.NumArgs()
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
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.CheckNumArgsException(_min_num_args, _max_num_args)
					} else {
						zend.CheckNumArgsError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = executeData.Arg(0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgArray(_arg, &formdata, 0, 1) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgString(_arg, &prefix, &prefix_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgString(_arg, &arg_sep, &arg_sep_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &enc_type, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if _error_code != zend.ZPP_ERROR_OK {
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
			return_value.SetFalse()
			return
		}
		break
	}
	if PhpUrlEncodeHashEx(zend.HASH_OF(formdata), &formstr, prefix, prefix_len, nil, 0, nil, 0, b.Cond(formdata.IsType(zend.IS_OBJECT), formdata, nil), arg_sep, int(enc_type)) == zend.FAILURE {
		if formstr.GetS() != nil {
			formstr.Free()
		}
		return_value.SetFalse()
		return
	}
	if formstr.GetS() == nil {
		zend.ZVAL_EMPTY_STRING(return_value)
		return
	}
	formstr.ZeroTail()
	return_value.SetString(formstr.GetS())
	return
}
