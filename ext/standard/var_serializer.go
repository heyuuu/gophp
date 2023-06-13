package standard

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
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
func (se *VarSerializer) WriteUlong(i uint) {
	str := strconv.FormatUint(uint64(i), 10)
	se.buf.WriteString(str)
}
func (se *VarSerializer) SerializeLong(val zend.ZendLong) {
	se.WriteString("i:")
	se.WriteLong(val)
	se.WriteByte(';')
}
func (se *VarSerializer) SerializeString(str string) {
	se.WriteString("s:")
	se.WriteLong(len(str))
	se.WriteString(":\"")
	se.WriteString(str)
	se.WriteString("\";")
}
func (se *VarSerializer) SerializeClassName(struc *types.Zval) bool {
	className, incompleteClass := PhpClassAttributes(struc)
	PHP_SET_CLASS_ATTRIBUTES(struc)
	se.WriteString("O:")
	se.WriteLong(len(className))
	se.WriteString(":\"")
	se.WriteString(className)
	se.WriteString("\":")
	return incompleteClass
}
func (se *VarSerializer) AddVarHash(var_ *types.Zval) int {
	se.data.IncN()

	var isRef = var_.IsReference()
	if !isRef && !var_.IsObject() {
		return 0
	}

	/* References to objects are treated as if the reference didn't exist */
	if isRef && types.Z_REFVAL_P(var_).IsType(types.IS_OBJECT) {
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
func (se *VarSerializer) SerializeNestedData(struc *types.Zval, ht *types.Array, count int, incompleteClass bool) {
	se.WriteUlong(uint(count))
	se.WriteString(":{")
	if count > 0 {
		ht.ForeachIndirectEx(func(key types.ArrayKey, data *types.Zval) bool {
			if incompleteClass && key.StrKey() == MAGIC_MEMBER {
				return true
			}
			if key.IsStrKey() {
				se.SerializeString(key.StrKey())
			} else {
				se.SerializeLong(key.IdxKey())
			}

			/* we should still add element even if it's not OK,
			 * since we already wrote the length of the array before */
			if data.IsArray() {
				if data.Array().IsRecursive() || struc.IsArray() && data.Array() == struc.Array() {
					se.AddVarHash(struc)
					se.WriteString("N;")
				} else {
					data.Array().ProtectRecursive()
					se.Serialize(data)
					data.Array().UnprotectRecursive()
				}
			} else {
				se.Serialize(data)
			}

			return true
		})
	}
	se.WriteByte('}')
}
func (se *VarSerializer) SerializeClass(struc *types.Zval, retval_ptr *types.Zval) {
	var props types.Array
	if PhpVarSerializeGetSleepProps(&props, struc, zend.HASH_OF(retval_ptr)) == types.SUCCESS {
		se.SerializeClassName(struc)
		se.SerializeNestedData(struc, &props, props.Count(), false)
	}
	props.Destroy()
}
func (se *VarSerializer) Serialize(struc *types.Zval) {
	if zend.EG__().GetException() != nil {
		return
	}
	if se.data != nil {
		if varAlready := se.AddVarHash(struc); varAlready != 0 {
			if varAlready == -1 {
				/* Reference to an object that failed to serialize, replace with null. */
				se.WriteString("N;")
				return
			} else if struc.IsReference() {
				se.WriteString("R:")
				se.WriteLong(varAlready)
				se.WriteByte(';')
				return
			} else if struc.IsType(types.IS_OBJECT) {
				se.WriteString("r:")
				se.WriteLong(varAlready)
				se.WriteByte(';')
				return
			}
		}
	}
again:
	switch struc.GetType() {
	case types.IS_FALSE:
		se.WriteString("b:0;")
		return
	case types.IS_TRUE:
		se.WriteString("b:1;")
		return
	case types.IS_NULL:
		se.WriteString("N;")
		return
	case types.IS_LONG:
		se.SerializeLong(struc.Long())
		return
	case types.IS_DOUBLE:
		var tmp_str []byte
		se.WriteString("d:")
		core.PhpGcvt(struc.Double(), int(core.PG__().serialize_precision), '.', 'E', tmp_str)
		se.WriteString(b.CastStrAuto(tmp_str))
		se.WriteByte(';')
		return
	case types.IS_STRING:
		se.SerializeString(struc.StringVal())
		return
	case types.IS_OBJECT:
		var ce = types.Z_OBJCE_P(struc)
		if ce.FunctionTable().Exists("__serialize") {
			var retval types.Zval
			var obj types.Zval
			obj.SetObject(struc.Object())
			if PhpVarSerializeCallMagicSerialize(&retval, &obj) == types.FAILURE {
				if zend.EG__().GetException() == nil {
					se.WriteString("N;")
				}
				return
			}
			se.SerializeClassName(&obj)
			se.WriteLong(retval.Array().Count())
			se.WriteString(":{")
			retval.Array().ForeachIndirect(func(key types.ArrayKey, data *types.Zval) {
				if key.IsStrKey() {
					se.SerializeString(key.StrKey())
				} else {
					se.SerializeLong(key.IdxKey())
				}
				se.Serialize(data)
			})
			se.WriteByte('}')
			return
		}
		if ce.GetSerialize() != nil {

			/* has custom handler */
			var serialized_data *uint8 = nil
			var serialized_length int
			if ce.GetSerialize()(struc, &serialized_data, &serialized_length, se.data) == types.SUCCESS {
				se.WriteString("C:")
				se.WriteLong(types.Z_OBJCE_P(struc).GetName().GetLen())
				se.WriteString(":\"")
				se.WriteString(types.Z_OBJCE_P(struc).GetName().GetStr())
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
			if PhpVarSerializeCallSleep(&retval, &tmp) == types.FAILURE {
				if zend.EG__().GetException() == nil {

					/* we should still add element even if it's not OK,
					 * since we already wrote the length of the array before */
					se.WriteString("N;")
				}
				return
			}
			se.SerializeClass(&tmp, &retval)
			return
		}
		incompleteClass := se.SerializeClassName(struc)
		myht := zend.ZendGetPropertiesFor(struc, zend.ZEND_PROP_PURPOSE_SERIALIZE)

		/* count after serializing name, since php_var_serialize_class_name
		 * changes the count if the variable is incomplete class */

		count := myht.Count()
		if count > 0 && incompleteClass {
			count--
		}
		se.SerializeNestedData(struc, myht, count, incompleteClass)
		return
	case types.IS_ARRAY:
		se.WriteString("a:")
		myht := struc.Array()
		se.SerializeNestedData(struc, myht, myht.Count(), false)
		return
	case types.IS_REFERENCE:
		struc = types.Z_REFVAL_P(struc)
		goto again
	default:
		se.WriteString("i:0;")
		return
	}
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
	b.Assert(zv.IsReference() || zv.IsObject())
	if zv.IsReference() {
		return unsafe.Pointer(zv.Reference())
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
