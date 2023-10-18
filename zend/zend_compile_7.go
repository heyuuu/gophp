package zend

import (
	"fmt"
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
	"os"
)

func ZendAddImport(typ uint32, lookupName string, oldName string) bool {
	switch typ {
	case ZEND_SYMBOL_CLASS:
		return FC__().AddImport(lookupName, oldName)
	case ZEND_SYMBOL_FUNCTION:
		return FC__().AddImportFunction(lookupName, oldName)
	case ZEND_SYMBOL_CONST:
		return FC__().AddImportConst(lookupName, oldName)
	default:
		panic("unreachable")
	}
}

func ZendGetUseTypeStr(type_ uint32) string {
	switch type_ {
	case ZEND_SYMBOL_CLASS:
		return ""
	case ZEND_SYMBOL_FUNCTION:
		return " function"
	case ZEND_SYMBOL_CONST:
		return " const"
	default:
	}
	return " unknown"
}
func ZendCheckAlreadyInUse(typ uint32, oldName string, newName string, checkName string) {
	if ascii.StrCaseEquals(oldName, checkName) {
		return
	}
	faults.ErrorNoreturn(faults.E_COMPILE_ERROR, fmt.Sprintf("Cannot use%s %s as %s because the name is already in use", ZendGetUseTypeStr(typ), oldName, newName))
}
func (compiler *Compiler) CompileUse(ast *ZendAst) {
	var list *ZendAstList = ast.AsAstList()
	var i uint32
	var current_ns = FC__().CurrentNamespace()
	var type_ uint32 = ast.Attr()
	var case_sensitive bool = type_ == ZEND_SYMBOL_CONST
	for i = 0; i < list.GetChildren(); i++ {
		var use_ast *ZendAst = list.Children()[i]
		var old_name_ast *ZendAst = use_ast.Child(0)
		var new_name_ast *ZendAst = use_ast.Child(1)
		var old_name *types.String = ZendAstGetStr(old_name_ast)
		var new_name string
		var lookup_name string
		if new_name_ast != nil {
			new_name = ZendAstGetStr(new_name_ast).GetStr()
		} else {
			if unqualifiedName, ok := ZendGetUnqualifiedNameEx(old_name.GetStr()); ok {
				/* The form "use A\B" is equivalent to "use A\B as B" */
				new_name = unqualifiedName
			} else {
				new_name = old_name.GetStr()
				if current_ns == "" {
					if type_ == T_CLASS && new_name == "strict" {
						faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "You seem to be trying to use a different language...")
					}
					faults.Error(faults.E_WARNING, fmt.Sprintf("The use statement with non-compound name '%s' has no effect", new_name))
				}
			}
		}
		if case_sensitive {
			lookup_name = new_name
		} else {
			lookup_name = ascii.StrToLower(new_name)
		}
		if type_ == ZEND_SYMBOL_CLASS && ZendIsReservedClassName(new_name) {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, fmt.Sprintf("Cannot use %s as %s because '%s' is a special class name", old_name.GetStr(), new_name, new_name))
		}
		if current_ns != "" {
			nsName := ascii.StrToLower(current_ns) + "\\" + lookup_name
			if FC__().HaveSeenSymbol(nsName, type_) {
				ZendCheckAlreadyInUse(type_, old_name.GetStr(), new_name, nsName)
			}
		} else {
			if FC__().HaveSeenSymbol(lookup_name, type_) {
				ZendCheckAlreadyInUse(type_, old_name.GetStr(), new_name, lookup_name)
			}
		}
		if ZendAddImport(type_, lookup_name, old_name.GetStr()) {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, fmt.Sprintf("Cannot use%s %s as %s because the name is already in use", ZendGetUseTypeStr(type_), old_name.GetStr(), new_name))
		}
	}
}

