package php

import (
	"github.com/heyuuu/gophp/compile/ast"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
)

type executeState uint8

const (
	stateNormal executeState = iota
	stateReturn
	stateBreak
	stateContinue
	stateGoto
)

type (
	executeResult interface {
		state() executeState
	}

	returnResult   struct{ retVal Val }
	breakResult    struct{ num int }
	continueResult struct{ num int }
	gotoResult     struct{ label string }
)

func (r returnResult) state() executeState   { return stateReturn }
func (r breakResult) state() executeState    { return stateBreak }
func (r continueResult) state() executeState { return stateContinue }
func (r gotoResult) state() executeState     { return stateGoto }

// private

func (e *Executor) userFunction(fn *types.Function, args []Val) Val {
	// todo 初始化 args
	res := e.stmtList(fn.Stmts())
	switch r := res.(type) {
	case *returnResult:
		return r.retVal
	case *continueResult, *breakResult, *gotoResult:
		panic(perr.Unreachable())
	}
	return types.Null
}

func (e *Executor) stmtList(stmts []ast.Stmt) executeResult {
	var labels = map[string]int{}
	for i, stmt := range stmts {
		if label, ok := stmt.(*ast.LabelStmt); ok {
			labels[label.Name.Name] = i
		}
	}

	l := len(stmts)
	for i := 0; i < l; i++ {
		var result executeResult
		switch x := stmts[i].(type) {
		case *ast.EmptyStmt: // pass
		case *ast.ExprStmt:
			_ = e.expr(x.Expr)
		case *ast.ReturnStmt:
			retVal := e.expr(x.Expr)
			return returnResult{retVal: retVal}
		case *ast.LabelStmt:
			// pass
		case *ast.GotoStmt:
			labelName := x.Name.Name
			result = gotoResult{labelName}
		case *ast.BreakStmt:
			num := 1
			if x.Num != nil {
				num = ZvalGetLong(e.ctx, e.expr(x.Num))
			}
			result = breakResult{num: num}
		case *ast.ContinueStmt:
			num := 1
			if x.Num != nil {
				num = ZvalGetLong(e.ctx, e.expr(x.Num))
			}
			result = continueResult{num: num}
		case *ast.EchoStmt:
			values := e.exprList(x.Exprs)
			for _, value := range values {
				vmEcho(e.ctx, value)
			}
		case *ast.IfStmt:
			result = e.ifStmt(x)
		case *ast.SwitchStmt:
			result = e.switchStmt(x)
		case *ast.ForStmt:
			result = e.forStmt(x)
		case *ast.ForeachStmt:
			result = e.foreachStmt(x)
		case *ast.WhileStmt:
			result = e.whileStmt(x)
		case *ast.DoStmt:
			result = e.doStmt(x)
		case *ast.TryCatchStmt:
			result = e.tryCatchStmt(x)
		case *ast.ConstStmt:
			result = e.constStmt(x)
		case *ast.HaltCompilerStmt:
			// 中断执行，返回结果
			return nil
		case *ast.InlineHTMLStmt:
			e.ctx.WriteString(x.Value)
		// todo
		default:
			panic(perr.Todof("todo executor.stmtList(%T)", x))
		}
		if result != nil {
			switch r := result.(type) {
			case gotoResult:
				// todo goto 能跳到非循环结构内(比如 if)
				if v, ok := labels[r.label]; ok {
					i = v
				} else {
					return r
				}
			default:
				return r
			}
		}
	}
	return nil
}

