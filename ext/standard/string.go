package standard

import (
	"encoding/hex"
	"fmt"
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/builtin/ascii"
	"github.com/heyuuu/gophp/builtin/strutil"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
	"github.com/heyuuu/gophp/zend/zpp"
	"strconv"
	"strings"
)

/**
 * Constants
 */

const STR_PAD_LEFT = 0
const STR_PAD_RIGHT = 1
const STR_PAD_BOTH = 2
const PHP_PATHINFO_DIRNAME = 1
const PHP_PATHINFO_BASENAME = 2
const PHP_PATHINFO_EXTENSION = 4
const PHP_PATHINFO_FILENAME = 8
const PHP_PATHINFO_ALL = PHP_PATHINFO_DIRNAME | PHP_PATHINFO_BASENAME | PHP_PATHINFO_EXTENSION | PHP_PATHINFO_FILENAME

const _HEB_BLOCK_TYPE_ENG = 1
const _HEB_BLOCK_TYPE_HEB = 2

const PHP_TAG_BUF_SIZE = 1023

/**
 * helpers
 */
func PhpStringToupper(s *types.String) *types.String {
	return types.NewString(ascii.StrToUpper(s.GetStr()))
}

func PhpStringTolower(s *types.String) *types.String {
	return types.NewString(ascii.StrToLower(s.GetStr()))
}

func substr(str string, offset int, length *int) (string, bool) {
	negativeOffset := offset < 0
	rawLen := len(str)
	// handle offset
	if offset > len(str) {
		return "", false
	} else if offset < 0 {
		/* if "offset" position is negative, count start position from the end
		 * of the string
		 */
		offset += len(str)
		if offset < 0 {
			offset = 0
		}
	}
	if offset > 0 {
		str = str[offset:]
	}

	// handle length
	if length != nil {
		l := *length
		if l < 0 {
			/* if "length" position is negative, set it to the length
			 * needed to stop that many chars from the end of the string
			 */
			if -l > len(str) {
				if negativeOffset && -l <= rawLen {
					l = 0
				} else {
					return "", false
				}
			} else {
				l = len(str) + l
			}
		}

		if l < len(str) {
			str = str[:l]
		}
	}

	return str, true
}

