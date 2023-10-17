package standard

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
	"strconv"
	"strings"
	"unsafe"
)

type VarSerializer struct {
	buf  strings.Builder
	data *PhpSerializeData
}

func InitSerializer() *VarSerializer {
	return &VarSerializer{data: PhpVarSerializeInit()}
}

func (se *VarSerializer) DestroyData() {
	PhpVarSerializeDestroy(se.data)
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

func (se *VarSerializer) Serialize(struc *types.Zval) {
	se.serializeIntern(struc)
}

func (se *VarSerializer) serializeLong(val zend.ZendLong) {
	se.WriteString("i:")
	se.WriteLong(val)
	se.WriteByte(';')
}
func (se *VarSerializer) serializeString(str string) {
	se.WriteString("s:")
	se.WriteLong(len(str))
	se.WriteString(":\"")
	se.WriteString(str)
	se.WriteString("\";")
}
func (se *VarSerializer) serializeClassName(struc *types.Zval) bool {
	className, incompleteClass := PhpClassAttributes(struc)
	se.WriteString("O:")
	se.WriteLong(len(className))
	se.WriteString(":\"")
	se.WriteString(className)
	se.WriteString("\":")
	return incompleteClass
}
func (se *VarSerializer) addVarHash(var_ *types.Zval) int {
	se.data.IncN()

	var isRef = var_.IsRef()
	if !isRef && !var_.IsObject() {
		return 0
	}

	/* References to objects are treated as if the reference didn't exist */
	if isRef && types.Z_REFVAL_P(var_).IsType(types.IsObject) {
		var_ = types.Z_REFVAL_P(var_)
	}

	/* Index for the variable is stored using the numeric value of the pointer to
	 * the zend_refcounted struct */
	if n, exists := se.data.FindMark(var_); exists {
		/* References are only counted once, undo the data->n increment above */
		if isRef && n != -1 {
			se.data.DecN()
		}
		return n
	} else {
		se.data.Mark(var_)
		return 0
	}
}
func (se *VarSerializer) serializeNestedData(struc *types.Zval, ht *types.Array, count int, incompleteClass bool) {
	se.WriteLong(count)
	se.WriteString(":{")
	if count > 0 {
		ht.ForeachIndirectEx(func(key types.ArrayKey, data *types.Zval) bool {
			if incompleteClass && key.StrKey() == MAGIC_MEMBER {
				return true
			}
			if key.IsStrKey() {
				se.serializeString(key.StrKey())
			} else {
				se.serializeLong(key.IdxKey())
			}

			/* we should still add element even if it's not OK,
			 * since we already wrote the length of the array before */
			if data.IsArray() {
				if data.Array().IsRecursive() || struc.IsArray() && data.Array() == struc.Array() {
					se.addVarHash(struc)
					se.WriteString("N;")
				} else {
					data.Array().ProtectRecursive()
					se.serializeIntern(data)
					data.Array().UnprotectRecursive()
				}
			} else {
				se.serializeIntern(data)
			}

			return true
		})
	}
	se.WriteByte('}')
}
func (se *VarSerializer) tryAddSleepProp(ht *types.Array, props *types.Array, name *types.String, errorName *types.String, struc *types.Zval) int {
	var val = props.KeyFind(name.GetStr())
	if val == nil {
		return types.FAILURE
	}
	if val.IsIndirect() {
		val = val.Indirect()
		if val.IsUndef() {
			var info = zend.ZendGetTypedPropertyInfoForSlot(struc.Object(), val)
			if info != nil {
				return types.SUCCESS
			}
			return types.FAILURE
		}
	}
	if ht.KeyAdd(name.GetStr(), val) == nil {
		core.PhpErrorDocref("", faults.E_NOTICE, "\"%s\" is returned from __sleep multiple times", errorName.GetVal())
		return types.SUCCESS
	}
	return types.SUCCESS
}
func (se *VarSerializer) getSleepProps(ht *types.Array, struc *types.Zval, sleep_retval *types.Array) int {
	var ce = types.Z_OBJCE_P(struc)
	var props = zend.ZendGetPropertiesFor(struc, zend.ZEND_PROP_PURPOSE_SERIALIZE)
	var name_val *types.Zval
	var retval = types.SUCCESS
	*ht = *types.NewArrayCap(sleep_retval.Len())

	/* TODO: Rewrite this by fetching the property info instead of trying out different
	 * name manglings? */

	var __ht = sleep_retval
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()
		if _z.IsIndirect() {
			_z = _z.Indirect()
			if _z.IsUndef() {
				continue
			}
		}
		name_val = _z
		var name *types.String
		var priv_name *types.String
		var prot_name *types.String
		name_val = types.ZVAL_DEREF(name_val)
		if !name_val.IsString() {
			core.PhpErrorDocref("", faults.E_NOTICE, "__sleep should return an array only containing the names of instance-variables to serialize.")
		}
		name = operators.ZvalGetString(name_val)
		if se.tryAddSleepProp(ht, props, name, name, struc) == types.SUCCESS {
			continue
		}
		if zend.EG__().HasException() {
			retval = types.FAILURE
			break
		}
		priv_name = zend.ZendManglePropertyName_ZStr(ce.Name(), name.GetStr())
		if se.tryAddSleepProp(ht, props, priv_name, name, struc) == types.SUCCESS {
			continue
		}
		if zend.EG__().HasException() {
			retval = types.FAILURE
			break
		}
		prot_name = zend.ZendManglePropertyName_ZStr("*", name.GetStr())
		if se.tryAddSleepProp(ht, props, prot_name, name, struc) == types.SUCCESS {
			continue
		}
		if zend.EG__().HasException() {
			retval = types.FAILURE
			break
		}
		core.PhpErrorDocref("", faults.E_NOTICE, "\"%s\" returned as member variable from __sleep() but does not exist", name.GetVal())
		ht.KeyAdd(name.GetStr(), zend.UninitializedZval())
	}
	return retval
}
func (se *VarSerializer) serializeClass(struc *types.Zval, retval_ptr *types.Zval) {
	var props types.Array
	if se.getSleepProps(&props, struc, zend.HASH_OF(retval_ptr)) == types.SUCCESS {
		se.serializeClassName(struc)
		se.serializeNestedData(struc, &props, props.Count(), false)
	}
	props.Destroy()
}
func (se *VarSerializer) serializeIntern(struc *types.Zval) {
	if zend.EG__().HasException() {
		return
	}
	if se.data != nil {
		if varAlready := se.addVarHash(struc); varAlready != 0 {
			if varAlready == -1 {
				/* Reference to an object that failed to serialize, replace with null. */
				se.WriteString("N;")
				return
			} else if struc.IsRef() {
				se.WriteString("R:")
				se.WriteLong(varAlready)
				se.WriteByte(';')
				return
			} else if struc.IsType(types.IsObject) {
				se.WriteString("r:")
				se.WriteLong(varAlready)
				se.WriteByte(';')
				return
			}
		}
	}
again:
	switch struc.Type() {
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
		se.serializeLong(struc.Long())
		return
	case types.IsDouble:
		var tmp_str []byte
		se.WriteString("d:")
		core.PhpGcvt(struc.Double(), int(core.PG__().serialize_precision), '.', 'E', tmp_str)
		se.WriteString(b.CastStrAuto(tmp_str))
		se.WriteByte(';')
		return
	case types.IsString:
		se.serializeString(struc.String())
		return
	case types.IsObject:
		var ce = types.Z_OBJCE_P(struc)
		if ce.FunctionTable().Exists("__serialize") {
			var retval types.Zval
			var obj types.Zval
			obj.SetObject(struc.Object())
			if !se.callSerialize(&retval, &obj) {
				if zend.EG__().NoException() {
					se.WriteString("N;")
				}
				return
			}
			se.serializeClassName(&obj)
			se.WriteLong(retval.Array().Count())
			se.WriteString(":{")
			retval.Array().ForeachIndirect(func(key types.ArrayKey, data *types.Zval) {
				if key.IsStrKey() {
					se.serializeString(key.StrKey())
				} else {
					se.serializeLong(key.IdxKey())
				}
				se.serializeIntern(data)
			})
			se.WriteByte('}')
			return
		}
		if ce.GetSerialize() != nil {

			/* has custom handler */
			var serialized_data *uint8 = nil
			var serialized_length int
			if ce.GetSerialize()(struc, &serialized_data, &serialized_length, se.data) == types.SUCCESS {
				ceName := types.Z_OBJCE_P(struc).Name()
				se.WriteString("C:")
				se.WriteLong(len(ceName))
				se.WriteString(":\"")
				se.WriteString(ceName)
				se.WriteString("\":")
				se.WriteLong(serialized_length)
				se.WriteString(":{")
				se.WriteString(b.CastStr((*byte)(serialized_data), serialized_length))
				se.WriteByte('}')
			} else {
				/* Mark this value in the var_hash, to avoid creating references to it. */
				se.data.MarkUsed(struc)
				se.WriteString("N;")
			}
			if serialized_data != nil {
				zend.Efree(serialized_data)
			}
			return
		}
		if ce != PHP_IC_ENTRY && ce.FunctionTable().Exists("__sleep") {
			var retval types.Zval
			var tmp types.Zval
			tmp.SetObject(struc.Object())
			if !se.callSleep(&retval, &tmp) {
				if zend.EG__().NoException() {
					/* we should still add element even if it's not OK,
					 * since we already wrote the length of the array before */
					se.WriteString("N;")
				}
				return
			}
			se.serializeClass(&tmp, &retval)
			return
		}
		incompleteClass := se.serializeClassName(struc)
		myht := zend.ZendGetPropertiesFor(struc, zend.ZEND_PROP_PURPOSE_SERIALIZE)

		/* count after serializing name, since php_var_serialize_class_name
		 * changes the count if the variable is incomplete class */

		count := myht.Count()
		if count > 0 && incompleteClass {
			count--
		}
		se.serializeNestedData(struc, myht, count, incompleteClass)
		return
	case types.IsArray:
		se.WriteString("a:")
		myht := struc.Array()
		se.serializeNestedData(struc, myht, myht.Count(), false)
		return
	case types.IsRef:
		struc = types.Z_REFVAL_P(struc)
		goto again
	default:
		se.WriteString("i:0;")
		return
	}
}