func (e *Executor) ifStmt(x *ast.IfStmt) executeResult {
	cond := e.expr(x.Cond)
	if ZvalIsTrue(e.ctx, cond) {
		return e.stmtList(x.Stmts)
	}

	for _, elseIfStmt := range x.Elseifs {
		elseIfCond := e.expr(elseIfStmt.Cond)
		if ZvalIsTrue(e.ctx, elseIfCond) {
			return e.stmtList(x.Stmts)
		}
	}

	if x.Else != nil {
		return e.stmtList(x.Else.Stmts)
	}

	return nil
}
func (e *Executor) switchStmt(x *ast.SwitchStmt) executeResult {
	cond := e.expr(x.Cond)

	matchIndex := -1
	for i, caseStmt := range x.Cases {
		if caseStmt.Cond == nil {
			matchIndex = i
			continue
		}
		caseCond := e.expr(caseStmt.Cond)
		if ZvalEquals(e.ctx, cond, caseCond) {
			matchIndex = i
			break
		}
	}
	if matchIndex < 0 {
		return nil
	}

	for _, caseStmt := range x.Cases[matchIndex:] {
		result, stop := e.handleSwitchStmts(caseStmt.Stmts)
		if stop {
			return result
		}
	}
	return nil
}
func (e *Executor) forStmt(x *ast.ForStmt) executeResult {
	// init
	for _, expr := range x.Init {
		e.expr(expr)
	}

	for {
		// cond
		conds := e.exprList(x.Cond)
		if len(conds) != 0 && !ZvalIsTrue(e.ctx, conds[len(conds)-1]) {
			break
		}

		// body
		result, stop := e.handleLoopStmts(x.Stmts)
		if stop {
			return result
		}

		// step
		e.exprList(x.Loop)
	}

	return nil
}
func (e *Executor) foreachStmt(x *ast.ForeachStmt) executeResult {
	variable := e.expr(x.Expr)
	if variable.IsArray() {
		// todo array 懒复制，避免循环时修改数组
		var result executeResult
		var stop bool
		variable.Array().EachEx(func(key types.ArrayKey, value types.Zval) error {
			// foreach(x as $key => $value)
			if x.KeyVar != nil {
				e.assignVariable(x.KeyVar, key.ToZval())
			}
			// todo byRef
			e.assignVariable(x.ValueVar, value)

			// body
			result, stop = e.handleLoopStmts(x.Stmts)
			if stop {
				return lang.BreakErr
			}
			return nil
		})
		return result
	} else {
		panic(perr.Todof("暂未支持非数组的 foreach 操作"))
	}
}
func (e *Executor) whileStmt(x *ast.WhileStmt) executeResult {
	for {
		if !ZvalIsTrue(e.ctx, e.expr(x.Cond)) {
			break
		}

		result, stop := e.handleLoopStmts(x.Stmts)
		if stop {
			return result
		}
	}
	return nil
}
func (e *Executor) doStmt(x *ast.DoStmt) executeResult {
	first := true
	for {
		if !first && !ZvalIsTrue(e.ctx, e.expr(x.Cond)) {
			break
		}
		first = false

		result, stop := e.handleLoopStmts(x.Stmts)
		if stop {
			return result
		}
	}
	return nil
}

func (e *Executor) handleLoopStmts(stmts []ast.Stmt) (result executeResult, stop bool) {
	switch r := e.stmtList(stmts).(type) {
	case nil:
		return nil, false
	case breakResult:
		if r.num > 1 {
			return breakResult{num: r.num - 1}, true
		}
		return nil, true
	case continueResult:
		if r.num > 1 {
			return continueResult{num: r.num - 1}, true
		}
		return nil, false
	default:
		return r, true
	}
}

func (e *Executor) handleSwitchStmts(stmts []ast.Stmt) (result executeResult, stop bool) {
	switch r := e.stmtList(stmts).(type) {
	case nil:
		return nil, false
	case breakResult:
		if r.num > 1 {
			return breakResult{num: r.num - 1}, true
		}
		return nil, true
	default:
		return r, true
	}
}

func (e *Executor) tryCatchStmt(x *ast.TryCatchStmt) executeResult { panic(perr.Todof("TryCatchStmt")) }
func (e *Executor) constStmt(x *ast.ConstStmt) executeResult       { panic(perr.Todof("ConstStmt")) }

func (e *Executor) exprList(exprs []ast.Expr) []Val {
	values := make([]Val, len(exprs))
	for i, expr := range exprs {
		values[i] = e.expr(expr)
	}
	return values
}

