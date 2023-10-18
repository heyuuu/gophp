package zend

import (
	"fmt"
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
	"strconv"
	"strings"
)

const YYCURSOR *uint8 = LANG_SCNG__().yy_cursor
const YYLIMIT *uint8 = LANG_SCNG__().yy_limit

func YYSETCONDITION(s int) __auto__ {
	LANG_SCNG__().yy_state = s
	return LANG_SCNG__().yy_state
}

func BEGIN(state __auto__) __auto__ { return YYSETCONDITION(yycstate) }

var LanguageScannerGlobals ZendPhpScannerGlobals

/* To save initial string length after scanning to first variable */

func isLabelStart(c uint8) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c == '_' || c >= 0x80
}
func isLabelSuccessor(c uint8) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9' || c == '_' || c >= 0x80
}
func zendIsOct(c byte) bool { return c >= '0' && c <= '7' }
func zendIsHex(c byte) bool {
	return c >= '0' && c <= '9' || c >= 'a' && c <= 'f' || c >= 'A' && c <= 'F'
}

func YyScanBuffer(str *byte, len_ int) {
	YYCURSOR = (*uint8)(str)
	YYLIMIT = YYCURSOR + len_
	if !(LANG_SCNG__().yy_start) {
		LANG_SCNG__().yy_start = YYCURSOR
	}
}
func StartupScanner() {
	var sc *LangScanner
	CG__().parse_error = 0
	CG__().doc_comment = nil
	CG__().extra_fn_flags = 0
	sc.stateStack.Clean()
	sc.heredocLabelStack.Clean()
	sc.heredocScanAhead = false
}
func ShutdownScanner() {
	var sc *LangScanner
	CG__().parse_error = 0
	sc.docComment = nil
	sc.stateStack.Clean()
	sc.heredocLabelStack.Clean()
	sc.heredocScanAhead = false
	sc.onEvent = nil
}
func ZendSaveLexicalState(lexState *ZendLexState) {
	var sc *LangScanner
	*lexState = *sc.saveLexState()
}
func ZendRestoreLexicalState(lexState *ZendLexState) {
	var sc *LangScanner
	sc.restoreLexState(lexState)
}
func ZendDestroyFileHandle(fileHandle *FileHandle) {
	CG__().DelOpenFile(fileHandle)

	fileHandle.openedPath = ""
	fileHandle.filename = ""
}

