package types

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend"
)

func ZEND_TYPE_NAME(t ZendType) *String   { return t.Name() }
func ZEND_TYPE_CE(t ZendType) *ClassEntry { return t.Ce() }

func ZEND_TYPE_ENCODE(code uint32, allow_null int) ZendType {
	if allow_null != 0 {
		return ZendType(code<<2 | 0x1)
	} else {
		return ZendType(code<<2 | 0x0)
	}
}
func ZEND_TYPE_ENCODE_CE(ce *ClassEntry, allow_null bool) ZendType {
	var ptr = b.CastUintptr(ce)
	if allow_null {
		return ZendType(ptr | 0x3)
	} else {
		return ZendType(ptr | 0x2)
	}
}
func ZEND_TYPE_ENCODE_CLASS(class_name *String, allow_null ZendBool) ZendType {
	var ptr = b.CastUintptr(class_name)
	if allow_null != 0 {
		return ZendType(ptr | 0x1)
	} else {
		return ZendType(ptr | 0x0)
	}
}
func ZEND_TYPE_ENCODE_CLASS_CONST(class_name string, allow_null int) ZendType {
	var fullClassName string
	if allow_null != 0 {
		fullClassName = "?" + class_name
	} else {
		fullClassName = class_name
	}
	var ptr = b.CastUintptr(&fullClassName)
	return ZendType(ptr)
}
func ZEND_PROPERTY_INFO_SOURCE_FROM_LIST(list *ZendPropertyInfoList) int { return 0x1 | uintPtr(list) }
func ZEND_PROPERTY_INFO_SOURCE_TO_LIST(list uintPtr) *ZendPropertyInfoList {
	return (*ZendPropertyInfoList)(list & ^0x1)
}
func ZEND_PROPERTY_INFO_SOURCE_IS_LIST(list uintPtr) int { return list & 0x1 }
func ZEND_SAME_FAKE_TYPE(faketype int, realtype uint8) bool {
	return faketype == realtype || faketype == IS_BOOL && (realtype == IS_TRUE || realtype == IS_FALSE)
}

func Z_OBJ_P(zval *Zval) *ZendObject                 { return zval.Object() }
func Z_OBJ_HT(zval Zval) *zend.ZendObjectHandlers    { return zval.Object().GetHandlers() }
func Z_OBJ_HT_P(zval *Zval) *zend.ZendObjectHandlers { return zval.Object().GetHandlers() }
func Z_OBJCE(zval Zval) *ClassEntry                  { return zval.Object().GetCe() }
func Z_OBJCE_P(zval *Zval) *ClassEntry               { return zval.Object().GetCe() }
func Z_OBJPROP(zval Zval) *Array {
	return Z_OBJ_HT(zval).GetGetProperties()(&zval)
}
func Z_OBJPROP_P(zval_p *Zval) *Array       { return Z_OBJPROP(*zval_p) }
func Z_RES(zval Zval) *ZendResource         { return zval.Resource() }
func Z_RES_P(zval_p *Zval) *ZendResource    { return zval_p.Resource() }
func Z_RES_HANDLE(zval Zval) int            { return Z_RES(zval).GetHandle() }
func Z_RES_HANDLE_P(zval_p *Zval) int       { return Z_RES_HANDLE(*zval_p) }
func Z_RES_TYPE(zval Zval) int              { return Z_RES(zval).GetType() }
func Z_RES_TYPE_P(zval_p *Zval) int         { return Z_RES_TYPE(*zval_p) }
func Z_REF_P(zval_p *Zval) *ZendReference   { return zval_p.Reference() }
func Z_REFVAL(zval Zval) *Zval              { return zval.Reference().GetVal() }
func Z_REFVAL_P(zval_p *Zval) *Zval         { return zval_p.Reference().GetVal() }
func GC_AST(p *ZendAstRef) *zend.ZendAst    { return p.GcAst() }
func Z_ASTVAL(zval Zval) *zend.ZendAst      { return GC_AST(zval.ConstantAst()) }
func Z_ASTVAL_P(zval_p *Zval) *zend.ZendAst { return Z_ASTVAL(*zval_p) }
func Z_INDIRECT_P(zval_p *Zval) *Zval       { return zval_p.Indirect() }
func Z_PTR(zval Zval) any                   { return zval.Ptr() }

func ZVAL_NEW_PERSISTENT_ARR(z *Zval) {
	var arr = NewArray(0)
	z.SetArray(arr)
}
func ZVAL_NEW_REF(z *Zval, r *Zval) { z.SetNewRef(r) }
func ZVAL_MAKE_REF_EX(z *Zval, refcount uint32) {
	var ref *ZendReference = NewZendReference(z)
	ref.SetRefcount(refcount)
	z.SetReference(ref)
}
func ZVAL_PTR(z *Zval, p any)      { z.SetPtr(p) }
func Z_REFCOUNT_P(pz *Zval) uint32 { return pz.GetRefcount() }

func ZVAL_COPY_VALUE(z *Zval, v *Zval)  { z.CopyValueFrom(v) }
func ZVAL_COPY(z *Zval, v *Zval)        { z.CopyFrom(v) }
func ZVAL_COPY_OR_DUP(z *Zval, v *Zval) { z.CopyOrDupFrom(v) }
func ZVAL_DEREF(z *Zval) *Zval          { return z.DeRef() }
func ZVAL_DEINDIRECT(z *Zval) *Zval     { return z.DeIndirect() }

func ZVAL_MAKE_REF(zv *Zval) {
	var __zv *Zval = zv
	if !(__zv.IsReference()) {
		ZVAL_NEW_REF(__zv, __zv)
	}
}
func ZVAL_UNREF(z *Zval) {
	var _z *Zval = z
	var ref *ZendReference
	b.Assert(_z.IsReference())
	ref = _z.Reference()
	ZVAL_COPY_VALUE(_z, ref.GetVal())
	zend.EfreeSize(ref, b.SizeOf("zend_reference"))
}
func ZVAL_COPY_DEREF(z *Zval, v *Zval) {
	z.CopyFrom(v.DeRef())
}
func SeparateArray(zv *Zval) {
	b.Assert(zv.IsArray())
	zv.SetArray(ArrayRealDup(zv.Array()))
}
func SeparateZval(zv *Zval) {
	// 解 Ref
	if zv.IsReference() {
		zv.CopyValueFrom(zv.DeRef())
	}
	// 仅数组需要分离
	if zv.IsArray() {
		zv.SetArray(ArrayRealDup(zv.Array()))
	}
}
func SEPARATE_ARG_IF_REF(varptr *Zval) *Zval {
	return varptr.DeRef()
}
func ZVAL_COPY_VALUE_PROP(z *Zval, v *Zval) { *z = *v }
func ZVAL_COPY_PROP(z *Zval, v *Zval) {
	ZVAL_COPY(z, v)
	z.SetU2Extra(v.GetU2Extra())
}
func ZVAL_COPY_OR_DUP_PROP(z *Zval, v *Zval) {
	ZVAL_COPY_OR_DUP(z, v)
	z.SetU2Extra(v.GetU2Extra())
}