func (e *Executor) expr(expr ast.Expr) Val {
	switch x := expr.(type) {
	case *ast.IntLit:
		return Long(x.Value)
	case *ast.FloatLit:
		return Double(x.Value)
	case *ast.StringLit:
		return String(x.Value)
	case *ast.ArrayExpr:
		return e.executeArrayExpr(x)
	case *ast.ClosureExpr:
		return e.executeClosureExpr(x)
	case *ast.ClosureUseExpr:
		return e.executeClosureUseExpr(x)
	case *ast.ArrowFunctionExpr:
		return e.executeArrowFunctionExpr(x)
	case *ast.IndexExpr:
		return e.executeIndexExpr(x)
	case *ast.CastExpr:
		return e.executeCastExpr(x)
	case *ast.UnaryExpr:
		return e.executeUnaryExpr(x)
	case *ast.BinaryOpExpr:
		return e.executeBinaryOpExpr(x)
	case *ast.AssignExpr:
		return e.executeAssignExpr(x)
	case *ast.AssignOpExpr:
		return e.executeAssignOpExpr(x)
	case *ast.AssignRefExpr:
		return e.executeAssignRefExpr(x)
	case *ast.IssetExpr:
		return e.executeIssetExpr(x)
	case *ast.EmptyExpr:
		return e.executeEmptyExpr(x)
	case *ast.EvalExpr:
		return e.executeEvalExpr(x)
	case *ast.IncludeExpr:
		return e.executeIncludeExpr(x)
	case *ast.CloneExpr:
		return e.executeCloneExpr(x)
	case *ast.ErrorSuppressExpr:
		return e.executeErrorSuppressExpr(x)
	case *ast.ExitExpr:
		return e.executeExitExpr(x)
	case *ast.ConstFetchExpr:
		return e.executeConstFetchExpr(x)
	case *ast.ClassConstFetchExpr:
		return e.executeClassConstFetchExpr(x)
	case *ast.MagicConstExpr:
		return e.executeMagicConstExpr(x)
	case *ast.InstanceofExpr:
		return e.executeInstanceofExpr(x)
	case *ast.ListExpr:
		return e.executeListExpr(x)
	case *ast.PrintExpr:
		return e.executePrintExpr(x)
	case *ast.PropertyFetchExpr:
		return e.executePropertyFetchExpr(x)
	case *ast.StaticPropertyFetchExpr:
		return e.executeStaticPropertyFetchExpr(x)
	case *ast.ShellExecExpr:
		return e.executeShellExecExpr(x)
	case *ast.TernaryExpr:
		return e.executeTernaryExpr(x)
	case *ast.ThrowExpr:
		return e.executeThrowExpr(x)
	case *ast.VariableExpr:
		return e.executeVariableExpr(x)
	case *ast.YieldExpr:
		return e.executeYieldExpr(x)
	case *ast.YieldFromExpr:
		return e.executeYieldFromExpr(x)
	case *ast.FuncCallExpr:
		return e.executeFuncCallExpr(x)
	case *ast.NewExpr:
		return e.executeNewExpr(x)
	case *ast.MethodCallExpr:
		return e.executeMethodCallExpr(x)
	case *ast.StaticCallExpr:
		return e.executeStaticCallExpr(x)
	case *ast.ArrayItemExpr:
		panic(perr.Todof("unexpected execute type: %T", x))
	default:
		panic(perr.Todof("todo executor.executeExpr(%T)", x))
	}
}

