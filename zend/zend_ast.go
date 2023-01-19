// <<generate>>

package zend

import g "sik/runtime/grammar"

// Source: <Zend/zend_ast.h>

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
   | Authors: Bob Weinand <bwoebi@php.net>                                |
   |          Dmitry Stogov <dmitry@php.net>                              |
   |          Nikita Popov <nikic@php.net>                                |
   +----------------------------------------------------------------------+
*/

// #define ZEND_AST_H

// # include "zend.h"

// #define ZEND_AST_SPEC       1

// #define ZEND_AST_SPECIAL_SHIFT       6

// #define ZEND_AST_IS_LIST_SHIFT       7

// #define ZEND_AST_NUM_CHILDREN_SHIFT       8

type _zendAstKind = int

const (
	ZEND_AST_ZVAL _zendAstKind = 1 << 6
	ZEND_AST_CONSTANT
	ZEND_AST_ZNODE
	ZEND_AST_FUNC_DECL
	ZEND_AST_CLOSURE
	ZEND_AST_METHOD
	ZEND_AST_CLASS
	ZEND_AST_ARROW_FUNC
	ZEND_AST_ARG_LIST _zendAstKind = 1 << 7
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
	ZEND_AST_MAGIC_CONST _zendAstKind = 0 << 8
	ZEND_AST_TYPE
	ZEND_AST_CONSTANT_CLASS
	ZEND_AST_VAR _zendAstKind = 1 << 8
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
	ZEND_AST_DIM _zendAstKind = 2 << 8
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
	ZEND_AST_METHOD_CALL _zendAstKind = 3 << 8
	ZEND_AST_STATIC_CALL
	ZEND_AST_CONDITIONAL
	ZEND_AST_TRY
	ZEND_AST_CATCH
	ZEND_AST_PARAM
	ZEND_AST_PROP_ELEM
	ZEND_AST_CONST_ELEM
	ZEND_AST_FOR _zendAstKind = 4 << 8
	ZEND_AST_FOREACH
)

type ZendAstKind = uint16
type ZendAstAttr = uint16

// @type ZendAst struct

/* Same as zend_ast, but with children count, which is updated dynamically */

// @type ZendAstList struct

/* Lineno is stored in val.u2.lineno */

// @type ZendAstZval struct

/* Separate structure for function and class declaration, as they need extra information. */

// @type ZendAstDecl struct

type ZendAstProcessT func(ast *ZendAst)

// #define ZEND_AST_SPEC_CALL(name) ZEND_EXPAND_VA ( ZEND_AST_SPEC_CALL_ ( name , __VA_ARGS__ , _4 , _3 , _2 , _1 , _0 ) ( __VA_ARGS__ ) )

// #define ZEND_AST_SPEC_CALL_(name,_,_4,_3,_2,_1,suffix) name ## suffix

// #define ZEND_AST_SPEC_CALL_EX(name) ZEND_EXPAND_VA ( ZEND_AST_SPEC_CALL_EX_ ( name , __VA_ARGS__ , _4 , _3 , _2 , _1 , _0 ) ( __VA_ARGS__ ) )

// #define ZEND_AST_SPEC_CALL_EX_(name,_,_5,_4,_3,_2,_1,suffix) name ## suffix

func ZendAstCreateEx0(kind ZendAstKind, attr ZendAstAttr) *ZendAst {
	var ast *ZendAst = ZendAstCreate0(kind)
	ast.SetAttr(attr)
	return ast
}
func ZendAstCreateEx1(kind ZendAstKind, attr ZendAstAttr, child *ZendAst) *ZendAst {
	var ast *ZendAst = ZendAstCreate1(kind, child)
	ast.SetAttr(attr)
	return ast
}
func ZendAstCreateEx2(kind ZendAstKind, attr ZendAstAttr, child1 *ZendAst, child2 *ZendAst) *ZendAst {
	var ast *ZendAst = ZendAstCreate2(kind, child1, child2)
	ast.SetAttr(attr)
	return ast
}
func ZendAstCreateEx3(kind ZendAstKind, attr ZendAstAttr, child1 *ZendAst, child2 *ZendAst, child3 *ZendAst) *ZendAst {
	var ast *ZendAst = ZendAstCreate3(kind, child1, child2, child3)
	ast.SetAttr(attr)
	return ast
}
func ZendAstCreateEx4(kind ZendAstKind, attr ZendAstAttr, child1 *ZendAst, child2 *ZendAst, child3 *ZendAst, child4 *ZendAst) *ZendAst {
	var ast *ZendAst = ZendAstCreate4(kind, child1, child2, child3, child4)
	ast.SetAttr(attr)
	return ast
}

// #define zend_ast_create() ZEND_AST_SPEC_CALL ( zend_ast_create , __VA_ARGS__ )

// #define zend_ast_create_ex() ZEND_AST_SPEC_CALL_EX ( zend_ast_create_ex , __VA_ARGS__ )

// #define zend_ast_create_list(init_children) ZEND_AST_SPEC_CALL ( zend_ast_create_list , __VA_ARGS__ )

type ZendAstApplyFunc func(ast_ptr **ZendAst)

func ZendAstIsSpecial(ast *ZendAst) ZendBool { return ast.GetKind() >> 6 & 1 }
func ZendAstIsList(ast *ZendAst) ZendBool    { return ast.GetKind() >> 7 & 1 }
func ZendAstGetList(ast *ZendAst) *ZendAstList {
	assert(ZendAstIsList(ast) != 0)
	return (*ZendAstList)(ast)
}
func ZendAstGetZval(ast *ZendAst) *Zval {
	assert(ast.GetKind() == ZEND_AST_ZVAL)
	return &((*ZendAstZval)(ast)).val
}
func ZendAstGetStr(ast *ZendAst) *ZendString {
	var zv *Zval = ZendAstGetZval(ast)
	assert(zv.GetType() == 6)
	return zv.GetValue().GetStr()
}
func ZendAstGetConstantName(ast *ZendAst) *ZendString {
	assert(ast.GetKind() == ZEND_AST_CONSTANT)
	assert((*ZendAstZval)(ast).GetVal().GetType() == 6)
	return (*ZendAstZval)(ast).GetVal().GetValue().GetStr()
}
func ZendAstGetNumChildren(ast *ZendAst) uint32 {
	assert(ZendAstIsList(ast) == 0)
	return ast.GetKind() >> 8
}
func ZendAstGetLineno(ast *ZendAst) uint32 {
	if ast.GetKind() == ZEND_AST_ZVAL {
		var zv *Zval = ZendAstGetZval(ast)
		return zv.GetLineno()
	} else {
		return ast.GetLineno()
	}
}
func ZendAstCreateBinaryOp(opcode uint32, op0 *ZendAst, op1 *ZendAst) *ZendAst {
	return ZendAstCreateEx2(ZEND_AST_BINARY_OP, opcode, op0, op1)
}
func ZendAstCreateAssignOp(opcode uint32, op0 *ZendAst, op1 *ZendAst) *ZendAst {
	return ZendAstCreateEx2(ZEND_AST_ASSIGN_OP, opcode, op0, op1)
}
func ZendAstCreateCast(type_ uint32, op0 *ZendAst) *ZendAst {
	return ZendAstCreateEx1(ZEND_AST_CAST, type_, op0)
}
func ZendAstListRtrim(ast *ZendAst) *ZendAst {
	var list *ZendAstList = ZendAstGetList(ast)
	if list.GetChildren() != 0 && list.GetChild()[list.GetChildren()-1] == nil {
		list.GetChildren()--
	}
	return ast
}

// Source: <Zend/zend_ast.c>

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
   | Authors: Bob Weinand <bwoebi@php.net>                                |
   |          Dmitry Stogov <dmitry@php.net>                              |
   +----------------------------------------------------------------------+
*/

// # include "zend_ast.h"

// # include "zend_API.h"

// # include "zend_operators.h"

// # include "zend_language_parser.h"

// # include "zend_smart_str.h"

// # include "zend_exceptions.h"

// # include "zend_constants.h"

var ZendAstProcess ZendAstProcessT = nil

