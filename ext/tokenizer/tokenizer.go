package tokenizer

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/ext/standard"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
	"github.com/heyuuu/gophp/zend/zpp"
)

var TokenizerFunctions = []types.FunctionEntry{
	DefZifTokenGetAll,
	DefZifTokenName,
}
var TokenizerModuleEntry zend.ModuleEntry = zend.MakeZendModuleEntry(
	"tokenizer",
	TokenizerFunctions,
	ZmStartupTokenizer,
	nil,
	nil,
	nil,
	ZmInfoTokenizer,
	core.PHP_VERSION,
	0,
	nil,
	nil,
	nil,
)

/**
 * functions
 */
func TokenizerTokenGetAllRegisterConstants(type_ int, module_number int) {
	zend.RegisterLongConstant("TOKEN_PARSE", 1, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
}
func ZmStartupTokenizer(type_ int, module_number int) int {
	TokenizerRegisterConstants(type_, module_number)
	TokenizerTokenGetAllRegisterConstants(type_, module_number)
	return types.SUCCESS
}
func ZmInfoTokenizer(zend_module *zend.ModuleEntry) {
	standard.PhpInfoPrintTableStart()
	standard.PhpInfoPrintTableRow(2, "Tokenizer Support", "enabled")
	standard.PhpInfoPrintTableEnd()
}
func AddToken(return_value *types.Zval, token_type int, text *uint8, leng int, lineno int) {
	if token_type >= 256 {
		arr := types.NewArray(0)
		arr.Append(types.NewZvalLong(token_type))
		arr.Append(types.NewZvalString(b.CastStr(text, leng)))
		arr.Append(types.NewZvalLong(lineno))

		return_value.Array().Append(types.NewZvalArray(arr))
	} else {
		return_value.Array().Append(types.NewZvalString(b.CastStr(text, leng)))
	}
}
func Tokenize(return_value *types.Zval, source string) bool {
	var source_zval types.Zval
	var original_lex_state zend.ZendLexState
	var token types.Zval
	var token_line int = 1
	var need_tokens int = -1
	source_zval.SetStringVal(source)
	zend.ZendSaveLexicalState(&original_lex_state)
	if zend.ZendPrepareStringForScanning(&source_zval, "") == types.FAILURE {
		zend.ZendRestoreLexicalState(&original_lex_state)
		return 0
	}
	sc := zend.NewLangScanner(source, false)
	sc.Begin(zend.LANG_YYC_INITIAL)

	zend.ArrayInit(return_value)
	for {
		token_type := sc.LexScanEx(&token, nil)
		if token_type == 0 {
			break
		}

		AddToken(return_value, token_type, zend.INI_SCNG__().GetYyText(), zend.INI_SCNG__().GetYyLeng(), token_line)
		if token.IsNotUndef() {
			// zend.ZvalPtrDtorNogc(&token)
			token.SetUndef()
		}

		/* after T_HALT_COMPILER collect the next three non-dropped tokens */

		if need_tokens != -1 {
			if token_type != zend.T_WHITESPACE && token_type != zend.T_OPEN_TAG && token_type != zend.T_COMMENT && token_type != zend.T_DOC_COMMENT && b.PreDec(&need_tokens) == 0 {

				/* fetch the rest into a T_INLINE_HTML */

				if zend.INI_SCNG__().GetYyCursor() != zend.INI_SCNG__().GetYyLimit() {
					AddToken(return_value, zend.T_INLINE_HTML, zend.INI_SCNG__().GetYyCursor(), zend.INI_SCNG__().GetYyLimit()-zend.INI_SCNG__().GetYyCursor(), token_line)
				}
				break
			}
		} else if token_type == zend.T_HALT_COMPILER {
			need_tokens = 3
		}
		if zend.CG__().GetIncrementLineno() != 0 {
			zend.CG__().GetZendLineno()++
			zend.CG__().SetIncrementLineno(0)
		}
		token_line = zend.CG__().GetZendLineno()
	}
	zend.ZendRestoreLexicalState(&original_lex_state)
	return 1
}
func OnEvent(event zend.ZendPhpScannerEvent, token int, line int, context any) {
	var token_stream *types.Zval = (*types.Zval)(context)
	var tokens_ht *types.Array
	var token_zv *types.Zval
	switch event {
	case zend.ON_TOKEN:
		if token == zend.END {
			break
		}

		/* Special cases */
		if token == ';' && zend.INI_SCNG__().GetYyLeng() > 1 {
			token = zend.T_CLOSE_TAG
		} else if token == zend.T_ECHO && zend.INI_SCNG__().GetYyLeng() == b.SizeOf("\"<?=\"")-1 {
			token = zend.T_OPEN_TAG_WITH_ECHO
		}
		AddToken(token_stream, token, zend.INI_SCNG__().GetYyText(), zend.INI_SCNG__().GetYyLeng(), line)
	case zend.ON_FEEDBACK:
		tokens_ht = token_stream.Array()
		token_zv = tokens_ht.IndexFindH(tokens_ht.GetNNumOfElements() - 1)
		if token_zv != nil && token_zv.IsType(zend.IS_ARRAY) {
			token_zv.Array().IndexFindH(0).SetLong(token)
		}
	case zend.ON_STOP:
		if zend.INI_SCNG__().GetYyCursor() != zend.INI_SCNG__().GetYyLimit() {
			AddToken(token_stream, zend.T_INLINE_HTML, zend.INI_SCNG__().GetYyCursor(), zend.INI_SCNG__().GetYyLimit()-zend.INI_SCNG__().GetYyCursor(), zend.CG__().GetZendLineno())
		}
	}
}
func TokenizeParse(return_value *types.Zval, source string) bool {
	var source_zval types.Zval
	var original_lex_state zend.ZendLexState
	var original_in_compilation bool
	var success bool
	source_zval.SetStringVal(source)
	original_in_compilation = zend.CG__().GetInCompilation()
	zend.CG__().SetInCompilation(1)
	zend.ZendSaveLexicalState(&original_lex_state)
	if b.Assign(&success, zend.ZendPrepareStringForScanning(&source_zval, "") == zend.SUCCESS) {
		var token_stream types.Zval
		zend.ArrayInit(&token_stream)
		zend.CG__().SetAst(nil)
		zend.LANG_SCNG__().SetYyState(zend.yycINITIAL)
		zend.LANG_SCNG__().on_event = OnEvent
		zend.LANG_SCNG__().on_event_context = &token_stream
		if b.Assign(&success, zend.Zendparse() == zend.SUCCESS) {
			zend.ZVAL_COPY_VALUE(return_value, &token_stream)
		} else {
			types.ZvalPtrDtor(&token_stream)
		}
		zend.ZendAstDestroy(zend.CG__().GetAst())
	}

	/* restore compiler and scanner global states */

	zend.ZendRestoreLexicalState(&original_lex_state)
	zend.CG__().SetInCompilation(original_in_compilation)
	types.ZvalPtrDtorStr(&source_zval)
	return success
}

func ZifTokenGetAll(return_value zpp.Ret, source string, _ zpp.Opt, flags int) {
	var success bool
	if (flags & 1) != 0 {
		success = TokenizeParse(return_value, source)
	} else {
		success = Tokenize(return_value, source)

		/* Normal token_get_all() should not throw. */
		faults.ClearException()
	}
	if success == 0 {
		return_value.SetFalse()
		return
	}
}
func ZifTokenName(token int) string {
	return GetTokenTypeName(token)
}
