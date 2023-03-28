package phpparse

import (
	"gophp/php/ast"
	"gophp/php/token"
)

func decodeNode(data map[string]any) (node ast.Node, err error) {
	nodeType := data["nodeType"].(string)
	switch nodeType {
	case "Arg":
		node = &ast.Arg{
			Name:   asTypeOrNil[*ast.Identifier](data["name"]),
			Value:  data["value"].(ast.Expr),
			ByRef:  data["byRef"].(bool),
			Unpack: data["unpack"].(bool),
		}
	case "Attribute":
		node = &ast.Attribute{
			Name: data["name"].(*ast.Name),
			Args: asSlice[*ast.Arg](data["args"]),
		}
	case "AttributeGroup":
		node = &ast.AttributeGroup{
			Attrs: asSlice[*ast.Attribute](data["attrs"]),
		}
	case "Const":
		node = &ast.Const{
			Name:           data["name"].(*ast.Identifier),
			Value:          data["value"].(ast.Expr),
			NamespacedName: asTypeOrNil[*ast.Name](data["namespacedName"]),
		}
	case "ExprArray":
		node = &ast.ArrayExpr{
			Items: asSlice[*ast.ArrayItemExpr](data["items"]),
		}
	case "ExprArrayDimFetch":
		node = &ast.ArrayDimFetchExpr{
			Var: data["var"].(ast.Expr),
			Dim: asTypeOrNil[ast.Expr](data["dim"]),
		}
	case "ExprArrayItem":
		node = &ast.ArrayItemExpr{
			Key:    asTypeOrNil[ast.Expr](data["key"]),
			Value:  data["value"].(ast.Expr),
			ByRef:  data["byRef"].(bool),
			Unpack: data["unpack"].(bool),
		}
	case "ExprArrowFunction":
		node = &ast.ArrowFunctionExpr{
			Static:     data["static"].(bool),
			ByRef:      data["byRef"].(bool),
			Params:     asSlice[*ast.Param](data["params"]),
			ReturnType: data["returnType"],
			Expr:       data["expr"].(ast.Expr),
			AttrGroups: asSlice[*ast.AttributeGroup](data["attrGroups"]),
		}
	case "ExprAssign":
		node = &ast.AssignExpr{
			Var:  data["var"].(ast.Expr),
			Expr: data["expr"].(ast.Expr),
		}
	case "ExprAssignOpBitwiseAnd":
		node = &ast.AssignOpExpr{
			Op:   token.AndAssign,
			Var:  data["var"].(ast.Expr),
			Expr: data["expr"].(ast.Expr),
		}
	case "ExprAssignOpBitwiseOr":
		node = &ast.AssignOpExpr{
			Op:   token.OrAssign,
			Var:  data["var"].(ast.Expr),
			Expr: data["expr"].(ast.Expr),
		}
	case "ExprAssignOpBitwiseXor":
		node = &ast.AssignOpExpr{
			Op:   token.XorAssign,
			Var:  data["var"].(ast.Expr),
			Expr: data["expr"].(ast.Expr),
		}
	case "ExprAssignOpCoalesce":
		node = &ast.AssignOpExpr{
			Op:   token.CoalesceAssign,
			Var:  data["var"].(ast.Expr),
			Expr: data["expr"].(ast.Expr),
		}
	case "ExprAssignOpConcat":
		node = &ast.AssignOpExpr{
			Op:   token.ConcatAssign,
			Var:  data["var"].(ast.Expr),
			Expr: data["expr"].(ast.Expr),
		}
	case "ExprAssignOpDiv":
		node = &ast.AssignOpExpr{
			Op:   token.DivAssign,
			Var:  data["var"].(ast.Expr),
			Expr: data["expr"].(ast.Expr),
		}
	case "ExprAssignOpMinus":
		node = &ast.AssignOpExpr{
			Op:   token.SubAssign,
			Var:  data["var"].(ast.Expr),
			Expr: data["expr"].(ast.Expr),
		}
	case "ExprAssignOpMod":
		node = &ast.AssignOpExpr{
			Op:   token.ModAssign,
			Var:  data["var"].(ast.Expr),
			Expr: data["expr"].(ast.Expr),
		}
	case "ExprAssignOpMul":
		node = &ast.AssignOpExpr{
			Op:   token.MulAssign,
			Var:  data["var"].(ast.Expr),
			Expr: data["expr"].(ast.Expr),
		}
	case "ExprAssignOpPlus":
		node = &ast.AssignOpExpr{
			Op:   token.AddAssign,
			Var:  data["var"].(ast.Expr),
			Expr: data["expr"].(ast.Expr),
		}
	case "ExprAssignOpPow":
		node = &ast.AssignOpExpr{
			Op:   token.PowAssign,
			Var:  data["var"].(ast.Expr),
			Expr: data["expr"].(ast.Expr),
		}
	case "ExprAssignOpShiftLeft":
		node = &ast.AssignOpExpr{
			Op:   token.ShiftLeftAssign,
			Var:  data["var"].(ast.Expr),
			Expr: data["expr"].(ast.Expr),
		}
	case "ExprAssignOpShiftRight":
		node = &ast.AssignOpExpr{
			Op:   token.ShiftRightAssign,
			Var:  data["var"].(ast.Expr),
			Expr: data["expr"].(ast.Expr),
		}
	case "ExprAssignRef":
		node = &ast.AssignRefExpr{
			Var:  data["var"].(ast.Expr),
			Expr: data["expr"].(ast.Expr),
		}
	case "ExprBinaryOpBitwiseAnd":
		node = &ast.BinaryOpExpr{
			Op:    token.And,
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "ExprBinaryOpBitwiseOr":
		node = &ast.BinaryOpExpr{
			Op:    token.Or,
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "ExprBinaryOpBitwiseXor":
		node = &ast.BinaryOpExpr{
			Op:    token.Xor,
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "ExprBinaryOpBooleanAnd":
		node = &ast.BinaryOpExpr{
			Op:    token.BooleanAnd,
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "ExprBinaryOpBooleanOr":
		node = &ast.BinaryOpExpr{
			Op:    token.BooleanOr,
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "ExprBinaryOpCoalesce":
		node = &ast.BinaryOpExpr{
			Op:    token.Coalesce,
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "ExprBinaryOpConcat":
		node = &ast.BinaryOpExpr{
			Op:    token.Concat,
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "ExprBinaryOpDiv":
		node = &ast.BinaryOpExpr{
			Op:    token.Div,
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "ExprBinaryOpEqual":
		node = &ast.BinaryOpExpr{
			Op:    token.Equal,
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "ExprBinaryOpGreater":
		node = &ast.BinaryOpExpr{
			Op:    token.Greater,
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "ExprBinaryOpGreaterOrEqual":
		node = &ast.BinaryOpExpr{
			Op:    token.GreaterOrEqual,
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "ExprBinaryOpIdentical":
		node = &ast.BinaryOpExpr{
			Op:    token.Identical,
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "ExprBinaryOpLogicalAnd":
		node = &ast.BinaryOpExpr{
			Op:    token.LogicalAnd,
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "ExprBinaryOpLogicalOr":
		node = &ast.BinaryOpExpr{
			Op:    token.LogicalOr,
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "ExprBinaryOpLogicalXor":
		node = &ast.BinaryOpExpr{
			Op:    token.LogicalXor,
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "ExprBinaryOpMinus":
		node = &ast.BinaryOpExpr{
			Op:    token.Sub,
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "ExprBinaryOpMod":
		node = &ast.BinaryOpExpr{
			Op:    token.Mod,
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "ExprBinaryOpMul":
		node = &ast.BinaryOpExpr{
			Op:    token.Mul,
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "ExprBinaryOpNotEqual":
		node = &ast.BinaryOpExpr{
			Op:    token.NotEqual,
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "ExprBinaryOpNotIdentical":
		node = &ast.BinaryOpExpr{
			Op:    token.NotIdentical,
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "ExprBinaryOpPlus":
		node = &ast.BinaryOpExpr{
			Op:    token.Add,
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "ExprBinaryOpPow":
		node = &ast.BinaryOpExpr{
			Op:    token.Pow,
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "ExprBinaryOpShiftLeft":
		node = &ast.BinaryOpExpr{
			Op:    token.ShiftLeft,
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "ExprBinaryOpShiftRight":
		node = &ast.BinaryOpExpr{
			Op:    token.ShiftRight,
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "ExprBinaryOpSmaller":
		node = &ast.BinaryOpExpr{
			Op:    token.Smaller,
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "ExprBinaryOpSmallerOrEqual":
		node = &ast.BinaryOpExpr{
			Op:    token.SmallerOrEqual,
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "ExprBinaryOpSpaceship":
		node = &ast.BinaryOpExpr{
			Op:    token.Spaceship,
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "ExprBitwiseNot":
		node = &ast.BitwiseNotExpr{
			Expr: data["expr"].(ast.Expr),
		}
	case "ExprBooleanNot":
		node = &ast.BooleanNotExpr{
			Expr: data["expr"].(ast.Expr),
		}
	case "ExprCastArray":
		node = &ast.CastExpr{
			Op:   token.ArrayCast,
			Expr: data["expr"].(ast.Expr),
		}
	case "ExprCastBool":
		node = &ast.CastExpr{
			Op:   token.BoolCast,
			Expr: data["expr"].(ast.Expr),
		}
	case "ExprCastDouble":
		node = &ast.CastExpr{
			Op:   token.DoubleCast,
			Expr: data["expr"].(ast.Expr),
		}
	case "ExprCastInt":
		node = &ast.CastExpr{
			Op:   token.IntCast,
			Expr: data["expr"].(ast.Expr),
		}
	case "ExprCastObject":
		node = &ast.CastExpr{
			Op:   token.ObjectCast,
			Expr: data["expr"].(ast.Expr),
		}
	case "ExprCastString":
		node = &ast.CastExpr{
			Op:   token.StringCast,
			Expr: data["expr"].(ast.Expr),
		}
	case "ExprCastUnset":
		node = &ast.CastExpr{
			Op:   token.UnsetCast,
			Expr: data["expr"].(ast.Expr),
		}
	case "ExprClassConstFetch":
		node = &ast.ClassConstFetchExpr{
			Class: data["class"],
			Name:  data["name"],
		}
	case "ExprClone":
		node = &ast.CloneExpr{
			Expr: data["expr"].(ast.Expr),
		}
	case "ExprClosure":
		node = &ast.ClosureExpr{
			Static:     data["static"].(bool),
			ByRef:      data["byRef"].(bool),
			Params:     asSlice[*ast.Param](data["params"]),
			Uses:       asSlice[*ast.ClosureUseExpr](data["uses"]),
			ReturnType: data["returnType"],
			Stmts:      asSlice[ast.Stmt](data["stmts"]),
			AttrGroups: asSlice[*ast.AttributeGroup](data["attrGroups"]),
		}
	case "ExprClosureUse":
		node = &ast.ClosureUseExpr{
			Var:   data["var"].(*ast.VariableExpr),
			ByRef: data["byRef"].(bool),
		}
	case "ExprConstFetch":
		node = &ast.ConstFetchExpr{
			Name: data["name"].(*ast.Name),
		}
	case "ExprEmpty":
		node = &ast.EmptyExpr{
			Expr: data["expr"].(ast.Expr),
		}
	case "ExprError":
		node = &ast.ErrorExpr{}
	case "ExprErrorSuppress":
		node = &ast.ErrorSuppressExpr{
			Expr: data["expr"].(ast.Expr),
		}
	case "ExprEval":
		node = &ast.EvalExpr{
			Expr: data["expr"].(ast.Expr),
		}
	case "ExprExit":
		node = &ast.ExitExpr{
			Expr: asTypeOrNil[ast.Expr](data["expr"]),
		}
	case "ExprFuncCall":
		node = &ast.FuncCallExpr{
			Name: data["name"],
			Args: asSlice[any](data["args"]),
		}
	case "ExprInclude":
		node = &ast.IncludeExpr{
			Expr: data["expr"].(ast.Expr),
			Type: asInt(data["type"]),
		}
	case "ExprInstanceof":
		node = &ast.InstanceofExpr{
			Expr:  data["expr"].(ast.Expr),
			Class: data["class"],
		}
	case "ExprIsset":
		node = &ast.IssetExpr{
			Vars: asSlice[ast.Expr](data["vars"]),
		}
	case "ExprList":
		node = &ast.ListExpr{
			Items: asSlice[*ast.ArrayItemExpr](data["items"]),
		}
	case "ExprMatch":
		node = &ast.MatchExpr{
			Cond: data["cond"].(ast.Expr),
			Arms: asSlice[*ast.MatchArm](data["arms"]),
		}
	case "ExprMethodCall":
		node = &ast.MethodCallExpr{
			Var:  data["var"].(ast.Expr),
			Name: data["name"],
			Args: asSlice[any](data["args"]),
		}
	case "ExprNew":
		node = &ast.NewExpr{
			Class: data["class"],
			Args:  asSlice[any](data["args"]),
		}
	case "ExprNullsafeMethodCall":
		node = &ast.NullsafeMethodCallExpr{
			Var:  data["var"].(ast.Expr),
			Name: data["name"],
			Args: asSlice[any](data["args"]),
		}
	case "ExprNullsafePropertyFetch":
		node = &ast.NullsafePropertyFetchExpr{
			Var:  data["var"].(ast.Expr),
			Name: data["name"],
		}
	case "ExprPostDec":
		node = &ast.PostDecExpr{
			Var: data["var"].(ast.Expr),
		}
	case "ExprPostInc":
		node = &ast.PostIncExpr{
			Var: data["var"].(ast.Expr),
		}
	case "ExprPreDec":
		node = &ast.PreDecExpr{
			Var: data["var"].(ast.Expr),
		}
	case "ExprPreInc":
		node = &ast.PreIncExpr{
			Var: data["var"].(ast.Expr),
		}
	case "ExprPrint":
		node = &ast.PrintExpr{
			Expr: data["expr"].(ast.Expr),
		}
	case "ExprPropertyFetch":
		node = &ast.PropertyFetchExpr{
			Var:  data["var"].(ast.Expr),
			Name: data["name"],
		}
	case "ExprShellExec":
		node = &ast.ShellExecExpr{
			Parts: asSlice[any](data["parts"]),
		}
	case "ExprStaticCall":
		node = &ast.StaticCallExpr{
			Class: data["class"],
			Name:  data["name"],
			Args:  asSlice[any](data["args"]),
		}
	case "ExprStaticPropertyFetch":
		node = &ast.StaticPropertyFetchExpr{
			Class: data["class"],
			Name:  data["name"],
		}
	case "ExprTernary":
		node = &ast.TernaryExpr{
			Cond: data["cond"].(ast.Expr),
			If:   asTypeOrNil[ast.Expr](data["if"]),
			Else: data["else"].(ast.Expr),
		}
	case "ExprThrow":
		node = &ast.ThrowExpr{
			Expr: data["expr"].(ast.Expr),
		}
	case "ExprUnaryMinus":
		node = &ast.UnaryMinusExpr{
			Expr: data["expr"].(ast.Expr),
		}
	case "ExprUnaryPlus":
		node = &ast.UnaryPlusExpr{
			Expr: data["expr"].(ast.Expr),
		}
	case "ExprVariable":
		node = &ast.VariableExpr{
			Name: data["name"],
		}
	case "ExprYield":
		node = &ast.YieldExpr{
			Key:   asTypeOrNil[ast.Expr](data["key"]),
			Value: asTypeOrNil[ast.Expr](data["value"]),
		}
	case "ExprYieldFrom":
		node = &ast.YieldFromExpr{
			Expr: data["expr"].(ast.Expr),
		}
	case "Identifier":
		node = &ast.Identifier{
			Name: data["name"].(string),
		}
	case "IntersectionType":
		node = &ast.IntersectionType{
			Types: asSlice[any](data["types"]),
		}
	case "MatchArm":
		node = &ast.MatchArm{
			Conds: asSlice[ast.Expr](data["conds"]),
			Body:  data["body"].(ast.Expr),
		}
	case "Name":
		node = &ast.Name{
			Parts: asSlice[string](data["parts"]),
		}
	case "NameFullyQualified":
		node = &ast.NameFullyQualified{
			Parts: asSlice[string](data["parts"]),
		}
	case "NameRelative":
		node = &ast.NameRelative{
			Parts: asSlice[string](data["parts"]),
		}
	case "NullableType":
		node = &ast.NullableType{
			Type: data["type"],
		}
	case "Param":
		node = &ast.Param{
			Type:       data["type"],
			ByRef:      data["byRef"].(bool),
			Variadic:   data["variadic"].(bool),
			Var:        data["var"],
			Default:    asTypeOrNil[ast.Expr](data["default"]),
			Flags:      asInt(data["flags"]),
			AttrGroups: asSlice[*ast.AttributeGroup](data["attrGroups"]),
		}
	case "ScalarDNumber":
		node = &ast.DNumberScalar{
			Value: asFloat(data["value"]),
		}
	case "ScalarEncapsed":
		node = &ast.EncapsedScalar{
			Parts: asSlice[ast.Expr](data["parts"]),
		}
	case "ScalarEncapsedStringPart":
		node = &ast.EncapsedStringPartScalar{
			Value: data["value"].(string),
		}
	case "ScalarLNumber":
		node = &ast.LNumberScalar{
			Value: asInt(data["value"]),
		}
	case "ScalarMagicConstClass":
		node = &ast.MagicConstScalar{Op: token.ClassConst}
	case "ScalarMagicConstDir":
		node = &ast.MagicConstScalar{Op: token.DirConst}
	case "ScalarMagicConstFile":
		node = &ast.MagicConstScalar{Op: token.FileConst}
	case "ScalarMagicConstFunction":
		node = &ast.MagicConstScalar{Op: token.FunctionConst}
	case "ScalarMagicConstLine":
		node = &ast.MagicConstScalar{Op: token.LineConst}
	case "ScalarMagicConstMethod":
		node = &ast.MagicConstScalar{Op: token.MethodConst}
	case "ScalarMagicConstNamespace":
		node = &ast.MagicConstScalar{Op: token.NamespaceConst}
	case "ScalarMagicConstTrait":
		node = &ast.MagicConstScalar{Op: token.TraitConst}
	case "ScalarString":
		node = &ast.StringScalar{
			Value: data["value"].(string),
		}
	case "StmtBreak":
		node = &ast.BreakStmt{
			Num: asTypeOrNil[ast.Expr](data["num"]),
		}
	case "StmtCase":
		node = &ast.CaseStmt{
			Cond:  asTypeOrNil[ast.Expr](data["cond"]),
			Stmts: asSlice[ast.Stmt](data["stmts"]),
		}
	case "StmtCatch":
		node = &ast.CatchStmt{
			Types: asSlice[*ast.Name](data["types"]),
			Var:   asTypeOrNil[*ast.VariableExpr](data["var"]),
			Stmts: asSlice[ast.Stmt](data["stmts"]),
		}
	case "StmtClass":
		node = &ast.ClassStmt{
			Flags:          asInt(data["flags"]),
			Extends:        asTypeOrNil[*ast.Name](data["extends"]),
			Implements:     asSlice[*ast.Name](data["implements"]),
			Name:           asTypeOrNil[*ast.Identifier](data["name"]),
			Stmts:          asSlice[ast.Stmt](data["stmts"]),
			AttrGroups:     asSlice[*ast.AttributeGroup](data["attrGroups"]),
			NamespacedName: asTypeOrNil[*ast.Name](data["namespacedName"]),
		}
	case "StmtClassConst":
		node = &ast.ClassConstStmt{
			Flags:      asInt(data["flags"]),
			Consts:     asSlice[*ast.Const](data["consts"]),
			AttrGroups: asSlice[*ast.AttributeGroup](data["attrGroups"]),
		}
	case "StmtClassMethod":
		node = &ast.ClassMethodStmt{
			Flags:      asInt(data["flags"]),
			ByRef:      data["byRef"].(bool),
			Name:       data["name"].(*ast.Identifier),
			Params:     asSlice[*ast.Param](data["params"]),
			ReturnType: data["returnType"],
			Stmts:      asSlice[ast.Stmt](data["stmts"]),
			AttrGroups: asSlice[*ast.AttributeGroup](data["attrGroups"]),
		}
	case "StmtConst":
		node = &ast.ConstStmt{
			Consts: asSlice[*ast.Const](data["consts"]),
		}
	case "StmtContinue":
		node = &ast.ContinueStmt{
			Num: asTypeOrNil[ast.Expr](data["num"]),
		}
	case "StmtDeclare":
		node = &ast.DeclareStmt{
			Declares: asSlice[*ast.DeclareDeclareStmt](data["declares"]),
			Stmts:    asSlice[ast.Stmt](data["stmts"]),
		}
	case "StmtDeclareDeclare":
		node = &ast.DeclareDeclareStmt{
			Key:   data["key"].(*ast.Identifier),
			Value: data["value"].(ast.Expr),
		}
	case "StmtDo":
		node = &ast.DoStmt{
			Stmts: asSlice[ast.Stmt](data["stmts"]),
			Cond:  data["cond"].(ast.Expr),
		}
	case "StmtEcho":
		node = &ast.EchoStmt{
			Exprs: asSlice[ast.Expr](data["exprs"]),
		}
	case "StmtElse":
		node = &ast.ElseStmt{
			Stmts: asSlice[ast.Stmt](data["stmts"]),
		}
	case "StmtElseIf":
		node = &ast.ElseIfStmt{
			Cond:  data["cond"].(ast.Expr),
			Stmts: asSlice[ast.Stmt](data["stmts"]),
		}
	case "StmtEnum":
		node = &ast.EnumStmt{
			ScalarType:     asTypeOrNil[*ast.Identifier](data["scalarType"]),
			Implements:     asSlice[*ast.Name](data["implements"]),
			Name:           asTypeOrNil[*ast.Identifier](data["name"]),
			Stmts:          asSlice[ast.Stmt](data["stmts"]),
			AttrGroups:     asSlice[*ast.AttributeGroup](data["attrGroups"]),
			NamespacedName: asTypeOrNil[*ast.Name](data["namespacedName"]),
		}
	case "StmtEnumCase":
		node = &ast.EnumCaseStmt{
			Name:       data["name"].(*ast.Identifier),
			Expr:       asTypeOrNil[ast.Expr](data["expr"]),
			AttrGroups: asSlice[*ast.AttributeGroup](data["attrGroups"]),
		}
	case "StmtExpression":
		node = &ast.ExpressionStmt{
			Expr: data["expr"].(ast.Expr),
		}
	case "StmtFinally":
		node = &ast.FinallyStmt{
			Stmts: asSlice[ast.Stmt](data["stmts"]),
		}
	case "StmtFor":
		node = &ast.ForStmt{
			Init:  asSlice[ast.Expr](data["init"]),
			Cond:  asSlice[ast.Expr](data["cond"]),
			Loop:  asSlice[ast.Expr](data["loop"]),
			Stmts: asSlice[ast.Stmt](data["stmts"]),
		}
	case "StmtForeach":
		node = &ast.ForeachStmt{
			Expr:     data["expr"].(ast.Expr),
			KeyVar:   asTypeOrNil[ast.Expr](data["keyVar"]),
			ByRef:    data["byRef"].(bool),
			ValueVar: data["valueVar"].(ast.Expr),
			Stmts:    asSlice[ast.Stmt](data["stmts"]),
		}
	case "StmtFunction":
		node = &ast.FunctionStmt{
			ByRef:          data["byRef"].(bool),
			Name:           data["name"].(*ast.Identifier),
			Params:         asSlice[*ast.Param](data["params"]),
			ReturnType:     data["returnType"],
			Stmts:          asSlice[ast.Stmt](data["stmts"]),
			AttrGroups:     asSlice[*ast.AttributeGroup](data["attrGroups"]),
			NamespacedName: asTypeOrNil[*ast.Name](data["namespacedName"]),
		}
	case "StmtGlobal":
		node = &ast.GlobalStmt{
			Vars: asSlice[ast.Expr](data["vars"]),
		}
	case "StmtGoto":
		node = &ast.GotoStmt{
			Name: data["name"].(*ast.Identifier),
		}
	case "StmtGroupUse":
		node = &ast.GroupUseStmt{
			Type:   asInt(data["type"]),
			Prefix: data["prefix"].(*ast.Name),
			Uses:   asSlice[*ast.UseUseStmt](data["uses"]),
		}
	case "StmtHaltCompiler":
		node = &ast.HaltCompilerStmt{
			Remaining: data["remaining"].(string),
		}
	case "StmtIf":
		node = &ast.IfStmt{
			Cond:    data["cond"].(ast.Expr),
			Stmts:   asSlice[ast.Stmt](data["stmts"]),
			Elseifs: asSlice[*ast.ElseIfStmt](data["elseifs"]),
			Else:    asTypeOrNil[*ast.ElseStmt](data["else"]),
		}
	case "StmtInlineHTML":
		node = &ast.InlineHTMLStmt{
			Value: data["value"].(string),
		}
	case "StmtInterface":
		node = &ast.InterfaceStmt{
			Extends:        asSlice[*ast.Name](data["extends"]),
			Name:           asTypeOrNil[*ast.Identifier](data["name"]),
			Stmts:          asSlice[ast.Stmt](data["stmts"]),
			AttrGroups:     asSlice[*ast.AttributeGroup](data["attrGroups"]),
			NamespacedName: asTypeOrNil[*ast.Name](data["namespacedName"]),
		}
	case "StmtLabel":
		node = &ast.LabelStmt{
			Name: data["name"].(*ast.Identifier),
		}
	case "StmtNamespace":
		node = &ast.NamespaceStmt{
			Name:  asTypeOrNil[*ast.Name](data["name"]),
			Stmts: asSlice[ast.Stmt](data["stmts"]),
		}
	case "StmtNop":
		node = &ast.NopStmt{}
	case "StmtProperty":
		node = &ast.PropertyStmt{
			Flags:      asInt(data["flags"]),
			Props:      asSlice[*ast.PropertyPropertyStmt](data["props"]),
			Type:       data["type"],
			AttrGroups: asSlice[*ast.AttributeGroup](data["attrGroups"]),
		}
	case "StmtPropertyProperty":
		node = &ast.PropertyPropertyStmt{
			Name:    data["name"].(*ast.VarLikeIdentifier),
			Default: asTypeOrNil[ast.Expr](data["default"]),
		}
	case "StmtReturn":
		node = &ast.ReturnStmt{
			Expr: asTypeOrNil[ast.Expr](data["expr"]),
		}
	case "StmtStatic":
		node = &ast.StaticStmt{
			Vars: asSlice[*ast.StaticVarStmt](data["vars"]),
		}
	case "StmtStaticVar":
		node = &ast.StaticVarStmt{
			Var:     data["var"].(*ast.VariableExpr),
			Default: asTypeOrNil[ast.Expr](data["default"]),
		}
	case "StmtSwitch":
		node = &ast.SwitchStmt{
			Cond:  data["cond"].(ast.Expr),
			Cases: asSlice[*ast.CaseStmt](data["cases"]),
		}
	case "StmtThrow":
		node = &ast.ThrowStmt{
			Expr: data["expr"].(ast.Expr),
		}
	case "StmtTrait":
		node = &ast.TraitStmt{
			Name:           asTypeOrNil[*ast.Identifier](data["name"]),
			Stmts:          asSlice[ast.Stmt](data["stmts"]),
			AttrGroups:     asSlice[*ast.AttributeGroup](data["attrGroups"]),
			NamespacedName: asTypeOrNil[*ast.Name](data["namespacedName"]),
		}
	case "StmtTraitUse":
		node = &ast.TraitUseStmt{
			Traits:      asSlice[*ast.Name](data["traits"]),
			Adaptations: asSlice[*ast.TraitUseAdaptationStmt](data["adaptations"]),
		}
	case "StmtTraitUseAdaptationAlias":
		node = &ast.TraitUseAdaptationAliasStmt{
			NewModifier: asInt(data["newModifier"]),
			NewName:     asTypeOrNil[*ast.Identifier](data["newName"]),
			Trait:       asTypeOrNil[*ast.Name](data["trait"]),
			Method:      data["method"].(*ast.Identifier),
		}
	case "StmtTraitUseAdaptationPrecedence":
		node = &ast.TraitUseAdaptationPrecedenceStmt{
			Insteadof: asSlice[*ast.Name](data["insteadof"]),
			Trait:     asTypeOrNil[*ast.Name](data["trait"]),
			Method:    data["method"].(*ast.Identifier),
		}
	case "StmtTryCatch":
		node = &ast.TryCatchStmt{
			Stmts:   asSlice[ast.Stmt](data["stmts"]),
			Catches: asSlice[*ast.CatchStmt](data["catches"]),
			Finally: asTypeOrNil[*ast.FinallyStmt](data["finally"]),
		}
	case "StmtUnset":
		node = &ast.UnsetStmt{
			Vars: asSlice[ast.Expr](data["vars"]),
		}
	case "StmtUse":
		node = &ast.UseStmt{
			Type: asInt(data["type"]),
			Uses: asSlice[*ast.UseUseStmt](data["uses"]),
		}
	case "StmtUseUse":
		node = &ast.UseUseStmt{
			Type:  asInt(data["type"]),
			Name:  data["name"].(*ast.Name),
			Alias: asTypeOrNil[*ast.Identifier](data["alias"]),
		}
	case "StmtWhile":
		node = &ast.WhileStmt{
			Cond:  data["cond"].(ast.Expr),
			Stmts: asSlice[ast.Stmt](data["stmts"]),
		}
	case "UnionType":
		node = &ast.UnionType{
			Types: asSlice[any](data["types"]),
		}
	case "VarLikeIdentifier":
		node = &ast.VarLikeIdentifier{
			Name: data["name"].(string),
		}
	case "VariadicPlaceholder":
		node = &ast.VariadicPlaceholder{}
	}

	return node, nil
}