func ZendAstAlloc(size int) any {
	return ZendArenaAlloc(&CG.ast_arena, size)
}
func ZendAstRealloc(old any, old_size int, new_size int) any {
	var new_ any = ZendAstAlloc(new_size)
	memcpy(new_, old, old_size)
	return new_
}
func ZendAstSize(children uint32) int {
	return g.SizeOf("zend_ast") - g.SizeOf("zend_ast *") + g.SizeOf("zend_ast *")*children
}
func ZendAstListSize(children uint32) int {
	return g.SizeOf("zend_ast_list") - g.SizeOf("zend_ast *") + g.SizeOf("zend_ast *")*children
}
func ZendAstCreateZnode(node *Znode) *ZendAst {
	var ast *ZendAstZnode
	ast = ZendAstAlloc(g.SizeOf("zend_ast_znode"))
	ast.SetKind(ZEND_AST_ZNODE)
	ast.SetAttr(0)
	ast.SetLineno(CG.GetZendLineno())
	ast.SetNode(*node)
	return (*ZendAst)(ast)
}
func ZendAstCreateZvalInt(zv *Zval, attr uint32, lineno uint32) *ZendAst {
	var ast *ZendAstZval
	ast = ZendAstAlloc(g.SizeOf("zend_ast_zval"))
	ast.SetKind(ZEND_AST_ZVAL)
	ast.SetAttr(attr)
	var _z1 *Zval = &ast.val
	var _z2 *Zval = zv
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	ast.GetVal().SetLineno(lineno)
	return (*ZendAst)(ast)
}
func ZendAstCreateZvalWithLineno(zv *Zval, lineno uint32) *ZendAst {
	return ZendAstCreateZvalInt(zv, 0, lineno)
}
func ZendAstCreateZvalEx(zv *Zval, attr ZendAstAttr) *ZendAst {
	return ZendAstCreateZvalInt(zv, attr, CG.GetZendLineno())
}
func ZendAstCreateZval(zv *Zval) *ZendAst {
	return ZendAstCreateZvalInt(zv, 0, CG.GetZendLineno())
}
func ZendAstCreateZvalFromStr(str *ZendString) *ZendAst {
	var zv Zval
	var __z *Zval = &zv
	var __s *ZendString = str
	__z.GetValue().SetStr(__s)
	if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
		__z.SetTypeInfo(6)
	} else {
		__z.SetTypeInfo(6 | 1<<0<<8)
	}
	return ZendAstCreateZvalInt(&zv, 0, CG.GetZendLineno())
}
func ZendAstCreateZvalFromLong(lval ZendLong) *ZendAst {
	var zv Zval
	var __z *Zval = &zv
	__z.GetValue().SetLval(lval)
	__z.SetTypeInfo(4)
	return ZendAstCreateZvalInt(&zv, 0, CG.GetZendLineno())
}
func ZendAstCreateConstant(name *ZendString, attr ZendAstAttr) *ZendAst {
	var ast *ZendAstZval
	ast = ZendAstAlloc(g.SizeOf("zend_ast_zval"))
	ast.SetKind(ZEND_AST_CONSTANT)
	ast.SetAttr(attr)
	var __z *Zval = &ast.val
	var __s *ZendString = name
	__z.GetValue().SetStr(__s)
	if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
		__z.SetTypeInfo(6)
	} else {
		__z.SetTypeInfo(6 | 1<<0<<8)
	}
	ast.GetVal().SetLineno(CG.GetZendLineno())
	return (*ZendAst)(ast)
}
func ZendAstCreateClassConstOrName(class_name *ZendAst, name *ZendAst) *ZendAst {
	var name_str *ZendString = ZendAstGetStr(name)
	if name_str.GetLen() == g.SizeOf("\"class\"")-1 && ZendBinaryStrcasecmp(name_str.GetVal(), name_str.GetLen(), "class", g.SizeOf("\"class\"")-1) == 0 {
		ZendStringRelease(name_str)
		return ZendAstCreate1(ZEND_AST_CLASS_NAME, class_name)
	} else {
		return ZendAstCreate2(ZEND_AST_CLASS_CONST, class_name, name)
	}
}
func ZendAstCreateDecl(kind ZendAstKind, flags uint32, start_lineno uint32, doc_comment *ZendString, name *ZendString, child0 *ZendAst, child1 *ZendAst, child2 *ZendAst, child3 *ZendAst) *ZendAst {
	var ast *ZendAstDecl
	ast = ZendAstAlloc(g.SizeOf("zend_ast_decl"))
	ast.SetKind(kind)
	ast.SetAttr(0)
	ast.SetStartLineno(start_lineno)
	ast.SetEndLineno(CG.GetZendLineno())
	ast.SetFlags(flags)
	ast.SetLexPos(LANG_SCNG.GetYyText())
	ast.SetDocComment(doc_comment)
	ast.SetName(name)
	ast.GetChild()[0] = child0
	ast.GetChild()[1] = child1
	ast.GetChild()[2] = child2
	ast.GetChild()[3] = child3
	return (*ZendAst)(ast)
}
func ZendAstCreate0(kind ZendAstKind) *ZendAst {
	var ast *ZendAst
	assert(kind>>8 == 0)
	ast = ZendAstAlloc(ZendAstSize(0))
	ast.SetKind(kind)
	ast.SetAttr(0)
	ast.SetLineno(CG.GetZendLineno())
	return ast
}
func ZendAstCreate1(kind ZendAstKind, child *ZendAst) *ZendAst {
	var ast *ZendAst
	var lineno uint32
	assert(kind>>8 == 1)
	ast = ZendAstAlloc(ZendAstSize(1))
	ast.SetKind(kind)
	ast.SetAttr(0)
	ast.GetChild()[0] = child
	if child != nil {
		lineno = ZendAstGetLineno(child)
	} else {
		lineno = CG.GetZendLineno()
	}
	ast.SetLineno(lineno)
	ast.SetLineno(lineno)
	return ast
}
func ZendAstCreate2(kind ZendAstKind, child1 *ZendAst, child2 *ZendAst) *ZendAst {
	var ast *ZendAst
	var lineno uint32
	assert(kind>>8 == 2)
	ast = ZendAstAlloc(ZendAstSize(2))
	ast.SetKind(kind)
	ast.SetAttr(0)
	ast.GetChild()[0] = child1
	ast.GetChild()[1] = child2
	if child1 != nil {
		lineno = ZendAstGetLineno(child1)
	} else if child2 != nil {
		lineno = ZendAstGetLineno(child2)
	} else {
		lineno = CG.GetZendLineno()
	}
	ast.SetLineno(lineno)
	return ast
}
func ZendAstCreate3(kind ZendAstKind, child1 *ZendAst, child2 *ZendAst, child3 *ZendAst) *ZendAst {
	var ast *ZendAst
	var lineno uint32
	assert(kind>>8 == 3)
	ast = ZendAstAlloc(ZendAstSize(3))
	ast.SetKind(kind)
	ast.SetAttr(0)
	ast.GetChild()[0] = child1
	ast.GetChild()[1] = child2
	ast.GetChild()[2] = child3
	if child1 != nil {
		lineno = ZendAstGetLineno(child1)
	} else if child2 != nil {
		lineno = ZendAstGetLineno(child2)
	} else if child3 != nil {
		lineno = ZendAstGetLineno(child3)
	} else {
		lineno = CG.GetZendLineno()
	}
	ast.SetLineno(lineno)
	return ast
}
func ZendAstCreate4(kind ZendAstKind, child1 *ZendAst, child2 *ZendAst, child3 *ZendAst, child4 *ZendAst) *ZendAst {
	var ast *ZendAst
	var lineno uint32
	assert(kind>>8 == 4)
	ast = ZendAstAlloc(ZendAstSize(4))
	ast.SetKind(kind)
	ast.SetAttr(0)
	ast.GetChild()[0] = child1
	ast.GetChild()[1] = child2
	ast.GetChild()[2] = child3
	ast.GetChild()[3] = child4
	if child1 != nil {
		lineno = ZendAstGetLineno(child1)
	} else if child2 != nil {
		lineno = ZendAstGetLineno(child2)
	} else if child3 != nil {
		lineno = ZendAstGetLineno(child3)
	} else if child4 != nil {
		lineno = ZendAstGetLineno(child4)
	} else {
		lineno = CG.GetZendLineno()
	}
	ast.SetLineno(lineno)
	return ast
}
func ZendAstCreateList0(kind ZendAstKind) *ZendAst {
	var ast *ZendAst
	var list *ZendAstList
	ast = ZendAstAlloc(ZendAstListSize(4))
	list = (*ZendAstList)(ast)
	list.SetKind(kind)
	list.SetAttr(0)
	list.SetLineno(CG.GetZendLineno())
	list.SetChildren(0)
	return ast
}
func ZendAstCreateList1(kind ZendAstKind, child *ZendAst) *ZendAst {
	var ast *ZendAst
	var list *ZendAstList
	var lineno uint32
	ast = ZendAstAlloc(ZendAstListSize(4))
	list = (*ZendAstList)(ast)
	list.SetKind(kind)
	list.SetAttr(0)
	list.SetChildren(1)
	list.GetChild()[0] = child
	if child != nil {
		lineno = ZendAstGetLineno(child)
		if lineno > CG.GetZendLineno() {
			lineno = CG.GetZendLineno()
		}
	} else {
		lineno = CG.GetZendLineno()
	}
	list.SetLineno(lineno)
	return ast
}
func ZendAstCreateList2(kind ZendAstKind, child1 *ZendAst, child2 *ZendAst) *ZendAst {
	var ast *ZendAst
	var list *ZendAstList
	var lineno uint32
	ast = ZendAstAlloc(ZendAstListSize(4))
	list = (*ZendAstList)(ast)
	list.SetKind(kind)
	list.SetAttr(0)
	list.SetChildren(2)
	list.GetChild()[0] = child1
	list.GetChild()[1] = child2
	if child1 != nil {
		lineno = ZendAstGetLineno(child1)
		if lineno > CG.GetZendLineno() {
			lineno = CG.GetZendLineno()
		}
	} else if child2 != nil {
		lineno = ZendAstGetLineno(child2)
		if lineno > CG.GetZendLineno() {
			lineno = CG.GetZendLineno()
		}
	} else {
		list.SetChildren(0)
		lineno = CG.GetZendLineno()
	}
	list.SetLineno(lineno)
	return ast
}
func IsPowerOfTwo(n uint32) ZendBool { return n != 0 && n == (n & ^n + 1) }
func ZendAstListAdd(ast *ZendAst, op *ZendAst) *ZendAst {
	var list *ZendAstList = ZendAstGetList(ast)
	if list.GetChildren() >= 4 && IsPowerOfTwo(list.GetChildren()) != 0 {
		list = ZendAstRealloc(list, ZendAstListSize(list.GetChildren()), ZendAstListSize(list.GetChildren()*2))
	}
	list.GetChild()[g.PostInc(&(list.GetChildren()))] = op
	return (*ZendAst)(list)
}
func ZendAstAddArrayElement(result *Zval, offset *Zval, expr *Zval) int {
	switch offset.GetType() {
	case 0:
		if ZendHashNextIndexInsert(result.GetValue().GetArr(), expr) == nil {
			ZendError(1<<1, "Cannot add element to the array as the next element is already occupied")
			ZvalPtrDtorNogc(expr)
		}
		break
	case 6:
		ZendSymtableUpdate(result.GetValue().GetArr(), offset.GetValue().GetStr(), expr)
		ZvalPtrDtorStr(offset)
		break
	case 1:
		ZendSymtableUpdate(result.GetValue().GetArr(), ZendEmptyString, expr)
		break
	case 4:
		ZendHashIndexUpdate(result.GetValue().GetArr(), offset.GetValue().GetLval(), expr)
		break
	case 2:
		ZendHashIndexUpdate(result.GetValue().GetArr(), 0, expr)
		break
	case 3:
		ZendHashIndexUpdate(result.GetValue().GetArr(), 1, expr)
		break
	case 5:
		ZendHashIndexUpdate(result.GetValue().GetArr(), ZendDvalToLval(offset.GetValue().GetDval()), expr)
		break
	case 9:
		ZendError(1<<3, "Resource ID#%d used as offset, casting to integer (%d)", offset.GetValue().GetRes().GetHandle(), offset.GetValue().GetRes().GetHandle())
		ZendHashIndexUpdate(result.GetValue().GetArr(), offset.GetValue().GetRes().GetHandle(), expr)
		break
	default:
		ZendThrowError(nil, "Illegal offset type")
		return FAILURE
	}
	return SUCCESS
}
func ZendAstAddUnpackedElement(result *Zval, expr *Zval) int {
	if expr.GetType() == 7 {
		var ht *HashTable = expr.GetValue().GetArr()
		var val *Zval
		var key *ZendString
		for {
			var __ht *HashTable = ht
			var _p *Bucket = __ht.GetArData()
			var _end *Bucket = _p + __ht.GetNNumUsed()
			for ; _p != _end; _p++ {
				var _z *Zval = &_p.val

				if _z.GetType() == 0 {
					continue
				}
				key = _p.GetKey()
				val = _z
				if key != nil {
					ZendThrowError(nil, "Cannot unpack array with string keys")
					return FAILURE
				} else {
					if ZendHashNextIndexInsert(result.GetValue().GetArr(), val) == nil {
						ZendError(1<<1, "Cannot add element to the array as the next element is already occupied")
						break
					}
					if val.GetTypeFlags() != 0 {
						ZvalAddrefP(val)
					}
				}
			}
			break
		}
		return SUCCESS
	}

	/* Objects or references cannot occur in a constant expression. */

	ZendThrowError(nil, "Only arrays and Traversables can be unpacked")
	return FAILURE
}
func ZendAstEvaluate(result *Zval, ast *ZendAst, scope *ZendClassEntry) int {
	var op1 Zval
	var op2 Zval
	var ret int = SUCCESS
	switch ast.GetKind() {
	case ZEND_AST_BINARY_OP:
		if ZendAstEvaluate(&op1, ast.GetChild()[0], scope) != SUCCESS {
			ret = FAILURE
		} else if ZendAstEvaluate(&op2, ast.GetChild()[1], scope) != SUCCESS {
			ZvalPtrDtorNogc(&op1)
			ret = FAILURE
		} else {
			var op BinaryOpType = GetBinaryOp(ast.GetAttr())
			ret = op(result, &op1, &op2)
			ZvalPtrDtorNogc(&op1)
			ZvalPtrDtorNogc(&op2)
		}
		break
	case ZEND_AST_GREATER:

	case ZEND_AST_GREATER_EQUAL:
		if ZendAstEvaluate(&op1, ast.GetChild()[0], scope) != SUCCESS {
			ret = FAILURE
		} else if ZendAstEvaluate(&op2, ast.GetChild()[1], scope) != SUCCESS {
			ZvalPtrDtorNogc(&op1)
			ret = FAILURE
		} else {

			/* op1 > op2 is the same as op2 < op1 */

			var op BinaryOpType = g.Cond(ast.GetKind() == ZEND_AST_GREATER, IsSmallerFunction, IsSmallerOrEqualFunction)
			ret = op(result, &op2, &op1)
			ZvalPtrDtorNogc(&op1)
			ZvalPtrDtorNogc(&op2)
		}
		break
	case ZEND_AST_UNARY_OP:
		if ZendAstEvaluate(&op1, ast.GetChild()[0], scope) != SUCCESS {
			ret = FAILURE
		} else {
			var op UnaryOpType = GetUnaryOp(ast.GetAttr())
			ret = op(result, &op1)
			ZvalPtrDtorNogc(&op1)
		}
		break
	case ZEND_AST_ZVAL:
		var zv *Zval = ZendAstGetZval(ast)
		var _z1 *Zval = result
		var _z2 *Zval = zv
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if (_t & 0xff00) != 0 {
			ZendGcAddref(&_gc.gc)
		}
		break
	case ZEND_AST_CONSTANT:
		var name *ZendString = ZendAstGetConstantName(ast)
		var zv *Zval = ZendGetConstantEx(name, scope, ast.GetAttr())
		if zv == nil {
			result.SetTypeInfo(0)
			ret = ZendUseUndefinedConstant(name, ast.GetAttr(), result)
			break
		}
		var _z1 *Zval = result
		var _z2 *Zval = zv
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if (_t & 0xff00) != 0 {
			if (ZvalGcFlags(_gc.GetGc().GetTypeInfo()) & 1 << 7) == 0 {
				ZendGcAddref(&_gc.gc)
			} else {
				ZvalCopyCtorFunc(_z1)
			}
		}
		break
	case ZEND_AST_CONSTANT_CLASS:
		if scope != nil {
			var __z *Zval = result
			var __s *ZendString = scope.GetName()
			__z.GetValue().SetStr(__s)
			if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
				__z.SetTypeInfo(6)
			} else {
				ZendGcAddref(&__s.gc)
				__z.SetTypeInfo(6 | 1<<0<<8)
			}
		} else {
			var __z *Zval = result
			var __s *ZendString = ZendEmptyString
			__z.GetValue().SetStr(__s)
			__z.SetTypeInfo(6)
		}
		break
	case ZEND_AST_CLASS_NAME:
		if scope == nil {
			ZendThrowError(nil, "Cannot use \"self\" when no class scope is active")
			return FAILURE
		}
		if ast.GetAttr() == 1 {
			var __z *Zval = result
			var __s *ZendString = scope.GetName()
			__z.GetValue().SetStr(__s)
			if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
				__z.SetTypeInfo(6)
			} else {
				ZendGcAddref(&__s.gc)
				__z.SetTypeInfo(6 | 1<<0<<8)
			}
		} else if ast.GetAttr() == 2 {
			if !(scope.parent) {
				ZendThrowError(nil, "Cannot use \"parent\" when current class scope has no parent")
				return FAILURE
			}
			var __z *Zval = result
			var __s *ZendString = scope.parent.name
			__z.GetValue().SetStr(__s)
			if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
				__z.SetTypeInfo(6)
			} else {
				ZendGcAddref(&__s.gc)
				__z.SetTypeInfo(6 | 1<<0<<8)
			}
		} else {
			assert(false)
		}
		break
	case ZEND_AST_AND:
		if ZendAstEvaluate(&op1, ast.GetChild()[0], scope) != SUCCESS {
			ret = FAILURE
			break
		}
		if ZendIsTrue(&op1) != 0 {
			if ZendAstEvaluate(&op2, ast.GetChild()[1], scope) != SUCCESS {
				ZvalPtrDtorNogc(&op1)
				ret = FAILURE
				break
			}
			if ZendIsTrue(&op2) != 0 {
				result.SetTypeInfo(3)
			} else {
				result.SetTypeInfo(2)
			}
			ZvalPtrDtorNogc(&op2)
		} else {
			result.SetTypeInfo(2)
		}
		ZvalPtrDtorNogc(&op1)
		break
	case ZEND_AST_OR:
		if ZendAstEvaluate(&op1, ast.GetChild()[0], scope) != SUCCESS {
			ret = FAILURE
			break
		}
		if ZendIsTrue(&op1) != 0 {
			result.SetTypeInfo(3)
		} else {
			if ZendAstEvaluate(&op2, ast.GetChild()[1], scope) != SUCCESS {
				ZvalPtrDtorNogc(&op1)
				ret = FAILURE
				break
			}
			if ZendIsTrue(&op2) != 0 {
				result.SetTypeInfo(3)
			} else {
				result.SetTypeInfo(2)
			}
			ZvalPtrDtorNogc(&op2)
		}
		ZvalPtrDtorNogc(&op1)
		break
	case ZEND_AST_CONDITIONAL:
		if ZendAstEvaluate(&op1, ast.GetChild()[0], scope) != SUCCESS {
			ret = FAILURE
			break
		}
		if ZendIsTrue(&op1) != 0 {
			if ast.GetChild()[1] == nil {
				*result = op1
			} else {
				if ZendAstEvaluate(result, ast.GetChild()[1], scope) != SUCCESS {
					ZvalPtrDtorNogc(&op1)
					ret = FAILURE
					break
				}
				ZvalPtrDtorNogc(&op1)
			}
		} else {
			if ZendAstEvaluate(result, ast.GetChild()[2], scope) != SUCCESS {
				ZvalPtrDtorNogc(&op1)
				ret = FAILURE
				break
			}
			ZvalPtrDtorNogc(&op1)
		}
		break
	case ZEND_AST_COALESCE:
		if ZendAstEvaluate(&op1, ast.GetChild()[0], scope) != SUCCESS {
			ret = FAILURE
			break
		}
		if op1.GetType() > 1 {
			*result = op1
		} else {
			if ZendAstEvaluate(result, ast.GetChild()[1], scope) != SUCCESS {
				ZvalPtrDtorNogc(&op1)
				ret = FAILURE
				break
			}
			ZvalPtrDtorNogc(&op1)
		}
		break
	case ZEND_AST_UNARY_PLUS:
		if ZendAstEvaluate(&op2, ast.GetChild()[0], scope) != SUCCESS {
			ret = FAILURE
		} else {
			var __z *Zval = &op1
			__z.GetValue().SetLval(0)
			__z.SetTypeInfo(4)
			ret = AddFunction(result, &op1, &op2)
			ZvalPtrDtorNogc(&op2)
		}
		break
	case ZEND_AST_UNARY_MINUS:
		if ZendAstEvaluate(&op2, ast.GetChild()[0], scope) != SUCCESS {
			ret = FAILURE
		} else {
			var __z *Zval = &op1
			__z.GetValue().SetLval(0)
			__z.SetTypeInfo(4)
			ret = SubFunction(result, &op1, &op2)
			ZvalPtrDtorNogc(&op2)
		}
		break
	case ZEND_AST_ARRAY:
		var i uint32
		var list *ZendAstList = ZendAstGetList(ast)
		if list.GetChildren() == 0 {
			var __z *Zval = result
			__z.GetValue().SetArr((*ZendArray)(&ZendEmptyArray))
			__z.SetTypeInfo(7)
			break
		}
		var __arr *ZendArray = _zendNewArray(0)
		var __z *Zval = result
		__z.GetValue().SetArr(__arr)
		__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
		for i = 0; i < list.GetChildren(); i++ {
			var elem *ZendAst = list.GetChild()[i]
			if elem.GetKind() == ZEND_AST_UNPACK {
				if ZendAstEvaluate(&op1, elem.GetChild()[0], scope) != SUCCESS {
					ZvalPtrDtorNogc(result)
					return FAILURE
				}
				if ZendAstAddUnpackedElement(result, &op1) != SUCCESS {
					ZvalPtrDtorNogc(&op1)
					ZvalPtrDtorNogc(result)
					return FAILURE
				}
				ZvalPtrDtorNogc(&op1)
				continue
			}
			if elem.GetChild()[1] != nil {
				if ZendAstEvaluate(&op1, elem.GetChild()[1], scope) != SUCCESS {
					ZvalPtrDtorNogc(result)
					return FAILURE
				}
			} else {
				&op1.SetTypeInfo(0)
			}
			if ZendAstEvaluate(&op2, elem.GetChild()[0], scope) != SUCCESS {
				ZvalPtrDtorNogc(&op1)
				ZvalPtrDtorNogc(result)
				return FAILURE
			}
			if ZendAstAddArrayElement(result, &op1, &op2) != SUCCESS {
				ZvalPtrDtorNogc(&op1)
				ZvalPtrDtorNogc(&op2)
				ZvalPtrDtorNogc(result)
				return FAILURE
			}
		}
		break
	case ZEND_AST_DIM:
		if ast.GetChild()[1] == nil {
			ZendErrorNoreturn(1<<6, "Cannot use [] for reading")
		}
		if ZendAstEvaluate(&op1, ast.GetChild()[0], scope) != SUCCESS {
			ret = FAILURE
		} else if ZendAstEvaluate(&op2, ast.GetChild()[1], scope) != SUCCESS {
			ZvalPtrDtorNogc(&op1)
			ret = FAILURE
		} else {
			ZendFetchDimensionConst(result, &op1, &op2, g.Cond((ast.GetAttr()&1<<0) != 0, 3, 0))
			ZvalPtrDtorNogc(&op1)
			ZvalPtrDtorNogc(&op2)
		}
		break
	default:
		ZendThrowError(nil, "Unsupported constant expression")
		ret = FAILURE
	}
	return ret
}
func ZendAstTreeSize(ast *ZendAst) int {
	var size int
	if ast.GetKind() == ZEND_AST_ZVAL || ast.GetKind() == ZEND_AST_CONSTANT {
		size = g.SizeOf("zend_ast_zval")
	} else if ZendAstIsList(ast) != 0 {
		var i uint32
		var list *ZendAstList = ZendAstGetList(ast)
		size = ZendAstListSize(list.GetChildren())
		for i = 0; i < list.GetChildren(); i++ {
			if list.GetChild()[i] != nil {
				size += ZendAstTreeSize(list.GetChild()[i])
			}
		}
	} else {
		var i uint32
		var children uint32 = ZendAstGetNumChildren(ast)
		size = ZendAstSize(children)
		for i = 0; i < children; i++ {
			if ast.GetChild()[i] != nil {
				size += ZendAstTreeSize(ast.GetChild()[i])
			}
		}
	}
	return size
}
func ZendAstTreeCopy(ast *ZendAst, buf any) any {
	if ast.GetKind() == ZEND_AST_ZVAL {
		var new_ *ZendAstZval = (*ZendAstZval)(buf)
		new_.SetKind(ZEND_AST_ZVAL)
		new_.SetAttr(ast.GetAttr())
		var _z1 *Zval = &new_.val
		var _z2 *Zval = ZendAstGetZval(ast)
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if (_t & 0xff00) != 0 {
			ZendGcAddref(&_gc.gc)
		}
		buf = any((*byte)(buf + g.SizeOf("zend_ast_zval")))
	} else if ast.GetKind() == ZEND_AST_CONSTANT {
		var new_ *ZendAstZval = (*ZendAstZval)(buf)
		new_.SetKind(ZEND_AST_CONSTANT)
		new_.SetAttr(ast.GetAttr())
		var __z *Zval = &new_.val
		var __s *ZendString = ZendAstGetConstantName(ast)
		__z.GetValue().SetStr(__s)
		if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
			__z.SetTypeInfo(6)
		} else {
			ZendGcAddref(&__s.gc)
			__z.SetTypeInfo(6 | 1<<0<<8)
		}
		buf = any((*byte)(buf + g.SizeOf("zend_ast_zval")))
	} else if ZendAstIsList(ast) != 0 {
		var list *ZendAstList = ZendAstGetList(ast)
		var new_ *ZendAstList = (*ZendAstList)(buf)
		var i uint32
		new_.SetKind(list.GetKind())
		new_.SetAttr(list.GetAttr())
		new_.SetChildren(list.GetChildren())
		buf = any((*byte)(buf + ZendAstListSize(list.GetChildren())))
		for i = 0; i < list.GetChildren(); i++ {
			if list.GetChild()[i] != nil {
				new_.GetChild()[i] = (*ZendAst)(buf)
				buf = ZendAstTreeCopy(list.GetChild()[i], buf)
			} else {
				new_.GetChild()[i] = nil
			}
		}
	} else {
		var i uint32
		var children uint32 = ZendAstGetNumChildren(ast)
		var new_ *ZendAst = (*ZendAst)(buf)
		new_.SetKind(ast.GetKind())
		new_.SetAttr(ast.GetAttr())
		buf = any((*byte)(buf + ZendAstSize(children)))
		for i = 0; i < children; i++ {
			if ast.GetChild()[i] != nil {
				new_.GetChild()[i] = (*ZendAst)(buf)
				buf = ZendAstTreeCopy(ast.GetChild()[i], buf)
			} else {
				new_.GetChild()[i] = nil
			}
		}
	}
	return buf
}
func ZendAstCopy(ast *ZendAst) *ZendAstRef {
	var tree_size int
	var ref *ZendAstRef
	assert(ast != nil)
	tree_size = ZendAstTreeSize(ast) + g.SizeOf("zend_ast_ref")
	ref = _emalloc(tree_size)
	ZendAstTreeCopy(ast, (*ZendAst)((*byte)(ref)+g.SizeOf("zend_ast_ref")))
	ZendGcSetRefcount(&ref.gc, 1)
	ref.GetGc().SetTypeInfo(11)
	return ref
}
func ZendAstDestroy(ast *ZendAst) {
tail_call:
	if ast == nil {
		return
	}
	if ast.GetKind() >= ZEND_AST_VAR {
		var i uint32
		var children uint32 = ZendAstGetNumChildren(ast)
		for i = 1; i < children; i++ {
			ZendAstDestroy(ast.GetChild()[i])
		}
		ast = ast.GetChild()[0]
		goto tail_call
	} else if ast.GetKind() == ZEND_AST_ZVAL {
		ZvalPtrDtorNogc(ZendAstGetZval(ast))
	} else if ZendAstIsList(ast) != 0 {
		var list *ZendAstList = ZendAstGetList(ast)
		if list.GetChildren() != 0 {
			var i uint32
			for i = 1; i < list.GetChildren(); i++ {
				ZendAstDestroy(list.GetChild()[i])
			}
			ast = list.GetChild()[0]
			goto tail_call
		}
	} else if ast.GetKind() == ZEND_AST_CONSTANT {
		ZendStringReleaseEx(ZendAstGetConstantName(ast), 0)
	} else if ast.GetKind() >= ZEND_AST_FUNC_DECL {
		var decl *ZendAstDecl = (*ZendAstDecl)(ast)
		if decl.GetName() != nil {
			ZendStringReleaseEx(decl.GetName(), 0)
		}
		if decl.GetDocComment() != nil {
			ZendStringReleaseEx(decl.GetDocComment(), 0)
		}
		ZendAstDestroy(decl.GetChild()[0])
		ZendAstDestroy(decl.GetChild()[1])
		ZendAstDestroy(decl.GetChild()[2])
		ast = decl.GetChild()[3]
		goto tail_call
	}
}
func ZendAstRefDestroy(ast *ZendAstRef) {
	ZendAstDestroy((*ZendAst)((*byte)(ast) + g.SizeOf("zend_ast_ref")))
	_efree(ast)
}
func ZendAstApply(ast *ZendAst, fn ZendAstApplyFunc) {
	if ZendAstIsList(ast) != 0 {
		var list *ZendAstList = ZendAstGetList(ast)
		var i uint32
		for i = 0; i < list.GetChildren(); i++ {
			fn(&list.child[i])
		}
	} else {
		var i uint32
		var children uint32 = ZendAstGetNumChildren(ast)
		for i = 0; i < children; i++ {
			fn(&ast.child[i])
		}
	}
}

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

