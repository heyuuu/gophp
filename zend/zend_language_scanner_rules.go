package zend

import (
	"math"
	b "sik/builtin"
	"strings"
)

func (sc *LangScanner) parseNumStr(str string, base int) (lVal int, dVal float64, overflow bool) {
	// todo 越界的边际条件需再确认
	overflowLimit := math.MaxInt / base
	for _, c := range str {
		var digit int
		if c == '_' {
			continue
		} else if c >= '0' && c <= '9' {
			digit = int(c - '0')
		} else if c >= 'a' && c <= 'f' {
			digit = int(c - 'a')
		} else if c >= 'A' && c <= 'F' {
			digit = int(c - 'a')
		} else {
			// todo
		}

		if digit >= base {
			ZendThrowException(ZendCeParseError, "Invalid numeric literal", 0)
		}

		// 预期数字会越界时，改用浮点数
		if !overflow && (lVal >= overflowLimit) {
			overflow = true
			dVal = float64(lVal)
		}

		if overflow {
			dVal = dVal*float64(base) + float64(digit)
		} else {
			lVal = lVal*base + digit
		}
	}
	return
}

func (sc *LangScanner) tokenWithNumStr(offset int, base int) (int, bool) {
	/* skips "0b" or "0x" */
	var str = sc.seg()[offset:]

	// base == 8 的单独字符检查
	if base == 8 {
		for i, c := range str {
			if c == '8' || c == '9' {
				ZendThrowException(ZendCeParseError, "Invalid numeric literal", 0)
				if sc.isParserMode() {
					sc.zendlval.SetUndef()
					return sc.token(T_ERROR)
				}

				/* Continue in order to determine if this is T_LNUMBER or T_DNUMBER. */
				str = str[:i]
				break
			}
		}
	}

	// 字符串传数值
	lVal, dVal, overflow := sc.parseNumStr(str, base)

	// 记录并 zval 并返回 token
	if !overflow {
		sc.zendlval.SetLong(lVal)
		return sc.tokenWithVal(T_LNUMBER)
	} else {
		sc.zendlval.SetDouble(dVal)
		return sc.tokenWithVal(T_DNUMBER)
	}
}

func (sc *LangScanner) lexerOffsetNum() (int, bool) {
	str := sc.seg()
	sc.zendlval.SetRawString(str)
	return sc.tokenWithVal(T_NUM_STRING)
}
func (sc *LangScanner) lexerDNum() (int, bool) {
	str := sc.seg()
	str = strings.ReplaceAll(str, "_", "")
	dVal := ZendStrtod(str)
	sc.zendlval.SetDouble(dVal)
	return sc.tokenWithVal(T_DNUMBER)
}

func (sc *LangScanner) lexerRule12() (int, bool) {
	/* Allow <?php followed by end of file. */
	if sc.yyCursor == sc.yyLimit {
		sc.begin(yycST_IN_SCRIPTING)
		return sc.returnOrSkipToken(T_OPEN_TAG)
	}

	/* Degenerate case: <?phpX is interpreted as <? phpX with short tags. */
	if sc.shortTags {
		sc.setLen(2)
		sc.begin(yycST_IN_SCRIPTING)
		return sc.returnOrSkipToken(T_OPEN_TAG)
	}
	return sc.inlineCharHandler()
}
func (sc *LangScanner) lexerRule13() (int, bool) {
	if sc.shortTags {
		sc.begin(yycST_IN_SCRIPTING)
		return sc.returnOrSkipToken(T_OPEN_TAG)
	} else {
		return sc.inlineCharHandler()
	}
}
func (sc *LangScanner) lexerRule14() (int, bool) {
	if sc.isEnd() {
		return sc.token(END)
	}
	return sc.inlineCharHandler()
}

func (sc *LangScanner) inlineCharHandler() (int, bool) {
	// 读取到 openTag 或文件结尾，可优化
	for {
		pos := strings.IndexByte(sc.tail(), '<')
		if pos >= 0 {
			sc.cursor += uint(pos) + 1
		} else {
			sc.cursor = sc.limit
		}
		if !sc.canRead() {
			break
		}
		if sc.peek() == '?' {
			isOpenTag := false
			if sc.shortTags {
				isOpenTag = true
			} else if sc.peekStrIs("?=") {
				isOpenTag = true
			} else if sc.peekStrIs("?php") && (sc.cursor+4 == sc.limit || strings.IndexByte(" \t\r\n", sc.peekOffset(4)) >= 0) {
				isOpenTag = true
			}
			if isOpenTag {
				sc.fallback()
				break
			}
		}
	}

	sc.resetLen()
	var str = sc.matched()
	sc.setStrFiltered(str)
	sc.handleNewlinesEx(str)
	return sc.tokenWithVal(T_INLINE_HTML)
}

