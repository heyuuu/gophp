package zend

import (
	b "sik/builtin"
	"sik/zend/faults"
	"sik/zend/types"
)

func ZEND_AST_SPEC_CALL(name __auto__) __auto__ {
	return ZEND_EXPAND_VA(ZEND_AST_SPEC_CALL_(name, __VA_ARGS__, _4, _3, _2, _1, _0)(__VA_ARGS__))
}
func ZEND_AST_SPEC_CALL_EX(name func() __auto__) __auto__ {
	return ZEND_EXPAND_VA(ZEND_AST_SPEC_CALL_EX_(name, __VA_ARGS__, _4, _3, _2, _1, _0)(__VA_ARGS__))
}
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
func ZendAstCreateEx4(
	kind ZendAstKind,
	attr ZendAstAttr,
	child1 *ZendAst,
	child2 *ZendAst,
	child3 *ZendAst,
	child4 *ZendAst,
) *ZendAst {
	var ast *ZendAst = ZendAstCreate4(kind, child1, child2, child3, child4)
	ast.SetAttr(attr)
	return ast
}
func ZendAstCreate() __auto__ {
	return ZEND_AST_SPEC_CALL(ZendAstCreate, __VA_ARGS__)
}
func ZendAstCreateEx() __auto__ {
	return ZEND_AST_SPEC_CALL_EX(ZendAstCreateEx, __VA_ARGS__)
}
func ZendAstCreateList(init_children int) __auto__ {
	return ZEND_AST_SPEC_CALL(ZendAstCreateList, __VA_ARGS__)
}
func ZendAstIsSpecial(ast *ZendAst) types.ZendBool {
	return ast.GetKind() >> ZEND_AST_SPECIAL_SHIFT & 1
}
func ZendAstIsList(ast *ZendAst) types.ZendBool {
	return ast.GetKind() >> ZEND_AST_IS_LIST_SHIFT & 1
}
func ZendAstGetList(ast *ZendAst) *ZendAstList {
	b.Assert(ZendAstIsList(ast) != 0)
	return (*ZendAstList)(ast)
}
func ZendAstGetZval(ast *ZendAst) *types.Zval {
	b.Assert(ast.GetKind() == ZEND_AST_ZVAL)
	return (*ZendAstZval)(ast).GetVal()
}
func ZendAstGetStr(ast *ZendAst) *types.String {
	var zv *types.Zval = ZendAstGetZval(ast)
	b.Assert(zv.IsString())
	return zv.GetStr()
}
func ZendAstGetConstantName(ast *ZendAst) *types.String {
	b.Assert(ast.GetKind() == ZEND_AST_CONSTANT)
	b.Assert((*ZendAstZval)(ast).GetVal().IsString())
	return (*ZendAstZval)(ast).GetVal().GetStr()
}
func ZendAstGetNumChildren(ast *ZendAst) uint32 {
	b.Assert(ZendAstIsList(ast) == 0)
	return ast.GetKind() >> ZEND_AST_NUM_CHILDREN_SHIFT
}
func ZendAstGetLineno(ast *ZendAst) uint32 {
	if ast.GetKind() == ZEND_AST_ZVAL {
		var zv *types.Zval = ZendAstGetZval(ast)
		return zv.GetLineno()
	} else {
		return ast.GetLineno()
	}
}
func ZendAstCreateBinaryOp(opcode uint32, op0 *ZendAst, op1 *ZendAst) *ZendAst {
	return ZendAstCreateEx(ZEND_AST_BINARY_OP, opcode, op0, op1)
}
func ZendAstCreateAssignOp(opcode uint32, op0 *ZendAst, op1 *ZendAst) *ZendAst {
	return ZendAstCreateEx(ZEND_AST_ASSIGN_OP, opcode, op0, op1)
}
func ZendAstCreateCast(type_ uint32, op0 *ZendAst) *ZendAst {
	return ZendAstCreateEx(ZEND_AST_CAST, type_, op0)
}
func ZendAstListRtrim(ast *ZendAst) *ZendAst {
	var list *ZendAstList = ZendAstGetList(ast)
	if list.GetChildren() != 0 && list.GetChild()[list.GetChildren()-1] == nil {
		list.GetChildren()--
	}
	return ast
}
func ZendAstAlloc(size int) any {
	return ZendArenaAlloc(CG__().GetAstArena(), size)
}
func ZendAstRealloc(old any, old_size int, new_size int) any {
	var new_ any = ZendAstAlloc(new_size)
	memcpy(new_, old, old_size)
	return new_
}
func ZendAstSize(children uint32) int {
	return b.SizeOf("zend_ast") - b.SizeOf("zend_ast *") + b.SizeOf("zend_ast *")*children
}
func ZendAstListSize(children uint32) int {
	return b.SizeOf("zend_ast_list") - b.SizeOf("zend_ast *") + b.SizeOf("zend_ast *")*children
}
func ZendAstCreateZnode(node *Znode) *ZendAst {
	var ast *ZendAstZnode
	ast = ZendAstAlloc(b.SizeOf("zend_ast_znode"))
	ast.SetKind(ZEND_AST_ZNODE)
	ast.SetAttr(0)
	ast.SetLineno(CG__().GetZendLineno())
	ast.SetNode(*node)
	return (*ZendAst)(ast)
}
func ZendAstCreateZvalInt(zv *types.Zval, attr uint32, lineno uint32) *ZendAst {
	var ast *ZendAstZval
	ast = ZendAstAlloc(b.SizeOf("zend_ast_zval"))
	ast.SetKind(ZEND_AST_ZVAL)
	ast.SetAttr(attr)
	types.ZVAL_COPY_VALUE(ast.GetVal(), zv)
	ast.GetVal().GetLineno() = lineno
	return (*ZendAst)(ast)
}
func ZendAstCreateZvalWithLineno(zv *types.Zval, lineno uint32) *ZendAst {
	return ZendAstCreateZvalInt(zv, 0, lineno)
}
func ZendAstCreateZvalEx(zv *types.Zval, attr ZendAstAttr) *ZendAst {
	return ZendAstCreateZvalInt(zv, attr, CG__().GetZendLineno())
}
func ZendAstCreateZval(zv *types.Zval) *ZendAst {
	return ZendAstCreateZvalInt(zv, 0, CG__().GetZendLineno())
}
func ZendAstCreateZvalFromStr(str *types.String) *ZendAst {
	var zv types.Zval
	zv.SetString(str)
	return ZendAstCreateZvalInt(&zv, 0, CG__().GetZendLineno())
}
func ZendAstCreateZvalFromLong(lval ZendLong) *ZendAst {
	var zv types.Zval
	zv.SetLong(lval)
	return ZendAstCreateZvalInt(&zv, 0, CG__().GetZendLineno())
}
func ZendAstCreateConstant(name *types.String, attr ZendAstAttr) *ZendAst {
	var ast *ZendAstZval
	ast = ZendAstAlloc(b.SizeOf("zend_ast_zval"))
	ast.SetKind(ZEND_AST_CONSTANT)
	ast.SetAttr(attr)
	ast.GetVal().SetString(name)
	ast.GetVal().GetLineno() = CG__().GetZendLineno()
	return (*ZendAst)(ast)
}
func ZendAstCreateClassConstOrName(class_name *ZendAst, name *ZendAst) *ZendAst {
	var name_str *types.String = ZendAstGetStr(name)
	if types.ZendStringEqualsLiteralCi(name_str, "class") {
		types.ZendStringRelease(name_str)
		return ZendAstCreate(ZEND_AST_CLASS_NAME, class_name)
	} else {
		return ZendAstCreate(ZEND_AST_CLASS_CONST, class_name, name)
	}
}
func ZendAstCreateDecl(
	kind ZendAstKind,
	flags uint32,
	start_lineno uint32,
	doc_comment *types.String,
	name *types.String,
	child0 *ZendAst,
	child1 *ZendAst,
	child2 *ZendAst,
	child3 *ZendAst,
) *ZendAst {
	var ast *ZendAstDecl
	ast = ZendAstAlloc(b.SizeOf("zend_ast_decl"))
	ast.SetKind(kind)
	ast.SetAttr(0)
	ast.SetStartLineno(start_lineno)
	ast.SetEndLineno(CG__().GetZendLineno())
	ast.SetFlags(flags)
	ast.SetLexPos(INI_SCNG__().GetYyText())
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
	b.Assert(kind>>ZEND_AST_NUM_CHILDREN_SHIFT == 0)
	ast = ZendAstAlloc(ZendAstSize(0))
	ast.SetKind(kind)
	ast.SetAttr(0)
	ast.SetLineno(CG__().GetZendLineno())
	return ast
}
func ZendAstCreate1(kind ZendAstKind, child *ZendAst) *ZendAst {
	var ast *ZendAst
	var lineno uint32
	b.Assert(kind>>ZEND_AST_NUM_CHILDREN_SHIFT == 1)
	ast = ZendAstAlloc(ZendAstSize(1))
	ast.SetKind(kind)
	ast.SetAttr(0)
	ast.GetChild()[0] = child
	if child != nil {
		lineno = ZendAstGetLineno(child)
	} else {
		lineno = CG__().GetZendLineno()
	}
	ast.SetLineno(lineno)
	ast.SetLineno(lineno)
	return ast
}
func ZendAstCreate2(kind ZendAstKind, child1 *ZendAst, child2 *ZendAst) *ZendAst {
	var ast *ZendAst
	var lineno uint32
	b.Assert(kind>>ZEND_AST_NUM_CHILDREN_SHIFT == 2)
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
		lineno = CG__().GetZendLineno()
	}
	ast.SetLineno(lineno)
	return ast
}
func ZendAstCreate3(kind ZendAstKind, child1 *ZendAst, child2 *ZendAst, child3 *ZendAst) *ZendAst {
	var ast *ZendAst
	var lineno uint32
	b.Assert(kind>>ZEND_AST_NUM_CHILDREN_SHIFT == 3)
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
		lineno = CG__().GetZendLineno()
	}
	ast.SetLineno(lineno)
	return ast
}
func ZendAstCreate4(kind ZendAstKind, child1 *ZendAst, child2 *ZendAst, child3 *ZendAst, child4 *ZendAst) *ZendAst {
	var ast *ZendAst
	var lineno uint32
	b.Assert(kind>>ZEND_AST_NUM_CHILDREN_SHIFT == 4)
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
		lineno = CG__().GetZendLineno()
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
	list.SetLineno(CG__().GetZendLineno())
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
		if lineno > CG__().GetZendLineno() {
			lineno = CG__().GetZendLineno()
		}
	} else {
		lineno = CG__().GetZendLineno()
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
		if lineno > CG__().GetZendLineno() {
			lineno = CG__().GetZendLineno()
		}
	} else if child2 != nil {
		lineno = ZendAstGetLineno(child2)
		if lineno > CG__().GetZendLineno() {
			lineno = CG__().GetZendLineno()
		}
	} else {
		list.SetChildren(0)
		lineno = CG__().GetZendLineno()
	}
	list.SetLineno(lineno)
	return ast
}
func IsPowerOfTwo(n uint32) types.ZendBool { return n != 0 && n == (n & ^n + 1) }
func ZendAstListAdd(ast *ZendAst, op *ZendAst) *ZendAst {
	var list *ZendAstList = ZendAstGetList(ast)
	if list.GetChildren() >= 4 && IsPowerOfTwo(list.GetChildren()) != 0 {
		list = ZendAstRealloc(list, ZendAstListSize(list.GetChildren()), ZendAstListSize(list.GetChildren()*2))
	}
	list.GetChild()[b.PostInc(&(list.GetChildren()))] = op
	return (*ZendAst)(list)
}
func ZendAstAddArrayElement(result *types.Zval, offset *types.Zval, expr *types.Zval) int {
	switch offset.GetType() {
	case types.IS_UNDEF:
		if result.GetArr().NextIndexInsert(expr) == nil {
			faults.Error(faults.E_WARNING, "Cannot add element to the array as the next element is already occupied")
			ZvalPtrDtorNogc(expr)
		}
	case types.IS_STRING:
		result.GetArr().SymtableUpdate(offset.GetStr().GetStr(), expr)

	case types.IS_NULL:
		result.GetArr().SymtableUpdate(types.ZSTR_EMPTY_ALLOC().GetStr(), expr)
	case types.IS_LONG:
		result.GetArr().IndexUpdate(offset.GetLval(), expr)
	case types.IS_FALSE:
		result.GetArr().IndexUpdate(0, expr)
	case types.IS_TRUE:
		result.GetArr().IndexUpdate(1, expr)
	case types.IS_DOUBLE:
		result.GetArr().IndexUpdate(ZendDvalToLval(offset.GetDval()), expr)
	case types.IS_RESOURCE:
		faults.Error(faults.E_NOTICE, "Resource ID#%d used as offset, casting to integer (%d)", types.Z_RES_HANDLE_P(offset), types.Z_RES_HANDLE_P(offset))
		result.GetArr().IndexUpdate(types.Z_RES_HANDLE_P(offset), expr)
	default:
		faults.ThrowError(nil, "Illegal offset type")
		return types.FAILURE
	}
	return types.SUCCESS
}
func ZendAstAddUnpackedElement(result *types.Zval, expr *types.Zval) int {
	if expr.IsArray() {
		var ht *types.Array = expr.GetArr()
		var val *types.Zval
		var key *types.String
		var __ht *types.Array = ht
		for _, _p := range __ht.ForeachData() {
			var _z *types.Zval = _p.GetVal()

			key = _p.GetKey()
			val = _z
			if key != nil {
				faults.ThrowError(nil, "Cannot unpack array with string keys")
				return types.FAILURE
			} else {
				if result.GetArr().NextIndexInsert(val) == nil {
					faults.Error(faults.E_WARNING, "Cannot add element to the array as the next element is already occupied")
					break
				}
				val.TryAddRefcount()
			}
		}
		return types.SUCCESS
	}

	/* Objects or references cannot occur in a constant expression. */

	faults.ThrowError(nil, "Only arrays and Traversables can be unpacked")
	return types.FAILURE
}
func ZendAstEvaluate(result *types.Zval, ast *ZendAst, scope *types.ClassEntry) int {
	var op1 types.Zval
	var op2 types.Zval
	var ret int = types.SUCCESS
	switch ast.GetKind() {
	case ZEND_AST_BINARY_OP:
		if ZendAstEvaluate(&op1, ast.GetChild()[0], scope) != types.SUCCESS {
			ret = types.FAILURE
		} else if ZendAstEvaluate(&op2, ast.GetChild()[1], scope) != types.SUCCESS {
			ZvalPtrDtorNogc(&op1)
			ret = types.FAILURE
		} else {
			var op BinaryOpType = GetBinaryOp(ast.GetAttr())
			ret = op(result, &op1, &op2)
			ZvalPtrDtorNogc(&op1)
			ZvalPtrDtorNogc(&op2)
		}
	case ZEND_AST_GREATER:
		fallthrough
	case ZEND_AST_GREATER_EQUAL:
		if ZendAstEvaluate(&op1, ast.GetChild()[0], scope) != types.SUCCESS {
			ret = types.FAILURE
		} else if ZendAstEvaluate(&op2, ast.GetChild()[1], scope) != types.SUCCESS {
			ZvalPtrDtorNogc(&op1)
			ret = types.FAILURE
		} else {

			/* op1 > op2 is the same as op2 < op1 */

			var op BinaryOpType = b.Cond(ast.GetKind() == ZEND_AST_GREATER, IsSmallerFunction, IsSmallerOrEqualFunction)
			ret = op(result, &op2, &op1)
			ZvalPtrDtorNogc(&op1)
			ZvalPtrDtorNogc(&op2)
		}
	case ZEND_AST_UNARY_OP:
		if ZendAstEvaluate(&op1, ast.GetChild()[0], scope) != types.SUCCESS {
			ret = types.FAILURE
		} else {
			var op UnaryOpType = GetUnaryOp(ast.GetAttr())
			ret = op(result, &op1)
			ZvalPtrDtorNogc(&op1)
		}
	case ZEND_AST_ZVAL:
		var zv *types.Zval = ZendAstGetZval(ast)
		types.ZVAL_COPY(result, zv)
	case ZEND_AST_CONSTANT:
		var name *types.String = ZendAstGetConstantName(ast)
		var zv *types.Zval = ZendGetConstantEx(name, scope, ast.GetAttr())
		if zv == nil {
			result.SetUndef()
			ret = ZendUseUndefinedConstant(name, ast.GetAttr(), result)
			break
		}
		types.ZVAL_COPY_OR_DUP(result, zv)
	case ZEND_AST_CONSTANT_CLASS:
		if scope != nil {
			result.SetStringCopy(scope.GetName())
		} else {
			ZVAL_EMPTY_STRING(result)
		}
	case ZEND_AST_CLASS_NAME:
		if scope == nil {
			faults.ThrowError(nil, "Cannot use \"self\" when no class scope is active")
			return types.FAILURE
		}
		if ast.GetAttr() == ZEND_FETCH_CLASS_SELF {
			result.SetStringCopy(scope.GetName())
		} else if ast.GetAttr() == ZEND_FETCH_CLASS_PARENT {
			if !(scope.GetParent()) {
				faults.ThrowError(nil, "Cannot use \"parent\" when current class scope has no parent")
				return types.FAILURE
			}
			result.SetStringCopy(scope.GetParent().name)
		} else {
			b.Assert(false)
		}
	case ZEND_AST_AND:
		if ZendAstEvaluate(&op1, ast.GetChild()[0], scope) != types.SUCCESS {
			ret = types.FAILURE
			break
		}
		if ZendIsTrue(&op1) != 0 {
			if ZendAstEvaluate(&op2, ast.GetChild()[1], scope) != types.SUCCESS {
				ZvalPtrDtorNogc(&op1)
				ret = types.FAILURE
				break
			}
			types.ZVAL_BOOL(result, ZendIsTrue(&op2) != 0)
			ZvalPtrDtorNogc(&op2)
		} else {
			result.SetFalse()
		}
		ZvalPtrDtorNogc(&op1)
	case ZEND_AST_OR:
		if ZendAstEvaluate(&op1, ast.GetChild()[0], scope) != types.SUCCESS {
			ret = types.FAILURE
			break
		}
		if ZendIsTrue(&op1) != 0 {
			result.SetTrue()
		} else {
			if ZendAstEvaluate(&op2, ast.GetChild()[1], scope) != types.SUCCESS {
				ZvalPtrDtorNogc(&op1)
				ret = types.FAILURE
				break
			}
			types.ZVAL_BOOL(result, ZendIsTrue(&op2) != 0)
			ZvalPtrDtorNogc(&op2)
		}
		ZvalPtrDtorNogc(&op1)
	case ZEND_AST_CONDITIONAL:
		if ZendAstEvaluate(&op1, ast.GetChild()[0], scope) != types.SUCCESS {
			ret = types.FAILURE
			break
		}
		if ZendIsTrue(&op1) != 0 {
			if ast.GetChild()[1] == nil {
				*result = op1
			} else {
				if ZendAstEvaluate(result, ast.GetChild()[1], scope) != types.SUCCESS {
					ZvalPtrDtorNogc(&op1)
					ret = types.FAILURE
					break
				}
				ZvalPtrDtorNogc(&op1)
			}
		} else {
			if ZendAstEvaluate(result, ast.GetChild()[2], scope) != types.SUCCESS {
				ZvalPtrDtorNogc(&op1)
				ret = types.FAILURE
				break
			}
			ZvalPtrDtorNogc(&op1)
		}
	case ZEND_AST_COALESCE:
		if ZendAstEvaluate(&op1, ast.GetChild()[0], scope) != types.SUCCESS {
			ret = types.FAILURE
			break
		}
		if op1.GetType() > types.IS_NULL {
			*result = op1
		} else {
			if ZendAstEvaluate(result, ast.GetChild()[1], scope) != types.SUCCESS {
				ZvalPtrDtorNogc(&op1)
				ret = types.FAILURE
				break
			}
			ZvalPtrDtorNogc(&op1)
		}
	case ZEND_AST_UNARY_PLUS:
		if ZendAstEvaluate(&op2, ast.GetChild()[0], scope) != types.SUCCESS {
			ret = types.FAILURE
		} else {
			op1.SetLong(0)
			ret = AddFunction(result, &op1, &op2)
			ZvalPtrDtorNogc(&op2)
		}
	case ZEND_AST_UNARY_MINUS:
		if ZendAstEvaluate(&op2, ast.GetChild()[0], scope) != types.SUCCESS {
			ret = types.FAILURE
		} else {
			op1.SetLong(0)
			ret = SubFunction(result, &op1, &op2)
			ZvalPtrDtorNogc(&op2)
		}
	case ZEND_AST_ARRAY:
		var i uint32
		var list *ZendAstList = ZendAstGetList(ast)
		if list.GetChildren() == 0 {
			result.SetEmptyArray()
			break
		}
		ArrayInit(result)
		for i = 0; i < list.GetChildren(); i++ {
			var elem *ZendAst = list.GetChild()[i]
			if elem.GetKind() == ZEND_AST_UNPACK {
				if ZendAstEvaluate(&op1, elem.GetChild()[0], scope) != types.SUCCESS {
					ZvalPtrDtorNogc(result)
					return types.FAILURE
				}
				if ZendAstAddUnpackedElement(result, &op1) != types.SUCCESS {
					ZvalPtrDtorNogc(&op1)
					ZvalPtrDtorNogc(result)
					return types.FAILURE
				}
				ZvalPtrDtorNogc(&op1)
				continue
			}
			if elem.GetChild()[1] != nil {
				if ZendAstEvaluate(&op1, elem.GetChild()[1], scope) != types.SUCCESS {
					ZvalPtrDtorNogc(result)
					return types.FAILURE
				}
			} else {
				op1.SetUndef()
			}
			if ZendAstEvaluate(&op2, elem.GetChild()[0], scope) != types.SUCCESS {
				ZvalPtrDtorNogc(&op1)
				ZvalPtrDtorNogc(result)
				return types.FAILURE
			}
			if ZendAstAddArrayElement(result, &op1, &op2) != types.SUCCESS {
				ZvalPtrDtorNogc(&op1)
				ZvalPtrDtorNogc(&op2)
				ZvalPtrDtorNogc(result)
				return types.FAILURE
			}
		}
	case ZEND_AST_DIM:
		if ast.GetChild()[1] == nil {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot use [] for reading")
		}
		if ZendAstEvaluate(&op1, ast.GetChild()[0], scope) != types.SUCCESS {
			ret = types.FAILURE
		} else if ZendAstEvaluate(&op2, ast.GetChild()[1], scope) != types.SUCCESS {
			ZvalPtrDtorNogc(&op1)
			ret = types.FAILURE
		} else {
			ZendFetchDimensionConst(result, &op1, &op2, b.Cond((ast.GetAttr()&ZEND_DIM_IS) != 0, BP_VAR_IS, BP_VAR_R))
			ZvalPtrDtorNogc(&op1)
			ZvalPtrDtorNogc(&op2)
		}
	default:
		faults.ThrowError(nil, "Unsupported constant expression")
		ret = types.FAILURE
	}
	return ret
}
func ZendAstTreeSize(ast *ZendAst) int {
	var size int
	if ast.GetKind() == ZEND_AST_ZVAL || ast.GetKind() == ZEND_AST_CONSTANT {
		size = b.SizeOf("zend_ast_zval")
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
		types.ZVAL_COPY(new_.GetVal(), ZendAstGetZval(ast))
		buf = any((*byte)(buf + b.SizeOf("zend_ast_zval")))
	} else if ast.GetKind() == ZEND_AST_CONSTANT {
		var new_ *ZendAstZval = (*ZendAstZval)(buf)
		new_.SetKind(ZEND_AST_CONSTANT)
		new_.SetAttr(ast.GetAttr())
		new_.GetVal().SetStringCopy(ZendAstGetConstantName(ast))
		buf = any((*byte)(buf + b.SizeOf("zend_ast_zval")))
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
func ZendAstCopy(ast *ZendAst) *types.ZendAstRef {
	var tree_size int
	var ref *types.ZendAstRef
	b.Assert(ast != nil)
	tree_size = ZendAstTreeSize(ast) + b.SizeOf("zend_ast_ref")
	ref = Emalloc(tree_size)
	ZendAstTreeCopy(ast, types.GC_AST(ref))
	ref.SetRefcount(1)
	ref.GetGcTypeInfo() = types.IS_CONSTANT_AST
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
		types.ZendStringReleaseEx(ZendAstGetConstantName(ast), 0)
	} else if ast.GetKind() >= ZEND_AST_FUNC_DECL {
		var decl *ZendAstDecl = (*ZendAstDecl)(ast)
		if decl.GetName() != nil {
			types.ZendStringReleaseEx(decl.GetName(), 0)
		}
		if decl.GetDocComment() != nil {
			types.ZendStringReleaseEx(decl.GetDocComment(), 0)
		}
		ZendAstDestroy(decl.GetChild()[0])
		ZendAstDestroy(decl.GetChild()[1])
		ZendAstDestroy(decl.GetChild()[2])
		ast = decl.GetChild()[3]
		goto tail_call
	}
}
func ZendAstRefDestroy(ast *types.ZendAstRef) {
	ZendAstDestroy(types.GC_AST(ast))
	Efree(ast)
}
func ZendAstApply(ast *ZendAst, fn ZendAstApplyFunc) {
	if ZendAstIsList(ast) != 0 {
		var list *ZendAstList = ZendAstGetList(ast)
		var i uint32
		for i = 0; i < list.GetChildren(); i++ {
			fn(list.GetChild()[i])
		}
	} else {
		var i uint32
		var children uint32 = ZendAstGetNumChildren(ast)
		for i = 0; i < children; i++ {
			fn(ast.GetChild()[i])
		}
	}
}
func ZendAstExportStr(str *SmartStr, s *types.String) {
	var i int
	for i = 0; i < s.GetLen(); i++ {
		var c uint8 = s.GetVal()[i]
		if c == '\'' || c == '\\' {
			str.AppendByte('\\')
			str.AppendByte(c)
		} else {
			str.AppendByte(c)
		}
	}
}
func ZendAstExportQstr(str *SmartStr, quote byte, s *types.String) {
	var i int
	for i = 0; i < s.GetLen(); i++ {
		var c uint8 = s.GetVal()[i]
		if c < ' ' {
			switch c {
			case '\n':
				str.AppendString("\\n")
			case '\r':
				str.AppendString("\\r")
			case '\t':
				str.AppendString("\\t")
			case 'f':
				str.AppendString("\\f")
			case 'v':
				str.AppendString("\\v")
			case 'e':
				str.AppendString("\\e")
			default:
				str.AppendString("\\0")
				str.AppendByte('0' + c/8)
				str.AppendByte('0' + c%8)
			}
		} else {
			if c == quote || c == '$' || c == '\\' {
				str.AppendByte('\\')
			}
			str.AppendByte(c)
		}
	}
}
func ZendAstExportIndent(str *SmartStr, indent int) {
	for indent > 0 {
		str.AppendString("    ")
		indent--
	}
}
func ZendAstExportName(str *SmartStr, ast *ZendAst, priority int, indent int) {
	if ast.GetKind() == ZEND_AST_ZVAL {
		var zv *types.Zval = ZendAstGetZval(ast)
		if zv.IsString() {
			str.AppendString(zv.GetStr().GetStr())
			return
		}
	}
	ZendAstExportEx(str, ast, priority, indent)
}
func ZendAstExportNsName(str *SmartStr, ast *ZendAst, priority int, indent int) {
	if ast.GetKind() == ZEND_AST_ZVAL {
		var zv *types.Zval = ZendAstGetZval(ast)
		if zv.IsString() {
			if ast.GetAttr() == ZEND_NAME_FQ {
				str.AppendByte('\\')
			} else if ast.GetAttr() == ZEND_NAME_RELATIVE {
				str.AppendString("namespace\\")
			}
			str.AppendString(zv.GetStr().GetStr())
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
		var zv *types.Zval = ZendAstGetZval(ast)
		if zv.IsString() && ZendAstValidVarName(zv.GetStr().GetVal(), zv.GetStr().GetLen()) != 0 {
			str.AppendString(zv.GetStr().GetStr())
			return
		}
	} else if ast.GetKind() == ZEND_AST_VAR {
		ZendAstExportEx(str, ast, 0, indent)
		return
	}
	str.AppendByte('{')
	ZendAstExportName(str, ast, 0, indent)
	str.AppendByte('}')
}
func ZendAstExportList(str *SmartStr, list *ZendAstList, separator int, priority int, indent int) {
	var i uint32 = 0
	for i < list.GetChildren() {
		if i != 0 && separator != 0 {
			str.AppendString(", ")
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
			var zv *types.Zval = ZendAstGetZval(ast)
			b.Assert(zv.IsString())
			ZendAstExportQstr(str, quote, zv.GetStr())
		} else if ast.GetKind() == ZEND_AST_VAR && ast.GetChild()[0].GetKind() == ZEND_AST_ZVAL && (i+1 == list.GetChildren() || list.GetChild()[i+1].GetKind() != ZEND_AST_ZVAL || ZendAstVarNeedsBraces((*types.Z_STRVAL_P)(ZendAstGetZval(list.GetChild()[i+1]))) == 0) {
			ZendAstExportEx(str, ast, 0, indent)
		} else {
			str.AppendByte('{')
			ZendAstExportEx(str, ast, 0, indent)
			str.AppendByte('}')
		}
		i++
	}
}
func ZendAstExportNameListEx(str *SmartStr, list *ZendAstList, indent int, separator string) {
	var i uint32 = 0
	for i < list.GetChildren() {
		if i != 0 {
			str.AppendString(b.CastStrAuto(separator))
		}
		ZendAstExportName(str, list.GetChild()[i], 0, indent)
		i++
	}
}
func ZendAstExportNameList(s *SmartStr, l *ZendAstList, i int) {
	ZendAstExportNameListEx(s, l, i, ", ")
}
func ZendAstExportCatchNameList(s *SmartStr, l *ZendAstList, i int) {
	ZendAstExportNameListEx(s, l, i, "|")
}
func ZendAstExportVarList(str *SmartStr, list *ZendAstList, indent int) {
	var i uint32 = 0
	for i < list.GetChildren() {
		if i != 0 {
			str.AppendString(", ")
		}
		if (list.GetChild()[i].GetAttr() & ZEND_BIND_REF) != 0 {
			str.AppendByte('&')
		}
		str.AppendByte('$')
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
			fallthrough
		case ZEND_AST_IF:
			fallthrough
		case ZEND_AST_SWITCH:
			fallthrough
		case ZEND_AST_WHILE:
			fallthrough
		case ZEND_AST_TRY:
			fallthrough
		case ZEND_AST_FOR:
			fallthrough
		case ZEND_AST_FOREACH:
			fallthrough
		case ZEND_AST_FUNC_DECL:
			fallthrough
		case ZEND_AST_METHOD:
			fallthrough
		case ZEND_AST_CLASS:
			fallthrough
		case ZEND_AST_USE_TRAIT:
			fallthrough
		case ZEND_AST_NAMESPACE:
			fallthrough
		case ZEND_AST_DECLARE:

		default:
			str.AppendByte(';')
		}
		str.AppendByte('\n')
	}
}
func ZendAstExportIfStmt(str *SmartStr, list *ZendAstList, indent int) {
	var i uint32
	var ast *ZendAst
tail_call:
	i = 0
	for i < list.GetChildren() {
		ast = list.GetChild()[i]
		b.Assert(ast.GetKind() == ZEND_AST_IF_ELEM)
		if ast.GetChild()[0] != nil {
			if i == 0 {
				str.AppendString("if (")
			} else {
				ZendAstExportIndent(str, indent)
				str.AppendString("} elseif (")
			}
			ZendAstExportEx(str, ast.GetChild()[0], 0, indent)
			str.AppendString(") {\n")
			ZendAstExportStmt(str, ast.GetChild()[1], indent+1)
		} else {
			ZendAstExportIndent(str, indent)
			str.AppendString("} else ")
			if ast.GetChild()[1] != nil && ast.GetChild()[1].GetKind() == ZEND_AST_IF {
				list = (*ZendAstList)(ast.GetChild()[1])
				goto tail_call
			} else {
				str.AppendString("{\n")
				ZendAstExportStmt(str, ast.GetChild()[1], indent+1)
			}
		}
		i++
	}
	ZendAstExportIndent(str, indent)
	str.AppendByte('}')
}
func ZendAstExportZval(str *SmartStr, zv *types.Zval, priority int, indent int) {
	var idx ZendLong
	var key *types.String
	var val *types.Zval
	var first int
	zv = types.ZVAL_DEREF(zv)
	switch zv.GetType() {
	case types.IS_NULL:
		str.AppendString("null")
	case types.IS_FALSE:
		str.AppendString("false")
	case types.IS_TRUE:
		str.AppendString("true")
	case types.IS_LONG:
		str.AppendLong(zv.GetLval())
	case types.IS_DOUBLE:
		key = ZendStrpprintf(0, "%.*G", int(EG__().GetPrecision()), zv.GetDval())
		str.AppendString(key.GetStr())
		types.ZendStringReleaseEx(key, 0)
	case types.IS_STRING:
		str.AppendByte('\'')
		ZendAstExportStr(str, zv.GetStr())
		str.AppendByte('\'')
	case types.IS_ARRAY:
		str.AppendByte('[')
		first = 1
		var __ht *types.Array = types.Z_ARRVAL_P(zv)
		for _, _p := range __ht.ForeachData() {
			var _z *types.Zval = _p.GetVal()

			idx = _p.GetH()
			key = _p.GetKey()
			val = _z
			if first != 0 {
				first = 0
			} else {
				str.AppendString(", ")
			}
			if key != nil {
				str.AppendByte('\'')
				ZendAstExportStr(str, key)
				str.AppendString("' => ")
			} else {
				str.AppendLong(idx)
				str.AppendString(" => ")
			}
			ZendAstExportZval(str, val, 0, indent)
		}
		str.AppendByte(']')
	case types.IS_CONSTANT_AST:
		ZendAstExportEx(str, types.Z_ASTVAL_P(zv), priority, indent)
	default:

	}
}
func ZendAstExportClassNoHeader(str *SmartStr, decl *ZendAstDecl, indent int) {
	if decl.GetChild()[0] != nil {
		str.AppendString(" extends ")
		ZendAstExportNsName(str, decl.GetChild()[0], 0, indent)
	}
	if decl.GetChild()[1] != nil {
		str.AppendString(" implements ")
		ZendAstExportEx(str, decl.GetChild()[1], 0, indent)
	}
	str.AppendString(" {\n")
	ZendAstExportStmt(str, decl.GetChild()[2], indent+1)
	ZendAstExportIndent(str, indent)
	str.AppendString("}")
}
func BINARY_OP(_op string, _p int, _pl int, _pr int) {
	op = _op
	p = _p
	pl = _pl
	pr = _pr
	goto binary_op
}
func PREFIX_OP(_op string, _p int, _pl int) {
	op = _op
	p = _p
	pl = _pl
	goto prefix_op
}
func FUNC_OP(_op string) {
	op = _op
	goto func_op
}
func POSTFIX_OP(_op string, _p int, _pl int) {
	op = _op
	p = _p
	pl = _pl
	goto postfix_op
}
func APPEND_NODE_1(_op string) {
	op = _op
	goto append_node_1
}
func APPEND_STR(_op string) {
	op = _op
	goto append_str
}
func APPEND_DEFAULT_VALUE(n int) {
	p = n
	goto append_default_value
}
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
	case ZEND_AST_CONSTANT:
		var name *types.String = ZendAstGetConstantName(ast)
		str.AppendString(name.GetStr())
	case ZEND_AST_CONSTANT_CLASS:
		str.AppendString("__CLASS__")
	case ZEND_AST_ZNODE:

		/* This AST kind is only used for temporary nodes during compilation */

		b.Assert(false)
	case ZEND_AST_FUNC_DECL:
		fallthrough
	case ZEND_AST_CLOSURE:
		fallthrough
	case ZEND_AST_ARROW_FUNC:
		fallthrough
	case ZEND_AST_METHOD:
		decl = (*ZendAstDecl)(ast)
		if decl.IsPublic() {
			str.AppendString("public ")
		} else if decl.IsProtected() {
			str.AppendString("protected ")
		} else if decl.IsPrivate() {
			str.AppendString("private ")
		}
		if decl.IsStatic() {
			str.AppendString("static ")
		}
		if decl.IsAbstract() {
			str.AppendString("abstract ")
		}
		if decl.IsFinal() {
			str.AppendString("final ")
		}
		if decl.GetKind() == ZEND_AST_ARROW_FUNC {
			str.AppendString("fn")
		} else {
			str.AppendString("function ")
		}
		if decl.IsReturnReference() {
			str.AppendByte('&')
		}
		if ast.GetKind() != ZEND_AST_CLOSURE && ast.GetKind() != ZEND_AST_ARROW_FUNC {
			str.AppendString(decl.GetName().GetStr())
		}
		str.AppendByte('(')
		ZendAstExportEx(str, decl.GetChild()[0], 0, indent)
		str.AppendByte(')')
		ZendAstExportEx(str, decl.GetChild()[1], 0, indent)
		if decl.GetChild()[3] != nil {
			str.AppendString(": ")
			if (decl.GetChild()[3].GetAttr() & ZEND_TYPE_NULLABLE) != 0 {
				str.AppendByte('?')
			}
			ZendAstExportNsName(str, decl.GetChild()[3], 0, indent)
		}
		if decl.GetChild()[2] != nil {
			if decl.GetKind() == ZEND_AST_ARROW_FUNC {
				b.Assert(decl.GetChild()[2].GetKind() == ZEND_AST_RETURN)
				str.AppendString(" => ")
				ZendAstExportEx(str, decl.GetChild()[2].GetChild()[0], 0, indent)
				break
			}
			str.AppendString(" {\n")
			ZendAstExportStmt(str, decl.GetChild()[2], indent+1)
			ZendAstExportIndent(str, indent)
			str.AppendByte('}')
			if ast.GetKind() != ZEND_AST_CLOSURE {
				str.AppendByte('\n')
			}
		} else {
			str.AppendString(";\n")
		}
	case ZEND_AST_CLASS:
		decl = (*ZendAstDecl)(ast)
		if decl.IsInterface() {
			str.AppendString("interface ")
		} else if decl.IsTrait() {
			str.AppendString("trait ")
		} else {
			if decl.IsExplicitAbstractClass() {
				str.AppendString("abstract ")
			}
			if decl.IsFinal() {
				str.AppendString("final ")
			}
			str.AppendString("class ")
		}
		str.AppendString(decl.GetName().GetStr())
		ZendAstExportClassNoHeader(str, decl, indent)
		str.AppendByte('\n')
	case ZEND_AST_ARG_LIST:
		fallthrough
	case ZEND_AST_EXPR_LIST:
		fallthrough
	case ZEND_AST_PARAM_LIST:
	simple_list:
		ZendAstExportList(str, (*ZendAstList)(ast), 1, 20, indent)
	case ZEND_AST_ARRAY:
		str.AppendByte('[')
		ZendAstExportList(str, (*ZendAstList)(ast), 1, 20, indent)
		str.AppendByte(']')
	case ZEND_AST_ENCAPS_LIST:
		str.AppendByte('"')
		ZendAstExportEncapsList(str, '"', (*ZendAstList)(ast), indent)
		str.AppendByte('"')
	case ZEND_AST_STMT_LIST:
		fallthrough
	case ZEND_AST_TRAIT_ADAPTATIONS:
		ZendAstExportStmt(str, ast, indent)
	case ZEND_AST_IF:
		ZendAstExportIfStmt(str, (*ZendAstList)(ast), indent)
	case ZEND_AST_SWITCH_LIST:
		fallthrough
	case ZEND_AST_CATCH_LIST:
		ZendAstExportList(str, (*ZendAstList)(ast), 0, 0, indent)
	case ZEND_AST_CLOSURE_USES:
		str.AppendString(" use(")
		ZendAstExportVarList(str, (*ZendAstList)(ast), indent)
		str.AppendByte(')')
	case ZEND_AST_PROP_GROUP:
		var type_ast *ZendAst = ast.GetChild()[0]
		var prop_ast *ZendAst = ast.GetChild()[1]
		if (ast.GetAttr() & ZEND_ACC_PUBLIC) != 0 {
			str.AppendString("public ")
		} else if (ast.GetAttr() & ZEND_ACC_PROTECTED) != 0 {
			str.AppendString("protected ")
		} else if (ast.GetAttr() & ZEND_ACC_PRIVATE) != 0 {
			str.AppendString("private ")
		}
		if (ast.GetAttr() & ZEND_ACC_STATIC) != 0 {
			str.AppendString("static ")
		}
		if type_ast != nil {
			if (type_ast.GetAttr() & ZEND_TYPE_NULLABLE) != 0 {
				str.AppendByte('?')
			}
			ZendAstExportNsName(str, type_ast, 0, indent)
			str.AppendByte(' ')
		}
		ast = prop_ast
		goto simple_list
	case ZEND_AST_CONST_DECL:
		fallthrough
	case ZEND_AST_CLASS_CONST_DECL:
		str.AppendString("const ")
		goto simple_list
	case ZEND_AST_NAME_LIST:
		ZendAstExportNameList(str, (*ZendAstList)(ast), indent)
	case ZEND_AST_USE:
		str.AppendString("use ")
		if ast.GetAttr() == T_FUNCTION {
			str.AppendString("function ")
		} else if ast.GetAttr() == T_CONST {
			str.AppendString("const ")
		}
		goto simple_list
	case ZEND_AST_MAGIC_CONST:
		switch ast.GetAttr() {
		case T_LINE:
			APPEND_STR("__LINE__")
			fallthrough
		case T_FILE:
			APPEND_STR("__FILE__")
			fallthrough
		case T_DIR:
			APPEND_STR("__DIR__")
			fallthrough
		case T_TRAIT_C:
			APPEND_STR("__TRAIT__")
			fallthrough
		case T_METHOD_C:
			APPEND_STR("__METHOD__")
			fallthrough
		case T_FUNC_C:
			APPEND_STR("__FUNCTION__")
			fallthrough
		case T_NS_C:
			APPEND_STR("__NAMESPACE__")
			fallthrough
		case T_CLASS_C:
			APPEND_STR("__CLASS__")
			fallthrough
		default:

		}
	case ZEND_AST_TYPE:
		switch ast.GetAttr() & ^ZEND_TYPE_NULLABLE {
		case types.IS_ARRAY:
			APPEND_STR("array")
			fallthrough
		case types.IS_CALLABLE:
			APPEND_STR("callable")
			fallthrough
		default:

		}
	case ZEND_AST_VAR:
		str.AppendByte('$')
		ZendAstExportVar(str, ast.GetChild()[0], 0, indent)
	case ZEND_AST_CONST:
		ZendAstExportNsName(str, ast.GetChild()[0], 0, indent)
	case ZEND_AST_UNPACK:
		str.AppendString("...")
		ast = ast.GetChild()[0]
		goto tail_call
	case ZEND_AST_UNARY_PLUS:
		PREFIX_OP("+", 240, 241)
		fallthrough
	case ZEND_AST_UNARY_MINUS:
		PREFIX_OP("-", 240, 241)
		fallthrough
	case ZEND_AST_CAST:
		switch ast.GetAttr() {
		case types.IS_NULL:
			PREFIX_OP("(unset)", 240, 241)
			fallthrough
		case types.IS_BOOL:
			PREFIX_OP("(bool)", 240, 241)
			fallthrough
		case types.IS_LONG:
			PREFIX_OP("(int)", 240, 241)
			fallthrough
		case types.IS_DOUBLE:
			PREFIX_OP("(double)", 240, 241)
			fallthrough
		case types.IS_STRING:
			PREFIX_OP("(string)", 240, 241)
			fallthrough
		case types.IS_ARRAY:
			PREFIX_OP("(array)", 240, 241)
			fallthrough
		case types.IS_OBJECT:
			PREFIX_OP("(object)", 240, 241)
			fallthrough
		default:

		}
	case ZEND_AST_EMPTY:
		FUNC_OP("empty")
		fallthrough
	case ZEND_AST_ISSET:
		FUNC_OP("isset")
		fallthrough
	case ZEND_AST_SILENCE:
		PREFIX_OP("@", 240, 241)
		fallthrough
	case ZEND_AST_SHELL_EXEC:
		str.AppendByte('`')
		if ast.GetChild()[0].GetKind() == ZEND_AST_ENCAPS_LIST {
			ZendAstExportEncapsList(str, '`', (*ZendAstList)(ast.GetChild()[0]), indent)
		} else {
			var zv *types.Zval
			b.Assert(ast.GetChild()[0].GetKind() == ZEND_AST_ZVAL)
			zv = ZendAstGetZval(ast.GetChild()[0])
			b.Assert(zv.IsString())
			ZendAstExportQstr(str, '`', zv.GetStr())
		}
		str.AppendByte('`')
	case ZEND_AST_CLONE:
		PREFIX_OP("clone ", 270, 271)
		fallthrough
	case ZEND_AST_EXIT:
		if ast.GetChild()[0] != nil {
			FUNC_OP("exit")
		} else {
			APPEND_STR("exit")
		}
	case ZEND_AST_PRINT:
		PREFIX_OP("print ", 60, 61)
		fallthrough
	case ZEND_AST_INCLUDE_OR_EVAL:
		switch ast.GetAttr() {
		case ZEND_INCLUDE_ONCE:
			FUNC_OP("include_once")
			fallthrough
		case ZEND_INCLUDE:
			FUNC_OP("include")
			fallthrough
		case ZEND_REQUIRE_ONCE:
			FUNC_OP("require_once")
			fallthrough
		case ZEND_REQUIRE:
			FUNC_OP("require")
			fallthrough
		case ZEND_EVAL:
			FUNC_OP("eval")
			fallthrough
		default:

		}
	case ZEND_AST_UNARY_OP:
		switch ast.GetAttr() {
		case ZEND_BW_NOT:
			PREFIX_OP("~", 240, 241)
			fallthrough
		case ZEND_BOOL_NOT:
			PREFIX_OP("!", 240, 241)
			fallthrough
		default:

		}
	case ZEND_AST_PRE_INC:
		PREFIX_OP("++", 240, 241)
		fallthrough
	case ZEND_AST_PRE_DEC:
		PREFIX_OP("--", 240, 241)
		fallthrough
	case ZEND_AST_POST_INC:
		POSTFIX_OP("++", 240, 241)
		fallthrough
	case ZEND_AST_POST_DEC:
		POSTFIX_OP("--", 240, 241)
		fallthrough
	case ZEND_AST_GLOBAL:
		APPEND_NODE_1("global")
		fallthrough
	case ZEND_AST_UNSET:
		FUNC_OP("unset")
		fallthrough
	case ZEND_AST_RETURN:
		APPEND_NODE_1("return")
		fallthrough
	case ZEND_AST_LABEL:
		ZendAstExportName(str, ast.GetChild()[0], 0, indent)
		str.AppendByte(':')
	case ZEND_AST_REF:
		str.AppendByte('&')
		ast = ast.GetChild()[0]
		goto tail_call
	case ZEND_AST_HALT_COMPILER:
		APPEND_STR("__HALT_COMPILER()")
		fallthrough
	case ZEND_AST_ECHO:
		APPEND_NODE_1("echo")
		fallthrough
	case ZEND_AST_THROW:
		APPEND_NODE_1("throw")
		fallthrough
	case ZEND_AST_GOTO:
		str.AppendString("goto ")
		ZendAstExportName(str, ast.GetChild()[0], 0, indent)
	case ZEND_AST_BREAK:
		APPEND_NODE_1("break")
		fallthrough
	case ZEND_AST_CONTINUE:
		APPEND_NODE_1("continue")
		fallthrough
	case ZEND_AST_DIM:
		ZendAstExportEx(str, ast.GetChild()[0], 260, indent)
		str.AppendByte('[')
		if ast.GetChild()[1] != nil {
			ZendAstExportEx(str, ast.GetChild()[1], 0, indent)
		}
		str.AppendByte(']')
	case ZEND_AST_PROP:
		ZendAstExportEx(str, ast.GetChild()[0], 0, indent)
		str.AppendString("->")
		ZendAstExportVar(str, ast.GetChild()[1], 0, indent)
	case ZEND_AST_STATIC_PROP:
		ZendAstExportNsName(str, ast.GetChild()[0], 0, indent)
		str.AppendString("::$")
		ZendAstExportVar(str, ast.GetChild()[1], 0, indent)
	case ZEND_AST_CALL:
		ZendAstExportNsName(str, ast.GetChild()[0], 0, indent)
		str.AppendByte('(')
		ZendAstExportEx(str, ast.GetChild()[1], 0, indent)
		str.AppendByte(')')
	case ZEND_AST_CLASS_CONST:
		ZendAstExportNsName(str, ast.GetChild()[0], 0, indent)
		str.AppendString("::")
		ZendAstExportName(str, ast.GetChild()[1], 0, indent)
	case ZEND_AST_CLASS_NAME:
		ZendAstExportNsName(str, ast.GetChild()[0], 0, indent)
		str.AppendString("::class")
	case ZEND_AST_ASSIGN:
		BINARY_OP(" = ", 90, 91, 90)
		fallthrough
	case ZEND_AST_ASSIGN_REF:
		BINARY_OP(" =& ", 90, 91, 90)
		fallthrough
	case ZEND_AST_ASSIGN_OP:
		switch ast.GetAttr() {
		case ZEND_ADD:
			BINARY_OP(" += ", 90, 91, 90)
			fallthrough
		case ZEND_SUB:
			BINARY_OP(" -= ", 90, 91, 90)
			fallthrough
		case ZEND_MUL:
			BINARY_OP(" *= ", 90, 91, 90)
			fallthrough
		case ZEND_DIV:
			BINARY_OP(" /= ", 90, 91, 90)
			fallthrough
		case ZEND_MOD:
			BINARY_OP(" %= ", 90, 91, 90)
			fallthrough
		case ZEND_SL:
			BINARY_OP(" <<= ", 90, 91, 90)
			fallthrough
		case ZEND_SR:
			BINARY_OP(" >>= ", 90, 91, 90)
			fallthrough
		case ZEND_CONCAT:
			BINARY_OP(" .= ", 90, 91, 90)
			fallthrough
		case ZEND_BW_OR:
			BINARY_OP(" |= ", 90, 91, 90)
			fallthrough
		case ZEND_BW_AND:
			BINARY_OP(" &= ", 90, 91, 90)
			fallthrough
		case ZEND_BW_XOR:
			BINARY_OP(" ^= ", 90, 91, 90)
			fallthrough
		case ZEND_POW:
			BINARY_OP(" **= ", 90, 91, 90)
			fallthrough
		default:

		}
	case ZEND_AST_ASSIGN_COALESCE:
		BINARY_OP(" ??= ", 90, 91, 90)
		fallthrough
	case ZEND_AST_BINARY_OP:
		switch ast.GetAttr() {
		case ZEND_ADD:
			BINARY_OP(" + ", 200, 200, 201)
			fallthrough
		case ZEND_SUB:
			BINARY_OP(" - ", 200, 200, 201)
			fallthrough
		case ZEND_MUL:
			BINARY_OP(" * ", 210, 210, 211)
			fallthrough
		case ZEND_DIV:
			BINARY_OP(" / ", 210, 210, 211)
			fallthrough
		case ZEND_MOD:
			BINARY_OP(" % ", 210, 210, 211)
			fallthrough
		case ZEND_SL:
			BINARY_OP(" << ", 190, 190, 191)
			fallthrough
		case ZEND_SR:
			BINARY_OP(" >> ", 190, 190, 191)
			fallthrough
		case ZEND_PARENTHESIZED_CONCAT:
			fallthrough
		case ZEND_CONCAT:
			BINARY_OP(" . ", 200, 200, 201)
			fallthrough
		case ZEND_BW_OR:
			BINARY_OP(" | ", 140, 140, 141)
			fallthrough
		case ZEND_BW_AND:
			BINARY_OP(" & ", 160, 160, 161)
			fallthrough
		case ZEND_BW_XOR:
			BINARY_OP(" ^ ", 150, 150, 151)
			fallthrough
		case ZEND_IS_IDENTICAL:
			BINARY_OP(" === ", 170, 171, 171)
			fallthrough
		case ZEND_IS_NOT_IDENTICAL:
			BINARY_OP(" !== ", 170, 171, 171)
			fallthrough
		case ZEND_IS_EQUAL:
			BINARY_OP(" == ", 170, 171, 171)
			fallthrough
		case ZEND_IS_NOT_EQUAL:
			BINARY_OP(" != ", 170, 171, 171)
			fallthrough
		case ZEND_IS_SMALLER:
			BINARY_OP(" < ", 180, 181, 181)
			fallthrough
		case ZEND_IS_SMALLER_OR_EQUAL:
			BINARY_OP(" <= ", 180, 181, 181)
			fallthrough
		case ZEND_POW:
			BINARY_OP(" ** ", 250, 251, 250)
			fallthrough
		case ZEND_BOOL_XOR:
			BINARY_OP(" xor ", 40, 40, 41)
			fallthrough
		case ZEND_SPACESHIP:
			BINARY_OP(" <=> ", 180, 181, 181)
			fallthrough
		default:

		}
	case ZEND_AST_GREATER:
		BINARY_OP(" > ", 180, 181, 181)
		fallthrough
	case ZEND_AST_GREATER_EQUAL:
		BINARY_OP(" >= ", 180, 181, 181)
		fallthrough
	case ZEND_AST_AND:
		BINARY_OP(" && ", 130, 130, 131)
		fallthrough
	case ZEND_AST_OR:
		BINARY_OP(" || ", 120, 120, 121)
		fallthrough
	case ZEND_AST_ARRAY_ELEM:
		if ast.GetChild()[1] != nil {
			ZendAstExportEx(str, ast.GetChild()[1], 80, indent)
			str.AppendString(" => ")
		}
		if ast.GetAttr() != 0 {
			str.AppendByte('&')
		}
		ZendAstExportEx(str, ast.GetChild()[0], 80, indent)
	case ZEND_AST_NEW:
		str.AppendString("new ")
		if ast.GetChild()[0].GetKind() == ZEND_AST_CLASS {
			str.AppendString("class")
			if ZendAstGetList(ast.GetChild()[1]).GetChildren() != 0 {
				str.AppendByte('(')
				ZendAstExportEx(str, ast.GetChild()[1], 0, indent)
				str.AppendByte(')')
			}
			ZendAstExportClassNoHeader(str, (*ZendAstDecl)(ast.GetChild()[0]), indent)
		} else {
			ZendAstExportNsName(str, ast.GetChild()[0], 0, indent)
			str.AppendByte('(')
			ZendAstExportEx(str, ast.GetChild()[1], 0, indent)
			str.AppendByte(')')
		}
	case ZEND_AST_INSTANCEOF:
		ZendAstExportEx(str, ast.GetChild()[0], 0, indent)
		str.AppendString(" instanceof ")
		ZendAstExportNsName(str, ast.GetChild()[1], 0, indent)
	case ZEND_AST_YIELD:
		if priority > 70 {
			str.AppendByte('(')
		}
		str.AppendString("yield ")
		if ast.GetChild()[0] != nil {
			if ast.GetChild()[1] != nil {
				ZendAstExportEx(str, ast.GetChild()[1], 70, indent)
				str.AppendString(" => ")
			}
			ZendAstExportEx(str, ast.GetChild()[0], 70, indent)
		}
		if priority > 70 {
			str.AppendByte(')')
		}
	case ZEND_AST_YIELD_FROM:
		PREFIX_OP("yield from ", 85, 86)
		fallthrough
	case ZEND_AST_COALESCE:
		BINARY_OP(" ?? ", 110, 111, 110)
		fallthrough
	case ZEND_AST_STATIC:
		str.AppendString("static $")
		ZendAstExportName(str, ast.GetChild()[0], 0, indent)
		APPEND_DEFAULT_VALUE(1)
		fallthrough
	case ZEND_AST_WHILE:
		str.AppendString("while (")
		ZendAstExportEx(str, ast.GetChild()[0], 0, indent)
		str.AppendString(") {\n")
		ZendAstExportStmt(str, ast.GetChild()[1], indent+1)
		ZendAstExportIndent(str, indent)
		str.AppendByte('}')
	case ZEND_AST_DO_WHILE:
		str.AppendString("do {\n")
		ZendAstExportStmt(str, ast.GetChild()[0], indent+1)
		ZendAstExportIndent(str, indent)
		str.AppendString("} while (")
		ZendAstExportEx(str, ast.GetChild()[1], 0, indent)
		str.AppendByte(')')
	case ZEND_AST_IF_ELEM:
		if ast.GetChild()[0] != nil {
			str.AppendString("if (")
			ZendAstExportEx(str, ast.GetChild()[0], 0, indent)
			str.AppendString(") {\n")
			ZendAstExportStmt(str, ast.GetChild()[1], indent+1)
		} else {
			str.AppendString("else {\n")
			ZendAstExportStmt(str, ast.GetChild()[1], indent+1)
		}
		ZendAstExportIndent(str, indent)
		str.AppendByte('}')
	case ZEND_AST_SWITCH:
		str.AppendString("switch (")
		ZendAstExportEx(str, ast.GetChild()[0], 0, indent)
		str.AppendString(") {\n")
		ZendAstExportEx(str, ast.GetChild()[1], 0, indent+1)
		ZendAstExportIndent(str, indent)
		str.AppendByte('}')
	case ZEND_AST_SWITCH_CASE:
		ZendAstExportIndent(str, indent)
		if ast.GetChild()[0] != nil {
			str.AppendString("case ")
			ZendAstExportEx(str, ast.GetChild()[0], 0, indent)
			str.AppendString(":\n")
		} else {
			str.AppendString("default:\n")
		}
		ZendAstExportStmt(str, ast.GetChild()[1], indent+1)
	case ZEND_AST_DECLARE:
		str.AppendString("declare(")
		b.Assert(ast.GetChild()[0].GetKind() == ZEND_AST_CONST_DECL)
		ZendAstExportList(str, (*ZendAstList)(ast.GetChild()[0]), 1, 0, indent)
		str.AppendByte(')')
		if ast.GetChild()[1] != nil {
			str.AppendString(" {\n")
			ZendAstExportStmt(str, ast.GetChild()[1], indent+1)
			ZendAstExportIndent(str, indent)
			str.AppendByte('}')
		} else {
			str.AppendByte(';')
		}
	case ZEND_AST_PROP_ELEM:
		str.AppendByte('$')
		fallthrough
	case ZEND_AST_CONST_ELEM:
		ZendAstExportName(str, ast.GetChild()[0], 0, indent)
		APPEND_DEFAULT_VALUE(1)
		fallthrough
	case ZEND_AST_USE_TRAIT:
		str.AppendString("use ")
		ZendAstExportEx(str, ast.GetChild()[0], 0, indent)
		if ast.GetChild()[1] != nil {
			str.AppendString(" {\n")
			ZendAstExportEx(str, ast.GetChild()[1], 0, indent+1)
			ZendAstExportIndent(str, indent)
			str.AppendString("}")
		} else {
			str.AppendString(";")
		}
	case ZEND_AST_TRAIT_PRECEDENCE:
		ZendAstExportEx(str, ast.GetChild()[0], 0, indent)
		str.AppendString(" insteadof ")
		ZendAstExportEx(str, ast.GetChild()[1], 0, indent)
	case ZEND_AST_METHOD_REFERENCE:
		if ast.GetChild()[0] != nil {
			ZendAstExportName(str, ast.GetChild()[0], 0, indent)
			str.AppendString("::")
		}
		ZendAstExportName(str, ast.GetChild()[1], 0, indent)
	case ZEND_AST_NAMESPACE:
		str.AppendString("namespace")
		if ast.GetChild()[0] != nil {
			str.AppendByte(' ')
			ZendAstExportName(str, ast.GetChild()[0], 0, indent)
		}
		if ast.GetChild()[1] != nil {
			str.AppendString(" {\n")
			ZendAstExportStmt(str, ast.GetChild()[1], indent+1)
			ZendAstExportIndent(str, indent)
			str.AppendString("}\n")
		} else {
			str.AppendByte(';')
		}
	case ZEND_AST_USE_ELEM:
		fallthrough
	case ZEND_AST_TRAIT_ALIAS:
		ZendAstExportName(str, ast.GetChild()[0], 0, indent)
		if (ast.GetAttr() & ZEND_ACC_PUBLIC) != 0 {
			str.AppendString(" as public")
		} else if (ast.GetAttr() & ZEND_ACC_PROTECTED) != 0 {
			str.AppendString(" as protected")
		} else if (ast.GetAttr() & ZEND_ACC_PRIVATE) != 0 {
			str.AppendString(" as private")
		} else if ast.GetChild()[1] != nil {
			str.AppendString(" as")
		}
		if ast.GetChild()[1] != nil {
			str.AppendByte(' ')
			ZendAstExportName(str, ast.GetChild()[1], 0, indent)
		}
	case ZEND_AST_METHOD_CALL:
		ZendAstExportEx(str, ast.GetChild()[0], 0, indent)
		str.AppendString("->")
		ZendAstExportVar(str, ast.GetChild()[1], 0, indent)
		str.AppendByte('(')
		ZendAstExportEx(str, ast.GetChild()[2], 0, indent)
		str.AppendByte(')')
	case ZEND_AST_STATIC_CALL:
		ZendAstExportNsName(str, ast.GetChild()[0], 0, indent)
		str.AppendString("::")
		ZendAstExportVar(str, ast.GetChild()[1], 0, indent)
		str.AppendByte('(')
		ZendAstExportEx(str, ast.GetChild()[2], 0, indent)
		str.AppendByte(')')
	case ZEND_AST_CONDITIONAL:
		if priority > 100 {
			str.AppendByte('(')
		}
		ZendAstExportEx(str, ast.GetChild()[0], 100, indent)
		if ast.GetChild()[1] != nil {
			str.AppendString(" ? ")
			ZendAstExportEx(str, ast.GetChild()[1], 101, indent)
			str.AppendString(" : ")
		} else {
			str.AppendString(" ?: ")
		}
		ZendAstExportEx(str, ast.GetChild()[2], 101, indent)
		if priority > 100 {
			str.AppendByte(')')
		}
	case ZEND_AST_TRY:
		str.AppendString("try {\n")
		ZendAstExportStmt(str, ast.GetChild()[0], indent+1)
		ZendAstExportIndent(str, indent)
		ZendAstExportEx(str, ast.GetChild()[1], 0, indent)
		if ast.GetChild()[2] != nil {
			str.AppendString("} finally {\n")
			ZendAstExportStmt(str, ast.GetChild()[2], indent+1)
			ZendAstExportIndent(str, indent)
		}
		str.AppendByte('}')
	case ZEND_AST_CATCH:
		str.AppendString("} catch (")
		ZendAstExportCatchNameList(str, ZendAstGetList(ast.GetChild()[0]), indent)
		str.AppendString(" $")
		ZendAstExportVar(str, ast.GetChild()[1], 0, indent)
		str.AppendString(") {\n")
		ZendAstExportStmt(str, ast.GetChild()[2], indent+1)
		ZendAstExportIndent(str, indent)
	case ZEND_AST_PARAM:
		if ast.GetChild()[0] != nil {
			if (ast.GetChild()[0].GetAttr() & ZEND_TYPE_NULLABLE) != 0 {
				str.AppendByte('?')
			}
			ZendAstExportNsName(str, ast.GetChild()[0], 0, indent)
			str.AppendByte(' ')
		}
		if (ast.GetAttr() & ZEND_PARAM_REF) != 0 {
			str.AppendByte('&')
		}
		if (ast.GetAttr() & ZEND_PARAM_VARIADIC) != 0 {
			str.AppendString("...")
		}
		str.AppendByte('$')
		ZendAstExportName(str, ast.GetChild()[1], 0, indent)
		APPEND_DEFAULT_VALUE(2)
		fallthrough
	case ZEND_AST_FOR:
		str.AppendString("for (")
		ZendAstExportEx(str, ast.GetChild()[0], 0, indent)
		str.AppendByte(';')
		if ast.GetChild()[1] != nil {
			str.AppendByte(' ')
			ZendAstExportEx(str, ast.GetChild()[1], 0, indent)
		}
		str.AppendByte(';')
		if ast.GetChild()[2] != nil {
			str.AppendByte(' ')
			ZendAstExportEx(str, ast.GetChild()[2], 0, indent)
		}
		str.AppendString(") {\n")
		ZendAstExportStmt(str, ast.GetChild()[3], indent+1)
		ZendAstExportIndent(str, indent)
		str.AppendByte('}')
	case ZEND_AST_FOREACH:
		str.AppendString("foreach (")
		ZendAstExportEx(str, ast.GetChild()[0], 0, indent)
		str.AppendString(" as ")
		if ast.GetChild()[2] != nil {
			ZendAstExportEx(str, ast.GetChild()[2], 0, indent)
			str.AppendString(" => ")
		}
		ZendAstExportEx(str, ast.GetChild()[1], 0, indent)
		str.AppendString(") {\n")
		ZendAstExportStmt(str, ast.GetChild()[3], indent+1)
		ZendAstExportIndent(str, indent)
		str.AppendByte('}')
	default:

	}
	return