func (e *Executor) executeBinaryOpExpr(expr *ast.BinaryOpExpr) (val Val) {
	// && / || / ?? 操作比较特殊，右表达式节点可能不会执行
	switch expr.Op {
	case ast.BinaryOpBooleanAnd: // &&
		left := e.expr(expr.Left)
		right := func() Val { return e.expr(expr.Right) }
		return OpBooleanAnd(e.ctx, left, right)
	case ast.BinaryOpBooleanOr: // ||
		left := e.expr(expr.Left)
		right := func() Val { return e.expr(expr.Right) }
		return OpBooleanAnd(e.ctx, left, right)
	case ast.BinaryOpCoalesce: // ??
		left := e.expr(expr.Left)
		right := func() Val { return e.expr(expr.Right) }
		return OpCoalesce(e.ctx, left, right)
	}

	// common
	left := e.expr(expr.Left)
	right := e.expr(expr.Right)

	switch expr.Op {
	case ast.BinaryOpPlus: // +
		return OpAdd(e.ctx, left, right)
	case ast.BinaryOpMinus: // -
		return OpSub(e.ctx, left, right)
	case ast.BinaryOpMul: // *
		return OpMul(e.ctx, left, right)
	case ast.BinaryOpDiv: // /
		return OpDiv(e.ctx, left, right)
	case ast.BinaryOpMod: // %
		return OpMod(e.ctx, left, right)
	case ast.BinaryOpPow: // **
		return OpPow(e.ctx, left, right)
	case ast.BinaryOpBitwiseAnd: // &
		return OpBitwiseAnd(e.ctx, left, right)
	case ast.BinaryOpBitwiseOr: // n|
		return OpBitwiseOr(e.ctx, left, right)
	case ast.BinaryOpBitwiseXor: // ^
		return OpBitwiseXor(e.ctx, left, right)
	case ast.BinaryOpConcat: // .
		return OpConcat(e.ctx, left, right)
	case ast.BinaryOpEqual: // ==
		return OpEqual(e.ctx, left, right)
	case ast.BinaryOpGreater: // >
		return OpGreater(e.ctx, left, right)
	case ast.BinaryOpGreaterOrEqual: // >=
		return OpGreaterOrEqual(e.ctx, left, right)
	case ast.BinaryOpIdentical: // ===
		return OpIdentical(e.ctx, left, right)
	case ast.BinaryOpBooleanXor: // xor
		return OpBooleanXor(e.ctx, left, right)
	case ast.BinaryOpNotEqual: // !=
		return OpNotEqual(e.ctx, left, right)
	case ast.BinaryOpNotIdentical: // !==
		return OpNotIdentical(e.ctx, left, right)
	case ast.BinaryOpShiftLeft: // <<
		return OpSL(e.ctx, left, right)
	case ast.BinaryOpShiftRight: // >>
		return OpSR(e.ctx, left, right)
	case ast.BinaryOpSmaller: // <
		return OpSmaller(e.ctx, left, right)
	case ast.BinaryOpSmallerOrEqual: // <=
		return OpSmallerOrEqual(e.ctx, left, right)
	case ast.BinaryOpSpaceship: // <=>
		return OpSpaceship(e.ctx, left, right)
	default:
		panic(perr.Unreachable())
	}
}
func (e *Executor) executeArrayExpr(expr *ast.ArrayExpr) Val {
	arr := types.NewArrayCap(len(expr.Items))
	for _, item := range expr.Items {
		if item.ByRef {
			// todo item byref
			perr.Panic("todo item byref")
		} else if item.Unpack && item.Key != nil {
			// todo item unpack with key
			perr.Panic("todo item unpack with key")
		}

		if item.Key != nil {
			key := e.expr(item.Key)
			value := e.expr(item.Value)
			arrayKey := ZvalToArrayKey(e.ctx, key)
			arr.Add(arrayKey, value)
		} else {
			value := e.expr(item.Value)
			arr.Append(value)
		}
	}
	return Array(arr)
}
func (e *Executor) executeClosureExpr(expr *ast.ClosureExpr) Val {
	panic(perr.Todof("executeClosureExpr"))
}
func (e *Executor) executeClosureUseExpr(expr *ast.ClosureUseExpr) Val {
	panic(perr.Todof("executeClosureUseExpr"))
}
func (e *Executor) executeArrowFunctionExpr(expr *ast.ArrowFunctionExpr) Val {
	panic(perr.Todof("executeArrowFunctionExpr"))
}
func (e *Executor) executeIndexExpr(expr *ast.IndexExpr) Val {
	if expr.Dim == nil {
		panic(perr.Todof("PHP Fatal error:  Cannot use [] for reading"))
	}

	arr := e.expr(expr.Var)
	dim := e.expr(expr.Dim)
	key := ZvalToArrayKey(e.ctx, dim)
	value := e.arrayGet(arr, key)
	if value.IsUndef() {
		if key.IsStrKey() {
			panic(perr.Todof(`Warning: Undefined array key "%v" in`, key.StrKey()))
		} else {
			panic(perr.Todof(`Warning: Undefined array key %d in`, key.IdxKey()))
		}
		return types.Null
	}
	return value
}
func (e *Executor) executeCastExpr(expr *ast.CastExpr) Val {
	switch expr.Kind {
	case ast.CastArray:
	case ast.CastBool:
	case ast.CastDouble:
	case ast.CastInt:
	case ast.CastObject:
	case ast.CastString:
	case ast.CastUnset:
	}
	panic(perr.Todof("executeCastExpr: %v", expr))
}
func (e *Executor) executeUnaryExpr(expr *ast.UnaryExpr) Val {
	// todo 考虑是否需要用 unary 原生替代 v = v + 1 的模拟

	var oldValue = e.expr(expr.Var)
	var newValue Val
	var useOldValue = false

	switch expr.Op {
	case ast.UnaryOpPlus:
		newValue = OpMul(e.ctx, oldValue, Long(1))
	case ast.UnaryOpMinus:
		newValue = OpMul(e.ctx, oldValue, Long(-1))
	case ast.UnaryOpBooleanNot:
		newValue = OpBooleanNot(e.ctx, oldValue)
	case ast.UnaryOpBitwiseNot:
		newValue = OpBitwiseNot(e.ctx, oldValue)
	case ast.UnaryOpPreInc:
		newValue = OpAdd(e.ctx, oldValue, Long(1))
	case ast.UnaryOpPreDec:
		newValue = OpSub(e.ctx, oldValue, Long(1))
	case ast.UnaryOpPostInc:
		newValue = OpAdd(e.ctx, oldValue, Long(1))
		useOldValue = true
	case ast.UnaryOpPostDec:
		newValue = OpSub(e.ctx, oldValue, Long(1))
		useOldValue = true
	default:
		panic(perr.Internalf("Unexpected ast.UnaryExpr.Op: %+v", expr.Op))
	}

	e.assignVariable(expr.Var, newValue)

	if useOldValue {
		return oldValue
	} else {
		return newValue
	}
}
func (e *Executor) executeAssignExpr(expr *ast.AssignExpr) Val {
	value := e.expr(expr.Expr)
	e.assignVariable(expr.Var, value)
	return value
}
func (e *Executor) executeAssignOpExpr(expr *ast.AssignOpExpr) Val {
	panic(perr.Todof("executeAssignOpExpr"))
}
func (e *Executor) executeAssignRefExpr(expr *ast.AssignRefExpr) Val {
	panic(perr.Todof("executeAssignRefExpr"))
}
func (e *Executor) executeIssetExpr(expr *ast.IssetExpr) Val {
	panic(perr.Todof("executeIssetExpr"))
}
func (e *Executor) executeEmptyExpr(expr *ast.EmptyExpr) Val {
	panic(perr.Todof("executeEmptyExpr"))
}
func (e *Executor) executeEvalExpr(expr *ast.EvalExpr) Val {
	panic(perr.Todof("executeEvalExpr"))
}
func (e *Executor) executeIncludeExpr(expr *ast.IncludeExpr) Val {
	panic(perr.Todof("executeIncludeExpr"))
}
func (e *Executor) executeCloneExpr(expr *ast.CloneExpr) Val {
	panic(perr.Todof("executeCloneExpr"))
}
func (e *Executor) executeErrorSuppressExpr(expr *ast.ErrorSuppressExpr) Val {
	panic(perr.Todof("executeErrorSuppressExpr"))
}
func (e *Executor) executeExitExpr(expr *ast.ExitExpr) Val {
	panic(perr.Todof("executeExitExpr"))
}
func (e *Executor) executeConstFetchExpr(expr *ast.ConstFetchExpr) Val {
	name := expr.Name.ToCodeString()
	c := GetConstant(e.ctx, name)
	if c == nil {
		panic(perr.Todof("const not defined: " + name))
	}
	return c.Value()
}
func (e *Executor) executeClassConstFetchExpr(expr *ast.ClassConstFetchExpr) Val {
	panic(perr.Todof("executeClassConstFetchExpr"))
}
func (e *Executor) executeMagicConstExpr(expr *ast.MagicConstExpr) Val {
	panic(perr.Todof("executeMagicConstExpr"))
}
func (e *Executor) executeInstanceofExpr(expr *ast.InstanceofExpr) Val {
	panic(perr.Todof("executeInstanceofExpr"))
}
func (e *Executor) executeListExpr(expr *ast.ListExpr) Val {
	panic(perr.Todof("executeListExpr"))
}
func (e *Executor) executePrintExpr(expr *ast.PrintExpr) Val {
	panic(perr.Todof("executePrintExpr"))
}
func (e *Executor) executePropertyFetchExpr(expr *ast.PropertyFetchExpr) Val {
	panic(perr.Todof("executePropertyFetchExpr"))
}
func (e *Executor) executeStaticPropertyFetchExpr(expr *ast.StaticPropertyFetchExpr) Val {
	panic(perr.Todof("executeStaticPropertyFetchExpr"))
}
func (e *Executor) executeShellExecExpr(expr *ast.ShellExecExpr) Val {
	panic(perr.Todof("executeShellExecExpr"))
}
func (e *Executor) executeTernaryExpr(expr *ast.TernaryExpr) Val {
	panic(perr.Todof("executeTernaryExpr"))
}
func (e *Executor) executeThrowExpr(expr *ast.ThrowExpr) Val {
	panic(perr.Todof("executeThrowExpr"))
}
func (e *Executor) executeVariableExpr(expr *ast.VariableExpr) Val {
	name := e.executeVariableName(expr.Name)
	// todo undefined warning
	symbols := e.executeData.symbols
	return symbols.Get(name)
}

