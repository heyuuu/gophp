package standard

import (
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/php/zpp"
	"strings"
)

type CsvSetting struct {
	Delimiter byte
	Enclosure byte
	Escape    byte
	NoEscape  bool
}

func buildCsvSettingByParams(ctx *php.Context, delimiter *string, enclosure *string, escape *string) (setting CsvSetting, ok bool) {
	var delimiterStr string = lang.Option(delimiter, `,`)
	var enclosureStr string = lang.Option(enclosure, `"`)
	var escapeStr string = lang.Option(escape, `\`)
	if delimiterStr == "" {
		php.ErrorDocRef(ctx, "", perr.E_WARNING, "delimiter must be a character")
		return
	} else if len(delimiterStr) > 1 {
		php.ErrorDocRef(ctx, "", perr.E_NOTICE, "delimiter must be a single character")
	}
	if enclosureStr == "" {
		php.ErrorDocRef(ctx, "", perr.E_WARNING, "enclosure must be a character")
		return
	} else if len(enclosureStr) > 1 {
		php.ErrorDocRef(ctx, "", perr.E_NOTICE, "enclosure must be a single character")
	}
	if len(escapeStr) > 1 {
		php.ErrorDocRef(ctx, "", perr.E_NOTICE, "escape must be empty or a single character")
	}

	setting.Delimiter = delimiterStr[0]
	setting.Enclosure = enclosureStr[0]
	if len(escapeStr) < 1 {
		setting.NoEscape = true
	} else {
		/* use first character from string */
		setting.Escape = escapeStr[0]
	}
	return setting, true
}

func SprintCsvLine(ctx *php.Context, fields *types.Array, setting CsvSetting) string {
	var escapeChars string
	if setting.NoEscape {
		escapeChars = string([]byte{'\r', '\n', '\t', ' ', setting.Delimiter, setting.Enclosure})
	} else {
		escapeChars = string([]byte{'\r', '\n', '\t', ' ', setting.Delimiter, setting.Enclosure, setting.Escape})
	}

	var buf strings.Builder
	var first bool
	fields.Each(func(_ types.ArrayKey, field types.Zval) {
		// delimiter
		if !first {
			buf.WriteByte(setting.Delimiter)
		}
		first = true

		// field
		var fieldStr = php.ZvalGetStrVal(ctx, field)
		if strings.ContainsAny(fieldStr, escapeChars) {
			var escaped = 0
			buf.WriteByte(setting.Enclosure)
			for _, ch := range []byte(fieldStr) {
				if !setting.NoEscape && ch == setting.Escape {
					escaped = 1
				} else if escaped == 0 && ch == setting.Enclosure {
					buf.WriteByte(setting.Enclosure)
				} else {
					escaped = 0
				}
				buf.WriteByte(ch)
			}
			buf.WriteByte(setting.Enclosure)
		} else {
			buf.WriteString(fieldStr)
		}
	})
	buf.WriteByte('\n')
	return buf.String()
}

func ZifStrGetcsv(string_ string, _ zpp.Opt, delimiter string, enclosure string, escape *string) *types.Array {
	if string_ == "" {
		return types.NewArrayOf(types.Null)
	}

	settings := CsvSetting{Delimiter: ',', Enclosure: '"', Escape: '\\'}
	if delimiter != "" {
		settings.Delimiter = delimiter[0]
	}
	if enclosure != "" {
		settings.Enclosure = enclosure[0]
	}
	if escape != nil {
		if *escape != "" {
			settings.Escape = (*escape)[0]
		} else {
			settings.NoEscape = true
		}
	}

	return types.NewArrayOfString(strGetCsvLine(string_, settings))
}

func strGetCsvLine(s string, setting CsvSetting) []string {
	delimiter := setting.Delimiter
	enclosure := setting.Enclosure
	escape := setting.Escape
	noEscape := setting.NoEscape

	if strings.HasSuffix(s, "\r\n") {
		s = s[:len(s)-2]
	} else if strings.HasSuffix(s, "\n") {
		s = s[:len(s)-1]
	}

	var result = make([]string, 0, strings.Count(s, string(delimiter))+1)
	var buf strings.Builder
	start := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == enclosure { // `"` 进入
			buf.Reset()
			for i++; i < len(s); i++ {
				c = s[i]
				if c == enclosure && i+1 < len(s) && s[i+1] == enclosure { // `""` 转义
					buf.WriteByte(enclosure)
					i++
				} else if !noEscape && c == escape && i+1 < len(s) && s[i+1] == enclosure { // `\"` 转义
					buf.WriteByte(escape)
					buf.WriteByte(enclosure)
					i++
				} else if c == enclosure { // `"` 退出
					i++
					break
				} else {
					buf.WriteByte(c)
				}
			}
			for ; i < len(s) && s[i] != delimiter; i++ {
				buf.WriteByte(s[i])
			}
			result = append(result, buf.String())
			start = i + 1
		} else if c == delimiter {
			result = append(result, s[start:i])
			start = i + 1
		}
	}
	if start <= len(s) {
		result = append(result, s[start:])
	}

	return result
}