func OpenFileForScanning(fileHandle *FileHandle) int {
	buf, ok := fileHandle.Fixup()
	if !ok {
		/* Still add it to open_files to make destroy_file_handle work */
		CG__().AddOpenFile(fileHandle)
		return types.FAILURE
	}
	size := len(buf)

	b.Assert(!(EG__().exception) && "stream_fixup() should have failed")

	CG__().AddOpenFile(fileHandle)

	/* Reset the scanner for scanning the new file */
	sc := NewLangScanner(buf)
	if CG__().skip_shebang {
		CG__().skip_shebang = false
		sc.begin(yycSHEBANG)
	} else {
		sc.begin(yycINITIAL)
	}

	var compiled_filename string
	if len(fileHandle.openedPath) != 0 {
		compiled_filename = fileHandle.openedPath
	} else {
		compiled_filename = fileHandle.filename
	}
	ZendSetCompiledFilename(compiled_filename)
	return types.SUCCESS
}
func ZendCompile(type_ int) *types.ZendOpArray {
	// backup
	var originalInCompilation = CG__().in_compilation

	// reset
	CG__().in_compilation = 1
	CG__().ast = nil

	var opArray *types.ZendOpArray = nil
	if !(Zendparse()) {
		var last_lineno int = CG__().zend_lineno
		var originalFileContext ZendFileContext
		var originalOparrayContext ZendOparrayContext
		var originalActiveOpArray = CG__().GetActiveOpArray()
		opArray = InitOpArrayEx()
		CG__().active_op_array = opArray

		compiler := CurrCompiler()

		/* Use heap to not waste arena memory */
		opArray.AddFnFlags(types.AccHeapRtCache)
		originalFileContext = CG__().FileContextBegin()
		ZendOparrayContextBegin(&originalOparrayContext)
		compiler.CompileTopStmt(CG__().ast)
		CG__().zend_lineno = last_lineno
		ZendEmitFinalReturn(type_ == ZEND_USER_FUNCTION)
		opArray.line_start = 1
		opArray.line_end = last_lineno
		compiler.PassTwo(opArray)
		ZendOparrayContextEnd(&originalOparrayContext)
		CG__().FileContextEnd(originalFileContext)
		CG__().active_op_array = originalActiveOpArray
	}
	// ZendAstDestroy(CG__().ast)

	// restore
	CG__().in_compilation = originalInCompilation

	return opArray
}
func CompileFilename(type_ int, filename string) *types.ZendOpArray {
	fh := NewFileHandleByFilename(filename)
	opArray := CompileFile(fh, type_)
	if opArray != nil && fh.GetStream().GetHandle() != nil {
		if fh.GetOpenedPath() == "" {
			fh.SetOpenedPath(filename)
		}
		EG__().AddIncludedFile(fh.GetOpenedPath())
	}
	ZendDestroyFileHandle(fh)
	return opArray
}
func ZendPrepareStringForScanning(str *types.Zval, filename string) int {
	/* enforce ZEND_MMAP_AHEAD trailing NULLs for flex... */
	buf := str.String() + strings.Repeat("\x00", ZEND_MMAP_AHEAD)
	size := len(str.String())

	LANG_SCNG__().yy_start = nil
	YyScanBuffer(buf, size)
	ZendSetCompiledFilename(filename)
	CG__().zend_lineno = 1
	CG__().increment_lineno = 0
	CG__().doc_comment = nil
	return types.SUCCESS
}
func CompileString(source_string *types.Zval, filename *byte) *types.ZendOpArray {
	var original_lex_state ZendLexState
	var op_array int = nil
	var tmp types.Zval
	tmp.SetString(operators.ZvalGetStrVal(source_string))
	iflen(tmp.String()) == 0
	{
		return nil
	}
	ZendSaveLexicalState(&original_lex_state)
	if ZendPrepareStringForScanning(&tmp, filename) == types.SUCCESS {
		BEGIN(ST_IN_SCRIPTING)
		op_array = ZendCompile(ZEND_EVAL_CODE)
	}
	ZendRestoreLexicalState(&original_lex_state)
	zval_ptr_dtor(&tmp)
	return op_array
}
func HighlightFile(filename *byte, syntax_highlighter_ini *zend_syntax_highlighter_ini) int {
	var original_lex_state ZendLexState
	var file_handle FileHandle
	zend_stream_init_filename(&file_handle, filename)
	ZendSaveLexicalState(&original_lex_state)
	if OpenFileForScanning(&file_handle) == types.FAILURE {
		zend_message_dispatcher(ZMSG_FAILED_HIGHLIGHT_FOPEN, filename)
		ZendRestoreLexicalState(&original_lex_state)
		return types.FAILURE
	}
	zend_highlight(syntax_highlighter_ini)
	if LANG_SCNG__().script_filtered {
		efree(LANG_SCNG__().script_filtered)
		LANG_SCNG__().script_filtered = nil
	}
	ZendDestroyFileHandle(&file_handle)
	ZendRestoreLexicalState(&original_lex_state)
	return types.SUCCESS
}
func HighlightString(str *types.Zval, syntaxHighlighterIni *ZendSyntaxHighlighterIni, strName string) int {
	var original_lex_state ZendLexState
	var tmp types.Zval
	if Z_TYPE_P(str) != types.IsString {
		str = types.NewZvalString(operators.ZvalGetStrVal(str))
	}
	ZendSaveLexicalState(&original_lex_state)
	if ZendPrepareStringForScanning(str, strName) == types.FAILURE {
		ZendRestoreLexicalState(&original_lex_state)
		if str == &tmp {
			zval_ptr_dtor(&tmp)
		}
		return types.FAILURE
	}
	BEGIN(INITIAL)
	zend_highlight(syntaxHighlighterIni)
	if LANG_SCNG__().script_filtered {
		efree(LANG_SCNG__().script_filtered)
		LANG_SCNG__().script_filtered = nil
	}
	ZendRestoreLexicalState(&original_lex_state)
	if str == &tmp {
		zval_ptr_dtor(&tmp)
	}
	return types.SUCCESS
}

