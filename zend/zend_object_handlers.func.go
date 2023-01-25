// <<generate>>

package zend

import (
	b "sik/builtin"
	"sik/core"
)

func IS_VALID_PROPERTY_OFFSET(offset uintPtr) bool   { return intptr_t(offset) > 0 }
func IS_WRONG_PROPERTY_OFFSET(offset uintPtr) bool   { return intptr_t(offset) == 0 }
func IS_DYNAMIC_PROPERTY_OFFSET(offset uintPtr) bool { return intptr_t(offset) < 0 }
func IS_UNKNOWN_DYNAMIC_PROPERTY_OFFSET(offset uintPtr) bool {
	return offset == ZEND_DYNAMIC_PROPERTY_OFFSET
}
func ZEND_DECODE_DYN_PROP_OFFSET(offset uintPtr) __auto__ {
	return uintptr_t(-(intptr_t(offset)) - 2)
}
func ZEND_ENCODE_DYN_PROP_OFFSET(offset uintPtr) __auto__ {
	return uintptr_t(-(intptr_t(offset) + 2))
}
func ZendGetStdObjectHandlers() *ZendObjectHandlers { return &StdObjectHandlers }
func ZendGetFunctionRootClass(fbc *ZendFunction) *ZendClassEntry {
	if fbc.GetPrototype() != nil {
		return fbc.GetPrototype().GetScope()
	} else {
		return fbc.GetScope()
	}
}
func ZendReleaseProperties(ht *HashTable) {
	if ht != nil && (GC_FLAGS(ht)&GC_IMMUTABLE) == 0 && GC_DELREF(ht) == 0 {
		ZendArrayDestroy(ht)
	}
}
func ZendFreeTrampoline(func_ any) {
	if func_ == &EG(trampoline) {
		EG(trampoline).common.function_name = nil
	} else {
		Efree(func_)
	}
}
func RebuildObjectProperties(zobj *ZendObject) {
	if zobj.GetProperties() == nil {
		var prop_info *ZendPropertyInfo
		var ce *ZendClassEntry = zobj.GetCe()
		var flags uint32 = 0
		zobj.SetProperties(ZendNewArray(ce.GetDefaultPropertiesCount()))
		if ce.GetDefaultPropertiesCount() != 0 {
			ZendHashRealInitMixed(zobj.GetProperties())
			for {
				var __ht *HashTable = &ce.properties_info
				var _p *Bucket = __ht.GetArData()
				var _end *Bucket = _p + __ht.GetNNumUsed()
				for ; _p != _end; _p++ {
					var _z *Zval = &_p.val

					if UNEXPECTED(Z_TYPE_P(_z) == IS_UNDEF) {
						continue
					}
					prop_info = Z_PTR_P(_z)
					if (prop_info.GetFlags() & ZEND_ACC_STATIC) == 0 {
						flags |= prop_info.GetFlags()
						if UNEXPECTED(Z_TYPE_P(OBJ_PROP(zobj, prop_info.GetOffset())) == IS_UNDEF) {
							HT_FLAGS(zobj.GetProperties()) |= HASH_FLAG_HAS_EMPTY_IND
						}
						_zendHashAppendInd(zobj.GetProperties(), prop_info.GetName(), OBJ_PROP(zobj, prop_info.GetOffset()))
					}
				}
				break
			}
			if (flags & ZEND_ACC_CHANGED) != 0 {
				for ce.parent && ce.parent.default_properties_count {
					ce = ce.parent
					for {
						var __ht *HashTable = &ce.properties_info
						var _p *Bucket = __ht.GetArData()
						var _end *Bucket = _p + __ht.GetNNumUsed()
						for ; _p != _end; _p++ {
							var _z *Zval = &_p.val

							if UNEXPECTED(Z_TYPE_P(_z) == IS_UNDEF) {
								continue
							}
							prop_info = Z_PTR_P(_z)
							if prop_info.GetCe() == ce && (prop_info.GetFlags()&ZEND_ACC_STATIC) == 0 && (prop_info.GetFlags()&ZEND_ACC_PRIVATE) != 0 {
								var zv Zval
								if UNEXPECTED(Z_TYPE_P(OBJ_PROP(zobj, prop_info.GetOffset())) == IS_UNDEF) {
									HT_FLAGS(zobj.GetProperties()) |= HASH_FLAG_HAS_EMPTY_IND
								}
								ZVAL_INDIRECT(&zv, OBJ_PROP(zobj, prop_info.GetOffset()))
								ZendHashAdd(zobj.GetProperties(), prop_info.GetName(), &zv)
							}
						}
						break
					}
				}
			}
		}
	}
}
func ZendStdGetProperties(object *Zval) *HashTable {
	var zobj *ZendObject
	zobj = Z_OBJ_P(object)
	if zobj.GetProperties() == nil {
		RebuildObjectProperties(zobj)
	}
	return zobj.GetProperties()
}
func ZendStdGetGc(object *Zval, table **Zval, n *int) *HashTable {
	if Z_OBJ_HANDLER_P(object, get_properties) != ZendStdGetProperties {
		*table = nil
		*n = 0
		return Z_OBJ_HANDLER_P(object, get_properties)(object)
	} else {
		var zobj *ZendObject = Z_OBJ_P(object)
		if zobj.GetProperties() != nil {
			*table = nil
			*n = 0
			if UNEXPECTED(GC_REFCOUNT(zobj.GetProperties()) > 1) && EXPECTED((GC_FLAGS(zobj.GetProperties())&IS_ARRAY_IMMUTABLE) == 0) {
				GC_DELREF(zobj.GetProperties())
				zobj.SetProperties(ZendArrayDup(zobj.GetProperties()))
			}
			return zobj.GetProperties()
		} else {
			*table = zobj.GetPropertiesTable()
			*n = zobj.GetCe().GetDefaultPropertiesCount()
			return nil
		}
	}
}
func ZendStdGetDebugInfo(object *Zval, is_temp *int) *HashTable {
	var ce *ZendClassEntry = Z_OBJCE_P(object)
	var retval Zval
	var ht *HashTable
	if ce.GetDebugInfo() == nil {
		*is_temp = 0
		return Z_OBJ_HANDLER_P(object, get_properties)(object)
	}
	ZendCallMethodWith0Params(object, ce, &ce.__debugInfo, ZEND_DEBUGINFO_FUNC_NAME, &retval)
	if Z_TYPE(retval) == IS_ARRAY {
		if !(Z_REFCOUNTED(retval)) {
			*is_temp = 1
			return ZendArrayDup(Z_ARRVAL(retval))
		} else if Z_REFCOUNT(retval) <= 1 {
			*is_temp = 1
			ht = Z_ARR(retval)
			return ht
		} else {
			*is_temp = 0
			ZvalPtrDtor(&retval)
			return Z_ARRVAL(retval)
		}
	} else if Z_TYPE(retval) == IS_NULL {
		*is_temp = 1
		ht = ZendNewArray(0)
		return ht
	}
	ZendErrorNoreturn(E_ERROR, ZEND_DEBUGINFO_FUNC_NAME+"() must return an array")
	return nil
}
func ZendStdCallGetter(zobj *ZendObject, prop_name *ZendString, retval *Zval) {
	var ce *ZendClassEntry = zobj.GetCe()
	var orig_fake_scope *ZendClassEntry = ExecutorGlobals.GetFakeScope()
	var fci ZendFcallInfo
	var fcic ZendFcallInfoCache
	var member Zval
	ExecutorGlobals.SetFakeScope(nil)

	/* __get handler is called with one argument:
	      property name

	   it should return whether the call was successful or not
	*/

	ZVAL_STR(&member, prop_name)
	fci.SetSize(b.SizeOf("fci"))
	fci.SetObject(zobj)
	fci.SetRetval(retval)
	fci.SetParamCount(1)
	fci.SetParams(&member)
	fci.SetNoSeparation(1)
	ZVAL_UNDEF(&fci.function_name)
	fcic.SetFunctionHandler(ce.GetGet())
	fcic.SetCalledScope(ce)
	fcic.SetObject(zobj)
	ZendCallFunction(&fci, &fcic)
	ExecutorGlobals.SetFakeScope(orig_fake_scope)
}
func ZendStdCallSetter(zobj *ZendObject, prop_name *ZendString, value *Zval) {
	var ce *ZendClassEntry = zobj.GetCe()
	var orig_fake_scope *ZendClassEntry = ExecutorGlobals.GetFakeScope()
	var fci ZendFcallInfo
	var fcic ZendFcallInfoCache
	var args []Zval
	var ret Zval
	ExecutorGlobals.SetFakeScope(nil)

	/* __set handler is called with two arguments:
	   property name
	   value to be set
	*/

	ZVAL_STR(&args[0], prop_name)
	ZVAL_COPY_VALUE(&args[1], value)
	ZVAL_UNDEF(&ret)
	fci.SetSize(b.SizeOf("fci"))
	fci.SetObject(zobj)
	fci.SetRetval(&ret)
	fci.SetParamCount(2)
	fci.SetParams(args)
	fci.SetNoSeparation(1)
	ZVAL_UNDEF(&fci.function_name)
	fcic.SetFunctionHandler(ce.GetSet())
	fcic.SetCalledScope(ce)
	fcic.SetObject(zobj)
	ZendCallFunction(&fci, &fcic)
	ZvalPtrDtor(&ret)
	ExecutorGlobals.SetFakeScope(orig_fake_scope)
}
func ZendStdCallUnsetter(zobj *ZendObject, prop_name *ZendString) {
	var ce *ZendClassEntry = zobj.GetCe()
	var orig_fake_scope *ZendClassEntry = ExecutorGlobals.GetFakeScope()
	var fci ZendFcallInfo
	var fcic ZendFcallInfoCache
	var ret Zval
	var member Zval
	ExecutorGlobals.SetFakeScope(nil)

	/* __unset handler is called with one argument:
	   property name
	*/

	ZVAL_STR(&member, prop_name)
	ZVAL_UNDEF(&ret)
	fci.SetSize(b.SizeOf("fci"))
	fci.SetObject(zobj)
	fci.SetRetval(&ret)
	fci.SetParamCount(1)
	fci.SetParams(&member)
	fci.SetNoSeparation(1)
	ZVAL_UNDEF(&fci.function_name)
	fcic.SetFunctionHandler(ce.GetUnset())
	fcic.SetCalledScope(ce)
	fcic.SetObject(zobj)
	ZendCallFunction(&fci, &fcic)
	ZvalPtrDtor(&ret)
	ExecutorGlobals.SetFakeScope(orig_fake_scope)
}
func ZendStdCallIssetter(zobj *ZendObject, prop_name *ZendString, retval *Zval) {
	var ce *ZendClassEntry = zobj.GetCe()
	var orig_fake_scope *ZendClassEntry = ExecutorGlobals.GetFakeScope()
	var fci ZendFcallInfo
	var fcic ZendFcallInfoCache
	var member Zval
	ExecutorGlobals.SetFakeScope(nil)

	/* __isset handler is called with one argument:
	      property name

	   it should return whether the property is set or not
	*/

	ZVAL_STR(&member, prop_name)
	fci.SetSize(b.SizeOf("fci"))
	fci.SetObject(zobj)
	fci.SetRetval(retval)
	fci.SetParamCount(1)
	fci.SetParams(&member)
	fci.SetNoSeparation(1)
	ZVAL_UNDEF(&fci.function_name)
	fcic.SetFunctionHandler(ce.GetIsset())
	fcic.SetCalledScope(ce)
	fcic.SetObject(zobj)
	ZendCallFunction(&fci, &fcic)
	ExecutorGlobals.SetFakeScope(orig_fake_scope)
}
func IsDerivedClass(child_class *ZendClassEntry, parent_class *ZendClassEntry) ZendBool {
	child_class = child_class.parent
	for child_class != nil {
		if child_class == parent_class {
			return 1
		}
		child_class = child_class.parent
	}
	return 0
}
func IsProtectedCompatibleScope(ce *ZendClassEntry, scope *ZendClassEntry) int {
	return scope != nil && (IsDerivedClass(ce, scope) != 0 || IsDerivedClass(scope, ce) != 0)
}
func ZendGetParentPrivateProperty(scope *ZendClassEntry, ce *ZendClassEntry, member *ZendString) *ZendPropertyInfo {
	var zv *Zval
	var prop_info *ZendPropertyInfo
	if scope != ce && scope != nil && IsDerivedClass(ce, scope) != 0 {
		zv = ZendHashFind(&scope.properties_info, member)
		if zv != nil {
			prop_info = (*ZendPropertyInfo)(Z_PTR_P(zv))
			if (prop_info.GetFlags()&ZEND_ACC_PRIVATE) != 0 && prop_info.GetCe() == scope {
				return prop_info
			}
		}
	}
	return nil
}
func ZendBadPropertyAccess(property_info *ZendPropertyInfo, ce *ZendClassEntry, member *ZendString) {
	ZendThrowError(nil, "Cannot access %s property %s::$%s", ZendVisibilityString(property_info.GetFlags()), ZSTR_VAL(ce.GetName()), ZSTR_VAL(member))
}
func ZendBadPropertyName() {
	ZendThrowError(nil, "Cannot access property started with '\\0'")
}
func ZendGetPropertyOffset(ce *ZendClassEntry, member *ZendString, silent int, cache_slot *any, info_ptr **ZendPropertyInfo) uintPtr {
	var zv *Zval
	var property_info *ZendPropertyInfo
	var flags uint32
	var scope *ZendClassEntry
	var offset uintPtr
	if cache_slot != nil && EXPECTED(ce == CACHED_PTR_EX(cache_slot)) {
		*info_ptr = CACHED_PTR_EX(cache_slot + 2)
		return uintPtr(CACHED_PTR_EX(cache_slot + 1))
	}
	if UNEXPECTED(ZendHashNumElements(&ce.properties_info) == 0) || UNEXPECTED(b.Assign(&zv, ZendHashFind(&ce.properties_info, member)) == nil) {
		if UNEXPECTED(ZSTR_VAL(member)[0] == '0') && ZSTR_LEN(member) != 0 {
			if silent == 0 {
				ZendBadPropertyName()
			}
			return ZEND_WRONG_PROPERTY_OFFSET
		}
	dynamic:
		if cache_slot != nil {
			CACHE_POLYMORPHIC_PTR_EX(cache_slot, ce, any(ZEND_DYNAMIC_PROPERTY_OFFSET))
			CACHE_PTR_EX(cache_slot+2, nil)
		}
		return ZEND_DYNAMIC_PROPERTY_OFFSET
	}
	property_info = (*ZendPropertyInfo)(Z_PTR_P(zv))
	flags = property_info.GetFlags()
	if (flags & (ZEND_ACC_CHANGED | ZEND_ACC_PRIVATE | ZEND_ACC_PROTECTED)) != 0 {
		if UNEXPECTED(ExecutorGlobals.GetFakeScope() != nil) {
			scope = ExecutorGlobals.GetFakeScope()
		} else {
			scope = ZendGetExecutedScope()
		}
		if property_info.GetCe() != scope {
			if (flags & ZEND_ACC_CHANGED) != 0 {
				var p *ZendPropertyInfo = ZendGetParentPrivateProperty(scope, ce, member)

				/* If there is a public/protected instance property on ce, don't try to use a
				 * private static property on scope. If both are static, prefer the static
				 * property on scope. This will throw a static property notice, rather than
				 * a visibility error. */

				if p != nil && ((p.GetFlags()&ZEND_ACC_STATIC) == 0 || (flags&ZEND_ACC_STATIC) != 0) {
					property_info = p
					flags = property_info.GetFlags()
					goto found
				} else if (flags & ZEND_ACC_PUBLIC) != 0 {
					goto found
				}

				/* If there is a public/protected instance property on ce, don't try to use a
				 * private static property on scope. If both are static, prefer the static
				 * property on scope. This will throw a static property notice, rather than
				 * a visibility error. */

			}
			if (flags & ZEND_ACC_PRIVATE) != 0 {
				if property_info.GetCe() != ce {
					goto dynamic
				} else {
				wrong:

					/* Information was available, but we were denied access.  Error out. */

					if silent == 0 {
						ZendBadPropertyAccess(property_info, ce, member)
					}
					return ZEND_WRONG_PROPERTY_OFFSET
				}
			} else {
				ZEND_ASSERT((flags & ZEND_ACC_PROTECTED) != 0)
				if UNEXPECTED(IsProtectedCompatibleScope(property_info.GetCe(), scope) == 0) {
					goto wrong
				}
			}
		}
	}
found:
	if UNEXPECTED((flags & ZEND_ACC_STATIC) != 0) {
		if silent == 0 {
			ZendError(E_NOTICE, "Accessing static property %s::$%s as non static", ZSTR_VAL(ce.GetName()), ZSTR_VAL(member))
		}
		return ZEND_DYNAMIC_PROPERTY_OFFSET
	}
	offset = property_info.GetOffset()
	if EXPECTED(property_info.GetType() == 0) {
		property_info = nil
	} else {
		*info_ptr = property_info
	}
	if cache_slot != nil {
		CACHE_POLYMORPHIC_PTR_EX(cache_slot, ce, any(uintPtr(offset)))
		CACHE_PTR_EX(cache_slot+2, property_info)
	}
	return offset
}
func ZendWrongOffset(ce *ZendClassEntry, member *ZendString) {
	var dummy *ZendPropertyInfo

	/* Trigger the correct error */

	ZendGetPropertyOffset(ce, member, 0, nil, &dummy)

	/* Trigger the correct error */
}
func ZendGetPropertyInfo(ce *ZendClassEntry, member *ZendString, silent int) *ZendPropertyInfo {
	var zv *Zval
	var property_info *ZendPropertyInfo
	var flags uint32
	var scope *ZendClassEntry
	if UNEXPECTED(ZendHashNumElements(&ce.properties_info) == 0) || EXPECTED(b.Assign(&zv, ZendHashFind(&ce.properties_info, member)) == nil) {
		if UNEXPECTED(ZSTR_VAL(member)[0] == '0') && ZSTR_LEN(member) != 0 {
			if silent == 0 {
				ZendBadPropertyName()
			}
			return ZEND_WRONG_PROPERTY_INFO
		}
	dynamic:
		return nil
	}
	property_info = (*ZendPropertyInfo)(Z_PTR_P(zv))
	flags = property_info.GetFlags()
	if (flags & (ZEND_ACC_CHANGED | ZEND_ACC_PRIVATE | ZEND_ACC_PROTECTED)) != 0 {
		if UNEXPECTED(ExecutorGlobals.GetFakeScope() != nil) {
			scope = ExecutorGlobals.GetFakeScope()
		} else {
			scope = ZendGetExecutedScope()
		}
		if property_info.GetCe() != scope {
			if (flags & ZEND_ACC_CHANGED) != 0 {
				var p *ZendPropertyInfo = ZendGetParentPrivateProperty(scope, ce, member)
				if p != nil {
					property_info = p
					flags = property_info.GetFlags()
					goto found
				} else if (flags & ZEND_ACC_PUBLIC) != 0 {
					goto found
				}
			}
			if (flags & ZEND_ACC_PRIVATE) != 0 {
				if property_info.GetCe() != ce {
					goto dynamic
				} else {
				wrong:

					/* Information was available, but we were denied access.  Error out. */

					if silent == 0 {
						ZendBadPropertyAccess(property_info, ce, member)
					}
					return ZEND_WRONG_PROPERTY_INFO
				}
			} else {
				ZEND_ASSERT((flags & ZEND_ACC_PROTECTED) != 0)
				if UNEXPECTED(IsProtectedCompatibleScope(property_info.GetCe(), scope) == 0) {
					goto wrong
				}
			}
		}
	}
found:
	if UNEXPECTED((flags & ZEND_ACC_STATIC) != 0) {
		if silent == 0 {
			ZendError(E_NOTICE, "Accessing static property %s::$%s as non static", ZSTR_VAL(ce.GetName()), ZSTR_VAL(member))
		}
	}
	return property_info
}
func ZendCheckPropertyAccess(zobj *ZendObject, prop_info_name *ZendString, is_dynamic ZendBool) int {
	var property_info *ZendPropertyInfo
	var class_name *byte = nil
	var prop_name *byte
	var member *ZendString
	var prop_name_len int
	if ZSTR_VAL(prop_info_name)[0] == 0 {
		if is_dynamic != 0 {
			return SUCCESS
		}
		ZendUnmanglePropertyNameEx(prop_info_name, &class_name, &prop_name, &prop_name_len)
		member = ZendStringInit(prop_name, prop_name_len, 0)
		property_info = ZendGetPropertyInfo(zobj.GetCe(), member, 1)
		ZendStringReleaseEx(member, 0)
		if property_info == nil || property_info == ZEND_WRONG_PROPERTY_INFO {
			return FAILURE
		}
		if class_name[0] != '*' {
			if (property_info.GetFlags() & ZEND_ACC_PRIVATE) == 0 {

				/* we we're looking for a private prop but found a non private one of the same name */

				return FAILURE

				/* we we're looking for a private prop but found a non private one of the same name */

			} else if strcmp(ZSTR_VAL(prop_info_name)+1, ZSTR_VAL(property_info.GetName())+1) {

				/* we we're looking for a private prop but found a private one of the same name but another class */

				return FAILURE

				/* we we're looking for a private prop but found a private one of the same name but another class */

			}
		} else {
			ZEND_ASSERT((property_info.GetFlags() & ZEND_ACC_PROTECTED) != 0)
		}
		return SUCCESS
	} else {
		property_info = ZendGetPropertyInfo(zobj.GetCe(), prop_info_name, 1)
		if property_info == nil {
			ZEND_ASSERT(is_dynamic != 0)
			return SUCCESS
		} else if property_info == ZEND_WRONG_PROPERTY_INFO {
			return FAILURE
		}
		if (property_info.GetFlags() & ZEND_ACC_PUBLIC) != 0 {
			return SUCCESS
		} else {
			return FAILURE
		}
	}
}
func ZendPropertyGuardDtor(el *Zval) {
	var ptr *uint32 = (*uint32)(Z_PTR_P(el))
	if EXPECTED((ZendUintptrT(ptr) & 1) == 0) {
		EfreeSize(ptr, b.SizeOf("uint32_t"))
	}
}
func ZendGetPropertyGuard(zobj *ZendObject, member *ZendString) *uint32 {
	var guards *HashTable
	var zv *Zval
	var ptr *uint32
	ZEND_ASSERT((zobj.GetCe().GetCeFlags() & ZEND_ACC_USE_GUARDS) != 0)
	zv = zobj.GetPropertiesTable() + zobj.GetCe().GetDefaultPropertiesCount()
	if EXPECTED(Z_TYPE_P(zv) == IS_STRING) {
		var str *ZendString = Z_STR_P(zv)
		if EXPECTED(str == member) || EXPECTED(ZSTR_H(str) == ZendStringHashVal(member)) && EXPECTED(ZendStringEqualContent(str, member) != 0) {
			return &Z_PROPERTY_GUARD_P(zv)
		} else if EXPECTED(Z_PROPERTY_GUARD_P(zv) == 0) {
			ZvalPtrDtorStr(zv)
			ZVAL_STR_COPY(zv, member)
			return &Z_PROPERTY_GUARD_P(zv)
		} else {
			ALLOC_HASHTABLE(guards)
			ZendHashInit(guards, 8, nil, ZendPropertyGuardDtor, 0)

			/* mark pointer as "special" using low bit */

			ZendHashAddNewPtr(guards, str, any(zend_uintptr_t&Z_PROPERTY_GUARD_P(zv)|1))
			ZvalPtrDtorStr(zv)
			ZVAL_ARR(zv, guards)
		}
	} else if EXPECTED(Z_TYPE_P(zv) == IS_ARRAY) {
		guards = Z_ARRVAL_P(zv)
		ZEND_ASSERT(guards != nil)
		zv = ZendHashFind(guards, member)
		if zv != nil {
			return (*uint32)(ZendUintptrT(Z_PTR_P(zv)) & ^1)
		}
	} else {
		ZEND_ASSERT(Z_TYPE_P(zv) == IS_UNDEF)
		ZVAL_STR_COPY(zv, member)
		Z_PROPERTY_GUARD_P(zv) = 0
		return &Z_PROPERTY_GUARD_P(zv)
	}

	/* we have to allocate uint32_t separately because ht->arData may be reallocated */

	ptr = (*uint32)(Emalloc(b.SizeOf("uint32_t")))
	*ptr = 0
	return (*uint32)(ZendHashAddNewPtr(guards, member, ptr))
}
func ZendStdReadProperty(object *Zval, member *Zval, type_ int, cache_slot *any, rv *Zval) *Zval {
	var zobj *ZendObject
	var name *ZendString
	var tmp_name *ZendString
	var retval *Zval
	var property_offset uintPtr
	var prop_info *ZendPropertyInfo = nil
	var guard *uint32 = nil
	zobj = Z_OBJ_P(object)
	name = ZvalTryGetTmpString(member, &tmp_name)
	if UNEXPECTED(name == nil) {
		return &(ExecutorGlobals.GetUninitializedZval())
	}

	/* make zend_get_property_info silent if we have getter - we may want to use it */

	property_offset = ZendGetPropertyOffset(zobj.GetCe(), name, type_ == BP_VAR_IS || zobj.GetCe().GetGet() != nil, cache_slot, &prop_info)
	if EXPECTED(IS_VALID_PROPERTY_OFFSET(property_offset)) {
		retval = OBJ_PROP(zobj, property_offset)
		if EXPECTED(Z_TYPE_P(retval) != IS_UNDEF) {
			goto exit
		}
		if UNEXPECTED(Z_PROP_FLAG_P(retval) == IS_PROP_UNINIT) {

			/* Skip __get() for uninitialized typed properties */

			goto uninit_error

			/* Skip __get() for uninitialized typed properties */

		}
	} else if EXPECTED(IS_DYNAMIC_PROPERTY_OFFSET(property_offset)) {
		if EXPECTED(zobj.GetProperties() != nil) {
			if !(IS_UNKNOWN_DYNAMIC_PROPERTY_OFFSET(property_offset)) {
				var idx uintPtr = ZEND_DECODE_DYN_PROP_OFFSET(property_offset)
				if EXPECTED(idx < zobj.GetProperties().GetNNumUsed()*b.SizeOf("Bucket")) {
					var p *Bucket = (*Bucket)((*byte)(zobj.GetProperties().GetArData() + idx))
					if EXPECTED(Z_TYPE(p.GetVal()) != IS_UNDEF) && (EXPECTED(p.GetKey() == name) || EXPECTED(p.GetH() == ZSTR_H(name)) && EXPECTED(p.GetKey() != nil) && EXPECTED(ZendStringEqualContent(p.GetKey(), name) != 0)) {
						retval = &p.val
						goto exit
					}
				}
				CACHE_PTR_EX(cache_slot+1, any(ZEND_DYNAMIC_PROPERTY_OFFSET))
			}
			retval = ZendHashFind(zobj.GetProperties(), name)
			if EXPECTED(retval != nil) {
				if cache_slot != nil {
					var idx uintPtr = (*byte)(retval - (*byte)(zobj.GetProperties().GetArData()))
					CACHE_PTR_EX(cache_slot+1, any(ZEND_ENCODE_DYN_PROP_OFFSET(idx)))
				}
				goto exit
			}
		}
	} else if UNEXPECTED(ExecutorGlobals.GetException() != nil) {
		retval = &(ExecutorGlobals.GetUninitializedZval())
		goto exit
	}

	/* magic isset */

	if type_ == BP_VAR_IS && zobj.GetCe().GetIsset() != nil {
		var tmp_result Zval
		guard = ZendGetPropertyGuard(zobj, name)
		if ((*guard) & IN_ISSET) == 0 {
			if tmp_name == nil && ZSTR_IS_INTERNED(name) == 0 {
				tmp_name = ZendStringCopy(name)
			}
			GC_ADDREF(zobj)
			ZVAL_UNDEF(&tmp_result)
			*guard |= IN_ISSET
			ZendStdCallIssetter(zobj, name, &tmp_result)
			*guard &= ^IN_ISSET
			if ZendIsTrue(&tmp_result) == 0 {
				retval = &(ExecutorGlobals.GetUninitializedZval())
				OBJ_RELEASE(zobj)
				ZvalPtrDtor(&tmp_result)
				goto exit
			}
			ZvalPtrDtor(&tmp_result)
			if zobj.GetCe().GetGet() != nil && ((*guard)&IN_GET) == 0 {
				goto call_getter
			}
			OBJ_RELEASE(zobj)
		} else if zobj.GetCe().GetGet() != nil && ((*guard)&IN_GET) == 0 {
			goto call_getter_addref
		}
	} else if zobj.GetCe().GetGet() != nil {

		/* magic get */

		guard = ZendGetPropertyGuard(zobj, name)
		if ((*guard) & IN_GET) == 0 {

			/* have getter - try with it! */

		call_getter_addref:
			GC_ADDREF(zobj)
		call_getter:
			*guard |= IN_GET
			ZendStdCallGetter(zobj, name, rv)
			*guard &= ^IN_GET
			if Z_TYPE_P(rv) != IS_UNDEF {
				retval = rv
				if !(Z_ISREF_P(rv)) && (type_ == BP_VAR_W || type_ == BP_VAR_RW || type_ == BP_VAR_UNSET) {
					if UNEXPECTED(Z_TYPE_P(rv) != IS_OBJECT) {
						ZendError(E_NOTICE, "Indirect modification of overloaded property %s::$%s has no effect", ZSTR_VAL(zobj.GetCe().GetName()), ZSTR_VAL(name))
					}
				}
			} else {
				retval = &(ExecutorGlobals.GetUninitializedZval())
			}
			if UNEXPECTED(prop_info != nil) {
				ZendVerifyPropAssignableByRef(prop_info, retval, (zobj.GetCe().GetGet().GetFnFlags()&ZEND_ACC_STRICT_TYPES) != 0)
			}
			OBJ_RELEASE(zobj)
			goto exit
		} else if UNEXPECTED(IS_WRONG_PROPERTY_OFFSET(property_offset)) {

			/* Trigger the correct error */

			ZendGetPropertyOffset(zobj.GetCe(), name, 0, nil, &prop_info)
			ZEND_ASSERT(ExecutorGlobals.GetException() != nil)
			retval = &(ExecutorGlobals.GetUninitializedZval())
			goto exit
		}
	}
uninit_error:
	if type_ != BP_VAR_IS {
		if UNEXPECTED(prop_info != nil) {
			ZendThrowError(nil, "Typed property %s::$%s must not be accessed before initialization", ZSTR_VAL(prop_info.GetCe().GetName()), ZSTR_VAL(name))
		} else {
			ZendError(E_NOTICE, "Undefined property: %s::$%s", ZSTR_VAL(zobj.GetCe().GetName()), ZSTR_VAL(name))
		}
	}
	retval = &(ExecutorGlobals.GetUninitializedZval())
exit:
	ZendTmpStringRelease(tmp_name)
	return retval
}
func PropertyUsesStrictTypes() ZendBool {
	var execute_data *ZendExecuteData = ExecutorGlobals.GetCurrentExecuteData()
	return execute_data != nil && execute_data.GetFunc() != nil && ZEND_CALL_USES_STRICT_TYPES(ExecutorGlobals.GetCurrentExecuteData())
}
func ZendStdWriteProperty(object *Zval, member *Zval, value *Zval, cache_slot *any) *Zval {
	var zobj *ZendObject
	var name *ZendString
	var tmp_name *ZendString
	var variable_ptr *Zval
	var tmp Zval
	var property_offset uintPtr
	var prop_info *ZendPropertyInfo = nil
	ZEND_ASSERT(!(Z_ISREF_P(value)))
	zobj = Z_OBJ_P(object)
	name = ZvalTryGetTmpString(member, &tmp_name)
	if UNEXPECTED(name == nil) {
		return value
	}
	property_offset = ZendGetPropertyOffset(zobj.GetCe(), name, zobj.GetCe().GetSet() != nil, cache_slot, &prop_info)
	if EXPECTED(IS_VALID_PROPERTY_OFFSET(property_offset)) {
		variable_ptr = OBJ_PROP(zobj, property_offset)
		if Z_TYPE_P(variable_ptr) != IS_UNDEF {
			Z_TRY_ADDREF_P(value)
			if UNEXPECTED(prop_info != nil) {
				ZVAL_COPY_VALUE(&tmp, value)
				if UNEXPECTED(ZendVerifyPropertyType(prop_info, &tmp, PropertyUsesStrictTypes()) == 0) {
					Z_TRY_DELREF_P(value)
					variable_ptr = &(ExecutorGlobals.GetErrorZval())
					goto exit
				}
				value = &tmp
			}
		found:
			variable_ptr = ZendAssignToVariable(variable_ptr, value, IS_TMP_VAR, PropertyUsesStrictTypes())
			goto exit
		}
		if Z_PROP_FLAG_P(variable_ptr) == IS_PROP_UNINIT {

			/* Writes to uninitializde typed properties bypass __set(). */

			Z_PROP_FLAG_P(variable_ptr) = 0
			goto write_std_property
		}
	} else if EXPECTED(IS_DYNAMIC_PROPERTY_OFFSET(property_offset)) {
		if EXPECTED(zobj.GetProperties() != nil) {
			if UNEXPECTED(GC_REFCOUNT(zobj.GetProperties()) > 1) {
				if EXPECTED((GC_FLAGS(zobj.GetProperties()) & IS_ARRAY_IMMUTABLE) == 0) {
					GC_DELREF(zobj.GetProperties())
				}
				zobj.SetProperties(ZendArrayDup(zobj.GetProperties()))
			}
			if b.Assign(&variable_ptr, ZendHashFind(zobj.GetProperties(), name)) != nil {
				Z_TRY_ADDREF_P(value)
				goto found
			}
		}
	} else if UNEXPECTED(ExecutorGlobals.GetException() != nil) {
		variable_ptr = &(ExecutorGlobals.GetErrorZval())
		goto exit
	}

	/* magic set */

	if zobj.GetCe().GetSet() != nil {
		var guard *uint32 = ZendGetPropertyGuard(zobj, name)
		if ((*guard) & IN_SET) == 0 {
			GC_ADDREF(zobj)
			*guard |= IN_SET
			ZendStdCallSetter(zobj, name, value)
			*guard &= ^IN_SET
			OBJ_RELEASE(zobj)
			variable_ptr = value
		} else if EXPECTED(!(IS_WRONG_PROPERTY_OFFSET(property_offset))) {
			goto write_std_property
		} else {

			/* Trigger the correct error */

			ZendWrongOffset(zobj.GetCe(), name)
			ZEND_ASSERT(ExecutorGlobals.GetException() != nil)
			variable_ptr = &(ExecutorGlobals.GetErrorZval())
			goto exit
		}
	} else {
		ZEND_ASSERT(!(IS_WRONG_PROPERTY_OFFSET(property_offset)))
	write_std_property:
		Z_TRY_ADDREF_P(value)
		if EXPECTED(IS_VALID_PROPERTY_OFFSET(property_offset)) {
			variable_ptr = OBJ_PROP(zobj, property_offset)
			if UNEXPECTED(prop_info != nil) {
				ZVAL_COPY_VALUE(&tmp, value)
				if UNEXPECTED(ZendVerifyPropertyType(prop_info, &tmp, PropertyUsesStrictTypes()) == 0) {
					ZvalPtrDtor(value)
					goto exit
				}
				value = &tmp
				goto found
			}
			ZVAL_COPY_VALUE(variable_ptr, value)
		} else {
			if zobj.GetProperties() == nil {
				RebuildObjectProperties(zobj)
			}
			variable_ptr = ZendHashAddNew(zobj.GetProperties(), name, value)
		}
	}
exit:
	ZendTmpStringRelease(tmp_name)
	return variable_ptr
}
func ZendBadArrayAccess(ce *ZendClassEntry) {
	ZendThrowError(nil, "Cannot use object of type %s as array", ZSTR_VAL(ce.GetName()))
}
func ZendStdReadDimension(object *Zval, offset *Zval, type_ int, rv *Zval) *Zval {
	var ce *ZendClassEntry = Z_OBJCE_P(object)
	var tmp_offset Zval
	var tmp_object Zval
	if EXPECTED(InstanceofFunctionEx(ce, ZendCeArrayaccess, 1) != 0) {
		if offset == nil {

			/* [] construct */

			ZVAL_NULL(&tmp_offset)

			/* [] construct */

		} else {
			ZVAL_COPY_DEREF(&tmp_offset, offset)
		}
		Z_ADDREF_P(object)
		ZVAL_OBJ(&tmp_object, Z_OBJ_P(object))
		if type_ == BP_VAR_IS {
			ZendCallMethodWith1Params(&tmp_object, ce, nil, "offsetexists", rv, &tmp_offset)
			if UNEXPECTED(Z_ISUNDEF_P(rv)) {
				ZvalPtrDtor(&tmp_object)
				ZvalPtrDtor(&tmp_offset)
				return nil
			}
			if IZendIsTrue(rv) == 0 {
				ZvalPtrDtor(&tmp_object)
				ZvalPtrDtor(&tmp_offset)
				ZvalPtrDtor(rv)
				return &(ExecutorGlobals.GetUninitializedZval())
			}
			ZvalPtrDtor(rv)
		}
		ZendCallMethodWith1Params(&tmp_object, ce, nil, "offsetget", rv, &tmp_offset)
		ZvalPtrDtor(&tmp_object)
		ZvalPtrDtor(&tmp_offset)
		if UNEXPECTED(Z_TYPE_P(rv) == IS_UNDEF) {
			if UNEXPECTED(ExecutorGlobals.GetException() == nil) {
				ZendThrowError(nil, "Undefined offset for object of type %s used as array", ZSTR_VAL(ce.GetName()))
			}
			return nil
		}
		return rv
	} else {
		ZendBadArrayAccess(ce)
		return nil
	}
}
func ZendStdWriteDimension(object *Zval, offset *Zval, value *Zval) {
	var ce *ZendClassEntry = Z_OBJCE_P(object)
	var tmp_offset Zval
	var tmp_object Zval
	if EXPECTED(InstanceofFunctionEx(ce, ZendCeArrayaccess, 1) != 0) {
		if offset == nil {
			ZVAL_NULL(&tmp_offset)
		} else {
			ZVAL_COPY_DEREF(&tmp_offset, offset)
		}
		Z_ADDREF_P(object)
		ZVAL_OBJ(&tmp_object, Z_OBJ_P(object))
		ZendCallMethodWith2Params(&tmp_object, ce, nil, "offsetset", nil, &tmp_offset, value)
		ZvalPtrDtor(&tmp_object)
		ZvalPtrDtor(&tmp_offset)
	} else {
		ZendBadArrayAccess(ce)
	}
}
func ZendStdHasDimension(object *Zval, offset *Zval, check_empty int) int {
	var ce *ZendClassEntry = Z_OBJCE_P(object)
	var retval Zval
	var tmp_offset Zval
	var tmp_object Zval
	var result int
	if EXPECTED(InstanceofFunctionEx(ce, ZendCeArrayaccess, 1) != 0) {
		ZVAL_COPY_DEREF(&tmp_offset, offset)
		Z_ADDREF_P(object)
		ZVAL_OBJ(&tmp_object, Z_OBJ_P(object))
		ZendCallMethodWith1Params(&tmp_object, ce, nil, "offsetexists", &retval, &tmp_offset)
		result = IZendIsTrue(&retval)
		ZvalPtrDtor(&retval)
		if check_empty != 0 && result != 0 && EXPECTED(ExecutorGlobals.GetException() == nil) {
			ZendCallMethodWith1Params(&tmp_object, ce, nil, "offsetget", &retval, &tmp_offset)
			result = IZendIsTrue(&retval)
			ZvalPtrDtor(&retval)
		}
		ZvalPtrDtor(&tmp_object)
		ZvalPtrDtor(&tmp_offset)
	} else {
		ZendBadArrayAccess(ce)
		return 0
	}
	return result
}
func ZendStdGetPropertyPtrPtr(object *Zval, member *Zval, type_ int, cache_slot *any) *Zval {
	var zobj *ZendObject
	var name *ZendString
	var tmp_name *ZendString
	var retval *Zval = nil
	var property_offset uintPtr
	var prop_info *ZendPropertyInfo = nil
	zobj = Z_OBJ_P(object)
	name = ZvalTryGetTmpString(member, &tmp_name)
	if UNEXPECTED(name == nil) {
		return &(ExecutorGlobals.GetErrorZval())
	}
	property_offset = ZendGetPropertyOffset(zobj.GetCe(), name, zobj.GetCe().GetGet() != nil, cache_slot, &prop_info)
	if EXPECTED(IS_VALID_PROPERTY_OFFSET(property_offset)) {
		retval = OBJ_PROP(zobj, property_offset)
		if UNEXPECTED(Z_TYPE_P(retval) == IS_UNDEF) {
			if EXPECTED(zobj.GetCe().GetGet() == nil) || UNEXPECTED(((*ZendGetPropertyGuard)(zobj, name)&IN_GET) != 0) || UNEXPECTED(prop_info != nil && Z_PROP_FLAG_P(retval) == IS_PROP_UNINIT) {
				if UNEXPECTED(type_ == BP_VAR_RW || type_ == BP_VAR_R) {
					if UNEXPECTED(prop_info != nil) {
						ZendThrowError(nil, "Typed property %s::$%s must not be accessed before initialization", ZSTR_VAL(prop_info.GetCe().GetName()), ZSTR_VAL(name))
						retval = &(ExecutorGlobals.GetErrorZval())
					} else {
						ZVAL_NULL(retval)
						ZendError(E_NOTICE, "Undefined property: %s::$%s", ZSTR_VAL(zobj.GetCe().GetName()), ZSTR_VAL(name))
					}
				}
			} else {

				/* we do have getter - fail and let it try again with usual get/set */

				retval = nil

				/* we do have getter - fail and let it try again with usual get/set */

			}
		}
	} else if EXPECTED(IS_DYNAMIC_PROPERTY_OFFSET(property_offset)) {
		if EXPECTED(zobj.GetProperties() != nil) {
			if UNEXPECTED(GC_REFCOUNT(zobj.GetProperties()) > 1) {
				if EXPECTED((GC_FLAGS(zobj.GetProperties()) & IS_ARRAY_IMMUTABLE) == 0) {
					GC_DELREF(zobj.GetProperties())
				}
				zobj.SetProperties(ZendArrayDup(zobj.GetProperties()))
			}
			if EXPECTED(b.Assign(&retval, ZendHashFind(zobj.GetProperties(), name)) != nil) {
				ZendTmpStringRelease(tmp_name)
				return retval
			}
		}
		if EXPECTED(zobj.GetCe().GetGet() == nil) || UNEXPECTED(((*ZendGetPropertyGuard)(zobj, name)&IN_GET) != 0) {
			if UNEXPECTED(zobj.GetProperties() == nil) {
				RebuildObjectProperties(zobj)
			}
			retval = ZendHashUpdate(zobj.GetProperties(), name, &(ExecutorGlobals.GetUninitializedZval()))

			/* Notice is thrown after creation of the property, to avoid EG(std_property_info)
			 * being overwritten in an error handler. */

			if UNEXPECTED(type_ == BP_VAR_RW || type_ == BP_VAR_R) {
				ZendError(E_NOTICE, "Undefined property: %s::$%s", ZSTR_VAL(zobj.GetCe().GetName()), ZSTR_VAL(name))
			}

			/* Notice is thrown after creation of the property, to avoid EG(std_property_info)
			 * being overwritten in an error handler. */

		}
	} else if zobj.GetCe().GetGet() == nil {
		retval = &(ExecutorGlobals.GetErrorZval())
	}
	ZendTmpStringRelease(tmp_name)
	return retval
}
func ZendStdUnsetProperty(object *Zval, member *Zval, cache_slot *any) {
	var zobj *ZendObject
	var name *ZendString
	var tmp_name *ZendString
	var property_offset uintPtr
	var prop_info *ZendPropertyInfo = nil
	zobj = Z_OBJ_P(object)
	name = ZvalTryGetTmpString(member, &tmp_name)
	if UNEXPECTED(name == nil) {
		return
	}
	property_offset = ZendGetPropertyOffset(zobj.GetCe(), name, zobj.GetCe().GetUnset() != nil, cache_slot, &prop_info)
	if EXPECTED(IS_VALID_PROPERTY_OFFSET(property_offset)) {
		var slot *Zval = OBJ_PROP(zobj, property_offset)
		if Z_TYPE_P(slot) != IS_UNDEF {
			if UNEXPECTED(Z_ISREF_P(slot)) && (core.ZEND_DEBUG != 0 || ZEND_REF_HAS_TYPE_SOURCES(Z_REF_P(slot))) {
				if prop_info != nil {
					ZEND_REF_DEL_TYPE_SOURCE(Z_REF_P(slot), prop_info)
				}
			}
			var tmp Zval
			ZVAL_COPY_VALUE(&tmp, slot)
			ZVAL_UNDEF(slot)
			ZvalPtrDtor(&tmp)
			if zobj.GetProperties() != nil {
				HT_FLAGS(zobj.GetProperties()) |= HASH_FLAG_HAS_EMPTY_IND
			}
			goto exit
		}
		if UNEXPECTED(Z_PROP_FLAG_P(slot) == IS_PROP_UNINIT) {

			/* Reset the IS_PROP_UNINIT flag, if it exists and bypass __unset(). */

			Z_PROP_FLAG_P(slot) = 0
			goto exit
		}
	} else if EXPECTED(IS_DYNAMIC_PROPERTY_OFFSET(property_offset)) && EXPECTED(zobj.GetProperties() != nil) {
		if UNEXPECTED(GC_REFCOUNT(zobj.GetProperties()) > 1) {
			if EXPECTED((GC_FLAGS(zobj.GetProperties()) & IS_ARRAY_IMMUTABLE) == 0) {
				GC_DELREF(zobj.GetProperties())
			}
			zobj.SetProperties(ZendArrayDup(zobj.GetProperties()))
		}
		if EXPECTED(ZendHashDel(zobj.GetProperties(), name) != FAILURE) {
			goto exit
		}
	} else if UNEXPECTED(ExecutorGlobals.GetException() != nil) {
		goto exit
	}

	/* magic unset */

	if zobj.GetCe().GetUnset() != nil {
		var guard *uint32 = ZendGetPropertyGuard(zobj, name)
		if ((*guard) & IN_UNSET) == 0 {

			/* have unseter - try with it! */

			*guard |= IN_UNSET
			ZendStdCallUnsetter(zobj, name)
			*guard &= ^IN_UNSET
		} else if UNEXPECTED(IS_WRONG_PROPERTY_OFFSET(property_offset)) {

			/* Trigger the correct error */

			ZendWrongOffset(zobj.GetCe(), name)
			ZEND_ASSERT(ExecutorGlobals.GetException() != nil)
			goto exit
		}
	}
exit:
	ZendTmpStringRelease(tmp_name)
}
func ZendStdUnsetDimension(object *Zval, offset *Zval) {
	var ce *ZendClassEntry = Z_OBJCE_P(object)
	var tmp_offset Zval
	var tmp_object Zval
	if InstanceofFunctionEx(ce, ZendCeArrayaccess, 1) != 0 {
		ZVAL_COPY_DEREF(&tmp_offset, offset)
		Z_ADDREF_P(object)
		ZVAL_OBJ(&tmp_object, Z_OBJ_P(object))
		ZendCallMethodWith1Params(&tmp_object, ce, nil, "offsetunset", nil, &tmp_offset)
		ZvalPtrDtor(&tmp_object)
		ZvalPtrDtor(&tmp_offset)
	} else {
		ZendBadArrayAccess(ce)
	}
}
func ZendGetParentPrivateMethod(scope *ZendClassEntry, ce *ZendClassEntry, function_name *ZendString) *ZendFunction {
	var func_ *Zval
	var fbc *ZendFunction
	if scope != ce && scope != nil && IsDerivedClass(ce, scope) != 0 {
		func_ = ZendHashFind(&scope.function_table, function_name)
		if func_ != nil {
			fbc = Z_FUNC_P(func_)
			if (fbc.GetFnFlags()&ZEND_ACC_PRIVATE) != 0 && fbc.GetScope() == scope {
				return fbc
			}
		}
	}
	return nil
}
func ZendCheckProtected(ce *ZendClassEntry, scope *ZendClassEntry) int {
	var fbc_scope *ZendClassEntry = ce

	/* Is the context that's calling the function, the same as one of
	 * the function's parents?
	 */

	for fbc_scope != nil {
		if fbc_scope == scope {
			return 1
		}
		fbc_scope = fbc_scope.parent
	}

	/* Is the function's scope the same as our current object context,
	 * or any of the parents of our context?
	 */

	for scope != nil {
		if scope == ce {
			return 1
		}
		scope = scope.parent
	}
	return 0
}
func ZendGetCallTrampolineFunc(ce *ZendClassEntry, method_name *ZendString, is_static int) *ZendFunction {
	var mname_len int
	var func_ *ZendOpArray
	var fbc *ZendFunction = b.CondF(is_static != 0, func() *ZendFunction { return ce.GetCallstatic() }, func() *ZendFunction { return ce.GetCall() })

	/* We use non-NULL value to avoid useless run_time_cache allocation.
	 * The low bit must be zero, to not be interpreted as a MAP_PTR offset.
	 */

	var dummy any = any(intPtr(2))
	ZEND_ASSERT(fbc != nil)
	if EXPECTED(ExecutorGlobals.GetTrampoline().GetFunctionName() == nil) {
		func_ = &(ExecutorGlobals.GetTrampoline()).op_array
	} else {
		func_ = Ecalloc(1, b.SizeOf("zend_op_array"))
	}
	func_.SetType(ZEND_USER_FUNCTION)
	func_.GetArgFlags()[0] = 0
	func_.GetArgFlags()[1] = 0
	func_.GetArgFlags()[2] = 0
	func_.SetFnFlags(ZEND_ACC_CALL_VIA_TRAMPOLINE | ZEND_ACC_PUBLIC)
	if is_static != 0 {
		func_.SetFnFlags(func_.GetFnFlags() | ZEND_ACC_STATIC)
	}
	func_.SetOpcodes(&(ExecutorGlobals.GetCallTrampolineOp()))
	ZEND_MAP_PTR_INIT(func_.run_time_cache, (**any)(&dummy))
	func_.SetScope(fbc.GetScope())

	/* reserve space for arguments, local and temporary variables */

	if fbc.GetType() == ZEND_USER_FUNCTION {
		func_.SetT(MAX(fbc.GetOpArray().GetLastVar()+fbc.GetOpArray().GetT(), 2))
	} else {
		func_.SetT(2)
	}
	if fbc.GetType() == ZEND_USER_FUNCTION {
		func_.SetFilename(fbc.GetOpArray().GetFilename())
	} else {
		func_.SetFilename(ZSTR_EMPTY_ALLOC())
	}
	if fbc.GetType() == ZEND_USER_FUNCTION {
		func_.SetLineStart(fbc.GetOpArray().GetLineStart())
	} else {
		func_.SetLineStart(0)
	}
	if fbc.GetType() == ZEND_USER_FUNCTION {
		func_.SetLineEnd(fbc.GetOpArray().GetLineEnd())
	} else {
		func_.SetLineEnd(0)
	}

	//??? keep compatibility for "\0" characters

	if UNEXPECTED(b.Assign(&mname_len, strlen(ZSTR_VAL(method_name))) != ZSTR_LEN(method_name)) {
		func_.SetFunctionName(ZendStringInit(ZSTR_VAL(method_name), mname_len, 0))
	} else {
		func_.SetFunctionName(ZendStringCopy(method_name))
	}
	func_.SetPrototype(nil)
	func_.SetNumArgs(0)
	func_.SetRequiredNumArgs(0)
	func_.SetArgInfo(0)
	return (*ZendFunction)(func_)
}
func ZendGetUserCallFunction(ce *ZendClassEntry, method_name *ZendString) *ZendFunction {
	return ZendGetCallTrampolineFunc(ce, method_name, 0)
}
func ZendBadMethodCall(fbc *ZendFunction, method_name *ZendString, scope *ZendClassEntry) {
	ZendThrowError(nil, "Call to %s method %s::%s() from context '%s'", ZendVisibilityString(fbc.GetFnFlags()), ZEND_FN_SCOPE_NAME(fbc), ZSTR_VAL(method_name), b.CondF1(scope != nil, func() []byte { return ZSTR_VAL(scope.GetName()) }, ""))
}
func ZendStdGetMethod(obj_ptr **ZendObject, method_name *ZendString, key *Zval) *ZendFunction {
	var zobj *ZendObject = *obj_ptr
	var func_ *Zval
	var fbc *ZendFunction
	var lc_method_name *ZendString
	var scope *ZendClassEntry
	if EXPECTED(key != nil) {
		lc_method_name = Z_STR_P(key)
	} else {
		ZSTR_ALLOCA_ALLOC(lc_method_name, ZSTR_LEN(method_name), use_heap)
		ZendStrTolowerCopy(ZSTR_VAL(lc_method_name), ZSTR_VAL(method_name), ZSTR_LEN(method_name))
	}
	if UNEXPECTED(b.Assign(&func_, ZendHashFind(&zobj.ce.GetFunctionTable(), lc_method_name)) == nil) {
		if UNEXPECTED(key == nil) {
			ZSTR_ALLOCA_FREE(lc_method_name, use_heap)
		}
		if zobj.GetCe().GetCall() != nil {
			return ZendGetUserCallFunction(zobj.GetCe(), method_name)
		} else {
			return nil
		}
	}
	fbc = Z_FUNC_P(func_)

	/* Check access level */

	if (fbc.GetOpArray().GetFnFlags() & (ZEND_ACC_CHANGED | ZEND_ACC_PRIVATE | ZEND_ACC_PROTECTED)) != 0 {
		scope = ZendGetExecutedScope()
		if fbc.GetScope() != scope {
			if (fbc.GetOpArray().GetFnFlags() & ZEND_ACC_CHANGED) != 0 {
				var updated_fbc *ZendFunction = ZendGetParentPrivateMethod(scope, zobj.GetCe(), lc_method_name)
				if EXPECTED(updated_fbc != nil) {
					fbc = updated_fbc
					goto exit
				} else if (fbc.GetOpArray().GetFnFlags() & ZEND_ACC_PUBLIC) != 0 {
					goto exit
				}
			}
			if UNEXPECTED((fbc.GetOpArray().GetFnFlags()&ZEND_ACC_PRIVATE) != 0) || UNEXPECTED(ZendCheckProtected(ZendGetFunctionRootClass(fbc), scope) == 0) {
				if zobj.GetCe().GetCall() != nil {
					fbc = ZendGetUserCallFunction(zobj.GetCe(), method_name)
				} else {
					ZendBadMethodCall(fbc, method_name, scope)
					fbc = nil
				}
			}
		}
	}
exit:
	if UNEXPECTED(key == nil) {
		ZSTR_ALLOCA_FREE(lc_method_name, use_heap)
	}
	return fbc
}
func ZendGetUserCallstaticFunction(ce *ZendClassEntry, method_name *ZendString) *ZendFunction {
	return ZendGetCallTrampolineFunc(ce, method_name, 1)
}
func ZendStdGetStaticMethod(ce *ZendClassEntry, function_name *ZendString, key *Zval) *ZendFunction {
	var fbc *ZendFunction = nil
	var lc_function_name *ZendString
	var object *ZendObject
	var scope *ZendClassEntry
	if EXPECTED(key != nil) {
		lc_function_name = Z_STR_P(key)
	} else {
		lc_function_name = ZendStringTolower(function_name)
	}
	var func_ *Zval = ZendHashFind(&ce.function_table, lc_function_name)
	if EXPECTED(func_ != nil) {
		fbc = Z_FUNC_P(func_)
	} else if ce.GetConstructor() != nil && ZSTR_LEN(lc_function_name) == ZSTR_LEN(ce.GetName()) && ZendBinaryStrncasecmp(ZSTR_VAL(lc_function_name), ZSTR_LEN(lc_function_name), ZSTR_VAL(ce.GetName()), ZSTR_LEN(lc_function_name), ZSTR_LEN(lc_function_name)) == 0 && (ZSTR_VAL(ce.GetConstructor().GetFunctionName())[0] != '_' || ZSTR_VAL(ce.GetConstructor().GetFunctionName())[1] != '_') {
		fbc = ce.GetConstructor()
	} else {
		if UNEXPECTED(key == nil) {
			ZendStringReleaseEx(lc_function_name, 0)
		}
		if ce.GetCall() != nil && b.Assign(&object, ZendGetThisObject(ExecutorGlobals.GetCurrentExecuteData())) != nil && InstanceofFunction(object.GetCe(), ce) != 0 {

			/* Call the top-level defined __call().
			 * see: tests/classes/__call_004.phpt  */

			var call_ce *ZendClassEntry = object.GetCe()
			for call_ce.GetCall() == nil {
				call_ce = call_ce.parent
			}
			return ZendGetUserCallFunction(call_ce, function_name)
		} else if ce.GetCallstatic() != nil {
			return ZendGetUserCallstaticFunction(ce, function_name)
		} else {
			return nil
		}
	}
	if (fbc.GetOpArray().GetFnFlags() & ZEND_ACC_PUBLIC) == 0 {
		scope = ZendGetExecutedScope()
		if UNEXPECTED(fbc.GetScope() != scope) {
			if UNEXPECTED((fbc.GetOpArray().GetFnFlags()&ZEND_ACC_PRIVATE) != 0) || UNEXPECTED(ZendCheckProtected(ZendGetFunctionRootClass(fbc), scope) == 0) {
				if ce.GetCallstatic() != nil {
					fbc = ZendGetUserCallstaticFunction(ce, function_name)
				} else {
					ZendBadMethodCall(fbc, function_name, scope)
					fbc = nil
				}
			}
		}
	}
	if UNEXPECTED(key == nil) {
		ZendStringReleaseEx(lc_function_name, 0)
	}
	return fbc
}
func ZendClassInitStatics(class_type *ZendClassEntry) {
	var i int
	var p *Zval
	if class_type.GetDefaultStaticMembersCount() != 0 && CE_STATIC_MEMBERS(class_type) == nil {
		if class_type.parent {
			ZendClassInitStatics(class_type.parent)
		}
		ZEND_MAP_PTR_SET(class_type.static_members_table, Emalloc(b.SizeOf("zval")*class_type.GetDefaultStaticMembersCount()))
		for i = 0; i < class_type.GetDefaultStaticMembersCount(); i++ {
			p = &class_type.default_static_members_table[i]
			if Z_TYPE_P(p) == IS_INDIRECT {
				var q *Zval = &CE_STATIC_MEMBERS(class_type.parent)[i]
				ZVAL_DEINDIRECT(q)
				ZVAL_INDIRECT(&CE_STATIC_MEMBERS(class_type)[i], q)
			} else {
				ZVAL_COPY_OR_DUP(&CE_STATIC_MEMBERS(class_type)[i], p)
			}
		}
	}
}
func ZendStdGetStaticPropertyWithInfo(ce *ZendClassEntry, property_name *ZendString, type_ int, property_info_ptr **ZendPropertyInfo) *Zval {
	var ret *Zval
	var scope *ZendClassEntry
	var property_info *ZendPropertyInfo = ZendHashFindPtr(&ce.properties_info, property_name)
	*property_info_ptr = property_info
	if UNEXPECTED(property_info == nil) {
		goto undeclared_property
	}
	if (property_info.GetFlags() & ZEND_ACC_PUBLIC) == 0 {
		if UNEXPECTED(ExecutorGlobals.GetFakeScope() != nil) {
			scope = ExecutorGlobals.GetFakeScope()
		} else {
			scope = ZendGetExecutedScope()
		}
		if property_info.GetCe() != scope {
			if UNEXPECTED((property_info.GetFlags()&ZEND_ACC_PRIVATE) != 0) || UNEXPECTED(IsProtectedCompatibleScope(property_info.GetCe(), scope) == 0) {
				if type_ != BP_VAR_IS {
					ZendBadPropertyAccess(property_info, ce, property_name)
				}
				return nil
			}
		}
	}
	if UNEXPECTED((property_info.GetFlags() & ZEND_ACC_STATIC) == 0) {
		goto undeclared_property
	}
	if UNEXPECTED((ce.GetCeFlags() & ZEND_ACC_CONSTANTS_UPDATED) == 0) {
		if UNEXPECTED(ZendUpdateClassConstants(ce) != 0) != SUCCESS {
			return nil
		}
	}

	/* check if static properties were destroyed */

	if UNEXPECTED(CE_STATIC_MEMBERS(ce) == nil) {
		if ce.GetType() == ZEND_INTERNAL_CLASS || (ce.GetCeFlags()&(ZEND_ACC_IMMUTABLE|ZEND_ACC_PRELOADED)) != 0 {
			ZendClassInitStatics(ce)
		} else {
		undeclared_property:
			if type_ != BP_VAR_IS {
				ZendThrowError(nil, "Access to undeclared static property: %s::$%s", ZSTR_VAL(ce.GetName()), ZSTR_VAL(property_name))
			}
			return nil
		}
	}
	ret = CE_STATIC_MEMBERS(ce) + property_info.GetOffset()
	ZVAL_DEINDIRECT(ret)
	if UNEXPECTED((type_ == BP_VAR_R || type_ == BP_VAR_RW) && Z_TYPE_P(ret) == IS_UNDEF && property_info.GetType() != 0) {
		ZendThrowError(nil, "Typed static property %s::$%s must not be accessed before initialization", ZSTR_VAL(property_info.GetCe().GetName()), ZendGetUnmangledPropertyName(property_name))
		return nil
	}
	return ret
}
func ZendStdGetStaticProperty(ce *ZendClassEntry, property_name *ZendString, type_ int) *Zval {
	var prop_info *ZendPropertyInfo
	return ZendStdGetStaticPropertyWithInfo(ce, property_name, type_, &prop_info)
}
func ZendStdUnsetStaticProperty(ce *ZendClassEntry, property_name *ZendString) ZendBool {
	ZendThrowError(nil, "Attempt to unset static property %s::$%s", ZSTR_VAL(ce.GetName()), ZSTR_VAL(property_name))
	return 0
}
func ZendBadConstructorCall(constructor *ZendFunction, scope *ZendClassEntry) {
	if scope != nil {
		ZendThrowError(nil, "Call to %s %s::%s() from context '%s'", ZendVisibilityString(constructor.GetFnFlags()), ZSTR_VAL(constructor.GetScope().GetName()), ZSTR_VAL(constructor.GetFunctionName()), ZSTR_VAL(scope.GetName()))
	} else {
		ZendThrowError(nil, "Call to %s %s::%s() from invalid context", ZendVisibilityString(constructor.GetFnFlags()), ZSTR_VAL(constructor.GetScope().GetName()), ZSTR_VAL(constructor.GetFunctionName()))
	}
}
func ZendStdGetConstructor(zobj *ZendObject) *ZendFunction {
	var constructor *ZendFunction = zobj.GetCe().GetConstructor()
	var scope *ZendClassEntry
	if constructor != nil {
		if UNEXPECTED((constructor.GetOpArray().GetFnFlags() & ZEND_ACC_PUBLIC) == 0) {
			if UNEXPECTED(ExecutorGlobals.GetFakeScope() != nil) {
				scope = ExecutorGlobals.GetFakeScope()
			} else {
				scope = ZendGetExecutedScope()
			}
			if UNEXPECTED(constructor.GetScope() != scope) {
				if UNEXPECTED((constructor.GetOpArray().GetFnFlags()&ZEND_ACC_PRIVATE) != 0) || UNEXPECTED(ZendCheckProtected(ZendGetFunctionRootClass(constructor), scope) == 0) {
					ZendBadConstructorCall(constructor, scope)
					constructor = nil
				}
			}
		}
	}
	return constructor
}
func ZendStdCompareObjects(o1 *Zval, o2 *Zval) int {
	var zobj1 *ZendObject
	var zobj2 *ZendObject
	zobj1 = Z_OBJ_P(o1)
	zobj2 = Z_OBJ_P(o2)
	if zobj1 == zobj2 {
		return 0
	}
	if zobj1.GetCe() != zobj2.GetCe() {
		return 1
	}
	if zobj1.GetProperties() == nil && zobj2.GetProperties() == nil {
		var info *ZendPropertyInfo
		if zobj1.GetCe().GetDefaultPropertiesCount() == 0 {
			return 0
		}

		/* It's enough to protect only one of the objects.
		 * The second one may be referenced from the first and this may cause
		 * false recursion detection.
		 */

		if UNEXPECTED(Z_IS_RECURSIVE_P(o1) != 0) {
			ZendErrorNoreturn(E_ERROR, "Nesting level too deep - recursive dependency?")
		}
		Z_PROTECT_RECURSION_P(o1)
		for {
			var __ht *HashTable = &zobj1.ce.GetPropertiesInfo()
			var _p *Bucket = __ht.GetArData()
			var _end *Bucket = _p + __ht.GetNNumUsed()
			for ; _p != _end; _p++ {
				var _z *Zval = &_p.val

				if UNEXPECTED(Z_TYPE_P(_z) == IS_UNDEF) {
					continue
				}
				info = Z_PTR_P(_z)
				var p1 *Zval = OBJ_PROP(zobj1, info.GetOffset())
				var p2 *Zval = OBJ_PROP(zobj2, info.GetOffset())
				if (info.GetFlags() & ZEND_ACC_STATIC) != 0 {
					continue
				}
				if Z_TYPE_P(p1) != IS_UNDEF {
					if Z_TYPE_P(p2) != IS_UNDEF {
						var result Zval
						if CompareFunction(&result, p1, p2) == FAILURE {
							Z_UNPROTECT_RECURSION_P(o1)
							return 1
						}
						if Z_LVAL(result) != 0 {
							Z_UNPROTECT_RECURSION_P(o1)
							return Z_LVAL(result)
						}
					} else {
						Z_UNPROTECT_RECURSION_P(o1)
						return 1
					}
				} else {
					if Z_TYPE_P(p2) != IS_UNDEF {
						Z_UNPROTECT_RECURSION_P(o1)
						return 1
					}
				}
			}
			break
		}
		Z_UNPROTECT_RECURSION_P(o1)
		return 0
	} else {
		if zobj1.GetProperties() == nil {
			RebuildObjectProperties(zobj1)
		}
		if zobj2.GetProperties() == nil {
			RebuildObjectProperties(zobj2)
		}
		return ZendCompareSymbolTables(zobj1.GetProperties(), zobj2.GetProperties())
	}
}
func ZendStdHasProperty(object *Zval, member *Zval, has_set_exists int, cache_slot *any) int {
	var zobj *ZendObject
	var result int
	var value *Zval = nil
	var name *ZendString
	var tmp_name *ZendString
	var property_offset uintPtr
	var prop_info *ZendPropertyInfo = nil
	zobj = Z_OBJ_P(object)
	name = ZvalTryGetTmpString(member, &tmp_name)
	if UNEXPECTED(name == nil) {
		return 0
	}
	property_offset = ZendGetPropertyOffset(zobj.GetCe(), name, 1, cache_slot, &prop_info)
	if EXPECTED(IS_VALID_PROPERTY_OFFSET(property_offset)) {
		value = OBJ_PROP(zobj, property_offset)
		if Z_TYPE_P(value) != IS_UNDEF {
			goto found
		}
		if UNEXPECTED(Z_PROP_FLAG_P(value) == IS_PROP_UNINIT) {

			/* Skip __isset() for uninitialized typed properties */

			result = 0
			goto exit
		}
	} else if EXPECTED(IS_DYNAMIC_PROPERTY_OFFSET(property_offset)) {
		if EXPECTED(zobj.GetProperties() != nil) {
			if !(IS_UNKNOWN_DYNAMIC_PROPERTY_OFFSET(property_offset)) {
				var idx uintPtr = ZEND_DECODE_DYN_PROP_OFFSET(property_offset)
				if EXPECTED(idx < zobj.GetProperties().GetNNumUsed()*b.SizeOf("Bucket")) {
					var p *Bucket = (*Bucket)((*byte)(zobj.GetProperties().GetArData() + idx))
					if EXPECTED(Z_TYPE(p.GetVal()) != IS_UNDEF) && (EXPECTED(p.GetKey() == name) || EXPECTED(p.GetH() == ZSTR_H(name)) && EXPECTED(p.GetKey() != nil) && EXPECTED(ZendStringEqualContent(p.GetKey(), name) != 0)) {
						value = &p.val
						goto found
					}
				}
				CACHE_PTR_EX(cache_slot+1, any(ZEND_DYNAMIC_PROPERTY_OFFSET))
			}
			value = ZendHashFind(zobj.GetProperties(), name)
			if value != nil {
				if cache_slot != nil {
					var idx uintPtr = (*byte)(value - (*byte)(zobj.GetProperties().GetArData()))
					CACHE_PTR_EX(cache_slot+1, any(ZEND_ENCODE_DYN_PROP_OFFSET(idx)))
				}
			found:
				if has_set_exists == ZEND_PROPERTY_NOT_EMPTY {
					result = ZendIsTrue(value)
				} else if has_set_exists < ZEND_PROPERTY_NOT_EMPTY {
					ZEND_ASSERT(has_set_exists == ZEND_PROPERTY_ISSET)
					ZVAL_DEREF(value)
					result = Z_TYPE_P(value) != IS_NULL
				} else {
					ZEND_ASSERT(has_set_exists == ZEND_PROPERTY_EXISTS)
					result = 1
				}
				goto exit
			}
		}
	} else if UNEXPECTED(ExecutorGlobals.GetException() != nil) {
		result = 0
		goto exit
	}
	result = 0
	if has_set_exists != ZEND_PROPERTY_EXISTS && zobj.GetCe().GetIsset() != nil {
		var guard *uint32 = ZendGetPropertyGuard(zobj, name)
		if ((*guard) & IN_ISSET) == 0 {
			var rv Zval

			/* have issetter - try with it! */

			if tmp_name == nil && ZSTR_IS_INTERNED(name) == 0 {
				tmp_name = ZendStringCopy(name)
			}
			GC_ADDREF(zobj)
			*guard |= IN_ISSET
			ZendStdCallIssetter(zobj, name, &rv)
			result = ZendIsTrue(&rv)
			ZvalPtrDtor(&rv)
			if has_set_exists == ZEND_PROPERTY_NOT_EMPTY && result != 0 {
				if EXPECTED(ExecutorGlobals.GetException() == nil) && zobj.GetCe().GetGet() != nil && ((*guard)&IN_GET) == 0 {
					*guard |= IN_GET
					ZendStdCallGetter(zobj, name, &rv)
					*guard &= ^IN_GET
					result = IZendIsTrue(&rv)
					ZvalPtrDtor(&rv)
				} else {
					result = 0
				}
			}
			*guard &= ^IN_ISSET
			OBJ_RELEASE(zobj)
		}
	}
exit:
	ZendTmpStringRelease(tmp_name)
	return result
}
func ZendStdGetClassName(zobj *ZendObject) *ZendString {
	return ZendStringCopy(zobj.GetCe().GetName())
}
func ZendStdCastObjectTostring(readobj *Zval, writeobj *Zval, type_ int) int {
	var retval Zval
	var ce *ZendClassEntry
	switch type_ {
	case IS_STRING:
		ce = Z_OBJCE_P(readobj)
		if ce.GetTostring() != nil {
			var fake_scope *ZendClassEntry = ExecutorGlobals.GetFakeScope()
			ExecutorGlobals.SetFakeScope(nil)
			ZendCallMethodWith0Params(readobj, ce, &ce.__tostring, "__tostring", &retval)
			ExecutorGlobals.SetFakeScope(fake_scope)
			if EXPECTED(Z_TYPE(retval) == IS_STRING) {
				ZVAL_COPY_VALUE(writeobj, &retval)
				return SUCCESS
			}
			ZvalPtrDtor(&retval)
			if ExecutorGlobals.GetException() == nil {
				ZendThrowError(nil, "Method %s::__toString() must return a string value", ZSTR_VAL(ce.GetName()))
			}
		}
		return FAILURE
	case _IS_BOOL:
		ZVAL_TRUE(writeobj)
		return SUCCESS
	case IS_LONG:
		ce = Z_OBJCE_P(readobj)
		ZendError(E_NOTICE, "Object of class %s could not be converted to int", ZSTR_VAL(ce.GetName()))
		ZVAL_LONG(writeobj, 1)
		return SUCCESS
	case IS_DOUBLE:
		ce = Z_OBJCE_P(readobj)
		ZendError(E_NOTICE, "Object of class %s could not be converted to float", ZSTR_VAL(ce.GetName()))
		ZVAL_DOUBLE(writeobj, 1)
		return SUCCESS
	case _IS_NUMBER:
		ce = Z_OBJCE_P(readobj)
		ZendError(E_NOTICE, "Object of class %s could not be converted to number", ZSTR_VAL(ce.GetName()))
		ZVAL_LONG(writeobj, 1)
		return SUCCESS
	default:
		ZVAL_NULL(writeobj)
		break
	}
	return FAILURE
}
func ZendStdGetClosure(obj *Zval, ce_ptr **ZendClassEntry, fptr_ptr **ZendFunction, obj_ptr **ZendObject) int {
	var func_ *Zval
	var ce *ZendClassEntry = Z_OBJCE_P(obj)
	if b.Assign(&func_, ZendHashFindEx(&ce.function_table, ZSTR_KNOWN(ZEND_STR_MAGIC_INVOKE), 1)) == nil {
		return FAILURE
	}
	*fptr_ptr = Z_FUNC_P(func_)
	*ce_ptr = ce
	if ((*fptr_ptr).GetFnFlags() & ZEND_ACC_STATIC) != 0 {
		if obj_ptr != nil {
			*obj_ptr = nil
		}
	} else {
		if obj_ptr != nil {
			*obj_ptr = Z_OBJ_P(obj)
		}
	}
	return SUCCESS
}
func ZendStdGetPropertiesFor(obj *Zval, purpose ZendPropPurpose) *HashTable {
	var ht *HashTable
	switch purpose {
	case ZEND_PROP_PURPOSE_DEBUG:
		if Z_OBJ_HT_P(obj).GetGetDebugInfo() != nil {
			var is_temp int
			ht = Z_OBJ_HT_P(obj).GetGetDebugInfo()(obj, &is_temp)
			if ht != nil && is_temp == 0 && (GC_FLAGS(ht)&GC_IMMUTABLE) == 0 {
				GC_ADDREF(ht)
			}
			return ht
		}
	case ZEND_PROP_PURPOSE_ARRAY_CAST:

	case ZEND_PROP_PURPOSE_SERIALIZE:

	case ZEND_PROP_PURPOSE_VAR_EXPORT:

	case ZEND_PROP_PURPOSE_JSON:

	case _ZEND_PROP_PURPOSE_ARRAY_KEY_EXISTS:
		ht = Z_OBJ_HT_P(obj).GetGetProperties()(obj)
		if ht != nil && (GC_FLAGS(ht)&GC_IMMUTABLE) == 0 {
			GC_ADDREF(ht)
		}
		return ht
	default:
		ZEND_ASSERT(false)
		return nil
	}
}
func ZendGetPropertiesFor(obj *Zval, purpose ZendPropPurpose) *HashTable {
	if Z_OBJ_HT_P(obj).GetGetPropertiesFor() != nil {
		return Z_OBJ_HT_P(obj).GetGetPropertiesFor()(obj, purpose)
	}
	return ZendStdGetPropertiesFor(obj, purpose)
}
