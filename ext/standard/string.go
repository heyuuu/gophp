package standard

import (
	"encoding/hex"
	"fmt"
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/builtin/ascii"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
	"github.com/heyuuu/gophp/zend/zpp"
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
