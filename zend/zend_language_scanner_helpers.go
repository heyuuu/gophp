package zend

import (
	b "sik/builtin"
	"strings"
)

type LangScanner struct {
	code string // 代码原文

	lineno    uint32 // CG__().zend_lineno
	startLine uint32
	shortTags bool // CG__().short_tags

	zendlval *Zval                // argument zendlval
	elem     *ZendParserStackElem // argument: elem

	scannedStringLen uint // LANG_SCNG__().scanned_string_len

	len_       uint         // LANG_SCNG__().yy_leng
	start      uint         // LANG_SCNG__().yy_start *byte
	text       uint         // LANG_SCNG__().yy_text *byte
	cursor     uint         // LANG_SCNG__().yy_cursor *byte
	marker     uint         // LANG_SCNG__().yy_marker *byte
	limit      uint         // LANG_SCNG__().yy_limit *byte
	state      int          // LANG_SCNG__().yy_state
	stateStack b.Stack[int] // LANG_SCNG__().stateStack

	yyText   *byte // LANG_SCNG__().yy_text
	yyCursor *byte // LANG_SCNG__().yy_cursor

	heredocLabelStack            b.Stack[*ZendHeredocLabel] // LANG_SCNG__().heredocLabelStack
	heredocScanAhead             bool                       // LANG_SCNG__().heredoc_scan_ahead
	heredocIndentation           int                        // LANG_SCNG__().heredoc_indentation
	heredocIndentationUsesSpaces bool                       // LANG_SCNG__().heredoc_indentation_uses_spaces

	onEvent    func(event ZendPhpScannerEvent, token int, line int, context any) // LANG_SCNG__().on_event
	onEventCxt any                                                               // LANG_SCNG__().on_event_context

	inputFilter  func(string) string // LANG_SCNG__().input_filter  函数参数类型有差异
	outputFilter func(string) string // LANG_SCNG__().output_filter 函数参数类型有差异

	docComment *string // CG__().doc_comment *ZendString
}

func (sc *LangScanner) LexScan(elem *ZendParserStackElem) (int, *Zval) {
	var zv Zval
	zv.SetUndef()
	sc.zendlval = &zv
	token := sc.LexScanEx(&zv, elem)
	return token, &zv
}

func (sc *LangScanner) LexScanEx(zendlval *Zval, elem *ZendParserStackElem) int {
	sc.limit = uint(len(sc.code)) - 1

	// 初始化
	sc.zendlval = zendlval
	sc.elem = elem
	sc.startLine = sc.lineno
	sc.zendlval.SetUndef()

	// 第一次执行
	sc.text = sc.cursor
	token, restart := _lexScan(zendlval, elem, sc)

	// 重新执行，对标原 goto restart
	for restart {
		sc.text = sc.cursor
		token, restart = _lexScan(zendlval, elem, sc)
	}

	return token
}

func (sc *LangScanner) charAt(pos int) byte {
	if pos >= 0 && pos < len(sc.code) {
		return sc.code[pos]
	}
	return 0
}
func (sc *LangScanner) textChar(offset int) byte { return sc.charAt(int(sc.text) + offset) }

func (sc *LangScanner) isEnd() bool   { return sc.cursor > sc.limit }
func (sc *LangScanner) canRead() bool { return sc.cursor < sc.limit }
func (sc *LangScanner) read() byte {
	c := sc.code[sc.cursor]
	sc.cursor++
	return c
}
func (sc *LangScanner) readStr(l uint) string {
	s := sc.peekStr(l)
	sc.cursor += uint(len(s))
	return s
}

