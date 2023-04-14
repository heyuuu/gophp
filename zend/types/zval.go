package types

import (
	b "github.com/heyuuu/gophp/builtin"
)

/**
 * Zval
 */
type Zval struct {
	value     any
	typ       ZendUchar
	typeFlags ZendUchar
	u2        uint32
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

/** new */
func NewZvalUndef() *Zval                     { var tmp Zval; tmp.SetUndef(); return &tmp }
func NewZvalNull() *Zval                      { var tmp Zval; tmp.SetNull(); return &tmp }
func NewZvalBool(b bool) *Zval                { var tmp Zval; tmp.SetBool(b); return &tmp }
func NewZvalString(str string) *Zval          { var tmp Zval; tmp.SetStringVal(str); return &tmp }
func NewZvalLong(l int) *Zval                 { var tmp Zval; tmp.SetLong(l); return &tmp }
func NewZvalDouble(d float64) *Zval           { var tmp Zval; tmp.SetDouble(d); return &tmp }
func NewZvalArray(arr *Array) *Zval           { var tmp Zval; tmp.SetArray(arr); return &tmp }
func NewZvalResource(res *ZendResource) *Zval { var tmp Zval; tmp.SetResource(res); return &tmp }
func NewZvalPtr(ptr any) *Zval                { var tmp Zval; tmp.SetPtr(ptr); return &tmp }

/** value 的 isType/ getter / setter 判断 */
func (zv *Zval) IsType(value ZendUchar) bool { return zv.typ == value }
func (zv *Zval) IsUndef() bool               { return zv.typ == IS_UNDEF }
func (zv *Zval) IsNotUndef() bool            { return zv.typ != IS_UNDEF }
func (zv *Zval) IsNull() bool                { return zv.typ == IS_NULL }
func (zv *Zval) IsFalse() bool               { return zv.typ == IS_FALSE }
func (zv *Zval) IsTrue() bool                { return zv.typ == IS_TRUE }
func (zv *Zval) IsLong() bool                { return zv.typ == IS_LONG }
func (zv *Zval) IsDouble() bool              { return zv.typ == IS_DOUBLE }
func (zv *Zval) IsString() bool              { return zv.typ == IS_STRING }
func (zv *Zval) IsArray() bool               { return zv.typ == IS_ARRAY }
func (zv *Zval) IsObject() bool              { return zv.typ == IS_OBJECT }
func (zv *Zval) IsResource() bool            { return zv.typ == IS_RESOURCE }
func (zv *Zval) IsReference() bool           { return zv.typ == IS_REFERENCE }
func (zv *Zval) IsConstantAst() bool         { return zv.typ == IS_CONSTANT_AST }
func (zv *Zval) IsIndirect() bool            { return zv.typ == IS_INDIRECT }
func (zv *Zval) IsError() bool               { return zv.typ == IS_ERROR }

func (zv *Zval) Long() int                   { return zv.value.(int) }
func (zv *Zval) Double() float64             { return zv.value.(float64) }
func (zv *Zval) String() *String             { return zv.value.(*String) }
func (zv *Zval) StringVal() string           { return zv.value.(*String).GetStr() }
func (zv *Zval) Array() *Array               { return zv.value.(*Array) }
func (zv *Zval) Object() *ZendObject         { return zv.value.(*ZendObject) }
func (zv *Zval) Resource() *ZendResource     { return zv.value.(*ZendResource) }
func (zv *Zval) Reference() *ZendReference   { return zv.value.(*ZendReference) }
func (zv *Zval) RefCounted() *ZendRefcounted { return zv.value.(*ZendRefcounted) }
func (zv *Zval) ConstantAst() *ZendAstRef    { return zv.value.(*ZendAstRef) }
func (zv *Zval) Indirect() *Zval             { return zv.value.(*Zval) }
func (zv *Zval) Ptr() any                    { return zv.value }
func (zv *Zval) Class() *ClassEntry          { return zv.value.(*ClassEntry) }
func (zv *Zval) Func() IFunction             { return zv.value.(IFunction) }

func (zv *Zval) SetUndef()                       { zv.typ, zv.value = IS_UNDEF, nil }
func (zv *Zval) SetNull()                        { zv.typ, zv.value = IS_NULL, nil }
func (zv *Zval) SetFalse()                       { zv.typ, zv.value = IS_FALSE, nil }
func (zv *Zval) SetTrue()                        { zv.typ, zv.value = IS_TRUE, nil }
func (zv *Zval) SetBool(v bool)                  { zv.typ, zv.value = b.Cond(v, IS_TRUE, IS_FALSE), nil }
func (zv *Zval) SetLong(l int)                   { zv.typ, zv.value = IS_LONG, l }
func (zv *Zval) SetDouble(d float64)             { zv.typ, zv.value = IS_DOUBLE, d }
func (zv *Zval) SetStringVal(s string)           { zv.typ, zv.value = IS_STRING, NewString(s) }
func (zv *Zval) SetString(s *String)             { zv.typ, zv.value = IS_STRING, s }
func (zv *Zval) SetStringCopy(s *String)         { zv.typ, zv.value = IS_STRING, NewString(s.GetStr()) }
func (zv *Zval) SetEmptyArray()                  { zv.typ, zv.value = IS_ARRAY, NewArray(0) }
func (zv *Zval) SetArray(arr *Array)             { zv.typ, zv.value = IS_ARRAY, arr }
func (zv *Zval) SetImmutableArray(arr *Array)    { zv.typ, zv.value = _IS_IMMUTABLE_ARRAY, arr }
func (zv *Zval) SetArrayOfInt(arr []int)         { zv.SetArray(NewArrayOfInt(arr)) }
func (zv *Zval) SetArrayOfString(arr []string)   { zv.SetArray(NewArrayOfString(arr)) }
func (zv *Zval) SetArrayOfZval(arr []*Zval)      { zv.SetArray(NewArrayOfZval(arr)) }
func (zv *Zval) SetObject(obj *ZendObject)       { zv.typ, zv.value = IS_OBJECT, obj }
func (zv *Zval) SetResource(res *ZendResource)   { zv.typ, zv.value = IS_RESOURCE, res }
func (zv *Zval) SetReference(ref *ZendReference) { zv.typ, zv.value = IS_REFERENCE, ref }
func (zv *Zval) SetNewEmptyRef()                 { zv.SetReference(NewZendReference(nil)) }
func (zv *Zval) SetNewRef(val *Zval)             { zv.SetReference(NewZendReference(val)) }
func (zv *Zval) SetConstantAst(ast *ZendAstRef)  { zv.typ, zv.value = IS_CONSTANT_AST, ast }
func (zv *Zval) SetIndirect(v *Zval)             { zv.typ, zv.value = IS_INDIRECT, v }
func (zv *Zval) SetPtr(ptr any)                  { zv.typ, zv.value = IS_PTR, ptr }
func (zv *Zval) SetCe(value *ClassEntry)         { zv.typ, zv.value = IS_PTR, value }
func (zv *Zval) SetFunc(value IFunction)         { zv.typ, zv.value = IS_PTR, value }

/** Zval.u1 -> type & typeFlags */
func (zv *Zval) GetType() ZvalType {
	if zv.typ == _IS_IMMUTABLE_ARRAY {
		return IS_ARRAY
	}
	return zv.typ
}
func (zv *Zval) GetTypeFlags() ZendUchar      { return zv.typeFlags }
func (zv *Zval) SetTypeFlags(value ZendUchar) { zv.typeFlags = value } /* todo remove */
func (zv *Zval) GetTypeInfo() uint32 {
	return uint32(zv.typ) | uint32(zv.typeFlags)<<Z_TYPE_FLAGS_SHIFT
}
func (zv *Zval) SetTypeInfo(value uint32) {
	zv.typ = uint8(value & Z_TYPE_MASK)
	zv.typeFlags = uint8((value & Z_TYPE_FLAGS_MASK) >> Z_TYPE_FLAGS_SHIFT)
}

// 所有对类型的设置操作都集中到这里
func (zv *Zval) SetTypeUndef()  { zv.typ, zv.typeFlags = IS_UNDEF, 0 }
func (zv *Zval) SetTypeNull()   { zv.typ, zv.typeFlags = IS_NULL, 0 }
func (zv *Zval) SetTypeFalse()  { zv.typ, zv.typeFlags = IS_FALSE, 0 }
func (zv *Zval) SetTypeTrue()   { zv.typ, zv.typeFlags = IS_TRUE, 0 }
func (zv *Zval) SetTypeLong()   { zv.typ, zv.typeFlags = IS_LONG, 0 }
func (zv *Zval) SetTypeDouble() { zv.typ, zv.typeFlags = IS_DOUBLE, 0 }
func (zv *Zval) SetTypeString() { zv.typ, zv.typeFlags = IS_STRING, 0 }
func (zv *Zval) SetTypeArray() {
	zv.typ, zv.typeFlags = IS_ARRAY, IS_TYPE_REFCOUNTED|IS_TYPE_COLLECTABLE
}
func (zv *Zval) SetTypeImmutableArray() {
	zv.typ, zv.typeFlags = IS_ARRAY, 0
}

func (zv *Zval) SetTypeObject() {
	zv.typ, zv.typeFlags = IS_OBJECT, IS_TYPE_REFCOUNTED|IS_TYPE_COLLECTABLE
}
func (zv *Zval) SetTypeResource()  { zv.typ, zv.typeFlags = IS_RESOURCE, IS_TYPE_REFCOUNTED }
func (zv *Zval) SetTypeReference() { zv.typ, zv.typeFlags = IS_REFERENCE, IS_TYPE_REFCOUNTED }
func (zv *Zval) SetTypeConstant()  { zv.typ, zv.typeFlags = IS_CONSTANT_AST, IS_TYPE_REFCOUNTED }

/** Zval.u2 */
func (zv *Zval) GetCacheSlot() uint32          { return zv.u2 }
func (zv *Zval) SetCacheSlot(value uint32)     { zv.u2 = value }
func (zv *Zval) GetOplineNum() uint32          { return zv.u2 }
func (zv *Zval) SetOplineNum(value uint32)     { zv.u2 = value }
func (zv *Zval) GetLineno() uint32             { return zv.u2 }
func (zv *Zval) SetLineno(value uint32)        { zv.u2 = value }
func (zv *Zval) GetNumArgs() uint32            { return zv.u2 }
func (zv *Zval) SetNumArgs(value uint32)       { zv.u2 = value }
func (zv *Zval) GetFePos() uint32              { return zv.u2 }
func (zv *Zval) SetFePos(value uint32)         { zv.u2 = value }
func (zv *Zval) GetFeIterIdx() uint32          { return zv.u2 }
func (zv *Zval) SetFeIterIdx(value uint32)     { zv.u2 = value }
func (zv *Zval) GetAccessFlags() uint32        { return zv.u2 }
func (zv *Zval) SetAccessFlags(value uint32)   { zv.u2 = value }
func (zv *Zval) GetPropertyGuard() uint32      { return zv.u2 }
func (zv *Zval) SetPropertyGuard(value uint32) { zv.u2 = value }
func (zv *Zval) GetU2Extra() uint32            { return zv.u2 }
func (zv *Zval) SetU2Extra(value uint32)       { zv.u2 = value }

/* Zval.u2.access_flags */
func (zv *Zval) AddAccessFlags(value uint32)      { zv.u2 |= value }
func (zv *Zval) SubAccessFlags(value uint32)      { zv.u2 &^= value }
func (zv *Zval) HasAccessFlags(value uint32) bool { return zv.u2&value != 0 }
func (zv *Zval) SwitchAccessFlags(value uint32, cond bool) {
	if cond {
		zv.AddAccessFlags(value)
	} else {
		zv.SubAccessFlags(value)
	}
}

/* Zval.u2.constant_flags */
func (zv *Zval) AddConstantFlags(value uint32)      { zv.u2 |= value }
func (zv *Zval) SubConstantFlags(value uint32)      { zv.u2 &^= value }
func (zv *Zval) HasConstantFlags(value uint32) bool { return zv.u2&value != 0 }
func (zv *Zval) SwitchConstantFlags(value uint32, cond bool) {
	if cond {
		zv.AddConstantFlags(value)
	} else {
		zv.SubConstantFlags(value)
	}
}

/**
 * init
 */
func (zv *Zval) SetBy(val *Zval) {
	ZVAL_COPY_VALUE(zv, val)
}
