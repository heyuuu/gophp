package zend

import (
	b "sik/builtin"
	r "sik/runtime"
)

type ctype = byte
type LangScanner struct {
	lineno           int // CG__().zend_lineno
	scannedStringLen int // LANG_SCNG__().scanned_string_len

	yyLen    uint  // LANG_SCNG__().yy_leng
	yyStart  *byte // LANG_SCNG__().yy_start
	yyText   *byte // LANG_SCNG__().yy_text
	yyCursor *byte // LANG_SCNG__().yy_cursor
	yyMarker *byte // LANG_SCNG__().yy_marker
	yyLimit  *byte // LANG_SCNG__().yy_limit
	yyState  int   // LANG_SCNG__().yy_state
}

func (sc *LangScanner) handleNewlines(l *byte) {
	var p = sc.yyText
	var boundary *byte = p + l
	for p < boundary {
		if (*p) == '\n' || (*p) == '\r' && (*(p + 1)) != '\n' {
			sc.lineno++
		}
		p++
	}
	LANG_SCNG__()
}
func (sc *LangScanner) handleNewline(c byte) {
	if c == '\n' || c == '\r' {
		sc.lineno++
	}
}

func (sc *LangScanner) SetDoubleQuotesScannedLength(len_ int) int {
	sc.scannedStringLen = len_
	return sc.scannedStringLen
}
func (sc *LangScanner) GetDoubleQuotesScannedLength() int {
	return sc.scannedStringLen
}
func (sc *LangScanner) LabelStart(c ctype) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c == '_' || c >= 0x80
}
func (sc *LangScanner) IsLabelSuccessor(c ctype) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9' || c == '_' || c >= 0x80
}
func (sc *LangScanner) IsOct(c ctype) bool { return c >= '0' && c <= '7' }
func (sc *LangScanner) IsHex(c ctype) bool {
	return c >= '0' && c <= '9' || c >= 'a' && c <= 'f' || c >= 'A' && c <= 'F'
}
func (sc *LangScanner) StripUnderscores(str *byte, len_ *int) {
	var src *byte = str
	var dest *byte = str
	for (*src) != '0' {
		if (*src) != '_' {
			*dest = *src
			dest++
		} else {
			*len_--
		}
		src++
	}
	*dest = '0'
}
func EncodingFilterScriptToInternal(to **uint8, to_length *int, from *uint8, from_length int) int {
	var internal_encoding *ZendEncoding = ZendMultibyteGetInternalEncoding()
	ZEND_ASSERT(internal_encoding != nil)
	return ZendMultibyteEncodingConverter(to, to_length, from, from_length, internal_encoding, LANG_SCNG__().script_encoding)
}
func EncodingFilterScriptToIntermediate(to **uint8, to_length *int, from *uint8, from_length int) int {
	return ZendMultibyteEncodingConverter(to, to_length, from, from_length, ZendMultibyteEncodingUtf8, LANG_SCNG__().script_encoding)
}
func EncodingFilterIntermediateToScript(to **uint8, to_length *int, from *uint8, from_length int) int {
	return ZendMultibyteEncodingConverter(to, to_length, from, from_length, LANG_SCNG__().script_encoding, ZendMultibyteEncodingUtf8)
}
func EncodingFilterIntermediateToInternal(to **uint8, to_length *int, from *uint8, from_length int) int {
	var internal_encoding *ZendEncoding = ZendMultibyteGetInternalEncoding()
	ZEND_ASSERT(internal_encoding != nil)
	return ZendMultibyteEncodingConverter(to, to_length, from, from_length, internal_encoding, ZendMultibyteEncodingUtf8)
}

// #define yy_push_state(state_and_tsrm) _yy_push_state ( yyc ## state_and_tsrm )

