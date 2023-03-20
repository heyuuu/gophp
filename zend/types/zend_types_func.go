package types

import (
	b "sik/builtin"
	"sik/zend"
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
func HT_IDX_TO_HASH(idx uint32) uint32 { return idx }
func HT_HASH(ht *HashTable, idx HashPosition) uint32 {
	// todo 待移除 - 在旧 arData 上返回通过 idx 获取对应的 pos 位置
	return 0
}
func HT_HASH_SIZE(nTableMask uint32) int {
	// todo 待移除 - 旧 hash 数组内存大小
	return 0
}
func HT_SIZE(ht *HashTable) int {
	// todo 待移除 - 返回 HashTable 中旧 arData 内存大小 (含 hash 内存 + item 内存)
	return 0
}
func HT_USED_SIZE(ht *HashTable) int {
	return HT_HASH_SIZE(ht.GetNTableMask()) + size_t(ht).nNumUsed*b.SizeOf("Bucket")
}
func HT_HASH_TO_BUCKET(ht *HashTable, idx uint32) *Bucket { return ht.Bucket(idx) }
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
	return faketype == realtype || faketype == IS_BOOL && (realtype == IS_TRUE || realtype == IS_FALSE)
}
func Z_FE_ITER_P(zval_p *Zval) uint32      { return zval_p.GetFeIterIdx() }
func Z_TYPE_INFO_REFCOUNTED(t uint32) bool { return (t & Z_TYPE_FLAGS_MASK) != 0 }

func Z_OPT_REFCOUNTED_P(zval_p *Zval) bool             { return zval_p.IsRefcounted() }
func Z_ISREF_P(zval_p *Zval) bool                      { return zval_p.IsReference() }
func Z_STR_P(zval_p *Zval) *String                     { return zval_p.GetStr() }
func Z_STRVAL_P(zval_p *Zval) []byte                   { return zval_p.GetStr().GetVal() }
func Z_ARRVAL(zval Zval) *Array                        { return zval.GetArr() }
func Z_ARRVAL_P(zval_p *Zval) *Array                   { return zval_p.GetArr() }
func Z_OBJ(zval Zval) *ZendObject                      { return zval.GetObj() }
func Z_OBJ_P(zval_p *Zval) *ZendObject                 { return zval_p.GetObj() }
func Z_OBJ_HT(zval Zval) *zend.ZendObjectHandlers      { return Z_OBJ(zval).GetHandlers() }
func Z_OBJ_HT_P(zval_p *Zval) *zend.ZendObjectHandlers { return Z_OBJ_HT(*zval_p) }
func Z_OBJCE(zval Zval) *ClassEntry                    { return zval.GetObj().GetCe() }
func Z_OBJCE_P(zval_p *Zval) *ClassEntry               { return zval_p.GetObj().GetCe() }
func Z_OBJPROP(zval Zval) *HashTable {
	return Z_OBJ_HT(zval).GetGetProperties()(&zval)
}
func Z_OBJPROP_P(zval_p *Zval) *HashTable   { return Z_OBJPROP(*zval_p) }
func Z_RES(zval Zval) *ZendResource         { return zval.GetRes() }
func Z_RES_P(zval_p *Zval) *ZendResource    { return zval_p.GetRes() }
func Z_RES_HANDLE(zval Zval) int            { return Z_RES(zval).GetHandle() }
func Z_RES_HANDLE_P(zval_p *Zval) int       { return Z_RES_HANDLE(*zval_p) }
func Z_RES_TYPE(zval Zval) int              { return Z_RES(zval).GetType() }
func Z_RES_TYPE_P(zval_p *Zval) int         { return Z_RES_TYPE(*zval_p) }
func Z_REF(zval Zval) *ZendReference        { return zval.GetRef() }
func Z_REF_P(zval_p *Zval) *ZendReference   { return zval_p.GetRef() }
func Z_REFVAL(zval Zval) *Zval              { return zval.GetRef().GetVal() }
func Z_REFVAL_P(zval_p *Zval) *Zval         { return zval_p.GetRef().GetVal() }
func GC_AST(p *ZendAstRef) *zend.ZendAst    { return p.GcAst() }
func Z_ASTVAL(zval Zval) *zend.ZendAst      { return GC_AST(zval.GetAst()) }
func Z_ASTVAL_P(zval_p *Zval) *zend.ZendAst { return Z_ASTVAL(*zval_p) }
func Z_INDIRECT(zval Zval) *Zval            { return zval.GetZv() }
func Z_INDIRECT_P(zval_p *Zval) *Zval       { return zval_p.GetZv() }
func Z_CE(zval Zval) *ClassEntry            { return zval.GetCe() }
func Z_PTR(zval Zval) any                   { return zval.GetPtr() }

