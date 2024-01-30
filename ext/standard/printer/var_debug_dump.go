package printer

import (
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"io"
)

type VarDebugPrinter struct {
	basePrinter
	ctx *php.Context
}

func NewVarDebugPrinter(ctx *php.Context, w io.Writer) *VarDebugPrinter {
	p := &VarDebugPrinter{}
	p.ctx = ctx
	p.w = w
	return p
}

func (p *VarDebugPrinter) Zval(zv types.Zval, level int) {
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
		p.printf("\"\n")
	case types.IsArray:
		myht := zv.Array()
		if myht.IsRecursive() {
			p.print("*RECURSION*\n")
			return
		}
		myht.ProtectRecursive()
		count := myht.Count()
		p.printf("%sarray(%d){\n", common, count)
		myht.Each(func(key types.ArrayKey, value types.Zval) {
			p.arrayElement(value, key, level)
		})
		myht.UnprotectRecursive()
		p.printIdent(level)
		p.print("}\n")
	case types.IsObject:
		var myht *types.Array
		myht = php.ZendGetPropertiesFor(zv, types.PropPurposeDebug)
		className := zv.Object().ClassName()

		if myht == nil {
			p.printf("%sobject(%s)#%d (%d) {\n", common, className, zv.Object().Handle(), 0) // types.ZendStringReleaseEx(class_name, 0)
		} else {
			if myht.IsRecursive() {
				p.print("*RECURSION*\n")
				//zend.ZendReleaseProperties(myht)
				return
			}

			myht.ProtectRecursive()
			p.printf("%sobject(%s)#%d (%d) {\n", common, className, zv.Object().Handle(), myht.Count())
			myht.Each(func(key types.ArrayKey, value types.Zval) {
				var propInfo *types.PropertyInfo = nil
				//if value.IsIndirect() {
				//	value = value.Indirect()
				//	if key.IsStrKey() {
				//		propInfo = php.ZendGetTypedPropertyInfoForSlot(zv.Object(), value)
				//	}
				//}
				if !value.IsUndef() || propInfo != nil {
					p.objectProperty(propInfo, value, key, level)
				}
			})
			myht.UnprotectRecursive()
		}

		p.printIdent(level)
		p.print("}\n")
	case types.IsResource:
		typeName := lang.Option(php.ZendRsrcListGetRsrcTypeEx(zv.Resource()), "Unknown")
		p.printf("%sresource(%d) of type (%s)\n", common, zv.ResourceHandle(), typeName)
	default:
		p.printf("%sUNKNOWN:0\n", common)
	}
}

func (p *VarDebugPrinter) arrayElement(zv types.Zval, key types.ArrayKey, level int) {
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

func (p *VarDebugPrinter) objectProperty(propInfo *types.PropertyInfo, zv types.Zval, key types.ArrayKey, level int) {
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
			p.printf(`"%s"`, propName)
		}
		p.print("]=>\n")
	}
	if propInfo != nil && zv.IsUndef() {
		php.Assert(propInfo.Type() != nil)
		typ := propInfo.Type().FormatType()
		p.printIdent(level)
		p.printf("uninitialized(%s)\n", typ)
	} else {
		p.Zval(zv, level+2)
	}
}
