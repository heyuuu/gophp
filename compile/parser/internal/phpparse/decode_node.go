package phpparse

import (
	"fmt"
	"github.com/heyuuu/gophp/compile/ast"
	"github.com/heyuuu/gophp/kits/slicekit"
)

func decodeNode(data map[string]any) (node ast.Node, err error) {
	nodeType := data["nodeType"].(string)
	switch nodeType {
	case "Arg":
		name := asTypeOrNil[*ast.Ident](data["name"])
		if name != nil {
			err = unsupported("unsupported high version php feature: php8.0 named arguments")
			break
		}

		byRef := data["byRef"].(bool)
		if byRef {
			err = unsupported("lower version php feature: Call-time pass-by-reference has been removed in PHP 5.4")
			break
		}

		node = &ast.Arg{
			Value:  data["value"].(ast.Expr),
			Unpack: data["unpack"].(bool),
		}
	case "Const":
		node = &ast.ConstStmt{
			Name:           data["name"].(*ast.Ident),
			Value:          data["value"].(ast.Expr),
			NamespacedName: asTypeOrNil[*ast.Name](data["namespacedName"]),
		}
	case "ArrayExpr":
		node = &ast.ArrayExpr{
			Items: asSlice[*ast.ArrayItemExpr](data["items"]),
		}
	case "ArrayDimFetchExpr":
		node = &ast.IndexExpr{
			Var: data["var"].(ast.Expr),
			Dim: asTypeOrNil[ast.Expr](data["dim"]),
		}
	case "ArrayItemExpr":
		node = &ast.ArrayItemExpr{
			Key:    asTypeOrNil[ast.Expr](data["key"]),
			Value:  data["value"].(ast.Expr),
			ByRef:  data["byRef"].(bool),
			Unpack: data["unpack"].(bool),
		}
	case "ArrowFunctionExpr":
		node = &ast.ArrowFunctionExpr{
			Static:     data["static"].(bool),
			ByRef:      data["byRef"].(bool),
			Params:     asSlice[*ast.Param](data["params"]),
			ReturnType: asTypeHint(data["returnType"]),
			Expr:       data["expr"].(ast.Expr),
		}
	case "AssignExpr":
		node = &ast.AssignExpr{
			Var:  data["var"].(ast.Expr),
			Expr: data["expr"].(ast.Expr),
		}
	case "AssignOpBitwiseAndExpr":
		node = &ast.AssignOpExpr{
			Op:   ast.AssignOpBitwiseAnd,
			Var:  data["var"].(ast.Expr),
			Expr: data["expr"].(ast.Expr),
		}
	case "AssignOpBitwiseOrExpr":
		node = &ast.AssignOpExpr{
			Op:   ast.AssignOpBitwiseOr,
			Var:  data["var"].(ast.Expr),
			Expr: data["expr"].(ast.Expr),
		}
	case "AssignOpBitwiseXorExpr":
		node = &ast.AssignOpExpr{
			Op:   ast.AssignOpBitwiseXor,
			Var:  data["var"].(ast.Expr),
			Expr: data["expr"].(ast.Expr),
		}
	case "AssignOpCoalesceExpr":
		node = &ast.AssignOpExpr{
			Op:   ast.AssignOpCoalesce,
			Var:  data["var"].(ast.Expr),
			Expr: data["expr"].(ast.Expr),
		}
	case "AssignOpConcatExpr":
		node = &ast.AssignOpExpr{
			Op:   ast.AssignOpConcat,
			Var:  data["var"].(ast.Expr),
			Expr: data["expr"].(ast.Expr),
		}
	case "AssignOpDivExpr":
		node = &ast.AssignOpExpr{
			Op:   ast.AssignOpDiv,
			Var:  data["var"].(ast.Expr),
			Expr: data["expr"].(ast.Expr),
		}
	case "AssignOpMinusExpr":
		node = &ast.AssignOpExpr{
			Op:   ast.AssignOpMinus,
			Var:  data["var"].(ast.Expr),
			Expr: data["expr"].(ast.Expr),
		}
	case "AssignOpModExpr":
		node = &ast.AssignOpExpr{
			Op:   ast.AssignOpMod,
			Var:  data["var"].(ast.Expr),
			Expr: data["expr"].(ast.Expr),
		}
	case "AssignOpMulExpr":
		node = &ast.AssignOpExpr{
			Op:   ast.AssignOpMul,
			Var:  data["var"].(ast.Expr),
			Expr: data["expr"].(ast.Expr),
		}
	case "AssignOpPlusExpr":
		node = &ast.AssignOpExpr{
			Op:   ast.AssignOpPlus,
			Var:  data["var"].(ast.Expr),
			Expr: data["expr"].(ast.Expr),
		}
	case "AssignOpPowExpr":
		node = &ast.AssignOpExpr{
			Op:   ast.AssignOpPow,
			Var:  data["var"].(ast.Expr),
			Expr: data["expr"].(ast.Expr),
		}
	case "AssignOpShiftLeftExpr":
		node = &ast.AssignOpExpr{
			Op:   ast.AssignOpShiftLeft,
			Var:  data["var"].(ast.Expr),
			Expr: data["expr"].(ast.Expr),
		}
	case "AssignOpShiftRightExpr":
		node = &ast.AssignOpExpr{
			Op:   ast.AssignOpShiftRight,
			Var:  data["var"].(ast.Expr),
			Expr: data["expr"].(ast.Expr),
		}
	case "AssignRefExpr":
		node = &ast.AssignRefExpr{
			Var:  data["var"].(ast.Expr),
			Expr: data["expr"].(ast.Expr),
		}
	case "BinaryOpBitwiseAndExpr":
		node = &ast.BinaryOpExpr{
			Op:    ast.BinaryOpBitwiseAnd,
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "BinaryOpBitwiseOrExpr":
		node = &ast.BinaryOpExpr{
			Op:    ast.BinaryOpBitwiseOr,
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "BinaryOpBitwiseXorExpr":
		node = &ast.BinaryOpExpr{
			Op:    ast.BinaryOpBitwiseXor,
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "BinaryOpBooleanAndExpr":
		node = &ast.BinaryOpExpr{
			Op:    ast.BinaryOpBooleanAnd,
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "BinaryOpBooleanOrExpr":
		node = &ast.BinaryOpExpr{
			Op:    ast.BinaryOpBooleanOr,
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "BinaryOpCoalesceExpr":
		node = &ast.BinaryOpExpr{
			Op:    ast.BinaryOpCoalesce,
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "BinaryOpConcatExpr":
		node = &ast.BinaryOpExpr{
			Op:    ast.BinaryOpConcat,
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "BinaryOpDivExpr":
		node = &ast.BinaryOpExpr{
			Op:    ast.BinaryOpDiv,
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "BinaryOpEqualExpr":
		node = &ast.BinaryOpExpr{
			Op:    ast.BinaryOpEqual,
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "BinaryOpGreaterExpr":
		node = &ast.BinaryOpExpr{
			Op:    ast.BinaryOpGreater,
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "BinaryOpGreaterOrEqualExpr":
		node = &ast.BinaryOpExpr{
			Op:    ast.BinaryOpGreaterOrEqual,
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "BinaryOpIdenticalExpr":
		node = &ast.BinaryOpExpr{
			Op:    ast.BinaryOpIdentical,
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "BinaryOpLogicalAndExpr":
		node = &ast.BinaryOpExpr{
			Op:    ast.BinaryOpBooleanAnd, // 使用 `&&` 代替 `and`
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "BinaryOpLogicalOrExpr":
		node = &ast.BinaryOpExpr{
			Op:    ast.BinaryOpBooleanOr, // 使用 `||` 代替 `or`
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "BinaryOpLogicalXorExpr":
		node = &ast.BinaryOpExpr{
			Op:    ast.BinaryOpBooleanXor,
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "BinaryOpMinusExpr":
		node = &ast.BinaryOpExpr{
			Op:    ast.BinaryOpMinus,
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "BinaryOpModExpr":
		node = &ast.BinaryOpExpr{
			Op:    ast.BinaryOpMod,
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "BinaryOpMulExpr":
		node = &ast.BinaryOpExpr{
			Op:    ast.BinaryOpMul,
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "BinaryOpNotEqualExpr":
		node = &ast.BinaryOpExpr{
			Op:    ast.BinaryOpNotEqual,
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "BinaryOpNotIdenticalExpr":
		node = &ast.BinaryOpExpr{
			Op:    ast.BinaryOpNotIdentical,
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "BinaryOpPlusExpr":
		node = &ast.BinaryOpExpr{
			Op:    ast.BinaryOpPlus,
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "BinaryOpPowExpr":
		node = &ast.BinaryOpExpr{
			Op:    ast.BinaryOpPow,
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "BinaryOpShiftLeftExpr":
		node = &ast.BinaryOpExpr{
			Op:    ast.BinaryOpShiftLeft,
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "BinaryOpShiftRightExpr":
		node = &ast.BinaryOpExpr{
			Op:    ast.BinaryOpShiftRight,
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "BinaryOpSmallerExpr":
		node = &ast.BinaryOpExpr{
			Op:    ast.BinaryOpSmaller,
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "BinaryOpSmallerOrEqualExpr":
		node = &ast.BinaryOpExpr{
			Op:    ast.BinaryOpSmallerOrEqual,
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "BinaryOpSpaceshipExpr":
		node = &ast.BinaryOpExpr{
			Op:    ast.BinaryOpSpaceship,
			Left:  data["left"].(ast.Expr),
			Right: data["right"].(ast.Expr),
		}
	case "BitwiseNotExpr":
		node = &ast.UnaryExpr{
			Op:  ast.UnaryOpBitwiseNot,
			Var: data["expr"].(ast.Expr),
		}
	case "BooleanNotExpr":
		node = &ast.UnaryExpr{
			Op:  ast.UnaryOpBooleanNot,
			Var: data["expr"].(ast.Expr),
		}
	case "CastArrayExpr":
		node = &ast.CastExpr{
			Kind: ast.CastArray,
			Expr: data["expr"].(ast.Expr),
		}
	case "CastBoolExpr":
		node = &ast.CastExpr{
			Kind: ast.CastBool,
			Expr: data["expr"].(ast.Expr),
		}
	case "CastDoubleExpr":
		node = &ast.CastExpr{
			Kind: ast.CastDouble,
			Expr: data["expr"].(ast.Expr),
		}
	case "CastIntExpr":
		node = &ast.CastExpr{
			Kind: ast.CastInt,
			Expr: data["expr"].(ast.Expr),
		}
	case "CastObjectExpr":
		node = &ast.CastExpr{
			Kind: ast.CastObject,
			Expr: data["expr"].(ast.Expr),
		}
	case "CastStringExpr":
		node = &ast.CastExpr{
			Kind: ast.CastString,
			Expr: data["expr"].(ast.Expr),
		}
	case "CastUnsetExpr":
		node = &ast.CastExpr{
			Kind: ast.CastUnset,
			Expr: data["expr"].(ast.Expr),
		}
	case "ClassConstFetchExpr":
		name := data["name"]
		if _, ok := name.(ast.Expr); ok {
			err = unsupported("unsupported high version php feature: php8.3 dynamic class const fetch name")
			break
		}

		node = &ast.ClassConstFetchExpr{
			Class: data["class"].(ast.Node),
			Name:  name.(*ast.Ident),
		}
	case "CloneExpr":
		node = &ast.CloneExpr{
			Expr: data["expr"].(ast.Expr),
		}
	case "ClosureExpr":
		node = &ast.ClosureExpr{
			Static:     data["static"].(bool),
			ByRef:      data["byRef"].(bool),
			Params:     asSlice[*ast.Param](data["params"]),
			Uses:       asSlice[*ast.ClosureUseExpr](data["uses"]),
			ReturnType: asTypeHint(data["returnType"]),
			Stmts:      asStmtList(data["stmts"]),
		}
	case "ClosureUseExpr":
		node = &ast.ClosureUseExpr{
			Var:   data["var"].(*ast.VariableExpr),
			ByRef: data["byRef"].(bool),
		}
	case "ConstFetchExpr":
		node = &ast.ConstFetchExpr{
			Name: data["name"].(*ast.Name),
		}
	case "EmptyExpr":
		node = &ast.EmptyExpr{
			Expr: data["expr"].(ast.Expr),
		}
	case "ErrorSuppressExpr":
		node = &ast.ErrorSuppressExpr{
			Expr: data["expr"].(ast.Expr),
		}
	case "EvalExpr":
		node = &ast.EvalExpr{
			Expr: data["expr"].(ast.Expr),
		}
	case "ExitExpr":
		node = &ast.ExitExpr{
			Expr: asTypeOrNil[ast.Expr](data["expr"]),
		}
	case "FuncCallExpr":
		node = &ast.FuncCallExpr{
			Name: data["name"].(ast.Node),
			Args: asSlice[*ast.Arg](data["args"]),
		}
	case "IncludeExpr":
		var kind ast.IncludeKind
		typ := asInt(data["type"])
		switch typ {
		case 1:
			kind = ast.KindInclude
		case 2:
			kind = ast.KindIncludeOnce
		case 3:
			kind = ast.KindRequire
		case 4:
			kind = ast.KindRequireOnce
		default:
			return nil, fmt.Errorf("unexpteted ExprInclude.type: %d", typ)
		}
		node = &ast.IncludeExpr{
			Kind: kind,
			Expr: data["expr"].(ast.Expr),
		}
	case "InstanceofExpr":
		node = &ast.InstanceofExpr{
			Expr:  data["expr"].(ast.Expr),
			Class: data["class"].(ast.Node),
		}
	case "IssetExpr":
		node = &ast.IssetExpr{
			Vars: asSlice[ast.Expr](data["vars"]),
		}
	case "ListExpr":
		node = &ast.ListExpr{
			Items: asSliceItemNullable[*ast.ArrayItemExpr](data["items"]),
		}
	case "MethodCallExpr":
		node = &ast.MethodCallExpr{
			Var:  data["var"].(ast.Expr),
			Name: data["name"].(ast.Node),
			Args: asSlice[*ast.Arg](data["args"]),
		}
	case "NewExpr":
		node = &ast.NewExpr{
			Class: data["class"].(ast.Node),
			Args:  asSlice[*ast.Arg](data["args"]),
		}
	case "NullsafeMethodCallExpr":
		node = &ast.MethodCallExpr{
			Var:      data["var"].(ast.Expr),
			Name:     data["name"].(ast.Node),
			Args:     asSlice[*ast.Arg](data["args"]),
			Nullsafe: true,
		}
	case "NullsafePropertyFetchExpr":
		node = &ast.PropertyFetchExpr{
			Var:      data["var"].(ast.Expr),
			Name:     data["name"].(ast.Node),
			Nullsafe: true,
		}
	case "PostDecExpr":
		node = &ast.UnaryExpr{
			Op:  ast.UnaryOpPostDec,
			Var: data["var"].(ast.Expr),
		}
	case "PostIncExpr":
		node = &ast.UnaryExpr{
			Op:  ast.UnaryOpPostInc,
			Var: data["var"].(ast.Expr),
		}
	case "PreDecExpr":
		node = &ast.UnaryExpr{
			Op:  ast.UnaryOpPreDec,
			Var: data["var"].(ast.Expr),
		}
	case "PreIncExpr":
		node = &ast.UnaryExpr{
			Op:  ast.UnaryOpPreInc,
			Var: data["var"].(ast.Expr),
		}
	case "PrintExpr":
		node = &ast.PrintExpr{
			Expr: data["expr"].(ast.Expr),
		}
	case "PropertyFetchExpr":
		node = &ast.PropertyFetchExpr{
			Var:  data["var"].(ast.Expr),
			Name: data["name"].(ast.Node),
		}
	case "ShellExecExpr":
		node = &ast.ShellExecExpr{
			Parts: asSlice[ast.Expr](data["parts"]),
		}
	case "StaticCallExpr":
		node = &ast.StaticCallExpr{
			Class: data["class"].(ast.Node),
			Name:  data["name"].(ast.Node),
			Args:  asSlice[*ast.Arg](data["args"]),
		}
	case "StaticPropertyFetchExpr":
		node = &ast.StaticPropertyFetchExpr{
			Class: data["class"].(ast.Node),
			Name:  data["name"].(ast.Node),
		}
	case "TernaryExpr":
		node = &ast.TernaryExpr{
			Cond: data["cond"].(ast.Expr),
			If:   asTypeOrNil[ast.Expr](data["if"]),
			Else: data["else"].(ast.Expr),
		}
	case "ThrowExpr":
		node = &ast.ThrowExpr{
			Expr: data["expr"].(ast.Expr),
		}
	case "UnaryMinusExpr":
		node = &ast.UnaryExpr{
			Op:  ast.UnaryOpMinus,
			Var: data["expr"].(ast.Expr),
		}
	case "UnaryPlusExpr":
		node = &ast.UnaryExpr{
			Op:  ast.UnaryOpPlus,
			Var: data["expr"].(ast.Expr),
		}
	case "VariableExpr":
		var nameExpr ast.Node
		switch name := data["name"].(type) {
		case string:
			nameExpr = &ast.Ident{Name: name}
		default:
			nameExpr = name.(ast.Expr)
		}
		node = &ast.VariableExpr{
			Name: nameExpr,
		}
	case "YieldExpr":
		node = &ast.YieldExpr{
			Key:   asTypeOrNil[ast.Expr](data["key"]),
			Value: asTypeOrNil[ast.Expr](data["value"]),
		}
	case "YieldFromExpr":
		node = &ast.YieldFromExpr{
			Expr: data["expr"].(ast.Expr),
		}
	case "Identifier":
		node = &ast.Ident{
			Name: data["name"].(string),
		}
	case "IntersectionType":
		node = &ast.IntersectionType{
			Types: asTypeHints(data["types"]),
		}
	case "Name":
		node = &ast.Name{
			Kind:  ast.NameNormal,
			Parts: asSlice[string](data["parts"]),
		}
	case "NameFullyQualified":
		node = &ast.Name{
			Kind:  ast.NameFullyQualified,
			Parts: asSlice[string](data["parts"]),
		}
	case "NameRelative":
		node = &ast.Name{
			Kind:  ast.NameRelative,
			Parts: asSlice[string](data["parts"]),
		}
	case "NullableType":
		node = &ast.NullableType{
			Type: asTypeHint(data["type"]).(*ast.SimpleType),
		}
	case "Param":
		flags := asFlags(data["flags"])
		if flags != 0 {
			err = unsupported("unsupported high version php feature: php8.0 constructor promotion")
			break
		}

		node = &ast.Param{
			Type:     asTypeHint(data["type"]),
			ByRef:    data["byRef"].(bool),
			Variadic: data["variadic"].(bool),
			Var:      data["var"].(*ast.VariableExpr),
			Default:  asTypeOrNil[ast.Expr](data["default"]),
		}
	case "DNumberScalar":
		node = &ast.FloatLit{
			Value: asFloat(data["value"]),
		}
	case "EncapsedScalar":
		parts := asSlice[ast.Expr](data["parts"])
		if len(parts) == 0 {
			return nil, fmt.Errorf("EncapsedScalar need at least 1 part")
		}
		expr := parts[0]
		for _, next := range parts[1:] {
			expr = &ast.BinaryOpExpr{
				Op:    ast.BinaryOpConcat,
				Left:  expr,
				Right: next,
			}
		}
		node = expr
	case "EncapsedStringPartScalar":
		node = &ast.StringLit{
			Value: base64decode(data["value"].(string)),
		}
	case "LNumberScalar":
		node = &ast.IntLit{
			Value: asInt(data["value"]),
		}
	case "MagicConstClassScalar":
		node = &ast.MagicConstExpr{Kind: ast.MagicConstClass}
	case "MagicConstDirScalar":
		node = &ast.MagicConstExpr{Kind: ast.MagicConstDir}
	case "MagicConstFileScalar":
		node = &ast.MagicConstExpr{Kind: ast.MagicConstFile}
	case "MagicConstFunctionScalar":
		node = &ast.MagicConstExpr{Kind: ast.MagicConstFunction}
	case "MagicConstLineScalar":
		node = &ast.MagicConstExpr{Kind: ast.MagicConstLine}
	case "MagicConstMethodScalar":
		node = &ast.MagicConstExpr{Kind: ast.MagicConstMethod}
	case "MagicConstNamespaceScalar":
		node = &ast.MagicConstExpr{Kind: ast.MagicConstNamespace}
	case "MagicConstTraitScalar":
		node = &ast.MagicConstExpr{Kind: ast.MagicConstTrait}
	case "StringScalar":
		node = &ast.StringLit{
			Value: base64decode(data["value"].(string)),
		}
	case "BreakStmt":
		node = &ast.BreakStmt{
			Num: asTypeOrNil[ast.Expr](data["num"]),
		}
	case "CaseStmt":
		node = &ast.CaseStmt{
			Cond:  asTypeOrNil[ast.Expr](data["cond"]),
			Stmts: asStmtList(data["stmts"]),
		}
	case "CatchStmt":
		var_ := asTypeOrNil[*ast.VariableExpr](data["var"])
		if var_ == nil {
			err = unsupported("php8.0 catch an exception without storing it in a variable.")
			break
		}

		node = &ast.CatchStmt{
			Types: asSlice[*ast.Name](data["types"]),
			Var:   var_,
			Stmts: asStmtList(data["stmts"]),
		}
	case "ClassStmt":
		node = &ast.ClassStmt{
			Flags:          asFlags(data["flags"]),
			Extends:        asTypeOrNil[*ast.Name](data["extends"]),
			Implements:     asSlice[*ast.Name](data["implements"]),
			Name:           asTypeOrNil[*ast.Ident](data["name"]),
			Stmts:          asStmtList(data["stmts"]),
			NamespacedName: asTypeOrNil[*ast.Name](data["namespacedName"]),
		}
	case "ClassConstStmt":
		flags := asFlags(data["flags"])
		consts := asSlice[*ast.ConstStmt](data["consts"])
		typ := asTypeHint(data["type"])
		stmts := slicekit.Map(consts, func(c *ast.ConstStmt) ast.Stmt {
			if c.NamespacedName != nil {
				err = unsupported("ClassConst 不应有 NamespacedName")
			}

			return &ast.ClassConstStmt{
				Flags: flags,
				Type:  typ,
				Name:  c.Name,
				Value: c.Value,
			}
		})
		node = multiStmt(stmts)
	case "ClassMethodStmt":
		node = &ast.ClassMethodStmt{
			Flags:      asFlags(data["flags"]),
			ByRef:      data["byRef"].(bool),
			Name:       data["name"].(*ast.Ident),
			Params:     asSlice[*ast.Param](data["params"]),
			ReturnType: asTypeHint(data["returnType"]),
			Stmts:      asStmtList(data["stmts"]),
		}
	case "ConstStmt":
		stmts := asStmtList(data["consts"])
		node = multiStmt(stmts)
	case "ContinueStmt":
		node = &ast.ContinueStmt{
			Num: asTypeOrNil[ast.Expr](data["num"]),
		}
	case "DeclareStmt":
		node = &ast.DeclareStmt{
			Declares: asSlice[*ast.DeclareDeclareStmt](data["declares"]),
			Stmts:    asStmtList(data["stmts"]),
		}
	case "DeclareDeclareStmt":
		node = &ast.DeclareDeclareStmt{
			Key:   data["key"].(*ast.Ident),
			Value: data["value"].(ast.Expr),
		}
	case "DoStmt":
		node = &ast.DoStmt{
			Stmts: asStmtList(data["stmts"]),
			Cond:  data["cond"].(ast.Expr),
		}
	case "EchoStmt":
		node = &ast.EchoStmt{
			Exprs: asSlice[ast.Expr](data["exprs"]),
		}
	case "ElseStmt":
		node = &ast.ElseStmt{
			Stmts: asStmtList(data["stmts"]),
		}
	case "ElseIfStmt":
		node = &ast.ElseIfStmt{
			Cond:  data["cond"].(ast.Expr),
			Stmts: asStmtList(data["stmts"]),
		}
	case "ExpressionStmt":
		node = &ast.ExprStmt{
			Expr: data["expr"].(ast.Expr),
		}
	case "FinallyStmt":
		node = &ast.FinallyStmt{
			Stmts: asStmtList(data["stmts"]),
		}
	case "ForStmt":
		node = &ast.ForStmt{
			Init:  asSlice[ast.Expr](data["init"]),
			Cond:  asSlice[ast.Expr](data["cond"]),
			Loop:  asSlice[ast.Expr](data["loop"]),
			Stmts: asStmtList(data["stmts"]),
		}
	case "ForeachStmt":
		node = &ast.ForeachStmt{
			Expr:     data["expr"].(ast.Expr),
			KeyVar:   asTypeOrNil[ast.Expr](data["keyVar"]),
			ByRef:    data["byRef"].(bool),
			ValueVar: data["valueVar"].(ast.Expr),
			Stmts:    asStmtList(data["stmts"]),
		}
	case "FunctionStmt":
		node = &ast.FunctionStmt{
			ByRef:          data["byRef"].(bool),
			Name:           data["name"].(*ast.Ident),
			Params:         asSlice[*ast.Param](data["params"]),
			ReturnType:     asTypeHint(data["returnType"]),
			Stmts:          asStmtList(data["stmts"]),
			NamespacedName: asTypeOrNil[*ast.Name](data["namespacedName"]),
		}
	case "GlobalStmt":
		vars := asSlice[ast.Expr](data["vars"])
		stmts := slicekit.Map(vars, func(v ast.Expr) ast.Stmt {
			return &ast.GlobalStmt{
				Var: v,
			}
		})
		node = multiStmt(stmts)
	case "GotoStmt":
		node = &ast.GotoStmt{
			Name: data["name"].(*ast.Ident),
		}
	case "GroupUseStmt":
		typ := asUseType(data["type"])
		prefix := data["prefix"].(*ast.Name)
		uses := asSlice[*ast.UseStmt](data["uses"])
		stmts := slicekit.Map(uses, func(useStmt *ast.UseStmt) ast.Stmt {
			if typ != ast.UseNormal {
				useStmt.Type = typ
			}
			useStmt.Name = ast.NewName(slicekit.Concat(prefix.Parts, useStmt.Name.Parts)...)

			return useStmt
		})
		node = multiStmt(stmts)
	case "HaltCompilerStmt":
		node = &ast.HaltCompilerStmt{
			Remaining: data["remaining"].(string),
		}
	case "IfStmt":
		node = &ast.IfStmt{
			Cond:    data["cond"].(ast.Expr),
			Stmts:   asStmtList(data["stmts"]),
			Elseifs: asSlice[*ast.ElseIfStmt](data["elseifs"]),
			Else:    asTypeOrNil[*ast.ElseStmt](data["else"]),
		}
	case "InlineHTMLStmt":
		node = &ast.InlineHTMLStmt{
			Value: data["value"].(string),
		}
	case "InterfaceStmt":
		node = &ast.InterfaceStmt{
			Extends:        asSlice[*ast.Name](data["extends"]),
			Name:           asTypeOrNil[*ast.Ident](data["name"]),
			Stmts:          asStmtList(data["stmts"]),
			NamespacedName: asTypeOrNil[*ast.Name](data["namespacedName"]),
		}
	case "LabelStmt":
		node = &ast.LabelStmt{
			Name: data["name"].(*ast.Ident),
		}
	case "NamespaceStmt":
		node = &ast.NamespaceStmt{
			Name:  asTypeOrNil[*ast.Name](data["name"]),
			Stmts: asStmtList(data["stmts"]),
		}
	case "NopStmt":
		node = &ast.EmptyStmt{}
	case "PropertyStmt":
		flags := asFlags(data["flags"])
		typeHint := asTypeHint(data["type"])
		properties := asSlice[*ast.PropertyStmt](data["props"])
		stmts := slicekit.Map(properties, func(t *ast.PropertyStmt) ast.Stmt {
			t.Flags = flags
			t.Type = typeHint
			return t
		})
		node = multiStmt(stmts)
	case "PropertyPropertyStmt":
		node = &ast.PropertyStmt{
			Name:    data["name"].(*ast.Ident),
			Default: asTypeOrNil[ast.Expr](data["default"]),
		}
	case "ReturnStmt":
		node = &ast.ReturnStmt{
			Expr: asTypeOrNil[ast.Expr](data["expr"]),
		}
	case "StaticStmt":
		vars := asSlice[*ast.StaticStmt](data["vars"])
		node = multiStmt(vars)
	case "StaticVarStmt":
		node = &ast.StaticStmt{
			Var:     data["var"].(*ast.VariableExpr),
			Default: asTypeOrNil[ast.Expr](data["default"]),
		}
	case "SwitchStmt":
		node = &ast.SwitchStmt{
			Cond:  data["cond"].(ast.Expr),
			Cases: asSlice[*ast.CaseStmt](data["cases"]),
		}
	case "ThrowStmt":
		node = &ast.ExprStmt{
			Expr: &ast.ThrowExpr{
				Expr: data["expr"].(ast.Expr),
			},
		}
	case "TraitStmt":
		node = &ast.TraitStmt{
			Name:           asTypeOrNil[*ast.Ident](data["name"]),
			Stmts:          asStmtList(data["stmts"]),
			NamespacedName: asTypeOrNil[*ast.Name](data["namespacedName"]),
		}
	case "TraitUseStmt":
		node = &ast.TraitUseStmt{
			Traits:      asSlice[*ast.Name](data["traits"]),
			Adaptations: asSlice[ast.TraitUseAdaptationStmt](data["adaptations"]),
		}
	case "TraitUseAdaptationAliasStmt":
		node = &ast.TraitUseAdaptationAliasStmt{
			NewModifier: asFlags(data["newModifier"]),
			NewName:     asTypeOrNil[*ast.Ident](data["newName"]),
			Trait:       asTypeOrNil[*ast.Name](data["trait"]),
			Method:      data["method"].(*ast.Ident),
		}
	case "TraitUseAdaptationPrecedenceStmt":
		node = &ast.TraitUseAdaptationPrecedenceStmt{
			Insteadof: asSlice[*ast.Name](data["insteadof"]),
			Trait:     asTypeOrNil[*ast.Name](data["trait"]),
			Method:    data["method"].(*ast.Ident),
		}
	case "TryCatchStmt":
		node = &ast.TryCatchStmt{
			Stmts:   asStmtList(data["stmts"]),
			Catches: asSlice[*ast.CatchStmt](data["catches"]),
			Finally: asTypeOrNil[*ast.FinallyStmt](data["finally"]),
		}
	case "UnsetStmt":
		vars := asSlice[ast.Expr](data["vars"])
		stmts := slicekit.Map(vars, func(v ast.Expr) ast.Stmt {
			return &ast.UnsetStmt{
				Var: v,
			}
		})
		node = multiStmt(stmts)
	case "UseStmt":
		typ := asUseType(data["type"])
		uses := asSlice[*ast.UseStmt](data["uses"])
		stmts := slicekit.Map(uses, func(useStmt *ast.UseStmt) ast.Stmt {
			if typ != ast.UseNormal {
				useStmt.Type = typ
			}
			return useStmt
		})
		node = multiStmt(stmts)
	case "UseUseStmt":
		node = &ast.UseStmt{
			Type:  asUseType(data["type"]),
			Name:  data["name"].(*ast.Name),
			Alias: asTypeOrNil[*ast.Ident](data["alias"]),
		}
	case "WhileStmt":
		node = &ast.WhileStmt{
			Cond:  data["cond"].(ast.Expr),
			Stmts: asStmtList(data["stmts"]),
		}
	case "UnionType":
		node = &ast.UnionType{
			Types: asTypeHints(data["types"]),
		}
	case "VarLikeIdentifier":
		node = &ast.Ident{
			Name:    data["name"].(string),
			VarLike: true,
		}
	case "Comment":
		node = &ast.Comment{
			Type: ast.CommentLine,
			Text: data["text"].(string),
		}
	case "AttributeGroup", "Attribute":
		err = unsupported("unsupported high version php feature: php8.0 attribute")
	case "MatchExpr", "MatchArm":
		err = unsupported("unsupported high version php feature: php8.0 match")
	case "EnumStmt", "EnumCaseStmt":
		err = unsupported("unsupported high version php feature: php8.0 match")
	case "VariadicPlaceholder":
		err = unsupported("unsupported high version php feature: php8.2 first class callable syntax")
	default:
		err = unsupported("unexpected node type: " + nodeType)
	}
	if err == nil && node != nil && data["attributes"] != nil {
		trySetMeta(node, data["attributes"])
	}
	return
}
