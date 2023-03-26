package zend

const ZEND_AST_SPEC = 1
const ZEND_AST_SPECIAL_SHIFT = 6
const ZEND_AST_IS_LIST_SHIFT = 7
const ZEND_AST_NUM_CHILDREN_SHIFT = 8

type _zendAstKind = int

const (
	ZEND_AST_ZVAL _zendAstKind = 1 << ZEND_AST_SPECIAL_SHIFT
	ZEND_AST_CONSTANT
	ZEND_AST_ZNODE
	ZEND_AST_FUNC_DECL
	ZEND_AST_CLOSURE
	ZEND_AST_METHOD
	ZEND_AST_CLASS
	ZEND_AST_ARROW_FUNC
	ZEND_AST_ARG_LIST _zendAstKind = 1 << ZEND_AST_IS_LIST_SHIFT
	ZEND_AST_ARRAY
	ZEND_AST_ENCAPS_LIST
	ZEND_AST_EXPR_LIST
	ZEND_AST_STMT_LIST
	ZEND_AST_IF
	ZEND_AST_SWITCH_LIST
	ZEND_AST_CATCH_LIST
	ZEND_AST_PARAM_LIST
	ZEND_AST_CLOSURE_USES
	ZEND_AST_PROP_DECL
	ZEND_AST_CONST_DECL
	ZEND_AST_CLASS_CONST_DECL
	ZEND_AST_NAME_LIST
	ZEND_AST_TRAIT_ADAPTATIONS
	ZEND_AST_USE
	ZEND_AST_MAGIC_CONST _zendAstKind = 0 << ZEND_AST_NUM_CHILDREN_SHIFT
	ZEND_AST_TYPE
	ZEND_AST_CONSTANT_CLASS
	ZEND_AST_VAR _zendAstKind = 1 << ZEND_AST_NUM_CHILDREN_SHIFT
	ZEND_AST_CONST
	ZEND_AST_UNPACK
	ZEND_AST_UNARY_PLUS
	ZEND_AST_UNARY_MINUS
	ZEND_AST_CAST
	ZEND_AST_EMPTY
	ZEND_AST_ISSET
	ZEND_AST_SILENCE
	ZEND_AST_SHELL_EXEC
	ZEND_AST_CLONE
	ZEND_AST_EXIT
	ZEND_AST_PRINT
	ZEND_AST_INCLUDE_OR_EVAL
	ZEND_AST_UNARY_OP
	ZEND_AST_PRE_INC
	ZEND_AST_PRE_DEC
	ZEND_AST_POST_INC
	ZEND_AST_POST_DEC
	ZEND_AST_YIELD_FROM
	ZEND_AST_CLASS_NAME
	ZEND_AST_GLOBAL
	ZEND_AST_UNSET
	ZEND_AST_RETURN
	ZEND_AST_LABEL
	ZEND_AST_REF
	ZEND_AST_HALT_COMPILER
	ZEND_AST_ECHO
	ZEND_AST_THROW
	ZEND_AST_GOTO
	ZEND_AST_BREAK
	ZEND_AST_CONTINUE
	ZEND_AST_DIM _zendAstKind = 2 << ZEND_AST_NUM_CHILDREN_SHIFT
	ZEND_AST_PROP
	ZEND_AST_STATIC_PROP
	ZEND_AST_CALL
	ZEND_AST_CLASS_CONST
	ZEND_AST_ASSIGN
	ZEND_AST_ASSIGN_REF
	ZEND_AST_ASSIGN_OP
	ZEND_AST_BINARY_OP
	ZEND_AST_GREATER
	ZEND_AST_GREATER_EQUAL
	ZEND_AST_AND
	ZEND_AST_OR
	ZEND_AST_ARRAY_ELEM
	ZEND_AST_NEW
	ZEND_AST_INSTANCEOF
	ZEND_AST_YIELD
	ZEND_AST_COALESCE
	ZEND_AST_ASSIGN_COALESCE
	ZEND_AST_STATIC
	ZEND_AST_WHILE
	ZEND_AST_DO_WHILE
	ZEND_AST_IF_ELEM
	ZEND_AST_SWITCH
	ZEND_AST_SWITCH_CASE
	ZEND_AST_DECLARE
	ZEND_AST_USE_TRAIT
	ZEND_AST_TRAIT_PRECEDENCE
	ZEND_AST_METHOD_REFERENCE
	ZEND_AST_NAMESPACE
	ZEND_AST_USE_ELEM
	ZEND_AST_TRAIT_ALIAS
	ZEND_AST_GROUP_USE
	ZEND_AST_PROP_GROUP
	ZEND_AST_METHOD_CALL _zendAstKind = 3 << ZEND_AST_NUM_CHILDREN_SHIFT
	ZEND_AST_STATIC_CALL
	ZEND_AST_CONDITIONAL
	ZEND_AST_TRY
	ZEND_AST_CATCH
	ZEND_AST_PARAM
	ZEND_AST_PROP_ELEM
	ZEND_AST_CONST_ELEM
	ZEND_AST_FOR _zendAstKind = 4 << ZEND_AST_NUM_CHILDREN_SHIFT
	ZEND_AST_FOREACH
)

type ZendAstKind = uint16
type ZendAstAttr = uint16

/* Same as zend_ast, but with children count, which is updated dynamically */

/* Lineno is stored in val.u2.lineno */

/* Separate structure for function and class declaration, as they need extra information. */

type ZendAstProcessT func(ast *ZendAst)
type ZendAstApplyFunc func(ast_ptr **ZendAst)

var ZendAstProcess ZendAstProcessT = nil

/*
 * Operator Precedence
 * ====================
 * priority  associativity  operators
 * ----------------------------------
 *   10     left            include, include_once, eval, require, require_once
 *   20     left            ,
 *   30     left            or
 *   40     left            xor
 *   50     left            and
 *   60     right           print
 *   70     right           yield
 *   80     right           =>
 *   85     right           yield from
 *   90     right           = += -= *= /= .= %= &= |= ^= <<= >>= **=
 *  100     left            ? :
 *  110     right           ??
 *  120     left            ||
 *  130     left            &&
 *  140     left            |
 *  150     left            ^
 *  160     left            &
 *  170     non-associative == != === !==
 *  180     non-associative < <= > >= <=>
 *  190     left            << >>
 *  200     left            + - .
 *  210     left            * / %
 *  220     right           !
 *  230     non-associative instanceof
 *  240     right           + - ++ -- ~ (type) @
 *  250     right           **
 *  260     left            [
 *  270     non-associative clone new
 */
