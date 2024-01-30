package vari

import (
	"fmt"
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/types"
)

type VarDebugPrinter struct {
	ctx *php.Context
}

func NewVarDebugPrinter(ctx *php.Context) *VarDebugPrinter {
	return &VarDebugPrinter{ctx: ctx}
}

func (p *VarDebugPrinter) ArrayElement(zv types.Zval, key types.ArrayKey, level int) {
	ctx := p.ctx
	if key.IsStrKey() {
		ctx.WriteString(fmt.Sprintf(`%*c["`, level+1, ' '))
		ctx.WriteString(key.StrKey())
		ctx.WriteString("\"]=>\n")
	} else {
		ctx.WriteString(fmt.Sprintf("%*c[%d]=>\n", level+1, ' ', key.IdxKey()))
	}
	p.Zval(zv, level+2)
}
func (p *VarDebugPrinter) ObjectProperty(propInfo *types.PropertyInfo, zv *types.Zval, key types.ArrayKey, level int) {
	//ctx := p.ctx
	//if !key.IsStrKey() {
	//	ctx.WriteString(fmt.Sprintf("%*c[%d]=>\n", level+1, ' ', key.IdxKey()))
	//} else {
	//	className, propName, ok := php.UnmanglePropertyName(ctx, key.StrKey())
	//	ctx.WriteString(fmt.Sprintf("%*c[", level+1, ' '))
	//	if ok {
	//		if className[0] == '*' {
	//			ctx.WriteString(fmt.Sprintf(`"%s":protected`, propName))
	//		} else {
	//			ctx.WriteString(fmt.Sprintf(`"%s":"%s":private`, propName, className))
	//		}
	//	} else {
	//		ctx.WriteString(fmt.Sprintf(`"%s"`, propName))
	//	}
	//	ctx.WriteString("]=>\n")
	//}
	//if propInfo != nil && zv.IsUndef() {
	//	php.Assert(propInfo.Type() != nil)
	//
	//	typ := propInfo.Type().FormatType()
	//	ctx.WriteString(fmt.Sprintf("%*cuninitialized(%s)\n", level+1, ' ', typ))
	//} else {
	//	p.Zval(zv, level+2)
	//}
}
func (p *VarDebugPrinter) Zval(struc types.Zval, level int) {
	ctx := p.ctx
	if level > 1 {
		ctx.WriteString(fmt.Sprintf("%*c", level-1, ' '))
	}

	// deref
	common := ""
	if struc.IsRef() {
		common = "&"
		struc = struc.DeRef()
	}

	switch struc.Type() {
	case types.IsFalse:
		ctx.WriteString(fmt.Sprintf("%sbool(false)\n", common))
	case types.IsTrue:
		ctx.WriteString(fmt.Sprintf("%sbool(true)\n", common))
	case types.IsNull:
		ctx.WriteString(fmt.Sprintf("%sNULL\n", common))
	case types.IsLong:
		ctx.WriteString(fmt.Sprintf("%sint(%d)\n", common, struc.Long()))
	case types.IsDouble:
		ctx.WriteString(fmt.Sprintf("%sfloat(%.*G)\n", common, ctx.EG().Precision(), struc.Double()))
	case types.IsString:
		ctx.WriteString(fmt.Sprintf(`%sstring(%d) "`, common, len(struc.String())))
		ctx.WriteString(struc.String())
		ctx.WriteString(fmt.Sprintf("\"\n"))
	case types.IsArray:
		myht := struc.Array()
		if level > 1 {
			if myht.IsRecursive() {
				ctx.WriteString("*RECURSION*\n")
				return
			}
			myht.ProtectRecursive()
		}
		count := myht.Count()
		ctx.WriteString(fmt.Sprintf("%sarray(%d){\n", common, count))
		myht.Each(func(key types.ArrayKey, value types.Zval) {
			p.ArrayElement(value, key, level)
		})
		if level > 1 {
			myht.UnprotectRecursive()
		}
		if level > 1 {
			ctx.WriteString(fmt.Sprintf("%*c", level-1, ' '))
		}
		ctx.WriteString("}\n")
	case types.IsObject:
		//myht := php.ZendGetPropertiesFor(struc, types.PropPurposeDebug)
		//if myht != nil {
		//	if myht.IsRecursive() {
		//		ctx.WriteString("*RECURSION*\n")
		//		//zend.ZendReleaseProperties(myht)
		//		return
		//	}
		//	myht.ProtectRecursive()
		//}
		//className := struc.Object().ClassName()
		//ctx.WriteString(fmt.Sprintf("%sobject(%s)#%d (%d) {\n", common, className, struc.Object().Handle(), lang.CondF1(myht != nil, func() int { return myht.Count() }, 0))) // types.ZendStringReleaseEx(class_name, 0)
		//if myht != nil {
		//	myht.Each(func(key types.ArrayKey, value types.Zval) {
		//		var propInfo *types.PropertyInfo = nil
		//		if value.IsIndirect() {
		//			value = value.Indirect()
		//			if key.IsStrKey() {
		//				propInfo = php.ZendGetTypedPropertyInfoForSlot(struc.Object(), value)
		//			}
		//		}
		//		if !value.IsUndef() || propInfo != nil {
		//			p.ObjectProperty(propInfo, value, key, level)
		//		}
		//	})
		//	myht.UnprotectRecursive()
		//}
		//if level > 1 {
		//	ctx.WriteString(fmt.Sprintf("%*c", level-1, ' '))
		//}
		//ctx.WriteString("}\n")
	case types.IsResource:
		//typeName := lang.Option(php.ZendRsrcListGetRsrcTypeEx(struc.Resource()), "Unknown")
		//ctx.WriteString(fmt.Sprintf("%sresource(%d) of type (%s)\n", common, struc.ResourceHandle(), typeName))
	default:
		ctx.WriteString(fmt.Sprintf("%sUNKNOWN:0\n", common))
	}
}