func (e *Executor) executeVariableName(nameNode ast.Node) string {
	switch x := nameNode.(type) {
	case *ast.Ident:
		return x.Name
	case *ast.VariableExpr:
		nameVal := e.expr(x)
		return nameVal.String()
	default:
		panic(perr.Todof("unexpected VariableExpr.Name type: %T, %+v", nameNode, nameNode))
	}
}

func (e *Executor) executeYieldExpr(expr *ast.YieldExpr) Val {
	panic(perr.Todof("executeYieldExpr"))
}
func (e *Executor) executeYieldFromExpr(expr *ast.YieldFromExpr) Val {
	panic(perr.Todof("executeYieldFromExpr"))
}
func (e *Executor) executeFuncCallExpr(expr *ast.FuncCallExpr) Val {
	var name Val
	switch nameAst := expr.Name.(type) {
	case *ast.Name:
		name = String(nameAst.ToString())
	case ast.Expr:
		name = e.expr(nameAst)
	default:
		panic(perr.Unreachable())
	}

	var fn *types.Function
	if name.IsString() {
		fn = e.initStringCall(name.String())
	} else {
		// todo 各种类型的 function name 处理
		panic(perr.Todof("executeFuncCallExpr"))
	}

	args := make([]Val, 0, len(expr.Args))
	for _, arg := range expr.Args {
		argVal := e.expr(arg.Value)

		if !arg.Unpack {
			args = append(args, argVal)
		} else {
			// todo unpack args
			panic("todo unpack args")
		}
	}

	return e.function(fn, args)
}
func (e *Executor) executeNewExpr(expr *ast.NewExpr) Val {
	panic(perr.Todof("executeNewExpr"))
}
func (e *Executor) executeMethodCallExpr(expr *ast.MethodCallExpr) Val {
	panic(perr.Todof("executeMethodCallExpr"))
}
func (e *Executor) executeStaticCallExpr(expr *ast.StaticCallExpr) Val {
	panic(perr.Todof("executeStaticCallExpr"))
}

