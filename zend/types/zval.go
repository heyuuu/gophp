// <<generate>>

package types

import "sik/zend"

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
	func_   *zend.ZendFunction
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
func (this *ZendValue) GetFunc() *zend.ZendFunction { return this.val.(*zend.ZendFunction) }

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
func (this *ZendValue) SetFunc(value *zend.ZendFunction) { this.val = value }

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
	//	num_args       uint32 /* arguments number for EX(This) 函数调用时传入参数的个数 */
	//	fe_pos         uint32 /* foreach position  遍历数组时的当前位置*/
	//	fe_iter_idx    uint32 /* foreach iterator index */
	//	access_flags   uint32 /* class constant access flags 对象类的访问标志*/
	//	property_guard uint32 /* single property guard  防止类中魔术方法的循环调用*/
	//	constant_flags uint32
	//	extra          uint32
	//}
}

/** Zval.value */
func (this *Zval) GetValue() ZendValue      { return this.value }
func (this *Zval) SetValue(value ZendValue) { this.value = value }
func (this *Zval) GetStrVal() string        { return this.value.GetStr().GetStr() }

// Zval 从 ZendValue 继承的方法。暂时不写成组合模式，方便开发中查找对应调用调用代码。
func (this *Zval) GetLval() int                { return this.value.GetLval() }
func (this *Zval) GetDval() float64            { return this.value.GetDval() }
func (this *Zval) GetCounted() *ZendRefcounted { return this.value.GetCounted() }
func (this *Zval) GetStr() *String             { return this.value.GetStr() }
func (this *Zval) GetArr() *Array              { return this.value.GetArr() }
func (this *Zval) GetObj() *ZendObject         { return this.value.GetObj() }
func (this *Zval) GetRes() *ZendResource       { return this.value.GetRes() }
func (this *Zval) GetRef() *ZendReference      { return this.value.GetRef() }
func (this *Zval) GetAst() *ZendAstRef         { return this.value.GetAst() }
func (this *Zval) GetZv() *Zval                { return this.value.GetZv() }
func (this *Zval) GetPtr() any                 { return this.value.GetPtr() }
func (this *Zval) GetCe() *ClassEntry          { return this.value.GetCe() }
func (this *Zval) GetFunc() *zend.ZendFunction { return this.value.GetFunc() }

func (this *Zval) SetLval(value int)                { this.SetLval(value) }
func (this *Zval) SetDval(value float64)            { this.SetDval(value) }
func (this *Zval) SetCounted(value *ZendRefcounted) { this.SetCounted(value) }
func (this *Zval) SetStr(value *String)             { this.SetStr(value) }
func (this *Zval) SetArr(value *Array)              { this.SetArr(value) }
func (this *Zval) SetObj(value *ZendObject)         { this.SetObj(value) }
func (this *Zval) SetRes(value *ZendResource)       { this.SetRes(value) }
func (this *Zval) SetRef(value *ZendReference)      { this.SetRef(value) }
func (this *Zval) SetAst(value *ZendAstRef)         { this.SetAst(value) }
func (this *Zval) SetZv(value *Zval)                { this.SetZv(value) }
func (this *Zval) SetPtr(value any)                 { this.SetPtr(value) }
func (this *Zval) SetCe(value *ClassEntry)          { this.SetCe(value) }
func (this *Zval) SetFunc(value *zend.ZendFunction) { this.SetFunc(value) }

/** Zval.u1 -> type & typeFlags */
func (this *Zval) GetType() ZendUchar           { return this.typ }
func (this *Zval) SetType(value ZendUchar)      { this.typ = value }
func (this *Zval) GetTypeFlags() ZendUchar      { return this.typeFlags }
func (this *Zval) SetTypeFlags(value ZendUchar) { this.typeFlags = value }
func (this *Zval) GetTypeInfo() uint32 {
	return uint32(this.typ) | uint32(this.typeFlags)<<Z_TYPE_FLAGS_SHIFT
}
func (this *Zval) SetTypeInfo(value uint32) {
	this.typ = uint8(value & Z_TYPE_MASK)
	this.typeFlags = uint8((value & Z_TYPE_FLAGS_MASK) >> Z_TYPE_FLAGS_SHIFT)
}