func PhpAddslashes(str string) string {
	if str == "" {
		return ""
	}
	if pos := strings.IndexByte(str, '\\'); pos < 0 {
		return str
	}
	replacer := strings.NewReplacer(
		"\\000", "\\0",
		`'`, `\'`,
		`"`, `\"`,
		`\`, `\\`,
	)
	return replacer.Replace(str)
}

func PhpAddcslashes(str string, what string) string {
	mask, _ := PhpCharmaskEx(what)

	strings.NewReplacer()

	var buf strings.Builder
	for _, c := range []byte(str) {
		if strings.ContainsRune(mask, rune(c)) {
			if c < 32 || c > 126 {
				buf.WriteByte('\\')
				switch c {
				case '\n':
					buf.WriteByte('n')
				case '\t':
					buf.WriteByte('t')
				case '\r':
					buf.WriteByte('r')
				case '\a':
					buf.WriteByte('a')
				case '\v':
					buf.WriteByte('v')
				case '\b':
					buf.WriteByte('b')
				case '\f':
					buf.WriteByte('f')
				default:
					buf.WriteString(fmt.Sprintf("%03o", c))
				}
				continue
			}
			buf.WriteByte('\\')
		}
		buf.WriteByte(c)
	}

	return buf.String()
}

/**
 * Zif functions
 */

func ZifBin2hex(data string) string {
	return hex.EncodeToString([]byte(data))
}
func ZifHex2bin(data string) (string, bool) {
	if len(data)%2 != 0 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Hexadecimal input string must have an even length")
		return "", false
	}

	bin, err := hex.DecodeString(data)
	if err != nil {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Input string must be hexadecimal string")
		return "", false
	}
	return string(bin), true
}

func ZifStrspn(str string, mask string, _ zpp.Opt, offset int, length *int) (int, bool) {
	if offset > len(str) {
		return 0, false
	}

	// 此处忽略了 substr 中 length 为负数时可能返回 false 的情况
	str, _ = substr(str, offset, length)

	return PhpStrspnEx(str, mask), true
}
func ZifStrcspn(str string, mask string, _ zpp.Opt, offset int, length *int) (int, bool) {
	if offset > len(str) {
		return 0, false
	}

	// 此处忽略了 substr 中 length 为负数时可能返回 false 的情况
	str, _ = substr(str, offset, length)

	return PhpStrcspnEx(str, mask), true
}
func ZifStrcoll(str1 string, str2 string) int {
	return strings.Compare(str1, str2)
}
func PhpCharmaskEx(input string) (string, bool) {
	if pos := strings.Index(input, ".."); pos < 0 {
		return input, true
	}

	var buf strings.Builder
	for {
		pos := strings.Index(input, "..")
		if pos < 0 {
			buf.WriteString(input)
			break
		}

		// e.g. "a..z"
		if pos > 0 && pos+2 < len(input)-1 && input[pos-1] <= input[pos+2] {
			buf.WriteString(input[:pos-1])
			for c := input[pos-1]; c < input[pos+2]; c++ {
				buf.WriteByte(c)
			}
			input = input[pos+3:]
		} else {
			/* Error, try to be as helpful as possible:
			   (a range ending/starting with '.' won't be captured here) */
			if pos == 0 {
				core.PhpErrorDocref(nil, faults.E_WARNING, "Invalid '..'-range, no character to the left of '..'")
				return "", false
			}
			if pos+2 >= len(input)-1 {
				core.PhpErrorDocref(nil, faults.E_WARNING, "Invalid '..'-range, no character to the right of '..'")
				return "", false
			}
			if input[pos-1] > input[pos+2] {
				core.PhpErrorDocref(nil, faults.E_WARNING, "Invalid '..'-range, '..'-range needs to be incrementing")
				return "", false
			}

			/* FIXME: better error (a..b..c is the only left possibility?) */
			core.PhpErrorDocref(nil, faults.E_WARNING, "Invalid '..'-range")
			return "", false
		}

	}
	return buf.String(), true
}
func PhpTrimAll(str string, what *string) string {
	var cutset = " \n\r\t\v\x00"
	if what != nil {
		cutset, _ = PhpCharmaskEx(*what)
	}
	return strings.Trim(str, cutset)
}
func PhpTrimLeft(str string, what *string) string {
	var cutset = " \n\r\t\v\x00"
	if what != nil {
		cutset, _ = PhpCharmaskEx(*what)
	}
	return strings.TrimLeft(str, cutset)
}
func PhpTrimRight(str string, what *string) string {
	var cutset = " \n\r\t\v\x00"
	if what != nil {
		cutset, _ = PhpCharmaskEx(*what)
	}
	return strings.TrimRight(str, cutset)
}

func ZifTrim(str string, _ zpp.Opt, characterMask *string) string {
	return PhpTrimAll(str, characterMask)
}

//@zif -alias chop
func ZifRtrim(str string, _ zpp.Opt, characterMask *string) string {
	return PhpTrimLeft(str, characterMask)
}
func ZifLtrim(str string, _ zpp.Opt, characterMask *string) string {
	return PhpTrimRight(str, characterMask)
}

func ZifWordwrap(str string, _ zpp.Opt, width *int, break_ *string, cut bool) (string, bool) {
	var linelength = b.Option(width, 75)
	var breakchar = b.Option(break_, "\n")

	if str == "" {
		return "", true
	}
	if breakchar == "" {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Break string cannot be empty")
		return "", false
	}
	if linelength == 0 && cut {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Can't force cut when width is zero")
		return "", false
	}

	/* Special case for a single-character break as it needs no
	   additional storage space */

	if len(breakchar) == 1 && !cut {
		lastspace := 0
		laststart := 0

		bin := []byte(str)
		for i, c := range bin {
			if c == breakchar[0] {
				lastspace = i + 1
				laststart = i + 1
			} else if str[i] == ' ' {
				if i-laststart >= linelength {
					bin[i] = breakchar[0]
					laststart = i + 1
				}
				lastspace = i
			} else if i-laststart >= linelength && laststart != lastspace {
				bin[lastspace] = breakchar[0]
				laststart = lastspace + 1
			}
		}
		return string(bin), true
	} else {
		/* Multiple character line break or forced cut */

		var buf strings.Builder

		lastspace := 0
		laststart := 0
		current := 0
		for current = 0; current < len(str); current++ {
			/* when we hit an existing break, copy to new buffer, and
			 * fix up laststart and lastspace */
			c := str[current]

			if strings.HasPrefix(str[current:], breakchar) {
				buf.WriteString(str[laststart : current+len(breakchar)])
				current += len(breakchar) - 1
				lastspace = current + 1
				laststart = lastspace
			} else if c == ' ' {
				if current-laststart >= linelength {
					buf.WriteString(str[laststart:current])
					buf.WriteString(breakchar)

					laststart = current + 1
				}
				lastspace = current
			} else if current-laststart >= linelength && cut && laststart >= lastspace {
				buf.WriteString(str[laststart:current])
				buf.WriteString(breakchar)

				lastspace = current
				laststart = lastspace
			} else if current-laststart >= linelength && laststart < lastspace {
				buf.WriteString(str[laststart:lastspace])
				buf.WriteString(breakchar)

				lastspace = lastspace + 1
				laststart = lastspace
			}
		}

		/* copy over any stragglers */
		if laststart != current {
			buf.WriteString(str[laststart:current])
		}

		return buf.String(), true
	}
}
func ZifExplode(separator string, str string, _ zpp.Opt, limit_ *int) ([]string, bool) {
	var limit = zend.ZEND_LONG_MAX
	if limit_ != nil {
		limit = *limit_
	}

	if len(separator) == 0 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Empty delimiter")
		return nil, false
	}

	// doc: If the limit parameter is zero, then this is treated as 1.
	if limit == 0 {
		limit = 1
	}

	var arr []string
	if limit > 0 {
		arr = strings.SplitN(str, separator, limit)
	} else { // limit < 0
		arr = strings.Split(str, separator)
		if len(arr) > -limit {
			arr = arr[:len(arr)+limit]
		} else {
			arr = nil
		}
	}
	return arr, true
}

//@zif -alias join
func ZifImplode(glue_ *types.Zval, _ zpp.Opt, pieces_ *types.Zval) string {
	var arg1 = glue_
	var arg2 = pieces_
	var pieces *types.Array
	var glue string

	// 兼容多种参数传递方法，但后两种会有 E_DEPRECATED 提示
	// - implode(string $separator, array $array)
	// - implode(array $array)
	// - implode(array $array, string $separator)
	if arg2 == nil {
		if !arg1.IsArray() {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Argument must be an array")
			return ""
		}
		glue = ""
		pieces = arg1.GetArr()
	} else {
		if arg1.IsType(types.IS_ARRAY) {
			glue = zend.ZvalGetStrVal(arg2)
			pieces = arg1.GetArr()
			core.PhpErrorDocref(nil, faults.E_DEPRECATED, "Passing glue string after array is deprecated. Swap the parameters")
		} else if arg2.IsType(types.IS_ARRAY) {
			glue = zend.ZvalGetStrVal(arg1)
			pieces = arg2.GetArr()
		} else {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Invalid arguments passed")
			return ""
		}
	}
	return PhpImplode(glue, pieces)
}
func PhpImplode(glue string, pieces *types.Array) string {
	var parts []string
	pieces.ForeachIndirect(func(_ types.ArrayKey, value *types.Zval) {
		parts = append(parts, zend.ZvalGetStrVal(value))
	})
	return strings.Join(parts, glue)
}

type strTokState struct {
	str string
}

func ZifStrtok(str string, _ zpp.Opt, token_ *string) (string, bool) {
	// todo 改为 state 参数传入
	var state *strTokState = &BG__().strTokState

	// 两种参数形式
	// - strtok(string $string, string $token): string|false
	// - strtok(string $token): string|false
	var token string
	if token_ != nil {
		*state = strTokState{
			str: str,
		}
		token = *token_
	} else {
		token = str
	}

	if state.str == "" {
		return "", false
	}

	/* Skip leading delimiters */
	state.str = strings.TrimLeft(state.str, token)
	if state.str == "" {
		return "", false
	}

	/* We know at this place that *p is no delimiter, so skip it */
	if pos := strings.IndexAny(state.str, token); pos >= 0 {
		result := state.str[:pos]
		state.str = state.str[pos+1:]
		return result, true
	} else {
		result := state.str
		state.str = ""
		return result, true
	}
}
func ZifStrtoupper(str string) string {
	return ascii.StrToUpper(str)
}
func ZifStrtolower(str string) string {
	return ascii.StrToLower(str)
}
func PhpBasenameZStr(str string, suffix string) *types.String {
	return types.NewString(PhpBasename(str, suffix))
}
func PhpBasename(s string, suffix string) string {
	if s == "" {
		return s
	}
	if s[len(s)-1] == '/' {
		s = s[:len(s)-1]
	}
	if pos := strings.LastIndexByte(s, '/'); pos >= 0 {
		s = s[pos+1:]
	}
	if suffix != "" && strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}
func ZifBasename(path string, _ zpp.Opt, suffix string) string {
	return PhpBasename(path, suffix)
}
func ZifDirname(path string, _ zpp.Opt, levels_ *int) string {
	var levels = 1
	if levels_ != nil {
		levels = *levels_
	}

	if levels == 1 {
		/* Default case */
		return zend.ZendDirname(path)
	} else if levels < 1 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Invalid argument, levels must be >= 1")
		return ""
	} else {
		/* Some levels up */
		dir := path
		for i := 0; i < levels; i++ {
			newDir := zend.ZendDirname(dir)
			if newDir == dir {
				break
			}
			dir = newDir
		}
		return dir
	}
}

func ZifPathinfo(path string, _ zpp.Opt, options *int) *types.Zval {
	opt := b.Option(options, PHP_PATHINFO_ALL)

	var tmp *types.Zval = types.NewZvalEmptyArray()
	if (opt & PHP_PATHINFO_DIRNAME) == PHP_PATHINFO_DIRNAME {
		dirname := zend.ZendDirname(path)
		if dirname != "" {
			zend.AddAssocStr(tmp, "dirname", dirname)
		}
	}

	basename := PhpBasename(path, "")
	if (opt & PHP_PATHINFO_BASENAME) == PHP_PATHINFO_BASENAME {
		zend.AddAssocStr(tmp, "basename", basename)
	}
	if (opt & PHP_PATHINFO_EXTENSION) == PHP_PATHINFO_EXTENSION {
		if pos := strings.LastIndexByte(basename, '.'); pos >= 0 {
			zend.AddAssocStr(tmp, "extension", basename[pos+1:])
		}
	}
	if (opt & PHP_PATHINFO_FILENAME) == PHP_PATHINFO_FILENAME {
		if pos := strings.LastIndexByte(basename, '.'); pos >= 0 {
			zend.AddAssocStr(tmp, "filename", basename[:pos])
		} else {
			zend.AddAssocStr(tmp, "filename", basename)
		}
	}

	if opt == PHP_PATHINFO_ALL {
		return tmp
	} else {
		var element *types.Zval = types.ZendHashGetCurrentData(tmp.GetArr())
		if element != nil {
			return element
		} else {
			return types.NewZvalString("")
		}
	}
}

func PhpStrspnEx(s1 string, s2 string) int {
	if s1 == "" {
		return 0
	}
	for i, c := range []byte(s1) {
		if !strings.ContainsRune(s2, rune(c)) {
			return i
		}
	}
	return len(s1)
}
func PhpStrcspnEx(s1 string, s2 string) int {
	if s1 == "" {
		return 0
	}
	for i, c := range []byte(s1) {
		if strings.ContainsRune(s2, rune(c)) {
			return i
		}
	}
	return len(s1)
}
func PhpNeedleChar(needle *types.Zval) (byte, bool) {
	switch needle.GetType() {
	case types.IS_LONG:
		return byte(needle.GetLval()), true
	case types.IS_NULL, types.IS_FALSE:
		return 0, true
	case types.IS_TRUE:
		return 1, true
	case types.IS_DOUBLE, types.IS_OBJECT:
		return byte(zend.ZvalGetLong(needle)), true
	default:
		core.PhpErrorDocref(nil, faults.E_WARNING, "needle is not a string or an integer")
		return 0, false
	}
}
func ZifStristr(haystack string, needle *types.Zval, _ zpp.Opt, part bool) (string, bool) {
	needleStr, ok := parseNeedle(needle)
	if !ok {
		return "", false
	}
	if needleStr == "" {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Empty needle")
		return "", false
	}

	haystackLc := ascii.StrToLower(haystack)
	needleStrLc := ascii.StrToLower(needleStr)
	if pos := strings.Index(haystackLc, needleStrLc); pos >= 0 {
		if part {
			return haystack[:pos], true
		} else {
			return haystack[pos:], true
		}
	} else {
		return "", false
	}
}

//@zif -alias strchr
func ZifStrstr(haystack string, needle *types.Zval, _ zpp.Opt, part bool) (string, bool) {
	needleStr, ok := parseNeedle(needle)
	if !ok {
		return "", false
	}
	if needleStr == "" {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Empty needle")
		return "", false
	}

	if pos := strings.Index(haystack, needleStr); pos >= 0 {
		if part {
			return haystack[:pos], true
		} else {
			return haystack[pos:], true
		}
	} else {
		return "", false
	}
}

func posSubstr(str string, offset int) (string, bool) {
	if offset < 0 {
		offset += len(str)
	}
	if offset < 0 || offset > len(str) {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Offset not contained in string")
		return "", false
	}
	if offset == 0 {
		str = str[offset:]
	}
	return str, true
}
func parseNeedle(needle *types.Zval) (string, bool) {
	if needle.IsString() {
		return needle.GetStrVal(), true
	} else {
		/*
		 * 在 PHP 8.0.0 之前，如果 needle 不是字符串，它将被转换为整数并作为字符的序数值应用。
		 * 从 PHP 7.3.0 开始，这种行为已被废弃，不鼓励依赖它。根据预期的行为，应该明确地将 needle 转换成字符串，或者明确地调用 chr()。
		 */
		needleChar, ok := PhpNeedleChar(needle)
		if !ok {
			return "", false
		}

		core.PhpErrorDocref(nil, faults.E_DEPRECATED, "Non-string needles will be interpreted as strings in the future. "+"Use an explicit chr() call to preserve the current behavior")

		return string([]byte{needleChar}), true
	}
}

func ZifStrpos(haystack string, needle *types.Zval, _ zpp.Opt, offset int) (int, bool) {
	haystack, ok := posSubstr(haystack, offset)
	if !ok {
		return 0, false
	}
	if len(haystack) == 0 {
		return 0, false
	}

	needleStr, ok := parseNeedle(needle)
	if !ok {
		return 0, false
	}

	if pos := strings.Index(haystack, needleStr); pos < 0 {
		return pos, true
	} else {
		return 0, false
	}
}
func ZifStripos(haystack string, needle *types.Zval, _ zpp.Opt, offset int) (int, bool) {
	haystack, ok := posSubstr(haystack, offset)
	if !ok {
		return 0, false
	}
	if len(haystack) == 0 {
		return 0, false
	}

	needleStr, ok := parseNeedle(needle)
	if !ok {
		return 0, false
	}

	haystack = ascii.StrToLower(haystack)
	needleStr = ascii.StrToLower(needleStr)
	if pos := strings.Index(haystack, needleStr); pos < 0 {
		return pos, true
	} else {
		return 0, false
	}
}
func ZifStrrpos(haystack string, needle *types.Zval, _ zpp.Opt, offset int) (int, bool) {
	needleStr, ok := parseNeedle(needle)
	if !ok {
		return 0, false
	}

	if len(haystack) == 0 {
		return 0, false
	}

	if offset >= 0 {
		haystack, ok = posSubstr(haystack, offset)
		if !ok {
			return 0, false
		}
		if pos := strings.Index(haystack, needleStr); pos >= 0 {
			return pos, true
		} else {
			return 0, false
		}
	} else { // offset < 0
		offset += len(haystack)
		if offset < 0 {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Offset is greater than the length of haystack string")
			return 0, false
		}

		if offset+len(needleStr) < len(haystack) {
			haystack = haystack[:offset+len(needleStr)]
		}
		if pos := strings.LastIndex(haystack, needleStr); pos >= 0 {
			return pos, true
		} else {
			return 0, false
		}
	}
}
func ZifStrripos(haystack string, needle *types.Zval, _ zpp.Opt, offset int) (int, bool) {
	needleStr, ok := parseNeedle(needle)
	if !ok {
		return 0, false
	}

	if len(haystack) == 0 {
		return 0, false
	}

	haystack = ascii.StrToLower(haystack)
	needleStr = ascii.StrToLower(needleStr)
	if offset >= 0 {
		haystack, ok = posSubstr(haystack, offset)
		if !ok {
			return 0, false
		}
		if pos := strings.Index(haystack, needleStr); pos >= 0 {
			return pos, true
		} else {
			return 0, false
		}
	} else { // offset < 0
		offset += len(haystack)
		if offset < 0 {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Offset is greater than the length of haystack string")
			return 0, false
		}

		if offset+len(needleStr) < len(haystack) {
			haystack = haystack[:offset+len(needleStr)]
		}
		if pos := strings.LastIndex(haystack, needleStr); pos >= 0 {
			return pos, true
		} else {
			return 0, false
		}
	}
}
func ZifStrrchr(haystack string, needle *types.Zval) (string, bool) {
	needleStr, ok := parseNeedle(needle)
	if !ok || needleStr == "" {
		return "", false
	}

	if pos := strings.LastIndexByte(haystack, needleStr[0]); pos >= 0 {
		return haystack[pos:], true
	} else {
		return "", false
	}
}
func ZifChunkSplit(str string, _ zpp.Opt, chunklen_ *int, ending_ *string) (string, bool) {
	chunklen := b.Option(chunklen_, 76)
	ending := b.Option(ending_, "\r\n")

	if chunklen <= 0 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Chunk length should be greater than zero")
		return "", false
	}

	// fast
	if ending == "" {
		return str, true
	}
	if chunklen >= len(str) {
		return str + ending, true
	}

	// common
	var buf strings.Builder
	for i := 0; i < len(str); i += chunklen {
		if i+chunklen <= len(str) {
			buf.WriteString(str[i : i+chunklen])
		} else {
			buf.WriteString(str[i:])
		}
		buf.WriteString(ending)
	}
	return buf.String(), true
}
func ZifSubstr(str string, offset int, _ zpp.Opt, length *int) (string, bool) {
	return substr(str, offset, length)
}

func substrReplaceSingle(str string, replace string, start int, l int) string {
	/* if "start" position is negative, count start position start the end
	 * of the string
	 */
	f := start
	if f < 0 {
		f = f + len(str)
		if f < 0 {
			f = 0
		}
	} else if f > len(str) {
		f = len(str)
	}

	if l < 0 {
		l = len(str) - f + l
		if l < 0 {
			l = 0
		}
	} else if l > len(str) {
		l = len(str)
	}
	b.Assert(0 <= f && f <= zend.ZEND_LONG_MAX)
	b.Assert(0 <= l && l <= zend.ZEND_LONG_MAX)
	if f+l > len(str) {
		l = len(str) - f
	}

	return str[:f] + replace + str[f+l:]
}

func substrReplaceStr(str string, replace *types.Zval, start *types.Zval, length *types.Zval) string {
	// str 为字符串时，允许的参数类型:
	// - substr_replace(string, array|string, int, int|null)
	// 其他情况都会触发 warning 并返回原字符串
	if start.IsArray() {
		if length == nil || !length.IsArray() {
			core.PhpErrorDocref(nil, faults.E_WARNING, "'start' and 'length' should be of same type - numerical or array ")
		} else if length.IsArray() && start.GetArr().Len() != length.GetArr().Len() {
			core.PhpErrorDocref(nil, faults.E_WARNING, "'start' and 'length' should have the same number of elements")
		} else {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Functionality of 'start' and 'length' as arrays is not implemented")
		}
		return str
	}
	if length != nil && length.IsArray() {
		core.PhpErrorDocref(nil, faults.E_WARNING, "'start' and 'length' should be of same type - numerical or array ")
		return str
	}

	// 正常处理 substr_replace(string, array|string, int, int|null)
	l := len(str)
	if length != nil && !length.IsArray() {
		l = length.GetLval()
	}

	//
	var replStr string
	if replace.IsArray() {
		_, replZval := replace.GetArr().First()
		if replZval != nil {
			replStr = zend.ZvalGetStrVal(replZval)
		} else {
			replStr = ""
		}
	} else {
		replStr = replace.GetStrVal()
	}

	res := substrReplaceSingle(str, replStr, start.GetLval(), l)
	return res
}

func substrReplaceArray(str *types.Array, replace *types.Zval, start *types.Zval, length *types.Zval) *types.Array {
	arr := types.NewArray(str.Len())

	var replaceStr []string
	if replace.IsArray() {
		replace.GetArr().Foreach(func(_ types.ArrayKey, value *types.Zval) {
			replaceStr = append(replaceStr, zend.ZvalGetStrVal(value))
		})
	}

	var startPoints []int
	if start.IsArray() {
		start.GetArr().Foreach(func(_ types.ArrayKey, value *types.Zval) {
			startPoints = append(startPoints, zend.ZvalGetLong(value))
		})
	}

	var lengthPoints []int
	if length != nil && length.IsArray() {
		length.GetArr().Foreach(func(_ types.ArrayKey, value *types.Zval) {
			lengthPoints = append(lengthPoints, zend.ZvalGetLong(value))
		})
	}

	idx := -1
	str.ForeachIndirect(func(key types.ArrayKey, value *types.Zval) {
		idx++

		origStr := zend.ZvalGetStrVal(value)

		// f
		var f int = 0
		if start.IsArray() {
			if idx < len(startPoints) {
				f = startPoints[idx]
			}
		} else {
			f = start.GetLval()
		}

		// l
		var l int = len(origStr)
		if length != nil {
			if length.IsArray() {
				if idx < len(lengthPoints) {
					l = lengthPoints[idx]
				}
			} else {
				l = length.GetLval()
			}
		}

		// repl
		var replStr string = ""
		if replace.IsArray() {
			if idx <= len(replaceStr) {
				replStr = replaceStr[idx]
			}
		} else {
			replStr = replace.GetStrVal()
		}

		ret := substrReplaceSingle(origStr, replStr, f, l)
		if key.IsStrKey() {
			arr.SymtableUpdate(key.StrKey(), types.NewZvalString(ret))
		} else {
			arr.IndexUpdate(key.IndexKey(), types.NewZvalString(ret))
		}
	})
	return arr
}

func ZifSubstrReplace(returnValue zpp.Ret, str *types.Zval, replace *types.Zval, start *types.Zval, _ zpp.Opt, length *types.Zval) {
	// 限定参数类型
	// - substr_replace(array|string $str, array|string $replace, array|int $start, array|int|null $length = null)
	if !str.IsArray() {
		zend.ConvertToStringEx(str)
	}
	if !replace.IsArray() {
		zend.ConvertToStringEx(replace)
	}
	if !start.IsArray() {
		zend.ConvertToLong(start)
	}
	if length != nil && !length.IsArray() {
		zend.ConvertToLong(length)
	}
	if zend.EG__().GetException() != nil {
		return
	}

	if str.IsString() {
		res := substrReplaceStr(str.GetStrVal(), replace, start, length)
		returnValue.SetStringVal(res)
		return
	} else {
		res := substrReplaceArray(str.GetArr(), replace, str, length)
		returnValue.SetArray(res)
		return
	}
}

func ZifQuotemeta(str string) (string, bool) {
	if str == "" {
		return "", false
	}

	replacer := strings.NewReplacer(
		`.`, `\.`,
		`\`, `\\`,
		`+`, `\+`,
		`*`, `\*`,
		`?`, `\?`,
		`[`, `\[`,
		`^`, `\^`,
		`]`, `\]`,
		`$`, `\$`,
		`(`, `\(`,
		`)`, `\)`,
	)
	return replacer.Replace(str), true
}

func ZifOrd(character string) int {
	if character == "" {
		return 0
	}
	return int(character[0])
}
func ZifChr(codepoint int) string {
	c := byte(codepoint & 0xff)
	return string(c)
}
func ZifUcfirst(str string) string {
	if str != "" && ascii.IsLower(str[0]) {
		return string(ascii.ToUpper(str[0])) + str[1:]
	}
	return str
}
func ZifLcfirst(str string) string {
	if str != "" && ascii.IsUpper(str[0]) {
		return string(ascii.ToLower(str[0])) + str[1:]
	}
	return str
}
func ZifUcwords(str string, _ zpp.Opt, delimiters *string) string {
	var mask = " \t\r\n\f\v"
	if delimiters != nil {
		mask, _ = PhpCharmaskEx(*delimiters)
	}

	if str == "" {
		return ""
	}

	chars := []byte(str)
	chars[0] = ascii.ToUpper(chars[0])
	for i := 1; i < len(str)-1; i++ {
		if strings.ContainsRune(mask, rune(chars[i-1])) {
			chars[i] = ascii.ToUpper(chars[i])
		}
	}
	return string(chars)
}

func phpStrtrEx(str string, from string, to string) string {
	l := b.Min(len(from), len(to))
	if len(from) == 0 {
		return str
	}

	runeMap := make(map[byte]byte)
	for i := 0; i < l; i++ {
		runeMap[from[i]] = to[i]
	}
	return strutil.MapByte(func(b byte) byte {
		if r, ok := runeMap[b]; ok {
			return r
		}
		return b
	}, str)
}
func phpStrtrArray(str string, pats *types.Array) (string, bool) {
	// 扫描替换数组
	var minLen int = 128 * 1024        // 最长扫描字符串长度
	var maxLen int = 0                 // 最短扫描字符串长度
	numBitset := b.NewBitset(len(str)) // 标记是否有对应长度的扫描字符串
	bitset := ascii.NewAsciiSet()      // 标记是否有对应字符开头的扫描字符串

	var strMap = make(map[string]*types.Zval, pats.Len())
	pats.ForeachIndirect(func(key types.ArrayKey, value *types.Zval) {
		var strKey string
		if !key.IsStrKey() {
			strKey = strconv.Itoa(key.IndexKey())
		} else {
			strKey = key.StrKey()
		}

		/* skip long patterns */
		if len(strKey) > len(str) {
			return
		}

		maxLen = b.Max(maxLen, len(strKey))
		minLen = b.Min(minLen, len(strKey))

		/* remember possible key length */
		numBitset.Mark(len(strKey))
		bitset.Mark(strKey[0])

		strMap[strKey] = value
	})

	// 特殊case, key 为空字符串时，直接返回false (此行为在PHP8中有所不同)
	if _, ok := strMap[""]; ok {
		return "", false
	}

	if len(strMap) == 0 {
		/* return the original string */
		return str, true
	}

	oldPos := 0
	var result strings.Builder
	for pos := 0; pos <= len(str)-minLen; pos++ {
		if !bitset.Contains(str[pos]) {
			continue
		}

		for len_ := b.Min(maxLen, len(str)-pos); len_ >= minLen; len_-- {
			if numBitset.Marked(len_) {
				continue
			}

			search := str[pos : pos+len_]
			replaceZval, ok := strMap[search]
			if !ok {
				continue
			}

			replaceStr := zend.ZvalGetStrVal(replaceZval)
			result.WriteString(str[oldPos:pos])
			result.WriteString(replaceStr)
			oldPos = pos + len_
			pos = oldPos - 1
			break
		}
	}
	if oldPos != 0 {
		result.WriteString(str[oldPos:])
		return result.String(), true
	} else {
		return str, true
	}
}

func PhpCharToStr(str string, from byte, to string, caseSensitivity bool) (string, int) {
	// 预计算替换个数
	var count int
	if caseSensitivity || !ascii.IsAscii(from) {
		count = strings.Count(str, string(from))
	} else {
		count = strings.Count(str, string(ascii.ToUpper(from))) + strings.Count(str, string(ascii.ToLower(from)))
	}
	if count == 0 {
		return str, 0
	}

	// 替换
	var result string
	if caseSensitivity || !ascii.IsAscii(from) {
		result = strings.ReplaceAll(str, string(from), to)
	} else {
		replacer := strings.NewReplacer(
			string(ascii.ToUpper(from)), to,
			string(ascii.ToLower(from)), to,
		)
		result = replacer.Replace(str)
	}

	return result, count
}

func PhpCharToStrEx(str *types.String, from byte, to *byte, toLen int, caseSensitivity int, replaceCount *zend.ZendLong) *types.String {
	result, count := PhpCharToStr(str.GetStr(), from, b.CastStr(to, toLen), caseSensitivity != 0)
	if replaceCount != nil {
		*replaceCount = count
	}
	return types.NewString(result)
}

func PhpStrToStrEx_Ex(haystack string, needle string, str string) (string, int) {
	if len(needle) > len(haystack) {
		return haystack, 0
	} else if needle == haystack {
		return str, 1
	}

	count := strings.Count(haystack, needle)
	result := strings.ReplaceAll(haystack, needle, str)
	return result, count
}

func PhpStrToStrEx(
	haystack *types.String,
	needle *byte,
	needle_len int,
	str *byte,
	str_len int,
	replace_count *zend.ZendLong,
) *types.String {
	result, count := PhpStrToStrEx_Ex(haystack.GetStr(), b.CastStr(needle, needle_len), b.CastStr(str, str_len))
	if replace_count != nil {
		*replace_count += count
	}
	return types.NewString(result)
}
func PhpStrToStrIEx_Ex(haystack string, lcHaystack string, needle string, str string) (string, int) {
	b.Assert(len(needle) != 0)

	if len(needle) > len(haystack) {
		return haystack, 0
	}

	lcNeedle := ascii.StrToLower(needle)
	if lcHaystack == lcNeedle {
		return str, 1
	}

	var buf strings.Builder
	lastPos := 0
	count := 0
	for {
		pos := strings.Index(lcHaystack[lastPos:], needle)
		if pos < 0 {
			break
		}

		buf.WriteString(haystack[lastPos : lastPos+pos])
		buf.WriteString(str)
		lastPos = lastPos + pos + len(needle)
		count++
	}
	if count == 0 {
		return haystack, 0
	} else {
		buf.WriteString(haystack[lastPos:])
		return buf.String(), count
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
	result, count := PhpStrToStrIEx_Ex(haystack.GetStr(), b.CastStrAuto(lc_haystack), needle.GetStr(), b.CastStr(str, str_len))
	if replace_count != nil {
		*replace_count += count
	}
	return types.NewString(result)
}

func ZifStrtr(str string, from *types.Zval, _ zpp.Opt, to_ *string) (string, bool) {
	// 支持两种参数形式:
	// - strtr(string, array)
	// - strtr(string, string, string)
	if to_ != nil && !from.IsArray() {
		core.PhpErrorDocref(nil, faults.E_WARNING, "The second argument is not an array")
		return "", false
	}

	/* shortcut for empty string */
	if str == "" {
		return "", true
	}
	if to_ == nil {
		var pats *types.Array = from.GetArr()
		switch pats.Len() {
		case 0:
			return str, true
		case 1:
			key, val := pats.First()

			var strKey string
			if !key.IsStrKey() {
				strKey = strconv.Itoa(key.IndexKey())
			} else {
				strKey = key.StrKey()
			}

			if strKey == "" {
				return str, true
			}

			replace := zend.ZvalGetStrVal(val)
			return strings.ReplaceAll(str, strKey, replace), true
		default:
			return phpStrtrArray(str, pats)
		}
	} else {
		if zend.TryConvertToString(from) == 0 {
			// unreachable, 触发 fatal error
			return "", false
		}
		return phpStrtrEx(str, from.GetStrVal(), *to_), true
	}
}
