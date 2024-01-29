package standard

import (
	"encoding/hex"
	"fmt"
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/kits/strkit"
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/php/zpp"
	"math"
	"math/rand"
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

const spaceCutset = " \n\r\t\v\x00"

const (
	CHAR_MAX    = 127
	LC_CTYPE    = 2
	LC_NUMERIC  = 4
	LC_TIME     = 5
	LC_COLLATE  = 1
	LC_MONETARY = 3
	LC_ALL      = 0
)

func RegisterStringConstants(ctx *php.Context, moduleNumber int) {
	php.RegisterConstant(ctx, moduleNumber, "STR_PAD_LEFT", php.Long(STR_PAD_LEFT))
	php.RegisterConstant(ctx, moduleNumber, "STR_PAD_RIGHT", php.Long(STR_PAD_RIGHT))
	php.RegisterConstant(ctx, moduleNumber, "STR_PAD_BOTH", php.Long(STR_PAD_BOTH))
	php.RegisterConstant(ctx, moduleNumber, "PATHINFO_DIRNAME", php.Long(PHP_PATHINFO_DIRNAME))
	php.RegisterConstant(ctx, moduleNumber, "PATHINFO_BASENAME", php.Long(PHP_PATHINFO_BASENAME))
	php.RegisterConstant(ctx, moduleNumber, "PATHINFO_EXTENSION", php.Long(PHP_PATHINFO_EXTENSION))
	php.RegisterConstant(ctx, moduleNumber, "PATHINFO_FILENAME", php.Long(PHP_PATHINFO_FILENAME))

	/* If last members of struct lconv equal CHAR_MAX, no grouping is done */

	php.RegisterConstant(ctx, moduleNumber, "CHAR_MAX", php.Long(CHAR_MAX))
	php.RegisterConstant(ctx, moduleNumber, "LC_CTYPE", php.Long(LC_CTYPE))
	php.RegisterConstant(ctx, moduleNumber, "LC_NUMERIC", php.Long(LC_NUMERIC))
	php.RegisterConstant(ctx, moduleNumber, "LC_TIME", php.Long(LC_TIME))
	php.RegisterConstant(ctx, moduleNumber, "LC_COLLATE", php.Long(LC_COLLATE))
	php.RegisterConstant(ctx, moduleNumber, "LC_MONETARY", php.Long(LC_MONETARY))
	php.RegisterConstant(ctx, moduleNumber, "LC_ALL", php.Long(LC_ALL))
}

//func ZifParseStr(ctx *php.Context, encodedString string, _ zpp.Opt, result zpp.RefZval) {
//	if result == nil {
//		if !php.ForbidDynamicCall(ctx, "parse_str() with a single argument") {
//			return
//		}
//		php.ErrorDocRef(ctx, "", perr.E_DEPRECATED, "Calling parse_str() without the result argument is deprecated")
//		var symbolTable = php.ZendRebuildSymbolTable(ctx)
//		var tmp = php.Array(symbolTable.Ht())
//		php.PhpDefaultTreatStringData(ctx, encodedString, tmp)
//		if symbolTable.KeyDelete(types.STR_THIS) {
//			php.ThrowError(ctx, nil, "Cannot re-assign $this")
//		}
//	} else {
//		result = php.ZendTryArrayInit(ctx, result)
//		if result == nil {
//			return
//		}
//		php.PhpDefaultTreatStringData(ctx, encodedString, result)
//	}
//}
//
//func ZifSscanf(ctx *php.Context, str string, format string, vars []zpp.RefZval) *types.Zval {
//	retval, result := SscanfInternal(ctx, str, format, vars)
//	if SCAN_ERROR_WRONG_PARAM_COUNT == result {
//		php.ZendWrongParamCount(ctx)
//	}
//	return retval
//}

func ZifUtf8Encode(data string) string {
	var buf strings.Builder
	for _, c := range []byte(data) {
		if c < 0x80 {
			buf.WriteByte(c)
		} else {
			buf.WriteByte(0xc0 | c>>6)
			buf.WriteByte(0x80 | c&0x3f)
		}
	}
	return buf.String()
}

//func ZifUtf8Decode(data string) string {
//	var pos = 0
//	var buf strings.Builder
//	for pos < len(data) {
//		var status = types.FAILURE
//		c := PhpNextUtf8Char((*uint8)(data), len(data), &pos, &status)
//
//		/* The lower 256 codepoints of Unicode are identical to Latin-1,
//		 * so we don't need to do any mapping here beyond replacing non-Latin-1
//		 * characters. */
//		if status == types.FAILURE || c > 0xff {
//			c = '?'
//		}
//		buf.WriteRune(rune(c))
//	}
//	return buf.String()
//}

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
	replacer := strings.NewReplacer(
		"\000", "\\0",
		`'`, `\'`,
		`"`, `\"`,
		`\`, `\\`,
	)
	result := replacer.Replace(str)
	return result
}

func PhpAddcslashes(ctx *php.Context, str string, what string) string {
	if str == "" {
		return ""
	}
	if what == "" {
		return str
	}

	mask, _ := PhpCharmaskEx(ctx, what)

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

func PhpStripslashes(str string) string {
	if str == "" {
		return ""
	}
	if idx := strings.IndexByte(str, '\\'); idx < 0 {
		return str
	}
	var buf strings.Builder
	for i := 0; i < len(str); i++ {
		if str[i] == '\\' {
			i++ /* skip the slash */
			if i < len(str) {
				if str[i] == '0' {
					buf.WriteByte(0)
				} else {
					buf.WriteByte(str[i])
				}
			}
		} else {
			buf.WriteByte(str[i])
		}
	}
	return buf.String()
}

func PhpStripcslashes(str string) string {
	if pos := strings.IndexByte(str, '\\'); pos < 0 {
		return str
	}

	var buf strings.Builder
	for i := 0; i < len(str); i++ {
		if str[i] != '\\' || i+1 >= len(str) {
			buf.WriteByte(str[i])
			continue
		}

		// 处理转义字符
		i++
		switch str[i] {
		case 'n':
			buf.WriteByte('\n')
		case 'r':
			buf.WriteByte('\r')
		case 'a':
			buf.WriteByte('\a')
		case 't':
			buf.WriteByte('\t')
		case 'v':
			buf.WriteByte('\v')
		case 'b':
			buf.WriteByte('\b')
		case 'f':
			buf.WriteByte('\f')
		case '\\':
			buf.WriteByte('\\')
		case 'x':
			// try \x[0-9a-fA-F]{1,2}
			hexSize := 0
			for hexSize < 2 && i+hexSize+1 < len(str) && ascii.IsXDigit(str[i+hexSize+1]) {
				hexSize++
			}
			if hexSize > 0 {
				hexNum, _ := strconv.ParseInt(str[i+1:i+hexSize+1], 16, 0)
				buf.WriteByte(byte(hexNum))
				i += hexSize
				break
			}

			// fallback
			buf.WriteString("\\x")
		default:
			// try \[0-7]{1,3}
			octSize := 0
			for octSize < 3 && i+octSize < len(str) && '0' <= str[i+octSize] && str[i+octSize] <= '7' {
				octSize++
			}
			if octSize > 0 {
				hexNum, _ := strconv.ParseInt(str[i:i+octSize], 8, 0)
				buf.WriteByte(byte(hexNum))
				i += octSize - 1
				break
			}

			// fallback
			buf.WriteByte(str[i])
		}
	}
	return buf.String()
}

/**
 * Zif functions
 */

func ZifBin2hex(data string) string {
	return hex.EncodeToString([]byte(data))
}
func ZifHex2bin(ctx *php.Context, data string) (string, bool) {
	if len(data)%2 != 0 {
		php.ErrorDocRef(ctx, "", perr.E_WARNING, "Hexadecimal input string must have an even length")
		return "", false
	}

	bin, err := hex.DecodeString(data)
	if err != nil {
		php.ErrorDocRef(ctx, "", perr.E_WARNING, "Input string must be hexadecimal string")
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
func PhpCharmaskEx(ctx *php.Context, input string) (string, bool) {
	return charmaskEx(input, func(err string) {
		php.ErrorDocRef(ctx, "", perr.E_WARNING, err)
	})
}
func charmaskEx(input string, onError func(string)) (string, bool) {
	if pos := strings.Index(input, ".."); pos < 0 {
		return input, true
	}

	var result = true
	var buf strings.Builder
	var length = len(input)
	for i := 0; i < length; i++ {
		if i+3 < length && input[i+1] == '.' && input[i+2] == '.' && input[i+3] >= input[i] {
			for c := input[i]; c <= input[i+3]; c++ {
				buf.WriteByte(c)
			}
			i += 3
		} else if i+1 < length && input[i] == '.' && input[i+1] == '.' {
			/* Error, try to be as helpful as possible:
			   (a range ending/starting with '.' won't be captured here) */
			if i == 0 {
				result = false
				onError("Invalid '..'-range, no character to the left of '..'")
				continue
			}
			if i+2 >= len(input) {
				result = false
				onError("Invalid '..'-range, no character to the right of '..'")
				continue
			}
			if input[i-1] > input[i+2] {
				result = false
				onError("Invalid '..'-range, '..'-range needs to be incrementing")
				continue
			}

			/* FIXME: better error (a..b..c is the only left possibility?) */
			result = false
			onError("Invalid '..'-range")
			continue
		} else {
			buf.WriteByte(input[i])
		}
	}
	return buf.String(), result
}

func ZifTrim(ctx *php.Context, ex *php.ExecuteData, str string, _ zpp.Opt, characterMask string) string {
	var cutset = spaceCutset
	if ex.NumArgs() >= 2 {
		cutset, _ = PhpCharmaskEx(ctx, characterMask)
	}
	return strings.Trim(str, cutset)
}

//@zif(alias="chop")
func ZifRtrim(ctx *php.Context, ex *php.ExecuteData, str string, _ zpp.Opt, characterMask string) string {
	var cutset = spaceCutset
	if ex.NumArgs() >= 2 {
		cutset, _ = PhpCharmaskEx(ctx, characterMask)
	}
	return strings.TrimRight(str, cutset)
}
func ZifLtrim(ctx *php.Context, ex *php.ExecuteData, str string, _ zpp.Opt, characterMask string) string {
	var cutset = spaceCutset
	if ex.NumArgs() >= 2 {
		cutset, _ = PhpCharmaskEx(ctx, characterMask)
	}
	return strings.TrimLeft(str, cutset)
}

func ZifWordwrap(ctx *php.Context, str string, _ zpp.Opt, width *int, break_ *string, cut bool) (string, bool) {
	var linelength = lang.Option(width, 75)
	var breakchar = lang.Option(break_, "\n")

	if str == "" {
		return "", true
	}
	if breakchar == "" {
		php.ErrorDocRef(ctx, "", perr.E_WARNING, "Break string cannot be empty")
		return "", false
	}
	if linelength == 0 && cut {
		php.ErrorDocRef(ctx, "", perr.E_WARNING, "Can't force cut when width is zero")
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
func ZifExplode(ctx *php.Context, separator string, str string, _ zpp.Opt, limit *int) ([]string, bool) {
	if len(separator) == 0 {
		php.ErrorDocRef(ctx, "", perr.E_WARNING, "Empty delimiter")
		return nil, false
	}

	var arr []string
	if limit == nil {
		arr = strings.Split(str, separator)
	} else if *limit >= 0 {
		// doc: If the limit parameter is zero, then this is treated as 1.
		limitVal := lang.Max(*limit, 1)
		limitVal = lang.Min(limitVal, strings.Count(str, separator))
		arr = strings.SplitN(str, separator, limitVal)
	} else {
		limitVal := *limit // limitVal < 0
		arr = strings.Split(str, separator)
		if len(arr) > -limitVal {
			arr = arr[:len(arr)+limitVal]
		} else {
			arr = nil
		}
	}
	return arr, true
}

//@zif(alias="join")
func ZifImplode(ctx *php.Context, glue_ *types.Zval, _ zpp.Opt, pieces_ *types.Zval) types.Zval {
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
			php.ErrorDocRef(ctx, "", perr.E_WARNING, "Argument must be an array")
			return types.Null
		}
		glue = ""
		pieces = arg1.Array()
	} else {
		if arg1.IsArray() {
			glue = php.ZvalGetStrVal(ctx, *arg2)
			pieces = arg1.Array()
			php.ErrorDocRef(ctx, "", perr.E_DEPRECATED, "Passing glue string after array is deprecated. Swap the parameters")
		} else if arg2.IsArray() {
			glue = php.ZvalGetStrVal(ctx, *arg1)
			pieces = arg2.Array()
		} else {
			php.ErrorDocRef(ctx, "", perr.E_WARNING, "Invalid arguments passed")
			return types.Null
		}
	}
	result := PhpImplode(ctx, glue, pieces)
	return php.String(result)
}
func PhpImplode(ctx *php.Context, glue string, pieces *types.Array) string {
	var parts []string
	pieces.Each(func(_ types.ArrayKey, value types.Zval) {
		parts = append(parts, php.ZvalGetStrVal(ctx, value))
	})
	return strings.Join(parts, glue)
}

func ZifStrtok(ctx *php.Context, str string, _ zpp.Opt, token_ *string) (string, bool) {
	var state = BG(ctx).StrTokState()

	// 两种参数形式
	// - strtok(string $string, string $token): string|false
	// - strtok(string $token): string|false
	var token string
	if token_ != nil {
		*state = StrTokState{
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
	if suffix != "" && len(suffix) < len(s) && strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}
func ZifBasename(path string, _ zpp.Opt, suffix string) string {
	return PhpBasename(path, suffix)
}
func ZifDirname(ctx *php.Context, path string, _ zpp.Opt, levels_ *int) types.Zval {
	var levels = 1
	if levels_ != nil {
		levels = *levels_
		if levels < 1 {
			php.ErrorDocRef(ctx, "", perr.E_WARNING, "Invalid argument, levels must be >= 1")
			return types.Null
		}
	}

	dir := path
	for i := 0; i < levels; i++ {
		newDir := php.ZendDirname(dir)
		if newDir == dir {
			break
		}
		dir = newDir
	}
	return types.ZvalString(dir)
}

func ZifPathinfo(path string, _ zpp.Opt, options *int) types.Zval {
	opt := lang.Option(options, PHP_PATHINFO_ALL)

	arr := types.NewArray()
	if (opt & PHP_PATHINFO_DIRNAME) == PHP_PATHINFO_DIRNAME {
		dirname := php.ZendDirname(path)
		if dirname != "" {
			arr.KeyUpdate("dirname", types.ZvalString(dirname))
		}
	}

	basename := PhpBasename(path, "")
	if (opt & PHP_PATHINFO_BASENAME) == PHP_PATHINFO_BASENAME {
		arr.KeyUpdate("basename", types.ZvalString(basename))
	}
	if (opt & PHP_PATHINFO_EXTENSION) == PHP_PATHINFO_EXTENSION {
		if pos := strings.LastIndexByte(basename, '.'); pos >= 0 {
			arr.KeyUpdate("extension", types.ZvalString(basename[pos+1:]))
		}
	}
	if (opt & PHP_PATHINFO_FILENAME) == PHP_PATHINFO_FILENAME {
		if pos := strings.LastIndexByte(basename, '.'); pos >= 0 {
			arr.KeyUpdate("filename", types.ZvalString(basename[:pos]))
		} else {
			arr.KeyUpdate("filename", types.ZvalString(basename))
		}
	}

	if opt == PHP_PATHINFO_ALL {
		return types.ZvalArray(arr)
	} else {
		p := arr.First()
		if p.IsValid() {
			return p.Val
		} else {
			return types.ZvalString("")
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

func strLenCheckNull(s string) int {
	if idx := strings.IndexByte(s, 0); idx >= 0 {
		return idx
	} else {
		return len(s)
	}
}

func PhpStrcspnEx(s1 string, s2 string) int {
	if s1 == "" {
		return 0
	}
	if s2 == "" {
		return strLenCheckNull(s1)
	}
	for i, c := range []byte(s1) {
		if strings.ContainsRune(s2, rune(c)) {
			return i
		}
	}
	return len(s1)
}
func PhpNeedleChar(ctx *php.Context, needle types.Zval) (byte, bool) {
	switch needle.Type() {
	case types.IsLong:
		return byte(needle.Long()), true
	case types.IsNull, types.IsFalse:
		return 0, true
	case types.IsTrue:
		return 1, true
	case types.IsDouble, types.IsObject:
		return byte(php.ZvalGetLong(ctx, needle)), true
	default:
		php.ErrorDocRef(ctx, "", perr.E_WARNING, "needle is not a string or an integer")
		return 0, false
	}
}
func ZifStristr(ctx *php.Context, haystack string, needle types.Zval, _ zpp.Opt, part bool) (string, bool) {
	needleStr, ok := parseNeedle(ctx, needle)
	if !ok {
		return "", false
	}
	if needleStr == "" {
		php.ErrorDocRef(ctx, "", perr.E_WARNING, "Empty needle")
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

//@zif(alias="strchr")
func ZifStrstr(ctx *php.Context, haystack string, needle types.Zval, _ zpp.Opt, part bool) (string, bool) {
	needleStr, ok := parseNeedle(ctx, needle)
	if !ok {
		return "", false
	}
	if needleStr == "" {
		php.ErrorDocRef(ctx, "", perr.E_WARNING, "Empty needle")
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

func offsetSubstr(ctx *php.Context, str string, offset int) (result string, realOffset int, ok bool) {
	if offset < 0 {
		offset += len(str)
	}
	if offset < 0 || offset > len(str) {
		php.ErrorDocRef(ctx, "", perr.E_WARNING, "Offset not contained in string")
		return "", 0, false
	}
	return str[offset:], offset, true
}

func parseNeedle(ctx *php.Context, needle types.Zval) (string, bool) {
	if needle.IsString() {
		return needle.String(), true
	} else {
		/*
		 * 在 PHP 8.0.0 之前，如果 needle 不是字符串，它将被转换为整数并作为字符的序数值应用。
		 * 从 PHP 7.3.0 开始，这种行为已被废弃，不鼓励依赖它。根据预期的行为，应该明确地将 needle 转换成字符串，或者明确地调用 chr()。
		 */
		needleChar, ok := PhpNeedleChar(ctx, needle)
		if !ok {
			return "", false
		}

		php.ErrorDocRef(ctx, "", perr.E_DEPRECATED, "Non-string needles will be interpreted as strings in the future. Use an explicit chr() call to preserve the current behavior")

		return string([]byte{needleChar}), true
	}
}

func ZifStrpos(ctx *php.Context, haystack string, needle types.Zval, _ zpp.Opt, offset int) (int, bool) {
	haystack, offset, ok := offsetSubstr(ctx, haystack, offset)
	if !ok {
		return 0, false
	}
	if len(haystack) == 0 {
		return 0, false
	}

	needleStr, ok := parseNeedle(ctx, needle)
	if !ok {
		return 0, false
	}

	if pos := strings.Index(haystack, needleStr); pos >= 0 {
		return offset + pos, true
	} else {
		return 0, false
	}
}
func ZifStripos(ctx *php.Context, haystack string, needle types.Zval, _ zpp.Opt, offset int) (int, bool) {
	haystack, offset, ok := offsetSubstr(ctx, haystack, offset)
	if !ok {
		return 0, false
	}
	if len(haystack) == 0 {
		return 0, false
	}

	needleStr, ok := parseNeedle(ctx, needle)
	if !ok || needleStr == "" {
		return 0, false
	}

	haystack = ascii.StrToLower(haystack)
	needleStr = ascii.StrToLower(needleStr)
	if pos := strings.Index(haystack, needleStr); pos >= 0 {
		return offset + pos, true
	} else {
		return 0, false
	}
}

//@zif(onError=1)
func ZifStrrpos(ctx *php.Context, haystack string, needle types.Zval, _ zpp.Opt, offset int) (int, bool) {
	needleStr, ok := parseNeedle(ctx, needle)
	if !ok {
		return 0, false
	}

	if len(haystack) == 0 {
		return 0, false
	}

	if offset >= 0 {
		haystack, offset, ok = offsetSubstr(ctx, haystack, offset)
		if !ok {
			return 0, false
		}
		if pos := strings.LastIndex(haystack, needleStr); pos >= 0 {
			return offset + pos, true
		} else {
			return 0, false
		}
	} else { // offset < 0
		offset += len(haystack)
		if offset < 0 {
			php.ErrorDocRef(ctx, "", perr.E_WARNING, "Offset is greater than the length of haystack string")
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

//@zif(onError=1)
func ZifStrripos(ctx *php.Context, haystack string, needle types.Zval, _ zpp.Opt, offset int) (int, bool) {
	needleStr, ok := parseNeedle(ctx, needle)
	if !ok {
		return 0, false
	}

	if len(haystack) == 0 {
		return 0, false
	}

	haystack = ascii.StrToLower(haystack)
	needleStr = ascii.StrToLower(needleStr)
	if offset >= 0 {
		var ok bool
		haystack, offset, ok = offsetSubstr(ctx, haystack, offset)
		if !ok {
			return 0, false
		}
		if pos := strings.Index(haystack, needleStr); pos >= 0 {
			return offset + pos, true
		} else {
			return 0, false
		}
	} else { // offset < 0
		offset += len(haystack)
		if offset < 0 {
			php.ErrorDocRef(ctx, "", perr.E_WARNING, "Offset is greater than the length of haystack string")
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
func ZifStrrchr(ctx *php.Context, haystack string, needle types.Zval) (string, bool) {
	needleStr, ok := parseNeedle(ctx, needle)
	if !ok {
		return "", false
	}
	var needleChar byte
	if needleStr != "" {
		needleChar = needleStr[0]
	} else {
		needleChar = 0
	}
	if pos := strings.LastIndexByte(haystack, needleChar); pos >= 0 {
		return haystack[pos:], true
	} else {
		return "", false
	}
}
func ZifChunkSplit(ctx *php.Context, str string, _ zpp.Opt, chunklen_ *int, ending_ *string) (string, bool) {
	chunklen := lang.Option(chunklen_, 76)
	ending := lang.Option(ending_, "\r\n")

	if chunklen <= 0 {
		php.ErrorDocRef(ctx, "", perr.E_WARNING, "Chunk length should be greater than zero")
		return "", false
	}

	// fast
	if ending == "" {
		return str, true
	}
	if chunklen >= len(str) {
		return str + ending, true
	}

	// 确保结果长度不溢出
	strLen, endingLen := len(str), len(ending)
	chunks := strLen / chunklen
	if endingLen > (types.MaxStrLen-strLen)/chunks {
		return "", false
	}
	outLen := strLen + chunks*endingLen

	// common
	var buf strings.Builder
	buf.Grow(outLen)
	for i := 0; i < strLen; i += chunklen {
		if i+chunklen <= strLen {
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
	php.Assert(0 <= f && f <= php.LongMax)
	php.Assert(0 <= l && l <= php.LongMax)
	if f+l > len(str) {
		l = len(str) - f
	}

	return str[:f] + replace + str[f+l:]
}

func substrReplaceStr(ctx *php.Context, str string, replace types.Zval, start types.Zval, length *types.Zval) string {
	// str 为字符串时，允许的参数类型:
	// - substr_replace(string, array|string, int, int|null)
	// 其他情况都会触发 warning 并返回原字符串
	if start.IsArray() {
		if length == nil || !length.IsArray() {
			php.ErrorDocRef(ctx, "", perr.E_WARNING, "'start' and 'length' should be of same type - numerical or array ")
		} else if length.IsArray() && start.Array().Len() != length.Array().Len() {
			php.ErrorDocRef(ctx, "", perr.E_WARNING, "'start' and 'length' should have the same number of elements")
		} else {
			php.ErrorDocRef(ctx, "", perr.E_WARNING, "Functionality of 'start' and 'length' as arrays is not implemented")
		}
		return str
	}
	if length != nil && length.IsArray() {
		php.ErrorDocRef(ctx, "", perr.E_WARNING, "'start' and 'length' should be of same type - numerical or array ")
		return str
	}

	// 正常处理 substr_replace(string, array|string, int, int|null)
	l := len(str)
	if length != nil && !length.IsArray() {
		l = length.Long()
	}

	//
	var replStr string
	if replace.IsArray() {
		replPair := replace.Array().First()
		if replPair.IsValid() {
			replStr = php.ZvalGetStrVal(ctx, replPair.Val)
		} else {
			replStr = ""
		}
	} else {
		replStr = replace.String()
	}

	res := substrReplaceSingle(str, replStr, start.Long(), l)
	return res
}

func substrReplaceArray(ctx *php.Context, str *types.Array, replace types.Zval, start types.Zval, length *types.Zval) *types.Array {
	arr := types.NewArrayCap(str.Len())

	var replaceStr []string
	if replace.IsArray() {
		replace.Array().Each(func(_ types.ArrayKey, value types.Zval) {
			replaceStr = append(replaceStr, php.ZvalGetStrVal(ctx, value))
		})
	}

	var startPoints []int
	if start.IsArray() {
		start.Array().Each(func(_ types.ArrayKey, value types.Zval) {
			startPoints = append(startPoints, php.ZvalGetLong(ctx, value))
		})
	}

	var lengthPoints []int
	if length != nil && length.IsArray() {
		length.Array().Each(func(_ types.ArrayKey, value types.Zval) {
			lengthPoints = append(lengthPoints, php.ZvalGetLong(ctx, value))
		})
	}

	idx := -1
	str.Each(func(key types.ArrayKey, value types.Zval) {
		idx++

		origStr := php.ZvalGetStrVal(ctx, value)

		// f
		var f = 0
		if start.IsArray() {
			if idx < len(startPoints) {
				f = startPoints[idx]
			}
		} else {
			f = start.Long()
		}

		// l
		var l = len(origStr)
		if length != nil {
			if length.IsArray() {
				if idx < len(lengthPoints) {
					l = lengthPoints[idx]
				}
			} else {
				l = length.Long()
			}
		}

		// repl
		var replStr = ""
		if replace.IsArray() {
			if idx <= len(replaceStr) {
				replStr = replaceStr[idx]
			}
		} else {
			replStr = replace.String()
		}

		ret := substrReplaceSingle(origStr, replStr, f, l)
		if key.IsStrKey() {
			arr.SymtableUpdate(key.StrKey(), types.ZvalString(ret))
		} else {
			arr.IndexUpdate(key.IdxKey(), types.ZvalString(ret))
		}
	})
	return arr
}

func ZifSubstrReplace(ctx *php.Context, returnValue zpp.Ret, str types.Zval, replace types.Zval, start types.Zval, _ zpp.Opt, length *types.Zval) {
	// 限定参数类型
	// - substr_replace(array|string $str, array|string $replace, array|int $start, array|int|null $length = null)
	if !str.IsArray() {
		php.ConvertToString(ctx, &str)
	}
	if !replace.IsArray() {
		php.ConvertToString(ctx, &replace)
	}
	if !start.IsArray() {
		start.SetLong(php.ZvalGetLong(ctx, start))
	}
	if length != nil && !length.IsArray() {
		length.SetLong(php.ZvalGetLong(ctx, *length))
	}
	if ctx.EG().HasException() {
		return
	}

	if str.IsString() {
		res := substrReplaceStr(ctx, str.String(), replace, start, length)
		returnValue.SetString(res)
		return
	} else {
		res := substrReplaceArray(ctx, str.Array(), replace, str, length)
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

//@zif(onError=2)
func ZifChr(codepoint int) string {
	c := byte(codepoint & 0xff)
	return string([]byte{c})
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
func ZifUcwords(ctx *php.Context, str string, _ zpp.Opt, delimiters *string) string {
	var mask = " \t\r\n\f\v"
	if delimiters != nil {
		mask, _ = PhpCharmaskEx(ctx, *delimiters)
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

func Strtr(str string, from string, to string) string {
	l := lang.Min(len(from), len(to))
	if l == 0 {
		return str
	}

	byteMap := make(map[byte]byte)
	for i := 0; i < l; i++ {
		byteMap[from[i]] = to[i]
	}
	return strkit.MapByte(func(b byte) byte {
		if r, ok := byteMap[b]; ok {
			return r
		}
		return b
	}, str)
}

func phpStrtrArray(ctx *php.Context, str string, pats *types.Array) (string, bool) {
	// 扫描替换数组
	var minLen = 128 * 1024               // 最长扫描字符串长度
	var maxLen = 0                        // 最短扫描字符串长度
	lenExists := make(map[int]struct{})   // 标记是否有对应长度的扫描字符串
	headExists := make(map[byte]struct{}) // 标记是否有对应字符开头的扫描字符串

	var strMap = make(map[string]types.Zval, pats.Len())
	pats.Each(func(key types.ArrayKey, value types.Zval) {
		var strKey string
		if !key.IsStrKey() {
			strKey = strconv.Itoa(key.IdxKey())
		} else {
			strKey = key.StrKey()
		}

		/* skip empty patterns */
		if strKey == "" {
			return
		}

		/* skip long patterns */
		if len(strKey) > len(str) {
			return
		}

		maxLen = lang.Max(maxLen, len(strKey))
		minLen = lang.Min(minLen, len(strKey))

		/* remember possible key length */
		lenExists[len(strKey)] = struct{}{}
		headExists[strKey[0]] = struct{}{}

		strMap[strKey] = value
	})

	if len(strMap) == 0 {
		/* return the original string */
		return str, true
	}

	oldPos := 0
	var result strings.Builder
	for pos := 0; pos <= len(str)-minLen; pos++ {
		if _, exists := headExists[str[pos]]; !exists {
			continue
		}

		for l := lang.Min(maxLen, len(str)-pos); l >= minLen; l-- {
			if _, exists := lenExists[l]; !exists {
				continue
			}

			search := str[pos : pos+l]
			replaceZval, ok := strMap[search]
			if !ok {
				continue
			}

			replaceStr := php.ZvalGetStrVal(ctx, replaceZval)
			result.WriteString(str[oldPos:pos])
			result.WriteString(replaceStr)
			oldPos = pos + l
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

// @see php_str_to_str()
func stringReplace(s string, search string, replace string) (string, int) {
	if s == "" || search == "" || len(search) > len(s) {
		return s, 0
	}

	count := strings.Count(s, search)
	if count == 0 {
		return s, 0
	}

	result := strings.ReplaceAll(s, search, replace)
	return result, count
}

// @see php_str_to_str_i_ex
func stringReplaceIgnoreCase(s string, search string, replace string) (string, int) {
	if s == "" || search == "" || len(search) > len(s) {
		return s, 0
	}

	lcStr := ascii.StrToLower(s)
	lcSearch := ascii.StrToLower(search)
	if lcStr == lcSearch {
		return replace, 1
	}

	count := strings.Count(lcStr, lcSearch)
	if count == 0 {
		return s, 0
	}

	var buf strings.Builder
	buf.Grow(len(s) + count*(len(replace)-len(search)))
	start := 0
	for i := 0; i < count; i++ {
		j := start + strings.Index(lcStr[start:], lcSearch) // 已确定有 count 个 lcSearch, 此时 strings.Index() 结果肯定 >= 0
		buf.WriteString(s[start:j])
		buf.WriteString(replace)
		start = j + len(search)
	}
	buf.WriteString(s[start:])
	return buf.String(), count
}

func ZifStrtr(ctx *php.Context, str string, from types.Zval, _ zpp.Opt, to_ *string) (string, bool) {
	// 支持两种参数形式:
	// - strtr(string, array)
	// - strtr(string, string, string)
	if to_ == nil && !from.IsArray() {
		php.ErrorDocRef(ctx, "", perr.E_WARNING, "The second argument is not an array")
		return "", false
	}

	/* shortcut for empty string */
	if str == "" {
		return "", true
	}
	if to_ == nil {
		var pats = from.Array()
		switch pats.Len() {
		case 0:
			return str, true
		case 1:
			p := pats.First()
			key, val := p.Key, p.Val

			var strKey string
			if !key.IsStrKey() {
				strKey = strconv.Itoa(key.IdxKey())
			} else {
				strKey = key.StrKey()
			}

			if strKey == "" {
				return str, true
			}

			replace := php.ZvalGetStrVal(ctx, val)
			return strings.ReplaceAll(str, strKey, replace), true
		default:
			return phpStrtrArray(ctx, str, pats)
		}
	} else {
		if s, ok := php.ZvalTryGetStr(ctx, from); ok {
			from.SetString(s)
		} else {
			// unreachable, 触发 fatal error
			return "", false
		}
		return Strtr(str, from.String(), *to_), true
	}
}

func ZifStrrev(str string) string {
	l := len(str)
	result := make([]byte, l)
	for i := 0; i < l; i++ {
		result[i] = str[l-i-1]
	}
	return string(result)
}

func phpSimilarStr(txt1 string, txt2 string) (max int, count int, pos1 int, pos2 int) {
	for i := range []byte(txt1) {
		for j := range []byte(txt2) {
			l := 0
			for l < len(txt1) && l < len(txt2) && txt1[l] == txt2[l] {
				l++
			}
			if l > max {
				max = l
				count++
				pos1 = i
				pos2 = j
			}
		}
	}
	return
}

func phpSimilarChar(txt1 string, txt2 string) int {
	max, count, pos1, pos2 := phpSimilarStr(txt1, txt2)
	sum := max
	if max != 0 {
		if pos1 != 0 && pos2 != 0 && count > 1 {
			sum += phpSimilarChar(txt1[:pos1], txt2[:pos2])
		}
		if pos1+max < len(txt1) && pos2+max < len(txt2) {
			sum += phpSimilarChar(txt1[pos1+max:], txt2[pos2+max:])
		}
	}
	return sum
}

// 计算两个字符串的相似度, O(n^3) 算法为: <Programming Classics: Implementing the World's Best Algorithms by Oliver> (ISBN 0-131-00413-1)
func ZifSimilarText(ctx *php.Context, str1 string, str2 string, _ zpp.Opt, percent zpp.RefZval) int {
	// 特例
	if str1 == "" && str2 == "" {
		if percent != nil {
			php.ZendTryAssignRefDouble(ctx, percent, 0)
		}
		return 0
	}

	// 常规情况
	sim := phpSimilarChar(str1, str2)
	if percent != nil {
		simFloat := float64(sim) * 200 / float64(len(str1)+len(str2))
		php.ZendTryAssignRefDouble(ctx, percent, simFloat)
	}
	return sim
}

func ZifAddslashes(str string) string { return PhpAddslashes(str) }
func ZifAddcslashes(ctx *php.Context, str string, charlist string) string {
	return PhpAddcslashes(ctx, str, charlist)
}
func ZifStripslashes(str string) string  { return PhpStripslashes(str) }
func ZifStripcslashes(str string) string { return PhpStripcslashes(str) }

func strReplaceStr(ctx *php.Context, subject string, search types.Zval, replace types.Zval, caseSensitivity bool) (string, int) {
	if subject == "" {
		return "", 0
	}

	if search.IsArray() {
		var replaceStrings []string
		var replaceStr string
		if replace.IsArray() {
			replaceStrings = make([]string, 0, replace.Array().Len())
			replace.Array().Each(func(key types.ArrayKey, value types.Zval) {
				replaceStrings = append(replaceStrings, php.ZvalGetStrVal(ctx, value))
			})
		} else {
			replaceStr = php.ZvalGetStrVal(ctx, replace)
		}

		var result = subject
		var replaceCount = 0
		i := -1
		search.Array().Each(func(key types.ArrayKey, val types.Zval) {
			if subject == "" {
				return
			}

			searchStr := php.ZvalGetStrVal(ctx, val)
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
					tmpResult, count = stringReplace(result, searchStr, replaceStr)
				} else {
					tmpResult, count = stringReplaceIgnoreCase(result, searchStr, replaceStr)
				}
			}

			result = tmpResult
			replaceCount += count
		})
		return result, replaceCount
	} else {
		php.Assert(search.IsString())
		searchStr := search.String()
		if searchStr == "" {
			return subject, 0
		}
		replaceStr := replace.String()

		if len(searchStr) == 1 {
			return PhpCharToStr(subject, searchStr[0], replaceStr, caseSensitivity)
		} else {
			if caseSensitivity {
				return stringReplace(subject, searchStr, replaceStr)
			} else {
				return stringReplaceIgnoreCase(subject, searchStr, replaceStr)
			}
		}
	}
}
func strReplaceArray(ctx *php.Context, subject *types.Array, search types.Zval, replace types.Zval, caseSensitivity bool) (*types.Array, int) {
	arr := types.NewArrayCap(subject.Len())
	replaceCount := 0
	subject.Each(func(key types.ArrayKey, value types.Zval) {
		value = value.DeRef()

		var result types.Zval
		if !value.IsArray() && !value.IsObject() {
			tmpResult, count := strReplaceStr(ctx, php.ZvalGetStrVal(ctx, value), search, replace, caseSensitivity)
			result.SetString(tmpResult)
			replaceCount += count
		} else {
			result = value
		}

		arr.Add(key, result)
	})
	return arr, replaceCount
}
func strReplace(ctx *php.Context, returnValue *types.Zval, search types.Zval, replace types.Zval, subject types.Zval, replaceCount zpp.RefZval, caseSensitivity bool) {
	// 限定参数类型
	// - str_replace(array|string $search, array|string $replace, array|string $subject, int|null &$count == null): string|array
	if !search.IsArray() {
		php.ConvertToString(ctx, &search)
		php.ConvertToString(ctx, &replace)
	} else if !replace.IsArray() {
		php.ConvertToString(ctx, &replace)
	}
	if ctx.EG().HasException() {
		return
	}

	// 主逻辑
	var count int
	if subject.IsArray() {
		var arr *types.Array
		arr, count = strReplaceArray(ctx, subject.Array(), search, replace, caseSensitivity)
		returnValue.SetArray(arr)
	} else {
		var str string
		str, count = strReplaceStr(ctx, php.ZvalGetStrVal(ctx, subject), search, replace, caseSensitivity)
		returnValue.SetString(str)
	}

	if replaceCount != nil {
		php.ZendTryAssignRefLong(ctx, replaceCount, count)
	}
}

func ZifStrReplace(ctx *php.Context, returnValue zpp.Ret, search types.Zval, replace types.Zval, subject types.Zval, _ zpp.Opt, replaceCount zpp.RefZval) {
	strReplace(ctx, returnValue, search, replace, subject, replaceCount, true)
}
func ZifStrIreplace(ctx *php.Context, returnValue zpp.Ret, search types.Zval, replace types.Zval, subject types.Zval, _ zpp.Opt, replaceCount zpp.RefZval) {
	strReplace(ctx, returnValue, search, replace, subject, replaceCount, false)
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
			var newCharCount = char_count
			var newBegin = begin
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
	var isXhtml = lang.Option(isXhtml_, true)

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

	var normalize strings.Builder
	state := 0
	for i, c := range []byte(tag) {
		c = ascii.ToLower(c)
		if c == '<' {
			normalize.WriteByte(c)
		} else if c == '>' {
			break
		} else {
			if !ascii.IsSpace(c) {
				if state == 0 {
					state = 1
				}
				if c != '/' || ((i == 0 || tag[i-1] != '<') && (i == len(tag)-1 || tag[i+1] != '>')) {
					normalize.WriteByte(c)
				}
			} else {
				if state == 1 {
					break
				}
			}
		}
	}
	normalize.WriteByte('>')

	return strings.Contains(set, normalize.String())
}

func StripTagsEx(str string, state uint8, allowTags string) (string, uint8) {
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
func ZifStripTags(ctx *php.Context, str string, _ zpp.Opt, allowableTags *types.Zval) string {
	var allow = allowableTags
	var allowTagsStr string
	if allow != nil {
		if allow.IsArray() {
			var buf strings.Builder
			allow.Array().Each(func(key types.ArrayKey, value types.Zval) {
				tag := php.ZvalGetStrVal(ctx, value)
				buf.WriteByte('<')
				buf.WriteString(tag)
				buf.WriteByte('>')
			})
			allowTagsStr = buf.String()
		} else {
			/* To maintain a certain BC, we allow anything for the second parameter and return original string */
			allow.SetString(php.ZvalTryGetStrVal(ctx, *allow))
			allowTagsStr = allowableTags.String()
		}
	}

	result, _ := StripTagsEx(str, 0, allowTagsStr)
	return result
}

func ZifStrRepeat(ctx *php.Context, input string, mult int) (string, bool) {
	if mult < 0 {
		php.ErrorDocRef(ctx, "", perr.E_WARNING, "Second argument has to be greater than or equal to 0")
		return "", false
	}
	/* Don't waste our time if it's empty */
	if input == "" || mult == 0 {
		return "", true
	}

	return strings.Repeat(input, mult), true
}
func ZifCountChars(ctx *php.Context, input string, _ zpp.Opt, mode int) (*types.Zval, bool) {
	if mode < 0 || mode > 4 {
		php.ErrorDocRef(ctx, "", perr.E_WARNING, "Unknown mode")
		return nil, false
	}

	var charCount [256]int

	for _, c := range []byte(input) {
		charCount[c]++
	}

	if mode < 3 { // mode=0,1,2 以数组返回
		arr := types.NewArray()
		for i := 0; i < 256; i++ {
			count := charCount[i]
			if mode == 0 || (mode == 1 && count != 0) || (mode == 2 && count == 0) {
				arr.IndexUpdate(i, types.ZvalLong(charCount[i]))
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
func ZifSubstrCount(ctx *php.Context, haystack string, needle string, _ zpp.Opt, offset int, length_ *int) (int, bool) {
	// check needle
	if needle == "" {
		php.ErrorDocRef(ctx, "", perr.E_WARNING, "Empty substring")
		return 0, false
	}

	// check offset
	if offset < 0 {
		offset += len(haystack)
	}
	if offset < 0 || offset > len(haystack) {
		php.ErrorDocRef(ctx, "", perr.E_WARNING, "Offset not contained in string")
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
			php.ErrorDocRef(ctx, "", perr.E_WARNING, "Invalid length value")
			return 0, false
		}
	}

	// 截取目标字符串范围
	searchStr := haystack[offset : offset+length]

	return strings.Count(searchStr, needle), true
}

func ZifStrPad(ctx *php.Context, input string, padLength int, _ zpp.Opt, padString_ *string, padType_ *int) (string, bool) {
	padString := lang.Option(padString_, " ")
	padType := lang.Option(padType_, STR_PAD_RIGHT)

	/* If resulting string turns out to be shorter than input string,
	   we simply copy the input and return. */
	if padLength < 0 || padLength < len(input) {
		return input, true
	}
	if padString == "" {
		php.ErrorDocRef(ctx, "", perr.E_WARNING, "Padding string cannot be empty")
		return "", false
	}
	if padType < STR_PAD_LEFT || padType > STR_PAD_BOTH {
		php.ErrorDocRef(ctx, "", perr.E_WARNING, "Padding type has to be STR_PAD_LEFT, STR_PAD_RIGHT, or STR_PAD_BOTH")
		return "", false
	}
	numPadChars := padLength - len(input)
	if numPadChars >= math.MaxInt {
		php.ErrorDocRef(ctx, "", perr.E_WARNING, "Padding length is too long")
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
func ZifStrShuffle(str string) string {
	if len(str) <= 1 {
		return str
	}

	bin := []byte(str)
	for i := len(str) - 1; i >= 0; i-- {
		rndIdx := rand.Intn(i + 1)
		if rndIdx != i {
			bin[i], bin[rndIdx] = bin[rndIdx], bin[i]
		}
	}
	return string(bin)
}
func ZifStrWordCount(ctx *php.Context, str string, _ zpp.Opt, format int, charlist *string) (*types.Zval, bool) {
	var mask = ""
	if charlist != nil {
		mask, _ = PhpCharmaskEx(ctx, *charlist)
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
		arr := types.NewArrayCap(len(spans))
		for _, span := range spans {
			arr.Append(types.ZvalString(str[span.start:span.end]))
		}
		return types.NewZvalArray(arr), true
	case 2:
		arr := types.NewArrayCap(len(spans))
		for _, span := range spans {
			arr.IndexUpdate(span.start, types.ZvalString(str[span.start:span.end]))
		}
		return types.NewZvalArray(arr), true
	default:
		php.ErrorDocRef(ctx, "", perr.E_WARNING, fmt.Sprintf("Invalid format value %d", format))
		return nil, false
	}
}
func ZifStrSplit(ctx *php.Context, str string, _ zpp.Opt, splitLength_ *int) ([]string, bool) {
	var splitLength = lang.Option(splitLength_, 1)

	if splitLength <= 0 {
		php.ErrorDocRef(ctx, "", perr.E_WARNING, "The length of each segment must be greater than zero")
		return nil, false
	}

	if str == "" || splitLength >= len(str) {
		return []string{str}, true
	}

	size := (len(str) + splitLength - 1) / splitLength
	result := make([]string, size)
	for i := 0; i < len(str); i += splitLength {
		if i+splitLength <= len(str) {
			result = append(result, str[i:])
		} else {
			result = append(result, str[i:i+splitLength])
		}
	}
	return result, true
}

//@zif(onError=1)
func ZifStrpbrk(ctx *php.Context, haystack string, charList string) (string, bool) {
	if charList == "" {
		php.ErrorDocRef(ctx, "", perr.E_WARNING, "The character list cannot be empty")
		return "", false
	}
	if pos := strings.IndexAny(haystack, charList); pos >= 0 {
		return haystack[pos:], true
	}
	return "", false
}

//@zif(onError=1)
func ZifSubstrCompare(ctx *php.Context, return_value zpp.Ret, haystack string, needle string, offset int, _ zpp.Opt, length *int, caseInsensitivity bool) (int, bool) {
	// check length
	if length != nil && *length <= 0 {
		if *length == 0 {
			return_value.SetLong(0)
			return 0, true
		} else {
			php.ErrorDocRef(ctx, "", perr.E_WARNING, "The length must be greater than or equal to zero")
			return 0, false
		}
	}

	// check offset
	if offset < 0 {
		offset = len(haystack) + offset
		if offset < 0 {
			offset = 0
		}
	}
	if offset > len(haystack) {
		php.ErrorDocRef(ctx, "", perr.E_WARNING, "The start position cannot exceed initial string length")
		return_value.SetFalse()
		return 0, false
	}

	// cut
	s1 := haystack[offset:]
	s2 := needle
	if length != nil {
		if len(s1) > *length {
			s1 = s1[:*length]
		}
		if len(s2) > *length {
			s2 = s2[:*length]
		}
	}

	if caseInsensitivity {
		return strings.Compare(s1, s2), true
	} else {
		return ascii.StrCaseCompare(s1, s2), true
	}
}
