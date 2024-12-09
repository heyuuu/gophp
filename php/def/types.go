package def

type Val interface {
	val()
}

type Variable interface {
	Val
	Set(v Val)
}

type Arg interface {
	Value() Val
	Unpack() bool
}

type Param interface {
	ByRef() bool
	Variadic() bool
	Default() Val
}

type TypeHint interface {
	typeHint()
}

type ArrayItem struct {
	Key Val
	Val Val
}

type ForeachIterator interface {
	Valid() bool
	Next()
	Key() Val
	Value() Val
}
