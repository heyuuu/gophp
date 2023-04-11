package standard

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/builtin/ascii"
	"github.com/heyuuu/gophp/core"
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

func PhpStrtr(str *byte, len_ int, from string, to string) {
	s := b.CastStr(str, len_)
	newStr := phpStrtrEx(s, from, to)
	str *= newStr
}

func strReplaceStr(subject string, search *types.Zval, replace *types.Zval, caseSensitivity bool) (string, int) {
	if subject == "" {
		return "", 0
	}

	if search.IsArray() {
		var replaceStrings []string
		var replaceStr string
		if replace.IsArray() {
			replaceStrings = make([]string, replace.GetArr().Len())
			replace.GetArr().Foreach(func(key types.ArrayKey, value *types.Zval) {
				replaceStrings = append(replaceStrings, zend.ZvalGetStrVal(value))
			})
		} else {
			replaceStr = zend.ZvalGetStrVal(replace)
		}

		var result = subject
		var replaceCount = 0
		i := -1
		search.GetArr().ForeachIndirect(func(key types.ArrayKey, val *types.Zval) {
			if subject == "" {
				return
			}

			searchStr := zend.ZvalGetStrVal(val)
			if searchStr == "" {
				return
			}

			i++
			if replace.IsArray() {
				if i < len(replaceStrings) {
					replaceStr = replaceStrings[i]
				} else {
					replaceStr = ""
				}
			}

			var tmpResult string
			var count int
			if len(searchStr) == 1 {
				tmpResult, count = PhpCharToStr(result, searchStr[0], replaceStr, caseSensitivity)
			} else {
				if caseSensitivity {
					tmpResult, count = PhpStrToStrEx_Ex(result, searchStr, replaceStr)
				} else {
					lcSubjectStr := ascii.StrToLower(result)
					tmpResult, count = PhpStrToStrIEx_Ex(result, lcSubjectStr, searchStr, replaceStr)
				}
			}

			result = tmpResult
			replaceCount += count
		})
		return result, replaceCount
	} else {
		b.Assert(search.IsString())
		searchStr := search.GetStrVal()
		if searchStr == "" {
			return subject, 0
		}
		replaceStr := replace.GetStrVal()

		if len(searchStr) == 1 {
			return PhpCharToStr(subject, searchStr[0], replaceStr, caseSensitivity)
		} else {
			if caseSensitivity {
				return PhpStrToStrEx_Ex(subject, searchStr, replaceStr)
			} else {
				lcSubject := ascii.StrToLower(subject)
				return PhpStrToStrIEx_Ex(subject, lcSubject, searchStr, replaceStr)
			}
		}
	}
}
func strReplaceArray(subject *types.Array, search *types.Zval, replace *types.Zval, caseSensitivity bool) (*types.Array, int) {
	arr := types.NewArray(subject.Len())
	replaceCount := 0
	subject.ForeachIndirect(func(key types.ArrayKey, value *types.Zval) {
		value = types.ZVAL_DEREF(value)

		var result types.Zval
		if !value.IsArray() && !value.IsObject() {
			tmpResult, count := strReplaceStr(zend.ZvalGetStrVal(value), search, replace, caseSensitivity)
			result.SetStringVal(tmpResult)
			replaceCount += count
		} else {
			types.ZVAL_COPY(&result, value)
		}

		if key.IsStrKey() {
			arr.KeyAddNew(key.StrKey(), &result)
		} else {
			arr.IndexAddNew(key.IndexKey(), &result)
		}
	})
	return arr, replaceCount
}
func strReplace(returnValue *types.Zval, search *types.Zval, replace *types.Zval, subject *types.Zval, replaceCount zpp.RefZval, caseSensitivity bool) {
	// 限定参数类型
	// - str_replace(array|string $search, array|string $replace, array|string $subject, int|null &$count == null): string|array
	if !search.IsArray() {
		zend.ConvertToStringEx(search)
		zend.ConvertToStringEx(replace)
	} else if !replace.IsArray() {
		zend.ConvertToStringEx(replace)
	}
	if zend.EG__().GetException() != nil {
		return
	}

	// 主逻辑
	var count int
	if subject.IsType(types.IS_ARRAY) {
		var arr *types.Array
		arr, count = strReplaceArray(subject.GetArr(), search, replace, caseSensitivity)
		returnValue.SetArray(arr)
	} else {
		var str string
		str, count = strReplaceStr(zend.ZvalGetStrVal(subject), search, replace, caseSensitivity)
		returnValue.SetStringVal(str)
	}

	if replaceCount != nil {
		zend.ZEND_TRY_ASSIGN_REF_LONG(replaceCount, count)
	}
}

