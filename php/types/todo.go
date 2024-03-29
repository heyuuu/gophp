package types

import (
	"github.com/heyuuu/gophp/php/assert"
	"github.com/heyuuu/gophp/php/perr"
)

func SeparateArray(zv *Zval) {
	assert.Assert(zv.IsArray())
	zv.SetArray(zv.Array().RealClone())
}
func SeparateZval(zv *Zval) {
	// 解 Ref
	if zv.IsRef() {
		*zv = zv.DeRef()
	}
	// 仅数组需要分离
	if zv.IsArray() {
		zv.SetArray(zv.Array().RealClone())
	}
}
func ZVAL_MAKE_REF(zv *Zval) {
	var __zv *Zval = zv
	if !(__zv.IsRef()) {
		ZVAL_NEW_REF(__zv, __zv)
	}
}

func ZVAL_NEW_REF(z *Zval, r *Zval) {
	//z.SetNewRef(r)
	panic(perr.Todof("ZVAL_NEW_REF"))
}

func ZVAL_COPY_VALUE(z *Zval, v *Zval) { *z = *v }

func Z_OBJPROP_P(v Zval) *Array {
	return nil
}
