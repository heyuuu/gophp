// <<generate>>

package zend

import (
	b "sik/builtin"
	"sik/zend/faults"
	"sik/zend/types"
)

func IS_VALID_PROPERTY_OFFSET(offset uintPtr) bool   { return intptr_t(offset) > 0 }
func IS_WRONG_PROPERTY_OFFSET(offset uintPtr) bool   { return intptr_t(offset) == 0 }
func IS_DYNAMIC_PROPERTY_OFFSET(offset uintPtr) bool { return intptr_t(offset) < 0 }
func IS_UNKNOWN_DYNAMIC_PROPERTY_OFFSET(offset uintPtr) bool {
	return offset == ZEND_DYNAMIC_PROPERTY_OFFSET
}
func ZEND_DECODE_DYN_PROP_OFFSET(offset uintPtr) __auto__ { return uintPtr(-(intptr_t(offset)) - 2) }
func ZEND_ENCODE_DYN_PROP_OFFSET(offset uintPtr) __auto__ { return uintPtr(-(intptr_t(offset) + 2)) }
func ZendGetStdObjectHandlers() *ZendObjectHandlers       { return &StdObjectHandlers }
func ZendGetFunctionRootClass(fbc *ZendFunction) *types.ClassEntry {
	if fbc.GetPrototype() != nil {
		return fbc.GetPrototype().GetScope()
	} else {
		return fbc.GetScope()
	}
}
func ZendReleaseProperties(ht *types.Array) {
	if ht != nil && (ht.GetGcFlags()&types.GC_IMMUTABLE) == 0 && ht.DelRefcount() == 0 {
		ht.DestroyEx()
	}
}
func ZendFreeTrampoline(func_ any) {
	if func_ == EG__().GetTrampoline() {
		EG__().GetTrampoline().SetFunctionName(nil)
	} else {
		Efree(func_)
	}
}
func RebuildObjectProperties(zobj *types.ZendObject) {
	if zobj.GetProperties() == nil {
		var prop_info *ZendPropertyInfo
		var ce *types.ClassEntry = zobj.GetCe()
		var flags uint32 = 0
		zobj.SetProperties(types.NewZendArray(ce.GetDefaultPropertiesCount()))
		if ce.GetDefaultPropertiesCount() != 0 {
			types.ZendHashRealInitMixed(zobj.GetProperties())
			var __ht *types.Array = ce.GetPropertiesInfo()
			for _, _p := range __ht.foreachData() {
				var _z *types.Zval = _p.GetVal()

				prop_info = _z.GetPtr()
				if !prop_info.IsStatic() {
					flags |= prop_info.GetFlags()
					if OBJ_PROP(zobj, prop_info.GetOffset()).IsUndef() {
						zobj.GetProperties().GetUFlags() |= types.HASH_FLAG_HAS_EMPTY_IND
					}
					types._zendHashAppendInd(zobj.GetProperties(), prop_info.GetName(), OBJ_PROP(zobj, prop_info.GetOffset()))
				}
			}
			if (flags & ZEND_ACC_CHANGED) != 0 {
				for ce.GetParent() && ce.GetParent().default_properties_count {
					ce = ce.GetParent()
					var __ht *types.Array = ce.GetPropertiesInfo()
					for _, _p := range __ht.foreachData() {
						var _z *types.Zval = _p.GetVal()

						prop_info = _z.GetPtr()
						if prop_info.GetCe() == ce && !prop_info.IsStatic() && prop_info.IsPrivate() {
							var zv types.Zval
							if OBJ_PROP(zobj, prop_info.GetOffset()).IsUndef() {
								zobj.GetProperties().GetUFlags() |= types.HASH_FLAG_HAS_EMPTY_IND
							}
							zv.SetIndirect(OBJ_PROP(zobj, prop_info.GetOffset()))
							zobj.GetProperties().KeyAdd(prop_info.GetName().GetStr(), &zv)
						}
					}
				}
			}
		}
	}
}
func ZendStdGetProperties(object *types.Zval) *types.Array {
	var zobj *types.ZendObject
	zobj = object.GetObj()
	if zobj.GetProperties() == nil {
		RebuildObjectProperties(zobj)
	}
	return zobj.GetProperties()
}
func ZendStdGetGc(object *types.Zval, table **types.Zval, n *int) *types.Array {
	if types.Z_OBJ_HT(*object).GetGetProperties() != ZendStdGetProperties {
		*table = nil
		*n = 0
		return types.Z_OBJ_HT(*object).GetGetProperties()(object)
	} else {
		var zobj *types.ZendObject = object.GetObj()
		if zobj.GetProperties() != nil {
			*table = nil
			*n = 0
			if zobj.GetProperties().GetRefcount() > 1 && (zobj.GetProperties().GetGcFlags()&types.IS_ARRAY_IMMUTABLE) == 0 {
				zobj.GetProperties().DelRefcount()
				zobj.SetProperties(types.ZendArrayDup(zobj.GetProperties()))
			}
			return zobj.GetProperties()
		} else {
			*table = zobj.GetPropertiesTable()
			*n = zobj.GetCe().GetDefaultPropertiesCount()
			return nil
		}
	}
}
func ZendStdGetDebugInfo(object *types.Zval, is_temp *int) *types.Array {
	var ce *types.ClassEntry = types.Z_OBJCE_P(object)
	var retval types.Zval
	var ht *types.Array
	if ce.GetDebugInfo() == nil {
		*is_temp = 0
		return types.Z_OBJ_HT(*object).GetGetProperties()(object)
	}
	ZendCallMethodWith0Params(object, ce, ce.GetDebugInfo(), ZEND_DEBUGINFO_FUNC_NAME, &retval)
	if retval.IsArray() {
		if !(retval.IsRefcounted()) {
			*is_temp = 1
			return types.ZendArrayDup(retval.GetArr())
		} else if retval.GetRefcount() <= 1 {
			*is_temp = 1
			ht = retval.GetArr()
			return ht
		} else {
			*is_temp = 0
			ZvalPtrDtor(&retval)
			return retval.GetArr()
		}
	} else if retval.IsNull() {
		*is_temp = 1
		ht = types.NewZendArray(0)
		return ht
	}
	faults.ErrorNoreturn(faults.E_ERROR, ZEND_DEBUGINFO_FUNC_NAME+"() must return an array")
	return nil
}
func ZendStdCallGetter(zobj *types.ZendObject, prop_name *types.String, retval *types.Zval) {
	var ce *types.ClassEntry = zobj.GetCe()
	var orig_fake_scope *types.ClassEntry = EG__().GetFakeScope()
	var fci types.ZendFcallInfo
	var fcic types.ZendFcallInfoCache
	var member types.Zval
	EG__().SetFakeScope(nil)

	/* __get handler is called with one argument:
	      property name

	   it should return whether the call was successful or not
	*/

	member.SetString(prop_name)
	fci.SetSize(b.SizeOf("fci"))
	fci.SetObject(zobj)
	fci.SetRetval(retval)
	fci.SetParamCount(1)
	fci.SetParams(&member)
	fci.SetNoSeparation(1)
	fci.GetFunctionName().SetUndef()
	fcic.SetFunctionHandler(ce.GetGet())
	fcic.SetCalledScope(ce)
	fcic.SetObject(zobj)
	ZendCallFunction(&fci, &fcic)
	EG__().SetFakeScope(orig_fake_scope)
}
func ZendStdCallSetter(zobj *types.ZendObject, prop_name *types.String, value *types.Zval) {
	var ce *types.ClassEntry = zobj.GetCe()
	var orig_fake_scope *types.ClassEntry = EG__().GetFakeScope()
	var fci types.ZendFcallInfo
	var fcic types.ZendFcallInfoCache
	var args []types.Zval
	var ret types.Zval
	EG__().SetFakeScope(nil)

	/* __set handler is called with two arguments:
	   property name
	   value to be set
	*/

	args[0].SetString(prop_name)
	types.ZVAL_COPY_VALUE(&args[1], value)
	ret.SetUndef()
	fci.SetSize(b.SizeOf("fci"))
	fci.SetObject(zobj)
	fci.SetRetval(&ret)
	fci.SetParamCount(2)
	fci.SetParams(args)
	fci.SetNoSeparation(1)
	fci.GetFunctionName().SetUndef()
	fcic.SetFunctionHandler(ce.GetSet())
	fcic.SetCalledScope(ce)
	fcic.SetObject(zobj)
	ZendCallFunction(&fci, &fcic)
	ZvalPtrDtor(&ret)
	EG__().SetFakeScope(orig_fake_scope)
}
func ZendStdCallUnsetter(zobj *types.ZendObject, prop_name *types.String) {
	var ce *types.ClassEntry = zobj.GetCe()
	var orig_fake_scope *types.ClassEntry = EG__().GetFakeScope()
	var fci types.ZendFcallInfo
	var fcic types.ZendFcallInfoCache
	var ret types.Zval
	var member types.Zval
	EG__().SetFakeScope(nil)

	/* __unset handler is called with one argument:
	   property name
	*/

	member.SetString(prop_name)
	ret.SetUndef()
	fci.SetSize(b.SizeOf("fci"))
	fci.SetObject(zobj)
	fci.SetRetval(&ret)
	fci.SetParamCount(1)
	fci.SetParams(&member)
	fci.SetNoSeparation(1)
	fci.GetFunctionName().SetUndef()
	fcic.SetFunctionHandler(ce.GetUnset())
	fcic.SetCalledScope(ce)
	fcic.SetObject(zobj)
	ZendCallFunction(&fci, &fcic)
	ZvalPtrDtor(&ret)
	EG__().SetFakeScope(orig_fake_scope)
}
func ZendStdCallIssetter(zobj *types.ZendObject, prop_name *types.String, retval *types.Zval) {
	var ce *types.ClassEntry = zobj.GetCe()
	var orig_fake_scope *types.ClassEntry = EG__().GetFakeScope()
	var fci types.ZendFcallInfo
	var fcic types.ZendFcallInfoCache
	var member types.Zval
	EG__().SetFakeScope(nil)

	/* __isset handler is called with one argument:
	      property name

	   it should return whether the property is set or not
	*/

	member.SetString(prop_name)
	fci.SetSize(b.SizeOf("fci"))
	fci.SetObject(zobj)
	fci.SetRetval(retval)
	fci.SetParamCount(1)
	fci.SetParams(&member)
	fci.SetNoSeparation(1)
	fci.GetFunctionName().SetUndef()
	fcic.SetFunctionHandler(ce.GetIsset())
	fcic.SetCalledScope(ce)
	fcic.SetObject(zobj)
	ZendCallFunction(&fci, &fcic)
	EG__().SetFakeScope(orig_fake_scope)
}
func IsDerivedClass(child_class *types.ClassEntry, parent_class *types.ClassEntry) types.ZendBool {
	child_class = child_class.GetParent()
	for child_class != nil {
		if child_class == parent_class {
			return 1
		}
		child_class = child_class.GetParent()
	}
	return 0
}
func IsProtectedCompatibleScope(ce *types.ClassEntry, scope *types.ClassEntry) int {
	return scope != nil && (IsDerivedClass(ce, scope) != 0 || IsDerivedClass(scope, ce) != 0)
}
func ZendGetParentPrivateProperty(scope *types.ClassEntry, ce *types.ClassEntry, member *types.String) *ZendPropertyInfo {
	var zv *types.Zval
	var prop_info *ZendPropertyInfo
	if scope != ce && scope != nil && IsDerivedClass(ce, scope) != 0 {
		zv = scope.GetPropertiesInfo().KeyFind(member.GetStr())
		if zv != nil {
			prop_info = (*ZendPropertyInfo)(zv.GetPtr())
			if prop_info.IsPrivate() && prop_info.GetCe() == scope {
				return prop_info
			}
		}
	}
	return nil
}
func ZendBadPropertyAccess(property_info *ZendPropertyInfo, ce *types.ClassEntry, member *types.String) {
	faults.ThrowError(nil, "Cannot access %s property %s::$%s", ZendVisibilityString(property_info.GetFlags()), ce.GetName().GetVal(), member.GetVal())
}
func ZendBadPropertyName() {
	faults.ThrowError(nil, "Cannot access property started with '\\0'")
}
func ZendGetPropertyOffset(ce *types.ClassEntry, member *types.String, silent int, cache_slot *any, info_ptr **ZendPropertyInfo) uintPtr {
	var zv *types.Zval
	var property_info *ZendPropertyInfo
	var flags uint32
	var scope *types.ClassEntry
	var offset uintPtr
	if cache_slot != nil && ce == CACHED_PTR_EX(cache_slot) {
		*info_ptr = CACHED_PTR_EX(cache_slot + 2)
		return uintPtr(CACHED_PTR_EX(cache_slot + 1))
	}
	if ce.GetPropertiesInfo().GetNNumOfElements() == 0 || b.Assign(&zv, ce.GetPropertiesInfo().KeyFind(member.GetStr())) == nil {
		if member.GetVal()[0] == '0' && member.GetLen() != 0 {
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
	property_info = (*ZendPropertyInfo)(zv.GetPtr())
	flags = property_info.GetFlags()
	if (flags & (ZEND_ACC_CHANGED | ZEND_ACC_PRIVATE | ZEND_ACC_PROTECTED)) != 0 {
		if EG__().GetFakeScope() != nil {
			scope = EG__().GetFakeScope()
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

				if p != nil && (!p.IsStatic() || (flags&ZEND_ACC_STATIC) != 0) {
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
				b.Assert((flags & ZEND_ACC_PROTECTED) != 0)
				if IsProtectedCompatibleScope(property_info.GetCe(), scope) == 0 {
					goto wrong
				}
			}
		}
	}
found:
	if (flags & ZEND_ACC_STATIC) != 0 {
		if silent == 0 {
			faults.Error(faults.E_NOTICE, "Accessing static property %s::$%s as non static", ce.GetName().GetVal(), member.GetVal())
		}
		return ZEND_DYNAMIC_PROPERTY_OFFSET
	}
	offset = property_info.GetOffset()
	if property_info.GetType() == 0 {
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
func ZendWrongOffset(ce *types.ClassEntry, member *types.String) {
	var dummy *ZendPropertyInfo

	/* Trigger the correct error */

	ZendGetPropertyOffset(ce, member, 0, nil, &dummy)

	/* Trigger the correct error */
}
func ZendGetPropertyInfo(ce *types.ClassEntry, member *types.String, silent int) *ZendPropertyInfo {
	var zv *types.Zval
	var property_info *ZendPropertyInfo
	var flags uint32
	var scope *types.ClassEntry
	if ce.GetPropertiesInfo().GetNNumOfElements() == 0 || b.Assign(&zv, ce.GetPropertiesInfo().KeyFind(member.GetStr())) == nil {
		if member.GetVal()[0] == '0' && member.GetLen() != 0 {
			if silent == 0 {
				ZendBadPropertyName()
			}
			return ZEND_WRONG_PROPERTY_INFO
		}
	dynamic:
		return nil
	}
	property_info = (*ZendPropertyInfo)(zv.GetPtr())
	flags = property_info.GetFlags()
	if (flags & (ZEND_ACC_CHANGED | ZEND_ACC_PRIVATE | ZEND_ACC_PROTECTED)) != 0 {
		if EG__().GetFakeScope() != nil {
			scope = EG__().GetFakeScope()
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
				b.Assert((flags & ZEND_ACC_PROTECTED) != 0)
				if IsProtectedCompatibleScope(property_info.GetCe(), scope) == 0 {
					goto wrong
				}
			}
		}
	}
found:
	if (flags & ZEND_ACC_STATIC) != 0 {
		if silent == 0 {
			faults.Error(faults.E_NOTICE, "Accessing static property %s::$%s as non static", ce.GetName().GetVal(), member.GetVal())
		}
	}
	return property_info
}
func ZendCheckPropertyAccess(zobj *types.ZendObject, prop_info_name *types.String, is_dynamic types.ZendBool) int {
	var property_info *ZendPropertyInfo
	var class_name *byte = nil
	var prop_name *byte
	var member *types.String
	var prop_name_len int
	if prop_info_name.GetVal()[0] == 0 {
		if is_dynamic != 0 {
			return types.SUCCESS
		}
		ZendUnmanglePropertyNameEx(prop_info_name, &class_name, &prop_name, &prop_name_len)
		member = types.NewString(b.CastStr(prop_name, prop_name_len))
		property_info = ZendGetPropertyInfo(zobj.GetCe(), member, 1)
		types.ZendStringReleaseEx(member, 0)
		if property_info == nil || property_info == ZEND_WRONG_PROPERTY_INFO {
			return types.FAILURE
		}
		if class_name[0] != '*' {
			if !property_info.IsPrivate() {

				/* we we're looking for a private prop but found a non private one of the same name */

				return types.FAILURE

				/* we we're looking for a private prop but found a non private one of the same name */

			} else if strcmp(prop_info_name.GetVal()+1, property_info.GetName().GetVal()+1) {

				/* we we're looking for a private prop but found a private one of the same name but another class */

				return types.FAILURE

				/* we we're looking for a private prop but found a private one of the same name but another class */

			}
		} else {
			b.Assert(property_info.IsProtected())
		}
		return types.SUCCESS
	} else {
		property_info = ZendGetPropertyInfo(zobj.GetCe(), prop_info_name, 1)
		if property_info == nil {
			b.Assert(is_dynamic != 0)
			return types.SUCCESS
		} else if property_info == ZEND_WRONG_PROPERTY_INFO {
			return types.FAILURE
		}
		if property_info.IsPublic() {
			return types.SUCCESS
		} else {
			return types.FAILURE
		}
	}
}
func ZendPropertyGuardDtor(el *types.Zval) {
	var ptr *uint32 = (*uint32)(el.GetPtr())
	if (types.ZendUintptrT(ptr) & 1) == 0 {
		EfreeSize(ptr, b.SizeOf("uint32_t"))
	}
}
func ZendGetPropertyGuard(zobj *types.ZendObject, member *types.String) *uint32 {
	var guards *types.Array
	var zv *types.Zval
	var ptr *uint32
	b.Assert(zobj.GetCe().IsUseGuards())
	zv = zobj.GetPropertiesTable() + zobj.GetCe().GetDefaultPropertiesCount()
	if zv.IsString() {
		var str *types.String = zv.GetStr()
		if str == member || str.GetH() == member.GetHash() && types.ZendStringEqualContent(str, member) != 0 {
			return &(zv.GetPropertyGuard())
		} else if zv.GetPropertyGuard() == 0 {
			ZvalPtrDtorStr(zv)
			zv.SetStringCopy(member)
			return &(zv.GetPropertyGuard())
		} else {
			ALLOC_HASHTABLE(guards)
			guards = types.MakeArrayEx(8, ZendPropertyGuardDtor, 0)

			/* mark pointer as "special" using low bit */

			types.ZendHashAddNewPtr(guards, str.GetStr(), any(zend_uintptr_t&zv.GetPropertyGuard()|1))
			ZvalPtrDtorStr(zv)
			zv.SetArray(guards)
		}
	} else if zv.IsArray() {
		guards = zv.GetArr()
		b.Assert(guards != nil)
		zv = guards.KeyFind(member.GetStr())
		if zv != nil {
			return (*uint32)(types.ZendUintptrT(zv.GetPtr()) & ^1)
		}
	} else {
		b.Assert(zv.IsUndef())
		zv.SetStringCopy(member)
		zv.SetPropertyGuard(0)
		return &(zv.GetPropertyGuard())
	}

	/* we have to allocate uint32_t separately because ht->arData may be reallocated */

	ptr = (*uint32)(Emalloc(b.SizeOf("uint32_t")))
	*ptr = 0
	return (*uint32)(types.ZendHashAddNewPtr(guards, member.GetStr(), ptr))
}
func ZendStdReadProperty(object *types.Zval, member *types.Zval, type_ int, cache_slot *any, rv *types.Zval) *types.Zval {
	var zobj *types.ZendObject
	var name *types.String
	var tmp_name *types.String
	var retval *types.Zval
	var property_offset uintPtr
	var prop_info *ZendPropertyInfo = nil
	var guard *uint32 = nil
	zobj = object.GetObj()
	name = ZvalTryGetTmpString(member, &tmp_name)
	if name == nil {
		return EG__().GetUninitializedZval()
	}

	/* make zend_get_property_info silent if we have getter - we may want to use it */

	property_offset = ZendGetPropertyOffset(zobj.GetCe(), name, type_ == BP_VAR_IS || zobj.GetCe().GetGet() != nil, cache_slot, &prop_info)
	if IS_VALID_PROPERTY_OFFSET(property_offset) {
		retval = OBJ_PROP(zobj, property_offset)
		if retval.GetType() != types.IS_UNDEF {
			goto exit
		}
		if retval.GetU2Extra() == types.IS_PROP_UNINIT {

			/* Skip __get() for uninitialized typed properties */

			goto uninit_error

			/* Skip __get() for uninitialized typed properties */

		}
	} else if IS_DYNAMIC_PROPERTY_OFFSET(property_offset) {
		if zobj.GetProperties() != nil {
			if !(IS_UNKNOWN_DYNAMIC_PROPERTY_OFFSET(property_offset)) {
				var idx uintPtr = ZEND_DECODE_DYN_PROP_OFFSET(property_offset)
				if idx < zobj.GetProperties().GetNNumUsed()*b.SizeOf("Bucket") {
					var p *types.Bucket = (*types.Bucket)((*byte)(zobj.GetProperties().GetArData() + idx))
					if p.GetVal().GetType() != types.IS_UNDEF && (p.GetKey() == name || p.GetH() == name.GetH() && p.GetKey() != nil && types.ZendStringEqualContent(p.GetKey(), name) != 0) {
						retval = p.GetVal()
						goto exit
					}
				}
				CACHE_PTR_EX(cache_slot+1, any(ZEND_DYNAMIC_PROPERTY_OFFSET))
			}
			retval = zobj.GetProperties().KeyFind(name.GetStr())
			if retval != nil {
				if cache_slot != nil {
					var idx uintPtr = (*byte)(retval - (*byte)(zobj.GetProperties().GetArData()))
					CACHE_PTR_EX(cache_slot+1, any(ZEND_ENCODE_DYN_PROP_OFFSET(idx)))
				}
				goto exit
			}
		}
	} else if EG__().GetException() != nil {
		retval = EG__().GetUninitializedZval()
		goto exit
	}

	/* magic isset */

	if type_ == BP_VAR_IS && zobj.GetCe().GetIsset() != nil {
		var tmp_result types.Zval
		guard = ZendGetPropertyGuard(zobj, name)
		if ((*guard) & IN_ISSET) == 0 {
			if tmp_name == nil {
				tmp_name = name.Copy()
			}
			zobj.AddRefcount()
			tmp_result.SetUndef()
			*guard |= IN_ISSET
			ZendStdCallIssetter(zobj, name, &tmp_result)
			*guard &= ^IN_ISSET
			if ZendIsTrue(&tmp_result) == 0 {
				retval = EG__().GetUninitializedZval()
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
			zobj.AddRefcount()
		call_getter:
			*guard |= IN_GET
			ZendStdCallGetter(zobj, name, rv)
			*guard &= ^IN_GET
			if rv.GetType() != types.IS_UNDEF {
				retval = rv
				if !(rv.IsReference()) && (type_ == BP_VAR_W || type_ == BP_VAR_RW || type_ == BP_VAR_UNSET) {
					if rv.GetType() != types.IS_OBJECT {
						faults.Error(faults.E_NOTICE, "Indirect modification of overloaded property %s::$%s has no effect", zobj.GetCe().GetName().GetVal(), name.GetVal())
					}
				}
			} else {
				retval = EG__().GetUninitializedZval()
			}
			if prop_info != nil {
				ZendVerifyPropAssignableByRef(prop_info, retval, zobj.GetCe().GetGet().IsStrictTypes())
			}
			OBJ_RELEASE(zobj)
			goto exit
		} else if IS_WRONG_PROPERTY_OFFSET(property_offset) {

			/* Trigger the correct error */

			ZendGetPropertyOffset(zobj.GetCe(), name, 0, nil, &prop_info)
			b.Assert(EG__().GetException() != nil)
			retval = EG__().GetUninitializedZval()
			goto exit
		}
	}
uninit_error:
	if type_ != BP_VAR_IS {
		if prop_info != nil {
			faults.ThrowError(nil, "Typed property %s::$%s must not be accessed before initialization", prop_info.GetCe().GetName().GetVal(), name.GetVal())
		} else {
			faults.Error(faults.E_NOTICE, "Undefined property: %s::$%s", zobj.GetCe().GetName().GetVal(), name.GetVal())
		}
	}
	retval = EG__().GetUninitializedZval()
exit:
	ZendTmpStringRelease(tmp_name)
	return retval
}
func PropertyUsesStrictTypes() types.ZendBool {
	var executeData *ZendExecuteData = CurrEX()
	return executeData != nil && executeData.GetFunc() != nil && CurrEX().IsCallUseStrictTypes()
}
func ZendStdWriteProperty(object *types.Zval, member *types.Zval, value *types.Zval, cache_slot *any) *types.Zval {
	var zobj *types.ZendObject
	var name *types.String
	var tmp_name *types.String
	var variable_ptr *types.Zval
	var tmp types.Zval
	var property_offset uintPtr
	var prop_info *ZendPropertyInfo = nil
	b.Assert(!(value.IsReference()))
	zobj = object.GetObj()
	name = ZvalTryGetTmpString(member, &tmp_name)
	if name == nil {
		return value
	}
	property_offset = ZendGetPropertyOffset(zobj.GetCe(), name, zobj.GetCe().GetSet() != nil, cache_slot, &prop_info)
	if IS_VALID_PROPERTY_OFFSET(property_offset) {
		variable_ptr = OBJ_PROP(zobj, property_offset)
		if variable_ptr.GetType() != types.IS_UNDEF {
			value.TryAddRefcount()
			if prop_info != nil {
				types.ZVAL_COPY_VALUE(&tmp, value)
				if ZendVerifyPropertyType(prop_info, &tmp, PropertyUsesStrictTypes()) == 0 {
					value.TryDelRefcount()
					variable_ptr = EG__().GetErrorZval()
					goto exit
				}
				value = &tmp
			}
		found:
			variable_ptr = ZendAssignToVariable(variable_ptr, value, IS_TMP_VAR, PropertyUsesStrictTypes())
			goto exit
		}
		if variable_ptr.GetU2Extra() == types.IS_PROP_UNINIT {

			/* Writes to uninitializde typed properties bypass __set(). */

			variable_ptr.SetU2Extra(0)
			goto write_std_property
		}
	} else if IS_DYNAMIC_PROPERTY_OFFSET(property_offset) {
		if zobj.GetProperties() != nil {
			if zobj.GetProperties().GetRefcount() > 1 {
				if (zobj.GetProperties().GetGcFlags() & types.IS_ARRAY_IMMUTABLE) == 0 {
					zobj.GetProperties().DelRefcount()
				}
				zobj.SetProperties(types.ZendArrayDup(zobj.GetProperties()))
			}
			if b.Assign(&variable_ptr, zobj.GetProperties().KeyFind(name.GetStr())) != nil {
				value.TryAddRefcount()
				goto found
			}
		}
	} else if EG__().GetException() != nil {
		variable_ptr = EG__().GetErrorZval()
		goto exit
	}

	/* magic set */

	if zobj.GetCe().GetSet() != nil {
		var guard *uint32 = ZendGetPropertyGuard(zobj, name)
		if ((*guard) & IN_SET) == 0 {
			zobj.AddRefcount()
			*guard |= IN_SET
			ZendStdCallSetter(zobj, name, value)
			*guard &= ^IN_SET
			OBJ_RELEASE(zobj)
			variable_ptr = value
		} else if !(IS_WRONG_PROPERTY_OFFSET(property_offset)) {
			goto write_std_property
		} else {

			/* Trigger the correct error */

			ZendWrongOffset(zobj.GetCe(), name)
			b.Assert(EG__().GetException() != nil)
			variable_ptr = EG__().GetErrorZval()
			goto exit
		}
	} else {
		b.Assert(!(IS_WRONG_PROPERTY_OFFSET(property_offset)))
	write_std_property:
		value.TryAddRefcount()
		if IS_VALID_PROPERTY_OFFSET(property_offset) {
			variable_ptr = OBJ_PROP(zobj, property_offset)
			if prop_info != nil {
				types.ZVAL_COPY_VALUE(&tmp, value)
				if ZendVerifyPropertyType(prop_info, &tmp, PropertyUsesStrictTypes()) == 0 {
					ZvalPtrDtor(value)
					goto exit
				}
				value = &tmp
				goto found
			}
			types.ZVAL_COPY_VALUE(variable_ptr, value)
		} else {
			if zobj.GetProperties() == nil {
				RebuildObjectProperties(zobj)
			}
			variable_ptr = zobj.GetProperties().KeyAddNew(name.GetStr(), value)
		}
	}
exit:
	ZendTmpStringRelease(tmp_name)
	return variable_ptr
}
func ZendBadArrayAccess(ce *types.ClassEntry) {
	faults.ThrowError(nil, "Cannot use object of type %s as array", ce.GetName().GetVal())
}
func ZendStdReadDimension(object *types.Zval, offset *types.Zval, type_ int, rv *types.Zval) *types.Zval {
	var ce *types.ClassEntry = types.Z_OBJCE_P(object)
	var tmp_offset types.Zval
	var tmp_object types.Zval
	if InstanceofFunctionEx(ce, ZendCeArrayaccess, 1) != 0 {
		if offset == nil {

			/* [] construct */

			tmp_offset.SetNull()

			/* [] construct */

		} else {
			types.ZVAL_COPY_DEREF(&tmp_offset, offset)
		}
		object.AddRefcount()
		tmp_object.SetObject(object.GetObj())
		if type_ == BP_VAR_IS {
			ZendCallMethodWith1Params(&tmp_object, ce, nil, "offsetexists", rv, &tmp_offset)
			if rv.IsUndef() {
				ZvalPtrDtor(&tmp_object)
				ZvalPtrDtor(&tmp_offset)
				return nil
			}
			if IZendIsTrue(rv) == 0 {
				ZvalPtrDtor(&tmp_object)
				ZvalPtrDtor(&tmp_offset)
				ZvalPtrDtor(rv)
				return EG__().GetUninitializedZval()
			}
			ZvalPtrDtor(rv)
		}
		ZendCallMethodWith1Params(&tmp_object, ce, nil, "offsetget", rv, &tmp_offset)
		ZvalPtrDtor(&tmp_object)
		ZvalPtrDtor(&tmp_offset)
		if rv.IsUndef() {
			if EG__().GetException() == nil {
				faults.ThrowError(nil, "Undefined offset for object of type %s used as array", ce.GetName().GetVal())
			}
			return nil
		}
		return rv
	} else {
		ZendBadArrayAccess(ce)
		return nil
	}
}
func ZendStdWriteDimension(object *types.Zval, offset *types.Zval, value *types.Zval) {
	var ce *types.ClassEntry = types.Z_OBJCE_P(object)
	var tmp_offset types.Zval
	var tmp_object types.Zval
	if InstanceofFunctionEx(ce, ZendCeArrayaccess, 1) != 0 {
		if offset == nil {
			tmp_offset.SetNull()
		} else {
			types.ZVAL_COPY_DEREF(&tmp_offset, offset)
		}
		object.AddRefcount()
		tmp_object.SetObject(object.GetObj())
		ZendCallMethodWith2Params(&tmp_object, ce, nil, "offsetset", nil, &tmp_offset, value)
		ZvalPtrDtor(&tmp_object)
		ZvalPtrDtor(&tmp_offset)
	} else {
		ZendBadArrayAccess(ce)
	}
}
func ZendStdHasDimension(object *types.Zval, offset *types.Zval, check_empty int) int {
	var ce *types.ClassEntry = types.Z_OBJCE_P(object)
	var retval types.Zval
	var tmp_offset types.Zval
	var tmp_object types.Zval
	var result int
	if InstanceofFunctionEx(ce, ZendCeArrayaccess, 1) != 0 {
		types.ZVAL_COPY_DEREF(&tmp_offset, offset)
		object.AddRefcount()
		tmp_object.SetObject(object.GetObj())
		ZendCallMethodWith1Params(&tmp_object, ce, nil, "offsetexists", &retval, &tmp_offset)
		result = IZendIsTrue(&retval)
		ZvalPtrDtor(&retval)
		if check_empty != 0 && result != 0 && EG__().GetException() == nil {
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
func ZendStdGetPropertyPtrPtr(object *types.Zval, member *types.Zval, type_ int, cache_slot *any) *types.Zval {
	var zobj *types.ZendObject
	var name *types.String
	var tmp_name *types.String
	var retval *types.Zval = nil
	var property_offset uintPtr
	var prop_info *ZendPropertyInfo = nil
	zobj = object.GetObj()
	name = ZvalTryGetTmpString(member, &tmp_name)
	if name == nil {
		return EG__().GetErrorZval()
	}
	property_offset = ZendGetPropertyOffset(zobj.GetCe(), name, zobj.GetCe().GetGet() != nil, cache_slot, &prop_info)
	if IS_VALID_PROPERTY_OFFSET(property_offset) {
		retval = OBJ_PROP(zobj, property_offset)
		if retval.IsUndef() {
			if zobj.GetCe().GetGet() == nil || ((*ZendGetPropertyGuard)(zobj, name)&IN_GET) != 0 || prop_info != nil && retval.GetU2Extra() == types.IS_PROP_UNINIT {
				if type_ == BP_VAR_RW || type_ == BP_VAR_R {
					if prop_info != nil {
						faults.ThrowError(nil, "Typed property %s::$%s must not be accessed before initialization", prop_info.GetCe().GetName().GetVal(), name.GetVal())
						retval = EG__().GetErrorZval()
					} else {
						retval.SetNull()
						faults.Error(faults.E_NOTICE, "Undefined property: %s::$%s", zobj.GetCe().GetName().GetVal(), name.GetVal())
					}
				}
			} else {

				/* we do have getter - fail and let it try again with usual get/set */

				retval = nil

				/* we do have getter - fail and let it try again with usual get/set */

			}
		}
	} else if IS_DYNAMIC_PROPERTY_OFFSET(property_offset) {
		if zobj.GetProperties() != nil {
			if zobj.GetProperties().GetRefcount() > 1 {
				if (zobj.GetProperties().GetGcFlags() & types.IS_ARRAY_IMMUTABLE) == 0 {
					zobj.GetProperties().DelRefcount()
				}
				zobj.SetProperties(types.ZendArrayDup(zobj.GetProperties()))
			}
			if b.Assign(&retval, zobj.GetProperties().KeyFind(name.GetStr())) != nil {
				ZendTmpStringRelease(tmp_name)
				return retval
			}
		}
		if zobj.GetCe().GetGet() == nil || ((*ZendGetPropertyGuard)(zobj, name)&IN_GET) != 0 {
			if zobj.GetProperties() == nil {
				RebuildObjectProperties(zobj)
			}
			retval = zobj.GetProperties().KeyUpdate(name.GetStr(), EG__().GetUninitializedZval())

			/* Notice is thrown after creation of the property, to avoid EG(std_property_info)
			 * being overwritten in an error handler. */

			if type_ == BP_VAR_RW || type_ == BP_VAR_R {
				faults.Error(faults.E_NOTICE, "Undefined property: %s::$%s", zobj.GetCe().GetName().GetVal(), name.GetVal())
			}

			/* Notice is thrown after creation of the property, to avoid EG(std_property_info)
			 * being overwritten in an error handler. */

		}
	} else if zobj.GetCe().GetGet() == nil {
		retval = EG__().GetErrorZval()
	}
	ZendTmpStringRelease(tmp_name)
	return retval
}
func ZendStdUnsetProperty(object *types.Zval, member *types.Zval, cache_slot *any) {
	var zobj *types.ZendObject
	var name *types.String
	var tmp_name *types.String
	var property_offset uintPtr
	var prop_info *ZendPropertyInfo = nil
	zobj = object.GetObj()
	name = ZvalTryGetTmpString(member, &tmp_name)
	if name == nil {
		return
	}
	property_offset = ZendGetPropertyOffset(zobj.GetCe(), name, zobj.GetCe().GetUnset() != nil, cache_slot, &prop_info)
	if IS_VALID_PROPERTY_OFFSET(property_offset) {
		var slot *types.Zval = OBJ_PROP(zobj, property_offset)
		if slot.GetType() != types.IS_UNDEF {
			if slot.IsReference() && ZEND_REF_HAS_TYPE_SOURCES(slot.GetRef()) {
				if prop_info != nil {
					ZEND_REF_DEL_TYPE_SOURCE(slot.GetRef(), prop_info)
				}
			}
			var tmp types.Zval
			types.ZVAL_COPY_VALUE(&tmp, slot)
			slot.SetUndef()
			ZvalPtrDtor(&tmp)
			if zobj.GetProperties() != nil {
				zobj.GetProperties().GetUFlags() |= types.HASH_FLAG_HAS_EMPTY_IND
			}
			goto exit
		}
		if slot.GetU2Extra() == types.IS_PROP_UNINIT {

			/* Reset the IS_PROP_UNINIT flag, if it exists and bypass __unset(). */

			slot.SetU2Extra(0)
			goto exit
		}
	} else if IS_DYNAMIC_PROPERTY_OFFSET(property_offset) && zobj.GetProperties() != nil {
		if zobj.GetProperties().GetRefcount() > 1 {
			if (zobj.GetProperties().GetGcFlags() & types.IS_ARRAY_IMMUTABLE) == 0 {
				zobj.GetProperties().DelRefcount()
			}
			zobj.SetProperties(types.ZendArrayDup(zobj.GetProperties()))
		}
		if types.ZendHashDel(zobj.GetProperties(), name) != types.FAILURE {
			goto exit
		}
	} else if EG__().GetException() != nil {
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
		} else if IS_WRONG_PROPERTY_OFFSET(property_offset) {

			/* Trigger the correct error */

			ZendWrongOffset(zobj.GetCe(), name)
			b.Assert(EG__().GetException() != nil)
			goto exit
		}
	}
exit:
	ZendTmpStringRelease(tmp_name)
}
func ZendStdUnsetDimension(object *types.Zval, offset *types.Zval) {
	var ce *types.ClassEntry = types.Z_OBJCE_P(object)
	var tmp_offset types.Zval
	var tmp_object types.Zval
	if InstanceofFunctionEx(ce, ZendCeArrayaccess, 1) != 0 {
		types.ZVAL_COPY_DEREF(&tmp_offset, offset)
		object.AddRefcount()
		tmp_object.SetObject(object.GetObj())
		ZendCallMethodWith1Params(&tmp_object, ce, nil, "offsetunset", nil, &tmp_offset)
		ZvalPtrDtor(&tmp_object)
		ZvalPtrDtor(&tmp_offset)
	} else {
		ZendBadArrayAccess(ce)
	}
}
func ZendGetParentPrivateMethod(scope *types.ClassEntry, ce *types.ClassEntry, function_name *types.String) *ZendFunction {
	var func_ *types.Zval
	var fbc *ZendFunction
	if scope != ce && scope != nil && IsDerivedClass(ce, scope) != 0 {
		func_ = scope.GetFunctionTable().KeyFind(function_name.GetStr())
		if func_ != nil {
			fbc = func_.GetFunc()
			if fbc.IsPrivate() && fbc.GetScope() == scope {
				return fbc
			}
		}
	}
	return nil
}
func ZendCheckProtected(ce *types.ClassEntry, scope *types.ClassEntry) int {
	var fbc_scope *types.ClassEntry = ce

	/* Is the context that's calling the function, the same as one of
	 * the function's parents?
	 */

	for fbc_scope != nil {
		if fbc_scope == scope {
			return 1
		}
		fbc_scope = fbc_scope.GetParent()
	}

	/* Is the function's scope the same as our current object context,
	 * or any of the parents of our context?
	 */

	for scope != nil {
		if scope == ce {
			return 1
		}
		scope = scope.GetParent()
	}
	return 0
}
func ZendGetCallTrampolineFunc(ce *types.ClassEntry, method_name *types.String, is_static int) *ZendFunction {
	var mname_len int
	var func_ *ZendOpArray
	var fbc *ZendFunction = b.CondF(is_static != 0, func() *ZendFunction { return ce.GetCallstatic() }, func() *ZendFunction { return ce.GetCall() })

	/* We use non-NULL value to avoid useless run_time_cache allocation.
	 * The low bit must be zero, to not be interpreted as a MAP_PTR offset.
	 */

	var dummy any = any(intPtr(2))
	b.Assert(fbc != nil)
	if EG__().GetTrampoline().GetFunctionName() == nil {
		func_ = EG__().GetTrampoline().GetOpArray()
	} else {
		func_ = Ecalloc(1, b.SizeOf("zend_op_array"))
	}
	func_.SetType(ZEND_USER_FUNCTION)
	func_.GetArgFlags()[0] = 0
	func_.GetArgFlags()[1] = 0
	func_.GetArgFlags()[2] = 0
	func_.SetFnFlags(ZEND_ACC_CALL_VIA_TRAMPOLINE | ZEND_ACC_PUBLIC)
	if is_static != 0 {
		func_.SetIsStatic(true)
	}
	func_.SetOpcodes(EG__().GetCallTrampolineOp())
	ZEND_MAP_PTR_INIT(func_.run_time_cache, (**any)(&dummy))
	func_.SetScope(fbc.GetScope())

	/* reserve space for arguments, local and temporary variables */

	if fbc.GetType() == ZEND_USER_FUNCTION {
		func_.SetT(b.Max(fbc.GetOpArray().GetLastVar()+fbc.GetOpArray().GetT(), 2))
	} else {
		func_.SetT(2)
	}
	if fbc.GetType() == ZEND_USER_FUNCTION {
		func_.SetFilename(fbc.GetOpArray().GetFilename())
	} else {
		func_.SetFilename(types.ZSTR_EMPTY_ALLOC())
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

	if b.Assign(&mname_len, strlen(method_name.GetVal())) != method_name.GetLen() {
		func_.SetFunctionName(types.NewString(b.CastStr(method_name.GetVal(), mname_len)))
	} else {
		func_.SetFunctionName(method_name.Copy())
	}
	func_.SetPrototype(nil)
	func_.SetNumArgs(0)
	func_.SetRequiredNumArgs(0)
	func_.SetArgInfo(0)
	return (*ZendFunction)(func_)
}
func ZendGetUserCallFunction(ce *types.ClassEntry, method_name *types.String) *ZendFunction {
	return ZendGetCallTrampolineFunc(ce, method_name, 0)
}
func ZendBadMethodCall(fbc *ZendFunction, method_name *types.String, scope *types.ClassEntry) {
	faults.ThrowError(nil, "Call to %s method %s::%s() from context '%s'", ZendVisibilityString(fbc.GetFnFlags()), ZEND_FN_SCOPE_NAME(fbc), method_name.GetVal(), b.CondF1(scope != nil, func() []byte { return scope.GetName().GetVal() }, ""))
}
func ZendStdGetMethod(obj_ptr **types.ZendObject, method_name *types.String, key *types.Zval) *ZendFunction {
	var zobj *types.ZendObject = *obj_ptr
	var func_ *types.Zval
	var fbc *ZendFunction
	var lc_method_name *types.String
	var scope *types.ClassEntry
	if key != nil {
		lc_method_name = key.GetStr()
	} else {
		types.ZSTR_ALLOCA_ALLOC(lc_method_name, method_name.GetLen())
		ZendStrTolowerCopy(lc_method_name.GetVal(), method_name.GetVal(), method_name.GetLen())
	}
	if b.Assign(&func_, zobj.GetCe().GetFunctionTable().KeyFind(lc_method_name.GetStr())) == nil {
		if key == nil {
			lc_method_name.Free()
		}
		if zobj.GetCe().GetCall() != nil {
			return ZendGetUserCallFunction(zobj.GetCe(), method_name)
		} else {
			return nil
		}
	}
	fbc = func_.GetFunc()

	/* Check access level */

	if fbc.GetOpArray().HasFnFlags(ZEND_ACC_CHANGED | ZEND_ACC_PRIVATE | ZEND_ACC_PROTECTED) {
		scope = ZendGetExecutedScope()
		if fbc.GetScope() != scope {
			if fbc.GetOpArray().IsChanged() {
				var updated_fbc *ZendFunction = ZendGetParentPrivateMethod(scope, zobj.GetCe(), lc_method_name)
				if updated_fbc != nil {
					fbc = updated_fbc
					goto exit
				} else if fbc.GetOpArray().IsPublic() {
					goto exit
				}
			}
			if fbc.GetOpArray().IsPrivate() || ZendCheckProtected(ZendGetFunctionRootClass(fbc), scope) == 0 {
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
	if key == nil {
		lc_method_name.Free()
	}
	return fbc
}
func ZendGetUserCallstaticFunction(ce *types.ClassEntry, method_name *types.String) *ZendFunction {
	return ZendGetCallTrampolineFunc(ce, method_name, 1)
}
func ZendStdGetStaticMethod(ce *types.ClassEntry, function_name *types.String, key *types.Zval) *ZendFunction {
	var fbc *ZendFunction = nil
	var lc_function_name *types.String
	var object *types.ZendObject
	var scope *types.ClassEntry
	if key != nil {
		lc_function_name = key.GetStr()
	} else {
		lc_function_name = ZendStringTolower(function_name)
	}
	var func_ *types.Zval = ce.GetFunctionTable().KeyFind(lc_function_name.GetStr())
	if func_ != nil {
		fbc = func_.GetFunc()
	} else if ce.GetConstructor() != nil && lc_function_name.GetLen() == ce.GetName().GetLen() && ZendBinaryStrncasecmp(lc_function_name.GetStr(), b.CastStr(ce.GetName().GetVal(), lc_function_name.GetLen()), lc_function_name.GetLen()) == 0 && (ce.GetConstructor().GetFunctionName().GetVal()[0] != '_' || ce.GetConstructor().GetFunctionName().GetVal()[1] != '_') {
		fbc = ce.GetConstructor()
	} else {
		if key == nil {
			types.ZendStringReleaseEx(lc_function_name, 0)
		}
		if ce.GetCall() != nil && b.Assign(&object, ZendGetThisObject(CurrEX())) != nil && InstanceofFunction(object.GetCe(), ce) != 0 {

			/* Call the top-level defined __call().
			 * see: tests/classes/__call_004.phpt  */

			var call_ce *types.ClassEntry = object.GetCe()
			for call_ce.GetCall() == nil {
				call_ce = call_ce.GetParent()
			}
			return ZendGetUserCallFunction(call_ce, function_name)
		} else if ce.GetCallstatic() != nil {
			return ZendGetUserCallstaticFunction(ce, function_name)
		} else {
			return nil
		}
	}
	if !fbc.GetOpArray().IsPublic() {
		scope = ZendGetExecutedScope()
		if fbc.GetScope() != scope {
			if fbc.GetOpArray().IsPrivate() || ZendCheckProtected(ZendGetFunctionRootClass(fbc), scope) == 0 {
				if ce.GetCallstatic() != nil {
					fbc = ZendGetUserCallstaticFunction(ce, function_name)
				} else {
					ZendBadMethodCall(fbc, function_name, scope)
					fbc = nil
				}
			}
		}
	}
	if key == nil {
		types.ZendStringReleaseEx(lc_function_name, 0)
	}
	return fbc
}
func ZendClassInitStatics(class_type *types.ClassEntry) {
	var i int
	var p *types.Zval
	if class_type.GetDefaultStaticMembersCount() != 0 && CE_STATIC_MEMBERS(class_type) == nil {
		if class_type.GetParent() {
			ZendClassInitStatics(class_type.GetParent())
		}
		ZEND_MAP_PTR_SET(class_type.static_members_table, Emalloc(b.SizeOf("zval")*class_type.GetDefaultStaticMembersCount()))
		for i = 0; i < class_type.GetDefaultStaticMembersCount(); i++ {
			p = class_type.GetDefaultStaticMembersTable()[i]
			if p.IsIndirect() {
				var q *types.Zval = &CE_STATIC_MEMBERS(class_type.GetParent())[i]
				q = types.ZVAL_DEINDIRECT(q)
				CE_STATIC_MEMBERS(class_type)[i].SetIndirect(q)
			} else {
				types.ZVAL_COPY_OR_DUP(&CE_STATIC_MEMBERS(class_type)[i], p)
			}
		}
	}
}
func ZendStdGetStaticPropertyWithInfo(ce *types.ClassEntry, property_name *types.String, type_ int, property_info_ptr **ZendPropertyInfo) *types.Zval {
	var ret *types.Zval
	var scope *types.ClassEntry
	var property_info *ZendPropertyInfo = types.ZendHashFindPtr(ce.GetPropertiesInfo(), property_name.GetStr())
	*property_info_ptr = property_info
	if property_info == nil {
		goto undeclared_property
	}
	if !property_info.IsPublic() {
		if EG__().GetFakeScope() != nil {
			scope = EG__().GetFakeScope()
		} else {
			scope = ZendGetExecutedScope()
		}
		if property_info.GetCe() != scope {
			if property_info.IsPrivate() || IsProtectedCompatibleScope(property_info.GetCe(), scope) == 0 {
				if type_ != BP_VAR_IS {
					ZendBadPropertyAccess(property_info, ce, property_name)
				}
				return nil
			}
		}
	}
	if !property_info.IsStatic() {
		goto undeclared_property
	}
	if !ce.IsConstantsUpdated() {
		if ZendUpdateClassConstants(ce) != types.SUCCESS {
			return nil
		}
	}

	/* check if static properties were destroyed */

	if CE_STATIC_MEMBERS(ce) == nil {
		if ce.GetType() == ZEND_INTERNAL_CLASS || ce.HasCeFlags(ZEND_ACC_IMMUTABLE|ZEND_ACC_PRELOADED) {
			ZendClassInitStatics(ce)
		} else {
		undeclared_property:
			if type_ != BP_VAR_IS {
				faults.ThrowError(nil, "Access to undeclared static property: %s::$%s", ce.GetName().GetVal(), property_name.GetVal())
			}
			return nil
		}
	}
	ret = CE_STATIC_MEMBERS(ce) + property_info.GetOffset()
	ret = types.ZVAL_DEINDIRECT(ret)
	if (type_ == BP_VAR_R || type_ == BP_VAR_RW) && ret.IsUndef() && property_info.GetType() != 0 {
		faults.ThrowError(nil, "Typed static property %s::$%s must not be accessed before initialization", property_info.GetCe().GetName().GetVal(), ZendGetUnmangledPropertyName(property_name))
		return nil
	}
	return ret
}
func ZendStdGetStaticProperty(ce *types.ClassEntry, property_name *types.String, type_ int) *types.Zval {
	var prop_info *ZendPropertyInfo
	return ZendStdGetStaticPropertyWithInfo(ce, property_name, type_, &prop_info)
}
func ZendStdUnsetStaticProperty(ce *types.ClassEntry, property_name *types.String) types.ZendBool {
	faults.ThrowError(nil, "Attempt to unset static property %s::$%s", ce.GetName().GetVal(), property_name.GetVal())
	return 0
}
func ZendBadConstructorCall(constructor *ZendFunction, scope *types.ClassEntry) {
	if scope != nil {
		faults.ThrowError(nil, "Call to %s %s::%s() from context '%s'", ZendVisibilityString(constructor.GetFnFlags()), constructor.GetScope().GetName().GetVal(), constructor.GetFunctionName().GetVal(), scope.GetName().GetVal())
	} else {
		faults.ThrowError(nil, "Call to %s %s::%s() from invalid context", ZendVisibilityString(constructor.GetFnFlags()), constructor.GetScope().GetName().GetVal(), constructor.GetFunctionName().GetVal())
	}
}
func ZendStdGetConstructor(zobj *types.ZendObject) *ZendFunction {
	var constructor *ZendFunction = zobj.GetCe().GetConstructor()
	var scope *types.ClassEntry
	if constructor != nil {
		if !constructor.GetOpArray().IsPublic() {
			if EG__().GetFakeScope() != nil {
				scope = EG__().GetFakeScope()
			} else {
				scope = ZendGetExecutedScope()
			}
			if constructor.GetScope() != scope {
				if constructor.GetOpArray().IsPrivate() || ZendCheckProtected(ZendGetFunctionRootClass(constructor), scope) == 0 {
					ZendBadConstructorCall(constructor, scope)
					constructor = nil
				}
			}
		}
	}
	return constructor
}
func ZendStdCompareObjects(o1 *types.Zval, o2 *types.Zval) int {
	var zobj1 *types.ZendObject
	var zobj2 *types.ZendObject
	zobj1 = o1.GetObj()
	zobj2 = o2.GetObj()
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

		if o1.IsRecursive() {
			faults.ErrorNoreturn(faults.E_ERROR, "Nesting level too deep - recursive dependency?")
		}
		o1.ProtectRecursive()
		var __ht *types.Array = zobj1.GetCe().GetPropertiesInfo()
		for _, _p := range __ht.foreachData() {
			var _z *types.Zval = _p.GetVal()

			info = _z.GetPtr()
			var p1 *types.Zval = OBJ_PROP(zobj1, info.GetOffset())
			var p2 *types.Zval = OBJ_PROP(zobj2, info.GetOffset())
			if info.IsStatic() {
				continue
			}
			if p1.GetType() != types.IS_UNDEF {
				if p2.GetType() != types.IS_UNDEF {
					var result types.Zval
					if CompareFunction(&result, p1, p2) == types.FAILURE {
						o1.UnprotectRecursive()
						return 1
					}
					if result.GetLval() != 0 {
						o1.UnprotectRecursive()
						return result.GetLval()
					}
				} else {
					o1.UnprotectRecursive()
					return 1
				}
			} else {
				if p2.GetType() != types.IS_UNDEF {
					o1.UnprotectRecursive()
					return 1
				}
			}
		}
		o1.UnprotectRecursive()
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
func ZendStdHasProperty(object *types.Zval, member *types.Zval, has_set_exists int, cache_slot *any) int {
	var zobj *types.ZendObject
	var result int
	var value *types.Zval = nil
	var name *types.String
	var tmp_name *types.String
	var property_offset uintPtr
	var prop_info *ZendPropertyInfo = nil
	zobj = object.GetObj()
	name = ZvalTryGetTmpString(member, &tmp_name)
	if name == nil {
		return 0
	}
	property_offset = ZendGetPropertyOffset(zobj.GetCe(), name, 1, cache_slot, &prop_info)
	if IS_VALID_PROPERTY_OFFSET(property_offset) {
		value = OBJ_PROP(zobj, property_offset)
		if value.GetType() != types.IS_UNDEF {
			goto found
		}
		if value.GetU2Extra() == types.IS_PROP_UNINIT {

			/* Skip __isset() for uninitialized typed properties */

			result = 0
			goto exit
		}
	} else if IS_DYNAMIC_PROPERTY_OFFSET(property_offset) {
		if zobj.GetProperties() != nil {
			if !(IS_UNKNOWN_DYNAMIC_PROPERTY_OFFSET(property_offset)) {
				var idx uintPtr = ZEND_DECODE_DYN_PROP_OFFSET(property_offset)
				if idx < zobj.GetProperties().GetNNumUsed()*b.SizeOf("Bucket") {
					var p *types.Bucket = (*types.Bucket)((*byte)(zobj.GetProperties().GetArData() + idx))
					if p.GetVal().GetType() != types.IS_UNDEF && (p.GetKey() == name || p.GetH() == name.GetH() && p.GetKey() != nil && types.ZendStringEqualContent(p.GetKey(), name) != 0) {
						value = p.GetVal()
						goto found
					}
				}
				CACHE_PTR_EX(cache_slot+1, any(ZEND_DYNAMIC_PROPERTY_OFFSET))
			}
			value = zobj.GetProperties().KeyFind(name.GetStr())
			if value != nil {
				if cache_slot != nil {
					var idx uintPtr = (*byte)(value - (*byte)(zobj.GetProperties().GetArData()))
					CACHE_PTR_EX(cache_slot+1, any(ZEND_ENCODE_DYN_PROP_OFFSET(idx)))
				}
			found:
				if has_set_exists == ZEND_PROPERTY_NOT_EMPTY {
					result = ZendIsTrue(value)
				} else if has_set_exists < ZEND_PROPERTY_NOT_EMPTY {
					b.Assert(has_set_exists == ZEND_PROPERTY_ISSET)
					value = types.ZVAL_DEREF(value)
					result = value.GetType() != types.IS_NULL
				} else {
					b.Assert(has_set_exists == ZEND_PROPERTY_EXISTS)
					result = 1
				}
				goto exit
			}
		}
	} else if EG__().GetException() != nil {
		result = 0
		goto exit
	}
	result = 0
	if has_set_exists != ZEND_PROPERTY_EXISTS && zobj.GetCe().GetIsset() != nil {
		var guard *uint32 = ZendGetPropertyGuard(zobj, name)
		if ((*guard) & IN_ISSET) == 0 {
			var rv types.Zval

			/* have issetter - try with it! */

			if tmp_name == nil {
				tmp_name = name.Copy()
			}
			zobj.AddRefcount()
			*guard |= IN_ISSET
			ZendStdCallIssetter(zobj, name, &rv)
			result = ZendIsTrue(&rv)
			ZvalPtrDtor(&rv)
			if has_set_exists == ZEND_PROPERTY_NOT_EMPTY && result != 0 {
				if EG__().GetException() == nil && zobj.GetCe().GetGet() != nil && ((*guard)&IN_GET) == 0 {
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
func ZendStdGetClassName(zobj *types.ZendObject) *types.String {
	return zobj.GetCe().GetName().Copy()
}
func ZendStdCastObjectTostring(readobj *types.Zval, writeobj *types.Zval, type_ int) int {
	var retval types.Zval
	var ce *types.ClassEntry
	switch type_ {
	case types.IS_STRING:
		ce = types.Z_OBJCE_P(readobj)
		if ce.GetTostring() != nil {
			var fake_scope *types.ClassEntry = EG__().GetFakeScope()
			EG__().SetFakeScope(nil)
			ZendCallMethodWith0Params(readobj, ce, ce.GetTostring(), "__tostring", &retval)
			EG__().SetFakeScope(fake_scope)
			if retval.IsString() {
				types.ZVAL_COPY_VALUE(writeobj, &retval)
				return types.SUCCESS
			}
			ZvalPtrDtor(&retval)
			if EG__().GetException() == nil {
				faults.ThrowError(nil, "Method %s::__toString() must return a string value", ce.GetName().GetVal())
			}
		}
		return types.FAILURE
	case types.IS_BOOL:
		writeobj.SetTrue()
		return types.SUCCESS
	case types.IS_LONG:
		ce = types.Z_OBJCE_P(readobj)
		faults.Error(faults.E_NOTICE, "Object of class %s could not be converted to int", ce.GetName().GetVal())
		writeobj.SetLong(1)
		return types.SUCCESS
	case types.IS_DOUBLE:
		ce = types.Z_OBJCE_P(readobj)
		faults.Error(faults.E_NOTICE, "Object of class %s could not be converted to float", ce.GetName().GetVal())
		writeobj.SetDouble(1)
		return types.SUCCESS
	case types.IS_NUMBER:
		ce = types.Z_OBJCE_P(readobj)
		faults.Error(faults.E_NOTICE, "Object of class %s could not be converted to number", ce.GetName().GetVal())
		writeobj.SetLong(1)
		return types.SUCCESS
	default:
		writeobj.SetNull()
	}
	return types.FAILURE
}
func ZendStdGetClosure(obj *types.Zval, ce_ptr **types.ClassEntry, fptr_ptr **ZendFunction, obj_ptr **types.ZendObject) int {
	var func_ *types.Zval
	var ce *types.ClassEntry = types.Z_OBJCE_P(obj)
	if b.Assign(&func_, ce.GetFunctionTable().KeyFind(types.ZSTR_MAGIC_INVOKE.GetStr())) == nil {
		return types.FAILURE
	}
	*fptr_ptr = func_.GetFunc()
	*ce_ptr = ce
	if fptr_ptr.IsStatic() {
		if obj_ptr != nil {
			*obj_ptr = nil
		}
	} else {
		if obj_ptr != nil {
			*obj_ptr = obj.GetObj()
		}
	}
	return types.SUCCESS
}
func ZendStdGetPropertiesFor(obj *types.Zval, purpose ZendPropPurpose) *types.Array {
	var ht *types.Array
	switch purpose {
	case ZEND_PROP_PURPOSE_DEBUG:
		if types.Z_OBJ_HT_P(obj).GetGetDebugInfo() != nil {
			var is_temp int
			ht = types.Z_OBJ_HT_P(obj).GetGetDebugInfo()(obj, &is_temp)
			if ht != nil && is_temp == 0 && (ht.GetGcFlags()&types.GC_IMMUTABLE) == 0 {
				ht.AddRefcount()
			}
			return ht
		}
		fallthrough
	case ZEND_PROP_PURPOSE_ARRAY_CAST:
		fallthrough
	case ZEND_PROP_PURPOSE_SERIALIZE:
		fallthrough
	case ZEND_PROP_PURPOSE_VAR_EXPORT:
		fallthrough
	case ZEND_PROP_PURPOSE_JSON:
		fallthrough
	case _ZEND_PROP_PURPOSE_ARRAY_KEY_EXISTS:
		ht = types.Z_OBJ_HT_P(obj).GetGetProperties()(obj)
		if ht != nil && (ht.GetGcFlags()&types.GC_IMMUTABLE) == 0 {
			ht.AddRefcount()
		}
		return ht
	default:
		b.Assert(false)
		return nil
	}
}
func ZendGetPropertiesFor(obj *types.Zval, purpose ZendPropPurpose) *types.Array {
	if types.Z_OBJ_HT_P(obj).GetGetPropertiesFor() != nil {
		return types.Z_OBJ_HT_P(obj).GetGetPropertiesFor()(obj, purpose)
	}
	return ZendStdGetPropertiesFor(obj, purpose)
}
