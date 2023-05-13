package types

import (
	b "github.com/heyuuu/gophp/builtin"
)

// ZVAL_COPY_VALUE
func (zv *Zval) CopyValueFrom(v *Zval) {
	// 复制除 u2 外所有数据
	zv.typ, zv.value = v.typ, v.value
}
func (zv *Zval) CopyFrom(v *Zval) {
	zv.CopyValueFrom(v)
	// 除数组外，基础类型都复制了值，引用类型都复制了指针；仅数组需要做写时复制
	if v.IsArray() {
		zv.SetArray(zv.Array().Copy())
	}
}
func (zv *Zval) CopyOrDupFrom(v *Zval) {
	zv.CopyValueFrom(v)
	// 除数组外，基础类型都复制了值，引用类型都复制了指针；仅数组需要做写时复制
	if zv.IsArray() {
		zv.SetArray(ZendArrayDup(zv.Array()))
	}
}

// ZVAL_DEREF(zv)
func (zv *Zval) DeRef() *Zval {
	if zv.IsReference() {
		return zv.Reference().GetVal()
	}
	return zv
}

// ZVAL_DEINDIRECT(zv)
func (zv *Zval) DeIndirect() *Zval {
	if zv.IsIndirect() {
		return zv.Indirect()
	}
	return zv
}

/**
 * GC - Refcount
 */
func (zv *Zval) IsRefcounted() bool {
	switch zv.typ {
	case IS_ARRAY, // 不包含 _IS_IMMUTABLE_ARRAY
		IS_OBJECT,
		IS_RESOURCE,
		IS_REFERENCE:
		return true
	default:
		return false
	}
}

func (zv *Zval) GetRefcount() uint32 {
	b.Assert(zv.IsRefcounted())
	return zv.RefCounted().GetRefcount()
}

/**
 * GC - GC_PROTECTED
 */
func (zv *Zval) IsRecursive() bool   { return zv.RefCounted().IsRecursive() }
func (zv *Zval) ProtectRecursive()   { zv.RefCounted().ProtectRecursive() }
func (zv *Zval) UnprotectRecursive() { zv.RefCounted().UnprotectRecursive() }