func (sc *LangScanner) lexerRule16() (int, bool) {
	for sc.canRead() {
		switch sc.read() {
		case '\r':
			if sc.peek() == '\n' {
				sc.skip()
			}
			fallthrough
		case '\n':
			sc.lineno++
		case '?':
			if sc.peek() == '>' {
				sc.skip()
				break
			}
			fallthrough
		default:
			continue
		}
		break
	}
	sc.resetLen()
	return sc.returnOrSkipToken(T_COMMENT)
}
func (sc *LangScanner) lexerRule17() (int, bool) {
	isDocComment := false
	if sc.len_ > 2 {
		isDocComment = true
		RESET_DOC_COMMENT()
	}

	for sc.canRead() {
		if sc.read() == '*' && sc.peek() == '/' {
			break
		}
	}

	if sc.canRead() {
		sc.skip()
	} else if !sc.heredocScanAhead {
		ZendError(E_COMPILE_WARNING, "Unterminated comment starting line %d", CG__().zend_lineno)
	}

	sc.resetLen()
	sc.handleNewlinesEx(sc.seg())
	if isDocComment {
		CG__().doc_comment = NewZendString(sc.seg())
		return sc.returnOrSkipToken(T_DOC_COMMENT)
	}
	return sc.returnOrSkipToken(T_COMMENT)
}
func (sc *LangScanner) lexerRule18() (int, bool) {
	sc.begin(yycINITIAL)
	seg := sc.seg()
	if seg[len(seg)-1] != '>' {
		CG__().increment_lineno = 1
	}
	if sc.isParserMode() {
		return sc.token(';')
	}
	return sc.token(T_CLOSE_TAG)
}
func (sc *LangScanner) lexerRule19() (int, bool) {
	bPrefix := 0
	if sc.yyText0() != '\'' {
		bPrefix = 1
	}
	for {
		if sc.canRead() {
			if sc.peek() == '\'' {
				sc.skip()
				sc.resetLen()
				break
			} else if sc.read() == '\\' && sc.canRead() {
				sc.skip()
			}
		} else {
			/* Unclosed single quotes; treat similar to double quotes, but without a separate token
			 * for ' (unrecognized by parser), instead of old flex fallback to "Unexpected character..."
			 * rule, which continued in ST_IN_SCRIPTING state after the quote */
			sc.zendlval.SetNull()
			return sc.tokenWithVal(T_ENCAPSED_AND_WHITESPACE)
		}
	}

	str := sc.seg()
	str = str[bPrefix+1 : len(str)-bPrefix-2]
	str = strEscape(str)
	sc.handleNewlinesEx(str)
	sc.setStrFiltered(str)
	return sc.tokenWithVal(T_CONSTANT_ENCAPSED_STRING)
}

