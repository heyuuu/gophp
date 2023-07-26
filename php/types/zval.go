package types

import (
	b "github.com/heyuuu/gophp/builtin"
)

/** value 的 isType/ getter / setter 判断 */
func (zv *Zval) IsType(value ZvalType) bool { return zv.typ == value }
func (zv *Zval) IsUndef() bool              { return zv.typ == IsUndef }
func (zv *Zval) IsNotUndef() bool           { return zv.typ != IsUndef }
func (zv *Zval) IsNull() bool               { return zv.typ == IsNull }
func (zv *Zval) IsFalse() bool              { return zv.typ == IsFalse }
func (zv *Zval) IsTrue() bool               { return zv.typ == IsTrue }
func (zv *Zval) IsBool() bool               { return zv.typ == IsFalse || zv.typ == IsTrue }
func (zv *Zval) IsLong() bool               { return zv.typ == IsLong }
func (zv *Zval) IsDouble() bool             { return zv.typ == IsDouble }
func (zv *Zval) IsString() bool             { return zv.typ == IsString }
func (zv *Zval) IsArray() bool              { return zv.typ == IsArray }
func (zv *Zval) IsObject() bool             { return zv.typ == IsObject }
func (zv *Zval) IsResource() bool           { return zv.typ == IsResource }
func (zv *Zval) IsReference() bool          { return zv.typ == IsRef }
func (zv *Zval) IsConstantAst() bool        { return zv.typ == IsConstantAst }
func (zv *Zval) IsIndirect() bool           { return zv.typ == IsIndirect }
func (zv *Zval) IsError() bool              { return zv.typ == IsError }

// 返回是否为 undef、null、false，用于快速类型判断
func (zv *Zval) IsSignFalse() bool { return zv.typ <= IsFalse }

// 返回是否为 undef、null、false 或 true，用于快速类型判断
func (zv *Zval) IsSignType() bool { return zv.typ <= IsTrue }

func (zv *Zval) Long() int                 { return zv.value.(int) }
func (zv *Zval) Double() float64           { return zv.value.(float64) }
func (zv *Zval) String() *String           { return zv.value.(*String) }
func (zv *Zval) StringVal() string         { return zv.value.(*String).GetStr() }
func (zv *Zval) Array() *Array             { return zv.value.(*Array) }
func (zv *Zval) Object() *Object           { return zv.value.(*Object) }
func (zv *Zval) Resource() *Resource       { return zv.value.(*Resource) }
func (zv *Zval) Reference() *ZendReference { return zv.value.(*ZendReference) }
func (zv *Zval) ConstantAst() *ZendAstRef  { return zv.value.(*ZendAstRef) }
func (zv *Zval) Indirect() *Zval           { return zv.value.(*Zval) }
func (zv *Zval) Ptr() any                  { return zv.value }
func (zv *Zval) Class() *ClassEntry        { return zv.value.(*ClassEntry) }
func (zv *Zval) Func() IFunction           { return zv.value.(IFunction) }

// fast property
func (zv *Zval) ResourceHandle() int { return zv.Resource().GetHandle() }
func (zv *Zval) ResourceType() int   { return zv.Resource().GetType() }

/** Zval.u1 -> type & typeFlags */
func (zv *Zval) GetType() ZvalType { return zv.typ }
func (zv *Zval) SetType(typ ZvalType) {
	b.Assert(typ <= IsTrue)
	zv.typ, zv.value = typ, nil
}

/** Zval.u2 */
func (zv *Zval) GetCacheSlot() uint32      { return zv.u2 }
func (zv *Zval) SetCacheSlot(value uint32) { zv.u2 = value }
func (zv *Zval) GetOplineNum() uint32      { return zv.u2 }
func (zv *Zval) SetOplineNum(value uint32) { zv.u2 = value }
func (zv *Zval) GetFePos() uint32          { return zv.u2 }
func (zv *Zval) SetFePos(value uint32)     { zv.u2 = value }
func (zv *Zval) GetFeIterIdx() uint32      { return zv.u2 }
func (zv *Zval) SetFeIterIdx(value uint32) { zv.u2 = value }
func (zv *Zval) GetU2Extra() uint32        { return zv.u2 }
func (zv *Zval) SetU2Extra(value uint32)   { zv.u2 = value }
