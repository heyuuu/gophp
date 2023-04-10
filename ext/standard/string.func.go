package standard

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/builtin/ascii"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/sapi/cli"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
	"github.com/heyuuu/gophp/zend/zpp"
	"strings"
)

func RegisterStringConstants(type_ int, module_number int) {
	zend.RegisterLongConstant("STR_PAD_LEFT", STR_PAD_LEFT, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STR_PAD_RIGHT", STR_PAD_RIGHT, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STR_PAD_BOTH", STR_PAD_BOTH, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("PATHINFO_DIRNAME", PHP_PATHINFO_DIRNAME, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("PATHINFO_BASENAME", PHP_PATHINFO_BASENAME, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("PATHINFO_EXTENSION", PHP_PATHINFO_EXTENSION, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("PATHINFO_FILENAME", PHP_PATHINFO_FILENAME, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)

	/* If last members of struct lconv equal CHAR_MAX, no grouping is done */

	zend.RegisterLongConstant("CHAR_MAX", CHAR_MAX, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("LC_CTYPE", LC_CTYPE, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("LC_NUMERIC", LC_NUMERIC, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("LC_TIME", LC_TIME, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("LC_COLLATE", LC_COLLATE, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("LC_MONETARY", LC_MONETARY, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("LC_ALL", LC_ALL, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
}

func PhpStrtolower(s *byte, len_ int) *byte {
	var c *uint8
	var e *uint8
	c = (*uint8)(s)
	e = c + len_
	for c < e {
		*c = tolower(*c)
		c++
	}
	return s
}

func PhpStrtr(str *byte, len_ int, str_from *byte, str_to *byte, trlen int) *byte {
	var i int
	if trlen < 1 {
		return str
	} else if trlen == 1 {
		var ch_from byte = *str_from
		var ch_to byte = *str_to
		for i = 0; i < len_; i++ {
			if str[i] == ch_from {
				str[i] = ch_to
			}
		}
	} else {
		var xlat []uint8
		var j uint8 = 0
		for {
			xlat[j] = j
			if b.PreInc(&j) == 0 {
				break
			}
		}
		for i = 0; i < trlen; i++ {
			xlat[int(uint8(str_from[i]))] = str_to[i]
		}
		for i = 0; i < len_; i++ {
			str[i] = xlat[int(uint8(str[i]))]
		}
	}
	return str
}
func PhpStrtrEx(str *types.String, str_from *byte, str_to *byte, trlen int) *types.String {
	var new_str *types.String = nil
	var i int
	if trlen < 1 {
		return str.Copy()
	} else if trlen == 1 {
		var ch_from byte = *str_from
		var ch_to byte = *str_to
		for i = 0; i < str.GetLen(); i++ {
			if str.GetVal()[i] == ch_from {
				new_str = types.ZendStringAlloc(str.GetLen(), 0)
				memcpy(new_str.GetVal(), str.GetVal(), i)
				new_str.GetVal()[i] = ch_to
				break
			}
		}
		for ; i < str.GetLen(); i++ {
			if str.GetVal()[i] != ch_from {
				new_str.GetVal()[i] = str.GetVal()[i]
			} else {
				new_str.GetVal()[i] = ch_to
			}
		}
	} else {
		var xlat []uint8
		var j uint8 = 0
		for {
			xlat[j] = j
			if b.PreInc(&j) == 0 {
				break
			}
		}
		for i = 0; i < trlen; i++ {
			xlat[int(uint8(str_from[i]))] = str_to[i]
		}
		for i = 0; i < str.GetLen(); i++ {
			if str.GetVal()[i] != xlat[int(uint8(str.GetVal()[i]))] {
				new_str = types.ZendStringAlloc(str.GetLen(), 0)
				memcpy(new_str.GetVal(), str.GetVal(), i)
				new_str.GetVal()[i] = xlat[int(uint8(str.GetVal()[i]))]
				break
			}
		}
		for ; i < str.GetLen(); i++ {
			new_str.GetVal()[i] = xlat[int(uint8(str.GetVal()[i]))]
		}
	}
	if new_str == nil {
		return str.Copy()
	}
	new_str.GetVal()[new_str.GetLen()] = 0
	return new_str
}
func PhpStrtrArray(return_value *types.Zval, input *types.String, pats *types.Array) {
	var str *byte = input.GetVal()
	var slen int = input.GetLen()
	var num_key zend.ZendUlong
	var str_key *types.String
	var len_ int
	var pos int
	var old_pos int
	var num_keys int = 0
	var minlen int = 128 * 1024
	var maxlen int = 0
	var str_hash types.Array
	var entry *types.Zval
	var key *byte
	var result zend.SmartStr = zend.MakeSmartStr(0)
	var bitset []zend.ZendUlong
	var num_bitset *zend.ZendUlong

	/* we will collect all possible key lengths */

	num_bitset = zend.Ecalloc((slen+b.SizeOf("zend_ulong"))/b.SizeOf("zend_ulong"), b.SizeOf("zend_ulong"))
	memset(bitset, 0, b.SizeOf("bitset"))

	/* check if original array has numeric keys */

	var __ht *types.Array = pats
	for _, _p := range __ht.ForeachData() {
		var _z *types.Zval = _p.GetVal()

		str_key = _p.GetKey()
		if str_key == nil {
			num_keys = 1
		} else {
			len_ = str_key.GetLen()
			if len_ < 1 {
				zend.Efree(num_bitset)
				return_value.SetFalse()
				return
			} else if len_ > slen {

				/* skip long patterns */

				continue

				/* skip long patterns */

			}
			if len_ > maxlen {
				maxlen = len_
			}
			if len_ < minlen {
				minlen = len_
			}

			/* remember possible key length */

			num_bitset[len_/b.SizeOf("zend_ulong")] |= uint64(1) << len_ % b.SizeOf("zend_ulong")
			bitset[uint8(str_key.GetVal()[0])/b.SizeOf("zend_ulong")] |= uint64(1) << uint8(str_key.GetVal()[0]) % b.SizeOf("zend_ulong")
		}
	}
	if num_keys != 0 {
		var key_used *types.String

		/* we have to rebuild HashTable with numeric keys */

		&str_hash = types.MakeArrayEx(pats.Len(), nil, 0)
		var __ht *types.Array = pats
		for _, _p := range __ht.ForeachData() {
			var _z *types.Zval = _p.GetVal()
			if _z.IsIndirect() {
				_z = _z.GetZv()
				if _z.IsUndef() {
					continue
				}
			}
			num_key = _p.GetH()
			str_key = _p.GetKey()
			entry = _z
			if str_key == nil {
				key_used = zend.ZendLongToStr(num_key)
				len_ = key_used.GetLen()
				if len_ > slen {

					/* skip long patterns */

					// types.ZendStringRelease(key_used)
					continue
				}
				if len_ > maxlen {
					maxlen = len_
				}
				if len_ < minlen {
					minlen = len_
				}

				/* remember possible key length */

				num_bitset[len_/b.SizeOf("zend_ulong")] |= uint64(1) << len_ % b.SizeOf("zend_ulong")
				bitset[uint8(key_used.GetVal()[0])/b.SizeOf("zend_ulong")] |= uint64(1) << uint8(key_used.GetVal()[0]) % b.SizeOf("zend_ulong")
			} else {
				key_used = str_key
				len_ = key_used.GetLen()
				if len_ > slen {

					/* skip long patterns */

					continue

					/* skip long patterns */

				}
			}
			str_hash.KeyAdd(key_used.GetStr(), entry)
			if str_key == nil {
				// types.ZendStringReleaseEx(key_used, 0)
			}
		}
		pats = &str_hash
	}
	if minlen > maxlen {

		/* return the original string */

		if pats == &str_hash {
			str_hash.Destroy()
		}
		zend.Efree(num_bitset)
		return_value.SetStringCopy(input)
		return
	}
	pos = 0
	old_pos = pos
	for pos <= slen-minlen {
		key = str + pos
		if (bitset[uint8(key[0])/b.SizeOf("zend_ulong")] & uint64(1) << uint8(key[0]) % b.SizeOf("zend_ulong")) != 0 {
			len_ = maxlen
			if len_ > slen-pos {
				len_ = slen - pos
			}
			for len_ >= minlen {
				if (num_bitset[len_/b.SizeOf("zend_ulong")] & uint64(1) << len_ % b.SizeOf("zend_ulong")) != 0 {
					entry = pats.KeyFind(b.CastStr(key, len_))
					if entry != nil {
						var tmp *types.String
						var s *types.String = zend.ZvalGetTmpString(entry, &tmp)
						result.AppendString(b.CastStr(str+old_pos, pos-old_pos))
						result.AppendString(s.GetStr())
						old_pos = pos + len_
						pos = old_pos - 1
						zend.ZendTmpStringRelease(tmp)
						break
					}
				}
				len_--
			}
		}
		pos++
	}
	if result.GetS() != nil {
		result.AppendString(b.CastStr(str+old_pos, slen-old_pos))
		result.ZeroTail()
		return_value.SetString(result.GetS())
	} else {
		result.Free()
		return_value.SetStringCopy(input)
	}
	if pats == &str_hash {
		str_hash.Destroy()
	}
	zend.Efree(num_bitset)
}
func PhpCharToStrEx(
	str *types.String,
	from byte,
	to *byte,
	to_len int,
	case_sensitivity int,
	replace_count *zend.ZendLong,
) *types.String {
	var result *types.String
	var char_count int = 0
	var lc_from int = 0
	var source *byte
	var source_end *byte = str.GetVal() + str.GetLen()
	var target *byte
	if case_sensitivity != 0 {
		var p *byte = str.GetVal()
		var e *byte = p + str.GetLen()
		for b.Assign(&p, memchr(p, from, e-p)) {
			char_count++
			p++
		}
	} else {
		lc_from = tolower(from)
		for source = str.GetVal(); source < source_end; source++ {
			if tolower(*source) == lc_from {
				char_count++
			}
		}
	}
	if char_count == 0 {
		return str.Copy()
	}
	if to_len > 0 {
		result = types.ZendStringSafeAlloc(char_count, to_len-1, str.GetLen(), 0)
	} else {
		result = types.ZendStringAlloc(str.GetLen()-char_count, 0)
	}
	target = result.GetVal()
	if case_sensitivity != 0 {
		var p *byte = str.GetVal()
		var e *byte = p + str.GetLen()
		var s *byte = str.GetVal()
		for b.Assign(&p, memchr(p, from, e-p)) {
			memcpy(target, s, p-s)
			target += p - s
			memcpy(target, to, to_len)
			target += to_len
			p++
			s = p
			if replace_count != nil {
				*replace_count += 1
			}
		}
		if s < e {
			memcpy(target, s, e-s)
			target += e - s
		}
	} else {
		for source = str.GetVal(); source < source_end; source++ {
			if tolower(*source) == lc_from {
				if replace_count != nil {
					*replace_count += 1
				}
				memcpy(target, to, to_len)
				target += to_len
			} else {
				*target = *source
				target++
			}
		}
	}
	*target = 0
	return result
}
func PhpStrToStrEx(
	haystack *types.String,
	needle *byte,
	needle_len int,
	str *byte,
	str_len int,
	replace_count *zend.ZendLong,
) *types.String {
	var new_str *types.String
	if needle_len < haystack.GetLen() {
		var end *byte
		var p *byte
		var r *byte
		var e *byte
		if needle_len == str_len {
			new_str = nil
			end = haystack.GetVal() + haystack.GetLen()
			for p = haystack.GetVal(); b.Assign(&r, (*byte)(core.PhpMemnstr(p, needle, needle_len, end))); p = r + needle_len {
				if new_str == nil {
					new_str = types.NewString(haystack.GetStr())
				}
				memcpy(new_str.GetVal()+(r-haystack.GetVal()), str, str_len)
				*replace_count++
			}
			if new_str == nil {
				goto nothing_todo
			}
			return new_str
		} else {
			var count int = 0
			var o *byte = haystack.GetVal()
			var n *byte = needle
			var endp *byte = o + haystack.GetLen()
			for b.Assign(&o, (*byte)(core.PhpMemnstr(o, n, needle_len, endp))) {
				o += needle_len
				count++
			}
			if count == 0 {

				/* Needle doesn't occur, shortcircuit the actual replacement. */

				goto nothing_todo

				/* Needle doesn't occur, shortcircuit the actual replacement. */

			}
			if str_len > needle_len {
				new_str = types.ZendStringSafeAlloc(count, str_len-needle_len, haystack.GetLen(), 0)
			} else {
				new_str = types.ZendStringAlloc(count*(str_len-needle_len)+haystack.GetLen(), 0)
			}
			e = new_str.GetVal()
			end = haystack.GetVal() + haystack.GetLen()
			for p = haystack.GetVal(); b.Assign(&r, (*byte)(core.PhpMemnstr(p, needle, needle_len, end))); p = r + needle_len {
				memcpy(e, p, r-p)
				e += r - p
				memcpy(e, str, str_len)
				e += str_len
				*replace_count++
			}
			if p < end {
				memcpy(e, p, end-p)
				e += end - p
			}
			*e = '0'
			return new_str
		}
	} else if needle_len > haystack.GetLen() || memcmp(haystack.GetVal(), needle, haystack.GetLen()) {
	nothing_todo:
		return haystack.Copy()
	} else {
		if str_len == 0 {
			new_str = types.NewString("")
		} else if str_len == 1 {
			new_str = types.NewString(string(*str))
		} else {
			new_str = types.NewString(b.CastStr(str, str_len))
		}
		*replace_count++
		return new_str
	}
}
func PhpStrToStrIEx(
	haystack *types.String,
	lc_haystack *byte,
	needle *types.String,
	str *byte,
	str_len int,
	replace_count *zend.ZendLong,
) *types.String {
	var new_str *types.String = nil
	var lc_needle *types.String
	if needle.GetLen() < haystack.GetLen() {
		var end *byte
		var p *byte
		var r *byte
		var e *byte
		if needle.GetLen() == str_len {
			lc_needle = PhpStringTolower(needle)
			end = lc_haystack + haystack.GetLen()
			for p = lc_haystack; b.Assign(&r, (*byte)(core.PhpMemnstr(p, lc_needle.GetVal(), lc_needle.GetLen(), end))); p = r + lc_needle.GetLen() {
				if new_str == nil {
					new_str = types.NewString(haystack.GetStr())
				}
				memcpy(new_str.GetVal()+(r-lc_haystack), str, str_len)
				*replace_count++
			}
			// types.ZendStringReleaseEx(lc_needle, 0)
			if new_str == nil {
				goto nothing_todo
			}
			return new_str
		} else {
			var count int = 0
			var o *byte = lc_haystack
			var n *byte
			var endp *byte = o + haystack.GetLen()
			lc_needle = PhpStringTolower(needle)
			n = lc_needle.GetVal()
			for b.Assign(&o, (*byte)(core.PhpMemnstr(o, n, lc_needle.GetLen(), endp))) {
				o += lc_needle.GetLen()
				count++
			}
			if count == 0 {

				/* Needle doesn't occur, shortcircuit the actual replacement. */

				// types.ZendStringReleaseEx(lc_needle, 0)
				goto nothing_todo
			}
			if str_len > lc_needle.GetLen() {
				new_str = types.ZendStringSafeAlloc(count, str_len-lc_needle.GetLen(), haystack.GetLen(), 0)
			} else {
				new_str = types.ZendStringAlloc(count*(str_len-lc_needle.GetLen())+haystack.GetLen(), 0)
			}
			e = new_str.GetVal()
			end = lc_haystack + haystack.GetLen()
			for p = lc_haystack; b.Assign(&r, (*byte)(core.PhpMemnstr(p, lc_needle.GetVal(), lc_needle.GetLen(), end))); p = r + lc_needle.GetLen() {
				memcpy(e, haystack.GetVal()+(p-lc_haystack), r-p)
				e += r - p
				memcpy(e, str, str_len)
				e += str_len
				*replace_count++
			}
			if p < end {
				memcpy(e, haystack.GetVal()+(p-lc_haystack), end-p)
				e += end - p
			}
			*e = '0'
			// types.ZendStringReleaseEx(lc_needle, 0)
			return new_str
		}
	} else if needle.GetLen() > haystack.GetLen() {
	nothing_todo:
		return haystack.Copy()
	} else {
		lc_needle = PhpStringTolower(needle)
		if memcmp(lc_haystack, lc_needle.GetVal(), lc_needle.GetLen()) {
			// types.ZendStringReleaseEx(lc_needle, 0)
			goto nothing_todo
		}
		// types.ZendStringReleaseEx(lc_needle, 0)
		new_str = types.NewString(b.CastStr(str, str_len))
		*replace_count++
		return new_str
	}
}
func PhpStrToStr(
	haystack *byte,
	length int,
	needle string,
	needle_len int,
	str string,
	str_len int,
) *types.String {
	var new_str *types.String
	if needle_len < length {
		var end *byte
		var s *byte
		var p *byte
		var e *byte
		var r *byte
		if needle_len == str_len {
			new_str = types.NewString(b.CastStr(haystack, length))
			end = new_str.GetVal() + length
			for p = new_str.GetVal(); b.Assign(&r, (*byte)(core.PhpMemnstr(p, needle, needle_len, end))); p = r + needle_len {
				memcpy(r, str, str_len)
			}
			return new_str
		} else {
			if str_len < needle_len {
				new_str = types.ZendStringAlloc(length, 0)
			} else {
				var count int = 0
				var o *byte = haystack
				var n *byte = needle
				var endp *byte = o + length
				for b.Assign(&o, (*byte)(core.PhpMemnstr(o, n, needle_len, endp))) {
					o += needle_len
					count++
				}
				if count == 0 {

					/* Needle doesn't occur, shortcircuit the actual replacement. */

					new_str = types.NewString(b.CastStr(haystack, length))
					return new_str
				} else {
					if str_len > needle_len {
						new_str = types.ZendStringSafeAlloc(count, str_len-needle_len, length, 0)
					} else {
						new_str = types.ZendStringAlloc(count*(str_len-needle_len)+length, 0)
					}
				}
			}
			e = new_str.GetVal()
			s = e
			end = haystack + length
			for p = haystack; b.Assign(&r, (*byte)(core.PhpMemnstr(p, needle, needle_len, end))); p = r + needle_len {
				memcpy(e, p, r-p)
				e += r - p
				memcpy(e, str, str_len)
				e += str_len
			}
			if p < end {
				memcpy(e, p, end-p)
				e += end - p
			}
			*e = '0'
			new_str = types.ZendStringTruncate(new_str, e-s)
			return new_str
		}
	} else if needle_len > length || memcmp(haystack, needle, length) {
		new_str = types.NewString(b.CastStr(haystack, length))
		return new_str
	} else {
		new_str = types.NewString(b.CastStr(str, str_len))
		return new_str
	}
}
func ZifStrtr(executeData zpp.Ex, return_value zpp.Ret, str *types.Zval, from *types.Zval, _ zpp.Opt, to *types.Zval) {
	var from *types.Zval
	var str *types.String
	var to *byte = nil
	var to_len int = 0
	var ac int = executeData.NumArgs()
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 3, 0)
			str = fp.ParseStr()
			from = fp.ParseZval()
			fp.StartOptional()
			to, to_len = fp.ParseString()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if ac == 2 && from.GetType() != types.IS_ARRAY {
		core.PhpErrorDocref(nil, faults.E_WARNING, "The second argument is not an array")
		return_value.SetFalse()
		return
	}

	/* shortcut for empty string */

	if str.GetLen() == 0 {
		return_value.SetStringVal("")
		return
	}
	if ac == 2 {
		var pats *types.Array = from.GetArr()
		if pats.Len() < 1 {
			return_value.SetStringCopy(str)
			return
		} else if pats.Len() == 1 {
			var num_key zend.ZendLong
			var str_key *types.String
			var tmp_str *types.String
			var replace *types.String
			var tmp_replace *types.String
			var entry *types.Zval
			var __ht *types.Array = pats
			for _, _p := range __ht.ForeachData() {
				var _z *types.Zval = _p.GetVal()
				if _z.IsIndirect() {
					_z = _z.GetZv()
					if _z.IsUndef() {
						continue
					}
				}
				num_key = _p.GetH()
				str_key = _p.GetKey()
				entry = _z
				tmp_str = nil
				if str_key == nil {
					tmp_str = zend.ZendLongToStr(num_key)
					str_key = tmp_str
				}
				replace = zend.ZvalGetTmpString(entry, &tmp_replace)
				if str_key.GetLen() < 1 {
					return_value.SetStringCopy(str)
				} else if str_key.GetLen() == 1 {
					return_value.SetString(PhpCharToStrEx(str, str_key.GetVal()[0], replace.GetVal(), replace.GetLen(), 1, nil))
				} else {
					var dummy zend.ZendLong
					return_value.SetString(PhpStrToStrEx(str, str_key.GetVal(), str_key.GetLen(), replace.GetVal(), replace.GetLen(), &dummy))
				}
				zend.ZendTmpStringRelease(tmp_str)
				zend.ZendTmpStringRelease(tmp_replace)
				return
			}
		} else {
			PhpStrtrArray(return_value, str, pats)
		}
	} else {
		if zend.TryConvertToString(from) == 0 {
			return
		}
		return_value.SetString(PhpStrtrEx(str, from.GetStr().GetVal(), to, cli.MIN(from.GetStr().GetLen(), to_len)))
		return
	}
}
func ZifStrrev(executeData zpp.Ex, return_value zpp.Ret, str *types.Zval) {
	var str *types.String
	var s *byte
	var e *byte
	var p *byte
	var n *types.String
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			str = fp.ParseStr()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	n = types.ZendStringAlloc(str.GetLen(), 0)
	p = n.GetVal()
	s = str.GetVal()
	e = s + str.GetLen()
	e--
	for e >= s {
		*e--
		b.PostInc(&(*p)) = (*e) + 1
	}
	*p = '0'
	return_value.SetString(n)
}
func PhpSimilarStr(
	txt1 *byte,
	len1 int,
	txt2 *byte,
	len2 int,
	pos1 *int,
	pos2 *int,
	max *int,
	count *int,
) {
	var p *byte
	var q *byte
	var end1 *byte = (*byte)(txt1 + len1)
	var end2 *byte = (*byte)(txt2 + len2)
	var l int
	*max = 0
	*count = 0
	for p = (*byte)(txt1); p < end1; p++ {
		for q = (*byte)(txt2); q < end2; q++ {
			for l = 0; p+l < end1 && q+l < end2 && p[l] == q[l]; l++ {

			}
			if l > (*max) {
				*max = l
				*count += 1
				*pos1 = p - txt1
				*pos2 = q - txt2
			}
		}
	}
}
func PhpSimilarChar(txt1 *byte, len1 int, txt2 *byte, len2 int) int {
	var sum int
	var pos1 int = 0
	var pos2 int = 0
	var max int
	var count int
	PhpSimilarStr(txt1, len1, txt2, len2, &pos1, &pos2, &max, &count)
	if b.Assign(&sum, max) {
		if pos1 != 0 && pos2 != 0 && count > 1 {
			sum += PhpSimilarChar(txt1, pos1, txt2, pos2)
		}
		if pos1+max < len1 && pos2+max < len2 {
			sum += PhpSimilarChar(txt1+pos1+max, len1-pos1-max, txt2+pos2+max, len2-pos2-max)
		}
	}
	return sum
}
func ZifSimilarText(executeData zpp.Ex, return_value zpp.Ret, str1 *types.Zval, str2 *types.Zval, _ zpp.Opt, percent zpp.RefZval) {
	var t1 *types.String
	var t2 *types.String
	var percent *types.Zval = nil
	var ac int = executeData.NumArgs()
	var sim int
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 3, 0)
			t1 = fp.ParseStr()
			t2 = fp.ParseStr()
			fp.StartOptional()
			percent = fp.ParseZval()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if t1.GetLen()+t2.GetLen() == 0 {
		if ac > 2 {
			zend.ZEND_TRY_ASSIGN_REF_DOUBLE(percent, 0)
		}
		return_value.SetLong(0)
		return
	}
	sim = PhpSimilarChar(t1.GetVal(), t1.GetLen(), t2.GetVal(), t2.GetLen())
	if ac > 2 {
		zend.ZEND_TRY_ASSIGN_REF_DOUBLE(percent, sim*200.0/(t1.GetLen()+t2.GetLen()))
	}
	return_value.SetLong(sim)
	return
}
func ZifAddcslashes(str string, charlist string) string {
	if str == "" {
		return ""
	}
	if charlist == "" {
		return str
	}
	return PhpAddcslashes(str, charlist)
}
func ZifAddslashes(str string) string {
	if str == "" {
		return ""
	}
	return PhpAddslashes(str)
}
func ZifStripcslashes(executeData zpp.Ex, return_value zpp.Ret, str *types.Zval) {
	var str *types.String
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			str = fp.ParseStr()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	return_value.SetStringVal(str.GetStr())
	PhpStripcslashes(return_value.GetStr())
}
func ZifStripslashes(executeData zpp.Ex, return_value zpp.Ret, str *types.Zval) {
	var str *types.String
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			str = fp.ParseStr()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	return_value.SetStringVal(str.GetStr())
	PhpStripslashes(return_value.GetStr())
}
func PhpStripcslashes(str *types.String) {
	var source *byte
	var end *byte
	var target *byte
	var nlen int = str.GetLen()
	var i int
	var numtmp []byte
	source = (*byte)(str.GetVal())
	end = source + str.GetLen()
	target = str.GetVal()
	for ; source < end; source++ {
		if (*source) == '\\' && source+1 < end {
			source++
			switch *source {
			case 'n':
				b.PostInc(&(*target)) = '\n'
				nlen--
			case 'r':
				b.PostInc(&(*target)) = '\r'
				nlen--
			case 'a':
				b.PostInc(&(*target)) = 'a'
				nlen--
			case 't':
				b.PostInc(&(*target)) = '\t'
				nlen--
			case 'v':
				b.PostInc(&(*target)) = 'v'
				nlen--
			case 'b':
				b.PostInc(&(*target)) = 'b'
				nlen--
			case 'f':
				b.PostInc(&(*target)) = 'f'
				nlen--
			case '\\':
				b.PostInc(&(*target)) = '\\'
				nlen--
			case 'x':
				if source+1 < end && isxdigit(int(*(source + 1))) {
					numtmp[0] = *(b.PreInc(&source))
					if source+1 < end && isxdigit(int(*(source + 1))) {
						numtmp[1] = *(b.PreInc(&source))
						numtmp[2] = '0'
						nlen -= 3
					} else {
						numtmp[1] = '0'
						nlen -= 2
					}
					b.PostInc(&(*target)) = byte(strtol(numtmp, nil, 16))
					break
				}
				fallthrough
			default:
				i = 0
				for source < end && (*source) >= '0' && (*source) <= '7' && i < 3 {
					*source++
					numtmp[b.PostInc(&i)] = (*source) - 1
				}
				if i != 0 {
					numtmp[i] = '0'
					b.PostInc(&(*target)) = byte(strtol(numtmp, nil, 8))
					nlen -= i
					source--
				} else {
					b.PostInc(&(*target)) = *source
					nlen--
				}
			}
		} else {
			b.PostInc(&(*target)) = *source
		}
	}
	if nlen != 0 {
		*target = '0'
	}
	str.SetLen(nlen)
}
func PhpStripslashesImpl(str *byte, out *byte, len_ int) *byte {
	for len_ > 0 {
		if (*str) == '\\' {
			str++
			len_--
			if len_ > 0 {
				if (*str) == '0' {
					b.PostInc(&(*out)) = '0'
					str++
				} else {
					*str++
					b.PostInc(&(*out)) = (*str) - 1
				}
				len_--
			}
		} else {
			*str++
			b.PostInc(&(*out)) = (*str) - 1
			len_--
		}
	}
	return out
}
func PhpStripslashes(str *types.String) {
	var t *byte = PhpStripslashesImpl(str.GetVal(), str.GetVal(), str.GetLen())
	if t != str.GetVal()+str.GetLen() {
		str.SetLen(t - str.GetVal())
		str.GetVal()[str.GetLen()] = '0'
	}
}
func Isheb(c __auto__) int {
	if uint8(c) >= 224 && uint8(c) <= 250 {
		return 1
	} else {
		return 0
	}
}
func _isblank(c __auto__) int {
	if uint8(c) == ' ' || uint8(c) == '\t' {
		return 1
	} else {
		return 0
	}
}
func _isnewline(c byte) int {
	if uint8(c) == '\n' || uint8(c) == '\r' {
		return 1
	} else {
		return 0
	}
}
func PhpStrReplaceInSubject(search *types.Zval, replace *types.Zval, subject *types.Zval, result *types.Zval, case_sensitivity int) zend.ZendLong {
	var search_entry *types.Zval
	var tmp_result *types.String
	var tmp_subject_str *types.String
	var replace_value *byte = nil
	var replace_len int = 0
	var replace_count zend.ZendLong = 0
	var subject_str *types.String
	var lc_subject_str *types.String = nil
	var replace_idx uint32

	/* Make sure we're dealing with strings. */

	subject_str = zend.ZvalGetTmpString(subject, &tmp_subject_str)
	if subject_str.GetLen() == 0 {
		zend.ZendTmpStringRelease(tmp_subject_str)
		result.SetStringVal("")
		return 0
	}

	/* If search is an array */

	if search.IsType(types.IS_ARRAY) {

		/* Duplicate subject string for repeated replacement */

		//subject_str.AddRefcount()
		if replace.IsType(types.IS_ARRAY) {
			replace_idx = 0
		} else {
			/* Set replacement value to the passed one */
			replace_value = replace.GetStr().GetVal()
			replace_len = replace.GetStr().GetLen()
		}

		var __ht *types.Array = search.GetArr()
		for _, _p := range __ht.ForeachData() {
			var _z *types.Zval = _p.GetVal()
			if _z.IsIndirect() {
				_z = _z.GetZv()
				if _z.IsUndef() {
					continue
				}
			}
			search_entry = _z

			/* Make sure we're dealing with strings. */

			var tmp_search_str *types.String
			var search_str *types.String = zend.ZvalGetTmpString(search_entry, &tmp_search_str)
			var replace_entry_str *types.String
			var tmp_replace_entry_str *types.String = nil

			/* If replace is an array. */

			if replace.IsType(types.IS_ARRAY) {

				/* Get current entry */

				var replace_entry *types.Zval = nil
				for replace_idx < types.Z_ARRVAL_P(replace).GetNNumUsed() {
					replace_entry = types.Z_ARRVAL_P(replace).GetArData()[replace_idx].GetVal()
					if replace_entry.IsNotUndef() {
						break
					}
					replace_idx++
				}
				if replace_idx < types.Z_ARRVAL_P(replace).GetNNumUsed() {

					/* Make sure we're dealing with strings. */

					replace_entry_str = zend.ZvalGetTmpString(replace_entry, &tmp_replace_entry_str)

					/* Set replacement value to the one we got from array */

					replace_value = replace_entry_str.GetVal()
					replace_len = replace_entry_str.GetLen()
					replace_idx++
				} else {

					/* We've run out of replacement strings, so use an empty one. */

					replace_value = ""
					replace_len = 0
				}
			}
			if search_str.GetLen() == 1 {
				var old_replace_count zend.ZendLong = replace_count
				tmp_result = PhpCharToStrEx(subject_str, search_str.GetVal()[0], replace_value, replace_len, case_sensitivity, &replace_count)
				if lc_subject_str != nil && replace_count != old_replace_count {
					// types.ZendStringReleaseEx(lc_subject_str, 0)
					lc_subject_str = nil
				}
			} else if search_str.GetLen() > 1 {
				if case_sensitivity != 0 {
					tmp_result = PhpStrToStrEx(subject_str, search_str.GetVal(), search_str.GetLen(), replace_value, replace_len, &replace_count)
				} else {
					var old_replace_count zend.ZendLong = replace_count
					if lc_subject_str == nil {
						lc_subject_str = PhpStringTolower(subject_str)
					}
					tmp_result = PhpStrToStrIEx(subject_str, lc_subject_str.GetVal(), search_str, replace_value, replace_len, &replace_count)
					if replace_count != old_replace_count {
						// types.ZendStringReleaseEx(lc_subject_str, 0)
						lc_subject_str = nil
					}
				}
			} else {
				zend.ZendTmpStringRelease(tmp_search_str)
				zend.ZendTmpStringRelease(tmp_replace_entry_str)
				continue
			}
			zend.ZendTmpStringRelease(tmp_search_str)
			zend.ZendTmpStringRelease(tmp_replace_entry_str)
			if subject_str == tmp_result {
				//subject_str.DelRefcount()
			} else {
				// types.ZendStringReleaseEx(subject_str, 0)
				subject_str = tmp_result
				if subject_str.GetLen() == 0 {
					// types.ZendStringReleaseEx(subject_str, 0)
					result.SetStringVal("")
					if lc_subject_str != nil {
						// types.ZendStringReleaseEx(lc_subject_str, 0)
					}
					zend.ZendTmpStringRelease(tmp_subject_str)
					return replace_count
				}
			}
		}
		result.SetString(subject_str)
		if lc_subject_str != nil {
			// types.ZendStringReleaseEx(lc_subject_str, 0)
		}
	} else {
		b.Assert(search.IsType(types.IS_STRING))
		if search.GetStr().GetLen() == 1 {
			result.SetString(PhpCharToStrEx(subject_str, search.GetStr().GetVal()[0], replace.GetStr().GetVal(), replace.GetStr().GetLen(), case_sensitivity, &replace_count))
		} else if search.GetStr().GetLen() > 1 {
			if case_sensitivity != 0 {
				result.SetString(PhpStrToStrEx(subject_str, search.GetStr().GetVal(), search.GetStr().GetLen(), replace.GetStr().GetVal(), replace.GetStr().GetLen(), &replace_count))
			} else {
				lc_subject_str = PhpStringTolower(subject_str)
				result.SetString(PhpStrToStrIEx(subject_str, lc_subject_str.GetVal(), search.GetStr(), replace.GetStr().GetVal(), replace.GetStr().GetLen(), &replace_count))
				// types.ZendStringReleaseEx(lc_subject_str, 0)
			}
		} else {
			result.SetStringCopy(subject_str)
		}
	}
	zend.ZendTmpStringRelease(tmp_subject_str)
	return replace_count
}
func PhpStrReplaceCommon(executeData *zend.ZendExecuteData, return_value *types.Zval, case_sensitivity int) {
	var subject *types.Zval
	var search *types.Zval
	var replace *types.Zval
	var subject_entry *types.Zval
	var zcount *types.Zval = nil
	var result types.Zval
	var string_key *types.String
	var num_key zend.ZendUlong
	var count zend.ZendLong = 0
	var argc int = executeData.NumArgs()
	for {
		for {
			fp := zpp.FastParseStart(executeData, 3, 4, 0)
			search = fp.ParseZval()
			replace = fp.ParseZval()
			subject = fp.ParseZval()
			fp.StartOptional()
			zcount = fp.ParseZval()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}

	/* Make sure we're dealing with strings and do the replacement. */

	if search.GetType() != types.IS_ARRAY {
		zend.ConvertToStringEx(search)
		if replace.GetType() != types.IS_STRING {
			zend.ConvertToStringEx(replace)
		}
	} else if replace.GetType() != types.IS_ARRAY {
		zend.ConvertToStringEx(replace)
	}
	if zend.EG__().GetException() != nil {
		return
	}

	/* if subject is an array */

	if subject.IsType(types.IS_ARRAY) {
		zend.ArrayInit(return_value)

		/* For each subject entry, convert it to string, then perform replacement
		   and add the result to the return_value array. */

		var __ht *types.Array = subject.GetArr()
		for _, _p := range __ht.ForeachData() {
			var _z *types.Zval = _p.GetVal()
			if _z.IsIndirect() {
				_z = _z.GetZv()
				if _z.IsUndef() {
					continue
				}
			}
			num_key = _p.GetH()
			string_key = _p.GetKey()
			subject_entry = _z
			subject_entry = types.ZVAL_DEREF(subject_entry)
			if subject_entry.GetType() != types.IS_ARRAY && subject_entry.GetType() != types.IS_OBJECT {
				count += PhpStrReplaceInSubject(search, replace, subject_entry, &result, case_sensitivity)
			} else {
				types.ZVAL_COPY(&result, subject_entry)
			}

			/* Add to return array */

			if string_key != nil {
				return_value.GetArr().KeyAddNew(string_key.GetStr(), &result)
			} else {
				return_value.GetArr().IndexAddNew(num_key, &result)
			}

			/* Add to return array */

		}

		/* For each subject entry, convert it to string, then perform replacement
		   and add the result to the return_value array. */

	} else {
		count = PhpStrReplaceInSubject(search, replace, subject, return_value, case_sensitivity)
	}
	if argc > 3 {
		zend.ZEND_TRY_ASSIGN_REF_LONG(zcount, count)
	}
}
func ZifStrReplace(executeData zpp.Ex, return_value zpp.Ret, search *types.Zval, replace *types.Zval, subject *types.Zval, _ zpp.Opt, replaceCount zpp.RefZval) {
	PhpStrReplaceCommon(executeData, return_value, 1)
}
func ZifStrIreplace(executeData zpp.Ex, return_value zpp.Ret, search *types.Zval, replace *types.Zval, subject *types.Zval, _ zpp.Opt, replaceCount zpp.RefZval) {
	PhpStrReplaceCommon(executeData, return_value, 0)
}
func PhpHebrev(executeData *zend.ZendExecuteData, return_value *types.Zval, convert_newlines int) {
	var str *byte
	var heb_str *byte
	var target *byte
	var tmp *byte
	var block_start int
	var block_end int
	var block_type int
	var block_length int
	var i int
	var max_chars zend.ZendLong = 0
	var char_count zend.ZendLong
	var begin int
	var end int
	var orig_begin int
	var str_len int
	var broken_str *types.String
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 2, 0)
			str, str_len = fp.ParseString()
			fp.StartOptional()
			max_chars = fp.ParseLong()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if str_len == 0 {
		return_value.SetFalse()
		return
	}
	tmp = str
	block_end = 0
	block_start = block_end
	heb_str = (*byte)(zend.Emalloc(str_len + 1))
	target = heb_str + str_len
	*target = 0
	target--
	block_length = 0
	if Isheb(*tmp) != 0 {
		block_type = _HEB_BLOCK_TYPE_HEB
	} else {
		block_type = _HEB_BLOCK_TYPE_ENG
	}
	for {
		if block_type == _HEB_BLOCK_TYPE_HEB {
			for (Isheb(int(*(tmp + 1))) != 0 || _isblank(int(*(tmp + 1))) != 0 || ispunct(int(*(tmp + 1))) || int((*(tmp + 1)) == '\n') != 0) && block_end < str_len-1 {
				tmp++
				block_end++
				block_length++
			}
			for i = block_start + 1; i <= block_end+1; i++ {
				*target = str[i-1]
				switch *target {
				case '(':
					*target = ')'
				case ')':
					*target = '('
				case '[':
					*target = ']'
				case ']':
					*target = '['
				case '{':
					*target = '}'
				case '}':
					*target = '{'
				case '<':
					*target = '>'
				case '>':
					*target = '<'
				case '\\':
					*target = '/'
				case '/':
					*target = '\\'
				default:

				}
				target--
			}
			block_type = _HEB_BLOCK_TYPE_ENG
		} else {
			for Isheb(*(tmp + 1)) == 0 && int((*(tmp + 1)) != '\n' && block_end < str_len-1) != 0 {
				tmp++
				block_end++
				block_length++
			}
			for (_isblank(int(*tmp)) != 0 || ispunct(int(*tmp))) && (*tmp) != '/' && (*tmp) != '-' && block_end > block_start {
				tmp--
				block_end--
			}
			for i = block_end + 1; i >= block_start+1; i-- {
				*target = str[i-1]
				target--
			}
			block_type = _HEB_BLOCK_TYPE_HEB
		}
		block_start = block_end + 1
		if block_end >= str_len-1 {
			break
		}
	}
	broken_str = types.ZendStringAlloc(str_len, 0)
	end = str_len - 1
	begin = end
	target = broken_str.GetVal()
	for true {
		char_count = 0
		for (max_chars == 0 || max_chars > 0 && char_count < max_chars) && begin > 0 {
			char_count++
			begin--
			if _isnewline(heb_str[begin]) != 0 {
				for begin > 0 && _isnewline(heb_str[begin-1]) != 0 {
					begin--
					char_count++
				}
				break
			}
		}
		if max_chars >= 0 && char_count == max_chars {
			var new_char_count int = char_count
			var new_begin int = begin
			for new_char_count > 0 {
				if _isblank(heb_str[new_begin]) != 0 || _isnewline(heb_str[new_begin]) != 0 {
					break
				}
				new_begin++
				new_char_count--
			}
			if new_char_count > 0 {
				begin = new_begin
			}
		}
		orig_begin = begin
		if _isblank(heb_str[begin]) != 0 {
			heb_str[begin] = '\n'
		}
		for begin <= end && _isnewline(heb_str[begin]) != 0 {
			begin++
		}
		for i = begin; i <= end; i++ {
			*target = heb_str[i]
			target++
		}
		for i = orig_begin; i <= end && _isnewline(heb_str[i]) != 0; i++ {
			*target = heb_str[i]
			target++
		}
		begin = orig_begin
		if begin == 0 {
			*target = 0
			break
		}
		begin--
		end = begin
	}
	zend.Efree(heb_str)
	if convert_newlines != 0 {
		return_value.SetString(PhpCharToStrEx(broken_str, '\n', "<br />\n", 7, 1, nil))
		// types.ZendStringReleaseEx(broken_str, 0)
	} else {
		return_value.SetString(broken_str)
		return
	}
}
func ZifHebrev(executeData zpp.Ex, return_value zpp.Ret, str *types.Zval, _ zpp.Opt, maxCharsPerLine *types.Zval) {
	PhpHebrev(executeData, return_value, 0)
}
func ZifHebrevc(executeData zpp.Ex, return_value zpp.Ret, str *types.Zval, _ zpp.Opt, maxCharsPerLine *types.Zval) {
	PhpHebrev(executeData, return_value, 1)
}

/* in brief this inserts <br /> or <br> before matched regexp \n\r?|\r\n? */
func ZifNl2br(executeData zpp.Ex, return_value zpp.Ret, str *types.Zval, _ zpp.Opt, isXhtml *types.Zval) {

	var tmp *byte
	var end *byte
	var str *types.String
	var target *byte
	var repl_cnt int = 0
	var is_xhtml types.ZendBool = 1
	var result *types.String
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 2, 0)
			str = fp.ParseStr()
			fp.StartOptional()
			is_xhtml = fp.ParseBool()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	tmp = str.GetVal()
	end = str.GetVal() + str.GetLen()

	/* it is really faster to scan twice and allocate mem once instead of scanning once
	   and constantly reallocing */

	for tmp < end {
		if (*tmp) == '\r' {
			if (*(tmp + 1)) == '\n' {
				tmp++
			}
			repl_cnt++
		} else if (*tmp) == '\n' {
			if (*(tmp + 1)) == '\r' {
				tmp++
			}
			repl_cnt++
		}
		tmp++
	}
	if repl_cnt == 0 {
		return_value.SetStringCopy(str)
		return
	}
	var repl_len int = b.CondF(is_xhtml != 0, func() int { return b.SizeOf("\"<br />\"") - 1 }, func() int { return b.SizeOf("\"<br>\"") - 1 })
	result = types.ZendStringSafeAlloc(repl_cnt, repl_len, str.GetLen(), 0)
	target = result.GetVal()
	tmp = str.GetVal()
	for tmp < end {
		switch *tmp {
		case '\r':
			fallthrough
		case '\n':
			b.PostInc(&(*target)) = '<'
			b.PostInc(&(*target)) = 'b'
			b.PostInc(&(*target)) = 'r'
			if is_xhtml != 0 {
				b.PostInc(&(*target)) = ' '
				b.PostInc(&(*target)) = '/'
			}
			b.PostInc(&(*target)) = '>'
			if (*tmp) == '\r' && (*(tmp + 1)) == '\n' || (*tmp) == '\n' && (*(tmp + 1)) == '\r' {
				*tmp++
				b.PostInc(&(*target)) = (*tmp) - 1
			}
			fallthrough
		default:
			b.PostInc(&(*target)) = *tmp
		}
		tmp++
	}
	*target = '0'
	return_value.SetString(result)
	return
}
func ZifStripTags(executeData zpp.Ex, return_value zpp.Ret, str *types.Zval, _ zpp.Opt, allowableTags *types.Zval) {
	var buf *types.String
	var str *types.String
	var allow *types.Zval = nil
	var allowed_tags *byte = nil
	var allowed_tags_len int = 0
	var tags_ss zend.SmartStr = zend.MakeSmartStr(0)
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 2, 0)
			str = fp.ParseStr()
			fp.StartOptional()
			allow = fp.ParseZval()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if allow != nil {
		if allow.IsType(types.IS_ARRAY) {
			var tmp *types.Zval
			var tag *types.String
			var __ht *types.Array = allow.GetArr()
			for _, _p := range __ht.ForeachData() {
				var _z *types.Zval = _p.GetVal()

				tmp = _z
				tag = zend.ZvalGetString(tmp)
				tags_ss.AppendByte('<')
				tags_ss.AppendString(tag.GetStr())
				tags_ss.AppendByte('>')
				// types.ZendStringRelease(tag)
			}
			if tags_ss.GetS() != nil {
				tags_ss.ZeroTail()
				allowed_tags = tags_ss.GetS().GetVal()
				allowed_tags_len = tags_ss.GetS().GetLen()
			}
		} else {

			/* To maintain a certain BC, we allow anything for the second parameter and return original string */

			zend.ConvertToString(allow)
			allowed_tags = allow.GetStr().GetVal()
			allowed_tags_len = allow.GetStr().GetLen()
		}
	}
	buf = types.NewString(str.GetStr())
	buf.SetLen(PhpStripTagsEx(buf.GetVal(), str.GetLen(), nil, allowed_tags, allowed_tags_len, 0))
	tags_ss.Free()
	return_value.SetString(buf)
	return
}
func ZifParseStr(executeData zpp.Ex, return_value zpp.Ret, encodedString *types.Zval, _ zpp.Opt, result zpp.RefZval) {
	var arg *byte
	var arrayArg *types.Zval = nil
	var res *byte = nil
	var arglen int
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 2, 0)
			arg, arglen = fp.ParseString()
			fp.StartOptional()
			arrayArg = fp.ParseZval()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	res = zend.Estrndup(arg, arglen)
	if arrayArg == nil {
		var tmp types.Zval
		var symbol_table *types.Array
		if zend.ZendForbidDynamicCall("parse_str() with a single argument") == types.FAILURE {
			zend.Efree(res)
			return
		}
		core.PhpErrorDocref(nil, faults.E_DEPRECATED, "Calling parse_str() without the result argument is deprecated")
		symbol_table = zend.ZendRebuildSymbolTable()
		tmp.SetArray(symbol_table)
		core.SM__().GetTreatData()(core.PARSE_STRING, res, &tmp)
		if types.ZendHashDel(symbol_table, types.STR_THIS) == types.SUCCESS {
			faults.ThrowError(nil, "Cannot re-assign $this")
		}
	} else {
		arrayArg = zend.ZendTryArrayInit(arrayArg)
		if arrayArg == nil {
			zend.Efree(res)
			return
		}
		core.SM__().GetTreatData()(core.PARSE_STRING, res, arrayArg)
	}
}
func PhpTagFind(tag *byte, len_ int, set *byte) int {
	var c byte
	var n *byte
	var t *byte
	var state int = 0
	var done int = 0
	var norm *byte
	if len_ == 0 {
		return 0
	}
	norm = zend.Emalloc(len_ + 1)
	n = norm
	t = tag
	c = tolower(*t)

	/*
	   normalize the tag removing leading and trailing whitespace
	   and turn any <a whatever...> into just <a> and any </tag>
	   into <tag>
	*/

	for done == 0 {
		switch c {
		case '<':
			*(b.PostInc(&n)) = c
		case '>':
			done = 1
		default:
			if !(isspace(int(c))) {
				if state == 0 {
					state = 1
				}
				if c != '/' || (*(t - 1)) != '<' && (*(t + 1)) != '>' {
					*(b.PostInc(&n)) = c
				}
			} else {
				if state == 1 {
					done = 1
				}
			}
		}
		c = tolower(*(b.PreInc(&t)))
	}
	*(b.PostInc(&n)) = '>'
	*n = '0'
	if strstr(set, norm) {
		done = 1
	} else {
		done = 0
	}
	zend.Efree(norm)
	return done
}
func PhpStripTags(rbuf *byte, len_ int, stateptr *uint8, allow *byte, allow_len int) int {
	return PhpStripTagsEx(rbuf, len_, stateptr, allow, allow_len, 0)
}
func PhpStripTagsEx(
	rbuf *byte,
	len_ int,
	stateptr *uint8,
	allow *byte,
	allow_len int,
	allow_tag_spaces types.ZendBool,
) int {
	var tbuf *byte
	var tp *byte
	var rp *byte
	var c byte
	var lc byte
	var buf *byte
	var p *byte
	var end *byte
	var br int
	var depth int = 0
	var in_q int = 0
	var state uint8 = 0
	var pos int
	var allow_free *byte = nil
	var is_xml byte = 0
	buf = zend.Estrndup(rbuf, len_)
	end = buf + len_
	lc = '0'
	p = buf
	rp = rbuf
	br = 0
	if allow != nil {
		allow_free = zend.ZendStrTolowerDupEx(allow, allow_len)
		if allow_free != nil {
			allow = allow_free
		} else {
			allow = allow
		}
		tbuf = zend.Emalloc(PHP_TAG_BUF_SIZE + 1)
		tp = tbuf
	} else {
		tp = nil
		tbuf = tp
	}
	if stateptr != nil {
		state = *stateptr
		switch state {
		case 1:
			goto state_1
		case 2:
			goto state_2
		case 3:
			goto state_3
		case 4:
			goto state_4
		default:

		}
	}
state_0:
	if p >= end {
		goto finish
	}
	c = *p
	switch c {
	case '0':

	case '<':
		if in_q != 0 {
			break
		}
		if isspace(*(p + 1)) && allow_tag_spaces == 0 {
			*(b.PostInc(&rp)) = c
			break
		}
		lc = '<'
		state = 1
		if allow != nil {
			if tp-tbuf >= PHP_TAG_BUF_SIZE {
				pos = tp - tbuf
				tbuf = zend.Erealloc(tbuf, tp-tbuf+PHP_TAG_BUF_SIZE+1)
				tp = tbuf + pos
			}
			*(b.PostInc(&tp)) = '<'
		}
		p++
		goto state_1
	case '>':
		if depth != 0 {
			depth--
			break
		}
		if in_q != 0 {
			break
		}
		*(b.PostInc(&rp)) = c
	default:
		*(b.PostInc(&rp)) = c
	}
	p++
	goto state_0
state_1:
	if p >= end {
		goto finish
	}
	c = *p
	switch c {
	case '0':

	case '<':
		if in_q != 0 {
			break
		}
		if isspace(*(p + 1)) && allow_tag_spaces == 0 {
			goto reg_char_1
		}
		depth++
	case '>':
		if depth != 0 {
			depth--
			break
		}
		if in_q != 0 {
			break
		}
		lc = '>'
		if is_xml && p >= buf+1 && (*(p - 1)) == '-' {
			break
		}
		is_xml = 0
		state = is_xml
		in_q = state
		if allow != nil {
			if tp-tbuf >= PHP_TAG_BUF_SIZE {
				pos = tp - tbuf
				tbuf = zend.Erealloc(tbuf, tp-tbuf+PHP_TAG_BUF_SIZE+1)
				tp = tbuf + pos
			}
			*(b.PostInc(&tp)) = '>'
			*tp = '0'
			if PhpTagFind(tbuf, tp-tbuf, allow) != 0 {
				memcpy(rp, tbuf, tp-tbuf)
				rp += tp - tbuf
			}
			tp = tbuf
		}
		p++
		goto state_0
	case '"':
		fallthrough
	case '\'':
		if p != buf && (in_q == 0 || (*p) == in_q) {
			if in_q != 0 {
				in_q = 0
			} else {
				in_q = *p
			}
		}
		goto reg_char_1
	case '!':

		/* JavaScript & Other HTML scripting languages */

		if p >= buf+1 && (*(p - 1)) == '<' {
			state = 3
			lc = c
			p++
			goto state_3
		} else {
			goto reg_char_1
		}
	case '?':
		if p >= buf+1 && (*(p - 1)) == '<' {
			br = 0
			state = 2
			p++
			goto state_2
		} else {
			goto reg_char_1
		}
	default:
	reg_char_1:
		if allow != nil {
			if tp-tbuf >= PHP_TAG_BUF_SIZE {
				pos = tp - tbuf
				tbuf = zend.Erealloc(tbuf, tp-tbuf+PHP_TAG_BUF_SIZE+1)
				tp = tbuf + pos
			}
			*(b.PostInc(&tp)) = c
		}
	}
	p++
	goto state_1
state_2:
	if p >= end {
		goto finish
	}
	c = *p
	switch c {
	case '(':
		if lc != '"' && lc != '\'' {
			lc = '('
			br++
		}
	case ')':
		if lc != '"' && lc != '\'' {
			lc = ')'
			br--
		}
	case '>':
		if depth != 0 {
			depth--
			break
		}
		if in_q != 0 {
			break
		}
		if br == 0 && p >= buf+1 && lc != '"' && (*(p - 1)) == '?' {
			state = 0
			in_q = state
			tp = tbuf
			p++
			goto state_0
		}
	case '"':
		fallthrough
	case '\'':
		if p >= buf+1 && (*(p - 1)) != '\\' {
			if lc == c {
				lc = '0'
			} else if lc != '\\' {
				lc = c
			}
			if p != buf && (in_q == 0 || (*p) == in_q) {
				if in_q != 0 {
					in_q = 0
				} else {
					in_q = *p
				}
			}
		}
	case 'l':
		fallthrough
	case 'L':

		/* swm: If we encounter '<?xml' then we shouldn't be in
		 * state == 2 (PHP). Switch back to HTML.
		 */

		if state == 2 && p > buf+4 && ((*(p - 1)) == 'm' || (*(p - 1)) == 'M') && ((*(p - 2)) == 'x' || (*(p - 2)) == 'X') && (*(p - 3)) == '?' && (*(p - 4)) == '<' {
			state = 1
			is_xml = 1
			p++
			goto state_1
		}
	default:

	}
	p++
	goto state_2
state_3:
	if p >= end {
		goto finish
	}
	c = *p
	switch c {
	case '>':
		if depth != 0 {
			depth--
			break
		}
		if in_q != 0 {
			break
		}
		state = 0
		in_q = state
		tp = tbuf
		p++
		goto state_0
	case '"':
		fallthrough
	case '\'':
		if p != buf && (*(p - 1)) != '\\' && (in_q == 0 || (*p) == in_q) {
			if in_q != 0 {
				in_q = 0
			} else {
				in_q = *p
			}
		}
	case '-':
		if p >= buf+2 && (*(p - 1)) == '-' && (*(p - 2)) == '!' {
			state = 4
			p++
			goto state_4
		}
	case 'E':
		fallthrough
	case 'e':

		/* !DOCTYPE exception */

		if p > buf+6 && ((*(p - 1)) == 'p' || (*(p - 1)) == 'P') && ((*(p - 2)) == 'y' || (*(p - 2)) == 'Y') && ((*(p - 3)) == 't' || (*(p - 3)) == 'T') && ((*(p - 4)) == 'c' || (*(p - 4)) == 'C') && ((*(p - 5)) == 'o' || (*(p - 5)) == 'O') && ((*(p - 6)) == 'd' || (*(p - 6)) == 'D') {
			state = 1
			p++
			goto state_1
		}
	default:

	}
	p++
	goto state_3
state_4:
	for p < end {
		c = *p
		if c == '>' && in_q == 0 {
			if p >= buf+2 && (*(p - 1)) == '-' && (*(p - 2)) == '-' {
				state = 0
				in_q = state
				tp = tbuf
				p++
				goto state_0
			}
		}
		p++
	}
finish:
	if rp < rbuf+len_ {
		*rp = '0'
	}
	zend.Efree(any(buf))
	if tbuf != nil {
		zend.Efree(tbuf)
	}
	if allow_free != nil {
		zend.Efree(allow_free)
	}
	if stateptr != nil {
		*stateptr = state
	}
	return size_t(rp - rbuf)
}
func ZifStrGetcsv(executeData zpp.Ex, return_value zpp.Ret, string *types.Zval, _ zpp.Opt, delimiter *types.Zval, enclosure *types.Zval, escape *types.Zval) {
	var str *types.String
	var delim byte = ','
	var enc byte = '"'
	var esc int = uint8('\\')
	var delim_str *byte = nil
	var enc_str *byte = nil
	var esc_str *byte = nil
	var delim_len int = 0
	var enc_len int = 0
	var esc_len int = 0
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 4, 0)
			str = fp.ParseStr()
			fp.StartOptional()
			delim_str, delim_len = fp.ParseString()
			enc_str, enc_len = fp.ParseString()
			esc_str, esc_len = fp.ParseString()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if delim_len != 0 {
		delim = delim_str[0]
	} else {
		delim = delim
	}
	if enc_len != 0 {
		enc = enc_str[0]
	} else {
		enc = enc
	}
	if esc_str != nil {
		if esc_len != 0 {
			esc = uint8(esc_str[0])
		} else {
			esc = PHP_CSV_NO_ESCAPE
		}
	}
	PhpFgetcsv(nil, delim, enc, esc, str.GetLen(), str.GetVal(), return_value)
}
func ZifStrRepeat(input string, mult int) (string, bool) {
	if mult < 0 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Second argument has to be greater than or equal to 0")
		return "", false
	}
	/* Don't waste our time if it's empty */
	if input == "" || mult == 0 {
		return "", true
	}

	return strings.Repeat(input, mult), true
}
func ZifCountChars(executeData zpp.Ex, return_value zpp.Ret, input *types.Zval, _ zpp.Opt, mode *types.Zval) {
	var input *types.String
	var chars []int
	var mymode zend.ZendLong = 0
	var buf *uint8
	var inx int
	var retstr []byte
	var retlen int = 0
	var tmp int = 0
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 2, 0)
			input = fp.ParseStr()
			fp.StartOptional()
			mymode = fp.ParseLong()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if mymode < 0 || mymode > 4 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Unknown mode")
		return_value.SetFalse()
		return
	}
	buf = (*uint8)(input.GetVal())
	memset(any(chars), 0, b.SizeOf("chars"))
	for tmp < input.GetLen() {
		chars[*buf]++
		buf++
		tmp++
	}
	if mymode < 3 {
		zend.ArrayInit(return_value)
	}
	for inx = 0; inx < 256; inx++ {
		switch mymode {
		case 0:
			zend.AddIndexLong(return_value, inx, chars[inx])
		case 1:
			if chars[inx] != 0 {
				zend.AddIndexLong(return_value, inx, chars[inx])
			}
		case 2:
			if chars[inx] == 0 {
				zend.AddIndexLong(return_value, inx, chars[inx])
			}
		case 3:
			if chars[inx] != 0 {
				retstr[b.PostInc(&retlen)] = inx
			}
		case 4:
			if chars[inx] == 0 {
				retstr[b.PostInc(&retlen)] = inx
			}
		}
	}
	if mymode >= 3 && mymode <= 4 {
		return_value.SetStringVal(b.CastStr(retstr, retlen))
		return
	}
}
func PhpStrnatcmp(executeData *zend.ZendExecuteData, return_value *types.Zval, fold_case int) {
	var s1 *types.String
	var s2 *types.String
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 2, 0)
			s1 = fp.ParseStr()
			s2 = fp.ParseStr()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	return_value.SetLong(StrnatcmpEx(s1.GetVal(), s1.GetLen(), s2.GetVal(), s2.GetLen(), fold_case))
	return
}
func ZifStrnatcmp(executeData zpp.Ex, return_value zpp.Ret, s1 *types.Zval, s2 *types.Zval) {
	PhpStrnatcmp(executeData, return_value, 0)
}
func ZifStrnatcasecmp(executeData zpp.Ex, return_value zpp.Ret, s1 *types.Zval, s2 *types.Zval) {
	PhpStrnatcmp(executeData, return_value, 1)
}
func ZifSubstrCount(executeData zpp.Ex, return_value zpp.Ret, haystack *types.Zval, needle *types.Zval, _ zpp.Opt, offset *types.Zval, length *types.Zval) {
	var haystack *byte
	var needle *byte
	var offset zend.ZendLong = 0
	var length zend.ZendLong = 0
	var ac int = executeData.NumArgs()
	var count zend.ZendLong = 0
	var haystack_len int
	var needle_len int
	var p *byte
	var endp *byte
	var cmp byte
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 4, 0)
			haystack, haystack_len = fp.ParseString()
			needle, needle_len = fp.ParseString()
			fp.StartOptional()
			offset = fp.ParseLong()
			length = fp.ParseLong()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if needle_len == 0 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Empty substring")
		return_value.SetFalse()
		return
	}
	p = haystack
	endp = p + haystack_len
	if offset < 0 {
		offset += zend.ZendLong(haystack_len)
	}
	if offset < 0 || int(offset > haystack_len) != 0 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Offset not contained in string")
		return_value.SetFalse()
		return
	}
	p += offset
	if ac == 4 {
		if length < 0 {
			length += haystack_len - offset
		}
		if length < 0 || int(length > haystack_len-offset) != 0 {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Invalid length value")
			return_value.SetFalse()
			return
		}
		endp = p + length
	}
	if needle_len == 1 {
		cmp = needle[0]
		for b.Assign(&p, memchr(p, cmp, endp-p)) {
			count++
			p++
		}
	} else {
		for b.Assign(&p, (*byte)(core.PhpMemnstr(p, needle, needle_len, endp))) {
			p += needle_len
			count++
		}
	}
	return_value.SetLong(count)
	return
}