func (sc *LangScanner) peek() byte {
	return sc.code[sc.cursor]
}
func (sc *LangScanner) peekStr(l uint) string {
	// 确认读取位置
	if sc.cursor > sc.limit {
		return ""
	}

	// 确认结束读取位置
	end := sc.cursor + l
	if end > sc.limit {
		end = sc.limit
	}

	return sc.code[sc.cursor:end]
}
func (sc *LangScanner) peekIs(bytes ...byte) bool {
	c := sc.peek()
	for _, b := range bytes {
		if c == b {
			return true
		}
	}
	return false
}
func (sc *LangScanner) peekStrIs(str string) bool {
	return sc.peekStr(uint(len(str))) == str
}
func (sc *LangScanner) peekStrIsIgnoreCase(str string) bool {
	return strings.EqualFold(sc.peekStr(uint(len(str))), str)
}

func (sc *LangScanner) peekOffset(offset int) byte {
	return sc.code[int(sc.cursor)+offset]
}

func (sc *LangScanner) skip()         { sc.cursor++ }
func (sc *LangScanner) fallback()     { sc.cursor-- }
func (sc *LangScanner) backup()       { sc.cursor = sc.marker }
func (sc *LangScanner) restore()      { sc.marker = sc.cursor }
func (sc *LangScanner) yyfill()       {}
func (sc *LangScanner) getState() int { return sc.state }

func (sc *LangScanner) HasPrefix(str string) bool {
	return sc.peekStr(uint(len(str))) == str
}

// 读取剩余字符串
func (sc *LangScanner) tail() string {
	if sc.cursor > sc.limit {
		return ""
	}
	return sc.code[sc.cursor:]
}

func (sc *LangScanner) resetLen() {
	sc.len_ = sc.cursor - sc.text
}

func (sc *LangScanner) setLen(len_ uint) {
	sc.cursor = sc.text + len_
	sc.len_ = len_
}

func (sc *LangScanner) matched() string {
	return sc.code[sc.text:sc.cursor]
}

func (sc *LangScanner) seg() string {
	return sc.code[sc.text : sc.text+sc.len_]
}

func (sc *LangScanner) segLen(len_ uint) string {
	return sc.code[sc.text : sc.text+len_]
}

func (sc *LangScanner) setStr(str string) {
	sc.zendlval.SetRawString(str)
}

func (sc *LangScanner) setStrFiltered(str string) {
	if sc.outputFilter != nil {
		str = sc.outputFilter(str)
	}
	sc.zendlval.SetRawString(str)
}

func (sc *LangScanner) isParserMode() bool {
	return sc.elem != nil
}

func (sc *LangScanner) emitTokenEvent(token int) {
	if sc.onEvent != nil {
		sc.onEvent(ON_TOKEN, token, sc.startLine, sc.onEventCxt)
	}
}

func (sc *LangScanner) token(token int) (int, bool) {
	sc.emitTokenEvent(token)
	return token, false
}

func (sc *LangScanner) tokenWithVal(token int) (int, bool) {
	if sc.isParserMode() {
		ZEND_ASSERT(!sc.zendlval.IsUndef())
		sc.elem.ast = ZendAstCreateZvalWithLineno(sc.zendlval, uint32(sc.startLine))
	}
	return sc.token(token)
}

func (sc *LangScanner) tokenWithStr(token int, offset uint) (int, bool) {
	var str = sc.seg()[offset:]
	sc.setStrFiltered(str)
	return sc.tokenWithVal(token)
}

func (sc *LangScanner) returnOrSkipToken(token int) (int, bool) {
	if sc.isParserMode() {
		sc.emitTokenEvent(token)
		sc.startLine = sc.lineno
		return 0, true
	}
	return sc.token(token)
}

func (sc *LangScanner) begin(state int) {
	sc.state = state
}

func (sc *LangScanner) pushState(state int) {
	sc.stateStack.Push(sc.state)
	sc.state = state
}

func (sc *LangScanner) popState() {
	sc.state = sc.stateStack.Pop()
}

func (sc *LangScanner) tryPopState() {
	if !sc.stateStack.IsEmpty() {
		sc.state = sc.stateStack.Pop()
	}
}

func (sc *LangScanner) isStateStackEmpty() bool {
	return sc.stateStack.IsEmpty()
}