func ZVAL_BOOL(z *Zval, b int)             { z.SetBool(b != 0) }
func ZVAL_LONG(z *Zval, l zend.ZendLong)   { z.SetLong(l) }
func ZVAL_STR(z *Zval, s *String)          { z.SetString(s) }
func ZVAL_INTERNED_STR(z *Zval, s *String) { z.SetInternedString(s) }
func ZVAL_NEW_STR(z *Zval, s *String)      { z.SetString(s) }
func ZVAL_STR_COPY(z *Zval, s *String)     { z.SetStringCopy(s) }
func ZVAL_ARR(z *Zval, a *Array)           { z.SetArray(a) }
func ZVAL_NEW_PERSISTENT_ARR(z *Zval) {
	var arr = NewZendArray(0)
	z.SetArray(arr)
}
func ZVAL_OBJ(z *Zval, o *ZendObject)                      { z.SetObject(o) }
func ZVAL_RES(z *Zval, r *ZendResource)                    { z.SetResource(r) }
func ZVAL_NEW_RES(z *Zval, h int, p any, t int)            { z.SetNewResource(h, p, t) }
func ZVAL_NEW_PERSISTENT_RES(z *Zval, h int, p any, t int) { z.SetNewResourcePersistent(h, p, t) }
func ZVAL_REF(z *Zval, r *ZendReference)                   { z.SetReference(r) }
func ZVAL_NEW_EMPTY_REF(z *Zval)                           { z.SetNewEmptyRef() }
func ZVAL_NEW_REF(z *Zval, r *Zval)                        { z.SetNewRef(r) }
func ZVAL_MAKE_REF_EX(z *Zval, refcount uint32) {
	var ref *ZendReference = NewZendReference(z)
	ref.SetRefcount(refcount)
	z.SetReference(ref)
}
func ZVAL_AST(z *Zval, ast *ZendAstRef) { z.SetConstantAst(ast) }
func ZVAL_INDIRECT(z *Zval, v *Zval)    { z.SetIndirect(v) }
func ZVAL_PTR(z *Zval, p any)           { z.SetAsPtr(p) }
func ZVAL_ALIAS_PTR(z *Zval, p *ClassEntry) {
	z.SetPtr(p)
	z.SetTypeInfo(IS_ALIAS_PTR)
}
func Z_REFCOUNT_P(pz *Zval) uint32       { return pz.GetRefcount() }
func Z_ADDREF_P(pz *Zval) uint32         { return pz.AddRefcount() }
func Z_DELREF_P(pz *Zval) uint32         { return pz.DelRefcount() }
func GC_MAKE_PERSISTENT_LOCAL(p *String) {}

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
			zend.ZvalCopyCtorFunc(z)
		}
	}
}

