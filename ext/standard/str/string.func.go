package str

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/ext/standard"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
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
	newStr := Strtr(s, from, to)
	str *= newStr
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
		if symbol_table.KeyDelete(types.STR_THIS) {
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
			esc = standard.PHP_CSV_NO_ESCAPE
		}
	}

	standard.PhpFgetcsv(nil, delim, enc, esc, str.GetLen(), str.GetVal(), return_value)
}

func ZifSscanf(str string, format string, vars []zpp.RefZval) *types.Zval {
	var args *types.Zval = nil
	var result int
	var num_args int = 0
	result = standard.PhpSscanfInternal(str, format, num_args, args, 0, return_value)
	if standard.SCAN_ERROR_WRONG_PARAM_COUNT == result {
		zend.ZendWrongParamCount()
		return
	}
}

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
func ZifUtf8Decode(data string) string {
	var pos int = 0
	var buf strings.Builder
	for pos < len(data) {
		var status int = types.FAILURE
		c := standard.PhpNextUtf8Char((*uint8)(data), len(data), &pos, &status)

		/* The lower 256 codepoints of Unicode are identical to Latin-1,
		 * so we don't need to do any mapping here beyond replacing non-Latin-1
		 * characters. */
		if status == types.FAILURE || c > 0xff {
			c = '?'
		}
		buf.WriteRune(rune(c))
	}
	return buf.String()
}