func StartupScanner() {
	CG__().SetParseError(0)
	CG__().SetDocComment(nil)
	CG__().SetExtraFnFlags(0)
	ZendStackInit(&(LANG_SCNG__().state_stack), b.SizeOf("int"))
	ZendPtrStackInit(&(LANG_SCNG__().heredoc_label_stack))
	LANG_SCNG__().heredoc_scan_ahead = 0
}
func HeredocLabelDtor(heredoc_label *ZendHeredocLabel) { Efree(heredoc_label.GetLabel()) }
func ShutdownScanner() {
	CG__().SetParseError(0)
	RESET_DOC_COMMENT()
	ZendStackDestroy(&(LANG_SCNG__().state_stack))
	ZendPtrStackClean(&(LANG_SCNG__().heredoc_label_stack), (func(any))(&HeredocLabelDtor), 1)
	ZendPtrStackDestroy(&(LANG_SCNG__().heredoc_label_stack))
	LANG_SCNG__().heredoc_scan_ahead = 0
	LANG_SCNG__().on_event = nil
}
func ZendSaveLexicalState(lex_state *ZendLexState) {
	lex_state.SetYyLeng(LANG_SCNG__().yy_leng)
	lex_state.SetYyStart(LANG_SCNG__().yy_start)
	lex_state.SetYyText(LANG_SCNG__().yy_text)
	lex_state.SetYyCursor(LANG_SCNG__().yy_cursor)
	lex_state.SetYyMarker(LANG_SCNG__().yy_marker)
	lex_state.SetYyLimit(LANG_SCNG__().yy_limit)
	lex_state.SetStateStack(LANG_SCNG__().state_stack)
	ZendStackInit(&(LANG_SCNG__().state_stack), b.SizeOf("int"))
	lex_state.SetHeredocLabelStack(LANG_SCNG__().heredoc_label_stack)
	ZendPtrStackInit(&(LANG_SCNG__().heredoc_label_stack))
	lex_state.SetIn(LANG_SCNG__().yy_in)
	lex_state.SetYyState(YYSTATE)
	lex_state.SetFilename(ZendGetCompiledFilename())
	lex_state.SetLineno(CG__().GetZendLineno())
	lex_state.SetScriptOrg(LANG_SCNG__().script_org)
	lex_state.SetScriptOrgSize(LANG_SCNG__().script_org_size)
	lex_state.SetScriptFiltered(LANG_SCNG__().script_filtered)
	lex_state.SetScriptFilteredSize(LANG_SCNG__().script_filtered_size)
	lex_state.SetInputFilter(LANG_SCNG__().input_filter)
	lex_state.SetOutputFilter(LANG_SCNG__().output_filter)
	lex_state.SetScriptEncoding(LANG_SCNG__().script_encoding)
	lex_state.SetOnEvent(LANG_SCNG__().on_event)
	lex_state.SetOnEventContext(LANG_SCNG__().on_event_context)
	lex_state.SetAst(CG__().GetAst())
	lex_state.SetAstArena(CG__().GetAstArena())
}
func ZendRestoreLexicalState(lex_state *ZendLexState) {
	LANG_SCNG__().yy_leng = lex_state.GetYyLeng()
	LANG_SCNG__().yy_start = lex_state.GetYyStart()
	LANG_SCNG__().yy_text = lex_state.GetYyText()
	LANG_SCNG__().yy_cursor = lex_state.GetYyCursor()
	LANG_SCNG__().yy_marker = lex_state.GetYyMarker()
	LANG_SCNG__().yy_limit = lex_state.GetYyLimit()
	ZendStackDestroy(&(LANG_SCNG__().state_stack))
	LANG_SCNG__().state_stack = lex_state.GetStateStack()
	ZendPtrStackClean(&(LANG_SCNG__().heredoc_label_stack), (func(any))(&HeredocLabelDtor), 1)
	ZendPtrStackDestroy(&(LANG_SCNG__().heredoc_label_stack))
	LANG_SCNG__().heredoc_label_stack = lex_state.GetHeredocLabelStack()
	LANG_SCNG__().yy_in = lex_state.GetIn()
	YYSETCONDITION(lex_state.GetYyState())
	CG__().SetZendLineno(lex_state.GetLineno())
	ZendRestoreCompiledFilename(lex_state.GetFilename())
	if LANG_SCNG__().script_filtered {
		Efree(LANG_SCNG__().script_filtered)
		LANG_SCNG__().script_filtered = nil
	}
	LANG_SCNG__().script_org = lex_state.GetScriptOrg()
	LANG_SCNG__().script_org_size = lex_state.GetScriptOrgSize()
	LANG_SCNG__().script_filtered = lex_state.GetScriptFiltered()
	LANG_SCNG__().script_filtered_size = lex_state.GetScriptFilteredSize()
	LANG_SCNG__().input_filter = lex_state.GetInputFilter()
	LANG_SCNG__().output_filter = lex_state.GetOutputFilter()
	LANG_SCNG__().script_encoding = lex_state.GetScriptEncoding()
	LANG_SCNG__().on_event = lex_state.GetOnEvent()
	LANG_SCNG__().on_event_context = lex_state.GetOnEventContext()
	CG__().SetAst(lex_state.GetAst())
	CG__().SetAstArena(lex_state.GetAstArena())
	RESET_DOC_COMMENT()
}
func ZendDestroyFileHandle(file_handle *ZendFileHandle) {
	ZendLlistDelElement(CG__().GetOpenFiles(), file_handle, (func(any, any) int)(ZendCompareFileHandles))

	/* zend_file_handle_dtor() operates on the copy, so we have to NULLify the original here */

	file_handle.SetOpenedPath(nil)
	if file_handle.GetFreeFilename() != 0 {
		file_handle.SetFilename(nil)
	}
}
func ZendLexTstring(zv *Zval) {
	if LANG_SCNG__().on_event {
		LANG_SCNG__().on_event(ON_FEEDBACK, T_STRING, 0, LANG_SCNG__().on_event_context)
	}
	ZVAL_STRINGL(zv, (*byte)(LANG_SCNG__().yy_text), LANG_SCNG__().yy_leng)
}

const BOM_UTF32_BE = "x00x00xfexff"
const BOM_UTF32_LE = "xffxfex00x00"
const BOM_UTF16_BE = "xfexff"
const BOM_UTF16_LE = "xffxfe"
const BOM_UTF8 = "xefxbbxbf"

