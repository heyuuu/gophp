package php

import (
	"fmt"
	"github.com/heyuuu/gophp/compile/ast"
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/php/assert"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
)

type (
	execResult interface {
		execResult()
	}

	returnResult   struct{ retVal types.Zval }
	breakResult    struct{ num int }
	continueResult struct{ num int }
	gotoResult     struct{ label string }
)

func (r returnResult) execResult()   {}
func (r breakResult) execResult()    {}
func (r continueResult) execResult() {}
func (r gotoResult) execResult()     {}

// executor
type Executor struct {
	ctx         *Context
	executeData *ExecuteData
}

func NewExecutor(ctx *Context) *Executor {
	return &Executor{ctx: ctx}
}

func (e *Executor) Execute(fn *types.Function) (retVal types.Zval, ret error) {
	return e.doCall(fn, nil, nil), nil
}

func (e *Executor) function(fn *types.Function, args []types.Zval) types.Zval {
	assert.Assert(fn != nil)

	// push && pop executeData
	ex := NewExecuteData(e.ctx, fn, args)
	e.ctx.EG().PushExecuteData(ex)
	defer func() {
		popEx := e.ctx.EG().PopExecuteData()
		assert.AssertEx(ex == popEx, "push & pop executeData 的 executeData 对象不相同，执行堆栈可能有错误")
	}()

	var retval types.Zval
	if fn.IsInternalFunction() {
		if handler, ok := fn.Handler().(ZifHandler); ok {
			handler(ex, &retval)
		} else {
			perr.Panic(fmt.Sprintf("不支持的内部函数 handler 类型: %T", fn.Handler()))
		}
	} else {
		retval = e.userFunction(fn, args)
	}
	if retval.IsUndef() {
		retval = types.Null
	}
	return retval
}

func (e *Executor) initStringCall(name string) *types.Function {
	// todo ZendInitDynamicCallString
	fn := e.ctx.EG().FindFunction(name)
	if fn == nil {
		ThrowError(e.ctx, nil, fmt.Sprintf("Call to undefined function %s()", name))
		return nil
	}
	return fn
}

func (e *Executor) userFunction(fn *types.Function, args []types.Zval) types.Zval {
	// todo 初始化 args
	for i, argInfo := range fn.ArgInfos() {
		if i < len(args) {
			e.currSymbols().Set(argInfo.Name, args[i])
		}
	}

	res := e.stmtList(fn.Stmts())
	switch r := res.(type) {
	case returnResult:
		return r.retVal
	case continueResult, breakResult, gotoResult:
		panic(perr.Unreachable())
	}
	return types.Null
}

func (e *Executor) globalSymbols() ISymtable {
	panic(perr.Todof("globalSymbols"))
}

func (e *Executor) currSymbols() ISymtable {
	return e.ctx.CurrEX().Symbols()
}

// -- echo node types

