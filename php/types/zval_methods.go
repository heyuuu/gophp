package types

// ZVAL_COPY_VALUE
func (zv *Zval) CopyValueFrom(v *Zval) {
	// 复制除 u2 外所有数据
	zv.v = v.v
}
func (zv *Zval) CopyValue() *Zval {
	var tmp Zval
	tmp.CopyValueFrom(zv)
	return &tmp
}

func (zv *Zval) CopyFrom(v *Zval) {
	zv.CopyValueFrom(v)
	// 除数组外，基础类型都复制了值，引用类型都复制了指针；仅数组需要做写时复制
	if v.IsArray() {
		zv.SetArray(zv.Array().Copy())
	}
}
func (zv *Zval) Copy() *Zval {
	var tmp Zval
	tmp.CopyFrom(zv)
	return &tmp
}

func (zv *Zval) CopyOrDupFrom(v *Zval) {
	zv.CopyValueFrom(v)
	// 除数组外，基础类型都复制了值，引用类型都复制了指针；仅数组需要做写时复制
	if zv.IsArray() {
		zv.SetArray(ZendArrayDup(zv.Array()))
	}
}

// ZVAL_COPY_PROP
func (zv *Zval) CopyPropFrom(v *Zval) {
	zv.CopyFrom(v)
	zv.u2 = v.u2
}

// ZVAL_COPY_OR_DUP_PROP
func (zv *Zval) CopyOrDupPropFrom(v *Zval) {
	zv.CopyOrDupFrom(v)
	zv.u2 = v.u2
}

// ZVAL_DEREF(zv)
func (zv *Zval) DeRef() *Zval {
	if zv.IsRef() {
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
	switch zv.Type() {
	case IsArray, // 不包含 _IS_IMMUTABLE_ARRAY
		IsObject,
		IsResource,
		IsRef:
		return true
	default:
		return false
	}
}
