// <<generate>>

package zend

import (
	b "sik/builtin"
)

func ZvalPtrDtorNogc(zval_ptr *Zval) {
	if Z_REFCOUNTED_P(zval_ptr) && Z_DELREF_P(zval_ptr) == 0 {
		RcDtorFunc(Z_COUNTED_P(zval_ptr))
	}
}
func IZvalPtrDtor(zval_ptr *Zval) {
	if Z_REFCOUNTED_P(zval_ptr) {
		var ref *ZendRefcounted = Z_COUNTED_P(zval_ptr)
		if GC_DELREF(ref) == 0 {
			RcDtorFunc(ref)
		} else {
			GcCheckPossibleRoot(ref)
		}
	}
}
func ZvalCopyCtor(zvalue *Zval) {
	if zvalue.GetType() == IS_ARRAY {
		ZVAL_ARR(zvalue, ZendArrayDup(Z_ARR_P(zvalue)))
	} else if Z_REFCOUNTED_P(zvalue) {
		Z_ADDREF_P(zvalue)
	}
}
func ZvalOptCopyCtor(zvalue *Zval) {
	if Z_OPT_TYPE_P(zvalue) == IS_ARRAY {
		ZVAL_ARR(zvalue, ZendArrayDup(Z_ARR_P(zvalue)))
	} else if Z_OPT_REFCOUNTED_P(zvalue) {
		Z_ADDREF_P(zvalue)
	}
}
func ZvalPtrDtorStr(zval_ptr *Zval) {
	if Z_REFCOUNTED_P(zval_ptr) && Z_DELREF_P(zval_ptr) == 0 {
		ZEND_ASSERT(zval_ptr.GetType() == IS_STRING)
		ZEND_ASSERT(true)
		ZEND_ASSERT((GC_FLAGS(Z_STR_P(zval_ptr)) & IS_STR_PERSISTENT) == 0)
		Efree(Z_STR_P(zval_ptr))
	}
}
func ZvalDtor(zvalue *Zval)         { ZvalPtrDtorNogc(zvalue) }
func ZvalInternalDtor(zvalue *Zval) { ZvalInternalPtrDtor(zvalue) }
func RcDtorFunc(p *ZendRefcounted) {
	ZEND_ASSERT(GC_TYPE(p) <= IS_CONSTANT_AST)
	ZendRcDtorFunc[GC_TYPE(p)](p)
}
func ZendReferenceDestroy(ref *ZendReference) {
	ZEND_ASSERT(!(ZEND_REF_HAS_TYPE_SOURCES(ref)))
	IZvalPtrDtor(ref.GetVal())
	EfreeSize(ref, b.SizeOf("zend_reference"))
}
func ZendEmptyDestroy(ref *ZendReference) {}
func ZvalPtrDtor(zval_ptr *Zval)          { IZvalPtrDtor(zval_ptr) }
func ZvalInternalPtrDtor(zval_ptr *Zval) {
	if Z_REFCOUNTED_P(zval_ptr) {
		var ref *ZendRefcounted = Z_COUNTED_P(zval_ptr)
		if GC_DELREF(ref) == 0 {
			if zval_ptr.GetType() == IS_STRING {
				var str *ZendString = (*ZendString)(ref)
				ZEND_ASSERT(true)
				ZEND_ASSERT((GC_FLAGS(str) & IS_STR_PERSISTENT) != 0)
				Free(str)
			} else {
				ZendErrorNoreturn(E_CORE_ERROR, "Internal zval's can't be arrays, objects, resources or reference")
			}
		}
	}
}
func ZvalAddRef(p *Zval) {
	if Z_REFCOUNTED_P(p) {
		if Z_ISREF_P(p) && Z_REFCOUNT_P(p) == 1 {
			ZVAL_COPY(p, Z_REFVAL_P(p))
		} else {
			Z_ADDREF_P(p)
		}
	}
}
func ZvalCopyCtorFunc(zvalue *Zval) {
	if zvalue.GetType() == IS_ARRAY {
		ZVAL_ARR(zvalue, ZendArrayDup(Z_ARRVAL_P(zvalue)))
	} else if zvalue.GetType() == IS_STRING {
		ZEND_ASSERT(true)
		ZVAL_NEW_STR(zvalue, ZendStringDup(Z_STR_P(zvalue), 0))
	}
}