func ZVAL_DEREF(z *Zval) *Zval {
	if z.IsReference() {
		return Z_REFVAL_P(z)
	}
	return z
}
func ZVAL_DEINDIRECT(z *Zval) *Zval {
	if z.IsIndirect() {
		return z.GetZv()
	}
	return z
}
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
	ref = _z.GetRef()
	ZVAL_COPY_VALUE(_z, ref.GetVal())
	zend.EfreeSize(ref, b.SizeOf("zend_reference"))
}
func ZVAL_COPY_DEREF(z *Zval, v *Zval) {
	var _z3 *Zval = v
	if _z3.IsRefcounted() {
		if _z3.IsReference() {
			_z3 = Z_REFVAL_P(_z3)
			if _z3.IsRefcounted() {
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
	var _arr *Array = _zv.GetArr()
	if _arr.GetRefcount() > 1 {
		if _zv.IsRefcounted() {
			_arr.DelRefcount()
		}
		ZVAL_ARR(_zv, ZendArrayDup(_arr))
	}
}
func SEPARATE_ZVAL_IF_NOT_REF(zv *Zval) {
	var __zv *Zval = zv
	if __zv.IsArray() {
		if Z_REFCOUNT_P(__zv) > 1 {
			if __zv.IsRefcounted() {
				Z_DELREF_P(__zv)
			}
			ZVAL_ARR(__zv, ZendArrayDup(__zv.GetArr()))
		}
	}
}
func SEPARATE_ZVAL_NOREF(zv *Zval) {
	var _zv *Zval = zv
	b.Assert(_zv.GetType() != IS_REFERENCE)
	SEPARATE_ZVAL_IF_NOT_REF(_zv)
}
func SEPARATE_ZVAL(zv *Zval) {
	for {
		var _zv *Zval = zv
		if _zv.IsReference() {
			var _r *ZendReference = _zv.GetRef()
			ZVAL_COPY_VALUE(_zv, _r.GetVal())
			if _r.DelRefcount() == 0 {
				zend.EfreeSize(_r, b.SizeOf("zend_reference"))
			} else if _zv.IsArray() {
				ZVAL_ARR(_zv, ZendArrayDup(_zv.GetArr()))
				break
			} else if _zv.IsRefcounted() {
				Z_ADDREF_P(_zv)
				break
			}
		}
		SEPARATE_ZVAL_IF_NOT_REF(_zv)
		break
	}
}
func SEPARATE_ARG_IF_REF(varptr *Zval) {
	varptr = ZVAL_DEREF(varptr)
	if varptr.IsRefcounted() {
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

func ZendGetTypeByConst(type_ ZendUchar) string {
	switch type_ {
	case IS_FALSE, IS_TRUE, IS_BOOL:
		return "bool"
	case IS_LONG:
		return "int"
	case IS_DOUBLE:
		return "float"
	case IS_STRING:
		return "string"
	case IS_OBJECT:
		return "object"
	case IS_RESOURCE:
		return "resource"
	case IS_NULL:
		return "null"
	case IS_CALLABLE:
		return "callable"
	case IS_ITERABLE:
		return "iterable"
	case IS_ARRAY:
		return "array"
	case IS_VOID:
		return "void"
	case IS_NUMBER:
		return "number"
	default:
		return "unknown"
	}
}
func ZendZvalTypeName(arg *Zval) string {
	arg = ZVAL_DEREF(arg)
	return ZendGetTypeByConst(arg.GetType())
}
func ZendZvalGetType(arg *Zval) *String {
	switch arg.GetType() {
	case IS_NULL:
		return ZSTR_NULL
	case IS_FALSE, IS_TRUE:
		return ZSTR_BOOLEAN
	case IS_LONG:
		return ZSTR_INTEGER
	case IS_DOUBLE:
		return ZSTR_DOUBLE
	case IS_STRING:
		return ZSTR_STRING
	case IS_ARRAY:
		return ZSTR_ARRAY
	case IS_OBJECT:
		return ZSTR_OBJECT
	case IS_RESOURCE:
		if zend.ZendRsrcListGetRsrcType(arg.GetRes()) != nil {
			return ZSTR_RESOURCE
		} else {
			return ZSTR_CLOSED_RESOURCE
		}
	default:
		return nil
	}
}