func (this *Zval) IsType(value ZendUchar) bool { return this.typ == value }
func (this *Zval) IsUndef() bool               { return this.IsType(IS_UNDEF) }
func (this *Zval) IsNull() bool                { return this.IsType(IS_NULL) }
func (this *Zval) IsFalse() bool               { return this.IsType(IS_FALSE) }
func (this *Zval) IsTrue() bool                { return this.IsType(IS_TRUE) }
func (this *Zval) IsLong() bool                { return this.IsType(IS_LONG) }
func (this *Zval) IsDouble() bool              { return this.IsType(IS_DOUBLE) }
func (this *Zval) IsString() bool              { return this.IsType(IS_STRING) }
func (this *Zval) IsArray() bool               { return this.IsType(IS_ARRAY) }
func (this *Zval) IsObject() bool              { return this.IsType(IS_OBJECT) }
func (this *Zval) IsResource() bool            { return this.IsType(IS_RESOURCE) }
func (this *Zval) IsReference() bool           { return this.IsType(IS_REFERENCE) }
func (this *Zval) IsConstant() bool            { return this.IsType(IS_CONSTANT_AST) }
func (this *Zval) IsIndirect() bool            { return this.IsType(IS_INDIRECT) }
func (this *Zval) IsError() bool               { return this.IsType(IS_ERROR) }

/** Zval.u2 */
func (this *Zval) GetNext() uint32               { return this.u2 }
func (this *Zval) SetNext(value uint32)          { this.u2 = value }
func (this *Zval) GetCacheSlot() uint32          { return this.u2 }
func (this *Zval) SetCacheSlot(value uint32)     { this.u2 = value }
func (this *Zval) GetOplineNum() uint32          { return this.u2 }
func (this *Zval) SetOplineNum(value uint32)     { this.u2 = value }
func (this *Zval) GetLineno() uint32             { return this.u2 }
func (this *Zval) SetLineno(value uint32)        { this.u2 = value }
func (this *Zval) GetNumArgs() uint32            { return this.u2 }
func (this *Zval) SetNumArgs(value uint32)       { this.u2 = value }
func (this *Zval) GetFePos() uint32              { return this.u2 }
func (this *Zval) SetFePos(value uint32)         { this.u2 = value }
func (this *Zval) GetFeIterIdx() uint32          { return this.u2 }
func (this *Zval) SetFeIterIdx(value uint32)     { this.u2 = value }
func (this *Zval) GetAccessFlags() uint32        { return this.u2 }
func (this *Zval) SetAccessFlags(value uint32)   { this.u2 = value }
func (this *Zval) GetPropertyGuard() uint32      { return this.u2 }
func (this *Zval) SetPropertyGuard(value uint32) { this.u2 = value }
func (this *Zval) GetConstantFlags() uint32      { return this.u2 }
func (this *Zval) SetConstantFlags(value uint32) { this.u2 = value }
func (this *Zval) GetU2Extra() uint32            { return this.u2 }
func (this *Zval) SetU2Extra(value uint32)       { this.u2 = value }

/* Zval.u2.access_flags */
func (this *Zval) AddAccessFlags(value uint32)      { this.u2 |= value }
func (this *Zval) SubAccessFlags(value uint32)      { this.u2 &^= value }
func (this *Zval) HasAccessFlags(value uint32) bool { return this.u2&value != 0 }
func (this *Zval) SwitchAccessFlags(value uint32, cond bool) {
	if cond {
		this.AddAccessFlags(value)
	} else {
		this.SubAccessFlags(value)
	}
}

