package types

import (
	b "github.com/heyuuu/gophp/builtin"
)

/**
 * ZendValue
 */
type ZendValue struct {
	val any
}
type RawZendValue struct /* union */ {
	lval    int
	dval    float64
	counted *ZendRefcounted
	str     *String
	arr     *Array
	obj     *ZendObject
	res     *ZendResource
	ref     *ZendReference
	ast     *ZendAstRef
	zv      *Zval
	ptr     any
	ce      *ClassEntry
	func_   IFunction
}

func (this *ZendValue) GetLval() int                { return this.val.(int) }
func (this *ZendValue) GetDval() float64            { return this.val.(float64) }
func (this *ZendValue) GetCounted() *ZendRefcounted { return this.val.(*ZendRefcounted) }
func (this *ZendValue) GetStr() *String             { return this.val.(*String) }
func (this *ZendValue) GetArr() *Array              { return this.val.(*Array) }
func (this *ZendValue) GetObj() *ZendObject         { return this.val.(*ZendObject) }
func (this *ZendValue) GetRes() *ZendResource       { return this.val.(*ZendResource) }
func (this *ZendValue) GetRef() *ZendReference      { return this.val.(*ZendReference) }
func (this *ZendValue) GetAst() *ZendAstRef         { return this.val.(*ZendAstRef) }
func (this *ZendValue) GetZv() *Zval                { return this.val.(*Zval) }
func (this *ZendValue) GetPtr() any                 { return this.val }
func (this *ZendValue) GetCe() *ClassEntry          { return this.val.(*ClassEntry) }
func (this *ZendValue) GetFunc() IFunction          { return this.val.(IFunction) }

func (this *ZendValue) SetLval(value int)                { this.val = value }
func (this *ZendValue) SetDval(value float64)            { this.val = value }
func (this *ZendValue) SetCounted(value *ZendRefcounted) { this.val = value }
func (this *ZendValue) SetStr(value *String)             { this.val = value }
func (this *ZendValue) SetArr(value *Array)              { this.val = value }
func (this *ZendValue) SetObj(value *ZendObject)         { this.val = value }
func (this *ZendValue) SetRes(value *ZendResource)       { this.val = value }
func (this *ZendValue) SetRef(value *ZendReference)      { this.val = value }
func (this *ZendValue) SetAst(value *ZendAstRef)         { this.val = value }
func (this *ZendValue) SetZv(value *Zval)                { this.val = value }
func (this *ZendValue) SetPtr(value any)                 { this.val = value }
func (this *ZendValue) SetCe(value *ClassEntry)          { this.val = value }
func (this *ZendValue) SetFunc(value IFunction)          { this.val = value }

/**
 * Zval
 */