binary_op:
	if priority > p {
		str.AppendByte('(')
	}
	ZendAstExportEx(str, ast.GetChild()[0], pl, indent)
	str.AppendString(b.CastStrAuto(op))
	ZendAstExportEx(str, ast.GetChild()[1], pr, indent)
	if priority > p {
		str.AppendByte(')')
	}
	return
prefix_op:
	if priority > p {
		str.AppendByte('(')
	}
	str.AppendString(b.CastStrAuto(op))
	ZendAstExportEx(str, ast.GetChild()[0], pl, indent)
	if priority > p {
		str.AppendByte(')')
	}
	return
postfix_op:
	if priority > p {
		str.AppendByte('(')
	}
	ZendAstExportEx(str, ast.GetChild()[0], pl, indent)
	str.AppendString(b.CastStrAuto(op))
	if priority > p {
		str.AppendByte(')')
	}
	return
func_op:
	str.AppendString(b.CastStrAuto(op))
	str.AppendByte('(')
	ZendAstExportEx(str, ast.GetChild()[0], 0, indent)
	str.AppendByte(')')
	return
append_node_1:
	str.AppendString(b.CastStrAuto(op))
	if ast.GetChild()[0] != nil {
		str.AppendByte(' ')
		ast = ast.GetChild()[0]
		goto tail_call
	}
	return
append_str:
	str.AppendString(b.CastStrAuto(op))
	return
append_default_value:
	if ast.GetChild()[p] != nil {
		str.AppendString(" = ")
		ast = ast.GetChild()[p]
		goto tail_call
	}
	return
}
func ZendAstExport(prefix string, ast *ZendAst, suffix string) *types.String {
	var str SmartStr = MakeSmartStr(0)
	str.AppendString(b.CastStrAuto(prefix))
	ZendAstExportEx(&str, ast, 0, 0)
	str.AppendString(b.CastStrAuto(suffix))
	str.ZeroTail()
	return str.GetS()
}
