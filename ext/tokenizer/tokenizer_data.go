package tokenizer

import (
	"github.com/heyuuu/gophp/zend"
)

func TokenizerRegisterConstants(type_ int, module_number int) {
	zend.RegisterLongConstant("T_INCLUDE", zend.T_INCLUDE, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_INCLUDE_ONCE", zend.T_INCLUDE_ONCE, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_REQUIRE", zend.T_REQUIRE, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_REQUIRE_ONCE", zend.T_REQUIRE_ONCE, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_LOGICAL_OR", zend.T_LOGICAL_OR, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_LOGICAL_XOR", zend.T_LOGICAL_XOR, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_LOGICAL_AND", zend.T_LOGICAL_AND, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_PRINT", zend.T_PRINT, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_YIELD", zend.T_YIELD, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_DOUBLE_ARROW", zend.T_DOUBLE_ARROW, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_YIELD_FROM", zend.T_YIELD_FROM, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_PLUS_EQUAL", zend.T_PLUS_EQUAL, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_MINUS_EQUAL", zend.T_MINUS_EQUAL, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_MUL_EQUAL", zend.T_MUL_EQUAL, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_DIV_EQUAL", zend.T_DIV_EQUAL, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_CONCAT_EQUAL", zend.T_CONCAT_EQUAL, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_MOD_EQUAL", zend.T_MOD_EQUAL, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_AND_EQUAL", zend.T_AND_EQUAL, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_OR_EQUAL", zend.T_OR_EQUAL, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_XOR_EQUAL", zend.T_XOR_EQUAL, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_SL_EQUAL", zend.T_SL_EQUAL, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_SR_EQUAL", zend.T_SR_EQUAL, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_POW_EQUAL", zend.T_POW_EQUAL, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_COALESCE_EQUAL", zend.T_COALESCE_EQUAL, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_COALESCE", zend.T_COALESCE, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_BOOLEAN_OR", zend.T_BOOLEAN_OR, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_BOOLEAN_AND", zend.T_BOOLEAN_AND, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_IS_EQUAL", zend.T_IS_EQUAL, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_IS_NOT_EQUAL", zend.T_IS_NOT_EQUAL, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_IS_IDENTICAL", zend.T_IS_IDENTICAL, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_IS_NOT_IDENTICAL", zend.T_IS_NOT_IDENTICAL, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_SPACESHIP", zend.T_SPACESHIP, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_IS_SMALLER_OR_EQUAL", zend.T_IS_SMALLER_OR_EQUAL, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_IS_GREATER_OR_EQUAL", zend.T_IS_GREATER_OR_EQUAL, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_SL", zend.T_SL, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_SR", zend.T_SR, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_INSTANCEOF", zend.T_INSTANCEOF, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_INT_CAST", zend.T_INT_CAST, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_DOUBLE_CAST", zend.T_DOUBLE_CAST, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_STRING_CAST", zend.T_STRING_CAST, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_ARRAY_CAST", zend.T_ARRAY_CAST, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_OBJECT_CAST", zend.T_OBJECT_CAST, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_BOOL_CAST", zend.T_BOOL_CAST, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_UNSET_CAST", zend.T_UNSET_CAST, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_POW", zend.T_POW, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_NEW", zend.T_NEW, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_CLONE", zend.T_CLONE, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_ELSEIF", zend.T_ELSEIF, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_ELSE", zend.T_ELSE, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_LNUMBER", zend.T_LNUMBER, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_DNUMBER", zend.T_DNUMBER, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_STRING", zend.T_STRING, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_VARIABLE", zend.T_VARIABLE, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_INLINE_HTML", zend.T_INLINE_HTML, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_ENCAPSED_AND_WHITESPACE", zend.T_ENCAPSED_AND_WHITESPACE, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_CONSTANT_ENCAPSED_STRING", zend.T_CONSTANT_ENCAPSED_STRING, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_STRING_VARNAME", zend.T_STRING_VARNAME, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_NUM_STRING", zend.T_NUM_STRING, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_EVAL", zend.T_EVAL, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_INC", zend.T_INC, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_DEC", zend.T_DEC, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_EXIT", zend.T_EXIT, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_IF", zend.T_IF, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_ENDIF", zend.T_ENDIF, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_ECHO", zend.T_ECHO, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_DO", zend.T_DO, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_WHILE", zend.T_WHILE, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_ENDWHILE", zend.T_ENDWHILE, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_FOR", zend.T_FOR, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_ENDFOR", zend.T_ENDFOR, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_FOREACH", zend.T_FOREACH, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_ENDFOREACH", zend.T_ENDFOREACH, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_DECLARE", zend.T_DECLARE, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_ENDDECLARE", zend.T_ENDDECLARE, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_AS", zend.T_AS, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_SWITCH", zend.T_SWITCH, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_ENDSWITCH", zend.T_ENDSWITCH, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_CASE", zend.T_CASE, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_DEFAULT", zend.T_DEFAULT, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_BREAK", zend.T_BREAK, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_CONTINUE", zend.T_CONTINUE, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_GOTO", zend.T_GOTO, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_FUNCTION", zend.T_FUNCTION, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_FN", zend.T_FN, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_CONST", zend.T_CONST, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_RETURN", zend.T_RETURN, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_TRY", zend.T_TRY, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_CATCH", zend.T_CATCH, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_FINALLY", zend.T_FINALLY, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_THROW", zend.T_THROW, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_USE", zend.T_USE, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_INSTEADOF", zend.T_INSTEADOF, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_GLOBAL", zend.T_GLOBAL, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_STATIC", zend.T_STATIC, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_ABSTRACT", zend.T_ABSTRACT, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_FINAL", zend.T_FINAL, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_PRIVATE", zend.T_PRIVATE, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_PROTECTED", zend.T_PROTECTED, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_PUBLIC", zend.T_PUBLIC, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_VAR", zend.T_VAR, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_UNSET", zend.T_UNSET, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_ISSET", zend.T_ISSET, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_EMPTY", zend.T_EMPTY, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_HALT_COMPILER", zend.T_HALT_COMPILER, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_CLASS", zend.T_CLASS, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_TRAIT", zend.T_TRAIT, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_INTERFACE", zend.T_INTERFACE, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_EXTENDS", zend.T_EXTENDS, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_IMPLEMENTS", zend.T_IMPLEMENTS, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_OBJECT_OPERATOR", zend.T_OBJECT_OPERATOR, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_LIST", zend.T_LIST, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_ARRAY", zend.T_ARRAY, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_CALLABLE", zend.T_CALLABLE, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_LINE", zend.T_LINE, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_FILE", zend.T_FILE, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_DIR", zend.T_DIR, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_CLASS_C", zend.T_CLASS_C, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_TRAIT_C", zend.T_TRAIT_C, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_METHOD_C", zend.T_METHOD_C, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_FUNC_C", zend.T_FUNC_C, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_COMMENT", zend.T_COMMENT, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_DOC_COMMENT", zend.T_DOC_COMMENT, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_OPEN_TAG", zend.T_OPEN_TAG, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_OPEN_TAG_WITH_ECHO", zend.T_OPEN_TAG_WITH_ECHO, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_CLOSE_TAG", zend.T_CLOSE_TAG, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_WHITESPACE", zend.T_WHITESPACE, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_START_HEREDOC", zend.T_START_HEREDOC, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_END_HEREDOC", zend.T_END_HEREDOC, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_DOLLAR_OPEN_CURLY_BRACES", zend.T_DOLLAR_OPEN_CURLY_BRACES, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_CURLY_OPEN", zend.T_CURLY_OPEN, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_PAAMAYIM_NEKUDOTAYIM", zend.T_PAAMAYIM_NEKUDOTAYIM, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_NAMESPACE", zend.T_NAMESPACE, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_NS_C", zend.T_NS_C, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_NS_SEPARATOR", zend.T_NS_SEPARATOR, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_ELLIPSIS", zend.T_ELLIPSIS, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_BAD_CHARACTER", zend.T_BAD_CHARACTER, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("T_DOUBLE_COLON", zend.T_PAAMAYIM_NEKUDOTAYIM, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
}
func GetTokenTypeName(token_type int) string {
	switch token_type {
	case zend.T_INCLUDE:
		return "T_INCLUDE"
	case zend.T_INCLUDE_ONCE:
		return "T_INCLUDE_ONCE"
	case zend.T_REQUIRE:
		return "T_REQUIRE"
	case zend.T_REQUIRE_ONCE:
		return "T_REQUIRE_ONCE"
	case zend.T_LOGICAL_OR:
		return "T_LOGICAL_OR"
	case zend.T_LOGICAL_XOR:
		return "T_LOGICAL_XOR"
	case zend.T_LOGICAL_AND:
		return "T_LOGICAL_AND"
	case zend.T_PRINT:
		return "T_PRINT"
	case zend.T_YIELD:
		return "T_YIELD"
	case zend.T_DOUBLE_ARROW:
		return "T_DOUBLE_ARROW"
	case zend.T_YIELD_FROM:
		return "T_YIELD_FROM"
	case zend.T_PLUS_EQUAL:
		return "T_PLUS_EQUAL"
	case zend.T_MINUS_EQUAL:
		return "T_MINUS_EQUAL"
	case zend.T_MUL_EQUAL:
		return "T_MUL_EQUAL"
	case zend.T_DIV_EQUAL:
		return "T_DIV_EQUAL"
	case zend.T_CONCAT_EQUAL:
		return "T_CONCAT_EQUAL"
	case zend.T_MOD_EQUAL:
		return "T_MOD_EQUAL"
	case zend.T_AND_EQUAL:
		return "T_AND_EQUAL"
	case zend.T_OR_EQUAL:
		return "T_OR_EQUAL"
	case zend.T_XOR_EQUAL:
		return "T_XOR_EQUAL"
	case zend.T_SL_EQUAL:
		return "T_SL_EQUAL"
	case zend.T_SR_EQUAL:
		return "T_SR_EQUAL"
	case zend.T_POW_EQUAL:
		return "T_POW_EQUAL"
	case zend.T_COALESCE_EQUAL:
		return "T_COALESCE_EQUAL"
	case zend.T_COALESCE:
		return "T_COALESCE"
	case zend.T_BOOLEAN_OR:
		return "T_BOOLEAN_OR"
	case zend.T_BOOLEAN_AND:
		return "T_BOOLEAN_AND"
	case zend.T_IS_EQUAL:
		return "T_IS_EQUAL"
	case zend.T_IS_NOT_EQUAL:
		return "T_IS_NOT_EQUAL"
	case zend.T_IS_IDENTICAL:
		return "T_IS_IDENTICAL"
	case zend.T_IS_NOT_IDENTICAL:
		return "T_IS_NOT_IDENTICAL"
	case zend.T_SPACESHIP:
		return "T_SPACESHIP"
	case zend.T_IS_SMALLER_OR_EQUAL:
		return "T_IS_SMALLER_OR_EQUAL"
	case zend.T_IS_GREATER_OR_EQUAL:
		return "T_IS_GREATER_OR_EQUAL"
	case zend.T_SL:
		return "T_SL"
	case zend.T_SR:
		return "T_SR"
	case zend.T_INSTANCEOF:
		return "T_INSTANCEOF"
	case zend.T_INT_CAST:
		return "T_INT_CAST"
	case zend.T_DOUBLE_CAST:
		return "T_DOUBLE_CAST"
	case zend.T_STRING_CAST:
		return "T_STRING_CAST"
	case zend.T_ARRAY_CAST:
		return "T_ARRAY_CAST"
	case zend.T_OBJECT_CAST:
		return "T_OBJECT_CAST"
	case zend.T_BOOL_CAST:
		return "T_BOOL_CAST"
	case zend.T_UNSET_CAST:
		return "T_UNSET_CAST"
	case zend.T_POW:
		return "T_POW"
	case zend.T_NEW:
		return "T_NEW"
	case zend.T_CLONE:
		return "T_CLONE"
	case zend.T_ELSEIF:
		return "T_ELSEIF"
	case zend.T_ELSE:
		return "T_ELSE"
	case zend.T_LNUMBER:
		return "T_LNUMBER"
	case zend.T_DNUMBER:
		return "T_DNUMBER"
	case zend.T_STRING:
		return "T_STRING"
	case zend.T_VARIABLE:
		return "T_VARIABLE"
	case zend.T_INLINE_HTML:
		return "T_INLINE_HTML"
	case zend.T_ENCAPSED_AND_WHITESPACE:
		return "T_ENCAPSED_AND_WHITESPACE"
	case zend.T_CONSTANT_ENCAPSED_STRING:
		return "T_CONSTANT_ENCAPSED_STRING"
	case zend.T_STRING_VARNAME:
		return "T_STRING_VARNAME"
	case zend.T_NUM_STRING:
		return "T_NUM_STRING"
	case zend.T_EVAL:
		return "T_EVAL"
	case zend.T_INC:
		return "T_INC"
	case zend.T_DEC:
		return "T_DEC"
	case zend.T_EXIT:
		return "T_EXIT"
	case zend.T_IF:
		return "T_IF"
	case zend.T_ENDIF:
		return "T_ENDIF"
	case zend.T_ECHO:
		return "T_ECHO"
	case zend.T_DO:
		return "T_DO"
	case zend.T_WHILE:
		return "T_WHILE"
	case zend.T_ENDWHILE:
		return "T_ENDWHILE"
	case zend.T_FOR:
		return "T_FOR"
	case zend.T_ENDFOR:
		return "T_ENDFOR"
	case zend.T_FOREACH:
		return "T_FOREACH"
	case zend.T_ENDFOREACH:
		return "T_ENDFOREACH"
	case zend.T_DECLARE:
		return "T_DECLARE"
	case zend.T_ENDDECLARE:
		return "T_ENDDECLARE"
	case zend.T_AS:
		return "T_AS"
	case zend.T_SWITCH:
		return "T_SWITCH"
	case zend.T_ENDSWITCH:
		return "T_ENDSWITCH"
	case zend.T_CASE:
		return "T_CASE"
	case zend.T_DEFAULT:
		return "T_DEFAULT"
	case zend.T_BREAK:
		return "T_BREAK"
	case zend.T_CONTINUE:
		return "T_CONTINUE"
	case zend.T_GOTO:
		return "T_GOTO"
	case zend.T_FUNCTION:
		return "T_FUNCTION"
	case zend.T_FN:
		return "T_FN"
	case zend.T_CONST:
		return "T_CONST"
	case zend.T_RETURN:
		return "T_RETURN"
	case zend.T_TRY:
		return "T_TRY"
	case zend.T_CATCH:
		return "T_CATCH"
	case zend.T_FINALLY:
		return "T_FINALLY"
	case zend.T_THROW:
		return "T_THROW"
	case zend.T_USE:
		return "T_USE"
	case zend.T_INSTEADOF:
		return "T_INSTEADOF"
	case zend.T_GLOBAL:
		return "T_GLOBAL"
	case zend.T_STATIC:
		return "T_STATIC"
	case zend.T_ABSTRACT:
		return "T_ABSTRACT"
	case zend.T_FINAL:
		return "T_FINAL"
	case zend.T_PRIVATE:
		return "T_PRIVATE"
	case zend.T_PROTECTED:
		return "T_PROTECTED"
	case zend.T_PUBLIC:
		return "T_PUBLIC"
	case zend.T_VAR:
		return "T_VAR"
	case zend.T_UNSET:
		return "T_UNSET"
	case zend.T_ISSET:
		return "T_ISSET"
	case zend.T_EMPTY:
		return "T_EMPTY"
	case zend.T_HALT_COMPILER:
		return "T_HALT_COMPILER"
	case zend.T_CLASS:
		return "T_CLASS"
	case zend.T_TRAIT:
		return "T_TRAIT"
	case zend.T_INTERFACE:
		return "T_INTERFACE"
	case zend.T_EXTENDS:
		return "T_EXTENDS"
	case zend.T_IMPLEMENTS:
		return "T_IMPLEMENTS"
	case zend.T_OBJECT_OPERATOR:
		return "T_OBJECT_OPERATOR"
	case zend.T_LIST:
		return "T_LIST"
	case zend.T_ARRAY:
		return "T_ARRAY"
	case zend.T_CALLABLE:
		return "T_CALLABLE"
	case zend.T_LINE:
		return "T_LINE"
	case zend.T_FILE:
		return "T_FILE"
	case zend.T_DIR:
		return "T_DIR"
	case zend.T_CLASS_C:
		return "T_CLASS_C"
	case zend.T_TRAIT_C:
		return "T_TRAIT_C"
	case zend.T_METHOD_C:
		return "T_METHOD_C"
	case zend.T_FUNC_C:
		return "T_FUNC_C"
	case zend.T_COMMENT:
		return "T_COMMENT"
	case zend.T_DOC_COMMENT:
		return "T_DOC_COMMENT"
	case zend.T_OPEN_TAG:
		return "T_OPEN_TAG"
	case zend.T_OPEN_TAG_WITH_ECHO:
		return "T_OPEN_TAG_WITH_ECHO"
	case zend.T_CLOSE_TAG:
		return "T_CLOSE_TAG"
	case zend.T_WHITESPACE:
		return "T_WHITESPACE"
	case zend.T_START_HEREDOC:
		return "T_START_HEREDOC"
	case zend.T_END_HEREDOC:
		return "T_END_HEREDOC"
	case zend.T_DOLLAR_OPEN_CURLY_BRACES:
		return "T_DOLLAR_OPEN_CURLY_BRACES"
	case zend.T_CURLY_OPEN:
		return "T_CURLY_OPEN"
	case zend.T_PAAMAYIM_NEKUDOTAYIM:
		return "T_DOUBLE_COLON"
	case zend.T_NAMESPACE:
		return "T_NAMESPACE"
	case zend.T_NS_C:
		return "T_NS_C"
	case zend.T_NS_SEPARATOR:
		return "T_NS_SEPARATOR"
	case zend.T_ELLIPSIS:
		return "T_ELLIPSIS"
	case zend.T_BAD_CHARACTER:
		return "T_BAD_CHARACTER"
	}
	return "UNKNOWN"
}