func (sc *LangScanner) lexerRule20() (int, bool) {
	bprefix := 0
	if sc.yyText0() != '"' {
		bprefix = 1
	}
	for sc.canRead() {
		switch sc.read() {
		case '"':
			sc.resetLen()
			str := sc.seg()
			str = str[bprefix+1 : len(str)-bprefix-2]

			if sc.ScanEscapeString(str, '"') || !sc.isParserMode() {
				return sc.tokenWithVal(T_CONSTANT_ENCAPSED_STRING)
			} else {
				return sc.token(T_ERROR)
			}
		case '$':
			if isLabelStart(sc.peek()) || sc.peek() == '{' {
				break
			}
			continue
		case '{':
			if sc.peek() == '$' {
				break
			}
			continue
		case '\\':
			if sc.canRead() {
				sc.skip()
			}
			fallthrough
		default:
			continue
		}
		sc.fallback()
		break
	}

	/* Remember how much was scanned to save rescanning */

	sc.setDoubleQuotesScannedLength(sc.cursor - sc.text - sc.len_)
	sc.cursor = sc.text + sc.len_
	sc.begin(yycST_DOUBLE_QUOTES)
	return sc.token('"')
}
func (sc *LangScanner) lexerRule21() (int, bool) {
	var s *byte
	var savedCursor *uint8
	var bprefix int = b.Cond(sc.yyText0() != '<', 1, 0)
	var spacing int = 0
	var indentation int = 0
	var heredocLabel = &ZendHeredocLabel{}
	var isHeredoc zend_bool = 1
	sc.lineno++
	heredocLabel.length = sc.len_ - bprefix - 3 - 1 - b.Cond(sc.yyText[sc.len_-2] == '\r', 1, 0)
	s = sc.yyText + bprefix + 3
	for (*s) == ' ' || (*s) == '\t' {
		s++
		heredocLabel.length--
	}
	if (*s) == '\'' {
		s++
		heredocLabel.length -= 2
		isHeredoc = 0
		sc.begin(yycST_NOWDOC)
	} else {
		if (*s) == '"' {
			s++
			heredocLabel.length -= 2
		}
		sc.begin(yycST_HEREDOC)
	}
	heredocLabel.label = estrndup(s, heredocLabel.length)
	heredocLabel.indentation = 0
	savedCursor = sc.yyCursor
	sc.heredocLabelStack.Push(heredocLabel)
	for sc.canRead() && (sc.peek() == ' ' || sc.peek() == '\t') {
		if sc.peek() == '\t' {
			spacing |= HEREDOC_USING_TABS
		} else {
			spacing |= HEREDOC_USING_SPACES
		}
		sc.skip()
		indentation++
	}
	if sc.yyCursor == sc.yyLimit {
		sc.yyCursor = savedCursor
		return sc.token(T_START_HEREDOC)
	}

	/* Check for ending label on the next line */

	if heredocLabel.length < sc.yyLimit-sc.yyCursor && !(memcmp(sc.yyCursor, s, heredocLabel.length)) {
		if !(isLabelSuccessor(sc.yyCursor[heredocLabel.length])) {
			if spacing == (HEREDOC_USING_SPACES | HEREDOC_USING_TABS) {
				ZendThrowException(ZendCeParseError, "Invalid indentation - tabs and spaces cannot be mixed", 0)
				if sc.isParserMode() {
					return sc.token(T_ERROR)
				}
			}
			sc.yyCursor = savedCursor
			heredocLabel.indentation = indentation
			sc.begin(yycST_END_HEREDOC)
			return sc.token(T_START_HEREDOC)
		}
	}
	sc.yyCursor = savedCursor
	if isHeredoc && !sc.heredocScanAhead {
		var current_state zend_lex_state
		var saved_doc_comment *zend_string = CG__().doc_comment
		var heredoc_nesting_level int = 1
		var first_token int = 0
		var error int = 0
		ZendSaveLexicalState(&current_state)
		sc.heredocScanAhead = 1
		sc.heredocIndentation = 0
		sc.heredocIndentationUsesSpaces = 0
		sc.onEvent = nil
		CG__().doc_comment = nil
		zend_ptr_stack_reverse_apply(current_state.heredoc_label_stack, CopyHeredocLabelStack)
		zend_exception_save()
		for heredoc_nesting_level != 0 {
			var zv zval
			var retval int
			ZVAL_UNDEF(&zv)
			retval = lex_scan(&zv, nil)
			zval_ptr_dtor_nogc(&zv)
			if EG__().exception {
				zend_clear_exception()
				break
			}
			if first_token == 0 {
				first_token = retval
			}
			switch retval {
			case T_START_HEREDOC:
				heredoc_nesting_level++
			case T_END_HEREDOC:
				heredoc_nesting_level--
			case END:
				heredoc_nesting_level = 0
			}
		}
		zend_exception_restore()
		if (first_token == T_VARIABLE || first_token == T_DOLLAR_OPEN_CURLY_BRACES || first_token == T_CURLY_OPEN) && sc.heredocIndentation {
			ZendThrowExceptionEx(ZendCeParseError, 0, "Invalid body indentation level (expecting an indentation level of at least %d)", sc.heredocIndentation)
			error = 1
		}
		heredocLabel.indentation = sc.heredocIndentation
		heredocLabel.indentation_uses_spaces = sc.heredocIndentationUsesSpaces
		ZendRestoreLexicalState(&current_state)
		sc.heredocScanAhead = false
		CG__().increment_lineno = 0
		CG__().doc_comment = saved_doc_comment
		if sc.isParserMode() && error != 0 {
			return sc.token(T_ERROR)
		}
	}
	return sc.token(T_START_HEREDOC)
}
func (sc *LangScanner) lexerRule22() (int, bool) {
	var heredocLabel = sc.heredocLabelStack.Pop()
	sc.len_ = heredocLabel.indentation + heredocLabel.length
	sc.yyCursor += sc.len_ - 1
	HeredocLabelDtor(heredocLabel)
	efree(heredocLabel)
	sc.begin(yycST_IN_SCRIPTING)
	return sc.token(T_END_HEREDOC)
}
func (sc *LangScanner) lexerRule23() (int, bool) {
	if sc.getDoubleQuotesScannedLength() != 0 {
		sc.yyCursor += sc.getDoubleQuotesScannedLength() - 1
		sc.setDoubleQuotesScannedLength(0)
		goto double_quotes_scan_done
	}
	if sc.isEnd() {
		return sc.token(END)
	}
	if sc.yyText[0] == '\\' && sc.canRead() {
		sc.skip()
	}
	for sc.canRead() {
		switch sc.read() {
		case '"':

		case '$':
			if isLabelStart(sc.peek()) || sc.peek() == '{' {
				break
			}
			continue
		case '{':
			if sc.peek() == '$' {
				break
			}
			continue
		case '\\':
			if sc.canRead() {
				sc.skip()
			}
			fallthrough
		default:
			continue
		}
		sc.fallback()
		break
	}
double_quotes_scan_done:
	sc.resetLen()
	if ZendScanEscapeString(zendlval, sc.yyText, sc.len_, '"') == SUCCESS || !(sc.isParserMode()) {
		return sc.tokenWithVal(T_ENCAPSED_AND_WHITESPACE)
	} else {
		return sc.token(T_ERROR)
	}
}
func (sc *LangScanner) lexerRule24() (int, bool) {
	if sc.isEnd() {
		return sc.token(END)
	}
	if sc.yyText0() == '\\' && sc.canRead() {
		sc.skip()
	}
	for sc.canRead() {
		switch sc.read() {
		case '`':
		case '$':
			if isLabelStart(sc.peek()) || sc.peek() == '{' {
				break
			}
			continue
		case '{':
			if sc.peek() == '$' {
				break
			}
			continue
		case '\\':
			if sc.canRead() {
				sc.read()
			}
			fallthrough
		default:
			continue
		}
		sc.fallback()
		break
	}
	sc.resetLen()
	if ZendScanEscapeString(zendlval, sc.yyText, sc.len_, '`') == SUCCESS || !(sc.isParserMode()) {
		return sc.tokenWithVal(T_ENCAPSED_AND_WHITESPACE)
	} else {
		return sc.token(T_ERROR)
	}
}
func (sc *LangScanner) lexerRule25() (int, bool) {
	var heredocLabel *ZendHeredocLabel = sc.heredocLabelStack.Top()
	var newline int = 0
	var indentation int = 0
	var spacing int = 0
	if sc.isEnd() {
		return sc.token(END)
	}
	sc.fallback()
	for sc.canRead() {
		switch sc.read() {
		case '\r':
			if sc.peek() == '\n' {
				sc.skip()
			}
			fallthrough
		case '\n':
			spacing = 0
			indentation = spacing
			for sc.canRead() && (sc.peek() == ' ' || sc.peek() == '\t') {
				if sc.peek() == '\t' {
					spacing |= HEREDOC_USING_TABS
				} else {
					spacing |= HEREDOC_USING_SPACES
				}
				sc.skip()
				indentation++
			}
			if sc.yyCursor == sc.yyLimit {
				sc.resetLen()
				sc.handleNewlines(sc.yyText, sc.len_)
				sc.zendlval.SetNull()
				return sc.tokenWithVal(T_ENCAPSED_AND_WHITESPACE)
			}

			/* Check for ending label on the next line */

			if isLabelStart(sc.peek()) && heredocLabel.length < sc.yyLimit-sc.yyCursor && !(memcmp(sc.yyCursor, heredocLabel.label, heredocLabel.length)) {
				if isLabelSuccessor(sc.yyCursor[heredocLabel.length]) {
					continue
				}
				if spacing == (HEREDOC_USING_SPACES | HEREDOC_USING_TABS) {
					ZendThrowException(ZendCeParseError, "Invalid indentation - tabs and spaces cannot be mixed", 0)
					if sc.isParserMode() {
						return sc.token(T_ERROR)
					}
				}

				/* newline before label will be subtracted from returned text, but
				 * yyleng/yytext will include it, for zend_highlight/strip, tokenizer, etc. */

				if sc.yyCursor[-indentation-2] == '\r' && sc.yyCursor[-indentation-1] == '\n' {
					newline = 2
				} else {
					newline = 1
				}
				CG__().increment_lineno = 1
				if sc.heredocScanAhead {
					sc.heredocIndentation = indentation
					sc.heredocIndentationUsesSpaces = spacing == HEREDOC_USING_SPACES
				} else {
					sc.yyCursor -= indentation
				}
				sc.begin(yycST_END_HEREDOC)
				goto heredoc_scan_done
			}
			continue
		case '$':
			if isLabelStart(sc.peek()) || sc.peek() == '{' {
				break
			}
			continue
		case '{':
			if sc.peek() == '$' {
				break
			}
			continue
		case '\\':
			if sc.canRead() && sc.peek() != '\n' && sc.peek() != '\r' {
				sc.skip()
			}
			fallthrough
		default:
			continue
		}
		sc.fallback()
		break
	}
heredoc_scan_done:
	sc.resetLen()
	ZVAL_STRINGL(zendlval, sc.yyText, sc.len_-newline)
	if !(sc.heredocScanAhead) && !(EG__().exception) && sc.isParserMode() {
		var newline_at_start zend_bool = (*(sc.yyText - 1)) == '\n' || (*(sc.yyText - 1)) == '\r'
		var copy *zend_string = Z_STR_P(zendlval)
		if !(StripMultilineStringIndentation(zendlval, heredocLabel.indentation, heredocLabel.indentation_uses_spaces, newline_at_start, newline != 0)) {
			return sc.token(T_ERROR)
		}
		if ZendScanEscapeString(zendlval, ZSTR_VAL(copy), ZSTR_LEN(copy), 0) != SUCCESS {
			zend_string_efree(copy)
			return sc.token(T_ERROR)
		}
		zend_string_efree(copy)
	} else {
		sc.handleNewlines(sc.yyText, sc.len_-newline)
	}
	return sc.tokenWithVal(T_ENCAPSED_AND_WHITESPACE)
}
func (sc *LangScanner) lexerRule26() (int, bool) {
	if sc.isEnd() {
		return sc.token(END)
	}

	var heredocLabel = sc.heredocLabelStack.Top()
	var newline int = 0
	var indentation int = 0
	var spacing int = -1
	sc.fallback()
	for sc.canRead() {
		switch sc.read() {
		case '\r':
			if sc.peek() == '\n' {
				sc.skip()
			}
			fallthrough
		case '\n':
			spacing = 0
			indentation = 0
			for sc.canRead() && sc.peekIs(' ', '\t') {
				if sc.peek() == '\t' {
					spacing |= HEREDOC_USING_TABS
				} else {
					spacing |= HEREDOC_USING_SPACES
				}
				sc.skip()
				indentation++
			}
			if sc.yyCursor == sc.yyLimit {
				sc.resetLen()
				sc.handleNewlinesEx(sc.seg())
				sc.zendlval.SetNull()
				return sc.tokenWithVal(T_ENCAPSED_AND_WHITESPACE)
			}

			/* Check for ending label on the next line */

			if isLabelStart(sc.peek()) && sc.HasPrefix(heredocLabel.GetLabel()) {
				if isLabelSuccessor(sc.yyCursor[heredocLabel.length]) {
					continue
				}
				if spacing == (HEREDOC_USING_SPACES | HEREDOC_USING_TABS) {
					ZendThrowException(ZendCeParseError, "Invalid indentation - tabs and spaces cannot be mixed", 0)
					if sc.isParserMode() {
						return sc.token(T_ERROR)
					}
				}

				/* newline before label will be subtracted from returned text, but
				 * yyleng/yytext will include it, for zend_highlight/strip, tokenizer, etc. */
				if sc.yyCursor[-indentation-2] == '\r' && sc.yyCursor[-indentation-1] == '\n' {
					newline = 2
				} else {
					newline = 1
				}
				CG__().increment_lineno = 1
				sc.yyCursor -= indentation
				heredocLabel.indentation = indentation
				sc.begin(yycST_END_HEREDOC)
				goto nowdoc_scan_done
			}
			fallthrough
		default:
			continue
		}
	}
nowdoc_scan_done:
	sc.resetLen()
	ZVAL_STRINGL(zendlval, sc.yyText, sc.len_-newline)
	if !(EG__().exception) && spacing != -1 && sc.isParserMode() {
		var newline_at_start zend_bool = (*(sc.yyText - 1)) == '\n' || (*(sc.yyText - 1)) == '\r'
		if !(StripMultilineStringIndentation(zendlval, indentation, spacing == HEREDOC_USING_SPACES, newline_at_start, newline != 0)) {
			return sc.token(T_ERROR)
		}
	}
	sc.handleNewlines(sc.yyText, sc.len_-newline)
	return sc.tokenWithVal(T_ENCAPSED_AND_WHITESPACE)
}
func (sc *LangScanner) lexerRule27() (int, bool) {
	if sc.isEnd() {
		return sc.token(END)
	}
	if !sc.heredocScanAhead {
		ZendError(E_COMPILE_WARNING, "Unexpected character in input:  '%c' (ASCII=%d) state=%d", sc.yyText0(), sc.yyText0(), sc.yyState)
	}
	if sc.isParserMode() {
		return 0, true
	} else {
		return sc.token(T_BAD_CHARACTER)
	}
}