func ZifStrPad(input string, padLength int, _ zpp.Opt, padString_ *string, padType_ *int) (string, bool) {
	padString := b.Option(padString_, " ")
	padType := b.Option(padType_, STR_PAD_RIGHT)

	/* If resulting string turns out to be shorter than input string,
	   we simply copy the input and return. */
	if padLength < 0 || padLength < len(input) {
		return input, true
	}
	if padString == "" {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Padding string cannot be empty")
		return "", false
	}
	if padType < STR_PAD_LEFT || padType > STR_PAD_BOTH {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Padding type has to be STR_PAD_LEFT, STR_PAD_RIGHT, or STR_PAD_BOTH")
		return "", false
	}
	numPadChars := padLength - len(input)
	if numPadChars >= core.INT_MAX {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Padding length is too long")
		return "", false
	}

	/* We need to figure out the left/right padding lengths. */
	var leftPad, rightPad int
	switch padType {
	case STR_PAD_RIGHT:
		leftPad = 0
		rightPad = numPadChars
	case STR_PAD_LEFT:
		leftPad = numPadChars
		rightPad = 0
	case STR_PAD_BOTH:
		leftPad = numPadChars / 2
		rightPad = numPadChars - leftPad
	}

	var buf strings.Builder
	for i := 0; i < leftPad; i++ {
		buf.WriteByte(padString[i%len(padString)])
	}
	buf.WriteString(input)
	for i := 0; i < rightPad; i++ {
		buf.WriteByte(padString[i%len(padString)])
	}

	return buf.String(), true
}
func ZifSscanf(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var args *types.Zval = nil
	var str *byte
	var format *byte
	var str_len int
	var format_len int
	var result int
	var num_args int = 0
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = -1

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			str, str_len = fp.ParseString()
			format, format_len = fp.ParseString()
			args, num_args = fp.ParseVariadic0()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	result = PhpSscanfInternal(str, format, num_args, args, 0, return_value)
	if SCAN_ERROR_WRONG_PARAM_COUNT == result {
		zend.ZendWrongParamCount()
		return
	}
}
func PhpStrRot13(str *types.String) *types.String {
	var ret *types.String
	var p *byte
	var e *byte
	var target *byte
	if str.GetLen() == 0 {
		return types.NewString("")
	}
	ret = types.ZendStringAlloc(str.GetLen(), 0)
	p = str.GetVal()
	e = p + str.GetLen()
	target = ret.GetVal()
	for p < e {
		if (*p) >= 'a' && (*p) <= 'z' {
			b.PostInc(&(*target)) = 'a' + (b.PostInc(&(*p))-'a'+13)%26
		} else if (*p) >= 'A' && (*p) <= 'Z' {
			b.PostInc(&(*target)) = 'A' + (b.PostInc(&(*p))-'A'+13)%26
		} else {
			*p++
			b.PostInc(&(*target)) = (*p) - 1
		}
	}
	*target = '0'
	return ret
}
func ZifStrRot13(executeData zpp.Ex, return_value zpp.Ret, str *types.Zval) {
	var arg *types.String
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			arg = fp.ParseStr()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	return_value.SetString(PhpStrRot13(arg))
	return
}
func PhpStringShuffle(str *byte, len_ zend.ZendLong) {
	var n_elems zend.ZendLong
	var rnd_idx zend.ZendLong
	var n_left zend.ZendLong
	var temp byte

	/* The implementation is stolen from array_data_shuffle       */

	n_elems = len_
	if n_elems <= 1 {
		return
	}
	n_left = n_elems
	for b.PreDec(&n_left) {
		rnd_idx = PhpMtRandRange(0, n_left)
		if rnd_idx != n_left {
			temp = str[n_left]
			str[n_left] = str[rnd_idx]
			str[rnd_idx] = temp
		}
	}
}
func ZifStrShuffle(executeData zpp.Ex, return_value zpp.Ret, str *types.Zval) {
	var arg *types.String
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			arg = fp.ParseStr()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	return_value.SetStringVal(arg.GetStr())
	if return_value.GetStr().GetLen() > 1 {
		PhpStringShuffle(return_value.GetStr().GetVal(), zend.ZendLong(return_value.GetStr().GetLen()))
	}
}
func ZifStrWordCount(str string, _ zpp.Opt, format int, charlist *string) (*types.Zval, bool) {
	var mask = ""
	if charlist != nil {
		mask, _ = PhpCharmaskEx(*charlist)
	}

	// find spans
	type span struct {
		start int
		end   int
	}
	spans := make([]span, 0, 32)

	start := -1
	for end, c := range []byte(str) {
		if ascii.IsAscii(c) || (mask != "" && strings.ContainsRune(mask, rune(c))) {
			if start < 0 {
				start = end
			}
		} else {
			spans = append(spans, span{start, end})
			start = -1
		}
	}
	if start > 0 {
		spans = append(spans, span{start, len(str)})
	}

	// 区分三种输出格式返回
	switch format {
	case 0:
		count := len(spans)
		return types.NewZvalLong(count), true
	case 1:
		arr := types.NewArray(len(spans))
		for _, span := range spans {
			arr.NextIndexInsert(types.NewZvalString(str[span.start:span.end]))
		}
		return types.NewZvalArray(arr), true
	case 2:
		arr := types.NewArray(len(spans))
		for _, span := range spans {
			arr.IndexUpdate(span.start, types.NewZvalString(str[span.start:span.end]))
		}
		return types.NewZvalArray(arr), true
	default:
		core.PhpErrorDocref(nil, faults.E_WARNING, "Invalid format value "+zend.ZEND_LONG_FMT, format)
		return nil, false
	}
}
func ZifMoneyFormat(executeData zpp.Ex, return_value zpp.Ret, format *types.Zval, value *types.Zval) {
	var format_len int = 0
	var format *byte
	var p *byte
	var e *byte
	var value float64
	var check types.ZendBool = 0
	var str *types.String
	var res_len ssize_t
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 2, 0)
			format, format_len = fp.ParseString()
			value = fp.ParseDouble()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	p = format
	e = p + format_len
	for b.Assign(&p, memchr(p, '%', e-p)) {
		if (*(p + 1)) == '%' {
			p += 2
		} else if check == 0 {
			check = 1
			p++
		} else {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Only a single %%i or %%n token can be used")
			return_value.SetFalse()
			return
		}
	}
	str = types.ZendStringSafeAlloc(format_len, 1, 1024, 0)
	if b.Assign(&res_len, strfmon(str.GetVal(), str.GetLen(), format, value)) < 0 {
		// types.ZendStringEfree(str)
		return_value.SetFalse()
		return
	}
	str.SetLen(int(res_len))
	str.GetVal()[str.GetLen()] = '0'
	return_value.SetString(types.ZendStringTruncate(str, str.GetLen()))
	return
}
func ZifStrSplit(executeData zpp.Ex, return_value zpp.Ret, str *types.Zval, _ zpp.Opt, splitLength *types.Zval) {
	var str *types.String
	var split_length zend.ZendLong = 1
	var p *byte
	var n_reg_segments int
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 2, 0)
			str = fp.ParseStr()
			fp.StartOptional()
			split_length = fp.ParseLong()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if split_length <= 0 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "The length of each segment must be greater than zero")
		return_value.SetFalse()
		return
	}
	if 0 == str.GetLen() || int(split_length >= str.GetLen()) != 0 {
		zend.ArrayInitSize(return_value, 1)
		zend.AddNextIndexStringl(return_value, str.GetVal(), str.GetLen())
		return
	}
	zend.ArrayInitSize(return_value, uint32((str.GetLen()-1)/split_length+1))
	n_reg_segments = str.GetLen() / split_length
	p = str.GetVal()
	for b.PostDec(&n_reg_segments) > 0 {
		zend.AddNextIndexStringl(return_value, p, split_length)
		p += split_length
	}
	if p != str.GetVal()+str.GetLen() {
		zend.AddNextIndexStringl(return_value, p, str.GetVal()+str.GetLen()-p)
	}
}
func ZifStrpbrk(executeData zpp.Ex, return_value zpp.Ret, haystack *types.Zval, charList *types.Zval) {
	var haystack *types.String
	var char_list *types.String
	var haystack_ptr *byte
	var cl_ptr *byte
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 2, 0)
			haystack = fp.ParseStr()
			char_list = fp.ParseStr()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	if char_list.GetLen() == 0 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "The character list cannot be empty")
		return_value.SetFalse()
		return
	}
	for haystack_ptr = haystack.GetVal(); haystack_ptr < haystack.GetVal()+haystack.GetLen(); haystack_ptr++ {
		for cl_ptr = char_list.GetVal(); cl_ptr < char_list.GetVal()+char_list.GetLen(); cl_ptr++ {
			if (*cl_ptr) == (*haystack_ptr) {
				return_value.SetStringVal(b.CastStr(haystack_ptr, haystack.GetVal()+haystack.GetLen()-haystack_ptr))
				return
			}
		}
	}
	return_value.SetFalse()
	return
}
func ZifSubstrCompare(executeData zpp.Ex, return_value zpp.Ret, mainStr *types.Zval, str *types.Zval, offset *types.Zval, _ zpp.Opt, length *types.Zval, caseSensitivity *types.Zval) {
	var s1 *types.String
	var s2 *types.String
	var offset zend.ZendLong
	var len_ zend.ZendLong = 0
	var len_is_default types.ZendBool = 1
	var cs types.ZendBool = 0
	var cmp_len int
	for {
		for {
			fp := zpp.FastParseStart(executeData, 3, 5, 0)
			s1 = fp.ParseStr()
			s2 = fp.ParseStr()
			offset = fp.ParseLong()
			fp.StartOptional()
			len_, len_is_default = fp.ParseLongEx(true, false)
			cs = fp.ParseBool()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	if len_is_default == 0 && len_ <= 0 {
		if len_ == 0 {
			return_value.SetLong(0)
			return
		} else {
			core.PhpErrorDocref(nil, faults.E_WARNING, "The length must be greater than or equal to zero")
			return_value.SetFalse()
			return
		}
	}
	if offset < 0 {
		offset = s1.GetLen() + offset
		if offset < 0 {
			offset = 0
		} else {
			offset = offset
		}
	}
	if int(offset > s1.GetLen()) != 0 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "The start position cannot exceed initial string length")
		return_value.SetFalse()
		return
	}
	if len_ != 0 {
		cmp_len = int(len_)
	} else {
		cmp_len = b.Max(s2.GetLen(), s1.GetLen()-offset)
	}
	if cs == 0 {
		return_value.SetLong(zend.ZendBinaryStrncmp(b.CastStr(s1.GetVal()+offset, s1.GetLen()-offset), s2.GetStr(), cmp_len))
		return
	} else {
		return_value.SetLong(zend.ZendBinaryStrncasecmpL(b.CastStr(s1.GetVal()+offset, s1.GetLen()-offset), b.CastStr(s2.GetVal(), s2.GetLen()), cmp_len))
		return
	}
}
func PhpUtf8EncodeEx(s string) string {
	var buf strings.Builder
	for _, c := range []byte(s) {
		if c < 0x80 {
			buf.WriteByte(c)
		} else {
			buf.WriteByte(0xc0 | c>>6)
			buf.WriteByte(0x80 | c&0x3f)
		}
	}
	return buf.String()
}
func PhpUtf8Decode(s *byte, len_ int) *types.String {
	var pos int = 0
	var c uint
	var str *types.String
	str = types.ZendStringAlloc(len_, 0)
	str.SetLen(0)
	for pos < len_ {
		var status int = types.FAILURE
		c = PhpNextUtf8Char((*uint8)(s), int(len_), &pos, &status)

		/* The lower 256 codepoints of Unicode are identical to Latin-1,
		 * so we don't need to do any mapping here beyond replacing non-Latin-1
		 * characters. */

		if status == types.FAILURE || c > 0xff {
			c = '?'
		}
		str.GetVal()[b.PostInc(&(str.GetLen()))] = c
	}
	str.GetVal()[str.GetLen()] = '0'
	if str.GetLen() < len_ {
		str = types.ZendStringTruncate(str, str.GetLen())
	}
	return str
}
func ZifUtf8Encode(data string) string {
	return PhpUtf8EncodeEx(data)
}
func ZifUtf8Decode(data string) string {
	return PhpUtf8Decode(data)
}