func (e *Executor) stmtList(stmts []ast.Stmt) execResult {
	var labels = map[string]int{}
	for i, stmt := range stmts {
		if label, ok := stmt.(*ast.LabelStmt); ok {
			labels[label.Name.Name] = i
		}
	}

	l := len(stmts)
	for i := 0; i < l; i++ {
		result := e.stmt(stmts[i])
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

func (e *Executor) stmt(stmt ast.Stmt) execResult {
	switch x := stmt.(type) {
	case *ast.EmptyStmt:
		// pass
	case *ast.BlockStmt:
		return e.stmtList(x.List)
	case *ast.ExprStmt:
		_ = e.expr(x.Expr)
	case *ast.ReturnStmt:
		retVal := e.expr(x.Expr)
		return returnResult{retVal: retVal}
	case *ast.LabelStmt:
		// pass
	case *ast.GotoStmt:
		labelName := x.Name.Name
		return gotoResult{labelName}
	case *ast.IfStmt:
		return e.ifStmt(x)
	case *ast.SwitchStmt:
		return e.switchStmt(x)
	case *ast.ForStmt:
		return e.forStmt(x)
	case *ast.ForeachStmt:
		return e.foreachStmt(x)
	case *ast.BreakStmt:
		num := 1
		if x.Num != nil {
			num = ZvalGetLong(e.ctx, e.expr(x.Num))
		}
		return breakResult{num: num}
	case *ast.ContinueStmt:
		num := 1
		if x.Num != nil {
			num = ZvalGetLong(e.ctx, e.expr(x.Num))
		}
		return continueResult{num: num}
	case *ast.WhileStmt:
		return e.whileStmt(x)
	case *ast.DoStmt:
		return e.doStmt(x)
	case *ast.TryCatchStmt:
		return e.tryCatchStmt(x)
	case *ast.ConstStmt:
		return e.constStmt(x)
	case *ast.EchoStmt:
		return e.echoStmt(x)
	case *ast.GlobalStmt:
		return e.globalStmt(x)
	case *ast.HaltCompilerStmt:
		// 中断执行，返回结果
		return returnResult{}
	case *ast.InlineHTMLStmt:
		return e.inlineHTMLStmt(x)
	case *ast.StaticStmt:
		return e.staticStmt(x)
	case *ast.UnsetStmt:
		return e.unsetStmt(x)
	case *ast.DeclareStmt:
		return e.declareStmt(x)
	case *ast.DeclareDeclareStmt:
		return e.declareDeclareStmt(x)
	case *ast.NamespaceStmt:
		return e.namespaceStmt(x)
	case *ast.FunctionStmt:
		return e.functionStmt(x)
	case *ast.InterfaceStmt:
		return e.interfaceStmt(x)
	case *ast.ClassStmt:
		return e.classStmt(x)
	case *ast.TraitStmt:
		return e.traitStmt(x)
	case *ast.TraitUseStmt:
		return e.traitUseStmt(x)
	case *ast.TraitUseAdaptationAliasStmt:
		return e.traitUseAdaptationAliasStmt(x)
	case *ast.TraitUseAdaptationPrecedenceStmt:
		return e.traitUseAdaptationPrecedenceStmt(x)
	default:
		panic(perr.Internalf("Unexpected stmt type: %+v", x))
	}
	return nil
}

func (e *Executor) ifStmt(x *ast.IfStmt) execResult {
	cond := e.expr(x.Cond)
	if ZvalIsTrue(e.ctx, cond) {
		return e.stmtList(x.Stmts)
	}

	for _, elseIfStmt := range x.Elseifs {
		elseIfCond := e.expr(elseIfStmt.Cond)
		if ZvalIsTrue(e.ctx, elseIfCond) {
			return e.stmtList(elseIfStmt.Stmts)
		}
	}

	if x.Else != nil {
		return e.stmtList(x.Else.Stmts)
	}

	return nil
}

func (e *Executor) switchStmt(x *ast.SwitchStmt) execResult {
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

func (e *Executor) handleSwitchStmts(stmts []ast.Stmt) (result execResult, stop bool) {
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

func (e *Executor) forStmt(x *ast.ForStmt) execResult {
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

func (e *Executor) foreachStmt(x *ast.ForeachStmt) execResult {
	variable := e.exprDeref(x.Expr)
	if variable.IsArray() {
		// todo array 懒复制，避免循环时修改数组
		var result execResult
		var stop bool
		variable.Array().EachEx(func(key types.ArrayKey, value types.Zval) error {
			// foreach(x as $key => $value)
			if x.KeyVar != nil {
				e.variableSet(x.KeyVar, key.ToZval())
			}
			// todo byRef
			e.variableSet(x.ValueVar, value)

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

func (e *Executor) whileStmt(x *ast.WhileStmt) execResult {
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

func (e *Executor) doStmt(x *ast.DoStmt) execResult {
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

func (e *Executor) handleLoopStmts(stmts []ast.Stmt) (result execResult, stop bool) {
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

func (e *Executor) tryCatchStmt(x *ast.TryCatchStmt) execResult {
	panic(perr.Todof("e.tryCatchStmt"))
}

func (e *Executor) constStmt(x *ast.ConstStmt) execResult {
	// todo 确认必须在 global 域内执行
	assert.Assert(x.NamespacedName != nil)
	name := x.NamespacedName.ToCodeString()
	value := e.expr(x.Value)

	RegisterConstant(e.ctx, 0, name, value)
	return nil
}

func (e *Executor) echoStmt(x *ast.EchoStmt) execResult {
	for _, expr := range x.Exprs {
		PrintZval(e.ctx, e.expr(expr))
	}
	return nil
}

func (e *Executor) globalStmt(x *ast.GlobalStmt) execResult {
	panic(perr.Todof("e.globalStmt"))
}

func (e *Executor) inlineHTMLStmt(x *ast.InlineHTMLStmt) execResult {
	e.ctx.WriteString(x.Value)
	return nil
}

func (e *Executor) staticStmt(x *ast.StaticStmt) execResult {
	for _, staticVarStmt := range x.Vars {
		_ = staticVarStmt
	}
	panic(perr.Todof("e.staticStmt"))
}

func (e *Executor) unsetStmt(x *ast.UnsetStmt) execResult {
	for _, var_ := range x.Vars {
		e.variableRef(var_).Unset()
	}
	return nil
}

func (e *Executor) declareStmt(x *ast.DeclareStmt) execResult {
	panic(perr.Todof("e.declareStmt"))
}

func (e *Executor) declareDeclareStmt(x *ast.DeclareDeclareStmt) execResult {
	panic(perr.Todof("e.declareDeclareStmt"))
}

func (e *Executor) namespaceStmt(x *ast.NamespaceStmt) execResult {
	panic(perr.Todof("e.namespaceStmt"))
}

func (e *Executor) functionStmt(x *ast.FunctionStmt) execResult {
	var name string
	if x.NamespacedName != nil {
		name = x.NamespacedName.ToCodeString()
	} else {
		name = x.Name.Name
	}

	var argInfos []types.ArgInfo
	if len(x.Params) > 0 {
		argInfos = make([]types.ArgInfo, len(x.Params))
		for i, param := range x.Params {
			argInfos[i] = types.ArgInfo{
				Name:     param.Var.Name.(*ast.Ident).Name,
				ByRef:    param.ByRef,
				Variadic: param.Variadic,
			}
		}
	}

	fn := types.NewAstFunction(name, argInfos, x.Stmts)
	RegisterFunction(e.ctx, name, fn)

	return nil
}

func (e *Executor) interfaceStmt(x *ast.InterfaceStmt) execResult {
	panic(perr.Todof("e.interfaceStmt"))
}

func (e *Executor) classStmt(x *ast.ClassStmt) execResult {
	var decl types.UserClassDecl

	decl.Name = x.NamespacedName.ToCodeString()

	for _, stmt := range x.Stmts {
		switch s := stmt.(type) {
		case *ast.ClassConstStmt:
			constant := types.NewClassConstant(
				s.Name.Name,
				e.expr(s.Value),
				"",
				0,
			)
			decl.Constants = append(decl.Constants, constant)
		case *ast.PropertyStmt:
			var defaultValue types.Zval
			if s.Default != nil {
				defaultValue = e.expr(s.Default)
			} else {
				defaultValue = UninitializedZval()
			}
			property := types.NewPropertyInfo(
				s.Name.Name,
				0,
				nil,
				defaultValue,
			)
			decl.Properties = append(decl.Properties, property)
		case *ast.ClassMethodStmt:
			method := types.NewAstFunction(
				s.Name.Name,
				nil,
				s.Stmts,
			)
			decl.Methods = append(decl.Methods, method)
		default:
			panic(perr.Todof("class stmt type: %T", s))
		}
	}

	_ = RegisterUserClass(e.ctx, &decl)

	return nil
}

func (e *Executor) propertyStmt(x *ast.PropertyStmt) execResult {
	panic(perr.Todof("e.propertyStmt"))
}

func (e *Executor) classMethodStmt(x *ast.ClassMethodStmt) execResult {
	panic(perr.Todof("e.classMethodStmt"))
}

func (e *Executor) traitStmt(x *ast.TraitStmt) execResult {
	panic(perr.Todof("e.traitStmt"))
}

func (e *Executor) traitUseStmt(x *ast.TraitUseStmt) execResult {
	panic(perr.Todof("e.traitUseStmt"))
}

func (e *Executor) traitUseAdaptationAliasStmt(x *ast.TraitUseAdaptationAliasStmt) execResult {
	panic(perr.Todof("e.traitUseAdaptationAliasStmt"))
}

func (e *Executor) traitUseAdaptationPrecedenceStmt(x *ast.TraitUseAdaptationPrecedenceStmt) execResult {
	panic(perr.Todof("e.traitUseAdaptationPrecedenceStmt"))
}

func (e *Executor) exprList(exprs []ast.Expr) []types.Zval {
	values := make([]types.Zval, len(exprs))
	for i, expr := range exprs {
		values[i] = e.expr(expr)
	}
	return values
}

func (e *Executor) exprDeref(expr ast.Expr) types.Zval {
	return e.expr(expr).DeRef()
}

func (e *Executor) expr(expr ast.Expr) types.Zval {
	switch x := expr.(type) {
	case *ast.IntLit:
		return Long(x.Value)
	case *ast.FloatLit:
		return Double(x.Value)
	case *ast.StringLit:
		return String(x.Value)
	case *ast.ArrayExpr:
		return e.arrayExpr(x)
	case *ast.ClosureExpr:
		return e.closureExpr(x)
	case *ast.ClosureUseExpr:
		return e.closureUseExpr(x)
	case *ast.ArrowFunctionExpr:
		return e.arrowFunctionExpr(x)
	case *ast.IndexExpr:
		return e.indexExpr(x)
	case *ast.CastExpr:
		return e.castExpr(x)
	case *ast.UnaryExpr:
		return e.unaryExpr(x)
	case *ast.BinaryOpExpr:
		return e.binaryOpExpr(x)
	case *ast.AssignExpr:
		return e.assignExpr(x)
	case *ast.AssignOpExpr:
		return e.assignOpExpr(x)
	case *ast.AssignRefExpr:
		return e.assignRefExpr(x)
	case *ast.IssetExpr:
		return e.issetExpr(x)
	case *ast.EmptyExpr:
		return e.emptyExpr(x)
	case *ast.EvalExpr:
		return e.evalExpr(x)
	case *ast.IncludeExpr:
		return e.includeExpr(x)
	case *ast.CloneExpr:
		return e.cloneExpr(x)
	case *ast.ErrorSuppressExpr:
		return e.errorSuppressExpr(x)
	case *ast.ExitExpr:
		return e.exitExpr(x)
	case *ast.ConstFetchExpr:
		return e.constFetchExpr(x)
	case *ast.ClassConstFetchExpr:
		return e.classConstFetchExpr(x)
	case *ast.MagicConstExpr:
		return e.magicConstExpr(x)
	case *ast.InstanceofExpr:
		return e.instanceofExpr(x)
	case *ast.ListExpr:
		return e.listExpr(x)
	case *ast.PrintExpr:
		return e.printExpr(x)
	case *ast.PropertyFetchExpr:
		return e.propertyFetchExpr(x)
	case *ast.StaticPropertyFetchExpr:
		return e.staticPropertyFetchExpr(x)
	case *ast.ShellExecExpr:
		return e.shellExecExpr(x)
	case *ast.TernaryExpr:
		return e.ternaryExpr(x)
	case *ast.ThrowExpr:
		return e.throwExpr(x)
	case *ast.VariableExpr:
		return e.variableExpr(x)
	case *ast.YieldExpr:
		return e.yieldExpr(x)
	case *ast.YieldFromExpr:
		return e.yieldFromExpr(x)
	case *ast.FuncCallExpr:
		return e.funcCallExpr(x)
	case *ast.NewExpr:
		return e.newExpr(x)
	case *ast.MethodCallExpr:
		return e.methodCallExpr(x)
	case *ast.StaticCallExpr:
		return e.staticCallExpr(x)
	default:
		panic(perr.Todof("todo executor.executeExpr(%T)", x))
	}
}

func (e *Executor) arrayExpr(expr *ast.ArrayExpr) types.Zval {
	arr := types.NewArrayCap(len(expr.Items))
	for _, item := range expr.Items {
		if item.Unpack && item.Key != nil {
			// todo item unpack with key
			perr.Panic("todo item unpack with key")
		}

		var val types.Zval
		if item.ByRef {
			val = e.variableRefZval(item.Value)
		} else {
			val = e.expr(item.Value).DeRef()
		}

		if item.Key != nil {
			if key, ok := ZvalToArrayKey(e.ctx, e.expr(item.Key)); ok {
				arr.Update(key, val)
			}
		} else {
			arr.Append(val)
		}
	}

	return Array(arr)
}

func (e *Executor) closureExpr(expr *ast.ClosureExpr) types.Zval {
	panic(perr.Todof("e.closureExpr"))
}

func (e *Executor) closureUseExpr(expr *ast.ClosureUseExpr) types.Zval {
	panic(perr.Todof("e.closureUseExpr"))
}

func (e *Executor) arrowFunctionExpr(expr *ast.ArrowFunctionExpr) types.Zval {
	panic(perr.Todof("e.arrowFunctionExpr"))
}

func (e *Executor) indexExpr(expr *ast.IndexExpr) types.Zval {
	if expr.Dim == nil {
		panic(perr.Todof("PHP Fatal error:  Cannot use [] for reading"))
	}

	arr := e.expr(expr.Var)
	dim := e.expr(expr.Dim)
	key, ok := ZvalToArrayKey(e.ctx, dim)
	if !ok {
		return UninitializedZval()
	}

	value := e.arrayGet(arr, key)
	if value.IsUndef() {
		if key.IsStrKey() {
			panic(perr.Todof(`Warning: Undefined array key "%v" in`, key.StrKey()))
		} else {
			panic(perr.Todof(`Warning: Undefined array key %d in`, key.IdxKey()))
		}
		return UninitializedZval()
	}
	return value
}

func (e *Executor) castExpr(expr *ast.CastExpr) types.Zval {
	switch expr.Kind {
	case ast.CastArray:
		value := e.expr(expr.Expr)
		return Array(ZvalGetArray(e.ctx, value))
	case ast.CastBool:
		value := e.expr(expr.Expr)
		return Bool(ZvalIsTrue(e.ctx, value))
	case ast.CastDouble:
		value := e.expr(expr.Expr)
		return Double(ZvalGetDouble(e.ctx, value))
	case ast.CastInt:
		value := e.expr(expr.Expr)
		return Long(ZvalGetLong(e.ctx, value))
	case ast.CastObject:
		value := e.expr(expr.Expr)
		return types.ZvalObject(ZvalGetObject(e.ctx, value))
	case ast.CastString:
		value := e.expr(expr.Expr)
		return String(ZvalGetStrVal(e.ctx, value))
	case ast.CastUnset:
		// notice: deprecated in php >=7.2, trigger fatal error in php >= 8.0
		return types.Null
	}
	panic(perr.Todof("castExpr: %v", expr))
}

func (e *Executor) unaryExpr(expr *ast.UnaryExpr) types.Zval {
	var value = e.expr(expr.Var)
	switch expr.Op {
	case ast.UnaryOpPlus:
		return OpMul(e.ctx, value, Long(1))
	case ast.UnaryOpMinus:
		return OpMul(e.ctx, value, Long(-1))
	case ast.UnaryOpBooleanNot:
		return OpBooleanNot(e.ctx, value)
	case ast.UnaryOpBitwiseNot:
		return OpBitwiseNot(e.ctx, value)
	// todo 考虑是否需要用 unary 原生替代 v = v + 1 的模拟
	case ast.UnaryOpPreInc:
		newValue := OpAdd(e.ctx, value, Long(1))
		e.variableSet(expr.Var, newValue)
		return newValue
	case ast.UnaryOpPreDec:
		newValue := OpSub(e.ctx, value, Long(1))
		e.variableSet(expr.Var, newValue)
		return newValue
	case ast.UnaryOpPostInc:
		newValue := OpAdd(e.ctx, value, Long(1))
		e.variableSet(expr.Var, newValue)
		return value
	case ast.UnaryOpPostDec:
		newValue := OpSub(e.ctx, value, Long(1))
		e.variableSet(expr.Var, newValue)
		return value
	default:
		panic(perr.Internalf("Unexpected ast.UnaryExpr.Op: %+v", expr.Op))
	}
}

func (e *Executor) binaryOpExpr(expr *ast.BinaryOpExpr) (val types.Zval) {
	// && / || / ?? 操作比较特殊，右表达式节点可能不会执行
	switch expr.Op {
	case ast.BinaryOpBooleanAnd: // &&
		left := e.expr(expr.Left)
		right := func() types.Zval { return e.expr(expr.Right) }
		return OpBooleanAnd(e.ctx, left, right)
	case ast.BinaryOpBooleanOr: // ||
		left := e.expr(expr.Left)
		right := func() types.Zval { return e.expr(expr.Right) }
		return OpBooleanAnd(e.ctx, left, right)
	case ast.BinaryOpCoalesce: // ??
		left := e.expr(expr.Left)
		right := func() types.Zval { return e.expr(expr.Right) }
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
func (e *Executor) assignExpr(expr *ast.AssignExpr) types.Zval {
	value := e.expr(expr.Expr)
	e.variableSet(expr.Var, value)
	return value
}

func (e *Executor) assignOpExpr(expr *ast.AssignOpExpr) types.Zval {
	left := e.expr(expr.Var)

	var value types.Zval
	switch expr.Op {
	case ast.AssignOpBitwiseAnd: // &=
		value = OpBitwiseAnd(e.ctx, left, e.expr(expr.Expr))
	case ast.AssignOpBitwiseOr: // |=
		value = OpBitwiseOr(e.ctx, left, e.expr(expr.Expr))
	case ast.AssignOpBitwiseXor: // ^=
		value = OpBitwiseXor(e.ctx, left, e.expr(expr.Expr))
	case ast.AssignOpCoalesce: // ??=
		value = OpCoalesce(e.ctx, left, func() Val { return e.expr(expr.Expr) })
	case ast.AssignOpConcat: // .=
		value = OpConcat(e.ctx, left, e.expr(expr.Expr))
	case ast.AssignOpDiv: // /=
		value = OpDiv(e.ctx, left, e.expr(expr.Expr))
	case ast.AssignOpMinus: // -=
		value = OpSub(e.ctx, left, e.expr(expr.Expr))
	case ast.AssignOpMod: // %=
		value = OpMod(e.ctx, left, e.expr(expr.Expr))
	case ast.AssignOpMul: // *=
		value = OpMul(e.ctx, left, e.expr(expr.Expr))
	case ast.AssignOpPlus: // +=
		value = OpAdd(e.ctx, left, e.expr(expr.Expr))
	case ast.AssignOpPow: // **=
		value = OpPow(e.ctx, left, e.expr(expr.Expr))
	case ast.AssignOpShiftLeft: // <<=
		value = OpSL(e.ctx, left, e.expr(expr.Expr))
	case ast.AssignOpShiftRight: // >>=
		value = OpSR(e.ctx, left, e.expr(expr.Expr))
	default:
		panic(perr.Unreachable())
	}

	e.variableSet(expr.Var, value)
	return value
}

func (e *Executor) assignRefExpr(expr *ast.AssignRefExpr) types.Zval {
	variable := e.variableRef(expr.Var)
	valueVar := e.variableRef(expr.Expr)
	if valueVar == nil {
		return UninitializedZval()
	}
	value := types.ZvalRef(valueVar.MakeRef())
	variable.Set(value)
	return value
}

func (e *Executor) issetExpr(expr *ast.IssetExpr) types.Zval {
	for _, var_ := range expr.Vars {
		value := e.variableRef(var_).Get()
		if !ZvalIsTrue(e.ctx, value) {
			return types.False
		}
	}
	return types.True
}

func (e *Executor) emptyExpr(expr *ast.EmptyExpr) types.Zval {
	value := e.variableRef(expr.Expr).Get()
	if ZvalIsTrue(e.ctx, value) {
		return types.False
	}
	return types.True
}

func (e *Executor) evalExpr(expr *ast.EvalExpr) types.Zval {
	panic(perr.Todof("e.evalExpr"))
}

func (e *Executor) includeExpr(expr *ast.IncludeExpr) types.Zval {
	panic(perr.Todof("e.includeExpr"))
}

func (e *Executor) cloneExpr(expr *ast.CloneExpr) types.Zval {
	panic(perr.Todof("e.cloneExpr"))
}

func (e *Executor) errorSuppressExpr(expr *ast.ErrorSuppressExpr) types.Zval {
	var result types.Zval
	e.ctx.EG().ErrorSuppressScope(func() {
		result = e.expr(expr.Expr)
	})
	return result
}

func (e *Executor) exitExpr(expr *ast.ExitExpr) types.Zval {
	panic(perr.Todof("e.exitExpr"))
}

func (e *Executor) constFetchExpr(expr *ast.ConstFetchExpr) types.Zval {
	// todo 在命名空间内，使用非限定名时的常量获取逻辑
	name := expr.Name.ToCodeString()
	c := GetConstant(e.ctx, name)
	if c == nil {
		panic(perr.Todof("const not defined: " + name))
	}
	return c.Value()
}

func (e *Executor) classConstFetchExpr(expr *ast.ClassConstFetchExpr) types.Zval {
	className := e.nameOrExprAsString(expr.Class)
	constName := expr.Name.Name

	ce := ZendFetchClassByName(e.ctx, className, "", 0)
	if ce == nil {
		return UninitializedZval()
	}

	cc := ce.ConstantTable().Get(ascii.StrToLower(constName))
	if cc == nil {
		ThrowError(e.ctx, nil, fmt.Sprintf("Undefined class constant '%s'", constName))
		return UninitializedZval()
	}

	return cc.Value()
}

func (e *Executor) magicConstExpr(expr *ast.MagicConstExpr) types.Zval {
	panic(perr.Todof("e.magicConstExpr"))
}

func (e *Executor) instanceofExpr(expr *ast.InstanceofExpr) types.Zval {
	//val := e.expr(expr)
	//var className string
	//switch c := expr.Class.(type) {
	//case *ast.Name:
	//	className = c.ToCodeString()
	//case ast.Expr:
	//	className = ZvalGetStrVal(e.ctx, e.expr(c))
	//default:
	//	panic(perr.Internalf("预期外的 InstanceofExpr.Class 类型: %+v", c))
	//}
	panic(perr.Todof("e.instanceofExpr"))
}

func (e *Executor) listExpr(expr *ast.ListExpr) types.Zval {
	panic(perr.Todof("e.listExpr"))
}

func (e *Executor) printExpr(expr *ast.PrintExpr) types.Zval {
	PrintZval(e.ctx, e.expr(expr.Expr))
	return Long(1)
}

func (e *Executor) propertyFetchExpr(expr *ast.PropertyFetchExpr) types.Zval {
	obj := e.exprDeref(expr.Var)
	// todo 次数 propName 应为 Zval 而非 string, 有特殊的非 string 属性存在
	propName := e.identOrExprAsString(expr.Name)
	if !obj.IsObject() {
		Error(e.ctx, perr.E_NOTICE, fmt.Sprintf("Trying to get property '%s' of non-object", propName))
		return UninitializedZval()
	}

	ret := obj.Object().ReadProperty(types.ZvalString(propName), 0).DeRef()
	if ret.IsUndef() {
		Error(e.ctx, perr.E_NOTICE, fmt.Sprintf("Undefined property: %s::$%s", obj.Object().ClassName(), propName))
		return UninitializedZval()
	}

	return ret
}

func (e *Executor) staticPropertyFetchExpr(expr *ast.StaticPropertyFetchExpr) types.Zval {
	panic(perr.Todof("e.staticPropertyFetchExpr"))
}

func (e *Executor) shellExecExpr(expr *ast.ShellExecExpr) types.Zval {
	panic(perr.Todof("e.shellExecExpr"))
}

func (e *Executor) ternaryExpr(expr *ast.TernaryExpr) types.Zval {
	cond := e.expr(expr.Cond)
	if ZvalIsTrue(e.ctx, cond) {
		if expr.If == nil {
			return cond
		} else {
			return e.expr(expr.If)
		}
	} else {
		return e.expr(expr.Else)
	}
}

func (e *Executor) throwExpr(expr *ast.ThrowExpr) types.Zval {
	panic(perr.Todof("e.throwExpr"))
}

func (e *Executor) variableExpr(expr *ast.VariableExpr) types.Zval {
	name := e.variableName(expr.Name)
	// todo undefined warning
	value := e.currSymbols().Get(name)
	if value.IsUndef() {
		Error(e.ctx, perr.E_NOTICE, fmt.Sprintf("Undefined variable: %s", name))
		value = types.Null
	}
	return value
}

func (e *Executor) variableName(nameNode ast.Node) string {
	switch x := nameNode.(type) {
	case *ast.Ident:
		return x.Name
	case *ast.VariableExpr:
		return ZvalGetStrVal(e.ctx, e.expr(x))
	default:
		panic(perr.Todof("unexpected VariableExpr.Name type: %T, %+v", nameNode, nameNode))
	}
}

func (e *Executor) identOrExprAsString(node ast.Node) string {
	switch x := node.(type) {
	case *ast.Ident:
		return x.Name
	case ast.Expr:
		return ZvalGetStrVal(e.ctx, e.expr(x))
	default:
		panic(perr.Todof("expected Ident or Expr, %T given", node))
	}
}

func (e *Executor) nameOrExprAsString(node ast.Node) string {
	switch x := node.(type) {
	case *ast.Name:
		return x.ToCodeString()
	case ast.Expr:
		return ZvalGetStrVal(e.ctx, e.expr(x))
	default:
		panic(perr.Todof("expected Name or Expr, %T given", node))
	}
}

func (e *Executor) yieldExpr(expr *ast.YieldExpr) types.Zval {
	panic(perr.Todof("e.yieldExpr"))
}

func (e *Executor) yieldFromExpr(expr *ast.YieldFromExpr) types.Zval {
	panic(perr.Todof("e.yieldFromExpr"))
}

func (e *Executor) funcCallExpr(expr *ast.FuncCallExpr) types.Zval {
	var name string
	switch nameNode := expr.Name.(type) {
	case *ast.Name:
		name = nameNode.ToString()
	case ast.Expr:
		name = ZvalGetStrVal(e.ctx, e.expr(nameNode))
	default:
		panic(perr.Unreachable())
	}

	fn := e.initStringCall(name)
	if fn == nil {
		panic(perr.Internalf("函数未找到: %s", name))
	}
	return e.doCall(fn, expr.Args, nil)
}

func (e *Executor) newExpr(expr *ast.NewExpr) types.Zval {
	var ce *types.Class
	switch x := expr.Class.(type) {
	case *ast.Name:
		ce = ZendFetchClassByName(e.ctx, x.ToCodeString(), "", 0)
	default:
		panic(perr.Todof("e.newExpr 为完成的"))
	}
	if ce == nil {
		assert.Assert(e.ctx.EG().HasException())
		return UninitializedZval()
	}

	obj := ObjectInit(e.ctx, ce)
	if obj == nil {
		return UninitializedZval()
	}

	// todo constructor
	return Object(obj)
}

func (e *Executor) methodCallExpr(expr *ast.MethodCallExpr) types.Zval {
	obj := e.exprDeref(expr.Var)
	methodName := e.identOrExprAsString(expr.Name)
	if !obj.IsObject() {
		ThrowError(e.ctx, nil, fmt.Sprintf("Call to a member function %s() on %s", methodName, types.ZendZvalTypeName(obj)))
		return UninitializedZval()
	}

	method := obj.Object().GetMethod(methodName)
	if method == nil {
		ThrowError(e.ctx, nil, fmt.Sprintf("Call to undefined method %s::%s()", obj.Object().Class().Name(), methodName))
		return UninitializedZval()
	}

	return e.doCall(method, expr.Args, nil)
}

func (e *Executor) staticCallExpr(expr *ast.StaticCallExpr) types.Zval {
	panic(perr.Todof("e.staticCallExpr"))
}

// ---

func (e *Executor) variableSet(variable ast.Expr, value types.Zval) {
	e.variableRef(variable).Set(value)
}

func (e *Executor) variableRefArray(variable ast.Expr) types.Zval {
	ref := e.variableRef(variable)
	value := ref.Get()
	if value.IsUndef() {
		value = types.InitZvalArray()
		ref.Set(value)
	}
	return value
}

func (e *Executor) variableRef(expr ast.Expr) iVariable {
	switch v := expr.(type) {
	case *ast.VariableExpr:
		name := e.variableName(v.Name)
		symbols := e.currSymbols()
		return newSymbolVariable(symbols, name)
	case *ast.IndexExpr:
		arr := e.variableRefArray(v.Var)
		if v.Dim == nil {
			return newArrayAppendVariable(e.ctx, arr)
		} else {
			dim := e.expr(v.Dim)
			key, ok := ZvalToArrayKey(e.ctx, dim)
			if !ok {
				return nil
			}
			return newArrayDimVariable(e.ctx, arr, key)
		}
	default:
		panic(perr.Todof("unsupported variableRef.Var type: %T, %+v", v, v))
	}
}
func (e *Executor) variableRefZval(expr ast.Expr) types.Zval {
	return types.ZvalRef(e.variableRef(expr).MakeRef())
}

func (e *Executor) arrayGet(arr types.Zval, key types.ArrayKey) types.Zval {
	switch arr.Type() {
	case types.IsArray:
		return arr.Array().Find(key)
	case types.IsString:
		var offset int
		if key.IsStrKey() {
			Error(e.ctx, perr.E_WARNING, fmt.Sprintf("Illegal string offset '%s'", key.StrKey()))
			offset = 0
		} else {
			offset = key.IdxKey()
		}

		str := arr.String()
		if 0 <= offset && offset < len(str) {
			return types.ZvalString(str[offset : offset+1])
		} else if offset < 0 && len(str)+offset >= 0 {
			offset += len(str)
			return types.ZvalString(str[offset : offset+1])
		} else {
			Error(e.ctx, perr.E_NOTICE, fmt.Sprintf("Uninitialized string offset: %d", offset))
		}
		return types.ZvalString("")
	default:
		panic(perr.Todof("unsupported e.arrayGet arr type: %s", types.ZvalGetType(arr)))
	}
}

func (e *Executor) arrayAppend(arr types.Zval, value types.Zval) {
	switch arr.Type() {
	case types.IsArray:
		arr.Array().Append(value)
	// todo ArrayAccess
	default:
		panic(perr.Todof("unsupported e.arrayAppend arr type: %s", types.ZvalGetType(arr)))
	}
}
func (e *Executor) arrayUpdate(arr types.Zval, key types.ArrayKey, value types.Zval) {
	switch arr.Type() {
	case types.IsArray:
		arr.Array().Update(key, value)
	// todo ArrayAccess
	default:
		panic(perr.Todof("unsupported e.arrayUpdate arr type: %s", types.ZvalGetType(arr)))
	}
}

func (e *Executor) doCall(fn *types.Function, argNodes []*ast.Arg, scope any) types.Zval {
	assert.Assert(fn != nil)

	// init args
	args := e.initCallArgs(fn, argNodes)

	// push && pop executeData
	ex := NewExecuteData(e.ctx, fn, args)
	ex.SetScope(scope)
	e.ctx.EG().PushExecuteData(ex)
	defer func() {
		popEx := e.ctx.EG().PopExecuteData()
		assert.AssertEx(ex == popEx, "push & pop executeData 的 executeData 对象不相同，执行堆栈可能有错误")
	}()

	// do call
	var retval types.Zval
	if fn.IsInternalFunction() {
		if handler, ok := fn.Handler().(ZifHandler); ok {
			handler(ex, &retval)
		} else {
			perr.Panic(fmt.Sprintf("不支持的内部函数 handler 类型: %T", fn.Handler()))
		}
	} else {
		retval = e.userFunction(fn, args)
	}
	if retval.IsUndef() {
		retval = UninitializedZval()
	}
	return retval
}

func (e *Executor) initCallArgs(fn *types.Function, argNodes []*ast.Arg) []types.Zval {
	if len(argNodes) == 0 {
		return nil
	}

	// todo 参数长度检查？

	getArgInfo := func(i int) *types.ArgInfo {
		// todo 变长参数下的类型获取
		if i < len(fn.ArgInfos()) {
			argInfo := fn.ArgInfos()[i]
			return &argInfo
		}
		return nil
	}

	args := make([]types.Zval, 0, len(argNodes))
	for _, arg := range argNodes {
		if !arg.Unpack {
			argInfo := getArgInfo(len(args))
			if argInfo != nil && argInfo.ByRef {
				args = append(args, e.variableRefZval(arg.Value))
			} else {
				argVal := e.expr(arg.Value).DeRef()
				args = append(args, argVal)
			}
		} else {
			argVal := e.expr(arg.Value)
			// todo Traversable
			if !argVal.IsArray() {
				panic(perr.Internalf("Only arrays and Traversables can be unpacked"))
			}
			argVal.Array().Each(func(key types.ArrayKey, value types.Zval) {
				if key.IsStrKey() {
					panic(perr.Internalf("Cannot unpack array with string keys"))
				}

				argInfo := getArgInfo(len(args))
				if argInfo != nil && argInfo.ByRef {
					args = append(args, types.ZvalRef(newArrayDimVariable(e.ctx, argVal, key).MakeRef()))
				} else {
					args = append(args, value.DeRef())
				}
			})
		}
	}
	return args
}