func (compiler *Compiler) CompileGroupUse(ast *ZendAst) {
	var i uint32
	var ns *types.String = ZendAstGetStr(ast.Child(0))
	var list *ZendAstList = ast.Child(1).AsAstList()
	for i = 0; i < list.GetChildren(); i++ {
		var inline_use *ZendAst
		var use *ZendAst = list.Children()[i]
		var name_zval *types.Zval = ZendAstGetZval(use.Children()[0])
		var name *types.String = name_zval.StringEx()
		var compound_ns string = ZendConcatNames(ns.GetStr(), name.GetStr())
		// types.ZendStringReleaseEx(name, 0)
		name_zval.SetString(compound_ns)
		inline_use = AstCreateList(ZEND_AST_USE, use)
		if ast.Attr() != 0 {
			inline_use.SetAttr(ast.Attr())
		} else {
			inline_use.SetAttr(use.Attr())
		}
		compiler.CompileUse(inline_use)
	}
}
func (compiler *Compiler) CompileConstDecl(ast *ZendAst) {
	var list *ZendAstList = ast.AsAstList()
	var i uint32
	for i = 0; i < list.GetChildren(); i++ {
		var const_ast *ZendAst = list.Children()[i]
		var name_ast *ZendAst = const_ast.Child(0)
		var value_ast *ZendAst = const_ast.Child(1)
		var unqualified_name *types.String = ZendAstGetStr(name_ast)
		var name *types.String
		var name_node Znode
		var value_node Znode
		var value_zv *types.Zval = value_node.GetConstant()
		value_node.SetOpType(IS_CONST)
		compiler.ConstExprToZval(value_zv, value_ast)
		if ZendLookupReservedConst(unqualified_name.GetStr()) != nil {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, fmt.Sprintf("Cannot redeclare constant '%s'", unqualified_name.GetStr()))
		}
		name = ZendPrefixWithNs(unqualified_name)
		//name = types.ZendNewInternedString(name)
		if importName := FC__().FindImportConst(unqualified_name.GetStr()); importName != "" && importName != name.GetStr() {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, fmt.Sprintf("Cannot declare const %s because the name is already in use", name.GetStr()))
		}
		name_node.SetOpType(IS_CONST)
		name_node.GetConstant().SetStringEx(name)
		ZendEmitOp(nil, ZEND_DECLARE_CONST, &name_node, &value_node)
		FC__().RegisterSeenSymbol(name.GetStr(), ZEND_SYMBOL_CONST)
	}
}
func (compiler *Compiler) CompileNamespace(ast *ZendAst) {
	var nameAst *ZendAst = ast.Child(0)
	var stmtAst *ZendAst = ast.Child(1)
	var name string
	var withBracket bool = stmtAst != nil

	/* handle mixed syntax declaration or nested namespaces */

	if !FC__().HasBracketedNamespaces() {
		if FC__().CurrentNamespace() != "" {
			if withBracket {
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot mix bracketed namespace declarations with unbracketed namespace declarations")
			}
		}
	} else {
		/* previous namespace declarations were bracketed */
		if !withBracket {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot mix bracketed namespace declarations with unbracketed namespace declarations")
		} else if FC__().CurrentNamespace() != "" || FC__().InNamespace() {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Namespace declarations cannot be nested")
		}
	}
	if (!withBracket && FC__().CurrentNamespace() == "" || withBracket && !FC__().HasBracketedNamespaces()) && CG__().GetActiveOpArray().GetLast() > 0 {

		/* ignore ZEND_EXT_STMT */

		var num uint32 = CG__().GetActiveOpArray().GetLast()
		for num > 0 && CG__().GetActiveOpArray().GetOpcodes()[num-1].GetOpcode() == ZEND_EXT_STMT {
			num--
		}
		if num > 0 {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Namespace declaration statement has to be the very first statement or after any declare call in the script")
		}
	}
	if nameAst != nil {
		name = ZendAstGetStr(nameAst).GetStr()
		if ZEND_FETCH_CLASS_DEFAULT != ZendGetClassFetchType(name) {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, fmt.Sprintf("Cannot use '%s' as namespace name", name))
		}
	}
	FC__().BeginNamespace(name, withBracket)
	if stmtAst != nil {
		compiler.CompileTopStmt(stmtAst)
		FC__().EndNamespace()
	}
}
func (compiler *Compiler) CompileHaltCompiler(ast *ZendAst) {
	var offsetAst *ZendAst = ast.Child(0)
	var offset ZendLong = offsetAst.Val().Long()
	var constName = "__COMPILER_HALT_OFFSET__"
	if FC__().HasBracketedNamespaces() != 0 && FC__().InNamespace() != 0 {
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "__HALT_COMPILER() can only be used from the outermost scope")
	}
	filename := ZendGetCompiledFilename()
	name := ZendManglePropertyName_Ex(constName, filename)
	RegisterLongConstant(name, offset, CONST_CS, 0)
}
func ZendTryCtEvalMagicConst(zv *types.Zval, ast *ZendAst) bool {
	var op_array *types.ZendOpArray = CG__().GetActiveOpArray()
	var ce *types.ClassEntry = CG__().GetActiveClassEntry()
	switch ast.Attr() {
	case T_LINE:
		zv.SetLong(ast.Lineno())
	case T_FILE:
		zv.SetString(CG__().GetCompiledFilename())
	case T_DIR:
		var filename = CG__().GetCompiledFilename()
		var dirname = ZendDirname(filename)
		if dirname == "." {
			dirname, _ = os.Getwd()
		}
		zv.SetString(dirname)
	case T_FUNC_C:
		if op_array != nil && op_array.FunctionName() != "" {
			zv.SetString(op_array.FunctionName())
		} else {
			zv.SetString("")
		}
	case T_METHOD_C:

		/* Detect whether we are directly inside a class (e.g. a class constant) and treat
		 * this as not being inside a function. */

		if op_array != nil && ce != nil && op_array.GetScope() == nil && !op_array.IsClosure() {
			op_array = nil
		}
		if op_array != nil && op_array.FunctionName() != "" {
			if op_array.GetScope() != nil {
				zv.SetString(op_array.GetScope().Name() + "::" + op_array.FunctionName())
			} else {
				zv.SetString(op_array.FunctionName())
			}
		} else {
			zv.SetString("")
		}
	case T_CLASS_C:
		if ce != nil {
			if ce.IsTrait() {
				return 0
			} else {
				zv.SetString(ce.Name())
			}
		} else {
			zv.SetString("")
		}
	case T_TRAIT_C:
		if ce != nil && ce.IsTrait() {
			zv.SetString(ce.Name())
		} else {
			zv.SetString("")
		}
	case T_NS_C:
		zv.SetString(FC__().CurrentNamespace())
	default:

	}
	return 1
}
func ZendBinaryOpProducesNumericStringError(opcode uint32, op1 *types.Zval, op2 *types.Zval) bool {
	if !(opcode == ZEND_ADD || opcode == ZEND_SUB || opcode == ZEND_MUL || opcode == ZEND_DIV || opcode == ZEND_POW || opcode == ZEND_MOD || opcode == ZEND_SL || opcode == ZEND_SR || opcode == ZEND_BW_OR || opcode == ZEND_BW_AND || opcode == ZEND_BW_XOR) {
		return 0
	}

	/* While basic arithmetic operators always produce numeric string errors,
	 * bitwise operators don't produce errors if both operands are strings */

	if (opcode == ZEND_BW_OR || opcode == ZEND_BW_AND || opcode == ZEND_BW_XOR) && op1.IsString() && op2.IsString() {
		return 0
	}
	if op1.IsString() && operators.IsNumericString(op1.String(), nil, nil, 0) == 0 {
		return 1
	}
	if op2.IsString() && operators.IsNumericString(op2.String(), nil, nil, 0) == 0 {
		return 1
	}
	return 0
}
func ZendBinaryOpProducesArrayConversionError(opcode uint32, op1 *types.Zval, op2 *types.Zval) bool {
	if opcode == ZEND_CONCAT && (op1.IsArray() || op2.IsArray()) {
		return 1
	}
	return 0
}
func ZendTryCtEvalBinaryOp(result *types.Zval, opcode uint32, op1 *types.Zval, op2 *types.Zval) bool {
	var fn BinaryOpType = GetBinaryOp(opcode)

	/* don't evaluate division by zero at compile-time */

	if (opcode == ZEND_DIV || opcode == ZEND_MOD) && operators.ZvalGetLong(op2) == 0 {
		return 0
	} else if (opcode == ZEND_SL || opcode == ZEND_SR) && operators.ZvalGetLong(op2) < 0 {
		return 0
	}

	/* don't evaluate numeric string error-producing operations at compile-time */

	if ZendBinaryOpProducesNumericStringError(opcode, op1, op2) != 0 {
		return 0
	}

	/* don't evaluate array to string conversions at compile-time */

	if ZendBinaryOpProducesArrayConversionError(opcode, op1, op2) != 0 {
		return 0
	}
	fn(result, op1, op2)
	return 1
}
func ZendCtEvalUnaryOp(result *types.Zval, opcode uint32, op *types.Zval) {
	var fn UnaryOpType = GetUnaryOp(opcode)
	fn(result, op)
}
func ZendTryCtEvalUnaryPm(result *types.Zval, kind ZendAstKind, op *types.Zval) bool {
	var left types.Zval
	left.SetLong(lang.Cond(kind == ZEND_AST_UNARY_PLUS, 1, -1))
	return ZendTryCtEvalBinaryOp(result, ZEND_MUL, &left, op)
}
func ZendCtEvalGreater(result *types.Zval, kind ZendAstKind, op1 *types.Zval, op2 *types.Zval) {
	var fn BinaryOpType = lang.Cond(kind == ZEND_AST_GREATER, operators.IsSmallerFunction, operators.IsSmallerOrEqualFunction)
	fn(result, op2, op1)
}