func ZendMultibyteDetectUtfEncoding(script *uint8, script_size int) *ZendEncoding {
	var p *uint8
	var wchar_size int = 2
	var le int = 0

	/* utf-16 or utf-32? */

	p = script
	r.Assert(p >= script)
	for size_t(p-script) < script_size {
		p = memchr(p, 0, script_size-(p-script)-2)
		if p == nil {
			break
		}
		if (*(p + 1)) == '0' && (*(p + 2)) == '0' {
			wchar_size = 4
			break
		}

		/* searching for UTF-32 specific byte orders, so this will do */
		p += 4
	}

	/* BE or LE? */

	p = script
	r.Assert(p >= script)
	for size_t(p-script) < script_size {
		if (*p) == '0' && (*(p + wchar_size - 1)) != '0' {

			/* BE */

			le = 0
			break
		} else if (*p) != '0' && (*(p + wchar_size - 1)) == '0' {

			/* LE* */

			le = 1
			break
		}
		p += wchar_size
	}
	if wchar_size == 2 {
		if le != 0 {
			return ZendMultibyteEncodingUtf16le
		} else {
			return ZendMultibyteEncodingUtf16be
		}
	} else {
		if le != 0 {
			return ZendMultibyteEncodingUtf32le
		} else {
			return ZendMultibyteEncodingUtf32be
		}
	}
	return nil
}
func ZendMultibyteDetectUnicode() *ZendEncoding {
	var script_encoding *ZendEncoding = nil
	var bom_size int
	var pos1 *uint8
	var pos2 *uint8
	if LANG_SCNG__().script_org_size < b.SizeOf("BOM_UTF32_LE")-1 {
		return nil
	}

	/* check out BOM */

	if !(memcmp(LANG_SCNG__().script_org, BOM_UTF32_BE, b.SizeOf("BOM_UTF32_BE")-1)) {
		script_encoding = ZendMultibyteEncodingUtf32be
		bom_size = b.SizeOf("BOM_UTF32_BE") - 1
	} else if !(memcmp(LANG_SCNG__().script_org, BOM_UTF32_LE, b.SizeOf("BOM_UTF32_LE")-1)) {
		script_encoding = ZendMultibyteEncodingUtf32le
		bom_size = b.SizeOf("BOM_UTF32_LE") - 1
	} else if !(memcmp(LANG_SCNG__().script_org, BOM_UTF16_BE, b.SizeOf("BOM_UTF16_BE")-1)) {
		script_encoding = ZendMultibyteEncodingUtf16be
		bom_size = b.SizeOf("BOM_UTF16_BE") - 1
	} else if !(memcmp(LANG_SCNG__().script_org, BOM_UTF16_LE, b.SizeOf("BOM_UTF16_LE")-1)) {
		script_encoding = ZendMultibyteEncodingUtf16le
		bom_size = b.SizeOf("BOM_UTF16_LE") - 1
	} else if !(memcmp(LANG_SCNG__().script_org, BOM_UTF8, b.SizeOf("BOM_UTF8")-1)) {
		script_encoding = ZendMultibyteEncodingUtf8
		bom_size = b.SizeOf("BOM_UTF8") - 1
	}
	if script_encoding != nil {

		/* remove BOM */

		LANG_SCNG__().script_org += bom_size
		LANG_SCNG__().script_org_size -= bom_size
		return script_encoding
	}

	/* script contains NULL bytes -> auto-detection */

	if b.Assign(&pos1, memchr(LANG_SCNG__().script_org, 0, LANG_SCNG__().script_org_size)) {

		/* check if the NULL byte is after the __HALT_COMPILER(); */

		pos2 = LANG_SCNG__().script_org
		for size_t(pos1-pos2) >= b.SizeOf("\"__HALT_COMPILER();\"")-1 {
			pos2 = memchr(pos2, '_', pos1-pos2)
			if pos2 == nil {
				break
			}
			pos2++
			if strncasecmp((*byte)(pos2), "_HALT_COMPILER", b.SizeOf("\"_HALT_COMPILER\"")-1) == 0 {
				pos2 += b.SizeOf("\"_HALT_COMPILER\"") - 1
				for (*pos2) == ' ' || (*pos2) == '\t' || (*pos2) == '\r' || (*pos2) == '\n' {
					pos2++
				}
				if (*pos2) == '(' {
					pos2++
					for (*pos2) == ' ' || (*pos2) == '\t' || (*pos2) == '\r' || (*pos2) == '\n' {
						pos2++
					}
					if (*pos2) == ')' {
						pos2++
						for (*pos2) == ' ' || (*pos2) == '\t' || (*pos2) == '\r' || (*pos2) == '\n' {
							pos2++
						}
						if (*pos2) == ';' {
							return nil
						}
					}
				}
			}
		}

		/* make best effort if BOM is missing */

		return ZendMultibyteDetectUtfEncoding(LANG_SCNG__().script_org, LANG_SCNG__().script_org_size)

		/* make best effort if BOM is missing */

	}
	return nil
}
func ZendMultibyteFindScriptEncoding() *ZendEncoding {
	var script_encoding *ZendEncoding
	if CG__().GetDetectUnicode() != 0 {

		/* check out bom(byte order mark) and see if containing wchars */

		script_encoding = ZendMultibyteDetectUnicode()
		if script_encoding != nil {

			/* bom or wchar detection is prior to 'script_encoding' option */

			return script_encoding

			/* bom or wchar detection is prior to 'script_encoding' option */

		}
	}

	/* if no script_encoding specified, just leave alone */

	if CG__().GetScriptEncodingList() == nil || CG__().GetScriptEncodingListSize() == 0 {
		return nil
	}

	/* if multiple encodings specified, detect automagically */

	if CG__().GetScriptEncodingListSize() > 1 {
		return ZendMultibyteEncodingDetector(LANG_SCNG__().script_org, LANG_SCNG__().script_org_size, CG__().GetScriptEncodingList(), CG__().GetScriptEncodingListSize())
	}
	return CG__().GetScriptEncodingList()[0]
}
func ZendMultibyteSetFilter(onetime_encoding *ZendEncoding) int {
	var internal_encoding *ZendEncoding = ZendMultibyteGetInternalEncoding()
	var script_encoding *ZendEncoding = b.CondF2(onetime_encoding != nil, onetime_encoding, func() *ZendEncoding { return ZendMultibyteFindScriptEncoding() })
	if script_encoding == nil {
		return FAILURE
	}

	/* judge input/output filter */

	LANG_SCNG__().script_encoding = script_encoding
	LANG_SCNG__().input_filter = nil
	LANG_SCNG__().output_filter = nil
	if internal_encoding == nil || LANG_SCNG__().script_encoding == internal_encoding {
		if ZendMultibyteCheckLexerCompatibility(LANG_SCNG__().script_encoding) == 0 {

			/* and if not, work around w/ script_encoding -> utf-8 -> script_encoding conversion */

			LANG_SCNG__().input_filter = EncodingFilterScriptToIntermediate
			LANG_SCNG__().output_filter = EncodingFilterIntermediateToScript
		} else {
			LANG_SCNG__().input_filter = nil
			LANG_SCNG__().output_filter = nil
		}
		return SUCCESS
	}
	if ZendMultibyteCheckLexerCompatibility(internal_encoding) != 0 {
		LANG_SCNG__().input_filter = EncodingFilterScriptToInternal
		LANG_SCNG__().output_filter = nil
	} else if ZendMultibyteCheckLexerCompatibility(LANG_SCNG__().script_encoding) != 0 {
		LANG_SCNG__().input_filter = nil
		LANG_SCNG__().output_filter = EncodingFilterScriptToInternal
	} else {

		/* both script and internal encodings are incompatible w/ flex */

		LANG_SCNG__().input_filter = EncodingFilterScriptToIntermediate
		LANG_SCNG__().output_filter = EncodingFilterIntermediateToInternal
	}
	return 0
}
func OpenFileForScanning(file_handle *ZendFileHandle) int {
	var buf *byte
	var size int
	var compiled_filename *ZendString
	if ZendStreamFixup(file_handle, &buf, &size) == FAILURE {

		/* Still add it to open_files to make destroy_file_handle work */

		ZendLlistAddElement(CG__().GetOpenFiles(), file_handle)
		return FAILURE
	}
	ZEND_ASSERT(EG__().GetException() == nil && "stream_fixup() should have failed")
	ZendLlistAddElement(CG__().GetOpenFiles(), file_handle)
	if file_handle.GetStream().GetHandle() >= any(file_handle != nil && file_handle.GetStream().GetHandle() <= any(file_handle+1)) {
		var fh *ZendFileHandle = (*ZendFileHandle)(ZendLlistGetLast(CG__().GetOpenFiles()))
		var diff int = (*byte)(file_handle.GetStream().GetHandle() - (*byte)(file_handle))
		fh.GetStream().SetHandle(any((*byte)(fh) + diff))
		file_handle.GetStream().SetHandle(fh.GetStream().GetHandle())
	}

	/* Reset the scanner for scanning the new file */

	LANG_SCNG__().yy_in = file_handle
	LANG_SCNG__().yy_start = nil
	if size != size_t-1 {
		if CG__().GetMultibyte() != 0 {
			LANG_SCNG__().script_org = (*uint8)(buf)
			LANG_SCNG__().script_org_size = size
			LANG_SCNG__().script_filtered = nil
			ZendMultibyteSetFilter(nil)
			if LANG_SCNG__().input_filter {
				if size_t-1 == LANG_SCNG__().input_filter(&(LANG_SCNG__().script_filtered), &(LANG_SCNG__().script_filtered_size), LANG_SCNG__().script_org, LANG_SCNG__().script_org_size) {
					ZendErrorNoreturn(E_COMPILE_ERROR, "Could not convert the script from the detected "+"encoding \"%s\" to a compatible encoding", ZendMultibyteGetEncodingName(LANG_SCNG__().script_encoding))
				}
				buf = (*byte)(LANG_SCNG__().script_filtered)
				size = LANG_SCNG__().script_filtered_size
			}
		}
		LANG_SCNG__().yy_start = (*uint8)(buf)
		YyScanBuffer(buf, size)
	} else {
		ZendErrorNoreturn(E_COMPILE_ERROR, "zend_stream_mmap() failed")
	}
	if CG__().GetSkipShebang() != 0 {
		CG__().SetSkipShebang(0)
		BEGIN(SHEBANG)
	} else {
		BEGIN(INITIAL)
	}
	if file_handle.GetOpenedPath() != nil {
		compiled_filename = file_handle.GetOpenedPath().Copy()
	} else {
		compiled_filename = ZendStringInit(file_handle.GetFilename(), strlen(file_handle.GetFilename()), 0)
	}
	ZendSetCompiledFilename(compiled_filename)
	ZendStringReleaseEx(compiled_filename, 0)
	RESET_DOC_COMMENT()
	CG__().SetZendLineno(1)
	CG__().SetIncrementLineno(0)
	return SUCCESS
}
func ZendCompile(type_ int) *ZendOpArray {
	var op_array *ZendOpArray = nil
	var original_in_compilation ZendBool = CG__().GetInCompilation()
	CG__().SetInCompilation(1)
	CG__().SetAst(nil)
	CG__().SetAstArena(ZendArenaCreate(1024 * 32))
	if Zendparse() == 0 {
		var last_lineno int = CG__().GetZendLineno()
		var original_file_context ZendFileContext
		var original_oparray_context ZendOparrayContext
		var original_active_op_array *ZendOpArray = CG__().GetActiveOpArray()
		op_array = Emalloc(b.SizeOf("zend_op_array"))
		InitOpArray(op_array, type_, INITIAL_OP_ARRAY_SIZE)
		CG__().SetActiveOpArray(op_array)

		/* Use heap to not waste arena memory */

		op_array.SetIsHeapRtCache(true)
		if ZendAstProcess != nil {
			ZendAstProcess(CG__().GetAst())
		}
		ZendFileContextBegin(&original_file_context)
		ZendOparrayContextBegin(&original_oparray_context)
		ZendCompileTopStmt(CG__().GetAst())
		CG__().SetZendLineno(last_lineno)
		ZendEmitFinalReturn(type_ == ZEND_USER_FUNCTION)
		op_array.SetLineStart(1)
		op_array.SetLineEnd(last_lineno)
		PassTwo(op_array)
		ZendOparrayContextEnd(&original_oparray_context)
		ZendFileContextEnd(&original_file_context)
		CG__().SetActiveOpArray(original_active_op_array)
	}
	ZendAstDestroy(CG__().GetAst())
	ZendArenaDestroy(CG__().GetAstArena())
	CG__().SetInCompilation(original_in_compilation)
	return op_array
}
func CompileFile(file_handle *ZendFileHandle, type_ int) *ZendOpArray {
	var original_lex_state ZendLexState
	var op_array *ZendOpArray = nil
	ZendSaveLexicalState(&original_lex_state)
	if OpenFileForScanning(file_handle) == FAILURE {
		if EG__().GetException() == nil {
			if type_ == ZEND_REQUIRE {
				ZendMessageDispatcher(ZMSG_FAILED_REQUIRE_FOPEN, file_handle.GetFilename())
				ZendBailout()
			} else {
				ZendMessageDispatcher(ZMSG_FAILED_INCLUDE_FOPEN, file_handle.GetFilename())
			}
		}
	} else {
		op_array = ZendCompile(ZEND_USER_FUNCTION)
	}
	ZendRestoreLexicalState(&original_lex_state)
	return op_array
}
func CompileFilename(type_ int, filename *Zval) *ZendOpArray {
	var file_handle ZendFileHandle
	var tmp Zval
	var retval *ZendOpArray
	var opened_path *ZendString = nil
	if filename.GetType() != IS_STRING {
		tmp.SetString(ZvalGetString(filename))
		filename = &tmp
	}
	ZendStreamInitFilename(&file_handle, Z_STRVAL_P(filename))
	retval = ZendCompileFile(&file_handle, type_)
	if retval != nil && file_handle.GetStream().GetHandle() {
		if file_handle.GetOpenedPath() == nil {
			opened_path = filename.GetStr().Copy()
			file_handle.SetOpenedPath(opened_path)
		}
		ZendHashAddEmptyElement(EG__().GetIncludedFiles(), file_handle.GetOpenedPath())
		if opened_path != nil {
			ZendStringReleaseEx(opened_path, 0)
		}
	}
	ZendDestroyFileHandle(&file_handle)
	if filename == &tmp {
		ZvalPtrDtor(&tmp)
	}
	return retval
}
func ZendPrepareStringForScanning(str *Zval, filename *byte) int {
	var buf *byte
	var size int
	var old_len int
	var new_compiled_filename *ZendString

	/* enforce ZEND_MMAP_AHEAD trailing NULLs for flex... */

	old_len = Z_STRLEN_P(str)
	str.SetStr(ZendStringExtend(str.GetStr(), old_len+ZEND_MMAP_AHEAD, 0))
	str.SetTypeInfo(IS_STRING_EX)
	memset(Z_STRVAL_P(str)+old_len, 0, ZEND_MMAP_AHEAD+1)
	LANG_SCNG__().yy_in = nil
	LANG_SCNG__().yy_start = nil
	buf = Z_STRVAL_P(str)
	size = old_len
	if CG__().GetMultibyte() != 0 {
		LANG_SCNG__().script_org = (*uint8)(buf)
		LANG_SCNG__().script_org_size = size
		LANG_SCNG__().script_filtered = nil
		ZendMultibyteSetFilter(ZendMultibyteGetInternalEncoding())
		if LANG_SCNG__().input_filter {
			if size_t-1 == LANG_SCNG__().input_filter(&(LANG_SCNG__().script_filtered), &(LANG_SCNG__().script_filtered_size), LANG_SCNG__().script_org, LANG_SCNG__().script_org_size) {
				ZendErrorNoreturn(E_COMPILE_ERROR, "Could not convert the script from the detected "+"encoding \"%s\" to a compatible encoding", ZendMultibyteGetEncodingName(LANG_SCNG__().script_encoding))
			}
			buf = (*byte)(LANG_SCNG__().script_filtered)
			size = LANG_SCNG__().script_filtered_size
		}
	}
	YyScanBuffer(buf, size)
	new_compiled_filename = ZendStringInit(filename, strlen(filename), 0)
	ZendSetCompiledFilename(new_compiled_filename)
	ZendStringReleaseEx(new_compiled_filename, 0)
	CG__().SetZendLineno(1)
	CG__().SetIncrementLineno(0)
	RESET_DOC_COMMENT()
	return SUCCESS
}
func ZendGetScannedFileOffset() int {
	var offset int = LANG_SCNG__().yy_cursor - LANG_SCNG__().yy_start
	if LANG_SCNG__().input_filter {
		var original_offset int = offset
		var length int = 0
		for {
			var p *uint8 = nil
			if size_t-1 == LANG_SCNG__().input_filter(&p, &length, LANG_SCNG__().script_org, offset) {
				return size_t - 1
			}
			Efree(p)
			if length > original_offset {
				offset--
			} else if length < original_offset {
				offset++
			}
			if original_offset == length {
				break
			}
		}
	}
	return offset
}
func CompileString(source_string *Zval, filename *byte) *ZendOpArray {
	var original_lex_state ZendLexState
	var op_array *ZendOpArray = nil
	var tmp Zval
	if source_string.GetType() != IS_STRING {
		tmp.SetString(ZvalGetStringFunc(source_string))
	} else {
		ZVAL_COPY(&tmp, source_string)
	}
	if Z_STRLEN(tmp) == 0 {
		ZvalPtrDtor(&tmp)
		return nil
	}
	ZendSaveLexicalState(&original_lex_state)
	if ZendPrepareStringForScanning(&tmp, filename) == SUCCESS {
		BEGIN(ST_IN_SCRIPTING)
		op_array = ZendCompile(ZEND_EVAL_CODE)
	}
	ZendRestoreLexicalState(&original_lex_state)
	ZvalPtrDtor(&tmp)
	return op_array
}
func HighlightFile(filename *byte, syntax_highlighter_ini *ZendSyntaxHighlighterIni) int {
	var original_lex_state ZendLexState
	var file_handle ZendFileHandle
	ZendStreamInitFilename(&file_handle, filename)
	ZendSaveLexicalState(&original_lex_state)
	if OpenFileForScanning(&file_handle) == FAILURE {
		ZendMessageDispatcher(ZMSG_FAILED_HIGHLIGHT_FOPEN, filename)
		ZendRestoreLexicalState(&original_lex_state)
		return FAILURE
	}
	ZendHighlight(syntax_highlighter_ini)
	if LANG_SCNG__().script_filtered {
		Efree(LANG_SCNG__().script_filtered)
		LANG_SCNG__().script_filtered = nil
	}
	ZendDestroyFileHandle(&file_handle)
	ZendRestoreLexicalState(&original_lex_state)
	return SUCCESS
}
func HighlightString(str *Zval, syntax_highlighter_ini *ZendSyntaxHighlighterIni, str_name *byte) int {
	var original_lex_state ZendLexState
	var tmp Zval
	if str.GetType() != IS_STRING {
		tmp.SetString(ZvalGetStringFunc(str))
		str = &tmp
	}
	ZendSaveLexicalState(&original_lex_state)
	if ZendPrepareStringForScanning(str, str_name) == FAILURE {
		ZendRestoreLexicalState(&original_lex_state)
		if str == &tmp {
			ZvalPtrDtor(&tmp)
		}
		return FAILURE
	}
	BEGIN(INITIAL)
	ZendHighlight(syntax_highlighter_ini)
	if LANG_SCNG__().script_filtered {
		Efree(LANG_SCNG__().script_filtered)
		LANG_SCNG__().script_filtered = nil
	}
	ZendRestoreLexicalState(&original_lex_state)
	if str == &tmp {
		ZvalPtrDtor(&tmp)
	}
	return SUCCESS
}
func ZendMultibyteYyinputAgain(old_input_filter ZendEncodingFilter, old_encoding *ZendEncoding) {
	var length int
	var new_yy_start *uint8

	/* convert and set */

	if !(LANG_SCNG__().input_filter) {
		if LANG_SCNG__().script_filtered {
			Efree(LANG_SCNG__().script_filtered)
			LANG_SCNG__().script_filtered = nil
		}
		LANG_SCNG__().script_filtered_size = 0
		length = LANG_SCNG__().script_org_size
		new_yy_start = LANG_SCNG__().script_org
	} else {
		if size_t-1 == LANG_SCNG__().input_filter(&new_yy_start, &length, LANG_SCNG__().script_org, LANG_SCNG__().script_org_size) {
			ZendErrorNoreturn(E_COMPILE_ERROR, "Could not convert the script from the detected "+"encoding \"%s\" to a compatible encoding", ZendMultibyteGetEncodingName(LANG_SCNG__().script_encoding))
		}
		if LANG_SCNG__().script_filtered {
			Efree(LANG_SCNG__().script_filtered)
		}
		LANG_SCNG__().script_filtered = new_yy_start
		LANG_SCNG__().script_filtered_size = length
	}
	LANG_SCNG__().yy_cursor = new_yy_start + (LANG_SCNG__().yy_cursor - LANG_SCNG__().yy_start)
	LANG_SCNG__().yy_marker = new_yy_start + (LANG_SCNG__().yy_marker - LANG_SCNG__().yy_start)
	LANG_SCNG__().yy_text = new_yy_start + (LANG_SCNG__().yy_text - LANG_SCNG__().yy_start)
	LANG_SCNG__().yy_limit = new_yy_start + length
	LANG_SCNG__().yy_start = new_yy_start
}

