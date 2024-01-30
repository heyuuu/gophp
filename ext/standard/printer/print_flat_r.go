package printer

import (
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/types"
	"io"
	"strconv"
)

type PrintFlatRPrinter struct {
	basePrinter
	ctx *php.Context
}

func NewPrintFlatRPrinter(ctx *php.Context, w io.Writer) *PrintFlatRPrinter {
	p := &PrintFlatRPrinter{ctx: ctx}
	p.ctx = ctx
	p.w = w
	return p
}

func (p *PrintFlatRPrinter) WriteString(str string) {
	p.ctx.WriteString(str)
}

func (p *PrintFlatRPrinter) Zval(v types.Zval) {
	v = v.DeRef()
	switch v.Type() {
	case types.IsArray:
		p.print("Array (")
		if v.Array().IsRecursive() {
			p.print(" *RECURSION*")
			return
		}
		v.Array().ProtectRecursive()
		p.printHash(v.Array())
		p.print(")")
		v.Array().UnprotectRecursive()
	case types.IsObject:
		var properties *types.Array
		p.printf("%s Object (", v.Object().ClassName())
		if v.Object().IsRecursive() {
			p.print(" *RECURSION*")
			return
		}
		properties = types.Z_OBJPROP_P(v)
		if properties != nil {
			v.Object().ProtectRecursive()
			p.printHash(properties)
			v.Object().UnprotectRecursive()
		}
		p.print(")")
	default:
		str := php.ZvalGetStrVal(p.ctx, v)
		p.print(str)
	}
}

func (p *PrintFlatRPrinter) printHash(ht *types.Array) {
	var i = 0
	ht.Each(func(key types.ArrayKey, value types.Zval) {
		if i > 0 {
			p.print(",")
		}
		i++

		p.print("[")
		if key.IsStrKey() {
			p.print(key.StrKey())
		} else {
			p.print(strconv.Itoa(key.IdxKey()))
		}
		p.print("] => ")
		p.Zval(value)
	})
}
