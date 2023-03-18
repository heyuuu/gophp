// <<generate>>

package zend

import (
	b "sik/builtin"
	"sik/zend/types"
)

func ZvalPtrDtorNogc(zval_ptr *types.Zval) {
	if zval_ptr.IsRefcounted() && zval_ptr.DelRefcount() == 0 {
		RcDtorFunc(zval_ptr.GetCounted())
	}
}
func IZvalPtrDtor(zval_ptr *types.Zval) {
	if zval_ptr.IsRefcounted() {
		var ref = zval_ptr.GetCounted()
		if ref.DelRefcount() == 0 {
			RcDtorFunc(ref)
		} else {
			GcCheckPossibleRoot(ref)
		}
	}
}
func ZvalPtrDtorStr(zval_ptr *types.Zval) {
	if zval_ptr.IsRefcounted() && zval_ptr.DelRefcount() == 0 {
		ZEND_ASSERT(zval_ptr.IsString())
		ZEND_ASSERT(true)
		ZEND_ASSERT((zval_ptr.GetStr().GetGcFlags() & types.IS_STR_PERSISTENT) == 0)
		Efree(zval_ptr.GetStr())
	}
}
func ZvalDtor(zvalue *types.Zval) { ZvalPtrDtorNogc(zvalue) }
func RcDtorFunc(p types.IRefcounted) {
	ZEND_ASSERT(p.GetGcType() <= types.IS_CONSTANT_AST)
	switch p.(type) {
	case *types.ZendArray:
		arr := p.(*types.ZendArray)
		arr.DestroyEx()
	case *types.ZendObject:
		obj := p.(*types.ZendObject)
		ZendObjectsStoreDel(obj)
	case *types.ZendResource:
		res := p.(*types.ZendResource)
		ZendListFree(res)
	case *types.ZendReference:
		ref := p.(*types.ZendReference)
		ZendReferenceDestroy(ref)
	case *types.ZendAstRef:
		ast := p.(*types.ZendAstRef)
		ZendAstRefDestroy(ast)
	}
}

func ZendReferenceDestroy(ref *types.ZendReference) {
	ZEND_ASSERT(!(ZEND_REF_HAS_TYPE_SOURCES(ref)))
	IZvalPtrDtor(ref.GetVal())
	EfreeSize(ref, b.SizeOf("zend_reference"))
}
func ZvalPtrDtor(zval_ptr *types.Zval) { IZvalPtrDtor(zval_ptr) }
func ZvalInternalPtrDtor(zval_ptr *types.Zval) {
	if zval_ptr.IsRefcounted() {
		var ref *types.ZendRefcounted = zval_ptr.GetCounted()
		if ref.DelRefcount() == 0 {
			if zval_ptr.IsString() {
				var str *types.ZendString = (*types.ZendString)(ref)
				ZEND_ASSERT(true)
				ZEND_ASSERT((str.GetGcFlags() & types.IS_STR_PERSISTENT) != 0)
				Free(str)
			} else {
				ZendErrorNoreturn(E_CORE_ERROR, "Internal zval's can't be arrays, objects, resources or reference")
			}
		}
	}
}
func ZvalAddRef(p *types.Zval) {
	if p.IsRefcounted() {
		if p.IsReference() && p.GetRefcount() == 1 {
			types.ZVAL_COPY(p, types.Z_REFVAL_P(p))
		} else {
			p.AddRefcount()
		}
	}
}
func ZvalCopyCtorFunc(zvalue *types.Zval) {
	if zvalue.IsArray() {
		zvalue.SetArray(ZendArrayDup(zvalue.GetArr()))
	} else if zvalue.IsString() {
		ZEND_ASSERT(true)
		zvalue.SetString(zvalue.GetStr().Dup(0))
	}
}
