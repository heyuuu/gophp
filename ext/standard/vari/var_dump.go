package vari

import (
	"fmt"
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
)

type VarDumpPrinter struct {
	ctx *php.Context
}

func NewVarDumpPrinter(ctx *php.Context) *VarDumpPrinter {
	return &VarDumpPrinter{ctx: ctx}
}

func (p *VarDumpPrinter) ArrayElement(zv types.Zval, key types.ArrayKey, level int) {
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

func (p *VarDumpPrinter) ObjectProperty(propInfo *types.PropertyInfo, zv types.Zval, key types.ArrayKey, level int) {
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
	//		ctx.WriteString(`"`)
	//		ctx.WriteString(key.StrKey())
	//		ctx.WriteString(`"`)
	//	}
	//	ctx.WriteString("]=>\n")
	//}
	//if zv.IsUndef() {
	//	//assert.Assert(propInfo.GetType() != nil)
	//	//typ := propInfo.GetType().FormatType()
	//	ctx.WriteString(fmt.Sprintf("%*cuninitialized(%s)\n", level+1, ' ', typ))
	//} else {
	//	p.Zval(zv, level+2)
	//}
}
func (p *VarDumpPrinter) Zval(struc types.Zval, level int) {
	ctx := p.ctx
	isRef := false
	if level > 1 {
		ctx.WriteString(fmt.Sprintf("%*c", level-1, ' '))
	}
again:
	common := ""
	if isRef {
		common = "&"
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
		ctx.WriteString(fmt.Sprintf("%sfloat(%.*G)\n", common, php.Precision, struc.Double()))
	case types.IsString:
		ctx.WriteString(fmt.Sprintf(`%sstring(%d) "`, common, len(struc.String())))
		ctx.WriteString(struc.String())
		ctx.WriteString("\"\n")
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
		ctx.WriteString(fmt.Sprintf("%sarray(%d) {\n", common, count))
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
		obj := struc.Object()
		if obj.IsRecursive() {
			ctx.WriteString("*RECURSION*\n")
			return
		}
		obj.ProtectRecursive()
		myht := obj.PropertiesFor(types.PropPurposeDebug)
		className := obj.ClassName()
		ctx.WriteString(fmt.Sprintf("%sobject(%s)#%d (%d) {\n", common, className, obj.Handle(), lang.CondF1(myht != nil, func() int { return myht.Count() }, 0)))
		if myht != nil {
			myht.Each(func(key types.ArrayKey, value types.Zval) {
				//		var prop_info *types.PropertyInfo = nil
				//		if value.IsIndirect() {
				//			value = value.Indirect()
				//			if key.IsStrKey() {
				//				prop_info = php.ZendGetTypedPropertyInfoForSlot(obj, value)
				//			}
				//		}
				//		if !value.IsUndef() || prop_info != nil {
				//			ObjectProperty(w, prop_info, value, key, level)
				//		}
			})
		}
		if level > 1 {
			ctx.WriteString(fmt.Sprintf("%*c", level-1, ' '))
		}
		ctx.WriteString("}\n")
		obj.UnprotectRecursive()
	case types.IsResource:
		//typeName := lang.Option(php.ZendRsrcListGetRsrcTypeEx(struc.Resource()), "Unknown")
		typeName := "Unknown"
		ctx.WriteString(fmt.Sprintf("%sresource(%d) of type (%s)\n", common, struc.ResourceHandle(), typeName))
	case types.IsRef:
		isRef = true
		struc = struc.DeRef()
		goto again
	default:
		ctx.WriteString(fmt.Sprintf("%sUNKNOWN:0\n", common))
	}
}
