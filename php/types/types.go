package types

import "github.com/heyuuu/gophp/php/lang"

// ZvalType
type ZvalType uint8

const (
	/* regular data types */
	IsUndef ZvalType = iota
	IsNull
	IsFalse
	IsTrue
	IsLong
	IsDouble
	IsString
	IsArray
	IsObject
	IsResource
	IsRef

	/* constant expressions */
	IsConstantAst ZvalType = 11

	/* internal types */
	IsIndirect ZvalType = 13
	IsPtr      ZvalType = 14
	IsError    ZvalType = 15 // _IS_ERROR

	/* fake types used only for type hinting (Z_TYPE(zv) can not use them) */
	IsBool     ZvalType = 16 // _IS_BOOL
	IsCallable ZvalType = 17
	IsIterable ZvalType = 18
	IsVoid     ZvalType = 19
	IsNumber   ZvalType = 20 // _IS_NUMBER
)

/**
 * Zval
 *
 * tips:
 * - 零值为合法的 Undef 类型，无需 SetUndef() 初始化
 */
type Zval struct {
	value any
	typ   ZvalType
	u2    uint32
	//u2 struct /* union */ {
	//	next           uint32 /* hash collision chain  用来解决哈希冲突问题，记录冲突的下一个元素位置*/
	//	cache_slot     uint32 /* literal cache slot  运行时缓存。在执行函数时会优先去缓存中查找，若缓存中没有，会在全局的function表中查找*/
	//	opline_num     uint32
	//	lineno         uint32 /* line number (for ast nodes) 文件执行的行号，应用在AST节点上 */
	//	numArgs       uint32 /* arguments number for EX(This) 函数调用时传入参数的个数 */
	//	fe_pos         uint32 /* foreach position  遍历数组时的当前位置*/
	//	fe_iter_idx    uint32 /* foreach iterator index */
	//	access_flags   uint32 /* class constant access flags 对象类的访问标志*/
	//	property_guard uint32 /* single property guard  防止类中魔术方法的循环调用*/
	//	constant_flags uint32
	//	extra          uint32
	//}
}

// new
func NewZvalUndef() *Zval                 { var zv Zval; zv.SetUndef(); return &zv }
func NewZvalNull() *Zval                  { var zv Zval; zv.SetNull(); return &zv }
func NewZvalFalse() *Zval                 { var zv Zval; zv.SetFalse(); return &zv }
func NewZvalTrue() *Zval                  { var zv Zval; zv.SetTrue(); return &zv }
func NewZvalBool(b bool) *Zval            { var zv Zval; zv.SetBool(b); return &zv }
func NewZvalLong(l int) *Zval             { var zv Zval; zv.SetLong(l); return &zv }
func NewZvalDouble(d float64) *Zval       { var zv Zval; zv.SetDouble(d); return &zv }
func NewZvalString(s string) *Zval        { var zv Zval; zv.SetStringVal(s); return &zv }
func NewZvalArray(a *Array) *Zval         { var zv Zval; zv.SetArray(a); return &zv }
func NewZvalEmptyArray() *Zval            { var zv Zval; zv.SetEmptyArray(); return &zv }
func NewZvalObject(o *Object) *Zval       { var zv Zval; zv.SetObject(o); return &zv }
func NewZvalResource(res *Resource) *Zval { var zv Zval; zv.SetResource(res); return &zv }
func NewZvalPtr(ptr any) *Zval            { var zv Zval; zv.SetPtr(ptr); return &zv }
func NewZvalIndirect(z *Zval) *Zval       { var zv Zval; zv.SetIndirect(z); return &zv }

// setter
func (zv *Zval) SetUndef()                 { zv.typ, zv.value = IsUndef, nil }
func (zv *Zval) SetNull()                  { zv.typ, zv.value = IsNull, nil }
func (zv *Zval) SetFalse()                 { zv.typ, zv.value = IsFalse, nil }
func (zv *Zval) SetTrue()                  { zv.typ, zv.value = IsTrue, nil }
func (zv *Zval) SetBool(b bool)            { zv.typ, zv.value = lang.Cond(b, IsTrue, IsFalse), nil }
func (zv *Zval) SetLong(l int)             { zv.typ, zv.value = IsLong, l }
func (zv *Zval) SetDouble(d float64)       { zv.typ, zv.value = IsDouble, d }
func (zv *Zval) SetStringVal(s string)     { zv.typ, zv.value = IsString, NewString(s) }
func (zv *Zval) SetString(s *String)       { zv.typ, zv.value = IsString, s }
func (zv *Zval) SetEmptyArray()            { zv.typ, zv.value = IsArray, NewArray(0) }
func (zv *Zval) SetArray(arr *Array)       { zv.typ, zv.value = IsArray, arr }
func (zv *Zval) SetObject(obj *Object)     { zv.typ, zv.value = IsObject, obj }
func (zv *Zval) SetResource(res *Resource) { zv.typ, zv.value = IsResource, res }

func (zv *Zval) SetArrayOfInt(arr []int)         { zv.SetArray(NewArrayOfInt(arr)) }
func (zv *Zval) SetArrayOfString(arr []string)   { zv.SetArray(NewArrayOfString(arr)) }
func (zv *Zval) SetArrayOfZval(arr []*Zval)      { zv.SetArray(NewArrayOfZval(arr)) }
func (zv *Zval) SetReference(ref *ZendReference) { zv.typ, zv.value = IsRef, ref }
func (zv *Zval) SetNewEmptyRef()                 { zv.SetReference(NewZendReference(nil)) }
func (zv *Zval) SetNewRef(val *Zval)             { zv.SetReference(NewZendReference(val)) }
func (zv *Zval) SetConstantAst(ast *ZendAstRef)  { zv.typ, zv.value = IsConstantAst, ast }
func (zv *Zval) SetIndirect(v *Zval)             { zv.typ, zv.value = IsIndirect, v }
func (zv *Zval) SetPtr(ptr any)                  { zv.typ, zv.value = IsPtr, ptr }
func (zv *Zval) SetCe(value *ClassEntry)         { zv.typ, zv.value = IsPtr, value }
func (zv *Zval) SetFunc(value IFunction)         { zv.typ, zv.value = IsPtr, value }
func (zv *Zval) SetIsError()                     { zv.typ, zv.value = IsError, nil }
func (zv *Zval) SetBy(val *Zval)                 { zv.CopyValueFrom(val) }
