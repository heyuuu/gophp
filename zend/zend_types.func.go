// <<generate>>

package zend

import (
	b "sik/builtin"
)

func ZEND_TYPE_IS_SET(t ZendType) bool        { return t.IsSet() }
func ZEND_TYPE_IS_CODE(t ZendType) bool       { return t.IsCode() }
func ZEND_TYPE_IS_CLASS(t ZendType) bool      { return t.IsClass() }
func ZEND_TYPE_IS_CE(t ZendType) bool         { return t.IsCe() }
func ZEND_TYPE_IS_NAME(t ZendType) bool       { return t.IsName() }
func ZEND_TYPE_NAME(t ZendType) *ZendString   { return t.Name() }
func ZEND_TYPE_CE(t ZendType) *ZendClassEntry { return t.Ce() }
func ZEND_TYPE_CODE(t ZendType) int           { return t.Code() }
func ZEND_TYPE_ALLOW_NULL(t ZendType) bool    { return t.AllowNull() }

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
func ZEND_TYPE_ENCODE_CLASS_CONST(class_name string, allow_null int) __auto__ {
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
func HT_HASH_TO_IDX(idx uint32) uint32                       { return idx }
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
func HT_HASH_RESET(ht *HashTable) __auto__ {
	return memset(&(HT_HASH(ht, ht.GetNTableMask())), HT_INVALID_IDX, HT_HASH_SIZE(ht.GetNTableMask()))
}
func HT_HASH_RESET_PACKED(ht *HashTable) {
	HT_HASH(ht, -2) = HT_INVALID_IDX
	HT_HASH(ht, -1) = HT_INVALID_IDX
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
func ZvalGetType(pz *Zval) ZendUchar                     { return pz.GetType() }
func ZEND_SAME_FAKE_TYPE(faketype int, realtype ZendUchar) bool {
	return faketype == realtype || faketype == _IS_BOOL && (realtype == IS_TRUE || realtype == IS_FALSE)
}
func Z_FE_ITER_P(zval_p *Zval) uint32                 { return zval_p.GetFeIterIdx() }
func GC_REFCOUNT(p *HashTable) uint32                 { return p.GetRefcount() }
func GC_SET_REFCOUNT(p IRefcounted, rc uint32) uint32 { return p.SetRefcount(rc) }
func GC_ADDREF(p IRefcounted) uint32                  { return p.AddRefcount() }
func GC_DELREF(p IRefcounted) uint32                  { return p.DelRefcount() }
func GC_ADDREF_EX(p IRefcounted, rc uint32) uint32    { return p.AddRefcountEx(rc) }
func GC_DELREF_EX(p IRefcounted, rc uint32) uint32    { return p.DelRefcountEx(rc) }
func ZvalGcType(gc_type_info uint32) ZendUchar        { return gc_type_info & GC_TYPE_MASK }
func ZvalGcFlags(gc_type_info uint32) uint32 {
	return gc_type_info >> GC_FLAGS_SHIFT & GC_FLAGS_MASK >> GC_FLAGS_SHIFT
}
func ZvalGcInfo(gc_type_info uint32) uint32    { return gc_type_info >> GC_INFO_SHIFT }
func GC_TYPE_INFO(p IRefcounted) uint32        { return p.GetGcTypeInfo() }
func GC_TYPE(p IRefcounted) ZendUchar          { return p.GetGcType() }
func GC_FLAGS(p IRefcounted) uint32            { return p.GetGcFlags() }
func GC_INFO(p IRefcounted) uint32             { return p.GetGcInfo() }
func GC_ADD_FLAGS(p IRefcounted, flags uint32) { p.AddGcFlags(flags) }
func GC_DEL_FLAGS(p IRefcounted, flags uint32) { p.DelGcFlags(flags) }
func Z_GC_TYPE(zval Zval) ZendUchar            { return zval.GetCounted().GetGcType() }
func Z_GC_TYPE_P(zval_p *Zval) ZendUchar       { return Z_GC_TYPE(*zval_p) }
func Z_GC_FLAGS(zval Zval) uint32              { return zval.GetCounted().GetGcFlags() }
func Z_GC_FLAGS_P(zval_p *Zval) uint32         { return Z_GC_FLAGS(*zval_p) }
func Z_GC_INFO(zval Zval) uint32               { return zval.GetCounted().GetGcInfo() }
func Z_GC_INFO_P(zval_p *Zval) uint32          { return Z_GC_INFO(*zval_p) }
func Z_GC_TYPE_INFO(zval Zval) __auto__        { return zval.GetCounted().GetGcTypeInfo() }
func Z_GC_TYPE_INFO_P(zval_p *Zval) __auto__   { return Z_GC_TYPE_INFO(*zval_p) }
func Z_TYPE_INFO_REFCOUNTED(t uint32) bool     { return (t & Z_TYPE_FLAGS_MASK) != 0 }
func OBJ_FLAGS(obj IRefcounted) uint32         { return obj.GetGcFlags() }
func GC_IS_RECURSIVE(p *HashTable) int         { return p.GetGcFlags() & GC_PROTECTED }
func GC_PROTECT_RECURSION(p *HashTable)        { p.AddGcFlags(GC_PROTECTED) }
func GC_UNPROTECT_RECURSION(p *HashTable)      { p.DelGcFlags(GC_PROTECTED) }
func GC_TRY_PROTECT_RECURSION(p *HashTable) {
	if (p.GetGcFlags() & GC_IMMUTABLE) == 0 {
		GC_PROTECT_RECURSION(p)
	}
}
func GC_TRY_UNPROTECT_RECURSION(p *HashTable) {
	if (p.GetGcFlags() & GC_IMMUTABLE) == 0 {
		GC_UNPROTECT_RECURSION(p)
	}
}
func Z_IS_RECURSIVE(zval Zval) int {
	return GC_IS_RECURSIVE(zval.GetCounted())
}
func Z_PROTECT_RECURSION(zval Zval) { GC_PROTECT_RECURSION(zval.GetCounted()) }
func Z_UNPROTECT_RECURSION(zval Zval) {
	GC_UNPROTECT_RECURSION(zval.GetCounted())
}
func Z_IS_RECURSIVE_P(zv *Zval) int    { return Z_IS_RECURSIVE(*zv) }
func Z_PROTECT_RECURSION_P(zv *Zval)   { Z_PROTECT_RECURSION(*zv) }
func Z_UNPROTECT_RECURSION_P(zv *Zval) { Z_UNPROTECT_RECURSION(*zv) }
func Z_CONSTANT(zval Zval) bool        { return zval.IsType(IS_CONSTANT_AST) }
func Z_CONSTANT_P(zval_p *Zval) bool   { return Z_CONSTANT(*zval_p) }
func Z_REFCOUNTED(zval Zval) bool      { return zval.GetTypeFlags() != 0 }
func Z_REFCOUNTED_P(zval_p *Zval) bool { return Z_REFCOUNTED(*zval_p) }
func Z_COLLECTABLE(zval Zval) bool {
	return zval.HasTypeFlags(IS_TYPE_COLLECTABLE)
}
func Z_COLLECTABLE_P(zval_p *Zval) bool   { return Z_COLLECTABLE(*zval_p) }
func Z_COPYABLE(zval __auto__) bool       { return zval.u1.v.type_ == IS_ARRAY }
func Z_COPYABLE_P(zval_p __auto__) bool   { return Z_COPYABLE(*zval_p) }
func Z_IMMUTABLE(zval Zval) bool          { return zval.GetTypeInfo() == IS_ARRAY }
func Z_IMMUTABLE_P(zval_p *Zval) bool     { return Z_IMMUTABLE(*zval_p) }
func Z_OPT_IMMUTABLE(zval __auto__) bool  { return Z_IMMUTABLE(zval_p) }
func Z_OPT_IMMUTABLE_P(zval_p *Zval) bool { return Z_IMMUTABLE(*zval_p) }
func Z_OPT_TYPE(zval Zval) int            { return zval.GetTypeInfo() & Z_TYPE_MASK }
func Z_OPT_TYPE_P(zval_p *Zval) int       { return Z_OPT_TYPE(*zval_p) }
func Z_OPT_CONSTANT(zval Zval) bool {
	return Z_OPT_TYPE(zval) == IS_CONSTANT_AST
}
func Z_OPT_CONSTANT_P(zval_p *Zval) bool { return Z_OPT_CONSTANT(*zval_p) }
func Z_OPT_REFCOUNTED(zval Zval) bool {
	return Z_TYPE_INFO_REFCOUNTED(zval.GetTypeInfo())
}
func Z_OPT_REFCOUNTED_P(zval_p *Zval) bool        { return Z_OPT_REFCOUNTED(*zval_p) }
func Z_OPT_COPYABLE(zval Zval) bool               { return Z_OPT_TYPE(zval) == IS_ARRAY }
func Z_OPT_COPYABLE_P(zval_p *Zval) bool          { return Z_OPT_COPYABLE(*zval_p) }
func Z_OPT_ISREF(zval Zval) bool                  { return Z_OPT_TYPE(zval) == IS_REFERENCE }
func Z_OPT_ISREF_P(zval_p *Zval) bool             { return Z_OPT_ISREF(*zval_p) }
func Z_ISREF(zval Zval) bool                      { return zval.IsType(IS_REFERENCE) }
func Z_ISREF_P(zval_p *Zval) bool                 { return Z_ISREF(*zval_p) }
func Z_ISUNDEF(zval Zval) bool                    { return zval.IsType(IS_UNDEF) }
func Z_ISUNDEF_P(zval_p *Zval) bool               { return Z_ISUNDEF(*zval_p) }
func Z_ISNULL(zval __auto__) bool                 { return zval.u1.v.type_ == IS_NULL }
func Z_ISNULL_P(zval_p __auto__) bool             { return Z_ISNULL(*zval_p) }
func Z_ISERROR(zval Zval) bool                    { return zval.IsType(_IS_ERROR) }
func Z_ISERROR_P(zval_p *Zval) bool               { return Z_ISERROR(*zval_p) }
func Z_LVAL(zval Zval) ZendLong                   { return zval.GetLval() }
func Z_LVAL_P(zval_p *Zval) ZendLong              { return zval_p.GetLval() }
func Z_DVAL(zval Zval) float64                    { return zval.GetDval() }
func Z_DVAL_P(zval_p *Zval) float64               { return zval_p.GetDval() }
func Z_STR(zval Zval) *ZendString                 { return zval.GetStr() }
func Z_STR_P(zval_p *Zval) *ZendString            { return zval_p.GetStr() }
func Z_STRVAL(zval Zval) []byte                   { return Z_STR(zval).GetVal() }
func Z_STRVAL_P(zval_p *Zval) []byte              { return Z_STRVAL(*zval_p) }
func Z_STRLEN(zval Zval) int                      { return Z_STR(zval).GetLen() }
func Z_STRLEN_P(zval_p *Zval) int                 { return Z_STRLEN(*zval_p) }
func Z_STRHASH(zval Zval) ZendUlong               { return ZSTR_HASH(zval.GetStr()) }
func Z_STRHASH_P(zval_p *Zval) ZendUlong          { return Z_STRHASH(*zval_p) }
func Z_ARR(zval Zval) *ZendArray                  { return zval.GetArr() }
func Z_ARR_P(zval_p *Zval) *ZendArray             { return zval_p.GetArr() }
func Z_ARRVAL(zval Zval) *ZendArray               { return zval.GetArr() }
func Z_ARRVAL_P(zval_p *Zval) *ZendArray          { return zval_p.GetArr() }
func Z_OBJ(zval Zval) *ZendObject                 { return zval.GetObj() }
func Z_OBJ_P(zval_p *Zval) *ZendObject            { return zval_p.GetObj() }
func Z_OBJ_HT(zval Zval) *ZendObjectHandlers      { return Z_OBJ(zval).GetHandlers() }
func Z_OBJ_HT_P(zval_p *Zval) *ZendObjectHandlers { return Z_OBJ_HT(*zval_p) }
func Z_OBJ_HANDLE(zval Zval) uint32               { return Z_OBJ(zval).GetHandle() }
func Z_OBJ_HANDLE_P(zval_p *Zval) uint32          { return Z_OBJ_HANDLE(*zval_p) }
func Z_OBJCE(zval Zval) *ZendClassEntry           { return Z_OBJ(zval).GetCe() }
func Z_OBJCE_P(zval_p *Zval) *ZendClassEntry      { return Z_OBJCE(*zval_p) }
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
func Z_RES_VAL(zval Zval) any             { return Z_RES(zval).GetPtr() }
func Z_RES_VAL_P(zval_p *Zval) any        { return Z_RES_VAL(*zval_p) }
func Z_REF(zval Zval) *ZendReference      { return zval.GetRef() }
func Z_REF_P(zval_p *Zval) *ZendReference { return zval_p.GetRef() }
func Z_REFVAL(zval Zval) Zval             { return Z_REF(zval).GetVal() }
func Z_REFVAL_P(zval_p *Zval) Zval        { return Z_REFVAL(*zval_p) }
func Z_AST(zval Zval) *ZendAstRef         { return zval.GetAst() }
func Z_AST_P(zval_p *Zval) *ZendAstRef    { return zval_p.GetAst() }
func GC_AST(p *ZendAstRef) *ZendAst {
	return (*ZendAst)((*byte)(p) + b.SizeOf("zend_ast_ref"))
}
func Z_ASTVAL(zval Zval) *ZendAst         { return GC_AST(zval.GetAst()) }
func Z_ASTVAL_P(zval_p *Zval) *ZendAst    { return Z_ASTVAL(*zval_p) }
func Z_INDIRECT(zval Zval) *Zval          { return zval.GetZv() }
func Z_INDIRECT_P(zval_p *Zval) *Zval     { return zval_p.GetZv() }
func Z_CE(zval Zval) *ZendClassEntry      { return zval.GetCe() }
func Z_CE_P(zval_p *Zval) *ZendClassEntry { return zval_p.GetCe() }
func Z_FUNC(zval Zval) *ZendFunction      { return zval.GetFunc() }
func Z_FUNC_P(zval_p *Zval) *ZendFunction { return zval_p.GetFunc() }
func Z_PTR(zval Zval) any                 { return zval.GetPtr() }
func Z_PTR_P(zval_p *Zval) any            { return zval_p.GetPtr() }
func ZVAL_UNDEF(z *Zval)                  { z.SetTypeInfo(IS_UNDEF) }
func ZVAL_NULL(z *Zval)                   { z.SetTypeInfo(IS_NULL) }
func ZVAL_FALSE(z *Zval)                  { z.SetTypeInfo(IS_FALSE) }
func ZVAL_TRUE(z *Zval)                   { z.SetTypeInfo(IS_TRUE) }
func ZVAL_BOOL(z *Zval, b int) {
	if b != 0 {
		z.SetTypeInfo(IS_TRUE)
	} else {
		z.SetTypeInfo(IS_FALSE)
	}
}
func ZVAL_LONG(z *Zval, l ZendLong) {
	var __z *Zval = z
	__z.SetLval(l)
	__z.SetTypeInfo(IS_LONG)
}
func ZVAL_DOUBLE(z *Zval, d float64) {
	var __z *Zval = z
	__z.SetDval(d)
	__z.SetTypeInfo(IS_DOUBLE)
}
func ZVAL_STR(z *Zval, s *ZendString) {
	var __z *Zval = z
	var __s *ZendString = s
	__z.SetStr(__s)
	__z.SetTypeInfo(IS_STRING_EX)
}
func ZVAL_INTERNED_STR(z *Zval, s *ZendString) {
	var __z *Zval = z
	var __s *ZendString = s
	__z.SetStr(__s)
	__z.SetTypeInfo(IS_INTERNED_STRING_EX)
}
func ZVAL_NEW_STR(z *Zval, s *ZendString) {
	var __z *Zval = z
	var __s *ZendString = s
	__z.SetStr(__s)
	__z.SetTypeInfo(IS_STRING_EX)
}
func ZVAL_STR_COPY(z *Zval, s *ZendString) {
	var __z *Zval = z
	var __s *ZendString = s
	__z.SetStr(__s)
	__s.AddRefcount()
	__z.SetTypeInfo(IS_STRING_EX)
}
func ZVAL_ARR(z *Zval, a *ZendArray) {
	var __arr *ZendArray = a
	var __z *Zval = z
	__z.SetArr(__arr)
	__z.SetTypeInfo(IS_ARRAY_EX)
}
func ZVAL_NEW_ARR(z *Zval) {
	var __z *Zval = z
	var _arr *ZendArray = (*ZendArray)(Emalloc(b.SizeOf("zend_array")))
	__z.SetArr(_arr)
	__z.SetTypeInfo(IS_ARRAY_EX)
}
func ZVAL_NEW_PERSISTENT_ARR(z *Zval) {
	var __z *Zval = z
	var _arr *ZendArray = (*ZendArray)(Malloc(b.SizeOf("zend_array")))
	__z.SetArr(_arr)
	__z.SetTypeInfo(IS_ARRAY_EX)
}
func ZVAL_OBJ(z *Zval, o *ZendObject) {
	var __z *Zval = z
	__z.SetObj(o)
	__z.SetTypeInfo(IS_OBJECT_EX)
}
func ZVAL_RES(z *Zval, r *ZendResource) {
	var __z *Zval = z
	__z.SetRes(r)
	__z.SetTypeInfo(IS_RESOURCE_EX)
}
func ZVAL_NEW_RES(z *Zval, h int, p any, t int) {
	var _res *ZendResource = (*ZendResource)(Emalloc(b.SizeOf("zend_resource")))
	var __z *Zval
	_res.SetRefcount(1)
	_res.GetGcTypeInfo() = IS_RESOURCE
	_res.SetHandle(h)
	_res.SetType(t)
	_res.SetPtr(p)
	__z = z
	__z.SetRes(_res)
	__z.SetTypeInfo(IS_RESOURCE_EX)
}
func ZVAL_NEW_PERSISTENT_RES(z *Zval, h int, p any, t int) {
	var _res *ZendResource = (*ZendResource)(Malloc(b.SizeOf("zend_resource")))
	var __z *Zval
	_res.SetRefcount(1)
	_res.GetGcTypeInfo() = IS_RESOURCE | GC_PERSISTENT<<GC_FLAGS_SHIFT
	_res.SetHandle(h)
	_res.SetType(t)
	_res.SetPtr(p)
	__z = z
	__z.SetRes(_res)
	__z.SetTypeInfo(IS_RESOURCE_EX)
}
func ZVAL_REF(z *Zval, r *ZendReference) {
	var __z *Zval = z
	__z.SetRef(r)
	__z.SetTypeInfo(IS_REFERENCE_EX)
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
func ZVAL_NEW_PERSISTENT_REF(z *Zval, r *Zval) {
	var _ref *ZendReference = (*ZendReference)(Malloc(b.SizeOf("zend_reference")))
	_ref.SetRefcount(1)
	_ref.GetGcTypeInfo() = IS_REFERENCE | GC_PERSISTENT<<GC_FLAGS_SHIFT
	ZVAL_COPY_VALUE(_ref.GetVal(), r)
	_ref.GetSources().SetPtr(nil)
	z.SetRef(_ref)
	z.SetTypeInfo(IS_REFERENCE_EX)
}
func ZVAL_AST(z *Zval, ast *ZendAstRef) {
	var __z *Zval = z
	__z.SetAst(ast)
	__z.SetTypeInfo(IS_CONSTANT_AST_EX)
}
func ZVAL_INDIRECT(z *Zval, v *Zval) {
	z.SetZv(v)
	z.SetTypeInfo(IS_INDIRECT)
}
func ZVAL_PTR(z *Zval, p any) {
	z.SetPtr(p)
	z.SetTypeInfo(IS_PTR)
}
func ZVAL_FUNC(z *Zval, f *ZendFunction) {
	z.SetFunc(f)
	z.SetTypeInfo(IS_PTR)
}
func ZVAL_CE(z *Zval, c *ZendClassEntry) {
	z.SetCe(c)
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
func Z_TRY_ADDREF(z Zval)                      { Z_TRY_ADDREF_P(&z) }
func Z_TRY_DELREF(z Zval)                      { Z_TRY_DELREF_P(&z) }
func ZEND_RC_MOD_CHECK(p *ZendRefcountedH)     {}
func GC_MAKE_PERSISTENT_LOCAL(p *ZendString)   {}
func ZendGcRefcount(p *ZendRefcountedH) uint32 { return p.GetRefcount() }
func ZendGcSetRefcount(p *ZendRefcountedH, rc uint32) uint32 {
	p.SetRefcount(rc)
	return p.GetRefcount()
}
func ZendGcAddref(p *ZendRefcountedH) uint32 {
	p.GetRefcount()++
	return p.GetRefcount()
}
func ZendGcDelref(p *ZendRefcountedH) uint32 {
	p.GetRefcount()--
	return p.GetRefcount()
}
func ZendGcAddrefEx(p *ZendRefcountedH, rc uint32) uint32 {
	p.SetRefcount(p.GetRefcount() + rc)
	return p.GetRefcount()
}
func ZendGcDelrefEx(p *ZendRefcountedH, rc uint32) uint32 {
	p.SetRefcount(p.GetRefcount() - rc)
	return p.GetRefcount()
}
func ZvalRefcountP(pz *Zval) uint32 { return pz.GetCounted().GetRefcount() }
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
	var _z1 *Zval = z
	var _z2 *Zval = v
	var _gc *ZendRefcounted = _z2.GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	ZVAL_COPY_VALUE_EX(_z1, _z2, _gc, _t)
}
func ZVAL_COPY(z *Zval, v *Zval) {
	var _z1 *Zval = z
	var _z2 *Zval = v
	var _gc *ZendRefcounted = _z2.GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	ZVAL_COPY_VALUE_EX(_z1, _z2, _gc, _t)
	if Z_TYPE_INFO_REFCOUNTED(_t) {
		_gc.AddRefcount()
	}
}
func ZVAL_DUP(z *Zval, v *Zval) {
	var _z1 *Zval = z
	var _z2 *Zval = v
	var _gc *ZendRefcounted = _z2.GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	if (_t & Z_TYPE_MASK) == IS_ARRAY {
		ZVAL_ARR(_z1, ZendArrayDup((*ZendArray)(_gc)))
	} else {
		if Z_TYPE_INFO_REFCOUNTED(_t) {
			_gc.AddRefcount()
		}
		ZVAL_COPY_VALUE_EX(_z1, _z2, _gc, _t)
	}
}
func ZVAL_COPY_OR_DUP(z *Zval, v *Zval) {
	var _z1 *Zval = z
	var _z2 *Zval = v
	var _gc *ZendRefcounted = _z2.GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	ZVAL_COPY_VALUE_EX(_z1, _z2, _gc, _t)
	if Z_TYPE_INFO_REFCOUNTED(_t) {
		if (_gc.GetGcFlags() & GC_PERSISTENT) == 0 {
			_gc.AddRefcount()
		} else {
			ZvalCopyCtorFunc(_z1)
		}
	}
}
func ZVAL_DEREF(z *Zval) {
	if Z_ISREF_P(z) {
		z = Z_REFVAL_P(z)
	}
}
func ZVAL_DEINDIRECT(z *Zval) {
	if z.IsType(IS_INDIRECT) {
		z = z.GetZv()
	}
}
func ZVAL_OPT_DEREF(z *Zval) {
	if Z_OPT_ISREF_P(z) {
		z = Z_REFVAL_P(z)
	}
}
func ZVAL_MAKE_REF(zv *Zval) {
	var __zv *Zval = zv
	if !(Z_ISREF_P(__zv)) {
		ZVAL_NEW_REF(__zv, __zv)
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
func SEPARATE_STRING(zv *Zval) {
	var _zv *Zval = zv
	if Z_REFCOUNT_P(_zv) > 1 {
		var _str *ZendString = _zv.GetStr()
		ZEND_ASSERT(Z_REFCOUNTED_P(_zv))
		ZEND_ASSERT(true)
		Z_DELREF_P(_zv)
		ZVAL_NEW_STR(_zv, ZendStringInit(_str.GetVal(), _str.GetLen(), 0))
	}
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
	var __zv *Zval = zv
	if __zv.IsType(IS_ARRAY) {
		if Z_REFCOUNT_P(__zv) > 1 {
			if Z_REFCOUNTED_P(__zv) {
				Z_DELREF_P(__zv)
			}
			ZVAL_ARR(__zv, ZendArrayDup(__zv.GetArr()))
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
func Z_PROP_FLAG_P(z *Zval) uint32          { return z.GetU2Extra() }
func ZVAL_COPY_VALUE_PROP(z *Zval, v *Zval) { *z = *v }
func ZVAL_COPY_PROP(z *Zval, v *Zval) {
	ZVAL_COPY(z, v)
	z.SetU2Extra(v.GetU2Extra())
}
func ZVAL_COPY_OR_DUP_PROP(z *Zval, v *Zval) {
	ZVAL_COPY_OR_DUP(z, v)
	z.SetU2Extra(v.GetU2Extra())
}
