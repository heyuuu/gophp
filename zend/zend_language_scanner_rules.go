package zend

import (
	"math"
	b "sik/builtin"
	"sik/zend/faults"
	"sik/zend/types"
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
			faults.ZendThrowException(faults.ZendCeParseError, "Invalid numeric literal", 0)
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
				faults.ZendThrowException(faults.ZendCeParseError, "Invalid numeric literal", 0)
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
	dVal := strToD(str)
	sc.zendlval.SetDouble(dVal)
	return sc.tokenWithVal(T_DNUMBER)
}

func (sc *LangScanner) lexerRule12() (int, bool) {
	/* Allow <?php followed by end of file. */
	if sc.cursor == sc.limit {
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
		sc.resetDocComment()
	}

	for sc.canRead() {
		if sc.read() == '*' && sc.peek() == '/' {
			break
		}
	}

	if sc.canRead() {
		sc.skip()
	} else if !sc.heredocScanAhead {
		faults.ZendError(faults.E_COMPILE_WARNING, "Unterminated comment starting line %d", CG__().zend_lineno)
	}

	sc.resetLen()
	sc.handleNewlinesEx(sc.seg())
	if isDocComment {
		sc.setDocComment(sc.seg())
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

			if sc.setEscapeString(str, '"') || !sc.isParserMode() {
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
	// 去除前缀 b?<<<，去除后缀 \r\n|\n, 去除空格和\t
	label := strings.Trim(sc.seg()[1:], "<\r\n \t")
	// 记录换行行号
	sc.lineno++

	isHeredoc := true
	if label[0] == '\'' {
		label = label[1 : len(label)-1]
		isHeredoc = false
		sc.begin(yycST_NOWDOC)
	} else {
		if label[0] == '"' {
			label = label[1 : len(label)-1]
		}
		sc.begin(yycST_HEREDOC)
	}

	heredocLabel := NewHeredocLabel(label)
	savedCursor := sc.cursor
	sc.heredocLabelStack.Push(heredocLabel)

	spacing := 0
	indentation := 0
	for sc.canRead() && sc.peekIs(' ', '\t') {
		if sc.peek() == '\t' {
			spacing |= HEREDOC_USING_TABS
		} else {
			spacing |= HEREDOC_USING_SPACES
		}
		sc.skip()
		indentation++
	}
	if sc.cursor == sc.limit {
		sc.cursor = savedCursor
		return sc.token(T_START_HEREDOC)
	}

	/* Check for ending label on the next line */
	if sc.peekStrIs(label) {
		if !(isLabelSuccessor(sc.peekOffset(len(label)))) {
			if spacing == (HEREDOC_USING_SPACES | HEREDOC_USING_TABS) {
				faults.ZendThrowException(faults.ZendCeParseError, "Invalid indentation - tabs and spaces cannot be mixed", 0)
				if sc.isParserMode() {
					return sc.token(T_ERROR)
				}
			}
			sc.cursor = savedCursor
			heredocLabel.indentation = indentation
			sc.begin(yycST_END_HEREDOC)
			return sc.token(T_START_HEREDOC)
		}
	}

	sc.cursor = savedCursor
	if isHeredoc && !sc.heredocScanAhead {
		var currentState ZendLexState
		var savedDocComment *types.ZendString = CG__().doc_comment
		heredocNestingLevel := 1
		firstToken := 0
		errno := 0
		ZendSaveLexicalState(&currentState)
		sc.heredocScanAhead = true
		sc.heredocIndentation = 0
		sc.heredocIndentationUsesSpaces = false
		sc.onEvent = nil
		CG__().doc_comment = nil
		currentState.heredocLabelStack.ApplyReverse(CopyHeredocLabelStack)
		faults.ZendExceptionSave()
		for heredocNestingLevel != 0 {
			retval, _ := sc.LexScan(nil)
			if EG__().exception != nil {
				faults.ZendClearException()
				break
			}
			if firstToken == 0 {
				firstToken = retval
			}
			switch retval {
			case T_START_HEREDOC:
				heredocNestingLevel++
			case T_END_HEREDOC:
				heredocNestingLevel--
			case END:
				heredocNestingLevel = 0
			}
		}
		faults.ZendExceptionRestore()
		if b.EqualsAny(firstToken, T_VARIABLE, T_DOLLAR_OPEN_CURLY_BRACES, T_CURLY_OPEN) && sc.heredocIndentation != 0 {
			faults.ZendThrowExceptionEx(faults.ZendCeParseError, 0, "Invalid body indentation level (expecting an indentation level of at least %d)", sc.heredocIndentation)
			errno = 1
		}
		heredocLabel.indentation = sc.heredocIndentation
		heredocLabel.indentationUsesSpaces = sc.heredocIndentationUsesSpaces
		ZendRestoreLexicalState(&currentState)
		sc.heredocScanAhead = false
		CG__().increment_lineno = 0
		CG__().doc_comment = savedDocComment
		if sc.isParserMode() && errno != 0 {
			return sc.token(T_ERROR)
		}
	}
	return sc.token(T_START_HEREDOC)
}
func (sc *LangScanner) lexerRule22() (int, bool) {
	var heredocLabel = sc.heredocLabelStack.Pop()
	sc.len_ = heredocLabel.indentation + heredocLabel.Length()
	sc.cursor += sc.len_ - 1
	sc.begin(yycST_IN_SCRIPTING)
	return sc.token(T_END_HEREDOC)
}
func (sc *LangScanner) lexerRule23() (int, bool) {
	if sc.getDoubleQuotesScannedLength() != 0 {
		sc.cursor += sc.getDoubleQuotesScannedLength() - 1
		sc.setDoubleQuotesScannedLength(0)
		goto double_quotes_scan_done
	}
	if sc.isEnd() {
		return sc.token(END)
	}
	if sc.yyText0() == '\\' && sc.canRead() {
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
	if sc.setEscapeString(sc.seg(), '"') || !(sc.isParserMode()) {
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
	if sc.setEscapeString(sc.seg(), '`') || !(sc.isParserMode()) {
		return sc.tokenWithVal(T_ENCAPSED_AND_WHITESPACE)
	} else {
		return sc.token(T_ERROR)
	}
}
func (sc *LangScanner) lexerRule25() (int, bool) {
	if sc.isEnd() {
		return sc.token(END)
	}

	var heredocLabel *ZendHeredocLabel = sc.heredocLabelStack.Top()
	var newline uint = 0
	var indentation int = 0
	var spacing int = 0

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
			if sc.cursor == sc.limit {
				sc.resetLen()
				sc.handleNewlinesEx(sc.seg())
				sc.zendlval.SetNull()
				return sc.tokenWithVal(T_ENCAPSED_AND_WHITESPACE)
			}

			/* Check for ending label on the next line */
			if isLabelStart(sc.peek()) && sc.peekStrIs(heredocLabel.Label()) {
				if isLabelSuccessor(sc.yyCursor[heredocLabel.Length()]) {
					continue
				}
				if spacing == (HEREDOC_USING_SPACES | HEREDOC_USING_TABS) {
					faults.ZendThrowException(faults.ZendCeParseError, "Invalid indentation - tabs and spaces cannot be mixed", 0)
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
	sc.setStr(sc.segLen(sc.len_ - newline))
	if !(sc.heredocScanAhead) && !(EG__().exception) && sc.isParserMode() {
		newlineAtStart := b.EqualsAny(sc.yyTextN(-1), '\r', '\n')
		var copy *types.ZendString = types.Z_STR_P(sc.zendlval)
		if !(StripMultilineStringIndentation(zendlval, heredocLabel.indentation, heredocLabel.indentationUsesSpaces, newlineAtStart, newline != 0)) {
			return sc.token(T_ERROR)
		}
		if sc.setEscapeString(copy.GetStr(), 0) {
			return sc.token(T_ERROR)
		}
	} else {
		sc.handleNewlinesEx(sc.segLen(sc.len_ - newline))
	}
	return sc.tokenWithVal(T_ENCAPSED_AND_WHITESPACE)
}
func (sc *LangScanner) lexerRule26() (int, bool) {
	if sc.isEnd() {
		return sc.token(END)
	}

	var heredocLabel = sc.heredocLabelStack.Top()
	var newline uint = 0
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
			if sc.cursor == sc.limit {
				sc.resetLen()
				sc.handleNewlinesEx(sc.seg())
				sc.zendlval.SetNull()
				return sc.tokenWithVal(T_ENCAPSED_AND_WHITESPACE)
			}

			/* Check for ending label on the next line */

			if isLabelStart(sc.peek()) && sc.HasPrefix(heredocLabel.Label()) {
				if isLabelSuccessor(sc.yyCursor[heredocLabel.Length()]) {
					continue
				}
				if spacing == (HEREDOC_USING_SPACES | HEREDOC_USING_TABS) {
					faults.ZendThrowException(faults.ZendCeParseError, "Invalid indentation - tabs and spaces cannot be mixed", 0)
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
	sc.setStr(sc.segLen(sc.len_ - newline))
	if !(EG__().exception) && spacing != -1 && sc.isParserMode() {
		newlineAtStart := b.EqualsAny(sc.yyTextN(-1), '\r', '\n')
		if !(StripMultilineStringIndentation(zendlval, indentation, spacing == HEREDOC_USING_SPACES, newlineAtStart, newline != 0)) {
			return sc.token(T_ERROR)
		}
	}
	sc.handleNewlinesEx(sc.segLen(sc.len_ - newline))
	return sc.tokenWithVal(T_ENCAPSED_AND_WHITESPACE)
}
func (sc *LangScanner) lexerRule27() (int, bool) {
	if sc.isEnd() {
		return sc.token(END)
	}
	if !sc.heredocScanAhead {
		faults.ZendError(faults.E_COMPILE_WARNING, "Unexpected character in input:  '%c' (ASCII=%d) state=%d", sc.yyText0(), sc.yyText0(), sc.state)
	}
	if sc.isParserMode() {
		return 0, true
	} else {
		return sc.token(T_BAD_CHARACTER)
	}
}
