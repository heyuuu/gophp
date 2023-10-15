package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
)

func ZendAstGetZval(ast *ZendAst) *types.Zval {
	return ast.Val()
}
func ZendAstGetStr(ast *ZendAst) *types.String {
	var zv *types.Zval = ast.Val()
	b.Assert(zv.IsString())
	return zv.StringEx()
}
func ZendAstGetStrVal(ast *ZendAst) string {
	var zv *types.Zval = ast.Val()
	b.Assert(zv.IsString())
	return zv.String()
}
func ZendAstGetConstantName(ast *ZendAst) *types.String {
	b.Assert(ast.Kind() == ZEND_AST_CONSTANT)
	b.Assert(ast.Val().IsString())
	return ast.Val().StringEx()
}
func ZendAstCreateZnode(node *Znode) *ZendAst {
	lineno := uint32(CG__().GetZendLineno())
	return NewAstZnode(lineno, node)
}
func ZendAstCreateZvalInt(zv *types.Zval, attr uint16, lineno uint32) *ZendAst {
	return NewAstZval(0, lineno, zv)
}
func ZendAstCreateZvalWithLineno(zv *types.Zval, lineno uint32) *ZendAst {
	return ZendAstCreateZvalInt(zv, 0, lineno)
}
func ZendAstCreateZval(zv *types.Zval) *ZendAst {
	return ZendAstCreateZvalInt(zv, 0, CG__().GetZendLineno())
}
func ZendAstCreateZvalFromStr(str string) *ZendAst {
	return ZendAstCreateZvalInt(types.NewZvalString(str), 0, CG__().GetZendLineno())
}
func ZendAstCreateConstant(name *types.String, attr ZendAstAttr) *ZendAst {
	zv := types.NewZvalString(name.GetStr())
	lineno := uint32(CG__().GetZendLineno())
	return NewAstConstant(attr, lineno, zv)
}

func AstCreateEx(kind ZendAstKind, attr ZendAstAttr, children ...*ZendAst) *ZendAst {
	lineno := uint32(CG__().GetZendLineno())
	for _, child := range children {
		if child != nil {
			lineno = child.Lineno()
			break
		}
	}
	return NewAst(kind, attr, lineno, children)
}
func AstCreate(kind ZendAstKind, children ...*ZendAst) *ZendAst {
	return AstCreateEx(kind, 0, children...)
}
func AstCreateList(kind ZendAstKind, children ...*ZendAst) *ZendAst {
	lineno := uint32(CG__().GetZendLineno())
	for _, child := range children {
		if child != nil {
			lineno = b.Min(lineno, child.Lineno())
			break
		}
	}
	return NewAstList(kind, 0, lineno, children)
}