// TODO: avoid reallocation ???

func ZendCopyValue(zendlval *Zval, yytext *byte, yyleng int) {
	if LANG_SCNG__().output_filter {
		var sz int = 0
		var s *byte = nil
		LANG_SCNG__().output_filter((**uint8)(&s), &sz, (*uint8)(yytext), int(yyleng))
		ZVAL_STRINGL(zendlval, s, sz)
		Efree(s)
	} else if yyleng == 1 {
		zendlval.SetInternedString(ZSTR_CHAR(zend_uchar * yytext))
	} else {
		ZVAL_STRINGL(zendlval, yytext, yyleng)
	}
}
func ZendScanEscapeString(zendlval *Zval, str *byte, len_ int, quote_type byte) int {
	var s *byte
	var t *byte
	var end *byte
	if len_ <= 1 {
		if len_ < 1 {
			ZVAL_EMPTY_STRING(zendlval)
		} else {
			var c ZendUchar = zend_uchar * str
			if c == '\n' || c == '\r' {
				CG__().GetZendLineno()++
			}
			zendlval.SetInternedString(ZSTR_CHAR(c))
		}
		goto skip_escape_conversion
	}
	ZVAL_STRINGL(zendlval, str, len_)

	/* convert escape sequences */

	s = Z_STRVAL_P(zendlval)
	end = s + Z_STRLEN_P(zendlval)
	for true {
		if (*s) == '\\' {
			break
		}
		if (*s) == '\n' || (*s) == '\r' && (*(s + 1)) != '\n' {
			CG__().GetZendLineno()++
		}
		s++
		if s == end {
			goto skip_escape_conversion
		}
	}
	t = s
	for s < end {
		if (*s) == '\\' {
			s++
			if s >= end {
				b.PostInc(&(*t)) = '\\'
				break
			}
			switch *s {
			case 'n':
				b.PostInc(&(*t)) = '\n'
			case 'r':
				b.PostInc(&(*t)) = '\r'
			case 't':
				b.PostInc(&(*t)) = '\t'
			case 'f':
				b.PostInc(&(*t)) = 'f'
			case 'v':
				b.PostInc(&(*t)) = 'v'
			case 'e':
				b.PostInc(&(*t)) = 'e'
			case '"':
				fallthrough
			case '`':
				if (*s) != quote_type {
					b.PostInc(&(*t)) = '\\'
					b.PostInc(&(*t)) = *s
					break
				}
				fallthrough
			case '\\':
				fallthrough
			case '$':
				b.PostInc(&(*t)) = *s
			case 'x':
				fallthrough
			case 'X':
				if sc.ZEND_IS_HEX(*(s + 1)) {
					var hex_buf []byte = []byte{0, 0, 0}
					hex_buf[0] = *(b.PreInc(&s))
					if sc.ZEND_IS_HEX(*(s + 1)) {
						hex_buf[1] = *(b.PreInc(&s))
					}
					b.PostInc(&(*t)) = byte(ZEND_STRTOL(hex_buf, nil, 16))
				} else {
					b.PostInc(&(*t)) = '\\'
					b.PostInc(&(*t)) = *s
				}
			case 'u':

				/* cache where we started so we can parse after validating */

				var start *byte = s + 1
				var len_ int = 0
				var valid ZendBool = 1
				var codepoint uint64
				if (*start) != '{' {

					/* we silently let this pass to avoid breaking code
					 * with JSON in string literals (e.g. "\"\u202e\""
					 */

					b.PostInc(&(*t)) = '\\'
					b.PostInc(&(*t)) = 'u'
					break
				} else {

					/* on the other hand, invalid \u{blah} errors */

					s++
					len_++
					s++
					for (*s) != '}' {
						if !(sc.ZEND_IS_HEX(*s)) {
							valid = 0
							break
						} else {
							len_++
						}
						s++
					}
					if (*s) == '}' {
						valid = 1
						len_++
					}
				}

				/* \u{} is invalid */

				if len_ <= 2 {
					valid = 0
				}
				if valid == 0 {
					ZendThrowException(ZendCeParseError, "Invalid UTF-8 codepoint escape sequence", 0)
					ZvalPtrDtor(zendlval)
					zendlval.SetUndef()
					return FAILURE
				}
				errno = 0
				codepoint = strtoul(start+1, nil, 16)

				/* per RFC 3629, UTF-8 can only represent 21 bits */

				if codepoint > 0x10ffff || errno {
					ZendThrowException(ZendCeParseError, "Invalid UTF-8 codepoint escape sequence: Codepoint too large", 0)
					ZvalPtrDtor(zendlval)
					zendlval.SetUndef()
					return FAILURE
				}

				/* based on https://en.wikipedia.org/wiki/UTF-8#Sample_code */

				if codepoint < 0x80 {
					b.PostInc(&(*t)) = codepoint
				} else if codepoint <= 0x7ff {
					b.PostInc(&(*t)) = (codepoint >> 6) + 0xc0
					b.PostInc(&(*t)) = (codepoint & 0x3f) + 0x80
				} else if codepoint <= 0xffff {
					b.PostInc(&(*t)) = (codepoint >> 12) + 0xe0
					b.PostInc(&(*t)) = (codepoint >> 6 & 0x3f) + 0x80
					b.PostInc(&(*t)) = (codepoint & 0x3f) + 0x80
				} else if codepoint <= 0x10ffff {
					b.PostInc(&(*t)) = (codepoint >> 18) + 0xf0
					b.PostInc(&(*t)) = (codepoint >> 12 & 0x3f) + 0x80
					b.PostInc(&(*t)) = (codepoint >> 6 & 0x3f) + 0x80
					b.PostInc(&(*t)) = (codepoint & 0x3f) + 0x80
				}

				/* based on https://en.wikipedia.org/wiki/UTF-8#Sample_code */

			default:

				/* check for an octal */

				if sc.ZEND_IS_OCT(*s) {
					var octal_buf []byte = []byte{0, 0, 0, 0}
					octal_buf[0] = *s
					if sc.ZEND_IS_OCT(*(s + 1)) {
						octal_buf[1] = *(b.PreInc(&s))
						if sc.ZEND_IS_OCT(*(s + 1)) {
							octal_buf[2] = *(b.PreInc(&s))
						}
					}
					if octal_buf[2] && octal_buf[0] > '3' && !(LANG_SCNG__().heredoc_scan_ahead) {

						/* 3 octit values must not overflow 0xFF (\377) */

						ZendError(E_COMPILE_WARNING, "Octal escape sequence overflow \\%s is greater than \\377", octal_buf)

						/* 3 octit values must not overflow 0xFF (\377) */

					}
					b.PostInc(&(*t)) = byte(ZEND_STRTOL(octal_buf, nil, 8))
				} else {
					b.PostInc(&(*t)) = '\\'
					b.PostInc(&(*t)) = *s
				}
			}
		} else {
			b.PostInc(&(*t)) = *s
		}
		if (*s) == '\n' || (*s) == '\r' && (*(s + 1)) != '\n' {
			CG__().GetZendLineno()++
		}
		s++
	}
	*t = 0
	Z_STRLEN_P(zendlval) = t - Z_STRVAL_P(zendlval)
skip_escape_conversion:
	if LANG_SCNG__().output_filter {
		var sz int = 0
		var str *uint8

		// TODO: avoid realocation ???

		s = Z_STRVAL_P(zendlval)
		LANG_SCNG__().output_filter(&str, &sz, (*uint8)(s), int(Z_STRLEN_P(zendlval)))
		ZvalPtrDtor(zendlval)
		ZVAL_STRINGL(zendlval, (*byte)(str), sz)
		Efree(str)
	}
	return SUCCESS
}