func ZifStrReplace(returnValue zpp.Ret, search *types.Zval, replace *types.Zval, subject *types.Zval, _ zpp.Opt, replaceCount zpp.RefZval) {
	strReplace(returnValue, search, replace, subject, replaceCount, true)
}
func ZifStrIreplace(returnValue zpp.Ret, search *types.Zval, replace *types.Zval, subject *types.Zval, _ zpp.Opt, replaceCount zpp.RefZval) {
	strReplace(returnValue, search, replace, subject, replaceCount, false)
}

// 判断是否为希伯来文字符的首字节
func _isHeb(c byte) bool     { return 224 <= c && c <= 250 }
func _isBlank(c byte) bool   { return c == ' ' || c == '\t' }
func _isNewline(c byte) bool { return c == '\r' || c == '\n' }
func _isPunct(c byte) bool   { panic("todo") } // todo
func PhpHebrev(str string, maxChars int, convertNewlines bool) (string, bool) {
	if str == "" {
		return "", false
	}

	hebStr := make([]byte, len(str))
	target := len(str) - 1

	blockStart, blockEnd, blockLength := 0, 0, 0
	var blockType int
	if _isHeb(str[0]) {
		blockType = _HEB_BLOCK_TYPE_HEB
	} else {
		blockType = _HEB_BLOCK_TYPE_ENG
	}

	idx := 0
	tmp := str[idx]
	for {
		if blockType == _HEB_BLOCK_TYPE_HEB {
			c := str[idx+1]
			for (_isHeb(c) || _isBlank(c) || _isPunct(c) || c == '\n') && blockEnd < len(str)-1 {
				idx++
				tmp = str[idx]
				blockEnd++
				blockLength++
			}
			for i := blockStart + 1; i <= blockEnd+1; i++ {
				switch str[i-1] {
				case '(':
					hebStr[target] = ')'
				case ')':
					hebStr[target] = '('
				case '[':
					hebStr[target] = ']'
				case ']':
					hebStr[target] = '['
				case '{':
					hebStr[target] = '}'
				case '}':
					hebStr[target] = '{'
				case '<':
					hebStr[target] = '>'
				case '>':
					hebStr[target] = '<'
				case '\\':
					hebStr[target] = '/'
				case '/':
					hebStr[target] = '\\'
				default:
					hebStr[target] = str[i-1]
				}
				target--
			}
			blockType = _HEB_BLOCK_TYPE_ENG
		} else {
			c := str[idx+1]
			for !_isHeb(c) && c != '\n' && blockEnd < len(str)-1 {
				idx++
				tmp = str[idx]
				blockEnd++
				blockLength++
			}

			for (_isBlank(tmp) || _isPunct(tmp)) && tmp != '/' && tmp != '-' && blockEnd > blockStart {
				tmp--
				blockEnd--
			}

			for i := blockEnd + 1; i >= blockStart+1; i-- {
				hebStr[target] = str[i-1]
				target--
			}
			blockType = _HEB_BLOCK_TYPE_HEB
		}
		blockStart = blockEnd + 1
		if blockEnd >= len(str)-1 {
			break
		}
	}

	begin := len(str) - 1
	end := len(str) - 1
	var brokenStr strings.Builder
	for true {
		char_count := 0
		for (maxChars == 0 || maxChars > 0 && char_count < maxChars) && begin > 0 {
			char_count++
			begin--
			if _isNewline(hebStr[begin]) {
				for begin > 0 && _isNewline(hebStr[begin-1]) {
					begin--
					char_count++
				}
				break
			}
		}
		if maxChars >= 0 && char_count == maxChars {
			var newCharCount int = char_count
			var newBegin int = begin
			for newCharCount > 0 {
				if _isBlank(hebStr[newBegin]) || _isNewline(hebStr[newBegin]) {
					break
				}
				newBegin++
				newCharCount--
			}
			if newCharCount > 0 {
				begin = newBegin
			}
		}
		origBegin := begin
		if _isBlank(hebStr[begin]) {
			hebStr[begin] = '\n'
		}
		for begin <= end && _isNewline(hebStr[begin]) {
			begin++
		}
		for i := begin; i <= end; i++ {
			brokenStr.WriteByte(hebStr[i])
		}
		for i := origBegin; i <= end && _isNewline(hebStr[i]); i++ {
			brokenStr.WriteByte(hebStr[i])
		}
		begin = origBegin
		if begin == 0 {
			break
		}
		begin--
		end = begin
	}

	if convertNewlines {
		return strings.ReplaceAll(brokenStr.String(), "\n", "<br />\n"), true
	} else {
		return brokenStr.String(), true
	}
}
func ZifHebrev(str string, _ zpp.Opt, maxCharsPerLine int) (string, bool) {
	return PhpHebrev(str, maxCharsPerLine, false)
}
func ZifHebrevc(str string, _ zpp.Opt, maxCharsPerLine int) (string, bool) {
	return PhpHebrev(str, maxCharsPerLine, true)
}

