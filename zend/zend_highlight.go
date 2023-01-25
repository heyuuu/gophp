// <<generate>>

package zend

import (
	b "sik/builtin"
)

// Source: <Zend/zend_highlight.h>

/*
   +----------------------------------------------------------------------+
   | Zend Engine                                                          |
   +----------------------------------------------------------------------+
   | Copyright (c) Zend Technologies Ltd. (http://www.zend.com)           |
   +----------------------------------------------------------------------+
   | This source file is subject to version 2.00 of the Zend license,     |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.zend.com/license/2_00.txt.                                |
   | If you did not receive a copy of the Zend license and are unable to  |
   | obtain it through the world-wide-web, please send a note to          |
   | license@zend.com so we can mail you a copy immediately.              |
   +----------------------------------------------------------------------+
   | Authors: Andi Gutmans <andi@php.net>                                 |
   |          Zeev Suraski <zeev@php.net>                                 |
   +----------------------------------------------------------------------+
*/

// #define ZEND_HIGHLIGHT_H

const HL_COMMENT_COLOR = "#FF8000"
const HL_DEFAULT_COLOR = "#0000BB"
const HL_HTML_COLOR = "#000000"
const HL_STRING_COLOR = "#DD0000"
const HL_KEYWORD_COLOR = "#007700"

var SyntaxHighlighterIni ZendSyntaxHighlighterIni

// Source: <Zend/zend_highlight.c>

/*
   +----------------------------------------------------------------------+
   | Zend Engine                                                          |
   +----------------------------------------------------------------------+
   | Copyright (c) Zend Technologies Ltd. (http://www.zend.com)           |
   +----------------------------------------------------------------------+
   | This source file is subject to version 2.00 of the Zend license,     |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.zend.com/license/2_00.txt.                                |
   | If you did not receive a copy of the Zend license and are unable to  |
   | obtain it through the world-wide-web, please send a note to          |
   | license@zend.com so we can mail you a copy immediately.              |
   +----------------------------------------------------------------------+
   | Authors: Andi Gutmans <andi@php.net>                                 |
   |          Zeev Suraski <zeev@php.net>                                 |
   +----------------------------------------------------------------------+
*/

// # include "zend.h"

// # include < zend_language_parser . h >

// # include "zend_compile.h"

// # include "zend_highlight.h"

// # include "zend_ptr_stack.h"

// # include "zend_globals.h"

// # include "zend_exceptions.h"

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
	if LanguageScannerGlobals.GetOutputFilter() != nil {
		LanguageScannerGlobals.GetOutputFilter()(&filtered, &filtered_len, ptr, len_)
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
	if LanguageScannerGlobals.GetOutputFilter() != nil {
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
			ZendHtmlPuts((*byte)(LanguageScannerGlobals.GetYyText()), LanguageScannerGlobals.GetYyLeng())
			ZVAL_UNDEF(&token)
			continue
			break
		default:
			if Z_TYPE(token) == IS_UNDEF {
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
		ZendHtmlPuts((*byte)(LanguageScannerGlobals.GetYyText()), LanguageScannerGlobals.GetYyLeng())
		if Z_TYPE(token) == IS_STRING {
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
			ZendWrite((*byte)(LanguageScannerGlobals.GetYyText()), LanguageScannerGlobals.GetYyLeng())

			/* read the following character, either newline or ; */

			if LexScan(&token, nil) != T_WHITESPACE {
				ZendWrite((*byte)(LanguageScannerGlobals.GetYyText()), LanguageScannerGlobals.GetYyLeng())
			}
			ZendWrite("\n", b.SizeOf("\"\\n\"")-1)
			prev_space = 1
			ZVAL_UNDEF(&token)
			continue
		default:
			ZendWrite((*byte)(LanguageScannerGlobals.GetYyText()), LanguageScannerGlobals.GetYyLeng())
			break
		}
		if Z_TYPE(token) == IS_STRING {
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
