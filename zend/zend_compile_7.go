// <<generate>>

package zend

import (
	b "sik/builtin"
	"sik/zend/faults"
	"sik/zend/types"
)

func ZendGetImportHt(type_ uint32) *types.Array {
	switch type_ {
	case ZEND_SYMBOL_CLASS:
		if FC__().GetImports() == nil {
			FC__().SetImports(Emalloc(b.SizeOf("HashTable")))
			FC__().GetImports() = types.MakeArrayEx(8, StrDtor, 0)
		}
		return FC__().GetImports()
	case ZEND_SYMBOL_FUNCTION:
		if FC__().GetImportsFunction() == nil {
			FC__().SetImportsFunction(Emalloc(b.SizeOf("HashTable")))
			FC__().GetImportsFunction() = types.MakeArrayEx(8, StrDtor, 0)
		}
		return FC__().GetImportsFunction()
	case ZEND_SYMBOL_CONST:
		if FC__().GetImportsConst() == nil {
			FC__().SetImportsConst(Emalloc(b.SizeOf("HashTable")))
			FC__().GetImportsConst() = types.MakeArrayEx(8, StrDtor, 0)
		}
		return FC__().GetImportsConst()
	default:

	}
	return nil
}
func ZendGetUseTypeStr(type_ uint32) *byte {
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
func ZendCheckAlreadyInUse(type_ uint32, old_name *types.String, new_name *types.String, check_name *types.String) {
	if types.ZendStringEqualsCi(old_name, check_name) {
		return
	}
	faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot use%s %s as %s because the name "+"is already in use", ZendGetUseTypeStr(type_), old_name.GetVal(), new_name.GetVal())
}
func ZendCompileUse(ast *ZendAst) {
	var list *ZendAstList = ZendAstGetList(ast)
	var i uint32
	var current_ns *types.String = FC__().GetCurrentNamespace()
	var type_ uint32 = ast.GetAttr()
	var current_import *types.Array = ZendGetImportHt(type_)
	var case_sensitive types.ZendBool = type_ == ZEND_SYMBOL_CONST
	for i = 0; i < list.GetChildren(); i++ {
		var use_ast *ZendAst = list.GetChild()[i]
		var old_name_ast *ZendAst = use_ast.GetChild()[0]
		var new_name_ast *ZendAst = use_ast.GetChild()[1]
		var old_name *types.String = ZendAstGetStr(old_name_ast)
		var new_name *types.String
		var lookup_name *types.String
		if new_name_ast != nil {
			new_name = ZendAstGetStr(new_name_ast).Copy()
		} else {
			var unqualified_name *byte
			var unqualified_name_len int
			if ZendGetUnqualifiedName(old_name, &unqualified_name, &unqualified_name_len) != 0 {

				/* The form "use A\B" is equivalent to "use A\B as B" */

				new_name = types.NewString(b.CastStr(unqualified_name, unqualified_name_len))

				/* The form "use A\B" is equivalent to "use A\B as B" */

			} else {
				new_name = old_name.Copy()
				if current_ns == nil {
					if type_ == T_CLASS && types.ZendStringEqualsLiteral(new_name, "strict") {
						faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "You seem to be trying to use a different language...")
					}
					faults.Error(faults.E_WARNING, "The use statement with non-compound name '%s' "+"has no effect", new_name.GetVal())
				}
			}
		}
		if case_sensitive != 0 {
			lookup_name = new_name.Copy()
		} else {
			lookup_name = ZendStringTolower(new_name)
		}
		if type_ == ZEND_SYMBOL_CLASS && ZendIsReservedClassName(new_name) != 0 {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot use %s as %s because '%s' "+"is a special class name", old_name.GetVal(), new_name.GetVal(), new_name.GetVal())
		}
		if current_ns != nil {
			var ns_name *types.String = types.ZendStringAlloc(current_ns.GetLen()+1+new_name.GetLen(), 0)
			ZendStrTolowerCopy(ns_name.GetVal(), current_ns.GetVal(), current_ns.GetLen())
			ns_name.GetVal()[current_ns.GetLen()] = '\\'
			memcpy(ns_name.GetVal()+current_ns.GetLen()+1, lookup_name.GetVal(), lookup_name.GetLen()+1)
			if ZendHaveSeenSymbol(ns_name, type_) != 0 {
				ZendCheckAlreadyInUse(type_, old_name, new_name, ns_name)
			}
			types.ZendStringEfree(ns_name)
		} else {
			if ZendHaveSeenSymbol(lookup_name, type_) != 0 {
				ZendCheckAlreadyInUse(type_, old_name, new_name, lookup_name)
			}
		}
		old_name.AddRefcount()
		old_name = types.ZendNewInternedString(old_name)
		if !(types.ZendHashAddPtr(current_import, lookup_name.GetStr(), old_name)) {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot use%s %s as %s because the name "+"is already in use", ZendGetUseTypeStr(type_), old_name.GetVal(), new_name.GetVal())
		}
		types.ZendStringReleaseEx(lookup_name, 0)
		types.ZendStringReleaseEx(new_name, 0)
	}
}
func ZendCompileGroupUse(ast *ZendAst) {
	var i uint32
	var ns *types.String = ZendAstGetStr(ast.GetChild()[0])
	var list *ZendAstList = ZendAstGetList(ast.GetChild()[1])
	for i = 0; i < list.GetChildren(); i++ {
		var inline_use *ZendAst
		var use *ZendAst = list.GetChild()[i]
		var name_zval *types.Zval = ZendAstGetZval(use.GetChild()[0])
		var name *types.String = name_zval.GetStr()
		var compound_ns *types.String = ZendConcatNames(ns.GetVal(), ns.GetLen(), name.GetVal(), name.GetLen())
		types.ZendStringReleaseEx(name, 0)
		name_zval.SetString(compound_ns)
		inline_use = ZendAstCreateList(1, ZEND_AST_USE, use)
		if ast.GetAttr() != 0 {
			inline_use.SetAttr(ast.GetAttr())
		} else {
			inline_use.SetAttr(use.GetAttr())
		}
		ZendCompileUse(inline_use)
	}
}
func ZendCompileConstDecl(ast *ZendAst) {
	var list *ZendAstList = ZendAstGetList(ast)
	var i uint32
	for i = 0; i < list.GetChildren(); i++ {
		var const_ast *ZendAst = list.GetChild()[i]
		var name_ast *ZendAst = const_ast.GetChild()[0]
		var value_ast *ZendAst = const_ast.GetChild()[1]
		var unqualified_name *types.String = ZendAstGetStr(name_ast)
		var name *types.String
		var name_node Znode
		var value_node Znode
		var value_zv *types.Zval = value_node.GetConstant()
		value_node.SetOpType(IS_CONST)
		ZendConstExprToZval(value_zv, value_ast)
		if ZendLookupReservedConst(unqualified_name.GetVal(), unqualified_name.GetLen()) != nil {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot redeclare constant '%s'", unqualified_name.GetVal())
		}
		name = ZendPrefixWithNs(unqualified_name)
		name = types.ZendNewInternedString(name)
		if FC__().GetImportsConst() != nil {
			var import_name *types.String = types.ZendHashFindPtr(FC__().GetImportsConst(), unqualified_name.GetStr())
			if import_name != nil && types.ZendStringEquals(import_name, name) == 0 {
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot declare const %s because "+"the name is already in use", name.GetVal())
			}
		}
		name_node.SetOpType(IS_CONST)
		name_node.GetConstant().SetString(name)
		ZendEmitOp(nil, ZEND_DECLARE_CONST, &name_node, &value_node)
		ZendRegisterSeenSymbol(name, ZEND_SYMBOL_CONST)
	}
}
func ZendCompileNamespace(ast *ZendAst) {
	var name_ast *ZendAst = ast.GetChild()[0]
	var stmt_ast *ZendAst = ast.GetChild()[1]
	var name *types.String
	var with_bracket types.ZendBool = stmt_ast != nil

	/* handle mixed syntax declaration or nested namespaces */

	if FC__().GetHasBracketedNamespaces() == 0 {
		if FC__().GetCurrentNamespace() != nil {

			/* previous namespace declarations were unbracketed */

			if with_bracket != 0 {
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot mix bracketed namespace declarations "+"with unbracketed namespace declarations")
			}

			/* previous namespace declarations were unbracketed */

		}
	} else {

		/* previous namespace declarations were bracketed */

		if with_bracket == 0 {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot mix bracketed namespace declarations "+"with unbracketed namespace declarations")
		} else if FC__().GetCurrentNamespace() != nil || FC__().GetInNamespace() != 0 {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Namespace declarations cannot be nested")
		}

		/* previous namespace declarations were bracketed */

	}
	if (with_bracket == 0 && FC__().GetCurrentNamespace() == nil || with_bracket != 0 && FC__().GetHasBracketedNamespaces() == 0) && CG__().GetActiveOpArray().GetLast() > 0 {

		/* ignore ZEND_EXT_STMT and ZEND_TICKS */

		var num uint32 = CG__().GetActiveOpArray().GetLast()
		for num > 0 && (CG__().GetActiveOpArray().GetOpcodes()[num-1].GetOpcode() == ZEND_EXT_STMT || CG__().GetActiveOpArray().GetOpcodes()[num-1].GetOpcode() == ZEND_TICKS) {
			num--
		}
		if num > 0 {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Namespace declaration statement has to be "+"the very first statement or after any declare call in the script")
		}
	}
	if FC__().GetCurrentNamespace() != nil {
		types.ZendStringReleaseEx(FC__().GetCurrentNamespace(), 0)
	}
	if name_ast != nil {
		name = ZendAstGetStr(name_ast)
		if ZEND_FETCH_CLASS_DEFAULT != ZendGetClassFetchType(name) {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot use '%s' as namespace name", name.GetVal())
		}
		FC__().SetCurrentNamespace(name.Copy())
	} else {
		FC__().SetCurrentNamespace(nil)
	}
	ZendResetImportTables()
	FC__().SetInNamespace(1)
	if with_bracket != 0 {
		FC__().SetHasBracketedNamespaces(1)
	}
	if stmt_ast != nil {
		ZendCompileTopStmt(stmt_ast)
		ZendEndNamespace()
	}
}
func ZendCompileHaltCompiler(ast *ZendAst) {
	var offset_ast *ZendAst = ast.GetChild()[0]
	var offset ZendLong = ZendAstGetZval(offset_ast).GetLval()
	var filename *types.String
	var name *types.String
	var const_name = "__COMPILER_HALT_OFFSET__"
	if FC__().GetHasBracketedNamespaces() != 0 && FC__().GetInNamespace() != 0 {
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "__HALT_COMPILER() can only be used from the outermost scope")
	}
	filename = ZendGetCompiledFilename()
	name = ZendManglePropertyName_ZStr(const_name, filename.GetStr())
	ZendRegisterLongConstant(name.GetVal(), name.GetLen(), offset, CONST_CS, 0)
	types.ZendStringReleaseEx(name, 0)
}
func ZendTryCtEvalMagicConst(zv *types.Zval, ast *ZendAst) types.ZendBool {
	var op_array *ZendOpArray = CG__().GetActiveOpArray()
	var ce *types.ClassEntry = CG__().GetActiveClassEntry()
	switch ast.GetAttr() {
	case T_LINE:
		zv.SetLong(ast.GetLineno())
	case T_FILE:
		zv.SetStringCopy(CG__().GetCompiledFilename())
	case T_DIR:
		var filename *types.String = CG__().GetCompiledFilename()
		var dirname *types.String = types.NewString(filename.GetStr())
		dirname.SetLen(ZendDirname(dirname.GetVal(), dirname.GetLen()))
		if strcmp(dirname.GetVal(), ".") == 0 {
			dirname = types.ZendStringExtend(dirname, MAXPATHLEN, 0)
			ZEND_IGNORE_VALUE(VCWD_GETCWD(dirname.GetVal(), MAXPATHLEN))
			dirname.SetLen(strlen(dirname.GetVal()))
		}
		zv.SetString(dirname)
	case T_FUNC_C:
		if op_array != nil && op_array.GetFunctionName() != nil {
			zv.SetStringCopy(op_array.GetFunctionName())
		} else {
			ZVAL_EMPTY_STRING(zv)
		}
	case T_METHOD_C:

		/* Detect whether we are directly inside a class (e.g. a class constant) and treat
		 * this as not being inside a function. */

		if op_array != nil && ce != nil && op_array.GetScope() == nil && !op_array.IsClosure() {
			op_array = nil
		}
		if op_array != nil && op_array.GetFunctionName() != nil {
			if op_array.GetScope() != nil {
				zv.SetString(ZendConcat3(op_array.GetScope().GetName().GetVal(), op_array.GetScope().GetName().GetLen(), "::", 2, op_array.GetFunctionName().GetVal(), op_array.GetFunctionName().GetLen()))
			} else {
				zv.SetStringCopy(op_array.GetFunctionName())
			}
		} else {
			ZVAL_EMPTY_STRING(zv)
		}
	case T_CLASS_C:
		if ce != nil {
			if ce.IsTrait() {
				return 0
			} else {
				zv.SetStringCopy(ce.GetName())
			}
		} else {
			ZVAL_EMPTY_STRING(zv)
		}
	case T_TRAIT_C:
		if ce != nil && ce.IsTrait() {
			zv.SetStringCopy(ce.GetName())
		} else {
			ZVAL_EMPTY_STRING(zv)
		}
	case T_NS_C:
		if FC__().GetCurrentNamespace() != nil {
			zv.SetStringCopy(FC__().GetCurrentNamespace())
		} else {
			ZVAL_EMPTY_STRING(zv)
		}
	default:

	}
	return 1
}
func ZendBinaryOpProducesNumericStringError(opcode uint32, op1 *types.Zval, op2 *types.Zval) types.ZendBool {
	if !(opcode == ZEND_ADD || opcode == ZEND_SUB || opcode == ZEND_MUL || opcode == ZEND_DIV || opcode == ZEND_POW || opcode == ZEND_MOD || opcode == ZEND_SL || opcode == ZEND_SR || opcode == ZEND_BW_OR || opcode == ZEND_BW_AND || opcode == ZEND_BW_XOR) {
		return 0
	}

	/* While basic arithmetic operators always produce numeric string errors,
	 * bitwise operators don't produce errors if both operands are strings */

	if (opcode == ZEND_BW_OR || opcode == ZEND_BW_AND || opcode == ZEND_BW_XOR) && op1.IsString() && op2.IsString() {
		return 0
	}
	if op1.IsString() && IsNumericString(op1.GetStr().GetStr(), nil, nil, 0) == 0 {
		return 1
	}
	if op2.IsString() && IsNumericString(op2.GetStr().GetStr(), nil, nil, 0) == 0 {
		return 1
	}
	return 0
}
func ZendBinaryOpProducesArrayConversionError(opcode uint32, op1 *types.Zval, op2 *types.Zval) types.ZendBool {
	if opcode == ZEND_CONCAT && (op1.IsArray() || op2.IsArray()) {
		return 1
	}
	return 0
}
func ZendTryCtEvalBinaryOp(result *types.Zval, opcode uint32, op1 *types.Zval, op2 *types.Zval) types.ZendBool {
	var fn BinaryOpType = GetBinaryOp(opcode)

	/* don't evaluate division by zero at compile-time */

	if (opcode == ZEND_DIV || opcode == ZEND_MOD) && ZvalGetLong(op2) == 0 {
		return 0
	} else if (opcode == ZEND_SL || opcode == ZEND_SR) && ZvalGetLong(op2) < 0 {
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
func ZendTryCtEvalUnaryPm(result *types.Zval, kind ZendAstKind, op *types.Zval) types.ZendBool {
	var left types.Zval
	left.SetLong(b.Cond(kind == ZEND_AST_UNARY_PLUS, 1, -1))
	return ZendTryCtEvalBinaryOp(result, ZEND_MUL, &left, op)
}
func ZendCtEvalGreater(result *types.Zval, kind ZendAstKind, op1 *types.Zval, op2 *types.Zval) {
	var fn BinaryOpType = b.Cond(kind == ZEND_AST_GREATER, IsSmallerFunction, IsSmallerOrEqualFunction)
	fn(result, op2, op1)
}
func ZendTryCtEvalArray(result *types.Zval, ast *ZendAst) types.ZendBool {
	var list *ZendAstList = ZendAstGetList(ast)
	var last_elem_ast *ZendAst = nil
	var i uint32
	var is_constant types.ZendBool = 1
	if ast.GetAttr() == ZEND_ARRAY_SYNTAX_LIST {
		faults.Error(faults.E_COMPILE_ERROR, "Cannot use list() as standalone expression")
	}

	/* First ensure that *all* child nodes are constant and by-val */

	for i = 0; i < list.GetChildren(); i++ {
		var elem_ast *ZendAst = list.GetChild()[i]
		if elem_ast == nil {

			/* Report error at line of last non-empty element */

			if last_elem_ast != nil {
				CG__().SetZendLineno(ZendAstGetLineno(last_elem_ast))
			}
			faults.Error(faults.E_COMPILE_ERROR, "Cannot use empty array elements in arrays")
		}
		if elem_ast.GetKind() != ZEND_AST_UNPACK {
			ZendEvalConstExpr(elem_ast.GetChild()[0])
			ZendEvalConstExpr(elem_ast.GetChild()[1])
			if elem_ast.GetAttr() != 0 || elem_ast.GetChild()[0].GetKind() != ZEND_AST_ZVAL || elem_ast.GetChild()[1] != nil && elem_ast.GetChild()[1].GetKind() != ZEND_AST_ZVAL {
				is_constant = 0
			}
		} else {
			ZendEvalConstExpr(elem_ast.GetChild()[0])
			if elem_ast.GetChild()[0].GetKind() != ZEND_AST_ZVAL {
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
		var elem_ast *ZendAst = list.GetChild()[i]
		var value_ast *ZendAst = elem_ast.GetChild()[0]
		var key_ast *ZendAst
		var value *types.Zval = ZendAstGetZval(value_ast)
		if elem_ast.GetKind() == ZEND_AST_UNPACK {
			if value.IsArray() {
				var ht *types.Array = value.GetArr()
				var val *types.Zval
				var key *types.String
				var __ht *types.Array = ht
				for _, _p := range __ht.foreachData() {
					var _z *types.Zval = _p.GetVal()

					key = _p.GetKey()
					val = _z
					if key != nil {
						faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot unpack array with string keys")
					}
					if result.GetArr().NextIndexInsert(val) == nil {
						ZvalPtrDtor(result)
						return 0
					}
					val.TryAddRefcount()
				}
				continue
			} else {
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Only arrays and Traversables can be unpacked")
			}
		}
		value.TryAddRefcount()
		key_ast = elem_ast.GetChild()[1]
		if key_ast != nil {
			var key *types.Zval = ZendAstGetZval(key_ast)
			switch key.GetType() {
			case types.IS_LONG:
				result.GetArr().IndexUpdate(key.GetLval(), value)
			case types.IS_STRING:
				result.GetArr().SymtableUpdate(key.GetStr().GetStr(), value)
			case types.IS_DOUBLE:
				result.GetArr().IndexUpdate(ZendDvalToLval(key.GetDval()), value)
			case types.IS_FALSE:
				result.GetArr().IndexUpdate(0, value)
			case types.IS_TRUE:
				result.GetArr().IndexUpdate(1, value)
			case types.IS_NULL:
				result.GetArr().KeyUpdate(types.ZSTR_EMPTY_ALLOC().GetStr(), value)
			default:
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Illegal offset type")
			}
		} else {
			if result.GetArr().NextIndexInsert(value) == nil {
				ZvalPtrDtorNogc(value)
				ZvalPtrDtor(result)
				return 0
			}
		}
	}
	return 1
}
func ZendCompileBinaryOp(result *Znode, ast *ZendAst) {
	var left_ast *ZendAst = ast.GetChild()[0]
	var right_ast *ZendAst = ast.GetChild()[1]
	var opcode uint32 = ast.GetAttr()
	if (opcode == ZEND_ADD || opcode == ZEND_SUB) && left_ast.GetKind() == ZEND_AST_BINARY_OP && left_ast.GetAttr() == ZEND_CONCAT {
		faults.Error(faults.E_DEPRECATED, "The behavior of unparenthesized expressions containing both '.' and '+'/'-' will change in PHP 8: '+'/'-' will take a higher precedence")
	}
	if (opcode == ZEND_SL || opcode == ZEND_SR) && (left_ast.GetKind() == ZEND_AST_BINARY_OP && left_ast.GetAttr() == ZEND_CONCAT || right_ast.GetKind() == ZEND_AST_BINARY_OP && right_ast.GetAttr() == ZEND_CONCAT) {
		faults.Error(faults.E_DEPRECATED, "The behavior of unparenthesized expressions containing both '.' and '>>'/'<<' will change in PHP 8: '<<'/'>>' will take a higher precedence")
	}
	if opcode == ZEND_PARENTHESIZED_CONCAT {
		opcode = ZEND_CONCAT
	}
	var left_node Znode
	var right_node Znode
	ZendCompileExpr(&left_node, left_ast)
	ZendCompileExpr(&right_node, right_ast)
	if left_node.GetOpType() == IS_CONST && right_node.GetOpType() == IS_CONST {
		if ZendTryCtEvalBinaryOp(result.GetConstant(), opcode, left_node.GetConstant(), right_node.GetConstant()) != 0 {
			result.SetOpType(IS_CONST)
			ZvalPtrDtor(left_node.GetConstant())
			ZvalPtrDtor(right_node.GetConstant())
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
					ZendEmitOpTmp(&left_node, ZEND_CAST, &left_node, nil).SetExtendedValue(types.IS_STRING)
				} else {
					ConvertToString(left_node.GetConstant())
				}
			}
			if right_node.GetOpType() == IS_CONST {
				if right_node.GetConstant().IsArray() {
					ZendEmitOpTmp(&right_node, ZEND_CAST, &right_node, nil).SetExtendedValue(types.IS_STRING)
				} else {
					ConvertToString(right_node.GetConstant())
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
func ZendCompileGreater(result *Znode, ast *ZendAst) {
	var left_ast *ZendAst = ast.GetChild()[0]
	var right_ast *ZendAst = ast.GetChild()[1]
	var left_node Znode
	var right_node Znode
	b.Assert(ast.GetKind() == ZEND_AST_GREATER || ast.GetKind() == ZEND_AST_GREATER_EQUAL)
	ZendCompileExpr(&left_node, left_ast)
	ZendCompileExpr(&right_node, right_ast)
	if left_node.GetOpType() == IS_CONST && right_node.GetOpType() == IS_CONST {
		result.SetOpType(IS_CONST)
		ZendCtEvalGreater(result.GetConstant(), ast.GetKind(), left_node.GetConstant(), right_node.GetConstant())
		ZvalPtrDtor(left_node.GetConstant())
		ZvalPtrDtor(right_node.GetConstant())
		return
	}
	ZendEmitOpTmp(result, b.Cond(ast.GetKind() == ZEND_AST_GREATER, ZEND_IS_SMALLER, ZEND_IS_SMALLER_OR_EQUAL), &right_node, &left_node)
}
func ZendCompileUnaryOp(result *Znode, ast *ZendAst) {
	var expr_ast *ZendAst = ast.GetChild()[0]
	var opcode uint32 = ast.GetAttr()
	var expr_node Znode
	ZendCompileExpr(&expr_node, expr_ast)
	if expr_node.GetOpType() == IS_CONST {
		result.SetOpType(IS_CONST)
		ZendCtEvalUnaryOp(result.GetConstant(), opcode, expr_node.GetConstant())
		ZvalPtrDtor(expr_node.GetConstant())
		return
	}
	ZendEmitOpTmp(result, opcode, &expr_node, nil)
}