func (se *VarSerializer) callSleep(retval *types.Zval, struc *types.Zval) bool {
	BG__().serialize_lock++
	res := zend.CallUserFunction_Ex(struc, types.NewZvalString("__sleep"), retval, nil)
	BG__().serialize_lock--

	if res == types.FAILURE || retval.IsUndef() {
		return false
	}
	if zend.HASH_OF(retval) == nil {
		core.PhpErrorDocref("", faults.E_NOTICE, "__sleep should return an array only containing the names of instance-variables to serialize")
		return false
	}
	return true
}
func (se *VarSerializer) callSerialize(retval *types.Zval, obj *types.Zval) bool {
	BG__().serialize_lock++
	res := zend.CallUserFunction_Ex(obj, types.NewZvalString("__serialize"), retval, nil)
	BG__().serialize_lock--

	if res == types.FAILURE || retval.IsUndef() {
		return false
	}
	if !retval.IsArray() {
		faults.TypeError(fmt.Sprintf("%s::__serialize() must return an array", types.Z_OBJCE_P(obj).Name()))
		return false
	}
	return true
}

/**
 * types
 */
type PhpSerializeDataT = *PhpSerializeData

/**
 * PhpSerializeData
 */
type PhpSerializeData struct {
	nMap map[unsafe.Pointer]int
	vMap map[unsafe.Pointer]*types.Zval
	n    uint32
}