/* in brief this inserts <br /> or <br> before matched regexp \n\r?|\r\n? */
func ZifNl2br(str string, _ zpp.Opt, isXhtml_ *bool) string {
	var isXhtml bool = b.Option(isXhtml_, true)

	// 无换行符直接返回原字符串
	// notice: 这里不是查找 "\r\n" 而是查找 '\r' 或 '\n'
	if pos := strings.IndexAny(str, "\r\n"); pos < 0 {
		return str
	}

	var br string
	if isXhtml {
		br = `<br />`
	} else {
		br = `<br>`
	}

	var buf strings.Builder
	for i := 0; i < len(str); i++ {
		c := str[i]
		if c != '\r' && c != '\n' {
			buf.WriteByte(c)
			continue
		}

		buf.WriteString(br)
		buf.WriteByte(c)
		if i+1 < len(str) && (c == '\r' && str[i+1] == '\n') || (c == '\n' && str[i+1] == '\r') {
			i++
			buf.WriteByte(str[i])
		}
	}
	return buf.String()
}

func phpTagFind(tag string, set string) bool {
	if tag == "" {
		return false
	}
	var norm_ strings.Builder

	t := 0 // for tag
	c := ascii.ToLower(tag[t])
	done := false
	state := 0
	for !done {
		switch c {
		case '<':
			norm_.WriteByte(c)
		case '>':
			done = true
		default:
			if !ascii.IsSpace(c) {
				if state == 0 {
					state = 1
				}
				if c != '/' || (tag[t-1] != '<' && tag[t+1] != '>') {
					norm_.WriteByte(c)
				}
			} else {
				if state == 1 {
					done = true
				}
			}
		}
		t++
		c = ascii.ToLower(tag[t])
	}
	norm_.WriteByte('>')

	if pos := strings.Index(set, norm_.String()); pos >= 0 {
		return true
	} else {
		return false
	}
}

