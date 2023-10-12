package types

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
	v any

	u2 uint32
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

// Zval new
func NewZvalUndef() *Zval           { return &Zval{v: nil} }
func NewZvalNull() *Zval            { return &Zval{v: IsNull} }
func NewZvalFalse() *Zval           { return &Zval{v: false} }
func NewZvalTrue() *Zval            { return &Zval{v: true} }
func NewZvalBool(b bool) *Zval      { return &Zval{v: b} }
func NewZvalLong(l int) *Zval       { return &Zval{v: l} }
func NewZvalDouble(d float64) *Zval { return &Zval{v: d} }
func NewZvalString(s string) *Zval  { return &Zval{v: s} }
func NewZvalArray(arr *Array) *Zval {
	if arr == nil {
		arr = NewArray()
	}
	return &Zval{v: arr}
}
func NewZvalObject(obj *Object) *Zval     { return &Zval{v: obj} }
func NewZvalResource(res *Resource) *Zval { return &Zval{v: res} }
func NewZvalPtr(ptr any) *Zval            { return &Zval{v: Ptr(ptr)} }
func NewZvalIndirect(z *Zval) *Zval       { return &Zval{v: Indirect(z)} }

// Zval setter
func (zv *Zval) SetUndef()                 { zv.v = nil }
func (zv *Zval) SetNull()                  { zv.v = IsNull }
func (zv *Zval) SetFalse()                 { zv.v = false }
func (zv *Zval) SetTrue()                  { zv.v = true }
func (zv *Zval) SetBool(b bool)            { zv.v = b }
func (zv *Zval) SetLong(l int)             { zv.v = l }
func (zv *Zval) SetDouble(d float64)       { zv.v = d }
func (zv *Zval) SetString(s string)        { zv.v = s }
func (zv *Zval) SetArray(arr *Array)       { zv.v = arr }
func (zv *Zval) SetObject(obj *Object)     { zv.v = obj }
func (zv *Zval) SetResource(res *Resource) { zv.v = res }

func (zv *Zval) SetStringEx(s *String)          { zv.v = s.GetStr() }
func (zv *Zval) SetEmptyArray()                 { zv.v = NewArray() }
func (zv *Zval) SetArrayOfInt(arr []int)        { zv.SetArray(NewArrayOfInt(arr)) }
func (zv *Zval) SetArrayOfString(arr []string)  { zv.SetArray(NewArrayOfString(arr)) }
func (zv *Zval) SetArrayOfZval(arr []*Zval)     { zv.SetArray(NewArrayOfZval(arr)) }
func (zv *Zval) SetReference(ref *Reference)    { zv.v = ref }
func (zv *Zval) SetNewEmptyRef()                { zv.SetReference(NewZendReference(nil)) }
func (zv *Zval) SetNewRef(val *Zval)            { zv.SetReference(NewZendReference(val)) }
func (zv *Zval) SetConstantAst(ast *ZendAstRef) { zv.v = ast }
func (zv *Zval) SetIndirect(v *Zval)            { zv.v = Indirect(v) }
func (zv *Zval) SetPtr(ptr any)                 { zv.v = Ptr(ptr) }
func (zv *Zval) SetCe(value *ClassEntry)        { zv.SetPtr(value) }
func (zv *Zval) SetFunc(value IFunction)        { zv.SetPtr(value) }
func (zv *Zval) SetIsError()                    { zv.v = IsError }
func (zv *Zval) SetBy(val *Zval)                { zv.CopyValueFrom(val) }

// Zval getter
func (zv *Zval) Type() ZvalType {
	switch v := zv.v.(type) {
	case nil:
		return IsUndef
	case ZvalType: // IsNull | IsError
		return v
	case bool:
		if v {
			return IsTrue
		} else {
			return IsFalse
		}
	case int:
		return IsLong
	case float64:
		return IsDouble
	case string:
		return IsString
	case *Array:
		return IsArray
	case *Object:
		return IsObject
	case *Resource:
		return IsResource
	case *Ref:
		return IsRef
	case *ZendAstRef:
		return IsConstantAst
	case Indirect:
		return IsIndirect
	case Ptr:
		return IsPtr
	default:
		panic("unreachable")
	}
}

