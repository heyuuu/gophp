package standard

import (
	"fmt"
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/php/zpp"
)

type stringWriter interface {
	WriteString(s string)
}

func PhpArrayElementDump(w stringWriter, zv types.Zval, key types.ArrayKey, level int) {
	if key.IsStrKey() {
		w.WriteString(fmt.Sprintf(`%*c["`, level+1, ' '))
		w.WriteString(key.StrKey())
		w.WriteString("\"]=>\n")
	} else {
		w.WriteString(fmt.Sprintf("%*c[%d]=>\n", level+1, ' ', key.IdxKey()))
	}
	PhpVarDump(w, zv, level+2)
}

//func PhpObjectPropertyDump(w stringWriter, propInfo *types.PropertyInfo, zv *types.Zval, key_ types.ArrayKey, level int) {
//	if !key_.IsStrKey() {
//		w.WriteString(fmt.Sprintf("%*c[%d]=>\n", level+1, ' ', key_.IdxKey()))
//	} else {
//		className, propName, ok := php.ZendUnmanglePropertyName_Ex(key_.StrKey())
//		w.WriteString(fmt.Sprintf("%*c[", level+1, ' '))
//		if ok {
//			if className[0] == '*' {
//				w.WriteString(fmt.Sprintf(`"%s":protected`, propName))
//			} else {
//				w.WriteString(fmt.Sprintf(`"%s":"%s":private`, propName, className))
//			}
//		} else {
//			w.WriteString(`"`)
//			w.WriteString(key_.StrKey())
//			w.WriteString(`"`)
//		}
//		w.WriteString("]=>\n")
//	}
//	if zv.IsUndef() {
//		perr.Assert(propInfo.GetType() != nil)
//
//		typ := propInfo.GetType().FormatType()
//		w.WriteString(fmt.Sprintf("%*cuninitialized(%s)\n", level+1, ' ', typ))
//	} else {
//		PhpVarDump(w, zv, level+2)
//	}
//}
func PhpVarDump(w stringWriter, struc types.Zval, level int) {
	isRef := false
	if level > 1 {
		w.WriteString(fmt.Sprintf("%*c", level-1, ' '))
	}
again:
	common := ""
	if isRef {
		common = "&"
	}
	switch struc.Type() {
	case types.IsFalse:
		w.WriteString(fmt.Sprintf("%sbool(false)\n", common))
	case types.IsTrue:
		w.WriteString(fmt.Sprintf("%sbool(true)\n", common))
	case types.IsNull:
		w.WriteString(fmt.Sprintf("%sNULL\n", common))
	case types.IsLong:
		w.WriteString(fmt.Sprintf("%sint(%d)\n", common, struc.Long()))
	case types.IsDouble:
		w.WriteString(fmt.Sprintf("%sfloat(%.*G)\n", common, php.Precision, struc.Double()))
	case types.IsString:
		w.WriteString(fmt.Sprintf(`%sstring(%d) "`, common, len(struc.String())))
		w.WriteString(struc.String())
		w.WriteString("\"\n")
	case types.IsArray:
		myht := struc.Array()
		if level > 1 {
			if myht.IsRecursive() {
				w.WriteString("*RECURSION*\n")
				return
			}
			myht.ProtectRecursive()
		}
		count := myht.Count()
		w.WriteString(fmt.Sprintf("%sarray(%d) {\n", common, count))
		myht.Each(func(key types.ArrayKey, value types.Zval) {
			PhpArrayElementDump(w, value, key, level)
		})
		if level > 1 {
			myht.UnprotectRecursive()
		}
		if level > 1 {
			w.WriteString(fmt.Sprintf("%*c", level-1, ' '))
		}
		w.WriteString("}\n")
	case types.IsObject:
		obj := struc.Object()
		if obj.IsRecursive() {
			w.WriteString("*RECURSION*\n")
			return
		}
		obj.ProtectRecursive()
		myht := obj.PropertiesFor(types.PropPurposeDebug)
		className := obj.ClassName()
		w.WriteString(fmt.Sprintf("%sobject(%s)#%d (%d) {\n", common, className, obj.Handle(), lang.CondF1(myht != nil, func() int { return myht.Count() }, 0)))
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
				//			PhpObjectPropertyDump(w, prop_info, value, key, level)
				//		}
			})
		}
		if level > 1 {
			w.WriteString(fmt.Sprintf("%*c", level-1, ' '))
		}
		w.WriteString("}\n")
		obj.UnprotectRecursive()
	case types.IsResource:
		//typeName := lang.Option(php.ZendRsrcListGetRsrcTypeEx(struc.Resource()), "Unknown")
		typeName := "Unknown"
		w.WriteString(fmt.Sprintf("%sresource(%d) of type (%s)\n", common, struc.ResourceHandle(), typeName))
	case types.IsRef:
		isRef = true
		struc = struc.DeRef()
		goto again
	default:
		w.WriteString(fmt.Sprintf("%sUNKNOWN:0\n", common))
	}
}
func ZifVarDump(ctx zpp.Ctx, vars []types.Zval) {
	for _, zv := range vars {
		PhpVarDump(ctx, zv, 1)
	}
}
