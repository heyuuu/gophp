package types

import (
	b "github.com/heyuuu/gophp/builtin"
)

// 返回是否为 undef、null、false，用于快速类型判断
func (zv *Zval) IsSignFalse() bool { return zv.Type() <= IsFalse }

// 返回是否为 undef、null、false 或 true，用于快速类型判断
func (zv *Zval) IsSignType() bool { return zv.Type() <= IsTrue }

func (zv *Zval) Long() int                { return zv.v.(int) }
func (zv *Zval) Double() float64          { return zv.v.(float64) }
func (zv *Zval) String() *String          { return zv.v.(*String) }
func (zv *Zval) StringVal() string        { return zv.v.(*String).GetStr() }
func (zv *Zval) Array() *Array            { return zv.v.(*Array) }
func (zv *Zval) Object() *Object          { return zv.v.(*Object) }
func (zv *Zval) Resource() *Resource      { return zv.v.(*Resource) }
func (zv *Zval) Reference() *Reference    { return zv.v.(*Reference) }
func (zv *Zval) ConstantAst() *ZendAstRef { return zv.v.(*ZendAstRef) }
func (zv *Zval) Indirect() *Zval          { return zv.v.(*Zval) }
func (zv *Zval) Ptr() any                 { return zv.v }
func (zv *Zval) Class() *ClassEntry       { return zv.v.(*ClassEntry) }
func (zv *Zval) Func() IFunction          { return zv.v.(IFunction) }

// fast property
func (zv *Zval) ResourceHandle() int { return zv.Resource().GetHandle() }
func (zv *Zval) ResourceType() int   { return zv.Resource().GetType() }

/** Zval.u1 -> type & typeFlags */
func (zv *Zval) SetType(typ ZvalType) {
	b.Assert(typ <= IsTrue)
	zv.typ, zv.v = typ, nil
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
