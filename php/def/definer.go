package def

import "github.com/heyuuu/gophp/compile/ast"

type TopDefiner interface {
	Definer
	Use(useType ast.UseType, name string, alias string)
}

type Definer interface {
	IsTrue(v Val) bool
	AsStr(v Val, stage Stage) string
	/* mics */
	Arg(v Val, unpack bool) Arg
	Param(t TypeHint, name string, defaultVal Val, byRef bool, variadic bool)
	/* type */
	SimpleType(name string) TypeHint
	IntersectionType(types ...TypeHint) TypeHint
	UnionType(types ...TypeHint) TypeHint
	NullableType(typ TypeHint) TypeHint
	/* lit */
	Undef() Val
	Null() Val
	Bool(v bool) Val
	Int(v int) Val
	Float(v float64) Val
	String(s string) Val
	Array(v ...*ArrayItem) Val
	ArrayItem(key Val, val Val, byRef bool, unpack bool) *ArrayItem
	/* variable */
	Var(name string) Variable
	VarEx(name Val) Variable
	VarValue(name string) Val
	VarValueEx(name Val) Variable
	/* index */
	Index(arr Val, key Val) Val
	/* unary op */
	PrePlus(v Val) Val
	PreMinus(v Val) Val
	BooleanNot(v Val) Val
	BitwiseNot(v Val) Val
	PreInc(v Variable)
	PreDec(v Variable)
	PostInc(v Variable)
	PostDec(v Variable)
	/* binary op */
	BitwiseAnd(v1 Val, v2 Val) Val     // &
	BitwiseOr(v1 Val, v2 Val) Val      // |
	BitwiseXor(v1 Val, v2 Val) Val     // ^
	BooleanAnd(v1 Val, v2 Val) Val     // &&
	BooleanOr(v1 Val, v2 Val) Val      // ||
	BooleanXor(v1 Val, v2 Val) Val     // xor
	Coalesce(v1 Val, v2 Val) Val       // ??
	Concat(v1 Val, v2 Val) Val         // .
	Div(v1 Val, v2 Val) Val            // /
	Equal(v1 Val, v2 Val) Val          // ==
	Greater(v1 Val, v2 Val) Val        // >
	GreaterOrEqual(v1 Val, v2 Val) Val // >=
	Identical(v1 Val, v2 Val) Val      // ===
	Minus(v1 Val, v2 Val) Val          // -
	Mod(v1 Val, v2 Val) Val            // %
	Mul(v1 Val, v2 Val) Val            // *
	NotEqual(v1 Val, v2 Val) Val       // !=
	NotIdentical(v1 Val, v2 Val) Val   // !==
	Plus(v1 Val, v2 Val) Val           // +
	Pow(v1 Val, v2 Val) Val            // **
	ShiftLeft(v1 Val, v2 Val) Val      // <<
	ShiftRight(v1 Val, v2 Val) Val     // >>
	Smaller(v1 Val, v2 Val) Val        // <
	SmallerOrEqual(v1 Val, v2 Val) Val // <=
	Spaceship(v1 Val, v2 Val) Val      // <=>
	/* assign op */
	Assign(v Variable, value Val) Val
	AssignRef(v Variable, value Variable) Val
	Unset(v Variable)
	AssignBitwiseAnd(v Variable, value Val) Val // &=
	AssignBitwiseOr(v Variable, value Val) Val  // |=
	AssignBitwiseXor(v Variable, value Val) Val // ^=
	AssignCoalesce(v Variable, value Val) Val   // ??=
	AssignConcat(v Variable, value Val) Val     // .=
	AssignDiv(v Variable, value Val) Val        // /=
	AssignMinus(v Variable, value Val) Val      // -=
	AssignMod(v Variable, value Val) Val        // %=
	AssignMul(v Variable, value Val) Val        // *=
	AssignPlus(v Variable, value Val) Val       // +=
	AssignPow(v Variable, value Val) Val        // **=
	AssignShiftLeft(v Variable, value Val) Val  // <<=
	AssignShiftRight(v Variable, value Val) Val // >>=

	/* system call */
	Echo(v ...string)
	EchoVal(v ...Val)
	Cast(kind ast.CastKind, v Val) Val
	Isset(vars ...Variable) Val
	Empty(v Val) Val
	Eval(v Val) Val
	Include(kind ast.IncludeKind, v Val) Val
	Clone(v Val) Val
	ErrorSuppress(v Val) Val
	Exit(v Val)

	/* const */
	Const(name string) Val
	ClassConst(className string, name string) Val
	MagicConst(kind ast.MagicConstKind, value Val) Val

	/* decl */
	DeclConst(name string, v Val)
	DeclGlobal(name string)
	DeclStatic(name string, v Val)

	/* foreach */
	ForeachIterator(v Val) ForeachIterator

	/* call like */
	FuncCall(fnName string, args ...Val) Val
	New(className string, args ...Val) Val
	MethodCall(inst Val, methodName string, args ...Val)
	NullsafeMethodCall(inst Val, methodName string, args ...Val)
	StaticMethodCall(className string, methodName string, args ...Val)
}
