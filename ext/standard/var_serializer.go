package standard

import (
	"fmt"
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
	"strconv"
	"strings"
	"unsafe"
)

type VarSerializer struct {
	ctx  *php.Context
	buf  strings.Builder
	data *PhpSerializeData
}

func InitSerializer(ctx *php.Context) *VarSerializer {
	return &VarSerializer{
		ctx: ctx,
		//data: PhpVarSerializeInit(),
		data: NewPhpSerializeData(),
	}
}

func (se *VarSerializer) DestroyData() {
	//PhpVarSerializeDestroy(se.data)
}

func (se *VarSerializer) String() string {
	return se.buf.String()
}

func (se *VarSerializer) WriteString(str string) {
	se.buf.WriteString(str)
}
func (se *VarSerializer) WriteByte(c byte) {
	se.buf.WriteByte(c)
}
func (se *VarSerializer) WriteLong(i int) {
	str := strconv.FormatInt(int64(i), 10)
	se.buf.WriteString(str)
}

func (se *VarSerializer) Serialize(zv types.Zval) {
	se.serializeIntern(zv)
}

func (se *VarSerializer) serializeLong(val int) {
	se.WriteString("i:")
	se.WriteLong(val)
	se.WriteByte(';')
}
func (se *VarSerializer) serializeString(str string) {
	se.WriteString("s:")
	se.WriteLong(len(str))
	se.WriteString(`:"`)
	se.WriteString(str)
	se.WriteString(`";`)
}
func (se *VarSerializer) serializeClassName(obj *types.Object) bool {
	className, incompleteClass := obj.ClassName(), false
	se.WriteString("O:")
	se.WriteLong(len(className))
	se.WriteString(`:"`)
	se.WriteString(className)
	se.WriteString(`":`)
	return incompleteClass
}
func (se *VarSerializer) addVarHash(v types.Zval) int {
	se.data.IncN()

	var isRef = v.IsRef()
	if !isRef && !v.IsObject() {
		return 0
	}

	/* References to objects are treated as if the reference didn't exist */
	if isRef && v.DeRef().IsObject() {
		v = v.DeRef()
	}

	/* Index for the variable is stored using the numeric value of the pointer to
	 * the zend_refcounted struct */
	if n, exists := se.data.FindMark(v); exists {
		/* References are only counted once, undo the data->n increment above */
		if isRef && n != -1 {
			se.data.DecN()
		}
		return n
	} else {
		se.data.Mark(v)
		return 0
	}
}
func (se *VarSerializer) serializeNestedData(zv types.Zval, ht *types.Array, count int, incompleteClass bool) {
	se.WriteLong(count)
	se.WriteString(":{")
	if count > 0 {
		ht.Each(func(key types.ArrayKey, data types.Zval) {
			//if incompleteClass && key.StrKey() == MAGIC_MEMBER {
			//	return
			//}
			if key.IsStrKey() {
				se.serializeString(key.StrKey())
			} else {
				se.serializeLong(key.IdxKey())
			}

			/* we should still add element even if it's not OK,
			 * since we already wrote the length of the array before */
			if data.IsArray() {
				if data.Array().IsRecursive() || zv.IsArray() && data.Array() == zv.Array() {
					se.addVarHash(zv)
					se.WriteString("N;")
				} else {
					data.Array().ProtectRecursive()
					se.serializeIntern(data)
					data.Array().UnprotectRecursive()
				}
			} else {
				se.serializeIntern(data)
			}
		})
	}
	se.WriteByte('}')
}
func (se *VarSerializer) tryAddSleepProp(ht *types.Array, props *types.Array, name string, errorName string, obj *types.Object) bool {
	var val = props.KeyFind(name)
	if val.IsUndef() {
		return false
	}
	//if val.IsIndirect() {
	//	val = val.IndirectVal()
	//	if val.IsUndef() {
	//		var info = php.ZendGetTypedPropertyInfoForSlot(obj, &val)
	//		if info != nil {
	//			return true
	//		}
	//		return false
	//	}
	//}
	if !ht.KeyAdd(name, val) {
		php.ErrorDocRef(se.ctx, "", perr.E_NOTICE, fmt.Sprintf(`"%s" is returned from __sleep multiple times`, errorName))
		return true
	}
	return true
}
func (se *VarSerializer) getSleepProps(obj *types.Object, sleepRetval *types.Array) (*types.Array, bool) {
	var ce = obj.Class()
	var props = obj.GetPropertiesFor(types.PropPurposeSerialize)
	var retval = true
	var ht = types.NewArrayCap(sleepRetval.Len())

	/* TODO: Rewrite this by fetching the property info instead of trying out different
	 * name manglings? */
	_ = sleepRetval.EachEx(func(_ types.ArrayKey, nameVal types.Zval) error {
		nameVal = nameVal.DeRef()
		if !nameVal.IsString() {
			php.ErrorDocRef(se.ctx, "", perr.E_NOTICE, "__sleep should return an array only containing the names of instance-variables to serialize.")
		}

		name := php.ZvalGetStrVal(se.ctx, nameVal)
		if se.tryAddSleepProp(ht, props, name, name, obj) {
			return nil
		}
		if se.ctx.EG().HasException() {
			retval = false
			return lang.BreakErr
		}

		privateName := php.ManglePropertyName(ce.Name(), name)
		if se.tryAddSleepProp(ht, props, privateName, name, obj) {
			return nil
		}
		if se.ctx.EG().HasException() {
			retval = false
			return lang.BreakErr
		}

		protectedName := php.ManglePropertyName("*", name)
		if se.tryAddSleepProp(ht, props, protectedName, name, obj) {
			return nil
		}
		if se.ctx.EG().HasException() {
			retval = false
			return lang.BreakErr
		}
		php.ErrorDocRef(se.ctx, "", perr.E_NOTICE, fmt.Sprintf(`"%s" returned as member variable from __sleep() but does not exist`, name))
		ht.KeyAdd(name, php.UninitializedZval())
		return nil
	})
	return ht, retval
}
func (se *VarSerializer) serializeClass(obj *types.Object, retval types.Zval) {
	if props, ok := se.getSleepProps(obj, php.HashOf(retval)); ok {
		se.serializeClassName(obj)
		se.serializeNestedData(types.ZvalObject(obj), props, props.Count(), false)
	}
}
func (se *VarSerializer) serializeIntern(zv types.Zval) {
	if se.ctx.EG().HasException() {
		return
	}
	if se.data != nil {
		if varAlready := se.addVarHash(zv); varAlready != 0 {
			if varAlready == -1 {
				/* Reference to an object that failed to serialize, replace with null. */
				se.WriteString("N;")
				return
			} else if zv.IsRef() {
				se.WriteString("R:")
				se.WriteLong(varAlready)
				se.WriteByte(';')
				return
			} else if zv.IsObject() {
				se.WriteString("r:")
				se.WriteLong(varAlready)
				se.WriteByte(';')
				return
			}
		}
	}
again:
	switch zv.Type() {
	case types.IsFalse:
		se.WriteString("b:0;")
		return
	case types.IsTrue:
		se.WriteString("b:1;")
		return
	case types.IsNull:
		se.WriteString("N;")
		return
	case types.IsLong:
		se.serializeLong(zv.Long())
		return
	case types.IsDouble:
		doubleStr := php.FormatDouble(zv.Double(), 'E', se.ctx.PG().SerializePrecision())
		se.WriteString("d:")
		se.WriteString(doubleStr)
		se.WriteByte(';')
		return
	case types.IsString:
		se.serializeString(zv.String())
		return
	case types.IsObject:
		var obj = zv.Object()
		//var ce = obj.Class()
		//if ce.FunctionTable().Exists("__serialize") {
		//	var retval types.Zval
		//	retval, ok := se.callSerialize(obj)
		//	if !ok {
		//		if se.ctx.EG().HasException() {
		//			se.WriteString("N;")
		//		}
		//		return
		//	}
		//	se.serializeClassName(obj)
		//	se.WriteLong(retval.Array().Count())
		//	se.WriteString(":{")
		//	retval.Array().Each(func(key types.ArrayKey, data types.Zval) {
		//		if key.IsStrKey() {
		//			se.serializeString(key.StrKey())
		//		} else {
		//			se.serializeLong(key.IdxKey())
		//		}
		//		se.serializeIntern(data)
		//	})
		//	se.WriteByte('}')
		//	return
		//}
		//if ce.HasSerialize() {
		//	/* has custom handler */
		//	if serializedData, ok := ce.DoSerialize(&zv); ok {
		//		ceName := ce.Name()
		//		se.WriteString("C:")
		//		se.WriteLong(len(ceName))
		//		se.WriteString(`:"`)
		//		se.WriteString(ceName)
		//		se.WriteString(`":`)
		//		se.WriteLong(len(serializedData))
		//		se.WriteString(":{")
		//		se.WriteString(serializedData)
		//		se.WriteByte('}')
		//	} else {
		//		/* Mark this value in the var_hash, to avoid creating references to it. */
		//		se.data.MarkUsed(zv)
		//		se.WriteString("N;")
		//	}
		//	return
		//}
		//if ce.FunctionTable().Exists("__sleep") {
		//	retval, ok := se.callSleep(obj)
		//	if !ok {
		//		if se.ctx.EG().NoException() {
		//			/* we should still add element even if it's not OK,
		//			 * since we already wrote the length of the array before */
		//			se.WriteString("N;")
		//		}
		//		return
		//	}
		//	se.serializeClass(obj, retval)
		//	return
		//}
		incompleteClass := se.serializeClassName(obj)
		myht := obj.GetPropertiesFor(types.PropPurposeSerialize)

		/* count after serializing name, since php_var_serialize_class_name
		 * changes the count if the variable is incomplete class */

		count := myht.Count()
		if count > 0 && incompleteClass {
			count--
		}
		se.serializeNestedData(zv, myht, count, incompleteClass)
		return
	case types.IsArray:
		se.WriteString("a:")
		myht := zv.Array()
		se.serializeNestedData(zv, myht, myht.Count(), false)
		return
	case types.IsRef:
		zv = zv.DeRef()
		goto again
	default:
		se.WriteString("i:0;")
		return
	}
}

