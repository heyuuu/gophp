package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
)

func ZendHtmlPutc(c byte) {
	switch c {
	case '\n':
		ZEND_PUTS("<br />")
	case '<':
		ZEND_PUTS("&lt;")
	case '>':
		ZEND_PUTS("&gt;")
	case '&':
		ZEND_PUTS("&amp;")
	case ' ':
		ZEND_PUTS("&nbsp;")
	case '\t':
		ZEND_PUTS("&nbsp;&nbsp;&nbsp;&nbsp;")
	default:
		ZEND_PUTC(c)
	}
}
func ZendHtmlPuts(s *byte, len_ int) {
	ZendHtmlPutsEx(b.CastStr(s, len_))
}
func ZendHtmlPutsEx(s string) {
	//outputFilter := INI_SCNG__().output_filter
	outputFilter := getOutputFilter()
	if outputFilter != nil {
		s = outputFilter(s)
	}

	// 输出第一个空格前的内容及其后的连续空格
	for i, c := range []byte(s) {
		if c == ' ' {
			for _, c0 := range []byte(s[i:]) {
				if c0 == ' ' {
					ZendHtmlPutc(c0)
				} else {
					break
				}
			}

			break
		} else {
			ZendHtmlPutc(c)
		}
	}
}
func ZendHighlight(syntax_highlighter_ini *ZendSyntaxHighlighterIni) {
	var token types.Zval
	var token_type int
	var last_color *byte = syntax_highlighter_ini.GetHighlightHtml()
	var next_color *byte
	ZendPrintf("<code>")
	ZendPrintf("<span style=\"color: %s\">\n", last_color)

	/* highlight stuff coming back from zendlex() */

	for lang.Assign(&token_type, LexScan(&token, nil)) {
		switch token_type {
		case T_INLINE_HTML:
			next_color = syntax_highlighter_ini.GetHighlightHtml()
		case T_COMMENT:
			fallthrough
		case T_DOC_COMMENT:
			next_color = syntax_highlighter_ini.GetHighlightComment()
		case T_OPEN_TAG:
			fallthrough
		case T_OPEN_TAG_WITH_ECHO:
			fallthrough
		case T_CLOSE_TAG:
			fallthrough
		case T_LINE:
			fallthrough
		case T_FILE:
			fallthrough
		case T_DIR:
			fallthrough
		case T_TRAIT_C:
			fallthrough
		case T_METHOD_C:
			fallthrough
		case T_FUNC_C:
			fallthrough
		case T_NS_C:
			fallthrough
		case T_CLASS_C:
			next_color = syntax_highlighter_ini.GetHighlightDefault()
		case '"':
			fallthrough
		case T_ENCAPSED_AND_WHITESPACE:
			fallthrough
		case T_CONSTANT_ENCAPSED_STRING:
			next_color = syntax_highlighter_ini.GetHighlightString()
		case T_WHITESPACE:
			ZendHtmlPuts((*byte)(INI_SCNG__().GetYyText()), INI_SCNG__().GetYyLeng())
			token.SetUndef()
			continue
		default:
			if token.IsUndef() {
				next_color = syntax_highlighter_ini.GetHighlightKeyword()
			} else {
				next_color = syntax_highlighter_ini.GetHighlightDefault()
			}
		}
		if last_color != next_color {
			if last_color != syntax_highlighter_ini.GetHighlightHtml() {
				ZendPrintf("</span>")
			}
			last_color = next_color
			if last_color != syntax_highlighter_ini.GetHighlightHtml() {
				ZendPrintf("<span style=\"color: %s\">", last_color)
			}
		}
		ZendHtmlPuts((*byte)(INI_SCNG__().GetYyText()), INI_SCNG__().GetYyLeng())
		if token.IsString() {
			switch token_type {
			case T_OPEN_TAG:
				fallthrough
			case T_OPEN_TAG_WITH_ECHO:
				fallthrough
			case T_CLOSE_TAG:
				fallthrough
			case T_WHITESPACE:
				fallthrough
			case T_COMMENT:
				fallthrough
			case T_DOC_COMMENT:

			default:

			}
		}
		token.SetUndef()
	}
	if last_color != syntax_highlighter_ini.GetHighlightHtml() {
		ZendPrintf("</span>\n")
	}
	ZendPrintf("</span>\n")
	ZendPrintf("</code>")

	/* Discard parse errors thrown during tokenization */
	EG__().ClearException()
}
func ZendStrip() {
	var token types.Zval
	var token_type int
	var prev_space int = 0
	for lang.Assign(&token_type, LexScan(&token, nil)) {
		switch token_type {
		case T_WHITESPACE:
			if prev_space == 0 {
				ZendWrite(" ")
				prev_space = 1
			}
			fallthrough
		case T_COMMENT:
			fallthrough
		case T_DOC_COMMENT:
			token.SetUndef()
			continue
		case T_END_HEREDOC:
			ZendWrite(b.CastStr(INI_SCNG__().GetYyText(), INI_SCNG__().GetYyLeng()))

			/* read the following character, either newline or ; */

			if LexScan(&token, nil) != T_WHITESPACE {
				ZendWrite(b.CastStr(INI_SCNG__().GetYyText(), INI_SCNG__().GetYyLeng()))
			}
			ZendWrite("\n")
			prev_space = 1
			token.SetUndef()
			continue
		default:
			ZendWrite(b.CastStr(INI_SCNG__().GetYyText(), INI_SCNG__().GetYyLeng()))
		}
		if token.IsString() {
			switch token_type {
			case T_OPEN_TAG:
				fallthrough
			case T_OPEN_TAG_WITH_ECHO:
				fallthrough
			case T_CLOSE_TAG:
				fallthrough
			case T_WHITESPACE:
				fallthrough
			case T_COMMENT:
				fallthrough
			case T_DOC_COMMENT:

			default:

			}
		}
		prev_space = 0
		token.SetUndef()
	}

	/* Discard parse errors thrown during tokenization */
	EG__().ClearException()
}
