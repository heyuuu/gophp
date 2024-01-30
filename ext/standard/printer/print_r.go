package printer

import (
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/types"
	"io"
	"strconv"
)

type PrintRPrinter struct {
	basePrinter
	ctx *php.Context
}

func NewPrintRPrinter(ctx *php.Context, w io.Writer) *PrintRPrinter {
	p := &PrintRPrinter{ctx: ctx}
	p.ctx = ctx
	p.w = w
	return p
}

func (p *PrintRPrinter) Zval(v types.Zval, indent int) {
	v = v.DeRef()
	switch v.Type() {
	case types.IsArray:
		p.print("Array\n")
		if v.Array().IsRecursive() {
			p.print(" *RECURSION*")
			return
		}
		v.Array().ProtectRecursive()
		p.printHash(v.Array(), indent, false)
		v.Array().UnprotectRecursive()
	case types.IsObject:
		var properties *types.Array
		p.print(v.Object().ClassName())
		p.print(" Object\n")
		if v.Object().IsRecursive() {
			p.print(" *RECURSION*")
			return
		}
		properties = php.ZendGetPropertiesFor(v, types.PropPurposeDebug)
		if properties == nil {
			break
		}
		v.Object().ProtectRecursive()
		p.printHash(properties, indent, true)
		v.Object().UnprotectRecursive()
	default:
		str := php.ZvalGetStrVal(p.ctx, v)
		p.print(str)
	}
}

func (p *PrintRPrinter) printHash(ht *types.Array, indent int, isObject bool) {
	p.printIdent(indent)
	p.print("(\n")
	indent += php.PrintZvalIndent
	ht.Each(func(key types.ArrayKey, value types.Zval) {
		p.printIdent(indent)
		p.print("[")
		if key.IsStrKey() {
			if isObject {
				className, propName, mangled := php.UnmanglePropertyName(p.ctx, key.StrKey())
				p.print(propName)
				if className != "" && mangled {
					if className[0] == '*' {
						p.print(":protected")
					} else {
						p.print(":")
						p.print(className)
						p.print(":private")
					}
				}
			} else {
				p.print(key.StrKey())
			}
		} else {
			p.print(strconv.Itoa(key.IdxKey()))
		}
		p.print("] => ")
		p.Zval(value, indent+php.PrintZvalIndent)
		p.print("\n")
	})
	indent -= php.PrintZvalIndent
	p.printIdent(indent)
	p.print(")\n")
}