//	func (se *VarSerializer) callSleep(obj *types.Object) (types.Zval, bool) {
//		BG__().serializeLock++
//		retval, ok := php.CallUserFunction(se.ctx, obj, "__sleep")
//		BG__().serializeLock--
//
//		if !ok || retval.IsUndef() {
//			return types.Undef, false
//		}
//		if php.HashOf(retval) == nil {
//			php.ErrorDocRef(se.ctx, "", perr.E_NOTICE, "__sleep should return an array only containing the names of instance-variables to serialize")
//			return types.Undef, false
//		}
//		return retval, true
//	}
//func (se *VarSerializer) callSerialize(obj *types.Object) (types.Zval, bool) {
//	BG__().serializeLock++
//	retval, ok := php.CallUserFunction(se.ctx, obj, "__serialize")
//	BG__().serializeLock--
//
//	if !ok || retval.IsUndef() {
//		return types.Undef, false
//	}
//	if !retval.IsArray() {
//		php.TypeError(se.ctx, fmt.Sprintf("%s::__serialize() must return an array", obj.ClassName()))
//		return types.Undef, false
//	}
//	return retval, true
//}

// PhpSerializeData
type PhpSerializeData struct {
	nMap map[unsafe.Pointer]int
	n    uint32
}

func NewPhpSerializeData() *PhpSerializeData {
	return &PhpSerializeData{
		nMap: make(map[unsafe.Pointer]int),
		n:    0,
	}
}

func (d *PhpSerializeData) Destroy() {
	d.nMap = nil
	d.n = 0
}

func (d *PhpSerializeData) IncN() { d.n++ }
func (d *PhpSerializeData) DecN() { d.n-- }

func (d *PhpSerializeData) zvalKey(zv types.Zval) unsafe.Pointer {
	php.Assert(zv.IsRef() || zv.IsObject())
	if zv.IsRef() {
		return unsafe.Pointer(zv.Ref())
	} else if zv.IsObject() {
		return unsafe.Pointer(zv.Object())
	} else {
		panic("unreachable")
	}
}

func (d *PhpSerializeData) Mark(zv types.Zval) {
	key := d.zvalKey(zv)
	d.nMap[key] = int(d.n)
}

func (d *PhpSerializeData) MarkUsed(zv types.Zval) {
	key := d.zvalKey(zv)
	d.nMap[key] = -1
}

func (d *PhpSerializeData) FindMark(zv types.Zval) (n int, exists bool) {
	key := d.zvalKey(zv)
	n, ok := d.nMap[key]
	return n, ok
}
