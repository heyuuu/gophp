// <<generate>>

package zend

import (
	b "sik/builtin"
)

func ZEND_TYPE_NAME(t ZendType) *ZendString   { return t.Name() }
func ZEND_TYPE_CE(t ZendType) *ZendClassEntry { return t.Ce() }

func ZEND_TYPE_ENCODE(code uint32, allow_null int) ZendType {
	if allow_null != 0 {
		return ZendType(code<<2 | 0x1)
	} else {
		return ZendType(code<<2 | 0x0)
	}
}
func ZEND_TYPE_ENCODE_CE(ce *ZendClassEntry, allow_null bool) ZendType {
	var ptr = b.CastUintptr(ce)
	if allow_null {
		return ZendType(ptr | 0x3)
	} else {
		return ZendType(ptr | 0x2)
	}
}
func ZEND_TYPE_ENCODE_CLASS(class_name *ZendString, allow_null ZendBool) ZendType {
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
func HT_HASH_TO_BUCKET_EX(data *Bucket, idx uint32) __auto__ { return data + idx }
func HT_IDX_TO_HASH(idx __auto__) __auto__                   { return idx }
func HT_HASH_EX(data __auto__, idx __auto__) uint32          { return (*uint32)(data)[int32(idx)] }
func HT_HASH(ht *HashTable, idx __auto__) uint32             { return HT_HASH_EX(ht.GetArData(), idx) }
func HT_SIZE_TO_MASK(nTableSize uint32) __auto__ {
	return uint32(-(nTableSize + nTableSize))
}
func HT_HASH_SIZE(nTableMask uint32) int {
	return (size_t(uint32) - int32(nTableMask)) * b.SizeOf("uint32_t")
}
func HT_DATA_SIZE(nTableSize uint32) int {
	return size_t(nTableSize) * b.SizeOf("Bucket")
}
func HT_SIZE_EX(nTableSize uint32, nTableMask uint32) int {
	return HT_DATA_SIZE(nTableSize) + HT_HASH_SIZE(nTableMask)
}
func HT_SIZE(ht *HashTable) int {
	return HT_SIZE_EX(ht.GetNTableSize(), ht.GetNTableMask())
}
func HT_USED_SIZE(ht *HashTable) int {
	return HT_HASH_SIZE(ht.GetNTableMask()) + size_t(ht).nNumUsed*b.SizeOf("Bucket")
}
func HT_HASH_TO_BUCKET(ht *HashTable, idx uint32) __auto__ {
	return HT_HASH_TO_BUCKET_EX(ht.GetArData(), idx)
}
func HT_SET_DATA_ADDR(ht *HashTable, ptr __auto__) {
	ht.SetArData((*Bucket)((*byte)(ptr) + HT_HASH_SIZE(ht.GetNTableMask())))
}
func HT_GET_DATA_ADDR(ht *HashTable) *byte {
	return (*byte)(ht.GetArData() - HT_HASH_SIZE(ht.GetNTableMask()))
}
func ZEND_PROPERTY_INFO_SOURCE_FROM_LIST(list *ZendPropertyInfoList) int { return 0x1 | uintPtr(list) }
func ZEND_PROPERTY_INFO_SOURCE_TO_LIST(list uintPtr) *ZendPropertyInfoList {
	return (*ZendPropertyInfoList)(list & ^0x1)
}
func ZEND_PROPERTY_INFO_SOURCE_IS_LIST(list uintPtr) int { return list & 0x1 }
func ZEND_SAME_FAKE_TYPE(faketype int, realtype ZendUchar) bool {
	return faketype == realtype || faketype == _IS_BOOL && (realtype == IS_TRUE || realtype == IS_FALSE)
}
func Z_FE_ITER_P(zval_p *Zval) uint32      { return zval_p.GetFeIterIdx() }
func Z_TYPE_INFO_REFCOUNTED(t uint32) bool { return (t & Z_TYPE_FLAGS_MASK) != 0 }

func GC_IS_RECURSIVE(p *HashTable) bool   { return p.IsRecursive() }
func GC_PROTECT_RECURSION(p *HashTable)   { p.ProtectRecursive() }
func GC_UNPROTECT_RECURSION(p *HashTable) { p.UnprotectRecursive() }

func Z_IS_RECURSIVE_P(zv *Zval) bool   { return zv.GetCounted().IsRecursive() }
func Z_PROTECT_RECURSION_P(zv *Zval)   { zv.GetCounted().ProtectRecursive() }
func Z_UNPROTECT_RECURSION_P(zv *Zval) { zv.GetCounted().UnprotectRecursive() }
func Z_CONSTANT(zval Zval) bool        { return zval.IsType(IS_CONSTANT_AST) }
func Z_REFCOUNTED(zval Zval) bool      { return zval.GetTypeFlags() != 0 }
func Z_REFCOUNTED_P(zval_p *Zval) bool { return Z_REFCOUNTED(*zval_p) }
func Z_COLLECTABLE(zval Zval) bool {
	return zval.HasTypeFlags(IS_TYPE_COLLECTABLE)
}
func Z_COLLECTABLE_P(zval_p *Zval) bool { return Z_COLLECTABLE(*zval_p) }
func Z_OPT_TYPE(zval Zval) int          { return zval.GetTypeInfo() & Z_TYPE_MASK }
func Z_OPT_TYPE_P(zval_p *Zval) int     { return Z_OPT_TYPE(*zval_p) }
func Z_OPT_CONSTANT(zval Zval) bool {
	return Z_OPT_TYPE(zval) == IS_CONSTANT_AST
}
func Z_OPT_REFCOUNTED(zval Zval) bool {
	return Z_TYPE_INFO_REFCOUNTED(zval.GetTypeInfo())
}
func Z_OPT_REFCOUNTED_P(zval_p *Zval) bool        { return Z_OPT_REFCOUNTED(*zval_p) }
func Z_OPT_ISREF(zval Zval) bool                  { return Z_OPT_TYPE(zval) == IS_REFERENCE }
func Z_OPT_ISREF_P(zval_p *Zval) bool             { return Z_OPT_ISREF(*zval_p) }
func Z_ISREF_P(zval_p *Zval) bool                 { return zval_p.IsReference() }
func Z_STR_P(zval_p *Zval) *ZendString            { return zval_p.GetStr() }
func Z_STRVAL(zval Zval) []byte                   { return zval.GetStr().GetVal() }
func Z_STRVAL_P(zval_p *Zval) []byte              { return zval_p.GetStr().GetVal() }
func Z_STRLEN(zval Zval) int                      { return zval.GetStr().GetLen() }
func Z_STRLEN_P(zval_p *Zval) int                 { return zval_p.GetStr().GetLen() }
func Z_ARRVAL(zval Zval) *ZendArray               { return zval.GetArr() }
func Z_ARRVAL_P(zval_p *Zval) *ZendArray          { return zval_p.GetArr() }
func Z_OBJ(zval Zval) *ZendObject                 { return zval.GetObj() }
func Z_OBJ_P(zval_p *Zval) *ZendObject            { return zval_p.GetObj() }
func Z_OBJ_HT(zval Zval) *ZendObjectHandlers      { return zval.GetObj().GetHandlers() }
func Z_OBJ_HT_P(zval_p *Zval) *ZendObjectHandlers { return zval_p.GetObj().GetHandlers() }
func Z_OBJ_HANDLE_P(zval_p *Zval) uint32          { return zval_p.GetObj().GetHandle() }
func Z_OBJCE(zval Zval) *ZendClassEntry           { return zval.GetObj().GetCe() }
func Z_OBJCE_P(zval_p *Zval) *ZendClassEntry      { return zval_p.GetObj().GetCe() }
func Z_OBJPROP(zval Zval) *HashTable {
	return Z_OBJ_HT(zval).GetGetProperties()(&zval)
}
func Z_OBJPROP_P(zval_p *Zval) *HashTable { return Z_OBJPROP(*zval_p) }
func Z_RES(zval Zval) *ZendResource       { return zval.GetRes() }
func Z_RES_P(zval_p *Zval) *ZendResource  { return zval_p.GetRes() }
func Z_RES_HANDLE(zval Zval) int          { return Z_RES(zval).GetHandle() }
func Z_RES_HANDLE_P(zval_p *Zval) int     { return Z_RES_HANDLE(*zval_p) }
func Z_RES_TYPE(zval Zval) int            { return Z_RES(zval).GetType() }
func Z_RES_TYPE_P(zval_p *Zval) int       { return Z_RES_TYPE(*zval_p) }
func Z_REF(zval Zval) *ZendReference      { return zval.GetRef() }
func Z_REF_P(zval_p *Zval) *ZendReference { return zval_p.GetRef() }
func Z_REFVAL(zval Zval) Zval             { return Z_REF(zval).GetVal() }
func Z_REFVAL_P(zval_p *Zval) Zval        { return Z_REFVAL(*zval_p) }
func GC_AST(p *ZendAstRef) *ZendAst {
	return (*ZendAst)((*byte)(p) + b.SizeOf("zend_ast_ref"))
}
func Z_ASTVAL(zval Zval) *ZendAst      { return GC_AST(zval.GetAst()) }
func Z_ASTVAL_P(zval_p *Zval) *ZendAst { return Z_ASTVAL(*zval_p) }
func Z_INDIRECT(zval Zval) *Zval       { return zval.GetZv() }
func Z_INDIRECT_P(zval_p *Zval) *Zval  { return zval_p.GetZv() }
func Z_CE(zval Zval) *ZendClassEntry   { return zval.GetCe() }
func Z_PTR(zval Zval) any              { return zval.GetPtr() }

func ZVAL_UNDEF(z *Zval)             { z.SetUndef() }
func ZVAL_NULL(z *Zval)              { z.SetNull() }
func ZVAL_FALSE(z *Zval)             { z.SetFalse() }
func ZVAL_TRUE(z *Zval)              { z.SetTrue() }
func ZVAL_BOOL(z *Zval, b int)       { z.SetBool(b != 0) }
func ZVAL_LONG(z *Zval, l ZendLong)  { z.SetLong(l) }
func ZVAL_DOUBLE(z *Zval, d float64) { z.SetDouble(d) }
func ZVAL_STR(z *Zval, s *ZendString) {
	z.SetStr(s)
	z.SetTypeInfo(IS_STRING_EX)
}
func ZVAL_INTERNED_STR(z *Zval, s *ZendString) {
	var z *Zval = z
	var __s *ZendString = s
	z.SetStr(__s)
	z.SetTypeInfo(IS_INTERNED_STRING_EX)
}
func ZVAL_NEW_STR(z *Zval, s *ZendString) {
	var z *Zval = z
	var __s *ZendString = s
	z.SetStr(__s)
	z.SetTypeInfo(IS_STRING_EX)
}
func ZVAL_STR_COPY(z *Zval, s *ZendString) {
	var z *Zval = z
	var __s *ZendString = s
	z.SetStr(__s)
	__s.AddRefcount()
	z.SetTypeInfo(IS_STRING_EX)
}
func ZVAL_ARR(z *Zval, a *ZendArray) {
	var __arr *ZendArray = a
	var z *Zval = z
	z.SetArr(__arr)
	z.SetTypeInfo(IS_ARRAY_EX)
}
func ZVAL_NEW_PERSISTENT_ARR(z *Zval) {
	var z *Zval = z
	var _arr *ZendArray = (*ZendArray)(Malloc(b.SizeOf("zend_array")))
	z.SetArr(_arr)
	z.SetTypeInfo(IS_ARRAY_EX)
}
func ZVAL_OBJ(z *Zval, o *ZendObject) {
	var z *Zval = z
	z.SetObj(o)
	z.SetTypeInfo(IS_OBJECT_EX)
}
func ZVAL_RES(z *Zval, r *ZendResource) {
	var z *Zval = z
	z.SetRes(r)
	z.SetTypeInfo(IS_RESOURCE_EX)
}
func ZVAL_NEW_RES(z *Zval, h int, p any, t int) {
	var _res *ZendResource = (*ZendResource)(Emalloc(b.SizeOf("zend_resource")))
	var z *Zval
	_res.SetRefcount(1)
	_res.GetGcTypeInfo() = IS_RESOURCE
	_res.SetHandle(h)
	_res.SetType(t)
	_res.SetPtr(p)
	z = z
	z.SetRes(_res)
	z.SetTypeInfo(IS_RESOURCE_EX)
}
func ZVAL_NEW_PERSISTENT_RES(z *Zval, h int, p any, t int) {
	var _res *ZendResource = (*ZendResource)(Malloc(b.SizeOf("zend_resource")))
	var z *Zval
	_res.SetRefcount(1)
	_res.GetGcTypeInfo() = IS_RESOURCE | GC_PERSISTENT<<GC_FLAGS_SHIFT
	_res.SetHandle(h)
	_res.SetType(t)
	_res.SetPtr(p)
	z = z
	z.SetRes(_res)
	z.SetTypeInfo(IS_RESOURCE_EX)
}
func ZVAL_REF(z *Zval, r *ZendReference) {
	var z *Zval = z
	z.SetRef(r)
	z.SetTypeInfo(IS_REFERENCE_EX)
}
func ZVAL_NEW_EMPTY_REF(z *Zval) {
	var _ref *ZendReference = (*ZendReference)(Emalloc(b.SizeOf("zend_reference")))
	_ref.SetRefcount(1)
	_ref.GetGcTypeInfo() = IS_REFERENCE
	_ref.GetSources().SetPtr(nil)
	z.SetRef(_ref)
	z.SetTypeInfo(IS_REFERENCE_EX)
}
func ZVAL_NEW_REF(z *Zval, r *Zval) {
	var _ref *ZendReference = (*ZendReference)(Emalloc(b.SizeOf("zend_reference")))
	_ref.SetRefcount(1)
	_ref.GetGcTypeInfo() = IS_REFERENCE
	ZVAL_COPY_VALUE(_ref.GetVal(), r)
	_ref.GetSources().SetPtr(nil)
	z.SetRef(_ref)
	z.SetTypeInfo(IS_REFERENCE_EX)
}
func ZVAL_MAKE_REF_EX(z *Zval, refcount uint32) {
	var _z *Zval = z
	var _ref *ZendReference = (*ZendReference)(Emalloc(b.SizeOf("zend_reference")))
	_ref.SetRefcount(refcount)
	_ref.GetGcTypeInfo() = IS_REFERENCE
	ZVAL_COPY_VALUE(_ref.GetVal(), _z)
	_ref.GetSources().SetPtr(nil)
	_z.SetRef(_ref)
	_z.SetTypeInfo(IS_REFERENCE_EX)
}
func ZVAL_AST(z *Zval, ast *ZendAstRef) {
	var z *Zval = z
	z.SetAst(ast)
	z.SetTypeInfo(IS_CONSTANT_AST_EX)
}
func ZVAL_INDIRECT(z *Zval, v *Zval) {
	z.SetZv(v)
	z.SetTypeInfo(IS_INDIRECT)
}
func ZVAL_PTR(z *Zval, p any) {
	z.SetPtr(p)
	z.SetTypeInfo(IS_PTR)
}
func ZVAL_ALIAS_PTR(z *Zval, p *ZendClassEntry) {
	z.SetPtr(p)
	z.SetTypeInfo(IS_ALIAS_PTR)
}
func ZVAL_ERROR(z *Zval)                          { z.SetTypeInfo(_IS_ERROR) }
func Z_REFCOUNT_P(pz *Zval) uint32                { return ZvalRefcountP(pz) }
func Z_SET_REFCOUNT_P(pz *Zval, rc uint32) uint32 { return ZvalSetRefcountP(pz, rc) }
func Z_ADDREF_P(pz *Zval) uint32                  { return ZvalAddrefP(pz) }
func Z_DELREF_P(pz *Zval) uint32                  { return ZvalDelrefP(pz) }
func Z_REFCOUNT(z Zval) uint32                    { return Z_REFCOUNT_P(&z) }
func Z_SET_REFCOUNT(z Zval, rc uint32) uint32     { return Z_SET_REFCOUNT_P(&z, rc) }
func Z_ADDREF(z Zval) uint32                      { return Z_ADDREF_P(&z) }
func Z_DELREF(z Zval) uint32                      { return Z_DELREF_P(&z) }
func Z_TRY_ADDREF_P(pz *Zval) {
	if Z_REFCOUNTED_P(pz) {
		Z_ADDREF_P(pz)
	}
}
func Z_TRY_DELREF_P(pz *Zval) {
	if Z_REFCOUNTED_P(pz) {
		Z_DELREF_P(pz)
	}
}
func Z_TRY_ADDREF(z Zval)                    { Z_TRY_ADDREF_P(&z) }
func Z_TRY_DELREF(z Zval)                    { Z_TRY_DELREF_P(&z) }
func GC_MAKE_PERSISTENT_LOCAL(p *ZendString) {}
func ZvalRefcountP(pz *Zval) uint32          { return pz.GetCounted().GetRefcount() }
func ZvalSetRefcountP(pz *Zval, rc uint32) uint32 {
	ZEND_ASSERT(Z_REFCOUNTED_P(pz))
	return pz.GetCounted().SetRefcount(rc)
}
func ZvalAddrefP(pz *Zval) uint32 {
	ZEND_ASSERT(Z_REFCOUNTED_P(pz))
	return pz.GetCounted().AddRefcount()
}
func ZvalDelrefP(pz *Zval) uint32 {
	ZEND_ASSERT(Z_REFCOUNTED_P(pz))
	return pz.GetCounted().DelRefcount()
}
func ZVAL_COPY_VALUE_EX(z *Zval, v *Zval, gc *ZendRefcounted, t uint32) {
	z.SetCounted(gc)
	z.SetTypeInfo(t)
}
func ZVAL_COPY_VALUE(z *Zval, v *Zval) {
	// 复制除 u2 外所有数据
	var temp = z.u2
	*z = *v
	z.u2 = temp
}
func ZVAL_COPY(z *Zval, v *Zval) {
	ZVAL_COPY_VALUE(z, v)
	// 若支持引用计数，则增加计数；此时 z、v 指向同一个 value，增加哪个都一样
	if v.IsRefcounted() {
		z.GetCounted().AddRefcount()
	}
}
func ZVAL_COPY_OR_DUP(z *Zval, v *Zval) {
	ZVAL_COPY_VALUE(z, v)
	if v.IsRefcounted() {
		if v.GetCounted().HasGcFlags(GC_PERSISTENT) {
			v.GetCounted().AddRefcount()
		} else {
			ZvalCopyCtorFunc(z)
		}
	}
}
func ZVAL_DEREF(z *Zval) {
	if Z_ISREF_P(z) {
		z = Z_REFVAL_P(z)
	}
}
func ZVAL_DEINDIRECT(z *Zval) {
	if z.IsIndirect() {
		z = z.GetZv()
	}
}
func ZVAL_MAKE_REF(zv *Zval) {
	var zv *Zval = zv
	if !(Z_ISREF_P(zv)) {
		ZVAL_NEW_REF(zv, zv)
	}
}
func ZVAL_UNREF(z *Zval) {
	var _z *Zval = z
	var ref *ZendReference
	ZEND_ASSERT(Z_ISREF_P(_z))
	ref = _z.GetRef()
	ZVAL_COPY_VALUE(_z, ref.GetVal())
	EfreeSize(ref, b.SizeOf("zend_reference"))
}
func ZVAL_COPY_DEREF(z *Zval, v *Zval) {
	var _z3 *Zval = v
	if Z_OPT_REFCOUNTED_P(_z3) {
		if Z_OPT_ISREF_P(_z3) {
			_z3 = Z_REFVAL_P(_z3)
			if Z_OPT_REFCOUNTED_P(_z3) {
				Z_ADDREF_P(_z3)
			}
		} else {
			Z_ADDREF_P(_z3)
		}
	}
	ZVAL_COPY_VALUE(z, _z3)
}
func SEPARATE_ARRAY(zv *Zval) {
	var _zv *Zval = zv
	var _arr *ZendArray = _zv.GetArr()
	if _arr.GetRefcount() > 1 {
		if Z_REFCOUNTED_P(_zv) {
			_arr.DelRefcount()
		}
		ZVAL_ARR(_zv, ZendArrayDup(_arr))
	}
}
func SEPARATE_ZVAL_IF_NOT_REF(zv *Zval) {
	var zv *Zval = zv
	if zv.IsArray() {
		if Z_REFCOUNT_P(zv) > 1 {
			if Z_REFCOUNTED_P(zv) {
				Z_DELREF_P(zv)
			}
			ZVAL_ARR(zv, ZendArrayDup(zv.GetArr()))
		}
	}
}
func SEPARATE_ZVAL_NOREF(zv *Zval) {
	var _zv *Zval = zv
	ZEND_ASSERT(_zv.GetType() != IS_REFERENCE)
	SEPARATE_ZVAL_IF_NOT_REF(_zv)
}
func SEPARATE_ZVAL(zv *Zval) {
	for {
		var _zv *Zval = zv
		if Z_ISREF_P(_zv) {
			var _r *ZendReference = _zv.GetRef()
			ZVAL_COPY_VALUE(_zv, _r.GetVal())
			if _r.DelRefcount() == 0 {
				EfreeSize(_r, b.SizeOf("zend_reference"))
			} else if Z_OPT_TYPE_P(_zv) == IS_ARRAY {
				ZVAL_ARR(_zv, ZendArrayDup(_zv.GetArr()))
				break
			} else if Z_OPT_REFCOUNTED_P(_zv) {
				Z_ADDREF_P(_zv)
				break
			}
		}
		SEPARATE_ZVAL_IF_NOT_REF(_zv)
		break
	}
}
func SEPARATE_ARG_IF_REF(varptr *Zval) {
	ZVAL_DEREF(varptr)
	if Z_REFCOUNTED_P(varptr) {
		Z_ADDREF_P(varptr)
	}
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
