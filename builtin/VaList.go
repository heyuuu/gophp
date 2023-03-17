package builtin

import "log"

type VaList struct {
	args  []any
	index int
}

func NewVaList(args []any) *VaList {
	return &VaList{args, 0}
}

func (va *VaList) Reset() {
	va.index = 0
}

func (va *VaList) Pop() any {
	arg := va.args[va.index]
	va.index++
	return arg
}

func (va *VaList) Next() {
	va.index++
}

func VaArg[T any](va *VaList) T {
	arg, ok := va.Pop().(T)
	if !ok {
		log.Panic("VaArg 类型转化异常")
	}
	return arg
}
