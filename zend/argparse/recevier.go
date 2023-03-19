package argparse

import (
	"log"
	"sik/zend/types"
)

type VaArgsReceiver struct {
	args []any
	pos  int
}

func NewVaListReceiver(args []any, pos int) *VaArgsReceiver {
	return &VaArgsReceiver{args: args, pos: pos}
}

func PutVaArg[T any](r *VaArgsReceiver, val T) {
	if ptr, ok := r.Pop().(*T); ok {
		*ptr = val
	} else {
		log.Fatalf("解析参数异常: 类型不匹配，pos=%d", r.pos)
	}
}

func (r *VaArgsReceiver) Pop() any {
	if r.pos >= len(r.args) {
		log.Fatal("解析参数异常，超过获取长度")
	}

	arg := r.args[r.pos]
	r.pos++
	return arg
}

func (r *VaArgsReceiver) Bool(val bool)              { PutVaArg(r, types.IntBool(val)) }
func (r *VaArgsReceiver) ZendBool(val int)           { PutVaArg(r, val) }
func (r *VaArgsReceiver) Long(val int)               { PutVaArg(r, val) }
func (r *VaArgsReceiver) Double(val float64)         { PutVaArg(r, val) }
func (r *VaArgsReceiver) Str(val string)             { PutVaArg(r, types.NewZendString(val)) }
func (r *VaArgsReceiver) ZStr(val *types.ZendString) { PutVaArg(r, val) }
func (r *VaArgsReceiver) Array(val *types.ZendArray) { PutVaArg(r, val) }
func (r *VaArgsReceiver) Zval(val *types.Zval)       { PutVaArg(r, val) }
func (r *VaArgsReceiver) StrPtr(s *byte, l int) {
	PutVaArg[*byte](r, s)
	PutVaArg[int](r, l)
}
