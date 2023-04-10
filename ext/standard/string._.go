package standard

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
	"github.com/heyuuu/gophp/zend/zpp"
)

func ZifSubstrReplace(executeData zpp.Ex, return_value zpp.Ret, str *types.Zval, replace *types.Zval, start *types.Zval, _ zpp.Opt, length *types.Zval) {
	var str *types.Zval
	var from *types.Zval
	var len_ *types.Zval = nil
	var repl *types.Zval
	var l zend.ZendLong = 0
	var f zend.ZendLong
	var argc int = executeData.NumArgs()
	var result *types.String
	var from_idx types.ArrayPosition
	var repl_idx types.ArrayPosition
	var len_idx types.ArrayPosition
	var tmp_str *types.Zval = nil
	var tmp_from *types.Zval = nil
	var tmp_repl *types.Zval = nil
	var tmp_len *types.Zval = nil
	for {
		for {
			fp := zpp.FastParseStart(executeData, 3, 4, 0)
			str = fp.ParseZval()
			repl = fp.ParseZval()
			from = fp.ParseZval()
			fp.StartOptional()
			len_ = fp.ParseZval()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if str.GetType() != types.IS_ARRAY {
		zend.ConvertToStringEx(str)
	}
	if repl.GetType() != types.IS_ARRAY {
		zend.ConvertToStringEx(repl)
	}
	if from.GetType() != types.IS_ARRAY {
		if from.GetType() != types.IS_LONG {
			zend.ConvertToLong(from)
		}
	}
	if zend.EG__().GetException() != nil {
		return
	}
	if argc > 3 {
		if len_.GetType() != types.IS_ARRAY {
			if len_.GetType() != types.IS_LONG {
				zend.ConvertToLong(len_)
			}
			l = len_.GetLval()
		}
	} else {
		if str.GetType() != types.IS_ARRAY {
			l = str.GetStr().GetLen()
		}
	}
	if str.IsType(types.IS_STRING) {
		if argc == 3 && from.IsType(types.IS_ARRAY) || argc == 4 && from.GetType() != len_.GetType() {
			core.PhpErrorDocref(nil, faults.E_WARNING, "'start' and 'length' should be of same type - numerical or array ")
			return_value.SetStringCopy(str.GetStr())
			return
		}
		if argc == 4 && from.IsType(types.IS_ARRAY) {
			if types.Z_ARRVAL_P(from).Len() != types.Z_ARRVAL_P(len_).Len() {
				core.PhpErrorDocref(nil, faults.E_WARNING, "'start' and 'length' should have the same number of elements")
				return_value.SetStringCopy(str.GetStr())
				return
			}
		}
	}
	if str.GetType() != types.IS_ARRAY {
		if from.GetType() != types.IS_ARRAY {
			var repl_str *types.String
			var tmp_repl_str *types.String = nil
			f = from.GetLval()

			/* if "from" position is negative, count start position from the end
			 * of the string
			 */

			if f < 0 {
				f = zend.ZendLong(str.GetStr().GetLen() + f)
				if f < 0 {
					f = 0
				}
			} else if int(f > str.GetStr().GetLen()) != 0 {
				f = str.GetStr().GetLen()
			}

			if l < 0 {
				l = zend.ZendLong(str.GetStr().GetLen()-f) + l
				if l < 0 {
					l = 0
				}
			}
			if int(l > str.GetStr().GetLen() || l < 0 && size_t(-l) > str.GetStr().GetLen()) != 0 {
				l = str.GetStr().GetLen()
			}
			if f+l > zend.ZendLong(str.GetStr().GetLen()) {
				l = str.GetStr().GetLen() - f
			}
			if repl.IsType(types.IS_ARRAY) {
				repl_idx = 0
				for repl_idx < types.Z_ARRVAL_P(repl).GetNNumUsed() {
					tmp_repl = types.Z_ARRVAL_P(repl).GetArData()[repl_idx].GetVal()
					if tmp_repl.IsNotUndef() {
						break
					}
					repl_idx++
				}
				if repl_idx < types.Z_ARRVAL_P(repl).GetNNumUsed() {
					repl_str = zend.ZvalGetTmpString(tmp_repl, &tmp_repl_str)
				} else {
					repl_str = types.NewString("")
				}
			} else {
				repl_str = repl.GetStr()
			}
			result = types.ZendStringSafeAlloc(1, str.GetStr().GetLen()-l+repl_str.GetLen(), 0, 0)
			memcpy(result.GetVal(), str.GetStr().GetVal(), f)
			if repl_str.GetLen() != 0 {
				memcpy(result.GetVal()+f, repl_str.GetVal(), repl_str.GetLen())
			}
			memcpy(result.GetVal()+f+repl_str.GetLen(), str.GetStr().GetVal()+f+l, str.GetStr().GetLen()-f-l)
			result.GetVal()[result.GetLen()] = '0'
			zend.ZendTmpStringRelease(tmp_repl_str)
			return_value.SetString(result)
			return
		} else {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Functionality of 'start' and 'length' as arrays is not implemented")
			return_value.SetStringCopy(str.GetStr())
			return
		}
	} else {
		var str_index *types.String = nil
		var result_len int
		var num_index zend.ZendUlong
		zend.ArrayInit(return_value)
		repl_idx = 0
		len_idx = repl_idx
		from_idx = len_idx
		var __ht *types.Array = str.GetArr()
		for _, _p := range __ht.ForeachData() {
			var _z *types.Zval = _p.GetVal()
			if _z.IsIndirect() {
				_z = _z.GetZv()
				if _z.IsUndef() {
					continue
				}
			}
			num_index = _p.GetH()
			str_index = _p.GetKey()
			tmp_str = _z
			var tmp_orig_str *types.String
			var orig_str *types.String = zend.ZvalGetTmpString(tmp_str, &tmp_orig_str)
			if from.IsType(types.IS_ARRAY) {
				for from_idx < types.Z_ARRVAL_P(from).GetNNumUsed() {
					tmp_from = types.Z_ARRVAL_P(from).GetArData()[from_idx].GetVal()
					if tmp_from.IsNotUndef() {
						break
					}
					from_idx++
				}
				if from_idx < types.Z_ARRVAL_P(from).GetNNumUsed() {
					f = zend.ZvalGetLong(tmp_from)
					if f < 0 {
						f = zend.ZendLong(orig_str.GetLen() + f)
						if f < 0 {
							f = 0
						}
					} else if f > zend.ZendLong(orig_str.GetLen()) {
						f = orig_str.GetLen()
					}
					from_idx++
				} else {
					f = 0
				}
			} else {
				f = from.GetLval()
				if f < 0 {
					f = zend.ZendLong(orig_str.GetLen() + f)
					if f < 0 {
						f = 0
					}
				} else if f > zend.ZendLong(orig_str.GetLen()) {
					f = orig_str.GetLen()
				}
			}
			if argc > 3 && len_.IsType(types.IS_ARRAY) {
				for len_idx < types.Z_ARRVAL_P(len_).GetNNumUsed() {
					tmp_len = types.Z_ARRVAL_P(len_).GetArData()[len_idx].GetVal()
					if tmp_len.IsNotUndef() {
						break
					}
					len_idx++
				}
				if len_idx < types.Z_ARRVAL_P(len_).GetNNumUsed() {
					l = zend.ZvalGetLong(tmp_len)
					len_idx++
				} else {
					l = orig_str.GetLen()
				}
			} else if argc > 3 {
				l = len_.GetLval()
			} else {
				l = orig_str.GetLen()
			}
			if l < 0 {
				l = orig_str.GetLen() - f + l
				if l < 0 {
					l = 0
				}
			}
			b.Assert(0 <= f && f <= zend.ZEND_LONG_MAX)
			b.Assert(0 <= l && l <= zend.ZEND_LONG_MAX)
			if int(f+l) > orig_str.GetLen() {
				l = orig_str.GetLen() - f
			}
			result_len = orig_str.GetLen() - l
			if repl.IsType(types.IS_ARRAY) {
				for repl_idx < types.Z_ARRVAL_P(repl).GetNNumUsed() {
					tmp_repl = types.Z_ARRVAL_P(repl).GetArData()[repl_idx].GetVal()
					if tmp_repl.IsNotUndef() {
						break
					}
					repl_idx++
				}
				if repl_idx < types.Z_ARRVAL_P(repl).GetNNumUsed() {
					var tmp_repl_str *types.String
					var repl_str *types.String = zend.ZvalGetTmpString(tmp_repl, &tmp_repl_str)
					result_len += repl_str.GetLen()
					repl_idx++
					result = types.ZendStringSafeAlloc(1, result_len, 0, 0)
					memcpy(result.GetVal(), orig_str.GetVal(), f)
					memcpy(result.GetVal()+f, repl_str.GetVal(), repl_str.GetLen())
					memcpy(result.GetVal()+f+repl_str.GetLen(), orig_str.GetVal()+f+l, orig_str.GetLen()-f-l)
					zend.ZendTmpStringRelease(tmp_repl_str)
				} else {
					result = types.ZendStringSafeAlloc(1, result_len, 0, 0)
					memcpy(result.GetVal(), orig_str.GetVal(), f)
					memcpy(result.GetVal()+f, orig_str.GetVal()+f+l, orig_str.GetLen()-f-l)
				}
			} else {
				result_len += repl.GetStr().GetLen()
				result = types.ZendStringSafeAlloc(1, result_len, 0, 0)
				memcpy(result.GetVal(), orig_str.GetVal(), f)
				memcpy(result.GetVal()+f, repl.GetStr().GetVal(), repl.GetStr().GetLen())
				memcpy(result.GetVal()+f+repl.GetStr().GetLen(), orig_str.GetVal()+f+l, orig_str.GetLen()-f-l)
			}
			result.GetVal()[result.GetLen()] = '0'
			if str_index != nil {
				var tmp types.Zval
				tmp.SetString(result)
				return_value.GetArr().SymtableUpdate(str_index.GetStr(), &tmp)
			} else {
				zend.AddIndexStr(return_value, num_index, result)
			}
			zend.ZendTmpStringRelease(tmp_orig_str)
		}
	}
}