func (sc *LangScanner) setEscapeString(str string, quoteType byte) bool {
	len_ := len(str)
	if len_ <= 1 || strings.IndexByte(str, '\\') >= 0 {
		// 无转义处理直接返回
		sc.setStrFiltered(str)
		return true
	}

	/* convert escape sequences */
	buf := strings.Builder{}
	for i := 0; i < len_; i++ {
		// 非转义字符直接计入结果
		if str[i] != '\\' || i == len_-1 {
			buf.WriteByte(str[i])
			continue
		}
		// 处理转义
		i++
		switch str[i] {
		case 'n':
			buf.WriteByte('\n')
		case 'r':
			buf.WriteByte('\r')
		case 't':
			buf.WriteByte('\t')
		case 'f':
			buf.WriteByte('\f')
		case 'v':
			buf.WriteByte('\v')
		case 'e':
			buf.WriteByte('\033') // golang 不支持 \e 转义符
		case quoteType, '\\', '$':
			buf.WriteByte(str[i])
		case 'x', 'X':
			if i+1 < len_ && zendIsHex(str[i+1]) {
				// 十六形式的字符 (e.g. 0x12)
				i++
				val := str[i]
				if i+1 < len_ && zendIsHex(str[i+1]) {
					i++
					val = val*16 + str[i+1]
				}
				buf.WriteByte(val)
			} else {
				buf.WriteByte('\\')
				buf.WriteByte(str[i])
			}
		case 'u':
			// 跳过非 \u{xxxx} 形式的值
			if i >= len_ || str[i+1] != '{' {
				/* we silently let this pass to avoid breaking code
				 * with JSON in string literals (e.g. "\"\u202e\""
				 */
				buf.WriteByte('\\')
				buf.WriteByte('u')
				break
			}

			/* \u{xxxx} 形式的值 */
			i += 2
			start := i
			valid := true
			for i < len_ && str[i] != '}' {
				if !zendIsHex(str[i]) {
					valid = false
					break
				}
				i++
			}
			if i == len_ {
				valid = false
			}

			/* \u{} is invalid */
			if i == start+1 {
				valid = false
			}
			if !valid { // 没找到或为 ${} 形式
				faults.ThrowException(faults.ZendCeParseError, "Invalid UTF-8 codepoint escape sequence", 0)
				sc.zendlval.SetUndef()
				return false
			}

			codepoint, _ := strconv.ParseUint(str[start:i], 16, 64)
			/* based on https://en.wikipedia.org/wiki/UTF-8#Sample_code */
			if codepoint < 0x80 {
				buf.WriteByte(uint8(codepoint))
			} else if codepoint <= 0x7ff {
				buf.WriteByte(uint8(codepoint>>6 + 0xc0))
				buf.WriteByte(uint8(codepoint&0x3f + 0x80))
			} else if codepoint <= 0xffff {
				buf.WriteByte(uint8(codepoint>>12 + 0xe0))
				buf.WriteByte(uint8(codepoint>>6&0x3f + 0x80))
				buf.WriteByte(uint8(codepoint&0x3f + 0x80))
			} else if codepoint <= 0x10ffff {
				buf.WriteByte(uint8(codepoint>>18 + 0xf0))
				buf.WriteByte(uint8(codepoint>>12&0x3f + 0x80))
				buf.WriteByte(uint8(codepoint>>6&0x3f + 0x80))
				buf.WriteByte(uint8(codepoint&0x3f + 0x80))
			} else {
				/* per RFC 3629, UTF-8 can only represent 21 bits */
				faults.ThrowException(faults.ZendCeParseError, "Invalid UTF-8 codepoint escape sequence: Codepoint too large", 0)
				sc.zendlval.SetUndef()
				return false
			}
		default:
			/* check for an octal */
			if zendIsOct(str[i]) {
				start := i
				var octal uint = uint(str[i])
				if i+1 < len_ && zendIsOct(str[i+1]) {
					i++
					octal = octal*8 + octal
					if i+1 < len_ && zendIsOct(str[i+1]) {
						i++
						octal = octal*8 + octal
					}
				}
				if octal > 0377 && !sc.heredocScanAhead {
					/* 3 octit values must not overflow 0xFF (\377) */
					faults.Error(faults.E_COMPILE_WARNING, fmt.Sprintf("Octal escape sequence overflow \\%s is greater than \\377", str[start:i+1]))
				}
				buf.WriteByte(byte(octal))
			} else {
				buf.WriteByte('\\')
				buf.WriteByte(str[i])
			}
		}
	}

	sc.setStrFiltered(buf.String())
	return true
}