func ZendAstAddArrayElement(result *types.Zval, offset *types.Zval, expr *types.Zval) int {
	switch offset.Type() {
	case types.IsUndef:
		if result.Array().Append(expr) == nil {
			faults.Error(faults.E_WARNING, "Cannot add element to the array as the next element is already occupied")
		}
	case types.IsString:
		result.Array().SymtableUpdate(offset.String(), expr)
	case types.IsNull:
		result.Array().SymtableUpdate(types.NewString("").GetStr(), expr)
	case types.IsLong:
		result.Array().IndexUpdate(offset.Long(), expr)
	case types.IsFalse:
		result.Array().IndexUpdate(0, expr)
	case types.IsTrue:
		result.Array().IndexUpdate(1, expr)
	case types.IsDouble:
		result.Array().IndexUpdate(operators.DvalToLval(offset.Double()), expr)
	case types.IsResource:
		faults.Error(faults.E_NOTICE, "Resource ID#%d used as offset, casting to integer (%d)", offset.ResourceHandle(), offset.ResourceHandle())
		result.Array().IndexUpdate(offset.ResourceHandle(), expr)
	default:
		faults.ThrowError(nil, "Illegal offset type")
		return types.FAILURE
	}
	return types.SUCCESS
}
func ZendAstAddUnpackedElement(result *types.Zval, expr *types.Zval) int {
	if expr.IsArray() {
		var ht *types.Array = expr.Array()
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
				if result.Array().Append(val) == nil {
					faults.Error(faults.E_WARNING, "Cannot add element to the array as the next element is already occupied")
					break
				}
				// val.TryAddRefcount()
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
	switch ast.Kind() {
	case ZEND_AST_BINARY_OP:
		if ZendAstEvaluate(&op1, ast.Child(0), scope) != types.SUCCESS {
			ret = types.FAILURE
		} else if ZendAstEvaluate(&op2, ast.Child(1), scope) != types.SUCCESS {
			ret = types.FAILURE
		} else {
			var op BinaryOpType = GetBinaryOp(OpCode(ast.Attr()))
			ret = op(result, &op1, &op2)
		}
	case ZEND_AST_GREATER:
		fallthrough
	case ZEND_AST_GREATER_EQUAL:
		if ZendAstEvaluate(&op1, ast.Child(0), scope) != types.SUCCESS {
			ret = types.FAILURE
		} else if ZendAstEvaluate(&op2, ast.Child(1), scope) != types.SUCCESS {
			// ZvalPtrDtorNogc(&op1)
			ret = types.FAILURE
		} else {

			/* op1 > op2 is the same as op2 < op1 */

			var op BinaryOpType = lang.Cond(ast.Kind() == ZEND_AST_GREATER, operators.IsSmallerFunction, operators.IsSmallerOrEqualFunction)
			ret = op(result, &op2, &op1)
			// ZvalPtrDtorNogc(&op1)
			// ZvalPtrDtorNogc(&op2)
		}
	case ZEND_AST_UNARY_OP:
		if ZendAstEvaluate(&op1, ast.Child(0), scope) != types.SUCCESS {
			ret = types.FAILURE
		} else {
			var op UnaryOpType = GetUnaryOp(ast.Attr())
			ret = op(result, &op1)
			// ZvalPtrDtorNogc(&op1)
		}
	case ZEND_AST_ZVAL:
		var zv *types.Zval = ast.Val()
		types.ZVAL_COPY(result, zv)
	case ZEND_AST_CONSTANT:
		var name *types.String = ZendAstGetConstantName(ast)
		var zv *types.Zval = ZendGetConstantEx(name.GetStr(), scope, ast.Attr())
		if zv == nil {
			result.SetUndef()
			ret = ZendUseUndefinedConstant(name, ast.Attr(), result)
			break
		}
		types.ZVAL_COPY_OR_DUP(result, zv)
	case ZEND_AST_CONSTANT_CLASS:
		if scope != nil {
			result.SetString(scope.Name())
		} else {
			result.SetString("")
		}
	case ZEND_AST_CLASS_NAME:
		if scope == nil {
			faults.ThrowError(nil, "Cannot use \"self\" when no class scope is active")
			return types.FAILURE
		}
		if ast.Attr() == ZEND_FETCH_CLASS_SELF {
			result.SetString(scope.Name())
		} else if ast.Attr() == ZEND_FETCH_CLASS_PARENT {
			if !(scope.GetParent()) {
				faults.ThrowError(nil, "Cannot use \"parent\" when current class scope has no parent")
				return types.FAILURE
			}
			result.SetString(scope.GetParent().name.GetStr())
		} else {
			b.Assert(false)
		}
	case ZEND_AST_AND:
		if ZendAstEvaluate(&op1, ast.Child(0), scope) != types.SUCCESS {
			ret = types.FAILURE
			break
		}
		if operators.ZvalIsTrue(&op1) {
			if ZendAstEvaluate(&op2, ast.Child(1), scope) != types.SUCCESS {
				// ZvalPtrDtorNogc(&op1)
				ret = types.FAILURE
				break
			}
			result.SetBool(operators.ZvalIsTrue(&op2))
			// ZvalPtrDtorNogc(&op2)
		} else {
			result.SetFalse()
		}
		// ZvalPtrDtorNogc(&op1)
	case ZEND_AST_OR:
		if ZendAstEvaluate(&op1, ast.Child(0), scope) != types.SUCCESS {
			ret = types.FAILURE
			break
		}
		if operators.ZvalIsTrue(&op1) {
			result.SetTrue()
		} else {
			if ZendAstEvaluate(&op2, ast.Child(1), scope) != types.SUCCESS {
				// ZvalPtrDtorNogc(&op1)
				ret = types.FAILURE
				break
			}
			result.SetBool(operators.ZvalIsTrue(&op2))
			// ZvalPtrDtorNogc(&op2)
		}
		// ZvalPtrDtorNogc(&op1)
	case ZEND_AST_CONDITIONAL:
		if ZendAstEvaluate(&op1, ast.Child(0), scope) != types.SUCCESS {
			ret = types.FAILURE
			break
		}
		if operators.ZvalIsTrue(&op1) {
			if ast.Child(1) == nil {
				*result = op1
			} else {
				if ZendAstEvaluate(result, ast.Child(1), scope) != types.SUCCESS {
					// ZvalPtrDtorNogc(&op1)
					ret = types.FAILURE
					break
				}
				// ZvalPtrDtorNogc(&op1)
			}
		} else {
			if ZendAstEvaluate(result, ast.Children()[2], scope) != types.SUCCESS {
				// ZvalPtrDtorNogc(&op1)
				ret = types.FAILURE
				break
			}
			// ZvalPtrDtorNogc(&op1)
		}
	case ZEND_AST_COALESCE:
		if ZendAstEvaluate(&op1, ast.Child(0), scope) != types.SUCCESS {
			ret = types.FAILURE
			break
		}
		if op1.Type() > types.IsNull {
			*result = op1
		} else {
			if ZendAstEvaluate(result, ast.Child(1), scope) != types.SUCCESS {
				// ZvalPtrDtorNogc(&op1)
				ret = types.FAILURE
				break
			}
			// ZvalPtrDtorNogc(&op1)
		}
	case ZEND_AST_UNARY_PLUS:
		if ZendAstEvaluate(&op2, ast.Child(0), scope) != types.SUCCESS {
			ret = types.FAILURE
		} else {
			op1.SetLong(0)
			ret = operators.AddFunction(result, &op1, &op2)
		}
	case ZEND_AST_UNARY_MINUS:
		if ZendAstEvaluate(&op2, ast.Child(0), scope) != types.SUCCESS {
			ret = types.FAILURE
		} else {
			op1.SetLong(0)
			ret = operators.SubFunction(result, &op1, &op2)
			// ZvalPtrDtorNogc(&op2)
		}
	case ZEND_AST_ARRAY:
		ArrayInit(result)
		for _, elem := range ast.Children() {
			if elem.Kind() == ZEND_AST_UNPACK {
				if ZendAstEvaluate(&op1, elem.Children()[0], scope) != types.SUCCESS {
					return types.FAILURE
				}
				if ZendAstAddUnpackedElement(result, &op1) != types.SUCCESS {
					return types.FAILURE
				}
				continue
			}
			if elem.Child(1) != nil {
				if ZendAstEvaluate(&op1, elem.Children()[1], scope) != types.SUCCESS {
					return types.FAILURE
				}
			} else {
				op1.SetUndef()
			}
			if ZendAstEvaluate(&op2, elem.Children()[0], scope) != types.SUCCESS {
				return types.FAILURE
			}
			if ZendAstAddArrayElement(result, &op1, &op2) != types.SUCCESS {
				return types.FAILURE
			}
		}
	case ZEND_AST_DIM:
		if ast.Child(1) == nil {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot use [] for reading")
		}
		if ZendAstEvaluate(&op1, ast.Child(0), scope) != types.SUCCESS {
			ret = types.FAILURE
		} else if ZendAstEvaluate(&op2, ast.Child(1), scope) != types.SUCCESS {
			ret = types.FAILURE
		} else {
			ZendFetchDimensionConst(result, &op1, &op2, lang.Cond((ast.Attr()&ZEND_DIM_IS) != 0, BP_VAR_IS, BP_VAR_R))
		}
	default:
		faults.ThrowError(nil, "Unsupported constant expression")
		ret = types.FAILURE
	}
	return ret
}
