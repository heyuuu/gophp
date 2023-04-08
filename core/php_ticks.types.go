package core

/**
 * StTickFunction
 */
type StTickFunction struct {
	func_ func(int, any)
	arg   any
}

func MakeStTickFunction(func_ func(int, any), arg any) StTickFunction {
	return StTickFunction{
		func_: func_,
		arg:   arg,
	}
}
func (this *StTickFunction) GetFunc() func(int, any) { return this.func_ }
func (this *StTickFunction) GetArg() any             { return this.arg }