func (sc *LangScanner) yyText0() byte           { return sc.code[sc.text] }
func (sc *LangScanner) yyTextN(offset int) byte { return sc.code[int(sc.text)+offset] }

func (sc *LangScanner) handleNewlinesEx(str string) {
	l := len(str)
	for i, c := range str {
		if c == '\n' || (c == '\r' && i+1 < l && str[i+1] != '\n') {
			sc.lineno++
		}
	}
}

func (sc *LangScanner) handleNewlines(s *byte, l uint) {
	var p *byte = s
	var boundary *byte = p + l
	for p < boundary {
		if (*p) == '\n' || (*p) == '\r' && (*(p + 1)) != '\n' {
			CG__().zend_lineno++
		}
		p++
	}
}

func (sc *LangScanner) handleNewline(c byte) {
	if c == '\n' || c == '\r' {
		sc.lineno++
	}
}
func (sc *LangScanner) setDoubleQuotesScannedLength(len_ uint) uint {
	sc.scannedStringLen = len_
	return sc.scannedStringLen
}

func (sc *LangScanner) getDoubleQuotesScannedLength() uint {
	return sc.scannedStringLen
}

func (sc *LangScanner) saveLexState() *ZendLexState {
	lexState := ZendLexState{}

	lexState.lineno = sc.lineno

	lexState.len_ = sc.len_
	lexState.start = sc.start
	lexState.text = sc.text
	lexState.cursor = sc.cursor
	lexState.marker = sc.marker
	lexState.limit = sc.limit
	lexState.state = sc.state
	lexState.stateStack = sc.stateStack.Copy()

	lexState.heredocLabelStack = sc.heredocLabelStack.Copy()

	lexState.on_event = sc.onEvent
	lexState.on_event_context = sc.onEventCxt

	lexState.input_filter = sc.inputFilter
	lexState.output_filter = sc.outputFilter

	lexState.in = LANG_SCNG__().yy_in
	lexState.filename = ZendGetCompiledFilename()
	lexState.script_org = LANG_SCNG__().script_org
	lexState.script_org_size = LANG_SCNG__().script_org_size
	lexState.script_filtered = LANG_SCNG__().script_filtered
	lexState.script_filtered_size = LANG_SCNG__().script_filtered_size
	lexState.ast = CG__().ast
	lexState.ast_arena = CG__().ast_arena

	sc.stateStack.Clean()
	sc.heredocLabelStack.Clean()

	return &lexState
}

func (sc *LangScanner) restoreLexState(lexState *ZendLexState) {
	sc.lineno = lexState.lineno

	sc.len_ = lexState.len_
	sc.start = lexState.start
	sc.text = lexState.text
	sc.cursor = lexState.cursor
	sc.marker = lexState.marker
	sc.limit = lexState.limit
	sc.state = lexState.state
	sc.stateStack = lexState.stateStack
	sc.heredocLabelStack = lexState.heredocLabelStack

	sc.onEvent = lexState.on_event
	sc.onEventCxt = lexState.on_event_context

	sc.inputFilter = lexState.input_filter
	sc.outputFilter = lexState.output_filter

	LANG_SCNG__().yy_in = lexState.in
	sc.lineno = lexState.lineno
	ZendRestoreCompiledFilename(lexState.filename)
	if LANG_SCNG__().script_filtered {
		efree(LANG_SCNG__().script_filtered)
		LANG_SCNG__().script_filtered = nil
	}
	LANG_SCNG__().script_org = lexState.script_org
	LANG_SCNG__().script_org_size = lexState.script_org_size
	LANG_SCNG__().script_filtered = lexState.script_filtered
	LANG_SCNG__().script_filtered_size = lexState.script_filtered_size
	CG__().ast = lexState.ast
	CG__().ast_arena = lexState.ast_arena
	sc.resetDocComment()
}

func (sc *LangScanner) setDocComment(docComment string) {
	sc.docComment = &docComment
}

func (sc *LangScanner) resetDocComment() {
	sc.docComment = nil
}
