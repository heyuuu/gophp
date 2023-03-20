package argparse

import (
	"log"
	"sik/zend/types"
)

/**
 * TypeSpec 的 Reader，方便读取
 */
type typeSpecReader struct{ str string }

func (r *typeSpecReader) curr() byte {
	if r.str != "" {
		return r.str[0]
	}
	return 0
}

func (r *typeSpecReader) read() byte {
	if r.str != "" {
		c := r.str[0]
		r.str = r.str[1:]
		return c
	}
	return 0
}

func (r *typeSpecReader) Next() (typ byte, checkNull bool, separate bool) {
	typ = r.read()
	if typ == '|' {
		typ = r.read()
	}
	for {
		if r.curr() == '/' {
			separate = true
		} else if r.curr() == '!' {
			checkNull = true
		} else {
			break
		}
		r.read()
	}
	return
}

/**
 * vaList 用于便捷操作传入 varargs 变量的类
 */
type vaList struct {
	args []any
	pos  int
}

func newVaList(args []any) *vaList { return &vaList{args: args} }

func PutVaArg[T any](r *vaList, val T) {
	if ptr, ok := r.Pop().(*T); ok {
		*ptr = val
	} else {
		log.Fatalf("解析参数异常: 类型不匹配，pos=%d", r.pos)
	}
}

func (l *vaList) Pop() any {
	if l.pos >= len(l.args) {
		log.Fatal("解析参数异常，超过获取长度")
	}

	arg := l.args[l.pos]
	l.pos++
	return arg
}

func (l *vaList) Bool(val bool)              { PutVaArg(l, types.IntBool(val)) }
func (l *vaList) ZendBool(val int)           { PutVaArg(l, val) }
func (l *vaList) Long(val int)               { PutVaArg(l, val) }
func (l *vaList) Double(val float64)         { PutVaArg(l, val) }
func (l *vaList) Str(val string)             { PutVaArg(l, types.NewString(val)) }
func (l *vaList) ZStr(val *types.String)     { PutVaArg(l, val) }
func (l *vaList) Array(val *types.ZendArray) { PutVaArg(l, val) }
func (l *vaList) Zval(val *types.Zval)       { PutVaArg(l, val) }
func (l *vaList) StrPtr(str *byte, len int) {
	PutVaArg[*byte](l, str)
	PutVaArg[int](l, len)
}

func vaArg[T any](va *[]any) *T {
	if len(*va) == 0 {
		log.Fatal("解析参数异常，超过获取长度")
	}

	ptr, ok := (*va)[0].(*T)
	if !ok {
		log.Fatalf("解析参数异常: 类型不匹配")
	}

	*va = (*va)[1:]
	return ptr
}
