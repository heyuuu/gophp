// <<generate>>

package zend

import (
	b "sik/builtin"
)

func ZvalPtrDtorNogc(zval_ptr *Zval) {
	if zval_ptr.IsRefcounted() && zval_ptr.DelRefcount() == 0 {
		RcDtorFunc(zval_ptr.GetCounted())
	}
}
func IZvalPtrDtor(zval_ptr *Zval) {
	if zval_ptr.IsRefcounted() {
		var ref *ZendRefcounted = zval_ptr.GetCounted()
		if ref.DelRefcount() == 0 {
			RcDtorFunc(ref)
		} else {
			GcCheckPossibleRoot(ref)
		}
	}
}
func ZvalCopyCtor(zvalue *Zval) {
	if zvalue.IsArray() {
		ZVAL_ARR(zvalue, ZendArrayDup(zvalue.GetArr()))
	} else if zvalue.IsRefcounted() {
		zvalue.AddRefcount()
	}
}
func ZvalOptCopyCtor(zvalue *Zval) {
	if zvalue.IsArray() {
		ZVAL_ARR(zvalue, ZendArrayDup(zvalue.GetArr()))
	} else if zvalue.IsRefcounted() {
		zvalue.AddRefcount()
	}
}
func ZvalPtrDtorStr(zval_ptr *Zval) {
	if zval_ptr.IsRefcounted() && zval_ptr.DelRefcount() == 0 {
		ZEND_ASSERT(zval_ptr.IsString())
		ZEND_ASSERT(true)
		ZEND_ASSERT((zval_ptr.GetStr().GetGcFlags() & IS_STR_PERSISTENT) == 0)
		Efree(zval_ptr.GetStr())
	}
}
func ZvalDtor(zvalue *Zval)         { ZvalPtrDtorNogc(zvalue) }
func ZvalInternalDtor(zvalue *Zval) { ZvalInternalPtrDtor(zvalue) }
func RcDtorFunc(p *ZendRefcounted) {
	ZEND_ASSERT(p.GetGcType() <= IS_CONSTANT_AST)
	ZendRcDtorFunc[p.GetGcType()](p)
}
func ZendReferenceDestroy(ref *ZendReference) {
	ZEND_ASSERT(!(ZEND_REF_HAS_TYPE_SOURCES(ref)))
	IZvalPtrDtor(ref.GetVal())
	EfreeSize(ref, b.SizeOf("zend_reference"))
}
func ZendEmptyDestroy(ref *ZendReference) {}
func ZvalPtrDtor(zval_ptr *Zval)          { IZvalPtrDtor(zval_ptr) }
func ZvalInternalPtrDtor(zval_ptr *Zval) {
	if zval_ptr.IsRefcounted() {
		var ref *ZendRefcounted = zval_ptr.GetCounted()
		if ref.DelRefcount() == 0 {
			if zval_ptr.IsString() {
				var str *ZendString = (*ZendString)(ref)
				ZEND_ASSERT(true)
				ZEND_ASSERT((str.GetGcFlags() & IS_STR_PERSISTENT) != 0)
				Free(str)
			} else {
				ZendErrorNoreturn(E_CORE_ERROR, "Internal zval's can't be arrays, objects, resources or reference")
			}
		}
	}
}
func ZvalAddRef(p *Zval) {
	if p.IsRefcounted() {
		if p.IsReference() && p.GetRefcount() == 1 {
			ZVAL_COPY(p, Z_REFVAL_P(p))
		} else {
			p.AddRefcount()
		}
	}
}
func ZvalCopyCtorFunc(zvalue *Zval) {
	if zvalue.IsArray() {
		ZVAL_ARR(zvalue, ZendArrayDup(zvalue.GetArr()))
	} else if zvalue.IsString() {
		ZEND_ASSERT(true)
		ZVAL_NEW_STR(zvalue, zvalue.GetStr().Dup(0))
	}
}