type executeVariable interface {
	Get() *types.Zval
	Set(v *types.Zval)
}

type executeVariableFunc struct {
	Getter func() *types.Zval
	Setter func(v *types.Zval)
}

func (e executeVariableFunc) Get() *types.Zval  { return e.Getter() }
func (e executeVariableFunc) Set(v *types.Zval) { e.Setter(v) }

func (e *Executor) assignVariable(variable ast.Expr, value types.Zval) {
	switch v := variable.(type) {
	case *ast.VariableExpr:
		name := e.executeVariableName(v.Name)
		symbols := e.executeData.symbols
		symbols.Set(name, value)
	case *ast.IndexExpr:
		arr := e.getOrInitArray(v.Var)
		// todo 转 arr 处理
		if v.Dim == nil {
			e.arrayAppend(arr, value)
		} else {
			dim := e.expr(v.Dim)
			key := e.zvalToArrayKey(dim)
			e.arrayUpdate(arr, key, value)
		}
	default:
		panic(perr.Todof("unsupported AssignExpr.Var type: %T, %+v", v, v))
	}
}

func (e *Executor) getOrInitArray(variable ast.Expr) Val {
	switch v := variable.(type) {
	case *ast.VariableExpr:
		name := e.executeVariableName(v.Name)
		symbols := e.executeData.symbols
		if !symbols.Isset(name) {
			symbols.Set(name, types.InitZvalArray())
		}
		return symbols.Get(name)
	case *ast.IndexExpr:
		arrVar := e.getOrInitArray(v.Var)
		if v.Dim == nil {
			result := types.InitZvalArray()
			e.arrayAppend(arrVar, result)
			return result
		} else {
			dim := e.expr(v.Dim)
			key := e.zvalToArrayKey(dim)
			result := e.arrayGet(arrVar, key)
			if result.IsUndef() {
				result = types.InitZvalArray()
				e.arrayUpdate(result, key, result)
			}
			return result
		}
	default:
		panic(perr.Todof("unsupported AssignExpr.Var type: %T, %+v", v, v))
	}
}

