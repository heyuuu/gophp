package args

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
	if r.pos >= len(r.args) {
		log.Fatal("解析参数异常，超过获取长度")
	}

	if ptr, ok := r.args[r.pos].(*T); ok {
		*ptr = val
	} else {
		log.Fatalf("解析参数异常: 类型不匹配，pos=%d", r.pos)
	}

	r.pos++
}

func (r *VaArgsReceiver) Bool(val bool)              { PutVaArg[types.ZendBool](r, types.IntBool(val)) }
func (r *VaArgsReceiver) ZendBool(val int)           { PutVaArg[types.ZendBool](r, val) }
func (r *VaArgsReceiver) Long(val int)               { PutVaArg[int](r, val) }
func (r *VaArgsReceiver) Double(val float64)         { PutVaArg[float64](r, val) }
func (r *VaArgsReceiver) Str(val string)             { PutVaArg[*types.ZendString](r, types.NewZendString(val)) }
func (r *VaArgsReceiver) ZStr(val *types.ZendString) { PutVaArg[*types.ZendString](r, val) }

func (r *VaArgsReceiver) StrPtr(s *byte, l int) {
	PutVaArg[*byte](r, s)
	PutVaArg[int](r, l)
}