/* Zval.u2.constant_flags */
func (this *Zval) AddConstantFlags(value uint32)      { this.u2 |= value }
func (this *Zval) SubConstantFlags(value uint32)      { this.u2 &^= value }
func (this *Zval) HasConstantFlags(value uint32) bool { return this.u2&value != 0 }
func (this *Zval) SwitchConstantFlags(value uint32, cond bool) {
	if cond {
		this.AddConstantFlags(value)
	} else {
		this.SubConstantFlags(value)
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
func NewZvalString(str string) *Zval          { var tmp Zval; tmp.SetRawString(str); return &tmp }
func NewZvalLong(l int) *Zval                 { var tmp Zval; tmp.SetLong(l); return &tmp }
func NewZvalDouble(d float64) *Zval           { var tmp Zval; tmp.SetDouble(d); return &tmp }
func NewZvalResource(res *ZendResource) *Zval { var tmp Zval; tmp.SetResource(res); return &tmp }
func NewZvalPtr(ptr any) *Zval                { var tmp Zval; tmp.SetAsPtr(ptr); return &tmp }

/**
 * init
 */
func (this *Zval) SetUndef() { this.SetTypeInfo(IS_UNDEF) }
func (this *Zval) SetNull()  { this.SetTypeInfo(IS_NULL) }
func (this *Zval) SetFalse() { this.SetTypeInfo(IS_FALSE) }
func (this *Zval) SetTrue()  { this.SetTypeInfo(IS_TRUE) }
func (this *Zval) SetBool(b bool) {
	if b {
		this.SetTypeInfo(IS_TRUE)
	} else {
		this.SetTypeInfo(IS_FALSE)
	}
}

func (this *Zval) SetLong(l int) {
	this.SetTypeInfo(IS_LONG)
	this.SetLval(l)
}
func (this *Zval) SetDouble(d float64) {
	this.SetTypeInfo(IS_DOUBLE)
	this.SetDval(d)
}

func (this *Zval) SetRawString(s string) {
	this.SetString(NewString(s))
}

func (this *Zval) SetString(s *String) {
	this.SetStr(s)
	this.SetTypeInfo(IS_STRING_EX)
}
func (this *Zval) SetStringVal(str string) {
	this.SetString(NewString(str))
}
func (this *Zval) SetInternedString(s *String) {
	this.SetStr(s)
	this.SetTypeInfo(IS_INTERNED_STRING_EX)
}
func (this *Zval) SetStringCopy(s *String) {
	s.AddRefcount()
	this.SetString(s)
}

func (this *Zval) SetArray(arr *Array) {
	this.SetArr(arr)
	this.SetTypeInfo(IS_ARRAY_EX)
}

func (this *Zval) SetObject(obj *ZendObject) {
	this.SetObj(obj)
	this.SetTypeInfo(IS_OBJECT_EX)
}

func (this *Zval) SetResource(res *ZendResource) {
	this.SetRes(res)
	this.SetTypeInfo(IS_RESOURCE_EX)
}
func (this *Zval) SetNewResource(handle int, ptr any, type_ int) {
	var res = NewZendResource(handle, ptr, type_)
	this.SetResource(res)
}
func (this *Zval) SetNewResourcePersistent(handle int, ptr any, type_ int) {
	var res = NewZendResourcePersistent(handle, ptr, type_, true)
	this.SetResource(res)
}

func (this *Zval) SetReference(ref *ZendReference) {
	this.SetRef(ref)
	this.SetTypeInfo(IS_REFERENCE_EX)
}
func (this *Zval) SetNewEmptyRef() {
	var ref *ZendReference = NewZendReference(nil)
	this.SetReference(ref)
}
func (this *Zval) SetNewRef(val *Zval) {
	var ref *ZendReference = NewZendReference(val)
	this.SetReference(ref)
}

func (this *Zval) SetConstantAst(ast *ZendAstRef) {
	this.SetAst(ast)
	this.SetTypeInfo(IS_CONSTANT_AST_EX)
}

func (this *Zval) SetIndirect(v *Zval) {
	this.SetZv(v)
	this.SetTypeInfo(IS_INDIRECT)
}

func (this *Zval) SetAsPtr(ptr any) {
	this.SetPtr(ptr)
	this.SetTypeInfo(IS_PTR)
}

func (this *Zval) SetAliasPtr(ptr any) {
	this.SetPtr(ptr)
	this.SetTypeInfo(IS_ALIAS_PTR)
}