func (e *Executor) arrayGet(arr Val, key types.ArrayKey) Val {
	switch arr.Type() {
	case types.IsArray:
		return arr.Array().Find(key)
	// todo ArrayAccess
	default:
		panic(perr.Todof("unsupported e.arrayGet arr type: %s", types.ZvalGetType(arr)))
	}
}

func (e *Executor) arrayAppend(arr Val, value Val) {
	switch arr.Type() {
	case types.IsArray:
		arr.Array().Append(value)
	// todo ArrayAccess
	default:
		panic(perr.Todof("unsupported e.arrayAppend arr type: %s", types.ZvalGetType(arr)))
	}
}
func (e *Executor) arrayUpdate(arr Val, key types.ArrayKey, value Val) {
	switch arr.Type() {
	case types.IsArray:
		arr.Array().Update(key, value)
	// todo ArrayAccess
	default:
		panic(perr.Todof("unsupported e.arrayUpdate arr type: %s", types.ZvalGetType(arr)))
	}
}

func (e *Executor) getVariable(variable ast.Expr) executeVariable {
	switch v := variable.(type) {
	case *ast.VariableExpr:
		var name string
		switch nameNode := v.Name.(type) {
		case *ast.Ident:
			name = nameNode.Name
		case ast.Expr:
			nameVal := e.expr(nameNode)
			name = nameVal.String()
		default:
			panic(perr.Todof("unexpected VariableExpr.Name type: %T, %+v", v, v))
		}
		return executeVariableFunc{
			Getter: func() *types.Zval {
				// todo
				return types.NewZvalString("$" + name)
			},
			Setter: func(v *types.Zval) {
				// todo
			},
		}
	case *ast.IndexExpr:
		if v.Dim == nil {
			return executeVariableFunc{
				Getter: func() *types.Zval {
					// todo
					//variable := e.getVariable(v.Var)
					return types.NewZvalArray(nil)
				},
				Setter: func(v *types.Zval) {
					// todo
				},
			}
		} else {
			//dim := e.expr(v.Dim)
			return executeVariableFunc{
				Getter: func() *types.Zval {
					// todo
					return types.NewZvalArray(nil)
				},
				Setter: func(v *types.Zval) {
					// todo
				},
			}
		}
	default:
		panic(perr.Todof("unsupported AssignExpr.Var type: %T, %+v", v, v))
	}
}