func (compiler *Compiler) TryCtEvalArray(result *types.Zval, ast *ZendAst) bool {
	var list *ZendAstList = ast.AsAstList()
	var last_elem_ast *ZendAst = nil
	var i uint32
	var is_constant bool = 1
	if ast.Attr() == ZEND_ARRAY_SYNTAX_LIST {
		faults.Error(faults.E_COMPILE_ERROR, "Cannot use list() as standalone expression")
	}

	/* First ensure that *all* child nodes are constant and by-val */

	for i = 0; i < list.GetChildren(); i++ {
		var elem_ast *ZendAst = list.Children()[i]
		if elem_ast == nil {

			/* Report error at line of last non-empty element */

			if last_elem_ast != nil {
				compiler.setLinenoByAst(last_elem_ast)
			}
			faults.Error(faults.E_COMPILE_ERROR, "Cannot use empty array elements in arrays")
		}
		if elem_ast.Kind() != ZEND_AST_UNPACK {
			compiler.EvalConstExpr(elem_ast.Child(0))
			compiler.EvalConstExpr(elem_ast.Child(1))
			if elem_ast.Attr() != 0 || elem_ast.Child(0).Kind() != ZEND_AST_ZVAL || elem_ast.Child(1) != nil && elem_ast.Child(1).Kind() != ZEND_AST_ZVAL {
				is_constant = 0
			}
		} else {
			compiler.EvalConstExpr(elem_ast.Child(0))
			if elem_ast.Child(0).Kind() != ZEND_AST_ZVAL {
				is_constant = 0
			}
		}
		last_elem_ast = elem_ast
	}
	if is_constant == 0 {
		return 0
	}
	if list.GetChildren() == 0 {
		result.SetEmptyArray()
		return 1
	}
	ArrayInitSize(result, list.GetChildren())
	for i = 0; i < list.GetChildren(); i++ {
		var elem_ast *ZendAst = list.Children()[i]
		var value_ast *ZendAst = elem_ast.Child(0)
		var key_ast *ZendAst
		var value *types.Zval = value_ast.Val()
		if elem_ast.Kind() == ZEND_AST_UNPACK {
			if value.IsArray() {
				var ht *types.Array = value.Array()
				ok := ht.ForeachEx(func(key types.ArrayKey, value *types.Zval) bool {
					if key.IsStrKey() {
						faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot unpack array with string keys")
					}
					if result.Array().Append(value) == nil {
						return false
					}
					return true
				})
				if !ok {
					return 0
				}

				continue
			} else {
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Only arrays and Traversables can be unpacked")
			}
		}
		// value.TryAddRefcount()
		key_ast = elem_ast.Child(1)
		if key_ast != nil {
			var key *types.Zval = key_ast.Val()
			switch key.Type() {
			case types.IsLong:
				result.Array().IndexUpdate(key.Long(), value)
			case types.IsString:
				result.Array().SymtableUpdate(key.String(), value)
			case types.IsDouble:
				result.Array().IndexUpdate(operators.DvalToLval(key.Double()), value)
			case types.IsFalse:
				result.Array().IndexUpdate(0, value)
			case types.IsTrue:
				result.Array().IndexUpdate(1, value)
			case types.IsNull:
				result.Array().KeyUpdate(types.NewString("").GetStr(), value)
			default:
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Illegal offset type")
			}
		} else {
			if result.Array().Append(value) == nil {
				// ZvalPtrDtorNogc(value)
				// ZvalPtrDtor(result)
				return 0
			}
		}
	}
	return 1
}
func (compiler *Compiler) CompileBinaryOp(result *Znode, ast *ZendAst) {
	var left_ast *ZendAst = ast.Child(0)
	var right_ast *ZendAst = ast.Child(1)
	var opcode uint32 = ast.Attr()
	if (opcode == ZEND_ADD || opcode == ZEND_SUB) && left_ast.Kind() == ZEND_AST_BINARY_OP && left_ast.Attr() == ZEND_CONCAT {
		faults.Error(faults.E_DEPRECATED, "The behavior of unparenthesized expressions containing both '.' and '+'/'-' will change in PHP 8: '+'/'-' will take a higher precedence")
	}
	if (opcode == ZEND_SL || opcode == ZEND_SR) && (left_ast.Kind() == ZEND_AST_BINARY_OP && left_ast.Attr() == ZEND_CONCAT || right_ast.Kind() == ZEND_AST_BINARY_OP && right_ast.Attr() == ZEND_CONCAT) {
		faults.Error(faults.E_DEPRECATED, "The behavior of unparenthesized expressions containing both '.' and '>>'/'<<' will change in PHP 8: '<<'/'>>' will take a higher precedence")
	}
	if opcode == ZEND_PARENTHESIZED_CONCAT {
		opcode = ZEND_CONCAT
	}
	var left_node Znode
	var right_node Znode
	compiler.CompileExpr(&left_node, left_ast)
	compiler.CompileExpr(&right_node, right_ast)
	if left_node.GetOpType() == IS_CONST && right_node.GetOpType() == IS_CONST {
		if ZendTryCtEvalBinaryOp(result.GetConstant(), opcode, left_node.GetConstant(), right_node.GetConstant()) != 0 {
			result.SetOpType(IS_CONST)
			// ZvalPtrDtor(left_node.GetConstant())
			// ZvalPtrDtor(right_node.GetConstant())
			return
		}
	}
	for {
		if opcode == ZEND_IS_EQUAL || opcode == ZEND_IS_NOT_EQUAL {
			if left_node.GetOpType() == IS_CONST {
				if left_node.GetConstant().IsFalse() {
					if opcode == ZEND_IS_NOT_EQUAL {
						opcode = ZEND_BOOL
					} else {
						opcode = ZEND_BOOL_NOT
					}
					ZendEmitOpTmp(result, opcode, &right_node, nil)
					break
				} else if left_node.GetConstant().IsTrue() {
					if opcode == ZEND_IS_EQUAL {
						opcode = ZEND_BOOL
					} else {
						opcode = ZEND_BOOL_NOT
					}
					ZendEmitOpTmp(result, opcode, &right_node, nil)
					break
				}
			} else if right_node.GetOpType() == IS_CONST {
				if right_node.GetConstant().IsFalse() {
					if opcode == ZEND_IS_NOT_EQUAL {
						opcode = ZEND_BOOL
					} else {
						opcode = ZEND_BOOL_NOT
					}
					ZendEmitOpTmp(result, opcode, &left_node, nil)
					break
				} else if right_node.GetConstant().IsTrue() {
					if opcode == ZEND_IS_EQUAL {
						opcode = ZEND_BOOL
					} else {
						opcode = ZEND_BOOL_NOT
					}
					ZendEmitOpTmp(result, opcode, &left_node, nil)
					break
				}
			}
		}
		if opcode == ZEND_CONCAT {

			/* convert constant operands to strings at compile-time */

			if left_node.GetOpType() == IS_CONST {
				if left_node.GetConstant().IsArray() {
					ZendEmitOpTmp(&left_node, ZEND_CAST, &left_node, nil).SetExtendedValue(types.IsString)
				} else {
					operators.ConvertToString(left_node.GetConstant())
				}
			}
			if right_node.GetOpType() == IS_CONST {
				if right_node.GetConstant().IsArray() {
					ZendEmitOpTmp(&right_node, ZEND_CAST, &right_node, nil).SetExtendedValue(types.IsString)
				} else {
					operators.ConvertToString(right_node.GetConstant())
				}
			}
			if left_node.GetOpType() == IS_CONST && right_node.GetOpType() == IS_CONST {
				opcode = ZEND_FAST_CONCAT
			}
		}
		ZendEmitOpTmp(result, opcode, &left_node, &right_node)
		break
	}
}
func (compiler *Compiler) CompileGreater(result *Znode, ast *ZendAst) {
	var left_ast *ZendAst = ast.Child(0)
	var right_ast *ZendAst = ast.Child(1)
	var left_node Znode
	var right_node Znode
	b.Assert(ast.Kind() == ZEND_AST_GREATER || ast.Kind() == ZEND_AST_GREATER_EQUAL)
	compiler.CompileExpr(&left_node, left_ast)
	compiler.CompileExpr(&right_node, right_ast)
	if left_node.GetOpType() == IS_CONST && right_node.GetOpType() == IS_CONST {
		result.SetOpType(IS_CONST)
		ZendCtEvalGreater(result.GetConstant(), ast.Kind(), left_node.GetConstant(), right_node.GetConstant())
		// ZvalPtrDtor(left_node.GetConstant())
		// ZvalPtrDtor(right_node.GetConstant())
		return
	}
	ZendEmitOpTmp(result, lang.Cond(ast.Kind() == ZEND_AST_GREATER, ZEND_IS_SMALLER, ZEND_IS_SMALLER_OR_EQUAL), &right_node, &left_node)
}
func (compiler *Compiler) CompileUnaryOp(result *Znode, ast *ZendAst) {
	var expr_ast *ZendAst = ast.Child(0)
	var opcode uint32 = ast.Attr()
	var expr_node Znode
	compiler.CompileExpr(&expr_node, expr_ast)
	if expr_node.GetOpType() == IS_CONST {
		result.SetOpType(IS_CONST)
		ZendCtEvalUnaryOp(result.GetConstant(), opcode, expr_node.GetConstant())
		// ZvalPtrDtor(expr_node.GetConstant())
		return
	}
	ZendEmitOpTmp(result, opcode, &expr_node, nil)
}
