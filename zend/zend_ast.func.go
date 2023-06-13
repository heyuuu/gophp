package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
)

func ZendAstGetZval(ast *ZendAst) *types.Zval {
	return ast.AsAstZval().GetVal()
}
func ZendAstGetStr(ast *ZendAst) *types.String {
	var zv *types.Zval = ZendAstGetZval(ast)
	b.Assert(zv.IsString())
	return zv.String()
}
func ZendAstGetStrVal(ast *ZendAst) string {
	var zv *types.Zval = ZendAstGetZval(ast)
	b.Assert(zv.IsString())
	return zv.StringVal()
}
func ZendAstGetConstantName(ast *ZendAst) *types.String {
	b.Assert(ast.GetKind() == ZEND_AST_CONSTANT)
	b.Assert((*ZendAstZval)(ast).GetVal().IsString())
	return (*ZendAstZval)(ast).GetVal().String()
}
func ZendAstGetLineno(ast *ZendAst) uint32 {
	if ast.GetKind() == ZEND_AST_ZVAL {
		astZval := (*ZendAstZval)(ast)
		return astZval.GetLineno()
	} else {
		return ast.GetLineno()
	}
}
func ZendAstAlloc(size int) any {
	return b.Malloc(size)
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
	var ast *ZendAstZval = NewAstZval(ZEND_AST_ZVAL, attr, zv, lineno)
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
func ZendAstCreateConstant(name *types.String, attr ZendAstAttr) *ZendAst {
	zv := types.NewZvalString(name.GetStr())
	lineno := CG__().GetZendLineno()
	var ast *ZendAstZval = NewAstZval(ZEND_AST_CONSTANT, attr, zv, lineno)
	return (*ZendAst)(ast)
}

func astLinenoByChildrenEx(children []*ZendAst) uint32 {
	lineno := uint32(CG__().GetZendLineno())
	for _, child := range children {
		if child != nil {
			childLineno := ZendAstGetLineno(child)
			if childLineno < lineno {
				lineno = childLineno
			}
			return lineno
		}
	}
	return lineno
}

func AstCreateEx(kind ZendAstKind, attr ZendAstAttr, children ...*ZendAst) *ZendAst {
	lineno := uint32(CG__().GetZendLineno())
	for _, child := range children {
		if child != nil {
			lineno = ZendAstGetLineno(child)
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
			lineno = b.Min(lineno, ZendAstGetLineno(child))
			break
		}
	}
	return NewAst(kind, 0, lineno, children)
}

func ZendAstAddArrayElement(result *types.Zval, offset *types.Zval, expr *types.Zval) int {
	switch offset.GetType() {
	case types.IS_UNDEF:
		if result.Array().Append(expr) == nil {
			faults.Error(faults.E_WARNING, "Cannot add element to the array as the next element is already occupied")
			// ZvalPtrDtorNogc(expr)
		}
	case types.IS_STRING:
		result.Array().SymtableUpdate(offset.String().GetStr(), expr)

	case types.IS_NULL:
		result.Array().SymtableUpdate(types.NewString("").GetStr(), expr)
	case types.IS_LONG:
		result.Array().IndexUpdate(offset.Long(), expr)
	case types.IS_FALSE:
		result.Array().IndexUpdate(0, expr)
	case types.IS_TRUE:
		result.Array().IndexUpdate(1, expr)
	case types.IS_DOUBLE:
		result.Array().IndexUpdate(operators.DvalToLval(offset.Double()), expr)
	case types.IS_RESOURCE:
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
	switch ast.GetKind() {
	case ZEND_AST_BINARY_OP:
		if ZendAstEvaluate(&op1, ast.GetChild()[0], scope) != types.SUCCESS {
			ret = types.FAILURE
		} else if ZendAstEvaluate(&op2, ast.GetChild()[1], scope) != types.SUCCESS {
			// ZvalPtrDtorNogc(&op1)
			ret = types.FAILURE
		} else {
			var op BinaryOpType = GetBinaryOp(ast.GetAttr())
			ret = op(result, &op1, &op2)
			// ZvalPtrDtorNogc(&op1)
			// ZvalPtrDtorNogc(&op2)
		}
	case ZEND_AST_GREATER:
		fallthrough
	case ZEND_AST_GREATER_EQUAL:
		if ZendAstEvaluate(&op1, ast.GetChild()[0], scope) != types.SUCCESS {
			ret = types.FAILURE
		} else if ZendAstEvaluate(&op2, ast.GetChild()[1], scope) != types.SUCCESS {
			// ZvalPtrDtorNogc(&op1)
			ret = types.FAILURE
		} else {

			/* op1 > op2 is the same as op2 < op1 */

			var op BinaryOpType = b.Cond(ast.GetKind() == ZEND_AST_GREATER, operators.IsSmallerFunction, operators.IsSmallerOrEqualFunction)
			ret = op(result, &op2, &op1)
			// ZvalPtrDtorNogc(&op1)
			// ZvalPtrDtorNogc(&op2)
		}
	case ZEND_AST_UNARY_OP:
		if ZendAstEvaluate(&op1, ast.GetChild()[0], scope) != types.SUCCESS {
			ret = types.FAILURE
		} else {
			var op UnaryOpType = GetUnaryOp(ast.GetAttr())
			ret = op(result, &op1)
			// ZvalPtrDtorNogc(&op1)
		}
	case ZEND_AST_ZVAL:
		var zv *types.Zval = ZendAstGetZval(ast)
		types.ZVAL_COPY(result, zv)
	case ZEND_AST_CONSTANT:
		var name *types.String = ZendAstGetConstantName(ast)
		var zv *types.Zval = ZendGetConstantEx(name.GetStr(), scope, ast.GetAttr())
		if zv == nil {
			result.SetUndef()
			ret = ZendUseUndefinedConstant(name, ast.GetAttr(), result)
			break
		}
		types.ZVAL_COPY_OR_DUP(result, zv)
	case ZEND_AST_CONSTANT_CLASS:
		if scope != nil {
			result.SetStringVal(scope.Name())
		} else {
			result.SetStringVal("")
		}
	case ZEND_AST_CLASS_NAME:
		if scope == nil {
			faults.ThrowError(nil, "Cannot use \"self\" when no class scope is active")
			return types.FAILURE
		}
		if ast.GetAttr() == ZEND_FETCH_CLASS_SELF {
			result.SetStringVal(scope.Name())
		} else if ast.GetAttr() == ZEND_FETCH_CLASS_PARENT {
			if !(scope.GetParent()) {
				faults.ThrowError(nil, "Cannot use \"parent\" when current class scope has no parent")
				return types.FAILURE
			}
			result.SetStringVal(scope.GetParent().name.GetStr())
		} else {
			b.Assert(false)
		}
	case ZEND_AST_AND:
		if ZendAstEvaluate(&op1, ast.GetChild()[0], scope) != types.SUCCESS {
			ret = types.FAILURE
			break
		}
		if operators.ZvalIsTrue(&op1) {
			if ZendAstEvaluate(&op2, ast.GetChild()[1], scope) != types.SUCCESS {
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
		if ZendAstEvaluate(&op1, ast.GetChild()[0], scope) != types.SUCCESS {
			ret = types.FAILURE
			break
		}
		if operators.ZvalIsTrue(&op1) {
			result.SetTrue()
		} else {
			if ZendAstEvaluate(&op2, ast.GetChild()[1], scope) != types.SUCCESS {
				// ZvalPtrDtorNogc(&op1)
				ret = types.FAILURE
				break
			}
			result.SetBool(operators.ZvalIsTrue(&op2))
			// ZvalPtrDtorNogc(&op2)
		}
		// ZvalPtrDtorNogc(&op1)
	case ZEND_AST_CONDITIONAL:
		if ZendAstEvaluate(&op1, ast.GetChild()[0], scope) != types.SUCCESS {
			ret = types.FAILURE
			break
		}
		if operators.ZvalIsTrue(&op1) {
			if ast.GetChild()[1] == nil {
				*result = op1
			} else {
				if ZendAstEvaluate(result, ast.GetChild()[1], scope) != types.SUCCESS {
					// ZvalPtrDtorNogc(&op1)
					ret = types.FAILURE
					break
				}
				// ZvalPtrDtorNogc(&op1)
			}
		} else {
			if ZendAstEvaluate(result, ast.GetChild()[2], scope) != types.SUCCESS {
				// ZvalPtrDtorNogc(&op1)
				ret = types.FAILURE
				break
			}
			// ZvalPtrDtorNogc(&op1)
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
				// ZvalPtrDtorNogc(&op1)
				ret = types.FAILURE
				break
			}
			// ZvalPtrDtorNogc(&op1)
		}
	case ZEND_AST_UNARY_PLUS:
		if ZendAstEvaluate(&op2, ast.GetChild()[0], scope) != types.SUCCESS {
			ret = types.FAILURE
		} else {
			op1.SetLong(0)
			ret = operators.AddFunction(result, &op1, &op2)
		}
	case ZEND_AST_UNARY_MINUS:
		if ZendAstEvaluate(&op2, ast.GetChild()[0], scope) != types.SUCCESS {
			ret = types.FAILURE
		} else {
			op1.SetLong(0)
			ret = operators.SubFunction(result, &op1, &op2)
			// ZvalPtrDtorNogc(&op2)
		}
	case ZEND_AST_ARRAY:
		var i uint32
		var list *ZendAstList = ast.AsAstList()
		if list.GetChildren() == 0 {
			result.SetEmptyArray()
			break
		}
		ArrayInit(result)
		for i = 0; i < list.GetChildren(); i++ {
			var elem *ZendAst = list.GetChild()[i]
			if elem.GetKind() == ZEND_AST_UNPACK {
				if ZendAstEvaluate(&op1, elem.GetChild()[0], scope) != types.SUCCESS {
					return types.FAILURE
				}
				if ZendAstAddUnpackedElement(result, &op1) != types.SUCCESS {
					return types.FAILURE
				}
				continue
			}
			if elem.GetChild()[1] != nil {
				if ZendAstEvaluate(&op1, elem.GetChild()[1], scope) != types.SUCCESS {
					// ZvalPtrDtorNogc(result)
					return types.FAILURE
				}
			} else {
				op1.SetUndef()
			}
			if ZendAstEvaluate(&op2, elem.GetChild()[0], scope) != types.SUCCESS {
				// ZvalPtrDtorNogc(&op1)
				// ZvalPtrDtorNogc(result)
				return types.FAILURE
			}
			if ZendAstAddArrayElement(result, &op1, &op2) != types.SUCCESS {
				// ZvalPtrDtorNogc(&op1)
				// ZvalPtrDtorNogc(&op2)
				// ZvalPtrDtorNogc(result)
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
			ret = types.FAILURE
		} else {
			ZendFetchDimensionConst(result, &op1, &op2, b.Cond((ast.GetAttr()&ZEND_DIM_IS) != 0, BP_VAR_IS, BP_VAR_R))
		}
	default:
		faults.ThrowError(nil, "Unsupported constant expression")
		ret = types.FAILURE
	}
	return ret
}

func AstTreeCopy(ast *ZendAst) *ZendAst {
	if ast.GetKind() == ZEND_AST_ZVAL {
		newAst := CopyAstZval(ast.AsAstZval())
		return (*ZendAst)(newAst)
	} else if ast.GetKind() == ZEND_AST_CONSTANT {
		constantName := ZendAstGetConstantName(ast).GetStr()
		newAst := NewAstZval(ZEND_AST_CONSTANT, ast.GetAttr(), types.NewZvalString(constantName), 0)
		return (*ZendAst)(newAst)
	} else {
		return CopyAst(ast, func(child *ZendAst) *ZendAst {
			return AstTreeCopy(child)
		})
	}
}

func ZendAstApply(ast *ZendAst, fn ZendAstApplyFunc) {
	if ast.IsList() {
		var list *ZendAstList = ast.AsAstList()
		var i uint32
		for i = 0; i < list.GetChildren(); i++ {
			fn(list.GetChild()[i])
		}
	} else {
		var i uint32
		var children uint32 = ast.GetChildren()
		for i = 0; i < children; i++ {
			fn(ast.GetChild()[i])
		}
	}
}
func ZendAstExportStr(str *SmartStr, s string) {
	for _, c := range []byte(s) {
		if c == '\'' || c == '\\' {
			str.WriteByte('\\')
			str.WriteByte(c)
		} else {
			str.WriteByte(c)
		}
	}
}
func ZendAstExportQstr(str *SmartStr, quote byte, s *types.String) {
	var i int
	for i = 0; i < s.GetLen(); i++ {
		var c uint8 = s.GetStr()[i]
		if c < ' ' {
			switch c {
			case '\n':
				str.WriteString("\\n")
			case '\r':
				str.WriteString("\\r")
			case '\t':
				str.WriteString("\\t")
			case 'f':
				str.WriteString("\\f")
			case 'v':
				str.WriteString("\\v")
			case 'e':
				str.WriteString("\\e")
			default:
				str.WriteString("\\0")
				str.WriteByte('0' + c/8)
				str.WriteByte('0' + c%8)
			}
		} else {
			if c == quote || c == '$' || c == '\\' {
				str.WriteByte('\\')
			}
			str.WriteByte(c)
		}
	}
}
func ZendAstExportIndent(str *SmartStr, indent int) {
	for indent > 0 {
		str.WriteString("    ")
		indent--
	}
}
func ZendAstExportName(str *SmartStr, ast *ZendAst, priority int, indent int) {
	if ast.GetKind() == ZEND_AST_ZVAL {
		var zv *types.Zval = ZendAstGetZval(ast)
		if zv.IsString() {
			str.WriteString(zv.String().GetStr())
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
				str.WriteByte('\\')
			} else if ast.GetAttr() == ZEND_NAME_RELATIVE {
				str.WriteString("namespace\\")
			}
			str.WriteString(zv.String().GetStr())
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
		if zv.IsString() && ZendAstValidVarName(zv.String().GetVal(), zv.String().GetLen()) != 0 {
			str.WriteString(zv.String().GetStr())
			return
		}
	} else if ast.GetKind() == ZEND_AST_VAR {
		ZendAstExportEx(str, ast, 0, indent)
		return
	}
	str.WriteByte('{')
	ZendAstExportName(str, ast, 0, indent)
	str.WriteByte('}')
}
func ZendAstExportList(str *SmartStr, list *ZendAstList, separator bool, priority int, indent int) {
	var i uint32 = 0
	for i < list.GetChildren() {
		if i != 0 && separator {
			str.WriteString(", ")
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
			ZendAstExportQstr(str, quote, zv.String())
		} else if ast.GetKind() == ZEND_AST_VAR && ast.GetChild()[0].GetKind() == ZEND_AST_ZVAL && (i+1 == list.GetChildren() || list.GetChild()[i+1].GetKind() != ZEND_AST_ZVAL || ZendAstVarNeedsBraces(ZendAstGetZval(list.GetChild()[i+1]).StringVal()[0]) == 0) {
			ZendAstExportEx(str, ast, 0, indent)
		} else {
			str.WriteByte('{')
			ZendAstExportEx(str, ast, 0, indent)
			str.WriteByte('}')
		}
		i++
	}
}
func ZendAstExportNameListEx(str *SmartStr, list *ZendAstList, indent int, separator string) {
	var i uint32 = 0
	for i < list.GetChildren() {
		if i != 0 {
			str.WriteString(b.CastStrAuto(separator))
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
			str.WriteString(", ")
		}
		if (list.GetChild()[i].GetAttr() & ZEND_BIND_REF) != 0 {
			str.WriteByte('&')
		}
		str.WriteByte('$')
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
			str.WriteByte(';')
		}
		str.WriteByte('\n')
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
				str.WriteString("if (")
			} else {
				ZendAstExportIndent(str, indent)
				str.WriteString("} elseif (")
			}
			ZendAstExportEx(str, ast.GetChild()[0], 0, indent)
			str.WriteString(") {\n")
			ZendAstExportStmt(str, ast.GetChild()[1], indent+1)
		} else {
			ZendAstExportIndent(str, indent)
			str.WriteString("} else ")
			if ast.GetChild()[1] != nil && ast.GetChild()[1].GetKind() == ZEND_AST_IF {
				list = (*ZendAstList)(ast.GetChild()[1])
				goto tail_call
			} else {
				str.WriteString("{\n")
				ZendAstExportStmt(str, ast.GetChild()[1], indent+1)
			}
		}
		i++
	}
	ZendAstExportIndent(str, indent)
	str.WriteByte('}')
}
func ZendAstExportZval(str *SmartStr, zv *types.Zval, priority int, indent int) {
	zv = types.ZVAL_DEREF(zv)
	switch zv.GetType() {
	case types.IS_NULL:
		str.WriteString("null")
	case types.IS_FALSE:
		str.WriteString("false")
	case types.IS_TRUE:
		str.WriteString("true")
	case types.IS_LONG:
		str.AppendLong(zv.Long())
	case types.IS_DOUBLE:
		key := ZendSprintf("%.*G", int(EG__().GetPrecision()), zv.Double())
		str.WriteString(key)
	case types.IS_STRING:
		str.WriteByte('\'')
		ZendAstExportStr(str, zv.StringVal())
		str.WriteByte('\'')
	case types.IS_ARRAY:
		str.WriteByte('[')
		first := 1
		zv.Array().Foreach(func(key types.ArrayKey, value *types.Zval) {
			if first != 0 {
				first = 0
			} else {
				str.WriteString(", ")
			}
			if key.IsStrKey() {
				str.WriteByte('\'')
				ZendAstExportStr(str, key.StrKey())
				str.WriteString("' => ")
			} else {
				str.AppendLong(key.IdxKey())
				str.WriteString(" => ")
			}
			ZendAstExportZval(str, value, 0, indent)
		})
		str.WriteByte(']')
	case types.IS_CONSTANT_AST:
		ZendAstExportEx(str, types.Z_ASTVAL_P(zv), priority, indent)
	default:

	}
}
func ZendAstExportClassNoHeader(str *SmartStr, decl *ZendAstDecl, indent int) {
	if decl.GetChild()[0] != nil {
		str.WriteString(" extends ")
		ZendAstExportNsName(str, decl.GetChild()[0], 0, indent)
	}
	if decl.GetChild()[1] != nil {
		str.WriteString(" implements ")
		ZendAstExportEx(str, decl.GetChild()[1], 0, indent)
	}
	str.WriteString(" {\n")
	ZendAstExportStmt(str, decl.GetChild()[2], indent+1)
	ZendAstExportIndent(str, indent)
	str.WriteString("}")
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
		str.WriteString(name.GetStr())
	case ZEND_AST_CONSTANT_CLASS:
		str.WriteString("__CLASS__")
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
			str.WriteString("public ")
		} else if decl.IsProtected() {
			str.WriteString("protected ")
		} else if decl.IsPrivate() {
			str.WriteString("private ")
		}
		if decl.IsStatic() {
			str.WriteString("static ")
		}
		if decl.IsAbstract() {
			str.WriteString("abstract ")
		}
		if decl.IsFinal() {
			str.WriteString("final ")
		}
		if decl.GetKind() == ZEND_AST_ARROW_FUNC {
			str.WriteString("fn")
		} else {
			str.WriteString("function ")
		}
		if decl.IsReturnReference() {
			str.WriteByte('&')
		}
		if ast.GetKind() != ZEND_AST_CLOSURE && ast.GetKind() != ZEND_AST_ARROW_FUNC {
			str.WriteString(decl.GetName().GetStr())
		}
		str.WriteByte('(')
		ZendAstExportEx(str, decl.GetChild()[0], 0, indent)
		str.WriteByte(')')
		ZendAstExportEx(str, decl.GetChild()[1], 0, indent)
		if decl.GetChild()[3] != nil {
			str.WriteString(": ")
			if (decl.GetChild()[3].GetAttr() & ZEND_TYPE_NULLABLE) != 0 {
				str.WriteByte('?')
			}
			ZendAstExportNsName(str, decl.GetChild()[3], 0, indent)
		}
		if decl.GetChild()[2] != nil {
			if decl.GetKind() == ZEND_AST_ARROW_FUNC {
				b.Assert(decl.GetChild()[2].GetKind() == ZEND_AST_RETURN)
				str.WriteString(" => ")
				ZendAstExportEx(str, decl.GetChild()[2].GetChild()[0], 0, indent)
				break
			}
			str.WriteString(" {\n")
			ZendAstExportStmt(str, decl.GetChild()[2], indent+1)
			ZendAstExportIndent(str, indent)
			str.WriteByte('}')
			if ast.GetKind() != ZEND_AST_CLOSURE {
				str.WriteByte('\n')
			}
		} else {
			str.WriteString(";\n")
		}
	case ZEND_AST_CLASS:
		decl = (*ZendAstDecl)(ast)
		if decl.IsInterface() {
			str.WriteString("interface ")
		} else if decl.IsTrait() {
			str.WriteString("trait ")
		} else {
			if decl.IsExplicitAbstractClass() {
				str.WriteString("abstract ")
			}
			if decl.IsFinal() {
				str.WriteString("final ")
			}
			str.WriteString("class ")
		}
		str.WriteString(decl.GetName().GetStr())
		ZendAstExportClassNoHeader(str, decl, indent)
		str.WriteByte('\n')
	case ZEND_AST_ARG_LIST:
		fallthrough
	case ZEND_AST_EXPR_LIST:
		fallthrough
	case ZEND_AST_PARAM_LIST:
	simple_list:
		ZendAstExportList(str, (*ZendAstList)(ast), true, 20, indent)
	case ZEND_AST_ARRAY:
		str.WriteByte('[')
		ZendAstExportList(str, (*ZendAstList)(ast), true, 20, indent)
		str.WriteByte(']')
	case ZEND_AST_ENCAPS_LIST:
		str.WriteByte('"')
		ZendAstExportEncapsList(str, '"', (*ZendAstList)(ast), indent)
		str.WriteByte('"')
	case ZEND_AST_STMT_LIST:
		fallthrough
	case ZEND_AST_TRAIT_ADAPTATIONS:
		ZendAstExportStmt(str, ast, indent)
	case ZEND_AST_IF:
		ZendAstExportIfStmt(str, (*ZendAstList)(ast), indent)
	case ZEND_AST_SWITCH_LIST:
		fallthrough
	case ZEND_AST_CATCH_LIST:
		ZendAstExportList(str, (*ZendAstList)(ast), false, 0, indent)
	case ZEND_AST_CLOSURE_USES:
		str.WriteString(" use(")
		ZendAstExportVarList(str, (*ZendAstList)(ast), indent)
		str.WriteByte(')')
	case ZEND_AST_PROP_GROUP:
		var type_ast *ZendAst = ast.GetChild()[0]
		var prop_ast *ZendAst = ast.GetChild()[1]
		if (ast.GetAttr() & types.AccPublic) != 0 {
			str.WriteString("public ")
		} else if (ast.GetAttr() & types.AccProtected) != 0 {
			str.WriteString("protected ")
		} else if (ast.GetAttr() & types.AccPrivate) != 0 {
			str.WriteString("private ")
		}
		if (ast.GetAttr() & types.AccStatic) != 0 {
			str.WriteString("static ")
		}
		if type_ast != nil {
			if (type_ast.GetAttr() & ZEND_TYPE_NULLABLE) != 0 {
				str.WriteByte('?')
			}
			ZendAstExportNsName(str, type_ast, 0, indent)
			str.WriteByte(' ')
		}
		ast = prop_ast
		goto simple_list
	case ZEND_AST_CONST_DECL:
		fallthrough
	case ZEND_AST_CLASS_CONST_DECL:
		str.WriteString("const ")
		goto simple_list
	case ZEND_AST_NAME_LIST:
		ZendAstExportNameList(str, (*ZendAstList)(ast), indent)
	case ZEND_AST_USE:
		str.WriteString("use ")
		if ast.GetAttr() == T_FUNCTION {
			str.WriteString("function ")
		} else if ast.GetAttr() == T_CONST {
			str.WriteString("const ")
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
		str.WriteByte('$')
		ZendAstExportVar(str, ast.GetChild()[0], 0, indent)
	case ZEND_AST_CONST:
		ZendAstExportNsName(str, ast.GetChild()[0], 0, indent)
	case ZEND_AST_UNPACK:
		str.WriteString("...")
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
		str.WriteByte('`')
		if ast.GetChild()[0].GetKind() == ZEND_AST_ENCAPS_LIST {
			ZendAstExportEncapsList(str, '`', (*ZendAstList)(ast.GetChild()[0]), indent)
		} else {
			var zv *types.Zval
			b.Assert(ast.GetChild()[0].GetKind() == ZEND_AST_ZVAL)
			zv = ZendAstGetZval(ast.GetChild()[0])
			b.Assert(zv.IsString())
			ZendAstExportQstr(str, '`', zv.String())
		}
		str.WriteByte('`')
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
		str.WriteByte(':')
	case ZEND_AST_REF:
		str.WriteByte('&')
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
		str.WriteString("goto ")
		ZendAstExportName(str, ast.GetChild()[0], 0, indent)
	case ZEND_AST_BREAK:
		APPEND_NODE_1("break")
		fallthrough
	case ZEND_AST_CONTINUE:
		APPEND_NODE_1("continue")
		fallthrough
	case ZEND_AST_DIM:
		ZendAstExportEx(str, ast.GetChild()[0], 260, indent)
		str.WriteByte('[')
		if ast.GetChild()[1] != nil {
			ZendAstExportEx(str, ast.GetChild()[1], 0, indent)
		}
		str.WriteByte(']')
	case ZEND_AST_PROP:
		ZendAstExportEx(str, ast.GetChild()[0], 0, indent)
		str.WriteString("->")
		ZendAstExportVar(str, ast.GetChild()[1], 0, indent)
	case ZEND_AST_STATIC_PROP:
		ZendAstExportNsName(str, ast.GetChild()[0], 0, indent)
		str.WriteString("::$")
		ZendAstExportVar(str, ast.GetChild()[1], 0, indent)
	case ZEND_AST_CALL:
		ZendAstExportNsName(str, ast.GetChild()[0], 0, indent)
		str.WriteByte('(')
		ZendAstExportEx(str, ast.GetChild()[1], 0, indent)
		str.WriteByte(')')
	case ZEND_AST_CLASS_CONST:
		ZendAstExportNsName(str, ast.GetChild()[0], 0, indent)
		str.WriteString("::")
		ZendAstExportName(str, ast.GetChild()[1], 0, indent)
	case ZEND_AST_CLASS_NAME:
		ZendAstExportNsName(str, ast.GetChild()[0], 0, indent)
		str.WriteString("::class")
	case ZEND_AST_ASSIGN:
		BINARY_OP(" = ", 90, 91, 90)
		fallthrough
	case ZEND_AST_ASSIGN_REF:
		BINARY_OP(" =& ", 90, 91, 90)
		fallthrough
	case ZEND_AST_ASSIGN_OP:
		switch OpCode(ast.GetAttr()) {
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
			str.WriteString(" => ")
		}
		if ast.GetAttr() != 0 {
			str.WriteByte('&')
		}
		ZendAstExportEx(str, ast.GetChild()[0], 80, indent)
	case ZEND_AST_NEW:
		str.WriteString("new ")
		if ast.GetChild()[0].GetKind() == ZEND_AST_CLASS {
			str.WriteString("class")
			if ast.GetChild()[1].AsAstList().GetChildren() != 0 {
				str.WriteByte('(')
				ZendAstExportEx(str, ast.GetChild()[1], 0, indent)
				str.WriteByte(')')
			}
			ZendAstExportClassNoHeader(str, (*ZendAstDecl)(ast.GetChild()[0]), indent)
		} else {
			ZendAstExportNsName(str, ast.GetChild()[0], 0, indent)
			str.WriteByte('(')
			ZendAstExportEx(str, ast.GetChild()[1], 0, indent)
			str.WriteByte(')')
		}
	case ZEND_AST_INSTANCEOF:
		ZendAstExportEx(str, ast.GetChild()[0], 0, indent)
		str.WriteString(" instanceof ")
		ZendAstExportNsName(str, ast.GetChild()[1], 0, indent)
	case ZEND_AST_YIELD:
		if priority > 70 {
			str.WriteByte('(')
		}
		str.WriteString("yield ")
		if ast.GetChild()[0] != nil {
			if ast.GetChild()[1] != nil {
				ZendAstExportEx(str, ast.GetChild()[1], 70, indent)
				str.WriteString(" => ")
			}
			ZendAstExportEx(str, ast.GetChild()[0], 70, indent)
		}
		if priority > 70 {
			str.WriteByte(')')
		}
	case ZEND_AST_YIELD_FROM:
		PREFIX_OP("yield from ", 85, 86)
		fallthrough
	case ZEND_AST_COALESCE:
		BINARY_OP(" ?? ", 110, 111, 110)
		fallthrough
	case ZEND_AST_STATIC:
		str.WriteString("static $")
		ZendAstExportName(str, ast.GetChild()[0], 0, indent)
		APPEND_DEFAULT_VALUE(1)
		fallthrough
	case ZEND_AST_WHILE:
		str.WriteString("while (")
		ZendAstExportEx(str, ast.GetChild()[0], 0, indent)
		str.WriteString(") {\n")
		ZendAstExportStmt(str, ast.GetChild()[1], indent+1)
		ZendAstExportIndent(str, indent)
		str.WriteByte('}')
	case ZEND_AST_DO_WHILE:
		str.WriteString("do {\n")
		ZendAstExportStmt(str, ast.GetChild()[0], indent+1)
		ZendAstExportIndent(str, indent)
		str.WriteString("} while (")
		ZendAstExportEx(str, ast.GetChild()[1], 0, indent)
		str.WriteByte(')')
	case ZEND_AST_IF_ELEM:
		if ast.GetChild()[0] != nil {
			str.WriteString("if (")
			ZendAstExportEx(str, ast.GetChild()[0], 0, indent)
			str.WriteString(") {\n")
			ZendAstExportStmt(str, ast.GetChild()[1], indent+1)
		} else {
			str.WriteString("else {\n")
			ZendAstExportStmt(str, ast.GetChild()[1], indent+1)
		}
		ZendAstExportIndent(str, indent)
		str.WriteByte('}')
	case ZEND_AST_SWITCH:
		str.WriteString("switch (")
		ZendAstExportEx(str, ast.GetChild()[0], 0, indent)
		str.WriteString(") {\n")
		ZendAstExportEx(str, ast.GetChild()[1], 0, indent+1)
		ZendAstExportIndent(str, indent)
		str.WriteByte('}')
	case ZEND_AST_SWITCH_CASE:
		ZendAstExportIndent(str, indent)
		if ast.GetChild()[0] != nil {
			str.WriteString("case ")
			ZendAstExportEx(str, ast.GetChild()[0], 0, indent)
			str.WriteString(":\n")
		} else {
			str.WriteString("default:\n")
		}
		ZendAstExportStmt(str, ast.GetChild()[1], indent+1)
	case ZEND_AST_DECLARE:
		str.WriteString("declare(")
		b.Assert(ast.GetChild()[0].GetKind() == ZEND_AST_CONST_DECL)
		ZendAstExportList(str, (*ZendAstList)(ast.GetChild()[0]), true, 0, indent)
		str.WriteByte(')')
		if ast.GetChild()[1] != nil {
			str.WriteString(" {\n")
			ZendAstExportStmt(str, ast.GetChild()[1], indent+1)
			ZendAstExportIndent(str, indent)
			str.WriteByte('}')
		} else {
			str.WriteByte(';')
		}
	case ZEND_AST_PROP_ELEM:
		str.WriteByte('$')
		fallthrough
	case ZEND_AST_CONST_ELEM:
		ZendAstExportName(str, ast.GetChild()[0], 0, indent)
		APPEND_DEFAULT_VALUE(1)
		fallthrough
	case ZEND_AST_USE_TRAIT:
		str.WriteString("use ")
		ZendAstExportEx(str, ast.GetChild()[0], 0, indent)
		if ast.GetChild()[1] != nil {
			str.WriteString(" {\n")
			ZendAstExportEx(str, ast.GetChild()[1], 0, indent+1)
			ZendAstExportIndent(str, indent)
			str.WriteString("}")
		} else {
			str.WriteString(";")
		}
	case ZEND_AST_TRAIT_PRECEDENCE:
		ZendAstExportEx(str, ast.GetChild()[0], 0, indent)
		str.WriteString(" insteadof ")
		ZendAstExportEx(str, ast.GetChild()[1], 0, indent)
	case ZEND_AST_METHOD_REFERENCE:
		if ast.GetChild()[0] != nil {
			ZendAstExportName(str, ast.GetChild()[0], 0, indent)
			str.WriteString("::")
		}
		ZendAstExportName(str, ast.GetChild()[1], 0, indent)
	case ZEND_AST_NAMESPACE:
		str.WriteString("namespace")
		if ast.GetChild()[0] != nil {
			str.WriteByte(' ')
			ZendAstExportName(str, ast.GetChild()[0], 0, indent)
		}
		if ast.GetChild()[1] != nil {
			str.WriteString(" {\n")
			ZendAstExportStmt(str, ast.GetChild()[1], indent+1)
			ZendAstExportIndent(str, indent)
			str.WriteString("}\n")
		} else {
			str.WriteByte(';')
		}
	case ZEND_AST_USE_ELEM:
		fallthrough
	case ZEND_AST_TRAIT_ALIAS:
		ZendAstExportName(str, ast.GetChild()[0], 0, indent)
		if (ast.GetAttr() & types.AccPublic) != 0 {
			str.WriteString(" as public")
		} else if (ast.GetAttr() & types.AccProtected) != 0 {
			str.WriteString(" as protected")
		} else if (ast.GetAttr() & types.AccPrivate) != 0 {
			str.WriteString(" as private")
		} else if ast.GetChild()[1] != nil {
			str.WriteString(" as")
		}
		if ast.GetChild()[1] != nil {
			str.WriteByte(' ')
			ZendAstExportName(str, ast.GetChild()[1], 0, indent)
		}
	case ZEND_AST_METHOD_CALL:
		ZendAstExportEx(str, ast.GetChild()[0], 0, indent)
		str.WriteString("->")
		ZendAstExportVar(str, ast.GetChild()[1], 0, indent)
		str.WriteByte('(')
		ZendAstExportEx(str, ast.GetChild()[2], 0, indent)
		str.WriteByte(')')
	case ZEND_AST_STATIC_CALL:
		ZendAstExportNsName(str, ast.GetChild()[0], 0, indent)
		str.WriteString("::")
		ZendAstExportVar(str, ast.GetChild()[1], 0, indent)
		str.WriteByte('(')
		ZendAstExportEx(str, ast.GetChild()[2], 0, indent)
		str.WriteByte(')')
	case ZEND_AST_CONDITIONAL:
		if priority > 100 {
			str.WriteByte('(')
		}
		ZendAstExportEx(str, ast.GetChild()[0], 100, indent)
		if ast.GetChild()[1] != nil {
			str.WriteString(" ? ")
			ZendAstExportEx(str, ast.GetChild()[1], 101, indent)
			str.WriteString(" : ")
		} else {
			str.WriteString(" ?: ")
		}
		ZendAstExportEx(str, ast.GetChild()[2], 101, indent)
		if priority > 100 {
			str.WriteByte(')')
		}
	case ZEND_AST_TRY:
		str.WriteString("try {\n")
		ZendAstExportStmt(str, ast.GetChild()[0], indent+1)
		ZendAstExportIndent(str, indent)
		ZendAstExportEx(str, ast.GetChild()[1], 0, indent)
		if ast.GetChild()[2] != nil {
			str.WriteString("} finally {\n")
			ZendAstExportStmt(str, ast.GetChild()[2], indent+1)
			ZendAstExportIndent(str, indent)
		}
		str.WriteByte('}')
	case ZEND_AST_CATCH:
		str.WriteString("} catch (")
		ZendAstExportCatchNameList(str, ast.GetChild()[0].AsAstList(), indent)
		str.WriteString(" $")
		ZendAstExportVar(str, ast.GetChild()[1], 0, indent)
		str.WriteString(") {\n")
		ZendAstExportStmt(str, ast.GetChild()[2], indent+1)
		ZendAstExportIndent(str, indent)
	case ZEND_AST_PARAM:
		if ast.GetChild()[0] != nil {
			if (ast.GetChild()[0].GetAttr() & ZEND_TYPE_NULLABLE) != 0 {
				str.WriteByte('?')
			}
			ZendAstExportNsName(str, ast.GetChild()[0], 0, indent)
			str.WriteByte(' ')
		}
		if (ast.GetAttr() & ZEND_PARAM_REF) != 0 {
			str.WriteByte('&')
		}
		if (ast.GetAttr() & ZEND_PARAM_VARIADIC) != 0 {
			str.WriteString("...")
		}
		str.WriteByte('$')
		ZendAstExportName(str, ast.GetChild()[1], 0, indent)
		APPEND_DEFAULT_VALUE(2)
		fallthrough
	case ZEND_AST_FOR:
		str.WriteString("for (")
		ZendAstExportEx(str, ast.GetChild()[0], 0, indent)
		str.WriteByte(';')
		if ast.GetChild()[1] != nil {
			str.WriteByte(' ')
			ZendAstExportEx(str, ast.GetChild()[1], 0, indent)
		}
		str.WriteByte(';')
		if ast.GetChild()[2] != nil {
			str.WriteByte(' ')
			ZendAstExportEx(str, ast.GetChild()[2], 0, indent)
		}
		str.WriteString(") {\n")
		ZendAstExportStmt(str, ast.GetChild()[3], indent+1)
		ZendAstExportIndent(str, indent)
		str.WriteByte('}')
	case ZEND_AST_FOREACH:
		str.WriteString("foreach (")
		ZendAstExportEx(str, ast.GetChild()[0], 0, indent)
		str.WriteString(" as ")
		if ast.GetChild()[2] != nil {
			ZendAstExportEx(str, ast.GetChild()[2], 0, indent)
			str.WriteString(" => ")
		}
		ZendAstExportEx(str, ast.GetChild()[1], 0, indent)
		str.WriteString(") {\n")
		ZendAstExportStmt(str, ast.GetChild()[3], indent+1)
		ZendAstExportIndent(str, indent)
		str.WriteByte('}')
	default:

	}
	return
binary_op:
	if priority > p {
		str.WriteByte('(')
	}
	ZendAstExportEx(str, ast.GetChild()[0], pl, indent)
	str.WriteString(b.CastStrAuto(op))
	ZendAstExportEx(str, ast.GetChild()[1], pr, indent)
	if priority > p {
		str.WriteByte(')')
	}
	return
prefix_op:
	if priority > p {
		str.WriteByte('(')
	}
	str.WriteString(b.CastStrAuto(op))
	ZendAstExportEx(str, ast.GetChild()[0], pl, indent)
	if priority > p {
		str.WriteByte(')')
	}
	return
postfix_op:
	if priority > p {
		str.WriteByte('(')
	}
	ZendAstExportEx(str, ast.GetChild()[0], pl, indent)
	str.WriteString(b.CastStrAuto(op))
	if priority > p {
		str.WriteByte(')')
	}
	return
func_op:
	str.WriteString(b.CastStrAuto(op))
	str.WriteByte('(')
	ZendAstExportEx(str, ast.GetChild()[0], 0, indent)
	str.WriteByte(')')
	return
append_node_1:
	str.WriteString(b.CastStrAuto(op))
	if ast.GetChild()[0] != nil {
		str.WriteByte(' ')
		ast = ast.GetChild()[0]
		goto tail_call
	}
	return
append_str:
	str.WriteString(b.CastStrAuto(op))
	return
append_default_value:
	if ast.GetChild()[p] != nil {
		str.WriteString(" = ")
		ast = ast.GetChild()[p]
		goto tail_call
	}
	return
}
func ZendAstExport(prefix string, ast *ZendAst, suffix string) *types.String {
	var str SmartStr = MakeSmartStr(0)
	str.WriteString(b.CastStrAuto(prefix))
	ZendAstExportEx(&str, ast, 0, 0)
	str.WriteString(b.CastStrAuto(suffix))
	str.ZeroTail()
	return str.GetS()
}
