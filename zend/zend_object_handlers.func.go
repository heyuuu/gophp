package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
	"strings"
)

func PropFindAndCache(zobj *types.Object, key string) *types.Zval {
	return zobj.GetProperties().KeyFind(key)
}

func SymbolFindAndCache(symbolTable *types.Array, key string, executeData *ZendExecuteData) *types.Zval {
	retval, idx := symbolTable.KeyFindValAndPos(key)
	if retval != nil {
		/* Store "hash slot index" + 1 (NULL is a mark of uninitialized cache slot) */
		CACHE_PTR(executeData.GetOpline().GetExtendedValue(), any(idx+1))
	}
	return retval
}

func IS_VALID_PROPERTY_OFFSET(offset uintPtr) bool   { return intptr_t(offset) > 0 }
func IS_WRONG_PROPERTY_OFFSET(offset uintPtr) bool   { return intptr_t(offset) == 0 }
func IS_DYNAMIC_PROPERTY_OFFSET(offset uintPtr) bool { return intptr_t(offset) < 0 }
func IS_UNKNOWN_DYNAMIC_PROPERTY_OFFSET(offset uintPtr) bool {
	return offset == ZEND_DYNAMIC_PROPERTY_OFFSET
}
func ZEND_DECODE_DYN_PROP_OFFSET(offset int) uintptr { return uintptr(-offset - 2) }
func ZendGetFunctionRootClass(fbc types.IFunction) *types.ClassEntry {
	if fbc.GetPrototype() != nil {
		return fbc.GetPrototype().GetScope()
	} else {
		return fbc.GetScope()
	}
}
func ZendFreeTrampoline(func_ any) {
	if func_ == EG__().GetTrampoline() {
		EG__().GetTrampoline().SetFunctionName("")
	} else {
		Efree(func_)
	}
}
func RebuildObjectProperties(zobj *types.Object) {
	if zobj.GetProperties() == nil {
		var prop_info *types.PropertyInfo
		var ce *types.ClassEntry = zobj.GetCe()
		var flags uint32 = 0
		zobj.SetProperties(types.NewArrayCap(ce.GetDefaultPropertiesCount()))
		if ce.GetDefaultPropertiesCount() != 0 {
			//types.ZendHashRealInitMixed(zobj.GetProperties())

			ce.PropertyTable().Foreach(func(key string, prop_info *types.PropertyInfo) {
				if !prop_info.IsStatic() {
					flags |= prop_info.GetFlags()
					if OBJ_PROP(zobj, prop_info.GetOffset()).IsUndef() {
						zobj.GetProperties().MarkHasEmptyIndex()
					}
					types.ZendHashAppendInd(zobj.GetProperties(), prop_info.GetName(), OBJ_PROP(zobj, prop_info.GetOffset()))
				}
			})

			if (flags & types.AccChanged) != 0 {
				for ce.GetParent() && ce.GetParent().default_properties_count {
					ce = ce.GetParent()
					ce.PropertyTable().Foreach(func(key string, prop_info *types.PropertyInfo) {
						if prop_info.GetCe() == ce && !prop_info.IsStatic() && prop_info.IsPrivate() {
							var zv types.Zval
							if OBJ_PROP(zobj, prop_info.GetOffset()).IsUndef() {
								zobj.GetProperties().MarkHasEmptyIndex()
							}
							zv.SetIndirect(OBJ_PROP(zobj, prop_info.GetOffset()))
							zobj.GetProperties().KeyAdd(prop_info.GetName(), &zv)
						}
					})
				}
			}
		}
	}
}
func ZendStdGetProperties(object *types.Zval) *types.Array {
	return ZendStdGetPropertiesEx(object.Object())
}
func ZendStdGetPropertiesEx(zobj *types.Object) *types.Array {
	if zobj.GetProperties() == nil {
		RebuildObjectProperties(zobj)
	}
	return zobj.GetProperties()
}
func ZendStdCallGetter(zobj *types.Object, propName string, retval *types.Zval) {
	var ce *types.ClassEntry = zobj.GetCe()
	var origFakeScope *types.ClassEntry = EG__().GetFakeScope()
	EG__().SetFakeScope(nil)

	/* __get handler is called with one argument:
	      property name

	   it should return whether the call was successful or not
	*/
	var fci = types.InitFCallInfo(zobj, retval, types.NewZvalString(propName))

	var fcic types.ZendFcallInfoCache
	fcic.SetFunctionHandler(ce.GetGet())
	fcic.SetCalledScope(ce)
	fcic.SetObject(zobj)

	ZendCallFunction(fci, &fcic)

	EG__().SetFakeScope(origFakeScope)
}
func ZendStdCallSetter(zobj *types.Object, propName string, value *types.Zval) {
	var ce *types.ClassEntry = zobj.GetCe()
	var origFakeScope *types.ClassEntry = EG__().GetFakeScope()
	EG__().SetFakeScope(nil)

	/* __set handler is called with two arguments:
	   property name
	   value to be set
	*/

	var fci = types.InitFCallInfo(zobj, nil, types.NewZvalString(propName), value)

	var fcic types.ZendFcallInfoCache
	fcic.SetFunctionHandler(ce.GetSet())
	fcic.SetCalledScope(ce)
	fcic.SetObject(zobj)

	ZendCallFunction(fci, &fcic)

	EG__().SetFakeScope(origFakeScope)
}
func ZendStdCallUnsetter(zobj *types.Object, propName string) {
	var ce *types.ClassEntry = zobj.GetCe()
	var origFakeScope *types.ClassEntry = EG__().GetFakeScope()
	EG__().SetFakeScope(nil)

	/* __unset handler is called with one argument:
	   property name
	*/
	var fci = types.InitFCallInfo(zobj, nil, types.NewZvalString(propName))

	var fcic types.ZendFcallInfoCache
	fcic.SetFunctionHandler(ce.GetUnset())
	fcic.SetCalledScope(ce)
	fcic.SetObject(zobj)
	ZendCallFunction(fci, &fcic)

	EG__().SetFakeScope(origFakeScope)
}
func ZendStdCallIssetter(zobj *types.Object, propName string, retval *types.Zval) {
	var ce = zobj.GetCe()
	var origFakeScope = EG__().GetFakeScope()
	EG__().SetFakeScope(nil)

	/* __isset handler is called with one argument:
	      property name

	   it should return whether the property is set or not
	*/
	var fci = types.InitFCallInfo(zobj, retval, types.NewZvalString(propName))

	var fcic types.ZendFcallInfoCache
	fcic.SetFunctionHandler(ce.GetIsset())
	fcic.SetCalledScope(ce)
	fcic.SetObject(zobj)

	ZendCallFunction(fci, &fcic)
	EG__().SetFakeScope(origFakeScope)
}
func IsDerivedClass(child_class *types.ClassEntry, parent_class *types.ClassEntry) bool {
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
func ZendGetParentPrivateProperty(scope *types.ClassEntry, ce *types.ClassEntry, member string) *types.PropertyInfo {
	var prop_info *types.PropertyInfo
	if scope != ce && scope != nil && IsDerivedClass(ce, scope) != 0 {
		prop_info = scope.PropertyTable().Get(member)
		if prop_info != nil {
			if prop_info.IsPrivate() && prop_info.GetCe() == scope {
				return prop_info
			}
		}
	}
	return nil
}

func ZendBadPropertyAccess(propInfo *types.PropertyInfo, ce *types.ClassEntry, member string) {
	faults.ThrowError(nil, "Cannot access %s property %s::$%s", ZendVisibilityString(propInfo.GetFlags()), ce.Name(), member)
}
func ZendBadPropertyName() {
	faults.ThrowError(nil, "Cannot access property started with '\\0'")
}
func _zendGetPropertyOffset(ce *types.ClassEntry, member string, silent bool) (uint32, *types.ClassEntry, *types.PropertyInfo) {
	var propertyInfo *types.PropertyInfo
	var flags uint32
	var scope *types.ClassEntry
	var offset uint32

	if propertyInfo = ce.GetProperty(member); propertyInfo == nil {
		if member != "" && member[0] == '\x00' {
			goto badPropName
		}
		goto dynamic
	}
	flags = propertyInfo.GetFlags()
	if (flags & (types.AccChanged | types.AccPrivate | types.AccProtected)) != 0 {
		if EG__().GetFakeScope() != nil {
			scope = EG__().GetFakeScope()
		} else {
			scope = ZendGetExecutedScope()
		}
		if propertyInfo.GetCe() != scope {
			if (flags & types.AccChanged) != 0 {
				var p = ZendGetParentPrivateProperty(scope, ce, member)

				/* If there is a public/protected instance property on ce, don't try to use a
				 * private static property on scope. If both are static, prefer the static
				 * property on scope. This will throw a static property notice, rather than
				 * a visibility error. */
				if p != nil && (!p.IsStatic() || (flags&types.AccStatic) != 0) {
					propertyInfo = p
					flags = propertyInfo.GetFlags()
					goto found
				} else if (flags & types.AccPublic) != 0 {
					goto found
				}
			}
			if (flags & types.AccPrivate) != 0 {
				if propertyInfo.GetCe() != ce {
					goto dynamic
				} else {
					goto wrong
				}
			} else {
				b.Assert((flags & types.AccProtected) != 0)
				if IsProtectedCompatibleScope(propertyInfo.GetCe(), scope) == 0 {
					goto wrong
				}
			}
		}
	}
found:
	if (flags & types.AccStatic) != 0 {
		if !silent {
			faults.Error(faults.E_NOTICE, fmt.Sprintf("Accessing static property %s::$%s as non static", ce.Name(), member))
		}
		return ZEND_DYNAMIC_PROPERTY_OFFSET, nil, nil
	}
	offset = propertyInfo.GetOffset()
	if propertyInfo.GetType() == nil {
		propertyInfo = nil
	}
	return offset, ce, propertyInfo
dynamic:
	return ZEND_DYNAMIC_PROPERTY_OFFSET, ce, nil
wrong:
	/* Information was available, but we were denied access.  Error out. */
	if !silent {
		ZendBadPropertyAccess(propertyInfo, ce, member)
	}
	return ZEND_WRONG_PROPERTY_OFFSET, nil, nil
badPropName:
	if !silent {
		ZendBadPropertyName()
	}
	return ZEND_WRONG_PROPERTY_OFFSET, nil, nil
}
func ZendGetPropertyOffset(ce *types.ClassEntry, member string, silent bool, cacheSlot *any, infoPtr **types.PropertyInfo) uint32 {
	var propertyInfo *types.PropertyInfo
	var flags uint32
	var scope *types.ClassEntry
	var offset uint32

	if propertyInfo = ce.GetProperty(member); propertyInfo == nil {
		if member != "" && member[0] == '\x00' {
			if !silent {
				ZendBadPropertyName()
			}
			return ZEND_WRONG_PROPERTY_OFFSET
		}
		goto dynamic
	}
	flags = propertyInfo.GetFlags()
	if (flags & (types.AccChanged | types.AccPrivate | types.AccProtected)) != 0 {
		if EG__().GetFakeScope() != nil {
			scope = EG__().GetFakeScope()
		} else {
			scope = ZendGetExecutedScope()
		}
		if propertyInfo.GetCe() != scope {
			if (flags & types.AccChanged) != 0 {
				var p = ZendGetParentPrivateProperty(scope, ce, member)

				/* If there is a public/protected instance property on ce, don't try to use a
				 * private static property on scope. If both are static, prefer the static
				 * property on scope. This will throw a static property notice, rather than
				 * a visibility error. */
				if p != nil && (!p.IsStatic() || (flags&types.AccStatic) != 0) {
					propertyInfo = p
					flags = propertyInfo.GetFlags()
					goto found
				} else if (flags & types.AccPublic) != 0 {
					goto found
				}
			}
			if (flags & types.AccPrivate) != 0 {
				if propertyInfo.GetCe() != ce {
					goto dynamic
				} else {
					goto wrong
				}
			} else {
				b.Assert((flags & types.AccProtected) != 0)
				if IsProtectedCompatibleScope(propertyInfo.GetCe(), scope) == 0 {
					goto wrong
				}
			}
		}
	}
found:
	if (flags & types.AccStatic) != 0 {
		if !silent {
			faults.Error(faults.E_NOTICE, fmt.Sprintf("Accessing static property %s::$%s as non static", ce.Name(), member))
		}
		return ZEND_DYNAMIC_PROPERTY_OFFSET
	}
	offset = propertyInfo.GetOffset()
	if propertyInfo.GetType() == nil {
		propertyInfo = nil
	} else {
		*infoPtr = propertyInfo
	}
	if cacheSlot != nil {
		_setCacheSlot(cacheSlot, ce, any(offset), propertyInfo)
	}
	return offset
dynamic:
	if cacheSlot != nil {
		_setCacheSlot(cacheSlot, ce, any(ZEND_DYNAMIC_PROPERTY_OFFSET), nil)
	}
	return ZEND_DYNAMIC_PROPERTY_OFFSET
wrong:
	/* Information was available, but we were denied access.  Error out. */
	if !silent {
		ZendBadPropertyAccess(propertyInfo, ce, member)
	}
	return ZEND_WRONG_PROPERTY_OFFSET
}
func ZendWrongOffset(ce *types.ClassEntry, member *types.String) {
	var dummy *types.PropertyInfo

	/* Trigger the correct error */
	ZendGetPropertyOffset(ce, member.GetStr(), false, nil, &dummy)
}

func ZendGetPropertyInfo(ce *types.ClassEntry, member string) *types.PropertyInfo {
	propInfo, _ := ZendGetPropertyInfoEx(ce, member)
	return propInfo
}

func ZendGetPropertyInfoEx(ce *types.ClassEntry, member string) (_ *types.PropertyInfo, forbidden bool) {
	propInfo := ce.PropertyTable().Get(member)
	if propInfo == nil {
		if member != "" && member[0] == '\x00' {
			return nil, true
		}
		return nil, false
	}

	flags := propInfo.GetFlags()
	if (flags & (types.AccChanged | types.AccPrivate | types.AccProtected)) != 0 {
		var scope *types.ClassEntry
		if EG__().GetFakeScope() != nil {
			scope = EG__().GetFakeScope()
		} else {
			scope = ZendGetExecutedScope()
		}
		if propInfo.GetCe() != scope {
			if (flags & types.AccChanged) != 0 {
				var p = ZendGetParentPrivateProperty(scope, ce, member)
				if p != nil {
					return p, false
				} else if (flags & types.AccPublic) != 0 {
					return propInfo, false
				}
			}
			if (flags & types.AccPrivate) != 0 {
				if propInfo.GetCe() != ce {
					return nil, false
				} else {
					return nil, true
				}
			} else {
				b.Assert((flags & types.AccProtected) != 0)
				if IsProtectedCompatibleScope(propInfo.GetCe(), scope) == 0 {
					return nil, true
				}
			}
		}
	}

	return propInfo, false
}
func ZendCheckPropertyAccess(zobj *types.Object, prop_info_name *types.String, is_dynamic bool) int {
	var property_info *types.PropertyInfo
	var class_name *byte = nil
	var prop_name *byte
	var member *types.String
	var prop_name_len int
	if prop_info_name.GetStr()[0] == 0 {
		if is_dynamic != 0 {
			return types.SUCCESS
		}
		ZendUnmanglePropertyNameEx(prop_info_name, &class_name, &prop_name, &prop_name_len)
		member = types.NewString(b.CastStr(prop_name, prop_name_len))
		property_info = ZendGetPropertyInfo(zobj.GetCe(), member.GetStr())
		if property_info == nil {
			return types.FAILURE
		}
		if class_name[0] != '*' {
			if !property_info.IsPrivate() {
				/* we we're looking for a private prop but found a non private one of the same name */
				return types.FAILURE
			} else if prop_info_name.GetStr()[1:] != property_info.GetName()[1:] {
				/* we we're looking for a private prop but found a private one of the same name but another class */
				return types.FAILURE
			}
		} else {
			b.Assert(property_info.IsProtected())
		}
		return types.SUCCESS
	} else {
		propertyInfo, forbidden := ZendGetPropertyInfoEx(zobj.GetCe(), prop_info_name.GetStr())
		if forbidden {
			return types.FAILURE
		} else if propertyInfo == nil {
			b.Assert(is_dynamic != 0)
			return types.SUCCESS
		}
		if propertyInfo.IsPublic() {
			return types.SUCCESS
		} else {
			return types.FAILURE
		}
	}
}

func ZendStdReadProperty(object *types.Zval, member *types.Zval, type_ int, cache_slot *any, rv *types.Zval) *types.Zval {
	return ZendStdReadPropertyEx(object.Object(), member, type_, cache_slot, rv)
}
func ZendStdReadPropertyEx(zobj *types.Object, member *types.Zval, typ int, cache_slot *any, rv *types.Zval) *types.Zval {
	var name *types.String
	var tmp_name *types.String
	var retval *types.Zval
	var guard *types.PropertyGuard = nil
	name = operators.ZvalTryGetString(member)
	if name == nil {
		return UninitializedZval()
	}

	/* make zend_get_property_info silent if we have getter - we may want to use it */

	property_offset, prop_ce, prop_info := _zendGetPropertyOffset(zobj.GetCe(), name.GetStr(), typ == BP_VAR_IS || zobj.GetCe().GetGet() != nil)
	if IS_VALID_PROPERTY_OFFSET(property_offset) {
		retval = OBJ_PROP(zobj, property_offset)
		if retval.IsNotUndef() {
			goto exit
		}
		if retval.GetU2Extra() == types.IS_PROP_UNINIT {
			/* Skip __get() for uninitialized typed properties */
			goto uninit_error
		}
	} else if IS_DYNAMIC_PROPERTY_OFFSET(property_offset) {
		if zobj.GetProperties() != nil {
			if !(IS_UNKNOWN_DYNAMIC_PROPERTY_OFFSET(property_offset)) {
				var idx uint = ZEND_DECODE_DYN_PROP_OFFSET(property_offset)
				if idx < zobj.GetProperties().GetNNumUsed()*b.SizeOf("Bucket") {
					var p *types.Bucket = (*types.Bucket)((*byte)(zobj.GetProperties().Bucket(idx)))
					if p.GetVal().IsNotUndef() && p.IsStrKey() && p.StrKey() == name.GetStr() {
						retval = p.GetVal()
						goto exit
					}
				}
				CACHE_PTR_EX(cache_slot+1, any(ZEND_DYNAMIC_PROPERTY_OFFSET))
			}
			retval = zobj.GetProperties().KeyFind(name.GetStr())
			if retval != nil {
				if cache_slot != nil {
					PropFindAndCache(zobj, name.GetStr())
				}
				goto exit
			}
		}
	} else if EG__().HasException() {
		retval = UninitializedZval()
		goto exit
	}

	/* magic isset */

	if typ == BP_VAR_IS && zobj.GetCe().GetIsset() != nil {
		var tmp_result types.Zval
		guard = zobj.Guard(name.GetStr())
		if !guard.InIsset() {
			if tmp_name == nil {
				tmp_name = name.Copy()
			}
			tmp_result.SetUndef()
			guard.MarkInIsset(true)
			ZendStdCallIssetter(zobj, name.GetStr(), &tmp_result)
			guard.MarkInIsset(false)
			if !operators.ZvalIsTrue(&tmp_result) {
				retval = UninitializedZval()
				goto exit
			}
			if zobj.GetCe().GetGet() != nil && !guard.InGet() {
				goto call_getter
			}
		} else if zobj.GetCe().GetGet() != nil && !guard.InGet() {
			goto call_getter_addref
		}
	} else if zobj.GetCe().GetGet() != nil {

		/* magic get */
		guard = zobj.Guard(name.GetStr())
		if !guard.InGet() {
		call_getter_addref:
		call_getter:
			guard.MarkInGet(true)
			ZendStdCallGetter(zobj, name.GetStr(), rv)
			guard.MarkInGet(false)
			if rv.IsNotUndef() {
				retval = rv
				if !(rv.IsRef()) && (typ == BP_VAR_W || typ == BP_VAR_RW || typ == BP_VAR_UNSET) {
					if !rv.IsObject() {
						faults.Error(faults.E_NOTICE, fmt.Sprintf("Indirect modification of overloaded property %s::$%s has no effect", zobj.GetCe().Name(), name.GetVal()))
					}
				}
			} else {
				retval = UninitializedZval()
			}
			if prop_info != nil {
				ZendVerifyPropAssignableByRef(prop_info, retval, zobj.GetCe().GetGet().IsStrictTypes())
			}
			// OBJ_RELEASE(zobj)
			goto exit
		} else if IS_WRONG_PROPERTY_OFFSET(property_offset) {

			/* Trigger the correct error */

			ZendGetPropertyOffset(zobj.GetCe(), name.GetStr(), false, nil, &prop_info)
			b.Assert(EG__().HasException())
			retval = UninitializedZval()
			goto exit
		}
	}
uninit_error:
	if typ != BP_VAR_IS {
		if prop_info != nil {
			faults.ThrowError(nil, "Typed property %s::$%s must not be accessed before initialization", prop_info.GetCe().Name(), name.GetVal())
		} else {
			faults.Error(faults.E_NOTICE, fmt.Sprintf("Undefined property: %s::$%s", zobj.GetCe().Name(), name.GetVal()))
		}
	}
	retval = UninitializedZval()
exit:
	return retval
}
func PropertyUsesStrictTypes() bool {
	var executeData *ZendExecuteData = CurrEX()
	return executeData != nil && executeData.GetFunc() != nil && CurrEX().IsCallUseStrictTypes()
}
func ZendStdWriteProperty(object *types.Zval, member *types.Zval, value *types.Zval, cache_slot *any) *types.Zval {
	return ZendStdWritePropertyEx(object.Object(), member, value, cache_slot)
}
func ZendStdWritePropertyEx(zobj *types.Object, member *types.Zval, value *types.Zval, cache_slot *any) *types.Zval {
	var name *types.String
	var variable_ptr *types.Zval
	var tmp types.Zval
	var property_offset uintPtr
	var prop_info *types.PropertyInfo = nil
	b.Assert(!(value.IsRef()))
	name = operators.ZvalTryGetString(member)
	if name == nil {
		return value
	}
	property_offset = ZendGetPropertyOffset(zobj.GetCe(), name.GetStr(), zobj.GetCe().GetSet() != nil, cache_slot, &prop_info)
	if IS_VALID_PROPERTY_OFFSET(property_offset) {
		variable_ptr = OBJ_PROP(zobj, property_offset)
		if variable_ptr.IsNotUndef() {
			if prop_info != nil {
				types.ZVAL_COPY_VALUE(&tmp, value)
				if ZendVerifyPropertyType(prop_info, &tmp, PropertyUsesStrictTypes()) == 0 {
					variable_ptr = EG__().GetErrorZval()
					goto exit
				}
				value = &tmp
			}
		found:
			variable_ptr = ZendAssignToVariable(variable_ptr, value, PropertyUsesStrictTypes())
			goto exit
		}
		if variable_ptr.GetU2Extra() == types.IS_PROP_UNINIT {

			/* Writes to uninitializde typed properties bypass __set(). */

			variable_ptr.SetU2Extra(0)
			goto write_std_property
		}
	} else if IS_DYNAMIC_PROPERTY_OFFSET(property_offset) {
		if zobj.GetProperties() != nil {
			zobj.DupProperties()
			if lang.Assign(&variable_ptr, zobj.GetProperties().KeyFind(name.GetStr())) != nil {
				goto found
			}
		}
	} else if EG__().HasException() {
		variable_ptr = EG__().GetErrorZval()
		goto exit
	}

	/* magic set */

	if zobj.GetCe().GetSet() != nil {
		var guard = zobj.Guard(name.GetStr())
		if !guard.InGet() {
			guard.MarkInSet(true)
			ZendStdCallSetter(zobj, name.GetStr(), value)
			guard.MarkInSet(false)
			variable_ptr = value
		} else if !(IS_WRONG_PROPERTY_OFFSET(property_offset)) {
			goto write_std_property
		} else {

			/* Trigger the correct error */

			ZendWrongOffset(zobj.GetCe(), name)
			b.Assert(EG__().HasException())
			variable_ptr = EG__().GetErrorZval()
			goto exit
		}
	} else {
		b.Assert(!(IS_WRONG_PROPERTY_OFFSET(property_offset)))
	write_std_property:
		// value.TryAddRefcount()
		if IS_VALID_PROPERTY_OFFSET(property_offset) {
			variable_ptr = OBJ_PROP(zobj, property_offset)
			if prop_info != nil {
				types.ZVAL_COPY_VALUE(&tmp, value)
				if ZendVerifyPropertyType(prop_info, &tmp, PropertyUsesStrictTypes()) == 0 {
					// ZvalPtrDtor(value)
					goto exit
				}
				value = &tmp
				goto found
			}
			variable_ptr.CopyValueFrom(value)
		} else {
			if zobj.GetProperties() == nil {
				RebuildObjectProperties(zobj)
			}
			variable_ptr = zobj.GetProperties().KeyAddNew(name.GetStr(), value)
		}
	}
exit:
	//ZendTmpStringRelease(tmp_name)
	return variable_ptr
}
func ZendBadArrayAccess(ce *types.ClassEntry) {
	faults.ThrowError(nil, "Cannot use object of type %s as array", ce.Name())
}
func ZendStdReadDimensionEx(object *types.Object, offset *types.Zval, type_ int, rv *types.Zval) *types.Zval {
	var ce *types.ClassEntry = object.GetCe()
	var tmp_offset types.Zval
	var tmp_object types.Zval
	if operators.InstanceofFunctionEx(ce, ZendCeArrayaccess, 1) != 0 {
		if offset == nil {
			/* [] construct */
			tmp_offset.SetNull()
		} else {
			types.ZVAL_COPY_DEREF(&tmp_offset, offset)
		}
		tmp_object.SetObject(object)
		if type_ == BP_VAR_IS {
			ZendCallMethodWith1Params(&tmp_object, ce, nil, "offsetexists", rv, &tmp_offset)
			if rv.IsUndef() {
				return nil
			}
			if !operators.ZvalIsTrue(rv) {
				return UninitializedZval()
			}
		}
		ZendCallMethodWith1Params(&tmp_object, ce, nil, "offsetget", rv, &tmp_offset)
		if rv.IsUndef() {
			if EG__().NoException() {
				faults.ThrowError(nil, "Undefined offset for object of type %s used as array", ce.Name())
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
	ZendStdWriteDimensionEx(object.Object(), offset, value)
}
func ZendStdWriteDimensionEx(object *types.Object, offset *types.Zval, value *types.Zval) {
	var ce *types.ClassEntry = object.GetCe()
	var tmp_offset types.Zval
	var tmp_object types.Zval
	if operators.InstanceofFunctionEx(ce, ZendCeArrayaccess, 1) != 0 {
		if offset == nil {
			tmp_offset.SetNull()
		} else {
			types.ZVAL_COPY_DEREF(&tmp_offset, offset)
		}
		tmp_object.SetObject(object)
		ZendCallMethodWith2Params(&tmp_object, ce, nil, "offsetset", nil, &tmp_offset, value)
	} else {
		ZendBadArrayAccess(ce)
	}
}
func ZendStdHasDimension(object *types.Zval, offset *types.Zval, check_empty int) int {
	return ZendStdHasDimensionEx(object.Object(), offset, check_empty)
}
func ZendStdHasDimensionEx(object *types.Object, offset *types.Zval, check_empty int) int {
	var ce *types.ClassEntry = object.GetCe()
	var retval types.Zval
	var tmp_offset types.Zval
	var tmp_object types.Zval
	var result int
	if operators.InstanceofFunctionEx(ce, ZendCeArrayaccess, 1) != 0 {
		types.ZVAL_COPY_DEREF(&tmp_offset, offset)
		tmp_object.SetObject(object)
		ZendCallMethodWith1Params(&tmp_object, ce, nil, "offsetexists", &retval, &tmp_offset)
		result = operators.IZendIsTrue(&retval)
		if check_empty != 0 && result != 0 && EG__().NoException() {
			ZendCallMethodWith1Params(&tmp_object, ce, nil, "offsetget", &retval, &tmp_offset)
			result = operators.IZendIsTrue(&retval)
		}
	} else {
		ZendBadArrayAccess(ce)
		return 0
	}
	return result
}
func ZendStdGetPropertyPtrPtr(object *types.Zval, member *types.Zval, type_ int, cache_slot *any) *types.Zval {
	return ZendStdGetPropertyPtrPtrEx(object.Object(), member, type_, cache_slot)
}
func ZendStdGetPropertyPtrPtrEx(zobj *types.Object, member *types.Zval, type_ int, cache_slot *any) *types.Zval {
	var name *types.String
	var retval *types.Zval = nil
	var property_offset uintPtr
	var prop_info *types.PropertyInfo = nil
	name = operators.ZvalTryGetString(member)
	if name == nil {
		return EG__().GetErrorZval()
	}
	property_offset = ZendGetPropertyOffset(zobj.GetCe(), name.GetStr(), zobj.GetCe().GetGet() != nil, cache_slot, &prop_info)
	if IS_VALID_PROPERTY_OFFSET(property_offset) {
		retval = OBJ_PROP(zobj, property_offset)
		if retval.IsUndef() {
			if zobj.GetCe().GetGet() == nil || zobj.Guard(name.GetStr()).InGet() || prop_info != nil && retval.GetU2Extra() == types.IS_PROP_UNINIT {
				if type_ == BP_VAR_RW || type_ == BP_VAR_R {
					if prop_info != nil {
						faults.ThrowError(nil, "Typed property %s::$%s must not be accessed before initialization", prop_info.GetCe().Name(), name.GetVal())
						retval = EG__().GetErrorZval()
					} else {
						retval.SetNull()
						faults.Error(faults.E_NOTICE, fmt.Sprintf("Undefined property: %s::$%s", zobj.GetCe().Name(), name.GetVal()))
					}
				}
			} else {
				/* we do have getter - fail and let it try again with usual get/set */
				retval = nil
			}
		}
	} else if IS_DYNAMIC_PROPERTY_OFFSET(property_offset) {
		if zobj.GetProperties() != nil {
			zobj.DupProperties()
			if lang.Assign(&retval, zobj.GetProperties().KeyFind(name.GetStr())) != nil {
				//ZendTmpStringRelease(tmp_name)
				return retval
			}
		}
		if zobj.GetCe().GetGet() == nil || zobj.Guard(name.GetStr()).InGet() {
			if zobj.GetProperties() == nil {
				RebuildObjectProperties(zobj)
			}
			retval = zobj.GetProperties().KeyUpdate(name.GetStr(), UninitializedZval())

			/* Notice is thrown after creation of the property, to avoid EG(std_property_info)
			 * being overwritten in an error handler. */

			if type_ == BP_VAR_RW || type_ == BP_VAR_R {
				faults.Error(faults.E_NOTICE, fmt.Sprintf("Undefined property: %s::$%s", zobj.GetCe().Name(), name.GetVal()))
			}

			/* Notice is thrown after creation of the property, to avoid EG(std_property_info)
			 * being overwritten in an error handler. */

		}
	} else if zobj.GetCe().GetGet() == nil {
		retval = EG__().GetErrorZval()
	}
	//ZendTmpStringRelease(tmp_name)
	return retval
}
func ZendStdUnsetProperty(object *types.Zval, member *types.Zval, cache_slot *any) {
	ZendStdUnsetPropertyEx(object.Object(), member, cache_slot)
}
func ZendStdUnsetPropertyEx(zobj *types.Object, member *types.Zval, cache_slot *any) {
	var name *types.String
	var property_offset uintPtr
	var prop_info *types.PropertyInfo = nil
	name = operators.ZvalTryGetString(member)
	if name == nil {
		return
	}
	property_offset = ZendGetPropertyOffset(zobj.GetCe(), name.GetStr(), zobj.GetCe().GetUnset() != nil, cache_slot, &prop_info)
	if IS_VALID_PROPERTY_OFFSET(property_offset) {
		var slot *types.Zval = OBJ_PROP(zobj, property_offset)
		if slot.IsNotUndef() {
			if slot.IsRef() && ZEND_REF_HAS_TYPE_SOURCES(slot.Ref()) {
				if prop_info != nil {
					ZEND_REF_DEL_TYPE_SOURCE(slot.Ref(), prop_info)
				}
			}
			var tmp types.Zval
			types.ZVAL_COPY_VALUE(&tmp, slot)
			slot.SetUndef()
			// ZvalPtrDtor(&tmp)
			if zobj.GetProperties() != nil {
				zobj.GetProperties().MarkHasEmptyIndex()
			}
			goto exit
		}
		if slot.GetU2Extra() == types.IS_PROP_UNINIT {

			/* Reset the IS_PROP_UNINIT flag, if it exists and bypass __unset(). */

			slot.SetU2Extra(0)
			goto exit
		}
	} else if IS_DYNAMIC_PROPERTY_OFFSET(property_offset) && zobj.GetProperties() != nil {
		zobj.DupProperties()
		if !zobj.GetProperties().KeyDelete(name.GetStr()) {
			goto exit
		}
	} else if EG__().HasException() {
		goto exit
	}

	/* magic unset */

	if zobj.GetCe().GetUnset() != nil {
		var guard = zobj.Guard(name.GetStr())
		if !guard.InUnset() {
			/* have unseter - try with it! */
			guard.MarkInUnset(true)
			ZendStdCallUnsetter(zobj, name.GetStr())
			guard.MarkInUnset(false)
		} else if IS_WRONG_PROPERTY_OFFSET(property_offset) {
			/* Trigger the correct error */
			ZendWrongOffset(zobj.GetCe(), name)
			b.Assert(EG__().HasException())
			goto exit
		}
	}
exit:
	//ZendTmpStringRelease(tmp_name)
}
func ZendStdUnsetDimensionEx(object *types.Object, offset *types.Zval) {
	var ce *types.ClassEntry = object.GetCe()
	var tmp_offset types.Zval
	var tmp_object types.Zval
	if operators.InstanceofFunctionEx(ce, ZendCeArrayaccess, 1) != 0 {
		types.ZVAL_COPY_DEREF(&tmp_offset, offset)
		tmp_object.SetObject(object)
		ZendCallMethodWith1Params(&tmp_object, ce, nil, "offsetunset", nil, &tmp_offset)
	} else {
		ZendBadArrayAccess(ce)
	}
}
func ZendGetParentPrivateMethod(scope *types.ClassEntry, ce *types.ClassEntry, functionName string) types.IFunction {
	var fbc types.IFunction
	if scope != ce && scope != nil && IsDerivedClass(ce, scope) != 0 {
		fbc = scope.FunctionTable().Get(functionName)
		if fbc != nil {
			if fbc.IsPrivate() && fbc.GetScope() == scope {
				return fbc
			}
		}
	}
	return nil
}
func ZendCheckProtected(ce *types.ClassEntry, scope *types.ClassEntry) bool {
	var fbc_scope *types.ClassEntry = ce

	/* Is the context that's calling the function, the same as one of
	 * the function's parents?
	 */

	for fbc_scope != nil {
		if fbc_scope == scope {
			return true
		}
		fbc_scope = fbc_scope.GetParent()
	}

	/* Is the function's scope the same as our current object context,
	 * or any of the parents of our context?
	 */

	for scope != nil {
		if scope == ce {
			return true
		}
		scope = scope.GetParent()
	}
	return false
}

func ZendGetCallTrampolineFunc(ce *types.ClassEntry, methodName string, isStatic bool) types.IFunction {
	var func_ *types.ZendOpArray
	var fbc types.IFunction = lang.Cond(isStatic, ce.GetCallstatic(), ce.GetCall())

	/* We use non-NULL value to avoid useless run_time_cache allocation.
	 * The low bit must be zero, to not be interpreted as a MAP_PTR offset.
	 */

	var dummy any = any(intPtr(2))
	b.Assert(fbc != nil)
	if EG__().GetTrampoline().FunctionName() == "" {
		func_ = EG__().GetTrampoline().GetOpArray()
	} else {
		func_ = types.NewOpArray()
	}
	func_.SetFnFlags(types.AccCallViaTrampoline | types.AccPublic)
	if isStatic {
		func_.SetIsStatic(true)
	}
	func_.SetOpcodes(EG__().GetCallTrampolineOp())
	ZEND_MAP_PTR_INIT(func_.GetRunTimeCache(), (**any)(&dummy))
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
		func_.SetFilename("")
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
	methodName = strings.TrimRight(methodName, "\000")
	func_.SetFunctionName(methodName)
	func_.SetPrototype(nil)
	func_.SetNumArgs(0)
	func_.SetRequiredNumArgs(0)
	func_.SetArgInfo(0)
	return (types.IFunction)(func_)
}
func ZendGetUserCallFunction(ce *types.ClassEntry, method_name *types.String) types.IFunction {
	return ZendGetCallTrampolineFunc(ce, method_name.GetStr(), false)
}
func ZendBadMethodCall(fbc types.IFunction, method_name *types.String, scope *types.ClassEntry) {
	faults.ThrowError(nil, "Call to %s method %s::%s() from context '%s'", ZendVisibilityString(fbc.GetFnFlags()), ZEND_FN_SCOPE_NAME(fbc), method_name.GetVal(), lang.CondF1(scope != nil, func() []byte { return scope.Name() }, ""))
}

func ZendStdGetMethod(obj_ptr **types.Object, method_name *types.String, key *types.Zval) types.IFunction {
	return ZendStdGetMethod_Ex(*obj_ptr, method_name.GetStr(), key)
}
func ZendStdGetMethod_Ex(zobj *types.Object, methodName string, key *types.Zval) types.IFunction {
	var fbc types.IFunction
	var lc_method_name string
	var scope *types.ClassEntry
	if key != nil {
		lc_method_name = key.String()
	} else {
		lc_method_name = ascii.StrToLower(methodName)
	}
	fbc = zobj.GetCe().FunctionTable().Get(lc_method_name)
	if fbc == nil {
		if zobj.GetCe().GetCall() != nil {
			return ZendGetUserCallFunction(zobj.GetCe(), types.NewString(methodName))
		} else {
			return nil
		}
	}

	/* Check access level */
	if fbc.GetOpArray().HasFnFlags(types.AccChanged | types.AccPrivate | types.AccProtected) {
		scope = ZendGetExecutedScope()
		if fbc.GetScope() != scope {
			if fbc.GetOpArray().IsChanged() {
				var updated_fbc types.IFunction = ZendGetParentPrivateMethod(scope, zobj.GetCe(), lc_method_name)
				if updated_fbc != nil {
					return updated_fbc
				} else if fbc.GetOpArray().IsPublic() {
					return fbc
				}
			}
			if fbc.GetOpArray().IsPrivate() || !ZendCheckProtected(ZendGetFunctionRootClass(fbc), scope) {
				if zobj.GetCe().GetCall() != nil {
					return ZendGetUserCallFunction(zobj.GetCe(), types.NewString(methodName))
				} else {
					ZendBadMethodCall(fbc, types.NewString(methodName), scope)
					return nil
				}
			}
		}
	}
	return fbc
}
func ZendGetUserCallstaticFunction(ce *types.ClassEntry, method_name *types.String) types.IFunction {
	return ZendGetCallTrampolineFunc(ce, method_name.GetStr(), true)
}
func ZendStdGetStaticMethod(ce *types.ClassEntry, function_name *types.String, key *types.Zval) types.IFunction {
	var lc_function_name *types.String
	var object *types.Object
	var scope *types.ClassEntry
	if key != nil {
		lc_function_name = key.StringEx()
	} else {
		lc_function_name = operators.ZendStringTolower(function_name)
	}

	fbc := ce.FunctionTable().Get(lc_function_name.GetStr())
	if fbc != nil {
		// pass
	} else if ce.GetConstructor() != nil && lc_function_name.GetLen() == len(ce.Name()) && operators.ZendBinaryStrncasecmp(lc_function_name.GetStr(), b.CastStr(ce.Name(), lc_function_name.GetLen()), lc_function_name.GetLen()) == 0 && (ce.GetConstructor().FunctionName()[0] != '_' || ce.GetConstructor().FunctionName()[1] != '_') {
		fbc = ce.GetConstructor()
	} else {
		if ce.GetCall() != nil && lang.Assign(&object, ZendGetThisObject(CurrEX())) != nil && operators.InstanceofFunction(object.GetCe(), ce) != 0 {
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
			if fbc.GetOpArray().IsPrivate() || !ZendCheckProtected(ZendGetFunctionRootClass(fbc), scope) {
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
		// types.ZendStringReleaseEx(lc_function_name, 0)
	}
	return fbc
}
func ZendClassInitStatics(class_type *types.ClassEntry) {
	var i int
	var p *types.Zval
	if class_type.GetDefaultStaticMembersCount() != 0 && CE_STATIC_MEMBERS(class_type) == nil {
		if class_type.GetParent() != nil {
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
func ZendStdGetStaticPropertyWithInfo(ce *types.ClassEntry, property_name *types.String, type_ int, property_info_ptr **types.PropertyInfo) *types.Zval {
	var ret *types.Zval
	var scope *types.ClassEntry
	var property_info *types.PropertyInfo = ce.PropertyTable().Get(property_name.GetStr())
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
					ZendBadPropertyAccess(property_info, ce, property_name.GetStr())
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
		if ce.IsInternalClass() || ce.HasCeFlags(types.AccImmutable|types.AccPreloaded) {
			ZendClassInitStatics(ce)
		} else {
		undeclared_property:
			if type_ != BP_VAR_IS {
				faults.ThrowError(nil, "Access to undeclared static property: %s::$%s", ce.Name(), property_name.GetVal())
			}
			return nil
		}
	}
	ret = CE_STATIC_MEMBERS(ce) + property_info.GetOffset()
	ret = types.ZVAL_DEINDIRECT(ret)
	if (type_ == BP_VAR_R || type_ == BP_VAR_RW) && ret.IsUndef() && property_info.GetType() != 0 {
		faults.ThrowError(nil, "Typed static property %s::$%s must not be accessed before initialization", property_info.GetCe().Name(), ZendGetUnmangledPropertyNameEx(property_name.GetStr()))
		return nil
	}
	return ret
}
func ZendStdGetStaticProperty(ce *types.ClassEntry, property_name *types.String, type_ int) *types.Zval {
	var prop_info *types.PropertyInfo
	return ZendStdGetStaticPropertyWithInfo(ce, property_name, type_, &prop_info)
}
func ZendStdUnsetStaticProperty(ce *types.ClassEntry, property_name *types.String) bool {
	faults.ThrowError(nil, "Attempt to unset static property %s::$%s", ce.Name(), property_name.GetVal())
	return 0
}
func ZendBadConstructorCall(constructor types.IFunction, scope *types.ClassEntry) {
	if scope != nil {
		faults.ThrowError(nil, "Call to %s %s::%s() from context '%s'", ZendVisibilityString(constructor.GetFnFlags()), constructor.GetScope().Name(), constructor.FunctionName(), scope.Name())
	} else {
		faults.ThrowError(nil, "Call to %s %s::%s() from invalid context", ZendVisibilityString(constructor.GetFnFlags()), constructor.GetScope().Name(), constructor.FunctionName())
	}
}
func ZendStdGetConstructor(zobj *types.Object) types.IFunction {
	var constructor types.IFunction = zobj.GetCe().GetConstructor()
	var scope *types.ClassEntry
	if constructor != nil {
		if !constructor.GetOpArray().IsPublic() {
			if EG__().GetFakeScope() != nil {
				scope = EG__().GetFakeScope()
			} else {
				scope = ZendGetExecutedScope()
			}
			if constructor.GetScope() != scope {
				if constructor.GetOpArray().IsPrivate() || !ZendCheckProtected(ZendGetFunctionRootClass(constructor), scope) {
					ZendBadConstructorCall(constructor, scope)
					constructor = nil
				}
			}
		}
	}
	return constructor
}
func ZendStdCompareObjects(o1 *types.Zval, o2 *types.Zval) int {
	return ZendStdCompareObjectsEx(o1.Object(), o2.Object())
}
func ZendStdCompareObjectsEx(zobj1 *types.Object, zobj2 *types.Object) int {
	if zobj1 == zobj2 {
		return 0
	}
	if zobj1.GetCe() != zobj2.GetCe() {
		return 1
	}
	if zobj1.GetProperties() == nil && zobj2.GetProperties() == nil {
		if zobj1.GetCe().GetDefaultPropertiesCount() == 0 {
			return 0
		}

		/* It's enough to protect only one of the objects.
		 * The second one may be referenced from the first and this may cause
		 * false recursion detection.
		 */

		if zobj1.IsRecursive() {
			faults.ErrorNoreturn(faults.E_ERROR, "Nesting level too deep - recursive dependency?")
		}
		zobj1.ProtectRecursive()

		var ret int
		zobj1.GetCe().PropertyTable().ForeachEx(func(key string, info *types.PropertyInfo) bool {
			var p1 *types.Zval = OBJ_PROP(zobj1, info.GetOffset())
			var p2 *types.Zval = OBJ_PROP(zobj2, info.GetOffset())
			if info.IsStatic() {
				return true
			}
			if p1.IsNotUndef() {
				if p2.IsNotUndef() {
					var result types.Zval
					if operators.CompareFunction(&result, p1, p2) == types.FAILURE {
						zobj1.UnprotectRecursive()
						return false
					}
					if result.Long() != 0 {
						zobj1.UnprotectRecursive()
						ret = result.Long()
						return false
					}
				} else {
					zobj1.UnprotectRecursive()
					ret = 1
					return false
				}
			} else {
				if p2.IsNotUndef() {
					zobj1.UnprotectRecursive()
					ret = 1
					return false
				}
			}
			return true
		})
		if ret != 0 {
			return ret
		}

		zobj1.UnprotectRecursive()
		return 0
	} else {
		if zobj1.GetProperties() == nil {
			RebuildObjectProperties(zobj1)
		}
		if zobj2.GetProperties() == nil {
			RebuildObjectProperties(zobj2)
		}
		return operators.ZendCompareSymbolTables(zobj1.GetProperties(), zobj2.GetProperties())
	}
}
func ZendStdHasProperty(object *types.Zval, member *types.Zval, has_set_exists int, cache_slot *any) int {
	return ZendStdHasPropertyEx(object.Object(), member, has_set_exists, cache_slot)
}
func ZendStdHasPropertyEx(zobj *types.Object, member *types.Zval, hasSetExists int, cache_slot *any) int {
	name, ok := operators.ZvalTryGetStr(member)
	if !ok {
		return 0
	}

	result := ZendStdHasPropertyExEx(zobj, name, hasSetExists)
	return types.IntBool(result)
}
func ZendStdHasPropertyExEx(zobj *types.Object, name string, hasSetExists int) bool {
	var result bool
	var value *types.Zval = nil

	propOffset, _, _ := _zendGetPropertyOffset(zobj.GetCe(), name, true)
	if propOffset > 0 {
		value = OBJ_PROP(zobj, propOffset)
		if value.IsNotUndef() {
			goto found
		}
		if value.GetU2Extra() == types.IS_PROP_UNINIT {
			/* Skip __isset() for uninitialized typed properties */
			return false
		}
	} else if propOffset < 0 {
		if zobj.GetProperties() != nil {
			if propOffset != ZEND_DYNAMIC_PROPERTY_OFFSET {
				var idx uint = ZEND_DECODE_DYN_PROP_OFFSET(propOffset)
				if idx < zobj.GetProperties().GetNNumUsed()*b.SizeOf("Bucket") {
					var p *types.Bucket = (*types.Bucket)((*byte)(zobj.GetProperties().Bucket(idx)))
					if p.GetVal().IsNotUndef() && p.IsStrKey() && p.StrKey() == name {
						value = p.GetVal()
						goto found
					}
				}
				CACHE_PTR_EX(cache_slot+1, any(ZEND_DYNAMIC_PROPERTY_OFFSET))
			}
			value = zobj.GetProperties().KeyFind(name)
			if value != nil {
				if cache_slot != nil {
					PropFindAndCache(zobj, name)
				}
				goto found
			}
		}
	} else if EG__().HasException() {
		return false
	}

	result = false
	if hasSetExists != ZEND_PROPERTY_EXISTS && zobj.GetCe().GetIsset() != nil {
		var guard = zobj.Guard(name)
		if !guard.InIsset() {
			var rv types.Zval

			/* have issetter - try with it! */
			guard.MarkInIsset(true)
			ZendStdCallIssetter(zobj, name, &rv)
			result = operators.ZvalIsTrue(&rv)
			if hasSetExists == ZEND_PROPERTY_NOT_EMPTY && result {
				if EG__().NoException() && zobj.GetCe().GetGet() != nil && !guard.InGet() {
					guard.MarkInGet(true)
					ZendStdCallGetter(zobj, name, &rv)
					guard.MarkInGet(false)
					result = operators.ZvalIsTrue(&rv)
				} else {
					result = false
				}
			}
			guard.MarkInIsset(false)
		}
	}
	return result

found:
	if hasSetExists == ZEND_PROPERTY_NOT_EMPTY {
		result = operators.ZvalIsTrue(value)
	} else if hasSetExists < ZEND_PROPERTY_NOT_EMPTY {
		b.Assert(hasSetExists == ZEND_PROPERTY_ISSET)
		value = types.ZVAL_DEREF(value)
		result = !value.IsNull()
	} else {
		b.Assert(hasSetExists == ZEND_PROPERTY_EXISTS)
		result = true
	}
	return result
}

func StdCastObjectToString(obj *types.Object) (string, bool) {
	ce := obj.GetCe()
	if ce.GetTostring() != nil {
		var fakeScope = EG__().GetFakeScope()
		EG__().SetFakeScope(nil)

		var retval types.Zval
		var fun = ce.GetTostring()
		ZendCallMethodWith0Params(types.NewZvalObject(obj), ce, &fun, "__tostring", &retval)
		EG__().SetFakeScope(fakeScope)
		if retval.IsString() {
			return retval.String(), true
		}
		if EG__().NoException() {
			faults.ThrowError(nil, "Method %s::__toString() must return a string value", ce.Name())
		}
	}
	return "", false
}

func ZendStdCastObject(obj *types.Object, retval *types.Zval, typ types.ZvalType) int {
	switch typ {
	case types.IsString:
		if str, ok := StdCastObjectToString(obj); ok {
			retval.SetString(str)
			return types.SUCCESS
		}
		return types.FAILURE
	case types.IsBool:
		retval.SetTrue()
		return types.SUCCESS
	case types.IsLong:
		className := obj.GetCe().Name()
		faults.Error(faults.E_NOTICE, fmt.Sprintf("Object of class %s could not be converted to int", className))
		retval.SetLong(1)
		return types.SUCCESS
	case types.IsDouble:
		className := obj.GetCe().Name()
		faults.Error(faults.E_NOTICE, fmt.Sprintf("Object of class %s could not be converted to float", className))
		retval.SetDouble(1)
		return types.SUCCESS
	case types.IsNumber:
		className := obj.GetCe().Name()
		faults.Error(faults.E_NOTICE, fmt.Sprintf("Object of class %s could not be converted to number", className))
		retval.SetLong(1)
		return types.SUCCESS
	default:
		retval.SetNull()
		return types.FAILURE
	}
}

func ZendStdCastObjectTostring(readobj *types.Zval, writeobj *types.Zval, typ types.ZvalType) int {
	return ZendStdCastObject(readobj.Object(), writeobj, typ)
}
func ZendStdGetClosure(obj *types.Zval, ce_ptr **types.ClassEntry, fptr_ptr *types.IFunction, obj_ptr **types.Object) int {
	var ce *types.ClassEntry = types.Z_OBJCE_P(obj)
	fptr := ce.FunctionTable().Get(types.STR_MAGIC_INVOKE)
	if fptr == nil {
		return types.FAILURE
	}
	*fptr_ptr = fptr
	*ce_ptr = ce
	if fptr.IsStatic() {
		if obj_ptr != nil {
			*obj_ptr = nil
		}
	} else {
		if obj_ptr != nil {
			*obj_ptr = obj.Object()
		}
	}
	return types.SUCCESS
}
func ZendStdGetPropertiesFor(obj *types.Zval, purpose ZendPropPurpose) *types.Array {
	var ht *types.Array
	switch purpose {
	case ZEND_PROP_PURPOSE_DEBUG:
		fallthrough
	case ZEND_PROP_PURPOSE_ARRAY_CAST:
		fallthrough
	case ZEND_PROP_PURPOSE_SERIALIZE:
		fallthrough
	case ZEND_PROP_PURPOSE_VAR_EXPORT:
		fallthrough
	case ZEND_PROP_PURPOSE_JSON:
		fallthrough
	case ZEND_PROP_PURPOSE_ARRAY_KEY_EXISTS:
		ht = obj.Object().GetPropertiesArray()
		return ht
	default:
		b.Assert(false)
		return nil
	}
}
func ZendGetPropertiesFor(obj *types.Zval, purpose ZendPropPurpose) *types.Array {
	if obj.Object().CanGetPropertiesFor() {
		return obj.Object().GetPropertiesFor(purpose)
	}
	return ZendStdGetPropertiesFor(obj, purpose)
}
