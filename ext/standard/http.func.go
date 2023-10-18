package standard

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/operators"
	"github.com/heyuuu/gophp/zend/zpp"
	"strconv"
)

func PhpUrlEncodeHashEx(
	ht *types.Array,
	formstr *zend.SmartStr,
	num_prefix *byte,
	num_prefix_len int,
	key_prefix *byte,
	key_prefix_len int,
	key_suffix string,
	key_suffix_len int,
	type_ *types.Zval,
	arg_sep *byte,
	enc_type int,
) int {
	var key *types.String = nil
	var newprefix string
	var p *byte
	var prop_name *byte
	var arg_sep_len int
	var newprefix_len int
	var prop_len int
	var idx zend.ZendUlong
	var zdata *types.Zval = nil
	if ht == nil {
		return types.FAILURE
	}
	if ht.IsRecursive() {

		/* Prevent recursion */

		return types.SUCCESS

		/* Prevent recursion */

	}
	if arg_sep == nil {
		arg_sep = zend.INI_STR("arg_separator.output")
		if arg_sep == nil || !(strlen(arg_sep)) {
			arg_sep = URL_DEFAULT_ARG_SEP
		}
	}
	arg_sep_len = strlen(arg_sep)
	var __ht *types.Array = ht
	for _, _p := range __ht.ForeachData() {
		var _z *types.Zval = _p.GetVal()

		idx = _p.GetH()
		key = _p.GetKey()
		zdata = _z
		var is_dynamic bool = 1
		if zdata.IsIndirect() {
			zdata = zdata.Indirect()
			if zdata.IsUndef() {
				continue
			}
			is_dynamic = 0
		}

		/* handling for private & protected object properties */

		if key != nil {
			prop_name = key.GetStr()
			prop_len = key.GetLen()
			if type_ != nil && zend.ZendCheckPropertyAccess(type_.Object(), key, is_dynamic) != types.SUCCESS {

				/* property not visible in this scope */

				continue

				/* property not visible in this scope */

			}
			if key.GetStr()[0] == '0' && type_ != nil {
				var tmp *byte
				zend.ZendUnmanglePropertyNameEx(key, &tmp, &prop_name, &prop_len)
			} else {
				prop_name = key.GetStr()
				prop_len = key.GetLen()
			}
		} else {
			prop_name = nil
			prop_len = 0
		}
		zdata = types.ZVAL_DEREF(zdata)
		if zdata.IsType(types.IsArray) || zdata.IsType(types.IsObject) {
			if key != nil {
				var ekey string
				if enc_type == PHP_QUERY_RFC3986 {
					ekey = PhpRawUrlEncode(b.CastStr(prop_name, prop_len))
				} else {
					ekey = PhpUrlEncode(b.CastStr(prop_name, prop_len))
				}
				newprefix = key_prefix + ekey + key_suffix + "%5B"
			} else {
				/* Is an integer key */
				ekey := strconv.Itoa(idx)
				newprefix = key_prefix[:key_prefix_len] + num_prefix[:num_prefix_len] + ekey + key_prefix[:key_suffix_len] + "%5B"
			}
			newprefix_len = len(newprefix)
			ht.ProtectRecursive()
			PhpUrlEncodeHashEx(zend.HASH_OF(zdata), formstr, nil, 0, newprefix, newprefix_len, "%5D", 3, lang.Cond(zdata.IsType(types.IsObject), zdata, nil), arg_sep, enc_type)
			ht.UnprotectRecursive()
			zend.Efree(newprefix)
		} else if zdata.IsType(types.IsNull) || zdata.IsType(types.IsResource) {

			/* Skip these types */

			continue

			/* Skip these types */

		} else {
			if formstr.GetS() != nil {
				formstr.WriteString(b.CastStr(arg_sep, arg_sep_len))
			}

			/* Simple key=value */

			if key_prefix != nil {
				formstr.WriteString(b.CastStr(key_prefix, key_prefix_len))
			}
			if key != nil {
				var ekey string
				if enc_type == PHP_QUERY_RFC3986 {
					ekey = PhpRawUrlEncode(b.CastStr(prop_name, prop_len))
				} else {
					ekey = PhpUrlEncode(b.CastStr(prop_name, prop_len))
				}
				formstr.WriteString(ekey)
			} else {

				/* Numeric key */

				if num_prefix != nil {
					formstr.WriteString(b.CastStr(num_prefix, num_prefix_len))
				}
				formstr.WriteLong(idx)
			}
			if key_suffix {
				formstr.WriteString(b.CastStr(key_suffix, key_suffix_len))
			}
			formstr.WriteString("=")
			switch zdata.Type() {
			case types.IsString:
				var ekey string
				if enc_type == PHP_QUERY_RFC3986 {
					ekey = PhpRawUrlEncode(zdata.String())
				} else {
					ekey = PhpUrlEncode(zdata.String())
				}
				formstr.WriteString(ekey)
			case types.IsLong:
				formstr.WriteLong(zdata.Long())
			case types.IsFalse:
				formstr.WriteString("0")
			case types.IsTrue:
				formstr.WriteString("1")
			default:
				var ekey string
				var str = operators.ZvalGetStrVal(zdata)
				if enc_type == PHP_QUERY_RFC3986 {
					ekey = PhpRawUrlEncode(str)
				} else {
					ekey = PhpUrlEncode(str)
				}
				formstr.WriteString(ekey)
			}
		}
	}
	return types.SUCCESS
}
func ZifHttpBuildQuery(executeData zpp.Ex, return_value zpp.Ret, formdata *types.Zval, _ zpp.Opt, prefix *types.Zval, argSeparator *types.Zval, encType *types.Zval) {
	var formdata *types.Zval
	var prefix *byte = nil
	var arg_sep *byte = nil
	var arg_sep_len int = 0
	var prefix_len int = 0
	var formstr zend.SmartStr
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
	if PhpUrlEncodeHashEx(zend.HASH_OF(formdata), &formstr, prefix, prefix_len, nil, 0, nil, 0, lang.Cond(formdata.IsType(types.IsObject), formdata, nil), arg_sep, int(enc_type)) == types.FAILURE {
		if formstr.GetS() != nil {
			formstr.Free()
		}
		return_value.SetFalse()
		return
	}
	if formstr.GetS() == nil {
		return_value.SetString("")
		return
	}
	//formstr.ZeroTail()
	return_value.SetString(formstr.GetStr())
	return
}
