package standard

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core"
	types2 "github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/zpp"
)

func PhpUrlEncodeHashEx(
	ht *types2.Array,
	formstr *zend.SmartStr,
	num_prefix *byte,
	num_prefix_len int,
	key_prefix *byte,
	key_prefix_len int,
	key_suffix string,
	key_suffix_len int,
	type_ *types2.Zval,
	arg_sep *byte,
	enc_type int,
) int {
	var key *types2.String = nil
	var newprefix *byte
	var p *byte
	var prop_name *byte
	var arg_sep_len int
	var newprefix_len int
	var prop_len int
	var idx zend.ZendUlong
	var zdata *types2.Zval = nil
	if ht == nil {
		return types2.FAILURE
	}
	if ht.IsRecursive() {

		/* Prevent recursion */

		return types2.SUCCESS

		/* Prevent recursion */

	}
	if arg_sep == nil {
		arg_sep = zend.INI_STR("arg_separator.output")
		if arg_sep == nil || !(strlen(arg_sep)) {
			arg_sep = URL_DEFAULT_ARG_SEP
		}
	}
	arg_sep_len = strlen(arg_sep)
	var __ht *types2.Array = ht
	for _, _p := range __ht.ForeachData() {
		var _z *types2.Zval = _p.GetVal()

		idx = _p.GetH()
		key = _p.GetKey()
		zdata = _z
		var is_dynamic types2.ZendBool = 1
		if zdata.IsIndirect() {
			zdata = zdata.Indirect()
			if zdata.IsUndef() {
				continue
			}
			is_dynamic = 0
		}

		/* handling for private & protected object properties */

		if key != nil {
			prop_name = key.GetVal()
			prop_len = key.GetLen()
			if type_ != nil && zend.ZendCheckPropertyAccess(type_.Object(), key, is_dynamic) != types2.SUCCESS {

				/* property not visible in this scope */

				continue

				/* property not visible in this scope */

			}
			if key.GetStr()[0] == '0' && type_ != nil {
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
		zdata = types2.ZVAL_DEREF(zdata)
		if zdata.IsType(types2.IS_ARRAY) || zdata.IsType(types2.IS_OBJECT) {
			if key != nil {
				var ekey *types2.String
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
				//types.ZendStringFree(ekey)
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
			if (ht.GetGcFlags() & types2.GC_IMMUTABLE) == 0 {
				ht.ProtectRecursive()
			}
			PhpUrlEncodeHashEx(zend.HASH_OF(zdata), formstr, nil, 0, newprefix, newprefix_len, "%5D", 3, b.Cond(zdata.IsType(types2.IS_OBJECT), zdata, nil), arg_sep, enc_type)
			if (ht.GetGcFlags() & types2.GC_IMMUTABLE) == 0 {
				ht.UnprotectRecursive()
			}
			zend.Efree(newprefix)
		} else if zdata.IsType(types2.IS_NULL) || zdata.IsType(types2.IS_RESOURCE) {

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
				var ekey *types2.String
				if enc_type == PHP_QUERY_RFC3986 {
					ekey = PhpRawUrlEncode(prop_name, prop_len)
				} else {
					ekey = PhpUrlEncode(prop_name, prop_len)
				}
				formstr.AppendString(ekey.GetStr())
				//types.ZendStringFree(ekey)
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
			case types2.IS_STRING:
				var ekey *types2.String
				if enc_type == PHP_QUERY_RFC3986 {
					ekey = PhpRawUrlEncode(zdata.String().GetVal(), zdata.String().GetLen())
				} else {
					ekey = PhpUrlEncode(zdata.String().GetVal(), zdata.String().GetLen())
				}
				formstr.AppendString(ekey.GetStr())
				//types.ZendStringFree(ekey)
			case types2.IS_LONG:
				formstr.AppendLong(zdata.Long())
			case types2.IS_FALSE:
				formstr.AppendString("0")
			case types2.IS_TRUE:
				formstr.AppendString("1")
			default:
				var ekey *types2.String
				var tmp *types2.String
				var str *types2.String = zend.ZvalGetTmpString(zdata, &tmp)
				if enc_type == PHP_QUERY_RFC3986 {
					ekey = PhpRawUrlEncode(str.GetVal(), str.GetLen())
				} else {
					ekey = PhpUrlEncode(str.GetVal(), str.GetLen())
				}
				formstr.AppendString(ekey.GetStr())
				// zend.ZendTmpStringRelease(tmp)
				//types.ZendStringFree(ekey)
			}
		}
	}
	return types2.SUCCESS
}
func ZifHttpBuildQuery(executeData zpp.Ex, return_value zpp.Ret, formdata *types2.Zval, _ zpp.Opt, prefix *types2.Zval, argSeparator *types2.Zval, encType *types2.Zval) {
	var formdata *types2.Zval
	var prefix *byte = nil
	var arg_sep *byte = nil
	var arg_sep_len int = 0
	var prefix_len int = 0
	var formstr zend.SmartStr = zend.MakeSmartStr(0)
	var enc_type zend.ZendLong = PHP_QUERY_RFC1738
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 4, 0)
			formdata = fp.ParseArrayOrObject()
			fp.StartOptional()
			prefix, prefix_len = fp.ParseString()
			arg_sep, arg_sep_len = fp.ParseString()
			enc_type = fp.ParseLong()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	if PhpUrlEncodeHashEx(zend.HASH_OF(formdata), &formstr, prefix, prefix_len, nil, 0, nil, 0, b.Cond(formdata.IsType(types2.IS_OBJECT), formdata, nil), arg_sep, int(enc_type)) == types2.FAILURE {
		if formstr.GetS() != nil {
			formstr.Free()
		}
		return_value.SetFalse()
		return
	}
	if formstr.GetS() == nil {
		return_value.SetStringVal("")
		return
	}
	formstr.ZeroTail()
	return_value.SetString(formstr.GetS())
	return
}
