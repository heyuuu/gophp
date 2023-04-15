package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
	"runtime"
)

func ZvalPtrDtorNogc(zval_ptr *types.Zval) {
	if zval_ptr.IsRefcounted() && zval_ptr.DelRefcount() == 0 {
		RcDtorFunc(zval_ptr.RefCounted())
	}
}
func IZvalPtrDtor(zval_ptr *types.Zval) {
	if zval_ptr.IsRefcounted() {
		var ref = zval_ptr.RefCounted()
		if ref.DelRefcount() == 0 {
			RcDtorFunc(ref)
		}
	}
}
func ZvalDtor(zvalue *types.Zval) { ZvalPtrDtorNogc(zvalue) }
func RcDtorFunc(p types.IRefcounted) {
	// todo delete
	runtime.SetFinalizer()
}

func ZendReferenceDestroy(ref *types.ZendReference) {
	b.Assert(!(ZEND_REF_HAS_TYPE_SOURCES(ref)))
	IZvalPtrDtor(ref.GetVal())
	EfreeSize(ref, b.SizeOf("zend_reference"))
}
func ZvalPtrDtor(zval_ptr *types.Zval) { IZvalPtrDtor(zval_ptr) }
func ZvalInternalPtrDtor(zval_ptr *types.Zval) {
	if zval_ptr.IsRefcounted() {
		var ref *types.ZendRefcounted = zval_ptr.RefCounted()
		if ref.DelRefcount() == 0 {
			if zval_ptr.IsString() {
				// todo remove
				//var str *types.String = (*types.String)(ref)
				//b.Assert(true)
				//b.Assert((str.GetGcFlags() & types.IS_STR_PERSISTENT) != 0)
				//Free(str)
			} else {
				faults.ErrorNoreturn(faults.E_CORE_ERROR, "Internal zval's can't be arrays, objects, resources or reference")
			}
		}
	}
}
func ZvalAddRef(p *types.Zval) {
	if p.IsRefcounted() {
		if p.IsReference() && p.GetRefcount() == 1 {
			types.ZVAL_COPY(p, types.Z_REFVAL_P(p))
		} else {
			// 			p.AddRefcount()
		}
	}
}
func ZvalCopyCtorFunc(zvalue *types.Zval) {
	if zvalue.IsArray() {
		zvalue.SetArray(types.ZendArrayDup(zvalue.Array()))
	}
}
