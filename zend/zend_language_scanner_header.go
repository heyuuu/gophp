// <<generate>>

package zend

import (
	b "sik/builtin"
	"strings"
)

const YYCURSOR *uint8 = LANG_SCNG__().yy_cursor
const YYLIMIT *uint8 = LANG_SCNG__().yy_limit
const YYMARKER = LANG_SCNG__().yy_marker

func YYSETCONDITION(s int) __auto__ {
	LANG_SCNG__().yy_state = s
	return LANG_SCNG__().yy_state
}

// #define STATE(name) yyc ## name

/* emulate flex constructs */

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
func HeredocLabelDtor(heredoc_label *ZendHeredocLabel) { efree(heredoc_label.label) }
func ShutdownScanner() {
	var sc *LangScanner
	CG__().parse_error = 0
	RESET_DOC_COMMENT()
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
func ZendDestroyFileHandle(file_handle *ZendFileHandle) {
	ZendLlistDelElement(CG__().open_files, file_handle, (func(any, any) int)(ZendCompareFileHandles))

	file_handle.openedPath = ""
	file_handle.filename = ""
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
	ZEND_ASSERT(p >= script)
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

		/* searching for UTF-32 specific byte orders, so this will do */

	}

	/* BE or LE? */

	p = script
	assert(p >= script)
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
			return zend_multibyte_encoding_utf16le
		} else {
			return zend_multibyte_encoding_utf16be
		}
	} else {
		if le != 0 {
			return zend_multibyte_encoding_utf32le
		} else {
			return zend_multibyte_encoding_utf32be
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
func ZendMultibyteFindScriptEncoding() *zend_encoding {
	var script_encoding *zend_encoding
	if CG__().detect_unicode {

		/* check out bom(byte order mark) and see if containing wchars */

		script_encoding = ZendMultibyteDetectUnicode()
		if script_encoding != nil {

			/* bom or wchar detection is prior to 'script_encoding' option */

			return script_encoding

			/* bom or wchar detection is prior to 'script_encoding' option */

		}
	}

	/* if no script_encoding specified, just leave alone */

	if !(CG__().script_encoding_list) || !(CG__().script_encoding_list_size) {
		return nil
	}

	/* if multiple encodings specified, detect automagically */

	if CG__().script_encoding_list_size > 1 {
		return zend_multibyte_encoding_detector(LANG_SCNG__().script_org, LANG_SCNG__().script_org_size, CG__().script_encoding_list, CG__().script_encoding_list_size)
	}
	return CG__().script_encoding_list[0]
}
func ZendMultibyteSetFilter(onetime_encoding *zend_encoding) int {
	var internal_encoding *zend_encoding = zend_multibyte_get_internal_encoding()
	var script_encoding *zend_encoding = b.CondF2(onetime_encoding != nil, onetime_encoding, func() *zend_encoding { return ZendMultibyteFindScriptEncoding() })
	if script_encoding == nil {
		return FAILURE
	}

	/* judge input/output filter */

	LANG_SCNG__().script_encoding = script_encoding
	LANG_SCNG__().input_filter = nil
	LANG_SCNG__().output_filter = nil
	if internal_encoding == nil || LANG_SCNG__().script_encoding == internal_encoding {
		if !(zend_multibyte_check_lexer_compatibility(LANG_SCNG__().script_encoding)) {

			/* and if not, work around w/ script_encoding -> utf-8 -> script_encoding conversion */

			LANG_SCNG__().input_filter = EncodingFilterScriptToIntermediate
			LANG_SCNG__().output_filter = EncodingFilterIntermediateToScript
		} else {
			LANG_SCNG__().input_filter = nil
			LANG_SCNG__().output_filter = nil
		}
		return SUCCESS
	}
	if zend_multibyte_check_lexer_compatibility(internal_encoding) {
		LANG_SCNG__().input_filter = EncodingFilterScriptToInternal
		LANG_SCNG__().output_filter = nil
	} else if zend_multibyte_check_lexer_compatibility(LANG_SCNG__().script_encoding) {
		LANG_SCNG__().input_filter = nil
		LANG_SCNG__().output_filter = EncodingFilterScriptToInternal
	} else {

		/* both script and internal encodings are incompatible w/ flex */

		LANG_SCNG__().input_filter = EncodingFilterScriptToIntermediate
		LANG_SCNG__().output_filter = EncodingFilterIntermediateToInternal
	}
	return 0
}
func OpenFileForScanning(fileHandle *ZendFileHandle) int {
	buf, ok := fileHandle.Fixup()
	if !ok {
		/* Still add it to open_files to make destroy_file_handle work */
		ZendLlistAddElement(CG__().open_files, fileHandle)
		return FAILURE
	}
	size := len(buf)

	var compiled_filename *ZendString
	ZEND_ASSERT(!(EG__().exception) && "stream_fixup() should have failed")

	ZendLlistAddElement(CG__().open_files, fileHandle)

	// todo 没看懂
	if fileHandle.stream.handle >= (any)(fileHandle) && fileHandle.stream.handle <= (any)(fileHandle+1)) {
		var fh *ZendFileHandle = (*ZendFileHandle)(ZendLlistGetLast(CG__().open_files))
		var diff int = (*byte)(fileHandle.stream.handle - (*byte)(fileHandle))
		fh.stream.handle = any((*byte)(fh) + diff)
		fileHandle.stream.handle = fh.stream.handle
	}

	/* Reset the scanner for scanning the new file */

	LANG_SCNG__().yy_in = fileHandle
	LANG_SCNG__().yy_start = nil
	if size != -1 {
		if CG__().multibyte {
			LANG_SCNG__().script_org = (*uint8)(buf)
			LANG_SCNG__().script_org_size = size
			LANG_SCNG__().script_filtered = nil
			ZendMultibyteSetFilter(nil)
			if LANG_SCNG__().input_filter {
				if size_t-1 == LANG_SCNG__().input_filter(&(LANG_SCNG__().script_filtered), &(LANG_SCNG__().script_filtered_size), LANG_SCNG__().script_org, LANG_SCNG__().script_org_size) {
					zend_error_noreturn(E_COMPILE_ERROR, "Could not convert the script from the detected "+"encoding \"%s\" to a compatible encoding", zend_multibyte_get_encoding_name(LANG_SCNG__().script_encoding))
				}
				buf = (*byte)(LANG_SCNG__().script_filtered)
				size = LANG_SCNG__().script_filtered_size
			}
		}
		LANG_SCNG__().yy_start = (*uint8)(buf)
		YyScanBuffer(buf, size)
	} else {
		zend_error_noreturn(E_COMPILE_ERROR, "zend_stream_mmap() failed")
	}
	if CG__().skip_shebang {
		CG__().skip_shebang = 0
		BEGIN(SHEBANG)
	} else {
		BEGIN(INITIAL)
	}
	if fileHandle.openedPath {
		compiled_filename = zend_string_copy(fileHandle.openedPath)
	} else {
		compiled_filename = zend_string_init(fileHandle.filename, strlen(fileHandle.filename), 0)
	}
	zend_set_compiled_filename(compiled_filename)
	zend_string_release_ex(compiled_filename, 0)
	RESET_DOC_COMMENT()
	CG__().zend_lineno = 1
	CG__().increment_lineno = 0
	return SUCCESS
}
func ZendCompile(type_ int) *zend_op_array {
	var op_array *zend_op_array = nil
	var original_in_compilation zend_bool = CG__().in_compilation
	CG__().in_compilation = 1
	CG__().ast = nil
	CG__().ast_arena = zend_arena_create(1024 * 32)
	if !(zendparse()) {
		var last_lineno int = CG__().zend_lineno
		var original_file_context zend_file_context
		var original_oparray_context zend_oparray_context
		var original_active_op_array *zend_op_array = CG__().active_op_array
		op_array = emalloc(b.SizeOf("zend_op_array"))
		init_op_array(op_array, type_, INITIAL_OP_ARRAY_SIZE)
		CG__().active_op_array = op_array

		/* Use heap to not waste arena memory */

		op_array.fn_flags |= ZEND_ACC_HEAP_RT_CACHE
		if zend_ast_process {
			zend_ast_process(CG__().ast)
		}
		zend_file_context_begin(&original_file_context)
		zend_oparray_context_begin(&original_oparray_context)
		zend_compile_top_stmt(CG__().ast)
		CG__().zend_lineno = last_lineno
		zend_emit_final_return(type_ == ZEND_USER_FUNCTION)
		op_array.line_start = 1
		op_array.line_end = last_lineno
		pass_two(op_array)
		zend_oparray_context_end(&original_oparray_context)
		zend_file_context_end(&original_file_context)
		CG__().active_op_array = original_active_op_array
	}
	zend_ast_destroy(CG__().ast)
	zend_arena_destroy(CG__().ast_arena)
	CG__().in_compilation = original_in_compilation
	return op_array
}
func CompileFile(file_handle *zend_file_handle, type_ int) *zend_op_array {
	var original_lex_state ZendLexState
	var op_array *zend_op_array = nil
	ZendSaveLexicalState(&original_lex_state)
	if OpenFileForScanning(file_handle) == FAILURE {
		if !(EG__().exception) {
			if type_ == ZEND_REQUIRE {
				zend_message_dispatcher(ZMSG_FAILED_REQUIRE_FOPEN, file_handle.filename)
				zend_bailout()
			} else {
				zend_message_dispatcher(ZMSG_FAILED_INCLUDE_FOPEN, file_handle.filename)
			}
		}
	} else {
		op_array = ZendCompile(ZEND_USER_FUNCTION)
	}
	ZendRestoreLexicalState(&original_lex_state)
	return op_array
}
func CompileFilename(type_ int, filename *zval) *zend_op_array {
	var file_handle zend_file_handle
	var tmp zval
	var retval *zend_op_array
	var opened_path *zend_string = nil
	if Z_TYPE_P(filename) != IS_STRING {
		ZVAL_STR(&tmp, zval_get_string(filename))
		filename = &tmp
	}
	zend_stream_init_filename(&file_handle, Z_STRVAL_P(filename))
	retval = zend_compile_file(&file_handle, type_)
	if retval != nil && file_handle.handle.stream.handle {
		if !(file_handle.opened_path) {
			opened_path = zend_string_copy(Z_STR_P(filename))
			file_handle.opened_path = opened_path
		}
		zend_hash_add_empty_element(EG__().included_files, file_handle.opened_path)
		if opened_path != nil {
			zend_string_release_ex(opened_path, 0)
		}
	}
	ZendDestroyFileHandle(&file_handle)
	if filename == &tmp {
		zval_ptr_dtor(&tmp)
	}
	return retval
}
func ZendPrepareStringForScanning(str *zval, filename *byte) int {
	var buf *byte
	var size int
	var old_len int
	var new_compiled_filename *zend_string

	/* enforce ZEND_MMAP_AHEAD trailing NULLs for flex... */

	old_len = Z_STRLEN_P(str)
	Z_STR_P(str) = zend_string_extend(Z_STR_P(str), old_len+ZEND_MMAP_AHEAD, 0)
	Z_TYPE_INFO_P(str) = IS_STRING_EX
	memset(Z_STRVAL_P(str)+old_len, 0, ZEND_MMAP_AHEAD+1)
	LANG_SCNG__().yy_in = nil
	LANG_SCNG__().yy_start = nil
	buf = Z_STRVAL_P(str)
	size = old_len
	if CG__().multibyte {
		LANG_SCNG__().script_org = (*uint8)(buf)
		LANG_SCNG__().script_org_size = size
		LANG_SCNG__().script_filtered = nil
		ZendMultibyteSetFilter(zend_multibyte_get_internal_encoding())
		if LANG_SCNG__().input_filter {
			if size_t-1 == LANG_SCNG__().input_filter(&(LANG_SCNG__().script_filtered), &(LANG_SCNG__().script_filtered_size), LANG_SCNG__().script_org, LANG_SCNG__().script_org_size) {
				zend_error_noreturn(E_COMPILE_ERROR, "Could not convert the script from the detected "+"encoding \"%s\" to a compatible encoding", zend_multibyte_get_encoding_name(LANG_SCNG__().script_encoding))
			}
			buf = (*byte)(LANG_SCNG__().script_filtered)
			size = LANG_SCNG__().script_filtered_size
		}
	}
	YyScanBuffer(buf, size)
	new_compiled_filename = zend_string_init(filename, strlen(filename), 0)
	zend_set_compiled_filename(new_compiled_filename)
	zend_string_release_ex(new_compiled_filename, 0)
	CG__().zend_lineno = 1
	CG__().increment_lineno = 0
	RESET_DOC_COMMENT()
	return SUCCESS
}
func CompileString(source_string *zval, filename *byte) *zend_op_array {
	var original_lex_state ZendLexState
	var op_array *zend_op_array = nil
	var tmp zval
	if Z_TYPE_P(source_string) != IS_STRING {
		ZVAL_STR(&tmp, zval_get_string_func(source_string))
	} else {
		ZVAL_COPY(&tmp, source_string)
	}
	if Z_STRLEN(tmp) == 0 {
		zval_ptr_dtor(&tmp)
		return nil
	}
	ZendSaveLexicalState(&original_lex_state)
	if ZendPrepareStringForScanning(&tmp, filename) == SUCCESS {
		BEGIN(ST_IN_SCRIPTING)
		op_array = ZendCompile(ZEND_EVAL_CODE)
	}
	ZendRestoreLexicalState(&original_lex_state)
	zval_ptr_dtor(&tmp)
	return op_array
}
func HighlightFile(filename *byte, syntax_highlighter_ini *zend_syntax_highlighter_ini) int {
	var original_lex_state ZendLexState
	var file_handle zend_file_handle
	zend_stream_init_filename(&file_handle, filename)
	ZendSaveLexicalState(&original_lex_state)
	if OpenFileForScanning(&file_handle) == FAILURE {
		zend_message_dispatcher(ZMSG_FAILED_HIGHLIGHT_FOPEN, filename)
		ZendRestoreLexicalState(&original_lex_state)
		return FAILURE
	}
	zend_highlight(syntax_highlighter_ini)
	if LANG_SCNG__().script_filtered {
		efree(LANG_SCNG__().script_filtered)
		LANG_SCNG__().script_filtered = nil
	}
	ZendDestroyFileHandle(&file_handle)
	ZendRestoreLexicalState(&original_lex_state)
	return SUCCESS
}
func HighlightString(str *zval, syntax_highlighter_ini *zend_syntax_highlighter_ini, str_name *byte) int {
	var original_lex_state ZendLexState
	var tmp zval
	if Z_TYPE_P(str) != IS_STRING {
		ZVAL_STR(&tmp, zval_get_string_func(str))
		str = &tmp
	}
	ZendSaveLexicalState(&original_lex_state)
	if ZendPrepareStringForScanning(str, str_name) == FAILURE {
		ZendRestoreLexicalState(&original_lex_state)
		if str == &tmp {
			zval_ptr_dtor(&tmp)
		}
		return FAILURE
	}
	BEGIN(INITIAL)
	zend_highlight(syntax_highlighter_ini)
	if LANG_SCNG__().script_filtered {
		efree(LANG_SCNG__().script_filtered)
		LANG_SCNG__().script_filtered = nil
	}
	ZendRestoreLexicalState(&original_lex_state)
	if str == &tmp {
		zval_ptr_dtor(&tmp)
	}
	return SUCCESS
}
func ZendMultibyteYyinputAgain(old_input_filter zend_encoding_filter, old_encoding *zend_encoding) {
	var length int
	var new_yy_start *uint8

	/* convert and set */

	if !(LANG_SCNG__().input_filter) {
		if LANG_SCNG__().script_filtered {
			efree(LANG_SCNG__().script_filtered)
			LANG_SCNG__().script_filtered = nil
		}
		LANG_SCNG__().script_filtered_size = 0
		length = LANG_SCNG__().script_org_size
		new_yy_start = LANG_SCNG__().script_org
	} else {
		if size_t-1 == LANG_SCNG__().input_filter(&new_yy_start, &length, LANG_SCNG__().script_org, LANG_SCNG__().script_org_size) {
			zend_error_noreturn(E_COMPILE_ERROR, "Could not convert the script from the detected "+"encoding \"%s\" to a compatible encoding", zend_multibyte_get_encoding_name(LANG_SCNG__().script_encoding))
		}
		if LANG_SCNG__().script_filtered {
			efree(LANG_SCNG__().script_filtered)
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

func ZendCopyValue(zendlval __auto__, yytext __auto__, yyleng __auto__) {
	if LANG_SCNG__().output_filter {
		var sz int = 0
		var s *byte = nil
		LANG_SCNG__().output_filter((**uint8)(&s), &sz, (*uint8)(yytext), int(yyleng))
		ZVAL_STRINGL(zendlval, s, sz)
		efree(s)
	} else if yyleng == 1 {
		ZVAL_INTERNED_STR(zendlval, ZSTR_CHAR(zend_uchar*yytext))
	} else {
		ZVAL_STRINGL(zendlval, yytext, yyleng)
	}
}

func escapeString(str string, quoteType byte) (string, bool) {
	len_ := len(str)
	if len_ <= 1 || strings.IndexByte(str, '\\') >= 0 {
		// 无转义直接返回
		return str, true
	}

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
			// todo
		case 'u':
			// todo
		default:
			// todo
		}
	}

	return buf.String(), true
}

func (sc *LangScanner) ScanEscapeString(str string, quoteType byte) bool {
	str, ok := escapeString(str, quoteType)
	if !ok {
		return false
	}
	sc.setStrFiltered(str)
	return true
}

func ZendScanEscapeString(zendlval *zval, str *byte, len_ int, quote_type byte) int {

	var s *byte
	var t *byte
	var end *byte
	if len_ <= 1 {
		if len_ < 1 {
			ZVAL_EMPTY_STRING(zendlval)
		} else {
			var c zend_uchar = zend_uchar * str
			if c == '\n' || c == '\r' {
				CG__().zend_lineno++
			}
			ZVAL_INTERNED_STR(zendlval, ZSTR_CHAR(c))
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
			CG__().zend_lineno++
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
				if zendIsHex(*(s + 1)) {
					var hex_buf []byte = []byte{0, 0, 0}
					hex_buf[0] = *(b.PreInc(&s))
					if zendIsHex(*(s + 1)) {
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
				var valid zend_bool = 1
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
						if !(zendIsHex(*s)) {
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
				if !valid {
					zend_throw_exception(zend_ce_parse_error, "Invalid UTF-8 codepoint escape sequence", 0)
					zval_ptr_dtor(zendlval)
					ZVAL_UNDEF(zendlval)
					return FAILURE
				}
				errno = 0
				codepoint = strtoul(start+1, nil, 16)

				/* per RFC 3629, UTF-8 can only represent 21 bits */

				if codepoint > 0x10ffff || errno {
					zend_throw_exception(zend_ce_parse_error, "Invalid UTF-8 codepoint escape sequence: Codepoint too large", 0)
					zval_ptr_dtor(zendlval)
					ZVAL_UNDEF(zendlval)
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

				if zendIsOct(*s) {
					var octal_buf []byte = []byte{0, 0, 0, 0}
					octal_buf[0] = *s
					if zendIsOct(*(s + 1)) {
						octal_buf[1] = *(b.PreInc(&s))
						if zendIsOct(*(s + 1)) {
							octal_buf[2] = *(b.PreInc(&s))
						}
					}
					if octal_buf[2] && octal_buf[0] > '3' && !(LANG_SCNG__().heredoc_scan_ahead) {

						/* 3 octit values must not overflow 0xFF (\377) */

						zend_error(E_COMPILE_WARNING, "Octal escape sequence overflow \\%s is greater than \\377", octal_buf)

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
			CG__().zend_lineno++
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
		zval_ptr_dtor(zendlval)
		ZVAL_STRINGL(zendlval, (*byte)(str), sz)
		efree(str)
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
func StripMultilineStringIndentation(zendlval *zval, indentation int, using_spaces zend_bool, newline_at_start zend_bool, newline_at_end zend_bool) zend_bool {
	var str *byte = Z_STRVAL_P(zendlval)
	var end *byte = str + Z_STRLEN_P(zendlval)
	var copy *byte = Z_STRVAL_P(zendlval)
	var newline_count int = 0
	var newline_len int
	var nl *byte
	if !newline_at_start {
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
		if nl == nil && newline_at_end {
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
				CG__().zend_lineno += newline_count
				zend_throw_exception_ex(zend_ce_parse_error, 0, "Invalid body indentation level (expecting an indentation level of at least %d)", indentation)
				goto error
			}
			if !using_spaces && (*str) == ' ' || using_spaces && (*str) == '\t' {
				CG__().zend_lineno += newline_count
				zend_throw_exception(zend_ce_parse_error, "Invalid indentation - tabs and spaces cannot be mixed", 0)
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
	zval_ptr_dtor_str(zendlval)
	ZVAL_UNDEF(zendlval)
	return 0
}
func CopyHeredocLabelStack(heredocLabel *ZendHeredocLabel) {
	newHeredocLabel := heredocLabel.Copy()
	LANG_SCNG__().heredoc_label_stack.Push(newHeredocLabel)
}