func ZendAstExportStr(str *SmartStr, s *ZendString) {
	var i int
	for i = 0; i < s.GetLen(); i++ {
		var c uint8 = s.GetVal()[i]
		if c == '\'' || c == '\\' {
			SmartStrAppendcEx(str, '\\', 0)
			SmartStrAppendcEx(str, c, 0)
		} else {
			SmartStrAppendcEx(str, c, 0)
		}
	}
}
func ZendAstExportQstr(str *SmartStr, quote byte, s *ZendString) {
	var i int
	for i = 0; i < s.GetLen(); i++ {
		var c uint8 = s.GetVal()[i]
		if c < ' ' {
			switch c {
			case '\n':
				SmartStrAppendlEx(str, "\\n", strlen("\\n"), 0)
				break
			case '\r':
				SmartStrAppendlEx(str, "\\r", strlen("\\r"), 0)
				break
			case '\t':
				SmartStrAppendlEx(str, "\\t", strlen("\\t"), 0)
				break
			case 'f':
				SmartStrAppendlEx(str, "\\f", strlen("\\f"), 0)
				break
			case 'v':
				SmartStrAppendlEx(str, "\\v", strlen("\\v"), 0)
				break
			case 'e':
				SmartStrAppendlEx(str, "\\e", strlen("\\e"), 0)
				break
			default:
				SmartStrAppendlEx(str, "\\0", strlen("\\0"), 0)
				SmartStrAppendcEx(str, '0'+c/8, 0)
				SmartStrAppendcEx(str, '0'+c%8, 0)
				break
			}
		} else {
			if c == quote || c == '$' || c == '\\' {
				SmartStrAppendcEx(str, '\\', 0)
			}
			SmartStrAppendcEx(str, c, 0)
		}
	}
}
func ZendAstExportIndent(str *SmartStr, indent int) {
	for indent > 0 {
		SmartStrAppendlEx(str, "    ", strlen("    "), 0)
		indent--
	}
}
func ZendAstExportName(str *SmartStr, ast *ZendAst, priority int, indent int) {
	if ast.GetKind() == ZEND_AST_ZVAL {
		var zv *Zval = ZendAstGetZval(ast)
		if zv.GetType() == 6 {
			SmartStrAppendEx(str, zv.GetValue().GetStr(), 0)
			return
		}
	}
	ZendAstExportEx(str, ast, priority, indent)
}
func ZendAstExportNsName(str *SmartStr, ast *ZendAst, priority int, indent int) {
	if ast.GetKind() == ZEND_AST_ZVAL {
		var zv *Zval = ZendAstGetZval(ast)
		if zv.GetType() == 6 {
			if ast.GetAttr() == 0 {
				SmartStrAppendcEx(str, '\\', 0)
			} else if ast.GetAttr() == 2 {
				SmartStrAppendlEx(str, "namespace\\", strlen("namespace\\"), 0)
			}
			SmartStrAppendEx(str, zv.GetValue().GetStr(), 0)
			return
		}
	}
	ZendAstExportEx(str, ast, priority, indent)
}
func ZendAstValidVarChar(ch byte) int {
	var c uint8 = uint8(ch)
	if c != '_' && c < 127 && (c < '0' || c > '9') && (c < 'A' || c > 'Z') && (c < 'a' || c > 'z') {
		return 0
	}
	return 1
}
func ZendAstValidVarName(s *byte, len_ int) int {
	var c uint8
	var i int
	if len_ == 0 {
		return 0
	}
	c = uint8(s[0])
	if c != '_' && c < 127 && (c < 'A' || c > 'Z') && (c < 'a' || c > 'z') {
		return 0
	}
	for i = 1; i < len_; i++ {
		c = uint8(s[i])
		if c != '_' && c < 127 && (c < '0' || c > '9') && (c < 'A' || c > 'Z') && (c < 'a' || c > 'z') {
			return 0
		}
	}
	return 1
}
func ZendAstVarNeedsBraces(ch byte) int {
	return ch == '[' || ZendAstValidVarChar(ch) != 0
}
func ZendAstExportVar(str *SmartStr, ast *ZendAst, priority int, indent int) {
	if ast.GetKind() == ZEND_AST_ZVAL {
		var zv *Zval = ZendAstGetZval(ast)
		if zv.GetType() == 6 && ZendAstValidVarName(zv.GetValue().GetStr().GetVal(), zv.GetValue().GetStr().GetLen()) != 0 {
			SmartStrAppendEx(str, zv.GetValue().GetStr(), 0)
			return
		}
	} else if ast.GetKind() == ZEND_AST_VAR {
		ZendAstExportEx(str, ast, 0, indent)
		return
	}
	SmartStrAppendcEx(str, '{', 0)
	ZendAstExportName(str, ast, 0, indent)
	SmartStrAppendcEx(str, '}', 0)
}
func ZendAstExportList(str *SmartStr, list *ZendAstList, separator int, priority int, indent int) {
	var i uint32 = 0
	for i < list.GetChildren() {
		if i != 0 && separator != 0 {
			SmartStrAppendlEx(str, ", ", strlen(", "), 0)
		}
		ZendAstExportEx(str, list.GetChild()[i], priority, indent)
		i++
	}
}
func ZendAstExportEncapsList(str *SmartStr, quote byte, list *ZendAstList, indent int) {
	var i uint32 = 0
	var ast *ZendAst
	for i < list.GetChildren() {
		ast = list.GetChild()[i]
		if ast.GetKind() == ZEND_AST_ZVAL {
			var zv *Zval = ZendAstGetZval(ast)
			assert(zv.GetType() == 6)
			ZendAstExportQstr(str, quote, zv.GetValue().GetStr())
		} else if ast.GetKind() == ZEND_AST_VAR && ast.GetChild()[0].GetKind() == ZEND_AST_ZVAL && (i+1 == list.GetChildren() || list.GetChild()[i+1].GetKind() != ZEND_AST_ZVAL || ZendAstVarNeedsBraces((*(ZendAstGetZval(list.GetChild()[i+1]).GetValue().GetStr())).val) == 0) {
			ZendAstExportEx(str, ast, 0, indent)
		} else {
			SmartStrAppendcEx(str, '{', 0)
			ZendAstExportEx(str, ast, 0, indent)
			SmartStrAppendcEx(str, '}', 0)
		}
		i++
	}
}
func ZendAstExportNameListEx(str *SmartStr, list *ZendAstList, indent int, separator string) {
	var i uint32 = 0
	for i < list.GetChildren() {
		if i != 0 {
			SmartStrAppendlEx(str, separator, strlen(separator), 0)
		}
		ZendAstExportName(str, list.GetChild()[i], 0, indent)
		i++
	}
}