const HEREDOC_USING_SPACES = 1
const HEREDOC_USING_TABS = 2

func NextNewline(str *byte, end *byte, newline_len *int) *byte {
	for ; str < end; str++ {
		if (*str) == '\r' {
			if str+1 < end && (*(str + 1)) == '\n' {
				*newline_len = 2
			} else {
				*newline_len = 1
			}
			return str
		} else if (*str) == '\n' {
			*newline_len = 1
			return str
		}
	}
	*newline_len = 0
	return nil
}
func StripMultilineStringIndentation(zendlval *Zval, indentation int, using_spaces ZendBool, newline_at_start ZendBool, newline_at_end ZendBool) ZendBool {
	var str *byte = Z_STRVAL_P(zendlval)
	var end *byte = str + Z_STRLEN_P(zendlval)
	var copy *byte = Z_STRVAL_P(zendlval)
	var newline_count int = 0
	var newline_len int
	var nl *byte
	if newline_at_start == 0 {
		nl = NextNewline(str, end, &newline_len)
		if nl == nil {
			return 1
		}
		str = nl + newline_len
		copy = (*byte)(nl + newline_len)
		newline_count++
	} else {
		nl = str
	}

	/* <= intentional */

	for str <= end && nl != nil {
		var skip int
		nl = NextNewline(str, end, &newline_len)
		if nl == nil && newline_at_end != 0 {
			nl = end
		}

		/* Try to skip indentation */

		for skip = 0; skip < indentation; {
			if str == nl {

				/* Don't require full indentation on whitespace-only lines */

				break

				/* Don't require full indentation on whitespace-only lines */

			}
			if str == end || (*str) != ' ' && (*str) != '\t' {
				CG__().SetZendLineno(CG__().GetZendLineno() + newline_count)
				ZendThrowExceptionEx(ZendCeParseError, 0, "Invalid body indentation level (expecting an indentation level of at least %d)", indentation)
				goto error
			}
			if using_spaces == 0 && (*str) == ' ' || using_spaces != 0 && (*str) == '\t' {
				CG__().SetZendLineno(CG__().GetZendLineno() + newline_count)
				ZendThrowException(ZendCeParseError, "Invalid indentation - tabs and spaces cannot be mixed", 0)
				goto error
			}
			skip++
			str++
		}
		if str == end {
			break
		}
		var len_ int = b.Cond(nl != nil, nl-str+newline_len, end-str)
		memmove(copy, str, len_)
		str += len_
		copy += len_
		newline_count++
	}
	*copy = '0'
	Z_STRLEN_P(zendlval) = copy - Z_STRVAL_P(zendlval)
	return 1
error:
	ZvalPtrDtorStr(zendlval)
	zendlval.SetUndef()
	return 0
}
func CopyHeredocLabelStack(void_heredoc_label any) {
	var heredoc_label *ZendHeredocLabel = void_heredoc_label
	var new_heredoc_label *ZendHeredocLabel = Emalloc(b.SizeOf("zend_heredoc_label"))
	*new_heredoc_label = *heredoc_label
	new_heredoc_label.SetLabel(Estrndup(heredoc_label.GetLabel(), heredoc_label.GetLength()))
	ZendPtrStackPush(&(LANG_SCNG__().heredoc_label_stack), any(new_heredoc_label))
}
func PARSER_MODE() bool { return elem != nil }

// #define RETURN_TOKEN(_token) do { token = _token ; goto emit_token ; } while ( 0 )

func RETURN_TOKEN_WITH_VAL(_token Yytokentype) {
	token = _token
	goto emit_token_with_val
}
func RETURN_TOKEN_WITH_STR(_token Yytokentype, _offset int) {
	token = _token
	offset = _offset
	goto emit_token_with_str
}
func RETURN_OR_SKIP_TOKEN(_token Yytokentype) {
	token = _token
	if PARSER_MODE() {
		goto skip_token
	}
	goto emit_token
}
