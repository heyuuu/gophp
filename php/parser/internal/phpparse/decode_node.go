package phpparse

import "gophp/php/ast"

func decodeNode(data map[string]any) (node ast.Node, err error) {
	nodeType := data["nodeType"].(string)
	switch nodeType {
    case "Arg":
        node = &ast.Arg{
            Name: asTypeOrNil[*ast.Identifier](data["name"]),
            Value: data["value"].(ast.Expr),
            ByRef: data["byRef"].(bool),
            Unpack: data["unpack"].(bool),
            Attributes: data["attributes"],

        }
    case "Attribute":
        node = &ast.Attribute{
            Name: data["name"].(*ast.Name),
            Args: asSlice[*ast.Arg](data["args"]),
            Attributes: data["attributes"],

        }
    case "AttributeGroup":
        node = &ast.AttributeGroup{
            Attrs: asSlice[*ast.Attribute](data["attrs"]),
            Attributes: data["attributes"],

        }
    case "Const":
        node = &ast.Const{
            Name: data["name"].(*ast.Identifier),
            Value: data["value"].(ast.Expr),
            NamespacedName: asTypeOrNil[*ast.Name](data["namespacedName"]),
            Attributes: data["attributes"],

        }
    case "ExprArray":
        node = &ast.ArrayExpr{
            Items: asSlice[*ast.ArrayItemExpr](data["items"]),
            Attributes: data["attributes"],

        }
    case "ExprArrayDimFetch":
        node = &ast.ArrayDimFetchExpr{
            Var: data["var"].(ast.Expr),
            Dim: asTypeOrNil[ast.Expr](data["dim"]),
            Attributes: data["attributes"],

        }
    case "ExprArrayItem":
        node = &ast.ArrayItemExpr{
            Key: asTypeOrNil[ast.Expr](data["key"]),
            Value: data["value"].(ast.Expr),
            ByRef: data["byRef"].(bool),
            Unpack: data["unpack"].(bool),
            Attributes: data["attributes"],

        }
    case "ExprArrowFunction":
        node = &ast.ArrowFunctionExpr{
            Static: data["static"].(bool),
            ByRef: data["byRef"].(bool),
            Params: asSlice[*ast.Param](data["params"]),
            ReturnType: data["returnType"],
            Expr: data["expr"].(ast.Expr),
            AttrGroups: asSlice[*ast.AttributeGroup](data["attrGroups"]),
            Attributes: data["attributes"],

        }
    case "ExprAssign":
        node = &ast.AssignExpr{
            Var: data["var"].(ast.Expr),
            Expr: data["expr"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprAssignOpBitwiseAnd":
        node = &ast.AssignOpBitwiseAndExpr{
            Var: data["var"].(ast.Expr),
            Expr: data["expr"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprAssignOpBitwiseOr":
        node = &ast.AssignOpBitwiseOrExpr{
            Var: data["var"].(ast.Expr),
            Expr: data["expr"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprAssignOpBitwiseXor":
        node = &ast.AssignOpBitwiseXorExpr{
            Var: data["var"].(ast.Expr),
            Expr: data["expr"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprAssignOpCoalesce":
        node = &ast.AssignOpCoalesceExpr{
            Var: data["var"].(ast.Expr),
            Expr: data["expr"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprAssignOpConcat":
        node = &ast.AssignOpConcatExpr{
            Var: data["var"].(ast.Expr),
            Expr: data["expr"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprAssignOpDiv":
        node = &ast.AssignOpDivExpr{
            Var: data["var"].(ast.Expr),
            Expr: data["expr"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprAssignOpMinus":
        node = &ast.AssignOpMinusExpr{
            Var: data["var"].(ast.Expr),
            Expr: data["expr"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprAssignOpMod":
        node = &ast.AssignOpModExpr{
            Var: data["var"].(ast.Expr),
            Expr: data["expr"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprAssignOpMul":
        node = &ast.AssignOpMulExpr{
            Var: data["var"].(ast.Expr),
            Expr: data["expr"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprAssignOpPlus":
        node = &ast.AssignOpPlusExpr{
            Var: data["var"].(ast.Expr),
            Expr: data["expr"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprAssignOpPow":
        node = &ast.AssignOpPowExpr{
            Var: data["var"].(ast.Expr),
            Expr: data["expr"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprAssignOpShiftLeft":
        node = &ast.AssignOpShiftLeftExpr{
            Var: data["var"].(ast.Expr),
            Expr: data["expr"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprAssignOpShiftRight":
        node = &ast.AssignOpShiftRightExpr{
            Var: data["var"].(ast.Expr),
            Expr: data["expr"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprAssignRef":
        node = &ast.AssignRefExpr{
            Var: data["var"].(ast.Expr),
            Expr: data["expr"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprBinaryOpBitwiseAnd":
        node = &ast.BinaryOpBitwiseAndExpr{
            Left: data["left"].(ast.Expr),
            Right: data["right"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprBinaryOpBitwiseOr":
        node = &ast.BinaryOpBitwiseOrExpr{
            Left: data["left"].(ast.Expr),
            Right: data["right"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprBinaryOpBitwiseXor":
        node = &ast.BinaryOpBitwiseXorExpr{
            Left: data["left"].(ast.Expr),
            Right: data["right"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprBinaryOpBooleanAnd":
        node = &ast.BinaryOpBooleanAndExpr{
            Left: data["left"].(ast.Expr),
            Right: data["right"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprBinaryOpBooleanOr":
        node = &ast.BinaryOpBooleanOrExpr{
            Left: data["left"].(ast.Expr),
            Right: data["right"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprBinaryOpCoalesce":
        node = &ast.BinaryOpCoalesceExpr{
            Left: data["left"].(ast.Expr),
            Right: data["right"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprBinaryOpConcat":
        node = &ast.BinaryOpConcatExpr{
            Left: data["left"].(ast.Expr),
            Right: data["right"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprBinaryOpDiv":
        node = &ast.BinaryOpDivExpr{
            Left: data["left"].(ast.Expr),
            Right: data["right"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprBinaryOpEqual":
        node = &ast.BinaryOpEqualExpr{
            Left: data["left"].(ast.Expr),
            Right: data["right"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprBinaryOpGreater":
        node = &ast.BinaryOpGreaterExpr{
            Left: data["left"].(ast.Expr),
            Right: data["right"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprBinaryOpGreaterOrEqual":
        node = &ast.BinaryOpGreaterOrEqualExpr{
            Left: data["left"].(ast.Expr),
            Right: data["right"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprBinaryOpIdentical":
        node = &ast.BinaryOpIdenticalExpr{
            Left: data["left"].(ast.Expr),
            Right: data["right"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprBinaryOpLogicalAnd":
        node = &ast.BinaryOpLogicalAndExpr{
            Left: data["left"].(ast.Expr),
            Right: data["right"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprBinaryOpLogicalOr":
        node = &ast.BinaryOpLogicalOrExpr{
            Left: data["left"].(ast.Expr),
            Right: data["right"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprBinaryOpLogicalXor":
        node = &ast.BinaryOpLogicalXorExpr{
            Left: data["left"].(ast.Expr),
            Right: data["right"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprBinaryOpMinus":
        node = &ast.BinaryOpMinusExpr{
            Left: data["left"].(ast.Expr),
            Right: data["right"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprBinaryOpMod":
        node = &ast.BinaryOpModExpr{
            Left: data["left"].(ast.Expr),
            Right: data["right"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprBinaryOpMul":
        node = &ast.BinaryOpMulExpr{
            Left: data["left"].(ast.Expr),
            Right: data["right"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprBinaryOpNotEqual":
        node = &ast.BinaryOpNotEqualExpr{
            Left: data["left"].(ast.Expr),
            Right: data["right"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprBinaryOpNotIdentical":
        node = &ast.BinaryOpNotIdenticalExpr{
            Left: data["left"].(ast.Expr),
            Right: data["right"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprBinaryOpPlus":
        node = &ast.BinaryOpPlusExpr{
            Left: data["left"].(ast.Expr),
            Right: data["right"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprBinaryOpPow":
        node = &ast.BinaryOpPowExpr{
            Left: data["left"].(ast.Expr),
            Right: data["right"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprBinaryOpShiftLeft":
        node = &ast.BinaryOpShiftLeftExpr{
            Left: data["left"].(ast.Expr),
            Right: data["right"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprBinaryOpShiftRight":
        node = &ast.BinaryOpShiftRightExpr{
            Left: data["left"].(ast.Expr),
            Right: data["right"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprBinaryOpSmaller":
        node = &ast.BinaryOpSmallerExpr{
            Left: data["left"].(ast.Expr),
            Right: data["right"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprBinaryOpSmallerOrEqual":
        node = &ast.BinaryOpSmallerOrEqualExpr{
            Left: data["left"].(ast.Expr),
            Right: data["right"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprBinaryOpSpaceship":
        node = &ast.BinaryOpSpaceshipExpr{
            Left: data["left"].(ast.Expr),
            Right: data["right"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprBitwiseNot":
        node = &ast.BitwiseNotExpr{
            Expr: data["expr"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprBooleanNot":
        node = &ast.BooleanNotExpr{
            Expr: data["expr"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprCastArray":
        node = &ast.CastArrayExpr{
            Expr: data["expr"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprCastBool":
        node = &ast.CastBoolExpr{
            Expr: data["expr"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprCastDouble":
        node = &ast.CastDoubleExpr{
            Expr: data["expr"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprCastInt":
        node = &ast.CastIntExpr{
            Expr: data["expr"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprCastObject":
        node = &ast.CastObjectExpr{
            Expr: data["expr"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprCastString":
        node = &ast.CastStringExpr{
            Expr: data["expr"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprCastUnset":
        node = &ast.CastUnsetExpr{
            Expr: data["expr"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprClassConstFetch":
        node = &ast.ClassConstFetchExpr{
            Class: data["class"],
            Name: data["name"],
            Attributes: data["attributes"],

        }
    case "ExprClone":
        node = &ast.CloneExpr{
            Expr: data["expr"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprClosure":
        node = &ast.ClosureExpr{
            Static: data["static"].(bool),
            ByRef: data["byRef"].(bool),
            Params: asSlice[*ast.Param](data["params"]),
            Uses: asSlice[*ast.ClosureUseExpr](data["uses"]),
            ReturnType: data["returnType"],
            Stmts: asSlice[ast.Stmt](data["stmts"]),
            AttrGroups: asSlice[*ast.AttributeGroup](data["attrGroups"]),
            Attributes: data["attributes"],

        }
    case "ExprClosureUse":
        node = &ast.ClosureUseExpr{
            Var: data["var"].(*ast.VariableExpr),
            ByRef: data["byRef"].(bool),
            Attributes: data["attributes"],

        }
    case "ExprConstFetch":
        node = &ast.ConstFetchExpr{
            Name: data["name"].(*ast.Name),
            Attributes: data["attributes"],

        }
    case "ExprEmpty":
        node = &ast.EmptyExpr{
            Expr: data["expr"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprError":
        node = &ast.ErrorExpr{
            Attributes: data["attributes"],

        }
    case "ExprErrorSuppress":
        node = &ast.ErrorSuppressExpr{
            Expr: data["expr"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprEval":
        node = &ast.EvalExpr{
            Expr: data["expr"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprExit":
        node = &ast.ExitExpr{
            Expr: asTypeOrNil[ast.Expr](data["expr"]),
            Attributes: data["attributes"],

        }
    case "ExprFuncCall":
        node = &ast.FuncCallExpr{
            Name: data["name"],
            Args: asSlice[any](data["args"]),
            Attributes: data["attributes"],

        }
    case "ExprInclude":
        node = &ast.IncludeExpr{
            Expr: data["expr"].(ast.Expr),
            Type: asInt(data["type"]),
            Attributes: data["attributes"],

        }
    case "ExprInstanceof":
        node = &ast.InstanceofExpr{
            Expr: data["expr"].(ast.Expr),
            Class: data["class"],
            Attributes: data["attributes"],

        }
    case "ExprIsset":
        node = &ast.IssetExpr{
            Vars: asSlice[ast.Expr](data["vars"]),
            Attributes: data["attributes"],

        }
    case "ExprList":
        node = &ast.ListExpr{
            Items: asSlice[*ast.ArrayItemExpr](data["items"]),
            Attributes: data["attributes"],

        }
    case "ExprMatch":
        node = &ast.MatchExpr{
            Cond: data["cond"].(ast.Expr),
            Arms: asSlice[*ast.MatchArm](data["arms"]),
            Attributes: data["attributes"],

        }
    case "ExprMethodCall":
        node = &ast.MethodCallExpr{
            Var: data["var"].(ast.Expr),
            Name: data["name"],
            Args: asSlice[any](data["args"]),
            Attributes: data["attributes"],

        }
    case "ExprNew":
        node = &ast.NewExpr{
            Class: data["class"],
            Args: asSlice[any](data["args"]),
            Attributes: data["attributes"],

        }
    case "ExprNullsafeMethodCall":
        node = &ast.NullsafeMethodCallExpr{
            Var: data["var"].(ast.Expr),
            Name: data["name"],
            Args: asSlice[any](data["args"]),
            Attributes: data["attributes"],

        }
    case "ExprNullsafePropertyFetch":
        node = &ast.NullsafePropertyFetchExpr{
            Var: data["var"].(ast.Expr),
            Name: data["name"],
            Attributes: data["attributes"],

        }
    case "ExprPostDec":
        node = &ast.PostDecExpr{
            Var: data["var"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprPostInc":
        node = &ast.PostIncExpr{
            Var: data["var"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprPreDec":
        node = &ast.PreDecExpr{
            Var: data["var"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprPreInc":
        node = &ast.PreIncExpr{
            Var: data["var"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprPrint":
        node = &ast.PrintExpr{
            Expr: data["expr"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprPropertyFetch":
        node = &ast.PropertyFetchExpr{
            Var: data["var"].(ast.Expr),
            Name: data["name"],
            Attributes: data["attributes"],

        }
    case "ExprShellExec":
        node = &ast.ShellExecExpr{
            Parts: asSlice[any](data["parts"]),
            Attributes: data["attributes"],

        }
    case "ExprStaticCall":
        node = &ast.StaticCallExpr{
            Class: data["class"],
            Name: data["name"],
            Args: asSlice[any](data["args"]),
            Attributes: data["attributes"],

        }
    case "ExprStaticPropertyFetch":
        node = &ast.StaticPropertyFetchExpr{
            Class: data["class"],
            Name: data["name"],
            Attributes: data["attributes"],

        }
    case "ExprTernary":
        node = &ast.TernaryExpr{
            Cond: data["cond"].(ast.Expr),
            If: asTypeOrNil[ast.Expr](data["if"]),
            Else: data["else"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprThrow":
        node = &ast.ThrowExpr{
            Expr: data["expr"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprUnaryMinus":
        node = &ast.UnaryMinusExpr{
            Expr: data["expr"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprUnaryPlus":
        node = &ast.UnaryPlusExpr{
            Expr: data["expr"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "ExprVariable":
        node = &ast.VariableExpr{
            Name: data["name"],
            Attributes: data["attributes"],

        }
    case "ExprYield":
        node = &ast.YieldExpr{
            Key: asTypeOrNil[ast.Expr](data["key"]),
            Value: asTypeOrNil[ast.Expr](data["value"]),
            Attributes: data["attributes"],

        }
    case "ExprYieldFrom":
        node = &ast.YieldFromExpr{
            Expr: data["expr"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "Identifier":
        node = &ast.Identifier{
            Name: data["name"].(string),
            SpecialClassNames: data["specialClassNames"],
            Attributes: data["attributes"],

        }
    case "IntersectionType":
        node = &ast.IntersectionType{
            Types: asSlice[any](data["types"]),
            Attributes: data["attributes"],

        }
    case "MatchArm":
        node = &ast.MatchArm{
            Conds: asSlice[ast.Expr](data["conds"]),
            Body: data["body"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "Name":
        node = &ast.Name{
            Parts: asSlice[string](data["parts"]),
            SpecialClassNames: data["specialClassNames"],
            Attributes: data["attributes"],

        }
    case "NameFullyQualified":
        node = &ast.NameFullyQualified{
            Parts: asSlice[string](data["parts"]),
            Attributes: data["attributes"],

        }
    case "NameRelative":
        node = &ast.NameRelative{
            Parts: asSlice[string](data["parts"]),
            Attributes: data["attributes"],

        }
    case "NullableType":
        node = &ast.NullableType{
            Type: data["type"],
            Attributes: data["attributes"],

        }
    case "Param":
        node = &ast.Param{
            Type: data["type"],
            ByRef: data["byRef"].(bool),
            Variadic: data["variadic"].(bool),
            Var: data["var"],
            Default: asTypeOrNil[ast.Expr](data["default"]),
            Flags: asInt(data["flags"]),
            AttrGroups: asSlice[*ast.AttributeGroup](data["attrGroups"]),
            Attributes: data["attributes"],

        }
    case "ScalarDNumber":
        node = &ast.DNumberScalar{
            Value: asFloat(data["value"]),
            Attributes: data["attributes"],

        }
    case "ScalarEncapsed":
        node = &ast.EncapsedScalar{
            Parts: asSlice[ast.Expr](data["parts"]),
            Attributes: data["attributes"],

        }
    case "ScalarEncapsedStringPart":
        node = &ast.EncapsedStringPartScalar{
            Value: data["value"].(string),
            Attributes: data["attributes"],

        }
    case "ScalarLNumber":
        node = &ast.LNumberScalar{
            Value: asInt(data["value"]),
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
            Value: data["value"].(string),
            Replacements: data["replacements"],
            Attributes: data["attributes"],

        }
    case "StmtBreak":
        node = &ast.BreakStmt{
            Num: asTypeOrNil[ast.Expr](data["num"]),
            Attributes: data["attributes"],

        }
    case "StmtCase":
        node = &ast.CaseStmt{
            Cond: asTypeOrNil[ast.Expr](data["cond"]),
            Stmts: asSlice[ast.Stmt](data["stmts"]),
            Attributes: data["attributes"],

        }
    case "StmtCatch":
        node = &ast.CatchStmt{
            Types: asSlice[*ast.Name](data["types"]),
            Var: asTypeOrNil[*ast.VariableExpr](data["var"]),
            Stmts: asSlice[ast.Stmt](data["stmts"]),
            Attributes: data["attributes"],

        }
    case "StmtClass":
        node = &ast.ClassStmt{
            Flags: asInt(data["flags"]),
            Extends: asTypeOrNil[*ast.Name](data["extends"]),
            Implements: asSlice[*ast.Name](data["implements"]),
            Name: asTypeOrNil[*ast.Identifier](data["name"]),
            Stmts: asSlice[ast.Stmt](data["stmts"]),
            AttrGroups: asSlice[*ast.AttributeGroup](data["attrGroups"]),
            NamespacedName: asTypeOrNil[*ast.Name](data["namespacedName"]),
            Attributes: data["attributes"],

        }
    case "StmtClassConst":
        node = &ast.ClassConstStmt{
            Flags: asInt(data["flags"]),
            Consts: asSlice[*ast.Const](data["consts"]),
            AttrGroups: asSlice[*ast.AttributeGroup](data["attrGroups"]),
            Attributes: data["attributes"],

        }
    case "StmtClassMethod":
        node = &ast.ClassMethodStmt{
            Flags: asInt(data["flags"]),
            ByRef: data["byRef"].(bool),
            Name: data["name"].(*ast.Identifier),
            Params: asSlice[*ast.Param](data["params"]),
            ReturnType: data["returnType"],
            Stmts: asSlice[ast.Stmt](data["stmts"]),
            AttrGroups: asSlice[*ast.AttributeGroup](data["attrGroups"]),
            MagicNames: data["magicNames"],
            Attributes: data["attributes"],

        }
    case "StmtConst":
        node = &ast.ConstStmt{
            Consts: asSlice[*ast.Const](data["consts"]),
            Attributes: data["attributes"],

        }
    case "StmtContinue":
        node = &ast.ContinueStmt{
            Num: asTypeOrNil[ast.Expr](data["num"]),
            Attributes: data["attributes"],

        }
    case "StmtDeclare":
        node = &ast.DeclareStmt{
            Declares: asSlice[*ast.DeclareDeclareStmt](data["declares"]),
            Stmts: asSlice[ast.Stmt](data["stmts"]),
            Attributes: data["attributes"],

        }
    case "StmtDeclareDeclare":
        node = &ast.DeclareDeclareStmt{
            Key: data["key"].(*ast.Identifier),
            Value: data["value"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "StmtDo":
        node = &ast.DoStmt{
            Stmts: asSlice[ast.Stmt](data["stmts"]),
            Cond: data["cond"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "StmtEcho":
        node = &ast.EchoStmt{
            Exprs: asSlice[ast.Expr](data["exprs"]),
            Attributes: data["attributes"],

        }
    case "StmtElse":
        node = &ast.ElseStmt{
            Stmts: asSlice[ast.Stmt](data["stmts"]),
            Attributes: data["attributes"],

        }
    case "StmtElseIf":
        node = &ast.ElseIfStmt{
            Cond: data["cond"].(ast.Expr),
            Stmts: asSlice[ast.Stmt](data["stmts"]),
            Attributes: data["attributes"],

        }
    case "StmtEnum":
        node = &ast.EnumStmt{
            ScalarType: asTypeOrNil[*ast.Identifier](data["scalarType"]),
            Implements: asSlice[*ast.Name](data["implements"]),
            Name: asTypeOrNil[*ast.Identifier](data["name"]),
            Stmts: asSlice[ast.Stmt](data["stmts"]),
            AttrGroups: asSlice[*ast.AttributeGroup](data["attrGroups"]),
            NamespacedName: asTypeOrNil[*ast.Name](data["namespacedName"]),
            Attributes: data["attributes"],

        }
    case "StmtEnumCase":
        node = &ast.EnumCaseStmt{
            Name: data["name"].(*ast.Identifier),
            Expr: asTypeOrNil[ast.Expr](data["expr"]),
            AttrGroups: asSlice[*ast.AttributeGroup](data["attrGroups"]),
            Attributes: data["attributes"],

        }
    case "StmtExpression":
        node = &ast.ExpressionStmt{
            Expr: data["expr"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "StmtFinally":
        node = &ast.FinallyStmt{
            Stmts: asSlice[ast.Stmt](data["stmts"]),
            Attributes: data["attributes"],

        }
    case "StmtFor":
        node = &ast.ForStmt{
            Init: asSlice[ast.Expr](data["init"]),
            Cond: asSlice[ast.Expr](data["cond"]),
            Loop: asSlice[ast.Expr](data["loop"]),
            Stmts: asSlice[ast.Stmt](data["stmts"]),
            Attributes: data["attributes"],

        }
    case "StmtForeach":
        node = &ast.ForeachStmt{
            Expr: data["expr"].(ast.Expr),
            KeyVar: asTypeOrNil[ast.Expr](data["keyVar"]),
            ByRef: data["byRef"].(bool),
            ValueVar: data["valueVar"].(ast.Expr),
            Stmts: asSlice[ast.Stmt](data["stmts"]),
            Attributes: data["attributes"],

        }
    case "StmtFunction":
        node = &ast.FunctionStmt{
            ByRef: data["byRef"].(bool),
            Name: data["name"].(*ast.Identifier),
            Params: asSlice[*ast.Param](data["params"]),
            ReturnType: data["returnType"],
            Stmts: asSlice[ast.Stmt](data["stmts"]),
            AttrGroups: asSlice[*ast.AttributeGroup](data["attrGroups"]),
            NamespacedName: asTypeOrNil[*ast.Name](data["namespacedName"]),
            Attributes: data["attributes"],

        }
    case "StmtGlobal":
        node = &ast.GlobalStmt{
            Vars: asSlice[ast.Expr](data["vars"]),
            Attributes: data["attributes"],

        }
    case "StmtGoto":
        node = &ast.GotoStmt{
            Name: data["name"].(*ast.Identifier),
            Attributes: data["attributes"],

        }
    case "StmtGroupUse":
        node = &ast.GroupUseStmt{
            Type: asInt(data["type"]),
            Prefix: data["prefix"].(*ast.Name),
            Uses: asSlice[*ast.UseUseStmt](data["uses"]),
            Attributes: data["attributes"],

        }
    case "StmtHaltCompiler":
        node = &ast.HaltCompilerStmt{
            Remaining: data["remaining"].(string),
            Attributes: data["attributes"],

        }
    case "StmtIf":
        node = &ast.IfStmt{
            Cond: data["cond"].(ast.Expr),
            Stmts: asSlice[ast.Stmt](data["stmts"]),
            Elseifs: asSlice[*ast.ElseIfStmt](data["elseifs"]),
            Else: asTypeOrNil[*ast.ElseStmt](data["else"]),
            Attributes: data["attributes"],

        }
    case "StmtInlineHTML":
        node = &ast.InlineHTMLStmt{
            Value: data["value"].(string),
            Attributes: data["attributes"],

        }
    case "StmtInterface":
        node = &ast.InterfaceStmt{
            Extends: asSlice[*ast.Name](data["extends"]),
            Name: asTypeOrNil[*ast.Identifier](data["name"]),
            Stmts: asSlice[ast.Stmt](data["stmts"]),
            AttrGroups: asSlice[*ast.AttributeGroup](data["attrGroups"]),
            NamespacedName: asTypeOrNil[*ast.Name](data["namespacedName"]),
            Attributes: data["attributes"],

        }
    case "StmtLabel":
        node = &ast.LabelStmt{
            Name: data["name"].(*ast.Identifier),
            Attributes: data["attributes"],

        }
    case "StmtNamespace":
        node = &ast.NamespaceStmt{
            Name: asTypeOrNil[*ast.Name](data["name"]),
            Stmts: asSlice[ast.Stmt](data["stmts"]),
            Attributes: data["attributes"],

        }
    case "StmtNop":
        node = &ast.NopStmt{
            Attributes: data["attributes"],

        }
    case "StmtProperty":
        node = &ast.PropertyStmt{
            Flags: asInt(data["flags"]),
            Props: asSlice[*ast.PropertyPropertyStmt](data["props"]),
            Type: data["type"],
            AttrGroups: asSlice[*ast.AttributeGroup](data["attrGroups"]),
            Attributes: data["attributes"],

        }
    case "StmtPropertyProperty":
        node = &ast.PropertyPropertyStmt{
            Name: data["name"].(*ast.VarLikeIdentifier),
            Default: asTypeOrNil[ast.Expr](data["default"]),
            Attributes: data["attributes"],

        }
    case "StmtReturn":
        node = &ast.ReturnStmt{
            Expr: asTypeOrNil[ast.Expr](data["expr"]),
            Attributes: data["attributes"],

        }
    case "StmtStatic":
        node = &ast.StaticStmt{
            Vars: asSlice[*ast.StaticVarStmt](data["vars"]),
            Attributes: data["attributes"],

        }
    case "StmtStaticVar":
        node = &ast.StaticVarStmt{
            Var: data["var"].(*ast.VariableExpr),
            Default: asTypeOrNil[ast.Expr](data["default"]),
            Attributes: data["attributes"],

        }
    case "StmtSwitch":
        node = &ast.SwitchStmt{
            Cond: data["cond"].(ast.Expr),
            Cases: asSlice[*ast.CaseStmt](data["cases"]),
            Attributes: data["attributes"],

        }
    case "StmtThrow":
        node = &ast.ThrowStmt{
            Expr: data["expr"].(ast.Expr),
            Attributes: data["attributes"],

        }
    case "StmtTrait":
        node = &ast.TraitStmt{
            Name: asTypeOrNil[*ast.Identifier](data["name"]),
            Stmts: asSlice[ast.Stmt](data["stmts"]),
            AttrGroups: asSlice[*ast.AttributeGroup](data["attrGroups"]),
            NamespacedName: asTypeOrNil[*ast.Name](data["namespacedName"]),
            Attributes: data["attributes"],

        }
    case "StmtTraitUse":
        node = &ast.TraitUseStmt{
            Traits: asSlice[*ast.Name](data["traits"]),
            Adaptations: asSlice[*ast.TraitUseAdaptationStmt](data["adaptations"]),
            Attributes: data["attributes"],

        }
    case "StmtTraitUseAdaptationAlias":
        node = &ast.TraitUseAdaptationAliasStmt{
            NewModifier: asInt(data["newModifier"]),
            NewName: asTypeOrNil[*ast.Identifier](data["newName"]),
            Trait: asTypeOrNil[*ast.Name](data["trait"]),
            Method: data["method"].(*ast.Identifier),
            Attributes: data["attributes"],

        }
    case "StmtTraitUseAdaptationPrecedence":
        node = &ast.TraitUseAdaptationPrecedenceStmt{
            Insteadof: asSlice[*ast.Name](data["insteadof"]),
            Trait: asTypeOrNil[*ast.Name](data["trait"]),
            Method: data["method"].(*ast.Identifier),
            Attributes: data["attributes"],

        }
    case "StmtTryCatch":
        node = &ast.TryCatchStmt{
            Stmts: asSlice[ast.Stmt](data["stmts"]),
            Catches: asSlice[*ast.CatchStmt](data["catches"]),
            Finally: asTypeOrNil[*ast.FinallyStmt](data["finally"]),
            Attributes: data["attributes"],

        }
    case "StmtUnset":
        node = &ast.UnsetStmt{
            Vars: asSlice[ast.Expr](data["vars"]),
            Attributes: data["attributes"],

        }
    case "StmtUse":
        node = &ast.UseStmt{
            Type: asInt(data["type"]),
            Uses: asSlice[*ast.UseUseStmt](data["uses"]),
            Attributes: data["attributes"],

        }
    case "StmtUseUse":
        node = &ast.UseUseStmt{
            Type: asInt(data["type"]),
            Name: data["name"].(*ast.Name),
            Alias: asTypeOrNil[*ast.Identifier](data["alias"]),
            Attributes: data["attributes"],

        }
    case "StmtWhile":
        node = &ast.WhileStmt{
            Cond: data["cond"].(ast.Expr),
            Stmts: asSlice[ast.Stmt](data["stmts"]),
            Attributes: data["attributes"],

        }
    case "UnionType":
        node = &ast.UnionType{
            Types: asSlice[any](data["types"]),
            Attributes: data["attributes"],

        }
    case "VarLikeIdentifier":
        node = &ast.VarLikeIdentifier{
            Name: data["name"].(string),
            Attributes: data["attributes"],

        }
    case "VariadicPlaceholder":
        node = &ast.VariadicPlaceholder{
            Attributes: data["attributes"],

        }
	}

	return node, nil
}