// #define zend_ast_export_name_list(s,l,i) zend_ast_export_name_list_ex ( s , l , i , ", " )

// #define zend_ast_export_catch_name_list(s,l,i) zend_ast_export_name_list_ex ( s , l , i , "|" )

func ZendAstExportVarList(str *SmartStr, list *ZendAstList, indent int) {
	var i uint32 = 0
	for i < list.GetChildren() {
		if i != 0 {
			SmartStrAppendlEx(str, ", ", strlen(", "), 0)
		}
		if (list.GetChild()[i].GetAttr() & 1) != 0 {
			SmartStrAppendcEx(str, '&', 0)
		}
		SmartStrAppendcEx(str, '$', 0)
		ZendAstExportName(str, list.GetChild()[i], 20, indent)
		i++
	}
}
func ZendAstExportStmt(str *SmartStr, ast *ZendAst, indent int) {
	if ast == nil {
		return
	}
	if ast.GetKind() == ZEND_AST_STMT_LIST || ast.GetKind() == ZEND_AST_TRAIT_ADAPTATIONS {
		var list *ZendAstList = (*ZendAstList)(ast)
		var i uint32 = 0
		for i < list.GetChildren() {
			ast = list.GetChild()[i]
			ZendAstExportStmt(str, ast, indent)
			i++
		}
	} else {
		ZendAstExportIndent(str, indent)
		ZendAstExportEx(str, ast, 0, indent)
		switch ast.GetKind() {
		case ZEND_AST_LABEL:

		case ZEND_AST_IF:

		case ZEND_AST_SWITCH:

		case ZEND_AST_WHILE:

		case ZEND_AST_TRY:

		case ZEND_AST_FOR:

		case ZEND_AST_FOREACH:

		case ZEND_AST_FUNC_DECL:

		case ZEND_AST_METHOD:

		case ZEND_AST_CLASS:

		case ZEND_AST_USE_TRAIT:

		case ZEND_AST_NAMESPACE:

		case ZEND_AST_DECLARE:
			break
		default:
			SmartStrAppendcEx(str, ';', 0)
			break
		}
		SmartStrAppendcEx(str, '\n', 0)
	}
}
func ZendAstExportIfStmt(str *SmartStr, list *ZendAstList, indent int) {
	var i uint32
	var ast *ZendAst
tail_call:
	i = 0
	for i < list.GetChildren() {
		ast = list.GetChild()[i]
		assert(ast.GetKind() == ZEND_AST_IF_ELEM)
		if ast.GetChild()[0] != nil {
			if i == 0 {
				SmartStrAppendlEx(str, "if (", strlen("if ("), 0)
			} else {
				ZendAstExportIndent(str, indent)
				SmartStrAppendlEx(str, "} elseif (", strlen("} elseif ("), 0)
			}
			ZendAstExportEx(str, ast.GetChild()[0], 0, indent)
			SmartStrAppendlEx(str, ") {\n", strlen(") {\n"), 0)
			ZendAstExportStmt(str, ast.GetChild()[1], indent+1)
		} else {
			ZendAstExportIndent(str, indent)
			SmartStrAppendlEx(str, "} else ", strlen("} else "), 0)
			if ast.GetChild()[1] != nil && ast.GetChild()[1].GetKind() == ZEND_AST_IF {
				list = (*ZendAstList)(ast.GetChild()[1])
				goto tail_call
			} else {
				SmartStrAppendlEx(str, "{\n", strlen("{\n"), 0)
				ZendAstExportStmt(str, ast.GetChild()[1], indent+1)
			}
		}
		i++
	}
	ZendAstExportIndent(str, indent)
	SmartStrAppendcEx(str, '}', 0)
}
func ZendAstExportZval(str *SmartStr, zv *Zval, priority int, indent int) {
	var idx ZendLong
	var key *ZendString
	var val *Zval
	var first int
	if zv.GetType() == 10 {
		zv = &(*zv).value.GetRef().GetVal()
	}
	switch zv.GetType() {
	case 1:
		SmartStrAppendlEx(str, "null", strlen("null"), 0)
		break
	case 2:
		SmartStrAppendlEx(str, "false", strlen("false"), 0)
		break
	case 3:
		SmartStrAppendlEx(str, "true", strlen("true"), 0)
		break
	case 4:
		SmartStrAppendLongEx(str, zv.GetValue().GetLval(), 0)
		break
	case 5:
		key = ZendStrpprintf(0, "%.*G", int(EG.GetPrecision()), zv.GetValue().GetDval())
		SmartStrAppendlEx(str, key.GetVal(), key.GetLen(), 0)
		ZendStringReleaseEx(key, 0)
		break
	case 6:
		SmartStrAppendcEx(str, '\'', 0)
		ZendAstExportStr(str, zv.GetValue().GetStr())
		SmartStrAppendcEx(str, '\'', 0)
		break
	case 7:
		SmartStrAppendcEx(str, '[', 0)
		first = 1
		for {
			var __ht *HashTable = zv.GetValue().GetArr()
			var _p *Bucket = __ht.GetArData()
			var _end *Bucket = _p + __ht.GetNNumUsed()
			for ; _p != _end; _p++ {
				var _z *Zval = &_p.val

				if _z.GetType() == 0 {
					continue
				}
				idx = _p.GetH()
				key = _p.GetKey()
				val = _z
				if first != 0 {
					first = 0
				} else {
					SmartStrAppendlEx(str, ", ", strlen(", "), 0)
				}
				if key != nil {
					SmartStrAppendcEx(str, '\'', 0)
					ZendAstExportStr(str, key)
					SmartStrAppendlEx(str, "' => ", strlen("' => "), 0)
				} else {
					SmartStrAppendLongEx(str, idx, 0)
					SmartStrAppendlEx(str, " => ", strlen(" => "), 0)
				}
				ZendAstExportZval(str, val, 0, indent)
			}
			break
		}
		SmartStrAppendcEx(str, ']', 0)
		break
	case 11:
		ZendAstExportEx(str, (*ZendAst)((*byte)(zv.GetValue().GetAst())+g.SizeOf("zend_ast_ref")), priority, indent)
		break
	default:
		break
	}
}
func ZendAstExportClassNoHeader(str *SmartStr, decl *ZendAstDecl, indent int) {
	if decl.GetChild()[0] != nil {
		SmartStrAppendlEx(str, " extends ", strlen(" extends "), 0)
		ZendAstExportNsName(str, decl.GetChild()[0], 0, indent)
	}
	if decl.GetChild()[1] != nil {
		SmartStrAppendlEx(str, " implements ", strlen(" implements "), 0)
		ZendAstExportEx(str, decl.GetChild()[1], 0, indent)
	}
	SmartStrAppendlEx(str, " {\n", strlen(" {\n"), 0)
	ZendAstExportStmt(str, decl.GetChild()[2], indent+1)
	ZendAstExportIndent(str, indent)
	SmartStrAppendlEx(str, "}", strlen("}"), 0)
}