const HEREDOC_USING_SPACES = 1
const HEREDOC_USING_TABS = 2

func nextNewLine(str string, start int) (int, int) {
	for i := start; i < len(str); i++ {
		if str[i] == '\r' {
			if i+1 < len(str) && str[i+1] == '\n' {
				return i, 2
			} else {
				return i, 1
			}
		} else if str[i] == '\n' {
			return i, 1
		}
	}
	return -1, 0
}

func StripMultilineStringIndentation(zendlval *types.Zval, indentation int, usingSpaces bool, newlineAtStart bool, newlineAtEnd bool) bool {
	if str, ok := StripMultilineStringIndentationEx(zendlval.String(), indentation, usingSpaces, newlineAtStart, newlineAtEnd); ok {
		zendlval.SetString(str)
		return true
	} else {
		zendlval.SetUndef()
		return false
	}
}

func splitLines(s string) []string {
	length := len(s)
	lineStart := 0
	var lines []string
	for i := 0; i < length; i++ {
		if s[i] == '\r' || s[i] == '\n' {
			if i+1 < len(s) && s[i] == '\r' && s[i+1] == '\n' {
				i++
			}

			lines = append(lines, s[lineStart:i+1])
			lineStart = i + 1
		}
	}
	if lineStart < length {
		lines = append(lines, s[lineStart:])
	}
	return lines
}

func StripMultilineStringIndentationEx(s string, indentation int, usingSpaces bool, newlineAtStart bool, newlineAtEnd bool) (string, bool) {
	var newlineCount int = 0
	var buf strings.Builder

	lines := splitLines(s)
	if !newlineAtStart {
		if len(lines) <= 1 {
			return s, true
		}
		lines = lines[1:]
		newlineCount++
	}

	/* <= intentional */
	for _, line := range lines {
		/* Try to skip indentation */
		for skip := 0; skip < indentation; skip++ {
			if line == "" || line[0] == '\r' || line[0] == '\n' {
				/* Don't require full indentation on whitespace-only lines */
				break
			}
			if line[0] != ' ' && line[0] != '\t' {
				CG__().zend_lineno += newlineCount
				faults.ThrowException(faults.ZendCeParseError, fmt.Sprintf("Invalid body indentation level (expecting an indentation level of at least %d)", indentation), 0)
				return "", false
			}
			if !usingSpaces && s[0] == ' ' || usingSpaces && s[0] == '\t' {
				CG__().zend_lineno += newlineCount
				faults.ThrowException(faults.ZendCeParseError, "Invalid indentation - tabs and spaces cannot be mixed", 0)
				return "", false
			}
			line = line[1:]
		}
		buf.WriteString(line)
		newlineCount++
	}
	return buf.String(), true
}

func CopyHeredocLabelStack(heredocLabel *ZendHeredocLabel) {
	newHeredocLabel := heredocLabel.Copy()
	LANG_SCNG__().heredoc_label_stack.Push(newHeredocLabel)
}
