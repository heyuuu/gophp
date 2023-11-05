package operators

type LazyVal = func() Val

type iOperator interface {
	// binary op
	Add(Val, Val) Val
	Sub(Val, Val) Val
	Mul(Val, Val) Val
	Div(Val, Val) Val
	Mod(Val, Val) Val
	SL(Val, Val) Val
	SR(Val, Val) Val
	Concat(Val, Val) Val
	Pow(Val, Val) Val
	BitwiseAnd(Val, Val) Val
	BitwiseOr(Val, Val) Val
	BitwiseXor(Val, Val) Val
	BooleanAnd(Val, LazyVal) Val
	BooleanOr(Val, LazyVal) Val
	BooleanXor(Val, Val) Val
	Identical(Val, Val) Val
	NotIdentical(Val, Val) Val
	Equal(Val, Val) Val
	NotEqual(Val, Val) Val
	Greater(Val, Val) Val
	GreaterOrEqual(Val, Val) Val
	Smaller(Val, Val) Val
	SmallerOrEqual(Val, Val) Val
	Spaceship(Val, Val) Val
}
