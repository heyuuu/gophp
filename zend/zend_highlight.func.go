// <<generate>>

package zend

import (
	b "sik/builtin"
)

func ZendHtmlPutc(c byte) {
	switch c {
	case '\n':
		ZEND_PUTS("<br />")
		break
	case '<':
		ZEND_PUTS("&lt;")
		break
	case '>':
		ZEND_PUTS("&gt;")
		break
	case '&':
		ZEND_PUTS("&amp;")
		break
	case ' ':
		ZEND_PUTS("&nbsp;")
		break
	case '\t':
		ZEND_PUTS("&nbsp;&nbsp;&nbsp;&nbsp;")
		break
	default:
		ZEND_PUTC(c)
		break
	}
}
func ZendHtmlPuts(s *byte, len_ int) {
	var ptr *uint8 = (*uint8)(s)
	var end *uint8 = ptr + len_
	var filtered *uint8 = nil
	var filtered_len int
	if INI_SCNG__().output_filter {
		INI_SCNG__().output_filter(&filtered, &filtered_len, ptr, len_)
		ptr = filtered
		end = filtered + filtered_len
	}
	for ptr < end {
		if (*ptr) == ' ' {
			for {
				ZendHtmlPutc(*ptr)
				if !(b.PreInc(&ptr) < end && (*ptr) == ' ') {
					break
				}
			}
		} else {
			ZendHtmlPutc(b.PostInc(&(*ptr)))
		}
	}
	if INI_SCNG__().output_filter {
		Efree(filtered)
	}
}
func ZendHighlight(syntax_highlighter_ini *ZendSyntaxHighlighterIni) {
	var token Zval
	var token_type int
	var last_color *byte = syntax_highlighter_ini.GetHighlightHtml()
	var next_color *byte
	ZendPrintf("<code>")
	ZendPrintf("<span style=\"color: %s\">\n", last_color)

	/* highlight stuff coming back from zendlex() */

	for b.Assign(&token_type, LexScan(&token, nil)) {
		switch token_type {
		case T_INLINE_HTML:
			next_color = syntax_highlighter_ini.GetHighlightHtml()
			break
		case T_COMMENT:

		case T_DOC_COMMENT:
			next_color = syntax_highlighter_ini.GetHighlightComment()
			break
		case T_OPEN_TAG:

		case T_OPEN_TAG_WITH_ECHO:

		case T_CLOSE_TAG:

		case T_LINE:

		case T_FILE:

		case T_DIR:

		case T_TRAIT_C:

		case T_METHOD_C:

		case T_FUNC_C:

		case T_NS_C:

		case T_CLASS_C:
			next_color = syntax_highlighter_ini.GetHighlightDefault()
			break
		case '"':

		case T_ENCAPSED_AND_WHITESPACE:

		case T_CONSTANT_ENCAPSED_STRING:
			next_color = syntax_highlighter_ini.GetHighlightString()
			break
		case T_WHITESPACE:
			ZendHtmlPuts((*byte)(INI_SCNG__().GetYyText()), INI_SCNG__().GetYyLeng())
			ZVAL_UNDEF(&token)
			continue
			break
		default:
			if token.IsUndef() {
				next_color = syntax_highlighter_ini.GetHighlightKeyword()
			} else {
				next_color = syntax_highlighter_ini.GetHighlightDefault()
			}
			break
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

			case T_OPEN_TAG_WITH_ECHO:

			case T_CLOSE_TAG:

			case T_WHITESPACE:

			case T_COMMENT:

			case T_DOC_COMMENT:
				break
			default:
				ZvalPtrDtorStr(&token)
				break
			}
		}
		ZVAL_UNDEF(&token)
	}
	if last_color != syntax_highlighter_ini.GetHighlightHtml() {
		ZendPrintf("</span>\n")
	}
	ZendPrintf("</span>\n")
	ZendPrintf("</code>")

	/* Discard parse errors thrown during tokenization */

	ZendClearException()

	/* Discard parse errors thrown during tokenization */
}
func ZendStrip() {
	var token Zval
	var token_type int
	var prev_space int = 0
	for b.Assign(&token_type, LexScan(&token, nil)) {
		switch token_type {
		case T_WHITESPACE:
			if prev_space == 0 {
				ZendWrite(" ", b.SizeOf("\" \"")-1)
				prev_space = 1
			}
		case T_COMMENT:

		case T_DOC_COMMENT:
			ZVAL_UNDEF(&token)
			continue
		case T_END_HEREDOC:
			ZendWrite((*byte)(INI_SCNG__().GetYyText()), INI_SCNG__().GetYyLeng())

			/* read the following character, either newline or ; */

			if LexScan(&token, nil) != T_WHITESPACE {
				ZendWrite((*byte)(INI_SCNG__().GetYyText()), INI_SCNG__().GetYyLeng())
			}
			ZendWrite("\n", b.SizeOf("\"\\n\"")-1)
			prev_space = 1
			ZVAL_UNDEF(&token)
			continue
		default:
			ZendWrite((*byte)(INI_SCNG__().GetYyText()), INI_SCNG__().GetYyLeng())
			break
		}
		if token.IsString() {
			switch token_type {
			case T_OPEN_TAG:

			case T_OPEN_TAG_WITH_ECHO:

			case T_CLOSE_TAG:

			case T_WHITESPACE:

			case T_COMMENT:

			case T_DOC_COMMENT:
				break
			default:
				ZvalPtrDtorStr(&token)
				break
			}
		}
		prev_space = 0
		ZVAL_UNDEF(&token)
	}

	/* Discard parse errors thrown during tokenization */

	ZendClearException()

	/* Discard parse errors thrown during tokenization */
}