func NewPhpSerializeData() *PhpSerializeData {
	return &PhpSerializeData{
		nMap: make(map[unsafe.Pointer]int),
		vMap: make(map[unsafe.Pointer]*types.Zval),
		n:    0,
	}
}

func (d *PhpSerializeData) Destroy() {
	d.nMap = nil
	d.vMap = nil
	d.n = 0
}

func (d *PhpSerializeData) IncN() { d.n++ }
func (d *PhpSerializeData) DecN() { d.n-- }

func (d *PhpSerializeData) zvalKey(zv *types.Zval) unsafe.Pointer {
	b.Assert(zv.IsRef() || zv.IsObject())
	if zv.IsRef() {
		return unsafe.Pointer(zv.Ref())
	} else if zv.IsObject() {
		return unsafe.Pointer(zv.Object())
	} else {
		panic("unreachable")
	}
}

func (d *PhpSerializeData) Mark(zv *types.Zval) {
	key := d.zvalKey(zv)

	//d.ht.IndexAdd(key, types.NewZvalLong(int(d.n)))
	d.nMap[key] = int(d.n)

	/* Additionally to the index, we also store the variable, to ensure that it is
	 * not destroyed during serialization and its pointer reused. The variable is
	 * stored at the numeric value of the pointer + 1, which cannot be the location
	 * of another zend_refcounted structure. */
	//d.ht.IndexAdd(key+1, zv)
	d.vMap[key] = zv
}

func (d *PhpSerializeData) MarkUsed(zv *types.Zval) {
	key := d.zvalKey(zv)
	d.nMap[key] = -1
}

func (d *PhpSerializeData) FindMark(zv *types.Zval) (n int, exists bool) {
	key := d.zvalKey(zv)
	n, ok := d.nMap[key]
	return n, ok
}