func (zv *Zval) IsType(typ ZvalType) bool { return zv.Type() == typ }
func (zv *Zval) IsUndef() bool            { return zv.v == nil }
func (zv *Zval) IsNotUndef() bool         { return zv.v != nil }
func (zv *Zval) IsNull() bool             { return zv.v == IsNull }
func (zv *Zval) IsFalse() bool            { return zv.v == false }
func (zv *Zval) IsTrue() bool             { return zv.v == true }
func (zv *Zval) IsBool() bool             { _, ok := zv.v.(bool); return ok }
func (zv *Zval) IsLong() bool             { _, ok := zv.v.(int); return ok }
func (zv *Zval) IsDouble() bool           { _, ok := zv.v.(float64); return ok }
func (zv *Zval) IsString() bool           { _, ok := zv.v.(string); return ok }
func (zv *Zval) IsArray() bool            { _, ok := zv.v.(*Array); return ok }
func (zv *Zval) IsObject() bool           { _, ok := zv.v.(*Object); return ok }
func (zv *Zval) IsResource() bool         { _, ok := zv.v.(*Resource); return ok }
func (zv *Zval) IsRef() bool              { _, ok := zv.v.(*Ref); return ok }
func (zv *Zval) IsConstantAst() bool      { _, ok := zv.v.(*ZendAstRef); return ok }
func (zv *Zval) IsIndirect() bool         { _, ok := zv.v.(Indirect); return ok }
func (zv *Zval) IsError() bool            { return zv.v == IsError }

// 返回是否为 undef、null、false，用于快速类型判断
func (zv *Zval) IsSignFalse() bool { return zv.Type() <= IsFalse }

// 返回是否为 undef、null、false 或 true，用于快速类型判断
func (zv *Zval) IsSignType() bool { return zv.Type() <= IsTrue }

func zvalValue[T any](zv *Zval) T {
	if v, ok := zv.v.(T); ok {
		return v
	}
	panic("Get Zval value by a mismatched type")
}
func (zv *Zval) Bool() bool          { return zvalValue[bool](zv) }
func (zv *Zval) Long() int           { return zvalValue[int](zv) }
func (zv *Zval) Double() float64     { return zvalValue[float64](zv) }
func (zv *Zval) String() string      { return zvalValue[string](zv) }
func (zv *Zval) StringEx() *String   { return NewString(zv.String()) }
func (zv *Zval) Array() *Array       { return zvalValue[*Array](zv) }
func (zv *Zval) Object() *Object     { return zvalValue[*Object](zv) }
func (zv *Zval) Resource() *Resource { return zvalValue[*Resource](zv) }
func (zv *Zval) Ref() *Ref           { return zvalValue[*Ref](zv) }
func (zv *Zval) DeRef() *Zval {
	if ref, ok := zv.v.(*Ref); ok {
		return ref.GetVal()
	}
	return zv
}
func (zv *Zval) SafeDeRef() *Zval {
	if zv == nil {
		return nil
	}

	if ref, ok := zv.v.(*Ref); ok {
		return ref.GetVal()
	}
	return zv
}

func (zv *Zval) ConstantAst() *ZendAstRef { return zvalValue[*ZendAstRef](zv) }
func (zv *Zval) Indirect() *Zval          { return zvalValue[Indirect](zv) }
func (zv *Zval) Ptr() any                 { return zvalValue[Ptr](zv) }
func (zv *Zval) Class() *ClassEntry       { return zvalValue[*ClassEntry](zv) }
func (zv *Zval) Func() IFunction          { return zvalValue[IFunction](zv) }

// fast property
func (zv *Zval) ResourceHandle() int { return zv.Resource().GetHandle() }
func (zv *Zval) ResourceType() int   { return zv.Resource().GetType() }

/* wrapper zval value */
type Indirect *Zval
type Ptr any