func PhpStripTags(str string, state uint8, allowTags string) (string, uint8) {
	var buf strings.Builder
	var tagBuf strings.Builder
	var br = 0
	var quote byte = 0 // 标识是否在字符串内，可能值有: 0, '"', '\''
	var lc byte = 0
	var depth = 0
	var isXml = false

	// 约束 state 范围
	if state < 1 || state > 4 {
		state = 0
	}

	allowTags = ascii.StrToLower(allowTags)
	for p, c := range []byte(str) {
		if state == 0 {
			switch c {
			case 0:
				break
			case '<':
				if quote != 0 {
					break
				}
				if ascii.IsSpace(str[p+1]) {
					buf.WriteByte(c)
					break
				}
				lc = '<'
				state = 1
				if allowTags != "" {
					tagBuf.WriteByte(c)
				}
			case '>':
				if depth != 0 {
					depth--
					break
				}
				if quote != 0 {
					break
				}
				buf.WriteByte(c)
			default:
				buf.WriteByte(c)
			}
		} else if state == 1 {
			switch c {
			case 0:
				break
			case '<':
				if quote != 0 {
					break
				}
				if ascii.IsSpace(str[p+1]) {
					if allowTags != "" {
						tagBuf.WriteByte(c)
					}
				} else {
					depth++
				}
			case '>':
				if depth != 0 {
					depth--
					break
				}
				if quote != 0 {
					break
				}
				lc = '>'
				if isXml && p >= 1 && str[p-1] == '-' {
					break
				}
				isXml = false
				state = 0
				quote = 0
				if allowTags != "" {
					tagBuf.WriteByte(c)
					if phpTagFind(tagBuf.String(), allowTags) {
						buf.WriteString(tagBuf.String())
					}
					tagBuf.Reset()
				}
			case '"', '\'':
				if p != 0 && (quote == 0 || str[p] == quote) {
					if quote != 0 {
						quote = 0
					} else {
						quote = str[p]
					}
				}
				if allowTags != "" {
					tagBuf.WriteByte(c)
				}
			case '!':
				/* JavaScript & Other HTML scripting languages */
				if p >= 1 && str[p-1] == '<' {
					state = 3
					lc = c
				} else {
					if allowTags != "" {
						tagBuf.WriteByte(c)
					}
				}
			case '?':
				if p >= 1 && str[p-1] == '<' {
					br = 0
					state = 2
				} else {
					if allowTags != "" {
						tagBuf.WriteByte(c)
					}
				}
			default:
				if allowTags != "" {
					tagBuf.WriteByte(c)
				}
			}
		} else if state == 2 {
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
				if quote != 0 {
					break
				}
				if br == 0 && p >= 1 && lc != '"' && str[p-1] == '?' {
					state = 0
					quote = 0
					tagBuf.Reset()
				}
			case '"', '\'':
				if p >= 1 && str[p-1] != '\\' {
					if lc == c {
						lc = 0
					} else if lc != '\\' {
						lc = c
					}
					if p != 0 && (quote == 0 || str[p] == quote) {
						if quote != 0 {
							quote = 0
						} else {
							quote = str[p]
						}
					}
				}
			case 'l', 'L':
				/* swm: If we encounter '<?xml' then we shouldn't be in
				 * state == 2 (PHP). Switch back to HTML.
				 */
				if state == 2 && p > 4 && ascii.StrToLower(str[p-4:p]) == "<?xm" {
					state = 1
					isXml = true
				}
			}
		} else if state == 3 {
			switch c {
			case '>':
				if depth != 0 {
					depth--
					break
				}
				if quote != 0 {
					break
				}
				state = 0
				quote = 0
				tagBuf.Reset()
			case '"', '\'':
				if p != 0 && str[p-1] != '\\' && (quote == 0 || str[p] == quote) {
					if quote != 0 {
						quote = 0
					} else {
						quote = str[p]
					}
				}
			case '-':
				if p >= 2 && str[p-2:p] == "!-" {
					state = 4
				}
			case 'E', 'e':
				/* !DOCTYPE exception */
				if p > 6 && ascii.StrToLower(str[p-6:p+1]) == "doctype" {
					state = 1
				}
			}
		} else { // stage == 4
			if c == '>' && quote == 0 {
				if p >= 2 && str[p-2:p] == "--" {
					state = 0
					quote = 0
					tagBuf.Reset()
				}
			}
		}
	}

	return buf.String(), state
}
func ZifStripTags(str string, _ zpp.Opt, allowableTags *types.Zval) string {
	var allow *types.Zval = allowableTags
	var allowTagsStr string
	if allow != nil {
		if allow.IsType(types.IS_ARRAY) {
			var buf strings.Builder
			allow.GetArr().Foreach(func(key types.ArrayKey, value *types.Zval) {
				tag := zend.ZvalGetStrVal(value)
				buf.WriteByte('<')
				buf.WriteString(tag)
				buf.WriteByte('>')
			})
			allowTagsStr = buf.String()
		} else {
			/* To maintain a certain BC, we allow anything for the second parameter and return original string */
			zend.ConvertToString(allow)
			allowTagsStr = allowableTags.GetStrVal()
		}
	}

	result, _ := PhpStripTags(str, 0, allowTagsStr)
	return result
}
func ZifParseStr(encodedString string, _ zpp.Opt, result zpp.RefZval) {
	if result == nil {
		if zend.ZendForbidDynamicCall("parse_str() with a single argument") == types.FAILURE {
			return
		}
		core.PhpErrorDocref(nil, faults.E_DEPRECATED, "Calling parse_str() without the result argument is deprecated")
		var tmp types.Zval
		var symbol_table *types.Array
		symbol_table = zend.ZendRebuildSymbolTable()
		tmp.SetArray(symbol_table)
		core.SM__().GetTreatData()(core.PARSE_STRING, encodedString, &tmp)
		if types.ZendHashDel(symbol_table, types.STR_THIS) == types.SUCCESS {
			faults.ThrowError(nil, "Cannot re-assign $this")
		}
	} else {
		result = zend.ZendTryArrayInit(result)
		if result == nil {
			return
		}
		core.SM__().GetTreatData()(core.PARSE_STRING, encodedString, result)
	}
}

