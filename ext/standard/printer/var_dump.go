package printer

import (
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/assert"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"io"
)

type VarDumpPrinter struct {
	basePrinter
	ctx *php.Context
}

func NewVarDumpPrinter(ctx *php.Context, w io.Writer) *VarDumpPrinter {
	p := &VarDumpPrinter{}
	p.ctx = ctx
	p.w = w
	return p
}

func (p *VarDumpPrinter) Zval(zv types.Zval, level int) {
	p.printIdent(level)

	// deref
	common := ""
	if zv.IsRef() {
		common = "&"
		zv = zv.DeRef()
	}

	switch zv.Type() {
	case types.IsFalse:
		p.printf("%sbool(false)\n", common)
	case types.IsTrue:
		p.printf("%sbool(true)\n", common)
	case types.IsNull:
		p.printf("%sNULL\n", common)
	case types.IsLong:
		p.printf("%sint(%d)\n", common, zv.Long())
	case types.IsDouble:
		doubleStr := php.FormatDouble(zv.Double(), 'G', p.ctx.EG().Precision())
		p.printf("%sfloat(%s)\n", common, doubleStr)
	case types.IsString:
		p.printf(`%sstring(%d) "`, common, len(zv.String()))
		p.print(zv.String())
		p.print("\"\n")
	case types.IsArray:
		myht := zv.Array()
		if myht.IsRecursive() {
			p.print("*RECURSION*\n")
			return
		}
		myht.ProtectRecursive()
		count := myht.Count()
		p.printf("%sarray(%d) {\n", common, count)
		myht.Each(func(key types.ArrayKey, value types.Zval) {
			p.arrayElement(value, key, level)
		})
		myht.UnprotectRecursive()
		p.printIdent(level)
		p.print("}\n")
	case types.IsObject:
		obj := zv.Object()
		if obj.IsRecursive() {
			p.print("*RECURSION*\n")
			return
		}
		obj.ProtectRecursive()
		myht := obj.PropertiesFor(types.PropPurposeDebug)
		className := obj.ClassName()
		p.printf("%sobject(%s)#%d (%d) {\n", common, className, obj.Handle(), lang.CondF1(myht != nil, func() int { return myht.Count() }, 0))
		if myht != nil {
			myht.Each(func(key types.ArrayKey, value types.Zval) {
				var propInfo *types.PropertyInfo = nil
				//if value.IsIndirect() {
				//	value = value.Indirect()
				//	if key.IsStrKey() {
				//		propInfo = php.ZendGetTypedPropertyInfoForSlot(obj, value)
				//	}
				//}
				if !value.IsUndef() || propInfo != nil {
					p.objectProperty(propInfo, value, key, level)
				}
			})
		}
		p.printIdent(level)
		p.print("}\n")
		obj.UnprotectRecursive()
	case types.IsResource:
		typeName := lang.Option(php.ZendRsrcListGetRsrcTypeEx(zv.Resource()), "Unknown")
		p.printf("%sresource(%d) of type (%s)\n", common, zv.ResourceHandle(), typeName)
	default:
		p.printf("%sUNKNOWN:0\n", common)
	}
}

func (p *VarDumpPrinter) arrayElement(zv types.Zval, key types.ArrayKey, level int) {
	if key.IsStrKey() {
		p.printIdent(level + 2)
		p.print(`["`)
		p.print(key.StrKey())
		p.println(`"]=>`)
	} else {
		p.printIdent(level + 2)
		p.printf("[%d]=>\n", key.IdxKey())
	}
	p.Zval(zv, level+2)
}
func (p *VarDumpPrinter) objectProperty(propInfo *types.PropertyInfo, zv types.Zval, key types.ArrayKey, level int) {
	if !key.IsStrKey() {
		p.printIdent(level)
		p.printf("[%d]=>\n", key.IdxKey())
	} else {
		className, propName, ok := php.UnmanglePropertyName(p.ctx, key.StrKey())

		p.printIdent(level)
		p.print("[")
		if ok {
			if className[0] == '*' {
				p.printf(`"%s":protected`, propName)
			} else {
				p.printf(`"%s":"%s":private`, propName, className)
			}
		} else {
			p.print(`"`)
			p.print(key.StrKey())
			p.print(`"`)
		}
		p.print("]=>\n")
	}
	if zv.IsUndef() {
		assert.Assert(propInfo.Type() != nil)
		typ := propInfo.Type().FormatType()
		p.printIdent(level)
		p.printf("uninitialized(%s)\n", typ)
	} else {
		p.Zval(zv, level+2)
	}
}
