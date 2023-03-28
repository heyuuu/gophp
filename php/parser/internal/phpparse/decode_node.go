package phpparse

import "gophp/php/ast"

func decodeNode(data map[string]any) (node ast.Node, err error) {
	nodeType := data["nodeType"].(string)
	switch nodeType {
	case "Arg":
		node = &ast.Arg{
			Name:       data["name"],
			Value:      data["value"],
			ByRef:      data["byRef"],
			Unpack:     data["unpack"],
			Attributes: data["attributes"],
		}
	case "Attribute":
		node = &ast.Attribute{
			Name:       data["name"],
			Args:       data["args"],
			Attributes: data["attributes"],
		}
	case "AttributeGroup":
		node = &ast.AttributeGroup{
			Attrs:      data["attrs"],
			Attributes: data["attributes"],
		}
	case "Const":
		node = &ast.Const{
			Name:           data["name"],
			Value:          data["value"],
			NamespacedName: data["namespacedName"],
			Attributes:     data["attributes"],
		}
	case "ExprArray":
		node = &ast.ArrayExpr{
			Items:      data["items"],
			Attributes: data["attributes"],
		}
	case "ExprArrayDimFetch":
		node = &ast.ArrayDimFetchExpr{
			Var:        data["var"],
			Dim:        data["dim"],
			Attributes: data["attributes"],
		}
	case "ExprArrayItem":
		node = &ast.ArrayItemExpr{
			Key:        data["key"],
			Value:      data["value"],
			ByRef:      data["byRef"],
			Unpack:     data["unpack"],
			Attributes: data["attributes"],
		}
	case "ExprArrowFunction":
		node = &ast.ArrowFunctionExpr{
			Static:     data["static"],
			ByRef:      data["byRef"],
			Params:     data["params"],
			ReturnType: data["returnType"],
			Expr:       data["expr"],
			AttrGroups: data["attrGroups"],
			Attributes: data["attributes"],
		}
	case "ExprAssign":
		node = &ast.AssignExpr{
			Var:        data["var"],
			Expr:       data["expr"],
			Attributes: data["attributes"],
		}
	case "ExprAssignOpBitwiseAnd":
		node = &ast.AssignOpBitwiseAndExpr{
			Var:        data["var"],
			Expr:       data["expr"],
			Attributes: data["attributes"],
		}
	case "ExprAssignOpBitwiseOr":
		node = &ast.AssignOpBitwiseOrExpr{
			Var:        data["var"],
			Expr:       data["expr"],
			Attributes: data["attributes"],
		}
	case "ExprAssignOpBitwiseXor":
		node = &ast.AssignOpBitwiseXorExpr{
			Var:        data["var"],
			Expr:       data["expr"],
			Attributes: data["attributes"],
		}
	case "ExprAssignOpCoalesce":
		node = &ast.AssignOpCoalesceExpr{
			Var:        data["var"],
			Expr:       data["expr"],
			Attributes: data["attributes"],
		}
	case "ExprAssignOpConcat":
		node = &ast.AssignOpConcatExpr{
			Var:        data["var"],
			Expr:       data["expr"],
			Attributes: data["attributes"],
		}
	case "ExprAssignOpDiv":
		node = &ast.AssignOpDivExpr{
			Var:        data["var"],
			Expr:       data["expr"],
			Attributes: data["attributes"],
		}
	case "ExprAssignOpMinus":
		node = &ast.AssignOpMinusExpr{
			Var:        data["var"],
			Expr:       data["expr"],
			Attributes: data["attributes"],
		}
	case "ExprAssignOpMod":
		node = &ast.AssignOpModExpr{
			Var:        data["var"],
			Expr:       data["expr"],
			Attributes: data["attributes"],
		}
	case "ExprAssignOpMul":
		node = &ast.AssignOpMulExpr{
			Var:        data["var"],
			Expr:       data["expr"],
			Attributes: data["attributes"],
		}
	case "ExprAssignOpPlus":
		node = &ast.AssignOpPlusExpr{
			Var:        data["var"],
			Expr:       data["expr"],
			Attributes: data["attributes"],
		}
	case "ExprAssignOpPow":
		node = &ast.AssignOpPowExpr{
			Var:        data["var"],
			Expr:       data["expr"],
			Attributes: data["attributes"],
		}
	case "ExprAssignOpShiftLeft":
		node = &ast.AssignOpShiftLeftExpr{
			Var:        data["var"],
			Expr:       data["expr"],
			Attributes: data["attributes"],
		}
	case "ExprAssignOpShiftRight":
		node = &ast.AssignOpShiftRightExpr{
			Var:        data["var"],
			Expr:       data["expr"],
			Attributes: data["attributes"],
		}
	case "ExprAssignRef":
		node = &ast.AssignRefExpr{
			Var:        data["var"],
			Expr:       data["expr"],
			Attributes: data["attributes"],
		}
	case "ExprBinaryOpBitwiseAnd":
		node = &ast.BinaryOpBitwiseAndExpr{
			Left:       data["left"],
			Right:      data["right"],
			Attributes: data["attributes"],
		}
	case "ExprBinaryOpBitwiseOr":
		node = &ast.BinaryOpBitwiseOrExpr{
			Left:       data["left"],
			Right:      data["right"],
			Attributes: data["attributes"],
		}
	case "ExprBinaryOpBitwiseXor":
		node = &ast.BinaryOpBitwiseXorExpr{
			Left:       data["left"],
			Right:      data["right"],
			Attributes: data["attributes"],
		}
	case "ExprBinaryOpBooleanAnd":
		node = &ast.BinaryOpBooleanAndExpr{
			Left:       data["left"],
			Right:      data["right"],
			Attributes: data["attributes"],
		}
	case "ExprBinaryOpBooleanOr":
		node = &ast.BinaryOpBooleanOrExpr{
			Left:       data["left"],
			Right:      data["right"],
			Attributes: data["attributes"],
		}
	case "ExprBinaryOpCoalesce":
		node = &ast.BinaryOpCoalesceExpr{
			Left:       data["left"],
			Right:      data["right"],
			Attributes: data["attributes"],
		}
	case "ExprBinaryOpConcat":
		node = &ast.BinaryOpConcatExpr{
			Left:       data["left"],
			Right:      data["right"],
			Attributes: data["attributes"],
		}
	case "ExprBinaryOpDiv":
		node = &ast.BinaryOpDivExpr{
			Left:       data["left"],
			Right:      data["right"],
			Attributes: data["attributes"],
		}
	case "ExprBinaryOpEqual":
		node = &ast.BinaryOpEqualExpr{
			Left:       data["left"],
			Right:      data["right"],
			Attributes: data["attributes"],
		}
	case "ExprBinaryOpGreater":
		node = &ast.BinaryOpGreaterExpr{
			Left:       data["left"],
			Right:      data["right"],
			Attributes: data["attributes"],
		}
	case "ExprBinaryOpGreaterOrEqual":
		node = &ast.BinaryOpGreaterOrEqualExpr{
			Left:       data["left"],
			Right:      data["right"],
			Attributes: data["attributes"],
		}
	case "ExprBinaryOpIdentical":
		node = &ast.BinaryOpIdenticalExpr{
			Left:       data["left"],
			Right:      data["right"],
			Attributes: data["attributes"],
		}
	case "ExprBinaryOpLogicalAnd":
		node = &ast.BinaryOpLogicalAndExpr{
			Left:       data["left"],
			Right:      data["right"],
			Attributes: data["attributes"],
		}
	case "ExprBinaryOpLogicalOr":
		node = &ast.BinaryOpLogicalOrExpr{
			Left:       data["left"],
			Right:      data["right"],
			Attributes: data["attributes"],
		}
	case "ExprBinaryOpLogicalXor":
		node = &ast.BinaryOpLogicalXorExpr{
			Left:       data["left"],
			Right:      data["right"],
			Attributes: data["attributes"],
		}
	case "ExprBinaryOpMinus":
		node = &ast.BinaryOpMinusExpr{
			Left:       data["left"],
			Right:      data["right"],
			Attributes: data["attributes"],
		}
	case "ExprBinaryOpMod":
		node = &ast.BinaryOpModExpr{
			Left:       data["left"],
			Right:      data["right"],
			Attributes: data["attributes"],
		}
	case "ExprBinaryOpMul":
		node = &ast.BinaryOpMulExpr{
			Left:       data["left"],
			Right:      data["right"],
			Attributes: data["attributes"],
		}
	case "ExprBinaryOpNotEqual":
		node = &ast.BinaryOpNotEqualExpr{
			Left:       data["left"],
			Right:      data["right"],
			Attributes: data["attributes"],
		}
	case "ExprBinaryOpNotIdentical":
		node = &ast.BinaryOpNotIdenticalExpr{
			Left:       data["left"],
			Right:      data["right"],
			Attributes: data["attributes"],
		}
	case "ExprBinaryOpPlus":
		node = &ast.BinaryOpPlusExpr{
			Left:       data["left"],
			Right:      data["right"],
			Attributes: data["attributes"],
		}
	case "ExprBinaryOpPow":
		node = &ast.BinaryOpPowExpr{
			Left:       data["left"],
			Right:      data["right"],
			Attributes: data["attributes"],
		}
	case "ExprBinaryOpShiftLeft":
		node = &ast.BinaryOpShiftLeftExpr{
			Left:       data["left"],
			Right:      data["right"],
			Attributes: data["attributes"],
		}
	case "ExprBinaryOpShiftRight":
		node = &ast.BinaryOpShiftRightExpr{
			Left:       data["left"],
			Right:      data["right"],
			Attributes: data["attributes"],
		}
	case "ExprBinaryOpSmaller":
		node = &ast.BinaryOpSmallerExpr{
			Left:       data["left"],
			Right:      data["right"],
			Attributes: data["attributes"],
		}
	case "ExprBinaryOpSmallerOrEqual":
		node = &ast.BinaryOpSmallerOrEqualExpr{
			Left:       data["left"],
			Right:      data["right"],
			Attributes: data["attributes"],
		}
	case "ExprBinaryOpSpaceship":
		node = &ast.BinaryOpSpaceshipExpr{
			Left:       data["left"],
			Right:      data["right"],
			Attributes: data["attributes"],
		}
	case "ExprBitwiseNot":
		node = &ast.BitwiseNotExpr{
			Expr:       data["expr"],
			Attributes: data["attributes"],
		}
	case "ExprBooleanNot":
		node = &ast.BooleanNotExpr{
			Expr:       data["expr"],
			Attributes: data["attributes"],
		}
	case "ExprCastArray":
		node = &ast.CastArrayExpr{
			Expr:       data["expr"],
			Attributes: data["attributes"],
		}
	case "ExprCastBool":
		node = &ast.CastBoolExpr{
			Expr:       data["expr"],
			Attributes: data["attributes"],
		}
	case "ExprCastDouble":
		node = &ast.CastDoubleExpr{
			Expr:       data["expr"],
			Attributes: data["attributes"],
		}
	case "ExprCastInt":
		node = &ast.CastIntExpr{
			Expr:       data["expr"],
			Attributes: data["attributes"],
		}
	case "ExprCastObject":
		node = &ast.CastObjectExpr{
			Expr:       data["expr"],
			Attributes: data["attributes"],
		}
	case "ExprCastString":
		node = &ast.CastStringExpr{
			Expr:       data["expr"],
			Attributes: data["attributes"],
		}
	case "ExprCastUnset":
		node = &ast.CastUnsetExpr{
			Expr:       data["expr"],
			Attributes: data["attributes"],
		}
	case "ExprClassConstFetch":
		node = &ast.ClassConstFetchExpr{
			Class:      data["class"],
			Name:       data["name"],
			Attributes: data["attributes"],
		}
	case "ExprClone":
		node = &ast.CloneExpr{
			Expr:       data["expr"],
			Attributes: data["attributes"],
		}
	case "ExprClosure":
		node = &ast.ClosureExpr{
			Static:     data["static"],
			ByRef:      data["byRef"],
			Params:     data["params"],
			Uses:       data["uses"],
			ReturnType: data["returnType"],
			Stmts:      data["stmts"],
			AttrGroups: data["attrGroups"],
			Attributes: data["attributes"],
		}
	case "ExprClosureUse":
		node = &ast.ClosureUseExpr{
			Var:        data["var"],
			ByRef:      data["byRef"],
			Attributes: data["attributes"],
		}
	case "ExprConstFetch":
		node = &ast.ConstFetchExpr{
			Name:       data["name"],
			Attributes: data["attributes"],
		}
	case "ExprEmpty":
		node = &ast.EmptyExpr{
			Expr:       data["expr"],
			Attributes: data["attributes"],
		}
	case "ExprError":
		node = &ast.ErrorExpr{
			Attributes: data["attributes"],
		}
	case "ExprErrorSuppress":
		node = &ast.ErrorSuppressExpr{
			Expr:       data["expr"],
			Attributes: data["attributes"],
		}
	case "ExprEval":
		node = &ast.EvalExpr{
			Expr:       data["expr"],
			Attributes: data["attributes"],
		}
	case "ExprExit":
		node = &ast.ExitExpr{
			Expr:       data["expr"],
			Attributes: data["attributes"],
		}
	case "ExprFuncCall":
		node = &ast.FuncCallExpr{
			Name:       data["name"],
			Args:       data["args"],
			Attributes: data["attributes"],
		}
	case "ExprInclude":
		node = &ast.IncludeExpr{
			Expr:       data["expr"],
			Type:       data["type"],
			Attributes: data["attributes"],
		}
	case "ExprInstanceof":
		node = &ast.InstanceofExpr{
			Expr:       data["expr"],
			Class:      data["class"],
			Attributes: data["attributes"],
		}
	case "ExprIsset":
		node = &ast.IssetExpr{
			Vars:       data["vars"],
			Attributes: data["attributes"],
		}
	case "ExprList":
		node = &ast.ListExpr{
			Items:      data["items"],
			Attributes: data["attributes"],
		}
	case "ExprMatch":
		node = &ast.MatchExpr{
			Cond:       data["cond"],
			Arms:       data["arms"],
			Attributes: data["attributes"],
		}
	case "ExprMethodCall":
		node = &ast.MethodCallExpr{
			Var:        data["var"],
			Name:       data["name"],
			Args:       data["args"],
			Attributes: data["attributes"],
		}
	case "ExprNew":
		node = &ast.NewExpr{
			Class:      data["class"],
			Args:       data["args"],
			Attributes: data["attributes"],
		}
	case "ExprNullsafeMethodCall":
		node = &ast.NullsafeMethodCallExpr{
			Var:        data["var"],
			Name:       data["name"],
			Args:       data["args"],
			Attributes: data["attributes"],
		}
	case "ExprNullsafePropertyFetch":
		node = &ast.NullsafePropertyFetchExpr{
			Var:        data["var"],
			Name:       data["name"],
			Attributes: data["attributes"],
		}
	case "ExprPostDec":
		node = &ast.PostDecExpr{
			Var:        data["var"],
			Attributes: data["attributes"],
		}
	case "ExprPostInc":
		node = &ast.PostIncExpr{
			Var:        data["var"],
			Attributes: data["attributes"],
		}
	case "ExprPreDec":
		node = &ast.PreDecExpr{
			Var:        data["var"],
			Attributes: data["attributes"],
		}
	case "ExprPreInc":
		node = &ast.PreIncExpr{
			Var:        data["var"],
			Attributes: data["attributes"],
		}
	case "ExprPrint":
		node = &ast.PrintExpr{
			Expr:       data["expr"],
			Attributes: data["attributes"],
		}
	case "ExprPropertyFetch":
		node = &ast.PropertyFetchExpr{
			Var:        data["var"],
			Name:       data["name"],
			Attributes: data["attributes"],
		}
	case "ExprShellExec":
		node = &ast.ShellExecExpr{
			Parts:      data["parts"],
			Attributes: data["attributes"],
		}
	case "ExprStaticCall":
		node = &ast.StaticCallExpr{
			Class:      data["class"],
			Name:       data["name"],
			Args:       data["args"],
			Attributes: data["attributes"],
		}
	case "ExprStaticPropertyFetch":
		node = &ast.StaticPropertyFetchExpr{
			Class:      data["class"],
			Name:       data["name"],
			Attributes: data["attributes"],
		}
	case "ExprTernary":
		node = &ast.TernaryExpr{
			Cond:       data["cond"],
			If:         data["if"],
			Else:       data["else"],
			Attributes: data["attributes"],
		}
	case "ExprThrow":
		node = &ast.ThrowExpr{
			Expr:       data["expr"],
			Attributes: data["attributes"],
		}
	case "ExprUnaryMinus":
		node = &ast.UnaryMinusExpr{
			Expr:       data["expr"],
			Attributes: data["attributes"],
		}
	case "ExprUnaryPlus":
		node = &ast.UnaryPlusExpr{
			Expr:       data["expr"],
			Attributes: data["attributes"],
		}
	case "ExprVariable":
		node = &ast.VariableExpr{
			Name:       data["name"],
			Attributes: data["attributes"],
		}
	case "ExprYield":
		node = &ast.YieldExpr{
			Key:        data["key"],
			Value:      data["value"],
			Attributes: data["attributes"],
		}
	case "ExprYieldFrom":
		node = &ast.YieldFromExpr{
			Expr:       data["expr"],
			Attributes: data["attributes"],
		}
	case "Identifier":
		node = &ast.Identifier{
			Name:              data["name"],
			SpecialClassNames: data["specialClassNames"],
			Attributes:        data["attributes"],
		}
	case "IntersectionType":
		node = &ast.IntersectionType{
			Types:      data["types"],
			Attributes: data["attributes"],
		}
	case "MatchArm":
		node = &ast.MatchArm{
			Conds:      data["conds"],
			Body:       data["body"],
			Attributes: data["attributes"],
		}
	case "Name":
		node = &ast.Name{
			Parts:             data["parts"],
			SpecialClassNames: data["specialClassNames"],
			Attributes:        data["attributes"],
		}
	case "NameFullyQualified":
		node = &ast.NameFullyQualified{
			Parts:      data["parts"],
			Attributes: data["attributes"],
		}
	case "NameRelative":
		node = &ast.NameRelative{
			Parts:      data["parts"],
			Attributes: data["attributes"],
		}
	case "NullableType":
		node = &ast.NullableType{
			Type:       data["type"],
			Attributes: data["attributes"],
		}
	case "Param":
		node = &ast.Param{
			Type:       data["type"],
			ByRef:      data["byRef"],
			Variadic:   data["variadic"],
			Var:        data["var"],
			Default:    data["default"],
			Flags:      data["flags"],
			AttrGroups: data["attrGroups"],
			Attributes: data["attributes"],
		}
	case "ScalarDNumber":
		node = &ast.DNumberScalar{
			Value:      data["value"],
			Attributes: data["attributes"],
		}
	case "ScalarEncapsed":
		node = &ast.EncapsedScalar{
			Parts:      data["parts"],
			Attributes: data["attributes"],
		}
	case "ScalarEncapsedStringPart":
		node = &ast.EncapsedStringPartScalar{
			Value:      data["value"],
			Attributes: data["attributes"],
		}
	case "ScalarLNumber":
		node = &ast.LNumberScalar{
			Value:      data["value"],
			Attributes: data["attributes"],
		}
	case "ScalarMagicConstClass":
		node = &ast.MagicConstClassScalar{
			Attributes: data["attributes"],
		}
	case "ScalarMagicConstDir":
		node = &ast.MagicConstDirScalar{
			Attributes: data["attributes"],
		}
	case "ScalarMagicConstFile":
		node = &ast.MagicConstFileScalar{
			Attributes: data["attributes"],
		}
	case "ScalarMagicConstFunction":
		node = &ast.MagicConstFunctionScalar{
			Attributes: data["attributes"],
		}
	case "ScalarMagicConstLine":
		node = &ast.MagicConstLineScalar{
			Attributes: data["attributes"],
		}
	case "ScalarMagicConstMethod":
		node = &ast.MagicConstMethodScalar{
			Attributes: data["attributes"],
		}
	case "ScalarMagicConstNamespace":
		node = &ast.MagicConstNamespaceScalar{
			Attributes: data["attributes"],
		}
	case "ScalarMagicConstTrait":
		node = &ast.MagicConstTraitScalar{
			Attributes: data["attributes"],
		}
	case "ScalarString":
		node = &ast.StringScalar{
			Value:        data["value"],
			Replacements: data["replacements"],
			Attributes:   data["attributes"],
		}
	case "StmtBreak":
		node = &ast.BreakStmt{
			Num:        data["num"],
			Attributes: data["attributes"],
		}
	case "StmtCase":
		node = &ast.CaseStmt{
			Cond:       data["cond"],
			Stmts:      data["stmts"],
			Attributes: data["attributes"],
		}
	case "StmtCatch":
		node = &ast.CatchStmt{
			Types:      data["types"],
			Var:        data["var"],
			Stmts:      data["stmts"],
			Attributes: data["attributes"],
		}
	case "StmtClass":
		node = &ast.ClassStmt{
			Flags:          data["flags"],
			Extends:        data["extends"],
			Implements:     data["implements"],
			Name:           data["name"],
			Stmts:          data["stmts"],
			AttrGroups:     data["attrGroups"],
			NamespacedName: data["namespacedName"],
			Attributes:     data["attributes"],
		}
	case "StmtClassConst":
		node = &ast.ClassConstStmt{
			Flags:      data["flags"],
			Consts:     data["consts"],
			AttrGroups: data["attrGroups"],
			Attributes: data["attributes"],
		}
	case "StmtClassMethod":
		node = &ast.ClassMethodStmt{
			Flags:      data["flags"],
			ByRef:      data["byRef"],
			Name:       data["name"],
			Params:     data["params"],
			ReturnType: data["returnType"],
			Stmts:      data["stmts"],
			AttrGroups: data["attrGroups"],
			MagicNames: data["magicNames"],
			Attributes: data["attributes"],
		}
	case "StmtConst":
		node = &ast.ConstStmt{
			Consts:     data["consts"],
			Attributes: data["attributes"],
		}
	case "StmtContinue":
		node = &ast.ContinueStmt{
			Num:        data["num"],
			Attributes: data["attributes"],
		}
	case "StmtDeclare":
		node = &ast.DeclareStmt{
			Declares:   data["declares"],
			Stmts:      data["stmts"],
			Attributes: data["attributes"],
		}
	case "StmtDeclareDeclare":
		node = &ast.DeclareDeclareStmt{
			Key:        data["key"],
			Value:      data["value"],
			Attributes: data["attributes"],
		}
	case "StmtDo":
		node = &ast.DoStmt{
			Stmts:      data["stmts"],
			Cond:       data["cond"],
			Attributes: data["attributes"],
		}
	case "StmtEcho":
		node = &ast.EchoStmt{
			Exprs:      data["exprs"],
			Attributes: data["attributes"],
		}
	case "StmtElse":
		node = &ast.ElseStmt{
			Stmts:      data["stmts"],
			Attributes: data["attributes"],
		}
	case "StmtElseIf":
		node = &ast.ElseIfStmt{
			Cond:       data["cond"],
			Stmts:      data["stmts"],
			Attributes: data["attributes"],
		}
	case "StmtEnum":
		node = &ast.EnumStmt{
			ScalarType:     data["scalarType"],
			Implements:     data["implements"],
			Name:           data["name"],
			Stmts:          data["stmts"],
			AttrGroups:     data["attrGroups"],
			NamespacedName: data["namespacedName"],
			Attributes:     data["attributes"],
		}
	case "StmtEnumCase":
		node = &ast.EnumCaseStmt{
			Name:       data["name"],
			Expr:       data["expr"],
			AttrGroups: data["attrGroups"],
			Attributes: data["attributes"],
		}
	case "StmtExpression":
		node = &ast.ExpressionStmt{
			Expr:       data["expr"],
			Attributes: data["attributes"],
		}
	case "StmtFinally":
		node = &ast.FinallyStmt{
			Stmts:      data["stmts"],
			Attributes: data["attributes"],
		}
	case "StmtFor":
		node = &ast.ForStmt{
			Init:       data["init"],
			Cond:       data["cond"],
			Loop:       data["loop"],
			Stmts:      data["stmts"],
			Attributes: data["attributes"],
		}
	case "StmtForeach":
		node = &ast.ForeachStmt{
			Expr:       data["expr"],
			KeyVar:     data["keyVar"],
			ByRef:      data["byRef"],
			ValueVar:   data["valueVar"],
			Stmts:      data["stmts"],
			Attributes: data["attributes"],
		}
	case "StmtFunction":
		node = &ast.FunctionStmt{
			ByRef:          data["byRef"],
			Name:           data["name"],
			Params:         data["params"],
			ReturnType:     data["returnType"],
			Stmts:          data["stmts"],
			AttrGroups:     data["attrGroups"],
			NamespacedName: data["namespacedName"],
			Attributes:     data["attributes"],
		}
	case "StmtGlobal":
		node = &ast.GlobalStmt{
			Vars:       data["vars"],
			Attributes: data["attributes"],
		}
	case "StmtGoto":
		node = &ast.GotoStmt{
			Name:       data["name"],
			Attributes: data["attributes"],
		}
	case "StmtGroupUse":
		node = &ast.GroupUseStmt{
			Type:       data["type"],
			Prefix:     data["prefix"],
			Uses:       data["uses"],
			Attributes: data["attributes"],
		}
	case "StmtHaltCompiler":
		node = &ast.HaltCompilerStmt{
			Remaining:  data["remaining"],
			Attributes: data["attributes"],
		}
	case "StmtIf":
		node = &ast.IfStmt{
			Cond:       data["cond"],
			Stmts:      data["stmts"],
			Elseifs:    data["elseifs"],
			Else:       data["else"],
			Attributes: data["attributes"],
		}
	case "StmtInlineHTML":
		node = &ast.InlineHTMLStmt{
			Value:      data["value"],
			Attributes: data["attributes"],
		}
	case "StmtInterface":
		node = &ast.InterfaceStmt{
			Extends:        data["extends"],
			Name:           data["name"],
			Stmts:          data["stmts"],
			AttrGroups:     data["attrGroups"],
			NamespacedName: data["namespacedName"],
			Attributes:     data["attributes"],
		}
	case "StmtLabel":
		node = &ast.LabelStmt{
			Name:       data["name"],
			Attributes: data["attributes"],
		}
	case "StmtNamespace":
		node = &ast.NamespaceStmt{
			Name:       data["name"],
			Stmts:      data["stmts"],
			Attributes: data["attributes"],
		}
	case "StmtNop":
		node = &ast.NopStmt{
			Attributes: data["attributes"],
		}
	case "StmtProperty":
		node = &ast.PropertyStmt{
			Flags:      data["flags"],
			Props:      data["props"],
			Type:       data["type"],
			AttrGroups: data["attrGroups"],
			Attributes: data["attributes"],
		}
	case "StmtPropertyProperty":
		node = &ast.PropertyPropertyStmt{
			Name:       data["name"],
			Default:    data["default"],
			Attributes: data["attributes"],
		}
	case "StmtReturn":
		node = &ast.ReturnStmt{
			Expr:       data["expr"],
			Attributes: data["attributes"],
		}
	case "StmtStatic":
		node = &ast.StaticStmt{
			Vars:       data["vars"],
			Attributes: data["attributes"],
		}
	case "StmtStaticVar":
		node = &ast.StaticVarStmt{
			Var:        data["var"],
			Default:    data["default"],
			Attributes: data["attributes"],
		}
	case "StmtSwitch":
		node = &ast.SwitchStmt{
			Cond:       data["cond"],
			Cases:      data["cases"],
			Attributes: data["attributes"],
		}
	case "StmtThrow":
		node = &ast.ThrowStmt{
			Expr:       data["expr"],
			Attributes: data["attributes"],
		}
	case "StmtTrait":
		node = &ast.TraitStmt{
			Name:           data["name"],
			Stmts:          data["stmts"],
			AttrGroups:     data["attrGroups"],
			NamespacedName: data["namespacedName"],
			Attributes:     data["attributes"],
		}
	case "StmtTraitUse":
		node = &ast.TraitUseStmt{
			Traits:      data["traits"],
			Adaptations: data["adaptations"],
			Attributes:  data["attributes"],
		}
	case "StmtTraitUseAdaptationAlias":
		node = &ast.TraitUseAdaptationAliasStmt{
			NewModifier: data["newModifier"],
			NewName:     data["newName"],
			Trait:       data["trait"],
			Method:      data["method"],
			Attributes:  data["attributes"],
		}
	case "StmtTraitUseAdaptationPrecedence":
		node = &ast.TraitUseAdaptationPrecedenceStmt{
			Insteadof:  data["insteadof"],
			Trait:      data["trait"],
			Method:     data["method"],
			Attributes: data["attributes"],
		}
	case "StmtTryCatch":
		node = &ast.TryCatchStmt{
			Stmts:      data["stmts"],
			Catches:    data["catches"],
			Finally:    data["finally"],
			Attributes: data["attributes"],
		}
	case "StmtUnset":
		node = &ast.UnsetStmt{
			Vars:       data["vars"],
			Attributes: data["attributes"],
		}
	case "StmtUse":
		node = &ast.UseStmt{
			Type:       data["type"],
			Uses:       data["uses"],
			Attributes: data["attributes"],
		}
	case "StmtUseUse":
		node = &ast.UseUseStmt{
			Type:       data["type"],
			Name:       data["name"],
			Alias:      data["alias"],
			Attributes: data["attributes"],
		}
	case "StmtWhile":
		node = &ast.WhileStmt{
			Cond:       data["cond"],
			Stmts:      data["stmts"],
			Attributes: data["attributes"],
		}
	case "UnionType":
		node = &ast.UnionType{
			Types:      data["types"],
			Attributes: data["attributes"],
		}
	case "VarLikeIdentifier":
		node = &ast.VarLikeIdentifier{
			Name:       data["name"],
			Attributes: data["attributes"],
		}
	case "VariadicPlaceholder":
		node = &ast.VariadicPlaceholder{
			Attributes: data["attributes"],
		}
	}

	return node, nil
}