func ZifStrGetcsv(return_value zpp.Ret, string_ string, _ zpp.Opt, delimiter *string, enclosure *string, escape *string) *types.Array {
	var str *types.String = string_
	var delim byte = ','
	var enc byte = '"'
	var esc byte = '\\'

	if delimiter != nil && *delimiter != "" {
		delim = (*delimiter)[0]
	}
	if enclosure != nil && *enclosure != "" {
		enc = (*enclosure)[0]
	}
	if escape != nil {
		if *escape != "" {
			esc = (*escape)[0]
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
func ZifCountChars(input string, _ zpp.Opt, mode int) (*types.Zval, bool) {
	if mode < 0 || mode > 4 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Unknown mode")
		return nil, false
	}

	var charCount [256]int

	for _, c := range []byte(input) {
		charCount[c]++
	}

	if mode < 3 { // mode=0,1,2 以数组返回
		arr := types.NewArray(0)
		for i := 0; i < 256; i++ {
			count := charCount[i]
			if mode == 0 || (mode == 1 && count != 0) || (mode == 2 && count == 0) {
				arr.IndexUpdate(i, types.NewZvalLong(charCount[i]))
			}
		}
		return types.NewZvalArray(arr), true
	} else { // mode=3,4 以字符串返回
		var str []byte
		for i := 0; i < 256; i++ {
			count := charCount[i]
			if (mode == 3 && count != 0) || (mode == 4 && count == 0) {
				str = append(str, byte(i))
			}
		}
		return types.NewZvalString(string(str)), true
	}
}
func ZifStrnatcmp(s1 string, s2 string) int {
	return Strnatcmp(s1, s2, false)
}
func ZifStrnatcasecmp(s1 string, s2 string) int {
	return Strnatcmp(s1, s2, true)
}
func ZifSubstrCount(haystack string, needle string, _ zpp.Opt, offset int, length_ *int) (int, bool) {
	// check needle
	if needle == "" {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Empty substring")
		return 0, false
	}

	// check offset
	if offset < 0 {
		offset += len(haystack)
	}
	if offset < 0 || offset > len(haystack) {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Offset not contained in string")
		return 0, false
	}

	// check length
	var length = len(haystack) - offset
	if length_ != nil {
		length = *length_
		if length < 0 {
			length += len(haystack) - offset
		}
		if length < 0 || length > len(haystack)-offset {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Invalid length value")
			return 0, false
		}
	}

	// 截取目标字符串范围
	searchStr := haystack[offset : offset+length]

	return strings.Count(searchStr, needle), true
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
func ZifSscanf(str string, format string, vars []zpp.RefZval) *types.Zval {
	var args *types.Zval = nil
	var result int
	var num_args int = 0
	result = PhpSscanfInternal(str, format, num_args, args, 0, return_value)
	if SCAN_ERROR_WRONG_PARAM_COUNT == result {
		zend.ZendWrongParamCount()
		return
	}
}

// 参数执行 ROT13 编码并将结果字符串返回。
func ZifStrRot13(str string) string {
	if str == "" {
		return ""
	}

	bin := []byte(str)
	for i, c := range bin {
		if 'a' <= c && c <= 'z' {
			bin[i] = 'a' + (c-'a'+13)%26
		} else if 'A' <= c && c <= 'Z' {
			bin[i] = 'A' + (c-'A'+13)%26
		}
	}
	return string(bin)
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