// #define BINARY_OP(_op,_p,_pl,_pr) do { op = _op ; p = _p ; pl = _pl ; pr = _pr ; goto binary_op ; } while ( 0 )

// #define PREFIX_OP(_op,_p,_pl) do { op = _op ; p = _p ; pl = _pl ; goto prefix_op ; } while ( 0 )

// #define FUNC_OP(_op) do { op = _op ; goto func_op ; } while ( 0 )

// #define POSTFIX_OP(_op,_p,_pl) do { op = _op ; p = _p ; pl = _pl ; goto postfix_op ; } while ( 0 )

// #define APPEND_NODE_1(_op) do { op = _op ; goto append_node_1 ; } while ( 0 )

// #define APPEND_STR(_op) do { op = _op ; goto append_str ; } while ( 0 )

// #define APPEND_DEFAULT_VALUE(n) do { p = n ; goto append_default_value ; } while ( 0 )

func ZendAstExportEx(str *SmartStr, ast *ZendAst, priority int, indent int) {
	var decl *ZendAstDecl
	var p int
	var pl int
	var pr int
	var op *byte
tail_call:
	if ast == nil {
		return
	}
	switch ast.GetKind() {
	case ZEND_AST_ZVAL:
		ZendAstExportZval(str, ZendAstGetZval(ast), priority, indent)
		break
	case ZEND_AST_CONSTANT:
		var name *ZendString = ZendAstGetConstantName(ast)
		SmartStrAppendlEx(str, name.GetVal(), name.GetLen(), 0)
		break
	case ZEND_AST_CONSTANT_CLASS:
		SmartStrAppendlEx(str, "__CLASS__", g.SizeOf("\"__CLASS__\"")-1, 0)
		break
	case ZEND_AST_ZNODE:

		/* This AST kind is only used for temporary nodes during compilation */

		assert(false)
		break
	case ZEND_AST_FUNC_DECL:

	case ZEND_AST_CLOSURE:

	case ZEND_AST_ARROW_FUNC:

	case ZEND_AST_METHOD:
		decl = (*ZendAstDecl)(ast)
		if (decl.GetFlags() & 1 << 0) != 0 {
			SmartStrAppendlEx(str, "public ", strlen("public "), 0)
		} else if (decl.GetFlags() & 1 << 1) != 0 {
			SmartStrAppendlEx(str, "protected ", strlen("protected "), 0)
		} else if (decl.GetFlags() & 1 << 2) != 0 {
			SmartStrAppendlEx(str, "private ", strlen("private "), 0)
		}
		if (decl.GetFlags() & 1 << 4) != 0 {
			SmartStrAppendlEx(str, "static ", strlen("static "), 0)
		}
		if (decl.GetFlags() & 1 << 6) != 0 {
			SmartStrAppendlEx(str, "abstract ", strlen("abstract "), 0)
		}
		if (decl.GetFlags() & 1 << 5) != 0 {
			SmartStrAppendlEx(str, "final ", strlen("final "), 0)
		}
		if decl.GetKind() == ZEND_AST_ARROW_FUNC {
			SmartStrAppendlEx(str, "fn", strlen("fn"), 0)
		} else {
			SmartStrAppendlEx(str, "function ", strlen("function "), 0)
		}
		if (decl.GetFlags() & 1 << 12) != 0 {
			SmartStrAppendcEx(str, '&', 0)
		}
		if ast.GetKind() != ZEND_AST_CLOSURE && ast.GetKind() != ZEND_AST_ARROW_FUNC {
			SmartStrAppendlEx(str, decl.GetName().GetVal(), decl.GetName().GetLen(), 0)
		}
		SmartStrAppendcEx(str, '(', 0)
		ZendAstExportEx(str, decl.GetChild()[0], 0, indent)
		SmartStrAppendcEx(str, ')', 0)
		ZendAstExportEx(str, decl.GetChild()[1], 0, indent)
		if decl.GetChild()[3] != nil {
			SmartStrAppendlEx(str, ": ", strlen(": "), 0)
			if (decl.GetChild()[3].GetAttr() & 1 << 8) != 0 {
				SmartStrAppendcEx(str, '?', 0)
			}
			ZendAstExportNsName(str, decl.GetChild()[3], 0, indent)
		}
		if decl.GetChild()[2] != nil {
			if decl.GetKind() == ZEND_AST_ARROW_FUNC {
				assert(decl.GetChild()[2].GetKind() == ZEND_AST_RETURN)
				SmartStrAppendlEx(str, " => ", strlen(" => "), 0)
				ZendAstExportEx(str, decl.GetChild()[2].GetChild()[0], 0, indent)
				break
			}
			SmartStrAppendlEx(str, " {\n", strlen(" {\n"), 0)
			ZendAstExportStmt(str, decl.GetChild()[2], indent+1)
			ZendAstExportIndent(str, indent)
			SmartStrAppendcEx(str, '}', 0)
			if ast.GetKind() != ZEND_AST_CLOSURE {
				SmartStrAppendcEx(str, '\n', 0)
			}
		} else {
			SmartStrAppendlEx(str, ";\n", strlen(";\n"), 0)
		}
		break
	case ZEND_AST_CLASS:
		decl = (*ZendAstDecl)(ast)
		if (decl.GetFlags() & 1 << 0) != 0 {
			SmartStrAppendlEx(str, "interface ", strlen("interface "), 0)
		} else if (decl.GetFlags() & 1 << 1) != 0 {
			SmartStrAppendlEx(str, "trait ", strlen("trait "), 0)
		} else {
			if (decl.GetFlags() & 1 << 6) != 0 {
				SmartStrAppendlEx(str, "abstract ", strlen("abstract "), 0)
			}
			if (decl.GetFlags() & 1 << 5) != 0 {
				SmartStrAppendlEx(str, "final ", strlen("final "), 0)
			}
			SmartStrAppendlEx(str, "class ", strlen("class "), 0)
		}
		SmartStrAppendlEx(str, decl.GetName().GetVal(), decl.GetName().GetLen(), 0)
		ZendAstExportClassNoHeader(str, decl, indent)
		SmartStrAppendcEx(str, '\n', 0)
		break
	case ZEND_AST_ARG_LIST:

	case ZEND_AST_EXPR_LIST:

	case ZEND_AST_PARAM_LIST:
	simple_list:
		ZendAstExportList(str, (*ZendAstList)(ast), 1, 20, indent)
		break
	case ZEND_AST_ARRAY:
		SmartStrAppendcEx(str, '[', 0)
		ZendAstExportList(str, (*ZendAstList)(ast), 1, 20, indent)
		SmartStrAppendcEx(str, ']', 0)
		break
	case ZEND_AST_ENCAPS_LIST:
		SmartStrAppendcEx(str, '"', 0)
		ZendAstExportEncapsList(str, '"', (*ZendAstList)(ast), indent)
		SmartStrAppendcEx(str, '"', 0)
		break
	case ZEND_AST_STMT_LIST:

	case ZEND_AST_TRAIT_ADAPTATIONS:
		ZendAstExportStmt(str, ast, indent)
		break
	case ZEND_AST_IF:
		ZendAstExportIfStmt(str, (*ZendAstList)(ast), indent)
		break
	case ZEND_AST_SWITCH_LIST:

	case ZEND_AST_CATCH_LIST:
		ZendAstExportList(str, (*ZendAstList)(ast), 0, 0, indent)
		break
	case ZEND_AST_CLOSURE_USES:
		SmartStrAppendlEx(str, " use(", strlen(" use("), 0)
		ZendAstExportVarList(str, (*ZendAstList)(ast), indent)
		SmartStrAppendcEx(str, ')', 0)
		break
	case ZEND_AST_PROP_GROUP:
		var type_ast *ZendAst = ast.GetChild()[0]
		var prop_ast *ZendAst = ast.GetChild()[1]
		if (ast.GetAttr() & 1 << 0) != 0 {
			SmartStrAppendlEx(str, "public ", strlen("public "), 0)
		} else if (ast.GetAttr() & 1 << 1) != 0 {
			SmartStrAppendlEx(str, "protected ", strlen("protected "), 0)
		} else if (ast.GetAttr() & 1 << 2) != 0 {
			SmartStrAppendlEx(str, "private ", strlen("private "), 0)
		}
		if (ast.GetAttr() & 1 << 4) != 0 {
			SmartStrAppendlEx(str, "static ", strlen("static "), 0)
		}
		if type_ast != nil {
			if (type_ast.GetAttr() & 1 << 8) != 0 {
				SmartStrAppendcEx(str, '?', 0)
			}
			ZendAstExportNsName(str, type_ast, 0, indent)
			SmartStrAppendcEx(str, ' ', 0)
		}
		ast = prop_ast
		goto simple_list
	case ZEND_AST_CONST_DECL:

	case ZEND_AST_CLASS_CONST_DECL:
		SmartStrAppendlEx(str, "const ", strlen("const "), 0)
		goto simple_list
	case ZEND_AST_NAME_LIST:
		ZendAstExportNameListEx(str, (*ZendAstList)(ast), indent, ", ")
		break
	case ZEND_AST_USE:
		SmartStrAppendlEx(str, "use ", strlen("use "), 0)
		if ast.GetAttr() == T_FUNCTION {
			SmartStrAppendlEx(str, "function ", strlen("function "), 0)
		} else if ast.GetAttr() == T_CONST {
			SmartStrAppendlEx(str, "const ", strlen("const "), 0)
		}
		goto simple_list
	case ZEND_AST_MAGIC_CONST:
		switch ast.GetAttr() {
		case T_LINE:
			op = "__LINE__"
			goto append_str
		case T_FILE:
			op = "__FILE__"
			goto append_str
		case T_DIR:
			op = "__DIR__"
			goto append_str
		case T_TRAIT_C:
			op = "__TRAIT__"
			goto append_str
		case T_METHOD_C:
			op = "__METHOD__"
			goto append_str
		case T_FUNC_C:
			op = "__FUNCTION__"
			goto append_str
		case T_NS_C:
			op = "__NAMESPACE__"
			goto append_str
		case T_CLASS_C:
			op = "__CLASS__"
			goto append_str
		default:
			break
		}
		break
	case ZEND_AST_TYPE:
		switch ast.GetAttr() & ^(1 << 8) {
		case 7:
			op = "array"
			goto append_str
		case 17:
			op = "callable"
			goto append_str
		default:
			break
		}
		break
	case ZEND_AST_VAR:
		SmartStrAppendcEx(str, '$', 0)
		ZendAstExportVar(str, ast.GetChild()[0], 0, indent)
		break
	case ZEND_AST_CONST:
		ZendAstExportNsName(str, ast.GetChild()[0], 0, indent)
		break
	case ZEND_AST_UNPACK:
		SmartStrAppendlEx(str, "...", strlen("..."), 0)
		ast = ast.GetChild()[0]
		goto tail_call
	case ZEND_AST_UNARY_PLUS:
		op = "+"
		p = 240
		pl = 241
		goto prefix_op
	case ZEND_AST_UNARY_MINUS:
		op = "-"
		p = 240
		pl = 241
		goto prefix_op
	case ZEND_AST_CAST:
		switch ast.GetAttr() {
		case 1:
			op = "(unset)"
			p = 240
			pl = 241
			goto prefix_op
		case 16:
			op = "(bool)"
			p = 240
			pl = 241
			goto prefix_op
		case 4:
			op = "(int)"
			p = 240
			pl = 241
			goto prefix_op
		case 5:
			op = "(double)"
			p = 240
			pl = 241
			goto prefix_op
		case 6:
			op = "(string)"
			p = 240
			pl = 241
			goto prefix_op
		case 7:
			op = "(array)"
			p = 240
			pl = 241
			goto prefix_op
		case 8:
			op = "(object)"
			p = 240
			pl = 241
			goto prefix_op
		default:
			break
		}
		break
	case ZEND_AST_EMPTY:
		op = "empty"
		goto func_op
	case ZEND_AST_ISSET:
		op = "isset"
		goto func_op
	case ZEND_AST_SILENCE:
		op = "@"
		p = 240
		pl = 241
		goto prefix_op
	case ZEND_AST_SHELL_EXEC:
		SmartStrAppendcEx(str, '`', 0)
		if ast.GetChild()[0].GetKind() == ZEND_AST_ENCAPS_LIST {
			ZendAstExportEncapsList(str, '`', (*ZendAstList)(ast.GetChild()[0]), indent)
		} else {
			var zv *Zval
			assert(ast.GetChild()[0].GetKind() == ZEND_AST_ZVAL)
			zv = ZendAstGetZval(ast.GetChild()[0])
			assert(zv.GetType() == 6)
			ZendAstExportQstr(str, '`', zv.GetValue().GetStr())
		}
		SmartStrAppendcEx(str, '`', 0)
		break
	case ZEND_AST_CLONE:
		op = "clone "
		p = 270
		pl = 271
		goto prefix_op
	case ZEND_AST_EXIT:
		if ast.GetChild()[0] != nil {
			op = "exit"
			goto func_op
		} else {
			op = "exit"
			goto append_str
		}
		break
	case ZEND_AST_PRINT:
		op = "print "
		p = 60
		pl = 61
		goto prefix_op
	case ZEND_AST_INCLUDE_OR_EVAL:
		switch ast.GetAttr() {
		case 1 << 2:
			op = "include_once"
			goto func_op
		case 1 << 1:
			op = "include"
			goto func_op
		case 1 << 4:
			op = "require_once"
			goto func_op
		case 1 << 3:
			op = "require"
			goto func_op
		case 1 << 0:
			op = "eval"
			goto func_op
		default:
			break
		}
		break
	case ZEND_AST_UNARY_OP:
		switch ast.GetAttr() {
		case 13:
			op = "~"
			p = 240
			pl = 241
			goto prefix_op
		case 14:
			op = "!"
			p = 240
			pl = 241
			goto prefix_op
		default:
			break
		}
		break
	case ZEND_AST_PRE_INC:
		op = "++"
		p = 240
		pl = 241
		goto prefix_op
	case ZEND_AST_PRE_DEC:
		op = "--"
		p = 240
		pl = 241
		goto prefix_op
	case ZEND_AST_POST_INC:
		op = "++"
		p = 240
		pl = 241
		goto postfix_op
	case ZEND_AST_POST_DEC:
		op = "--"
		p = 240
		pl = 241
		goto postfix_op
	case ZEND_AST_GLOBAL:
		op = "global"
		goto append_node_1
	case ZEND_AST_UNSET:
		op = "unset"
		goto func_op
	case ZEND_AST_RETURN:
		op = "return"
		goto append_node_1
	case ZEND_AST_LABEL:
		ZendAstExportName(str, ast.GetChild()[0], 0, indent)
		SmartStrAppendcEx(str, ':', 0)
		break
	case ZEND_AST_REF:
		SmartStrAppendcEx(str, '&', 0)
		ast = ast.GetChild()[0]
		goto tail_call
	case ZEND_AST_HALT_COMPILER:
		op = "__HALT_COMPILER()"
		goto append_str
	case ZEND_AST_ECHO:
		op = "echo"
		goto append_node_1
	case ZEND_AST_THROW:
		op = "throw"
		goto append_node_1
	case ZEND_AST_GOTO:
		SmartStrAppendlEx(str, "goto ", strlen("goto "), 0)
		ZendAstExportName(str, ast.GetChild()[0], 0, indent)
		break
	case ZEND_AST_BREAK:
		op = "break"
		goto append_node_1
	case ZEND_AST_CONTINUE:
		op = "continue"
		goto append_node_1
	case ZEND_AST_DIM:
		ZendAstExportEx(str, ast.GetChild()[0], 260, indent)
		SmartStrAppendcEx(str, '[', 0)
		if ast.GetChild()[1] != nil {
			ZendAstExportEx(str, ast.GetChild()[1], 0, indent)
		}
		SmartStrAppendcEx(str, ']', 0)
		break
	case ZEND_AST_PROP:
		ZendAstExportEx(str, ast.GetChild()[0], 0, indent)
		SmartStrAppendlEx(str, "->", strlen("->"), 0)
		ZendAstExportVar(str, ast.GetChild()[1], 0, indent)
		break
	case ZEND_AST_STATIC_PROP:
		ZendAstExportNsName(str, ast.GetChild()[0], 0, indent)
		SmartStrAppendlEx(str, "::$", strlen("::$"), 0)
		ZendAstExportVar(str, ast.GetChild()[1], 0, indent)
		break
	case ZEND_AST_CALL:
		ZendAstExportNsName(str, ast.GetChild()[0], 0, indent)
		SmartStrAppendcEx(str, '(', 0)
		ZendAstExportEx(str, ast.GetChild()[1], 0, indent)
		SmartStrAppendcEx(str, ')', 0)
		break
	case ZEND_AST_CLASS_CONST:
		ZendAstExportNsName(str, ast.GetChild()[0], 0, indent)
		SmartStrAppendlEx(str, "::", strlen("::"), 0)
		ZendAstExportName(str, ast.GetChild()[1], 0, indent)
		break
	case ZEND_AST_CLASS_NAME:
		ZendAstExportNsName(str, ast.GetChild()[0], 0, indent)
		SmartStrAppendlEx(str, "::class", strlen("::class"), 0)
		break
	case ZEND_AST_ASSIGN:
		op = " = "
		p = 90
		pl = 91
		pr = 90
		goto binary_op
	case ZEND_AST_ASSIGN_REF:
		op = " =& "
		p = 90
		pl = 91
		pr = 90
		goto binary_op
	case ZEND_AST_ASSIGN_OP:
		switch ast.GetAttr() {
		case 1:
			op = " += "
			p = 90
			pl = 91
			pr = 90
			goto binary_op
		case 2:
			op = " -= "
			p = 90
			pl = 91
			pr = 90
			goto binary_op
		case 3:
			op = " *= "
			p = 90
			pl = 91
			pr = 90
			goto binary_op
		case 4:
			op = " /= "
			p = 90
			pl = 91
			pr = 90
			goto binary_op
		case 5:
			op = " %= "
			p = 90
			pl = 91
			pr = 90
			goto binary_op
		case 6:
			op = " <<= "
			p = 90
			pl = 91
			pr = 90
			goto binary_op
		case 7:
			op = " >>= "
			p = 90
			pl = 91
			pr = 90
			goto binary_op
		case 8:
			op = " .= "
			p = 90
			pl = 91
			pr = 90
			goto binary_op
		case 9:
			op = " |= "
			p = 90
			pl = 91
			pr = 90
			goto binary_op
		case 10:
			op = " &= "
			p = 90
			pl = 91
			pr = 90
			goto binary_op
		case 11:
			op = " ^= "
			p = 90
			pl = 91
			pr = 90
			goto binary_op
		case 12:
			op = " **= "
			p = 90
			pl = 91
			pr = 90
			goto binary_op
		default:
			break
		}
		break
	case ZEND_AST_ASSIGN_COALESCE:
		op = " ??= "
		p = 90
		pl = 91
		pr = 90
		goto binary_op
	case ZEND_AST_BINARY_OP:
		switch ast.GetAttr() {
		case 1:
			op = " + "
			p = 200
			pl = 200
			pr = 201
			goto binary_op
		case 2:
			op = " - "
			p = 200
			pl = 200
			pr = 201
			goto binary_op
		case 3:
			op = " * "
			p = 210
			pl = 210
			pr = 211
			goto binary_op
		case 4:
			op = " / "
			p = 210
			pl = 210
			pr = 211
			goto binary_op
		case 5:
			op = " % "
			p = 210
			pl = 210
			pr = 211
			goto binary_op
		case 6:
			op = " << "
			p = 190
			pl = 190
			pr = 191
			goto binary_op
		case 7:
			op = " >> "
			p = 190
			pl = 190
			pr = 191
			goto binary_op
		case 252:

		case 8:
			op = " . "
			p = 200
			pl = 200
			pr = 201
			goto binary_op
		case 9:
			op = " | "
			p = 140
			pl = 140
			pr = 141
			goto binary_op
		case 10:
			op = " & "
			p = 160
			pl = 160
			pr = 161
			goto binary_op
		case 11:
			op = " ^ "
			p = 150
			pl = 150
			pr = 151
			goto binary_op
		case 16:
			op = " === "
			p = 170
			pl = 171
			pr = 171
			goto binary_op
		case 17:
			op = " !== "
			p = 170
			pl = 171
			pr = 171
			goto binary_op
		case 18:
			op = " == "
			p = 170
			pl = 171
			pr = 171
			goto binary_op
		case 19:
			op = " != "
			p = 170
			pl = 171
			pr = 171
			goto binary_op
		case 20:
			op = " < "
			p = 180
			pl = 181
			pr = 181
			goto binary_op
		case 21:
			op = " <= "
			p = 180
			pl = 181
			pr = 181
			goto binary_op
		case 12:
			op = " ** "
			p = 250
			pl = 251
			pr = 250
			goto binary_op
		case 15:
			op = " xor "
			p = 40
			pl = 40
			pr = 41
			goto binary_op
		case 170:
			op = " <=> "
			p = 180
			pl = 181
			pr = 181
			goto binary_op
		default:
			break
		}
		break
	case ZEND_AST_GREATER:
		op = " > "
		p = 180
		pl = 181
		pr = 181
		goto binary_op
	case ZEND_AST_GREATER_EQUAL:
		op = " >= "
		p = 180
		pl = 181
		pr = 181
		goto binary_op
	case ZEND_AST_AND:
		op = " && "
		p = 130
		pl = 130
		pr = 131
		goto binary_op
	case ZEND_AST_OR:
		op = " || "
		p = 120
		pl = 120
		pr = 121
		goto binary_op
	case ZEND_AST_ARRAY_ELEM:
		if ast.GetChild()[1] != nil {
			ZendAstExportEx(str, ast.GetChild()[1], 80, indent)
			SmartStrAppendlEx(str, " => ", strlen(" => "), 0)
		}
		if ast.GetAttr() != 0 {
			SmartStrAppendcEx(str, '&', 0)
		}
		ZendAstExportEx(str, ast.GetChild()[0], 80, indent)
		break
	case ZEND_AST_NEW:
		SmartStrAppendlEx(str, "new ", strlen("new "), 0)
		if ast.GetChild()[0].GetKind() == ZEND_AST_CLASS {
			SmartStrAppendlEx(str, "class", strlen("class"), 0)
			if ZendAstGetList(ast.GetChild()[1]).GetChildren() != 0 {
				SmartStrAppendcEx(str, '(', 0)
				ZendAstExportEx(str, ast.GetChild()[1], 0, indent)
				SmartStrAppendcEx(str, ')', 0)
			}
			ZendAstExportClassNoHeader(str, (*ZendAstDecl)(ast.GetChild()[0]), indent)
		} else {
			ZendAstExportNsName(str, ast.GetChild()[0], 0, indent)
			SmartStrAppendcEx(str, '(', 0)
			ZendAstExportEx(str, ast.GetChild()[1], 0, indent)
			SmartStrAppendcEx(str, ')', 0)
		}
		break
	case ZEND_AST_INSTANCEOF:
		ZendAstExportEx(str, ast.GetChild()[0], 0, indent)
		SmartStrAppendlEx(str, " instanceof ", strlen(" instanceof "), 0)
		ZendAstExportNsName(str, ast.GetChild()[1], 0, indent)
		break
	case ZEND_AST_YIELD:
		if priority > 70 {
			SmartStrAppendcEx(str, '(', 0)
		}
		SmartStrAppendlEx(str, "yield ", strlen("yield "), 0)
		if ast.GetChild()[0] != nil {
			if ast.GetChild()[1] != nil {
				ZendAstExportEx(str, ast.GetChild()[1], 70, indent)
				SmartStrAppendlEx(str, " => ", strlen(" => "), 0)
			}
			ZendAstExportEx(str, ast.GetChild()[0], 70, indent)
		}
		if priority > 70 {
			SmartStrAppendcEx(str, ')', 0)
		}
		break
	case ZEND_AST_YIELD_FROM:
		op = "yield from "
		p = 85
		pl = 86
		goto prefix_op
	case ZEND_AST_COALESCE:
		op = " ?? "
		p = 110
		pl = 111
		pr = 110
		goto binary_op
	case ZEND_AST_STATIC:
		SmartStrAppendlEx(str, "static $", strlen("static $"), 0)
		ZendAstExportName(str, ast.GetChild()[0], 0, indent)
		p = 1
		goto append_default_value
	case ZEND_AST_WHILE:
		SmartStrAppendlEx(str, "while (", strlen("while ("), 0)
		ZendAstExportEx(str, ast.GetChild()[0], 0, indent)
		SmartStrAppendlEx(str, ") {\n", strlen(") {\n"), 0)
		ZendAstExportStmt(str, ast.GetChild()[1], indent+1)
		ZendAstExportIndent(str, indent)
		SmartStrAppendcEx(str, '}', 0)
		break
	case ZEND_AST_DO_WHILE:
		SmartStrAppendlEx(str, "do {\n", strlen("do {\n"), 0)
		ZendAstExportStmt(str, ast.GetChild()[0], indent+1)
		ZendAstExportIndent(str, indent)
		SmartStrAppendlEx(str, "} while (", strlen("} while ("), 0)
		ZendAstExportEx(str, ast.GetChild()[1], 0, indent)
		SmartStrAppendcEx(str, ')', 0)
		break
	case ZEND_AST_IF_ELEM:
		if ast.GetChild()[0] != nil {
			SmartStrAppendlEx(str, "if (", strlen("if ("), 0)
			ZendAstExportEx(str, ast.GetChild()[0], 0, indent)
			SmartStrAppendlEx(str, ") {\n", strlen(") {\n"), 0)
			ZendAstExportStmt(str, ast.GetChild()[1], indent+1)
		} else {
			SmartStrAppendlEx(str, "else {\n", strlen("else {\n"), 0)
			ZendAstExportStmt(str, ast.GetChild()[1], indent+1)
		}
		ZendAstExportIndent(str, indent)
		SmartStrAppendcEx(str, '}', 0)
		break
	case ZEND_AST_SWITCH:
		SmartStrAppendlEx(str, "switch (", strlen("switch ("), 0)
		ZendAstExportEx(str, ast.GetChild()[0], 0, indent)
		SmartStrAppendlEx(str, ") {\n", strlen(") {\n"), 0)
		ZendAstExportEx(str, ast.GetChild()[1], 0, indent+1)
		ZendAstExportIndent(str, indent)
		SmartStrAppendcEx(str, '}', 0)
		break
	case ZEND_AST_SWITCH_CASE:
		ZendAstExportIndent(str, indent)
		if ast.GetChild()[0] != nil {
			SmartStrAppendlEx(str, "case ", strlen("case "), 0)
			ZendAstExportEx(str, ast.GetChild()[0], 0, indent)
			SmartStrAppendlEx(str, ":\n", strlen(":\n"), 0)
		} else {
			SmartStrAppendlEx(str, "default:\n", strlen("default:\n"), 0)
		}
		ZendAstExportStmt(str, ast.GetChild()[1], indent+1)
		break
	case ZEND_AST_DECLARE:
		SmartStrAppendlEx(str, "declare(", strlen("declare("), 0)
		assert(ast.GetChild()[0].GetKind() == ZEND_AST_CONST_DECL)
		ZendAstExportList(str, (*ZendAstList)(ast.GetChild()[0]), 1, 0, indent)
		SmartStrAppendcEx(str, ')', 0)
		if ast.GetChild()[1] != nil {
			SmartStrAppendlEx(str, " {\n", strlen(" {\n"), 0)
			ZendAstExportStmt(str, ast.GetChild()[1], indent+1)
			ZendAstExportIndent(str, indent)
			SmartStrAppendcEx(str, '}', 0)
		} else {
			SmartStrAppendcEx(str, ';', 0)
		}
		break
	case ZEND_AST_PROP_ELEM:
		SmartStrAppendcEx(str, '$', 0)
	case ZEND_AST_CONST_ELEM:
		ZendAstExportName(str, ast.GetChild()[0], 0, indent)
		p = 1
		goto append_default_value
	case ZEND_AST_USE_TRAIT:
		SmartStrAppendlEx(str, "use ", strlen("use "), 0)
		ZendAstExportEx(str, ast.GetChild()[0], 0, indent)
		if ast.GetChild()[1] != nil {
			SmartStrAppendlEx(str, " {\n", strlen(" {\n"), 0)
			ZendAstExportEx(str, ast.GetChild()[1], 0, indent+1)
			ZendAstExportIndent(str, indent)
			SmartStrAppendlEx(str, "}", strlen("}"), 0)
		} else {
			SmartStrAppendlEx(str, ";", strlen(";"), 0)
		}
		break
	case ZEND_AST_TRAIT_PRECEDENCE:
		ZendAstExportEx(str, ast.GetChild()[0], 0, indent)
		SmartStrAppendlEx(str, " insteadof ", strlen(" insteadof "), 0)
		ZendAstExportEx(str, ast.GetChild()[1], 0, indent)
		break
	case ZEND_AST_METHOD_REFERENCE:
		if ast.GetChild()[0] != nil {
			ZendAstExportName(str, ast.GetChild()[0], 0, indent)
			SmartStrAppendlEx(str, "::", strlen("::"), 0)
		}
		ZendAstExportName(str, ast.GetChild()[1], 0, indent)
		break
	case ZEND_AST_NAMESPACE:
		SmartStrAppendlEx(str, "namespace", strlen("namespace"), 0)
		if ast.GetChild()[0] != nil {
			SmartStrAppendcEx(str, ' ', 0)
			ZendAstExportName(str, ast.GetChild()[0], 0, indent)
		}
		if ast.GetChild()[1] != nil {
			SmartStrAppendlEx(str, " {\n", strlen(" {\n"), 0)
			ZendAstExportStmt(str, ast.GetChild()[1], indent+1)
			ZendAstExportIndent(str, indent)
			SmartStrAppendlEx(str, "}\n", strlen("}\n"), 0)
		} else {
			SmartStrAppendcEx(str, ';', 0)
		}
		break
	case ZEND_AST_USE_ELEM:

	case ZEND_AST_TRAIT_ALIAS:
		ZendAstExportName(str, ast.GetChild()[0], 0, indent)
		if (ast.GetAttr() & 1 << 0) != 0 {
			SmartStrAppendlEx(str, " as public", strlen(" as public"), 0)
		} else if (ast.GetAttr() & 1 << 1) != 0 {
			SmartStrAppendlEx(str, " as protected", strlen(" as protected"), 0)
		} else if (ast.GetAttr() & 1 << 2) != 0 {
			SmartStrAppendlEx(str, " as private", strlen(" as private"), 0)
		} else if ast.GetChild()[1] != nil {
			SmartStrAppendlEx(str, " as", strlen(" as"), 0)
		}
		if ast.GetChild()[1] != nil {
			SmartStrAppendcEx(str, ' ', 0)
			ZendAstExportName(str, ast.GetChild()[1], 0, indent)
		}
		break
	case ZEND_AST_METHOD_CALL:
		ZendAstExportEx(str, ast.GetChild()[0], 0, indent)
		SmartStrAppendlEx(str, "->", strlen("->"), 0)
		ZendAstExportVar(str, ast.GetChild()[1], 0, indent)
		SmartStrAppendcEx(str, '(', 0)
		ZendAstExportEx(str, ast.GetChild()[2], 0, indent)
		SmartStrAppendcEx(str, ')', 0)
		break
	case ZEND_AST_STATIC_CALL:
		ZendAstExportNsName(str, ast.GetChild()[0], 0, indent)
		SmartStrAppendlEx(str, "::", strlen("::"), 0)
		ZendAstExportVar(str, ast.GetChild()[1], 0, indent)
		SmartStrAppendcEx(str, '(', 0)
		ZendAstExportEx(str, ast.GetChild()[2], 0, indent)
		SmartStrAppendcEx(str, ')', 0)
		break
	case ZEND_AST_CONDITIONAL:
		if priority > 100 {
			SmartStrAppendcEx(str, '(', 0)
		}
		ZendAstExportEx(str, ast.GetChild()[0], 100, indent)
		if ast.GetChild()[1] != nil {
			SmartStrAppendlEx(str, " ? ", strlen(" ? "), 0)
			ZendAstExportEx(str, ast.GetChild()[1], 101, indent)
			SmartStrAppendlEx(str, " : ", strlen(" : "), 0)
		} else {
			SmartStrAppendlEx(str, " ?: ", strlen(" ?: "), 0)
		}
		ZendAstExportEx(str, ast.GetChild()[2], 101, indent)
		if priority > 100 {
			SmartStrAppendcEx(str, ')', 0)
		}
		break
	case ZEND_AST_TRY:
		SmartStrAppendlEx(str, "try {\n", strlen("try {\n"), 0)
		ZendAstExportStmt(str, ast.GetChild()[0], indent+1)
		ZendAstExportIndent(str, indent)
		ZendAstExportEx(str, ast.GetChild()[1], 0, indent)
		if ast.GetChild()[2] != nil {
			SmartStrAppendlEx(str, "} finally {\n", strlen("} finally {\n"), 0)
			ZendAstExportStmt(str, ast.GetChild()[2], indent+1)
			ZendAstExportIndent(str, indent)
		}
		SmartStrAppendcEx(str, '}', 0)
		break
	case ZEND_AST_CATCH:
		SmartStrAppendlEx(str, "} catch (", strlen("} catch ("), 0)
		ZendAstExportNameListEx(str, ZendAstGetList(ast.GetChild()[0]), indent, "|")
		SmartStrAppendlEx(str, " $", strlen(" $"), 0)
		ZendAstExportVar(str, ast.GetChild()[1], 0, indent)
		SmartStrAppendlEx(str, ") {\n", strlen(") {\n"), 0)
		ZendAstExportStmt(str, ast.GetChild()[2], indent+1)
		ZendAstExportIndent(str, indent)
		break
	case ZEND_AST_PARAM:
		if ast.GetChild()[0] != nil {
			if (ast.GetChild()[0].GetAttr() & 1 << 8) != 0 {
				SmartStrAppendcEx(str, '?', 0)
			}
			ZendAstExportNsName(str, ast.GetChild()[0], 0, indent)
			SmartStrAppendcEx(str, ' ', 0)
		}
		if (ast.GetAttr() & 1 << 0) != 0 {
			SmartStrAppendcEx(str, '&', 0)
		}
		if (ast.GetAttr() & 1 << 1) != 0 {
			SmartStrAppendlEx(str, "...", strlen("..."), 0)
		}
		SmartStrAppendcEx(str, '$', 0)
		ZendAstExportName(str, ast.GetChild()[1], 0, indent)
		p = 2
		goto append_default_value
	case ZEND_AST_FOR:
		SmartStrAppendlEx(str, "for (", strlen("for ("), 0)
		ZendAstExportEx(str, ast.GetChild()[0], 0, indent)
		SmartStrAppendcEx(str, ';', 0)
		if ast.GetChild()[1] != nil {
			SmartStrAppendcEx(str, ' ', 0)
			ZendAstExportEx(str, ast.GetChild()[1], 0, indent)
		}
		SmartStrAppendcEx(str, ';', 0)
		if ast.GetChild()[2] != nil {
			SmartStrAppendcEx(str, ' ', 0)
			ZendAstExportEx(str, ast.GetChild()[2], 0, indent)
		}
		SmartStrAppendlEx(str, ") {\n", strlen(") {\n"), 0)
		ZendAstExportStmt(str, ast.GetChild()[3], indent+1)
		ZendAstExportIndent(str, indent)
		SmartStrAppendcEx(str, '}', 0)
		break
	case ZEND_AST_FOREACH:
		SmartStrAppendlEx(str, "foreach (", strlen("foreach ("), 0)
		ZendAstExportEx(str, ast.GetChild()[0], 0, indent)
		SmartStrAppendlEx(str, " as ", strlen(" as "), 0)
		if ast.GetChild()[2] != nil {
			ZendAstExportEx(str, ast.GetChild()[2], 0, indent)
			SmartStrAppendlEx(str, " => ", strlen(" => "), 0)
		}
		ZendAstExportEx(str, ast.GetChild()[1], 0, indent)
		SmartStrAppendlEx(str, ") {\n", strlen(") {\n"), 0)
		ZendAstExportStmt(str, ast.GetChild()[3], indent+1)
		ZendAstExportIndent(str, indent)
		SmartStrAppendcEx(str, '}', 0)
		break
	default:
		break
	}
	return
binary_op:
	if priority > p {
		SmartStrAppendcEx(str, '(', 0)
	}
	ZendAstExportEx(str, ast.GetChild()[0], pl, indent)
	SmartStrAppendlEx(str, op, strlen(op), 0)
	ZendAstExportEx(str, ast.GetChild()[1], pr, indent)
	if priority > p {
		SmartStrAppendcEx(str, ')', 0)
	}
	return
prefix_op:
	if priority > p {
		SmartStrAppendcEx(str, '(', 0)
	}
	SmartStrAppendlEx(str, op, strlen(op), 0)
	ZendAstExportEx(str, ast.GetChild()[0], pl, indent)
	if priority > p {
		SmartStrAppendcEx(str, ')', 0)
	}
	return
postfix_op:
	if priority > p {
		SmartStrAppendcEx(str, '(', 0)
	}
	ZendAstExportEx(str, ast.GetChild()[0], pl, indent)
	SmartStrAppendlEx(str, op, strlen(op), 0)
	if priority > p {
		SmartStrAppendcEx(str, ')', 0)
	}
	return
func_op:
	SmartStrAppendlEx(str, op, strlen(op), 0)
	SmartStrAppendcEx(str, '(', 0)
	ZendAstExportEx(str, ast.GetChild()[0], 0, indent)
	SmartStrAppendcEx(str, ')', 0)
	return
append_node_1:
	SmartStrAppendlEx(str, op, strlen(op), 0)
	if ast.GetChild()[0] != nil {
		SmartStrAppendcEx(str, ' ', 0)
		ast = ast.GetChild()[0]
		goto tail_call
	}
	return
append_str:
	SmartStrAppendlEx(str, op, strlen(op), 0)
	return
append_default_value:
	if ast.GetChild()[p] != nil {
		SmartStrAppendlEx(str, " = ", strlen(" = "), 0)
		ast = ast.GetChild()[p]
		goto tail_call
	}
	return
}
func ZendAstExport(prefix string, ast *ZendAst, suffix string) *ZendString {
	var str SmartStr = SmartStr{0}
	SmartStrAppendlEx(&str, prefix, strlen(prefix), 0)
	ZendAstExportEx(&str, ast, 0, 0)
	SmartStrAppendlEx(&str, suffix, strlen(suffix), 0)
	SmartStr0(&str)
	return str.GetS()
}
