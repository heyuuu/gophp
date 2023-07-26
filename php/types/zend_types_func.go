package types

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend"
)

func ZEND_PROPERTY_INFO_SOURCE_FROM_LIST(list *ZendPropertyInfoList) int { return 0x1 | uintPtr(list) }
func ZEND_PROPERTY_INFO_SOURCE_TO_LIST(list uintPtr) *ZendPropertyInfoList {
	return (*ZendPropertyInfoList)(list & ^0x1)
}
func ZEND_PROPERTY_INFO_SOURCE_IS_LIST(list uintPtr) int { return list & 0x1 }
func ZEND_SAME_FAKE_TYPE(faketype uint8, realtype uint8) bool {
	return faketype == realtype || faketype == IsBool && (realtype == IsTrue || realtype == IsFalse)
}

func Z_OBJCE(zval Zval) *ClassEntry    { return zval.Object().GetCe() }
func Z_OBJCE_P(zval *Zval) *ClassEntry { return zval.Object().GetCe() }
func Z_OBJPROP(zval Zval) *Array {
	return zval.Object().GetPropertiesArray()
}
func Z_OBJPROP_P(zval_p *Zval) *Array       { return Z_OBJPROP(*zval_p) }
func Z_REFVAL(zval Zval) *Zval              { return zval.Reference().GetVal() }
func Z_REFVAL_P(zval_p *Zval) *Zval         { return zval_p.Reference().GetVal() }
func GC_AST(p *ZendAstRef) *zend.ZendAst    { return p.GcAst() }
func Z_ASTVAL(zval Zval) *zend.ZendAst      { return GC_AST(zval.ConstantAst()) }
func Z_ASTVAL_P(zval_p *Zval) *zend.ZendAst { return Z_ASTVAL(*zval_p) }
func Z_INDIRECT_P(zval_p *Zval) *Zval       { return zval_p.Indirect() }

func ZVAL_NEW_REF(z *Zval, r *Zval) { z.SetNewRef(r) }
func ZVAL_MAKE_REF_EX(z *Zval, refcount uint32) {
	var ref *Reference = NewZendReference(z)
	//ref.SetRefcount(refcount)
	z.SetReference(ref)
}
func ZVAL_PTR(z *Zval, p any) { z.SetPtr(p) }

func ZVAL_COPY_VALUE(z *Zval, v *Zval)  { z.CopyValueFrom(v) }
func ZVAL_COPY(z *Zval, v *Zval)        { z.CopyFrom(v) }
func ZVAL_COPY_OR_DUP(z *Zval, v *Zval) { z.CopyOrDupFrom(v) }
func ZVAL_DEREF(z *Zval) *Zval          { return z.DeRef() }
func ZVAL_DEINDIRECT(z *Zval) *Zval     { return z.DeIndirect() }

func ZVAL_MAKE_REF(zv *Zval) {
	var __zv *Zval = zv
	if !(__zv.IsRef()) {
		ZVAL_NEW_REF(__zv, __zv)
	}
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
	if zv.IsRef() {
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