type Zval struct {
	value     ZendValue
	typ       ZendUchar
	typeFlags ZendUchar
	u2        uint32

	//u1 struct /* union */ {
	//	v struct {
	//		type_      ZendUchar
	//		type_flags ZendUchar
	//		u          struct /* union */ {
	//			extra uint16
	//		}
	//	}
	//	type_info uint32
	//}
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

/** Zval.value */
func (zv *Zval) GetValue() ZendValue      { return zv.value }
func (zv *Zval) SetValue(value ZendValue) { zv.value = value }
func (zv *Zval) GetStrVal() string        { return zv.value.GetStr().GetStr() }

// Zval 从 ZendValue 继承的方法。暂时不写成组合模式，方便开发中查找对应调用调用代码。
func (zv *Zval) GetLval() int                { return zv.value.GetLval() }
func (zv *Zval) GetDval() float64            { return zv.value.GetDval() }
func (zv *Zval) GetCounted() *ZendRefcounted { return zv.value.GetCounted() }
func (zv *Zval) GetStr() *String             { return zv.value.GetStr() }
func (zv *Zval) GetArr() *Array              { return zv.value.GetArr() }
func (zv *Zval) GetObj() *ZendObject         { return zv.value.GetObj() }
func (zv *Zval) GetRes() *ZendResource       { return zv.value.GetRes() }
func (zv *Zval) GetRef() *ZendReference      { return zv.value.GetRef() }
func (zv *Zval) GetAst() *ZendAstRef         { return zv.value.GetAst() }
func (zv *Zval) GetZv() *Zval                { return zv.value.GetZv() }
func (zv *Zval) GetPtr() any                 { return zv.value.GetPtr() }
func (zv *Zval) GetCe() *ClassEntry          { return zv.value.GetCe() }
func (zv *Zval) GetFunc() IFunction          { return zv.value.GetFunc() }

func (zv *Zval) SetLval(value int)                { zv.value.SetLval(value) }
func (zv *Zval) SetDval(value float64)            { zv.value.SetDval(value) }
func (zv *Zval) SetCounted(value *ZendRefcounted) { zv.value.SetCounted(value) }
func (zv *Zval) SetStr(value *String)             { zv.value.SetStr(value) }
func (zv *Zval) SetArr(value *Array)              { zv.value.SetArr(value) }
func (zv *Zval) SetObj(value *ZendObject)         { zv.value.SetObj(value) }
func (zv *Zval) SetRes(value *ZendResource)       { zv.value.SetRes(value) }
func (zv *Zval) SetRef(value *ZendReference)      { zv.value.SetRef(value) }
func (zv *Zval) SetAst(value *ZendAstRef)         { zv.value.SetAst(value) }
func (zv *Zval) SetZv(value *Zval)                { zv.value.SetZv(value) }
func (zv *Zval) SetPtr(value any)                 { zv.value.SetPtr(value) }
func (zv *Zval) SetCe(value *ClassEntry)          { zv.value.SetCe(value) }
func (zv *Zval) SetFunc(value IFunction)          { zv.value.SetFunc(value) }

/** Zval.u1 -> type & typeFlags */
func (zv *Zval) GetType() ZvalType            { return zv.typ }
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
func (zv *Zval) SetTypeUndef()          { zv.typ, zv.typeFlags = IS_UNDEF, 0 }
func (zv *Zval) SetTypeNull()           { zv.typ, zv.typeFlags = IS_NULL, 0 }
func (zv *Zval) SetTypeFalse()          { zv.typ, zv.typeFlags = IS_FALSE, 0 }
func (zv *Zval) SetTypeTrue()           { zv.typ, zv.typeFlags = IS_TRUE, 0 }
func (zv *Zval) SetTypeLong()           { zv.typ, zv.typeFlags = IS_LONG, 0 }
func (zv *Zval) SetTypeDouble()         { zv.typ, zv.typeFlags = IS_DOUBLE, 0 }
func (zv *Zval) SetTypeString()         { zv.typ, zv.typeFlags = IS_STRING, IS_TYPE_REFCOUNTED }
func (zv *Zval) SetTypeInternedString() { zv.typ, zv.typeFlags = IS_STRING, 0 }
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
func (zv *Zval) SetTypeIndirect()  { zv.typ, zv.typeFlags = IS_INDIRECT, 0 }
func (zv *Zval) SetTypePtr()       { zv.typ, zv.typeFlags = IS_PTR, 0 }
func (zv *Zval) SetTypeAliasPtr()  { zv.typ, zv.typeFlags = IS_ALIAS_PTR, 0 }
func (zv *Zval) SetTypeError()     { zv.typ, zv.typeFlags = IS_ERROR, 0 }

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
func (zv *Zval) IsConstant() bool            { return zv.typ == IS_CONSTANT_AST }
func (zv *Zval) IsIndirect() bool            { return zv.typ == IS_INDIRECT }
func (zv *Zval) IsError() bool               { return zv.typ == IS_ERROR }

/** Zval.u2 */
func (zv *Zval) GetNext() uint32               { return zv.u2 }
func (zv *Zval) SetNext(value uint32)          { zv.u2 = value }
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
func (zv *Zval) GetConstantFlags() uint32      { return zv.u2 }
func (zv *Zval) SetConstantFlags(value uint32) { zv.u2 = value }
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
 * New
 */
func NewZvalUndef() *Zval                     { var tmp Zval; tmp.SetUndef(); return &tmp }
func NewZvalNull() *Zval                      { var tmp Zval; tmp.SetNull(); return &tmp }
func NewZvalFalse() *Zval                     { var tmp Zval; tmp.SetFalse(); return &tmp }
func NewZvalTrue() *Zval                      { var tmp Zval; tmp.SetTrue(); return &tmp }
func NewZvalBool(b bool) *Zval                { var tmp Zval; tmp.SetBool(b); return &tmp }
func NewZvalString(str string) *Zval          { var tmp Zval; tmp.SetStringVal(str); return &tmp }
func NewZvalLong(l int) *Zval                 { var tmp Zval; tmp.SetLong(l); return &tmp }
func NewZvalDouble(d float64) *Zval           { var tmp Zval; tmp.SetDouble(d); return &tmp }
func NewZvalArray(arr *Array) *Zval           { var tmp Zval; tmp.SetArray(arr); return &tmp }
func NewZvalResource(res *ZendResource) *Zval { var tmp Zval; tmp.SetResource(res); return &tmp }
func NewZvalPtr(ptr any) *Zval                { var tmp Zval; tmp.SetAsPtr(ptr); return &tmp }

/**
 * init
 */
func (zv *Zval) SetUndef() { zv.SetTypeUndef() }
func (zv *Zval) SetNull()  { zv.SetTypeNull() }
func (zv *Zval) SetFalse() { zv.SetTypeFalse() }
func (zv *Zval) SetTrue()  { zv.SetTypeTrue() }
func (zv *Zval) SetBool(b bool) {
	if b {
		zv.SetTypeTrue()
	} else {
		zv.SetTypeFalse()
	}
}
func (zv *Zval) SetLong(l int)               { zv.SetTypeLong(); zv.SetLval(l) }
func (zv *Zval) SetDouble(d float64)         { zv.SetTypeDouble(); zv.SetDval(d) }
func (zv *Zval) SetStringVal(s string)       { zv.SetString(NewString(s)) }
func (zv *Zval) SetString(s *String)         { zv.SetTypeString(); zv.SetStr(s) }
func (zv *Zval) SetInternedString(s *String) { zv.SetTypeInternedString(); zv.SetStr(s) }
func (zv *Zval) SetStringCopy(s *String) {
	//s.AddRefcount()
	zv.SetString(s)
}
func (zv *Zval) SetArray(arr *Array) { zv.SetTypeArray(); zv.SetArr(arr) }
func (zv *Zval) SetImmutableArray(arr *Array) {
	b.Assert(arr.IsImmutable())
	zv.SetTypeArray()
	zv.SetArr(arr)
}
func (zv *Zval) SetEmptyArray() {
	zv.SetImmutableArray(emptyArray)
}

func (zv *Zval) SetObject(obj *ZendObject)     { zv.SetTypeObject(); zv.SetObj(obj) }
func (zv *Zval) SetResource(res *ZendResource) { zv.SetTypeResource(); zv.SetRes(res) }
func (zv *Zval) SetNewResource(handle int, ptr any, type_ int) {
	var res = NewZendResource(handle, ptr, type_)
	zv.SetResource(res)
}
func (zv *Zval) SetNewResourcePersistent(handle int, ptr any, type_ int) {
	var res = NewZendResourcePersistent(handle, ptr, type_, true)
	zv.SetResource(res)
}
func (zv *Zval) SetReference(ref *ZendReference) { zv.SetTypeReference(); zv.SetRef(ref) }
func (zv *Zval) SetNewEmptyRef()                 { zv.SetReference(NewZendReference(nil)) }
func (zv *Zval) SetNewRef(val *Zval)             { zv.SetReference(NewZendReference(val)) }
func (zv *Zval) SetConstantAst(ast *ZendAstRef)  { zv.SetTypeConstant(); zv.SetAst(ast) }
func (zv *Zval) SetIndirect(v *Zval)             { zv.SetTypeIndirect(); zv.SetZv(v) }
func (zv *Zval) SetAsPtr(ptr any)                { zv.SetTypePtr(); zv.SetPtr(ptr) }
func (zv *Zval) SetAliasPtr(ptr any)             { zv.SetTypeAliasPtr(); zv.SetPtr(ptr) }

func (zv *Zval) SetBy(val *Zval) {
	ZVAL_COPY_VALUE(zv, val)
}
