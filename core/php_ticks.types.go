// <<generate>>

package core

/**
 * StTickFunction
 */
type StTickFunction struct {
	func_ func(int, any)
	arg   any
}

func (this *StTickFunction) GetFunc() func(int, any)      { return this.func_ }
func (this *StTickFunction) SetFunc(value func(int, any)) { this.func_ = value }
func (this *StTickFunction) GetArg() any                  { return this.arg }
func (this *StTickFunction) SetArg(value any)             { this.arg = value }
